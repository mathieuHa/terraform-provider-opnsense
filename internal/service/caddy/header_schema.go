package caddy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/caddy"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type headerResourceModel struct {
	Direction   types.String `tfsdk:"direction"`
	Header      types.String `tfsdk:"header"`
	Value       types.String `tfsdk:"value"`
	Replace     types.String `tfsdk:"replace"`
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func headerResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Header manipulation rules for Caddy handlers (**Services ▸ Caddy Web Server ▸ Reverse Proxy ▸ Headers**, plugin `os-caddy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"direction": schema.StringAttribute{
				MarkdownDescription: "\"header_up\" sets, adds (with the \"+\" prefix), deletes (with the \"-\" prefix), or performs a replacement (by using two arguments, a search and replacement) in a request header going upstream to the backend. \"header_down\" sets, adds (with the \"+\" prefix), deletes (with the \"-\" prefix), or performs a replacement (by using two arguments, a search and replacement) in a response header coming downstream from the backend. One of `header_up`, `header_down`. Defaults to `header_up`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("header_up"),
				Validators:          []validator.String{stringvalidator.OneOf("header_up", "header_down")},
			},
			"header": schema.StringAttribute{
				MarkdownDescription: "Enter a header, for example \"Host\". Use the \"+\" or \"-\" prefix to add or remove this header, for example \"-Host\" or \"+Host\". A suffix match like \"-Host-*\" is also supported. To replace a header, use \"Host\" without \"+\" or \"-\".",
				Required:            true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "Enter a value for the selected header. One of the most common options is \"{upstream_hostport}\". It is also possible to use a regular expression to search for a specific value in a header. For example: \"^prefix-([A-Za-z0-9]*)$\" which uses the regular expression language RE2 included in Go. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"replace": schema.StringAttribute{
				MarkdownDescription: "If a regular expression is used to search for a Header Value, the replacement string can be set here. For example: \"replaced-$1-suffix\" which expands the replacement string, allowing the use of captured values, \"$1\" being the first capture group. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Enter a description for this header. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the Caddy header rule.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func headerDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Header manipulation rules for Caddy handlers (**Services ▸ Caddy Web Server ▸ Reverse Proxy ▸ Headers**, plugin `os-caddy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the Caddy header rule.",
				Required:            true,
			},
			"direction":   dschema.StringAttribute{MarkdownDescription: "\"header_up\" sets, adds (with the \"+\" prefix), deletes (with the \"-\" prefix), or performs a replacement (by using two arguments, a search and replacement) in a request header going upstream to the backend. \"header_down\" sets, adds (with the \"+\" prefix), deletes (with the \"-\" prefix), or performs a replacement (by using two arguments, a search and replacement) in a response header coming downstream from the backend. One of `header_up`, `header_down`.", Computed: true},
			"header":      dschema.StringAttribute{MarkdownDescription: "Enter a header, for example \"Host\". Use the \"+\" or \"-\" prefix to add or remove this header, for example \"-Host\" or \"+Host\". A suffix match like \"-Host-*\" is also supported. To replace a header, use \"Host\" without \"+\" or \"-\".", Computed: true},
			"value":       dschema.StringAttribute{MarkdownDescription: "Enter a value for the selected header. One of the most common options is \"{upstream_hostport}\". It is also possible to use a regular expression to search for a specific value in a header. For example: \"^prefix-([A-Za-z0-9]*)$\" which uses the regular expression language RE2 included in Go.", Computed: true},
			"replace":     dschema.StringAttribute{MarkdownDescription: "If a regular expression is used to search for a Header Value, the replacement string can be set here. For example: \"replaced-$1-suffix\" which expands the replacement string, allowing the use of captured values, \"$1\" being the first capture group.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Enter a description for this header.", Computed: true},
		},
	}
}

func convertHeaderSchemaToStruct(d *headerResourceModel) (*caddy.Header, error) {
	return &caddy.Header{
		HeaderUpDown:  api.SelectedMap(d.Direction.ValueString()),
		HeaderType:    d.Header.ValueString(),
		HeaderValue:   d.Value.ValueString(),
		HeaderReplace: d.Replace.ValueString(),
		Description:   d.Description.ValueString(),
	}, nil
}

func convertHeaderStructToSchema(d *caddy.Header) (*headerResourceModel, error) {
	return &headerResourceModel{
		Direction:   types.StringValue(d.HeaderUpDown.String()),
		Header:      types.StringValue(d.HeaderType),
		Value:       types.StringValue(d.HeaderValue),
		Replace:     types.StringValue(d.HeaderReplace),
		Description: types.StringValue(d.Description),
	}, nil
}
