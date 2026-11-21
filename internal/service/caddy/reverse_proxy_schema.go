package caddy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/caddy"
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

type reverseProxyResourceModel struct {
	Enabled                    types.Bool   `tfsdk:"enabled"`
	FromDomain                 types.String `tfsdk:"from_domain"`
	FromPort                   types.Int64  `tfsdk:"from_port"`
	Accesslist                 types.String `tfsdk:"accesslist"`
	Basicauth                  types.Set    `tfsdk:"basicauth"`
	Description                types.String `tfsdk:"description"`
	DNSChallenge               types.Bool   `tfsdk:"dns_challenge"`
	DNSChallengeOverrideDomain types.String `tfsdk:"dns_challenge_override_domain"`
	CustomCertificate          types.String `tfsdk:"custom_certificate"`
	AccessLog                  types.Bool   `tfsdk:"access_log"`
	DynDNS                     types.Bool   `tfsdk:"dyn_dns"`
	AcmePassthrough            types.String `tfsdk:"acme_passthrough"`
	DisableTLS                 types.String `tfsdk:"disable_tls"`
	ClientAuthMode             types.String `tfsdk:"client_auth_mode"`
	ClientAuthTrustPool        types.Set    `tfsdk:"client_auth_trust_pool"`

	Id types.String `tfsdk:"id"`
}

func reverseProxyResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Caddy reverse proxy domains (**Services ▸ Caddy Web Server ▸ Reverse Proxy ▸ Domains**, plugin `os-caddy`). Handlers (`opnsense_caddy_handle`) route the requests of a domain to upstreams.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this domain. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"from_domain": schema.StringAttribute{
				MarkdownDescription: "Enter a domain name. For a base domain, use \"example.com\" or \"opn.example.com\". For a wildcard domain, use \"*.example.com\". Using a wildcard domain with subdomains requires a \"DNS Provider\" and the \"DNS-01 Challenge\" or a custom certificate.",
				Required:            true,
			},
			"from_port": schema.Int64Attribute{
				MarkdownDescription: "Leave empty to use ports 80 and 443 with automatic redirection from HTTP to HTTPS or choose a custom port. Don't forget to allow these ports with a Firewall rule. If the default ports have been changed in \"General Settings\", leaving this empty will use the chosen alternative ports instead. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"accesslist": schema.StringAttribute{
				MarkdownDescription: "Select an Access List to restrict access to this domain. If unset, any local or remote client is allowed access. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"basicauth": schema.SetAttribute{
				MarkdownDescription: "Select Users to restrict access to this domain. Basic Auth matches after Access Lists. If unset, any user is allowed access. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Enter a description for this domain. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_challenge": schema.BoolAttribute{
				MarkdownDescription: "Enable the DNS-01 Challenge for this domain. Requires a \"DNS Provider\" in \"General Settings\". This is only needed for wildcard domains, or when the default challenges can not be used due to restrictive firewall policies. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"dns_challenge_override_domain": schema.StringAttribute{
				MarkdownDescription: "If using CNAME delegation for DNS-01 challenges, enter the target domain here. The DNS provider will manage TXT records at '_acme-challenge.targetdomain' instead of the certificate domain. Your configured DNS Provider must have write access to this target domain. For wildcard certificates, ensure the CNAME delegation covers the _acme-challenge label. This enables least-privilege setups or can be used as a workaround when the primary DNS provider lacks an API. Leave empty if your DNS Provider has a dedicated field for this. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"custom_certificate": schema.StringAttribute{
				MarkdownDescription: "Choose \"Auto HTTPS\" to get automatic \"Let's Encrypt\" or \"ZeroSSL\" certificates; no additional plugin required. Alternatively, choose a custom certificate from \"System - Trust - Certificates\" for this domain, e.g., certificates managed by the optional os-acme-client plugin. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"access_log": schema.BoolAttribute{
				MarkdownDescription: "Enable HTTP request logging for this domain and its subdomains. This option is mostly used for troubleshooting. It will log every single request. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"dyn_dns": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable Dynamic DNS. Requires a \"DNS Provider\" in \"General Settings\". The DNS records of this domain will be automatically updated. A wildcard domain will create a \"*.example.com\" record. A base domain will create a \"example.com\" or \"opn.example.com\" record. Some providers need subdomains to set records for domains like \"opn.example.com\"; in that case use the checkbox in the subdomains tab. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"acme_passthrough": schema.StringAttribute{
				MarkdownDescription: "Enter a domain name or IP address. The HTTP-01 challenge will be redirected to that destination. This enables an ACME Client behind Caddy to serve \"/.well-known/acme-challenge/\" on port 80. Caddy will reverse proxy the HTTP-01 challenge for this domain, and will still issue a certificate using the TLS-ALPN-01 challenge or DNS-01 challenge for itself. This option can be used for High Availability when using Caddy with a master and backup OPNsense. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"disable_tls": schema.StringAttribute{
				MarkdownDescription: "When choosing HTTP, automatic certificate management will be disabled and all traffic to and from this domain will be unencrypted. One of `0`, `1`. Defaults to `0`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("0"),
				Validators:          []validator.String{stringvalidator.OneOf("0", "1")},
			},
			"client_auth_mode": schema.StringAttribute{
				MarkdownDescription: "\"request\" - Ask clients for a certificate, but allow even if there isn't one; do not verify it. \"require\" - Require clients to present a certificate, but do not verify it. \"verify_if_given\" - Ask clients for a certificate; allow even if there isn't one, but verify it if there is. \"require_and_verify\" - Require clients to present a valid certificate that is verified. One of `request`, `require`, `verify_if_given`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "request", "require", "verify_if_given")},
			},
			"client_auth_trust_pool": schema.SetAttribute{
				MarkdownDescription: "Choose multiple CAs or self-signed certificates from \"System - Trust - Authorities\". Client Auth is activated as soon as at least one certificate has been chosen. Important: Certificate revocation lists are not evaluated. If you need granular control, provide individual self-signed certificates for each device, and unset them to block access. Though keep in mind that if no certificate is left in this field, Client Auth will be deactivated. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the Caddy domain.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func reverseProxyDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Caddy reverse proxy domains (**Services ▸ Caddy Web Server ▸ Reverse Proxy ▸ Domains**, plugin `os-caddy`). Handlers (`opnsense_caddy_handle`) route the requests of a domain to upstreams.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the Caddy domain.",
				Required:            true,
			},
			"enabled":                       dschema.BoolAttribute{MarkdownDescription: "Enable this domain.", Computed: true},
			"from_domain":                   dschema.StringAttribute{MarkdownDescription: "Enter a domain name. For a base domain, use \"example.com\" or \"opn.example.com\". For a wildcard domain, use \"*.example.com\". Using a wildcard domain with subdomains requires a \"DNS Provider\" and the \"DNS-01 Challenge\" or a custom certificate.", Computed: true},
			"from_port":                     dschema.Int64Attribute{MarkdownDescription: "Leave empty to use ports 80 and 443 with automatic redirection from HTTP to HTTPS or choose a custom port. Don't forget to allow these ports with a Firewall rule. If the default ports have been changed in \"General Settings\", leaving this empty will use the chosen alternative ports instead.", Computed: true},
			"accesslist":                    dschema.StringAttribute{MarkdownDescription: "Select an Access List to restrict access to this domain. If unset, any local or remote client is allowed access.", Computed: true},
			"basicauth":                     dschema.SetAttribute{MarkdownDescription: "Select Users to restrict access to this domain. Basic Auth matches after Access Lists. If unset, any user is allowed access.", Computed: true, ElementType: types.StringType},
			"description":                   dschema.StringAttribute{MarkdownDescription: "Enter a description for this domain.", Computed: true},
			"dns_challenge":                 dschema.BoolAttribute{MarkdownDescription: "Enable the DNS-01 Challenge for this domain. Requires a \"DNS Provider\" in \"General Settings\". This is only needed for wildcard domains, or when the default challenges can not be used due to restrictive firewall policies.", Computed: true},
			"dns_challenge_override_domain": dschema.StringAttribute{MarkdownDescription: "If using CNAME delegation for DNS-01 challenges, enter the target domain here. The DNS provider will manage TXT records at '_acme-challenge.targetdomain' instead of the certificate domain. Your configured DNS Provider must have write access to this target domain. For wildcard certificates, ensure the CNAME delegation covers the _acme-challenge label. This enables least-privilege setups or can be used as a workaround when the primary DNS provider lacks an API. Leave empty if your DNS Provider has a dedicated field for this.", Computed: true},
			"custom_certificate":            dschema.StringAttribute{MarkdownDescription: "Choose \"Auto HTTPS\" to get automatic \"Let's Encrypt\" or \"ZeroSSL\" certificates; no additional plugin required. Alternatively, choose a custom certificate from \"System - Trust - Certificates\" for this domain, e.g., certificates managed by the optional os-acme-client plugin.", Computed: true},
			"access_log":                    dschema.BoolAttribute{MarkdownDescription: "Enable HTTP request logging for this domain and its subdomains. This option is mostly used for troubleshooting. It will log every single request.", Computed: true},
			"dyn_dns":                       dschema.BoolAttribute{MarkdownDescription: "Enable or disable Dynamic DNS. Requires a \"DNS Provider\" in \"General Settings\". The DNS records of this domain will be automatically updated. A wildcard domain will create a \"*.example.com\" record. A base domain will create a \"example.com\" or \"opn.example.com\" record. Some providers need subdomains to set records for domains like \"opn.example.com\"; in that case use the checkbox in the subdomains tab.", Computed: true},
			"acme_passthrough":              dschema.StringAttribute{MarkdownDescription: "Enter a domain name or IP address. The HTTP-01 challenge will be redirected to that destination. This enables an ACME Client behind Caddy to serve \"/.well-known/acme-challenge/\" on port 80. Caddy will reverse proxy the HTTP-01 challenge for this domain, and will still issue a certificate using the TLS-ALPN-01 challenge or DNS-01 challenge for itself. This option can be used for High Availability when using Caddy with a master and backup OPNsense.", Computed: true},
			"disable_tls":                   dschema.StringAttribute{MarkdownDescription: "When choosing HTTP, automatic certificate management will be disabled and all traffic to and from this domain will be unencrypted. One of `0`, `1`.", Computed: true},
			"client_auth_mode":              dschema.StringAttribute{MarkdownDescription: "\"request\" - Ask clients for a certificate, but allow even if there isn't one; do not verify it. \"require\" - Require clients to present a certificate, but do not verify it. \"verify_if_given\" - Ask clients for a certificate; allow even if there isn't one, but verify it if there is. \"require_and_verify\" - Require clients to present a valid certificate that is verified. One of `request`, `require`, `verify_if_given`.", Computed: true},
			"client_auth_trust_pool":        dschema.SetAttribute{MarkdownDescription: "Choose multiple CAs or self-signed certificates from \"System - Trust - Authorities\". Client Auth is activated as soon as at least one certificate has been chosen. Important: Certificate revocation lists are not evaluated. If you need granular control, provide individual self-signed certificates for each device, and unset them to block access. Though keep in mind that if no certificate is left in this field, Client Auth will be deactivated.", Computed: true, ElementType: types.StringType},
		},
	}
}

