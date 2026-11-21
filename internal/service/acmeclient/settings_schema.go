package acmeclient

import (
	"github.com/browningluke/opnsense-go/pkg/acmeclient"
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type settingsResourceModel struct {
	Enabled            types.Bool   `tfsdk:"enabled"`
	AutoRenewal        types.Bool   `tfsdk:"auto_renewal"`
	UpdateCron         types.String `tfsdk:"update_cron"`
	Environment        types.String `tfsdk:"environment"`
	ChallengePort      types.Int64  `tfsdk:"challenge_port"`
	TlsChallengePort   types.Int64  `tfsdk:"tls_challenge_port"`
	RestartTimeout     types.Int64  `tfsdk:"restart_timeout"`
	HaproxyIntegration types.Bool   `tfsdk:"haproxy_integration"`
	HaproxyAclRef      types.String `tfsdk:"haproxy_acl_ref"`
	HaproxyActionRef   types.String `tfsdk:"haproxy_action_ref"`
	HaproxyServerRef   types.String `tfsdk:"haproxy_server_ref"`
	HaproxyBackendRef  types.String `tfsdk:"haproxy_backend_ref"`
	LogLevel           types.String `tfsdk:"log_level"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable ACME client plugin.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"auto_renewal": schema.BoolAttribute{
			MarkdownDescription: "Enable automatic renewal for certificates to prevent expiration. This will add a cron job to the system. You may want to customize the cron job schedule to your needs, because re-issueing a certificate may lead to a short downtime, depending on the selected challenge type and service.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"update_cron": schema.StringAttribute{
			MarkdownDescription: "UpdateCron.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"environment": schema.StringAttribute{
			MarkdownDescription: "environment. One of `prod`, `stg`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"challenge_port": schema.Int64Attribute{
			MarkdownDescription: "When using HTTP-01 as challenge type, a local webserver is used to provide acme challenge data to the ACME CA. The local webserver is NOT directly exposed to the outside and should NOT use port 80 or any other well-known port. This setting allows you to change the local port of this webserver in case it interferes with another local service. Defaults to port 43580. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"tls_challenge_port": schema.Int64Attribute{
			MarkdownDescription: "The service port when using TLS-ALPN-01 as challenge type. It works similar to the HTTP-01 challenge type. Defaults to port 43581. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"restart_timeout": schema.Int64Attribute{
			MarkdownDescription: "The maximum time in seconds to wait for an automation to complete. When the timeout is reached the command is forcefully aborted. Defaults to 600 seconds. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"haproxy_integration": schema.BoolAttribute{
			MarkdownDescription: "Enable automatic integration with the OPNsense HAProxy plugin. Requires that the OPNsense HAProxy plugin is installed. This will automatically add the required backend, server, action and ACL for you. You just need to select your HAProxy frontend when configuring the certificate or challenge type. NOTE:This will only work for HTTP-01 validation and HAProxy frontends running in http mode; TCP frontends are not supported.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"haproxy_acl_ref": schema.StringAttribute{
			MarkdownDescription: "haproxyAclRef.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"haproxy_action_ref": schema.StringAttribute{
			MarkdownDescription: "haproxyActionRef.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"haproxy_server_ref": schema.StringAttribute{
			MarkdownDescription: "haproxyServerRef.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"haproxy_backend_ref": schema.StringAttribute{
			MarkdownDescription: "haproxyBackendRef.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"log_level": schema.StringAttribute{
			MarkdownDescription: "Specifies the log level for acme.sh, default is \"normal\". All other log levels add information for debug purposes, but be aware that this will break the log formatting in the GUI. Levels \"debug 2\" and \"debug 3\" log successively deeper log messages from the acme.sh including messages from DNS-01 DNSAPI scripts. One of `normal`, `extended`, `debug`, `debug2`, `debug3`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `acmeclient_settings`.",
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

func convertSettingsSchemaToStruct(d *settingsResourceModel) *acmeclient.Settings {
	return &acmeclient.Settings{
		Settings: acmeclient.SettingsSettings{
			Enabled:            settingsBool(d.Enabled),
			AutoRenewal:        settingsBool(d.AutoRenewal),
			UpdateCron:         api.SelectedMap(d.UpdateCron.ValueString()),
			Environment:        api.SelectedMap(d.Environment.ValueString()),
			ChallengePort:      settingsInt(d.ChallengePort),
			TlsChallengePort:   settingsInt(d.TlsChallengePort),
			RestartTimeout:     settingsInt(d.RestartTimeout),
			HaproxyIntegration: settingsBool(d.HaproxyIntegration),
			HaproxyAclRef:      api.SelectedMap(d.HaproxyAclRef.ValueString()),
			HaproxyActionRef:   api.SelectedMap(d.HaproxyActionRef.ValueString()),
			HaproxyServerRef:   api.SelectedMap(d.HaproxyServerRef.ValueString()),
			HaproxyBackendRef:  api.SelectedMap(d.HaproxyBackendRef.ValueString()),
			LogLevel:           api.SelectedMap(d.LogLevel.ValueString()),
		},
	}
}

func convertSettingsStructToSchema(s *acmeclient.Settings) *settingsResourceModel {
	return &settingsResourceModel{
		Enabled:            types.BoolValue(tools.StringToBool(s.Settings.Enabled)),
		AutoRenewal:        types.BoolValue(tools.StringToBool(s.Settings.AutoRenewal)),
		UpdateCron:         types.StringValue(s.Settings.UpdateCron.String()),
		Environment:        types.StringValue(s.Settings.Environment.String()),
		ChallengePort:      settingsReadInt(s.Settings.ChallengePort),
		TlsChallengePort:   settingsReadInt(s.Settings.TlsChallengePort),
		RestartTimeout:     settingsReadInt(s.Settings.RestartTimeout),
		HaproxyIntegration: types.BoolValue(tools.StringToBool(s.Settings.HaproxyIntegration)),
		HaproxyAclRef:      types.StringValue(s.Settings.HaproxyAclRef.String()),
		HaproxyActionRef:   types.StringValue(s.Settings.HaproxyActionRef.String()),
		HaproxyServerRef:   types.StringValue(s.Settings.HaproxyServerRef.String()),
		HaproxyBackendRef:  types.StringValue(s.Settings.HaproxyBackendRef.String()),
		LogLevel:           types.StringValue(s.Settings.LogLevel.String()),
	}
}
