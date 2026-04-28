package netbox

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// device_type_templates.go implements full-CRUD nested template lifecycle on
// netbox_device_type for all 10 NetBox device-type-component template families:
// power_port, interface, console_port, console_server_port, rear_port,
// device_bay, module_bay, power_outlet (FK -> power_port by name),
// front_port (FK -> rear_port by name), and inventory_item (parent tree +
// polymorphic component_type/component_id FK).
//
// Architecture
// ------------
// Each template family is exposed as a TypeSet block on netbox_device_type,
// hash-keyed by name. The schema for each family is built by a per-type
// <type>TemplateSchema() function. Sync between Terraform and NetBox is
// driven by a single helper, syncDeviceTypeTemplates, which orchestrates all
// 10 types in dependency order:
//
//   pass 1 — independent types (any order):
//       power_port, interface, console_port, console_server_port,
//       rear_port, device_bay, module_bay
//   pass 2 — depend on a sibling template ID resolved by name:
//       power_outlet (-> power_port), front_port (-> rear_port)
//   pass 3 — inventory_item, sorted by tree depth so parents are created
//       before children
//
// Per-type behavior is plugged in via a templateOps closure struct so the
// reconciliation logic stays in one place. Each per-type file section holds
// only the schema definition plus the small Expand / Flatten / List / Create
// / PartialUpdate / Delete glue functions.
//
// Coexistence with standalone netbox_interface_template / netbox_device_bay_template
// resources: any individual NetBox template object should be managed by either
// (a) a nested block on its parent device_type or (b) a standalone resource —
// never both. See AGENTS.md for the rationale.

// Top-level keys on the netbox_device_type schema (one per template family).
const (
	powerPortTemplatesKey         = "power_port_templates"
	interfaceTemplatesKey         = "interface_templates"
	consolePortTemplatesKey       = "console_port_templates"
	consoleServerPortTemplatesKey = "console_server_port_templates"
	rearPortTemplatesKey          = "rear_port_templates"
	deviceBayTemplatesKey         = "device_bay_templates"
	moduleBayTemplatesKey         = "module_bay_templates"
	powerOutletTemplatesKey       = "power_outlet_templates"
	frontPortTemplatesKey         = "front_port_templates"
	inventoryItemTemplatesKey     = "inventory_item_templates"
)

// templateListPageLimit is large enough to cover any realistic per-device-type
// template count in a single page. NetBox imposes its own ceiling but this
// matches what the existing data-source paginators use as a practical cap.
const templateListPageLimit int64 = 1000

// templateRef is what each per-type list returns: the bare information the
// reconciler needs to do a name-keyed diff. Type-specific data comes back via
// the closure's Flatten function below, which is only used during Read.
type templateRef struct {
	ID   int64
	Name string
}

// templateOps wires generic name-keyed reconciliation to a specific NetBox
// template type. Every per-type block in this file builds one of these.
type templateOps struct {
	// kind is a short human-readable name for logs / errors ("power_port", etc.).
	kind string

	// list returns all templates of this kind that are currently attached to
	// the given device_type, plus an opaque payload usable for name -> id
	// lookups during sibling FK resolution.
	list func(api *providerState, deviceTypeID int64) ([]templateRef, error)

	// expand converts one TF schema map into a writable model ready to send.
	// resolveSibling, when non-nil, is consulted by FK-bearing types to
	// resolve a referenced sibling template by name. resolveSibling returns
	// 0 if the name is not found in the current desired set.
	expand func(item map[string]interface{}, deviceTypeID int64, resolveSibling func(name string) int64) (interface{}, error)

	// create issues the CREATE call and returns the new ID.
	create func(api *providerState, payload interface{}) (int64, error)

	// update issues the partial-update call against an existing template.
	update func(api *providerState, id int64, payload interface{}) error

	// del removes a template by ID.
	del func(api *providerState, id int64) error
}

// reconcileTemplates walks the desired TF set against the current API list
// (both keyed by name) and issues the create / partial-update / delete calls
// to converge. The resolveSibling callback is wired through to each expand()
// so FK-bearing types can look up sibling template IDs by name.
//
// Ownership rule: a template is considered "ours" iff its name appears in
// either the new desired set OR the previousNames set (i.e. it was tracked
// in our prior state). Templates in NetBox that are not in either set are
// owned by something else (e.g. a standalone netbox_interface_template
// resource, an external manual edit) and we leave them strictly alone — we
// neither read them into our state nor delete them. This is what makes
// coexistence with the standalone template resources actually safe: without
// this gate, the reconciler would silently mass-delete any template attached
// to the device_type that did not appear in the user's nested HCL blocks.
//
// Returns a name->id map of the templates we now own (created or already
// existed and were updated), so dependent passes (power_outlet, front_port)
// can resolve their sibling FKs by name.
func reconcileTemplates(api *providerState, deviceTypeID int64, ops templateOps, desired *schema.Set, previousNames map[string]struct{}, resolveSibling func(name string) int64) (map[string]int64, error) {
	current, err := ops.list(api, deviceTypeID)
	if err != nil {
		return nil, fmt.Errorf("listing %s templates for device_type %d: %w", ops.kind, deviceTypeID, err)
	}

	currentByName := make(map[string]int64, len(current))
	for _, c := range current {
		currentByName[c.Name] = c.ID
	}

	desiredByName := make(map[string]map[string]interface{})
	if desired != nil {
		for _, raw := range desired.List() {
			item, ok := raw.(map[string]interface{})
			if !ok {
				continue
			}
			name, _ := item["name"].(string)
			if name == "" {
				continue
			}
			desiredByName[name] = item
		}
	}

	finalIDs := make(map[string]int64, len(desiredByName))

	// Create or update everything in the desired set.
	for name, item := range desiredByName {
		payload, err := ops.expand(item, deviceTypeID, resolveSibling)
		if err != nil {
			return nil, fmt.Errorf("expanding %s template %q: %w", ops.kind, name, err)
		}
		if existingID, ok := currentByName[name]; ok {
			if err := ops.update(api, existingID, payload); err != nil {
				return nil, fmt.Errorf("updating %s template %q (id=%d): %w", ops.kind, name, existingID, err)
			}
			finalIDs[name] = existingID
		} else {
			newID, err := ops.create(api, payload)
			if err != nil {
				return nil, fmt.Errorf("creating %s template %q: %w", ops.kind, name, err)
			}
			finalIDs[name] = newID
		}
	}

	// Delete only templates we previously owned that are not in the new
	// desired set. Anything in NetBox we never knew about (not in
	// previousNames, not in desired) belongs to someone else — leave it.
	for name, id := range currentByName {
		if _, keep := desiredByName[name]; keep {
			continue
		}
		if _, owned := previousNames[name]; !owned {
			continue
		}
		if err := ops.del(api, id); err != nil {
			return nil, fmt.Errorf("deleting %s template %q (id=%d): %w", ops.kind, name, id, err)
		}
	}

	return finalIDs, nil
}

// templateNamesFromSet pulls the `name` field out of every entry in a TF set
// and returns them as a set-keyed map. nil-safe.
func templateNamesFromSet(s *schema.Set) map[string]struct{} {
	if s == nil {
		return map[string]struct{}{}
	}
	out := make(map[string]struct{}, s.Len())
	for _, raw := range s.List() {
		m, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}
		name, _ := m["name"].(string)
		if name != "" {
			out[name] = struct{}{}
		}
	}
	return out
}

