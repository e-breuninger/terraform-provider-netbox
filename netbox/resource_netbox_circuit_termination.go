package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxCircuitTermination() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxCircuitTerminationCreate,
		Read:   resourceNetboxCircuitTerminationRead,
		Update: resourceNetboxCircuitTerminationUpdate,
		Delete: resourceNetboxCircuitTerminationDelete,

		Schema: map[string]*schema.Schema{
			"circuit_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"site_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"port_speed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"upstream_speed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"term_side": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"A", "Z"}, false),
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceNetboxCircuitTerminationCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableCircuitTermination{}

	termside := d.Get("term_side").(string)
	data.TermSide = &termside

	circuitIDValue, ok := d.GetOk("circuit_id")
	if ok {
		data.Circuit = int64ToPtr(int64(circuitIDValue.(int)))
	}

	siteIDValue, ok := d.GetOk("site_id")
	if ok {
		data.Site = int64ToPtr(int64(siteIDValue.(int)))
	}

	portspeedValue, ok := d.GetOk("port_speed")
	if ok {
		data.PortSpeed = int64ToPtr(int64(portspeedValue.(int)))
	}

	upstreamspeedValue, ok := d.GetOk("upstream_speed")
	if ok {
		data.UpstreamSpeed = int64ToPtr(int64(upstreamspeedValue.(int)))
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
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsCircuitTerminationsReadParams().WithID(id)

	res, err := api.Circuits.CircuitsCircuitTerminationsRead(params, nil)

	if err != nil {
		errorcode := err.(*circuits.CircuitsCircuitTerminationsReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("term_side", res.GetPayload().TermSide)

	if res.GetPayload().Circuit != nil {
		d.Set("circuit_id", res.GetPayload().Circuit.ID)
	} else {
		d.Set("circuit_id", nil)
	}

	if res.GetPayload().Site != nil {
		d.Set("site_id", res.GetPayload().Site.ID)
	} else {
		d.Set("site_id", nil)
	}

	if res.GetPayload().PortSpeed != nil {
		d.Set("port_speed", res.GetPayload().PortSpeed)
	} else {
		d.Set("port_speed", nil)
	}

	if res.GetPayload().UpstreamSpeed != nil {
		d.Set("upstream_speed", res.GetPayload().UpstreamSpeed)
	} else {
		d.Set("upstream_speed", nil)
	}

	return nil
}

func resourceNetboxCircuitTerminationUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableCircuitTermination{}

	termside := d.Get("term_side").(string)
	data.TermSide = &termside

	circuitIDValue, ok := d.GetOk("circuit_id")
	if ok {
		data.Circuit = int64ToPtr(int64(circuitIDValue.(int)))
	}

	siteIDValue, ok := d.GetOk("site_id")
	if ok {
		data.Site = int64ToPtr(int64(siteIDValue.(int)))
	}

	portspeedValue, ok := d.GetOk("port_speed")
	if ok {
		data.PortSpeed = int64ToPtr(int64(portspeedValue.(int)))
	}

	upstreamspeedValue, ok := d.GetOk("upstream_speed")
	if ok {
		data.UpstreamSpeed = int64ToPtr(int64(upstreamspeedValue.(int)))
	}
	params := circuits.NewCircuitsCircuitTerminationsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Circuits.CircuitsCircuitTerminationsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxCircuitTerminationRead(d, m)
}

func resourceNetboxCircuitTerminationDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsCircuitTerminationsDeleteParams().WithID(id)

	_, err := api.Circuits.CircuitsCircuitTerminationsDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
