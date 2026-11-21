package auth

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/auth"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type groupResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Gid         types.String `tfsdk:"gid"`
	Description types.String `tfsdk:"description"`
	Privileges  types.Set    `tfsdk:"privileges"`

	Id types.String `tfsdk:"id"`
}

func groupResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Local groups (**System ▸ Access ▸ Groups**). Membership is managed on the user side with `opnsense_auth_user.group_memberships`; updating a group never changes its members.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Group name.",
				Required:            true,
			},
			"gid": schema.StringAttribute{
				MarkdownDescription: "Group ID assigned by OPNsense; users refer to the group by this ID.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"privileges": schema.SetAttribute{
				MarkdownDescription: "Privileges granted to members, e.g. `page-all` or `page-firewall-rules`. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the group.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func groupDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Local groups (**System ▸ Access ▸ Groups**). Membership is managed on the user side with `opnsense_auth_user.group_memberships`; updating a group never changes its members.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the group.",
				Required:            true,
			},
			"name":        dschema.StringAttribute{MarkdownDescription: "Group name.", Computed: true},
			"gid":         dschema.StringAttribute{MarkdownDescription: "Group ID assigned by OPNsense; users refer to the group by this ID.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Optional description here for your reference (not parsed).", Computed: true},
			"privileges":  dschema.SetAttribute{MarkdownDescription: "Privileges granted to members, e.g. `page-all` or `page-firewall-rules`.", Computed: true, ElementType: types.StringType},
		},
	}
}

func convertGroupSchemaToStruct(d *groupResourceModel) (*auth.Group, error) {
	return &auth.Group{
		Name:        d.Name.ValueString(),
		GroupId:     d.Gid.ValueString(),
		Description: d.Description.ValueString(),
		Privilege:   api.SelectedMapList(setToSlice(d.Privileges)),
	}, nil
}

func convertGroupStructToSchema(d *auth.Group) (*groupResourceModel, error) {
	return &groupResourceModel{
		Name:        types.StringValue(d.Name),
		Gid:         tools.StringOrNull(d.GroupId),
		Description: tools.StringOrNull(d.Description),
		Privileges:  tools.StringSliceToSet(d.Privilege),
	}, nil
}
