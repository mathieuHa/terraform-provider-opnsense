package tailscale

import (
	"github.com/browningluke/opnsense-go/pkg/tailscale"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type subnetResourceModel struct {
	Subnet      types.String `tfsdk:"subnet"`
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func subnetResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Subnets this firewall advertises to the tailnet (**VPN ▸ Tailscale ▸ Settings ▸ Advertised routes**, plugin `os-tailscale`). Routes must also be approved in the Tailscale admin console unless auto-approved.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"subnet": schema.StringAttribute{
				MarkdownDescription: "Network in CIDR notation, e.g. `10.0.10.0/24`.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the Tailscale subnet.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func subnetDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Subnets this firewall advertises to the tailnet (**VPN ▸ Tailscale ▸ Settings ▸ Advertised routes**, plugin `os-tailscale`). Routes must also be approved in the Tailscale admin console unless auto-approved.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the Tailscale subnet.",
				Required:            true,
			},
			"subnet":      dschema.StringAttribute{MarkdownDescription: "Network in CIDR notation, e.g. `10.0.10.0/24`.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Description.", Computed: true},
		},
	}
}

func convertSubnetSchemaToStruct(d *subnetResourceModel) (*tailscale.Subnet, error) {
	return &tailscale.Subnet{
		Subnet:      d.Subnet.ValueString(),
		Description: d.Description.ValueString(),
	}, nil
}

func convertSubnetStructToSchema(d *tailscale.Subnet) (*subnetResourceModel, error) {
	return &subnetResourceModel{
		Subnet:      types.StringValue(d.Subnet),
		Description: tools.StringOrNull(d.Description),
	}, nil
}