// previousTemplateNames returns the set of template `name` values that were
// in this resource's prior state for `key`, before the current apply. It is
// our authoritative "we used to own these" record and drives both the
// Sync-time deletion gate and the Read-time filter.
func previousTemplateNames(d *schema.ResourceData, key string) map[string]struct{} {
	old, _ := d.GetChange(key)
	if oldSet, ok := old.(*schema.Set); ok {
		return templateNamesFromSet(oldSet)
	}
	return map[string]struct{}{}
}

// filterReadTemplatesByOwnership trims `fetched` (the full per-type list we
// just got from NetBox for this device_type) down to entries we actually
// own, and writes them to state. Ownership is the union of what's in current
// state for `key` and what's in `additional` (used for the post-Create read,
// where Sync has just created templates that don't exist in state yet).
//
// If we own zero templates of this kind, we leave state alone entirely
// (including not clearing it) — the caller probably never declared any
// nested blocks for this family and we should not advertise any drift on
// templates managed elsewhere.
func filterReadTemplatesByOwnership(d *schema.ResourceData, key string, fetched []map[string]interface{}, additional map[string]struct{}) error {
	owned := templateNamesFromSet(getTemplateSet(d, key))
	for name := range additional {
		owned[name] = struct{}{}
	}
	if len(owned) == 0 {
		return nil
	}
	out := make([]map[string]interface{}, 0, len(fetched))
	for _, m := range fetched {
		name, _ := m["name"].(string)
		if name == "" {
			continue
		}
		if _, ok := owned[name]; ok {
			out = append(out, m)
		}
	}
	return d.Set(key, out)
}

// syncDeviceTypeTemplates is the entry point called from
// resource_netbox_device_type.go Create/Update. It runs all 10 template
// reconcilers in dependency order against the live device_type and converges
// NetBox to match the user's HCL. Per-type ownership is derived from prior
// state via d.GetChange so we never delete templates we did not previously
// manage (see reconcileTemplates' coexistence-safety doc).
func syncDeviceTypeTemplates(d *schema.ResourceData, api *providerState, deviceTypeID int64) error {
	// Pass 1: independent types. No sibling FKs to resolve.
	ppIDs, err := reconcileTemplates(api, deviceTypeID, powerPortTemplateOps(), getTemplateSet(d, powerPortTemplatesKey), previousTemplateNames(d, powerPortTemplatesKey), nil)
	if err != nil {
		return err
	}
	if _, err := reconcileTemplates(api, deviceTypeID, interfaceTemplateOps(), getTemplateSet(d, interfaceTemplatesKey), previousTemplateNames(d, interfaceTemplatesKey), nil); err != nil {
		return err
	}
	if _, err := reconcileTemplates(api, deviceTypeID, consolePortTemplateOps(), getTemplateSet(d, consolePortTemplatesKey), previousTemplateNames(d, consolePortTemplatesKey), nil); err != nil {
		return err
	}
	if _, err := reconcileTemplates(api, deviceTypeID, consoleServerPortTemplateOps(), getTemplateSet(d, consoleServerPortTemplatesKey), previousTemplateNames(d, consoleServerPortTemplatesKey), nil); err != nil {
		return err
	}
	rpIDs, err := reconcileTemplates(api, deviceTypeID, rearPortTemplateOps(), getTemplateSet(d, rearPortTemplatesKey), previousTemplateNames(d, rearPortTemplatesKey), nil)
	if err != nil {
		return err
	}
	if _, err := reconcileTemplates(api, deviceTypeID, deviceBayTemplateOps(), getTemplateSet(d, deviceBayTemplatesKey), previousTemplateNames(d, deviceBayTemplatesKey), nil); err != nil {
		return err
	}
	if _, err := reconcileTemplates(api, deviceTypeID, moduleBayTemplateOps(), getTemplateSet(d, moduleBayTemplatesKey), previousTemplateNames(d, moduleBayTemplatesKey), nil); err != nil {
		return err
	}

	// Pass 2: dependent types. power_outlet references a sibling power_port
	// template by name; front_port references a sibling rear_port template
	// by name. Resolution closes over the IDs we just created above.
	ppLookup := func(name string) int64 { return ppIDs[name] }
	if _, err := reconcileTemplates(api, deviceTypeID, powerOutletTemplateOps(), getTemplateSet(d, powerOutletTemplatesKey), previousTemplateNames(d, powerOutletTemplatesKey), ppLookup); err != nil {
		return err
	}
	rpLookup := func(name string) int64 { return rpIDs[name] }
	if _, err := reconcileTemplates(api, deviceTypeID, frontPortTemplateOps(), getTemplateSet(d, frontPortTemplatesKey), previousTemplateNames(d, frontPortTemplatesKey), rpLookup); err != nil {
		return err
	}

	// Pass 3: inventory_item templates. These can reference sibling inventory
	// items as parents (forming a tree) and other templates as components via
	// the polymorphic component_type/component_id FK. The component_type/id
	// pair is taken as-is from the user (string + int) so callers can reference
	// any component object — no resolution needed here. Parents are resolved
	// after the fact by inventoryItemTemplateOps which walks the tree itself.
	if err := syncInventoryItemTemplates(api, deviceTypeID, getTemplateSet(d, inventoryItemTemplatesKey), previousTemplateNames(d, inventoryItemTemplatesKey)); err != nil {
		return err
	}

	return nil
}

// readDeviceTypeTemplates is called from Read after the parent device_type
// has been refreshed. It pulls every template family back from NetBox and
// flattens it into the corresponding TF set.
func readDeviceTypeTemplates(d *schema.ResourceData, api *providerState, deviceTypeID int64) error {
	if err := readPowerPortTemplates(d, api, deviceTypeID); err != nil {
		return err
	}
	if err := readInterfaceTemplates(d, api, deviceTypeID); err != nil {
		return err
	}
	if err := readConsolePortTemplates(d, api, deviceTypeID); err != nil {
		return err
	}
	if err := readConsoleServerPortTemplates(d, api, deviceTypeID); err != nil {
		return err
	}
	if err := readRearPortTemplates(d, api, deviceTypeID); err != nil {
		return err
	}
	if err := readDeviceBayTemplates(d, api, deviceTypeID); err != nil {
		return err
	}
	if err := readModuleBayTemplates(d, api, deviceTypeID); err != nil {
		return err
	}
	if err := readPowerOutletTemplates(d, api, deviceTypeID); err != nil {
		return err
	}
	if err := readFrontPortTemplates(d, api, deviceTypeID); err != nil {
		return err
	}
	if err := readInventoryItemTemplates(d, api, deviceTypeID); err != nil {
		return err
	}
	return nil
}

// getTemplateSet pulls the configured set out of ResourceData, returning nil
// if the user has not declared any blocks of this kind.
func getTemplateSet(d *schema.ResourceData, key string) *schema.Set {
	raw, ok := d.GetOk(key)
	if !ok {
		return nil
	}
	set, _ := raw.(*schema.Set)
	return set
}

// templateNameHash hashes a template block by name only. We expose this as
// the SetFunc for every TypeSet so adding/removing optional fields like
// `description` does not change the set membership identity.
func templateNameHash(v interface{}) int {
	m, ok := v.(map[string]interface{})
	if !ok {
		return 0
	}
	name, _ := m["name"].(string)
	return schema.HashString(name)
}

