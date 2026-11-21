package radvd

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/radvd"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type entryResourceModel struct {
	Enabled           types.Bool   `tfsdk:"enabled"`
	Interface         types.String `tfsdk:"interface"`
	BaseInterface     types.String `tfsdk:"base_interface"`
	Mode              types.String `tfsdk:"mode"`
	DeprecatePrefix   types.String `tfsdk:"deprecate_prefix"`
	RemoveAdvOnExit   types.String `tfsdk:"remove_adv_on_exit"`
	RemoveRoute       types.String `tfsdk:"remove_route"`
	Routes            types.Set    `tfsdk:"routes"`
	Rdnss             types.Set    `tfsdk:"rdnss"`
	Dnssl             types.Set    `tfsdk:"dnssl"`
	DNS               types.Bool   `tfsdk:"dns"`
	MinInterval       types.Int64  `tfsdk:"min_interval"`
	MaxInterval       types.Int64  `tfsdk:"max_interval"`
	DnsslLifetime     types.Int64  `tfsdk:"dnssl_lifetime"`
	DefaultLifetime   types.Int64  `tfsdk:"default_lifetime"`
	LinkMTU           types.Int64  `tfsdk:"link_mtu"`
	PreferredLifetime types.Int64  `tfsdk:"preferred_lifetime"`
	SourceAddress     types.String `tfsdk:"source_address"`
	RdnssLifetime     types.Int64  `tfsdk:"rdnss_lifetime"`
	RouteLifetime     types.Int64  `tfsdk:"route_lifetime"`
	ValidLifetime     types.Int64  `tfsdk:"valid_lifetime"`
	DefaultPreference types.String `tfsdk:"default_preference"`
	Nat64Prefix       types.String `tfsdk:"nat64_prefix"`
	CurHopLimit       types.Int64  `tfsdk:"cur_hop_limit"`

	Id types.String `tfsdk:"id"`
}

func entryResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "IPv6 router advertisements (**Services ▸ Router Advertisements**), one entry per interface.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether router advertisements are sent on this interface. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"interface": schema.StringAttribute{
				MarkdownDescription: "Interface to advertise on, e.g. `lan`.",
				Required:            true,
			},
			"base_interface": schema.StringAttribute{
				MarkdownDescription: "Interface whose IPv6 prefix is advertised, for tracked interfaces.",
				Optional:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Advertisement mode: `router` (router only), `unmanaged` (SLAAC), `managed` (DHCPv6), `assist` (SLAAC and DHCPv6), `stateless` (SLAAC and stateless DHCPv6). Defaults to `stateless`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("stateless"),
				Validators:          []validator.String{stringvalidator.OneOf("router", "unmanaged", "managed", "assist", "stateless")},
			},
			"deprecate_prefix": schema.StringAttribute{
				MarkdownDescription: "DeprecatePrefix option: `on`, `off`, or empty for the default. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "on", "off")},
			},
			"remove_adv_on_exit": schema.StringAttribute{
				MarkdownDescription: "RemoveAdvOnExit option: `on`, `off`, or empty for the default. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "on", "off")},
			},
			"remove_route": schema.StringAttribute{
				MarkdownDescription: "RemoveRoute option: `on`, `off`, or empty for the default. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "on", "off")},
			},
			"routes": schema.SetAttribute{
				MarkdownDescription: "Additional routes to advertise, in CIDR notation. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"rdnss": schema.SetAttribute{
				MarkdownDescription: "DNS servers to advertise. Empty advertises the firewall when `dns` is enabled. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"dnssl": schema.SetAttribute{
				MarkdownDescription: "DNS search domains to advertise. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"dns": schema.BoolAttribute{
				MarkdownDescription: "Advertise DNS servers and search domains. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"min_interval": schema.Int64Attribute{
				MarkdownDescription: "Minimum interval between unsolicited advertisements, in seconds. Defaults to `200`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(200),
				Validators:          []validator.Int64{int64validator.AtLeast(3)},
			},
			"max_interval": schema.Int64Attribute{
				MarkdownDescription: "Maximum interval between unsolicited advertisements, in seconds. Defaults to `600`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(600),
				Validators:          []validator.Int64{int64validator.AtLeast(4)},
			},
			"dnssl_lifetime": schema.Int64Attribute{
				MarkdownDescription: "Lifetime of the advertised search domains, in seconds. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.AtLeast(0))},
			},
			"default_lifetime": schema.Int64Attribute{
				MarkdownDescription: "Router lifetime, in seconds. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.AtLeast(0))},
			},
			"link_mtu": schema.Int64Attribute{
				MarkdownDescription: "MTU to advertise. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.AtLeast(0))},
			},
			"preferred_lifetime": schema.Int64Attribute{
				MarkdownDescription: "Preferred lifetime of the prefix, in seconds. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.AtLeast(0))},
			},
			"source_address": schema.StringAttribute{
				MarkdownDescription: "Source address of the advertisements.",
				Optional:            true,
			},
			"rdnss_lifetime": schema.Int64Attribute{
				MarkdownDescription: "Lifetime of the advertised DNS servers, in seconds. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.AtLeast(0))},
			},
			"route_lifetime": schema.Int64Attribute{
				MarkdownDescription: "Lifetime of the advertised routes, in seconds. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.AtLeast(0))},
			},
			"valid_lifetime": schema.Int64Attribute{
				MarkdownDescription: "Valid lifetime of the prefix, in seconds. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.AtLeast(0))},
			},
			"default_preference": schema.StringAttribute{
				MarkdownDescription: "Default router preference. Defaults to `medium`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("medium"),
				Validators:          []validator.String{stringvalidator.OneOf("low", "medium", "high")},
			},
			"nat64_prefix": schema.StringAttribute{
				MarkdownDescription: "NAT64 prefix to advertise (PREF64).",
				Optional:            true,
			},
			"cur_hop_limit": schema.Int64Attribute{
				MarkdownDescription: "Hop limit advertised to clients. Defaults to `64`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(64),
				Validators:          []validator.Int64{int64validator.Between(0, 255)},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the router advertisement.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func entryDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "IPv6 router advertisements (**Services ▸ Router Advertisements**), one entry per interface.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the router advertisement.",
				Required:            true,
			},
			"enabled":            dschema.BoolAttribute{MarkdownDescription: "Whether router advertisements are sent on this interface.", Computed: true},
			"interface":          dschema.StringAttribute{MarkdownDescription: "Interface to advertise on, e.g. `lan`.", Computed: true},
			"base_interface":     dschema.StringAttribute{MarkdownDescription: "Interface whose IPv6 prefix is advertised, for tracked interfaces.", Computed: true},
			"mode":               dschema.StringAttribute{MarkdownDescription: "Advertisement mode: `router` (router only), `unmanaged` (SLAAC), `managed` (DHCPv6), `assist` (SLAAC and DHCPv6), `stateless` (SLAAC and stateless DHCPv6).", Computed: true},
			"deprecate_prefix":   dschema.StringAttribute{MarkdownDescription: "DeprecatePrefix option: `on`, `off`, or empty for the default.", Computed: true},
			"remove_adv_on_exit": dschema.StringAttribute{MarkdownDescription: "RemoveAdvOnExit option: `on`, `off`, or empty for the default.", Computed: true},
			"remove_route":       dschema.StringAttribute{MarkdownDescription: "RemoveRoute option: `on`, `off`, or empty for the default.", Computed: true},
			"routes":             dschema.SetAttribute{MarkdownDescription: "Additional routes to advertise, in CIDR notation.", Computed: true, ElementType: types.StringType},
			"rdnss":              dschema.SetAttribute{MarkdownDescription: "DNS servers to advertise. Empty advertises the firewall when `dns` is enabled.", Computed: true, ElementType: types.StringType},
			"dnssl":              dschema.SetAttribute{MarkdownDescription: "DNS search domains to advertise.", Computed: true, ElementType: types.StringType},
			"dns":                dschema.BoolAttribute{MarkdownDescription: "Advertise DNS servers and search domains.", Computed: true},
			"min_interval":       dschema.Int64Attribute{MarkdownDescription: "Minimum interval between unsolicited advertisements, in seconds.", Computed: true},
			"max_interval":       dschema.Int64Attribute{MarkdownDescription: "Maximum interval between unsolicited advertisements, in seconds.", Computed: true},
			"dnssl_lifetime":     dschema.Int64Attribute{MarkdownDescription: "Lifetime of the advertised search domains, in seconds.", Computed: true},
			"default_lifetime":   dschema.Int64Attribute{MarkdownDescription: "Router lifetime, in seconds.", Computed: true},
			"link_mtu":           dschema.Int64Attribute{MarkdownDescription: "MTU to advertise.", Computed: true},
			"preferred_lifetime": dschema.Int64Attribute{MarkdownDescription: "Preferred lifetime of the prefix, in seconds.", Computed: true},
			"source_address":     dschema.StringAttribute{MarkdownDescription: "Source address of the advertisements.", Computed: true},
			"rdnss_lifetime":     dschema.Int64Attribute{MarkdownDescription: "Lifetime of the advertised DNS servers, in seconds.", Computed: true},
			"route_lifetime":     dschema.Int64Attribute{MarkdownDescription: "Lifetime of the advertised routes, in seconds.", Computed: true},
			"valid_lifetime":     dschema.Int64Attribute{MarkdownDescription: "Valid lifetime of the prefix, in seconds.", Computed: true},
			"default_preference": dschema.StringAttribute{MarkdownDescription: "Default router preference.", Computed: true},
			"nat64_prefix":       dschema.StringAttribute{MarkdownDescription: "NAT64 prefix to advertise (PREF64).", Computed: true},
			"cur_hop_limit":      dschema.Int64Attribute{MarkdownDescription: "Hop limit advertised to clients.", Computed: true},
		},
	}
}

