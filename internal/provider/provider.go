package provider

import (
	"context"
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/goware/urlx"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netbox-community/go-netbox/v4"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = (*netboxProvider)(nil)

func New() func() provider.Provider {
	return func() provider.Provider {
		return &netboxProvider{}
	}
}

type netboxProvider struct {
	client *netbox.APIClient
}
type netboxProviderModel struct {
	ApiToken                    types.String `tfsdk:"api_token"`
	ServerUrl                   types.String `tfsdk:"server_url"`
	SkipVersionCheck            types.Bool   `tfsdk:"skip_version_check"`
	AllowInsecureHttps          types.Bool   `tfsdk:"allow_insecure_https"`
	Headers                     types.Map    `tfsdk:"headers"`
	StripTrailingSlashesFromUrl types.Bool   `tfsdk:"strip_trailing_slashes_from_url"`
	RequestTimeout              types.Int32  `tfsdk:"request_timeout"`
}

func (p *netboxProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_token": schema.StringAttribute{
				MarkdownDescription: "Netbox API authentication token. Can be set via the `NETBOX_API_TOKEN` environment variable.",
				Optional:            true,
			},
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Location of Netbox server including scheme (http or https) and optional port. Can be set via the `NETBOX_SERVER_URL` environment variable.",
				Optional:            true,
			},
			"skip_version_check": schema.BoolAttribute{
				MarkdownDescription: "If true, do not try to determine the running Netbox version at provider startup. Disables warnings about possibly unsupported Netbox version. Also useful for local testing on terraform plans. Can be set via the `NETBOX_SKIP_VERSION_CHECK` environment variable. Defaults to `false`.",
				Optional:            true,
			},
			"allow_insecure_https": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Flag to set whether to allow https with invalid certificates. Can be set via the `NETBOX_ALLOW_INSECURE_HTTPS` environment variable. Defaults to `false`.",
			},
			"headers": schema.MapAttribute{
				Optional:            true,
				MarkdownDescription: "Set these header on all requests to Netbox. Can be set via the `NETBOX_HEADERS` environment variable.",
				ElementType:         types.StringType,
			},
			"strip_trailing_slashes_from_url": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "If true, strip trailing slashes from the `server_url` parameter and print a warning when doing so. Note that using trailing slashes in the `server_url` parameter will usually lead to errors. Can be set via the `NETBOX_STRIP_TRAILING_SLASHES_FROM_URL` environment variable. Defaults to `true`.",
			},
			"request_timeout": schema.Int32Attribute{
				Optional:            true,
				MarkdownDescription: "Netbox API HTTP request timeout in seconds. Can be set via the `NETBOX_REQUEST_TIMEOUT` environment variable.",
			},
		},
	}
}

