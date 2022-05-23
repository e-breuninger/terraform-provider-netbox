package netbox

import (
	"context"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxDevice() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxDeviceCreate,
		ReadContext:   resourceNetboxDeviceRead,
		UpdateContext: resourceNetboxDeviceUpdate,
		DeleteContext: resourceNetboxDeviceDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"device_type_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"role_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Set:      schema.HashString,
			},
			"primary_ipv4": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceNetboxDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)

	data := models.WritableDeviceWithConfigContext{
		Name: &name,
	}

	typeIDValue, ok := d.GetOk("device_type_id")
	if ok {
		typeID := int64(typeIDValue.(int))
		data.DeviceType = &typeID
	}

	comments := d.Get("comments").(string)
	data.Comments = comments

	serial := d.Get("serial").(string)
	data.Serial = serial

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		tenantID := int64(tenantIDValue.(int))
		data.Tenant = &tenantID
	}

	roleIDValue, ok := d.GetOk("role_id")
	if ok {
		roleID := int64(roleIDValue.(int))
		data.DeviceRole = &roleID
	}

	siteIDValue, ok := d.GetOk("site_id")
	if ok {
		siteID := int64(siteIDValue.(int))
		data.Site = &siteID
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	params := dcim.NewDcimDevicesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimDevicesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxDeviceRead(ctx, d, m)
}

func resourceNetboxDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := dcim.NewDcimDevicesReadParams().WithID(id)

	res, err := api.Dcim.DcimDevicesRead(params, nil)
	if err != nil {
		errorcode := err.(*dcim.DcimDevicesReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.Set("name", res.GetPayload().Name)

	if res.GetPayload().DeviceType != nil {
		d.Set("device_type_id", res.GetPayload().DeviceType.ID)
	}

	if res.GetPayload().PrimaryIp4 != nil {
		d.Set("primary_ipv4", res.GetPayload().PrimaryIp4.ID)
	} else {
		d.Set("primary_ipv4", nil)
	}

	if res.GetPayload().Tenant != nil {
		d.Set("tenant_id", res.GetPayload().Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	if res.GetPayload().DeviceRole != nil {
		d.Set("role_id", res.GetPayload().DeviceRole.ID)
	} else {
		d.Set("role_id", nil)
	}

	if res.GetPayload().Site != nil {
		d.Set("site_id", res.GetPayload().Site.ID)
	} else {
		d.Set("site_id", nil)
	}

	d.Set("comments", res.GetPayload().Comments)

	d.Set("serial", res.GetPayload().Serial)

	d.Set("tags", getTagListFromNestedTagList(res.GetPayload().Tags))
	return diags
}

func resourceNetboxDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableDeviceWithConfigContext{}

	name := d.Get("name").(string)
	data.Name = &name

	typeIDValue, ok := d.GetOk("device_type_id")
	if ok {
		tenantID := int64(typeIDValue.(int))
		data.Tenant = &tenantID
	}

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		tenantID := int64(tenantIDValue.(int))
		data.Tenant = &tenantID
	}

	roleIDValue, ok := d.GetOk("role_id")
	if ok {
		roleID := int64(roleIDValue.(int))
		data.DeviceRole = &roleID
	}

	siteIDValue, ok := d.GetOk("site_id")
	if ok {
		siteID := int64(siteIDValue.(int))
		data.Site = &siteID
	}

	commentsValue, ok := d.GetOk("comments")
	if ok {
		comments := commentsValue.(string)
		data.Comments = comments
	} else {
		comments := " "
		data.Comments = comments
	}

	primaryIPValue, ok := d.GetOk("primary_ipv4")
	if ok {
		primaryIP := int64(primaryIPValue.(int))
		data.PrimaryIp4 = &primaryIP
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	if d.HasChanges("comments") {
		// check if comment is set
		commentsValue, ok := d.GetOk("comments")
		comments := ""
		if !ok {
			// Setting an space string deletes the comment
			comments = " "
		} else {
			comments = commentsValue.(string)
		}
		data.Comments = comments
	}

	if d.HasChanges("serial") {
		// check if serial is set
		serialValue, ok := d.GetOk("serial")
		serial := ""
		if !ok {
			// Setting an space string deletes the serial
			serial = " "
		} else {
			serial = serialValue.(string)
		}
		data.Serial = serial
	}

	params := dcim.NewDcimDevicesUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimDevicesUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceNetboxDeviceRead(ctx, d, m)
}

func resourceNetboxDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimDevicesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimDevicesDelete(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
