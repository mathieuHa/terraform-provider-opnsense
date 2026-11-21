package kea

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/kea"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type dhcpv6SettingsResourceModel struct {
	Enabled                     types.Bool   `tfsdk:"enabled"`
	ManualConfig                types.Bool   `tfsdk:"manual_config"`
	Interfaces                  types.Set    `tfsdk:"interfaces"`
	ServiceSocketsMaxRetries    types.Int64  `tfsdk:"service_sockets_max_retries"`
	ServiceSocketsRetryWaitTime types.Int64  `tfsdk:"service_sockets_retry_wait_time"`
	ValidLifetime               types.Int64  `tfsdk:"valid_lifetime"`
	DeclineProbationPeriod      types.Int64  `tfsdk:"decline_probation_period"`
	FirewallRules               types.Bool   `tfsdk:"firewall_rules"`
	MacSources                  types.Set    `tfsdk:"mac_sources"`
	HoldReclaimedTime           types.Int64  `tfsdk:"hold_reclaimed_time"`
	ReclaimTimerWaitTime        types.Int64  `tfsdk:"reclaim_timer_wait_time"`
	FlushReclaimedTimerWaitTime types.Int64  `tfsdk:"flush_reclaimed_timer_wait_time"`
	MaxReclaimTime              types.Int64  `tfsdk:"max_reclaim_time"`
	MaxReclaimLeases            types.Int64  `tfsdk:"max_reclaim_leases"`
	UnwarnedReclaimCycles       types.Int64  `tfsdk:"unwarned_reclaim_cycles"`
	HaEnabled                   types.Bool   `tfsdk:"ha_enabled"`
	HaThisServerName            types.String `tfsdk:"ha_this_server_name"`
	HaMaxUnackedClients         types.Int64  `tfsdk:"ha_max_unacked_clients"`

	Id types.String `tfsdk:"id"`
}