// commonTemplateSchema returns the field set that every template family
// shares: name (the hash key), id (computed back from the API), label,
// description.
func commonTemplateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the template. Must be unique within the parent device_type and is used as the identity key for the nested set.",
		},
		"id": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "NetBox-assigned ID of the template, populated after Create.",
		},
		"label": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Optional physical label, e.g. text printed on the chassis next to the port.",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Free-form description shown in the NetBox UI.",
		},
	}
}

// withCommon merges per-type fields with the common base. Per-type maps
// override base keys when there is a clash (e.g. type-specific descriptions).
func withCommon(specific map[string]*schema.Schema) map[string]*schema.Schema {
	merged := commonTemplateSchema()
	for k, v := range specific {
		merged[k] = v
	}
	return merged
}

// strFromMap pulls a string value out of a TF set map, returning "" if absent.
func strFromMap(m map[string]interface{}, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}

// intFromMap pulls an int value out of a TF set map, returning 0 if absent.
func intFromMap(m map[string]interface{}, key string) int {
	if v, ok := m[key].(int); ok {
		return v
	}
	return 0
}

// boolFromMap pulls a bool value out of a TF set map, returning false if absent.
func boolFromMap(m map[string]interface{}, key string) bool {
	if v, ok := m[key].(bool); ok {
		return v
	}
	return false
}

// deviceTypeIDStr formats the device_type ID for the *string list filter on
// every Dcim*TemplatesList params type.
func deviceTypeIDStr(id int64) *string {
	s := strconv.FormatInt(id, 10)
	return &s
}

// =============================================================================
// power_port templates
// =============================================================================

func powerPortTemplateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: withCommon(map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Power port connector type, e.g. `iec-60320-c14`. See the NetBox docs for the full enumeration.",
			},
			"maximum_draw": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Maximum power draw in watts.",
			},
			"allocated_draw": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Allocated power draw in watts.",
			},
		}),
	}
}

func powerPortTemplateOps() templateOps {
	return templateOps{
		kind: "power_port",
		list: func(api *providerState, deviceTypeID int64) ([]templateRef, error) {
			params := dcim.NewDcimPowerPortTemplatesListParams().
				WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
			limit := templateListPageLimit
			params.Limit = &limit
			res, err := api.Dcim.DcimPowerPortTemplatesList(params, nil)
			if err != nil {
				return nil, err
			}
			out := make([]templateRef, 0, len(res.Payload.Results))
			for _, r := range res.Payload.Results {
				if r.Name != nil {
					out = append(out, templateRef{ID: r.ID, Name: *r.Name})
				}
			}
			return out, nil
		},
		expand: func(item map[string]interface{}, deviceTypeID int64, _ func(string) int64) (interface{}, error) {
			name := strFromMap(item, "name")
			model := &models.WritablePowerPortTemplate{
				Name:        &name,
				DeviceType:  int64ToPtr(deviceTypeID),
				Label:       strFromMap(item, "label"),
				Description: strFromMap(item, "description"),
				Type:        strFromMap(item, "type"),
			}
			if v := intFromMap(item, "maximum_draw"); v != 0 {
				model.MaximumDraw = int64ToPtr(int64(v))
			}
			if v := intFromMap(item, "allocated_draw"); v != 0 {
				model.AllocatedDraw = int64ToPtr(int64(v))
			}
			return model, nil
		},
		create: func(api *providerState, payload interface{}) (int64, error) {
			params := dcim.NewDcimPowerPortTemplatesCreateParams().WithData(payload.(*models.WritablePowerPortTemplate))
			res, err := api.Dcim.DcimPowerPortTemplatesCreate(params, nil)
			if err != nil {
				return 0, err
			}
			return res.Payload.ID, nil
		},
		update: func(api *providerState, id int64, payload interface{}) error {
			params := dcim.NewDcimPowerPortTemplatesPartialUpdateParams().WithID(id).WithData(payload.(*models.WritablePowerPortTemplate))
			_, err := api.Dcim.DcimPowerPortTemplatesPartialUpdate(params, nil)
			return err
		},
		del: func(api *providerState, id int64) error {
			params := dcim.NewDcimPowerPortTemplatesDeleteParams().WithID(id)
			_, err := api.Dcim.DcimPowerPortTemplatesDelete(params, nil)
			return err
		},
	}
}

func readPowerPortTemplates(d *schema.ResourceData, api *providerState, deviceTypeID int64) error {
	params := dcim.NewDcimPowerPortTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
	limit := templateListPageLimit
	params.Limit = &limit
	res, err := api.Dcim.DcimPowerPortTemplatesList(params, nil)
	if err != nil {
		return fmt.Errorf("reading power_port templates: %w", err)
	}
	out := make([]map[string]interface{}, 0, len(res.Payload.Results))
	for _, r := range res.Payload.Results {
		m := map[string]interface{}{
			"id":          int(r.ID),
			"name":        ptrToStr(r.Name),
			"label":       r.Label,
			"description": r.Description,
		}
		if r.Type != nil && r.Type.Value != nil {
			m["type"] = *r.Type.Value
		}
		if r.MaximumDraw != nil {
			m["maximum_draw"] = int(*r.MaximumDraw)
		}
		if r.AllocatedDraw != nil {
			m["allocated_draw"] = int(*r.AllocatedDraw)
		}
		out = append(out, m)
	}
	return filterReadTemplatesByOwnership(d, powerPortTemplatesKey, out, nil)
}

// =============================================================================
// interface templates
// =============================================================================

func interfaceTemplateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: withCommon(map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Interface type, e.g. `1000base-t`, `25gbase-x-sfp28`. See the NetBox docs for the full enumeration.",
			},
			"mgmt_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, this interface is for out-of-band management only.",
			},
			"poe_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "PoE mode (`pd`, `pse`).",
			},
			"poe_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "PoE type, e.g. `type1-ieee802.3af`.",
			},
		}),
	}
}

func interfaceTemplateOps() templateOps {
	return templateOps{
		kind: "interface",
		list: func(api *providerState, deviceTypeID int64) ([]templateRef, error) {
			params := dcim.NewDcimInterfaceTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
			limit := templateListPageLimit
			params.Limit = &limit
			res, err := api.Dcim.DcimInterfaceTemplatesList(params, nil)
			if err != nil {
				return nil, err
			}
			out := make([]templateRef, 0, len(res.Payload.Results))
			for _, r := range res.Payload.Results {
				if r.Name != nil {
					out = append(out, templateRef{ID: r.ID, Name: *r.Name})
				}
			}
			return out, nil
		},
		expand: func(item map[string]interface{}, deviceTypeID int64, _ func(string) int64) (interface{}, error) {
			name := strFromMap(item, "name")
			ifType := strFromMap(item, "type")
			model := &models.WritableInterfaceTemplate{
				Name:        &name,
				DeviceType:  int64ToPtr(deviceTypeID),
				Label:       strFromMap(item, "label"),
				Description: strFromMap(item, "description"),
				Type:        &ifType,
				MgmtOnly:    boolFromMap(item, "mgmt_only"),
				PoeMode:     strFromMap(item, "poe_mode"),
				PoeType:     strFromMap(item, "poe_type"),
			}
			return model, nil
		},
		create: func(api *providerState, payload interface{}) (int64, error) {
			params := dcim.NewDcimInterfaceTemplatesCreateParams().WithData(payload.(*models.WritableInterfaceTemplate))
			res, err := api.Dcim.DcimInterfaceTemplatesCreate(params, nil)
			if err != nil {
				return 0, err
			}
			return res.Payload.ID, nil
		},
		update: func(api *providerState, id int64, payload interface{}) error {
			params := dcim.NewDcimInterfaceTemplatesPartialUpdateParams().WithID(id).WithData(payload.(*models.WritableInterfaceTemplate))
			_, err := api.Dcim.DcimInterfaceTemplatesPartialUpdate(params, nil)
			return err
		},
		del: func(api *providerState, id int64) error {
			params := dcim.NewDcimInterfaceTemplatesDeleteParams().WithID(id)
			_, err := api.Dcim.DcimInterfaceTemplatesDelete(params, nil)
			return err
		},
	}
}

