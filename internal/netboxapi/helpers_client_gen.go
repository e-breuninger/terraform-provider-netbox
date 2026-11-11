// Code generated automatically. DO NOT EDIT.

// The connection to NetBox: building and configuring the go-netbox client, and helpers for
// issuing requests and interpreting NetBox's responses.

// Package netboxapi sits between the generated Terraform resources and go-netbox: per resource the
// generated request/response DTOs and adapters (<resource>_gen.go), and the hand-written code they
// all share (helpers_*_gen.go).
package netboxapi

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/status"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Options tunes the HTTP client built by New.
type Options struct {
	// AllowInsecureHTTPS skips TLS certificate verification.
	AllowInsecureHTTPS bool
	// CACertFile is a PEM file with the CA that signed the server
	// certificate ("" uses the system roots).
	CACertFile string
	// Headers are added to every request.
	Headers map[string]string
	// RequestTimeout bounds every request (0: go-openapi's default).
	RequestTimeout time.Duration
}

// New returns a go-netbox client for the NetBox at baseURL, authenticating
// with the API token: v2 tokens (nbt_ prefix) as "Authorization: Bearer",
// others as "Authorization: Token". baseURL is scheme://host[:port][/path],
// the path being the prefix NetBox is served under (e.g. /netbox behind a
// reverse proxy); the API base path is that prefix plus /api.
func New(baseURL, token string, opts Options) (*client.NetBoxAPI, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("parsing NetBox URL %q: %w", baseURL, err)
	}
	if u.Scheme == "" || u.Host == "" {
		return nil, fmt.Errorf("NetBox URL %q must include a scheme and host (e.g. https://netbox.example.com)", baseURL)
	}
	// A trailing slash on the prefix (strip_trailing_slashes_from_url off)
	// would double the separator before /api.
	basePath := strings.TrimRight(u.Path, "/") + client.DefaultBasePath
	rt, err := httptransport.TLSTransport(httptransport.TLSClientOptions{
		CA:                 opts.CACertFile,
		InsecureSkipVerify: opts.AllowInsecureHTTPS,
	})
	if err != nil {
		return nil, fmt.Errorf("configuring TLS: %w", err)
	}
	if len(opts.Headers) > 0 {
		rt = &headerTransport{next: rt, headers: opts.Headers}
	}
	rt = &retryTransport{next: rt, attempts: transientAttempts, backoff: transientBackoff}
	httpClient := &http.Client{Transport: rt, Timeout: opts.RequestTimeout}
	if opts.RequestTimeout > 0 {
		// go-swagger operations carry their own deadline (DefaultTimeout,
		// 30s) which would otherwise cap longer request_timeouts.
		httptransport.DefaultTimeout = opts.RequestTimeout
	}
	transport := httptransport.NewWithClient(u.Host, basePath, []string{u.Scheme}, httpClient)
	scheme := "Token "
	if strings.HasPrefix(token, "nbt_") {
		scheme = "Bearer "
	}
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", scheme+token)
	return client.New(transport, strfmt.Default), nil
}

// Client is the provider's client data: the go-netbox client plus the
// provider-level settings the resources apply themselves.
type Client struct {
	*client.NetBoxAPI
	// DefaultTags are tag slugs added to every resource (provider
	// default_tags).
	DefaultTags []string
}

// NetBox writes are not free of transient failures: a PATCH to a device
// while a cable of that device is deleted in parallel can deadlock in
// Postgres (NetBox refreshes cached fields on the device's cable
// terminations inside the device transaction), which NetBox surfaces as a
// 500 and which succeeds a moment later. retryTransport repeats such
// requests a few times with a short backoff.
const (
	transientAttempts = 3
	transientBackoff  = 500 * time.Millisecond
)

// retryTransport retries requests that failed transiently: a 5xx response,
// or no response at all (connection reset, EOF). POST is never retried, so
// a create that failed after the server committed is not repeated; every
// other method NetBox exposes is idempotent. A request whose body cannot be
// replayed (no GetBody) is sent once.
type retryTransport struct {
	next     http.RoundTripper
	attempts int
	backoff  time.Duration
}

func (t *retryTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	retryable := r.Method != http.MethodPost && (r.Body == nil || r.Body == http.NoBody || r.GetBody != nil)
	var (
		res *http.Response
		err error
	)
	for attempt := 1; ; attempt++ {
		res, err = t.next.RoundTrip(r)
		if !retryable || attempt >= t.attempts || !transientFailure(res, err) {
			return res, err
		}
		if res != nil {
			io.Copy(io.Discard, res.Body)
			res.Body.Close()
		}
		select {
		case <-r.Context().Done():
			return nil, r.Context().Err()
		case <-time.After(t.backoff * time.Duration(1<<(attempt-1))):
		}
		if r.GetBody != nil {
			body, berr := r.GetBody()
			if berr != nil {
				return nil, berr
			}
			r = r.Clone(r.Context())
			r.Body = body
		}
	}
}

