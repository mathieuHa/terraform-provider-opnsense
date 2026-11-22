package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type tlsFingerprintResourceModel struct {
	Description types.String `tfsdk:"description"`
	UserAgent   types.String `tfsdk:"user_agent"`
	Trusted     types.Bool   `tfsdk:"trusted"`
	Curves      types.String `tfsdk:"curves"`
	Ciphers     types.String `tfsdk:"ciphers"`

	Id types.String `tfsdk:"id"`
}

func tlsFingerprintResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx TLS fingerprints (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ TLS Fingerprint**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				MarkdownDescription: "Enter a useful description for your entry.",
				Required:            true,
			},
			"user_agent": schema.StringAttribute{
				MarkdownDescription: "User Agent.",
				Required:            true,
			},
			"trusted": schema.BoolAttribute{
				MarkdownDescription: "If you check this box, this entry will be added to the configuration of trusted fingerprints. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"curves": schema.StringAttribute{
				MarkdownDescription: "Enter a list of elliptic curves supported by the client in a standard configuration in OpenSSL syntax (curve1:curve2:0x0001). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ciphers": schema.StringAttribute{
				MarkdownDescription: "Enter a list of elliptic curves supported by the client in a standard configuration in OpenSSL syntax (cipher1:cipher2:0x0001).",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx TLS fingerprint.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func tlsFingerprintDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx TLS fingerprints (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ TLS Fingerprint**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx TLS fingerprint.",
				Required:            true,
			},
			"description": dschema.StringAttribute{MarkdownDescription: "Enter a useful description for your entry.", Computed: true},
			"user_agent":  dschema.StringAttribute{MarkdownDescription: "User Agent.", Computed: true},
			"trusted":     dschema.BoolAttribute{MarkdownDescription: "If you check this box, this entry will be added to the configuration of trusted fingerprints.", Computed: true},
			"curves":      dschema.StringAttribute{MarkdownDescription: "Enter a list of elliptic curves supported by the client in a standard configuration in OpenSSL syntax (curve1:curve2:0x0001).", Computed: true},
			"ciphers":     dschema.StringAttribute{MarkdownDescription: "Enter a list of elliptic curves supported by the client in a standard configuration in OpenSSL syntax (cipher1:cipher2:0x0001).", Computed: true},
		},
	}
}

func convertTlsFingerprintSchemaToStruct(d *tlsFingerprintResourceModel) (*nginx.TlsFingerprint, error) {
	return &nginx.TlsFingerprint{
		Description: d.Description.ValueString(),
		UserAgent:   d.UserAgent.ValueString(),
		Trusted:     tools.BoolToString(d.Trusted.ValueBool()),
		Curves:      d.Curves.ValueString(),
		Ciphers:     d.Ciphers.ValueString(),
	}, nil
}

func convertTlsFingerprintStructToSchema(d *nginx.TlsFingerprint) (*tlsFingerprintResourceModel, error) {
	return &tlsFingerprintResourceModel{
		Description: types.StringValue(d.Description),
		UserAgent:   types.StringValue(d.UserAgent),
		Trusted:     types.BoolValue(tools.StringToBool(d.Trusted)),
		Curves:      types.StringValue(d.Curves),
		Ciphers:     types.StringValue(d.Ciphers),
	}, nil
}
