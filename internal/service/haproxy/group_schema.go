package haproxy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type groupResourceModel struct {
	Enabled     types.Bool   `tfsdk:"enabled"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Members     types.Set    `tfsdk:"members"`
	AddUserlist types.Bool   `tfsdk:"add_userlist"`

	Id types.String `tfsdk:"id"`
}

func groupResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "HAProxy user lists (**Services ▸ HAProxy ▸ Settings ▸ Advanced ▸ User Lists**, plugin `os-haproxy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this group. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify this group.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this group. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"members": schema.SetAttribute{
				MarkdownDescription: "Members. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"add_userlist": schema.BoolAttribute{
				MarkdownDescription: "Usually HAproxy userlists are created automatically in a context sensitive way. This option adds this group as userlist, so that it can be referenced in rules/conditions. All special and non-alphanumeric characters will be removed from the userlist name. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the user list.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func groupDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "HAProxy user lists (**Services ▸ HAProxy ▸ Settings ▸ Advanced ▸ User Lists**, plugin `os-haproxy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the user list.",
				Required:            true,
			},
			"enabled":      dschema.BoolAttribute{MarkdownDescription: "Enable this group.", Computed: true},
			"name":         dschema.StringAttribute{MarkdownDescription: "Name to identify this group.", Computed: true},
			"description":  dschema.StringAttribute{MarkdownDescription: "Description for this group.", Computed: true},
			"members":      dschema.SetAttribute{MarkdownDescription: "Members.", Computed: true, ElementType: types.StringType},
			"add_userlist": dschema.BoolAttribute{MarkdownDescription: "Usually HAproxy userlists are created automatically in a context sensitive way. This option adds this group as userlist, so that it can be referenced in rules/conditions. All special and non-alphanumeric characters will be removed from the userlist name.", Computed: true},
		},
	}
}

func convertGroupSchemaToStruct(d *groupResourceModel) (*haproxy.Group, error) {
	return &haproxy.Group{
		Enabled:     tools.BoolToString(d.Enabled.ValueBool()),
		Name:        d.Name.ValueString(),
		Description: d.Description.ValueString(),
		Members:     api.SelectedMapList(setToSlice(d.Members)),
		AddUserlist: tools.BoolToString(d.AddUserlist.ValueBool()),
	}, nil
}

func convertGroupStructToSchema(d *haproxy.Group) (*groupResourceModel, error) {
	return &groupResourceModel{
		Enabled:     types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:        types.StringValue(d.Name),
		Description: types.StringValue(d.Description),
		Members:     tools.StringSliceToSet(d.Members),
		AddUserlist: types.BoolValue(tools.StringToBool(d.AddUserlist)),
	}, nil
}