func readInterfaceTemplates(d *schema.ResourceData, api *providerState, deviceTypeID int64) error {
	params := dcim.NewDcimInterfaceTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
	limit := templateListPageLimit
	params.Limit = &limit
	res, err := api.Dcim.DcimInterfaceTemplatesList(params, nil)
	if err != nil {
		return fmt.Errorf("reading interface templates: %w", err)
	}
	out := make([]map[string]interface{}, 0, len(res.Payload.Results))
	for _, r := range res.Payload.Results {
		m := map[string]interface{}{
			"id":          int(r.ID),
			"name":        ptrToStr(r.Name),
			"label":       r.Label,
			"description": r.Description,
			"mgmt_only":   r.MgmtOnly,
		}
		if r.Type != nil && r.Type.Value != nil {
			m["type"] = *r.Type.Value
		}
		if r.PoeMode != nil && r.PoeMode.Value != nil {
			m["poe_mode"] = *r.PoeMode.Value
		}
		if r.PoeType != nil && r.PoeType.Value != nil {
			m["poe_type"] = *r.PoeType.Value
		}
		out = append(out, m)
	}
	return filterReadTemplatesByOwnership(d, interfaceTemplatesKey, out, nil)
}

// =============================================================================
// console_port templates
// =============================================================================

func consolePortTemplateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: withCommon(map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Console port connector type, e.g. `de-9`, `rj-45`, `usb-c`. See the NetBox docs for the full enumeration.",
			},
		}),
	}
}

func consolePortTemplateOps() templateOps {
	return templateOps{
		kind: "console_port",
		list: func(api *providerState, deviceTypeID int64) ([]templateRef, error) {
			params := dcim.NewDcimConsolePortTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
			limit := templateListPageLimit
			params.Limit = &limit
			res, err := api.Dcim.DcimConsolePortTemplatesList(params, nil)
			if err != nil {
				return nil, err
			}
			out := make([]templateRef, 0, len(res.Payload.Results))
			for _, r := range res.Payload.Results {
				if r.Name != nil {
					out = append(out, templateRef{ID: r.ID, Name: *r.Name})
				}
			}
			return out, nil
		},
		expand: func(item map[string]interface{}, deviceTypeID int64, _ func(string) int64) (interface{}, error) {
			name := strFromMap(item, "name")
			model := &models.WritableConsolePortTemplate{
				Name:        &name,
				DeviceType:  int64ToPtr(deviceTypeID),
				Label:       strFromMap(item, "label"),
				Description: strFromMap(item, "description"),
				Type:        strFromMap(item, "type"),
			}
			return model, nil
		},
		create: func(api *providerState, payload interface{}) (int64, error) {
			params := dcim.NewDcimConsolePortTemplatesCreateParams().WithData(payload.(*models.WritableConsolePortTemplate))
			res, err := api.Dcim.DcimConsolePortTemplatesCreate(params, nil)
			if err != nil {
				return 0, err
			}
			return res.Payload.ID, nil
		},
		update: func(api *providerState, id int64, payload interface{}) error {
			params := dcim.NewDcimConsolePortTemplatesPartialUpdateParams().WithID(id).WithData(payload.(*models.WritableConsolePortTemplate))
			_, err := api.Dcim.DcimConsolePortTemplatesPartialUpdate(params, nil)
			return err
		},
		del: func(api *providerState, id int64) error {
			params := dcim.NewDcimConsolePortTemplatesDeleteParams().WithID(id)
			_, err := api.Dcim.DcimConsolePortTemplatesDelete(params, nil)
			return err
		},
	}
}

func readConsolePortTemplates(d *schema.ResourceData, api *providerState, deviceTypeID int64) error {
	params := dcim.NewDcimConsolePortTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
	limit := templateListPageLimit
	params.Limit = &limit
	res, err := api.Dcim.DcimConsolePortTemplatesList(params, nil)
	if err != nil {
		return fmt.Errorf("reading console_port templates: %w", err)
	}
	out := make([]map[string]interface{}, 0, len(res.Payload.Results))
	for _, r := range res.Payload.Results {
		m := map[string]interface{}{
			"id":          int(r.ID),
			"name":        ptrToStr(r.Name),
			"label":       r.Label,
			"description": r.Description,
		}
		if r.Type != nil && r.Type.Value != nil {
			m["type"] = *r.Type.Value
		}
		out = append(out, m)
	}
	return filterReadTemplatesByOwnership(d, consolePortTemplatesKey, out, nil)
}

// =============================================================================
// console_server_port templates
// =============================================================================

func consoleServerPortTemplateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: withCommon(map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Console server port connector type. See the NetBox docs for the full enumeration.",
			},
		}),
	}
}

func consoleServerPortTemplateOps() templateOps {
	return templateOps{
		kind: "console_server_port",
		list: func(api *providerState, deviceTypeID int64) ([]templateRef, error) {
			params := dcim.NewDcimConsoleServerPortTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
			limit := templateListPageLimit
			params.Limit = &limit
			res, err := api.Dcim.DcimConsoleServerPortTemplatesList(params, nil)
			if err != nil {
				return nil, err
			}
			out := make([]templateRef, 0, len(res.Payload.Results))
			for _, r := range res.Payload.Results {
				if r.Name != nil {
					out = append(out, templateRef{ID: r.ID, Name: *r.Name})
				}
			}
			return out, nil
		},
		expand: func(item map[string]interface{}, deviceTypeID int64, _ func(string) int64) (interface{}, error) {
			name := strFromMap(item, "name")
			model := &models.WritableConsoleServerPortTemplate{
				Name:        &name,
				DeviceType:  int64ToPtr(deviceTypeID),
				Label:       strFromMap(item, "label"),
				Description: strFromMap(item, "description"),
				Type:        strFromMap(item, "type"),
			}
			return model, nil
		},
		create: func(api *providerState, payload interface{}) (int64, error) {
			params := dcim.NewDcimConsoleServerPortTemplatesCreateParams().WithData(payload.(*models.WritableConsoleServerPortTemplate))
			res, err := api.Dcim.DcimConsoleServerPortTemplatesCreate(params, nil)
			if err != nil {
				return 0, err
			}
			return res.Payload.ID, nil
		},
		update: func(api *providerState, id int64, payload interface{}) error {
			params := dcim.NewDcimConsoleServerPortTemplatesPartialUpdateParams().WithID(id).WithData(payload.(*models.WritableConsoleServerPortTemplate))
			_, err := api.Dcim.DcimConsoleServerPortTemplatesPartialUpdate(params, nil)
			return err
		},
		del: func(api *providerState, id int64) error {
			params := dcim.NewDcimConsoleServerPortTemplatesDeleteParams().WithID(id)
			_, err := api.Dcim.DcimConsoleServerPortTemplatesDelete(params, nil)
			return err
		},
	}
}

