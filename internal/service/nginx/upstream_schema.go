package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type upstreamResourceModel struct {
	Description            types.String `tfsdk:"description"`
	Serverentries          types.Set    `tfsdk:"serverentries"`
	LoadBalancingAlgorithm types.String `tfsdk:"load_balancing_algorithm"`
	Keepalive              types.Int64  `tfsdk:"keepalive"`
	KeepaliveRequests      types.Int64  `tfsdk:"keepalive_requests"`
	KeepaliveTimeout       types.Int64  `tfsdk:"keepalive_timeout"`
	HostPort               types.Int64  `tfsdk:"host_port"`
	XForwardedHostVerbatim types.Bool   `tfsdk:"x_forwarded_host_verbatim"`
	ProxyProtocol          types.Bool   `tfsdk:"proxy_protocol"`
	Store                  types.Bool   `tfsdk:"store"`
	TLSEnable              types.Bool   `tfsdk:"tls_enable"`
	TLSClientCertificate   types.String `tfsdk:"tls_client_certificate"`
	TLSNameOverride        types.String `tfsdk:"tls_name_override"`
	TLSProtocolVersions    types.Set    `tfsdk:"tls_protocol_versions"`
	TLSSessionReuse        types.Bool   `tfsdk:"tls_session_reuse"`
	TLSTrustedCertificate  types.Set    `tfsdk:"tls_trusted_certificate"`
	TLSVerify              types.Bool   `tfsdk:"tls_verify"`
	TLSVerifyDepth         types.Int64  `tfsdk:"tls_verify_depth"`

	Id types.String `tfsdk:"id"`
}

func upstreamResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx upstreams (**Services ▸ Nginx ▸ Configuration ▸ Upstream ▸ Upstream**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				MarkdownDescription: "Description.",
				Required:            true,
			},
			"serverentries": schema.SetAttribute{
				MarkdownDescription: "Server Entries.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"load_balancing_algorithm": schema.StringAttribute{
				MarkdownDescription: "Load Balancing Algorithm. One of `ip_hash`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "ip_hash")},
			},
			"keepalive": schema.Int64Attribute{
				MarkdownDescription: "Activates the connections cache for upstream server and sets the max number of idle connections that are preserved in the cache of each worker process. Leave blank or set to 0 to disable. Keepalive support should be enabled at Location settings also. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"keepalive_requests": schema.Int64Attribute{
				MarkdownDescription: "Sets the maximum number of requests that can be served through one keepalive connection. NGINX default is 1000. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"keepalive_timeout": schema.Int64Attribute{
				MarkdownDescription: "Sets a timeout in seconds during which an idle keepalive connection to an upstream server will stay open. NGINX default is 75s. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"host_port": schema.Int64Attribute{
				MarkdownDescription: "Add this port value to the Host header. Not used by default. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"x_forwarded_host_verbatim": schema.BoolAttribute{
				MarkdownDescription: "Use Host header value from the client request ($http_host) for X-Forwarded-Host header. $host variable is used by default. Enabling this may cause incorrect behavior in case of malicious requests such as incorrect hostnames being logged or invalid redirects being performed. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"proxy_protocol": schema.BoolAttribute{
				MarkdownDescription: "If you enable the proxy protocol, an upstream proxy or server will get the client IP and the server port before the real traffic is sent. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"store": schema.BoolAttribute{
				MarkdownDescription: "Store the response on the local storage. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tls_enable": schema.BoolAttribute{
				MarkdownDescription: "Use TLS (HTTPS) to connect to the server. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tls_client_certificate": schema.StringAttribute{
				MarkdownDescription: "Authenticate on the Server using a client certificate. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tls_name_override": schema.StringAttribute{
				MarkdownDescription: "Force a specific hostname for the backend connection instead of passing the hostname from the downstream connection to the upstream connection. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tls_protocol_versions": schema.SetAttribute{
				MarkdownDescription: "TLS: Supported Versions. One of `TLSv1`, `TLSv1.1`, `TLSv1.2`, `TLSv1.3`. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
				Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("TLSv1", "TLSv1.1", "TLSv1.2", "TLSv1.3"))},
			},
			"tls_session_reuse": schema.BoolAttribute{
				MarkdownDescription: "TLS: Session Reuse. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"tls_trusted_certificate": schema.SetAttribute{
				MarkdownDescription: "TLS: Trusted Certificate. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"tls_verify": schema.BoolAttribute{
				MarkdownDescription: "Don't turn it off unless you really know what you are doing! Never do it because a random website tells you to do. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"tls_verify_depth": schema.Int64Attribute{
				MarkdownDescription: "Choose how many sub-CAs can be between the server certificate and a trusted CA. 1 means the certificate has to be signed directly by a CA. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx upstream.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func upstreamDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx upstreams (**Services ▸ Nginx ▸ Configuration ▸ Upstream ▸ Upstream**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx upstream.",
				Required:            true,
			},
			"description":               dschema.StringAttribute{MarkdownDescription: "Description.", Computed: true},
			"serverentries":             dschema.SetAttribute{MarkdownDescription: "Server Entries.", Computed: true, ElementType: types.StringType},
			"load_balancing_algorithm":  dschema.StringAttribute{MarkdownDescription: "Load Balancing Algorithm. One of `ip_hash`.", Computed: true},
			"keepalive":                 dschema.Int64Attribute{MarkdownDescription: "Activates the connections cache for upstream server and sets the max number of idle connections that are preserved in the cache of each worker process. Leave blank or set to 0 to disable. Keepalive support should be enabled at Location settings also.", Computed: true},
			"keepalive_requests":        dschema.Int64Attribute{MarkdownDescription: "Sets the maximum number of requests that can be served through one keepalive connection. NGINX default is 1000.", Computed: true},
			"keepalive_timeout":         dschema.Int64Attribute{MarkdownDescription: "Sets a timeout in seconds during which an idle keepalive connection to an upstream server will stay open. NGINX default is 75s.", Computed: true},
			"host_port":                 dschema.Int64Attribute{MarkdownDescription: "Add this port value to the Host header. Not used by default.", Computed: true},
			"x_forwarded_host_verbatim": dschema.BoolAttribute{MarkdownDescription: "Use Host header value from the client request ($http_host) for X-Forwarded-Host header. $host variable is used by default. Enabling this may cause incorrect behavior in case of malicious requests such as incorrect hostnames being logged or invalid redirects being performed.", Computed: true},
			"proxy_protocol":            dschema.BoolAttribute{MarkdownDescription: "If you enable the proxy protocol, an upstream proxy or server will get the client IP and the server port before the real traffic is sent.", Computed: true},
			"store":                     dschema.BoolAttribute{MarkdownDescription: "Store the response on the local storage.", Computed: true},
			"tls_enable":                dschema.BoolAttribute{MarkdownDescription: "Use TLS (HTTPS) to connect to the server.", Computed: true},
			"tls_client_certificate":    dschema.StringAttribute{MarkdownDescription: "Authenticate on the Server using a client certificate.", Computed: true},
			"tls_name_override":         dschema.StringAttribute{MarkdownDescription: "Force a specific hostname for the backend connection instead of passing the hostname from the downstream connection to the upstream connection.", Computed: true},
			"tls_protocol_versions":     dschema.SetAttribute{MarkdownDescription: "TLS: Supported Versions. One of `TLSv1`, `TLSv1.1`, `TLSv1.2`, `TLSv1.3`.", Computed: true, ElementType: types.StringType},
			"tls_session_reuse":         dschema.BoolAttribute{MarkdownDescription: "TLS: Session Reuse.", Computed: true},
			"tls_trusted_certificate":   dschema.SetAttribute{MarkdownDescription: "TLS: Trusted Certificate.", Computed: true, ElementType: types.StringType},
			"tls_verify":                dschema.BoolAttribute{MarkdownDescription: "Don't turn it off unless you really know what you are doing! Never do it because a random website tells you to do.", Computed: true},
			"tls_verify_depth":          dschema.Int64Attribute{MarkdownDescription: "Choose how many sub-CAs can be between the server certificate and a trusted CA. 1 means the certificate has to be signed directly by a CA.", Computed: true},
		},
	}
}

