package client

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	netboxlegacy "github.com/fbreckle/go-netbox/netbox/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/goware/urlx"
	netbox "github.com/netbox-community/go-netbox/v3"
	log "github.com/sirupsen/logrus"
)

// Config struct for the netbox provider
type Config struct {
	APIToken                    string
	ServerURL                   string
	AllowInsecureHTTPS          bool
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

// NewLegacyClient creates a NetBox API client based of github.com/fbreckle/go-netbox.
// It uses the legacy API client, which is based on the OpenAPI 2.0 specification.
// This client is deprecated and will be removed in the future.
func NewLegacyClient(cfg *Config) (*netboxlegacy.NetBoxAPI, error) {
	log.WithFields(log.Fields{
		"server_url": cfg.ServerURL,
	}).Debug("Initializing Netbox legacy client")

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
		InsecureSkipVerify: cfg.AllowInsecureHTTPS,
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

	transport := httptransport.NewWithClient(parsedURL.Host, parsedURL.Path+netboxlegacy.DefaultBasePath, desiredRuntimeClientSchemes, httpClient)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", fmt.Sprintf("Token %v", cfg.APIToken))
	transport.SetLogger(log.StandardLogger())
	netboxClient := netboxlegacy.New(transport, nil)

	return netboxClient, nil
}

// NewClient creates a NetBox API client based on github.com/netbox-community/go-netbox.
// This client is based on the OpenAPI 3.0 specification.
func NewClient(cfg *Config) (*netbox.APIClient, error) {
	log.WithFields(log.Fields{
		"server_url": cfg.ServerURL,
	}).Debug("Initializing Netbox client")
	if cfg.APIToken == "" {
		return nil, fmt.Errorf("missing netbox API key")
	}

	headers := map[string]string{}
	for k, v := range cfg.Headers {
		headers[k] = fmt.Sprintf("%v", v)
	}
	headers["Authorization"] = fmt.Sprintf("Token %v", cfg.APIToken)

	return netbox.NewAPIClient(&netbox.Configuration{
		Servers: []netbox.ServerConfiguration{{
			URL:         cfg.ServerURL,
			Description: "NetBox",
		}},
		HTTPClient: &http.Client{
			Timeout: time.Second * time.Duration(cfg.RequestTimeout),
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: cfg.AllowInsecureHTTPS},
			},
		},
		DefaultHeader: headers,
	}), nil
}

// RoundTrip adds the headers specified in the transport on every request.
func (t customHeaderTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	for key, value := range t.headers {
		r.Header.Add(key, fmt.Sprintf("%v", value))
	}

	resp, err := t.original.RoundTrip(r)
	return resp, err
}
