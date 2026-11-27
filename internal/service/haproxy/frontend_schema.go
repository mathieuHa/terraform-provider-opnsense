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

type frontendResourceModel struct {
	Enabled                      types.Bool   `tfsdk:"enabled"`
	Name                         types.String `tfsdk:"name"`
	Description                  types.String `tfsdk:"description"`
	Bind                         types.Set    `tfsdk:"bind"`
	BindOptions                  types.String `tfsdk:"bind_options"`
	Mode                         types.String `tfsdk:"mode"`
	DefaultBackend               types.String `tfsdk:"default_backend"`
	SslEnabled                   types.Bool   `tfsdk:"ssl_enabled"`
	SslCertificates              types.Set    `tfsdk:"ssl_certificates"`
	SslDefaultCertificate        types.String `tfsdk:"ssl_default_certificate"`
	SslCustomOptions             types.String `tfsdk:"ssl_custom_options"`
	SslAdvancedEnabled           types.Bool   `tfsdk:"ssl_advanced_enabled"`
	SslBindOptions               types.Set    `tfsdk:"ssl_bind_options"`
	SslMinVersion                types.String `tfsdk:"ssl_min_version"`
	SslMaxVersion                types.String `tfsdk:"ssl_max_version"`
	SslCipherList                types.String `tfsdk:"ssl_cipher_list"`
	SslCipherSuites              types.String `tfsdk:"ssl_cipher_suites"`
	SslHstsEnabled               types.Bool   `tfsdk:"ssl_hsts_enabled"`
	SslHstsIncludeSubDomains     types.Bool   `tfsdk:"ssl_hsts_include_sub_domains"`
	SslHstsPreload               types.Bool   `tfsdk:"ssl_hsts_preload"`
	SslHstsMaxAge                types.Int64  `tfsdk:"ssl_hsts_max_age"`
	SslClientAuthEnabled         types.Bool   `tfsdk:"ssl_client_auth_enabled"`
	SslClientAuthVerify          types.String `tfsdk:"ssl_client_auth_verify"`
	SslClientAuthCas             types.Set    `tfsdk:"ssl_client_auth_cas"`
	SslClientAuthCrls            types.Set    `tfsdk:"ssl_client_auth_crls"`
	BasicAuthEnabled             types.Bool   `tfsdk:"basic_auth_enabled"`
	BasicAuthUsers               types.Set    `tfsdk:"basic_auth_users"`
	BasicAuthGroups              types.Set    `tfsdk:"basic_auth_groups"`
	TuningMaxConnections         types.Int64  `tfsdk:"tuning_max_connections"`
	TuningTimeoutClient          types.String `tfsdk:"tuning_timeout_client"`
	TuningTimeoutHttpReq         types.String `tfsdk:"tuning_timeout_http_req"`
	TuningTimeoutHttpKeepAlive   types.String `tfsdk:"tuning_timeout_http_keep_alive"`
	LinkedCpuAffinityRules       types.Set    `tfsdk:"linked_cpu_affinity_rules"`
	TuningShards                 types.Int64  `tfsdk:"tuning_shards"`
	LoggingDontLogNull           types.Bool   `tfsdk:"logging_dont_log_null"`
	LoggingDontLogNormal         types.Bool   `tfsdk:"logging_dont_log_normal"`
	LoggingLogSeparateErrors     types.Bool   `tfsdk:"logging_log_separate_errors"`
	LoggingDetailedLog           types.Bool   `tfsdk:"logging_detailed_log"`
	LoggingSocketStats           types.Bool   `tfsdk:"logging_socket_stats"`
	StickinessPattern            types.String `tfsdk:"stickiness_pattern"`
	StickinessDataTypes          types.Set    `tfsdk:"stickiness_data_types"`
	StickinessExpire             types.String `tfsdk:"stickiness_expire"`
	StickinessSize               types.String `tfsdk:"stickiness_size"`
	StickinessCounter            types.Bool   `tfsdk:"stickiness_counter"`
	StickinessCounterKey         types.String `tfsdk:"stickiness_counter_key"`
	StickinessLength             types.Int64  `tfsdk:"stickiness_length"`
	StickinessConnRatePeriod     types.String `tfsdk:"stickiness_conn_rate_period"`
	StickinessSessRatePeriod     types.String `tfsdk:"stickiness_sess_rate_period"`
	StickinessHttpReqRatePeriod  types.String `tfsdk:"stickiness_http_req_rate_period"`
	StickinessHttpErrRatePeriod  types.String `tfsdk:"stickiness_http_err_rate_period"`
	StickinessBytesInRatePeriod  types.String `tfsdk:"stickiness_bytes_in_rate_period"`
	StickinessBytesOutRatePeriod types.String `tfsdk:"stickiness_bytes_out_rate_period"`
	StickinessGpcElements        types.Int64  `tfsdk:"stickiness_gpc_elements"`
	StickinessGpcRatePeriod      types.String `tfsdk:"stickiness_gpc_rate_period"`
	StickinessGptElements        types.Int64  `tfsdk:"stickiness_gpt_elements"`
	StickinessHttpFailRatePeriod types.String `tfsdk:"stickiness_http_fail_rate_period"`
	StickinessGlitchRatePeriod   types.String `tfsdk:"stickiness_glitch_rate_period"`
	Http2Enabled                 types.Bool   `tfsdk:"http2_enabled"`
	Http2EnabledNontls           types.Bool   `tfsdk:"http2_enabled_nontls"`
	AdvertisedProtocols          types.Set    `tfsdk:"advertised_protocols"`
	ForwardFor                   types.Bool   `tfsdk:"forward_for"`
	PrometheusEnabled            types.Bool   `tfsdk:"prometheus_enabled"`
	PrometheusPath               types.String `tfsdk:"prometheus_path"`
	ConnectionBehaviour          types.String `tfsdk:"connection_behaviour"`
	CustomOptions                types.String `tfsdk:"custom_options"`
	LinkedActions                types.Set    `tfsdk:"linked_actions"`
	LinkedErrorfiles             types.Set    `tfsdk:"linked_errorfiles"`

	Id types.String `tfsdk:"id"`
}

func frontendResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "HAProxy frontends (**Services ▸ HAProxy ▸ Settings ▸ Virtual Services ▸ Public Services**, plugin `os-haproxy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this Public Service. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify this Public Service.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this Public Service. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"bind": schema.SetAttribute{
				MarkdownDescription: "Configure listen addresses for this Public Service, i.e. 127.0.0.1:8080 or www.example.com:443 or unix@socket-name. Add quic4@ or quic6@ to enable QUIC for IPv4/IPv6, e.g. quic4@www.example.com:443. Use TAB key to complete typing a listen address.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"bind_options": schema.StringAttribute{
				MarkdownDescription: "A list of parameters that will be appended to every Listen Address line.Example: accept-proxy npn http/1.1. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Set the running mode or protocol for this Public Service. One of `http`, `ssl`, `tcp`. Defaults to `http`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http"),
				Validators:          []validator.String{stringvalidator.OneOf("http", "ssl", "tcp")},
			},
			"default_backend": schema.StringAttribute{
				MarkdownDescription: "Set the default Backend Pool to use for this Public Service. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable SSL offloading. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_certificates": schema.SetAttribute{
				MarkdownDescription: "Select certificates to use for SSL offloading. HAProxy's SNI recognition will determine the correct certificate automatically. If no SNI is provided by the client then the first certificate will be presented.To import additional certificates, go to Certificate Manager. Takes the `ref_id` of an `opnsense_trust_cert`, not its UUID. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"ssl_default_certificate": schema.StringAttribute{
				MarkdownDescription: "This certificate will be presented if no SNI is provided by the client or if the client provides an SNI hostname which does not match any certificate. This parameter is optional to enforce a certain sort order for certificates. Takes the `ref_id` of an `opnsense_trust_cert`, not its UUID. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_custom_options": schema.StringAttribute{
				MarkdownDescription: "Pass additional SSL parameters to the HAProxy configuration. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_advanced_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable advanced SSL settings. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_bind_options": schema.SetAttribute{
				MarkdownDescription: "Used to enforce or disable certain SSL options. One of `no-sslv3`, `no-tlsv10`, `no-tlsv11`, `no-tlsv12`, `no-tlsv13`, `no-tls-tickets`, `force-sslv3`, `force-tlsv10`, `force-tlsv11`, `force-tlsv12`, `force-tlsv13`, `prefer-client-ciphers`, `strict-sni`. Defaults to `[\"prefer-client-ciphers\"]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(types.SetValueMust(types.StringType, []attr.Value{types.StringValue("prefer-client-ciphers")})),
				Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("no-sslv3", "no-tlsv10", "no-tlsv11", "no-tlsv12", "no-tlsv13", "no-tls-tickets", "force-sslv3", "force-tlsv10", "force-tlsv11", "force-tlsv12", "force-tlsv13", "prefer-client-ciphers", "strict-sni"))},
			},
			"ssl_min_version": schema.StringAttribute{
				MarkdownDescription: "This option enforces use of the specified version (or higher) on SSL connections. One of `SSLv3`, `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, `TLSv1.3`. Defaults to `TLSv1.2`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("TLSv1.2"),
				Validators:          []validator.String{stringvalidator.OneOf("", "SSLv3", "TLSv1.0", "TLSv1.1", "TLSv1.2", "TLSv1.3")},
			},
			"ssl_max_version": schema.StringAttribute{
				MarkdownDescription: "This option enforces use of the specified version (or lower) on SSL connections. One of `SSLv3`, `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, `TLSv1.3`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "SSLv3", "TLSv1.0", "TLSv1.1", "TLSv1.2", "TLSv1.3")},
			},
			"ssl_cipher_list": schema.StringAttribute{
				MarkdownDescription: "It sets the default string describing the list of cipher algorithms (\"cipher suite\") that are negotiated during the SSL/TLS handshake up to TLSv1.2. Defaults to `ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA256`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA256"),
			},
			"ssl_cipher_suites": schema.StringAttribute{
				MarkdownDescription: "It sets the default string describing the list of cipher algorithms (\"cipher suite\") that are negotiated during the SSL/TLS handshake for TLSv1.3. Defaults to `TLS_AES_128_GCM_SHA256:TLS_AES_256_GCM_SHA384:TLS_CHACHA20_POLY1305_SHA256`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("TLS_AES_128_GCM_SHA256:TLS_AES_256_GCM_SHA384:TLS_CHACHA20_POLY1305_SHA256"),
			},
			"ssl_hsts_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable HTTP Strict Transport Security. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ssl_hsts_include_sub_domains": schema.BoolAttribute{
				MarkdownDescription: "Enable if all present and future subdomains will be HTTPS. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_hsts_preload": schema.BoolAttribute{
				MarkdownDescription: "Enable if you like this domain to be included in the HSTS preload list. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_hsts_max_age": schema.Int64Attribute{
				MarkdownDescription: "Future requests to the domain should use only HTTPS for the specified time (in seconds): 15768000 = 6 months. Defaults to `15768000`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(15768000),
			},
			"ssl_client_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable Client Certificate Authentication. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_client_auth_verify": schema.StringAttribute{
				MarkdownDescription: "If set to 'optional' or 'required', client certificate is requested. One of `none`, `optional`, `required`. Defaults to `required`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("required"),
				Validators:          []validator.String{stringvalidator.OneOf("", "none", "optional", "required")},
			},
			"ssl_client_auth_cas": schema.SetAttribute{
				MarkdownDescription: "Select CA certificates to use for client certificate authentication. To import additional CAs, go to CA Manager. Takes the `ref_id` of an `opnsense_trust_ca`, not its UUID. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"ssl_client_auth_crls": schema.SetAttribute{
				MarkdownDescription: "Select CRLs to use for client certificate authentication. To import additional CRLs, go to CRL Manager. Takes the `ref_id` of an `opnsense_trust_crl`, not its UUID. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
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
			"tuning_max_connections": schema.Int64Attribute{
				MarkdownDescription: "Set the maximum number of concurrent connections for this Public Service. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"tuning_timeout_client": schema.StringAttribute{
				MarkdownDescription: "Set the maximum inactivity time on the client side. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_timeout_http_req": schema.StringAttribute{
				MarkdownDescription: "Set the maximum allowed time to wait for a complete HTTP request. In order to offer DoS protection, it may be required to lower the maximum accepted time to receive a complete HTTP request without affecting the client timeout. This helps protecting against established connections on which nothing is sent. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_timeout_http_keep_alive": schema.StringAttribute{
				MarkdownDescription: "Set the maximum allowed time to wait for a new HTTP request to appear. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"linked_cpu_affinity_rules": schema.SetAttribute{
				MarkdownDescription: "Choose CPU affinity rules that should be applied to this Public Service. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"tuning_shards": schema.Int64Attribute{
				MarkdownDescription: "This option automatically creates the specified number of listeners for every IP:port combination and evenly distributes them among available threads. This can sometimes be useful when using very large thread counts where the in-kernel locking on a single socket starts to cause a significant overhead. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"logging_dont_log_null": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable logging of connections with no data. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"logging_dont_log_normal": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable logging of normal, successful connections. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"logging_log_separate_errors": schema.BoolAttribute{
				MarkdownDescription: "Allow HAProxy to automatically raise log level for non-completely successful connections to aid debugging. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"logging_detailed_log": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable verbose logging. Each log line turns into a much richer format. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"logging_socket_stats": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable collecting & providing separate statistics for each socket. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"stickiness_pattern": schema.StringAttribute{
				MarkdownDescription: "Choose the type of data that should be stored in this stick-table. Note that this stick-table cannot be used for session persistence, it is only used to store additional per-connection data (select below). See the HAProxy documentation for further information. One of `binary`, `integer`, `ipv4`, `ipv6`, `string`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "binary", "integer", "ipv4", "ipv6", "string")},
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
			"stickiness_counter": schema.BoolAttribute{
				MarkdownDescription: "Enable to be able to retrieve values from sticky counters. If disabled, all values will return 0, rendering many conditions useless. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"stickiness_counter_key": schema.StringAttribute{
				MarkdownDescription: "It describes what elements of the incoming request or connection will be analyzed, extracted, combined, and used to select which table entry to update the counters. Defaults to \"src\" to track elements of the source IP. See the HAProxy documentation for a full description. Defaults to `src`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("src"),
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
			"stickiness_gpc_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. Most of the time it will be used to measure the frequency of occurrence of certain events. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\". Defaults to `1m`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("1m"),
			},
			"stickiness_gpt_elements": schema.Int64Attribute{
				MarkdownDescription: "The number of array elements that will be used for the General Purpose Tag. This is limited to a maximum of 100 elements, ranging from 0 to 99. Defaults to `0`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
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
			"http2_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable support for HTTP/2. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http2_enabled_nontls": schema.BoolAttribute{
				MarkdownDescription: "Enable support for HTTP/2 even if TLS (SSL offloading) is not enabled. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"advertised_protocols": schema.SetAttribute{
				MarkdownDescription: "When using the TLS ALPN extension, HAProxy advertises the specified protocol list as supported on top of ALPN. SSL offloading must be enabled. Note that HTTP/3 will only be added to QUIC-compatible Listening Addresses. One of `h3`, `h2`, `http11`, `http10`. Defaults to `[\"h2\", \"http11\"]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(types.SetValueMust(types.StringType, []attr.Value{types.StringValue("h2"), types.StringValue("http11")})),
				Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("h3", "h2", "http11", "http10"))},
			},
			"forward_for": schema.BoolAttribute{
				MarkdownDescription: "This option is deprecated and should no longer be used. A new option is available in backend pool settings. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"prometheus_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable HAProxy's Prometheus exporter. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"prometheus_path": schema.StringAttribute{
				MarkdownDescription: "The path where the Prometheus exporter can be accessed. Defaults to `/metrics`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/metrics"),
			},
			"connection_behaviour": schema.StringAttribute{
				MarkdownDescription: "By default HAProxy operates in keep-alive mode with regards to persistent connections. Option \"httpclose\" configures HAProxy to close connections with the server and the client as soon as the request and the response are received. It will also check if a \"Connection: close\" header is already set in each direction, and will add one if missing. Option \"http-server-close\" enables HTTP connection-close mode on the server side while keeping the ability to support HTTP keep-alive and pipelining on the client side. One of `http-keep-alive`, `httpclose`, `http-server-close`. Defaults to `http-keep-alive`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http-keep-alive"),
				Validators:          []validator.String{stringvalidator.OneOf("http-keep-alive", "httpclose", "http-server-close")},
			},
			"custom_options": schema.StringAttribute{
				MarkdownDescription: "These lines will be added to the HAProxy frontend configuration. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"linked_actions": schema.SetAttribute{
				MarkdownDescription: "Choose rules to be included in this Public Service. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"linked_errorfiles": schema.SetAttribute{
				MarkdownDescription: "Choose error messages to be included in this Public Service. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the frontend.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func frontendDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "HAProxy frontends (**Services ▸ HAProxy ▸ Settings ▸ Virtual Services ▸ Public Services**, plugin `os-haproxy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the frontend.",
				Required:            true,
			},
			"enabled":                          dschema.BoolAttribute{MarkdownDescription: "Enable this Public Service.", Computed: true},
			"name":                             dschema.StringAttribute{MarkdownDescription: "Name to identify this Public Service.", Computed: true},
			"description":                      dschema.StringAttribute{MarkdownDescription: "Description for this Public Service.", Computed: true},
			"bind":                             dschema.SetAttribute{MarkdownDescription: "Configure listen addresses for this Public Service, i.e. 127.0.0.1:8080 or www.example.com:443 or unix@socket-name. Add quic4@ or quic6@ to enable QUIC for IPv4/IPv6, e.g. quic4@www.example.com:443. Use TAB key to complete typing a listen address.", Computed: true, ElementType: types.StringType},
			"bind_options":                     dschema.StringAttribute{MarkdownDescription: "A list of parameters that will be appended to every Listen Address line.Example: accept-proxy npn http/1.1.", Computed: true},
			"mode":                             dschema.StringAttribute{MarkdownDescription: "Set the running mode or protocol for this Public Service. One of `http`, `ssl`, `tcp`.", Computed: true},
			"default_backend":                  dschema.StringAttribute{MarkdownDescription: "Set the default Backend Pool to use for this Public Service.", Computed: true},
			"ssl_enabled":                      dschema.BoolAttribute{MarkdownDescription: "Enable SSL offloading.", Computed: true},
			"ssl_certificates":                 dschema.SetAttribute{MarkdownDescription: "Select certificates to use for SSL offloading. HAProxy's SNI recognition will determine the correct certificate automatically. If no SNI is provided by the client then the first certificate will be presented.To import additional certificates, go to Certificate Manager.", Computed: true, ElementType: types.StringType},
			"ssl_default_certificate":          dschema.StringAttribute{MarkdownDescription: "This certificate will be presented if no SNI is provided by the client or if the client provides an SNI hostname which does not match any certificate. This parameter is optional to enforce a certain sort order for certificates.", Computed: true},
			"ssl_custom_options":               dschema.StringAttribute{MarkdownDescription: "Pass additional SSL parameters to the HAProxy configuration.", Computed: true},
			"ssl_advanced_enabled":             dschema.BoolAttribute{MarkdownDescription: "Enable advanced SSL settings.", Computed: true},
			"ssl_bind_options":                 dschema.SetAttribute{MarkdownDescription: "Used to enforce or disable certain SSL options. One of `no-sslv3`, `no-tlsv10`, `no-tlsv11`, `no-tlsv12`, `no-tlsv13`, `no-tls-tickets`, `force-sslv3`, `force-tlsv10`, `force-tlsv11`, `force-tlsv12`, `force-tlsv13`, `prefer-client-ciphers`, `strict-sni`.", Computed: true, ElementType: types.StringType},
			"ssl_min_version":                  dschema.StringAttribute{MarkdownDescription: "This option enforces use of the specified version (or higher) on SSL connections. One of `SSLv3`, `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, `TLSv1.3`.", Computed: true},
			"ssl_max_version":                  dschema.StringAttribute{MarkdownDescription: "This option enforces use of the specified version (or lower) on SSL connections. One of `SSLv3`, `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, `TLSv1.3`.", Computed: true},
			"ssl_cipher_list":                  dschema.StringAttribute{MarkdownDescription: "It sets the default string describing the list of cipher algorithms (\"cipher suite\") that are negotiated during the SSL/TLS handshake up to TLSv1.2.", Computed: true},
			"ssl_cipher_suites":                dschema.StringAttribute{MarkdownDescription: "It sets the default string describing the list of cipher algorithms (\"cipher suite\") that are negotiated during the SSL/TLS handshake for TLSv1.3.", Computed: true},
			"ssl_hsts_enabled":                 dschema.BoolAttribute{MarkdownDescription: "Enable HTTP Strict Transport Security.", Computed: true},
			"ssl_hsts_include_sub_domains":     dschema.BoolAttribute{MarkdownDescription: "Enable if all present and future subdomains will be HTTPS.", Computed: true},
			"ssl_hsts_preload":                 dschema.BoolAttribute{MarkdownDescription: "Enable if you like this domain to be included in the HSTS preload list.", Computed: true},
			"ssl_hsts_max_age":                 dschema.Int64Attribute{MarkdownDescription: "Future requests to the domain should use only HTTPS for the specified time (in seconds): 15768000 = 6 months.", Computed: true},
			"ssl_client_auth_enabled":          dschema.BoolAttribute{MarkdownDescription: "Enable Client Certificate Authentication.", Computed: true},
			"ssl_client_auth_verify":           dschema.StringAttribute{MarkdownDescription: "If set to 'optional' or 'required', client certificate is requested. One of `none`, `optional`, `required`.", Computed: true},
			"ssl_client_auth_cas":              dschema.SetAttribute{MarkdownDescription: "Select CA certificates to use for client certificate authentication. To import additional CAs, go to CA Manager.", Computed: true, ElementType: types.StringType},
			"ssl_client_auth_crls":             dschema.SetAttribute{MarkdownDescription: "Select CRLs to use for client certificate authentication. To import additional CRLs, go to CRL Manager.", Computed: true, ElementType: types.StringType},
			"basic_auth_enabled":               dschema.BoolAttribute{MarkdownDescription: "Enable HTTP Basic Authentication.", Computed: true},
			"basic_auth_users":                 dschema.SetAttribute{MarkdownDescription: "Allowed Users.", Computed: true, ElementType: types.StringType},
			"basic_auth_groups":                dschema.SetAttribute{MarkdownDescription: "Allowed Groups.", Computed: true, ElementType: types.StringType},
			"tuning_max_connections":           dschema.Int64Attribute{MarkdownDescription: "Set the maximum number of concurrent connections for this Public Service.", Computed: true},
			"tuning_timeout_client":            dschema.StringAttribute{MarkdownDescription: "Set the maximum inactivity time on the client side. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"tuning_timeout_http_req":          dschema.StringAttribute{MarkdownDescription: "Set the maximum allowed time to wait for a complete HTTP request. In order to offer DoS protection, it may be required to lower the maximum accepted time to receive a complete HTTP request without affecting the client timeout. This helps protecting against established connections on which nothing is sent. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"tuning_timeout_http_keep_alive":   dschema.StringAttribute{MarkdownDescription: "Set the maximum allowed time to wait for a new HTTP request to appear. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"linked_cpu_affinity_rules":        dschema.SetAttribute{MarkdownDescription: "Choose CPU affinity rules that should be applied to this Public Service.", Computed: true, ElementType: types.StringType},
			"tuning_shards":                    dschema.Int64Attribute{MarkdownDescription: "This option automatically creates the specified number of listeners for every IP:port combination and evenly distributes them among available threads. This can sometimes be useful when using very large thread counts where the in-kernel locking on a single socket starts to cause a significant overhead.", Computed: true},
			"logging_dont_log_null":            dschema.BoolAttribute{MarkdownDescription: "Enable or disable logging of connections with no data.", Computed: true},
			"logging_dont_log_normal":          dschema.BoolAttribute{MarkdownDescription: "Enable or disable logging of normal, successful connections.", Computed: true},
			"logging_log_separate_errors":      dschema.BoolAttribute{MarkdownDescription: "Allow HAProxy to automatically raise log level for non-completely successful connections to aid debugging.", Computed: true},
			"logging_detailed_log":             dschema.BoolAttribute{MarkdownDescription: "Enable or disable verbose logging. Each log line turns into a much richer format.", Computed: true},
			"logging_socket_stats":             dschema.BoolAttribute{MarkdownDescription: "Enable or disable collecting & providing separate statistics for each socket.", Computed: true},
			"stickiness_pattern":               dschema.StringAttribute{MarkdownDescription: "Choose the type of data that should be stored in this stick-table. Note that this stick-table cannot be used for session persistence, it is only used to store additional per-connection data (select below). See the HAProxy documentation for further information. One of `binary`, `integer`, `ipv4`, `ipv6`, `string`.", Computed: true},
			"stickiness_data_types":            dschema.SetAttribute{MarkdownDescription: "This is used to store additional information in the stick-table. It may be used by ACLs in order to control various criteria related to the activity of the client matching the stick-table. Note that this directly impacts memory usage. See the HAProxy documentation for a full description. One of `bytes_in_cnt`, `bytes_in_rate`, `bytes_out_cnt`, `bytes_out_rate`, `conn_cnt`, `conn_cur`, `conn_rate`, `glitch_cnt`, `glitch_rate`, `gpc`, `gpc_rate`, `gpc0`, `gpc0_rate`, `gpc1`, `gpc1_rate`, `gpt`, `gpt0`, `http_err_cnt`, `http_err_rate`, `http_fail_cnt`, `http_fail_rate`, `http_req_cnt`, `http_req_rate`, `server_id`, `sess_cnt`, `sess_rate`.", Computed: true, ElementType: types.StringType},
			"stickiness_expire":                dschema.StringAttribute{MarkdownDescription: "Enter a number followed by one of the supported suffixes \"d\" (days), \"h\" (hour), \"m\" (minute), \"s\" (seconds), \"ms\" (miliseconds). This configures the maximum duration of an entry in the stick-table since it was last created, refreshed or matched. The maximum duration is slightly above 24 days.", Computed: true},
			"stickiness_size":                  dschema.StringAttribute{MarkdownDescription: "Enter a number followed by one of the supported suffixes \"k\", \"m\", \"g\". This configures the maximum number of entries that can fit in the table. This value directly impacts memory usage. Count approximately 50 bytes per entry, plus the size of a string if any.", Computed: true},
			"stickiness_counter":               dschema.BoolAttribute{MarkdownDescription: "Enable to be able to retrieve values from sticky counters. If disabled, all values will return 0, rendering many conditions useless.", Computed: true},
			"stickiness_counter_key":           dschema.StringAttribute{MarkdownDescription: "It describes what elements of the incoming request or connection will be analyzed, extracted, combined, and used to select which table entry to update the counters. Defaults to \"src\" to track elements of the source IP. See the HAProxy documentation for a full description.", Computed: true},
			"stickiness_length":                dschema.Int64Attribute{MarkdownDescription: "Specify the maximum length for a value in the stick-table. If the value is larger than this value it will be truncated before being stored. Depending on the stick-table type this repesents either characters or bytes.", Computed: true},
			"stickiness_conn_rate_period":      dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. It reports the average incoming connection rate over that period, in connections per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"stickiness_sess_rate_period":      dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. It reports the average incoming session rate over that period, in sessions per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"stickiness_http_req_rate_period":  dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. It reports the average HTTP request rate over that period, in requests per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"stickiness_http_err_rate_period":  dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. It reports the average HTTP request error rate over that period, in requests per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"stickiness_bytes_in_rate_period":  dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. It reports the average incoming bytes rate over that period, in bytes per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"stickiness_bytes_out_rate_period": dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. It reports the average outgoing bytes rate over that period, in bytes per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"stickiness_gpc_elements":          dschema.Int64Attribute{MarkdownDescription: "The number of array elements that will be used for the General Purpose Counter. This is limited to a maximum of 100 elements, ranging from 0 to 99.", Computed: true},
			"stickiness_gpc_rate_period":       dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. Most of the time it will be used to measure the frequency of occurrence of certain events. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"stickiness_gpt_elements":          dschema.Int64Attribute{MarkdownDescription: "The number of array elements that will be used for the General Purpose Tag. This is limited to a maximum of 100 elements, ranging from 0 to 99.", Computed: true},
			"stickiness_http_fail_rate_period": dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. It reports the average HTTP response fail error rate over that period, in requests per period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"stickiness_glitch_rate_period":    dschema.StringAttribute{MarkdownDescription: "The length of the period over which the average is measured. It reports the average front glitches rate over that period. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".", Computed: true},
			"http2_enabled":                    dschema.BoolAttribute{MarkdownDescription: "Enable support for HTTP/2.", Computed: true},
			"http2_enabled_nontls":             dschema.BoolAttribute{MarkdownDescription: "Enable support for HTTP/2 even if TLS (SSL offloading) is not enabled.", Computed: true},
			"advertised_protocols":             dschema.SetAttribute{MarkdownDescription: "When using the TLS ALPN extension, HAProxy advertises the specified protocol list as supported on top of ALPN. SSL offloading must be enabled. Note that HTTP/3 will only be added to QUIC-compatible Listening Addresses. One of `h3`, `h2`, `http11`, `http10`.", Computed: true, ElementType: types.StringType},
			"forward_for":                      dschema.BoolAttribute{MarkdownDescription: "This option is deprecated and should no longer be used. A new option is available in backend pool settings.", Computed: true},
			"prometheus_enabled":               dschema.BoolAttribute{MarkdownDescription: "Enable HAProxy's Prometheus exporter.", Computed: true},
			"prometheus_path":                  dschema.StringAttribute{MarkdownDescription: "The path where the Prometheus exporter can be accessed.", Computed: true},
			"connection_behaviour":             dschema.StringAttribute{MarkdownDescription: "By default HAProxy operates in keep-alive mode with regards to persistent connections. Option \"httpclose\" configures HAProxy to close connections with the server and the client as soon as the request and the response are received. It will also check if a \"Connection: close\" header is already set in each direction, and will add one if missing. Option \"http-server-close\" enables HTTP connection-close mode on the server side while keeping the ability to support HTTP keep-alive and pipelining on the client side. One of `http-keep-alive`, `httpclose`, `http-server-close`.", Computed: true},
			"custom_options":                   dschema.StringAttribute{MarkdownDescription: "These lines will be added to the HAProxy frontend configuration.", Computed: true},
			"linked_actions":                   dschema.SetAttribute{MarkdownDescription: "Choose rules to be included in this Public Service.", Computed: true, ElementType: types.StringType},
			"linked_errorfiles":                dschema.SetAttribute{MarkdownDescription: "Choose error messages to be included in this Public Service.", Computed: true, ElementType: types.StringType},
		},
	}
}

func convertFrontendSchemaToStruct(d *frontendResourceModel) (*haproxy.Frontend, error) {
	return &haproxy.Frontend{
		Enabled:                      tools.BoolToString(d.Enabled.ValueBool()),
		Name:                         d.Name.ValueString(),
		Description:                  d.Description.ValueString(),
		Bind:                         api.SelectedMapList(setToSlice(d.Bind)),
		BindOptions:                  d.BindOptions.ValueString(),
		Mode:                         api.SelectedMap(d.Mode.ValueString()),
		DefaultBackend:               api.SelectedMap(d.DefaultBackend.ValueString()),
		SslEnabled:                   tools.BoolToString(d.SslEnabled.ValueBool()),
		SslCertificates:              api.SelectedMapList(setToSlice(d.SslCertificates)),
		SslDefaultCertificate:        api.SelectedMap(d.SslDefaultCertificate.ValueString()),
		SslCustomOptions:             d.SslCustomOptions.ValueString(),
		SslAdvancedEnabled:           tools.BoolToString(d.SslAdvancedEnabled.ValueBool()),
		SslBindOptions:               api.SelectedMapList(setToSlice(d.SslBindOptions)),
		SslMinVersion:                api.SelectedMap(d.SslMinVersion.ValueString()),
		SslMaxVersion:                api.SelectedMap(d.SslMaxVersion.ValueString()),
		SslCipherList:                d.SslCipherList.ValueString(),
		SslCipherSuites:              d.SslCipherSuites.ValueString(),
		SslHstsEnabled:               tools.BoolToString(d.SslHstsEnabled.ValueBool()),
		SslHstsIncludeSubDomains:     tools.BoolToString(d.SslHstsIncludeSubDomains.ValueBool()),
		SslHstsPreload:               tools.BoolToString(d.SslHstsPreload.ValueBool()),
		SslHstsMaxAge:                tools.Int64ToString(d.SslHstsMaxAge.ValueInt64()),
		SslClientAuthEnabled:         tools.BoolToString(d.SslClientAuthEnabled.ValueBool()),
		SslClientAuthVerify:          api.SelectedMap(d.SslClientAuthVerify.ValueString()),
		SslClientAuthCas:             api.SelectedMapList(setToSlice(d.SslClientAuthCas)),
		SslClientAuthCrls:            api.SelectedMapList(setToSlice(d.SslClientAuthCrls)),
		BasicAuthEnabled:             tools.BoolToString(d.BasicAuthEnabled.ValueBool()),
		BasicAuthUsers:               api.SelectedMapList(setToSlice(d.BasicAuthUsers)),
		BasicAuthGroups:              api.SelectedMapList(setToSlice(d.BasicAuthGroups)),
		TuningMaxConnections:         tools.Int64ToStringNegative(d.TuningMaxConnections.ValueInt64()),
		TuningTimeoutClient:          d.TuningTimeoutClient.ValueString(),
		TuningTimeoutHttpReq:         d.TuningTimeoutHttpReq.ValueString(),
		TuningTimeoutHttpKeepAlive:   d.TuningTimeoutHttpKeepAlive.ValueString(),
		LinkedCpuAffinityRules:       api.SelectedMapList(setToSlice(d.LinkedCpuAffinityRules)),
		TuningShards:                 tools.Int64ToStringNegative(d.TuningShards.ValueInt64()),
		LoggingDontLogNull:           tools.BoolToString(d.LoggingDontLogNull.ValueBool()),
		LoggingDontLogNormal:         tools.BoolToString(d.LoggingDontLogNormal.ValueBool()),
		LoggingLogSeparateErrors:     tools.BoolToString(d.LoggingLogSeparateErrors.ValueBool()),
		LoggingDetailedLog:           tools.BoolToString(d.LoggingDetailedLog.ValueBool()),
		LoggingSocketStats:           tools.BoolToString(d.LoggingSocketStats.ValueBool()),
		StickinessPattern:            api.SelectedMap(d.StickinessPattern.ValueString()),
		StickinessDataTypes:          api.SelectedMapList(setToSlice(d.StickinessDataTypes)),
		StickinessExpire:             d.StickinessExpire.ValueString(),
		StickinessSize:               d.StickinessSize.ValueString(),
		StickinessCounter:            tools.BoolToString(d.StickinessCounter.ValueBool()),
		StickinessCounterKey:         d.StickinessCounterKey.ValueString(),
		StickinessLength:             tools.Int64ToStringNegative(d.StickinessLength.ValueInt64()),
		StickinessConnRatePeriod:     d.StickinessConnRatePeriod.ValueString(),
		StickinessSessRatePeriod:     d.StickinessSessRatePeriod.ValueString(),
		StickinessHttpReqRatePeriod:  d.StickinessHttpReqRatePeriod.ValueString(),
		StickinessHttpErrRatePeriod:  d.StickinessHttpErrRatePeriod.ValueString(),
		StickinessBytesInRatePeriod:  d.StickinessBytesInRatePeriod.ValueString(),
		StickinessBytesOutRatePeriod: d.StickinessBytesOutRatePeriod.ValueString(),
		StickinessGpcElements:        tools.Int64ToString(d.StickinessGpcElements.ValueInt64()),
		StickinessGpcRatePeriod:      d.StickinessGpcRatePeriod.ValueString(),
		StickinessGptElements:        tools.Int64ToString(d.StickinessGptElements.ValueInt64()),
		StickinessHttpFailRatePeriod: d.StickinessHttpFailRatePeriod.ValueString(),
		StickinessGlitchRatePeriod:   d.StickinessGlitchRatePeriod.ValueString(),
		Http2Enabled:                 tools.BoolToString(d.Http2Enabled.ValueBool()),
		Http2EnabledNontls:           tools.BoolToString(d.Http2EnabledNontls.ValueBool()),
		AdvertisedProtocols:          api.SelectedMapList(setToSlice(d.AdvertisedProtocols)),
		ForwardFor:                   tools.BoolToString(d.ForwardFor.ValueBool()),
		PrometheusEnabled:            tools.BoolToString(d.PrometheusEnabled.ValueBool()),
		PrometheusPath:               d.PrometheusPath.ValueString(),
		ConnectionBehaviour:          api.SelectedMap(d.ConnectionBehaviour.ValueString()),
		CustomOptions:                d.CustomOptions.ValueString(),
		LinkedActions:                api.SelectedMapList(setToSlice(d.LinkedActions)),
		LinkedErrorfiles:             api.SelectedMapList(setToSlice(d.LinkedErrorfiles)),
	}, nil
}

func convertFrontendStructToSchema(d *haproxy.Frontend) (*frontendResourceModel, error) {
	return &frontendResourceModel{
		Enabled:                      types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:                         types.StringValue(d.Name),
		Description:                  types.StringValue(d.Description),
		Bind:                         tools.StringSliceToSet(d.Bind),
		BindOptions:                  types.StringValue(d.BindOptions),
		Mode:                         types.StringValue(d.Mode.String()),
		DefaultBackend:               types.StringValue(d.DefaultBackend.String()),
		SslEnabled:                   types.BoolValue(tools.StringToBool(d.SslEnabled)),
		SslCertificates:              tools.StringSliceToSet(d.SslCertificates),
		SslDefaultCertificate:        types.StringValue(d.SslDefaultCertificate.String()),
		SslCustomOptions:             types.StringValue(d.SslCustomOptions),
		SslAdvancedEnabled:           types.BoolValue(tools.StringToBool(d.SslAdvancedEnabled)),
		SslBindOptions:               tools.StringSliceToSet(d.SslBindOptions),
		SslMinVersion:                types.StringValue(d.SslMinVersion.String()),
		SslMaxVersion:                types.StringValue(d.SslMaxVersion.String()),
		SslCipherList:                types.StringValue(d.SslCipherList),
		SslCipherSuites:              types.StringValue(d.SslCipherSuites),
		SslHstsEnabled:               types.BoolValue(tools.StringToBool(d.SslHstsEnabled)),
		SslHstsIncludeSubDomains:     types.BoolValue(tools.StringToBool(d.SslHstsIncludeSubDomains)),
		SslHstsPreload:               types.BoolValue(tools.StringToBool(d.SslHstsPreload)),
		SslHstsMaxAge:                types.Int64Value(tools.StringToInt64(d.SslHstsMaxAge)),
		SslClientAuthEnabled:         types.BoolValue(tools.StringToBool(d.SslClientAuthEnabled)),
		SslClientAuthVerify:          types.StringValue(d.SslClientAuthVerify.String()),
		SslClientAuthCas:             tools.StringSliceToSet(d.SslClientAuthCas),
		SslClientAuthCrls:            tools.StringSliceToSet(d.SslClientAuthCrls),
		BasicAuthEnabled:             types.BoolValue(tools.StringToBool(d.BasicAuthEnabled)),
		BasicAuthUsers:               tools.StringSliceToSet(d.BasicAuthUsers),
		BasicAuthGroups:              tools.StringSliceToSet(d.BasicAuthGroups),
		TuningMaxConnections:         types.Int64Value(tools.StringToInt64(d.TuningMaxConnections)),
		TuningTimeoutClient:          types.StringValue(d.TuningTimeoutClient),
		TuningTimeoutHttpReq:         types.StringValue(d.TuningTimeoutHttpReq),
		TuningTimeoutHttpKeepAlive:   types.StringValue(d.TuningTimeoutHttpKeepAlive),
		LinkedCpuAffinityRules:       tools.StringSliceToSet(d.LinkedCpuAffinityRules),
		TuningShards:                 types.Int64Value(tools.StringToInt64(d.TuningShards)),
		LoggingDontLogNull:           types.BoolValue(tools.StringToBool(d.LoggingDontLogNull)),
		LoggingDontLogNormal:         types.BoolValue(tools.StringToBool(d.LoggingDontLogNormal)),
		LoggingLogSeparateErrors:     types.BoolValue(tools.StringToBool(d.LoggingLogSeparateErrors)),
		LoggingDetailedLog:           types.BoolValue(tools.StringToBool(d.LoggingDetailedLog)),
		LoggingSocketStats:           types.BoolValue(tools.StringToBool(d.LoggingSocketStats)),
		StickinessPattern:            types.StringValue(d.StickinessPattern.String()),
		StickinessDataTypes:          tools.StringSliceToSet(d.StickinessDataTypes),
		StickinessExpire:             types.StringValue(d.StickinessExpire),
		StickinessSize:               types.StringValue(d.StickinessSize),
		StickinessCounter:            types.BoolValue(tools.StringToBool(d.StickinessCounter)),
		StickinessCounterKey:         types.StringValue(d.StickinessCounterKey),
		StickinessLength:             types.Int64Value(tools.StringToInt64(d.StickinessLength)),
		StickinessConnRatePeriod:     types.StringValue(d.StickinessConnRatePeriod),
		StickinessSessRatePeriod:     types.StringValue(d.StickinessSessRatePeriod),
		StickinessHttpReqRatePeriod:  types.StringValue(d.StickinessHttpReqRatePeriod),
		StickinessHttpErrRatePeriod:  types.StringValue(d.StickinessHttpErrRatePeriod),
		StickinessBytesInRatePeriod:  types.StringValue(d.StickinessBytesInRatePeriod),
		StickinessBytesOutRatePeriod: types.StringValue(d.StickinessBytesOutRatePeriod),
		StickinessGpcElements:        types.Int64Value(tools.StringToInt64(d.StickinessGpcElements)),
		StickinessGpcRatePeriod:      types.StringValue(d.StickinessGpcRatePeriod),
		StickinessGptElements:        types.Int64Value(tools.StringToInt64(d.StickinessGptElements)),
		StickinessHttpFailRatePeriod: types.StringValue(d.StickinessHttpFailRatePeriod),
		StickinessGlitchRatePeriod:   types.StringValue(d.StickinessGlitchRatePeriod),
		Http2Enabled:                 types.BoolValue(tools.StringToBool(d.Http2Enabled)),
		Http2EnabledNontls:           types.BoolValue(tools.StringToBool(d.Http2EnabledNontls)),
		AdvertisedProtocols:          tools.StringSliceToSet(d.AdvertisedProtocols),
		ForwardFor:                   types.BoolValue(tools.StringToBool(d.ForwardFor)),
		PrometheusEnabled:            types.BoolValue(tools.StringToBool(d.PrometheusEnabled)),
		PrometheusPath:               types.StringValue(d.PrometheusPath),
		ConnectionBehaviour:          types.StringValue(d.ConnectionBehaviour.String()),
		CustomOptions:                types.StringValue(d.CustomOptions),
		LinkedActions:                tools.StringSliceToSet(d.LinkedActions),
		LinkedErrorfiles:             tools.StringSliceToSet(d.LinkedErrorfiles),
	}, nil
}

func setToSlice(s types.Set) []string {
	out := tools.SetToStringSlice(s)
	if out == nil {
		return []string{}
	}
	return out
}
