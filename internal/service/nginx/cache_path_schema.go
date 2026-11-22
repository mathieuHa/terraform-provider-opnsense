package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type cachePathResourceModel struct {
	Path        types.String `tfsdk:"path"`
	Size        types.Int64  `tfsdk:"size"`
	Inactive    types.Int64  `tfsdk:"inactive"`
	UseTempPath types.Bool   `tfsdk:"use_temp_path"`
	MaxSize     types.Int64  `tfsdk:"max_size"`

	Id types.String `tfsdk:"id"`
}

func cachePathResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx cache paths (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ Cache Path**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"path": schema.StringAttribute{
				MarkdownDescription: "Path.",
				Required:            true,
			},
			"size": schema.Int64Attribute{
				MarkdownDescription: "Size (MB). Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"inactive": schema.Int64Attribute{
				MarkdownDescription: "Inactive Time (Minutes). Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"use_temp_path": schema.BoolAttribute{
				MarkdownDescription: "Use Temp Path. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"max_size": schema.Int64Attribute{
				MarkdownDescription: "Maximum Size (GB). Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx cache path.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func cachePathDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx cache paths (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ Cache Path**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx cache path.",
				Required:            true,
			},
			"path":          dschema.StringAttribute{MarkdownDescription: "Path.", Computed: true},
			"size":          dschema.Int64Attribute{MarkdownDescription: "Size (MB).", Computed: true},
			"inactive":      dschema.Int64Attribute{MarkdownDescription: "Inactive Time (Minutes).", Computed: true},
			"use_temp_path": dschema.BoolAttribute{MarkdownDescription: "Use Temp Path.", Computed: true},
			"max_size":      dschema.Int64Attribute{MarkdownDescription: "Maximum Size (GB).", Computed: true},
		},
	}
}

func convertCachePathSchemaToStruct(d *cachePathResourceModel) (*nginx.CachePath, error) {
	return &nginx.CachePath{
		Path:        d.Path.ValueString(),
		Size:        tools.Int64ToStringNegative(d.Size.ValueInt64()),
		Inactive:    tools.Int64ToStringNegative(d.Inactive.ValueInt64()),
		UseTempPath: tools.BoolToString(d.UseTempPath.ValueBool()),
		MaxSize:     tools.Int64ToStringNegative(d.MaxSize.ValueInt64()),
	}, nil
}

func convertCachePathStructToSchema(d *nginx.CachePath) (*cachePathResourceModel, error) {
	return &cachePathResourceModel{
		Path:        types.StringValue(d.Path),
		Size:        types.Int64Value(tools.StringToInt64(d.Size)),
		Inactive:    types.Int64Value(tools.StringToInt64(d.Inactive)),
		UseTempPath: types.BoolValue(tools.StringToBool(d.UseTempPath)),
		MaxSize:     types.Int64Value(tools.StringToInt64(d.MaxSize)),
	}, nil
}
