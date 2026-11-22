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

type locationResourceModel struct {
	Description             types.String `tfsdk:"description"`
	Urlpattern              types.String `tfsdk:"urlpattern"`
	Matchtype               types.String `tfsdk:"matchtype"`
	EnableSecrules          types.Bool   `tfsdk:"enable_secrules"`
	EnableLearningMode      types.Bool   `tfsdk:"enable_learning_mode"`
	SecrulesErrorpage       types.String `tfsdk:"secrules_errorpage"`
	XssBlockScore           types.Int64  `tfsdk:"xss_block_score"`
	SqliBlockScore          types.Int64  `tfsdk:"sqli_block_score"`
	CustomPolicy            types.Set    `tfsdk:"custom_policy"`
	Upstream                types.String `tfsdk:"upstream"`
	PathPrefix              types.String `tfsdk:"path_prefix"`
	CachePath               types.String `tfsdk:"cache_path"`
	CacheUseStale           types.Set    `tfsdk:"cache_use_stale"`
	CacheMethods            types.Set    `tfsdk:"cache_methods"`
	CacheMinUses            types.Int64  `tfsdk:"cache_min_uses"`
	CacheValid              types.Int64  `tfsdk:"cache_valid"`
	ProxyCacheValid         types.Set    `tfsdk:"proxy_cache_valid"`
	CacheBackgroundUpdate   types.Bool   `tfsdk:"cache_background_update"`
	CacheLock               types.Bool   `tfsdk:"cache_lock"`
	CacheRevalidate         types.Bool   `tfsdk:"cache_revalidate"`
	Root                    types.String `tfsdk:"root"`
	Rewrites                types.Set    `tfsdk:"rewrites"`
	Index                   types.Set    `tfsdk:"index"`
	Autoindex               types.Bool   `tfsdk:"autoindex"`
	Authbasic               types.String `tfsdk:"authbasic"`
	Authbasicuserfile       types.String `tfsdk:"authbasicuserfile"`
	AdvancedAcl             types.Bool   `tfsdk:"advanced_acl"`
	ForceHttps              types.String `tfsdk:"force_https"`
	PhpEnable               types.Bool   `tfsdk:"php_enable"`
	PhpOverrideScriptname   types.String `tfsdk:"php_override_scriptname"`
	LimitRequestConnections types.Set    `tfsdk:"limit_request_connections"`
	MaxBodySize             types.String `tfsdk:"max_body_size"`
	BodyBufferSize          types.String `tfsdk:"body_buffer_size"`
	Honeypot                types.Bool   `tfsdk:"honeypot"`
	Websocket               types.Bool   `tfsdk:"websocket"`
	UpstreamKeepalive       types.Bool   `tfsdk:"upstream_keepalive"`
	ProxyBufferSize         types.Int64  `tfsdk:"proxy_buffer_size"`
	ProxyBuffersCount       types.Int64  `tfsdk:"proxy_buffers_count"`
	ProxyBuffersSize        types.Int64  `tfsdk:"proxy_buffers_size"`
	ProxyBusyBuffersSize    types.Int64  `tfsdk:"proxy_busy_buffers_size"`
	ProxyIgnoreClientAbort  types.Bool   `tfsdk:"proxy_ignore_client_abort"`
	ProxyRequestBuffering   types.Bool   `tfsdk:"proxy_request_buffering"`
	ProxyBuffering          types.Bool   `tfsdk:"proxy_buffering"`
	ProxyReadTimeout        types.Int64  `tfsdk:"proxy_read_timeout"`
	ProxySendTimeout        types.Int64  `tfsdk:"proxy_send_timeout"`
	IPAcl                   types.String `tfsdk:"ip_acl"`
	Satisfy                 types.String `tfsdk:"satisfy"`
	ProxyMaxTempFileSize    types.Int64  `tfsdk:"proxy_max_temp_file_size"`
	ProxySslServerName      types.Bool   `tfsdk:"proxy_ssl_server_name"`
	Errorpages              types.Set    `tfsdk:"errorpages"`
	ProxyInterceptErrors    types.String `tfsdk:"proxy_intercept_errors"`

	Id types.String `tfsdk:"id"`
}

func locationResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx locations (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ Location**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				MarkdownDescription: "Description.",
				Required:            true,
			},
			"urlpattern": schema.StringAttribute{
				MarkdownDescription: "The URL pattern to match.",
				Required:            true,
			},
			"matchtype": schema.StringAttribute{
				MarkdownDescription: "Choose how to match on \"URL Pattern\". One of `=`, `~`, `~*`, `^~`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "=", "~", "~*", "^~")},
			},
			"enable_secrules": schema.BoolAttribute{
				MarkdownDescription: "Enable WAF. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"enable_learning_mode": schema.BoolAttribute{
				MarkdownDescription: "Enable learning mode means nothing is blocked but logged. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"secrules_errorpage": schema.StringAttribute{
				MarkdownDescription: "Select a custom error page to display instead of the default builtin error pages for security rule violations. Only the page content itself is used. Status code rewriting and redirection is not supported. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"xss_block_score": schema.Int64Attribute{
				MarkdownDescription: "Block XSS with a score above the specified value using generic detection. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sqli_block_score": schema.Int64Attribute{
				MarkdownDescription: "Block SQL injection with a score above the specified value using generic detection. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"custom_policy": schema.SetAttribute{
				MarkdownDescription: "Select custom security policies. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"upstream": schema.StringAttribute{
				MarkdownDescription: "Select an upstream to proxy to or connect via FastCGI if chosen. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"path_prefix": schema.StringAttribute{
				MarkdownDescription: "Define an optional path prefix for this location. It will be prepended to the path so http://www.example.com/index.php will be sent to http://www.example.com/app/index.php on your upstream if you choose \"/app\" as path prefix. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"cache_path": schema.StringAttribute{
				MarkdownDescription: "Choose a cache directory if you want to cache responses. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"cache_use_stale": schema.SetAttribute{
				MarkdownDescription: "If you enable this option, a stale response will be sent when in cases the original server is unable to (for example if it is down). One of `error`, `timeout`, `invalid_header`, `updating`, `http_403`, `http_404`, `http_429`, `http_500`, `http_502`, `http_503`, `http_504`. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
				Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("error", "timeout", "invalid_header", "updating", "http_403", "http_404", "http_429", "http_500", "http_502", "http_503", "http_504"))},
			},
			"cache_methods": schema.SetAttribute{
				MarkdownDescription: "Select the HTTP verbs to cache. GET and HEAD will be always cached. One of `POST`. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
				Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("POST"))},
			},
			"cache_min_uses": schema.Int64Attribute{
				MarkdownDescription: "Enter how often the resource must be hit before adding it to the cache. Defaults to `1`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1),
			},
			"cache_valid": schema.Int64Attribute{
				MarkdownDescription: "Force caching of 200, 301 and 302 responses according to the request methods enabled for caching. Given in minutes; leave empty to rely on request/response headers from client and upstream. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"proxy_cache_valid": schema.SetAttribute{
				MarkdownDescription: "Force caching of response codes specified at Response Code Caching page. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"cache_background_update": schema.BoolAttribute{
				MarkdownDescription: "Return a stale response to the client and update the cache. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"cache_lock": schema.BoolAttribute{
				MarkdownDescription: "Only allow a single request to call the backend on the same URL at a single time. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"cache_revalidate": schema.BoolAttribute{
				MarkdownDescription: "Only request a new version, if the content has changed since the last cache update. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"root": schema.StringAttribute{
				MarkdownDescription: "Enter the file system root from which the files are served. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"rewrites": schema.SetAttribute{
				MarkdownDescription: "Choose URL rewriting rules. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"index": schema.SetAttribute{
				MarkdownDescription: "Enter a list of file extensions, which are served instead of a directory. It is common to use index.html or index.php here. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"autoindex": schema.BoolAttribute{
				MarkdownDescription: "Check this to serve a file list when a directory is requested. This is useful for serving multiple files with someone without a web application. It is recommended to use this option only with authentication. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"authbasic": schema.StringAttribute{
				MarkdownDescription: "Enter the realm for this directory and enable authentication. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"authbasicuserfile": schema.StringAttribute{
				MarkdownDescription: "Select a credential list to use. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"advanced_acl": schema.BoolAttribute{
				MarkdownDescription: "Send an authentication request to the OPNsense backend for advanced access control. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"force_https": schema.StringAttribute{
				MarkdownDescription: "If the request scheme is not HTTPS, redirect to use HTTPS for this location. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"php_enable": schema.BoolAttribute{
				MarkdownDescription: "Only use this setting if you know what you are doing. If not, expect your machine to get compromised. Use this setting when you want to run a local PHP application or call an external upstream, if selected, via FastCGI. Example upstreams are PHP-FPM or Rails via FastCGI API. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"php_override_scriptname": schema.StringAttribute{
				MarkdownDescription: "If you set this setting, all requests are sent to this script instead of the request path (URL). Not using this setting on a remote instance can be dangerous. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"limit_request_connections": schema.SetAttribute{
				MarkdownDescription: "If you choose multiple limits, the strictest will be used. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"max_body_size": schema.StringAttribute{
				MarkdownDescription: "If the request is larger, it will be rejectet with error 413 (Request Entity Too Large). For example, you can enter 200m. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"body_buffer_size": schema.StringAttribute{
				MarkdownDescription: "If the request exceeds this size, it will be written to disk. Enter a number and a unit like 1m. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"honeypot": schema.BoolAttribute{
				MarkdownDescription: "If you enable the honeypot, all requests to this location will go to a special temporary log which will be used to block the IP. This is dangerous because you may accidentally block legitimate users or search engines. The result is available as a special alias in the firewall section. For example you can trigger on locations of Wordpress for phpMyAdmin if you are not using it. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"websocket": schema.BoolAttribute{
				MarkdownDescription: "If you enable the WebSocket Support option, nginx will pass the upgrade header to the backed server. Mutually exclusive with Keepalive support. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"upstream_keepalive": schema.BoolAttribute{
				MarkdownDescription: "Adds the directives required for upstream keepalive to work (enables HTTP 1.1 and clears the Connection header). Mutually exclusive with WebSocket support. Keepalive parameters must be set in Upstream settings too. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"proxy_buffer_size": schema.Int64Attribute{
				MarkdownDescription: "By default this is the size of one page (4 or 8 Kilobytes). It can be changed. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"proxy_buffers_count": schema.Int64Attribute{
				MarkdownDescription: "Count of buffers for a single connection. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"proxy_buffers_size": schema.Int64Attribute{
				MarkdownDescription: "Size of buffers of a single connection. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"proxy_busy_buffers_size": schema.Int64Attribute{
				MarkdownDescription: "When this buffer size is reached, nginx will send a part of the response to the client. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"proxy_ignore_client_abort": schema.BoolAttribute{
				MarkdownDescription: "If you enable this option, nginx will not terminate the connection to the backend server if the client connection is terminated. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"proxy_request_buffering": schema.BoolAttribute{
				MarkdownDescription: "If you enable the WebSocket option, nginx will pass the upgrade header to the backed server. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"proxy_buffering": schema.BoolAttribute{
				MarkdownDescription: "If you enable the Response Buffering option, nginx will buffer response from the backed server. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"proxy_read_timeout": schema.Int64Attribute{
				MarkdownDescription: "Enter a custom timout between data received from the client after which the connection is closed. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"proxy_send_timeout": schema.Int64Attribute{
				MarkdownDescription: "Enter a custom timout between data it sent to the client after which the connection is closed. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"ip_acl": schema.StringAttribute{
				MarkdownDescription: "If you select an IP ACL, the client can only access this service if it fulfills this requirement. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"satisfy": schema.StringAttribute{
				MarkdownDescription: "All: All access restrictions must be fulfilled; Any: Any of the access restrictions must be fulfilled. One of `any`, `all`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "any", "all")},
			},
			"proxy_max_temp_file_size": schema.Int64Attribute{
				MarkdownDescription: "The maximum temporary file size limit. This does not apply to cached requests and responses. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"proxy_ssl_server_name": schema.BoolAttribute{
				MarkdownDescription: "Check this box, if you want to passthrough the hostname you get from the downstream to the upstream connection. Can be overriden in the upstream's hostname configuration. Please note, that TLS server certificates are validated against this servername. If the server name passed by the client leads to different processing on the upstream server (i.e. different virtual host is choosen), this may cause security leaks. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"errorpages": schema.SetAttribute{
				MarkdownDescription: "Select custom error pages to display instead of the default builtin error pages. Selection will override error pages configured on HTTP server. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"proxy_intercept_errors": schema.StringAttribute{
				MarkdownDescription: "Intercept responses with codes greater than or equal to 300 and redirect to processing with custom error pages. One of `Inherit`, `on`, `off`. Defaults to `Inherit`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("Inherit"),
				Validators:          []validator.String{stringvalidator.OneOf("Inherit", "on", "off")},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx location.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func locationDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx locations (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ Location**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx location.",
				Required:            true,
			},
			"description":               dschema.StringAttribute{MarkdownDescription: "Description.", Computed: true},
			"urlpattern":                dschema.StringAttribute{MarkdownDescription: "The URL pattern to match.", Computed: true},
			"matchtype":                 dschema.StringAttribute{MarkdownDescription: "Choose how to match on \"URL Pattern\". One of `=`, `~`, `~*`, `^~`.", Computed: true},
			"enable_secrules":           dschema.BoolAttribute{MarkdownDescription: "Enable WAF.", Computed: true},
			"enable_learning_mode":      dschema.BoolAttribute{MarkdownDescription: "Enable learning mode means nothing is blocked but logged.", Computed: true},
			"secrules_errorpage":        dschema.StringAttribute{MarkdownDescription: "Select a custom error page to display instead of the default builtin error pages for security rule violations. Only the page content itself is used. Status code rewriting and redirection is not supported.", Computed: true},
			"xss_block_score":           dschema.Int64Attribute{MarkdownDescription: "Block XSS with a score above the specified value using generic detection.", Computed: true},
			"sqli_block_score":          dschema.Int64Attribute{MarkdownDescription: "Block SQL injection with a score above the specified value using generic detection.", Computed: true},
			"custom_policy":             dschema.SetAttribute{MarkdownDescription: "Select custom security policies.", Computed: true, ElementType: types.StringType},
			"upstream":                  dschema.StringAttribute{MarkdownDescription: "Select an upstream to proxy to or connect via FastCGI if chosen.", Computed: true},
			"path_prefix":               dschema.StringAttribute{MarkdownDescription: "Define an optional path prefix for this location. It will be prepended to the path so http://www.example.com/index.php will be sent to http://www.example.com/app/index.php on your upstream if you choose \"/app\" as path prefix.", Computed: true},
			"cache_path":                dschema.StringAttribute{MarkdownDescription: "Choose a cache directory if you want to cache responses.", Computed: true},
			"cache_use_stale":           dschema.SetAttribute{MarkdownDescription: "If you enable this option, a stale response will be sent when in cases the original server is unable to (for example if it is down). One of `error`, `timeout`, `invalid_header`, `updating`, `http_403`, `http_404`, `http_429`, `http_500`, `http_502`, `http_503`, `http_504`.", Computed: true, ElementType: types.StringType},
			"cache_methods":             dschema.SetAttribute{MarkdownDescription: "Select the HTTP verbs to cache. GET and HEAD will be always cached. One of `POST`.", Computed: true, ElementType: types.StringType},
			"cache_min_uses":            dschema.Int64Attribute{MarkdownDescription: "Enter how often the resource must be hit before adding it to the cache.", Computed: true},
			"cache_valid":               dschema.Int64Attribute{MarkdownDescription: "Force caching of 200, 301 and 302 responses according to the request methods enabled for caching. Given in minutes; leave empty to rely on request/response headers from client and upstream.", Computed: true},
			"proxy_cache_valid":         dschema.SetAttribute{MarkdownDescription: "Force caching of response codes specified at Response Code Caching page.", Computed: true, ElementType: types.StringType},
			"cache_background_update":   dschema.BoolAttribute{MarkdownDescription: "Return a stale response to the client and update the cache.", Computed: true},
			"cache_lock":                dschema.BoolAttribute{MarkdownDescription: "Only allow a single request to call the backend on the same URL at a single time.", Computed: true},
			"cache_revalidate":          dschema.BoolAttribute{MarkdownDescription: "Only request a new version, if the content has changed since the last cache update.", Computed: true},
			"root":                      dschema.StringAttribute{MarkdownDescription: "Enter the file system root from which the files are served.", Computed: true},
			"rewrites":                  dschema.SetAttribute{MarkdownDescription: "Choose URL rewriting rules.", Computed: true, ElementType: types.StringType},
			"index":                     dschema.SetAttribute{MarkdownDescription: "Enter a list of file extensions, which are served instead of a directory. It is common to use index.html or index.php here.", Computed: true, ElementType: types.StringType},
			"autoindex":                 dschema.BoolAttribute{MarkdownDescription: "Check this to serve a file list when a directory is requested. This is useful for serving multiple files with someone without a web application. It is recommended to use this option only with authentication.", Computed: true},
			"authbasic":                 dschema.StringAttribute{MarkdownDescription: "Enter the realm for this directory and enable authentication.", Computed: true},
			"authbasicuserfile":         dschema.StringAttribute{MarkdownDescription: "Select a credential list to use.", Computed: true},
			"advanced_acl":              dschema.BoolAttribute{MarkdownDescription: "Send an authentication request to the OPNsense backend for advanced access control.", Computed: true},
			"force_https":               dschema.StringAttribute{MarkdownDescription: "If the request scheme is not HTTPS, redirect to use HTTPS for this location.", Computed: true},
			"php_enable":                dschema.BoolAttribute{MarkdownDescription: "Only use this setting if you know what you are doing. If not, expect your machine to get compromised. Use this setting when you want to run a local PHP application or call an external upstream, if selected, via FastCGI. Example upstreams are PHP-FPM or Rails via FastCGI API.", Computed: true},
			"php_override_scriptname":   dschema.StringAttribute{MarkdownDescription: "If you set this setting, all requests are sent to this script instead of the request path (URL). Not using this setting on a remote instance can be dangerous.", Computed: true},
			"limit_request_connections": dschema.SetAttribute{MarkdownDescription: "If you choose multiple limits, the strictest will be used.", Computed: true, ElementType: types.StringType},
			"max_body_size":             dschema.StringAttribute{MarkdownDescription: "If the request is larger, it will be rejectet with error 413 (Request Entity Too Large). For example, you can enter 200m.", Computed: true},
			"body_buffer_size":          dschema.StringAttribute{MarkdownDescription: "If the request exceeds this size, it will be written to disk. Enter a number and a unit like 1m.", Computed: true},
			"honeypot":                  dschema.BoolAttribute{MarkdownDescription: "If you enable the honeypot, all requests to this location will go to a special temporary log which will be used to block the IP. This is dangerous because you may accidentally block legitimate users or search engines. The result is available as a special alias in the firewall section. For example you can trigger on locations of Wordpress for phpMyAdmin if you are not using it.", Computed: true},
			"websocket":                 dschema.BoolAttribute{MarkdownDescription: "If you enable the WebSocket Support option, nginx will pass the upgrade header to the backed server. Mutually exclusive with Keepalive support.", Computed: true},
			"upstream_keepalive":        dschema.BoolAttribute{MarkdownDescription: "Adds the directives required for upstream keepalive to work (enables HTTP 1.1 and clears the Connection header). Mutually exclusive with WebSocket support. Keepalive parameters must be set in Upstream settings too.", Computed: true},
			"proxy_buffer_size":         dschema.Int64Attribute{MarkdownDescription: "By default this is the size of one page (4 or 8 Kilobytes). It can be changed.", Computed: true},
			"proxy_buffers_count":       dschema.Int64Attribute{MarkdownDescription: "Count of buffers for a single connection.", Computed: true},
			"proxy_buffers_size":        dschema.Int64Attribute{MarkdownDescription: "Size of buffers of a single connection.", Computed: true},
			"proxy_busy_buffers_size":   dschema.Int64Attribute{MarkdownDescription: "When this buffer size is reached, nginx will send a part of the response to the client.", Computed: true},
			"proxy_ignore_client_abort": dschema.BoolAttribute{MarkdownDescription: "If you enable this option, nginx will not terminate the connection to the backend server if the client connection is terminated.", Computed: true},
			"proxy_request_buffering":   dschema.BoolAttribute{MarkdownDescription: "If you enable the WebSocket option, nginx will pass the upgrade header to the backed server.", Computed: true},
			"proxy_buffering":           dschema.BoolAttribute{MarkdownDescription: "If you enable the Response Buffering option, nginx will buffer response from the backed server.", Computed: true},
			"proxy_read_timeout":        dschema.Int64Attribute{MarkdownDescription: "Enter a custom timout between data received from the client after which the connection is closed.", Computed: true},
			"proxy_send_timeout":        dschema.Int64Attribute{MarkdownDescription: "Enter a custom timout between data it sent to the client after which the connection is closed.", Computed: true},
			"ip_acl":                    dschema.StringAttribute{MarkdownDescription: "If you select an IP ACL, the client can only access this service if it fulfills this requirement.", Computed: true},
			"satisfy":                   dschema.StringAttribute{MarkdownDescription: "All: All access restrictions must be fulfilled; Any: Any of the access restrictions must be fulfilled. One of `any`, `all`.", Computed: true},
			"proxy_max_temp_file_size":  dschema.Int64Attribute{MarkdownDescription: "The maximum temporary file size limit. This does not apply to cached requests and responses.", Computed: true},
			"proxy_ssl_server_name":     dschema.BoolAttribute{MarkdownDescription: "Check this box, if you want to passthrough the hostname you get from the downstream to the upstream connection. Can be overriden in the upstream's hostname configuration. Please note, that TLS server certificates are validated against this servername. If the server name passed by the client leads to different processing on the upstream server (i.e. different virtual host is choosen), this may cause security leaks.", Computed: true},
			"errorpages":                dschema.SetAttribute{MarkdownDescription: "Select custom error pages to display instead of the default builtin error pages. Selection will override error pages configured on HTTP server.", Computed: true, ElementType: types.StringType},
			"proxy_intercept_errors":    dschema.StringAttribute{MarkdownDescription: "Intercept responses with codes greater than or equal to 300 and redirect to processing with custom error pages. One of `Inherit`, `on`, `off`.", Computed: true},
		},
	}
}

