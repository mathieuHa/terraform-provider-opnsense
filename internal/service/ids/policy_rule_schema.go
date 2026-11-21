package ids

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/ids"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

type policyRuleResourceModel struct {
	Sid     types.Int64  `tfsdk:"sid"`
	Enabled types.Bool   `tfsdk:"enabled"`
	Action  types.String `tfsdk:"action"`

	Id types.String `tfsdk:"id"`
}

func policyRuleResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Per-rule adjustments of installed Suricata rules, by signature ID (**Services ▸ Intrusion Detection ▸ Policy ▸ Rule adjustments**). Prefer `opnsense_ids_policy` for more than a few rules. Changes are applied with a rule reload.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"sid": schema.Int64Attribute{
				MarkdownDescription: "Signature ID of the rule.",
				Required:            true,
				Validators:          []validator.Int64{int64validator.AtLeast(1)},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether the rule is enabled. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"action": schema.StringAttribute{
				MarkdownDescription: "Action: `alert` or `drop`. Defaults to `alert`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("alert"),
				Validators:          []validator.String{stringvalidator.OneOf("alert", "drop")},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the IDS rule adjustment.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func policyRuleDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Per-rule adjustments of installed Suricata rules, by signature ID (**Services ▸ Intrusion Detection ▸ Policy ▸ Rule adjustments**). Prefer `opnsense_ids_policy` for more than a few rules. Changes are applied with a rule reload.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the IDS rule adjustment.",
				Required:            true,
			},
			"sid":     dschema.Int64Attribute{MarkdownDescription: "Signature ID of the rule.", Computed: true},
			"enabled": dschema.BoolAttribute{MarkdownDescription: "Whether the rule is enabled.", Computed: true},
			"action":  dschema.StringAttribute{MarkdownDescription: "Action: `alert` or `drop`.", Computed: true},
		},
	}
}

func convertPolicyRuleSchemaToStruct(d *policyRuleResourceModel) (*ids.PolicyRule, error) {
	return &ids.PolicyRule{
		Sid:     tools.Int64ToString(d.Sid.ValueInt64()),
		Enabled: tools.BoolToString(d.Enabled.ValueBool()),
		Action:  api.SelectedMap(d.Action.ValueString()),
	}, nil
}

func convertPolicyRuleStructToSchema(d *ids.PolicyRule) (*policyRuleResourceModel, error) {
	return &policyRuleResourceModel{
		Sid:     types.Int64Value(tools.StringToInt64(d.Sid)),
		Enabled: types.BoolValue(tools.StringToBool(d.Enabled)),
		Action:  types.StringValue(d.Action.String()),
	}, nil
}
