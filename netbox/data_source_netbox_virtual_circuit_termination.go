package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxVirtualCircuitTermination() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxVirtualCircuitTerminationRead,
		Description: `:meta:subcategory:Circuits:`,
		Schema: map[string]*schema.Schema{
			"id": {
				Required: true,
				Type:     schema.TypeString,
			},
			"virtual_circuit_id": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"interface_id": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"role": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"description": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"custom_fields": {
				Computed: true,
				Type:     schema.TypeMap,
			},
			tagsKey: tagsSchemaRead,
		},
	}
}

func dataSourceNetboxVirtualCircuitTerminationRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, ok := d.Get("id").(string)
	if !ok || id == "" {
		return errors.New("id is required")
	}

	params := circuits.NewCircuitsVirtualCircuitTerminationsListParams()
	params.ID = &id

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Circuits.CircuitsVirtualCircuitTerminationsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no virtual circuit termination found matching filter")
	}
	result := res.GetPayload().Results[0]

	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("description", result.Description)

	if result.Role != nil {
		d.Set("role", result.Role.Value)
	}

	if result.VirtualCircuit != nil {
		d.Set("virtual_circuit_id", result.VirtualCircuit.ID)
	}

	if result.Interface != nil {
		d.Set("interface_id", result.Interface.ID)
	}

	if result.CustomFields != nil {
		d.Set("custom_fields", flattenCustomFields(result.CustomFields))
	}

	d.Set(tagsKey, getTagListFromNestedTagList(result.Tags))
	return nil
}
