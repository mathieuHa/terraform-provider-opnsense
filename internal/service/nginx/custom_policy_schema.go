package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type customPolicyResourceModel struct {
	Name       types.String `tfsdk:"name"`
	NaxsiRules types.Set    `tfsdk:"naxsi_rules"`
	Value      types.Int64  `tfsdk:"value"`
	Operator   types.String `tfsdk:"operator"`
	Action     types.String `tfsdk:"action"`

	Id types.String `tfsdk:"id"`
}

func customPolicyResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx NAXSI WAF policys (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ Naxsi WAF Policy**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name.",
				Required:            true,
			},
			"naxsi_rules": schema.SetAttribute{
				MarkdownDescription: "Rules.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"value": schema.Int64Attribute{
				MarkdownDescription: "If the sum of scores of all matching rules exceed the configured value the policy's action will get triggered.",
				Required:            true,
			},
			"operator": schema.StringAttribute{
				MarkdownDescription: "Operator. One of `>=`, `>`, `<`, `<=`. Defaults to `>=`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(">="),
				Validators:          []validator.String{stringvalidator.OneOf(">=", ">", "<", "<=")},
			},
			"action": schema.StringAttribute{
				MarkdownDescription: "Action. One of `BLOCK`, `ALLOW`, `DROP`, `LOG`. Defaults to `BLOCK`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("BLOCK"),
				Validators:          []validator.String{stringvalidator.OneOf("BLOCK", "ALLOW", "DROP", "LOG")},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx NAXSI WAF policy.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func customPolicyDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx NAXSI WAF policys (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ Naxsi WAF Policy**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx NAXSI WAF policy.",
				Required:            true,
			},
			"name":        dschema.StringAttribute{MarkdownDescription: "Name.", Computed: true},
			"naxsi_rules": dschema.SetAttribute{MarkdownDescription: "Rules.", Computed: true, ElementType: types.StringType},
			"value":       dschema.Int64Attribute{MarkdownDescription: "If the sum of scores of all matching rules exceed the configured value the policy's action will get triggered.", Computed: true},
			"operator":    dschema.StringAttribute{MarkdownDescription: "Operator. One of `>=`, `>`, `<`, `<=`.", Computed: true},
			"action":      dschema.StringAttribute{MarkdownDescription: "Action. One of `BLOCK`, `ALLOW`, `DROP`, `LOG`.", Computed: true},
		},
	}
}

func convertCustomPolicySchemaToStruct(d *customPolicyResourceModel) (*nginx.CustomPolicy, error) {
	return &nginx.CustomPolicy{
		Name:       d.Name.ValueString(),
		NaxsiRules: api.SelectedMapList(setToSlice(d.NaxsiRules)),
		Value:      tools.Int64ToString(d.Value.ValueInt64()),
		Operator:   api.SelectedMap(d.Operator.ValueString()),
		Action:     api.SelectedMap(d.Action.ValueString()),
	}, nil
}

func convertCustomPolicyStructToSchema(d *nginx.CustomPolicy) (*customPolicyResourceModel, error) {
	return &customPolicyResourceModel{
		Name:       types.StringValue(d.Name),
		NaxsiRules: tools.StringSliceToSet(d.NaxsiRules),
		Value:      types.Int64Value(tools.StringToInt64(d.Value)),
		Operator:   types.StringValue(d.Operator.String()),
		Action:     types.StringValue(d.Action.String()),
	}, nil
}
