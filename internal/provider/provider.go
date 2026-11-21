package provider

import (
	"context"
	"os"
	"strconv"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/acmeclient"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/auth"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/caddy"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/captiveportal"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/chrony"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/core"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/cron"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/crowdsec"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/dhcrelay"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/diagnostics"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/dnsmasq"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/firewall"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/haproxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/hasync"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/hostdiscovery"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/ids"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/interfaces"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/ipsec"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/kea"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/monit"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/nodeexporter"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/nut"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/openvpn"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/proxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/quagga"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/radvd"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/routes"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/routing"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/syslog"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/tailscale"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/telegraf"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/trafficshaper"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/trust"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/unbound"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/wireguard"
	"github.com/browningluke/terraform-provider-opnsense/internal/service/zerotier"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure OPNsenseProvider satisfies various provider interfaces.
var _ provider.Provider = &opnsenseProvider{}
var _ provider.ProviderWithEphemeralResources = &opnsenseProvider{}

// OPNsenseProvider defines the provider implementation.
type opnsenseProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// OPNsenseProviderModel describes the provider data model.
type OPNsenseProviderModel struct {
	Uri           types.String `tfsdk:"uri"`
	APIKey        types.String `tfsdk:"api_key"`
	APISecret     types.String `tfsdk:"api_secret"`
	AllowInsecure types.Bool   `tfsdk:"allow_insecure"`
	MaxBackoff    types.Int64  `tfsdk:"max_backoff"`
	MinBackoff    types.Int64  `tfsdk:"min_backoff"`
	MaxRetries    types.Int64  `tfsdk:"retries"`
	Timeout       types.Int64  `tfsdk:"request_timeout"`
}

func (p *opnsenseProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "opnsense"
	resp.Version = p.version
}

func (p *opnsenseProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"uri": schema.StringAttribute{
				// Required, but computed from environment variable if possible
				MarkdownDescription: "The URI to an OPNsense host. Alternatively, can be configured using the `OPNSENSE_URI` environment variable.",
				Optional:            true,
			},
			"api_key": schema.StringAttribute{
				// Required, but computed from environment variable if possible
				MarkdownDescription: "The API key for a user. Alternatively, can be configured using the `OPNSENSE_API_KEY` environment variable.",
				Optional:            true,
			},
			"api_secret": schema.StringAttribute{
				// Required, but computed from environment variable if possible
				MarkdownDescription: "The API secret for a user. Alternatively, can be configured using the `OPNSENSE_API_SECRET` environment variable.",
				Optional:            true,
			},
			"allow_insecure": schema.BoolAttribute{
				MarkdownDescription: "Allow insecure TLS connections. Alternatively, can be configured using the `OPNSENSE_ALLOW_INSECURE` environment variable. Defaults to `false`.",
				Optional:            true,
			},
			"max_backoff": schema.Int64Attribute{
				MarkdownDescription: "Maximum backoff period in seconds after failed API calls. Alternatively, can be configured using the `OPNSENSE_MAX_BACKOFF` environment variable.",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"min_backoff": schema.Int64Attribute{
				MarkdownDescription: "Minimum backoff period in seconds after failed API calls. Alternatively, can be configured using the `OPNSENSE_MIN_BACKOFF` environment variable.",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"retries": schema.Int64Attribute{
				MarkdownDescription: "Maximum number of retries to perform when an API request fails. Alternatively, can be configured using the `OPNSENSE_RETRIES` environment variable.",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 2147483647), // Since we convert the int64 to an int(32), set an upper bound.
				},
			},
			"request_timeout": schema.Int64Attribute{
				MarkdownDescription: "Timeout in seconds for a single API request, including reading the response. A request that times out is retried. Alternatively, can be configured using the `OPNSENSE_REQUEST_TIMEOUT` environment variable. Defaults to `120`.",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
		},
	}
}

