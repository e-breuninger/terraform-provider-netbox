//go:build acctest

// Object and multiobject custom fields: configured as related object ids, held as ids in state,
// and carried through updates when set out of band.
package provider_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccNetboxCustomField_objectValues(t *testing.T) {
	testName := testAccGetTestName("cf_object")
	base := strings.ReplaceAll(testName, "-", "_")
	config := func(siteSuffix, description, customFields string) string {
		return fmt.Sprintf(`
resource "netbox_custom_field" "object" {
  name                = "%[1]s_o"
  type                = "object"
  related_object_type = "dcim.site"
  object_types        = ["tenancy.tenant"]
}

resource "netbox_custom_field" "multiobject" {
  name                = "%[1]s_m"
  type                = "multiobject"
  related_object_type = "dcim.site"
  object_types        = ["tenancy.tenant"]
}

resource "netbox_site" "a" {
  name = "%[2]s-a%[3]s"
}

resource "netbox_site" "b" {
  name       = "%[2]s-b"
  depends_on = [netbox_site.a]
}

resource "netbox_tenant" "test" {
  name        = "%[2]s"
  description = "%[4]s"
  %[5]s
  depends_on  = [netbox_custom_field.object, netbox_custom_field.multiobject]
}

data "netbox_tenant" "test" {
  name       = netbox_tenant.test.name
  depends_on = [netbox_tenant.test]
}`, base, testName, siteSuffix, description, customFields)
	}
	// Site b is created after site a, so the configured list is not in id order (the state keeps it,
	// the data source answers the ids sorted).
	configured := `custom_fields = {
    (netbox_custom_field.object.name)      = tostring(netbox_site.a.id)
    (netbox_custom_field.multiobject.name) = jsonencode([netbox_site.b.id, netbox_site.a.id])
  }`
	var tenantID, siteAID, siteBID int64
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// NetBox answers related objects in its own order; the state keeps the configured ids.
				Config: config("", "one", configured),
				Check: resource.ComposeTestCheckFunc(
					stateID("netbox_tenant.test", &tenantID),
					stateID("netbox_site.a", &siteAID),
					stateID("netbox_site.b", &siteBID),
					resource.TestCheckResourceAttrPair("netbox_tenant.test", "custom_fields."+base+"_o", "netbox_site.a", "id"),
					checkCustomField("netbox_tenant.test", base+"_m", func() string { return fmt.Sprintf("[%d,%d]", siteBID, siteAID) }),
					checkCustomField("data.netbox_tenant.test", base+"_o", func() string { return fmt.Sprint(siteAID) }),
					checkCustomField("data.netbox_tenant.test", base+"_m", func() string { return fmt.Sprintf("[%d,%d]", siteAID, siteBID) }),
				),
			},
			{
				// An update sends the object fields again.
				Config: config("", "two", configured),
				Check:  resource.TestCheckResourceAttr("netbox_tenant.test", "description", "two"),
			},
			{
				// Renaming a related object changes nothing on the tenant.
				Config: config("-renamed", "two", configured),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction("netbox_tenant.test", plancheck.ResourceActionNoop)},
				},
			},
			{
				// custom_fields not configured and the object field set out of band: updates still work.
				PreConfig: func() {
					client, err := testAccClient()
					if err != nil {
						t.Fatal(err)
					}
					params := tenancy.NewTenancyTenantsPartialUpdateParams().WithID(tenantID).WithData(&models.WritableTenant{})
					body := map[string]any{"custom_fields": map[string]any{base + "_o": siteBID}}
					if _, err := client.Tenancy.TenancyTenantsPartialUpdate(params, nil, netboxapi.WithBody(body)); err != nil {
						t.Fatalf("out-of-band PATCH: %v", err)
					}
				},
				Config: config("-renamed", "three", ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tenant.test", "description", "three"),
					resource.TestCheckResourceAttrPair("netbox_tenant.test", "custom_fields."+base+"_o", "netbox_site.b", "id"),
				),
			},
		},
	})
}

// checkCustomField asserts a custom field value in state against want, read at check time.
func checkCustomField(name, field string, want func() string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return resource.TestCheckResourceAttr(name, "custom_fields."+field, want())(state)
	}
}
