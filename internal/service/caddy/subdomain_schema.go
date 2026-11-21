package caddy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/caddy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type subdomainResourceModel struct {
	Enabled             types.Bool   `tfsdk:"enabled"`
	Reverse             types.String `tfsdk:"reverse"`
	FromDomain          types.String `tfsdk:"from_domain"`
	Accesslist          types.String `tfsdk:"accesslist"`
	Basicauth           types.Set    `tfsdk:"basicauth"`
	Description         types.String `tfsdk:"description"`
	DynDNS              types.Bool   `tfsdk:"dyn_dns"`
	AcmePassthrough     types.String `tfsdk:"acme_passthrough"`
	ClientAuthMode      types.String `tfsdk:"client_auth_mode"`
	ClientAuthTrustPool types.Set    `tfsdk:"client_auth_trust_pool"`

	Id types.String `tfsdk:"id"`
}

func subdomainResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Subdomains of a wildcard Caddy domain (**Services ▸ Caddy Web Server ▸ Reverse Proxy ▸ Subdomains**, plugin `os-caddy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this subdomain. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"reverse": schema.StringAttribute{
				MarkdownDescription: "Select a wildcard domain for this subdomain.",
				Required:            true,
			},
			"from_domain": schema.StringAttribute{
				MarkdownDescription: "Enter a subdomain. For example, \"opn.example.com\" if the wildcard domain is \"*.example.com\". All subdomains use the same ports and protocols as their parent wildcard domain.",
				Required:            true,
			},
			"accesslist": schema.StringAttribute{
				MarkdownDescription: "Select an Access List to restrict access to this subdomain. If unset, any local or remote client is allowed access. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"basicauth": schema.SetAttribute{
				MarkdownDescription: "Select Users to restrict access to this subdomain. Basic Auth matches after Access Lists. If unset, any user is allowed access. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Enter a description for this subdomain. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dyn_dns": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable Dynamic DNS. Requires a \"DNS Provider\" in \"General Settings\". The DNS records of this subdomain will be automatically updated. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"acme_passthrough": schema.StringAttribute{
				MarkdownDescription: "Enter a domain name or IP address. The HTTP-01 challenge will be redirected to that destination. This enables an ACME Client behind Caddy to serve \"/.well-known/acme-challenge/\" on port 80. Caddy will reverse proxy the HTTP-01 challenge for this subdomain. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
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
				MarkdownDescription: "UUID of the Caddy subdomain.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func subdomainDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Subdomains of a wildcard Caddy domain (**Services ▸ Caddy Web Server ▸ Reverse Proxy ▸ Subdomains**, plugin `os-caddy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the Caddy subdomain.",
				Required:            true,
			},
			"enabled":                dschema.BoolAttribute{MarkdownDescription: "Enable this subdomain.", Computed: true},
			"reverse":                dschema.StringAttribute{MarkdownDescription: "Select a wildcard domain for this subdomain.", Computed: true},
			"from_domain":            dschema.StringAttribute{MarkdownDescription: "Enter a subdomain. For example, \"opn.example.com\" if the wildcard domain is \"*.example.com\". All subdomains use the same ports and protocols as their parent wildcard domain.", Computed: true},
			"accesslist":             dschema.StringAttribute{MarkdownDescription: "Select an Access List to restrict access to this subdomain. If unset, any local or remote client is allowed access.", Computed: true},
			"basicauth":              dschema.SetAttribute{MarkdownDescription: "Select Users to restrict access to this subdomain. Basic Auth matches after Access Lists. If unset, any user is allowed access.", Computed: true, ElementType: types.StringType},
			"description":            dschema.StringAttribute{MarkdownDescription: "Enter a description for this subdomain.", Computed: true},
			"dyn_dns":                dschema.BoolAttribute{MarkdownDescription: "Enable or disable Dynamic DNS. Requires a \"DNS Provider\" in \"General Settings\". The DNS records of this subdomain will be automatically updated.", Computed: true},
			"acme_passthrough":       dschema.StringAttribute{MarkdownDescription: "Enter a domain name or IP address. The HTTP-01 challenge will be redirected to that destination. This enables an ACME Client behind Caddy to serve \"/.well-known/acme-challenge/\" on port 80. Caddy will reverse proxy the HTTP-01 challenge for this subdomain.", Computed: true},
			"client_auth_mode":       dschema.StringAttribute{MarkdownDescription: "\"request\" - Ask clients for a certificate, but allow even if there isn't one; do not verify it. \"require\" - Require clients to present a certificate, but do not verify it. \"verify_if_given\" - Ask clients for a certificate; allow even if there isn't one, but verify it if there is. \"require_and_verify\" - Require clients to present a valid certificate that is verified. One of `request`, `require`, `verify_if_given`.", Computed: true},
			"client_auth_trust_pool": dschema.SetAttribute{MarkdownDescription: "Choose multiple CAs or self-signed certificates from \"System - Trust - Authorities\". Client Auth is activated as soon as at least one certificate has been chosen. Important: Certificate revocation lists are not evaluated. If you need granular control, provide individual self-signed certificates for each device, and unset them to block access. Though keep in mind that if no certificate is left in this field, Client Auth will be deactivated.", Computed: true, ElementType: types.StringType},
		},
	}
}

func convertSubdomainSchemaToStruct(d *subdomainResourceModel) (*caddy.Subdomain, error) {
	return &caddy.Subdomain{
		Enabled:             tools.BoolToString(d.Enabled.ValueBool()),
		Reverse:             api.SelectedMap(d.Reverse.ValueString()),
		FromDomain:          d.FromDomain.ValueString(),
		Accesslist:          api.SelectedMap(d.Accesslist.ValueString()),
		Basicauth:           api.SelectedMapList(setToSlice(d.Basicauth)),
		Description:         d.Description.ValueString(),
		DynDns:              tools.BoolToString(d.DynDNS.ValueBool()),
		AcmePassthrough:     d.AcmePassthrough.ValueString(),
		ClientAuthMode:      api.SelectedMap(d.ClientAuthMode.ValueString()),
		ClientAuthTrustPool: api.SelectedMapList(setToSlice(d.ClientAuthTrustPool)),
	}, nil
}

func convertSubdomainStructToSchema(d *caddy.Subdomain) (*subdomainResourceModel, error) {
	return &subdomainResourceModel{
		Enabled:             types.BoolValue(tools.StringToBool(d.Enabled)),
		Reverse:             types.StringValue(d.Reverse.String()),
		FromDomain:          types.StringValue(d.FromDomain),
		Accesslist:          types.StringValue(d.Accesslist.String()),
		Basicauth:           tools.StringSliceToSet(d.Basicauth),
		Description:         types.StringValue(d.Description),
		DynDNS:              types.BoolValue(tools.StringToBool(d.DynDns)),
		AcmePassthrough:     types.StringValue(d.AcmePassthrough),
		ClientAuthMode:      types.StringValue(d.ClientAuthMode.String()),
		ClientAuthTrustPool: tools.StringSliceToSet(d.ClientAuthTrustPool),
	}, nil
}
