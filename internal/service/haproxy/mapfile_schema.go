package haproxy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type mapfileResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Type        types.String `tfsdk:"type"`
	Content     types.String `tfsdk:"content"`
	URL         types.String `tfsdk:"url"`

	Id types.String `tfsdk:"id"`
}

func mapfileResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "HAProxy map files (**Services ▸ HAProxy ▸ Settings ▸ Advanced ▸ Map Files**, plugin `os-haproxy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify this map file.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this map file. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "The type of the map data. One of `beg`, `dom`, `end`, `int`, `ip`, `reg`, `str`, `sub`. Defaults to `dom`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("dom"),
				Validators:          []validator.String{stringvalidator.OneOf("beg", "dom", "end", "int", "ip", "reg", "str", "sub")},
			},
			"content": schema.StringAttribute{
				MarkdownDescription: "Paste the content of the map file here. See the HAProxy documentation for a full description. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "A URL from which the map file will be downloaded. The contents will be cached until next reload/restart of HAProxy. If the download fails, then the specified content will be used instead. In order to enable automatic updates of map files, setup a cron job that periodically reloads HAProxy. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the map file.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func mapfileDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "HAProxy map files (**Services ▸ HAProxy ▸ Settings ▸ Advanced ▸ Map Files**, plugin `os-haproxy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the map file.",
				Required:            true,
			},
			"name":        dschema.StringAttribute{MarkdownDescription: "Name to identify this map file.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Description for this map file.", Computed: true},
			"type":        dschema.StringAttribute{MarkdownDescription: "The type of the map data. One of `beg`, `dom`, `end`, `int`, `ip`, `reg`, `str`, `sub`.", Computed: true},
			"content":     dschema.StringAttribute{MarkdownDescription: "Paste the content of the map file here. See the HAProxy documentation for a full description.", Computed: true},
			"url":         dschema.StringAttribute{MarkdownDescription: "A URL from which the map file will be downloaded. The contents will be cached until next reload/restart of HAProxy. If the download fails, then the specified content will be used instead. In order to enable automatic updates of map files, setup a cron job that periodically reloads HAProxy.", Computed: true},
		},
	}
}

func convertMapfileSchemaToStruct(d *mapfileResourceModel) (*haproxy.Mapfile, error) {
	return &haproxy.Mapfile{
		Name:        d.Name.ValueString(),
		Description: d.Description.ValueString(),
		Type:        api.SelectedMap(d.Type.ValueString()),
		Content:     d.Content.ValueString(),
		Url:         d.URL.ValueString(),
	}, nil
}

func convertMapfileStructToSchema(d *haproxy.Mapfile) (*mapfileResourceModel, error) {
	return &mapfileResourceModel{
		Name:        types.StringValue(d.Name),
		Description: types.StringValue(d.Description),
		Type:        types.StringValue(d.Type.String()),
		Content:     types.StringValue(d.Content),
		URL:         types.StringValue(d.Url),
	}, nil
}
