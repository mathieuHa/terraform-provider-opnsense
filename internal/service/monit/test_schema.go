package monit

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/monit"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type testResourceModel struct {
	Name      types.String `tfsdk:"name"`
	Type      types.String `tfsdk:"type"`
	Condition types.String `tfsdk:"condition"`
	Action    types.String `tfsdk:"action"`
	Path      types.String `tfsdk:"path"`

	Id types.String `tfsdk:"id"`
}

func testResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Monit tests (**Services ▸ Monit ▸ Settings ▸ Service Tests Settings**) are reusable conditions that services refer to.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the test.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Test type. When unset, OPNsense derives it from `condition` (e.g. `cpu usage …` becomes `SystemResource`); a value that contradicts the condition is replaced by OPNsense.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"condition": schema.StringAttribute{
				MarkdownDescription: "Monit condition, e.g. `cpu usage is greater than 75%`.",
				Required:            true,
			},
			"action": schema.StringAttribute{
				MarkdownDescription: "Action taken when the condition matches.",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf("alert", "exec", "restart", "start", "stop", "unmonitor")},
			},
			"path": schema.StringAttribute{
				MarkdownDescription: "Script to run when `action` is `exec`.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the Monit test.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func testDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Monit tests (**Services ▸ Monit ▸ Settings ▸ Service Tests Settings**) are reusable conditions that services refer to.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the Monit test.",
				Required:            true,
			},
			"name":      dschema.StringAttribute{MarkdownDescription: "Name of the test.", Computed: true},
			"type":      dschema.StringAttribute{MarkdownDescription: "Test type. When unset, OPNsense derives it from `condition` (e.g. `cpu usage …` becomes `SystemResource`); a value that contradicts the condition is replaced by OPNsense.", Computed: true},
			"condition": dschema.StringAttribute{MarkdownDescription: "Monit condition, e.g. `cpu usage is greater than 75%`.", Computed: true},
			"action":    dschema.StringAttribute{MarkdownDescription: "Action taken when the condition matches.", Computed: true},
			"path":      dschema.StringAttribute{MarkdownDescription: "Script to run when `action` is `exec`.", Computed: true},
		},
	}
}

func convertTestSchemaToStruct(d *testResourceModel) (*monit.Test, error) {
	return &monit.Test{
		Name: d.Name.ValueString(),
		Type: func() api.SelectedMap {
			if d.Type.ValueString() == "" {
				return api.SelectedMap("Custom")
			}
			return api.SelectedMap(d.Type.ValueString())
		}(),
		Condition: d.Condition.ValueString(),
		Action:    api.SelectedMap(d.Action.ValueString()),
		Path:      d.Path.ValueString(),
	}, nil
}

func convertTestStructToSchema(d *monit.Test) (*testResourceModel, error) {
	return &testResourceModel{
		Name:      types.StringValue(d.Name),
		Type:      types.StringValue(d.Type.String()),
		Condition: types.StringValue(d.Condition),
		Action:    types.StringValue(d.Action.String()),
		Path:      tools.StringOrNull(d.Path),
	}, nil
}