func (p *opnsenseProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Get data from config
	var data OPNsenseProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If value is set, it must be known

	if data.Uri.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("uri"),
			"Unknown OPNsense API URI",
			"The provider cannot create the OPNsense API client as there is an unknown configuration value for the OPNsense API uri. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the OPNSENSE_URI environment variable.",
		)
	}

	if data.APIKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Unknown OPNsense API Key",
			"The provider cannot create the OPNsense API client as there is an unknown configuration value for the OPNsense API key. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the OPNSENSE_API_KEY environment variable.",
		)
	}

	if data.APISecret.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_secret"),
			"Unknown OPNsense API Secret",
			"The provider cannot create the OPNsense API client as there is an unknown configuration value for the OPNsense API secret. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the OPNSENSE_API_SECRET environment variable.",
		)
	}

	if data.AllowInsecure.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("allow_insecure"),
			"Unknown OPNsense API Value: allow_insecure",
			"The provider cannot create the OPNsense API client as there is an unknown configuration value for allow_insecure. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the OPNSENSE_ALLOW_INSECURE environment variable.",
		)
	}

	if data.MaxBackoff.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("max_backoff"),
			"Unknown OPNsense API Value: max_backoff",
			"The provider cannot create the OPNsense API client as there is an unknown configuration value for max_backoff. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the OPNSENSE_MAX_BACKOFF environment variable.",
		)
	}

	if data.MinBackoff.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("min_backoff"),
			"Unknown OPNsense API Value: min_backoff",
			"The provider cannot create the OPNsense API client as there is an unknown configuration value for min_backoff. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the OPNSENSE_MIN_BACKOFF environment variable.",
		)
	}

	if data.MaxRetries.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("retries"),
			"Unknown OPNsense API Value: retries",
			"The provider cannot create the OPNsense API client as there is an unknown configuration value for retries. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the OPNSENSE_RETRIES environment variable.",
		)
	}

	if data.Timeout.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("request_timeout"),
			"Unknown OPNsense API Value: request_timeout",
			"The provider cannot create the OPNsense API client as there is an unknown configuration value for request_timeout. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the OPNSENSE_REQUEST_TIMEOUT environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Attempt to load values from environment variables

	uri := os.Getenv("OPNSENSE_URI")
	if !data.Uri.IsNull() {
		uri = data.Uri.ValueString()
	}

	apiKey := os.Getenv("OPNSENSE_API_KEY")
	if !data.APIKey.IsNull() {
		apiKey = data.APIKey.ValueString()
	}

	apiSecret := os.Getenv("OPNSENSE_API_SECRET")
	if !data.APISecret.IsNull() {
		apiSecret = data.APISecret.ValueString()
	}

	allowInsecureStr := os.Getenv("OPNSENSE_ALLOW_INSECURE")
	allowInsecure, err := strconv.ParseBool(allowInsecureStr)
	if err != nil {
		// Set to default (false) if string is unparsable
		allowInsecure = false
	}
	if !data.AllowInsecure.IsNull() {
		allowInsecure = data.AllowInsecure.ValueBool()
	}

	maxBackoffStr := os.Getenv("OPNSENSE_MAX_BACKOFF")
	maxBackoff, err := strconv.ParseInt(maxBackoffStr, 10, 64)
	if err != nil {
		// Set to 0 to use client default downstream if string is unparsable
		maxBackoff = 0
	}
	if !data.MaxBackoff.IsNull() {
		maxBackoff = data.MaxBackoff.ValueInt64()
	}

	minBackoffStr := os.Getenv("OPNSENSE_MIN_BACKOFF")
	minBackoff, err := strconv.ParseInt(minBackoffStr, 10, 64)
	if err != nil {
		// Set to 0 to use client default downstream if string is unparsable
		minBackoff = 0
	}
	if !data.MinBackoff.IsNull() {
		minBackoff = data.MinBackoff.ValueInt64()
	}

	retriesStr := os.Getenv("OPNSENSE_RETRIES")
	retries, err := strconv.ParseInt(retriesStr, 10, 64)
	if err != nil {
		// Set to 0 to use client default downstream if string is unparsable
		retries = 0
	}
	if !data.MaxRetries.IsNull() {
		retries = data.MaxRetries.ValueInt64()
	}

	timeout, err := strconv.ParseInt(os.Getenv("OPNSENSE_REQUEST_TIMEOUT"), 10, 64)
	if err != nil {
		timeout = 0
	}
	if !data.Timeout.IsNull() {
		timeout = data.Timeout.ValueInt64()
	}

	// Ensure expected variables are not empty

	if uri == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("uri"),
			"Missing OPNsense API URI",
			"The provider cannot create the OPNsense API client as there is a missing or empty value for the OPNsense API uri. "+
				"Set the host value in the configuration or use the OPNSENSE_URI environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if apiKey == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing OPNsense API Key",
			"The provider cannot create the OPNsense API client as there is a missing or empty value for the OPNsense API key. "+
				"Set the host value in the configuration or use the OPNSENSE_API_KEY environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if apiSecret == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_secret"),
			"Missing OPNsense API Secret",
			"The provider cannot create the OPNsense API client as there is a missing or empty value for the OPNsense API secret. "+
				"Set the host value in the configuration or use the OPNSENSE_API_SECRET environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	// Create the OPNsense client
	opnOptions := api.Options{
		Uri:            uri,
		APIKey:         apiKey,
		APISecret:      apiSecret,
		AllowInsecure:  allowInsecure,
		MaxBackoff:     maxBackoff,
		MinBackoff:     minBackoff,
		MaxRetries:     retries,
		RequestTimeout: timeout,
	}
	client := api.NewClient(opnOptions)

	resp.DataSourceData = client
	resp.ResourceData = client
	resp.EphemeralResourceData = client
}

