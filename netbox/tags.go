package netbox

import (
	"fmt"
	"slices"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	tagsKey    = "tags"
	tagsAllKey = "tags_all"
)

var tagsSchema = &schema.Schema{
	Type: schema.TypeSet,
	Elem: &schema.Schema{
		Type: schema.TypeString,
	},
	Optional: true,
	Set:      schema.HashString,
}

var tagsAllSchema = &schema.Schema{
	Type: schema.TypeSet,
	Elem: &schema.Schema{
		Type: schema.TypeString,
	},
	Computed: true,
	Set:      schema.HashString,
}

var tagsSchemaRead = &schema.Schema{
	Type: schema.TypeSet,
	Elem: &schema.Schema{
		Type: schema.TypeString,
	},
	Computed: true,
	Set:      schema.HashString,
}

func getNestedTagListFromResourceDataSet(state *providerState, d interface{}) ([]*models.NestedTag, error) {
	tagList := d.(*schema.Set).List()
	tags := []*models.NestedTag{}
	for _, tag := range tagList {
		nbTag, ok := state.tagCache[tag.(string)]
		if !ok {
			var err error
			nbTag, err = findTag(state.legacyAPI, tag.(string))
			if err != nil {
				return tags, err
			}
		}

		tags = append(tags, nbTag)
	}

	return tags, nil
}

func findTag(client *client.NetBoxAPI, name string) (*models.NestedTag, error) {
	params := extras.NewExtrasTagsListParams()
	params.Name = &name

	limit := int64(2) // We search for a unique tag. Having two hits suffices to know its not unique.
	params.Limit = &limit

	res, err := client.Extras.ExtrasTagsList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("API Error trying to retrieve tag %q from netbox: %w", name, err)
	}

	payload := res.GetPayload()
	switch *payload.Count {
	case int64(0):
		return nil, fmt.Errorf("could not locate referenced tag %q in netbox, no results", name)
	case int64(1):
		return &models.NestedTag{
			Name: payload.Results[0].Name,
			Slug: payload.Results[0].Slug,
		}, nil
	default:
		return nil, fmt.Errorf("could not map tag %q to unique tag in netbox, %d results", name, *payload.Count)
	}
}

func getTagListFromNestedTagList(nestedTags []*models.NestedTag) []string {
	tags := []string{}
	for _, nestedTag := range nestedTags {
		tags = append(tags, *nestedTag.Name)
	}
	return tags
}

func (s *providerState) readTags(d *schema.ResourceData, apiTags []*models.NestedTag) {
	allTags := schema.NewSet(schema.HashString, nil)
	for _, t := range apiTags {
		allTags.Add(*t.Name)
	}
	d.Set(tagsAllKey, allTags.List())

	configTags := make([]string, len(apiTags))
	cf := d.GetRawConfig()
	if cf.IsNull() || !cf.IsKnown() {
		cf = d.GetRawState() // config is missing during refresh
	}
	if !cf.IsNull() && cf.IsKnown() { // there is some config
		c := cf.GetAttr(tagsKey)
		if !c.IsNull() && c.IsKnown() { // tags are configured
			for _, t := range c.AsValueSet().Values() {
				configTags = append(configTags, t.AsString())
			}
		}
	}

	resourceTags := schema.NewSet(schema.HashString, nil)
	// remove default tags (except when configured on the resource)
	for _, tag := range apiTags {
		if !s.defaultTags.Contains(*tag.Name) || slices.Contains(configTags, *tag.Name) {
			resourceTags.Add(*tag.Name)
		}
	}

	d.Set(tagsKey, resourceTags.List())
}
