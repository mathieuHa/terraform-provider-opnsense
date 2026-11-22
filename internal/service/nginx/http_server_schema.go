package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
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

type httpServerResourceModel struct {
	Servername                     types.Set    `tfsdk:"servername"`
	SyslogTargets                  types.Set    `tfsdk:"syslog_targets"`
	ListenHttpAddress              types.Set    `tfsdk:"listen_http_address"`
	ListenHttpsAddress             types.Set    `tfsdk:"listen_https_address"`
	DefaultServer                  types.Bool   `tfsdk:"default_server"`
	TLSRejectHandshake             types.Bool   `tfsdk:"tls_reject_handshake"`
	ProxyProtocol                  types.Bool   `tfsdk:"proxy_protocol"`
	TrustedProxies                 types.Set    `tfsdk:"trusted_proxies"`
	TrustedProxiesAlias            types.String `tfsdk:"trusted_proxies_alias"`
	RealIPSource                   types.String `tfsdk:"real_ip_source"`
	Locations                      types.Set    `tfsdk:"locations"`
	Rewrites                       types.Set    `tfsdk:"rewrites"`
	Root                           types.String `tfsdk:"root"`
	Certificate                    types.String `tfsdk:"certificate"`
	Ca                             types.Set    `tfsdk:"ca"`
	VerifyClient                   types.String `tfsdk:"verify_client"`
	AccessLogFormat                types.String `tfsdk:"access_log_format"`
	ErrorLogLevel                  types.String `tfsdk:"error_log_level"`
	LogHandshakes                  types.Bool   `tfsdk:"log_handshakes"`
	EnableAcmeSupport              types.Bool   `tfsdk:"enable_acme_support"`
	Charset                        types.String `tfsdk:"charset"`
	HttpsOnly                      types.Bool   `tfsdk:"https_only"`
	Http2                          types.Bool   `tfsdk:"http2"`
	EnableHttp3                    types.Bool   `tfsdk:"enable_http3"`
	TLSProtocols                   types.Set    `tfsdk:"tls_protocols"`
	TLSCiphers                     types.String `tfsdk:"tls_ciphers"`
	TLSEcdhCurve                   types.String `tfsdk:"tls_ecdh_curve"`
	TLSPreferServerCiphers         types.Bool   `tfsdk:"tls_prefer_server_ciphers"`
	Resolver                       types.String `tfsdk:"resolver"`
	OcspStapling                   types.Bool   `tfsdk:"ocsp_stapling"`
	OcspVerify                     types.Bool   `tfsdk:"ocsp_verify"`
	BlockNonpublicData             types.Bool   `tfsdk:"block_nonpublic_data"`
	DisableBotProtection           types.Bool   `tfsdk:"disable_bot_protection"`
	DisableGzip                    types.Bool   `tfsdk:"disable_gzip"`
	NaxsiWhitelistSrcip            types.Set    `tfsdk:"naxsi_whitelist_srcip"`
	NaxsiExtensiveLog              types.Bool   `tfsdk:"naxsi_extensive_log"`
	Sendfile                       types.Bool   `tfsdk:"sendfile"`
	ClientHeaderBufferSize         types.Int64  `tfsdk:"client_header_buffer_size"`
	LargeClientHeaderBuffersNumber types.Int64  `tfsdk:"large_client_header_buffers_number"`
	LargeClientHeaderBuffersSize   types.Int64  `tfsdk:"large_client_header_buffers_size"`
	SecurityHeader                 types.String `tfsdk:"security_header"`
	LimitRequestConnections        types.Set    `tfsdk:"limit_request_connections"`
	MaxBodySize                    types.String `tfsdk:"max_body_size"`
	BodyBufferSize                 types.String `tfsdk:"body_buffer_size"`
	IPAcl                          types.String `tfsdk:"ip_acl"`
	AdvancedAclServer              types.String `tfsdk:"advanced_acl_server"`
	Satisfy                        types.String `tfsdk:"satisfy"`
	ZeroRtt                        types.Bool   `tfsdk:"zero_rtt"`
	Errorpages                     types.Set    `tfsdk:"errorpages"`
	ProxyInterceptErrors           types.Bool   `tfsdk:"proxy_intercept_errors"`

	Id types.String `tfsdk:"id"`
}

func httpServerResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx HTTP servers (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ HTTP Server**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"servername": schema.SetAttribute{
				MarkdownDescription: "Server Name.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"syslog_targets": schema.SetAttribute{
				MarkdownDescription: "SYSLOG targets. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"listen_http_address": schema.SetAttribute{
				MarkdownDescription: "Enter a list of IP addresses and ports which can be used in nginx listen directives. To listen on a port on all IPs, use for example \"80,[::]:80\". Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"listen_https_address": schema.SetAttribute{
				MarkdownDescription: "Enter a list of IP addresses and ports which can be used in nginx listen directives. To listen on a port on all IPs, use for example \"443,[::]:443\". Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"default_server": schema.BoolAttribute{
				MarkdownDescription: "Default Server. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tls_reject_handshake": schema.BoolAttribute{
				MarkdownDescription: "If enabled, TLS handshakes for this server will be rejected. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"proxy_protocol": schema.BoolAttribute{
				MarkdownDescription: "If you enable the proxy protocol, a downstream proxy can send the client IP and port before the real traffic is set. Defaults to `false`.",
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
			"trusted_proxies_alias": schema.StringAttribute{
				MarkdownDescription: "Choose a Firewall Alias for trusted proxies. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"real_ip_source": schema.StringAttribute{
				MarkdownDescription: "X-Real-IP and X-Forwarded-For are HTTP headers, while PROXY protocol is a protocol which needs to be enabled. One of `X-Real-IP`, `X-Forwarded-For`, `proxy_protocol`, `CF-Connecting-IP`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "X-Real-IP", "X-Forwarded-For", "proxy_protocol", "CF-Connecting-IP")},
			},
			"locations": schema.SetAttribute{
				MarkdownDescription: "Locations. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"rewrites": schema.SetAttribute{
				MarkdownDescription: "Choose URL rewriting rules. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"root": schema.StringAttribute{
				MarkdownDescription: "File System Root. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"certificate": schema.StringAttribute{
				MarkdownDescription: "TLS Certificate. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ca": schema.SetAttribute{
				MarkdownDescription: "Trusted CA certificates. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
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
				MarkdownDescription: "Select Error Log Level. Log levels are listed in the order of increasing verbosity. Setting a certain log level will cause all messages of the specified and more severe log levels to be logged. One of `emerg`, `alert`, `crit`, `error`, `warn`, `notice`, `info`. Defaults to `error`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("error"),
				Validators:          []validator.String{stringvalidator.OneOf("emerg", "alert", "crit", "error", "warn", "notice", "info")},
			},
			"log_handshakes": schema.BoolAttribute{
				MarkdownDescription: "Log TLS handshakes to fill the User Agent fingerprint database and detect MITM attacks. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"enable_acme_support": schema.BoolAttribute{
				MarkdownDescription: "Enable Let's Encrypt Plugin Support. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"charset": schema.StringAttribute{
				MarkdownDescription: "Charset. One of `utf-8`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "utf-8")},
			},
			"https_only": schema.BoolAttribute{
				MarkdownDescription: "If the request scheme is not HTTPS, redirect to use HTTPS for this server. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http2": schema.BoolAttribute{
				MarkdownDescription: "Enable the HTTP/2 protocol. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"enable_http3": schema.BoolAttribute{
				MarkdownDescription: "Enable HTTP/3/QUIC for this server (adds QUIC listeners and Alt-Svc header). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tls_protocols": schema.SetAttribute{
				MarkdownDescription: "Enable specified protocols. Default: TLSv1.2 TLSv1.3. TLSv1 and TLSv1.1 are insecure and excluded from GUI. One of `TLSv1.2`, `TLSv1.3`. Defaults to `[\"TLSv1.2\", \"TLSv1.3\"]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(types.SetValueMust(types.StringType, []attr.Value{types.StringValue("TLSv1.2"), types.StringValue("TLSv1.3")})),
				Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("TLSv1.2", "TLSv1.3"))},
			},
			"tls_ciphers": schema.StringAttribute{
				MarkdownDescription: "Specify enabled ciphers in the TLS library cipher string format. Example: ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256. NGINX default (empty field): HIGH:!aNULL:!MD5. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tls_ecdh_curve": schema.StringAttribute{
				MarkdownDescription: "Specify a curve(s) for ECDHE ciphers. NGINX default (empty field): auto. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tls_prefer_server_ciphers": schema.BoolAttribute{
				MarkdownDescription: "Prefers server ciphers over client ciphers. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"resolver": schema.StringAttribute{
				MarkdownDescription: "Specify resolver for Upstream names resolution, OCSP stapling etc. Uses system resolver if not specified but can produce warning messages in log. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ocsp_stapling": schema.BoolAttribute{
				MarkdownDescription: "If you check this box, OCSP stapling is enabled. If your certificate has the TLS feature \"status_request\" enabled and don't enable OCSP Stapling, you could get an TLS error like \"MOZILLA_PKIX_ERROR_REQUIRED_TLS_FEATURE_MISSING\". This can be the case if you use Let's Encrypt with \"OCSP Must Staple\". More information about OCSP stapling can be found in RFC 7633. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ocsp_verify": schema.BoolAttribute{
				MarkdownDescription: "Enables or disables verification of OCSP responses by the server. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"block_nonpublic_data": schema.BoolAttribute{
				MarkdownDescription: "Blocks files like .htaccess files or other files not intended for the public. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"disable_bot_protection": schema.BoolAttribute{
				MarkdownDescription: "Blocks the request when a possibly bad bot is detected and adds the originating IP to the managed firewall alias for permanent blocking. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"disable_gzip": schema.BoolAttribute{
				MarkdownDescription: "Disables responses gzipping. Makes sense if TLS is used and this method of BREACH attack protection is preferred. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"naxsi_whitelist_srcip": schema.SetAttribute{
				MarkdownDescription: "Enter a list of IP addresses or CIDR networks which will be whitelisted for the Naxsi rules. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"naxsi_extensive_log": schema.BoolAttribute{
				MarkdownDescription: "Provide a more verbose WAF log for fixing false positives before going live. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"sendfile": schema.BoolAttribute{
				MarkdownDescription: "Allow the daemon to use the sendfile function. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"client_header_buffer_size": schema.Int64Attribute{
				MarkdownDescription: "Contains the size of a header Buffer. If a header or the request line does not fit into it, a large header buffer will be used. Defaults to `1`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1),
			},
			"large_client_header_buffers_number": schema.Int64Attribute{
				MarkdownDescription: "This field contains the amount of large header buffers. If there are not enough, the server will respond with an HTTP error. Defaults to `4`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(4),
			},
			"large_client_header_buffers_size": schema.Int64Attribute{
				MarkdownDescription: "This field contains the size of large header buffers. If a header or the request line does not fit into it, the server will respond with an HTTP error. Defaults to `8`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(8),
			},
			"security_header": schema.StringAttribute{
				MarkdownDescription: "Security Header. Defaults to ``.",
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
			"ip_acl": schema.StringAttribute{
				MarkdownDescription: "If you select an IP ACL, the client can only access this service if it fulfills this requirement. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"advanced_acl_server": schema.StringAttribute{
				MarkdownDescription: "Choose a backend for advanced authentication. If you don't have one, configure it under System -> Access -> Servers. If nothing is configured, the local database is used. Defaults to ``.",
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
			"zero_rtt": schema.BoolAttribute{
				MarkdownDescription: "If you enable this feature, your website will may get vulnerable to replay attacks, but you gain a big performance boost. You can read RFC 8446 for more information. The Backend will get a HTTP header \"Early-Data\" which can be used if a short handshake was used. If a response is too risky, answer with HTTP status code 425. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"errorpages": schema.SetAttribute{
				MarkdownDescription: "Select custom error pages to display instead of the default builtin error pages. If at least one error page is selected here, all default error pages will be disabled. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"proxy_intercept_errors": schema.BoolAttribute{
				MarkdownDescription: "Intercept responses with codes greater than or equal to 300 and redirect to processing with custom error pages. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx HTTP server.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func httpServerDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx HTTP servers (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ HTTP Server**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx HTTP server.",
				Required:            true,
			},
			"servername":                         dschema.SetAttribute{MarkdownDescription: "Server Name.", Computed: true, ElementType: types.StringType},
			"syslog_targets":                     dschema.SetAttribute{MarkdownDescription: "SYSLOG targets.", Computed: true, ElementType: types.StringType},
			"listen_http_address":                dschema.SetAttribute{MarkdownDescription: "Enter a list of IP addresses and ports which can be used in nginx listen directives. To listen on a port on all IPs, use for example \"80,[::]:80\".", Computed: true, ElementType: types.StringType},
			"listen_https_address":               dschema.SetAttribute{MarkdownDescription: "Enter a list of IP addresses and ports which can be used in nginx listen directives. To listen on a port on all IPs, use for example \"443,[::]:443\".", Computed: true, ElementType: types.StringType},
			"default_server":                     dschema.BoolAttribute{MarkdownDescription: "Default Server.", Computed: true},
			"tls_reject_handshake":               dschema.BoolAttribute{MarkdownDescription: "If enabled, TLS handshakes for this server will be rejected.", Computed: true},
			"proxy_protocol":                     dschema.BoolAttribute{MarkdownDescription: "If you enable the proxy protocol, a downstream proxy can send the client IP and port before the real traffic is set.", Computed: true},
			"trusted_proxies":                    dschema.SetAttribute{MarkdownDescription: "Enter a list of IP addresses or CIDR networks which are allowed to override the source IP address using the specified header.", Computed: true, ElementType: types.StringType},
			"trusted_proxies_alias":              dschema.StringAttribute{MarkdownDescription: "Choose a Firewall Alias for trusted proxies.", Computed: true},
			"real_ip_source":                     dschema.StringAttribute{MarkdownDescription: "X-Real-IP and X-Forwarded-For are HTTP headers, while PROXY protocol is a protocol which needs to be enabled. One of `X-Real-IP`, `X-Forwarded-For`, `proxy_protocol`, `CF-Connecting-IP`.", Computed: true},
			"locations":                          dschema.SetAttribute{MarkdownDescription: "Locations.", Computed: true, ElementType: types.StringType},
			"rewrites":                           dschema.SetAttribute{MarkdownDescription: "Choose URL rewriting rules.", Computed: true, ElementType: types.StringType},
			"root":                               dschema.StringAttribute{MarkdownDescription: "File System Root.", Computed: true},
			"certificate":                        dschema.StringAttribute{MarkdownDescription: "TLS Certificate.", Computed: true},
			"ca":                                 dschema.SetAttribute{MarkdownDescription: "Trusted CA certificates.", Computed: true, ElementType: types.StringType},
			"verify_client":                      dschema.StringAttribute{MarkdownDescription: "On: the certificate is requested and validated. Use this option to protect a service with TLS authentication.Off: The certificate is not requested. Choose this option for a normal website.Optional: The certificate is requested and validated if existing. Choose this option for websites, with TLS login support or mixed TLS protected API and web content.Optional, don't verify: Do accept the certificate and let the application choose what to do. Choose this option, for the same reasons as optional but in this case, the request is passed to the backend without rejecting untrusted certificates. One of `off`, `on`, `optional`, `optional_no_ca`.", Computed: true},
			"access_log_format":                  dschema.StringAttribute{MarkdownDescription: "Access Log Format. One of `main`, `main_ext`, `anonymized`, `disabled`.", Computed: true},
			"error_log_level":                    dschema.StringAttribute{MarkdownDescription: "Select Error Log Level. Log levels are listed in the order of increasing verbosity. Setting a certain log level will cause all messages of the specified and more severe log levels to be logged. One of `emerg`, `alert`, `crit`, `error`, `warn`, `notice`, `info`.", Computed: true},
			"log_handshakes":                     dschema.BoolAttribute{MarkdownDescription: "Log TLS handshakes to fill the User Agent fingerprint database and detect MITM attacks.", Computed: true},
			"enable_acme_support":                dschema.BoolAttribute{MarkdownDescription: "Enable Let's Encrypt Plugin Support.", Computed: true},
			"charset":                            dschema.StringAttribute{MarkdownDescription: "Charset. One of `utf-8`.", Computed: true},
			"https_only":                         dschema.BoolAttribute{MarkdownDescription: "If the request scheme is not HTTPS, redirect to use HTTPS for this server.", Computed: true},
			"http2":                              dschema.BoolAttribute{MarkdownDescription: "Enable the HTTP/2 protocol.", Computed: true},
			"enable_http3":                       dschema.BoolAttribute{MarkdownDescription: "Enable HTTP/3/QUIC for this server (adds QUIC listeners and Alt-Svc header).", Computed: true},
			"tls_protocols":                      dschema.SetAttribute{MarkdownDescription: "Enable specified protocols. Default: TLSv1.2 TLSv1.3. TLSv1 and TLSv1.1 are insecure and excluded from GUI. One of `TLSv1.2`, `TLSv1.3`.", Computed: true, ElementType: types.StringType},
			"tls_ciphers":                        dschema.StringAttribute{MarkdownDescription: "Specify enabled ciphers in the TLS library cipher string format. Example: ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256. NGINX default (empty field): HIGH:!aNULL:!MD5.", Computed: true},
			"tls_ecdh_curve":                     dschema.StringAttribute{MarkdownDescription: "Specify a curve(s) for ECDHE ciphers. NGINX default (empty field): auto.", Computed: true},
			"tls_prefer_server_ciphers":          dschema.BoolAttribute{MarkdownDescription: "Prefers server ciphers over client ciphers.", Computed: true},
			"resolver":                           dschema.StringAttribute{MarkdownDescription: "Specify resolver for Upstream names resolution, OCSP stapling etc. Uses system resolver if not specified but can produce warning messages in log.", Computed: true},
			"ocsp_stapling":                      dschema.BoolAttribute{MarkdownDescription: "If you check this box, OCSP stapling is enabled. If your certificate has the TLS feature \"status_request\" enabled and don't enable OCSP Stapling, you could get an TLS error like \"MOZILLA_PKIX_ERROR_REQUIRED_TLS_FEATURE_MISSING\". This can be the case if you use Let's Encrypt with \"OCSP Must Staple\". More information about OCSP stapling can be found in RFC 7633.", Computed: true},
			"ocsp_verify":                        dschema.BoolAttribute{MarkdownDescription: "Enables or disables verification of OCSP responses by the server.", Computed: true},
			"block_nonpublic_data":               dschema.BoolAttribute{MarkdownDescription: "Blocks files like .htaccess files or other files not intended for the public.", Computed: true},
			"disable_bot_protection":             dschema.BoolAttribute{MarkdownDescription: "Blocks the request when a possibly bad bot is detected and adds the originating IP to the managed firewall alias for permanent blocking.", Computed: true},
			"disable_gzip":                       dschema.BoolAttribute{MarkdownDescription: "Disables responses gzipping. Makes sense if TLS is used and this method of BREACH attack protection is preferred.", Computed: true},
			"naxsi_whitelist_srcip":              dschema.SetAttribute{MarkdownDescription: "Enter a list of IP addresses or CIDR networks which will be whitelisted for the Naxsi rules.", Computed: true, ElementType: types.StringType},
			"naxsi_extensive_log":                dschema.BoolAttribute{MarkdownDescription: "Provide a more verbose WAF log for fixing false positives before going live.", Computed: true},
			"sendfile":                           dschema.BoolAttribute{MarkdownDescription: "Allow the daemon to use the sendfile function.", Computed: true},
			"client_header_buffer_size":          dschema.Int64Attribute{MarkdownDescription: "Contains the size of a header Buffer. If a header or the request line does not fit into it, a large header buffer will be used.", Computed: true},
			"large_client_header_buffers_number": dschema.Int64Attribute{MarkdownDescription: "This field contains the amount of large header buffers. If there are not enough, the server will respond with an HTTP error.", Computed: true},
			"large_client_header_buffers_size":   dschema.Int64Attribute{MarkdownDescription: "This field contains the size of large header buffers. If a header or the request line does not fit into it, the server will respond with an HTTP error.", Computed: true},
			"security_header":                    dschema.StringAttribute{MarkdownDescription: "Security Header.", Computed: true},
			"limit_request_connections":          dschema.SetAttribute{MarkdownDescription: "If you choose multiple limits, the strictest will be used.", Computed: true, ElementType: types.StringType},
			"max_body_size":                      dschema.StringAttribute{MarkdownDescription: "If the request is larger, it will be rejectet with error 413 (Request Entity Too Large). For example, you can enter 200m.", Computed: true},
			"body_buffer_size":                   dschema.StringAttribute{MarkdownDescription: "If the request exceeds this size, it will be written to disk. Enter a number and a unit like 1m.", Computed: true},
			"ip_acl":                             dschema.StringAttribute{MarkdownDescription: "If you select an IP ACL, the client can only access this service if it fulfills this requirement.", Computed: true},
			"advanced_acl_server":                dschema.StringAttribute{MarkdownDescription: "Choose a backend for advanced authentication. If you don't have one, configure it under System -> Access -> Servers. If nothing is configured, the local database is used.", Computed: true},
			"satisfy":                            dschema.StringAttribute{MarkdownDescription: "All: All access restrictions must be fulfilled; Any: Any of the access restrictions must be fulfilled. One of `any`, `all`.", Computed: true},
			"zero_rtt":                           dschema.BoolAttribute{MarkdownDescription: "If you enable this feature, your website will may get vulnerable to replay attacks, but you gain a big performance boost. You can read RFC 8446 for more information. The Backend will get a HTTP header \"Early-Data\" which can be used if a short handshake was used. If a response is too risky, answer with HTTP status code 425.", Computed: true},
			"errorpages":                         dschema.SetAttribute{MarkdownDescription: "Select custom error pages to display instead of the default builtin error pages. If at least one error page is selected here, all default error pages will be disabled.", Computed: true, ElementType: types.StringType},
			"proxy_intercept_errors":             dschema.BoolAttribute{MarkdownDescription: "Intercept responses with codes greater than or equal to 300 and redirect to processing with custom error pages.", Computed: true},
		},
	}
}

func convertHttpServerSchemaToStruct(d *httpServerResourceModel) (*nginx.HttpServer, error) {
	return &nginx.HttpServer{
		Servername:                     api.SelectedMapList(setToSlice(d.Servername)),
		SyslogTargets:                  api.SelectedMapList(setToSlice(d.SyslogTargets)),
		ListenHttpAddress:              api.SelectedMapList(setToSlice(d.ListenHttpAddress)),
		ListenHttpsAddress:             api.SelectedMapList(setToSlice(d.ListenHttpsAddress)),
		DefaultServer:                  tools.BoolToString(d.DefaultServer.ValueBool()),
		TlsRejectHandshake:             tools.BoolToString(d.TLSRejectHandshake.ValueBool()),
		ProxyProtocol:                  tools.BoolToString(d.ProxyProtocol.ValueBool()),
		TrustedProxies:                 api.SelectedMapList(setToSlice(d.TrustedProxies)),
		TrustedProxiesAlias:            api.SelectedMap(d.TrustedProxiesAlias.ValueString()),
		RealIpSource:                   api.SelectedMap(d.RealIPSource.ValueString()),
		Locations:                      api.SelectedMapList(setToSlice(d.Locations)),
		Rewrites:                       api.SelectedMapList(setToSlice(d.Rewrites)),
		Root:                           d.Root.ValueString(),
		Certificate:                    api.SelectedMap(d.Certificate.ValueString()),
		Ca:                             api.SelectedMapList(setToSlice(d.Ca)),
		VerifyClient:                   api.SelectedMap(d.VerifyClient.ValueString()),
		AccessLogFormat:                api.SelectedMap(d.AccessLogFormat.ValueString()),
		ErrorLogLevel:                  api.SelectedMap(d.ErrorLogLevel.ValueString()),
		LogHandshakes:                  tools.BoolToString(d.LogHandshakes.ValueBool()),
		EnableAcmeSupport:              tools.BoolToString(d.EnableAcmeSupport.ValueBool()),
		Charset:                        api.SelectedMap(d.Charset.ValueString()),
		HttpsOnly:                      tools.BoolToString(d.HttpsOnly.ValueBool()),
		Http2:                          tools.BoolToString(d.Http2.ValueBool()),
		EnableHttp3:                    tools.BoolToString(d.EnableHttp3.ValueBool()),
		TlsProtocols:                   api.SelectedMapList(setToSlice(d.TLSProtocols)),
		TlsCiphers:                     d.TLSCiphers.ValueString(),
		TlsEcdhCurve:                   d.TLSEcdhCurve.ValueString(),
		TlsPreferServerCiphers:         tools.BoolToString(d.TLSPreferServerCiphers.ValueBool()),
		Resolver:                       api.SelectedMap(d.Resolver.ValueString()),
		OcspStapling:                   tools.BoolToString(d.OcspStapling.ValueBool()),
		OcspVerify:                     tools.BoolToString(d.OcspVerify.ValueBool()),
		BlockNonpublicData:             tools.BoolToString(d.BlockNonpublicData.ValueBool()),
		DisableBotProtection:           tools.BoolToString(d.DisableBotProtection.ValueBool()),
		DisableGzip:                    tools.BoolToString(d.DisableGzip.ValueBool()),
		NaxsiWhitelistSrcip:            api.SelectedMapList(setToSlice(d.NaxsiWhitelistSrcip)),
		NaxsiExtensiveLog:              tools.BoolToString(d.NaxsiExtensiveLog.ValueBool()),
		Sendfile:                       tools.BoolToString(d.Sendfile.ValueBool()),
		ClientHeaderBufferSize:         tools.Int64ToString(d.ClientHeaderBufferSize.ValueInt64()),
		LargeClientHeaderBuffersNumber: tools.Int64ToString(d.LargeClientHeaderBuffersNumber.ValueInt64()),
		LargeClientHeaderBuffersSize:   tools.Int64ToString(d.LargeClientHeaderBuffersSize.ValueInt64()),
		SecurityHeader:                 api.SelectedMap(d.SecurityHeader.ValueString()),
		LimitRequestConnections:        api.SelectedMapList(setToSlice(d.LimitRequestConnections)),
		MaxBodySize:                    d.MaxBodySize.ValueString(),
		BodyBufferSize:                 d.BodyBufferSize.ValueString(),
		IpAcl:                          api.SelectedMap(d.IPAcl.ValueString()),
		AdvancedAclServer:              api.SelectedMap(d.AdvancedAclServer.ValueString()),
		Satisfy:                        api.SelectedMap(d.Satisfy.ValueString()),
		ZeroRtt:                        tools.BoolToString(d.ZeroRtt.ValueBool()),
		Errorpages:                     api.SelectedMapList(setToSlice(d.Errorpages)),
		ProxyInterceptErrors:           tools.BoolToString(d.ProxyInterceptErrors.ValueBool()),
	}, nil
}

func convertHttpServerStructToSchema(d *nginx.HttpServer) (*httpServerResourceModel, error) {
	return &httpServerResourceModel{
		Servername:                     tools.StringSliceToSet(d.Servername),
		SyslogTargets:                  tools.StringSliceToSet(d.SyslogTargets),
		ListenHttpAddress:              tools.StringSliceToSet(d.ListenHttpAddress),
		ListenHttpsAddress:             tools.StringSliceToSet(d.ListenHttpsAddress),
		DefaultServer:                  types.BoolValue(tools.StringToBool(d.DefaultServer)),
		TLSRejectHandshake:             types.BoolValue(tools.StringToBool(d.TlsRejectHandshake)),
		ProxyProtocol:                  types.BoolValue(tools.StringToBool(d.ProxyProtocol)),
		TrustedProxies:                 tools.StringSliceToSet(d.TrustedProxies),
		TrustedProxiesAlias:            types.StringValue(d.TrustedProxiesAlias.String()),
		RealIPSource:                   types.StringValue(d.RealIpSource.String()),
		Locations:                      tools.StringSliceToSet(d.Locations),
		Rewrites:                       tools.StringSliceToSet(d.Rewrites),
		Root:                           types.StringValue(d.Root),
		Certificate:                    types.StringValue(d.Certificate.String()),
		Ca:                             tools.StringSliceToSet(d.Ca),
		VerifyClient:                   types.StringValue(d.VerifyClient.String()),
		AccessLogFormat:                types.StringValue(d.AccessLogFormat.String()),
		ErrorLogLevel:                  types.StringValue(d.ErrorLogLevel.String()),
		LogHandshakes:                  types.BoolValue(tools.StringToBool(d.LogHandshakes)),
		EnableAcmeSupport:              types.BoolValue(tools.StringToBool(d.EnableAcmeSupport)),
		Charset:                        types.StringValue(d.Charset.String()),
		HttpsOnly:                      types.BoolValue(tools.StringToBool(d.HttpsOnly)),
		Http2:                          types.BoolValue(tools.StringToBool(d.Http2)),
		EnableHttp3:                    types.BoolValue(tools.StringToBool(d.EnableHttp3)),
		TLSProtocols:                   tools.StringSliceToSet(d.TlsProtocols),
		TLSCiphers:                     types.StringValue(d.TlsCiphers),
		TLSEcdhCurve:                   types.StringValue(d.TlsEcdhCurve),
		TLSPreferServerCiphers:         types.BoolValue(tools.StringToBool(d.TlsPreferServerCiphers)),
		Resolver:                       types.StringValue(d.Resolver.String()),
		OcspStapling:                   types.BoolValue(tools.StringToBool(d.OcspStapling)),
		OcspVerify:                     types.BoolValue(tools.StringToBool(d.OcspVerify)),
		BlockNonpublicData:             types.BoolValue(tools.StringToBool(d.BlockNonpublicData)),
		DisableBotProtection:           types.BoolValue(tools.StringToBool(d.DisableBotProtection)),
		DisableGzip:                    types.BoolValue(tools.StringToBool(d.DisableGzip)),
		NaxsiWhitelistSrcip:            tools.StringSliceToSet(d.NaxsiWhitelistSrcip),
		NaxsiExtensiveLog:              types.BoolValue(tools.StringToBool(d.NaxsiExtensiveLog)),
		Sendfile:                       types.BoolValue(tools.StringToBool(d.Sendfile)),
		ClientHeaderBufferSize:         types.Int64Value(tools.StringToInt64(d.ClientHeaderBufferSize)),
		LargeClientHeaderBuffersNumber: types.Int64Value(tools.StringToInt64(d.LargeClientHeaderBuffersNumber)),
		LargeClientHeaderBuffersSize:   types.Int64Value(tools.StringToInt64(d.LargeClientHeaderBuffersSize)),
		SecurityHeader:                 types.StringValue(d.SecurityHeader.String()),
		LimitRequestConnections:        tools.StringSliceToSet(d.LimitRequestConnections),
		MaxBodySize:                    types.StringValue(d.MaxBodySize),
		BodyBufferSize:                 types.StringValue(d.BodyBufferSize),
		IPAcl:                          types.StringValue(d.IpAcl.String()),
		AdvancedAclServer:              types.StringValue(d.AdvancedAclServer.String()),
		Satisfy:                        types.StringValue(d.Satisfy.String()),
		ZeroRtt:                        types.BoolValue(tools.StringToBool(d.ZeroRtt)),
		Errorpages:                     tools.StringSliceToSet(d.Errorpages),
		ProxyInterceptErrors:           types.BoolValue(tools.StringToBool(d.ProxyInterceptErrors)),
	}, nil
}
