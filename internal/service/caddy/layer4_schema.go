package caddy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/caddy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type layer4ResourceModel struct {
	Enabled                               types.Bool   `tfsdk:"enabled"`
	Sequence                              types.Int64  `tfsdk:"sequence"`
	Type                                  types.String `tfsdk:"type"`
	Protocol                              types.String `tfsdk:"protocol"`
	FromPort                              types.Int64  `tfsdk:"from_port"`
	FromDomain                            types.Set    `tfsdk:"from_domain"`
	FromOpenvpnModes                      types.String `tfsdk:"from_openvpn_modes"`
	FromOpenvpnStaticKey                  types.Set    `tfsdk:"from_openvpn_static_key"`
	Matchers                              types.String `tfsdk:"matchers"`
	InvertMatchers                        types.Bool   `tfsdk:"invert_matchers"`
	ToDomain                              types.Set    `tfsdk:"to_domain"`
	TerminateTLS                          types.Bool   `tfsdk:"terminate_tls"`
	ToPort                                types.Int64  `tfsdk:"to_port"`
	OriginateTLS                          types.String `tfsdk:"originate_tls"`
	ProxyProtocol                         types.String `tfsdk:"proxy_protocol"`
	LbPolicy                              types.String `tfsdk:"lb_policy"`
	LbTryDuration                         types.Int64  `tfsdk:"lb_try_duration"`
	LbTryInterval                         types.Int64  `tfsdk:"lb_try_interval"`
	ActiveHealthPort                      types.Int64  `tfsdk:"active_health_port"`
	ActiveHealthInterval                  types.Int64  `tfsdk:"active_health_interval"`
	ActiveHealthTimeout                   types.Int64  `tfsdk:"active_health_timeout"`
	PassiveHealthFailDuration             types.Int64  `tfsdk:"passive_health_fail_duration"`
	PassiveHealthMaxFails                 types.Int64  `tfsdk:"passive_health_max_fails"`
	PassiveHealthUnhealthyConnectionCount types.Int64  `tfsdk:"passive_health_unhealthy_connection_count"`
	RemoteIP                              types.Set    `tfsdk:"remote_ip"`
	Description                           types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func layer4ResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Caddy layer 4 routes (**Services ▸ Caddy Web Server ▸ Layer 4 Proxy**, plugin `os-caddy`): TCP/UDP proxying by SNI, protocol matchers or port.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this Layer4 route. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"sequence": schema.Int64Attribute{
				MarkdownDescription: "Rules are sorted based on the sequence number, a higher number means lower priority. Rules without a sequence number will be processed first.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Choose either \"listener_wrappers\" for multiplexing protocols on the default HTTP and HTTPS ports on OSI Layer 7, or \"global\" for raw TCP/UDP traffic routing on a custom \"Local port\" on OSI Layer 4 with optional OSI Layer 7 protocol matching. One of `listener_wrappers`, `global`. Defaults to `listener_wrappers`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("listener_wrappers"),
				Validators:          []validator.String{stringvalidator.OneOf("listener_wrappers", "global")},
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: "Match the received traffic on OSI Layer 4, either TCP or UDP. When \"Routing Type\" is \"listener_wrappers\", currently only TCP will match. One of `tcp`, `udp`. Defaults to `tcp`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("tcp"),
				Validators:          []validator.String{stringvalidator.OneOf("tcp", "udp")},
			},
			"from_port": schema.Int64Attribute{
				MarkdownDescription: "Choose a custom local port to listen on. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"from_domain": schema.SetAttribute{
				MarkdownDescription: "Enter one or multiple domains to route via SNI or Host Header. Wildcard domains and host wildcards are allowed, e.g. \"*.example.com\" and \"*\". Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"from_openvpn_modes": schema.StringAttribute{
				MarkdownDescription: "Select the mode matching the OpenVPN Server or Client. One of `auth_sha256_normal`, `auth_sha256_inverse`, `auth_sha512_normal`, `auth_sha512_inverse`, `crypt`, `crypt2_client`, `crypt2_server`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "auth_sha256_normal", "auth_sha256_inverse", "auth_sha512_normal", "auth_sha512_inverse", "crypt", "crypt2_client", "crypt2_server")},
			},
			"from_openvpn_static_key": schema.SetAttribute{
				MarkdownDescription: "Select a Static Key to match. Multiple Static Keys are only supported for tls-crypt2_client mode. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"matchers": schema.StringAttribute{
				MarkdownDescription: "Match the received traffic on OSI Layer 7. When choosing a protocol like TLS, it will be matched and routed to the upstream destination. When \"Routing Type\" of the matcher is \"listener_wrapper\", any unmatched traffic will be received by the \"HTTP App\" (reverse_proxy). When \"Routing Type\" of the matcher is \"global\", unmatched traffic will be consumed and blocked. Choose \"ANY\" to not match on OSI Layer 7 and allow any traffic. One of `any`, `dns`, `http`, `httphost`, `openvpn`, `postgres`, `proxy_protocol`, `quic`, `quicsni`, `rdp`, `socks4`, `socks5`, `ssh`, `tls`, `tlssni`, `winbox`, `wireguard`, `xmpp`. Defaults to `tlssni`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("tlssni"),
				Validators:          []validator.String{stringvalidator.OneOf("any", "dns", "http", "httphost", "openvpn", "postgres", "proxy_protocol", "quic", "quicsni", "rdp", "socks4", "socks5", "ssh", "tls", "tlssni", "winbox", "wireguard", "xmpp")},
			},
			"invert_matchers": schema.BoolAttribute{
				MarkdownDescription: "Invert the sense of the matcher. E.g., if the protocol is TLS, inverting will match all traffic that is not TLS. When domains have been chosen, these will be equally inverted. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"to_domain": schema.SetAttribute{
				MarkdownDescription: "Enter a domain name or IP address of the upstream destination. If multiple are chosen, they will be load balanced with the default random policy.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"terminate_tls": schema.BoolAttribute{
				MarkdownDescription: "Terminate TLS before routing to the upstream. Since this requires a certificate, ensure there is a domain configured in \"Reverse Proxy\" that matches the SNI of \"Domain\". The best matching SAN or wildcard certificate will be used automatically for this route. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"to_port": schema.Int64Attribute{
				MarkdownDescription: "Choose a custom port for the upstream destination.",
				Required:            true,
			},
			"originate_tls": schema.StringAttribute{
				MarkdownDescription: "Setting this option will establish a new TLS connection to the upstream after it has been terminated. One of `tls`, `tls_insecure_skip_verify`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "tls", "tls_insecure_skip_verify")},
			},
			"proxy_protocol": schema.StringAttribute{
				MarkdownDescription: "Add the HA Proxy Protocol header. Either version 1 or 2 can be chosen. The default is off, since it is only needed when the upstream can use the Proxy Protocol header. One of `v1`, `v2`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "v1", "v2")},
			},
			"lb_policy": schema.StringAttribute{
				MarkdownDescription: "lb_policy is the name of the load balancing policy, along with any options. For policies that involve hashing, the highest-random-weight (HRW) algorithm is used to ensure that a client or request with the same hash key is mapped to the same upstream, even if the list of upstreams change. One of `first`, `ip_hash`, `least_conn`, `random_choose`, `round_robin`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "first", "ip_hash", "least_conn", "random_choose", "round_robin")},
			},
			"lb_try_duration": schema.Int64Attribute{
				MarkdownDescription: "lb_try_duration is a duration value that defines how long to try selecting available backends for each request if the next available host is down. Clients will wait for up to this long while the load balancer tries to find an available upstream host. A reasonable starting point might be 5s since the HTTP transport's default dial timeout is 3s, so that should allow for at least one retry if the first selected upstream cannot be reached. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"lb_try_interval": schema.Int64Attribute{
				MarkdownDescription: "lb_try_interval is a duration value that defines how long to wait between selecting the next host from the pool. Only relevant when a request to an upstream host fails. Be aware that setting this to 0 with a non-zero lb_try_duration can cause the CPU to spin if all backends are down and latency is very low. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"active_health_port": schema.Int64Attribute{
				MarkdownDescription: "health_port is the port to use for active health checks, if different from the upstream's port. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"active_health_interval": schema.Int64Attribute{
				MarkdownDescription: "health_interval is a duration value that defines how often to perform active health checks. Default: 30s. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"active_health_timeout": schema.Int64Attribute{
				MarkdownDescription: "health_timeout is a duration value that defines how long to wait for a reply before marking the backend as down. Default: 5s. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"passive_health_fail_duration": schema.Int64Attribute{
				MarkdownDescription: "fail_duration enables a passive health check when multiple destinations in \"Upstream Domain\" are set. It is a value that defines how long to remember a failed request. A duration of 1 or more seconds enables passive health checking. A reasonable starting point might be 30s to balance error rates with responsiveness when bringing an unhealthy upstream back online. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"passive_health_max_fails": schema.Int64Attribute{
				MarkdownDescription: "max_fails is the maximum number of failed requests within fail_duration that are needed before considering an upstream to be down. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"passive_health_unhealthy_connection_count": schema.Int64Attribute{
				MarkdownDescription: "unhealthy_connection_count is the permissible number of simultaneous connections to a backend before marking it as down. In other words, if a particular backend is currently handling this many connections, then it is considered \"overloaded\" and other backends will be preferred instead. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"remote_ip": schema.SetAttribute{
				MarkdownDescription: "Enter one or multiple IP addresses and networks. Leaving this empty will allow any remote client to connect. If populated and the remote IP address of a client matches, the connection to the \"Upstream Domain\" is allowed, otherwise the connection is dropped. Due to restrictions with this matcher, this list can not be inverted. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Enter a description for this Layer4 route. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the Caddy layer 4 route.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func layer4DataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Caddy layer 4 routes (**Services ▸ Caddy Web Server ▸ Layer 4 Proxy**, plugin `os-caddy`): TCP/UDP proxying by SNI, protocol matchers or port.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the Caddy layer 4 route.",
				Required:            true,
			},
			"enabled":                      dschema.BoolAttribute{MarkdownDescription: "Enable this Layer4 route.", Computed: true},
			"sequence":                     dschema.Int64Attribute{MarkdownDescription: "Rules are sorted based on the sequence number, a higher number means lower priority. Rules without a sequence number will be processed first.", Computed: true},
			"type":                         dschema.StringAttribute{MarkdownDescription: "Choose either \"listener_wrappers\" for multiplexing protocols on the default HTTP and HTTPS ports on OSI Layer 7, or \"global\" for raw TCP/UDP traffic routing on a custom \"Local port\" on OSI Layer 4 with optional OSI Layer 7 protocol matching. One of `listener_wrappers`, `global`.", Computed: true},
			"protocol":                     dschema.StringAttribute{MarkdownDescription: "Match the received traffic on OSI Layer 4, either TCP or UDP. When \"Routing Type\" is \"listener_wrappers\", currently only TCP will match. One of `tcp`, `udp`.", Computed: true},
			"from_port":                    dschema.Int64Attribute{MarkdownDescription: "Choose a custom local port to listen on.", Computed: true},
			"from_domain":                  dschema.SetAttribute{MarkdownDescription: "Enter one or multiple domains to route via SNI or Host Header. Wildcard domains and host wildcards are allowed, e.g. \"*.example.com\" and \"*\".", Computed: true, ElementType: types.StringType},
			"from_openvpn_modes":           dschema.StringAttribute{MarkdownDescription: "Select the mode matching the OpenVPN Server or Client. One of `auth_sha256_normal`, `auth_sha256_inverse`, `auth_sha512_normal`, `auth_sha512_inverse`, `crypt`, `crypt2_client`, `crypt2_server`.", Computed: true},
			"from_openvpn_static_key":      dschema.SetAttribute{MarkdownDescription: "Select a Static Key to match. Multiple Static Keys are only supported for tls-crypt2_client mode.", Computed: true, ElementType: types.StringType},
			"matchers":                     dschema.StringAttribute{MarkdownDescription: "Match the received traffic on OSI Layer 7. When choosing a protocol like TLS, it will be matched and routed to the upstream destination. When \"Routing Type\" of the matcher is \"listener_wrapper\", any unmatched traffic will be received by the \"HTTP App\" (reverse_proxy). When \"Routing Type\" of the matcher is \"global\", unmatched traffic will be consumed and blocked. Choose \"ANY\" to not match on OSI Layer 7 and allow any traffic. One of `any`, `dns`, `http`, `httphost`, `openvpn`, `postgres`, `proxy_protocol`, `quic`, `quicsni`, `rdp`, `socks4`, `socks5`, `ssh`, `tls`, `tlssni`, `winbox`, `wireguard`, `xmpp`.", Computed: true},
			"invert_matchers":              dschema.BoolAttribute{MarkdownDescription: "Invert the sense of the matcher. E.g., if the protocol is TLS, inverting will match all traffic that is not TLS. When domains have been chosen, these will be equally inverted.", Computed: true},
			"to_domain":                    dschema.SetAttribute{MarkdownDescription: "Enter a domain name or IP address of the upstream destination. If multiple are chosen, they will be load balanced with the default random policy.", Computed: true, ElementType: types.StringType},
			"terminate_tls":                dschema.BoolAttribute{MarkdownDescription: "Terminate TLS before routing to the upstream. Since this requires a certificate, ensure there is a domain configured in \"Reverse Proxy\" that matches the SNI of \"Domain\". The best matching SAN or wildcard certificate will be used automatically for this route.", Computed: true},
			"to_port":                      dschema.Int64Attribute{MarkdownDescription: "Choose a custom port for the upstream destination.", Computed: true},
			"originate_tls":                dschema.StringAttribute{MarkdownDescription: "Setting this option will establish a new TLS connection to the upstream after it has been terminated. One of `tls`, `tls_insecure_skip_verify`.", Computed: true},
			"proxy_protocol":               dschema.StringAttribute{MarkdownDescription: "Add the HA Proxy Protocol header. Either version 1 or 2 can be chosen. The default is off, since it is only needed when the upstream can use the Proxy Protocol header. One of `v1`, `v2`.", Computed: true},
			"lb_policy":                    dschema.StringAttribute{MarkdownDescription: "lb_policy is the name of the load balancing policy, along with any options. For policies that involve hashing, the highest-random-weight (HRW) algorithm is used to ensure that a client or request with the same hash key is mapped to the same upstream, even if the list of upstreams change. One of `first`, `ip_hash`, `least_conn`, `random_choose`, `round_robin`.", Computed: true},
			"lb_try_duration":              dschema.Int64Attribute{MarkdownDescription: "lb_try_duration is a duration value that defines how long to try selecting available backends for each request if the next available host is down. Clients will wait for up to this long while the load balancer tries to find an available upstream host. A reasonable starting point might be 5s since the HTTP transport's default dial timeout is 3s, so that should allow for at least one retry if the first selected upstream cannot be reached.", Computed: true},
			"lb_try_interval":              dschema.Int64Attribute{MarkdownDescription: "lb_try_interval is a duration value that defines how long to wait between selecting the next host from the pool. Only relevant when a request to an upstream host fails. Be aware that setting this to 0 with a non-zero lb_try_duration can cause the CPU to spin if all backends are down and latency is very low.", Computed: true},
			"active_health_port":           dschema.Int64Attribute{MarkdownDescription: "health_port is the port to use for active health checks, if different from the upstream's port.", Computed: true},
			"active_health_interval":       dschema.Int64Attribute{MarkdownDescription: "health_interval is a duration value that defines how often to perform active health checks. Default: 30s.", Computed: true},
			"active_health_timeout":        dschema.Int64Attribute{MarkdownDescription: "health_timeout is a duration value that defines how long to wait for a reply before marking the backend as down. Default: 5s.", Computed: true},
			"passive_health_fail_duration": dschema.Int64Attribute{MarkdownDescription: "fail_duration enables a passive health check when multiple destinations in \"Upstream Domain\" are set. It is a value that defines how long to remember a failed request. A duration of 1 or more seconds enables passive health checking. A reasonable starting point might be 30s to balance error rates with responsiveness when bringing an unhealthy upstream back online.", Computed: true},
			"passive_health_max_fails":     dschema.Int64Attribute{MarkdownDescription: "max_fails is the maximum number of failed requests within fail_duration that are needed before considering an upstream to be down.", Computed: true},
			"passive_health_unhealthy_connection_count": dschema.Int64Attribute{MarkdownDescription: "unhealthy_connection_count is the permissible number of simultaneous connections to a backend before marking it as down. In other words, if a particular backend is currently handling this many connections, then it is considered \"overloaded\" and other backends will be preferred instead.", Computed: true},
			"remote_ip":   dschema.SetAttribute{MarkdownDescription: "Enter one or multiple IP addresses and networks. Leaving this empty will allow any remote client to connect. If populated and the remote IP address of a client matches, the connection to the \"Upstream Domain\" is allowed, otherwise the connection is dropped. Due to restrictions with this matcher, this list can not be inverted.", Computed: true, ElementType: types.StringType},
			"description": dschema.StringAttribute{MarkdownDescription: "Enter a description for this Layer4 route.", Computed: true},
		},
	}
}

