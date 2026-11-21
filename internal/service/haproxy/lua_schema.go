package haproxy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type luaResourceModel struct {
	Enabled        types.Bool   `tfsdk:"enabled"`
	Name           types.String `tfsdk:"name"`
	Description    types.String `tfsdk:"description"`
	Preload        types.Bool   `tfsdk:"preload"`
	FilenameScheme types.String `tfsdk:"filename_scheme"`
	Content        types.String `tfsdk:"content"`

	Id types.String `tfsdk:"id"`
}

func luaResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "HAProxy Lua scripts (**Services ▸ HAProxy ▸ Settings ▸ Advanced ▸ Lua Scripts**, plugin `os-haproxy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this Lua script. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify this Lua script.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this Lua script. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"preload": schema.BoolAttribute{
				MarkdownDescription: "Whether HAProxy should load and execute this Lua script on startup. This is the default behaviour. However, if this Lua script is included by other Lua scripts using the \"require\" function, then preloading should be disabled to avoid HAProxy errors. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"filename_scheme": schema.StringAttribute{
				MarkdownDescription: "Specify the filename scheme for this Lua script. Usually using the ID is sufficient and the most fail-safe apparoach. However, when using Lua's \"require\" function this apparoach will not work and it becomes necessary to use the specified name as filename. In this case all non-alphanumeric characters are removed from the filename. Note that this may cause issues when creating multiple Lua scripts with the same name. One of `id`, `name`. Defaults to `id`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("id"),
				Validators:          []validator.String{stringvalidator.OneOf("id", "name")},
			},
			"content": schema.StringAttribute{
				MarkdownDescription: "Paste the content of your Lua script here.",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the Lua script.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func luaDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "HAProxy Lua scripts (**Services ▸ HAProxy ▸ Settings ▸ Advanced ▸ Lua Scripts**, plugin `os-haproxy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the Lua script.",
				Required:            true,
			},
			"enabled":         dschema.BoolAttribute{MarkdownDescription: "Enable this Lua script.", Computed: true},
			"name":            dschema.StringAttribute{MarkdownDescription: "Name to identify this Lua script.", Computed: true},
			"description":     dschema.StringAttribute{MarkdownDescription: "Description for this Lua script.", Computed: true},
			"preload":         dschema.BoolAttribute{MarkdownDescription: "Whether HAProxy should load and execute this Lua script on startup. This is the default behaviour. However, if this Lua script is included by other Lua scripts using the \"require\" function, then preloading should be disabled to avoid HAProxy errors.", Computed: true},
			"filename_scheme": dschema.StringAttribute{MarkdownDescription: "Specify the filename scheme for this Lua script. Usually using the ID is sufficient and the most fail-safe apparoach. However, when using Lua's \"require\" function this apparoach will not work and it becomes necessary to use the specified name as filename. In this case all non-alphanumeric characters are removed from the filename. Note that this may cause issues when creating multiple Lua scripts with the same name. One of `id`, `name`.", Computed: true},
			"content":         dschema.StringAttribute{MarkdownDescription: "Paste the content of your Lua script here.", Computed: true},
		},
	}
}

func convertLuaSchemaToStruct(d *luaResourceModel) (*haproxy.Lua, error) {
	return &haproxy.Lua{
		Enabled:        tools.BoolToString(d.Enabled.ValueBool()),
		Name:           d.Name.ValueString(),
		Description:    d.Description.ValueString(),
		Preload:        tools.BoolToString(d.Preload.ValueBool()),
		FilenameScheme: api.SelectedMap(d.FilenameScheme.ValueString()),
		Content:        d.Content.ValueString(),
	}, nil
}

func convertLuaStructToSchema(d *haproxy.Lua) (*luaResourceModel, error) {
	return &luaResourceModel{
		Enabled:        types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:           types.StringValue(d.Name),
		Description:    types.StringValue(d.Description),
		Preload:        types.BoolValue(tools.StringToBool(d.Preload)),
		FilenameScheme: types.StringValue(d.FilenameScheme.String()),
		Content:        types.StringValue(d.Content),
	}, nil
}
