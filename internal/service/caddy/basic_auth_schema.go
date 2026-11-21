package caddy

import (
	"github.com/browningluke/opnsense-go/pkg/caddy"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type basicAuthResourceModel struct {
	User        types.String `tfsdk:"user"`
	Password    types.String `tfsdk:"password"`
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func basicAuthResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Caddy basic authentication users (**Services ▸ Caddy Web Server ▸ Reverse Proxy ▸ Access**, plugin `os-caddy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"user": schema.StringAttribute{
				MarkdownDescription: "Enter a username. Afterwards, select it in a \"Domain\" or \"Subdomain\" to restrict access with Basic Auth.",
				Required:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Enter a password. It will be hashed with bcrypt. It can only be set and changed but will not be visible anymore.",
				Required:            true,
				Sensitive:           true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Enter a description for this user. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the Caddy basic auth user.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func basicAuthDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Caddy basic authentication users (**Services ▸ Caddy Web Server ▸ Reverse Proxy ▸ Access**, plugin `os-caddy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the Caddy basic auth user.",
				Required:            true,
			},
			"user":        dschema.StringAttribute{MarkdownDescription: "Enter a username. Afterwards, select it in a \"Domain\" or \"Subdomain\" to restrict access with Basic Auth.", Computed: true},
			"password":    dschema.StringAttribute{MarkdownDescription: "Enter a password. Never returned by OPNsense, always null.", Computed: true, Sensitive: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Enter a description for this user.", Computed: true},
		},
	}
}

func convertBasicAuthSchemaToStruct(d *basicAuthResourceModel) (*caddy.BasicAuth, error) {
	return &caddy.BasicAuth{
		Basicauthuser: d.User.ValueString(),
		Basicauthpass: d.Password.ValueString(),
		Description:   d.Description.ValueString(),
	}, nil
}

func convertBasicAuthStructToSchema(d *caddy.BasicAuth) (*basicAuthResourceModel, error) {
	return &basicAuthResourceModel{
		User:        types.StringValue(d.Basicauthuser),
		Password:    types.StringNull(),
		Description: types.StringValue(d.Description),
	}, nil
}
