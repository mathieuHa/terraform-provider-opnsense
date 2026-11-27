package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
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

type streamServerResourceModel struct {
	ListenAddress       types.Set    `tfsdk:"listen_address"`
	SyslogTargets       types.Set    `tfsdk:"syslog_targets"`
	Udp                 types.Bool   `tfsdk:"udp"`
	TrustedProxies      types.Set    `tfsdk:"trusted_proxies"`
	ProxyProtocol       types.Bool   `tfsdk:"proxy_protocol"`
	Certificate         types.String `tfsdk:"certificate"`
	Ca                  types.String `tfsdk:"ca"`
	VerifyClient        types.String `tfsdk:"verify_client"`
	AccessLogFormat     types.String `tfsdk:"access_log_format"`
	ErrorLogLevel       types.String `tfsdk:"error_log_level"`
	RouteField          types.String `tfsdk:"route_field"`
	Upstream            types.String `tfsdk:"upstream"`
	SniUpstreamMap      types.String `tfsdk:"sni_upstream_map"`
	IPAcl               types.String `tfsdk:"ip_acl"`
	ProxyResponses      types.Int64  `tfsdk:"proxy_responses"`
	ProxyConnectTimeout types.Int64  `tfsdk:"proxy_connect_timeout"`
	ProxyTimeout        types.Int64  `tfsdk:"proxy_timeout"`

	Id types.String `tfsdk:"id"`
}

func streamServerResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx stream servers (**Services ▸ Nginx ▸ Configuration ▸ Streams ▸ Stream Server**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"listen_address": schema.SetAttribute{
				MarkdownDescription: "Enter a list of IP addresses and ports which can be used in nginx listen directives. To listen on a port on all IPs, use for example \"22,[::]:22\". Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"syslog_targets": schema.SetAttribute{
				MarkdownDescription: "SYSLOG targets. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"udp": schema.BoolAttribute{
				MarkdownDescription: "UDP Port. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"trusted_proxies": schema.SetAttribute{
				MarkdownDescription: "Enter a list of IP addresses or CIDR networks which are allowed to override the source IP address using the specified header. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"proxy_protocol": schema.BoolAttribute{
				MarkdownDescription: "If you enable the proxy protocol, a downstream proxy can send the client IP and port before the real traffic is set. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"certificate": schema.StringAttribute{
				MarkdownDescription: "TLS Certificate. Takes the `ref_id` of an `opnsense_trust_cert`, not its UUID. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ca": schema.StringAttribute{
				MarkdownDescription: "CA Certificate. Takes the `ref_id` of an `opnsense_trust_ca`, not its UUID. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"verify_client": schema.StringAttribute{
				MarkdownDescription: "On: the certificate is requested and validated. Use this option to protect a service with TLS authentication.Off: The certificate is not requested. Choose this option for a normal website.Optional: The certificate is requested and validated if existing. Choose this option for websites, with TLS login support or mixed TLS protected API and web content.Optional, don't verify: Do accept the certificate and let the application choose what to do. Choose this option, for the same reasons as optional but in this case, the request is passed to the backend without rejecting untrusted certificates. One of `off`, `on`, `optional`, `optional_no_ca`. Defaults to `off`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("off"),
				Validators:          []validator.String{stringvalidator.OneOf("off", "on", "optional", "optional_no_ca")},
			},
			"access_log_format": schema.StringAttribute{
				MarkdownDescription: "Access Log Format. One of `main`, `main_ext`, `anonymized`, `disabled`. Defaults to `main`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("main"),
				Validators:          []validator.String{stringvalidator.OneOf("main", "main_ext", "anonymized", "disabled")},
			},
			"error_log_level": schema.StringAttribute{
				MarkdownDescription: "Select Error Log Level. Log levels are listed in the order of increasing verbosity. Setting a certain log level will cause all messages of the specified and more severe log levels to be logged. One of `emerg`, `alert`, `crit`, `error`, `warn`, `notice`, `info`. Defaults to `info`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("info"),
				Validators:          []validator.String{stringvalidator.OneOf("emerg", "alert", "crit", "error", "warn", "notice", "info")},
			},
			"route_field": schema.StringAttribute{
				MarkdownDescription: "Route With. One of `upstream`, `sni_upstream_map`. Defaults to `upstream`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("upstream"),
				Validators:          []validator.String{stringvalidator.OneOf("upstream", "sni_upstream_map")},
			},
			"upstream": schema.StringAttribute{
				MarkdownDescription: "Select an upstream to proxy to. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"sni_upstream_map": schema.StringAttribute{
				MarkdownDescription: "Select an upstream map to choose the host based on the name given by the client. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ip_acl": schema.StringAttribute{
				MarkdownDescription: "If you select an IP ACL, the client can only access this service if it fulfills this requirement. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"proxy_responses": schema.Int64Attribute{
				MarkdownDescription: "Due to the nature of UDP, nginx cannot know, when the communication ends and this helps as it tells nginx the number of datagrams the communication is expected to last on server side and it is expected to be closed afterwards. If you enter 0, it is expected, that the server never responds to a datagram. If nginx gets a datagram, it will still get forwarded to the client. Setting this option might be useful in (mostly) unidirectional communication as well. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"proxy_connect_timeout": schema.Int64Attribute{
				MarkdownDescription: "Defines a timeout (in seconds) for establishing a connection with a proxied server. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"proxy_timeout": schema.Int64Attribute{
				MarkdownDescription: "Sets the timeout (in seconds) between two successive read or write operations on client or proxied server connections. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx stream server.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func streamServerDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx stream servers (**Services ▸ Nginx ▸ Configuration ▸ Streams ▸ Stream Server**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx stream server.",
				Required:            true,
			},
			"listen_address":        dschema.SetAttribute{MarkdownDescription: "Enter a list of IP addresses and ports which can be used in nginx listen directives. To listen on a port on all IPs, use for example \"22,[::]:22\".", Computed: true, ElementType: types.StringType},
			"syslog_targets":        dschema.SetAttribute{MarkdownDescription: "SYSLOG targets.", Computed: true, ElementType: types.StringType},
			"udp":                   dschema.BoolAttribute{MarkdownDescription: "UDP Port.", Computed: true},
			"trusted_proxies":       dschema.SetAttribute{MarkdownDescription: "Enter a list of IP addresses or CIDR networks which are allowed to override the source IP address using the specified header.", Computed: true, ElementType: types.StringType},
			"proxy_protocol":        dschema.BoolAttribute{MarkdownDescription: "If you enable the proxy protocol, a downstream proxy can send the client IP and port before the real traffic is set.", Computed: true},
			"certificate":           dschema.StringAttribute{MarkdownDescription: "TLS Certificate.", Computed: true},
			"ca":                    dschema.StringAttribute{MarkdownDescription: "CA Certificate.", Computed: true},
			"verify_client":         dschema.StringAttribute{MarkdownDescription: "On: the certificate is requested and validated. Use this option to protect a service with TLS authentication.Off: The certificate is not requested. Choose this option for a normal website.Optional: The certificate is requested and validated if existing. Choose this option for websites, with TLS login support or mixed TLS protected API and web content.Optional, don't verify: Do accept the certificate and let the application choose what to do. Choose this option, for the same reasons as optional but in this case, the request is passed to the backend without rejecting untrusted certificates. One of `off`, `on`, `optional`, `optional_no_ca`.", Computed: true},
			"access_log_format":     dschema.StringAttribute{MarkdownDescription: "Access Log Format. One of `main`, `main_ext`, `anonymized`, `disabled`.", Computed: true},
			"error_log_level":       dschema.StringAttribute{MarkdownDescription: "Select Error Log Level. Log levels are listed in the order of increasing verbosity. Setting a certain log level will cause all messages of the specified and more severe log levels to be logged. One of `emerg`, `alert`, `crit`, `error`, `warn`, `notice`, `info`.", Computed: true},
			"route_field":           dschema.StringAttribute{MarkdownDescription: "Route With. One of `upstream`, `sni_upstream_map`.", Computed: true},
			"upstream":              dschema.StringAttribute{MarkdownDescription: "Select an upstream to proxy to.", Computed: true},
			"sni_upstream_map":      dschema.StringAttribute{MarkdownDescription: "Select an upstream map to choose the host based on the name given by the client.", Computed: true},
			"ip_acl":                dschema.StringAttribute{MarkdownDescription: "If you select an IP ACL, the client can only access this service if it fulfills this requirement.", Computed: true},
			"proxy_responses":       dschema.Int64Attribute{MarkdownDescription: "Due to the nature of UDP, nginx cannot know, when the communication ends and this helps as it tells nginx the number of datagrams the communication is expected to last on server side and it is expected to be closed afterwards. If you enter 0, it is expected, that the server never responds to a datagram. If nginx gets a datagram, it will still get forwarded to the client. Setting this option might be useful in (mostly) unidirectional communication as well.", Computed: true},
			"proxy_connect_timeout": dschema.Int64Attribute{MarkdownDescription: "Defines a timeout (in seconds) for establishing a connection with a proxied server.", Computed: true},
			"proxy_timeout":         dschema.Int64Attribute{MarkdownDescription: "Sets the timeout (in seconds) between two successive read or write operations on client or proxied server connections.", Computed: true},
		},
	}
}

