package haproxy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type settingsResourceModel struct {
	Enabled                              types.Bool   `tfsdk:"enabled"`
	GracefulStop                         types.Bool   `tfsdk:"graceful_stop"`
	HardStopAfter                        types.String `tfsdk:"hard_stop_after"`
	CloseSpreadTime                      types.String `tfsdk:"close_spread_time"`
	SeamlessReload                       types.Bool   `tfsdk:"seamless_reload"`
	StoreOcsp                            types.Bool   `tfsdk:"store_ocsp"`
	PeersEnabled                         types.Bool   `tfsdk:"peers_enabled"`
	PeersName1                           types.String `tfsdk:"peers_name1"`
	PeersListen1                         types.String `tfsdk:"peers_listen1"`
	PeersPort1                           types.Int64  `tfsdk:"peers_port1"`
	PeersName2                           types.String `tfsdk:"peers_name2"`
	PeersListen2                         types.String `tfsdk:"peers_listen2"`
	PeersPort2                           types.Int64  `tfsdk:"peers_port2"`
	TuningRoot                           types.Bool   `tfsdk:"tuning_root"`
	TuningMaxConnections                 types.Int64  `tfsdk:"tuning_max_connections"`
	TuningNbthread                       types.Int64  `tfsdk:"tuning_nbthread"`
	TuningResolversPrefer                types.String `tfsdk:"tuning_resolvers_prefer"`
	TuningSslServerVerify                types.String `tfsdk:"tuning_ssl_server_verify"`
	TuningMaxDhSize                      types.Int64  `tfsdk:"tuning_max_dh_size"`
	TuningBufferSize                     types.Int64  `tfsdk:"tuning_buffer_size"`
	TuningSpreadChecks                   types.Int64  `tfsdk:"tuning_spread_checks"`
	TuningBogusProxyEnabled              types.Bool   `tfsdk:"tuning_bogus_proxy_enabled"`
	TuningLuaMaxMem                      types.Int64  `tfsdk:"tuning_lua_max_mem"`
	TuningCustomOptions                  types.String `tfsdk:"tuning_custom_options"`
	TuningOcspUpdateEnabled              types.Bool   `tfsdk:"tuning_ocsp_update_enabled"`
	TuningOcspUpdateMinDelay             types.Int64  `tfsdk:"tuning_ocsp_update_min_delay"`
	TuningOcspUpdateMaxDelay             types.Int64  `tfsdk:"tuning_ocsp_update_max_delay"`
	TuningSslDefaultsEnabled             types.Bool   `tfsdk:"tuning_ssl_defaults_enabled"`
	TuningSslBindOptions                 types.Set    `tfsdk:"tuning_ssl_bind_options"`
	TuningSslMinVersion                  types.String `tfsdk:"tuning_ssl_min_version"`
	TuningSslMaxVersion                  types.String `tfsdk:"tuning_ssl_max_version"`
	TuningSslCipherList                  types.String `tfsdk:"tuning_ssl_cipher_list"`
	TuningSslCipherSuites                types.String `tfsdk:"tuning_ssl_cipher_suites"`
	TuningH2InitialWindowSize            types.Int64  `tfsdk:"tuning_h2_initial_window_size"`
	TuningH2InitialWindowSizeOutgoing    types.Int64  `tfsdk:"tuning_h2_initial_window_size_outgoing"`
	TuningH2InitialWindowSizeIncoming    types.Int64  `tfsdk:"tuning_h2_initial_window_size_incoming"`
	TuningH2MaxConcurrentStreams         types.Int64  `tfsdk:"tuning_h2_max_concurrent_streams"`
	TuningH2MaxConcurrentStreamsOutgoing types.Int64  `tfsdk:"tuning_h2_max_concurrent_streams_outgoing"`
	TuningH2MaxConcurrentStreamsIncoming types.Int64  `tfsdk:"tuning_h2_max_concurrent_streams_incoming"`
	DefaultsMaxConnections               types.Int64  `tfsdk:"defaults_max_connections"`
	DefaultsMaxConnectionsServers        types.Int64  `tfsdk:"defaults_max_connections_servers"`
	DefaultsTimeoutClient                types.String `tfsdk:"defaults_timeout_client"`
	DefaultsTimeoutConnect               types.String `tfsdk:"defaults_timeout_connect"`
	DefaultsTimeoutCheck                 types.String `tfsdk:"defaults_timeout_check"`
	DefaultsTimeoutServer                types.String `tfsdk:"defaults_timeout_server"`
	DefaultsRetries                      types.Int64  `tfsdk:"defaults_retries"`
	DefaultsRedispatch                   types.String `tfsdk:"defaults_redispatch"`
	DefaultsInitAddr                     types.Set    `tfsdk:"defaults_init_addr"`
	DefaultsCustomOptions                types.String `tfsdk:"defaults_custom_options"`
	LoggingHost                          types.String `tfsdk:"logging_host"`
	LoggingFacility                      types.String `tfsdk:"logging_facility"`
	LoggingLevel                         types.String `tfsdk:"logging_level"`
	LoggingLength                        types.Int64  `tfsdk:"logging_length"`
	StatsEnabled                         types.Bool   `tfsdk:"stats_enabled"`
	StatsPort                            types.Int64  `tfsdk:"stats_port"`
	StatsRemoteEnabled                   types.Bool   `tfsdk:"stats_remote_enabled"`
	StatsRemoteBind                      types.Set    `tfsdk:"stats_remote_bind"`
	StatsAuthEnabled                     types.Bool   `tfsdk:"stats_auth_enabled"`
	StatsUsers                           types.Set    `tfsdk:"stats_users"`
	StatsAllowedUsers                    types.Set    `tfsdk:"stats_allowed_users"`
	StatsAllowedGroups                   types.Set    `tfsdk:"stats_allowed_groups"`
	StatsCustomOptions                   types.String `tfsdk:"stats_custom_options"`
	StatsPrometheusEnabled               types.Bool   `tfsdk:"stats_prometheus_enabled"`
	StatsPrometheusBind                  types.Set    `tfsdk:"stats_prometheus_bind"`
	StatsPrometheusPath                  types.String `tfsdk:"stats_prometheus_path"`
	CacheEnabled                         types.Bool   `tfsdk:"cache_enabled"`
	CacheTotalMaxSize                    types.Int64  `tfsdk:"cache_total_max_size"`
	CacheMaxAge                          types.Int64  `tfsdk:"cache_max_age"`
	CacheMaxObjectSize                   types.Int64  `tfsdk:"cache_max_object_size"`
	CacheProcessVary                     types.Bool   `tfsdk:"cache_process_vary"`
	CacheMaxSecondaryEntries             types.Int64  `tfsdk:"cache_max_secondary_entries"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable or disable the HAProxy service.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"graceful_stop": schema.BoolAttribute{
			MarkdownDescription: "Enable HAProxy's graceful stop mode. In this mode HAProxy will continue to process existing connections until they close. Note that this may severely slow down HAProxy's shutdown, depending on the configured timeout values. If graceful stop mode is not enabled, HAProxy will use the hard stop mode where it immediately quits and all established connections are closed. Hard stop mode is recommended.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"hard_stop_after": schema.StringAttribute{
			MarkdownDescription: "Set the maximum time allowed to perform a clean graceful stop. HAProxy will terminate all open connections when the timeout is reached. This may be used to ensure that the instance will quit even if connections remain opened. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"close_spread_time": schema.StringAttribute{
			MarkdownDescription: "Specifies a time window during which connection closing will be spread during a soft-stop operation. Idle connections will all be closed at once if this option is not set, which may cause reconnecting clients to rush against the process. For best results, it should set lower than the \"Graceful stop timout\" option. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"seamless_reload": schema.BoolAttribute{
			MarkdownDescription: "HAProxy will handle service restarts in a way that no connections are dropped. This is the best restart mode, because it has no impact on user experience. That being said, there might be edge cases where seamless reloads lead to unexpected behaviour.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"store_ocsp": schema.BoolAttribute{
			MarkdownDescription: "storeOcsp.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"peers_enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable or disable HAProxy peers. This will propagate entries of any data-types in stick-tables between several HAProxy instances over TCP connections in a multi-master fashion.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"peers_name1": schema.StringAttribute{
			MarkdownDescription: "The name of the peer. This is usually the fully qualified domain name. If the name matches the system hostname, then this peer is automatically configured as local peer.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"peers_listen1": schema.StringAttribute{
			MarkdownDescription: "The listen address of the local peer or the address of the remote peer.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"peers_port1": schema.Int64Attribute{
			MarkdownDescription: "The TCP port that should be used for connections to this peer. It must not be used by any other service. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"peers_name2": schema.StringAttribute{
			MarkdownDescription: "The name of the peer. This is usually the fully qualified domain name. If the name matches the system hostname, then this peer is automatically configured as local peer.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"peers_listen2": schema.StringAttribute{
			MarkdownDescription: "The listen address of the local peer or the address of the remote peer.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"peers_port2": schema.Int64Attribute{
			MarkdownDescription: "The TCP port that should be used for connections to this peer. It must not be used by any other service. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tuning_root": schema.BoolAttribute{
			MarkdownDescription: "Enable or disable HAProxy running as user root. Enabling this option is strongly discouraged.NOTE: Running as user root could be a security issue but it may be required by some features.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"tuning_max_connections": schema.Int64Attribute{
			MarkdownDescription: "Sets the maximum number of concurrent connections per HAProxy process.NOTE: Consider raising the settings for kern.maxfiles and kern.maxfilesperproc in System: Settings: Tunables, otherwise HAProxy will fail to open the specified number of connections. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tuning_nbthread": schema.Int64Attribute{
			MarkdownDescription: "Number of threads to create for each HAProxy process. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tuning_resolvers_prefer": schema.StringAttribute{
			MarkdownDescription: "This option allows to choose which IP family is preferred when resolving DNS names. This is useful when IPv6 or IPv4 is not available. It solves a common issue with OCSP updates. One of `ipv4`, `ipv6`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"tuning_ssl_server_verify": schema.StringAttribute{
			MarkdownDescription: "This enforces a certain behavior for SSL verify on servers, ignoring per-server settings. If set to 'enforce verify', server certificates are verified. If set to 'disable verify', server certificates are not verified. The default is 'no preference' to only use per-server configurations and not enforce a global default for all servers. One of `ignore`, `required`, `none`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"tuning_max_dh_size": schema.Int64Attribute{
			MarkdownDescription: "Sets the maximum size of the Diffie-Hellman parameters used for generating the ephemeral/temporary Diffie-Hellman key in case of DHE key exchange (default is 1024).NOTE: Higher values will increase the CPU load. For more information about the \"tune.ssl.default-dh-param\" option please see the HAProxy Documentation. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tuning_buffer_size": schema.Int64Attribute{
			MarkdownDescription: "Change the buffer size (in bytes). Lower values allow more sessions to coexist in the same amount of RAM, and higher values allow some applications with very large cookies to work. The default value is 16384. NOTE: It is strongly recommended not to change this from the default value, as very low values will break some services such as statistics, and values larger than default size will increase memory usage, possibly causing the system to run out of memory. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tuning_spread_checks": schema.Int64Attribute{
			MarkdownDescription: "Add some randomness in the check interval between 0 and +/- 50%. A value between 2 and 5 seems to show good results. The default value is 0 (disabled). `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tuning_bogus_proxy_enabled": schema.BoolAttribute{
			MarkdownDescription: "A bug in the PROXY protocol v2 implementation was present in HAProxy up to version 2.1. Enabling this option reverts this old buggy behaviour.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"tuning_lua_max_mem": schema.Int64Attribute{
			MarkdownDescription: "Sets the maximum amount of RAM in megabytes per process usable by Lua. By default it is zero which means unlimited. It is important to set a limit to ensure that a bug in a script will not result in the system running out of memory. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tuning_custom_options": schema.StringAttribute{
			MarkdownDescription: "These lines will be added to the global settings of to the HAProxy configuration file.NOTE: The syntax will not be checked, use at your own risk!",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"tuning_ocsp_update_enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable automatic OCSP response updates. Each OCSP response will be updated at least once an hour, and even more frequently if a given OCSP response has an expire date earlier than this one hour limit.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"tuning_ocsp_update_min_delay": schema.Int64Attribute{
			MarkdownDescription: "Sets the minimum interval (in seconds) between two automatic updates of the same OCSP response. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tuning_ocsp_update_max_delay": schema.Int64Attribute{
			MarkdownDescription: "Sets the maximum interval (in seconds) between two automatic updates of the same OCSP response. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tuning_ssl_defaults_enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable global SSL default values.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"tuning_ssl_bind_options": schema.SetAttribute{
			MarkdownDescription: "Used to enforce or disable certain SSL options. One of `no-sslv3`, `no-tlsv10`, `no-tlsv11`, `no-tlsv12`, `no-tlsv13`, `no-tls-tickets`, `force-sslv3`, `force-tlsv10`, `force-tlsv11`, `force-tlsv12`, `force-tlsv13`, `prefer-client-ciphers`, `strict-sni`.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"tuning_ssl_min_version": schema.StringAttribute{
			MarkdownDescription: "This option enforces use of the specified version (or higher) on SSL connections. One of `SSLv3`, `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, `TLSv1.3`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"tuning_ssl_max_version": schema.StringAttribute{
			MarkdownDescription: "This option enforces use of the specified version (or lower) on SSL connections. One of `SSLv3`, `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, `TLSv1.3`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"tuning_ssl_cipher_list": schema.StringAttribute{
			MarkdownDescription: "It sets the default string describing the list of cipher algorithms (\"cipher suite\") that are negotiated during the SSL/TLS handshake up to TLSv1.2.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"tuning_ssl_cipher_suites": schema.StringAttribute{
			MarkdownDescription: "It sets the default string describing the list of cipher algorithms (\"cipher suite\") that are negotiated during the SSL/TLS handshake for TLSv1.3.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"tuning_h2_initial_window_size": schema.Int64Attribute{
			MarkdownDescription: "Sets the default value for the HTTP/2 initial window size, on both incoming and outgoing connections. The default value is 65536. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tuning_h2_initial_window_size_outgoing": schema.Int64Attribute{
			MarkdownDescription: "Sets the HTTP/2 initial window size for outgoing connections, which is the number of bytes the server can respond to before waiting for an acknowledgment from HAProxy. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tuning_h2_initial_window_size_incoming": schema.Int64Attribute{
			MarkdownDescription: "Sets the HTTP/2 initial window size for incoming connections, which is the number of bytes the client can upload before waiting for an acknowledgment from HAProxy. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tuning_h2_max_concurrent_streams": schema.Int64Attribute{
			MarkdownDescription: "Sets the default HTTP/2 maximum number of concurrent streams per connection (i.e. the number of outstanding requests on a single connection). The default value is 100. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tuning_h2_max_concurrent_streams_outgoing": schema.Int64Attribute{
			MarkdownDescription: "Sets the HTTP/2 maximum number of concurrent streams per outgoing connection (i.e. the number of outstanding requests on a single connection to a server). `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tuning_h2_max_concurrent_streams_incoming": schema.Int64Attribute{
			MarkdownDescription: "Sets the HTTP/2 maximum number of concurrent streams per incoming connection (i.e. the number of outstanding requests on a single connection from a client). `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"defaults_max_connections": schema.Int64Attribute{
			MarkdownDescription: "Set the maximum number of concurrent connections for public services. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"defaults_max_connections_servers": schema.Int64Attribute{
			MarkdownDescription: "Set the maximum number of concurrent connections for servers. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"defaults_timeout_client": schema.StringAttribute{
			MarkdownDescription: "Set the maximum inactivity time on the client side. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"defaults_timeout_connect": schema.StringAttribute{
			MarkdownDescription: "Set the maximum time to wait for a connection attempt to a server to succeed. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"defaults_timeout_check": schema.StringAttribute{
			MarkdownDescription: "Sets an additional read timeout for running health checks on a server. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"defaults_timeout_server": schema.StringAttribute{
			MarkdownDescription: "Set the maximum inactivity time on the server side. Defaults to milliseconds. Optionally the unit may be specified as either \"d\", \"h\", \"m\", \"s\", \"ms\" or \"us\".",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"defaults_retries": schema.Int64Attribute{
			MarkdownDescription: "Set the number of retries to perform on a server after a connection failure (default is 3). `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"defaults_redispatch": schema.StringAttribute{
			MarkdownDescription: "Enable or disable session redistribution in case of connection failure. One of `x3`, `x2`, `x1`, `x0`, `x-1`, `x-2`, `x-3`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"defaults_init_addr": schema.SetAttribute{
			MarkdownDescription: "Indicates in which order server addresses should be resolved upon startup. Method \"last\" suggests to pick the address which appears in the state file. Method \"libc\" uses the libc's internal resolver. Method \"none\" specifically indicates that the server should start without any valid IP address in a down state. It can be useful to ignore some DNS issues upon startup, waiting for the situation to get fixed later. Defaults to \"last,libc\". One of `last`, `libc`, `none`.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"defaults_custom_options": schema.StringAttribute{
			MarkdownDescription: "These lines will be added to the defaults settings of to the HAProxy configuration file.NOTE: The syntax will not be checked, use at your own risk!",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"logging_host": schema.StringAttribute{
			MarkdownDescription: "Indicates where to send the logs. Takes an IPv4 or IPv6 address optionally followed by a colon (':') and a UDP port, i.e. 127.0.0.1 or 10.0.0.1:514.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"logging_facility": schema.StringAttribute{
			MarkdownDescription: "Choose one of the 24 standard syslog facilities. The default value is local0. One of `alert`, `audit`, `auth2`, `auth`, `cron2`, `cron`, `daemon`, `ftp`, `kern`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`, `lpr`, `mail`, `news`, `ntp`, `syslog`, `user`, `uucp`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"logging_level": schema.StringAttribute{
			MarkdownDescription: "Can be specified to filter outgoing messages. By default, all messages are sent. If a level is specified, only messages with a severity at least as important as this level will be sent. One of `alert`, `crit`, `debug`, `emerg`, `err`, `info`, `notice`, `warning`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"logging_length": schema.Int64Attribute{
			MarkdownDescription: "Specify an optional maximum line length in characters. Log lines larger than this value will be truncated before being sent. The reason is that syslog servers act differently on log line length. All servers support the default value of 1024 characters, but some servers simply drop larger lines while others do log them. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"stats_enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable HAProxy's statistics page.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"stats_port": schema.Int64Attribute{
			MarkdownDescription: "Choose a TCP port to be used for the local statistics page. The default value is 8822. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"stats_remote_enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable remote access to HAProxy's statistics page. This may be a security risk if you do not enable authentication! Note that you need to add appropiate firewall rules for this to work.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"stats_remote_bind": schema.SetAttribute{
			MarkdownDescription: "Configure listen addresses for the statistics page to enable remote access, i.e. 10.0.0.1:8080 or haproxy.example.com:8999. Use TAB key to complete typing a listen address.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"stats_auth_enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable authentication.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"stats_users": schema.SetAttribute{
			MarkdownDescription: "Users of the statistics page, as `user:password`.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"stats_allowed_users": schema.SetAttribute{
			MarkdownDescription: "Allowed Users.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"stats_allowed_groups": schema.SetAttribute{
			MarkdownDescription: "Allowed Groups.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"stats_custom_options": schema.StringAttribute{
			MarkdownDescription: "These lines will be added to the statistics settings of to the HAProxy configuration file.NOTE: The syntax will not be checked, use at your own risk!",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"stats_prometheus_enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable HAProxy's Prometheus exporter.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"stats_prometheus_bind": schema.SetAttribute{
			MarkdownDescription: "Configure listen addresses for the prometheus exporter, i.e. 10.0.0.1:8404 or haproxy.example.com:8404. Use TAB key to complete typing a listen address.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"stats_prometheus_path": schema.StringAttribute{
			MarkdownDescription: "The path where the Prometheus exporter can be accessed.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"cache_enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable HAProxy's cache which was designed to perform cache on small objects (favicon, css...). This is a minimalist low-maintenance cache which runs in RAM.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"cache_total_max_size": schema.Int64Attribute{
			MarkdownDescription: "Define the size in RAM of the cache in megabytes. This size is split in blocks of 1kB which are used by the cache entries. Its maximum value is 4095. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"cache_max_age": schema.Int64Attribute{
			MarkdownDescription: "Define the maximum expiration duration. Cache-Control response headers will be respected if they are less than this value. The default value is 60 seconds. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"cache_max_object_size": schema.Int64Attribute{
			MarkdownDescription: "Define the maximum size of the objects to be cached. Must not be greater than an half of the maximum size of the cache. If not set, it equals to a 256th of the cache size. All objects with sizes larger than this value will not be cached. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"cache_process_vary": schema.BoolAttribute{
			MarkdownDescription: "Enable or disable the processing of the Vary header. When disabled, a response containing such a header will never be cached.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"cache_max_secondary_entries": schema.Int64Attribute{
			MarkdownDescription: "To add this cache parameter Vary support needs to be enabled. Define the maximum number of simultaneous secondary entries with the same primary key in the cache. Its default value is 10 and should be passed a strictly positive integer. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `haproxy_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func settingsBool(b types.Bool) string {
	if b.IsNull() || b.IsUnknown() {
		return ""
	}
	return tools.BoolToString(b.ValueBool())
}

func settingsInt(i types.Int64) string {
	if i.IsNull() || i.IsUnknown() || i.ValueInt64() == -1 {
		return ""
	}
	return tools.Int64ToString(i.ValueInt64())
}

func settingsReadInt(s string) types.Int64 {
	if s == "" {
		return types.Int64Value(-1)
	}
	return tools.StringToInt64Null(s)
}

func settingsList(s types.Set) []string {
	out := tools.SetToStringSlice(s)
	if out == nil {
		return []string{}
	}
	return out
}

func settingsNonEmpty(in []string) []string {
	out := []string{}
	for _, v := range in {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

func convertSettingsSchemaToStruct(d *settingsResourceModel) *haproxy.Settings {
	return &haproxy.Settings{
		General: haproxy.SettingsGeneral{
			Enabled:         settingsBool(d.Enabled),
			GracefulStop:    settingsBool(d.GracefulStop),
			HardStopAfter:   d.HardStopAfter.ValueString(),
			CloseSpreadTime: d.CloseSpreadTime.ValueString(),
			SeamlessReload:  settingsBool(d.SeamlessReload),
			StoreOcsp:       settingsBool(d.StoreOcsp),
			Peers: haproxy.SettingsGeneralPeers{
				PeersEnabled: settingsBool(d.PeersEnabled),
				PeersName1:   d.PeersName1.ValueString(),
				PeersListen1: d.PeersListen1.ValueString(),
				PeersPort1:   settingsInt(d.PeersPort1),
				PeersName2:   d.PeersName2.ValueString(),
				PeersListen2: d.PeersListen2.ValueString(),
				PeersPort2:   settingsInt(d.PeersPort2),
			},
			Tuning: haproxy.SettingsGeneralTuning{
				TuningRoot:                           settingsBool(d.TuningRoot),
				TuningMaxConnections:                 settingsInt(d.TuningMaxConnections),
				TuningNbthread:                       settingsInt(d.TuningNbthread),
				TuningResolversPrefer:                api.SelectedMap(d.TuningResolversPrefer.ValueString()),
				TuningSslServerVerify:                api.SelectedMap(d.TuningSslServerVerify.ValueString()),
				TuningMaxDhSize:                      settingsInt(d.TuningMaxDhSize),
				TuningBufferSize:                     settingsInt(d.TuningBufferSize),
				TuningSpreadChecks:                   settingsInt(d.TuningSpreadChecks),
				TuningBogusProxyEnabled:              settingsBool(d.TuningBogusProxyEnabled),
				TuningLuaMaxMem:                      settingsInt(d.TuningLuaMaxMem),
				TuningCustomOptions:                  d.TuningCustomOptions.ValueString(),
				TuningOcspUpdateEnabled:              settingsBool(d.TuningOcspUpdateEnabled),
				TuningOcspUpdateMinDelay:             settingsInt(d.TuningOcspUpdateMinDelay),
				TuningOcspUpdateMaxDelay:             settingsInt(d.TuningOcspUpdateMaxDelay),
				TuningSslDefaultsEnabled:             settingsBool(d.TuningSslDefaultsEnabled),
				TuningSslBindOptions:                 api.SelectedMapList(settingsList(d.TuningSslBindOptions)),
				TuningSslMinVersion:                  api.SelectedMap(d.TuningSslMinVersion.ValueString()),
				TuningSslMaxVersion:                  api.SelectedMap(d.TuningSslMaxVersion.ValueString()),
				TuningSslCipherList:                  d.TuningSslCipherList.ValueString(),
				TuningSslCipherSuites:                d.TuningSslCipherSuites.ValueString(),
				TuningH2InitialWindowSize:            settingsInt(d.TuningH2InitialWindowSize),
				TuningH2InitialWindowSizeOutgoing:    settingsInt(d.TuningH2InitialWindowSizeOutgoing),
				TuningH2InitialWindowSizeIncoming:    settingsInt(d.TuningH2InitialWindowSizeIncoming),
				TuningH2MaxConcurrentStreams:         settingsInt(d.TuningH2MaxConcurrentStreams),
				TuningH2MaxConcurrentStreamsOutgoing: settingsInt(d.TuningH2MaxConcurrentStreamsOutgoing),
				TuningH2MaxConcurrentStreamsIncoming: settingsInt(d.TuningH2MaxConcurrentStreamsIncoming),
			},
			Defaults: haproxy.SettingsGeneralDefaults{
				DefaultsMaxConnections:        settingsInt(d.DefaultsMaxConnections),
				DefaultsMaxConnectionsServers: settingsInt(d.DefaultsMaxConnectionsServers),
				DefaultsTimeoutClient:         d.DefaultsTimeoutClient.ValueString(),
				DefaultsTimeoutConnect:        d.DefaultsTimeoutConnect.ValueString(),
				DefaultsTimeoutCheck:          d.DefaultsTimeoutCheck.ValueString(),
				DefaultsTimeoutServer:         d.DefaultsTimeoutServer.ValueString(),
				DefaultsRetries:               settingsInt(d.DefaultsRetries),
				DefaultsRedispatch:            api.SelectedMap(d.DefaultsRedispatch.ValueString()),
				DefaultsInitAddr:              api.SelectedMapList(settingsList(d.DefaultsInitAddr)),
				DefaultsCustomOptions:         d.DefaultsCustomOptions.ValueString(),
			},
			Logging: haproxy.SettingsGeneralLogging{
				LoggingHost:     d.LoggingHost.ValueString(),
				LoggingFacility: api.SelectedMap(d.LoggingFacility.ValueString()),
				LoggingLevel:    api.SelectedMap(d.LoggingLevel.ValueString()),
				LoggingLength:   settingsInt(d.LoggingLength),
			},
			Stats: haproxy.SettingsGeneralStats{
				StatsEnabled:           settingsBool(d.StatsEnabled),
				StatsPort:              settingsInt(d.StatsPort),
				StatsRemoteEnabled:     settingsBool(d.StatsRemoteEnabled),
				StatsRemoteBind:        api.SelectedMapList(settingsList(d.StatsRemoteBind)),
				StatsAuthEnabled:       settingsBool(d.StatsAuthEnabled),
				StatsUsers:             api.SelectedMapList(settingsList(d.StatsUsers)),
				StatsAllowedUsers:      api.SelectedMapList(settingsList(d.StatsAllowedUsers)),
				StatsAllowedGroups:     api.SelectedMapList(settingsList(d.StatsAllowedGroups)),
				StatsCustomOptions:     d.StatsCustomOptions.ValueString(),
				StatsPrometheusEnabled: settingsBool(d.StatsPrometheusEnabled),
				StatsPrometheusBind:    api.SelectedMapList(settingsList(d.StatsPrometheusBind)),
				StatsPrometheusPath:    d.StatsPrometheusPath.ValueString(),
			},
			Cache: haproxy.SettingsGeneralCache{
				CacheEnabled:             settingsBool(d.CacheEnabled),
				CacheTotalMaxSize:        settingsInt(d.CacheTotalMaxSize),
				CacheMaxAge:              settingsInt(d.CacheMaxAge),
				CacheMaxObjectSize:       settingsInt(d.CacheMaxObjectSize),
				CacheProcessVary:         settingsBool(d.CacheProcessVary),
				CacheMaxSecondaryEntries: settingsInt(d.CacheMaxSecondaryEntries),
			},
		},
	}
}

func convertSettingsStructToSchema(s *haproxy.Settings) *settingsResourceModel {
	return &settingsResourceModel{
		Enabled:                              types.BoolValue(tools.StringToBool(s.General.Enabled)),
		GracefulStop:                         types.BoolValue(tools.StringToBool(s.General.GracefulStop)),
		HardStopAfter:                        types.StringValue(s.General.HardStopAfter),
		CloseSpreadTime:                      types.StringValue(s.General.CloseSpreadTime),
		SeamlessReload:                       types.BoolValue(tools.StringToBool(s.General.SeamlessReload)),
		StoreOcsp:                            types.BoolValue(tools.StringToBool(s.General.StoreOcsp)),
		PeersEnabled:                         types.BoolValue(tools.StringToBool(s.General.Peers.PeersEnabled)),
		PeersName1:                           types.StringValue(s.General.Peers.PeersName1),
		PeersListen1:                         types.StringValue(s.General.Peers.PeersListen1),
		PeersPort1:                           settingsReadInt(s.General.Peers.PeersPort1),
		PeersName2:                           types.StringValue(s.General.Peers.PeersName2),
		PeersListen2:                         types.StringValue(s.General.Peers.PeersListen2),
		PeersPort2:                           settingsReadInt(s.General.Peers.PeersPort2),
		TuningRoot:                           types.BoolValue(tools.StringToBool(s.General.Tuning.TuningRoot)),
		TuningMaxConnections:                 settingsReadInt(s.General.Tuning.TuningMaxConnections),
		TuningNbthread:                       settingsReadInt(s.General.Tuning.TuningNbthread),
		TuningResolversPrefer:                types.StringValue(s.General.Tuning.TuningResolversPrefer.String()),
		TuningSslServerVerify:                types.StringValue(s.General.Tuning.TuningSslServerVerify.String()),
		TuningMaxDhSize:                      settingsReadInt(s.General.Tuning.TuningMaxDhSize),
		TuningBufferSize:                     settingsReadInt(s.General.Tuning.TuningBufferSize),
		TuningSpreadChecks:                   settingsReadInt(s.General.Tuning.TuningSpreadChecks),
		TuningBogusProxyEnabled:              types.BoolValue(tools.StringToBool(s.General.Tuning.TuningBogusProxyEnabled)),
		TuningLuaMaxMem:                      settingsReadInt(s.General.Tuning.TuningLuaMaxMem),
		TuningCustomOptions:                  types.StringValue(s.General.Tuning.TuningCustomOptions),
		TuningOcspUpdateEnabled:              types.BoolValue(tools.StringToBool(s.General.Tuning.TuningOcspUpdateEnabled)),
		TuningOcspUpdateMinDelay:             settingsReadInt(s.General.Tuning.TuningOcspUpdateMinDelay),
		TuningOcspUpdateMaxDelay:             settingsReadInt(s.General.Tuning.TuningOcspUpdateMaxDelay),
		TuningSslDefaultsEnabled:             types.BoolValue(tools.StringToBool(s.General.Tuning.TuningSslDefaultsEnabled)),
		TuningSslBindOptions:                 tools.StringSliceToSet(settingsNonEmpty(s.General.Tuning.TuningSslBindOptions)),
		TuningSslMinVersion:                  types.StringValue(s.General.Tuning.TuningSslMinVersion.String()),
		TuningSslMaxVersion:                  types.StringValue(s.General.Tuning.TuningSslMaxVersion.String()),
		TuningSslCipherList:                  types.StringValue(s.General.Tuning.TuningSslCipherList),
		TuningSslCipherSuites:                types.StringValue(s.General.Tuning.TuningSslCipherSuites),
		TuningH2InitialWindowSize:            settingsReadInt(s.General.Tuning.TuningH2InitialWindowSize),
		TuningH2InitialWindowSizeOutgoing:    settingsReadInt(s.General.Tuning.TuningH2InitialWindowSizeOutgoing),
		TuningH2InitialWindowSizeIncoming:    settingsReadInt(s.General.Tuning.TuningH2InitialWindowSizeIncoming),
		TuningH2MaxConcurrentStreams:         settingsReadInt(s.General.Tuning.TuningH2MaxConcurrentStreams),
		TuningH2MaxConcurrentStreamsOutgoing: settingsReadInt(s.General.Tuning.TuningH2MaxConcurrentStreamsOutgoing),
		TuningH2MaxConcurrentStreamsIncoming: settingsReadInt(s.General.Tuning.TuningH2MaxConcurrentStreamsIncoming),
		DefaultsMaxConnections:               settingsReadInt(s.General.Defaults.DefaultsMaxConnections),
		DefaultsMaxConnectionsServers:        settingsReadInt(s.General.Defaults.DefaultsMaxConnectionsServers),
		DefaultsTimeoutClient:                types.StringValue(s.General.Defaults.DefaultsTimeoutClient),
		DefaultsTimeoutConnect:               types.StringValue(s.General.Defaults.DefaultsTimeoutConnect),
		DefaultsTimeoutCheck:                 types.StringValue(s.General.Defaults.DefaultsTimeoutCheck),
		DefaultsTimeoutServer:                types.StringValue(s.General.Defaults.DefaultsTimeoutServer),
		DefaultsRetries:                      settingsReadInt(s.General.Defaults.DefaultsRetries),
		DefaultsRedispatch:                   types.StringValue(s.General.Defaults.DefaultsRedispatch.String()),
		DefaultsInitAddr:                     tools.StringSliceToSet(settingsNonEmpty(s.General.Defaults.DefaultsInitAddr)),
		DefaultsCustomOptions:                types.StringValue(s.General.Defaults.DefaultsCustomOptions),
		LoggingHost:                          types.StringValue(s.General.Logging.LoggingHost),
		LoggingFacility:                      types.StringValue(s.General.Logging.LoggingFacility.String()),
		LoggingLevel:                         types.StringValue(s.General.Logging.LoggingLevel.String()),
		LoggingLength:                        settingsReadInt(s.General.Logging.LoggingLength),
		StatsEnabled:                         types.BoolValue(tools.StringToBool(s.General.Stats.StatsEnabled)),
		StatsPort:                            settingsReadInt(s.General.Stats.StatsPort),
		StatsRemoteEnabled:                   types.BoolValue(tools.StringToBool(s.General.Stats.StatsRemoteEnabled)),
		StatsRemoteBind:                      tools.StringSliceToSet(settingsNonEmpty(s.General.Stats.StatsRemoteBind)),
		StatsAuthEnabled:                     types.BoolValue(tools.StringToBool(s.General.Stats.StatsAuthEnabled)),
		StatsUsers:                           tools.StringSliceToSet(settingsNonEmpty(s.General.Stats.StatsUsers)),
		StatsAllowedUsers:                    tools.StringSliceToSet(settingsNonEmpty(s.General.Stats.StatsAllowedUsers)),
		StatsAllowedGroups:                   tools.StringSliceToSet(settingsNonEmpty(s.General.Stats.StatsAllowedGroups)),
		StatsCustomOptions:                   types.StringValue(s.General.Stats.StatsCustomOptions),
		StatsPrometheusEnabled:               types.BoolValue(tools.StringToBool(s.General.Stats.StatsPrometheusEnabled)),
		StatsPrometheusBind:                  tools.StringSliceToSet(settingsNonEmpty(s.General.Stats.StatsPrometheusBind)),
		StatsPrometheusPath:                  types.StringValue(s.General.Stats.StatsPrometheusPath),
		CacheEnabled:                         types.BoolValue(tools.StringToBool(s.General.Cache.CacheEnabled)),
		CacheTotalMaxSize:                    settingsReadInt(s.General.Cache.CacheTotalMaxSize),
		CacheMaxAge:                          settingsReadInt(s.General.Cache.CacheMaxAge),
		CacheMaxObjectSize:                   settingsReadInt(s.General.Cache.CacheMaxObjectSize),
		CacheProcessVary:                     types.BoolValue(tools.StringToBool(s.General.Cache.CacheProcessVary)),
		CacheMaxSecondaryEntries:             settingsReadInt(s.General.Cache.CacheMaxSecondaryEntries),
	}
}
