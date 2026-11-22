package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type userlistResourceModel struct {
	Name  types.String `tfsdk:"name"`
	Users types.Set    `tfsdk:"users"`

	Id types.String `tfsdk:"id"`
}

func userlistResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx user lists (**Services ▸ Nginx ▸ Configuration ▸ Access ▸ User List**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name.",
				Required:            true,
			},
			"users": schema.SetAttribute{
				MarkdownDescription: "Users.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx user list.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func userlistDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx user lists (**Services ▸ Nginx ▸ Configuration ▸ Access ▸ User List**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx user list.",
				Required:            true,
			},
			"name":  dschema.StringAttribute{MarkdownDescription: "Name.", Computed: true},
			"users": dschema.SetAttribute{MarkdownDescription: "Users.", Computed: true, ElementType: types.StringType},
		},
	}
}

func convertUserlistSchemaToStruct(d *userlistResourceModel) (*nginx.Userlist, error) {
	return &nginx.Userlist{
		Name:  d.Name.ValueString(),
		Users: api.SelectedMapList(setToSlice(d.Users)),
	}, nil
}

func convertUserlistStructToSchema(d *nginx.Userlist) (*userlistResourceModel, error) {
	return &userlistResourceModel{
		Name:  types.StringValue(d.Name),
		Users: tools.StringSliceToSet(d.Users),
	}, nil
}
