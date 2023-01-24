package netbox

import (
	"regexp"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxTag() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxTagCreate,
		Read:   resourceNetboxTagRead,
		Update: resourceNetboxTagUpdate,
		Delete: resourceNetboxTagDelete,

		Description: `:meta:subcategory:Extras:From the [official documentation](https://docs.netbox.dev/en/stable/models/extras/tag/):
> Tags are user-defined labels which can be applied to a variety of objects within NetBox. They can be used to establish dimensions of organization beyond the relationships built into NetBox. For example, you might create a tag to identify a particular ownership or condition across several types of objects.
>
> Each tag has a label, color, and a URL-friendly slug. For example, the slug for a tag named "Dunder Mifflin, Inc." would be dunder-mifflin-inc. The slug is generated automatically and makes tags easier to work with as URL parameters. Each tag can also be assigned a description indicating its purpose.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"color_hex": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "9e9e9e",
				ValidateFunc: validation.StringMatch(regexp.MustCompile("^[0-9a-f]{6}$"), "Must be hex color string"),
			},
			"description": {
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

func resourceNetboxTagCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)

	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}

	color := d.Get("color_hex").(string)
	description := d.Get("description").(string)
	params := extras.NewExtrasTagsCreateParams().WithData(
		&models.Tag{
			Name:        &name,
			Slug:        &slug,
			Color:       color,
			Description: description,
		},
	)

	res, err := api.Extras.ExtrasTagsCreate(params, nil)
	if err != nil {
		//return errors.New(getTextFromError(err))
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxTagRead(d, m)
}

func resourceNetboxTagRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasTagsReadParams().WithID(id)

	res, err := api.Extras.ExtrasTagsRead(params, nil)
	if err != nil {
		errorcode := err.(*extras.ExtrasTagsReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", res.GetPayload().Name)
	d.Set("slug", res.GetPayload().Slug)
	d.Set("color_hex", res.GetPayload().Color)
	d.Set("description", res.GetPayload().Description)
	return nil
}

func resourceNetboxTagUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.Tag{}

	name := d.Get("name").(string)
	color := d.Get("color_hex").(string)
	description := d.Get("description").(string)

	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}

	data.Slug = &slug
	data.Name = &name
	data.Color = color
	data.Description = description

	params := extras.NewExtrasTagsUpdateParams().WithID(id).WithData(&data)

	_, err := api.Extras.ExtrasTagsUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxTagRead(d, m)
}

func resourceNetboxTagDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasTagsDeleteParams().WithID(id)

	_, err := api.Extras.ExtrasTagsDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
