package dnsmasq

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dnsmasq"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
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

type dhcpRangeResourceModel struct {
	Interface        types.String `tfsdk:"interface"`
	Tag              types.String `tfsdk:"tag"`
	StartAddress     types.String `tfsdk:"start_address"`
	EndAddress       types.String `tfsdk:"end_address"`
	SubnetMask       types.String `tfsdk:"subnet_mask"`
	Constructor      types.String `tfsdk:"constructor"`
	Mode             types.Set    `tfsdk:"mode"`
	PrefixLength     types.Int64  `tfsdk:"prefix_length"`
	LeaseTime        types.Int64  `tfsdk:"lease_time"`
	DomainType       types.String `tfsdk:"domain_type"`
	Domain           types.String `tfsdk:"domain"`
	NoSync           types.Bool   `tfsdk:"no_sync"`
	RaMode           types.Set    `tfsdk:"ra_mode"`
	RaPriority       types.String `tfsdk:"ra_priority"`
	RaMTU            types.Int64  `tfsdk:"ra_mtu"`
	RaInterval       types.Int64  `tfsdk:"ra_interval"`
	RaRouterLifetime types.Int64  `tfsdk:"ra_router_lifetime"`
	Description      types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func dhcpRangeResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Dnsmasq DHCP ranges (**Services ▸ Dnsmasq DNS & DHCP ▸ DHCP ranges**) for DHCPv4 pools and DHCPv6 / router advertisements.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"interface": schema.StringAttribute{
				MarkdownDescription: "Interface the range is served on, e.g. `lan`.",
				Optional:            true,
			},
			"tag": schema.StringAttribute{
				MarkdownDescription: "UUID of the tag set on clients that get an address from this range.",
				Optional:            true,
			},
			"start_address": schema.StringAttribute{
				MarkdownDescription: "First address of the range.",
				Required:            true,
			},
			"end_address": schema.StringAttribute{
				MarkdownDescription: "Last address of the range.",
				Optional:            true,
			},
			"subnet_mask": schema.StringAttribute{
				MarkdownDescription: "Subnet mask handed out, when it differs from the interface.",
				Optional:            true,
			},
			"constructor": schema.StringAttribute{
				MarkdownDescription: "Interface whose IPv6 prefix the range is built from (DHCPv6 only).",
				Optional:            true,
			},
			"mode": schema.SetAttribute{
				MarkdownDescription: "Range flags; `static` only serves static reservations. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
				Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("static"))},
			},
			"prefix_length": schema.Int64Attribute{
				MarkdownDescription: "Prefix length for DHCPv6 ranges. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.Between(1, 64))},
			},
			"lease_time": schema.Int64Attribute{
				MarkdownDescription: "Lease time in seconds. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.AtLeast(0))},
			},
			"domain_type": schema.StringAttribute{
				MarkdownDescription: "Whether `domain` applies to this range or to the whole interface. Defaults to `range`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("range"),
				Validators:          []validator.String{stringvalidator.OneOf("range", "interface")},
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Domain handed out to clients of this range.",
				Optional:            true,
			},
			"no_sync": schema.BoolAttribute{
				MarkdownDescription: "When enabled, this range is not synchronised to an HA peer. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ra_mode": schema.SetAttribute{
				MarkdownDescription: "Router advertisement flags. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
				Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("ra-only", "slaac", "ra-names", "ra-stateless", "ra-advrouter", "off-link"))},
			},
			"ra_priority": schema.StringAttribute{
				MarkdownDescription: "Router advertisement priority; unset means normal.",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("high", "low")},
			},
			"ra_mtu": schema.Int64Attribute{
				MarkdownDescription: "MTU announced in router advertisements. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.Between(1280, 65535))},
			},
			"ra_interval": schema.Int64Attribute{
				MarkdownDescription: "Router advertisement interval in seconds. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.Between(4, 1800))},
			},
			"ra_router_lifetime": schema.Int64Attribute{
				MarkdownDescription: "Router lifetime announced, in seconds. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.Between(0, 9000))},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the dnsmasq DHCP range.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func dhcpRangeDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Dnsmasq DHCP ranges (**Services ▸ Dnsmasq DNS & DHCP ▸ DHCP ranges**) for DHCPv4 pools and DHCPv6 / router advertisements.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the dnsmasq DHCP range.",
				Required:            true,
			},
			"interface":          dschema.StringAttribute{MarkdownDescription: "Interface the range is served on, e.g. `lan`.", Computed: true},
			"tag":                dschema.StringAttribute{MarkdownDescription: "UUID of the tag set on clients that get an address from this range.", Computed: true},
			"start_address":      dschema.StringAttribute{MarkdownDescription: "First address of the range.", Computed: true},
			"end_address":        dschema.StringAttribute{MarkdownDescription: "Last address of the range.", Computed: true},
			"subnet_mask":        dschema.StringAttribute{MarkdownDescription: "Subnet mask handed out, when it differs from the interface.", Computed: true},
			"constructor":        dschema.StringAttribute{MarkdownDescription: "Interface whose IPv6 prefix the range is built from (DHCPv6 only).", Computed: true},
			"mode":               dschema.SetAttribute{MarkdownDescription: "Range flags; `static` only serves static reservations.", Computed: true, ElementType: types.StringType},
			"prefix_length":      dschema.Int64Attribute{MarkdownDescription: "Prefix length for DHCPv6 ranges.", Computed: true},
			"lease_time":         dschema.Int64Attribute{MarkdownDescription: "Lease time in seconds.", Computed: true},
			"domain_type":        dschema.StringAttribute{MarkdownDescription: "Whether `domain` applies to this range or to the whole interface.", Computed: true},
			"domain":             dschema.StringAttribute{MarkdownDescription: "Domain handed out to clients of this range.", Computed: true},
			"no_sync":            dschema.BoolAttribute{MarkdownDescription: "When enabled, this range is not synchronised to an HA peer.", Computed: true},
			"ra_mode":            dschema.SetAttribute{MarkdownDescription: "Router advertisement flags.", Computed: true, ElementType: types.StringType},
			"ra_priority":        dschema.StringAttribute{MarkdownDescription: "Router advertisement priority; unset means normal.", Computed: true},
			"ra_mtu":             dschema.Int64Attribute{MarkdownDescription: "MTU announced in router advertisements.", Computed: true},
			"ra_interval":        dschema.Int64Attribute{MarkdownDescription: "Router advertisement interval in seconds.", Computed: true},
			"ra_router_lifetime": dschema.Int64Attribute{MarkdownDescription: "Router lifetime announced, in seconds.", Computed: true},
			"description":        dschema.StringAttribute{MarkdownDescription: "Optional description here for your reference (not parsed).", Computed: true},
		},
	}
}