func convertLocationSchemaToStruct(d *locationResourceModel) (*nginx.Location, error) {
	return &nginx.Location{
		Description:             d.Description.ValueString(),
		Urlpattern:              d.Urlpattern.ValueString(),
		Matchtype:               api.SelectedMap(d.Matchtype.ValueString()),
		EnableSecrules:          tools.BoolToString(d.EnableSecrules.ValueBool()),
		EnableLearningMode:      tools.BoolToString(d.EnableLearningMode.ValueBool()),
		SecrulesErrorpage:       api.SelectedMap(d.SecrulesErrorpage.ValueString()),
		XssBlockScore:           tools.Int64ToStringNegative(d.XssBlockScore.ValueInt64()),
		SqliBlockScore:          tools.Int64ToStringNegative(d.SqliBlockScore.ValueInt64()),
		CustomPolicy:            api.SelectedMapList(setToSlice(d.CustomPolicy)),
		Upstream:                api.SelectedMap(d.Upstream.ValueString()),
		PathPrefix:              d.PathPrefix.ValueString(),
		CachePath:               api.SelectedMap(d.CachePath.ValueString()),
		CacheUseStale:           api.SelectedMapList(setToSlice(d.CacheUseStale)),
		CacheMethods:            api.SelectedMapList(setToSlice(d.CacheMethods)),
		CacheMinUses:            tools.Int64ToString(d.CacheMinUses.ValueInt64()),
		CacheValid:              tools.Int64ToStringNegative(d.CacheValid.ValueInt64()),
		ProxyCacheValid:         api.SelectedMapList(setToSlice(d.ProxyCacheValid)),
		CacheBackgroundUpdate:   tools.BoolToString(d.CacheBackgroundUpdate.ValueBool()),
		CacheLock:               tools.BoolToString(d.CacheLock.ValueBool()),
		CacheRevalidate:         tools.BoolToString(d.CacheRevalidate.ValueBool()),
		Root:                    d.Root.ValueString(),
		Rewrites:                api.SelectedMapList(setToSlice(d.Rewrites)),
		Index:                   api.SelectedMapList(setToSlice(d.Index)),
		Autoindex:               tools.BoolToString(d.Autoindex.ValueBool()),
		Authbasic:               d.Authbasic.ValueString(),
		Authbasicuserfile:       api.SelectedMap(d.Authbasicuserfile.ValueString()),
		AdvancedAcl:             tools.BoolToString(d.AdvancedAcl.ValueBool()),
		ForceHttps:              d.ForceHttps.ValueString(),
		PhpEnable:               tools.BoolToString(d.PhpEnable.ValueBool()),
		PhpOverrideScriptname:   d.PhpOverrideScriptname.ValueString(),
		LimitRequestConnections: api.SelectedMapList(setToSlice(d.LimitRequestConnections)),
		MaxBodySize:             d.MaxBodySize.ValueString(),
		BodyBufferSize:          d.BodyBufferSize.ValueString(),
		Honeypot:                tools.BoolToString(d.Honeypot.ValueBool()),
		Websocket:               tools.BoolToString(d.Websocket.ValueBool()),
		UpstreamKeepalive:       tools.BoolToString(d.UpstreamKeepalive.ValueBool()),
		ProxyBufferSize:         tools.Int64ToStringNegative(d.ProxyBufferSize.ValueInt64()),
		ProxyBuffersCount:       tools.Int64ToStringNegative(d.ProxyBuffersCount.ValueInt64()),
		ProxyBuffersSize:        tools.Int64ToStringNegative(d.ProxyBuffersSize.ValueInt64()),
		ProxyBusyBuffersSize:    tools.Int64ToStringNegative(d.ProxyBusyBuffersSize.ValueInt64()),
		ProxyIgnoreClientAbort:  tools.BoolToString(d.ProxyIgnoreClientAbort.ValueBool()),
		ProxyRequestBuffering:   tools.BoolToString(d.ProxyRequestBuffering.ValueBool()),
		ProxyBuffering:          tools.BoolToString(d.ProxyBuffering.ValueBool()),
		ProxyReadTimeout:        tools.Int64ToStringNegative(d.ProxyReadTimeout.ValueInt64()),
		ProxySendTimeout:        tools.Int64ToStringNegative(d.ProxySendTimeout.ValueInt64()),
		IpAcl:                   api.SelectedMap(d.IPAcl.ValueString()),
		Satisfy:                 api.SelectedMap(d.Satisfy.ValueString()),
		ProxyMaxTempFileSize:    tools.Int64ToStringNegative(d.ProxyMaxTempFileSize.ValueInt64()),
		ProxySslServerName:      tools.BoolToString(d.ProxySslServerName.ValueBool()),
		Errorpages:              api.SelectedMapList(setToSlice(d.Errorpages)),
		ProxyInterceptErrors:    api.SelectedMap(d.ProxyInterceptErrors.ValueString()),
	}, nil
}

