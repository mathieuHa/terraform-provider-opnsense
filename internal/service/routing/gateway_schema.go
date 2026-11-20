package routing

import (
	"regexp"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/routing"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
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

var gatewayNameRegexp = regexp.MustCompile(`^[a-zA-Z0-9_\-]{1,32}$`)

type gatewayResourceModel struct {
	Enabled                   types.Bool   `tfsdk:"enabled"`
	Name                      types.String `tfsdk:"name"`
	Description               types.String `tfsdk:"description"`
	Interface                 types.String `tfsdk:"interface"`
	IPProtocol                types.String `tfsdk:"ip_protocol"`
	Gateway                   types.String `tfsdk:"gateway"`
	DefaultGateway            types.Bool   `tfsdk:"default_gateway"`
	FarGateway                types.Bool   `tfsdk:"far_gateway"`
	MonitorDisable            types.Bool   `tfsdk:"monitor_disable"`
	MonitorNoRoute            types.Bool   `tfsdk:"monitor_no_route"`
	MonitorKillStates         types.Bool   `tfsdk:"monitor_kill_states"`
	MonitorKillStatesPriority types.Bool   `tfsdk:"monitor_kill_states_priority"`
	Monitor                   types.String `tfsdk:"monitor"`
	ForceDown                 types.Bool   `tfsdk:"force_down"`
	NoSync                    types.Bool   `tfsdk:"no_sync"`
	Priority                  types.Int64  `tfsdk:"priority"`
	Weight                    types.Int64  `tfsdk:"weight"`
	LatencyLow                types.Int64  `tfsdk:"latency_low"`
	LatencyHigh               types.Int64  `tfsdk:"latency_high"`
	LossLow                   types.Int64  `tfsdk:"loss_low"`
	LossHigh                  types.Int64  `tfsdk:"loss_high"`
	Interval                  types.Int64  `tfsdk:"interval"`
	TimePeriod                types.Int64  `tfsdk:"time_period"`
	LossInterval              types.Int64  `tfsdk:"loss_interval"`
	DataLength                types.Int64  `tfsdk:"data_length"`

	Id types.String `tfsdk:"id"`
}

func optionalSentinelInt64(description string, validators ...validator.Int64) schema.Int64Attribute {
	return schema.Int64Attribute{
		MarkdownDescription: description + " Defaults to `-1` (OPNsense default).",
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(-1),
		Validators:          append([]validator.Int64{int64validator.AtLeast(-1)}, validators...),
	}
}

func optionalBool(description string, def bool) schema.BoolAttribute {
	return schema.BoolAttribute{
		MarkdownDescription: description,
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(def),
	}
}

func gatewayResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Gateways define the next hop OPNsense uses to reach other networks, and optionally how that next hop is monitored.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": optionalBool("Whether this gateway is enabled. Defaults to `true`.", true),
			"name": schema.StringAttribute{
				MarkdownDescription: "Gateway name. Only `a-zA-Z0-9_-`, up to 32 characters. Routes, gateway groups and policy rules refer to the gateway by this name.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(gatewayNameRegexp, "must only contain a-zA-Z0-9_- and be at most 32 characters"),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"interface": schema.StringAttribute{
				MarkdownDescription: "Interface the gateway is reachable on, e.g. `wan` or `opt3`.",
				Required:            true,
			},
			"ip_protocol": schema.StringAttribute{
				MarkdownDescription: "Address family of the gateway, one of `inet` (IPv4) or `inet6` (IPv6). Defaults to `inet`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("inet"),
				Validators: []validator.String{
					stringvalidator.OneOf("inet", "inet6"),
				},
			},
			"gateway": schema.StringAttribute{
				MarkdownDescription: "IP address of the gateway. Leave unset for a dynamic gateway whose address the interface learns (DHCP, PPPoE, tunnels).",
				Optional:            true,
			},
			"default_gateway":              optionalBool("Whether this gateway is the upstream (default) gateway candidate for its address family. Defaults to `false`.", false),
			"far_gateway":                  optionalBool("When enabled, the gateway may be outside the interface subnet. Defaults to `false`.", false),
			"monitor_disable":              optionalBool("When enabled, gateway monitoring (dpinger) is disabled and the gateway is always considered up. Defaults to `true`.", true),
			"monitor_no_route":             optionalBool("When enabled, no static host route is added for the monitor IP. Defaults to `false`.", false),
			"monitor_kill_states":          optionalBool("When enabled, states using this gateway are killed when it goes down. Defaults to `false`.", false),
			"monitor_kill_states_priority": optionalBool("When enabled, states are also killed when a gateway with a higher priority comes back. Defaults to `false`.", false),
			"monitor": schema.StringAttribute{
				MarkdownDescription: "Alternative IP address to ping for monitoring. When unset, the gateway address itself is monitored.",
				Optional:            true,
			},
			"force_down": optionalBool("When enabled, the gateway is marked down regardless of monitoring. Defaults to `false`.", false),
			"no_sync":    optionalBool("When enabled, this gateway is not synchronised to an HA peer. Defaults to `false`.", false),
			"priority": schema.Int64Attribute{
				MarkdownDescription: "Gateway priority, lower wins when choosing the default gateway. Defaults to `255`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(255),
				Validators:          []validator.Int64{int64validator.Between(0, 255)},
			},
			"weight": schema.Int64Attribute{
				MarkdownDescription: "Weight of this gateway inside a load-balancing gateway group. Defaults to `1`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1),
				Validators:          []validator.Int64{int64validator.Between(1, 10)},
			},
			"latency_low":   optionalSentinelInt64("Low latency threshold in milliseconds."),
			"latency_high":  optionalSentinelInt64("High latency threshold in milliseconds."),
			"loss_low":      optionalSentinelInt64("Low packet loss threshold in percent.", int64validator.AtMost(99)),
			"loss_high":     optionalSentinelInt64("High packet loss threshold in percent.", int64validator.AtMost(100)),
			"interval":      optionalSentinelInt64("Probe interval in milliseconds."),
			"time_period":   optionalSentinelInt64("Time period over which results are averaged, in milliseconds."),
			"loss_interval": optionalSentinelInt64("Time after which a probe without reply counts as lost, in milliseconds."),
			"data_length":   optionalSentinelInt64("Size of the ICMP probe payload in bytes."),
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the gateway.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func gatewayDataSourceSchema() dschema.Schema {
	computedString := func(d string) dschema.StringAttribute {
		return dschema.StringAttribute{MarkdownDescription: d, Computed: true}
	}
	computedBool := func(d string) dschema.BoolAttribute {
		return dschema.BoolAttribute{MarkdownDescription: d, Computed: true}
	}
	computedInt64 := func(d string) dschema.Int64Attribute {
		return dschema.Int64Attribute{MarkdownDescription: d, Computed: true}
	}

	return dschema.Schema{
		MarkdownDescription: "Gateways define the next hop OPNsense uses to reach other networks, and optionally how that next hop is monitored.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the gateway.",
				Required:            true,
			},
			"enabled":                      computedBool("Whether this gateway is enabled."),
			"name":                         computedString("Gateway name."),
			"description":                  computedString("Optional description here for your reference (not parsed)."),
			"interface":                    computedString("Interface the gateway is reachable on."),
			"ip_protocol":                  computedString("Address family of the gateway, `inet` or `inet6`."),
			"gateway":                      computedString("IP address of the gateway, empty for a dynamic gateway."),
			"default_gateway":              computedBool("Whether this gateway is the upstream (default) gateway candidate."),
			"far_gateway":                  computedBool("Whether the gateway may be outside the interface subnet."),
			"monitor_disable":              computedBool("Whether gateway monitoring is disabled."),
			"monitor_no_route":             computedBool("Whether no static host route is added for the monitor IP."),
			"monitor_kill_states":          computedBool("Whether states are killed when the gateway goes down."),
			"monitor_kill_states_priority": computedBool("Whether states are killed when a higher-priority gateway comes back."),
			"monitor":                      computedString("Alternative IP address pinged for monitoring."),
			"force_down":                   computedBool("Whether the gateway is forced down."),
			"no_sync":                      computedBool("Whether this gateway is excluded from HA synchronisation."),
			"priority":                     computedInt64("Gateway priority."),
			"weight":                       computedInt64("Weight inside a load-balancing gateway group."),
			"latency_low":                  computedInt64("Low latency threshold in milliseconds, `-1` when unset."),
			"latency_high":                 computedInt64("High latency threshold in milliseconds, `-1` when unset."),
			"loss_low":                     computedInt64("Low packet loss threshold in percent, `-1` when unset."),
			"loss_high":                    computedInt64("High packet loss threshold in percent, `-1` when unset."),
			"interval":                     computedInt64("Probe interval in milliseconds, `-1` when unset."),
			"time_period":                  computedInt64("Averaging period in milliseconds, `-1` when unset."),
			"loss_interval":                computedInt64("Loss interval in milliseconds, `-1` when unset."),
			"data_length":                  computedInt64("ICMP probe payload size in bytes, `-1` when unset."),
		},
	}
}

