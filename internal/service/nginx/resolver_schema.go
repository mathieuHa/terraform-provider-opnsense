package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type resolverResourceModel struct {
	Description types.String `tfsdk:"description"`
	Address     types.Set    `tfsdk:"address"`
	Valid       types.Int64  `tfsdk:"valid"`
	Ipv4Off     types.Bool   `tfsdk:"ipv4_off"`
	Ipv6Off     types.Bool   `tfsdk:"ipv6_off"`
	Timeout     types.Int64  `tfsdk:"timeout"`

	Id types.String `tfsdk:"id"`
}

func resolverResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx DNS resolvers (**Services ▸ Nginx ▸ Configuration ▸ Other ▸ Resolver**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				MarkdownDescription: "Brief description for reference.",
				Required:            true,
			},
			"address": schema.SetAttribute{
				MarkdownDescription: "Enter a list of IP addresses and ports (optional). Uses port 53 if omitted. Like \"8.8.8.8, [2001:4860:4860::8888], 8.8.8.8:5353\".",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"valid": schema.Int64Attribute{
				MarkdownDescription: "Override response TTL in seconds. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"ipv4_off": schema.BoolAttribute{
				MarkdownDescription: "Do not request IPv4 addresses. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ipv6_off": schema.BoolAttribute{
				MarkdownDescription: "Do not request IPv6 addresses. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: "Resolver time out in seconds. NGINX default (if empty) is 30 seconds. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx DNS resolver.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func resolverDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx DNS resolvers (**Services ▸ Nginx ▸ Configuration ▸ Other ▸ Resolver**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx DNS resolver.",
				Required:            true,
			},
			"description": dschema.StringAttribute{MarkdownDescription: "Brief description for reference.", Computed: true},
			"address":     dschema.SetAttribute{MarkdownDescription: "Enter a list of IP addresses and ports (optional). Uses port 53 if omitted. Like \"8.8.8.8, [2001:4860:4860::8888], 8.8.8.8:5353\".", Computed: true, ElementType: types.StringType},
			"valid":       dschema.Int64Attribute{MarkdownDescription: "Override response TTL in seconds.", Computed: true},
			"ipv4_off":    dschema.BoolAttribute{MarkdownDescription: "Do not request IPv4 addresses.", Computed: true},
			"ipv6_off":    dschema.BoolAttribute{MarkdownDescription: "Do not request IPv6 addresses.", Computed: true},
			"timeout":     dschema.Int64Attribute{MarkdownDescription: "Resolver time out in seconds. NGINX default (if empty) is 30 seconds.", Computed: true},
		},
	}
}

func convertResolverSchemaToStruct(d *resolverResourceModel) (*nginx.Resolver, error) {
	return &nginx.Resolver{
		Description: d.Description.ValueString(),
		Address:     api.SelectedMapList(setToSlice(d.Address)),
		Valid:       tools.Int64ToStringNegative(d.Valid.ValueInt64()),
		Ipv4Off:     tools.BoolToString(d.Ipv4Off.ValueBool()),
		Ipv6Off:     tools.BoolToString(d.Ipv6Off.ValueBool()),
		Timeout:     tools.Int64ToStringNegative(d.Timeout.ValueInt64()),
	}, nil
}

func convertResolverStructToSchema(d *nginx.Resolver) (*resolverResourceModel, error) {
	return &resolverResourceModel{
		Description: types.StringValue(d.Description),
		Address:     tools.StringSliceToSet(d.Address),
		Valid:       types.Int64Value(tools.StringToInt64(d.Valid)),
		Ipv4Off:     types.BoolValue(tools.StringToBool(d.Ipv4Off)),
		Ipv6Off:     types.BoolValue(tools.StringToBool(d.Ipv6Off)),
		Timeout:     types.Int64Value(tools.StringToInt64(d.Timeout)),
	}, nil
}