func convertStreamServerSchemaToStruct(d *streamServerResourceModel) (*nginx.StreamServer, error) {
	return &nginx.StreamServer{
		ListenAddress:       api.SelectedMapList(setToSlice(d.ListenAddress)),
		SyslogTargets:       api.SelectedMapList(setToSlice(d.SyslogTargets)),
		Udp:                 tools.BoolToString(d.Udp.ValueBool()),
		TrustedProxies:      api.SelectedMapList(setToSlice(d.TrustedProxies)),
		ProxyProtocol:       tools.BoolToString(d.ProxyProtocol.ValueBool()),
		Certificate:         api.SelectedMap(d.Certificate.ValueString()),
		Ca:                  api.SelectedMap(d.Ca.ValueString()),
		VerifyClient:        api.SelectedMap(d.VerifyClient.ValueString()),
		AccessLogFormat:     api.SelectedMap(d.AccessLogFormat.ValueString()),
		ErrorLogLevel:       api.SelectedMap(d.ErrorLogLevel.ValueString()),
		RouteField:          api.SelectedMap(d.RouteField.ValueString()),
		Upstream:            api.SelectedMap(d.Upstream.ValueString()),
		SniUpstreamMap:      api.SelectedMap(d.SniUpstreamMap.ValueString()),
		IpAcl:               api.SelectedMap(d.IPAcl.ValueString()),
		ProxyResponses:      tools.Int64ToStringNegative(d.ProxyResponses.ValueInt64()),
		ProxyConnectTimeout: tools.Int64ToStringNegative(d.ProxyConnectTimeout.ValueInt64()),
		ProxyTimeout:        tools.Int64ToStringNegative(d.ProxyTimeout.ValueInt64()),
	}, nil
}

func convertStreamServerStructToSchema(d *nginx.StreamServer) (*streamServerResourceModel, error) {
	return &streamServerResourceModel{
		ListenAddress:       tools.StringSliceToSet(d.ListenAddress),
		SyslogTargets:       tools.StringSliceToSet(d.SyslogTargets),
		Udp:                 types.BoolValue(tools.StringToBool(d.Udp)),
		TrustedProxies:      tools.StringSliceToSet(d.TrustedProxies),
		ProxyProtocol:       types.BoolValue(tools.StringToBool(d.ProxyProtocol)),
		Certificate:         types.StringValue(d.Certificate.String()),
		Ca:                  types.StringValue(d.Ca.String()),
		VerifyClient:        types.StringValue(d.VerifyClient.String()),
		AccessLogFormat:     types.StringValue(d.AccessLogFormat.String()),
		ErrorLogLevel:       types.StringValue(d.ErrorLogLevel.String()),
		RouteField:          types.StringValue(d.RouteField.String()),
		Upstream:            types.StringValue(d.Upstream.String()),
		SniUpstreamMap:      types.StringValue(d.SniUpstreamMap.String()),
		IPAcl:               types.StringValue(d.IpAcl.String()),
		ProxyResponses:      types.Int64Value(tools.StringToInt64(d.ProxyResponses)),
		ProxyConnectTimeout: types.Int64Value(tools.StringToInt64(d.ProxyConnectTimeout)),
		ProxyTimeout:        types.Int64Value(tools.StringToInt64(d.ProxyTimeout)),
	}, nil
}
