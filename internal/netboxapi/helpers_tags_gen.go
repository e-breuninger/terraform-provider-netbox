// Code generated automatically. DO NOT EDIT.

// Handling of the tags attribute every resource carries, including the provider's default tags.

package netboxapi

import (
	"context"
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
)

// CheckTags verifies that every slug names an existing tag.
func CheckTags(ctx context.Context, c *client.NetBoxAPI, slugs []string) error {
	for _, slug := range slugs {
		slug := slug
		params := extras.NewExtrasTagsListParams().WithSlug([]string{slug})
		res, err := c.Extras.ExtrasTagsListContext(ctx, params, nil)
		if err != nil {
			return fmt.Errorf("looking up tag %q: %w", slug, err)
		}
		if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
			return fmt.Errorf("no tag with slug %q exists", slug)
		}
	}
	return nil
}

// MergeTags is configured plus the defaults it does not already list.
func MergeTags(configured, defaults []string) []string {
	out := append([]string{}, configured...)
	have := map[string]bool{}
	for _, slug := range configured {
		have[slug] = true
	}
	for _, slug := range defaults {
		if !have[slug] {
			out = append(out, slug)
			have[slug] = true
		}
	}
	return out
}

// AddDefaultTags appends the default tag slugs missing from the request
// body's tags.
func AddDefaultTags(body map[string]any, defaults []string) {
	if len(defaults) == 0 {
		return
	}
	refs, _ := body["tags"].([]map[string]string)
	slugs := make([]string, 0, len(refs))
	for _, ref := range refs {
		slugs = append(slugs, ref["slug"])
	}
	body["tags"] = tagRefs(MergeTags(slugs, defaults))
}

// ConfiguredTags is the tags attribute derived from the object's tags:
// every tag except a default one that the configuration did not list
// itself (all is what NetBox returned, configured the prior tags value).
func ConfiguredTags(all, configured, defaults []string) []string {
	isDefault := map[string]bool{}
	for _, slug := range defaults {
		isDefault[slug] = true
	}
	isConfigured := map[string]bool{}
	for _, slug := range configured {
		isConfigured[slug] = true
	}
	out := []string{}
	for _, tag := range all {
		if !isDefault[tag] || isConfigured[tag] {
			out = append(out, tag)
		}
	}
	return out
}

// tagRefs turns tag slugs into the objects NetBox accepts for the tags
// field. A nil input clears the tags.
func tagRefs(slugs []string) []map[string]string {
	out := make([]map[string]string, 0, len(slugs))
	for _, slug := range slugs {
		out = append(out, map[string]string{"slug": slug})
	}
	return out
}