func readConsoleServerPortTemplates(d *schema.ResourceData, api *providerState, deviceTypeID int64) error {
	params := dcim.NewDcimConsoleServerPortTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
	limit := templateListPageLimit
	params.Limit = &limit
	res, err := api.Dcim.DcimConsoleServerPortTemplatesList(params, nil)
	if err != nil {
		return fmt.Errorf("reading console_server_port templates: %w", err)
	}
	out := make([]map[string]interface{}, 0, len(res.Payload.Results))
	for _, r := range res.Payload.Results {
		m := map[string]interface{}{
			"id":          int(r.ID),
			"name":        ptrToStr(r.Name),
			"label":       r.Label,
			"description": r.Description,
		}
		if r.Type != nil && r.Type.Value != nil {
			m["type"] = *r.Type.Value
		}
		out = append(out, m)
	}
	return filterReadTemplatesByOwnership(d, consoleServerPortTemplatesKey, out, nil)
}

// =============================================================================
// rear_port templates
// =============================================================================

func rearPortTemplateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: withCommon(map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Rear port connector type, e.g. `8p8c`, `lc`, `mpo`. See the NetBox docs for the full enumeration.",
			},
			"positions": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "Number of front positions this rear port can be split into. Defaults to 1 if not set.",
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Hex color code (without leading `#`) used for the port in the UI.",
			},
		}),
	}
}

func rearPortTemplateOps() templateOps {
	return templateOps{
		kind: "rear_port",
		list: func(api *providerState, deviceTypeID int64) ([]templateRef, error) {
			params := dcim.NewDcimRearPortTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
			limit := templateListPageLimit
			params.Limit = &limit
			res, err := api.Dcim.DcimRearPortTemplatesList(params, nil)
			if err != nil {
				return nil, err
			}
			out := make([]templateRef, 0, len(res.Payload.Results))
			for _, r := range res.Payload.Results {
				if r.Name != nil {
					out = append(out, templateRef{ID: r.ID, Name: *r.Name})
				}
			}
			return out, nil
		},
		expand: func(item map[string]interface{}, deviceTypeID int64, _ func(string) int64) (interface{}, error) {
			name := strFromMap(item, "name")
			rpType := strFromMap(item, "type")
			model := &models.WritableRearPortTemplate{
				Name:        &name,
				DeviceType:  int64ToPtr(deviceTypeID),
				Label:       strFromMap(item, "label"),
				Description: strFromMap(item, "description"),
				Type:        &rpType,
				Color:       strFromMap(item, "color"),
			}
			if v := intFromMap(item, "positions"); v != 0 {
				model.Positions = int64(v)
			}
			return model, nil
		},
		create: func(api *providerState, payload interface{}) (int64, error) {
			params := dcim.NewDcimRearPortTemplatesCreateParams().WithData(payload.(*models.WritableRearPortTemplate))
			res, err := api.Dcim.DcimRearPortTemplatesCreate(params, nil)
			if err != nil {
				return 0, err
			}
			return res.Payload.ID, nil
		},
		update: func(api *providerState, id int64, payload interface{}) error {
			params := dcim.NewDcimRearPortTemplatesPartialUpdateParams().WithID(id).WithData(payload.(*models.WritableRearPortTemplate))
			_, err := api.Dcim.DcimRearPortTemplatesPartialUpdate(params, nil)
			return err
		},
		del: func(api *providerState, id int64) error {
			params := dcim.NewDcimRearPortTemplatesDeleteParams().WithID(id)
			_, err := api.Dcim.DcimRearPortTemplatesDelete(params, nil)
			return err
		},
	}
}

func readRearPortTemplates(d *schema.ResourceData, api *providerState, deviceTypeID int64) error {
	params := dcim.NewDcimRearPortTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
	limit := templateListPageLimit
	params.Limit = &limit
	res, err := api.Dcim.DcimRearPortTemplatesList(params, nil)
	if err != nil {
		return fmt.Errorf("reading rear_port templates: %w", err)
	}
	out := make([]map[string]interface{}, 0, len(res.Payload.Results))
	for _, r := range res.Payload.Results {
		m := map[string]interface{}{
			"id":          int(r.ID),
			"name":        ptrToStr(r.Name),
			"label":       r.Label,
			"description": r.Description,
			"positions":   int(r.Positions),
			"color":       r.Color,
		}
		if r.Type != nil && r.Type.Value != nil {
			m["type"] = *r.Type.Value
		}
		out = append(out, m)
	}
	return filterReadTemplatesByOwnership(d, rearPortTemplatesKey, out, nil)
}

// =============================================================================
// device_bay templates
// =============================================================================

func deviceBayTemplateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: withCommon(map[string]*schema.Schema{}),
	}
}

func deviceBayTemplateOps() templateOps {
	return templateOps{
		kind: "device_bay",
		list: func(api *providerState, deviceTypeID int64) ([]templateRef, error) {
			params := dcim.NewDcimDeviceBayTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
			limit := templateListPageLimit
			params.Limit = &limit
			res, err := api.Dcim.DcimDeviceBayTemplatesList(params, nil)
			if err != nil {
				return nil, err
			}
			out := make([]templateRef, 0, len(res.Payload.Results))
			for _, r := range res.Payload.Results {
				if r.Name != nil {
					out = append(out, templateRef{ID: r.ID, Name: *r.Name})
				}
			}
			return out, nil
		},
		expand: func(item map[string]interface{}, deviceTypeID int64, _ func(string) int64) (interface{}, error) {
			name := strFromMap(item, "name")
			model := &models.WritableDeviceBayTemplate{
				Name:        &name,
				DeviceType:  int64ToPtr(deviceTypeID),
				Label:       strFromMap(item, "label"),
				Description: strFromMap(item, "description"),
			}
			return model, nil
		},
		create: func(api *providerState, payload interface{}) (int64, error) {
			params := dcim.NewDcimDeviceBayTemplatesCreateParams().WithData(payload.(*models.WritableDeviceBayTemplate))
			res, err := api.Dcim.DcimDeviceBayTemplatesCreate(params, nil)
			if err != nil {
				return 0, err
			}
			return res.Payload.ID, nil
		},
		update: func(api *providerState, id int64, payload interface{}) error {
			params := dcim.NewDcimDeviceBayTemplatesPartialUpdateParams().WithID(id).WithData(payload.(*models.WritableDeviceBayTemplate))
			_, err := api.Dcim.DcimDeviceBayTemplatesPartialUpdate(params, nil)
			return err
		},
		del: func(api *providerState, id int64) error {
			params := dcim.NewDcimDeviceBayTemplatesDeleteParams().WithID(id)
			_, err := api.Dcim.DcimDeviceBayTemplatesDelete(params, nil)
			return err
		},
	}
}

func readDeviceBayTemplates(d *schema.ResourceData, api *providerState, deviceTypeID int64) error {
	params := dcim.NewDcimDeviceBayTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
	limit := templateListPageLimit
	params.Limit = &limit
	res, err := api.Dcim.DcimDeviceBayTemplatesList(params, nil)
	if err != nil {
		return fmt.Errorf("reading device_bay templates: %w", err)
	}
	out := make([]map[string]interface{}, 0, len(res.Payload.Results))
	for _, r := range res.Payload.Results {
		m := map[string]interface{}{
			"id":          int(r.ID),
			"name":        ptrToStr(r.Name),
			"label":       r.Label,
			"description": r.Description,
		}
		out = append(out, m)
	}
	return filterReadTemplatesByOwnership(d, deviceBayTemplatesKey, out, nil)
}

// =============================================================================
// module_bay templates
// =============================================================================

func moduleBayTemplateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: withCommon(map[string]*schema.Schema{
			"position": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Position designator inside the chassis, used by NetBox when {module} substitution is performed on child component template names.",
			},
		}),
	}
}

func moduleBayTemplateOps() templateOps {
	return templateOps{
		kind: "module_bay",
		list: func(api *providerState, deviceTypeID int64) ([]templateRef, error) {
			params := dcim.NewDcimModuleBayTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
			limit := templateListPageLimit
			params.Limit = &limit
			res, err := api.Dcim.DcimModuleBayTemplatesList(params, nil)
			if err != nil {
				return nil, err
			}
			out := make([]templateRef, 0, len(res.Payload.Results))
			for _, r := range res.Payload.Results {
				if r.Name != nil {
					out = append(out, templateRef{ID: r.ID, Name: *r.Name})
				}
			}
			return out, nil
		},
		expand: func(item map[string]interface{}, deviceTypeID int64, _ func(string) int64) (interface{}, error) {
			name := strFromMap(item, "name")
			model := &models.WritableModuleBayTemplate{
				Name:        &name,
				DeviceType:  int64ToPtr(deviceTypeID),
				Label:       strFromMap(item, "label"),
				Description: strFromMap(item, "description"),
				Position:    strFromMap(item, "position"),
			}
			return model, nil
		},
		create: func(api *providerState, payload interface{}) (int64, error) {
			params := dcim.NewDcimModuleBayTemplatesCreateParams().WithData(payload.(*models.WritableModuleBayTemplate))
			res, err := api.Dcim.DcimModuleBayTemplatesCreate(params, nil)
			if err != nil {
				return 0, err
			}
			return res.Payload.ID, nil
		},
		update: func(api *providerState, id int64, payload interface{}) error {
			params := dcim.NewDcimModuleBayTemplatesPartialUpdateParams().WithID(id).WithData(payload.(*models.WritableModuleBayTemplate))
			_, err := api.Dcim.DcimModuleBayTemplatesPartialUpdate(params, nil)
			return err
		},
		del: func(api *providerState, id int64) error {
			params := dcim.NewDcimModuleBayTemplatesDeleteParams().WithID(id)
			_, err := api.Dcim.DcimModuleBayTemplatesDelete(params, nil)
			return err
		},
	}
}

func readModuleBayTemplates(d *schema.ResourceData, api *providerState, deviceTypeID int64) error {
	params := dcim.NewDcimModuleBayTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
	limit := templateListPageLimit
	params.Limit = &limit
	res, err := api.Dcim.DcimModuleBayTemplatesList(params, nil)
	if err != nil {
		return fmt.Errorf("reading module_bay templates: %w", err)
	}
	out := make([]map[string]interface{}, 0, len(res.Payload.Results))
	for _, r := range res.Payload.Results {
		m := map[string]interface{}{
			"id":          int(r.ID),
			"name":        ptrToStr(r.Name),
			"label":       r.Label,
			"description": r.Description,
			"position":    r.Position,
		}
		out = append(out, m)
	}
	return filterReadTemplatesByOwnership(d, moduleBayTemplatesKey, out, nil)
}

// =============================================================================
// power_outlet templates  (FK -> power_port template by name)
// =============================================================================

func powerOutletTemplateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: withCommon(map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Power outlet connector type, e.g. `iec-60320-c13`. See the NetBox docs for the full enumeration.",
			},
			"power_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the sibling `power_port_templates` block this outlet is downstream of. Resolved to the corresponding template ID at apply time.",
			},
			"feed_leg": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Power feed leg this outlet is connected to. Valid values are `A`, `B`, `C`.",
			},
		}),
	}
}

func powerOutletTemplateOps() templateOps {
	return templateOps{
		kind: "power_outlet",
		list: func(api *providerState, deviceTypeID int64) ([]templateRef, error) {
			params := dcim.NewDcimPowerOutletTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
			limit := templateListPageLimit
			params.Limit = &limit
			res, err := api.Dcim.DcimPowerOutletTemplatesList(params, nil)
			if err != nil {
				return nil, err
			}
			out := make([]templateRef, 0, len(res.Payload.Results))
			for _, r := range res.Payload.Results {
				if r.Name != nil {
					out = append(out, templateRef{ID: r.ID, Name: *r.Name})
				}
			}
			return out, nil
		},
		expand: func(item map[string]interface{}, deviceTypeID int64, resolveSibling func(string) int64) (interface{}, error) {
			name := strFromMap(item, "name")
			model := &models.WritablePowerOutletTemplate{
				Name:        &name,
				DeviceType:  int64ToPtr(deviceTypeID),
				Label:       strFromMap(item, "label"),
				Description: strFromMap(item, "description"),
				Type:        strFromMap(item, "type"),
				FeedLeg:     strFromMap(item, "feed_leg"),
			}
			if ppName := strFromMap(item, "power_port"); ppName != "" {
				if resolveSibling == nil {
					return nil, fmt.Errorf("power_outlet %q references power_port %q but no sibling resolver was provided (this is a provider bug)", name, ppName)
				}
				ppID := resolveSibling(ppName)
				if ppID == 0 {
					return nil, fmt.Errorf("power_outlet %q references power_port template %q which is not declared on this device_type", name, ppName)
				}
				model.PowerPort = int64ToPtr(ppID)
			}
			return model, nil
		},
		create: func(api *providerState, payload interface{}) (int64, error) {
			params := dcim.NewDcimPowerOutletTemplatesCreateParams().WithData(payload.(*models.WritablePowerOutletTemplate))
			res, err := api.Dcim.DcimPowerOutletTemplatesCreate(params, nil)
			if err != nil {
				return 0, err
			}
			return res.Payload.ID, nil
		},
		update: func(api *providerState, id int64, payload interface{}) error {
			params := dcim.NewDcimPowerOutletTemplatesPartialUpdateParams().WithID(id).WithData(payload.(*models.WritablePowerOutletTemplate))
			_, err := api.Dcim.DcimPowerOutletTemplatesPartialUpdate(params, nil)
			return err
		},
		del: func(api *providerState, id int64) error {
			params := dcim.NewDcimPowerOutletTemplatesDeleteParams().WithID(id)
			_, err := api.Dcim.DcimPowerOutletTemplatesDelete(params, nil)
			return err
		},
	}
}

func readPowerOutletTemplates(d *schema.ResourceData, api *providerState, deviceTypeID int64) error {
	params := dcim.NewDcimPowerOutletTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
	limit := templateListPageLimit
	params.Limit = &limit
	res, err := api.Dcim.DcimPowerOutletTemplatesList(params, nil)
	if err != nil {
		return fmt.Errorf("reading power_outlet templates: %w", err)
	}
	out := make([]map[string]interface{}, 0, len(res.Payload.Results))
	for _, r := range res.Payload.Results {
		m := map[string]interface{}{
			"id":          int(r.ID),
			"name":        ptrToStr(r.Name),
			"label":       r.Label,
			"description": r.Description,
		}
		if r.Type != nil && r.Type.Value != nil {
			m["type"] = *r.Type.Value
		}
		if r.FeedLeg != nil && r.FeedLeg.Value != nil {
			m["feed_leg"] = *r.FeedLeg.Value
		}
		if r.PowerPort != nil && r.PowerPort.Name != nil {
			m["power_port"] = *r.PowerPort.Name
		}
		out = append(out, m)
	}
	return filterReadTemplatesByOwnership(d, powerOutletTemplatesKey, out, nil)
}

// =============================================================================
// front_port templates  (FK -> rear_port template by name)
// =============================================================================

func frontPortTemplateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: withCommon(map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Front port connector type, e.g. `8p8c`, `lc`, `mpo`. See the NetBox docs for the full enumeration.",
			},
			"rear_port": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the sibling `rear_port_templates` block this front port is mapped to. Resolved to the corresponding template ID at apply time.",
			},
			"rear_port_position": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "Which numbered position on the rear port this front port maps to. Defaults to `1`.",
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Hex color code (without leading `#`) used for the port in the UI.",
			},
		}),
	}
}

func frontPortTemplateOps() templateOps {
	return templateOps{
		kind: "front_port",
		list: func(api *providerState, deviceTypeID int64) ([]templateRef, error) {
			params := dcim.NewDcimFrontPortTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
			limit := templateListPageLimit
			params.Limit = &limit
			res, err := api.Dcim.DcimFrontPortTemplatesList(params, nil)
			if err != nil {
				return nil, err
			}
			out := make([]templateRef, 0, len(res.Payload.Results))
			for _, r := range res.Payload.Results {
				if r.Name != nil {
					out = append(out, templateRef{ID: r.ID, Name: *r.Name})
				}
			}
			return out, nil
		},
		expand: func(item map[string]interface{}, deviceTypeID int64, resolveSibling func(string) int64) (interface{}, error) {
			name := strFromMap(item, "name")
			fpType := strFromMap(item, "type")
			model := &models.WritableFrontPortTemplate{
				Name:        &name,
				DeviceType:  int64ToPtr(deviceTypeID),
				Label:       strFromMap(item, "label"),
				Description: strFromMap(item, "description"),
				Type:        &fpType,
				Color:       strFromMap(item, "color"),
			}
			if v := intFromMap(item, "rear_port_position"); v != 0 {
				model.RearPortPosition = int64(v)
			}
			rpName := strFromMap(item, "rear_port")
			if rpName == "" {
				return nil, fmt.Errorf("front_port %q is missing the required rear_port reference", name)
			}
			if resolveSibling == nil {
				return nil, fmt.Errorf("front_port %q references rear_port %q but no sibling resolver was provided (this is a provider bug)", name, rpName)
			}
			rpID := resolveSibling(rpName)
			if rpID == 0 {
				return nil, fmt.Errorf("front_port %q references rear_port template %q which is not declared on this device_type", name, rpName)
			}
			model.RearPort = int64ToPtr(rpID)
			return model, nil
		},
		create: func(api *providerState, payload interface{}) (int64, error) {
			params := dcim.NewDcimFrontPortTemplatesCreateParams().WithData(payload.(*models.WritableFrontPortTemplate))
			res, err := api.Dcim.DcimFrontPortTemplatesCreate(params, nil)
			if err != nil {
				return 0, err
			}
			return res.Payload.ID, nil
		},
		update: func(api *providerState, id int64, payload interface{}) error {
			params := dcim.NewDcimFrontPortTemplatesPartialUpdateParams().WithID(id).WithData(payload.(*models.WritableFrontPortTemplate))
			_, err := api.Dcim.DcimFrontPortTemplatesPartialUpdate(params, nil)
			return err
		},
		del: func(api *providerState, id int64) error {
			params := dcim.NewDcimFrontPortTemplatesDeleteParams().WithID(id)
			_, err := api.Dcim.DcimFrontPortTemplatesDelete(params, nil)
			return err
		},
	}
}

func readFrontPortTemplates(d *schema.ResourceData, api *providerState, deviceTypeID int64) error {
	params := dcim.NewDcimFrontPortTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
	limit := templateListPageLimit
	params.Limit = &limit
	res, err := api.Dcim.DcimFrontPortTemplatesList(params, nil)
	if err != nil {
		return fmt.Errorf("reading front_port templates: %w", err)
	}
	out := make([]map[string]interface{}, 0, len(res.Payload.Results))
	for _, r := range res.Payload.Results {
		m := map[string]interface{}{
			"id":                 int(r.ID),
			"name":               ptrToStr(r.Name),
			"label":              r.Label,
			"description":        r.Description,
			"color":              r.Color,
			"rear_port_position": int(r.RearPortPosition),
		}
		if r.Type != nil && r.Type.Value != nil {
			m["type"] = *r.Type.Value
		}
		if r.RearPort != nil && r.RearPort.Name != nil {
			m["rear_port"] = *r.RearPort.Name
		}
		out = append(out, m)
	}
	return filterReadTemplatesByOwnership(d, frontPortTemplatesKey, out, nil)
}

// =============================================================================
// inventory_item templates  (parent tree + polymorphic component_type/component_id)
// =============================================================================

func inventoryItemTemplateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: withCommon(map[string]*schema.Schema{
			"parent": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the sibling `inventory_item_templates` block that should be the parent of this item. Forms a tree; the root has no parent.",
			},
			"manufacturer_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Optional manufacturer ID for this inventory item.",
			},
			"part_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Manufacturer part number / SKU for this inventory item.",
			},
			"role_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Optional inventory item role ID.",
			},
			"component_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Polymorphic FK type, e.g. `dcim.interfacetemplate`, `dcim.consoleporttemplate`. Pair with `component_id` to attach this inventory item to another component on the same device_type.",
			},
			"component_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Polymorphic FK target ID. Use the `id` computed attribute of another nested template to wire this up.",
			},
		}),
	}
}

