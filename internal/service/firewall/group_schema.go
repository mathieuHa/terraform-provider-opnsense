package firewall

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/firewall"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type groupResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Members     types.Set    `tfsdk:"members"`
	NoGroup     types.Bool   `tfsdk:"no_group"`
	Sequence    types.Int64  `tfsdk:"sequence"`
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func groupResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Interface groups combine interfaces so firewall rules can be written once for all of them. Each group also gets its own tab under **Firewall ▸ Rules**.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Group name, used as the interface name in rules. At most 15 characters and must not end with a digit.",
				Required:            true,
			},
			"members": schema.SetAttribute{
				MarkdownDescription: "Interfaces in the group, e.g. `[\"lan\", \"opt1\"]`.",
				Required:            true,
				ElementType:         types.StringType,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"no_group": schema.BoolAttribute{
				MarkdownDescription: "When enabled, the group is hidden from the interface list of the rules menu. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"sequence": schema.Int64Attribute{
				MarkdownDescription: "Position of the group in the rules menu. Defaults to `0`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators:          []validator.Int64{int64validator.Between(0, 9999)},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the interface group.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func groupDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Interface groups combine interfaces so firewall rules can be written once for all of them.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the interface group.",
				Required:            true,
			},
			"name": dschema.StringAttribute{
				MarkdownDescription: "Group name, used as the interface name in rules.",
				Computed:            true,
			},
			"members": dschema.SetAttribute{
				MarkdownDescription: "Interfaces in the group.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"no_group": dschema.BoolAttribute{
				MarkdownDescription: "Whether the group is hidden from the interface list of the rules menu.",
				Computed:            true,
			},
			"sequence": dschema.Int64Attribute{
				MarkdownDescription: "Position of the group in the rules menu.",
				Computed:            true,
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Computed:            true,
			},
		},
	}
}

func convertGroupSchemaToStruct(d *groupResourceModel) (*firewall.Group, error) {
	return &firewall.Group{
		Name:        d.Name.ValueString(),
		Members:     api.SelectedMapList(tools.SetToStringSlice(d.Members)),
		NoGroup:     tools.BoolToString(d.NoGroup.ValueBool()),
		Sequence:    tools.Int64ToString(d.Sequence.ValueInt64()),
		Description: d.Description.ValueString(),
	}, nil
}

func convertGroupStructToSchema(d *firewall.Group) (*groupResourceModel, error) {
	return &groupResourceModel{
		Name:        types.StringValue(d.Name),
		Members:     tools.StringSliceToSet(d.Members),
		NoGroup:     types.BoolValue(tools.StringToBool(d.NoGroup)),
		Sequence:    types.Int64Value(tools.StringToInt64(d.Sequence)),
		Description: tools.StringOrNull(d.Description),
	}, nil
}