func dhcpv6SettingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable the Kea DHCPv6 server.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"manual_config": schema.BoolAttribute{
			MarkdownDescription: "Use a hand-written configuration file instead of the generated one.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"interfaces": schema.SetAttribute{
			MarkdownDescription: "Interfaces to serve, e.g. `[\"lan\"]`.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"service_sockets_max_retries": schema.Int64Attribute{
			MarkdownDescription: "Retries when binding a socket fails. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"service_sockets_retry_wait_time": schema.Int64Attribute{
			MarkdownDescription: "Milliseconds between socket bind retries. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"valid_lifetime": schema.Int64Attribute{
			MarkdownDescription: "Default lease lifetime in seconds. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"decline_probation_period": schema.Int64Attribute{
			MarkdownDescription: "Seconds a declined address is held back. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"firewall_rules": schema.BoolAttribute{
			MarkdownDescription: "Add the firewall rules that allow DHCP on the served interfaces.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"mac_sources": schema.SetAttribute{
			MarkdownDescription: "Sources of client MAC addresses, e.g. `[\"ipv6-link-local\", \"duid\"]`.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"hold_reclaimed_time": schema.Int64Attribute{
			MarkdownDescription: "Seconds an expired lease is kept before being removed. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"reclaim_timer_wait_time": schema.Int64Attribute{
			MarkdownDescription: "Seconds between lease reclamation runs. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"flush_reclaimed_timer_wait_time": schema.Int64Attribute{
			MarkdownDescription: "Seconds between removals of reclaimed leases. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"max_reclaim_time": schema.Int64Attribute{
			MarkdownDescription: "Maximum duration of a reclamation run, in milliseconds. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"max_reclaim_leases": schema.Int64Attribute{
			MarkdownDescription: "Maximum leases reclaimed per run. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"unwarned_reclaim_cycles": schema.Int64Attribute{
			MarkdownDescription: "Unfinished reclamation runs before a warning is logged. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"ha_enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable high availability with the peers defined in the HA peer resources.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"ha_this_server_name": schema.StringAttribute{
			MarkdownDescription: "Name of this server among the HA peers, empty for the hostname.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"ha_max_unacked_clients": schema.Int64Attribute{
			MarkdownDescription: "Clients the partner may leave unanswered before failover. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `kea_dhcpv6_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertDhcpv6SettingsSchemaToStruct(d *dhcpv6SettingsResourceModel) *kea.Dhcpv6Config {
	return &kea.Dhcpv6Config{
		General: kea.Dhcpv6ConfigGeneral{
			Enabled:                     settingsBool(d.Enabled),
			ManualConfig:                settingsBool(d.ManualConfig),
			Interfaces:                  api.SelectedMapList(settingsList(d.Interfaces)),
			ServiceSocketsMaxRetries:    settingsInt(d.ServiceSocketsMaxRetries),
			ServiceSocketsRetryWaitTime: settingsInt(d.ServiceSocketsRetryWaitTime),
			ValidLifetime:               settingsInt(d.ValidLifetime),
			DeclineProbationPeriod:      settingsInt(d.DeclineProbationPeriod),
			FirewallRules:               settingsBool(d.FirewallRules),
			MacSources:                  api.SelectedMapList(settingsList(d.MacSources)),
		},
		Lexpire: kea.Dhcpv6ConfigLexpire{
			HoldReclaimedTime:           settingsInt(d.HoldReclaimedTime),
			ReclaimTimerWaitTime:        settingsInt(d.ReclaimTimerWaitTime),
			FlushReclaimedTimerWaitTime: settingsInt(d.FlushReclaimedTimerWaitTime),
			MaxReclaimTime:              settingsInt(d.MaxReclaimTime),
			MaxReclaimLeases:            settingsInt(d.MaxReclaimLeases),
			UnwarnedReclaimCycles:       settingsInt(d.UnwarnedReclaimCycles),
		},
		Ha: kea.Dhcpv6ConfigHa{
			HaEnabled:           settingsBool(d.HaEnabled),
			HaThisServerName:    d.HaThisServerName.ValueString(),
			HaMaxUnackedClients: settingsInt(d.HaMaxUnackedClients),
		},
	}
}

func convertDhcpv6SettingsStructToSchema(s *kea.Dhcpv6Config) *dhcpv6SettingsResourceModel {
	return &dhcpv6SettingsResourceModel{
		Enabled:                     types.BoolValue(tools.StringToBool(s.General.Enabled)),
		ManualConfig:                types.BoolValue(tools.StringToBool(s.General.ManualConfig)),
		Interfaces:                  tools.StringSliceToSet(settingsNonEmpty(s.General.Interfaces)),
		ServiceSocketsMaxRetries:    settingsReadInt(s.General.ServiceSocketsMaxRetries),
		ServiceSocketsRetryWaitTime: settingsReadInt(s.General.ServiceSocketsRetryWaitTime),
		ValidLifetime:               settingsReadInt(s.General.ValidLifetime),
		DeclineProbationPeriod:      settingsReadInt(s.General.DeclineProbationPeriod),
		FirewallRules:               types.BoolValue(tools.StringToBool(s.General.FirewallRules)),
		MacSources:                  tools.StringSliceToSet(settingsNonEmpty(s.General.MacSources)),
		HoldReclaimedTime:           settingsReadInt(s.Lexpire.HoldReclaimedTime),
		ReclaimTimerWaitTime:        settingsReadInt(s.Lexpire.ReclaimTimerWaitTime),
		FlushReclaimedTimerWaitTime: settingsReadInt(s.Lexpire.FlushReclaimedTimerWaitTime),
		MaxReclaimTime:              settingsReadInt(s.Lexpire.MaxReclaimTime),
		MaxReclaimLeases:            settingsReadInt(s.Lexpire.MaxReclaimLeases),
		UnwarnedReclaimCycles:       settingsReadInt(s.Lexpire.UnwarnedReclaimCycles),
		HaEnabled:                   types.BoolValue(tools.StringToBool(s.Ha.HaEnabled)),
		HaThisServerName:            types.StringValue(s.Ha.HaThisServerName),
		HaMaxUnackedClients:         settingsReadInt(s.Ha.HaMaxUnackedClients),
	}
}