// syncInventoryItemTemplates is the inventory_item-specific reconciliation. It
// is broken out from the generic reconciler because (a) it has to walk the
// parent tree in dependency order, and (b) the parent FK resolves against the
// IDs of templates we are currently in the middle of creating.
// syncInventoryItemTemplates is the inventory_item-specific reconciler — it
// can't share reconcileTemplates because of parent/child tree ordering on
// both apply and delete. The previousNames argument plays the same
// "ownership gate" role as in reconcileTemplates: only inventory_item
// templates whose names are in either previousNames or the new desired set
// are touched. Anything else attached to the device_type is someone else's.
func syncInventoryItemTemplates(api *providerState, deviceTypeID int64, desired *schema.Set, previousNames map[string]struct{}) error {
	// Pull current state.
	listParams := dcim.NewDcimInventoryItemTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
	limit := templateListPageLimit
	listParams.Limit = &limit
	listRes, err := api.Dcim.DcimInventoryItemTemplatesList(listParams, nil)
	if err != nil {
		return fmt.Errorf("listing inventory_item templates for device_type %d: %w", deviceTypeID, err)
	}

	currentByName := make(map[string]int64, len(listRes.Payload.Results))
	for _, r := range listRes.Payload.Results {
		if r.Name != nil {
			currentByName[*r.Name] = r.ID
		}
	}

	desiredByName := make(map[string]map[string]interface{})
	if desired != nil {
		for _, raw := range desired.List() {
			item, ok := raw.(map[string]interface{})
			if !ok {
				continue
			}
			name, _ := item["name"].(string)
			if name == "" {
				continue
			}
			desiredByName[name] = item
		}
	}

	// Compute the parent tree depth for each desired item so we can sync from
	// roots down to leaves. Cycles are rejected.
	depths, err := inventoryItemDepths(desiredByName)
	if err != nil {
		return err
	}

	ordered := make([]string, 0, len(desiredByName))
	for name := range desiredByName {
		ordered = append(ordered, name)
	}
	sort.SliceStable(ordered, func(i, j int) bool {
		if depths[ordered[i]] != depths[ordered[j]] {
			return depths[ordered[i]] < depths[ordered[j]]
		}
		return ordered[i] < ordered[j]
	})

	finalIDs := make(map[string]int64, len(ordered))

	for _, name := range ordered {
		item := desiredByName[name]
		var parentID *int64
		if pname := strFromMap(item, "parent"); pname != "" {
			id, ok := finalIDs[pname]
			if !ok {
				return fmt.Errorf("inventory_item %q references parent %q which is not declared on this device_type", name, pname)
			}
			parentID = int64ToPtr(id)
		}
		nameCopy := name
		model := &models.WritableInventoryItemTemplate{
			Name:        &nameCopy,
			DeviceType:  int64ToPtr(deviceTypeID),
			Label:       strFromMap(item, "label"),
			Description: strFromMap(item, "description"),
			PartID:      strFromMap(item, "part_id"),
			Parent:      parentID,
		}
		if v := intFromMap(item, "manufacturer_id"); v != 0 {
			model.Manufacturer = int64ToPtr(int64(v))
		}
		if v := intFromMap(item, "role_id"); v != 0 {
			model.Role = int64ToPtr(int64(v))
		}
		if ct := strFromMap(item, "component_type"); ct != "" {
			ctCopy := ct
			model.ComponentType = &ctCopy
		}
		if v := intFromMap(item, "component_id"); v != 0 {
			model.ComponentID = int64ToPtr(int64(v))
		}

		if existingID, ok := currentByName[name]; ok {
			updateParams := dcim.NewDcimInventoryItemTemplatesPartialUpdateParams().WithID(existingID).WithData(model)
			if _, err := api.Dcim.DcimInventoryItemTemplatesPartialUpdate(updateParams, nil); err != nil {
				return fmt.Errorf("updating inventory_item template %q (id=%d): %w", name, existingID, err)
			}
			finalIDs[name] = existingID
		} else {
			createParams := dcim.NewDcimInventoryItemTemplatesCreateParams().WithData(model)
			res, err := api.Dcim.DcimInventoryItemTemplatesCreate(createParams, nil)
			if err != nil {
				return fmt.Errorf("creating inventory_item template %q: %w", name, err)
			}
			finalIDs[name] = res.Payload.ID
		}
	}

	// Delete only items we previously owned that are no longer desired.
	// Anything in NetBox that isn't in previousNames is owned elsewhere
	// (standalone inventory_item resource, manual NetBox edit, etc.) and
	// must be left alone — the same coexistence safety as reconcileTemplates.
	for _, name := range ordered {
		delete(currentByName, name)
	}
	for name := range currentByName {
		if _, owned := previousNames[name]; !owned {
			delete(currentByName, name)
		}
	}
	if len(currentByName) > 0 {
		// Re-pull with a cap to know depths/parents on the deletion candidates.
		listRes, err := api.Dcim.DcimInventoryItemTemplatesList(listParams, nil)
		if err != nil {
			return fmt.Errorf("re-listing inventory_item templates: %w", err)
		}
		// Build child-list: parentID -> []ID
		children := make(map[int64][]int64)
		toDelete := make(map[int64]string)
		for _, r := range listRes.Payload.Results {
			if r.Name == nil {
				continue
			}
			if _, isOrphan := currentByName[*r.Name]; !isOrphan {
				continue
			}
			toDelete[r.ID] = *r.Name
			if r.Parent != nil {
				children[*r.Parent] = append(children[*r.Parent], r.ID)
			}
		}
		// Iteratively delete leaves until nothing left to delete.
		for len(toDelete) > 0 {
			progressed := false
			for id, name := range toDelete {
				hasChildToDelete := false
				for _, childID := range children[id] {
					if _, stillThere := toDelete[childID]; stillThere {
						hasChildToDelete = true
						break
					}
				}
				if hasChildToDelete {
					continue
				}
				delParams := dcim.NewDcimInventoryItemTemplatesDeleteParams().WithID(id)
				if _, err := api.Dcim.DcimInventoryItemTemplatesDelete(delParams, nil); err != nil {
					return fmt.Errorf("deleting inventory_item template %q (id=%d): %w", name, id, err)
				}
				delete(toDelete, id)
				progressed = true
			}
			if !progressed {
				return fmt.Errorf("inventory_item delete pass made no progress; possible cycle in parent references")
			}
		}
	}

	return nil
}

// inventoryItemDepths assigns each desired inventory_item a tree depth, where
// roots are 0. It rejects cycles and parents that are not themselves declared
// in the desired set (since the parent FK is resolved by name from siblings,
// not by ID).
func inventoryItemDepths(items map[string]map[string]interface{}) (map[string]int, error) {
	depths := make(map[string]int, len(items))
	visiting := make(map[string]bool, len(items))

	// resolve returns the depth of `name`. `referrer` is the name that
	// triggered the lookup, used only for nicer error messages.
	var resolve func(name, referrer string) (int, error)
	resolve = func(name, referrer string) (int, error) {
		if d, ok := depths[name]; ok {
			return d, nil
		}
		if visiting[name] {
			return 0, fmt.Errorf("inventory_item %q is in a cycle of parent references", name)
		}
		item, ok := items[name]
		if !ok {
			return 0, fmt.Errorf("inventory_item %q references parent %q which is not declared on this device_type", referrer, name)
		}
		parent := strFromMap(item, "parent")
		if parent == "" {
			depths[name] = 0
			return 0, nil
		}
		visiting[name] = true
		pd, err := resolve(parent, name)
		visiting[name] = false
		if err != nil {
			return 0, err
		}
		depths[name] = pd + 1
		return depths[name], nil
	}

	for name := range items {
		if _, err := resolve(name, name); err != nil {
			return nil, err
		}
	}
	return depths, nil
}

func readInventoryItemTemplates(d *schema.ResourceData, api *providerState, deviceTypeID int64) error {
	params := dcim.NewDcimInventoryItemTemplatesListParams().WithDevicetypeID(deviceTypeIDStr(deviceTypeID))
	limit := templateListPageLimit
	params.Limit = &limit
	res, err := api.Dcim.DcimInventoryItemTemplatesList(params, nil)
	if err != nil {
		return fmt.Errorf("reading inventory_item templates: %w", err)
	}
	// Build id -> name map so we can flatten parent FKs back into names.
	idToName := make(map[int64]string, len(res.Payload.Results))
	for _, r := range res.Payload.Results {
		if r.Name != nil {
			idToName[r.ID] = *r.Name
		}
	}
	out := make([]map[string]interface{}, 0, len(res.Payload.Results))
	for _, r := range res.Payload.Results {
		m := map[string]interface{}{
			"id":          int(r.ID),
			"name":        ptrToStr(r.Name),
			"label":       r.Label,
			"description": r.Description,
			"part_id":     r.PartID,
		}
		if r.Manufacturer != nil {
			m["manufacturer_id"] = int(r.Manufacturer.ID)
		}
		if r.Role != nil {
			m["role_id"] = int(r.Role.ID)
		}
		if r.ComponentType != nil {
			m["component_type"] = *r.ComponentType
		}
		if r.ComponentID != nil {
			m["component_id"] = int(*r.ComponentID)
		}
		if r.Parent != nil {
			if pn, ok := idToName[*r.Parent]; ok {
				m["parent"] = pn
			}
		}
		out = append(out, m)
	}
	return filterReadTemplatesByOwnership(d, inventoryItemTemplatesKey, out, nil)
}

// ptrToStr is a small helper that dereferences a *string, returning "" for nil.
func ptrToStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
