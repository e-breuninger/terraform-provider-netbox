package netbox

import (
	"context"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxVirtualChassis() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxVirtualChassisCreate,
		ReadContext:   resourceNetboxVirtualChassisRead,
		UpdateContext: resourceNetboxVirtualChassisUpdate,
		DeleteContext: resourceNetboxVirtualChassisDelete,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/features/devices-cabling/#virtual-chassis):

		> Sometimes it is necessary to model a set of physical devices as sharing a single management plane. Perhaps the most common example of such a scenario is stackable switches. These can be modeled as virtual chassis in NetBox, with one device acting as the chassis master and the rest as members. All components of member devices will appear on the master.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			tagsKey:         tagsSchema,
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxVirtualChassisCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	name := d.Get("name").(string)

	data := models.WritableVirtualChassis{
		Name: &name,
	}

	domainValue, ok := d.GetOk("domain")
	if ok {
		domain := domainValue.(string)
		data.Domain = domain
	}

	descriptionValue, ok := d.GetOk("description")
	if ok {
		description := descriptionValue.(string)
		data.Description = description
	}

	commentsValue, ok := d.GetOk("comments")
	if ok {
		comments := commentsValue.(string)
		data.Comments = comments
	}

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return diag.FromErr(err)
	}

	params := dcim.NewDcimVirtualChassisCreateParams().WithData(&data)

	res, err := api.Dcim.DcimVirtualChassisCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVirtualChassisRead(ctx, d, m)
}

func resourceNetboxVirtualChassisRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := dcim.NewDcimVirtualChassisReadParams().WithID(id)

	res, err := api.Dcim.DcimVirtualChassisRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimVirtualChassisReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	virtualChassis := res.GetPayload()

	d.Set("name", virtualChassis.Name)
	d.Set("domain", virtualChassis.Domain)
	d.Set("description", virtualChassis.Description)
	d.Set("comments", virtualChassis.Comments)

	cf := flattenCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}

	api.readTags(d, virtualChassis.Tags)
	return nil
}

func resourceNetboxVirtualChassisUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableVirtualChassis{}

	name := d.Get("name").(string)
	data.Name = &name

	domainValue, ok := d.GetOk("domain")
	if ok {
		domain := domainValue.(string)
		data.Domain = domain
	}

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChanges("comments") {
		// check if comment is set
		if commentsValue, ok := d.GetOk("comments"); ok {
			data.Comments = commentsValue.(string)
		} else {
			data.Comments = " "
		}
	}
	if d.HasChanges("description") {
		// check if description is set
		if descriptionValue, ok := d.GetOk("description"); ok {
			data.Description = descriptionValue.(string)
		} else {
			data.Description = " "
		}
	}

	params := dcim.NewDcimVirtualChassisUpdateParams().WithID(id).WithData(&data)

	_, err = api.Dcim.DcimVirtualChassisUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceNetboxVirtualChassisRead(ctx, d, m)
}

func resourceNetboxVirtualChassisDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimVirtualChassisDeleteParams().WithID(id)

	_, err := api.Dcim.DcimVirtualChassisDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimVirtualChassisDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	return nil
}

func virtualChassisUpdateMaster(api *providerState, id int64, master *int64) error {
	// Need to read the virtual chassis because we cannot do a partial update
	// because setting `master` to nil would omit it entirely, so we need to
	// do a PUT request instead of PATCH
	vcRes, err := api.Dcim.DcimVirtualChassisRead(dcim.NewDcimVirtualChassisReadParams().WithID(id), nil)
	if err != nil {
		return err
	}
	vcData := vcRes.GetPayload()
	vcData.Master = nil
	newTags := make([]*models.NestedTag, 0)
	for _, Tag := range vcData.Tags {
		Tag.Display = ""
		Tag.ID = 0
		Tag.URL = ""
		newTags = append(newTags, Tag)
	}

	// Need to manually copy data because there is no automatic method to convert
	// from VirtualChassis to WritableVirtualChassis
	vcUpdateData := models.WritableVirtualChassis{
		ID:          vcData.ID,
		Description: vcData.Description,
		Domain:      vcData.Domain,
		Name:        vcData.Name,
		Comments:    vcData.Comments,
		Tags:        newTags,
		Master:      master,
	}

	_, err = api.Dcim.DcimVirtualChassisUpdate(dcim.NewDcimVirtualChassisUpdateParams().WithID(id).WithData(&vcUpdateData), nil)
	return err
}