// transientFailure reports whether a round trip is worth repeating: a
// transport error, or a 5xx other than 501.
func transientFailure(res *http.Response, err error) bool {
	if err != nil {
		return true
	}
	return res.StatusCode >= 500 && res.StatusCode != http.StatusNotImplemented
}

// headerTransport adds fixed headers to every request.
type headerTransport struct {
	next    http.RoundTripper
	headers map[string]string
}

func (t *headerTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r = r.Clone(r.Context())
	for name, value := range t.headers {
		r.Header.Set(name, value)
	}
	return t.next.RoundTrip(r)
}

// WithBody is a go-swagger client option that sets the operation's request
// body to body (a request DTO's Payload), keeping the path and query
// parameters the operation's params wrote.
func WithBody(body any) func(*runtime.ClientOperation) {
	return func(op *runtime.ClientOperation) {
		params := op.Params
		op.Params = runtime.ClientRequestWriterFunc(func(req runtime.ClientRequest, reg strfmt.Registry) error {
			if err := params.WriteToRequest(req, reg); err != nil {
				return err
			}
			return req.SetBodyParam(body)
		})
	}
}

// WithQuery is a go-swagger client option that adds values as query parameters after the
// operation's params wrote theirs; one call per name carries all of that name's values.
func WithQuery(values url.Values) func(*runtime.ClientOperation) {
	return func(op *runtime.ClientOperation) {
		params := op.Params
		op.Params = runtime.ClientRequestWriterFunc(func(req runtime.ClientRequest, reg strfmt.Registry) error {
			if err := params.WriteToRequest(req, reg); err != nil {
				return err
			}
			for name, value := range values {
				if err := req.SetQueryParam(name, value...); err != nil {
					return err
				}
			}
			return nil
		})
	}
}

var semanticVersion = regexp.MustCompile(`^v?(\d+\.\d+\.\d+)`)

// Version asks NetBox for its version (GET /api/status/, up to five
// attempts with a growing pause) and returns the reported string (e.g.
// "4.6.5-Docker-3.4.1") and the semantic version extracted from it.
func Version(ctx context.Context, c *client.NetBoxAPI) (reported, semantic string, err error) {
	const attempts = 5
	var res *status.StatusRetrieveOK
	for attempt := 1; attempt <= attempts; attempt++ {
		res, err = c.Status.StatusRetrieveContext(ctx, status.NewStatusRetrieveParams(), nil)
		if err == nil {
			break
		}
		if attempt < attempts {
			time.Sleep(time.Duration(attempt) * time.Second)
		}
	}
	if err != nil {
		return "", "", err
	}
	reported, _ = res.GetPayload()["netbox-version"].(string)
	if reported == "" {
		return "", "", fmt.Errorf("the status response has no netbox-version")
	}
	m := semanticVersion.FindStringSubmatch(reported)
	if len(m) < 2 {
		return reported, "", fmt.Errorf("no semantic version found in version string %q. try using the skip_version_check provider parameter to bypass this error", reported)
	}
	return reported, m[1], nil
}

// VersionErrorDetail explains a failed Version call to the user.
func VersionErrorDetail(serverURL string, err error) string {
	return fmt.Sprintf("The provider could not get a valid JSON response from a GET request to `%s/api/status/`.\n\n"+
		"This usually means one of the following:\n"+
		"  - `server_url` is misconfigured (wrong scheme, host or port, or it already includes a path such as `/api`)\n"+
		"  - Netbox (or a reverse proxy/ingress in front of it) returned a non-JSON response, such as an HTML error page, or a response missing a `Content-Type: application/json` header\n"+
		"  - the Netbox API is temporarily unreachable or returned a server error\n\n"+
		"You can set `skip_version_check = true` on the provider to bypass this check while troubleshooting.\n\n"+
		"Original error: %s", serverURL, err)
}

// IsNotFound reports whether err is a NetBox 404. go-swagger reports
// non-2xx responses as per-operation *XxxDefault errors exposing Code(), or
// as *runtime.APIError.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	var coded interface{ Code() int }
	if errors.As(err, &coded) {
		return coded.Code() == http.StatusNotFound
	}
	var apiErr *runtime.APIError
	if errors.As(err, &apiErr) {
		return apiErr.Code == http.StatusNotFound
	}
	return false
}
