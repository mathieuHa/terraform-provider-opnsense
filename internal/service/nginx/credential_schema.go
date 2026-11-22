package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/nginx"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type credentialResourceModel struct {
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`

	Id types.String `tfsdk:"id"`
}

func credentialResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx credentials (**Services ▸ Nginx ▸ Configuration ▸ Access ▸ Credential**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				MarkdownDescription: "Username.",
				Required:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password.",
				Required:            true,
				Sensitive:           true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx credential.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func credentialDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx credentials (**Services ▸ Nginx ▸ Configuration ▸ Access ▸ Credential**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx credential.",
				Required:            true,
			},
			"username": dschema.StringAttribute{MarkdownDescription: "Username.", Computed: true},
			"password": dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
		},
	}
}

func convertCredentialSchemaToStruct(d *credentialResourceModel) (*nginx.Credential, error) {
	return &nginx.Credential{
		Username: d.Username.ValueString(),
		Password: d.Password.ValueString(),
	}, nil
}

func convertCredentialStructToSchema(d *nginx.Credential) (*credentialResourceModel, error) {
	return &credentialResourceModel{
		Username: types.StringValue(d.Username),
		Password: types.StringValue(d.Password),
	}, nil
}
