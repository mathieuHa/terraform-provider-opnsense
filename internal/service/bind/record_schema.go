package bind

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/bind"
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

type recordResourceModel struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	Domain  types.String `tfsdk:"domain"`
	Name    types.String `tfsdk:"name"`
	Type    types.String `tfsdk:"type"`
	Value   types.String `tfsdk:"value"`

	Id types.String `tfsdk:"id"`
}

func recordResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Records of BIND primary zones (**Services ▸ BIND ▸ Configuration ▸ Records**, plugin `os-bind`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "This will enable or disable this record. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Select the zone for this record.",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Set the name for this record. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Set the type for this record. One of `A`, `AAAA`, `CAA`, `CNAME`, `DNAME`, `DNSKEY`, `DS`, `HTTPS`, `MX`, `NAPTR`, `NS`, `PTR`, `RP`, `RRSIG`, `SRV`, `SSHFP`, `SVCB`, `TLSA`, `TXT`. Defaults to `A`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("A"),
				Validators:          []validator.String{stringvalidator.OneOf("A", "AAAA", "CAA", "CNAME", "DNAME", "DNSKEY", "DS", "HTTPS", "MX", "NAPTR", "NS", "PTR", "RP", "RRSIG", "SRV", "SSHFP", "SVCB", "TLSA", "TXT")},
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "Set the value for this record.",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the BIND record.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func recordDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Records of BIND primary zones (**Services ▸ BIND ▸ Configuration ▸ Records**, plugin `os-bind`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the BIND record.",
				Required:            true,
			},
			"enabled": dschema.BoolAttribute{MarkdownDescription: "This will enable or disable this record.", Computed: true},
			"domain":  dschema.StringAttribute{MarkdownDescription: "Select the zone for this record.", Computed: true},
			"name":    dschema.StringAttribute{MarkdownDescription: "Set the name for this record.", Computed: true},
			"type":    dschema.StringAttribute{MarkdownDescription: "Set the type for this record. One of `A`, `AAAA`, `CAA`, `CNAME`, `DNAME`, `DNSKEY`, `DS`, `HTTPS`, `MX`, `NAPTR`, `NS`, `PTR`, `RP`, `RRSIG`, `SRV`, `SSHFP`, `SVCB`, `TLSA`, `TXT`.", Computed: true},
			"value":   dschema.StringAttribute{MarkdownDescription: "Set the value for this record.", Computed: true},
		},
	}
}

func convertRecordSchemaToStruct(d *recordResourceModel) (*bind.Record, error) {
	return &bind.Record{
		Enabled: tools.BoolToString(d.Enabled.ValueBool()),
		Domain:  api.SelectedMap(d.Domain.ValueString()),
		Name:    d.Name.ValueString(),
		Type:    api.SelectedMap(d.Type.ValueString()),
		Value:   d.Value.ValueString(),
	}, nil
}

func convertRecordStructToSchema(d *bind.Record) (*recordResourceModel, error) {
	return &recordResourceModel{
		Enabled: types.BoolValue(tools.StringToBool(d.Enabled)),
		Domain:  types.StringValue(d.Domain.String()),
		Name:    types.StringValue(d.Name),
		Type:    types.StringValue(d.Type.String()),
		Value:   types.StringValue(d.Value),
	}, nil
}
