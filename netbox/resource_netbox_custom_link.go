package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxCustomLinkButtonClassOptions = []string{"outline-dark", "blue", "indigo", "purple", "pink", "red", "orange", "yellow", "green", "teal", "cyan", "gray", "black", "white", "ghost-dark"}

func resourceNetboxCustomLink() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxCustomLinkCreate,
		Read:   resourceNetboxCustomLinkRead,
		Update: resourceNetboxCustomLinkUpdate,
		Delete: resourceNetboxCustomLinkDelete,

		Description: `:meta:subcategory:Extras:From the [official documentation](https://docs.netbox.dev/en/stable/customization/custom-links/):

> Custom links allow users to display arbitrary hyperlinks to external content within NetBox object views. These are helpful for cross-referencing related records in systems outside of NetBox. For example, you might want to link out to a customer-tracking system that supplies data on the client to which a device belongs. Custom links are shown as buttons at the top of the object view in NetBox.

> Custom links may be used with any object type in NetBox. Each is associated with a particular content type and displayed only for objects of that type. Note that custom links are also displayed for all instances of that type of object, so it may be desirable to make use of Jinja2 templating within the link to control link text or URL, or even whether the link is rendered at all.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"content_types": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"link_text": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Jinja2 template code for link text",
			},
			"link_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Jinja2 template code for link URL",
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"group_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 50),
				Description:  "Links with the same group will appear as a dropdown menu",
			},
			"button_class": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxCustomLinkButtonClassOptions, false),
				Description:  "The class of the first link in a group will be used for the dropdown button. " + buildValidValueDescription(resourceNetboxCustomLinkButtonClassOptions),
			},
			"new_window": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Force link to open in a new window",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxCustomLinkCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := &models.CustomLink{}

	name := d.Get("name").(string)
	data.Name = &name
	data.ObjectTypes = toStringList(d.Get("content_types"))
	linkText := d.Get("link_text").(string)
	data.LinkText = &linkText
	linkURL := d.Get("link_url").(string)
	data.LinkURL = &linkURL
	data.Enabled = d.Get("enabled").(bool)
	data.NewWindow = d.Get("new_window").(bool)
	data.GroupName = getOptionalStr(d, "group_name", false)
	data.ButtonClass = d.Get("button_class").(string)

	weight := int64(d.Get("weight").(int))
	data.Weight = &weight

	params := extras.NewExtrasCustomLinksCreateParams().WithData(data)

	res, err := api.Extras.ExtrasCustomLinksCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxCustomLinkRead(d, m)
}

func resourceNetboxCustomLinkRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasCustomLinksReadParams().WithID(id)

	res, err := api.Extras.ExtrasCustomLinksRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*extras.ExtrasCustomLinksReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	customLink := res.GetPayload()
	d.Set("name", customLink.Name)
	d.Set("content_types", customLink.ObjectTypes)
	d.Set("link_text", customLink.LinkText)
	d.Set("link_url", customLink.LinkURL)
	d.Set("enabled", customLink.Enabled)
	d.Set("new_window", customLink.NewWindow)
	d.Set("group_name", customLink.GroupName)
	d.Set("button_class", customLink.ButtonClass)
	d.Set("weight", customLink.Weight)

	return nil
}

func resourceNetboxCustomLinkUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.CustomLink{}

	name := d.Get("name").(string)
	data.Name = &name
	data.ObjectTypes = toStringList(d.Get("content_types"))
	linkText := d.Get("link_text").(string)
	data.LinkText = &linkText
	linkURL := d.Get("link_url").(string)
	data.LinkURL = &linkURL
	data.Enabled = d.Get("enabled").(bool)
	data.NewWindow = d.Get("new_window").(bool)
	data.GroupName = getOptionalStr(d, "group_name", true)
	data.ButtonClass = d.Get("button_class").(string)

	weight := int64(d.Get("weight").(int))
	data.Weight = &weight

	params := extras.NewExtrasCustomLinksUpdateParams().WithID(id).WithData(&data)

	_, err := api.Extras.ExtrasCustomLinksUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxCustomLinkRead(d, m)
}

func resourceNetboxCustomLinkDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasCustomLinksDeleteParams().WithID(id)

	_, err := api.Extras.ExtrasCustomLinksDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*extras.ExtrasCustomLinksDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
