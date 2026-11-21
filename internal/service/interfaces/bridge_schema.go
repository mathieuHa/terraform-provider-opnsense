package interfaces

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/interfaces"
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

type bridgeResourceModel struct {
	Device         types.String `tfsdk:"device"`
	Members        types.Set    `tfsdk:"members"`
	LinkLocal      types.Bool   `tfsdk:"link_local"`
	EnableStp      types.Bool   `tfsdk:"enable_stp"`
	StpProtocol    types.String `tfsdk:"stp_protocol"`
	StpInterfaces  types.Set    `tfsdk:"stp_interfaces"`
	MaxAge         types.Int64  `tfsdk:"max_age"`
	ForwardDelay   types.Int64  `tfsdk:"forward_delay"`
	HoldCount      types.Int64  `tfsdk:"hold_count"`
	MaxAddresses   types.Int64  `tfsdk:"max_addresses"`
	AddressTimeout types.Int64  `tfsdk:"address_timeout"`
	SpanInterface  types.String `tfsdk:"span_interface"`
	EdgePorts      types.Set    `tfsdk:"edge_ports"`
	AutoEdgePorts  types.Set    `tfsdk:"auto_edge_ports"`
	PtpPorts       types.Set    `tfsdk:"ptp_ports"`
	AutoPtpPorts   types.Set    `tfsdk:"auto_ptp_ports"`
	StaticPorts    types.Set    `tfsdk:"static_ports"`
	PrivatePorts   types.Set    `tfsdk:"private_ports"`
	Description    types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func bridgeResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Bridge devices (**Interfaces ▸ Devices ▸ Bridge**) join several interfaces into one layer-2 segment. Assign the resulting `device` to use it as an interface. Never bridge the interface you manage OPNsense through without console access.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "Bridge device name assigned by OPNsense, e.g. `bridge0`.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"members": schema.SetAttribute{
				MarkdownDescription: "Interfaces joined into the bridge, e.g. `[\"opt1\", \"opt2\"]`.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"link_local": schema.BoolAttribute{
				MarkdownDescription: "When enabled, an IPv6 link-local address is set on the bridge. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"enable_stp": schema.BoolAttribute{
				MarkdownDescription: "When enabled, the spanning tree protocol runs on the bridge. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"stp_protocol": schema.StringAttribute{
				MarkdownDescription: "Spanning tree protocol. Defaults to `rstp`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("rstp"),
				Validators:          []validator.String{stringvalidator.OneOf("rstp", "stp")},
			},
			"stp_interfaces": schema.SetAttribute{
				MarkdownDescription: "Members on which spanning tree is enabled. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"max_age": schema.Int64Attribute{
				MarkdownDescription: "Spanning tree max age, in seconds. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.Between(6, 40))},
			},
			"forward_delay": schema.Int64Attribute{
				MarkdownDescription: "Spanning tree forward delay, in seconds. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.Between(4, 30))},
			},
			"hold_count": schema.Int64Attribute{
				MarkdownDescription: "Spanning tree transmit hold count. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.Between(1, 10))},
			},
			"max_addresses": schema.Int64Attribute{
				MarkdownDescription: "Size of the bridge address cache. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.AtLeast(0))},
			},
			"address_timeout": schema.Int64Attribute{
				MarkdownDescription: "Timeout of address cache entries, in seconds. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.AtLeast(0))},
			},
			"span_interface": schema.StringAttribute{
				MarkdownDescription: "Interface that receives a copy of all bridged traffic (span port).",
				Optional:            true,
			},
			"edge_ports": schema.SetAttribute{
				MarkdownDescription: "Members marked as edge ports. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"auto_edge_ports": schema.SetAttribute{
				MarkdownDescription: "Members with automatic edge detection. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"ptp_ports": schema.SetAttribute{
				MarkdownDescription: "Members marked as point-to-point links. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"auto_ptp_ports": schema.SetAttribute{
				MarkdownDescription: "Members with automatic point-to-point detection. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"static_ports": schema.SetAttribute{
				MarkdownDescription: "Members marked as sticky (addresses never age out). Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"private_ports": schema.SetAttribute{
				MarkdownDescription: "Members that cannot exchange traffic with each other. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the bridge.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func bridgeDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Bridge devices (**Interfaces ▸ Devices ▸ Bridge**) join several interfaces into one layer-2 segment. Assign the resulting `device` to use it as an interface. Never bridge the interface you manage OPNsense through without console access.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the bridge.",
				Required:            true,
			},
			"device":          dschema.StringAttribute{MarkdownDescription: "Bridge device name assigned by OPNsense, e.g. `bridge0`.", Computed: true},
			"members":         dschema.SetAttribute{MarkdownDescription: "Interfaces joined into the bridge, e.g. `[\"opt1\", \"opt2\"]`.", Computed: true, ElementType: types.StringType},
			"link_local":      dschema.BoolAttribute{MarkdownDescription: "When enabled, an IPv6 link-local address is set on the bridge.", Computed: true},
			"enable_stp":      dschema.BoolAttribute{MarkdownDescription: "When enabled, the spanning tree protocol runs on the bridge.", Computed: true},
			"stp_protocol":    dschema.StringAttribute{MarkdownDescription: "Spanning tree protocol.", Computed: true},
			"stp_interfaces":  dschema.SetAttribute{MarkdownDescription: "Members on which spanning tree is enabled.", Computed: true, ElementType: types.StringType},
			"max_age":         dschema.Int64Attribute{MarkdownDescription: "Spanning tree max age, in seconds.", Computed: true},
			"forward_delay":   dschema.Int64Attribute{MarkdownDescription: "Spanning tree forward delay, in seconds.", Computed: true},
			"hold_count":      dschema.Int64Attribute{MarkdownDescription: "Spanning tree transmit hold count.", Computed: true},
			"max_addresses":   dschema.Int64Attribute{MarkdownDescription: "Size of the bridge address cache.", Computed: true},
			"address_timeout": dschema.Int64Attribute{MarkdownDescription: "Timeout of address cache entries, in seconds.", Computed: true},
			"span_interface":  dschema.StringAttribute{MarkdownDescription: "Interface that receives a copy of all bridged traffic (span port).", Computed: true},
			"edge_ports":      dschema.SetAttribute{MarkdownDescription: "Members marked as edge ports.", Computed: true, ElementType: types.StringType},
			"auto_edge_ports": dschema.SetAttribute{MarkdownDescription: "Members with automatic edge detection.", Computed: true, ElementType: types.StringType},
			"ptp_ports":       dschema.SetAttribute{MarkdownDescription: "Members marked as point-to-point links.", Computed: true, ElementType: types.StringType},
			"auto_ptp_ports":  dschema.SetAttribute{MarkdownDescription: "Members with automatic point-to-point detection.", Computed: true, ElementType: types.StringType},
			"static_ports":    dschema.SetAttribute{MarkdownDescription: "Members marked as sticky (addresses never age out).", Computed: true, ElementType: types.StringType},
			"private_ports":   dschema.SetAttribute{MarkdownDescription: "Members that cannot exchange traffic with each other.", Computed: true, ElementType: types.StringType},
			"description":     dschema.StringAttribute{MarkdownDescription: "Optional description here for your reference (not parsed).", Computed: true},
		},
	}
}