func convertLayer4SchemaToStruct(d *layer4ResourceModel) (*caddy.Layer4, error) {
	return &caddy.Layer4{
		Enabled: tools.BoolToString(d.Enabled.ValueBool()),
		Sequence: func() string {
			if d.Sequence.IsNull() || d.Sequence.IsUnknown() {
				return ""
			}
			return tools.Int64ToString(d.Sequence.ValueInt64())
		}(),
		Type:                                  api.SelectedMap(d.Type.ValueString()),
		Protocol:                              api.SelectedMap(d.Protocol.ValueString()),
		FromPort:                              tools.Int64ToStringNegative(d.FromPort.ValueInt64()),
		FromDomain:                            api.SelectedMapList(setToSlice(d.FromDomain)),
		FromOpenvpnModes:                      api.SelectedMap(d.FromOpenvpnModes.ValueString()),
		FromOpenvpnStaticKey:                  api.SelectedMapList(setToSlice(d.FromOpenvpnStaticKey)),
		Matchers:                              api.SelectedMap(d.Matchers.ValueString()),
		InvertMatchers:                        tools.BoolToString(d.InvertMatchers.ValueBool()),
		ToDomain:                              api.SelectedMapList(setToSlice(d.ToDomain)),
		TerminateTls:                          tools.BoolToString(d.TerminateTLS.ValueBool()),
		ToPort:                                tools.Int64ToString(d.ToPort.ValueInt64()),
		OriginateTls:                          api.SelectedMap(d.OriginateTLS.ValueString()),
		ProxyProtocol:                         api.SelectedMap(d.ProxyProtocol.ValueString()),
		LbPolicy:                              api.SelectedMap(d.LbPolicy.ValueString()),
		LbTryDuration:                         tools.Int64ToStringNegative(d.LbTryDuration.ValueInt64()),
		LbTryInterval:                         tools.Int64ToStringNegative(d.LbTryInterval.ValueInt64()),
		ActiveHealthPort:                      tools.Int64ToStringNegative(d.ActiveHealthPort.ValueInt64()),
		ActiveHealthInterval:                  tools.Int64ToStringNegative(d.ActiveHealthInterval.ValueInt64()),
		ActiveHealthTimeout:                   tools.Int64ToStringNegative(d.ActiveHealthTimeout.ValueInt64()),
		PassiveHealthFailDuration:             tools.Int64ToStringNegative(d.PassiveHealthFailDuration.ValueInt64()),
		PassiveHealthMaxFails:                 tools.Int64ToStringNegative(d.PassiveHealthMaxFails.ValueInt64()),
		PassiveHealthUnhealthyConnectionCount: tools.Int64ToStringNegative(d.PassiveHealthUnhealthyConnectionCount.ValueInt64()),
		RemoteIp:                              api.SelectedMapList(setToSlice(d.RemoteIP)),
		Description:                           d.Description.ValueString(),
	}, nil
}

