package netbox

import (
	"fmt"
	"net/http"
	"time"

	netboxclient "github.com/fbreckle/go-netbox/netbox/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/goware/urlx"
	log "github.com/sirupsen/logrus"
)

// Config struct for the netbox provider
type Config struct {
	APIToken                    string
	ServerURL                   string
	AllowInsecureHttps          bool
	Headers                     map[string]interface{}
	RequestTimeout              int
	StripTrailingSlashesFromURL bool
}

// customHeaderTransport is a transport that adds the specified headers on
// every request.
type customHeaderTransport struct {
	original http.RoundTripper
	headers  map[string]interface{}
}

// Client does the heavy lifting of establishing a base Open API client to Netbox.
func (cfg *Config) Client() (interface{}, error) {

	log.WithFields(log.Fields{
		"server_url": cfg.ServerURL,
	}).Debug("Initializing Netbox client")

	if cfg.APIToken == "" {
		return nil, fmt.Errorf("missing netbox API key")
	}

	// parse serverUrl
	parsedURL, urlParseError := urlx.Parse(cfg.ServerURL)
	if urlParseError != nil {
		return nil, fmt.Errorf("error while trying to parse URL: %s", urlParseError)
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

	trans, err := httptransport.TLSTransport(clientOpts)
	if err != nil {
		return nil, err
	}

	if cfg.Headers != nil && len(cfg.Headers) > 0 {
		log.WithFields(log.Fields{
			"custom_headers": cfg.Headers,
		}).Debug("Setting custom headers on every request to Netbox")

		trans = customHeaderTransport{
			original: trans,
			headers:  cfg.Headers,
		}
	}

	httpClient := &http.Client{
		Transport: trans,
		Timeout:   time.Second * time.Duration(cfg.RequestTimeout),
	}

	transport := httptransport.NewWithClient(parsedURL.Host, parsedURL.Path+netboxclient.DefaultBasePath, desiredRuntimeClientSchemes, httpClient)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", fmt.Sprintf("Token %v", cfg.APIToken))
	transport.SetLogger(log.StandardLogger())
	netboxClient := netboxclient.New(transport, nil)

	return netboxClient, nil
}

// RoundTrip adds the headers specified in the transport on every request.
func (t customHeaderTransport) RoundTrip(r *http.Request) (*http.Response, error) {

	for key, value := range t.headers {
		r.Header.Add(key, fmt.Sprintf("%v", value))
	}

	resp, err := t.original.RoundTrip(r)
	return resp, err
}
