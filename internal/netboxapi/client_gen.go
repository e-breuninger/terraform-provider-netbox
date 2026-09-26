// Code generated automatically. DO NOT EDIT.

// Package netboxapi constructs the go-netbox client for this provider and
// holds the generated request/response adapters.
package netboxapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
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

// CheckTags verifies that every slug names an existing tag.
func CheckTags(ctx context.Context, c *client.NetBoxAPI, slugs []string) error {
	for _, slug := range slugs {
		slug := slug
		params := extras.NewExtrasTagsListParams().WithSlug([]string{slug})
		res, err := c.Extras.ExtrasTagsListContext(ctx, params, nil)
		if err != nil {
			return fmt.Errorf("looking up tag %q: %w", slug, err)
		}
		if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
			return fmt.Errorf("no tag with slug %q exists", slug)
		}
	}
	return nil
}

// TypedCustomFields converts the string custom field values in body to the
// types their NetBox definitions require (integer, decimal, boolean,
// multiselect, json), looked up live from the custom field definitions.
// String-typed fields and unknown names pass through unchanged;
// object/multiobject fields are rejected. nil values (keys cleared by
// ClearRemovedCustomFields) stay null.
func TypedCustomFields(ctx context.Context, c *Client, body map[string]any) error {
	var fields map[string]any
	switch cf := body["custom_fields"].(type) {
	case map[string]string:
		fields = make(map[string]any, len(cf))
		for name, value := range cf {
			fields[name] = value
		}
	case map[string]any:
		fields = cf
	default:
		return nil
	}
	body["custom_fields"] = fields
	hasStrings := false
	for _, value := range fields {
		if _, ok := value.(string); ok {
			hasStrings = true
			break
		}
	}
	if !hasStrings {
		return nil
	}
	limit := int64(0)
	res, err := c.Extras.ExtrasCustomFieldsListContext(ctx, extras.NewExtrasCustomFieldsListParams().WithLimit(&limit), nil)
	if err != nil {
		return fmt.Errorf("listing custom field definitions: %w", err)
	}
	defs := map[string]string{}
	if res.Payload != nil {
		for _, definition := range res.Payload.Results {
			if definition != nil && definition.Name != nil && definition.Type != nil && definition.Type.Value != "" {
				defs[*definition.Name] = definition.Type.Value
			}
		}
	}
	for name, value := range fields {
		s, ok := value.(string)
		if !ok {
			continue
		}
		t, ok := defs[name]
		if !ok {
			continue
		}
		tv, err := typedCustomFieldValue(name, t, s)
		if err != nil {
			return err
		}
		fields[name] = tv
	}
	return nil
}

// typedCustomFieldValue converts one string value to the JSON type NetBox
// expects for the field type. The conversions are exact inverses of
// customFieldValues so values round-trip without drift.
func typedCustomFieldValue(name, fieldType, v string) (any, error) {
	switch fieldType {
	case "integer":
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("custom field %q: %q is not an integer", name, v)
		}
		return n, nil
	case "decimal":
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, fmt.Errorf("custom field %q: %q is not a decimal", name, v)
		}
		return f, nil
	case "boolean":
		switch v {
		case "true":
			return true, nil
		case "false":
			return false, nil
		}
		return nil, fmt.Errorf("custom field %q: %q is not a boolean (use \"true\" or \"false\")", name, v)
	case "multiselect":
		var xs []any
		if err := json.Unmarshal([]byte(v), &xs); err != nil {
			return nil, fmt.Errorf("custom field %q: %q is not a JSON array", name, v)
		}
		return xs, nil
	case "json":
		var x any
		if err := json.Unmarshal([]byte(v), &x); err != nil {
			return v, nil // bare text is sent as a JSON string
		}
		return x, nil
	case "object", "multiobject":
		return nil, fmt.Errorf("custom field %q: %s custom fields are not supported", name, fieldType)
	default:
		return v, nil
	}
}

// MergeTags is configured plus the defaults it does not already list.
func MergeTags(configured, defaults []string) []string {
	out := append([]string{}, configured...)
	have := map[string]bool{}
	for _, slug := range configured {
		have[slug] = true
	}
	for _, slug := range defaults {
		if !have[slug] {
			out = append(out, slug)
			have[slug] = true
		}
	}
	return out
}

// AddDefaultTags appends the default tag slugs missing from the request
// body's tags.
func AddDefaultTags(body map[string]any, defaults []string) {
	if len(defaults) == 0 {
		return
	}
	refs, _ := body["tags"].([]map[string]string)
	slugs := make([]string, 0, len(refs))
	for _, ref := range refs {
		slugs = append(slugs, ref["slug"])
	}
	body["tags"] = tagRefs(MergeTags(slugs, defaults))
}

