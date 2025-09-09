package client

import (
	"fmt"
	"net/http"
	"time"

	netboxlegacy "github.com/fbreckle/go-netbox/netbox/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/goware/urlx"
	netbox "github.com/netbox-community/go-netbox/v4"
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

func newHTTPClient(cfg *Config) (*http.Client, error) {
	if cfg.APIToken == "" {
		return nil, fmt.Errorf("missing NetBox API key")
	}

	// build http client
	clientOpts := httptransport.TLSClientOptions{
		InsecureSkipVerify: cfg.AllowInsecureHTTPS,
	}

	trans, err := httptransport.TLSTransport(clientOpts)
	if err != nil {
		return nil, err
	}

	trans.(*http.Transport).Proxy = http.ProxyFromEnvironment

	if len(cfg.Headers) > 0 {
		log.WithFields(log.Fields{
			"custom_headers": cfg.Headers,
		}).Debug("Setting custom headers on every request to Netbox")

		trans = customHeaderTransport{
			original: trans,
			headers:  cfg.Headers,
		}
	}

	return &http.Client{
		Transport: trans,
		Timeout:   time.Second * time.Duration(cfg.RequestTimeout),
	}, nil
}

// Client does the heavy lifting of establishing a base Open API client to Netbox.
func NewLegacyClient(cfg *Config) (*netboxlegacy.NetBoxAPI, error) {
	log.WithFields(log.Fields{
		"server_url": cfg.ServerURL,
	}).Debug("Initializing legacy NetBox client")

	httpClient, err := newHTTPClient(cfg)
	if err != nil {
		return nil, err
	}

	parsedURL, urlParseError := urlx.Parse(cfg.ServerURL)
	if urlParseError != nil {
		return nil, fmt.Errorf("error while trying to parse URL: %s", urlParseError)
	}

	desiredRuntimeClientSchemes := []string{parsedURL.Scheme}
	log.WithFields(log.Fields{
		"host":    parsedURL.Host,
		"schemes": desiredRuntimeClientSchemes,
	}).Debug("Initializing Netbox Open API runtime client")

	transport := httptransport.NewWithClient(parsedURL.Host, parsedURL.Path+netboxlegacy.DefaultBasePath, desiredRuntimeClientSchemes, httpClient)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", fmt.Sprintf("Token %v", cfg.APIToken))
	transport.SetLogger(log.StandardLogger())
	netboxClient := netboxlegacy.New(transport, nil)

	return netboxClient, nil
}

func NewClient(cfg *Config) (*netbox.APIClient, error) {
	log.WithFields(log.Fields{
		"server_url": cfg.ServerURL,
	}).Debug("Initializing NetBox client")

	httpClient, err := newHTTPClient(cfg)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Token %v", cfg.APIToken),
	}

	return netbox.NewAPIClient(&netbox.Configuration{
		Servers: []netbox.ServerConfiguration{{
			URL:         cfg.ServerURL,
			Description: "NetBox",
		}},
		HTTPClient:    httpClient,
		DefaultHeader: headers,
	}), nil
}

// customHeaderTransport is a transport that adds the specified headers on
// every request.
type customHeaderTransport struct {
	original http.RoundTripper
	headers  map[string]interface{}
}

// RoundTrip adds the headers specified in the transport on every request.
func (t customHeaderTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	for key, value := range t.headers {
		r.Header.Add(key, fmt.Sprintf("%v", value))
	}

	resp, err := t.original.RoundTrip(r)
	return resp, err
}
