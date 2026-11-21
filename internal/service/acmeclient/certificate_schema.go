package acmeclient

import (
	"github.com/browningluke/opnsense-go/pkg/acmeclient"
	"github.com/browningluke/opnsense-go/pkg/api"
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

type certificateResourceModel struct {
	Enabled          types.Bool   `tfsdk:"enabled"`
	Name             types.String `tfsdk:"name"`
	Description      types.String `tfsdk:"description"`
	AltNames         types.Set    `tfsdk:"alt_names"`
	Account          types.String `tfsdk:"account"`
	ValidationMethod types.String `tfsdk:"validation_method"`
	KeyLength        types.String `tfsdk:"key_length"`
	Ocsp             types.Bool   `tfsdk:"ocsp"`
	Profile          types.String `tfsdk:"profile"`
	RestartActions   types.Set    `tfsdk:"restart_actions"`
	AutoRenewal      types.Bool   `tfsdk:"auto_renewal"`
	RenewInterval    types.Int64  `tfsdk:"renew_interval"`
	Aliasmode        types.String `tfsdk:"aliasmode"`
	Domainalias      types.String `tfsdk:"domainalias"`
	Challengealias   types.String `tfsdk:"challengealias"`

	Id types.String `tfsdk:"id"`
}

func certificateResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "ACME certificates (**Services ▸ ACME Client ▸ Certificates**, plugin `os-acme-client`). OPNsense issues and renews them with its cron job; issued certificates appear in **System ▸ Trust ▸ Certificates**.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this certificate. When disabled, no attemps to issue or renew the certificate will be made. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Common Name (CN) and first Alt Name (subjectAltName) for this certificate.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this certificate. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"alt_names": schema.SetAttribute{
				MarkdownDescription: "Configure additional names that should be part of the certificate, i.e. www.example.com or mail.example.com. Use TAB key to complete typing a FQDN.NOTE:You need to forcefully re-issue the certificate if you change \"Alt Names\" after the certificate was signed by the ACME CA! Use the \"issue\" button in the Commands column in this case. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"account": schema.StringAttribute{
				MarkdownDescription: "Set the ACME CA account to use for this certificate.",
				Required:            true,
			},
			"validation_method": schema.StringAttribute{
				MarkdownDescription: "Set the ACME challenge type for this certificate.",
				Required:            true,
			},
			"key_length": schema.StringAttribute{
				MarkdownDescription: "Specify the domain key length: 2048, 3072, 4096, 8192 or ec-256, ec-384. One of `key_2048`, `key_3072`, `key_4096`, `key_ec256`, `key_ec384`. Defaults to `key_4096`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("key_4096"),
				Validators:          []validator.String{stringvalidator.OneOf("key_2048", "key_3072", "key_4096", "key_ec256", "key_ec384")},
			},
			"ocsp": schema.BoolAttribute{
				MarkdownDescription: "Generate and add OCSP Must Staple extension to the certificate. When this option is enabled and issueance/renewal requests fail, then this extension is probably not supported by the CA. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"profile": schema.StringAttribute{
				MarkdownDescription: "Optional. Specify the ACME Certificate Profile to use (e.g. \"shortlived\"). See CA documentation for available profiles. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"restart_actions": schema.SetAttribute{
				MarkdownDescription: "Choose the automations that should be run after certificate creation and renewal. Basically every application requires a quick restart to reload the updated certificate. If you don't configure an automation, the in-memory certificate may expire and cause security warnings and other issues. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"auto_renewal": schema.BoolAttribute{
				MarkdownDescription: "Enable automatic renewal for this certificate to prevent expiration. When disabled, the cron job will ignore this certificate. Note that it is still possible to renew the certificate from the GUI. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"renew_interval": schema.Int64Attribute{
				MarkdownDescription: "Specifies the days to renew the cert. The max value is 5000 days. Defaults to `60`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(60),
			},
			"aliasmode": schema.StringAttribute{
				MarkdownDescription: "Configure DNS alias mode to validate the certificate. One of `none`, `automatic`, `domain`, `challenge`. Defaults to `none`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("none"),
				Validators:          []validator.String{stringvalidator.OneOf("none", "automatic", "domain", "challenge")},
			},
			"domainalias": schema.StringAttribute{
				MarkdownDescription: "When setting DNS alias mode to \"Domain Alias\", enter the domain name that should be used for certificate validation. Please refer to the acme.sh documentation for further information. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"challengealias": schema.StringAttribute{
				MarkdownDescription: "When setting DNS alias mode to \"Challenge Alias\", enter the domain name that should be used for certificate validation. Please refer to the acme.sh documentation for further information. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the ACME certificate.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func certificateDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "ACME certificates (**Services ▸ ACME Client ▸ Certificates**, plugin `os-acme-client`). OPNsense issues and renews them with its cron job; issued certificates appear in **System ▸ Trust ▸ Certificates**.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the ACME certificate.",
				Required:            true,
			},
			"enabled":           dschema.BoolAttribute{MarkdownDescription: "Enable this certificate. When disabled, no attemps to issue or renew the certificate will be made.", Computed: true},
			"name":              dschema.StringAttribute{MarkdownDescription: "Common Name (CN) and first Alt Name (subjectAltName) for this certificate.", Computed: true},
			"description":       dschema.StringAttribute{MarkdownDescription: "Description for this certificate.", Computed: true},
			"alt_names":         dschema.SetAttribute{MarkdownDescription: "Configure additional names that should be part of the certificate, i.e. www.example.com or mail.example.com. Use TAB key to complete typing a FQDN.NOTE:You need to forcefully re-issue the certificate if you change \"Alt Names\" after the certificate was signed by the ACME CA! Use the \"issue\" button in the Commands column in this case.", Computed: true, ElementType: types.StringType},
			"account":           dschema.StringAttribute{MarkdownDescription: "Set the ACME CA account to use for this certificate.", Computed: true},
			"validation_method": dschema.StringAttribute{MarkdownDescription: "Set the ACME challenge type for this certificate.", Computed: true},
			"key_length":        dschema.StringAttribute{MarkdownDescription: "Specify the domain key length: 2048, 3072, 4096, 8192 or ec-256, ec-384. One of `key_2048`, `key_3072`, `key_4096`, `key_ec256`, `key_ec384`.", Computed: true},
			"ocsp":              dschema.BoolAttribute{MarkdownDescription: "Generate and add OCSP Must Staple extension to the certificate. When this option is enabled and issueance/renewal requests fail, then this extension is probably not supported by the CA.", Computed: true},
			"profile":           dschema.StringAttribute{MarkdownDescription: "Optional. Specify the ACME Certificate Profile to use (e.g. \"shortlived\"). See CA documentation for available profiles.", Computed: true},
			"restart_actions":   dschema.SetAttribute{MarkdownDescription: "Choose the automations that should be run after certificate creation and renewal. Basically every application requires a quick restart to reload the updated certificate. If you don't configure an automation, the in-memory certificate may expire and cause security warnings and other issues.", Computed: true, ElementType: types.StringType},
			"auto_renewal":      dschema.BoolAttribute{MarkdownDescription: "Enable automatic renewal for this certificate to prevent expiration. When disabled, the cron job will ignore this certificate. Note that it is still possible to renew the certificate from the GUI.", Computed: true},
			"renew_interval":    dschema.Int64Attribute{MarkdownDescription: "Specifies the days to renew the cert. The max value is 5000 days.", Computed: true},
			"aliasmode":         dschema.StringAttribute{MarkdownDescription: "Configure DNS alias mode to validate the certificate. One of `none`, `automatic`, `domain`, `challenge`.", Computed: true},
			"domainalias":       dschema.StringAttribute{MarkdownDescription: "When setting DNS alias mode to \"Domain Alias\", enter the domain name that should be used for certificate validation. Please refer to the acme.sh documentation for further information.", Computed: true},
			"challengealias":    dschema.StringAttribute{MarkdownDescription: "When setting DNS alias mode to \"Challenge Alias\", enter the domain name that should be used for certificate validation. Please refer to the acme.sh documentation for further information.", Computed: true},
		},
	}
}

