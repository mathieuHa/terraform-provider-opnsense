package interfaces

import (
	"github.com/browningluke/opnsense-go/pkg/interfaces"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type greResourceModel struct {
	Device              types.String `tfsdk:"device"`
	LocalAddress        types.String `tfsdk:"local_address"`
	RemoteAddress       types.String `tfsdk:"remote_address"`
	TunnelLocalAddress  types.String `tfsdk:"tunnel_local_address"`
	TunnelRemoteAddress types.String `tfsdk:"tunnel_remote_address"`
	TunnelRemoteNet     types.Int64  `tfsdk:"tunnel_remote_net"`
	Description         types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func greResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "GRE tunnel devices (**Interfaces ▸ Devices ▸ GRE**). OPNsense names the device (`greN`) and configures the inner addresses; assign the device to filter its traffic per interface.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "Device name, e.g. `gre0`. Assigned by OPNsense when unset.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: "Local outer address: an interface name such as `wan`, a CARP VIP (`wan_vip1`), an IP alias or an IP address.",
				Required:            true,
			},
			"remote_address": schema.StringAttribute{
				MarkdownDescription: "Remote outer address.",
				Required:            true,
			},
			"tunnel_local_address": schema.StringAttribute{
				MarkdownDescription: "Local inner address of the tunnel.",
				Required:            true,
			},
			"tunnel_remote_address": schema.StringAttribute{
				MarkdownDescription: "Remote inner address of the tunnel.",
				Required:            true,
			},
			"tunnel_remote_net": schema.Int64Attribute{
				MarkdownDescription: "Prefix length of the inner remote address. Defaults to `32`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(32),
				Validators:          []validator.Int64{int64validator.Between(1, 128)},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the GRE tunnel.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func greDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "GRE tunnel devices (**Interfaces ▸ Devices ▸ GRE**). OPNsense names the device (`greN`) and configures the inner addresses; assign the device to filter its traffic per interface.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the GRE tunnel.",
				Required:            true,
			},
			"device":                dschema.StringAttribute{MarkdownDescription: "Device name, e.g. `gre0`. Assigned by OPNsense when unset.", Computed: true},
			"local_address":         dschema.StringAttribute{MarkdownDescription: "Local outer address: an interface name such as `wan`, a CARP VIP (`wan_vip1`), an IP alias or an IP address.", Computed: true},
			"remote_address":        dschema.StringAttribute{MarkdownDescription: "Remote outer address.", Computed: true},
			"tunnel_local_address":  dschema.StringAttribute{MarkdownDescription: "Local inner address of the tunnel.", Computed: true},
			"tunnel_remote_address": dschema.StringAttribute{MarkdownDescription: "Remote inner address of the tunnel.", Computed: true},
			"tunnel_remote_net":     dschema.Int64Attribute{MarkdownDescription: "Prefix length of the inner remote address.", Computed: true},
			"description":           dschema.StringAttribute{MarkdownDescription: "Optional description here for your reference (not parsed).", Computed: true},
		},
	}
}

func convertGreSchemaToStruct(d *greResourceModel) (*interfaces.Gre, error) {
	return &interfaces.Gre{
		Device:              d.Device.ValueString(),
		LocalAddress:        d.LocalAddress.ValueString(),
		RemoteAddress:       d.RemoteAddress.ValueString(),
		TunnelLocalAddress:  d.TunnelLocalAddress.ValueString(),
		TunnelRemoteAddress: d.TunnelRemoteAddress.ValueString(),
		TunnelRemoteNet:     tools.Int64ToString(d.TunnelRemoteNet.ValueInt64()),
		Description:         d.Description.ValueString(),
	}, nil
}

func convertGreStructToSchema(d *interfaces.Gre) (*greResourceModel, error) {
	return &greResourceModel{
		Device:              tools.StringOrNull(d.Device),
		LocalAddress:        types.StringValue(d.LocalAddress),
		RemoteAddress:       types.StringValue(d.RemoteAddress),
		TunnelLocalAddress:  types.StringValue(d.TunnelLocalAddress),
		TunnelRemoteAddress: types.StringValue(d.TunnelRemoteAddress),
		TunnelRemoteNet:     types.Int64Value(tools.StringToInt64(d.TunnelRemoteNet)),
		Description:         tools.StringOrNull(d.Description),
	}, nil
}