func (p *opnsenseProvider) Resources(ctx context.Context) []func() resource.Resource {
	controllers := [][]func() resource.Resource{
		diagnostics.Resources(ctx),
		dhcrelay.Resources(ctx),
		dnsmasq.Resources(ctx),
		firewall.Resources(ctx),
		hasync.Resources(ctx),
		zerotier.Resources(ctx),
		telegraf.Resources(ctx),
		nut.Resources(ctx),
		crowdsec.Resources(ctx),
		haproxy.Resources(ctx),
		acmeclient.Resources(ctx),
		caddy.Resources(ctx),
		tailscale.Resources(ctx),
		chrony.Resources(ctx),
		nodeexporter.Resources(ctx),
		hostdiscovery.Resources(ctx),
		captiveportal.Resources(ctx),
		ids.Resources(ctx),
		interfaces.Resources(ctx),
		ipsec.Resources(ctx),
		auth.Resources(ctx),
		core.Resources(ctx),
		cron.Resources(ctx),
		kea.Resources(ctx),
		monit.Resources(ctx),
		openvpn.Resources(ctx),
		quagga.Resources(ctx),
		routes.Resources(ctx),
		routing.Resources(ctx),
		proxy.Resources(ctx),
		radvd.Resources(ctx),
		syslog.Resources(ctx),
		trafficshaper.Resources(ctx),
		trust.Resources(ctx),
		unbound.Resources(ctx),
		wireguard.Resources(ctx),
	}

	var resources []func() resource.Resource
	for _, s := range controllers {
		resources = append(resources, s...)
	}
	return resources
}

func (p *opnsenseProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	controllers := [][]func() datasource.DataSource{
		diagnostics.DataSources(ctx),
		dhcrelay.DataSources(ctx),
		dnsmasq.DataSources(ctx),
		firewall.DataSources(ctx),
		hasync.DataSources(ctx),
		zerotier.DataSources(ctx),
		telegraf.DataSources(ctx),
		nut.DataSources(ctx),
		crowdsec.DataSources(ctx),
		haproxy.DataSources(ctx),
		acmeclient.DataSources(ctx),
		caddy.DataSources(ctx),
		tailscale.DataSources(ctx),
		chrony.DataSources(ctx),
		nodeexporter.DataSources(ctx),
		hostdiscovery.DataSources(ctx),
		captiveportal.DataSources(ctx),
		ids.DataSources(ctx),
		interfaces.DataSources(ctx),
		ipsec.DataSources(ctx),
		auth.DataSources(ctx),
		core.DataSources(ctx),
		cron.DataSources(ctx),
		kea.DataSources(ctx),
		monit.DataSources(ctx),
		openvpn.DataSources(ctx),
		quagga.DataSources(ctx),
		routes.DataSources(ctx),
		routing.DataSources(ctx),
		proxy.DataSources(ctx),
		radvd.DataSources(ctx),
		syslog.DataSources(ctx),
		trafficshaper.DataSources(ctx),
		trust.DataSources(ctx),
		unbound.DataSources(ctx),
		wireguard.DataSources(ctx),
	}

	var dataSources []func() datasource.DataSource
	for _, s := range controllers {
		dataSources = append(dataSources, s...)
	}
	return dataSources
}

func (p *opnsenseProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	controllers := [][]func() ephemeral.EphemeralResource{
		openvpn.EphemeralResources(ctx),
	}

	var ephemerals []func() ephemeral.EphemeralResource
	for _, s := range controllers {
		ephemerals = append(ephemerals, s...)
	}
	return ephemerals
}

func NewProvider(ctx context.Context) (provider.Provider, error) {
	return &opnsenseProvider{}, nil
}
