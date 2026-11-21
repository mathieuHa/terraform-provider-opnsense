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

type errorfileResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Code        types.String `tfsdk:"code"`
	Content     types.String `tfsdk:"content"`

	Id types.String `tfsdk:"id"`
}

func errorfileResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "HAProxy error messages (**Services ▸ HAProxy ▸ Settings ▸ Advanced ▸ Error Messages**, plugin `os-haproxy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify this error message.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this error message. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"code": schema.StringAttribute{
				MarkdownDescription: "The HTTP status code.NOTE: It is important to understand that error messages are NOT meant to rewrite errors returned by the server, but errors detected and returned by HAProxy. This is why the list of supported errors is limited to a small set. One of `x200`, `x400`, `x403`, `x405`, `x408`, `x429`, `x500`, `x502`, `x503`, `x504`. Defaults to `503`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("503"),
				Validators:          []validator.String{stringvalidator.OneOf("x200", "x400", "x403", "x405", "x408", "x429", "x500", "x502", "x503", "x504")},
			},
			"content": schema.StringAttribute{
				MarkdownDescription: "Paste the content of your error messages here. The message must represent the full HTTP response and include required HTTP headers. It should not exceed the configured buffer size, which generally is 8 or 16 kB.",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the error message.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func errorfileDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "HAProxy error messages (**Services ▸ HAProxy ▸ Settings ▸ Advanced ▸ Error Messages**, plugin `os-haproxy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the error message.",
				Required:            true,
			},
			"name":        dschema.StringAttribute{MarkdownDescription: "Name to identify this error message.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Description for this error message.", Computed: true},
			"code":        dschema.StringAttribute{MarkdownDescription: "The HTTP status code.NOTE: It is important to understand that error messages are NOT meant to rewrite errors returned by the server, but errors detected and returned by HAProxy. This is why the list of supported errors is limited to a small set. One of `x200`, `x400`, `x403`, `x405`, `x408`, `x429`, `x500`, `x502`, `x503`, `x504`.", Computed: true},
			"content":     dschema.StringAttribute{MarkdownDescription: "Paste the content of your error messages here. The message must represent the full HTTP response and include required HTTP headers. It should not exceed the configured buffer size, which generally is 8 or 16 kB.", Computed: true},
		},
	}
}

func convertErrorfileSchemaToStruct(d *errorfileResourceModel) (*haproxy.Errorfile, error) {
	return &haproxy.Errorfile{
		Name:        d.Name.ValueString(),
		Description: d.Description.ValueString(),
		Code:        api.SelectedMap(d.Code.ValueString()),
		Content:     d.Content.ValueString(),
	}, nil
}

func convertErrorfileStructToSchema(d *haproxy.Errorfile) (*errorfileResourceModel, error) {
	return &errorfileResourceModel{
		Name:        types.StringValue(d.Name),
		Description: types.StringValue(d.Description),
		Code:        types.StringValue(d.Code.String()),
		Content:     types.StringValue(d.Content),
	}, nil
}