func convertReverseProxySchemaToStruct(d *reverseProxyResourceModel) (*caddy.ReverseProxy, error) {
	return &caddy.ReverseProxy{
		Enabled:                    tools.BoolToString(d.Enabled.ValueBool()),
		FromDomain:                 d.FromDomain.ValueString(),
		FromPort:                   tools.Int64ToStringNegative(d.FromPort.ValueInt64()),
		Accesslist:                 api.SelectedMap(d.Accesslist.ValueString()),
		Basicauth:                  api.SelectedMapList(setToSlice(d.Basicauth)),
		Description:                d.Description.ValueString(),
		DnsChallenge:               tools.BoolToString(d.DNSChallenge.ValueBool()),
		DnsChallengeOverrideDomain: d.DNSChallengeOverrideDomain.ValueString(),
		CustomCertificate:          api.SelectedMap(d.CustomCertificate.ValueString()),
		AccessLog:                  tools.BoolToString(d.AccessLog.ValueBool()),
		DynDns:                     tools.BoolToString(d.DynDNS.ValueBool()),
		AcmePassthrough:            d.AcmePassthrough.ValueString(),
		DisableTls:                 api.SelectedMap(d.DisableTLS.ValueString()),
		ClientAuthMode:             api.SelectedMap(d.ClientAuthMode.ValueString()),
		ClientAuthTrustPool:        api.SelectedMapList(setToSlice(d.ClientAuthTrustPool)),
	}, nil
}

func convertReverseProxyStructToSchema(d *caddy.ReverseProxy) (*reverseProxyResourceModel, error) {
	return &reverseProxyResourceModel{
		Enabled:                    types.BoolValue(tools.StringToBool(d.Enabled)),
		FromDomain:                 types.StringValue(d.FromDomain),
		FromPort:                   types.Int64Value(tools.StringToInt64(d.FromPort)),
		Accesslist:                 types.StringValue(d.Accesslist.String()),
		Basicauth:                  tools.StringSliceToSet(d.Basicauth),
		Description:                types.StringValue(d.Description),
		DNSChallenge:               types.BoolValue(tools.StringToBool(d.DnsChallenge)),
		DNSChallengeOverrideDomain: types.StringValue(d.DnsChallengeOverrideDomain),
		CustomCertificate:          types.StringValue(d.CustomCertificate.String()),
		AccessLog:                  types.BoolValue(tools.StringToBool(d.AccessLog)),
		DynDNS:                     types.BoolValue(tools.StringToBool(d.DynDns)),
		AcmePassthrough:            types.StringValue(d.AcmePassthrough),
		DisableTLS:                 types.StringValue(d.DisableTls.String()),
		ClientAuthMode:             types.StringValue(d.ClientAuthMode.String()),
		ClientAuthTrustPool:        tools.StringSliceToSet(d.ClientAuthTrustPool),
	}, nil
}

func setToSlice(s types.Set) []string {
	out := tools.SetToStringSlice(s)
	if out == nil {
		return []string{}
	}
	return out
}
