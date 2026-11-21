package ids

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/ids"
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

type userRuleResourceModel struct {
	Enabled     types.Bool   `tfsdk:"enabled"`
	Source      types.String `tfsdk:"source"`
	Destination types.String `tfsdk:"destination"`
	Fingerprint types.String `tfsdk:"fingerprint"`
	Description types.String `tfsdk:"description"`
	Action      types.String `tfsdk:"action"`
	Bypass      types.Bool   `tfsdk:"bypass"`

	Id types.String `tfsdk:"id"`
}

func userRuleResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Custom Suricata rules matching addresses or TLS fingerprints (**Services ▸ Intrusion Detection ▸ Administration ▸ User defined**). Changes are applied with a rule reload.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether the rule is active. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"source": schema.StringAttribute{
				MarkdownDescription: "Source address or network, empty for any.",
				Optional:            true,
			},
			"destination": schema.StringAttribute{
				MarkdownDescription: "Destination address or network, empty for any.",
				Optional:            true,
			},
			"fingerprint": schema.StringAttribute{
				MarkdownDescription: "SHA1 fingerprint of a TLS certificate, as 59 colon-separated hex characters.",
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description, used as the rule message.",
				Required:            true,
			},
			"action": schema.StringAttribute{
				MarkdownDescription: "Action: `alert`, `drop` (IPS modes only) or `pass`. Defaults to `alert`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("alert"),
				Validators:          []validator.String{stringvalidator.OneOf("alert", "drop", "pass")},
			},
			"bypass": schema.BoolAttribute{
				MarkdownDescription: "With `pass`, bypass the rest of the flow (no further inspection). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the IDS user rule.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func userRuleDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Custom Suricata rules matching addresses or TLS fingerprints (**Services ▸ Intrusion Detection ▸ Administration ▸ User defined**). Changes are applied with a rule reload.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the IDS user rule.",
				Required:            true,
			},
			"enabled":     dschema.BoolAttribute{MarkdownDescription: "Whether the rule is active.", Computed: true},
			"source":      dschema.StringAttribute{MarkdownDescription: "Source address or network, empty for any.", Computed: true},
			"destination": dschema.StringAttribute{MarkdownDescription: "Destination address or network, empty for any.", Computed: true},
			"fingerprint": dschema.StringAttribute{MarkdownDescription: "SHA1 fingerprint of a TLS certificate, as 59 colon-separated hex characters.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Description, used as the rule message.", Computed: true},
			"action":      dschema.StringAttribute{MarkdownDescription: "Action: `alert`, `drop` (IPS modes only) or `pass`.", Computed: true},
			"bypass":      dschema.BoolAttribute{MarkdownDescription: "With `pass`, bypass the rest of the flow (no further inspection).", Computed: true},
		},
	}
}

func convertUserRuleSchemaToStruct(d *userRuleResourceModel) (*ids.UserRule, error) {
	return &ids.UserRule{
		Enabled:     tools.BoolToString(d.Enabled.ValueBool()),
		Source:      d.Source.ValueString(),
		Destination: d.Destination.ValueString(),
		Fingerprint: d.Fingerprint.ValueString(),
		Description: d.Description.ValueString(),
		Action:      api.SelectedMap(d.Action.ValueString()),
		Bypass:      tools.BoolToString(d.Bypass.ValueBool()),
	}, nil
}

func convertUserRuleStructToSchema(d *ids.UserRule) (*userRuleResourceModel, error) {
	return &userRuleResourceModel{
		Enabled:     types.BoolValue(tools.StringToBool(d.Enabled)),
		Source:      tools.StringOrNull(d.Source),
		Destination: tools.StringOrNull(d.Destination),
		Fingerprint: tools.StringOrNull(d.Fingerprint),
		Description: types.StringValue(d.Description),
		Action:      types.StringValue(d.Action.String()),
		Bypass:      types.BoolValue(tools.StringToBool(d.Bypass)),
	}, nil
}
