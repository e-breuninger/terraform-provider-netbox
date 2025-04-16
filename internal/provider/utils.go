package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/netbox-community/go-netbox/v4"
	"regexp"
	"strings"
)

func testClient(client *netbox.APIClient) diag.Diagnostics {
	var diags = diag.Diagnostics{}
	//Validate that the APIClient exist.
	if client == nil {
		diags.AddError(
			"Create: Unconfigured API Client",
			"Expected configured API Client. Please report this issue to the provider developers.",
		)
		return nil
	}
	return diags
}

func getSlug(name string) string {
	var result string
	// \w = word characters (== [0-9A-Za-z_])
	// \s = whitespace (== [\t\n\f\r ])
	matchSpecial, _ := regexp.Compile(`[^\w\s-]`)
	matchMultiWhitespacesAndDashes, _ := regexp.Compile(`[\s-]+`)
	// Special chars are stripped
	result = matchSpecial.ReplaceAllString(name, "")
	// Blocks of multiple whitespaces and dashes will be replaced by a single dash
	result = matchMultiWhitespacesAndDashes.ReplaceAllString(result, "-")
	result = strings.Trim(result, "-")
	return strings.ToLower(result)
}
