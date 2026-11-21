package ids

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/ids"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
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

type policyResourceModel struct {
	Enabled     types.Bool   `tfsdk:"enabled"`
	Priority    types.Int64  `tfsdk:"priority"`
	Action      types.Set    `tfsdk:"action"`
	Rulesets    types.Set    `tfsdk:"rulesets"`
	Content     types.Set    `tfsdk:"content"`
	NewAction   types.String `tfsdk:"new_action"`
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func policyResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Policies that change the action of installed Suricata rules in bulk (**Services ▸ Intrusion Detection ▸ Policy**). Changes are applied with a rule reload.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether the policy is active. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "Priority, higher values are applied later and win. Defaults to `0`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators:          []validator.Int64{int64validator.AtLeast(0)},
			},
			"action": schema.SetAttribute{
				MarkdownDescription: "Only match rules whose default action is one of these: `disable`, `alert`, `drop`. Empty matches all. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
				Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("disable", "alert", "drop"))},
			},
			"rulesets": schema.SetAttribute{
				MarkdownDescription: "UUIDs of the enabled rulesets this policy applies to, e.g. `[opnsense_ids_ruleset.x.uuid]`. Empty applies to all. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"content": schema.SetAttribute{
				MarkdownDescription: "Rule metadata to match, as `property.value`, e.g. `[\"signature_severity.Major\"]`. Values must exist in the installed rules. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"new_action": schema.StringAttribute{
				MarkdownDescription: "Action given to matching rules: `default`, `alert`, `drop` or `disable`. Defaults to `alert`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("alert"),
				Validators:          []validator.String{stringvalidator.OneOf("default", "alert", "drop", "disable")},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the IDS policy.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func policyDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Policies that change the action of installed Suricata rules in bulk (**Services ▸ Intrusion Detection ▸ Policy**). Changes are applied with a rule reload.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the IDS policy.",
				Required:            true,
			},
			"enabled":     dschema.BoolAttribute{MarkdownDescription: "Whether the policy is active.", Computed: true},
			"priority":    dschema.Int64Attribute{MarkdownDescription: "Priority, higher values are applied later and win.", Computed: true},
			"action":      dschema.SetAttribute{MarkdownDescription: "Only match rules whose default action is one of these: `disable`, `alert`, `drop`. Empty matches all.", Computed: true, ElementType: types.StringType},
			"rulesets":    dschema.SetAttribute{MarkdownDescription: "UUIDs of the enabled rulesets this policy applies to, e.g. `[opnsense_ids_ruleset.x.uuid]`. Empty applies to all.", Computed: true, ElementType: types.StringType},
			"content":     dschema.SetAttribute{MarkdownDescription: "Rule metadata to match, as `property.value`, e.g. `[\"signature_severity.Major\"]`. Values must exist in the installed rules.", Computed: true, ElementType: types.StringType},
			"new_action":  dschema.StringAttribute{MarkdownDescription: "Action given to matching rules: `default`, `alert`, `drop` or `disable`.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Description.", Computed: true},
		},
	}
}

func convertPolicySchemaToStruct(d *policyResourceModel) (*ids.Policy, error) {
	return &ids.Policy{
		Enabled:     tools.BoolToString(d.Enabled.ValueBool()),
		Priority:    tools.Int64ToString(d.Priority.ValueInt64()),
		Action:      api.SelectedMapList(setToSlice(d.Action)),
		Rulesets:    api.SelectedMapList(setToSlice(d.Rulesets)),
		Content:     api.SelectedMapList(setToSlice(d.Content)),
		NewAction:   api.SelectedMap(d.NewAction.ValueString()),
		Description: d.Description.ValueString(),
	}, nil
}

func convertPolicyStructToSchema(d *ids.Policy) (*policyResourceModel, error) {
	return &policyResourceModel{
		Enabled:     types.BoolValue(tools.StringToBool(d.Enabled)),
		Priority:    types.Int64Value(tools.StringToInt64(d.Priority)),
		Action:      tools.StringSliceToSet(d.Action),
		Rulesets:    tools.StringSliceToSet(d.Rulesets),
		Content:     tools.StringSliceToSet(d.Content),
		NewAction:   types.StringValue(d.NewAction.String()),
		Description: tools.StringOrNull(d.Description),
	}, nil
}
