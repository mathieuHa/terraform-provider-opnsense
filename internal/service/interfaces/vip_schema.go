package interfaces

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/interfaces"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/browningluke/terraform-provider-opnsense/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// vipResourceModel describes the resource data model.
type vipResourceModel struct {
	Mode        types.String `tfsdk:"mode"`
	Interface   types.String `tfsdk:"interface"`
	Network     types.String `tfsdk:"network"`
	Gateway     types.String `tfsdk:"gateway"`
	Description types.String `tfsdk:"description"`
	VHID        types.Int64  `tfsdk:"vhid"`
	AdvBase     types.Int64  `tfsdk:"advbase"`
	AdvSkew     types.Int64  `tfsdk:"advskew"`
	Password    types.String `tfsdk:"password"`
	Peer        types.String `tfsdk:"peer"`
	Peer6       types.String `tfsdk:"peer6"`
	NoSync      types.Bool   `tfsdk:"no_sync"`
	NoBind      types.Bool   `tfsdk:"no_bind"`
	NoExpand    types.Bool   `tfsdk:"no_expand"`
	Id          types.String `tfsdk:"id"`
}

func vipResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Virtual IPs allow an OPNsense firewall to assign multiple IP addresses to the same network interface.",

		Attributes: map[string]schema.Attribute{
			"mode": schema.StringAttribute{
				MarkdownDescription: "Mode of the VIP. One of `ipalias`, `carp` or `proxyarp`. `proxyarp` cannot be bound to by anything running on the firewall, such as IPsec, OpenVPN, etc. A `carp` VIP needs `vhid` and `password`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("ipalias"),
				Validators: []validator.String{
					stringvalidator.OneOf("ipalias", "carp", "proxyarp"),
				},
			},
			"interface": schema.StringAttribute{
				MarkdownDescription: "Choose which interface this VIP applies to.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("wan"),
			},
			"network": schema.StringAttribute{
				MarkdownDescription: "Provide an address and subnet to use. (e.g 192.168.0.1/24)",
				Required:            true,
				Validators: []validator.String{
					validators.CIDR(),
				},
			},
			"gateway": schema.StringAttribute{
				MarkdownDescription: "For some interface types a gateway is required to configure an IP Alias (ppp/pppoe/tun), leave this field empty for all other interface types.",
				Optional:            true,
				Validators: []validator.String{
					validators.IpOrCIDR(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"vhid": schema.Int64Attribute{
				MarkdownDescription: "CARP virtual host ID, shared by all members of the CARP group (`carp` mode).",
				Optional:            true,
				Validators:          []validator.Int64{int64validator.Between(1, 255)},
			},
			"advbase": schema.Int64Attribute{
				MarkdownDescription: "CARP advertisement base in seconds (`carp` mode). Defaults to `1`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1),
				Validators:          []validator.Int64{int64validator.Between(1, 254)},
			},
			"advskew": schema.Int64Attribute{
				MarkdownDescription: "CARP advertisement skew (`carp` mode). The member with the lowest skew becomes master. Defaults to `0`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators:          []validator.Int64{int64validator.Between(0, 254)},
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "CARP group password, identical on all members (`carp` mode).",
				Optional:            true,
				Sensitive:           true,
			},
			"peer": schema.StringAttribute{
				MarkdownDescription: "Send CARP advertisements by unicast to this IPv4 peer instead of multicast (`carp` mode).",
				Optional:            true,
			},
			"peer6": schema.StringAttribute{
				MarkdownDescription: "Send CARP advertisements by unicast to this IPv6 peer instead of multicast (`carp` mode).",
				Optional:            true,
			},
			"no_sync": schema.BoolAttribute{
				MarkdownDescription: "Exclude this VIP from HA configuration synchronisation. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"no_bind": schema.BoolAttribute{
				MarkdownDescription: "Do not let services bind to this address, e.g. for addresses only used for NAT. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"no_expand": schema.BoolAttribute{
				MarkdownDescription: "Do not expand the subnet of this VIP into the interface network in aliases and rules. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the VIP.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func vipDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Virtual IPs allow an OPNsense firewall to assign multiple IP addresses to the same network interface.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Mode of the VIP. One of `ipalias`, `carp`, or `proxyarp`. `proxyarp` cannot be bound to by anything running on the firewall, such as IPsec, OpenVPN, etc. In most cases an `ipalias` should be used.",
				Computed:            true,
			},
			"interface": schema.StringAttribute{
				MarkdownDescription: "Choose which interface this VIP applies to.",
				Computed:            true,
			},
			"network": schema.StringAttribute{
				MarkdownDescription: "Provide an address and subnet to use. (e.g 192.168.0.1/24)",
				Computed:            true,
			},
			"gateway": schema.StringAttribute{
				MarkdownDescription: "For some interface types a gateway is required to configure an IP Alias (ppp/pppoe/tun), leave this field empty for all other interface types.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Computed:            true,
			},
			"vhid":      dschema.Int64Attribute{MarkdownDescription: "CARP virtual host ID.", Computed: true},
			"advbase":   dschema.Int64Attribute{MarkdownDescription: "CARP advertisement base in seconds.", Computed: true},
			"advskew":   dschema.Int64Attribute{MarkdownDescription: "CARP advertisement skew.", Computed: true},
			"password":  dschema.StringAttribute{MarkdownDescription: "CARP group password.", Computed: true, Sensitive: true},
			"peer":      dschema.StringAttribute{MarkdownDescription: "Unicast IPv4 CARP peer.", Computed: true},
			"peer6":     dschema.StringAttribute{MarkdownDescription: "Unicast IPv6 CARP peer.", Computed: true},
			"no_sync":   dschema.BoolAttribute{MarkdownDescription: "Whether the VIP is excluded from HA synchronisation.", Computed: true},
			"no_bind":   dschema.BoolAttribute{MarkdownDescription: "Whether services may not bind to this address.", Computed: true},
			"no_expand": dschema.BoolAttribute{MarkdownDescription: "Whether the subnet is not expanded in aliases and rules.", Computed: true},
		},
	}
}

func convertVipSchemaToStruct(d *vipResourceModel) (*interfaces.Vip, error) {
	return &interfaces.Vip{
		Description: d.Description.ValueString(),
		Mode:        api.SelectedMap(d.Mode.ValueString()),
		Interface:   api.SelectedMap(d.Interface.ValueString()),
		Network:     d.Network.ValueString(),
		Gateway:     d.Gateway.ValueString(),
		VHID:        vipIntOrEmpty(d.VHID),
		AdvBase:     tools.Int64ToString(d.AdvBase.ValueInt64()),
		AdvSkew:     tools.Int64ToString(d.AdvSkew.ValueInt64()),
		Password:    d.Password.ValueString(),
		Peer:        d.Peer.ValueString(),
		Peer6:       d.Peer6.ValueString(),
		NoSync:      tools.BoolToString(d.NoSync.ValueBool()),
		NoBind:      tools.BoolToString(d.NoBind.ValueBool()),
		NoExpand:    tools.BoolToString(d.NoExpand.ValueBool()),
	}, nil
}

func convertVipStructToSchema(d *interfaces.Vip) (*vipResourceModel, error) {
	return &vipResourceModel{
		Mode:        types.StringValue(d.Mode.String()),
		Interface:   types.StringValue(d.Interface.String()),
		Network:     types.StringValue(d.Network),
		Gateway:     tools.StringOrNull(d.Gateway),
		Description: tools.StringOrNull(d.Description),
		VHID:        tools.StringToInt64Null(d.VHID),
		AdvBase:     types.Int64Value(tools.StringToInt64(d.AdvBase)),
		AdvSkew:     types.Int64Value(tools.StringToInt64(d.AdvSkew)),
		Password:    tools.StringOrNull(d.Password),
		Peer:        tools.StringOrNull(d.Peer),
		Peer6:       tools.StringOrNull(d.Peer6),
		NoSync:      types.BoolValue(tools.StringToBool(d.NoSync)),
		NoBind:      types.BoolValue(tools.StringToBool(d.NoBind)),
		NoExpand:    types.BoolValue(tools.StringToBool(d.NoExpand)),
	}, nil
}

func vipIntOrEmpty(v types.Int64) string {
	if v.IsNull() || v.IsUnknown() {
		return ""
	}
	return tools.Int64ToString(v.ValueInt64())
}