func convertLocationStructToSchema(d *nginx.Location) (*locationResourceModel, error) {
	return &locationResourceModel{
		Description:             types.StringValue(d.Description),
		Urlpattern:              types.StringValue(d.Urlpattern),
		Matchtype:               types.StringValue(d.Matchtype.String()),
		EnableSecrules:          types.BoolValue(tools.StringToBool(d.EnableSecrules)),
		EnableLearningMode:      types.BoolValue(tools.StringToBool(d.EnableLearningMode)),
		SecrulesErrorpage:       types.StringValue(d.SecrulesErrorpage.String()),
		XssBlockScore:           types.Int64Value(tools.StringToInt64(d.XssBlockScore)),
		SqliBlockScore:          types.Int64Value(tools.StringToInt64(d.SqliBlockScore)),
		CustomPolicy:            tools.StringSliceToSet(d.CustomPolicy),
		Upstream:                types.StringValue(d.Upstream.String()),
		PathPrefix:              types.StringValue(d.PathPrefix),
		CachePath:               types.StringValue(d.CachePath.String()),
		CacheUseStale:           tools.StringSliceToSet(d.CacheUseStale),
		CacheMethods:            tools.StringSliceToSet(d.CacheMethods),
		CacheMinUses:            types.Int64Value(tools.StringToInt64(d.CacheMinUses)),
		CacheValid:              types.Int64Value(tools.StringToInt64(d.CacheValid)),
		ProxyCacheValid:         tools.StringSliceToSet(d.ProxyCacheValid),
		CacheBackgroundUpdate:   types.BoolValue(tools.StringToBool(d.CacheBackgroundUpdate)),
		CacheLock:               types.BoolValue(tools.StringToBool(d.CacheLock)),
		CacheRevalidate:         types.BoolValue(tools.StringToBool(d.CacheRevalidate)),
		Root:                    types.StringValue(d.Root),
		Rewrites:                tools.StringSliceToSet(d.Rewrites),
		Index:                   tools.StringSliceToSet(d.Index),
		Autoindex:               types.BoolValue(tools.StringToBool(d.Autoindex)),
		Authbasic:               types.StringValue(d.Authbasic),
		Authbasicuserfile:       types.StringValue(d.Authbasicuserfile.String()),
		AdvancedAcl:             types.BoolValue(tools.StringToBool(d.AdvancedAcl)),
		ForceHttps:              types.StringValue(d.ForceHttps),
		PhpEnable:               types.BoolValue(tools.StringToBool(d.PhpEnable)),
		PhpOverrideScriptname:   types.StringValue(d.PhpOverrideScriptname),
		LimitRequestConnections: tools.StringSliceToSet(d.LimitRequestConnections),
		MaxBodySize:             types.StringValue(d.MaxBodySize),
		BodyBufferSize:          types.StringValue(d.BodyBufferSize),
		Honeypot:                types.BoolValue(tools.StringToBool(d.Honeypot)),
		Websocket:               types.BoolValue(tools.StringToBool(d.Websocket)),
		UpstreamKeepalive:       types.BoolValue(tools.StringToBool(d.UpstreamKeepalive)),
		ProxyBufferSize:         types.Int64Value(tools.StringToInt64(d.ProxyBufferSize)),
		ProxyBuffersCount:       types.Int64Value(tools.StringToInt64(d.ProxyBuffersCount)),
		ProxyBuffersSize:        types.Int64Value(tools.StringToInt64(d.ProxyBuffersSize)),
		ProxyBusyBuffersSize:    types.Int64Value(tools.StringToInt64(d.ProxyBusyBuffersSize)),
		ProxyIgnoreClientAbort:  types.BoolValue(tools.StringToBool(d.ProxyIgnoreClientAbort)),
		ProxyRequestBuffering:   types.BoolValue(tools.StringToBool(d.ProxyRequestBuffering)),
		ProxyBuffering:          types.BoolValue(tools.StringToBool(d.ProxyBuffering)),
		ProxyReadTimeout:        types.Int64Value(tools.StringToInt64(d.ProxyReadTimeout)),
		ProxySendTimeout:        types.Int64Value(tools.StringToInt64(d.ProxySendTimeout)),
		IPAcl:                   types.StringValue(d.IpAcl.String()),
		Satisfy:                 types.StringValue(d.Satisfy.String()),
		ProxyMaxTempFileSize:    types.Int64Value(tools.StringToInt64(d.ProxyMaxTempFileSize)),
		ProxySslServerName:      types.BoolValue(tools.StringToBool(d.ProxySslServerName)),
		Errorpages:              tools.StringSliceToSet(d.Errorpages),
		ProxyInterceptErrors:    types.StringValue(d.ProxyInterceptErrors.String()),
	}, nil
}