func (p *netboxProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data netboxProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	apiToken := os.Getenv("NETBOX_API_TOKEN")
	serverUrl := os.Getenv("NETBOX_SERVER_URL")

	insecureHttps := false
	if os.Getenv("NETBOX_ALLOW_INSECURE_HTTPS") != "" {
		var err error
		insecureHttps, err = strconv.ParseBool(os.Getenv("NETBOX_ALLOW_INSECURE_HTTPS"))
		if err != nil {
			resp.Diagnostics.AddError("Unable to decode NETBOX_ALLOW_INSECURE_HTTPS", err.Error())
			return
		}
	} else if !data.AllowInsecureHttps.IsNull() {
		insecureHttps = data.AllowInsecureHttps.ValueBool()
	}

	requestTimeout := 10
	if os.Getenv("NETBOX_REQUEST_TIMEOUT") != "" {
		var err error
		requestTimeout, err = strconv.Atoi(os.Getenv("NETBOX_REQUEST_TIMEOUT"))
		if err != nil {
			resp.Diagnostics.AddError("Unable to decode NETBOX_REQUEST_TIMEOUT", err.Error())
			return
		}
	} else if !data.RequestTimeout.IsNull() {
		requestTimeout = int(data.RequestTimeout.ValueInt32())
	}

	stripTrailingSlashesFromURL := true
	if os.Getenv("NETBOX_STRIP_TRAILING_SLASHES_FROM_URL") != "" {
		var err error
		stripTrailingSlashesFromURL, err = strconv.ParseBool(os.Getenv("NETBOX_STRIP_TRAILING_SLASHES_FROM_URL"))
		if err != nil {
			resp.Diagnostics.AddError("Unable to decode NETBOX_STRIP_TRAILING_SLASHES_FROM_URL", err.Error())
			return
		}
	} else if !data.StripTrailingSlashesFromUrl.IsNull() {
		stripTrailingSlashesFromURL = data.StripTrailingSlashesFromUrl.ValueBool()
	}

	if !data.ApiToken.IsNull() {
		apiToken = data.ApiToken.ValueString()
	}

	if !data.ServerUrl.IsNull() {
		serverUrl = data.ServerUrl.ValueString()
	}

	if apiToken == "" {
		resp.Diagnostics.AddError(
			"Missing API Token Configuration",
			"TODO DETAIL")
		return
	}

	if serverUrl == "" {
		resp.Diagnostics.AddError(
			"Missing server URL configuration.",
			"TODO details")
		return
	}

	if stripTrailingSlashesFromURL {
		trimmed := false

		// This is Go's poor man's while loop
		for strings.HasSuffix(serverUrl, "/") {
			serverUrl = strings.TrimRight(serverUrl, "/")
			trimmed = true
		}
		if trimmed {
			resp.Diagnostics.AddWarning("Stripped trailing slashes from the `server_url` parameter",
				"Trailing slashes in the `server_url` parameter lead to problems in most setups, so all trailing slashes were stripped. Use the `strip_trailing_slashes_from_url` parameter to disable this feature or remove all trailing slashes in the `server_url` to disable this warning.")
		}
	}

	config := netbox.NewConfiguration()

	//Testing URL
	parsedURL, urlParseError := urlx.Parse(serverUrl)
	if urlParseError != nil {
		resp.Diagnostics.AddError("Unable to parse server URL",
			urlParseError.Error())
		return
	}
	config.Servers[0].URL = parsedURL.String()

	//Set Authorization token
	config.AddDefaultHeader("Authorization", fmt.Sprintf("Token %v", apiToken))
	config.AddDefaultHeader("Accept-Language", "en-US")

	//Import all the headers
	for index, entry := range data.Headers.Elements() {
		config.AddDefaultHeader(index, entry.String())
	}

	//Build http client
	clientOpts := httptransport.TLSClientOptions{
		InsecureSkipVerify: insecureHttps,
	}

	trans, err := httptransport.TLSTransport(clientOpts)
	if err != nil {
		resp.Diagnostics.AddError("Unable to set TLS transport mode", err.Error())
		return
	}

	trans.(*http.Transport).Proxy = http.ProxyFromEnvironment
	httpClient := &http.Client{
		Transport: trans,
		Timeout:   time.Second * time.Duration(requestTimeout),
	}
	config.HTTPClient = httpClient
	c := netbox.NewAPIClient(config)

	if !data.SkipVersionCheck.ValueBool() {
		response, _, err := c.StatusAPI.StatusRetrieve(ctx).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error getting netbox status.", err.Error())
			return
		}
		netboxVersion := response["netbox-version"].(string)
		supportedVersions := []string{"4.2.4"}
		if !slices.Contains(supportedVersions, netboxVersion) {
			resp.Diagnostics.AddWarning("Possibly unsupported Netbox version", fmt.Sprintf("Your Netbox version is v%v. The provider was successfully tested against the following versions:\n\n  %v\n\nUnexpected errors may occur.", netboxVersion, strings.Join(supportedVersions, ", ")))
		}
	}

	p.client = c
	resp.ResourceData = p
	resp.DataSourceData = p
}

func (p *netboxProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "netbox"
	resp.Version = "DEV"
}

func (p *netboxProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		func() datasource.DataSource {
			return &tagDataSource{}
		},
		func() datasource.DataSource {
			return &webhookDataSource{}
		},
	}
}

func (p *netboxProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		func() resource.Resource {
			return &webhookResource{}
		},
		func() resource.Resource {
			return &tagResource{}
		},
	}
}