func convertGatewaySchemaToStruct(d *gatewayResourceModel) (*routing.Gateway, error) {
	return &routing.Gateway{
		Disabled:                  tools.BoolToString(!d.Enabled.ValueBool()),
		Name:                      d.Name.ValueString(),
		Description:               d.Description.ValueString(),
		Interface:                 api.SelectedMap(d.Interface.ValueString()),
		IPProtocol:                api.SelectedMap(d.IPProtocol.ValueString()),
		Gateway:                   d.Gateway.ValueString(),
		DefaultGateway:            tools.BoolToString(d.DefaultGateway.ValueBool()),
		FarGateway:                tools.BoolToString(d.FarGateway.ValueBool()),
		MonitorDisable:            tools.BoolToString(d.MonitorDisable.ValueBool()),
		MonitorNoRoute:            tools.BoolToString(d.MonitorNoRoute.ValueBool()),
		MonitorKillStates:         tools.BoolToString(d.MonitorKillStates.ValueBool()),
		MonitorKillStatesPriority: tools.BoolToString(d.MonitorKillStatesPriority.ValueBool()),
		Monitor:                   d.Monitor.ValueString(),
		ForceDown:                 tools.BoolToString(d.ForceDown.ValueBool()),
		NoSync:                    tools.BoolToString(d.NoSync.ValueBool()),
		Priority:                  tools.Int64ToString(d.Priority.ValueInt64()),
		Weight:                    tools.Int64ToString(d.Weight.ValueInt64()),
		LatencyLow:                tools.Int64ToStringNegative(d.LatencyLow.ValueInt64()),
		LatencyHigh:               tools.Int64ToStringNegative(d.LatencyHigh.ValueInt64()),
		LossLow:                   tools.Int64ToStringNegative(d.LossLow.ValueInt64()),
		LossHigh:                  tools.Int64ToStringNegative(d.LossHigh.ValueInt64()),
		Interval:                  tools.Int64ToStringNegative(d.Interval.ValueInt64()),
		TimePeriod:                tools.Int64ToStringNegative(d.TimePeriod.ValueInt64()),
		LossInterval:              tools.Int64ToStringNegative(d.LossInterval.ValueInt64()),
		DataLength:                tools.Int64ToStringNegative(d.DataLength.ValueInt64()),
	}, nil
}