func convertUpstreamSchemaToStruct(d *upstreamResourceModel) (*nginx.Upstream, error) {
	return &nginx.Upstream{
		Description:            d.Description.ValueString(),
		Serverentries:          api.SelectedMapList(setToSlice(d.Serverentries)),
		LoadBalancingAlgorithm: api.SelectedMap(d.LoadBalancingAlgorithm.ValueString()),
		Keepalive:              tools.Int64ToStringNegative(d.Keepalive.ValueInt64()),
		KeepaliveRequests:      tools.Int64ToStringNegative(d.KeepaliveRequests.ValueInt64()),
		KeepaliveTimeout:       tools.Int64ToStringNegative(d.KeepaliveTimeout.ValueInt64()),
		HostPort:               tools.Int64ToStringNegative(d.HostPort.ValueInt64()),
		XForwardedHostVerbatim: tools.BoolToString(d.XForwardedHostVerbatim.ValueBool()),
		ProxyProtocol:          tools.BoolToString(d.ProxyProtocol.ValueBool()),
		Store:                  tools.BoolToString(d.Store.ValueBool()),
		TlsEnable:              tools.BoolToString(d.TLSEnable.ValueBool()),
		TlsClientCertificate:   api.SelectedMap(d.TLSClientCertificate.ValueString()),
		TlsNameOverride:        d.TLSNameOverride.ValueString(),
		TlsProtocolVersions:    api.SelectedMapList(setToSlice(d.TLSProtocolVersions)),
		TlsSessionReuse:        tools.BoolToString(d.TLSSessionReuse.ValueBool()),
		TlsTrustedCertificate:  api.SelectedMapList(setToSlice(d.TLSTrustedCertificate)),
		TlsVerify:              tools.BoolToString(d.TLSVerify.ValueBool()),
		TlsVerifyDepth:         tools.Int64ToStringNegative(d.TLSVerifyDepth.ValueInt64()),
	}, nil
}

func convertUpstreamStructToSchema(d *nginx.Upstream) (*upstreamResourceModel, error) {
	return &upstreamResourceModel{
		Description:            types.StringValue(d.Description),
		Serverentries:          tools.StringSliceToSet(d.Serverentries),
		LoadBalancingAlgorithm: types.StringValue(d.LoadBalancingAlgorithm.String()),
		Keepalive:              types.Int64Value(tools.StringToInt64(d.Keepalive)),
		KeepaliveRequests:      types.Int64Value(tools.StringToInt64(d.KeepaliveRequests)),
		KeepaliveTimeout:       types.Int64Value(tools.StringToInt64(d.KeepaliveTimeout)),
		HostPort:               types.Int64Value(tools.StringToInt64(d.HostPort)),
		XForwardedHostVerbatim: types.BoolValue(tools.StringToBool(d.XForwardedHostVerbatim)),
		ProxyProtocol:          types.BoolValue(tools.StringToBool(d.ProxyProtocol)),
		Store:                  types.BoolValue(tools.StringToBool(d.Store)),
		TLSEnable:              types.BoolValue(tools.StringToBool(d.TlsEnable)),
		TLSClientCertificate:   types.StringValue(d.TlsClientCertificate.String()),
		TLSNameOverride:        types.StringValue(d.TlsNameOverride),
		TLSProtocolVersions:    tools.StringSliceToSet(d.TlsProtocolVersions),
		TLSSessionReuse:        types.BoolValue(tools.StringToBool(d.TlsSessionReuse)),
		TLSTrustedCertificate:  tools.StringSliceToSet(d.TlsTrustedCertificate),
		TLSVerify:              types.BoolValue(tools.StringToBool(d.TlsVerify)),
		TLSVerifyDepth:         types.Int64Value(tools.StringToInt64(d.TlsVerifyDepth)),
	}, nil
}
