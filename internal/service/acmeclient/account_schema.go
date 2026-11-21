package acmeclient

import (
	"github.com/browningluke/opnsense-go/pkg/acmeclient"
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type accountResourceModel struct {
	Enabled     types.Bool   `tfsdk:"enabled"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Email       types.String `tfsdk:"email"`
	Ca          types.String `tfsdk:"ca"`
	CustomCa    types.String `tfsdk:"custom_ca"`
	EabKid      types.String `tfsdk:"eab_kid"`
	EabHmac     types.String `tfsdk:"eab_hmac"`

	Id types.String `tfsdk:"id"`
}

func accountResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "ACME accounts (**Services ▸ ACME Client ▸ Accounts**, plugin `os-acme-client`). The account is registered with the CA when the first certificate is issued.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this account. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify this account.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this account. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"email": schema.StringAttribute{
				MarkdownDescription: "Optional e-mail address for this account. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ca": schema.StringAttribute{
				MarkdownDescription: "The ACME CA that should be used for this account and all associated certificates. Note that some of them offer paid services and may require a subscription. Check the acme.sh documentation for a list of supported CAs. One of `buypass`, `buypass_test`, `google`, `google_test`, `letsencrypt`, `letsencrypt_test`, `sslcom`, `zerossl`, `custom`. Defaults to `letsencrypt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("letsencrypt"),
				Validators:          []validator.String{stringvalidator.OneOf("buypass", "buypass_test", "google", "google_test", "letsencrypt", "letsencrypt_test", "sslcom", "zerossl", "custom")},
			},
			"custom_ca": schema.StringAttribute{
				MarkdownDescription: "The HTTPS URL of the custom ACME CA that should be used for this account and all associated certificates. If using your own CA, make sure the Root CA is added to OPNsense's trust store. For example: https://acme-v02.api.letsencrypt.org/directory or https://ca.internal/acme/directory. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"eab_kid": schema.StringAttribute{
				MarkdownDescription: "An optional value provided by the CA when using ACME External Account Binding (EAB). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"eab_hmac": schema.StringAttribute{
				MarkdownDescription: "An optional value provided by the CA when using ACME External Account Binding (EAB). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the ACME account.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func accountDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "ACME accounts (**Services ▸ ACME Client ▸ Accounts**, plugin `os-acme-client`). The account is registered with the CA when the first certificate is issued.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the ACME account.",
				Required:            true,
			},
			"enabled":     dschema.BoolAttribute{MarkdownDescription: "Enable this account.", Computed: true},
			"name":        dschema.StringAttribute{MarkdownDescription: "Name to identify this account.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Description for this account.", Computed: true},
			"email":       dschema.StringAttribute{MarkdownDescription: "Optional e-mail address for this account.", Computed: true},
			"ca":          dschema.StringAttribute{MarkdownDescription: "The ACME CA that should be used for this account and all associated certificates. Note that some of them offer paid services and may require a subscription. Check the acme.sh documentation for a list of supported CAs. One of `buypass`, `buypass_test`, `google`, `google_test`, `letsencrypt`, `letsencrypt_test`, `sslcom`, `zerossl`, `custom`.", Computed: true},
			"custom_ca":   dschema.StringAttribute{MarkdownDescription: "The HTTPS URL of the custom ACME CA that should be used for this account and all associated certificates. If using your own CA, make sure the Root CA is added to OPNsense's trust store. For example: https://acme-v02.api.letsencrypt.org/directory or https://ca.internal/acme/directory.", Computed: true},
			"eab_kid":     dschema.StringAttribute{MarkdownDescription: "An optional value provided by the CA when using ACME External Account Binding (EAB).", Computed: true},
			"eab_hmac":    dschema.StringAttribute{MarkdownDescription: "An optional value provided by the CA when using ACME External Account Binding (EAB).", Computed: true, Sensitive: true},
		},
	}
}

func convertAccountSchemaToStruct(d *accountResourceModel) (*acmeclient.Account, error) {
	return &acmeclient.Account{
		Enabled:     tools.BoolToString(d.Enabled.ValueBool()),
		Name:        d.Name.ValueString(),
		Description: d.Description.ValueString(),
		Email:       d.Email.ValueString(),
		Ca:          api.SelectedMap(d.Ca.ValueString()),
		CustomCa:    d.CustomCa.ValueString(),
		EabKid:      d.EabKid.ValueString(),
		EabHmac:     d.EabHmac.ValueString(),
	}, nil
}

func convertAccountStructToSchema(d *acmeclient.Account) (*accountResourceModel, error) {
	return &accountResourceModel{
		Enabled:     types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:        types.StringValue(d.Name),
		Description: types.StringValue(d.Description),
		Email:       types.StringValue(d.Email),
		Ca:          types.StringValue(d.Ca.String()),
		CustomCa:    types.StringValue(d.CustomCa),
		EabKid:      types.StringValue(d.EabKid),
		EabHmac:     types.StringValue(d.EabHmac),
	}, nil
}