func convertEntrySchemaToStruct(d *entryResourceModel) (*radvd.Entry, error) {
	return &radvd.Entry{
		Enabled:              tools.BoolToString(d.Enabled.ValueBool()),
		Interface:            api.SelectedMap(d.Interface.ValueString()),
		BaseInterface:        api.SelectedMap(d.BaseInterface.ValueString()),
		Mode:                 api.SelectedMap(d.Mode.ValueString()),
		DeprecatePrefix:      api.SelectedMap(d.DeprecatePrefix.ValueString()),
		RemoveAdvOnExit:      api.SelectedMap(d.RemoveAdvOnExit.ValueString()),
		RemoveRoute:          api.SelectedMap(d.RemoveRoute.ValueString()),
		Routes:               api.SelectedMapList(setToSlice(d.Routes)),
		RDNSS:                api.SelectedMapList(setToSlice(d.Rdnss)),
		DNSSL:                api.SelectedMapList(setToSlice(d.Dnssl)),
		DNS:                  tools.BoolToString(d.DNS.ValueBool()),
		MinRtrAdvInterval:    tools.Int64ToString(d.MinInterval.ValueInt64()),
		MaxRtrAdvInterval:    tools.Int64ToString(d.MaxInterval.ValueInt64()),
		AdvDNSSLLifetime:     tools.Int64ToStringNegative(d.DnsslLifetime.ValueInt64()),
		AdvDefaultLifetime:   tools.Int64ToStringNegative(d.DefaultLifetime.ValueInt64()),
		AdvLinkMTU:           tools.Int64ToStringNegative(d.LinkMTU.ValueInt64()),
		AdvPreferredLifetime: tools.Int64ToStringNegative(d.PreferredLifetime.ValueInt64()),
		AdvRASrcAddress:      d.SourceAddress.ValueString(),
		AdvRDNSSLifetime:     tools.Int64ToStringNegative(d.RdnssLifetime.ValueInt64()),
		AdvRouteLifetime:     tools.Int64ToStringNegative(d.RouteLifetime.ValueInt64()),
		AdvValidLifetime:     tools.Int64ToStringNegative(d.ValidLifetime.ValueInt64()),
		AdvDefaultPreference: api.SelectedMap(d.DefaultPreference.ValueString()),
		NAT64Prefix:          d.Nat64Prefix.ValueString(),
		AdvCurHopLimit:       tools.Int64ToString(d.CurHopLimit.ValueInt64()),
	}, nil
}

func convertEntryStructToSchema(d *radvd.Entry) (*entryResourceModel, error) {
	return &entryResourceModel{
		Enabled:           types.BoolValue(tools.StringToBool(d.Enabled)),
		Interface:         types.StringValue(d.Interface.String()),
		BaseInterface:     tools.StringOrNull(d.BaseInterface.String()),
		Mode:              types.StringValue(d.Mode.String()),
		DeprecatePrefix:   types.StringValue(d.DeprecatePrefix.String()),
		RemoveAdvOnExit:   types.StringValue(d.RemoveAdvOnExit.String()),
		RemoveRoute:       types.StringValue(d.RemoveRoute.String()),
		Routes:            tools.StringSliceToSet(d.Routes),
		Rdnss:             tools.StringSliceToSet(d.RDNSS),
		Dnssl:             tools.StringSliceToSet(d.DNSSL),
		DNS:               types.BoolValue(tools.StringToBool(d.DNS)),
		MinInterval:       types.Int64Value(tools.StringToInt64(d.MinRtrAdvInterval)),
		MaxInterval:       types.Int64Value(tools.StringToInt64(d.MaxRtrAdvInterval)),
		DnsslLifetime:     types.Int64Value(tools.StringToInt64(d.AdvDNSSLLifetime)),
		DefaultLifetime:   types.Int64Value(tools.StringToInt64(d.AdvDefaultLifetime)),
		LinkMTU:           types.Int64Value(tools.StringToInt64(d.AdvLinkMTU)),
		PreferredLifetime: types.Int64Value(tools.StringToInt64(d.AdvPreferredLifetime)),
		SourceAddress:     tools.StringOrNull(d.AdvRASrcAddress),
		RdnssLifetime:     types.Int64Value(tools.StringToInt64(d.AdvRDNSSLifetime)),
		RouteLifetime:     types.Int64Value(tools.StringToInt64(d.AdvRouteLifetime)),
		ValidLifetime:     types.Int64Value(tools.StringToInt64(d.AdvValidLifetime)),
		DefaultPreference: types.StringValue(d.AdvDefaultPreference.String()),
		Nat64Prefix:       tools.StringOrNull(d.NAT64Prefix),
		CurHopLimit:       types.Int64Value(tools.StringToInt64(d.AdvCurHopLimit)),
	}, nil
}
