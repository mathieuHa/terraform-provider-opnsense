package bind

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/bind"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type aclResourceModel struct {
	Enabled  types.Bool   `tfsdk:"enabled"`
	Name     types.String `tfsdk:"name"`
	Networks types.Set    `tfsdk:"networks"`

	Id types.String `tfsdk:"id"`
}

func aclResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "BIND access control lists (**Services ▸ BIND ▸ Configuration ▸ ACLs**, plugin `os-bind`), used by the settings and the zones.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "This will enable or disable the ACL. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Set the name for this ACL.",
				Required:            true,
			},
			"networks": schema.SetAttribute{
				MarkdownDescription: "List of networks for this ACL.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the BIND ACL.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func aclDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "BIND access control lists (**Services ▸ BIND ▸ Configuration ▸ ACLs**, plugin `os-bind`), used by the settings and the zones.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the BIND ACL.",
				Required:            true,
			},
			"enabled":  dschema.BoolAttribute{MarkdownDescription: "This will enable or disable the ACL.", Computed: true},
			"name":     dschema.StringAttribute{MarkdownDescription: "Set the name for this ACL.", Computed: true},
			"networks": dschema.SetAttribute{MarkdownDescription: "List of networks for this ACL.", Computed: true, ElementType: types.StringType},
		},
	}
}

func convertAclSchemaToStruct(d *aclResourceModel) (*bind.Acl, error) {
	return &bind.Acl{
		Enabled:  tools.BoolToString(d.Enabled.ValueBool()),
		Name:     d.Name.ValueString(),
		Networks: api.SelectedMapList(setToSlice(d.Networks)),
	}, nil
}

func convertAclStructToSchema(d *bind.Acl) (*aclResourceModel, error) {
	return &aclResourceModel{
		Enabled:  types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:     types.StringValue(d.Name),
		Networks: tools.StringSliceToSet(d.Networks),
	}, nil
}