// ConfiguredTags is the tags attribute derived from the object's tags:
// every tag except a default one that the configuration did not list
// itself (all is what NetBox returned, configured the prior tags value).
func ConfiguredTags(all, configured, defaults []string) []string {
	isDefault := map[string]bool{}
	for _, slug := range defaults {
		isDefault[slug] = true
	}
	isConfigured := map[string]bool{}
	for _, slug := range configured {
		isConfigured[slug] = true
	}
	out := []string{}
	for _, tag := range all {
		if !isDefault[tag] || isConfigured[tag] {
			out = append(out, tag)
		}
	}
	return out
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

var (
	slugSpecial    = regexp.MustCompile(`[^\w\s-]`)
	slugWhitespace = regexp.MustCompile(`[\s-]+`)
)

// Slugify derives a NetBox slug from a name: special characters are stripped, runs of
// whitespace and dashes collapse to a single dash, and the result is trimmed and lowercased.
func Slugify(name string) string {
	s := slugSpecial.ReplaceAllString(name, "")
	s = slugWhitespace.ReplaceAllString(s, "-")
	return strings.ToLower(strings.Trim(s, "-"))
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

// reencode copies src into dst through JSON. It is how nested go-netbox
// values — typed nested models, slices of them, or the untyped interface{}
// NetBox uses for polymorphic fields such as assigned_object — become the
// generated nested response DTOs, matched by JSON field name.
func reencode(src, dst any) error {
	raw, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, dst)
}

// genericObjects renders one side of a cable for the wire: one
// {object_type, object_id} per id, all of the side's type.
func genericObjects(objectType *string, ids []int64) []map[string]any {
	out := make([]map[string]any, 0, len(ids))
	for _, id := range ids {
		var t any
		if objectType != nil {
			t = *objectType
		}
		out = append(out, map[string]any{"object_type": t, "object_id": id})
	}
	return out
}

// terminationSideOf folds NetBox's generic object list of one cable side back
// into its type and ids (through JSON, as the nested adapters do): the type
// is the first object's, nil when the side is empty.
func terminationSideOf(src any) (objectType *string, ids []int64) {
	var objs []map[string]any
	if err := reencode(src, &objs); err != nil {
		return nil, nil
	}
	ids = []int64{}
	for _, obj := range objs {
		if t, ok := obj["object_type"].(string); ok && objectType == nil {
			objectType = &t
		}
		if id, ok := obj["object_id"].(float64); ok {
			ids = append(ids, int64(id))
		}
	}
	return objectType, ids
}

// tagRefs turns tag slugs into the objects NetBox accepts for the tags
// field. A nil input clears the tags.
func tagRefs(slugs []string) []map[string]string {
	out := make([]map[string]string, 0, len(slugs))
	for _, slug := range slugs {
		out = append(out, map[string]string{"slug": slug})
	}
	return out
}

// customFieldValues converts NetBox's custom_fields object (arbitrary JSON
// values, unset fields null) to string values: strings as-is, numbers in
// their shortest form, booleans as true/false, anything else JSON-encoded.
// Null entries are dropped, so an object without values is an empty map
// (custom_fields = {} round-trips); only a missing object yields nil.
func customFieldValues(v any) map[string]string {
	m, ok := v.(map[string]any)
	if !ok {
		return nil
	}
	out := map[string]string{}
	for name, value := range m {
		switch x := value.(type) {
		case nil:
			continue
		case string:
			out[name] = x
		case float64:
			out[name] = strconv.FormatFloat(x, 'f', -1, 64)
		case bool:
			out[name] = strconv.FormatBool(x)
		default:
			raw, err := json.Marshal(x)
			if err != nil {
				continue
			}
			out[name] = string(raw)
		}
	}
	return out
}

// ClearRemovedCustomFields adds a null entry to body's custom_fields for
// every key of prior (the previous state) the plan no longer sets, so
// NetBox clears it instead of keeping the old value.
func ClearRemovedCustomFields(body map[string]any, prior map[string]string) {
	if len(prior) == 0 {
		return
	}
	cf, _ := body["custom_fields"].(map[string]string)
	merged := map[string]any{}
	for name, value := range cf {
		merged[name] = value
	}
	for name := range prior {
		if _, ok := cf[name]; !ok {
			merged[name] = nil
		}
	}
	body["custom_fields"] = merged
}

// choiceValue extracts the value of a go-netbox choice object, whose
// Value field is *T on most models and T on some; the zero value (an
// unset choice) becomes nil.
func choiceValue[T comparable](v any) *T {
	var zero T
	switch x := v.(type) {
	case *T:
		if x == nil || *x == zero {
			return nil
		}
		return x
	case T:
		if x == zero {
			return nil
		}
		return &x
	}
	return nil
}

// jsonText renders an untyped JSON value as JSON text; nil stays nil.
func jsonText(v any) *string {
	if v == nil {
		return nil
	}
	raw, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	s := string(raw)
	return &s
}

// parseJSONText decodes JSON text for an untyped field; text that is not
// valid JSON is sent as a JSON string.
func parseJSONText(s string) any {
	var v any
	if err := json.Unmarshal([]byte(s), &v); err != nil {
		return s
	}
	return v
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
