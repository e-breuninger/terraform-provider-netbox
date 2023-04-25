package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxRackReservation() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxRackReservationCreate,
		Read:   resourceNetboxRackReservationRead,
		Update: resourceNetboxRackReservationUpdate,
		Delete: resourceNetboxRackReservationDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/rackreservation/):

> Users can reserve specific units within a rack for future use. An arbitrary set of units within a rack can be associated with a single reservation, but reservations cannot span multiple racks. A description is required for each reservation, reservations may optionally be associated with a specific tenant.`,

		Schema: map[string]*schema.Schema{
			"rack_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"units": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Required: true,
				Set:      schema.HashInt,
			},
			"user_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxRackReservationCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := dcim.NewDcimRackReservationsCreateParams().WithData(
		&models.WritableRackReservation{
			Rack:        getOptionalInt(d, "rack_id"),
			Units:       toInt64PtrList(d.Get("units")),
			User:        getOptionalInt(d, "user_id"),
			Description: strToPtr(getOptionalStr(d, "description", false)),
			Tenant:      getOptionalInt(d, "tenant_id"),
			Comments:    getOptionalStr(d, "comments", false),
			Tags:        tags,
		},
	)

	res, err := api.Dcim.DcimRackReservationsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxRackReservationRead(d, m)
}

func resourceNetboxRackReservationRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimRackReservationsReadParams().WithID(id)

	res, err := api.Dcim.DcimRackReservationsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimRackReservationsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	rackRes := res.GetPayload()

	if rackRes.Rack != nil {
		d.Set("rack_id", rackRes.Rack.ID)
	}

	units := []int{}
	for _, unit := range rackRes.Units {
		units = append(units, int(*unit))
	}
	d.Set("units", units)

	if rackRes.User != nil {
		d.Set("user_id", rackRes.User.ID)
	}

	d.Set("description", rackRes.Description)

	if rackRes.Tenant != nil {
		d.Set("tenant_id", rackRes.Tenant.ID)
	}

	d.Set("comments", rackRes.Comments)

	d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))
	return nil
}

func resourceNetboxRackReservationUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	data := models.WritableRackReservation{
		Rack:        getOptionalInt(d, "rack_id"),
		Units:       toInt64PtrList(d.Get("units")),
		User:        getOptionalInt(d, "user_id"),
		Description: strToPtr(getOptionalStr(d, "description", false)),
		Tenant:      getOptionalInt(d, "tenant_id"),
		Comments:    getOptionalStr(d, "comments", false),
		Tags:        tags,
	}

	params := dcim.NewDcimRackReservationsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimRackReservationsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxRackReservationRead(d, m)
}

func resourceNetboxRackReservationDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimRackReservationsDeleteParams().WithID(id)

	_, err := api.Dcim.DcimRackReservationsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimRackReservationsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
