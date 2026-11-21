package ids

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/ids"
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

type settingsResourceModel struct {
	Enabled                 types.Bool   `tfsdk:"enabled"`
	Mode                    types.String `tfsdk:"mode"`
	DivertListeners         types.Int64  `tfsdk:"divert_listeners"`
	Promiscuous             types.Bool   `tfsdk:"promiscuous"`
	Interfaces              types.Set    `tfsdk:"interfaces"`
	HomeNetworks            types.Set    `tfsdk:"home_networks"`
	DefaultPacketSize       types.Int64  `tfsdk:"default_packet_size"`
	UpdateCron              types.String `tfsdk:"update_cron"`
	AlertLogrotate          types.String `tfsdk:"alert_logrotate"`
	AlertSaveLogs           types.Int64  `tfsdk:"alert_save_logs"`
	PatternMatcher          types.String `tfsdk:"pattern_matcher"`
	DetectProfile           types.String `tfsdk:"detect_profile"`
	DetectToclientGroups    types.Int64  `tfsdk:"detect_toclient_groups"`
	DetectToserverGroups    types.Int64  `tfsdk:"detect_toserver_groups"`
	SyslogAlerts            types.Bool   `tfsdk:"syslog_alerts"`
	SyslogEve               types.Bool   `tfsdk:"syslog_eve"`
	LogPayload              types.Bool   `tfsdk:"log_payload"`
	Verbosity               types.String `tfsdk:"verbosity"`
	EveHttpEnabled          types.Bool   `tfsdk:"eve_http_enabled"`
	EveHttpExtended         types.Bool   `tfsdk:"eve_http_extended"`
	EveHttpDumpAllHeaders   types.String `tfsdk:"eve_http_dump_all_headers"`
	EveTlsEnabled           types.Bool   `tfsdk:"eve_tls_enabled"`
	EveTlsExtended          types.Bool   `tfsdk:"eve_tls_extended"`
	EveTlsSessionResumption types.Bool   `tfsdk:"eve_tls_session_resumption"`
	EveTlsCustom            types.Set    `tfsdk:"eve_tls_custom"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable the Suricata intrusion detection service.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"mode": schema.StringAttribute{
			MarkdownDescription: "Capture mode: `pcap` (IDS only), `netmap` or `divert` (IPS, allows drop actions).",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"divert_listeners": schema.Int64Attribute{
			MarkdownDescription: "Number of divert listeners (1-256), divert mode only. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"promiscuous": schema.BoolAttribute{
			MarkdownDescription: "Put the interfaces in promiscuous mode, needed for VLANs in netmap mode.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"interfaces": schema.SetAttribute{
			MarkdownDescription: "Interfaces to inspect, e.g. `[\"wan\"]`.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"home_networks": schema.SetAttribute{
			MarkdownDescription: "Networks treated as `$HOME_NET`.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"default_packet_size": schema.Int64Attribute{
			MarkdownDescription: "Default packet size (82-65535), empty for the Suricata default. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"update_cron": schema.StringAttribute{
			MarkdownDescription: "UUID of the cron job that updates the rules.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"alert_logrotate": schema.StringAttribute{
			MarkdownDescription: "Alert log rotation: `W0D23` (weekly) or `D0` (daily).",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"alert_save_logs": schema.Int64Attribute{
			MarkdownDescription: "Number of rotated alert logs to keep (1-1000). `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"pattern_matcher": schema.StringAttribute{
			MarkdownDescription: "Multi-pattern matcher: empty for the default, `ac`, `ac-ks` or `hs`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"detect_profile": schema.StringAttribute{
			MarkdownDescription: "Detection profile: empty for the default, `low`, `medium`, `high` or `custom`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"detect_toclient_groups": schema.Int64Attribute{
			MarkdownDescription: "toclient-groups of the custom detection profile. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"detect_toserver_groups": schema.Int64Attribute{
			MarkdownDescription: "toserver-groups of the custom detection profile. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"syslog_alerts": schema.BoolAttribute{
			MarkdownDescription: "Send alerts to the system log in fast_log format.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"syslog_eve": schema.BoolAttribute{
			MarkdownDescription: "Send alerts to the system log in eve format.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"log_payload": schema.BoolAttribute{
			MarkdownDescription: "Include the packet payload in eve alerts.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"verbosity": schema.StringAttribute{
			MarkdownDescription: "Log verbosity: empty for the default, `v`, `vv`, `vvv` or `vvvv`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"eve_http_enabled": schema.BoolAttribute{
			MarkdownDescription: "Log HTTP transactions to the eve log.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"eve_http_extended": schema.BoolAttribute{
			MarkdownDescription: "Extended HTTP logging.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"eve_http_dump_all_headers": schema.StringAttribute{
			MarkdownDescription: "Headers dumped in HTTP logs: empty for none, `request`, `response` or `both`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"eve_tls_enabled": schema.BoolAttribute{
			MarkdownDescription: "Log TLS handshakes to the eve log.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"eve_tls_extended": schema.BoolAttribute{
			MarkdownDescription: "Extended TLS logging.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"eve_tls_session_resumption": schema.BoolAttribute{
			MarkdownDescription: "Log TLS session resumptions.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"eve_tls_custom": schema.SetAttribute{
			MarkdownDescription: "Custom TLS fields to log, e.g. `[\"sni\", \"ja4\"]`.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `ids_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func settingsBool(b types.Bool) string {
	if b.IsNull() || b.IsUnknown() {
		return ""
	}
	return tools.BoolToString(b.ValueBool())
}

func settingsInt(i types.Int64) string {
	if i.IsNull() || i.IsUnknown() || i.ValueInt64() == -1 {
		return ""
	}
	return tools.Int64ToString(i.ValueInt64())
}

func settingsReadInt(s string) types.Int64 {
	if s == "" {
		return types.Int64Value(-1)
	}
	return tools.StringToInt64Null(s)
}

func settingsList(s types.Set) []string {
	out := tools.SetToStringSlice(s)
	if out == nil {
		return []string{}
	}
	return out
}

func settingsNonEmpty(in []string) []string {
	out := []string{}
	for _, v := range in {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

func convertSettingsSchemaToStruct(d *settingsResourceModel) *ids.Settings {
	return &ids.Settings{
		General: ids.SettingsGeneral{
			Enabled:           settingsBool(d.Enabled),
			Mode:              api.SelectedMap(d.Mode.ValueString()),
			DivertListeners:   settingsInt(d.DivertListeners),
			Promiscuous:       settingsBool(d.Promiscuous),
			Interfaces:        api.SelectedMapList(settingsList(d.Interfaces)),
			HomeNetworks:      api.SelectedMapList(settingsList(d.HomeNetworks)),
			DefaultPacketSize: settingsInt(d.DefaultPacketSize),
			UpdateCron:        api.SelectedMap(d.UpdateCron.ValueString()),
			AlertLogrotate:    api.SelectedMap(d.AlertLogrotate.ValueString()),
			AlertSaveLogs:     settingsInt(d.AlertSaveLogs),
			PatternMatcher:    api.SelectedMap(d.PatternMatcher.ValueString()),
			SyslogAlerts:      settingsBool(d.SyslogAlerts),
			SyslogEve:         settingsBool(d.SyslogEve),
			LogPayload:        settingsBool(d.LogPayload),
			Verbosity:         api.SelectedMap(d.Verbosity.ValueString()),
			Detect: ids.SettingsGeneralDetect{
				DetectProfile:        api.SelectedMap(d.DetectProfile.ValueString()),
				DetectToclientGroups: settingsInt(d.DetectToclientGroups),
				DetectToserverGroups: settingsInt(d.DetectToserverGroups),
			},
			EveLog: ids.SettingsGeneralEveLog{
				Http: ids.SettingsGeneralEveLogHttp{
					EveHttpEnabled:        settingsBool(d.EveHttpEnabled),
					EveHttpExtended:       settingsBool(d.EveHttpExtended),
					EveHttpDumpAllHeaders: api.SelectedMap(d.EveHttpDumpAllHeaders.ValueString()),
				},
				Tls: ids.SettingsGeneralEveLogTls{
					EveTlsEnabled:           settingsBool(d.EveTlsEnabled),
					EveTlsExtended:          settingsBool(d.EveTlsExtended),
					EveTlsSessionResumption: settingsBool(d.EveTlsSessionResumption),
					EveTlsCustom:            api.SelectedMapList(settingsList(d.EveTlsCustom)),
				},
			},
		},
	}
}

func convertSettingsStructToSchema(s *ids.Settings) *settingsResourceModel {
	return &settingsResourceModel{
		Enabled:                 types.BoolValue(tools.StringToBool(s.General.Enabled)),
		Mode:                    types.StringValue(s.General.Mode.String()),
		DivertListeners:         settingsReadInt(s.General.DivertListeners),
		Promiscuous:             types.BoolValue(tools.StringToBool(s.General.Promiscuous)),
		Interfaces:              tools.StringSliceToSet(settingsNonEmpty(s.General.Interfaces)),
		HomeNetworks:            tools.StringSliceToSet(settingsNonEmpty(s.General.HomeNetworks)),
		DefaultPacketSize:       settingsReadInt(s.General.DefaultPacketSize),
		UpdateCron:              types.StringValue(s.General.UpdateCron.String()),
		AlertLogrotate:          types.StringValue(s.General.AlertLogrotate.String()),
		AlertSaveLogs:           settingsReadInt(s.General.AlertSaveLogs),
		PatternMatcher:          types.StringValue(s.General.PatternMatcher.String()),
		DetectProfile:           types.StringValue(s.General.Detect.DetectProfile.String()),
		DetectToclientGroups:    settingsReadInt(s.General.Detect.DetectToclientGroups),
		DetectToserverGroups:    settingsReadInt(s.General.Detect.DetectToserverGroups),
		SyslogAlerts:            types.BoolValue(tools.StringToBool(s.General.SyslogAlerts)),
		SyslogEve:               types.BoolValue(tools.StringToBool(s.General.SyslogEve)),
		LogPayload:              types.BoolValue(tools.StringToBool(s.General.LogPayload)),
		Verbosity:               types.StringValue(s.General.Verbosity.String()),
		EveHttpEnabled:          types.BoolValue(tools.StringToBool(s.General.EveLog.Http.EveHttpEnabled)),
		EveHttpExtended:         types.BoolValue(tools.StringToBool(s.General.EveLog.Http.EveHttpExtended)),
		EveHttpDumpAllHeaders:   types.StringValue(s.General.EveLog.Http.EveHttpDumpAllHeaders.String()),
		EveTlsEnabled:           types.BoolValue(tools.StringToBool(s.General.EveLog.Tls.EveTlsEnabled)),
		EveTlsExtended:          types.BoolValue(tools.StringToBool(s.General.EveLog.Tls.EveTlsExtended)),
		EveTlsSessionResumption: types.BoolValue(tools.StringToBool(s.General.EveLog.Tls.EveTlsSessionResumption)),
		EveTlsCustom:            tools.StringSliceToSet(settingsNonEmpty(s.General.EveLog.Tls.EveTlsCustom)),
	}
}
