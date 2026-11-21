package interfaces

import (
	"github.com/browningluke/opnsense-go/pkg/interfaces"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type gifResourceModel struct {
	Device               types.String `tfsdk:"device"`
	LocalAddress         types.String `tfsdk:"local_address"`
	RemoteAddress        types.String `tfsdk:"remote_address"`
	TunnelLocalAddress   types.String `tfsdk:"tunnel_local_address"`
	TunnelRemoteAddress  types.String `tfsdk:"tunnel_remote_address"`
	TunnelRemoteNet      types.Int64  `tfsdk:"tunnel_remote_net"`
	Description          types.String `tfsdk:"description"`
	EcnFriendly          types.Bool   `tfsdk:"ecn_friendly"`
	DisableIngressFilter types.Bool   `tfsdk:"disable_ingress_filter"`

	Id types.String `tfsdk:"id"`
}

func gifResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "GIF tunnel devices (**Interfaces ▸ Devices ▸ GIF**). OPNsense names the device (`gifN`) and configures the inner addresses; assign the device to filter its traffic per interface.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "Device name, e.g. `gif0`. Assigned by OPNsense when unset.",
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
			"ecn_friendly": schema.BoolAttribute{
				MarkdownDescription: "ECN friendly behaviour (`link1`). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"disable_ingress_filter": schema.BoolAttribute{
				MarkdownDescription: "Disable ingress filtering of the outer source address (`link2`). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the GIF tunnel.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func gifDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "GIF tunnel devices (**Interfaces ▸ Devices ▸ GIF**). OPNsense names the device (`gifN`) and configures the inner addresses; assign the device to filter its traffic per interface.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the GIF tunnel.",
				Required:            true,
			},
			"device":                 dschema.StringAttribute{MarkdownDescription: "Device name, e.g. `gif0`. Assigned by OPNsense when unset.", Computed: true},
			"local_address":          dschema.StringAttribute{MarkdownDescription: "Local outer address: an interface name such as `wan`, a CARP VIP (`wan_vip1`), an IP alias or an IP address.", Computed: true},
			"remote_address":         dschema.StringAttribute{MarkdownDescription: "Remote outer address.", Computed: true},
			"tunnel_local_address":   dschema.StringAttribute{MarkdownDescription: "Local inner address of the tunnel.", Computed: true},
			"tunnel_remote_address":  dschema.StringAttribute{MarkdownDescription: "Remote inner address of the tunnel.", Computed: true},
			"tunnel_remote_net":      dschema.Int64Attribute{MarkdownDescription: "Prefix length of the inner remote address.", Computed: true},
			"description":            dschema.StringAttribute{MarkdownDescription: "Optional description here for your reference (not parsed).", Computed: true},
			"ecn_friendly":           dschema.BoolAttribute{MarkdownDescription: "ECN friendly behaviour (`link1`).", Computed: true},
			"disable_ingress_filter": dschema.BoolAttribute{MarkdownDescription: "Disable ingress filtering of the outer source address (`link2`).", Computed: true},
		},
	}
}

func convertGifSchemaToStruct(d *gifResourceModel) (*interfaces.Gif, error) {
	return &interfaces.Gif{
		Device:              d.Device.ValueString(),
		LocalAddress:        d.LocalAddress.ValueString(),
		RemoteAddress:       d.RemoteAddress.ValueString(),
		TunnelLocalAddress:  d.TunnelLocalAddress.ValueString(),
		TunnelRemoteAddress: d.TunnelRemoteAddress.ValueString(),
		TunnelRemoteNet:     tools.Int64ToString(d.TunnelRemoteNet.ValueInt64()),
		Description:         d.Description.ValueString(),
		Link1:               tools.BoolToString(d.EcnFriendly.ValueBool()),
		Link2:               tools.BoolToString(d.DisableIngressFilter.ValueBool()),
	}, nil
}

func convertGifStructToSchema(d *interfaces.Gif) (*gifResourceModel, error) {
	return &gifResourceModel{
		Device:               tools.StringOrNull(d.Device),
		LocalAddress:         types.StringValue(d.LocalAddress),
		RemoteAddress:        types.StringValue(d.RemoteAddress),
		TunnelLocalAddress:   types.StringValue(d.TunnelLocalAddress),
		TunnelRemoteAddress:  types.StringValue(d.TunnelRemoteAddress),
		TunnelRemoteNet:      types.Int64Value(tools.StringToInt64(d.TunnelRemoteNet)),
		Description:          tools.StringOrNull(d.Description),
		EcnFriendly:          types.BoolValue(tools.StringToBool(d.Link1)),
		DisableIngressFilter: types.BoolValue(tools.StringToBool(d.Link2)),
	}, nil
}
