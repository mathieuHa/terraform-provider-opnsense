package nginx

import (
	"encoding/base64"
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type errorpageResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Statuscodes types.Set    `tfsdk:"statuscodes"`
	Pagecontent types.String `tfsdk:"pagecontent"`
	Redirect    types.String `tfsdk:"redirect"`
	Response    types.String `tfsdk:"response"`

	Id types.String `tfsdk:"id"`
}

func errorpageResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Custom nginx error pages (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ Error Pages**, plugin `os-nginx`), used by HTTP servers and locations.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Enter a name for your error page.",
				Required:            true,
			},
			"statuscodes": schema.SetAttribute{
				MarkdownDescription: "Enter the HTTP status codes for which the error page shall be shown. One of `400`, `401`, `403`, `404`, `405`, `407`, `408`, `410`, `415`, `429`, `431`, `500`, `501`, `502`, `503`, `504`.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1), setvalidator.ValueStringsAre(stringvalidator.OneOf("400", "401", "403", "404", "405", "407", "408", "410", "415", "429", "431", "500", "501", "502", "503", "504"))},
			},
			"pagecontent": schema.StringAttribute{
				MarkdownDescription: "The HTML page content to display. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"redirect": schema.StringAttribute{
				MarkdownDescription: "Instead of displaying a static error page redirect to a URL. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"response": schema.StringAttribute{
				MarkdownDescription: "Optionally enter a HTTP status code to change the response to. This does not work in combination with a redirection target. One of `200`, `300`, `301`, `302`, `401`, `403`, `404`, `451`, `500`, `503`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "200", "300", "301", "302", "401", "403", "404", "451", "500", "503")},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx error page.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func errorpageDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Custom nginx error pages (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ Error Pages**, plugin `os-nginx`), used by HTTP servers and locations.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx error page.",
				Required:            true,
			},
			"name":        dschema.StringAttribute{MarkdownDescription: "Enter a name for your error page.", Computed: true},
			"statuscodes": dschema.SetAttribute{MarkdownDescription: "Enter the HTTP status codes for which the error page shall be shown. One of `400`, `401`, `403`, `404`, `405`, `407`, `408`, `410`, `415`, `429`, `431`, `500`, `501`, `502`, `503`, `504`.", Computed: true, ElementType: types.StringType},
			"pagecontent": dschema.StringAttribute{MarkdownDescription: "The HTML page content to display.", Computed: true},
			"redirect":    dschema.StringAttribute{MarkdownDescription: "Instead of displaying a static error page redirect to a URL.", Computed: true},
			"response":    dschema.StringAttribute{MarkdownDescription: "Optionally enter a HTTP status code to change the response to. This does not work in combination with a redirection target. One of `200`, `300`, `301`, `302`, `401`, `403`, `404`, `451`, `500`, `503`.", Computed: true},
		},
	}
}

func convertErrorpageSchemaToStruct(d *errorpageResourceModel) (*nginx.Errorpage, error) {
	return &nginx.Errorpage{
		Name:        d.Name.ValueString(),
		Statuscodes: api.SelectedMapList(setToSlice(d.Statuscodes)),
		Pagecontent: d.Pagecontent.ValueString(),
		Redirect:    d.Redirect.ValueString(),
		Response:    api.SelectedMap(d.Response.ValueString()),
	}, nil
}

func convertErrorpageStructToSchema(d *nginx.Errorpage) (*errorpageResourceModel, error) {
	return &errorpageResourceModel{
		Name:        types.StringValue(d.Name),
		Statuscodes: tools.StringSliceToSet(d.Statuscodes),
		Pagecontent: types.StringValue(decodePageContent(d.Pagecontent)),
		Redirect:    types.StringValue(d.Redirect),
		Response:    types.StringValue(d.Response.String()),
	}, nil
}

func decodePageContent(s string) string {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return s
	}
	return string(b)
}
