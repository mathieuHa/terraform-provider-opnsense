package caddy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/caddy"
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
	Enabled                        types.Bool   `tfsdk:"enabled"`
	EnableLayer4                   types.Bool   `tfsdk:"enable_layer4"`
	StrictSniHost                  types.Bool   `tfsdk:"strict_sni_host"`
	HttpPort                       types.Int64  `tfsdk:"http_port"`
	HttpsPort                      types.Int64  `tfsdk:"https_port"`
	TlsEmail                       types.String `tfsdk:"tls_email"`
	TlsAutoHttps                   types.String `tfsdk:"tls_auto_https"`
	TlsDnsProvider                 types.String `tfsdk:"tls_dns_provider"`
	TlsDnsApiKey                   types.String `tfsdk:"tls_dns_api_key"`
	TlsDnsPropagationTimeout       types.Bool   `tfsdk:"tls_dns_propagation_timeout"`
	TlsDnsPropagationTimeoutPeriod types.Int64  `tfsdk:"tls_dns_propagation_timeout_period"`
	TlsDnsPropagationDelay         types.Int64  `tfsdk:"tls_dns_propagation_delay"`
	TlsDnsPropagationResolvers     types.String `tfsdk:"tls_dns_propagation_resolvers"`
	TlsDnsEchDomain                types.String `tfsdk:"tls_dns_ech_domain"`
	Accesslist                     types.String `tfsdk:"accesslist"`
	ClientIpHeaders                types.Set    `tfsdk:"client_ip_headers"`
	ProxyProtocolAllow             types.Set    `tfsdk:"proxy_protocol_allow"`
	DisableSuperuser               types.String `tfsdk:"disable_superuser"`
	GracePeriod                    types.Int64  `tfsdk:"grace_period"`
	HttpVersions                   types.Set    `tfsdk:"http_versions"`
	TimeoutReadBody                types.Int64  `tfsdk:"timeout_read_body"`
	TimeoutReadHeader              types.Int64  `tfsdk:"timeout_read_header"`
	TimeoutWrite                   types.Int64  `tfsdk:"timeout_write"`
	TimeoutIdle                    types.Int64  `tfsdk:"timeout_idle"`
	LogCredentials                 types.Bool   `tfsdk:"log_credentials"`
	LogAccessPlain                 types.Bool   `tfsdk:"log_access_plain"`
	LogAccessPlainKeep             types.Int64  `tfsdk:"log_access_plain_keep"`
	LogLevel                       types.String `tfsdk:"log_level"`
	DynDnsSimpleHttp               types.String `tfsdk:"dyn_dns_simple_http"`
	DynDnsInterface                types.String `tfsdk:"dyn_dns_interface"`
	DynDnsInterval                 types.Int64  `tfsdk:"dyn_dns_interval"`
	DynDnsIpVersions               types.String `tfsdk:"dyn_dns_ip_versions"`
	DynDnsTtl                      types.Int64  `tfsdk:"dyn_dns_ttl"`
	DynDnsUpdateOnly               types.Bool   `tfsdk:"dyn_dns_update_only"`
	AuthProvider                   types.String `tfsdk:"auth_provider"`
	AuthToDomain                   types.String `tfsdk:"auth_to_domain"`
	AuthToPort                     types.Int64  `tfsdk:"auth_to_port"`
	AuthToTls                      types.String `tfsdk:"auth_to_tls"`
	AuthToUri                      types.String `tfsdk:"auth_to_uri"`
	CopyHeaders                    types.Set    `tfsdk:"copy_headers"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable or disable the Caddy web server.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"enable_layer4": schema.BoolAttribute{
			MarkdownDescription: "Enable Layer4 Proxy to stream TCP/UDP. Layer4 Proxy matches before the Reverse Proxy; both can be used concurrently.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"strict_sni_host": schema.BoolAttribute{
			MarkdownDescription: "Reject TLS requests with HTTP status 421 when the Host header does not match the Server Name Indication (SNI). This can prevent HTTP/2 connection coalescing from bypassing Layer4 SNI routing.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"http_port": schema.Int64Attribute{
			MarkdownDescription: "If the default HTTP port is changed to e.g. 8080, then a port forward from port 80 to 8080 is necessary to issue automatic certificates with the HTTP-01 challenge and serve clients the reverse proxied resources. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"https_port": schema.Int64Attribute{
			MarkdownDescription: "If the default HTTPS port is changed to e.g. 8443, then a port forward from port 443 to 8443 is necessary to issue automatic certificates with the TLS-ALPN-01 challenge and serve clients the reverse proxied resources. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tls_email": schema.StringAttribute{
			MarkdownDescription: "Enter the email address for certificate notifications. This is required to receive automatic certificates from \"Let's Encrypt\" and \"ZeroSSL\".",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"tls_auto_https": schema.StringAttribute{
			MarkdownDescription: "Select the Auto HTTPS option. \"On (default)\" creates automatic certificates using \"Let's Encrypt\" or \"ZeroSSL\". \"Off\" turns all automatic certificate requests off. One of `off`, `disable_redirects`, `disable_certs`, `ignore_loaded_certs`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"tls_dns_provider": schema.StringAttribute{
			MarkdownDescription: "Select the DNS Provider. If you cannot find your provider here, consider using os-acme-client for the DNS-01 Challenge and os-ddclient for Dynamic DNS as alternatives. One of `cloudflare`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"tls_dns_api_key": schema.StringAttribute{
			MarkdownDescription: "This is the standard field for the API Key.",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"tls_dns_propagation_timeout": schema.BoolAttribute{
			MarkdownDescription: "This will disable propagation_timeout.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"tls_dns_propagation_timeout_period": schema.Int64Attribute{
			MarkdownDescription: "propagation_timeout is a duration value in seconds that sets the maximum time to wait for the DNS TXT records to appear when using the DNS challenge. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tls_dns_propagation_delay": schema.Int64Attribute{
			MarkdownDescription: "propagation_delay is a duration value in seconds that sets how long to wait before starting DNS TXT records propagation checks when using the DNS challenge. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tls_dns_propagation_resolvers": schema.StringAttribute{
			MarkdownDescription: "Leave empty to use the system resolvers (default). Resolvers customizes the DNS resolvers used when performing the DNS challenge; these take precedence over system resolvers or any default ones. If set here, the resolvers will propagate to all configured certificate issuers. If the system resolvers use DNS over TLS, setting an external resolver here is required or the DNS challenge will fail.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"tls_dns_ech_domain": schema.StringAttribute{
			MarkdownDescription: "Enables Encrypted ClientHello (ECH) by using the specified public domain name as the plaintext server name (SNI) in TLS handshakes.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"accesslist": schema.StringAttribute{
			MarkdownDescription: "Select an Access List to set IP ranges of Trusted Proxies. Access Lists can be added in \"Reverse Proxy - Access\". If Caddy is not the first server being connected to by clients (for example, when a \"CDN\" is in front of Caddy), configure \"Trusted Proxies\" with a list of IP ranges (CIDRs) from which incoming requests are trusted to have sent good values for these headers. Additionally, set the same Access List to the domains the Trusted Proxies connect to.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"client_ip_headers": schema.SetAttribute{
			MarkdownDescription: "Select one or multiple headers to extract the real client IP. Headers can be added in \"Reverse Proxy - Headers\". As example, setting \"X-Forwarded-For\" and \"Cf-Connecting-Ip\" can extract the real client IP from Cloudflare.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"proxy_protocol_allow": schema.SetAttribute{
			MarkdownDescription: "Enter one or multiple trusted proxy source IP addresses or networks in CIDR notation. PROXY protocol headers received from these sources will be accepted and used to determine the original client connection information. Only applies when Layer 4 Proxy is enabled.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"disable_superuser": schema.StringAttribute{
			MarkdownDescription: "Run this service as \"www\" user and group, instead of \"root\". This setting increases security, but comes with the hard restriction that the well-known port range can not be used anymore. After enabling and saving this setting, the service has to be totally restarted. For this, please disable Caddy and press Apply. Afterwards enable Caddy and press Apply. This setting is reversible by following the same steps. One of `0`, `1`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"grace_period": schema.Int64Attribute{
			MarkdownDescription: "Defines the grace period for shutting down Caddy during a reload in seconds. If clients do not finish their requests within the grace period, the server will be forcefully terminated to allow the reload to complete and free up resources. This can influence how long \"Apply\" of new configurations take, since Caddy waits for all open connections to close. If the grace period is over and Caddy is unresponsive, there will be a forced kill and service restart. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"http_versions": schema.SetAttribute{
			MarkdownDescription: "Select the HTTP versions for the frontend listeners. By default, QUIC (HTTP/3) is disabled. One of `h1`, `h2`, `h3`.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"timeout_read_body": schema.Int64Attribute{
			MarkdownDescription: "read_body is a duration value in seconds that sets how long to allow a read from a client's upload. Setting this to a short, non-zero value can mitigate slowloris attacks, but may also affect legitimately slow clients. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"timeout_read_header": schema.Int64Attribute{
			MarkdownDescription: "read_header is a duration value in seconds that sets how long to allow a read from a client's request headers. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"timeout_write": schema.Int64Attribute{
			MarkdownDescription: "write is a duration value in seconds that sets how long to allow a write to a client. Note that setting this to a small value when serving large files may negatively affect legitimately slow clients. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"timeout_idle": schema.Int64Attribute{
			MarkdownDescription: "idle is a duration value in seconds that sets the maximum time to wait for the next request when keep-alives are enabled. The default value helps to avoid resource exhaustion. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"log_credentials": schema.BoolAttribute{
			MarkdownDescription: "Log all cookies and authorization in HTTP request logging. Use this option combined with \"HTTP Access Log\" in a \"Domain\". Enable this option only for troubleshooting.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"log_access_plain": schema.BoolAttribute{
			MarkdownDescription: "Log HTTP access in a standard JSON logfile per domain in addition to the standard Log File. Use this option combined with \"HTTP Access Log\" in a \"Domain\". JSON logs can be found in the filesystem at \"/var/log/caddy/access/\".",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"log_access_plain_keep": schema.Int64Attribute{
			MarkdownDescription: "Specify how many days to keep the JSON access logs. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"log_level": schema.StringAttribute{
			MarkdownDescription: "Select the minimum global Log Level. \"INFO\" is the default and should not be changed without a reason, as it displays the ACME Client messages for automatic certificates. This setting does not influence the HTTP Access logs; they always use \"INFO\", which is their lowest supported Log Level. One of `DEBUG`, `WARN`, `ERROR`, `PANIC`, `FATAL`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"dyn_dns_simple_http": schema.StringAttribute{
			MarkdownDescription: "Enter a URL to test the current IP address of the firewall via the HTTP protocol. This is generally not needed as Caddy uses default providers to test the current IP addresses. If a custom provider is preferred, enter the \"https://\" link to an IP address testing website.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"dyn_dns_interface": schema.StringAttribute{
			MarkdownDescription: "Select an interface to extract the current IP addresses of the firewall. This is generally not needed as Caddy uses default providers to test the current IP addresses. Depending on the specified DynDns IP Version, at most one IPv6 Global Unicast Address and one IPv4 non-RFC1918 Address will be extracted.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"dyn_dns_interval": schema.Int64Attribute{
			MarkdownDescription: "Set the interval in seconds to poll for changes in the IP address. Leave empty to use system defaults. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"dyn_dns_ip_versions": schema.StringAttribute{
			MarkdownDescription: "Select the DynDns IP Version: \"IPv4+IPv6\" to set IPv4 A-Records and IPv6 AAAA-Records, \"IPv4 only\" for only A-Records, \"IPv6 only\" for only AAAA-Records. One of `ipv4`, `ipv6`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"dyn_dns_ttl": schema.Int64Attribute{
			MarkdownDescription: "Set the TTL (Time to Live) for DNS records in seconds. Leave empty to use the default of an already existing TTL (when updating only) or the default of the provider API (when creating new records). If explicitly set, values should be as defined in RFC 2181 Section 8.\u201d. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"dyn_dns_update_only": schema.BoolAttribute{
			MarkdownDescription: "If enabled, no new DNS records will be created. Only existing records will be updated. This means that the A or AAAA records need to be created manually ahead of time.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"auth_provider": schema.StringAttribute{
			MarkdownDescription: "Select a Forward Auth Provider. It can be added inside a \"Handler\" by enabling the \"Forward Auth\" checkbox. One of `authelia`, `authentik`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"auth_to_domain": schema.StringAttribute{
			MarkdownDescription: "Enter the domain name or IP address of the chosen Forward Auth Provider.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"auth_to_port": schema.Int64Attribute{
			MarkdownDescription: "Enter the listen port of the chosen Forward Auth Provider. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"auth_to_tls": schema.StringAttribute{
			MarkdownDescription: "Enable or disable HTTP over TLS (HTTPS) to communicate with the Forward Auth Provider. One of `0`, `1`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"auth_to_uri": schema.StringAttribute{
			MarkdownDescription: "Enter the URI of the authz api endpoint.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"copy_headers": schema.SetAttribute{
			MarkdownDescription: "Select headers to copy in addition to the default of the chosen provider. Headers can be added in \"Reverse Proxy - Headers\". As example, copying \"Authorization\" can pass Basic Auth credentials from the Auth Provider to the reverse proxied application.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `caddy_settings`.",
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

func convertSettingsSchemaToStruct(d *settingsResourceModel) *caddy.Settings {
	return &caddy.Settings{
		General: caddy.SettingsGeneral{
			Enabled:                        settingsBool(d.Enabled),
			EnableLayer4:                   settingsBool(d.EnableLayer4),
			StrictSniHost:                  settingsBool(d.StrictSniHost),
			HttpPort:                       settingsInt(d.HttpPort),
			HttpsPort:                      settingsInt(d.HttpsPort),
			TlsEmail:                       d.TlsEmail.ValueString(),
			TlsAutoHttps:                   api.SelectedMap(d.TlsAutoHttps.ValueString()),
			TlsDnsProvider:                 api.SelectedMap(d.TlsDnsProvider.ValueString()),
			TlsDnsApiKey:                   d.TlsDnsApiKey.ValueString(),
			TlsDnsPropagationTimeout:       settingsBool(d.TlsDnsPropagationTimeout),
			TlsDnsPropagationTimeoutPeriod: settingsInt(d.TlsDnsPropagationTimeoutPeriod),
			TlsDnsPropagationDelay:         settingsInt(d.TlsDnsPropagationDelay),
			TlsDnsPropagationResolvers:     d.TlsDnsPropagationResolvers.ValueString(),
			TlsDnsEchDomain:                d.TlsDnsEchDomain.ValueString(),
			Accesslist:                     api.SelectedMap(d.Accesslist.ValueString()),
			ClientIpHeaders:                api.SelectedMapList(settingsList(d.ClientIpHeaders)),
			ProxyProtocolAllow:             api.SelectedMapList(settingsList(d.ProxyProtocolAllow)),
			DisableSuperuser:               api.SelectedMap(d.DisableSuperuser.ValueString()),
			GracePeriod:                    settingsInt(d.GracePeriod),
			HttpVersions:                   api.SelectedMapList(settingsList(d.HttpVersions)),
			TimeoutReadBody:                settingsInt(d.TimeoutReadBody),
			TimeoutReadHeader:              settingsInt(d.TimeoutReadHeader),
			TimeoutWrite:                   settingsInt(d.TimeoutWrite),
			TimeoutIdle:                    settingsInt(d.TimeoutIdle),
			LogCredentials:                 settingsBool(d.LogCredentials),
			LogAccessPlain:                 settingsBool(d.LogAccessPlain),
			LogAccessPlainKeep:             settingsInt(d.LogAccessPlainKeep),
			LogLevel:                       api.SelectedMap(d.LogLevel.ValueString()),
			DynDnsSimpleHttp:               d.DynDnsSimpleHttp.ValueString(),
			DynDnsInterface:                api.SelectedMap(d.DynDnsInterface.ValueString()),
			DynDnsInterval:                 settingsInt(d.DynDnsInterval),
			DynDnsIpVersions:               api.SelectedMap(d.DynDnsIpVersions.ValueString()),
			DynDnsTtl:                      settingsInt(d.DynDnsTtl),
			DynDnsUpdateOnly:               settingsBool(d.DynDnsUpdateOnly),
			AuthProvider:                   api.SelectedMap(d.AuthProvider.ValueString()),
			AuthToDomain:                   d.AuthToDomain.ValueString(),
			AuthToPort:                     settingsInt(d.AuthToPort),
			AuthToTls:                      api.SelectedMap(d.AuthToTls.ValueString()),
			AuthToUri:                      d.AuthToUri.ValueString(),
			CopyHeaders:                    api.SelectedMapList(settingsList(d.CopyHeaders)),
		},
	}
}

func convertSettingsStructToSchema(s *caddy.Settings) *settingsResourceModel {
	return &settingsResourceModel{
		Enabled:                        types.BoolValue(tools.StringToBool(s.General.Enabled)),
		EnableLayer4:                   types.BoolValue(tools.StringToBool(s.General.EnableLayer4)),
		StrictSniHost:                  types.BoolValue(tools.StringToBool(s.General.StrictSniHost)),
		HttpPort:                       settingsReadInt(s.General.HttpPort),
		HttpsPort:                      settingsReadInt(s.General.HttpsPort),
		TlsEmail:                       types.StringValue(s.General.TlsEmail),
		TlsAutoHttps:                   types.StringValue(s.General.TlsAutoHttps.String()),
		TlsDnsProvider:                 types.StringValue(s.General.TlsDnsProvider.String()),
		TlsDnsApiKey:                   types.StringValue(s.General.TlsDnsApiKey),
		TlsDnsPropagationTimeout:       types.BoolValue(tools.StringToBool(s.General.TlsDnsPropagationTimeout)),
		TlsDnsPropagationTimeoutPeriod: settingsReadInt(s.General.TlsDnsPropagationTimeoutPeriod),
		TlsDnsPropagationDelay:         settingsReadInt(s.General.TlsDnsPropagationDelay),
		TlsDnsPropagationResolvers:     types.StringValue(s.General.TlsDnsPropagationResolvers),
		TlsDnsEchDomain:                types.StringValue(s.General.TlsDnsEchDomain),
		Accesslist:                     types.StringValue(s.General.Accesslist.String()),
		ClientIpHeaders:                tools.StringSliceToSet(settingsNonEmpty(s.General.ClientIpHeaders)),
		ProxyProtocolAllow:             tools.StringSliceToSet(settingsNonEmpty(s.General.ProxyProtocolAllow)),
		DisableSuperuser:               types.StringValue(s.General.DisableSuperuser.String()),
		GracePeriod:                    settingsReadInt(s.General.GracePeriod),
		HttpVersions:                   tools.StringSliceToSet(settingsNonEmpty(s.General.HttpVersions)),
		TimeoutReadBody:                settingsReadInt(s.General.TimeoutReadBody),
		TimeoutReadHeader:              settingsReadInt(s.General.TimeoutReadHeader),
		TimeoutWrite:                   settingsReadInt(s.General.TimeoutWrite),
		TimeoutIdle:                    settingsReadInt(s.General.TimeoutIdle),
		LogCredentials:                 types.BoolValue(tools.StringToBool(s.General.LogCredentials)),
		LogAccessPlain:                 types.BoolValue(tools.StringToBool(s.General.LogAccessPlain)),
		LogAccessPlainKeep:             settingsReadInt(s.General.LogAccessPlainKeep),
		LogLevel:                       types.StringValue(s.General.LogLevel.String()),
		DynDnsSimpleHttp:               types.StringValue(s.General.DynDnsSimpleHttp),
		DynDnsInterface:                types.StringValue(s.General.DynDnsInterface.String()),
		DynDnsInterval:                 settingsReadInt(s.General.DynDnsInterval),
		DynDnsIpVersions:               types.StringValue(s.General.DynDnsIpVersions.String()),
		DynDnsTtl:                      settingsReadInt(s.General.DynDnsTtl),
		DynDnsUpdateOnly:               types.BoolValue(tools.StringToBool(s.General.DynDnsUpdateOnly)),
		AuthProvider:                   types.StringValue(s.General.AuthProvider.String()),
		AuthToDomain:                   types.StringValue(s.General.AuthToDomain),
		AuthToPort:                     settingsReadInt(s.General.AuthToPort),
		AuthToTls:                      types.StringValue(s.General.AuthToTls.String()),
		AuthToUri:                      types.StringValue(s.General.AuthToUri),
		CopyHeaders:                    tools.StringSliceToSet(settingsNonEmpty(s.General.CopyHeaders)),
	}
}
