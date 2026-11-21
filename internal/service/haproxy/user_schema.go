package haproxy

import (
	"github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type userResourceModel struct {
	Enabled     types.Bool   `tfsdk:"enabled"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Password    types.String `tfsdk:"password"`

	Id types.String `tfsdk:"id"`
}

func userResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "HAProxy users (**Services ▸ HAProxy ▸ Settings ▸ Advanced ▸ Users**, plugin `os-haproxy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this user. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify this user.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this user. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Both encrypted and unencrypted passwords can be used. Most systems support MD5, SHA-256, SHA-512, and, of course, the classic DES-based method of encrypting passwords.NOTE: Avoid using unencrypted passwords that start with a $-sign, because this indicates an encrypted password and will make it impossible to authenticate.",
				Required:            true,
				Sensitive:           true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the user.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func userDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "HAProxy users (**Services ▸ HAProxy ▸ Settings ▸ Advanced ▸ Users**, plugin `os-haproxy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the user.",
				Required:            true,
			},
			"enabled":     dschema.BoolAttribute{MarkdownDescription: "Enable this user.", Computed: true},
			"name":        dschema.StringAttribute{MarkdownDescription: "Name to identify this user.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Description for this user.", Computed: true},
			"password":    dschema.StringAttribute{MarkdownDescription: "Both encrypted and unencrypted passwords can be used. Never returned by OPNsense, always null.", Computed: true, Sensitive: true},
		},
	}
}

func convertUserSchemaToStruct(d *userResourceModel) (*haproxy.User, error) {
	return &haproxy.User{
		Enabled:     tools.BoolToString(d.Enabled.ValueBool()),
		Name:        d.Name.ValueString(),
		Description: d.Description.ValueString(),
		Password:    d.Password.ValueString(),
	}, nil
}

func convertUserStructToSchema(d *haproxy.User) (*userResourceModel, error) {
	return &userResourceModel{
		Enabled:     types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:        types.StringValue(d.Name),
		Description: types.StringValue(d.Description),
		Password:    types.StringNull(),
	}, nil
}
