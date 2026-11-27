package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type httpRewriteResourceModel struct {
	Description types.String `tfsdk:"description"`
	Source      types.String `tfsdk:"source"`
	Destination types.String `tfsdk:"destination"`
	Flag        types.String `tfsdk:"flag"`

	Id types.String `tfsdk:"id"`
}

func httpRewriteResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx URL rewrites (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ URL Rewriting**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				MarkdownDescription: "Enter a short description like a name for this redirect. No spaces, tabs or double quotes.",
				Required:            true,
				Validators:          []validator.String{noSpaces},
			},
			"source": schema.StringAttribute{
				MarkdownDescription: "Enter a regular expression to match the URL.",
				Required:            true,
			},
			"destination": schema.StringAttribute{
				MarkdownDescription: "You may use replacement strings like $1 for the first match group from the source here. No spaces, tabs or double quotes.",
				Required:            true,
				Validators:          []validator.String{noSpaces},
			},
			"flag": schema.StringAttribute{
				MarkdownDescription: "Stop rule processing (break) or perform (internal) redirect (last). Return a moved permanently status code (301) or a temporary redirect (302). One of `break`, `last`, `redirect`, `permanent`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "break", "last", "redirect", "permanent")},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx URL rewrite.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func httpRewriteDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx URL rewrites (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ URL Rewriting**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx URL rewrite.",
				Required:            true,
			},
			"description": dschema.StringAttribute{MarkdownDescription: "Enter a short description like a name for this redirect.", Computed: true},
			"source":      dschema.StringAttribute{MarkdownDescription: "Enter a regular expression to match the URL.", Computed: true},
			"destination": dschema.StringAttribute{MarkdownDescription: "You may use replacement strings like $1 for the first match group from the source here.", Computed: true},
			"flag":        dschema.StringAttribute{MarkdownDescription: "Stop rule processing (break) or perform (internal) redirect (last). Return a moved permanently status code (301) or a temporary redirect (302). One of `break`, `last`, `redirect`, `permanent`.", Computed: true},
		},
	}
}

func convertHttpRewriteSchemaToStruct(d *httpRewriteResourceModel) (*nginx.HttpRewrite, error) {
	return &nginx.HttpRewrite{
		Description: d.Description.ValueString(),
		Source:      d.Source.ValueString(),
		Destination: d.Destination.ValueString(),
		Flag:        api.SelectedMap(d.Flag.ValueString()),
	}, nil
}

func convertHttpRewriteStructToSchema(d *nginx.HttpRewrite) (*httpRewriteResourceModel, error) {
	return &httpRewriteResourceModel{
		Description: types.StringValue(d.Description),
		Source:      types.StringValue(d.Source),
		Destination: types.StringValue(d.Destination),
		Flag:        types.StringValue(d.Flag.String()),
	}, nil
}