func convertLayer4StructToSchema(d *caddy.Layer4) (*layer4ResourceModel, error) {
	return &layer4ResourceModel{
		Enabled:                               types.BoolValue(tools.StringToBool(d.Enabled)),
		Sequence:                              tools.StringToInt64Null(d.Sequence),
		Type:                                  types.StringValue(d.Type.String()),
		Protocol:                              types.StringValue(d.Protocol.String()),
		FromPort:                              types.Int64Value(tools.StringToInt64(d.FromPort)),
		FromDomain:                            tools.StringSliceToSet(d.FromDomain),
		FromOpenvpnModes:                      types.StringValue(d.FromOpenvpnModes.String()),
		FromOpenvpnStaticKey:                  tools.StringSliceToSet(d.FromOpenvpnStaticKey),
		Matchers:                              types.StringValue(d.Matchers.String()),
		InvertMatchers:                        types.BoolValue(tools.StringToBool(d.InvertMatchers)),
		ToDomain:                              tools.StringSliceToSet(d.ToDomain),
		TerminateTLS:                          types.BoolValue(tools.StringToBool(d.TerminateTls)),
		ToPort:                                types.Int64Value(tools.StringToInt64(d.ToPort)),
		OriginateTLS:                          types.StringValue(d.OriginateTls.String()),
		ProxyProtocol:                         types.StringValue(d.ProxyProtocol.String()),
		LbPolicy:                              types.StringValue(d.LbPolicy.String()),
		LbTryDuration:                         types.Int64Value(tools.StringToInt64(d.LbTryDuration)),
		LbTryInterval:                         types.Int64Value(tools.StringToInt64(d.LbTryInterval)),
		ActiveHealthPort:                      types.Int64Value(tools.StringToInt64(d.ActiveHealthPort)),
		ActiveHealthInterval:                  types.Int64Value(tools.StringToInt64(d.ActiveHealthInterval)),
		ActiveHealthTimeout:                   types.Int64Value(tools.StringToInt64(d.ActiveHealthTimeout)),
		PassiveHealthFailDuration:             types.Int64Value(tools.StringToInt64(d.PassiveHealthFailDuration)),
		PassiveHealthMaxFails:                 types.Int64Value(tools.StringToInt64(d.PassiveHealthMaxFails)),
		PassiveHealthUnhealthyConnectionCount: types.Int64Value(tools.StringToInt64(d.PassiveHealthUnhealthyConnectionCount)),
		RemoteIP:                              tools.StringSliceToSet(d.RemoteIp),
		Description:                           types.StringValue(d.Description),
	}, nil
}