func convertCertificateSchemaToStruct(d *certificateResourceModel) (*acmeclient.Certificate, error) {
	return &acmeclient.Certificate{
		Enabled:          tools.BoolToString(d.Enabled.ValueBool()),
		Name:             d.Name.ValueString(),
		Description:      d.Description.ValueString(),
		AltNames:         api.SelectedMapList(setToSlice(d.AltNames)),
		Account:          api.SelectedMap(d.Account.ValueString()),
		ValidationMethod: api.SelectedMap(d.ValidationMethod.ValueString()),
		KeyLength:        api.SelectedMap(d.KeyLength.ValueString()),
		Ocsp:             tools.BoolToString(d.Ocsp.ValueBool()),
		Profile:          d.Profile.ValueString(),
		RestartActions:   api.SelectedMapList(setToSlice(d.RestartActions)),
		AutoRenewal:      tools.BoolToString(d.AutoRenewal.ValueBool()),
		RenewInterval:    tools.Int64ToString(d.RenewInterval.ValueInt64()),
		Aliasmode:        api.SelectedMap(d.Aliasmode.ValueString()),
		Domainalias:      d.Domainalias.ValueString(),
		Challengealias:   d.Challengealias.ValueString(),
	}, nil
}

func convertCertificateStructToSchema(d *acmeclient.Certificate) (*certificateResourceModel, error) {
	return &certificateResourceModel{
		Enabled:          types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:             types.StringValue(d.Name),
		Description:      types.StringValue(d.Description),
		AltNames:         tools.StringSliceToSet(d.AltNames),
		Account:          types.StringValue(d.Account.String()),
		ValidationMethod: types.StringValue(d.ValidationMethod.String()),
		KeyLength:        types.StringValue(d.KeyLength.String()),
		Ocsp:             types.BoolValue(tools.StringToBool(d.Ocsp)),
		Profile:          types.StringValue(d.Profile),
		RestartActions:   tools.StringSliceToSet(d.RestartActions),
		AutoRenewal:      types.BoolValue(tools.StringToBool(d.AutoRenewal)),
		RenewInterval:    types.Int64Value(tools.StringToInt64(d.RenewInterval)),
		Aliasmode:        types.StringValue(d.Aliasmode.String()),
		Domainalias:      types.StringValue(d.Domainalias),
		Challengealias:   types.StringValue(d.Challengealias),
	}, nil
}

func setToSlice(s types.Set) []string {
	out := tools.SetToStringSlice(s)
	if out == nil {
		return []string{}
	}
	return out
}
