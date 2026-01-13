package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxClustersDataSource_basic(t *testing.T) {
	testSlug := "clusters_ds_basic"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxClustersDataSourceDependencies(testName)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + testAccNetboxClustersDataSourceFilterName(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_clusters.test", "clusters.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_clusters.test", "clusters.0.name", testName+"_0"),
					resource.TestCheckResourceAttr("data.netbox_clusters.test", "clusters.0.comments", "thisisacomment"),
					resource.TestCheckResourceAttr("data.netbox_clusters.test", "clusters.0.description", "thisisadescription"),
					resource.TestCheckResourceAttrPair("data.netbox_clusters.test", "clusters.0.cluster_type_id", "netbox_cluster_type.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_clusters.test", "clusters.0.cluster_group_id", "netbox_cluster_group.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_clusters.test", "clusters.0.tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_clusters.test", "clusters.0.site_id", "netbox_site.test", "id"),
				),
			},
			{
				Config: dependencies + testAccNetboxClustersDataSourceFilterClusterTypeID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_clusters.test", "clusters.#", "4"),
					resource.TestCheckResourceAttrPair("data.netbox_clusters.test", "clusters.0.cluster_type_id", "netbox_cluster_type.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_clusters.test", "clusters.0.name", "netbox_cluster.test0", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_clusters.test", "clusters.1.name", "netbox_cluster.test1", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_clusters.test", "clusters.2.name", "netbox_cluster.test2", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_clusters.test", "clusters.3.name", "netbox_cluster.test3", "name"),
				),
			},
			{
				Config: dependencies + testAccNetboxClustersDataSourceFilterClusterGroupID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_clusters.test", "clusters.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_clusters.test", "clusters.0.cluster_group_id", "netbox_cluster_group.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_clusters.test", "clusters.0.name", "netbox_cluster.test0", "name"),
				),
			},
			{
				Config: dependencies + testAccNetboxClustersDataSourceFilterSiteID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_clusters.test", "clusters.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_clusters.test", "clusters.0.site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_clusters.test", "clusters.0.name", "netbox_cluster.test0", "name"),
				),
			},
			{
				Config: dependencies + testAccNetboxClustersDataSourceNameRegex,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_clusters.test", "clusters.#", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_clusters.test", "clusters.0.name", "netbox_cluster.test2", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_clusters.test", "clusters.1.name", "netbox_cluster.test3", "name"),
				),
			},
			{
				Config: dependencies + testAccNetboxClustersDataSourceLimit,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_clusters.test", "clusters.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_clusters.test", "clusters.0.cluster_type_id", "netbox_cluster_type.test", "id"),
				),
			},
		},
	})
}

func TestAccNetboxClustersDataSource_tags(t *testing.T) {
	testSlug := "clusters_ds_tags"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxClustersDataSourceDependenciesWithTags(testName)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + testAccNetboxClustersDataSourceTagA(testName) + testAccNetboxClustersDataSourceTagB(testName) + testAccNetboxClustersDataSourceTagAB(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_clusters.tag-a", "clusters.#", "2"),
					resource.TestCheckResourceAttr("data.netbox_clusters.tag-b", "clusters.#", "2"),
					resource.TestCheckResourceAttr("data.netbox_clusters.tag-ab", "clusters.#", "1"),
				),
			},
		},
	})
}

func testAccNetboxClustersDataSourceDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name   = "%[1]s"
  status = "active"
}

resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

resource "netbox_cluster_group" "test" {
  name = "%[1]s"
}

resource "netbox_cluster" "test0" {
  name             = "%[1]s_0"
  cluster_type_id  = netbox_cluster_type.test.id
  cluster_group_id = netbox_cluster_group.test.id
  site_id          = netbox_site.test.id
  tenant_id        = netbox_tenant.test.id
  comments         = "thisisacomment"
  description      = "thisisadescription"
}

resource "netbox_cluster" "test1" {
  name            = "%[1]s_1"
  cluster_type_id = netbox_cluster_type.test.id
}

resource "netbox_cluster" "test2" {
  name            = "%[1]s_2_regex"
  cluster_type_id = netbox_cluster_type.test.id
}

resource "netbox_cluster" "test3" {
  name            = "%[1]s_3_regex"
  cluster_type_id = netbox_cluster_type.test.id
}
`, testName)
}

func testAccNetboxClustersDataSourceDependenciesWithTags(testName string) string {
	return fmt.Sprintf(`
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

resource "netbox_tag" "servicea" {
  name = "%[1]s_service-a"
}

resource "netbox_tag" "serviceb" {
  name = "%[1]s_service-b"
}

resource "netbox_cluster" "test0" {
  name            = "%[1]s_0"
  cluster_type_id = netbox_cluster_type.test.id
  tags = [
    netbox_tag.servicea.name,
    netbox_tag.serviceb.name,
  ]
}

resource "netbox_cluster" "test1" {
  name            = "%[1]s_1"
  cluster_type_id = netbox_cluster_type.test.id
  tags = [
    netbox_tag.servicea.name,
  ]
}

resource "netbox_cluster" "test2" {
  name            = "%[1]s_2"
  cluster_type_id = netbox_cluster_type.test.id
  tags = [
    netbox_tag.serviceb.name,
  ]
}
`, testName)
}

func testAccNetboxClustersDataSourceFilterName(testName string) string {
	return fmt.Sprintf(`
data "netbox_clusters" "test" {
  filter {
    name  = "name"
    value = "%[1]s_0"
  }
}`, testName)
}

const testAccNetboxClustersDataSourceFilterClusterTypeID = `
data "netbox_clusters" "test" {
  filter {
    name  = "cluster_type_id"
    value = netbox_cluster_type.test.id
  }
}`

const testAccNetboxClustersDataSourceFilterClusterGroupID = `
data "netbox_clusters" "test" {
  filter {
    name  = "cluster_group_id"
    value = netbox_cluster_group.test.id
  }
}`

const testAccNetboxClustersDataSourceFilterSiteID = `
data "netbox_clusters" "test" {
  filter {
    name  = "site_id"
    value = netbox_site.test.id
  }
}`

const testAccNetboxClustersDataSourceNameRegex = `
data "netbox_clusters" "test" {
  name_regex = "test.*_regex"
}`

const testAccNetboxClustersDataSourceLimit = `
data "netbox_clusters" "test" {
  limit = 1
  filter {
    name  = "cluster_type_id"
    value = netbox_cluster_type.test.id
  }
}`

func testAccNetboxClustersDataSourceTagA(testName string) string {
	return fmt.Sprintf(`
data "netbox_clusters" "tag-a" {
  filter {
    name  = "tag"
    value = "%[1]s_service-a"
  }
}`, testName)
}

func testAccNetboxClustersDataSourceTagB(testName string) string {
	return fmt.Sprintf(`
data "netbox_clusters" "tag-b" {
  filter {
    name  = "tag"
    value = "%[1]s_service-b"
  }
}`, testName)
}

func testAccNetboxClustersDataSourceTagAB(testName string) string {
	return fmt.Sprintf(`
data "netbox_clusters" "tag-ab" {
  filter {
    name  = "tag"
    value = "%[1]s_service-a"
  }
  filter {
    name  = "tag"
    value = "%[1]s_service-b"
  }
}`, testName)
}
