package netbox

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// sweeperNetboxClients is a shared cache of netbox clients
// This prevents client re-initialization for every resource with no benefit
var sweeperNetboxClients map[string]interface{}

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// sharedClientForRegion returns a common provider client configured for the specified region
func sharedClientForRegion(region string) (interface{}, error) {
	if client, ok := sweeperNetboxClients[region]; ok {
		return client, nil
	}

	server := os.Getenv("NETBOX_SERVER")
	apiToken := os.Getenv("NETBOX_API_TOKEN")
	transport := httptransport.New(server, client.DefaultBasePath, []string{"http"})
	authScheme := "Token"
	if strings.HasPrefix(apiToken, "nbt_") {
		authScheme = "Bearer"
	}
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", fmt.Sprintf("%s %v", authScheme, apiToken))
	c := client.New(transport, nil)

	return c, nil
}