func convertBridgeSchemaToStruct(d *bridgeResourceModel) (*interfaces.Bridge, error) {
	return &interfaces.Bridge{
		Device:         d.Device.ValueString(),
		Members:        api.SelectedMapList(setToSlice(d.Members)),
		LinkLocal:      tools.BoolToString(d.LinkLocal.ValueBool()),
		EnableStp:      tools.BoolToString(d.EnableStp.ValueBool()),
		StpProto:       api.SelectedMap(d.StpProtocol.ValueString()),
		StpInterfaces:  api.SelectedMapList(setToSlice(d.StpInterfaces)),
		MaxAge:         tools.Int64ToStringNegative(d.MaxAge.ValueInt64()),
		ForwardDelay:   tools.Int64ToStringNegative(d.ForwardDelay.ValueInt64()),
		HoldCount:      tools.Int64ToStringNegative(d.HoldCount.ValueInt64()),
		MaxAddresses:   tools.Int64ToStringNegative(d.MaxAddresses.ValueInt64()),
		AddressTimeout: tools.Int64ToStringNegative(d.AddressTimeout.ValueInt64()),
		SpanInterface:  api.SelectedMap(d.SpanInterface.ValueString()),
		EdgePorts:      api.SelectedMapList(setToSlice(d.EdgePorts)),
		AutoEdgePorts:  api.SelectedMapList(setToSlice(d.AutoEdgePorts)),
		PtpPorts:       api.SelectedMapList(setToSlice(d.PtpPorts)),
		AutoPtpPorts:   api.SelectedMapList(setToSlice(d.AutoPtpPorts)),
		StaticPorts:    api.SelectedMapList(setToSlice(d.StaticPorts)),
		PrivatePorts:   api.SelectedMapList(setToSlice(d.PrivatePorts)),
		Description:    d.Description.ValueString(),
	}, nil
}

func convertBridgeStructToSchema(d *interfaces.Bridge) (*bridgeResourceModel, error) {
	return &bridgeResourceModel{
		Device:         tools.StringOrNull(d.Device),
		Members:        tools.StringSliceToSet(d.Members),
		LinkLocal:      types.BoolValue(tools.StringToBool(d.LinkLocal)),
		EnableStp:      types.BoolValue(tools.StringToBool(d.EnableStp)),
		StpProtocol:    types.StringValue(d.StpProto.String()),
		StpInterfaces:  tools.StringSliceToSet(d.StpInterfaces),
		MaxAge:         types.Int64Value(tools.StringToInt64(d.MaxAge)),
		ForwardDelay:   types.Int64Value(tools.StringToInt64(d.ForwardDelay)),
		HoldCount:      types.Int64Value(tools.StringToInt64(d.HoldCount)),
		MaxAddresses:   types.Int64Value(tools.StringToInt64(d.MaxAddresses)),
		AddressTimeout: types.Int64Value(tools.StringToInt64(d.AddressTimeout)),
		SpanInterface:  tools.StringOrNull(d.SpanInterface.String()),
		EdgePorts:      tools.StringSliceToSet(d.EdgePorts),
		AutoEdgePorts:  tools.StringSliceToSet(d.AutoEdgePorts),
		PtpPorts:       tools.StringSliceToSet(d.PtpPorts),
		AutoPtpPorts:   tools.StringSliceToSet(d.AutoPtpPorts),
		StaticPorts:    tools.StringSliceToSet(d.StaticPorts),
		PrivatePorts:   tools.StringSliceToSet(d.PrivatePorts),
		Description:    tools.StringOrNull(d.Description),
	}, nil
}
