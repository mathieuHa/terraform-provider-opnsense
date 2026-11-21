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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type handleResourceModel struct {
	Enabled                            types.Bool   `tfsdk:"enabled"`
	Reverse                            types.String `tfsdk:"reverse"`
	Subdomain                          types.String `tfsdk:"subdomain"`
	HandleType                         types.String `tfsdk:"handle_type"`
	HandlePath                         types.String `tfsdk:"handle_path"`
	Accesslist                         types.String `tfsdk:"accesslist"`
	Basicauth                          types.Set    `tfsdk:"basicauth"`
	Header                             types.Set    `tfsdk:"header"`
	HandleDirective                    types.String `tfsdk:"handle_directive"`
	ToDomain                           types.Set    `tfsdk:"to_domain"`
	ToPort                             types.Int64  `tfsdk:"to_port"`
	ToPath                             types.String `tfsdk:"to_path"`
	ForwardAuth                        types.Bool   `tfsdk:"forward_auth"`
	HttpTLS                            types.String `tfsdk:"http_tls"`
	HttpVersion                        types.String `tfsdk:"http_version"`
	HttpKeepalive                      types.Int64  `tfsdk:"http_keepalive"`
	HttpTLSInsecureSkipVerify          types.Bool   `tfsdk:"http_tls_insecure_skip_verify"`
	HttpTLSTrustedCaCerts              types.String `tfsdk:"http_tls_trusted_ca_certs"`
	HttpTLSServerName                  types.String `tfsdk:"http_tls_server_name"`
	LbPolicy                           types.String `tfsdk:"lb_policy"`
	LbRetries                          types.Int64  `tfsdk:"lb_retries"`
	LbTryDuration                      types.Int64  `tfsdk:"lb_try_duration"`
	LbTryInterval                      types.Int64  `tfsdk:"lb_try_interval"`
	PassiveHealthFailDuration          types.Int64  `tfsdk:"passive_health_fail_duration"`
	PassiveHealthMaxFails              types.Int64  `tfsdk:"passive_health_max_fails"`
	PassiveHealthUnhealthyStatus       types.String `tfsdk:"passive_health_unhealthy_status"`
	PassiveHealthUnhealthyLatency      types.Int64  `tfsdk:"passive_health_unhealthy_latency"`
	PassiveHealthUnhealthyRequestCount types.Int64  `tfsdk:"passive_health_unhealthy_request_count"`
	Description                        types.String `tfsdk:"description"`
	HealthUri                          types.String `tfsdk:"health_uri"`
	HealthUpstream                     types.String `tfsdk:"health_upstream"`
	HealthPort                         types.Int64  `tfsdk:"health_port"`
	HealthInterval                     types.Int64  `tfsdk:"health_interval"`
	HealthPasses                       types.Int64  `tfsdk:"health_passes"`
	HealthFails                        types.Int64  `tfsdk:"health_fails"`
	HealthTimeout                      types.Int64  `tfsdk:"health_timeout"`
	HealthStatus                       types.String `tfsdk:"health_status"`
	HealthBody                         types.String `tfsdk:"health_body"`
	HealthFollowRedirects              types.Bool   `tfsdk:"health_follow_redirects"`

	Id types.String `tfsdk:"id"`
}

func handleResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Caddy handlers (**Services ▸ Caddy Web Server ▸ Reverse Proxy ▸ Handlers**, plugin `os-caddy`): route the requests of a domain or subdomain, optionally by path, to upstream servers.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this handler. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"reverse": schema.StringAttribute{
				MarkdownDescription: "Select a domain to handle.",
				Required:            true,
			},
			"subdomain": schema.StringAttribute{
				MarkdownDescription: "Select a subdomain to handle. Make sure to additionally choose a wildcard domain as \"Domain\". Leave unset, if not using subdomains. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"handle_type": schema.StringAttribute{
				MarkdownDescription: "Choose a handling type: \"handle\" (default) will keep the path in all requests. \"handle_path\" will strip the path from all requests. One of `handle`, `handle_path`. Defaults to `handle`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("handle"),
				Validators:          []validator.String{stringvalidator.OneOf("handle", "handle_path")},
			},
			"handle_path": schema.StringAttribute{
				MarkdownDescription: "Enter a path to handle. Choose a pattern like \"/*\" or \"/example/*\". Leave empty to handle any paths (recommended). Any request matching this pattern will be handled by the chosen directive. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"accesslist": schema.StringAttribute{
				MarkdownDescription: "Select an Access List to restrict access to this handler. If unset, any local or remote client is allowed access. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"basicauth": schema.SetAttribute{
				MarkdownDescription: "Select Users to restrict access to this path. If unset, any user is allowed access. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"header": schema.SetAttribute{
				MarkdownDescription: "Select one or multiple headers. Caddy sets \"X-Forwarded-For\", \"X-Forwarded-Proto\" and \"X-Forwarded-Host\" by default, adding them here is not needed. Setting a wrong configuration can be a security risk or break functionality. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"handle_directive": schema.StringAttribute{
				MarkdownDescription: "Choose a handling directive: \"reverse_proxy\" will proxy all HTTP traffic to the selected upstream. \"redir\" will create a HTTP redirect. One of `reverse_proxy`, `redir`. Defaults to `reverse_proxy`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("reverse_proxy"),
				Validators:          []validator.String{stringvalidator.OneOf("reverse_proxy", "redir")},
			},
			"to_domain": schema.SetAttribute{
				MarkdownDescription: "Enter a domain name or IP address of the upstream destination. If multiple are chosen, they will be load balanced with the default random policy. A health check can be activated by populating \"Upstream Fail Duration\" in advanced mode. When directive is \"redir\", only the first domain will be evaluated.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"to_port": schema.Int64Attribute{
				MarkdownDescription: "Leave empty to use the default port or choose a custom port for the upstream destination. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"to_path": schema.StringAttribute{
				MarkdownDescription: "When directive is \"reverse_proxy\", enter a path prefix like \"/guacamole\" that should be prepended to the upstream request. This is useful for destinations that have a virtual directory as their base path. When directive is \"redir\", add the path the request should be redirected to; leaving it empty will append {uri}. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"forward_auth": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable Forward Auth. Requires an \"Auth Provider\" in \"General Settings\". Headers are set automatically to the standard of the chosen provider. Enabling this option will additionally generate the forward_auth directive in front of the reverse_proxy directive inside the scope of this handler. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_tls": schema.StringAttribute{
				MarkdownDescription: "Enable or disable HTTP over TLS (HTTPS) to communicate with the upstream destination. Caddy uses HTTP with the upstream destination by default. One of `0`, `1`, `2`. Defaults to `0`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("0"),
				Validators:          []validator.String{stringvalidator.OneOf("0", "1", "2")},
			},
			"http_version": schema.StringAttribute{
				MarkdownDescription: "The default versions are highly recommended. Choose a HTTP version for the upstream destination. HTTP/3 (HTTP over QUIC) requires HTTPS, and only establishes connections to web servers that also support HTTP/3. One of `http1`, `http2`, `http3`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "http1", "http2", "http3")},
			},
			"http_keepalive": schema.Int64Attribute{
				MarkdownDescription: "Leave empty to use default. Keepalive is either 0 (off) or a duration value that specifies how long to keep connections open (timeout) in seconds. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"http_tls_insecure_skip_verify": schema.BoolAttribute{
				MarkdownDescription: "Caddy uses HTTP by default to connect to the Upstream. If the Upstream is only reachable via HTTPS, this option disables the TLS handshake verification. This makes the connection insecure and vulnerable to man-in-the-middle attacks. In private networks the risk is low, though do not use in production if possible. It is advised to either use plain HTTP, or proper TLS handling by using the options in \"Trust\". Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_tls_trusted_ca_certs": schema.StringAttribute{
				MarkdownDescription: "Choose a CA or self-signed certificate to trust from \"System - Trust - Authorities\". Useful if the upstream destination only accepts TLS connections and offers a self signed certificate. Adding that certificate here will allow for the encrypted connection to succeed. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_tls_server_name": schema.StringAttribute{
				MarkdownDescription: "Enter a hostname or IP address that matches the SAN \"Subject Alternative Name\" of the offered upstream certificate. This will change the SNI \"Server Name Indication\" of Caddy. Setting an IP address as \"Upstream Domain\", enabling \"TLS\" and selecting a \"TLS Trust Pool\", would make the SAN of the offered upstream certificate not match with the SNI of Caddy, since it will be an IP address instead of a hostname. Setting the hostname of the certificate here, fixes this issue. Please note that only SAN certificates are supported; CN \"Common Name\" will not work. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"lb_policy": schema.StringAttribute{
				MarkdownDescription: "lb_policy is the name of the load balancing policy, along with any options. For policies that involve hashing, the highest-random-weight (HRW) algorithm is used to ensure that a client or request with the same hash key is mapped to the same upstream, even if the list of upstreams change. One of `first`, `round_robin`, `least_conn`, `ip_hash`, `client_ip_hash`, `uri_hash`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "first", "round_robin", "least_conn", "ip_hash", "client_ip_hash", "uri_hash")},
			},
			"lb_retries": schema.Int64Attribute{
				MarkdownDescription: "lb_retries is how many times to retry selecting available backends for each request if the next available host is down. If lb_try_duration is also configured, then retries may stop early if the duration is reached. In other words, the retry duration takes precedence over the retry count. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
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
			"passive_health_unhealthy_status": schema.StringAttribute{
				MarkdownDescription: "unhealthy_status counts a request as failed if the response comes back with one of these status codes. Can be a 3-digit status code or a status code class ending in xx, for example: 404 or 5xx. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"passive_health_unhealthy_latency": schema.Int64Attribute{
				MarkdownDescription: "unhealthy_latency is a duration value in ms that counts a request as failed if it takes this long to get a response. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"passive_health_unhealthy_request_count": schema.Int64Attribute{
				MarkdownDescription: "unhealthy_request_count is the permissible number of simultaneous requests to a backend before marking it as down. In other words, if a particular backend is currently handling this many requests, then it is considered \"overloaded\" and other backends will be preferred instead. This should be a reasonably large number; configuring this means that the proxy will have a limit of unhealthy_request_count × upstreams_count total simultaneous requests, and any requests after that point will result in an error due to no upstreams being available. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Enter a description for this handler. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"health_uri": schema.StringAttribute{
				MarkdownDescription: "health_uri is the URI path (and optional query) for active health checks. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"health_upstream": schema.StringAttribute{
				MarkdownDescription: "health_upstream is the ip:port to use for active health checks, if different from the upstream. This should be used in tandem with health_header and {http.reverse_proxy.active.target_upstream}. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"health_port": schema.Int64Attribute{
				MarkdownDescription: "health_port is the port to use for active health checks, if different from the upstream's port. Ignored if health_upstream is used. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"health_interval": schema.Int64Attribute{
				MarkdownDescription: "health_interval is a duration value that defines how often to perform active health checks. Default: 30s. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"health_passes": schema.Int64Attribute{
				MarkdownDescription: "health_passes is the number of consecutive health checks required before marking the backend as healthy again. Default: 1. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"health_fails": schema.Int64Attribute{
				MarkdownDescription: "health_fails is the number of consecutive health checks required before marking the backend as unhealthy. Default: 1. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"health_timeout": schema.Int64Attribute{
				MarkdownDescription: "health_timeout is a duration value that defines how long to wait for a reply before marking the backend as down. Default: 5s. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"health_status": schema.StringAttribute{
				MarkdownDescription: "health_status is the HTTP status code to expect from a healthy backend. Can be a 3-digit status code, or a status code class ending in xx. For example: 200 (which is the default), or 2xx. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"health_body": schema.StringAttribute{
				MarkdownDescription: "health_body is a substring or regular expression to match on the response body of an active health check. If the backend does not return a matching body, it will be marked as down. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"health_follow_redirects": schema.BoolAttribute{
				MarkdownDescription: "health_follow_redirects will cause the health check to follow redirects provided by upstream. By default, a redirect response would cause the health check to count as a fail. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the Caddy handler.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func handleDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Caddy handlers (**Services ▸ Caddy Web Server ▸ Reverse Proxy ▸ Handlers**, plugin `os-caddy`): route the requests of a domain or subdomain, optionally by path, to upstream servers.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the Caddy handler.",
				Required:            true,
			},
			"enabled":                                dschema.BoolAttribute{MarkdownDescription: "Enable this handler.", Computed: true},
			"reverse":                                dschema.StringAttribute{MarkdownDescription: "Select a domain to handle.", Computed: true},
			"subdomain":                              dschema.StringAttribute{MarkdownDescription: "Select a subdomain to handle. Make sure to additionally choose a wildcard domain as \"Domain\". Leave unset, if not using subdomains.", Computed: true},
			"handle_type":                            dschema.StringAttribute{MarkdownDescription: "Choose a handling type: \"handle\" (default) will keep the path in all requests. \"handle_path\" will strip the path from all requests. One of `handle`, `handle_path`.", Computed: true},
			"handle_path":                            dschema.StringAttribute{MarkdownDescription: "Enter a path to handle. Choose a pattern like \"/*\" or \"/example/*\". Leave empty to handle any paths (recommended). Any request matching this pattern will be handled by the chosen directive.", Computed: true},
			"accesslist":                             dschema.StringAttribute{MarkdownDescription: "Select an Access List to restrict access to this handler. If unset, any local or remote client is allowed access.", Computed: true},
			"basicauth":                              dschema.SetAttribute{MarkdownDescription: "Select Users to restrict access to this path. If unset, any user is allowed access.", Computed: true, ElementType: types.StringType},
			"header":                                 dschema.SetAttribute{MarkdownDescription: "Select one or multiple headers. Caddy sets \"X-Forwarded-For\", \"X-Forwarded-Proto\" and \"X-Forwarded-Host\" by default, adding them here is not needed. Setting a wrong configuration can be a security risk or break functionality.", Computed: true, ElementType: types.StringType},
			"handle_directive":                       dschema.StringAttribute{MarkdownDescription: "Choose a handling directive: \"reverse_proxy\" will proxy all HTTP traffic to the selected upstream. \"redir\" will create a HTTP redirect. One of `reverse_proxy`, `redir`.", Computed: true},
			"to_domain":                              dschema.SetAttribute{MarkdownDescription: "Enter a domain name or IP address of the upstream destination. If multiple are chosen, they will be load balanced with the default random policy. A health check can be activated by populating \"Upstream Fail Duration\" in advanced mode. When directive is \"redir\", only the first domain will be evaluated.", Computed: true, ElementType: types.StringType},
			"to_port":                                dschema.Int64Attribute{MarkdownDescription: "Leave empty to use the default port or choose a custom port for the upstream destination.", Computed: true},
			"to_path":                                dschema.StringAttribute{MarkdownDescription: "When directive is \"reverse_proxy\", enter a path prefix like \"/guacamole\" that should be prepended to the upstream request. This is useful for destinations that have a virtual directory as their base path. When directive is \"redir\", add the path the request should be redirected to; leaving it empty will append {uri}.", Computed: true},
			"forward_auth":                           dschema.BoolAttribute{MarkdownDescription: "Enable or disable Forward Auth. Requires an \"Auth Provider\" in \"General Settings\". Headers are set automatically to the standard of the chosen provider. Enabling this option will additionally generate the forward_auth directive in front of the reverse_proxy directive inside the scope of this handler.", Computed: true},
			"http_tls":                               dschema.StringAttribute{MarkdownDescription: "Enable or disable HTTP over TLS (HTTPS) to communicate with the upstream destination. Caddy uses HTTP with the upstream destination by default. One of `0`, `1`, `2`.", Computed: true},
			"http_version":                           dschema.StringAttribute{MarkdownDescription: "The default versions are highly recommended. Choose a HTTP version for the upstream destination. HTTP/3 (HTTP over QUIC) requires HTTPS, and only establishes connections to web servers that also support HTTP/3. One of `http1`, `http2`, `http3`.", Computed: true},
			"http_keepalive":                         dschema.Int64Attribute{MarkdownDescription: "Leave empty to use default. Keepalive is either 0 (off) or a duration value that specifies how long to keep connections open (timeout) in seconds.", Computed: true},
			"http_tls_insecure_skip_verify":          dschema.BoolAttribute{MarkdownDescription: "Caddy uses HTTP by default to connect to the Upstream. If the Upstream is only reachable via HTTPS, this option disables the TLS handshake verification. This makes the connection insecure and vulnerable to man-in-the-middle attacks. In private networks the risk is low, though do not use in production if possible. It is advised to either use plain HTTP, or proper TLS handling by using the options in \"Trust\".", Computed: true},
			"http_tls_trusted_ca_certs":              dschema.StringAttribute{MarkdownDescription: "Choose a CA or self-signed certificate to trust from \"System - Trust - Authorities\". Useful if the upstream destination only accepts TLS connections and offers a self signed certificate. Adding that certificate here will allow for the encrypted connection to succeed.", Computed: true},
			"http_tls_server_name":                   dschema.StringAttribute{MarkdownDescription: "Enter a hostname or IP address that matches the SAN \"Subject Alternative Name\" of the offered upstream certificate. This will change the SNI \"Server Name Indication\" of Caddy. Setting an IP address as \"Upstream Domain\", enabling \"TLS\" and selecting a \"TLS Trust Pool\", would make the SAN of the offered upstream certificate not match with the SNI of Caddy, since it will be an IP address instead of a hostname. Setting the hostname of the certificate here, fixes this issue. Please note that only SAN certificates are supported; CN \"Common Name\" will not work.", Computed: true},
			"lb_policy":                              dschema.StringAttribute{MarkdownDescription: "lb_policy is the name of the load balancing policy, along with any options. For policies that involve hashing, the highest-random-weight (HRW) algorithm is used to ensure that a client or request with the same hash key is mapped to the same upstream, even if the list of upstreams change. One of `first`, `round_robin`, `least_conn`, `ip_hash`, `client_ip_hash`, `uri_hash`.", Computed: true},
			"lb_retries":                             dschema.Int64Attribute{MarkdownDescription: "lb_retries is how many times to retry selecting available backends for each request if the next available host is down. If lb_try_duration is also configured, then retries may stop early if the duration is reached. In other words, the retry duration takes precedence over the retry count.", Computed: true},
			"lb_try_duration":                        dschema.Int64Attribute{MarkdownDescription: "lb_try_duration is a duration value that defines how long to try selecting available backends for each request if the next available host is down. Clients will wait for up to this long while the load balancer tries to find an available upstream host. A reasonable starting point might be 5s since the HTTP transport's default dial timeout is 3s, so that should allow for at least one retry if the first selected upstream cannot be reached.", Computed: true},
			"lb_try_interval":                        dschema.Int64Attribute{MarkdownDescription: "lb_try_interval is a duration value that defines how long to wait between selecting the next host from the pool. Only relevant when a request to an upstream host fails. Be aware that setting this to 0 with a non-zero lb_try_duration can cause the CPU to spin if all backends are down and latency is very low.", Computed: true},
			"passive_health_fail_duration":           dschema.Int64Attribute{MarkdownDescription: "fail_duration enables a passive health check when multiple destinations in \"Upstream Domain\" are set. It is a value that defines how long to remember a failed request. A duration of 1 or more seconds enables passive health checking. A reasonable starting point might be 30s to balance error rates with responsiveness when bringing an unhealthy upstream back online.", Computed: true},
			"passive_health_max_fails":               dschema.Int64Attribute{MarkdownDescription: "max_fails is the maximum number of failed requests within fail_duration that are needed before considering an upstream to be down.", Computed: true},
			"passive_health_unhealthy_status":        dschema.StringAttribute{MarkdownDescription: "unhealthy_status counts a request as failed if the response comes back with one of these status codes. Can be a 3-digit status code or a status code class ending in xx, for example: 404 or 5xx.", Computed: true},
			"passive_health_unhealthy_latency":       dschema.Int64Attribute{MarkdownDescription: "unhealthy_latency is a duration value in ms that counts a request as failed if it takes this long to get a response.", Computed: true},
			"passive_health_unhealthy_request_count": dschema.Int64Attribute{MarkdownDescription: "unhealthy_request_count is the permissible number of simultaneous requests to a backend before marking it as down. In other words, if a particular backend is currently handling this many requests, then it is considered \"overloaded\" and other backends will be preferred instead. This should be a reasonably large number; configuring this means that the proxy will have a limit of unhealthy_request_count × upstreams_count total simultaneous requests, and any requests after that point will result in an error due to no upstreams being available.", Computed: true},
			"description":                            dschema.StringAttribute{MarkdownDescription: "Enter a description for this handler.", Computed: true},
			"health_uri":                             dschema.StringAttribute{MarkdownDescription: "health_uri is the URI path (and optional query) for active health checks.", Computed: true},
			"health_upstream":                        dschema.StringAttribute{MarkdownDescription: "health_upstream is the ip:port to use for active health checks, if different from the upstream. This should be used in tandem with health_header and {http.reverse_proxy.active.target_upstream}.", Computed: true},
			"health_port":                            dschema.Int64Attribute{MarkdownDescription: "health_port is the port to use for active health checks, if different from the upstream's port. Ignored if health_upstream is used.", Computed: true},
			"health_interval":                        dschema.Int64Attribute{MarkdownDescription: "health_interval is a duration value that defines how often to perform active health checks. Default: 30s.", Computed: true},
			"health_passes":                          dschema.Int64Attribute{MarkdownDescription: "health_passes is the number of consecutive health checks required before marking the backend as healthy again. Default: 1.", Computed: true},
			"health_fails":                           dschema.Int64Attribute{MarkdownDescription: "health_fails is the number of consecutive health checks required before marking the backend as unhealthy. Default: 1.", Computed: true},
			"health_timeout":                         dschema.Int64Attribute{MarkdownDescription: "health_timeout is a duration value that defines how long to wait for a reply before marking the backend as down. Default: 5s.", Computed: true},
			"health_status":                          dschema.StringAttribute{MarkdownDescription: "health_status is the HTTP status code to expect from a healthy backend. Can be a 3-digit status code, or a status code class ending in xx. For example: 200 (which is the default), or 2xx.", Computed: true},
			"health_body":                            dschema.StringAttribute{MarkdownDescription: "health_body is a substring or regular expression to match on the response body of an active health check. If the backend does not return a matching body, it will be marked as down.", Computed: true},
			"health_follow_redirects":                dschema.BoolAttribute{MarkdownDescription: "health_follow_redirects will cause the health check to follow redirects provided by upstream. By default, a redirect response would cause the health check to count as a fail.", Computed: true},
		},
	}
}

func convertHandleSchemaToStruct(d *handleResourceModel) (*caddy.Handle, error) {
	return &caddy.Handle{
		Enabled:                            tools.BoolToString(d.Enabled.ValueBool()),
		Reverse:                            api.SelectedMap(d.Reverse.ValueString()),
		Subdomain:                          api.SelectedMap(d.Subdomain.ValueString()),
		HandleType:                         api.SelectedMap(d.HandleType.ValueString()),
		HandlePath:                         d.HandlePath.ValueString(),
		Accesslist:                         api.SelectedMap(d.Accesslist.ValueString()),
		Basicauth:                          api.SelectedMapList(setToSlice(d.Basicauth)),
		Header:                             api.SelectedMapList(setToSlice(d.Header)),
		HandleDirective:                    api.SelectedMap(d.HandleDirective.ValueString()),
		ToDomain:                           api.SelectedMapList(setToSlice(d.ToDomain)),
		ToPort:                             tools.Int64ToStringNegative(d.ToPort.ValueInt64()),
		ToPath:                             d.ToPath.ValueString(),
		ForwardAuth:                        tools.BoolToString(d.ForwardAuth.ValueBool()),
		HttpTls:                            api.SelectedMap(d.HttpTLS.ValueString()),
		HttpVersion:                        api.SelectedMap(d.HttpVersion.ValueString()),
		HttpKeepalive:                      tools.Int64ToStringNegative(d.HttpKeepalive.ValueInt64()),
		HttpTlsInsecureSkipVerify:          tools.BoolToString(d.HttpTLSInsecureSkipVerify.ValueBool()),
		HttpTlsTrustedCaCerts:              api.SelectedMap(d.HttpTLSTrustedCaCerts.ValueString()),
		HttpTlsServerName:                  d.HttpTLSServerName.ValueString(),
		LbPolicy:                           api.SelectedMap(d.LbPolicy.ValueString()),
		LbRetries:                          tools.Int64ToStringNegative(d.LbRetries.ValueInt64()),
		LbTryDuration:                      tools.Int64ToStringNegative(d.LbTryDuration.ValueInt64()),
		LbTryInterval:                      tools.Int64ToStringNegative(d.LbTryInterval.ValueInt64()),
		PassiveHealthFailDuration:          tools.Int64ToStringNegative(d.PassiveHealthFailDuration.ValueInt64()),
		PassiveHealthMaxFails:              tools.Int64ToStringNegative(d.PassiveHealthMaxFails.ValueInt64()),
		PassiveHealthUnhealthyStatus:       d.PassiveHealthUnhealthyStatus.ValueString(),
		PassiveHealthUnhealthyLatency:      tools.Int64ToStringNegative(d.PassiveHealthUnhealthyLatency.ValueInt64()),
		PassiveHealthUnhealthyRequestCount: tools.Int64ToStringNegative(d.PassiveHealthUnhealthyRequestCount.ValueInt64()),
		Description:                        d.Description.ValueString(),
		HealthUri:                          d.HealthUri.ValueString(),
		HealthUpstream:                     d.HealthUpstream.ValueString(),
		HealthPort:                         tools.Int64ToStringNegative(d.HealthPort.ValueInt64()),
		HealthInterval:                     tools.Int64ToStringNegative(d.HealthInterval.ValueInt64()),
		HealthPasses:                       tools.Int64ToStringNegative(d.HealthPasses.ValueInt64()),
		HealthFails:                        tools.Int64ToStringNegative(d.HealthFails.ValueInt64()),
		HealthTimeout:                      tools.Int64ToStringNegative(d.HealthTimeout.ValueInt64()),
		HealthStatus:                       d.HealthStatus.ValueString(),
		HealthBody:                         d.HealthBody.ValueString(),
		HealthFollowRedirects:              tools.BoolToString(d.HealthFollowRedirects.ValueBool()),
	}, nil
}

func convertHandleStructToSchema(d *caddy.Handle) (*handleResourceModel, error) {
	return &handleResourceModel{
		Enabled:                            types.BoolValue(tools.StringToBool(d.Enabled)),
		Reverse:                            types.StringValue(d.Reverse.String()),
		Subdomain:                          types.StringValue(d.Subdomain.String()),
		HandleType:                         types.StringValue(d.HandleType.String()),
		HandlePath:                         types.StringValue(d.HandlePath),
		Accesslist:                         types.StringValue(d.Accesslist.String()),
		Basicauth:                          tools.StringSliceToSet(d.Basicauth),
		Header:                             tools.StringSliceToSet(d.Header),
		HandleDirective:                    types.StringValue(d.HandleDirective.String()),
		ToDomain:                           tools.StringSliceToSet(d.ToDomain),
		ToPort:                             types.Int64Value(tools.StringToInt64(d.ToPort)),
		ToPath:                             types.StringValue(d.ToPath),
		ForwardAuth:                        types.BoolValue(tools.StringToBool(d.ForwardAuth)),
		HttpTLS:                            types.StringValue(d.HttpTls.String()),
		HttpVersion:                        types.StringValue(d.HttpVersion.String()),
		HttpKeepalive:                      types.Int64Value(tools.StringToInt64(d.HttpKeepalive)),
		HttpTLSInsecureSkipVerify:          types.BoolValue(tools.StringToBool(d.HttpTlsInsecureSkipVerify)),
		HttpTLSTrustedCaCerts:              types.StringValue(d.HttpTlsTrustedCaCerts.String()),
		HttpTLSServerName:                  types.StringValue(d.HttpTlsServerName),
		LbPolicy:                           types.StringValue(d.LbPolicy.String()),
		LbRetries:                          types.Int64Value(tools.StringToInt64(d.LbRetries)),
		LbTryDuration:                      types.Int64Value(tools.StringToInt64(d.LbTryDuration)),
		LbTryInterval:                      types.Int64Value(tools.StringToInt64(d.LbTryInterval)),
		PassiveHealthFailDuration:          types.Int64Value(tools.StringToInt64(d.PassiveHealthFailDuration)),
		PassiveHealthMaxFails:              types.Int64Value(tools.StringToInt64(d.PassiveHealthMaxFails)),
		PassiveHealthUnhealthyStatus:       types.StringValue(d.PassiveHealthUnhealthyStatus),
		PassiveHealthUnhealthyLatency:      types.Int64Value(tools.StringToInt64(d.PassiveHealthUnhealthyLatency)),
		PassiveHealthUnhealthyRequestCount: types.Int64Value(tools.StringToInt64(d.PassiveHealthUnhealthyRequestCount)),
		Description:                        types.StringValue(d.Description),
		HealthUri:                          types.StringValue(d.HealthUri),
		HealthUpstream:                     types.StringValue(d.HealthUpstream),
		HealthPort:                         types.Int64Value(tools.StringToInt64(d.HealthPort)),
		HealthInterval:                     types.Int64Value(tools.StringToInt64(d.HealthInterval)),
		HealthPasses:                       types.Int64Value(tools.StringToInt64(d.HealthPasses)),
		HealthFails:                        types.Int64Value(tools.StringToInt64(d.HealthFails)),
		HealthTimeout:                      types.Int64Value(tools.StringToInt64(d.HealthTimeout)),
		HealthStatus:                       types.StringValue(d.HealthStatus),
		HealthBody:                         types.StringValue(d.HealthBody),
		HealthFollowRedirects:              types.BoolValue(tools.StringToBool(d.HealthFollowRedirects)),
	}, nil
}
