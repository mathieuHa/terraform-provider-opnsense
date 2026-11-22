package nginx

import (
	"context"

	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type sniHostnameMapResourceModel struct {
	Description types.String `tfsdk:"description"`
	Entries     types.List   `tfsdk:"entries"`

	Id types.String `tfsdk:"id"`
}

type sniHostnameMapEntryModel struct {
	Hostname types.String `tfsdk:"hostname"`
	Upstream types.String `tfsdk:"upstream"`
}

var sniHostnameMapEntryTypes = map[string]attr.Type{
	"hostname": types.StringType,
	"upstream": types.StringType,
}

func sniHostnameMapResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx SNI upstream maps (**Services ▸ Nginx ▸ Configuration ▸ Streams ▸ SNI Based Routing**, plugin `os-nginx`): send TLS streams to an upstream chosen by the requested host name, without terminating TLS.",
		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				MarkdownDescription: "Name of the map, shown in the drop-downs.",
				Required:            true,
			},
			"entries": schema.ListNestedAttribute{
				MarkdownDescription: "Host names and the upstreams they go to.",
				Required:            true,
				Validators:          []validator.List{listvalidator.SizeAtLeast(1)},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"hostname": schema.StringAttribute{
							MarkdownDescription: "Host name sent by the client (SNI).",
							Required:            true,
						},
						"upstream": schema.StringAttribute{
							MarkdownDescription: "UUID of an `opnsense_nginx_upstream`.",
							Required:            true,
						},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the map.",
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func sniHostnameMapDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx SNI upstream maps (**Services ▸ Nginx ▸ Configuration ▸ Streams ▸ SNI Based Routing**, plugin `os-nginx`).",
		Attributes: map[string]dschema.Attribute{
			"id":          dschema.StringAttribute{MarkdownDescription: "UUID of the map.", Required: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Name of the map.", Computed: true},
			"entries": dschema.ListNestedAttribute{
				MarkdownDescription: "Host names and the upstreams they go to.",
				Computed:            true,
				NestedObject: dschema.NestedAttributeObject{
					Attributes: map[string]dschema.Attribute{
						"hostname": dschema.StringAttribute{MarkdownDescription: "Host name sent by the client.", Computed: true},
						"upstream": dschema.StringAttribute{MarkdownDescription: "UUID of the upstream.", Computed: true},
					},
				},
			},
		},
	}
}

func convertSniHostnameMapSchemaToStruct(d *sniHostnameMapResourceModel) (*nginx.SniHostnameMap, error) {
	var entries []sniHostnameMapEntryModel
	d.Entries.ElementsAs(context.Background(), &entries, false)
	data := []nginx.SniHostnameMapEntry{}
	for _, e := range entries {
		data = append(data, nginx.SniHostnameMapEntry{Hostname: e.Hostname.ValueString(), Upstream: e.Upstream.ValueString()})
	}
	return &nginx.SniHostnameMap{Description: d.Description.ValueString(), Data: data}, nil
}

func convertSniHostnameMapStructToSchema(d *nginx.SniHostnameMap) (*sniHostnameMapResourceModel, error) {
	entries := []sniHostnameMapEntryModel{}
	for _, e := range d.Data {
		entries = append(entries, sniHostnameMapEntryModel{Hostname: types.StringValue(e.Hostname), Upstream: types.StringValue(e.Upstream)})
	}
	list, _ := types.ListValueFrom(context.Background(), types.ObjectType{AttrTypes: sniHostnameMapEntryTypes}, entries)
	return &sniHostnameMapResourceModel{Description: types.StringValue(d.Description), Entries: list}, nil
}
