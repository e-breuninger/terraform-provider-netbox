package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxCircuitTerminationTermSideOptions = []string{"A", "Z"}

func resourceNetboxCircuitTermination() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxCircuitTerminationCreate,
		Read:   resourceNetboxCircuitTerminationRead,
		Update: resourceNetboxCircuitTerminationUpdate,
		Delete: resourceNetboxCircuitTerminationDelete,

		Description: `:meta:subcategory:Circuits:From the [official documentation](https://docs.netbox.dev/en/stable/features/circuits/#circuit-terminations):

> The association of a circuit with a particular site and/or device is modeled separately as a circuit termination. A circuit may have up to two terminations, labeled A and Z. A single-termination circuit can be used when you don't know (or care) about the far end of a circuit (for example, an Internet access circuit which connects to a transit provider). A dual-termination circuit is useful for tracking circuits which connect two sites.
>
> Each circuit termination is attached to either a site or to a provider network. Site terminations may optionally be connected via a cable to a specific device interface or port within that site. Each termination must be assigned a port speed, and can optionally be assigned an upstream speed if it differs from the downstream speed (a common scenario with e.g. DOCSIS cable modems). Fields are also available to track cross-connect and patch panel details.`,

		Schema: map[string]*schema.Schema{
			"circuit_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"location_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				ExactlyOneOf: []string{"site_id", "site_group_id", "region_id", "provider_network_id"},
			},
			"site_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				ExactlyOneOf: []string{"location_id", "site_group_id", "region_id", "provider_network_id"},
			},
			"site_group_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				ExactlyOneOf: []string{"location_id", "site_id", "region_id", "provider_network_id"},
			},
			"region_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				ExactlyOneOf: []string{"location_id", "site_id", "site_group_id", "provider_network_id"},
			},
			"provider_network_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				ExactlyOneOf: []string{"location_id", "site_id", "site_group_id", "region_id"},
			},
			"port_speed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"upstream_speed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"term_side": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxCircuitTerminationTermSideOptions, false),
				Description:  buildValidValueDescription(resourceNetboxCircuitTerminationTermSideOptions),
			},
			tagsKey:         tagsSchema,
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxCircuitTerminationCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableCircuitTermination{}

	termside := d.Get("term_side").(string)
	data.TermSide = &termside

	circuitIDValue, ok := d.GetOk("circuit_id")
	if ok {
		data.Circuit = int64ToPtr(int64(circuitIDValue.(int)))
	}

	siteID := getOptionalInt(d, "site_id")
	siteGroupID := getOptionalInt(d, "site_group_id")
	locationID := getOptionalInt(d, "location_id")
	regionID := getOptionalInt(d, "region_id")
	providerNetworkID := getOptionalInt(d, "provider_network_id")

	switch {
	case siteID != nil:
		data.TerminationType = strToPtr("dcim.site")
		data.TerminationID = siteID
	case siteGroupID != nil:
		data.TerminationType = strToPtr("dcim.sitegroup")
		data.TerminationID = siteGroupID
	case locationID != nil:
		data.TerminationType = strToPtr("dcim.location")
		data.TerminationID = locationID
	case regionID != nil:
		data.TerminationType = strToPtr("dcim.region")
		data.TerminationID = regionID
	case providerNetworkID != nil:
		data.TerminationType = strToPtr("circuits.providernetwork")
		data.TerminationID = providerNetworkID
	default:
		data.TerminationType = nil
		data.TerminationID = nil
	}

	portspeedValue, ok := d.GetOk("port_speed")
	if ok {
		data.PortSpeed = int64ToPtr(int64(portspeedValue.(int)))
	}

	upstreamspeedValue, ok := d.GetOk("upstream_speed")
	if ok {
		data.UpstreamSpeed = int64ToPtr(int64(upstreamspeedValue.(int)))
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := circuits.NewCircuitsCircuitTerminationsCreateParams().WithData(&data)

	res, err := api.Circuits.CircuitsCircuitTerminationsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxCircuitTerminationRead(d, m)
}

func resourceNetboxCircuitTerminationRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsCircuitTerminationsReadParams().WithID(id)

	res, err := api.Circuits.CircuitsCircuitTerminationsRead(params, nil)

	if err != nil {
		if errresp, ok := err.(*circuits.CircuitsCircuitTerminationsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	term := res.GetPayload()

	d.Set("term_side", term.TermSide)

	if term.Circuit != nil {
		d.Set("circuit_id", term.Circuit.ID)
	} else {
		d.Set("circuit_id", nil)
	}

	d.Set("site_id", nil)
	d.Set("site_group_id", nil)
	d.Set("location_id", nil)
	d.Set("region_id", nil)
	d.Set("provider_network_id", nil)

	if term.TerminationType != nil && term.TerminationID != nil {
		scopeID := term.TerminationID
		switch scopeType := term.TerminationType; *scopeType {
		case "dcim.site":
			d.Set("site_id", scopeID)
		case "dcim.sitegroup":
			d.Set("site_group_id", scopeID)
		case "dcim.location":
			d.Set("location_id", scopeID)
		case "dcim.region":
			d.Set("region_id", scopeID)
		case "circuits.providernetwork":
			d.Set("provider_network_id", scopeID)
		}
	}

	if term.PortSpeed != nil {
		d.Set("port_speed", term.PortSpeed)
	} else {
		d.Set("port_speed", nil)
	}

	if term.UpstreamSpeed != nil {
		d.Set("upstream_speed", term.UpstreamSpeed)
	} else {
		d.Set("upstream_speed", nil)
	}

	api.readTags(d, term.Tags)

	cf := getCustomFields(term.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}

	return nil
}

func resourceNetboxCircuitTerminationUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableCircuitTermination{}

	termside := d.Get("term_side").(string)
	data.TermSide = &termside

	circuitIDValue, ok := d.GetOk("circuit_id")
	if ok {
		data.Circuit = int64ToPtr(int64(circuitIDValue.(int)))
	}

	siteID := getOptionalInt(d, "site_id")
	siteGroupID := getOptionalInt(d, "site_group_id")
	locationID := getOptionalInt(d, "location_id")
	regionID := getOptionalInt(d, "region_id")
	providerNetworkID := getOptionalInt(d, "provider_network_id")

	switch {
	case siteID != nil:
		data.TerminationType = strToPtr("dcim.site")
		data.TerminationID = siteID
	case siteGroupID != nil:
		data.TerminationType = strToPtr("dcim.sitegroup")
		data.TerminationID = siteGroupID
	case locationID != nil:
		data.TerminationType = strToPtr("dcim.location")
		data.TerminationID = locationID
	case regionID != nil:
		data.TerminationType = strToPtr("dcim.region")
		data.TerminationID = regionID
	case providerNetworkID != nil:
		data.TerminationType = strToPtr("circuits.providernetwork")
		data.TerminationID = providerNetworkID
	default:
		data.TerminationType = nil
		data.TerminationID = nil
	}

	portspeedValue, ok := d.GetOk("port_speed")
	if ok {
		data.PortSpeed = int64ToPtr(int64(portspeedValue.(int)))
	}

	upstreamspeedValue, ok := d.GetOk("upstream_speed")
	if ok {
		data.UpstreamSpeed = int64ToPtr(int64(upstreamspeedValue.(int)))
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := circuits.NewCircuitsCircuitTerminationsPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Circuits.CircuitsCircuitTerminationsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxCircuitTerminationRead(d, m)
}

func resourceNetboxCircuitTerminationDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsCircuitTerminationsDeleteParams().WithID(id)

	_, err := api.Circuits.CircuitsCircuitTerminationsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*circuits.CircuitsCircuitTerminationsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