func convertGatewayStructToSchema(d *routing.Gateway) (*gatewayResourceModel, error) {
	return &gatewayResourceModel{
		Enabled:                   types.BoolValue(!tools.StringToBool(d.Disabled)),
		Name:                      types.StringValue(d.Name),
		Description:               tools.StringOrNull(d.Description),
		Interface:                 types.StringValue(d.Interface.String()),
		IPProtocol:                types.StringValue(d.IPProtocol.String()),
		Gateway:                   tools.StringOrNull(d.Gateway),
		DefaultGateway:            types.BoolValue(tools.StringToBool(d.DefaultGateway)),
		FarGateway:                types.BoolValue(tools.StringToBool(d.FarGateway)),
		MonitorDisable:            types.BoolValue(tools.StringToBool(d.MonitorDisable)),
		MonitorNoRoute:            types.BoolValue(tools.StringToBool(d.MonitorNoRoute)),
		MonitorKillStates:         types.BoolValue(tools.StringToBool(d.MonitorKillStates)),
		MonitorKillStatesPriority: types.BoolValue(tools.StringToBool(d.MonitorKillStatesPriority)),
		Monitor:                   tools.StringOrNull(d.Monitor),
		ForceDown:                 types.BoolValue(tools.StringToBool(d.ForceDown)),
		NoSync:                    types.BoolValue(tools.StringToBool(d.NoSync)),
		Priority:                  types.Int64Value(tools.StringToInt64(d.Priority)),
		Weight:                    types.Int64Value(tools.StringToInt64(d.Weight)),
		LatencyLow:                types.Int64Value(tools.StringToInt64(d.LatencyLow)),
		LatencyHigh:               types.Int64Value(tools.StringToInt64(d.LatencyHigh)),
		LossLow:                   types.Int64Value(tools.StringToInt64(d.LossLow)),
		LossHigh:                  types.Int64Value(tools.StringToInt64(d.LossHigh)),
		Interval:                  types.Int64Value(tools.StringToInt64(d.Interval)),
		TimePeriod:                types.Int64Value(tools.StringToInt64(d.TimePeriod)),
		LossInterval:              types.Int64Value(tools.StringToInt64(d.LossInterval)),
		DataLength:                types.Int64Value(tools.StringToInt64(d.DataLength)),
	}, nil
}
