package netbox

import (
	"os"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// sweeperNetboxClients is a shared cache of netbox clients
// This prevents client re-initialization for every resource with no benefit
var sweeperNetboxClients map[string]interface{}

func TestMain(m *testing.M) {
	// Initialize the client cache
	sweeperNetboxClients = make(map[string]interface{})
	resource.TestMain(m)
}

// sharedClientForRegion returns a common provider client configured for the specified region
func sharedClientForRegion(region string) (interface{}, error) {
	// Initialize map if it's nil (defensive programming)
	if sweeperNetboxClients == nil {
		sweeperNetboxClients = make(map[string]interface{})
	}

	if client, ok := sweeperNetboxClients[region]; ok {
		return client, nil
	}

	server := os.Getenv("NETBOX_SERVER")
	apiToken := os.Getenv("NETBOX_API_TOKEN")
	transport := httptransport.New(server, client.DefaultBasePath, []string{"http"})
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "Token "+apiToken)
	c := client.New(transport, nil)

	// Store the client in the cache for future use
	sweeperNetboxClients[region] = c

	return c, nil
}

func TestSharedClientForRegion(t *testing.T) {
	// Test with environment variables set
	originalServer := os.Getenv("NETBOX_SERVER")
	originalToken := os.Getenv("NETBOX_API_TOKEN")
	defer func() {
		os.Setenv("NETBOX_SERVER", originalServer)
		os.Setenv("NETBOX_API_TOKEN", originalToken)
	}()

	os.Setenv("NETBOX_SERVER", "http://localhost:8080")
	os.Setenv("NETBOX_API_TOKEN", "test-token")

	// Test first call
	client1, err := sharedClientForRegion("test-region")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if client1 == nil {
		t.Fatal("expected client to be non-nil")
	}

	// Test cached call
	client2, err := sharedClientForRegion("test-region")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if client2 == nil {
		t.Fatal("expected client to be non-nil")
	}

	// Should return the same client instance
	if client1 != client2 {
		t.Fatal("expected cached client to be the same instance")
	}
}

func TestSharedClientForRegion_MissingEnvVars(t *testing.T) {
	// Test with missing environment variables
	originalServer := os.Getenv("NETBOX_SERVER")
	originalToken := os.Getenv("NETBOX_API_TOKEN")
	defer func() {
		os.Setenv("NETBOX_SERVER", originalServer)
		os.Setenv("NETBOX_API_TOKEN", originalToken)
	}()

	os.Unsetenv("NETBOX_SERVER")
	os.Unsetenv("NETBOX_API_TOKEN")

	client, err := sharedClientForRegion("test-region")
	if err != nil {
		t.Fatalf("expected no error even with missing env vars, got %v", err)
	}
	if client == nil {
		t.Fatal("expected client to be created even with missing env vars")
	}
}