func convertRangeSchemaToStruct(d *dhcpRangeResourceModel) (*dnsmasq.Range, error) {
	return &dnsmasq.Range{
		Interface:        api.SelectedMap(d.Interface.ValueString()),
		Tag:              api.SelectedMap(d.Tag.ValueString()),
		StartAddress:     d.StartAddress.ValueString(),
		EndAddress:       d.EndAddress.ValueString(),
		SubnetMask:       d.SubnetMask.ValueString(),
		Constructor:      api.SelectedMap(d.Constructor.ValueString()),
		Mode:             api.SelectedMapList(setToSlice(d.Mode)),
		PrefixLength:     tools.Int64ToStringNegative(d.PrefixLength.ValueInt64()),
		LeaseTime:        tools.Int64ToStringNegative(d.LeaseTime.ValueInt64()),
		DomainType:       api.SelectedMap(d.DomainType.ValueString()),
		Domain:           d.Domain.ValueString(),
		DisableHASync:    tools.BoolToString(d.NoSync.ValueBool()),
		RaMode:           api.SelectedMapList(setToSlice(d.RaMode)),
		RaPriority:       api.SelectedMap(d.RaPriority.ValueString()),
		RaMTU:            tools.Int64ToStringNegative(d.RaMTU.ValueInt64()),
		RaInterval:       tools.Int64ToStringNegative(d.RaInterval.ValueInt64()),
		RaRouterLifetime: tools.Int64ToStringNegative(d.RaRouterLifetime.ValueInt64()),
		Description:      d.Description.ValueString(),
	}, nil
}

func convertRangeStructToSchema(d *dnsmasq.Range) (*dhcpRangeResourceModel, error) {
	return &dhcpRangeResourceModel{
		Interface:        tools.StringOrNull(d.Interface.String()),
		Tag:              tools.StringOrNull(d.Tag.String()),
		StartAddress:     types.StringValue(d.StartAddress),
		EndAddress:       tools.StringOrNull(d.EndAddress),
		SubnetMask:       tools.StringOrNull(d.SubnetMask),
		Constructor:      tools.StringOrNull(d.Constructor.String()),
		Mode:             tools.StringSliceToSet(d.Mode),
		PrefixLength:     types.Int64Value(tools.StringToInt64(d.PrefixLength)),
		LeaseTime:        types.Int64Value(tools.StringToInt64(d.LeaseTime)),
		DomainType:       types.StringValue(d.DomainType.String()),
		Domain:           tools.StringOrNull(d.Domain),
		NoSync:           types.BoolValue(tools.StringToBool(d.DisableHASync)),
		RaMode:           tools.StringSliceToSet(d.RaMode),
		RaPriority:       tools.StringOrNull(d.RaPriority.String()),
		RaMTU:            types.Int64Value(tools.StringToInt64(d.RaMTU)),
		RaInterval:       types.Int64Value(tools.StringToInt64(d.RaInterval)),
		RaRouterLifetime: types.Int64Value(tools.StringToInt64(d.RaRouterLifetime)),
		Description:      tools.StringOrNull(d.Description),
	}, nil
}
