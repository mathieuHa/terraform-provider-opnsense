package interfaces

import (
	"github.com/browningluke/opnsense-go/pkg/interfaces"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type neighborResourceModel struct {
	MacAddress  types.String `tfsdk:"mac_address"`
	IpAddress   types.String `tfsdk:"ip_address"`
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func neighborResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Static ARP (IPv4) and NDP (IPv6) entries (**Interfaces ▸ Neighbors**).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"mac_address": schema.StringAttribute{
				MarkdownDescription: "MAC address, e.g. `02:00:00:00:00:01`.",
				Required:            true,
			},
			"ip_address": schema.StringAttribute{
				MarkdownDescription: "IPv4 or IPv6 address of the neighbor.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the static neighbor.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func neighborDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Static ARP (IPv4) and NDP (IPv6) entries (**Interfaces ▸ Neighbors**).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the static neighbor.",
				Required:            true,
			},
			"mac_address": dschema.StringAttribute{MarkdownDescription: "MAC address, e.g. `02:00:00:00:00:01`.", Computed: true},
			"ip_address":  dschema.StringAttribute{MarkdownDescription: "IPv4 or IPv6 address of the neighbor.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Description.", Computed: true},
		},
	}
}

func convertNeighborSchemaToStruct(d *neighborResourceModel) (*interfaces.Neighbor, error) {
	return &interfaces.Neighbor{
		MacAddress:  d.MacAddress.ValueString(),
		IpAddress:   d.IpAddress.ValueString(),
		Description: d.Description.ValueString(),
	}, nil
}

func convertNeighborStructToSchema(d *interfaces.Neighbor) (*neighborResourceModel, error) {
	return &neighborResourceModel{
		MacAddress:  types.StringValue(d.MacAddress),
		IpAddress:   types.StringValue(d.IpAddress),
		Description: tools.StringOrNull(d.Description),
	}, nil
}
