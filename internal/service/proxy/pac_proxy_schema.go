package proxy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/proxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type pacProxyResourceModel struct {
	Name        types.String `tfsdk:"name"`
	ProxyType   types.String `tfsdk:"proxy_type"`
	URL         types.String `tfsdk:"url"`
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func pacProxyResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Proxies offered to clients by the proxy auto-config file (**Services ▸ Squid Web Proxy ▸ Proxy Auto Config ▸ Proxies**).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the proxy.",
				Required:            true,
			},
			"proxy_type": schema.StringAttribute{
				MarkdownDescription: "Proxy type written to the PAC file. Defaults to `PROXY`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("PROXY"),
				Validators:          []validator.String{stringvalidator.OneOf("PROXY", "DIRECT", "HTTP", "HTTPS", "SOCKS", "SOCKS4", "SOCKS5")},
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "Proxy as `host:port`. Not used with `DIRECT`.",
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the PAC proxy.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func pacProxyDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Proxies offered to clients by the proxy auto-config file (**Services ▸ Squid Web Proxy ▸ Proxy Auto Config ▸ Proxies**).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the PAC proxy.",
				Required:            true,
			},
			"name":        dschema.StringAttribute{MarkdownDescription: "Name of the proxy.", Computed: true},
			"proxy_type":  dschema.StringAttribute{MarkdownDescription: "Proxy type written to the PAC file.", Computed: true},
			"url":         dschema.StringAttribute{MarkdownDescription: "Proxy as `host:port`. Not used with `DIRECT`.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Optional description here for your reference (not parsed).", Computed: true},
		},
	}
}

func convertPacProxySchemaToStruct(d *pacProxyResourceModel) (*proxy.PacProxy, error) {
	return &proxy.PacProxy{
		Name:        d.Name.ValueString(),
		ProxyType:   api.SelectedMap(d.ProxyType.ValueString()),
		URL:         d.URL.ValueString(),
		Description: d.Description.ValueString(),
	}, nil
}

func convertPacProxyStructToSchema(d *proxy.PacProxy) (*pacProxyResourceModel, error) {
	return &pacProxyResourceModel{
		Name:        types.StringValue(d.Name),
		ProxyType:   types.StringValue(d.ProxyType.String()),
		URL:         tools.StringOrNull(d.URL),
		Description: tools.StringOrNull(d.Description),
	}, nil
}
