package haproxy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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

type backendResourceModel struct {
	Enabled                      types.Bool   `tfsdk:"enabled"`
	Name                         types.String `tfsdk:"name"`
	Description                  types.String `tfsdk:"description"`
	Mode                         types.String `tfsdk:"mode"`
	Algorithm                    types.String `tfsdk:"algorithm"`
	RandomDraws                  types.Int64  `tfsdk:"random_draws"`
	ProxyProtocol                types.String `tfsdk:"proxy_protocol"`
	LinkedServers                types.Set    `tfsdk:"linked_servers"`
	LinkedFcgi                   types.String `tfsdk:"linked_fcgi"`
	LinkedResolver               types.String `tfsdk:"linked_resolver"`
	ResolverOpts                 types.Set    `tfsdk:"resolver_opts"`
	ResolvePrefer                types.String `tfsdk:"resolve_prefer"`
	Source                       types.String `tfsdk:"source"`
	HealthCheckEnabled           types.Bool   `tfsdk:"health_check_enabled"`
	HealthCheck                  types.String `tfsdk:"health_check"`
	HealthCheckLogStatus         types.Bool   `tfsdk:"health_check_log_status"`
	CheckInterval                types.String `tfsdk:"check_interval"`
	CheckDownInterval            types.String `tfsdk:"check_down_interval"`
	HealthCheckFall              types.Int64  `tfsdk:"health_check_fall"`
	HealthCheckRise              types.Int64  `tfsdk:"health_check_rise"`
	LinkedMailer                 types.String `tfsdk:"linked_mailer"`
	HealthCheckProxyProto        types.String `tfsdk:"health_check_proxy_proto"`
	Http2Enabled                 types.Bool   `tfsdk:"http2_enabled"`
	Http2EnabledNontls           types.Bool   `tfsdk:"http2_enabled_nontls"`
	BaAdvertisedProtocols        types.Set    `tfsdk:"ba_advertised_protocols"`
	ForwardFor                   types.Bool   `tfsdk:"forward_for"`
	ForwardedHeader              types.Bool   `tfsdk:"forwarded_header"`
	ForwardedHeaderParameters    types.Set    `tfsdk:"forwarded_header_parameters"`
	Persistence                  types.String `tfsdk:"persistence"`
	PersistenceCookiemode        types.String `tfsdk:"persistence_cookiemode"`
	PersistenceCookiename        types.String `tfsdk:"persistence_cookiename"`
	PersistenceStripquotes       types.Bool   `tfsdk:"persistence_stripquotes"`
	StickinessPattern            types.String `tfsdk:"stickiness_pattern"`
	StickinessDataTypes          types.Set    `tfsdk:"stickiness_data_types"`
	StickinessExpire             types.String `tfsdk:"stickiness_expire"`
	StickinessSize               types.String `tfsdk:"stickiness_size"`
	StickinessCookiename         types.String `tfsdk:"stickiness_cookiename"`
	StickinessCookielength       types.Int64  `tfsdk:"stickiness_cookielength"`
	StickinessLength             types.Int64  `tfsdk:"stickiness_length"`
	StickinessConnRatePeriod     types.String `tfsdk:"stickiness_conn_rate_period"`
	StickinessSessRatePeriod     types.String `tfsdk:"stickiness_sess_rate_period"`
	StickinessHttpReqRatePeriod  types.String `tfsdk:"stickiness_http_req_rate_period"`
	StickinessHttpErrRatePeriod  types.String `tfsdk:"stickiness_http_err_rate_period"`
	StickinessBytesInRatePeriod  types.String `tfsdk:"stickiness_bytes_in_rate_period"`
	StickinessBytesOutRatePeriod types.String `tfsdk:"stickiness_bytes_out_rate_period"`
	StickinessGpcElements        types.Int64  `tfsdk:"stickiness_gpc_elements"`
	StickinessGptElements        types.Int64  `tfsdk:"stickiness_gpt_elements"`
	StickinessGpcRatePeriod      types.String `tfsdk:"stickiness_gpc_rate_period"`
	StickinessHttpFailRatePeriod types.String `tfsdk:"stickiness_http_fail_rate_period"`
	StickinessGlitchRatePeriod   types.String `tfsdk:"stickiness_glitch_rate_period"`
	BasicAuthEnabled             types.Bool   `tfsdk:"basic_auth_enabled"`
	BasicAuthUsers               types.Set    `tfsdk:"basic_auth_users"`
	BasicAuthGroups              types.Set    `tfsdk:"basic_auth_groups"`
	TuningTimeoutConnect         types.String `tfsdk:"tuning_timeout_connect"`
	TuningTimeoutCheck           types.String `tfsdk:"tuning_timeout_check"`
	TuningTimeoutServer          types.String `tfsdk:"tuning_timeout_server"`
	TuningRetries                types.Int64  `tfsdk:"tuning_retries"`
	CustomOptions                types.String `tfsdk:"custom_options"`
	TuningDefaultserver          types.String `tfsdk:"tuning_defaultserver"`
	TuningNoport                 types.Bool   `tfsdk:"tuning_noport"`
	TuningHttpreuse              types.String `tfsdk:"tuning_httpreuse"`
	TuningCaching                types.Bool   `tfsdk:"tuning_caching"`
	LinkedActions                types.Set    `tfsdk:"linked_actions"`
	LinkedErrorfiles             types.Set    `tfsdk:"linked_errorfiles"`

	Id types.String `tfsdk:"id"`
}

func backendResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "HAProxy backend pools (**Services ▸ HAProxy ▸ Settings ▸ Virtual Services ▸ Backend Pools**, plugin `os-haproxy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this Backend Pool. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify this Backend Pool.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this Backend Pool. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Set the running mode or protocol of the Backend Pool. Usually the Public Service and the Backend Pool are in the same mode. One of `http`, `tcp`. Defaults to `http`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http"),
				Validators:          []validator.String{stringvalidator.OneOf("http", "tcp")},
			},
			"algorithm": schema.StringAttribute{
				MarkdownDescription: "Define the load balancing algorithm to be used in a Backend Pool. See the HAProxy documentation for a full description. One of `source`, `roundrobin`, `static-rr`, `leastconn`, `uri`, `random`. Defaults to `source`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("source"),
				Validators:          []validator.String{stringvalidator.OneOf("source", "roundrobin", "static-rr", "leastconn", "uri", "random")},
			},
			"random_draws": schema.Int64Attribute{
				MarkdownDescription: "When using the Random Balancing Algorithm, this value indicates the number of draws before selecting the least loaded of these servers. Defaults to `2`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(2),
			},
			"proxy_protocol": schema.StringAttribute{
				MarkdownDescription: "Enforces use of the PROXY protocol over any connection established to the configured servers. The PROXY protocol informs the other end about the layer 3/4 addresses of the incoming connection, so that it can know the client's address or the public address it accessed to, whatever the upper layer protocol. This setting must not be used if the servers are not aware of the PROXY protocol. See the HAProxy documentation for a full description. One of `v1`, `v2`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "v1", "v2")},
			},
			"linked_servers": schema.SetAttribute{
				MarkdownDescription: "Add servers to this backend. Use TAB key to complete typing. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"linked_fcgi": schema.StringAttribute{
				MarkdownDescription: "Select the FastCGI application that should be used for all servers in this backend. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"linked_resolver": schema.StringAttribute{
				MarkdownDescription: "Select the custom resolver configuration that should be used for all servers in this backend. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"resolver_opts": schema.SetAttribute{
				MarkdownDescription: "Add resolver options. Use TAB key to complete typing. One of `allow-dup-ip`, `ignore-weight`, `prevent-dup-ip`. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
				Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("allow-dup-ip", "ignore-weight", "prevent-dup-ip"))},
			},
			"resolve_prefer": schema.StringAttribute{
				MarkdownDescription: "When DNS resolution is enabled for a server and multiple IP addresses from different families are returned, HAProxy will prefer using an IP address from the selected family. One of `ipv4`, `ipv6`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "ipv4", "ipv6")},
			},
			"source": schema.StringAttribute{
				MarkdownDescription: "Sets the source address which will be used when connecting to the server(s). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"health_check_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable health checking. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"health_check": schema.StringAttribute{
				MarkdownDescription: "Select Health Monitor for servers in this backend. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"health_check_log_status": schema.BoolAttribute{
				MarkdownDescription: "Enable to log health check status updates. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"check_interval": schema.StringAttribute{
				MarkdownDescription: "Sets the interval (in milliseconds) for running health checks on all configured servers. This setting takes precedence over default values in health monitors and real servers. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"check_down_interval": schema.StringAttribute{
				MarkdownDescription: "Sets the interval (in milliseconds) for running health checks on a configured server when the server state is DOWN. If it is not set HAProxy uses the check interval. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"health_check_fall": schema.Int64Attribute{
				MarkdownDescription: "The number of consecutive unsuccessful health checks before a server is considered as unavailable. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"health_check_rise": schema.Int64Attribute{
				MarkdownDescription: "The number of consecutive successful health checks before a server is considered as available. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"linked_mailer": schema.StringAttribute{
				MarkdownDescription: "Select an e-mail alert configuration. An e-mail is sent when the state of a server changes. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"health_check_proxy_proto": schema.StringAttribute{
				MarkdownDescription: "This option controls the emission of a PROXY protocol line with outgoing health checks. The default value is good for most situations. However, if either the server or the health check does not support the PROXY protocol, this option may be useful. One of `backend`, `enable`, `disable`. Defaults to `backend`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("backend"),
				Validators:          []validator.String{stringvalidator.OneOf("", "backend", "enable", "disable")},
			},
			"http2_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable support for end-to-end HTTP/2 communication. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http2_enabled_nontls": schema.BoolAttribute{
				MarkdownDescription: "Enable support for HTTP/2 even if TLS is not enabled. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ba_advertised_protocols": schema.SetAttribute{
				MarkdownDescription: "When using the TLS ALPN extension, HAProxy advertises the specified protocol list as supported on top of ALPN. TLS must be enabled. One of `h2`, `http11`, `http10`. Defaults to `[\"h2\", \"http11\"]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(types.SetValueMust(types.StringType, []attr.Value{types.StringValue("h2"), types.StringValue("http11")})),
				Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("h2", "http11", "http10"))},
			},
			"forward_for": schema.BoolAttribute{
				MarkdownDescription: "Enable insertion of the X-Forwarded-For header to requests sent to servers. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"forwarded_header": schema.BoolAttribute{
				MarkdownDescription: "Enable insertion of the RFC7239 forwarded header in requests sent to servers. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"forwarded_header_parameters": schema.SetAttribute{
				MarkdownDescription: "This may be used to add more information to the forwarded header. Default behavior enables proto parameter and injects original client IP. See he HAProxy documentation for a full description. One of `proto`, `host`, `by`, `by_port`, `for`, `for_port`. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
				Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("proto", "host", "by", "by_port", "for", "for_port"))},
			},
			"persistence": schema.StringAttribute{
				MarkdownDescription: "Choose how HAProxy should track user-to-server mappings. Stick-table persistence works with all protocols, but is broken in multi-process and multithreaded modes. Cookie-based persistence only works with HTTP/HTTPS protocols. One of `sticktable`, `cookie`. Defaults to `sticktable`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("sticktable"),
				Validators:          []validator.String{stringvalidator.OneOf("", "sticktable", "cookie")},
			},
			"persistence_cookiemode": schema.StringAttribute{
				MarkdownDescription: "Usually it is better to reuse an existing cookie. In this case HAProxy prefixes the cookie with the required information. See the HAProxy documentation for a full description. One of `piggyback`, `new`. Defaults to `piggyback`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("piggyback"),
				Validators:          []validator.String{stringvalidator.OneOf("piggyback", "new")},
			},
			"persistence_cookiename": schema.StringAttribute{
				MarkdownDescription: "Cookie name to use for persistence. Defaults to `SRVCOOKIE`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("SRVCOOKIE"),
			},
			"persistence_stripquotes": schema.BoolAttribute{
				MarkdownDescription: "Enable to automatically strip quotes from the cookie value. Some broken HTTP clients add quotes, which breaks persistence. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"stickiness_pattern": schema.StringAttribute{
				MarkdownDescription: "Choose the type of data that should be stored in this stick-table. See the HAProxy documentation for further information. One of `binary`, `cookievalue`, `integer`, `rdpcookie`, `sourceipv4`, `sourceipv6`, `string`. Defaults to `sourceipv4`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("sourceipv4"),
				Validators:          []validator.String{stringvalidator.OneOf("", "binary", "cookievalue", "integer", "rdpcookie", "sourceipv4", "sourceipv6", "string")},
			},
			"stickiness_data_types": schema.SetAttribute{
				MarkdownDescription: "This is used to store additional information in the stick-table. It may be used by ACLs in order to control various criteria related to the activity of the client matching the stick-table. Note that this directly impacts memory usage. See the HAProxy documentation for a full description. One of `bytes_in_cnt`, `bytes_in_rate`, `bytes_out_cnt`, `bytes_out_rate`, `conn_cnt`, `conn_cur`, `conn_rate`, `glitch_cnt`, `glitch_rate`, `gpc`, `gpc_rate`, `gpc0`, `gpc0_rate`, `gpc1`, `gpc1_rate`, `gpt`, `gpt0`, `http_err_cnt`, `http_err_rate`, `http_fail_cnt`, `http_fail_rate`, `http_req_cnt`, `http_req_rate`, `server_id`, `sess_cnt`, `sess_rate`. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
				Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("bytes_in_cnt", "bytes_in_rate", "bytes_out_cnt", "bytes_out_rate", "conn_cnt", "conn_cur", "conn_rate", "glitch_cnt", "glitch_rate", "gpc", "gpc_rate", "gpc0", "gpc0_rate", "gpc1", "gpc1_rate", "gpt", "gpt0", "http_err_cnt", "http_err_rate", "http_fail_cnt", "http_fail_rate", "http_req_cnt", "http_req_rate", "server_id", "sess_cnt", "sess_rate"))},
			},
			"stickiness_expire": schema.StringAttribute{
				MarkdownDescription: "Enter a number followed by one of the supported suffixes \"d\" (days), \"h\" (hour), \"m\" (minute), \"s\" (seconds), \"ms\" (miliseconds). This configures the maximum duration of an entry in the stick-table since it was last created, refreshed or matched. The maximum duration is slightly above 24 days. Defaults to `30m`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("30m"),
			},
			"stickiness_size": schema.StringAttribute{
				MarkdownDescription: "Enter a number followed by one of the supported suffixes \"k\", \"m\", \"g\". This configures the maximum number of entries that can fit in the table. This value directly impacts memory usage. Count approximately 50 bytes per entry, plus the size of a string if any. Defaults to `50k`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("50k"),
			},
			"stickiness_cookiename": schema.StringAttribute{
				MarkdownDescription: "Cookie name to use for stick table (if appropiate table type is selected). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"stickiness_cookielength": schema.Int64Attribute{
				MarkdownDescription: "The maximum number of characters that will be stored in the stick table (if appropiate table type is selected). Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"stickiness_length": schema.Int64Attribute{
				MarkdownDescription: "Specify the maximum length for a value in the stick-table. If the value is larger than this value it will be truncated before being stored. Depending on the stick-table type this repesents either characters or bytes. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"stickiness_conn_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average incoming connection rate over that period, in connections per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\". Defaults to `10s`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("10s"),
			},
			"stickiness_sess_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average incoming session rate over that period, in sessions per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\". Defaults to `10s`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("10s"),
			},
			"stickiness_http_req_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average HTTP request rate over that period, in requests per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\". Defaults to `10s`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("10s"),
			},
			"stickiness_http_err_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average HTTP request error rate over that period, in requests per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\". Defaults to `10s`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("10s"),
			},
			"stickiness_bytes_in_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average incoming bytes rate over that period, in bytes per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\". Defaults to `1m`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("1m"),
			},
			"stickiness_bytes_out_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average outgoing bytes rate over that period, in bytes per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\". Defaults to `1m`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("1m"),
			},
			"stickiness_gpc_elements": schema.Int64Attribute{
				MarkdownDescription: "The number of array elements that will be used for the General Purpose Counter. This is limited to a maximum of 100 elements, ranging from 0 to 99. Defaults to `0`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
			},
			"stickiness_gpt_elements": schema.Int64Attribute{
				MarkdownDescription: "The number of array elements that will be used for the General Purpose Tag. This is limited to a maximum of 100 elements, ranging from 0 to 99. Defaults to `0`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
			},
			"stickiness_gpc_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. Most of the time it will be used to measure the frequency of occurrence of certain events. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\". Defaults to `1m`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("1m"),
			},
			"stickiness_http_fail_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average HTTP response fail error rate over that period, in requests per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\". Defaults to `1m`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("1m"),
			},
			"stickiness_glitch_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average front glitches rate over that period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\". Defaults to `1m`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("1m"),
			},
			"basic_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable HTTP Basic Authentication. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"basic_auth_users": schema.SetAttribute{
				MarkdownDescription: "Allowed Users. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"basic_auth_groups": schema.SetAttribute{
				MarkdownDescription: "Allowed Groups. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"tuning_timeout_connect": schema.StringAttribute{
				MarkdownDescription: "Set the maximum time to wait for a connection attempt to a server to succeed. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_timeout_check": schema.StringAttribute{
				MarkdownDescription: "Sets an additional read timeout for running health checks on a server. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_timeout_server": schema.StringAttribute{
				MarkdownDescription: "Set the maximum inactivity time on the server side. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_retries": schema.Int64Attribute{
				MarkdownDescription: "Set the number of retries to perform on a server after a connection failure. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"custom_options": schema.StringAttribute{
				MarkdownDescription: "These lines will be added to the HAProxy backend configuration. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_defaultserver": schema.StringAttribute{
				MarkdownDescription: "Default option for all server entries. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_noport": schema.BoolAttribute{
				MarkdownDescription: "Don't use port on server, use the same port as frontend receive. If check enable, require port check in server. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tuning_httpreuse": schema.StringAttribute{
				MarkdownDescription: "Declare how idle HTTP connections may be shared between requests. One of `never`, `safe`, `aggressive`, `always`. Defaults to `safe`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("safe"),
				Validators:          []validator.String{stringvalidator.OneOf("", "never", "safe", "aggressive", "always")},
			},
			"tuning_caching": schema.BoolAttribute{
				MarkdownDescription: "Enable caching of responses from this backend. The HAProxy cache must be enabled under Settings before this will have any effect. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"linked_actions": schema.SetAttribute{
				MarkdownDescription: "Choose rules to be included in this Backend Pool. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"linked_errorfiles": schema.SetAttribute{
				MarkdownDescription: "Choose error messages to be included in this Backend Pool. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the backend pool.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func backendDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "HAProxy backend pools (**Services ▸ HAProxy ▸ Settings ▸ Virtual Services ▸ Backend Pools**, plugin `os-haproxy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the backend pool.",
				Required:            true,
			},
			"enabled":                          dschema.BoolAttribute{MarkdownDescription: "Enable this Backend Pool.", Computed: true},
			"name":                             dschema.StringAttribute{MarkdownDescription: "Name to identify this Backend Pool.", Computed: true},
			"description":                      dschema.StringAttribute{MarkdownDescription: "Description for this Backend Pool.", Computed: true},
			"mode":                             dschema.StringAttribute{MarkdownDescription: "Set the running mode or protocol of the Backend Pool. Usually the Public Service and the Backend Pool are in the same mode. One of `http`, `tcp`.", Computed: true},
			"algorithm":                        dschema.StringAttribute{MarkdownDescription: "Define the load balancing algorithm to be used in a Backend Pool. See the HAProxy documentation for a full description. One of `source`, `roundrobin`, `static-rr`, `leastconn`, `uri`, `random`.", Computed: true},
			"random_draws":                     dschema.Int64Attribute{MarkdownDescription: "When using the Random Balancing Algorithm, this value indicates the number of draws before selecting the least loaded of these servers.", Computed: true},
			"proxy_protocol":                   dschema.StringAttribute{MarkdownDescription: "Enforces use of the PROXY protocol over any connection established to the configured servers. The PROXY protocol informs the other end about the layer 3/4 addresses of the incoming connection, so that it can know the client's address or the public address it accessed to, whatever the upper layer protocol. This setting must not be used if the servers are not aware of the PROXY protocol. See the HAProxy documentation for a full description. One of `v1`, `v2`.", Computed: true},
			"linked_servers":                   dschema.SetAttribute{MarkdownDescription: "Add servers to this backend. Use TAB key to complete typing.", Computed: true, ElementType: types.StringType},
			"linked_fcgi":                      dschema.StringAttribute{MarkdownDescription: "Select the FastCGI application that should be used for all servers in this backend.", Computed: true},
			"linked_resolver":                  dschema.StringAttribute{MarkdownDescription: "Select the custom resolver configuration that should be used for all servers in this backend.", Computed: true},
			"resolver_opts":                    dschema.SetAttribute{MarkdownDescription: "Add resolver options. Use TAB key to complete typing. One of `allow-dup-ip`, `ignore-weight`, `prevent-dup-ip`.", Computed: true, ElementType: types.StringType},
			"resolve_prefer":                   dschema.StringAttribute{MarkdownDescription: "When DNS resolution is enabled for a server and multiple IP addresses from different families are returned, HAProxy will prefer using an IP address from the selected family. One of `ipv4`, `ipv6`.", Computed: true},
			"source":                           dschema.StringAttribute{MarkdownDescription: "Sets the source address which will be used when connecting to the server(s).", Computed: true},
			"health_check_enabled":             dschema.BoolAttribute{MarkdownDescription: "Enable or disable health checking.", Computed: true},
			"health_check":                     dschema.StringAttribute{MarkdownDescription: "Select Health Monitor for servers in this backend.", Computed: true},
			"health_check_log_status":          dschema.BoolAttribute{MarkdownDescription: "Enable to log health check status updates.", Computed: true},
			"check_interval":                   dschema.StringAttribute{MarkdownDescription: "Sets the interval (in milliseconds) for running health checks on all configured servers. This setting takes precedence over default values in health monitors and real servers.", Computed: true},
			"check_down_interval":              dschema.StringAttribute{MarkdownDescription: "Sets the interval (in milliseconds) for running health checks on a configured server when the server state is DOWN. If it is not set HAProxy uses the check interval.", Computed: true},
			"health_check_fall":                dschema.Int64Attribute{MarkdownDescription: "The number of consecutive unsuccessful health checks before a server is considered as unavailable.", Computed: true},
			"health_check_rise":                dschema.Int64Attribute{MarkdownDescription: "The number of consecutive successful health checks before a server is considered as available.", Computed: true},
			"linked_mailer":                    dschema.StringAttribute{MarkdownDescription: "Select an e-mail alert configuration. An e-mail is sent when the state of a server changes.", Computed: true},
			"health_check_proxy_proto":         dschema.StringAttribute{MarkdownDescription: "This option controls the emission of a PROXY protocol line with outgoing health checks. The default value is good for most situations. However, if either the server or the health check does not support the PROXY protocol, this option may be useful. One of `backend`, `enable`, `disable`.", Computed: true},
			"http2_enabled":                    dschema.BoolAttribute{MarkdownDescription: "Enable support for end-to-end HTTP/2 communication.", Computed: true},
			"http2_enabled_nontls":             dschema.BoolAttribute{MarkdownDescription: "Enable support for HTTP/2 even if TLS is not enabled.", Computed: true},
			"ba_advertised_protocols":          dschema.SetAttribute{MarkdownDescription: "When using the TLS ALPN extension, HAProxy advertises the specified protocol list as supported on top of ALPN. TLS must be enabled. One of `h2`, `http11`, `http10`.", Computed: true, ElementType: types.StringType},
			"forward_for":                      dschema.BoolAttribute{MarkdownDescription: "Enable insertion of the X-Forwarded-For header to requests sent to servers.", Computed: true},
			"forwarded_header":                 dschema.BoolAttribute{MarkdownDescription: "Enable insertion of the RFC7239 forwarded header in requests sent to servers.", Computed: true},
			"forwarded_header_parameters":      dschema.SetAttribute{MarkdownDescription: "This may be used to add more information to the forwarded header. Default behavior enables proto parameter and injects original client IP. See he HAProxy documentation for a full description. One of `proto`, `host`, `by`, `by_port`, `for`, `for_port`.", Computed: true, ElementType: types.StringType},
			"persistence":                      dschema.StringAttribute{MarkdownDescription: "Choose how HAProxy should track user-to-server mappings. Stick-table persistence works with all protocols, but is broken in multi-process and multithreaded modes. Cookie-based persistence only works with HTTP/HTTPS protocols. One of `sticktable`, `cookie`.", Computed: true},
			"persistence_cookiemode":           dschema.StringAttribute{MarkdownDescription: "Usually it is better to reuse an existing cookie. In this case HAProxy prefixes the cookie with the required information. See the HAProxy documentation for a full description. One of `piggyback`, `new`.", Computed: true},
			"persistence_cookiename":           dschema.StringAttribute{MarkdownDescription: "Cookie name to use for persistence.", Computed: true},
			"persistence_stripquotes":          dschema.BoolAttribute{MarkdownDescription: "Enable to automatically strip quotes from the cookie value. Some broken HTTP clients add quotes, which breaks persistence.", Computed: true},
			"stickiness_pattern":               dschema.StringAttribute{MarkdownDescription: "Choose the type of data that should be stored in this stick-table. See the HAProxy documentation for further information. One of `binary`, `cookievalue`, `integer`, `rdpcookie`, `sourceipv4`, `sourceipv6`, `string`.", Computed: true},
			"stickiness_data_types":            dschema.SetAttribute{MarkdownDescription: "This is used to store additional information in the stick-table. It may be used by ACLs in order to control various criteria related to the activity of the client matching the stick-table. Note that this directly impacts memory usage. See the HAProxy documentation for a full description. One of `bytes_in_cnt`, `bytes_in_rate`, `bytes_out_cnt`, `bytes_out_rate`, `conn_cnt`, `conn_cur`, `conn_rate`, `glitch_cnt`, `glitch_rate`, `gpc`, `gpc_rate`, `gpc0`, `gpc0_rate`, `gpc1`, `gpc1_rate`, `gpt`, `gpt0`, `http_err_cnt`, `http_err_rate`, `http_fail_cnt`, `http_fail_rate`, `http_req_cnt`, `http_req_rate`, `server_id`, `sess_cnt`, `sess_rate`.", Computed: true, ElementType: types.StringType},
			"stickiness_expire":                dschema.StringAttribute{MarkdownDescription: "Enter a number followed by one of the supported suffixes \"d\" (days), \"h\" (hour), \"m\" (minute), \"s\" (seconds), \"ms\" (miliseconds). This configures the maximum duration of an entry in the stick-table since it was last created, refreshed or matched. The maximum duration is slightly above 24 days.", Computed: true},
			"stickiness_size":                  dschema.StringAttribute{MarkdownDescription: "Enter a number followed by one of the supported suffixes \"k\", \"m\", \"g\". This configures the maximum number of entries that can fit in the table. This value directly impacts memory usage. Count approximately 50 bytes per entry, plus the size of a string if any.", Computed: true},
			"stickiness_cookiename":            dschema.StringAttribute{MarkdownDescription: "Cookie name to use for stick table (if appropiate table type is selected).", Computed: true},
			"stickiness_cookielength":          dschema.Int64Attribute{MarkdownDescription: "The maximum number of characters that will be stored in the stick table (if appropiate table type is selected).", Computed: true},
			"stickiness_length":                dschema.Int64Attribute{MarkdownDescription: "Specify the maximum length for a value in the stick-table. If the value is larger than this value it will be truncated before being stored. Depending on the stick-table type this repesents either characters or bytes.", Computed: true},
			"stickiness_conn_rate_period":      dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. It reports the average incoming connection rate over that period, in connections per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"stickiness_sess_rate_period":      dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. It reports the average incoming session rate over that period, in sessions per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"stickiness_http_req_rate_period":  dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. It reports the average HTTP request rate over that period, in requests per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"stickiness_http_err_rate_period":  dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. It reports the average HTTP request error rate over that period, in requests per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"stickiness_bytes_in_rate_period":  dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. It reports the average incoming bytes rate over that period, in bytes per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"stickiness_bytes_out_rate_period": dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. It reports the average outgoing bytes rate over that period, in bytes per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"stickiness_gpc_elements":          dschema.Int64Attribute{MarkdownDescription: "The number of array elements that will be used for the General Purpose Counter. This is limited to a maximum of 100 elements, ranging from 0 to 99.", Computed: true},
			"stickiness_gpt_elements":          dschema.Int64Attribute{MarkdownDescription: "The number of array elements that will be used for the General Purpose Tag. This is limited to a maximum of 100 elements, ranging from 0 to 99.", Computed: true},
			"stickiness_gpc_rate_period":       dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. Most of the time it will be used to measure the frequency of occurrence of certain events. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"stickiness_http_fail_rate_period": dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. It reports the average HTTP response fail error rate over that period, in requests per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"stickiness_glitch_rate_period":    dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. It reports the average front glitches rate over that period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"basic_auth_enabled":               dschema.BoolAttribute{MarkdownDescription: "Enable HTTP Basic Authentication.", Computed: true},
			"basic_auth_users":                 dschema.SetAttribute{MarkdownDescription: "Allowed Users.", Computed: true, ElementType: types.StringType},
			"basic_auth_groups":                dschema.SetAttribute{MarkdownDescription: "Allowed Groups.", Computed: true, ElementType: types.StringType},
			"tuning_timeout_connect":           dschema.StringAttribute{MarkdownDescription: "Set the maximum time to wait for a connection attempt to a server to succeed. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"tuning_timeout_check":             dschema.StringAttribute{MarkdownDescription: "Sets an additional read timeout for running health checks on a server. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"tuning_timeout_server":            dschema.StringAttribute{MarkdownDescription: "Set the maximum inactivity time on the server side. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"tuning_retries":                   dschema.Int64Attribute{MarkdownDescription: "Set the number of retries to perform on a server after a connection failure.", Computed: true},
			"custom_options":                   dschema.StringAttribute{MarkdownDescription: "These lines will be added to the HAProxy backend configuration.", Computed: true},
			"tuning_defaultserver":             dschema.StringAttribute{MarkdownDescription: "Default option for all server entries.", Computed: true},
			"tuning_noport":                    dschema.BoolAttribute{MarkdownDescription: "Don't use port on server, use the same port as frontend receive. If check enable, require port check in server.", Computed: true},
			"tuning_httpreuse":                 dschema.StringAttribute{MarkdownDescription: "Declare how idle HTTP connections may be shared between requests. One of `never`, `safe`, `aggressive`, `always`.", Computed: true},
			"tuning_caching":                   dschema.BoolAttribute{MarkdownDescription: "Enable caching of responses from this backend. The HAProxy cache must be enabled under Settings before this will have any effect.", Computed: true},
			"linked_actions":                   dschema.SetAttribute{MarkdownDescription: "Choose rules to be included in this Backend Pool.", Computed: true, ElementType: types.StringType},
			"linked_errorfiles":                dschema.SetAttribute{MarkdownDescription: "Choose error messages to be included in this Backend Pool.", Computed: true, ElementType: types.StringType},
		},
	}
}

func convertBackendSchemaToStruct(d *backendResourceModel) (*haproxy.Backend, error) {
	return &haproxy.Backend{
		Enabled:                      tools.BoolToString(d.Enabled.ValueBool()),
		Name:                         d.Name.ValueString(),
		Description:                  d.Description.ValueString(),
		Mode:                         api.SelectedMap(d.Mode.ValueString()),
		Algorithm:                    api.SelectedMap(d.Algorithm.ValueString()),
		RandomDraws:                  tools.Int64ToString(d.RandomDraws.ValueInt64()),
		ProxyProtocol:                api.SelectedMap(d.ProxyProtocol.ValueString()),
		LinkedServers:                api.SelectedMapList(setToSlice(d.LinkedServers)),
		LinkedFcgi:                   api.SelectedMap(d.LinkedFcgi.ValueString()),
		LinkedResolver:               api.SelectedMap(d.LinkedResolver.ValueString()),
		ResolverOpts:                 api.SelectedMapList(setToSlice(d.ResolverOpts)),
		ResolvePrefer:                api.SelectedMap(d.ResolvePrefer.ValueString()),
		Source:                       d.Source.ValueString(),
		HealthCheckEnabled:           tools.BoolToString(d.HealthCheckEnabled.ValueBool()),
		HealthCheck:                  api.SelectedMap(d.HealthCheck.ValueString()),
		HealthCheckLogStatus:         tools.BoolToString(d.HealthCheckLogStatus.ValueBool()),
		CheckInterval:                d.CheckInterval.ValueString(),
		CheckDownInterval:            d.CheckDownInterval.ValueString(),
		HealthCheckFall:              tools.Int64ToStringNegative(d.HealthCheckFall.ValueInt64()),
		HealthCheckRise:              tools.Int64ToStringNegative(d.HealthCheckRise.ValueInt64()),
		LinkedMailer:                 api.SelectedMap(d.LinkedMailer.ValueString()),
		HealthCheckProxyProto:        api.SelectedMap(d.HealthCheckProxyProto.ValueString()),
		Http2Enabled:                 tools.BoolToString(d.Http2Enabled.ValueBool()),
		Http2EnabledNontls:           tools.BoolToString(d.Http2EnabledNontls.ValueBool()),
		BaAdvertisedProtocols:        api.SelectedMapList(setToSlice(d.BaAdvertisedProtocols)),
		ForwardFor:                   tools.BoolToString(d.ForwardFor.ValueBool()),
		ForwardedHeader:              tools.BoolToString(d.ForwardedHeader.ValueBool()),
		ForwardedHeaderParameters:    api.SelectedMapList(setToSlice(d.ForwardedHeaderParameters)),
		Persistence:                  api.SelectedMap(d.Persistence.ValueString()),
		PersistenceCookiemode:        api.SelectedMap(d.PersistenceCookiemode.ValueString()),
		PersistenceCookiename:        d.PersistenceCookiename.ValueString(),
		PersistenceStripquotes:       tools.BoolToString(d.PersistenceStripquotes.ValueBool()),
		StickinessPattern:            api.SelectedMap(d.StickinessPattern.ValueString()),
		StickinessDataTypes:          api.SelectedMapList(setToSlice(d.StickinessDataTypes)),
		StickinessExpire:             d.StickinessExpire.ValueString(),
		StickinessSize:               d.StickinessSize.ValueString(),
		StickinessCookiename:         d.StickinessCookiename.ValueString(),
		StickinessCookielength:       tools.Int64ToStringNegative(d.StickinessCookielength.ValueInt64()),
		StickinessLength:             tools.Int64ToStringNegative(d.StickinessLength.ValueInt64()),
		StickinessConnRatePeriod:     d.StickinessConnRatePeriod.ValueString(),
		StickinessSessRatePeriod:     d.StickinessSessRatePeriod.ValueString(),
		StickinessHttpReqRatePeriod:  d.StickinessHttpReqRatePeriod.ValueString(),
		StickinessHttpErrRatePeriod:  d.StickinessHttpErrRatePeriod.ValueString(),
		StickinessBytesInRatePeriod:  d.StickinessBytesInRatePeriod.ValueString(),
		StickinessBytesOutRatePeriod: d.StickinessBytesOutRatePeriod.ValueString(),
		StickinessGpcElements:        tools.Int64ToString(d.StickinessGpcElements.ValueInt64()),
		StickinessGptElements:        tools.Int64ToString(d.StickinessGptElements.ValueInt64()),
		StickinessGpcRatePeriod:      d.StickinessGpcRatePeriod.ValueString(),
		StickinessHttpFailRatePeriod: d.StickinessHttpFailRatePeriod.ValueString(),
		StickinessGlitchRatePeriod:   d.StickinessGlitchRatePeriod.ValueString(),
		BasicAuthEnabled:             tools.BoolToString(d.BasicAuthEnabled.ValueBool()),
		BasicAuthUsers:               api.SelectedMapList(setToSlice(d.BasicAuthUsers)),
		BasicAuthGroups:              api.SelectedMapList(setToSlice(d.BasicAuthGroups)),
		TuningTimeoutConnect:         d.TuningTimeoutConnect.ValueString(),
		TuningTimeoutCheck:           d.TuningTimeoutCheck.ValueString(),
		TuningTimeoutServer:          d.TuningTimeoutServer.ValueString(),
		TuningRetries:                tools.Int64ToStringNegative(d.TuningRetries.ValueInt64()),
		CustomOptions:                d.CustomOptions.ValueString(),
		TuningDefaultserver:          d.TuningDefaultserver.ValueString(),
		TuningNoport:                 tools.BoolToString(d.TuningNoport.ValueBool()),
		TuningHttpreuse:              api.SelectedMap(d.TuningHttpreuse.ValueString()),
		TuningCaching:                tools.BoolToString(d.TuningCaching.ValueBool()),
		LinkedActions:                api.SelectedMapList(setToSlice(d.LinkedActions)),
		LinkedErrorfiles:             api.SelectedMapList(setToSlice(d.LinkedErrorfiles)),
	}, nil
}

func convertBackendStructToSchema(d *haproxy.Backend) (*backendResourceModel, error) {
	return &backendResourceModel{
		Enabled:                      types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:                         types.StringValue(d.Name),
		Description:                  types.StringValue(d.Description),
		Mode:                         types.StringValue(d.Mode.String()),
		Algorithm:                    types.StringValue(d.Algorithm.String()),
		RandomDraws:                  types.Int64Value(tools.StringToInt64(d.RandomDraws)),
		ProxyProtocol:                types.StringValue(d.ProxyProtocol.String()),
		LinkedServers:                tools.StringSliceToSet(d.LinkedServers),
		LinkedFcgi:                   types.StringValue(d.LinkedFcgi.String()),
		LinkedResolver:               types.StringValue(d.LinkedResolver.String()),
		ResolverOpts:                 tools.StringSliceToSet(d.ResolverOpts),
		ResolvePrefer:                types.StringValue(d.ResolvePrefer.String()),
		Source:                       types.StringValue(d.Source),
		HealthCheckEnabled:           types.BoolValue(tools.StringToBool(d.HealthCheckEnabled)),
		HealthCheck:                  types.StringValue(d.HealthCheck.String()),
		HealthCheckLogStatus:         types.BoolValue(tools.StringToBool(d.HealthCheckLogStatus)),
		CheckInterval:                types.StringValue(d.CheckInterval),
		CheckDownInterval:            types.StringValue(d.CheckDownInterval),
		HealthCheckFall:              types.Int64Value(tools.StringToInt64(d.HealthCheckFall)),
		HealthCheckRise:              types.Int64Value(tools.StringToInt64(d.HealthCheckRise)),
		LinkedMailer:                 types.StringValue(d.LinkedMailer.String()),
		HealthCheckProxyProto:        types.StringValue(d.HealthCheckProxyProto.String()),
		Http2Enabled:                 types.BoolValue(tools.StringToBool(d.Http2Enabled)),
		Http2EnabledNontls:           types.BoolValue(tools.StringToBool(d.Http2EnabledNontls)),
		BaAdvertisedProtocols:        tools.StringSliceToSet(d.BaAdvertisedProtocols),
		ForwardFor:                   types.BoolValue(tools.StringToBool(d.ForwardFor)),
		ForwardedHeader:              types.BoolValue(tools.StringToBool(d.ForwardedHeader)),
		ForwardedHeaderParameters:    tools.StringSliceToSet(d.ForwardedHeaderParameters),
		Persistence:                  types.StringValue(d.Persistence.String()),
		PersistenceCookiemode:        types.StringValue(d.PersistenceCookiemode.String()),
		PersistenceCookiename:        types.StringValue(d.PersistenceCookiename),
		PersistenceStripquotes:       types.BoolValue(tools.StringToBool(d.PersistenceStripquotes)),
		StickinessPattern:            types.StringValue(d.StickinessPattern.String()),
		StickinessDataTypes:          tools.StringSliceToSet(d.StickinessDataTypes),
		StickinessExpire:             types.StringValue(d.StickinessExpire),
		StickinessSize:               types.StringValue(d.StickinessSize),
		StickinessCookiename:         types.StringValue(d.StickinessCookiename),
		StickinessCookielength:       types.Int64Value(tools.StringToInt64(d.StickinessCookielength)),
		StickinessLength:             types.Int64Value(tools.StringToInt64(d.StickinessLength)),
		StickinessConnRatePeriod:     types.StringValue(d.StickinessConnRatePeriod),
		StickinessSessRatePeriod:     types.StringValue(d.StickinessSessRatePeriod),
		StickinessHttpReqRatePeriod:  types.StringValue(d.StickinessHttpReqRatePeriod),
		StickinessHttpErrRatePeriod:  types.StringValue(d.StickinessHttpErrRatePeriod),
		StickinessBytesInRatePeriod:  types.StringValue(d.StickinessBytesInRatePeriod),
		StickinessBytesOutRatePeriod: types.StringValue(d.StickinessBytesOutRatePeriod),
		StickinessGpcElements:        types.Int64Value(tools.StringToInt64(d.StickinessGpcElements)),
		StickinessGptElements:        types.Int64Value(tools.StringToInt64(d.StickinessGptElements)),
		StickinessGpcRatePeriod:      types.StringValue(d.StickinessGpcRatePeriod),
		StickinessHttpFailRatePeriod: types.StringValue(d.StickinessHttpFailRatePeriod),
		StickinessGlitchRatePeriod:   types.StringValue(d.StickinessGlitchRatePeriod),
		BasicAuthEnabled:             types.BoolValue(tools.StringToBool(d.BasicAuthEnabled)),
		BasicAuthUsers:               tools.StringSliceToSet(d.BasicAuthUsers),
		BasicAuthGroups:              tools.StringSliceToSet(d.BasicAuthGroups),
		TuningTimeoutConnect:         types.StringValue(d.TuningTimeoutConnect),
		TuningTimeoutCheck:           types.StringValue(d.TuningTimeoutCheck),
		TuningTimeoutServer:          types.StringValue(d.TuningTimeoutServer),
		TuningRetries:                types.Int64Value(tools.StringToInt64(d.TuningRetries)),
		CustomOptions:                types.StringValue(d.CustomOptions),
		TuningDefaultserver:          types.StringValue(d.TuningDefaultserver),
		TuningNoport:                 types.BoolValue(tools.StringToBool(d.TuningNoport)),
		TuningHttpreuse:              types.StringValue(d.TuningHttpreuse.String()),
		TuningCaching:                types.BoolValue(tools.StringToBool(d.TuningCaching)),
		LinkedActions:                tools.StringSliceToSet(d.LinkedActions),
		LinkedErrorfiles:             tools.StringSliceToSet(d.LinkedErrorfiles),
	}, nil
}
