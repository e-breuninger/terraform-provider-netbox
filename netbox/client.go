package netbox

import (
	"fmt"

	netboxclient "github.com/fbreckle/go-netbox/netbox/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/goware/urlx"
	log "github.com/sirupsen/logrus"
)

// Config struct for the netbox provider
type Config struct {
	APIToken           string
	ServerURL          string
	AllowInsecureHttps bool
}

// Client does the heavy lifting of establishing a base Open API client to Netbox.
func (cfg *Config) Client() (interface{}, error) {

	log.WithFields(log.Fields{
		"server_url": cfg.ServerURL,
	}).Debug("Initializing Netbox client")

	if cfg.APIToken == "" {
		return nil, fmt.Errorf("Missing netbox API key")
	}

	// parse serverUrl
	parsedURL, urlParseError := urlx.Parse(cfg.ServerURL)
	if urlParseError != nil {
		return nil, fmt.Errorf("Error while trying to parse URL: %s", urlParseError)
	}

	desiredRuntimeClientSchemes := []string{parsedURL.Scheme}
	log.WithFields(log.Fields{
		"host":    parsedURL.Host,
		"schemes": desiredRuntimeClientSchemes,
	}).Debug("Initializing Netbox Open API runtime client")

	// build http client
	clientOpts := httptransport.TLSClientOptions{
		InsecureSkipVerify: cfg.AllowInsecureHttps,
	}
	client, _ := httptransport.TLSClient(clientOpts)
	transport := httptransport.NewWithClient(parsedURL.Host, parsedURL.Path + netboxclient.DefaultBasePath, desiredRuntimeClientSchemes, client)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", fmt.Sprintf("Token %v", cfg.APIToken))
	transport.SetLogger(log.StandardLogger())
	netboxClient := netboxclient.New(transport, nil)

	return netboxClient, nil
}
