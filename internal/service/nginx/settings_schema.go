package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
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
	Enabled                       types.Bool   `tfsdk:"enabled"`
	BanTtl                        types.Int64  `tfsdk:"ban_ttl"`
	HttpWorkerProcesses           types.Int64  `tfsdk:"http_worker_processes"`
	HttpWorkerConnections         types.Int64  `tfsdk:"http_worker_connections"`
	HttpSendfile                  types.Bool   `tfsdk:"http_sendfile"`
	HttpKeepaliveTimeout          types.Int64  `tfsdk:"http_keepalive_timeout"`
	HttpResetTimedout             types.Bool   `tfsdk:"http_reset_timedout"`
	HttpDefaultType               types.String `tfsdk:"http_default_type"`
	HttpServerNamesHashBucketSize types.Int64  `tfsdk:"http_server_names_hash_bucket_size"`
	HttpServerNamesHashMaxSize    types.Int64  `tfsdk:"http_server_names_hash_max_size"`
	HttpVariablesHashMaxSize      types.Int64  `tfsdk:"http_variables_hash_max_size"`
	HttpVariablesHashBucketSize   types.Int64  `tfsdk:"http_variables_hash_bucket_size"`
	HttpBanResponse               types.String `tfsdk:"http_ban_response"`
	HttpLogPermBan                types.Bool   `tfsdk:"http_log_perm_ban"`
	HttpBotsUa                    types.Set    `tfsdk:"http_bots_ua"`
	HttpHeadersMoreEnable         types.Bool   `tfsdk:"http_headers_more_enable"`
	WebguiLimitNetworks           types.Bool   `tfsdk:"webgui_limit_networks"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable configured services.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"ban_ttl": schema.Int64Attribute{
			MarkdownDescription: "Set autoblock lifetime in minutes. 72 hours by default. Set to 0 for infinite. Please note that setting this to 0 may result in gradual system slowdown and the need to manually clear the entries. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"http_worker_processes": schema.Int64Attribute{
			MarkdownDescription: "Set the number of worker processes to run. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"http_worker_connections": schema.Int64Attribute{
			MarkdownDescription: "Set the maximum connections of each worker. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"http_sendfile": schema.BoolAttribute{
			MarkdownDescription: "Enable sendfile support (faster).",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"http_keepalive_timeout": schema.Int64Attribute{
			MarkdownDescription: "After this idle time, the client gets disconnected. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"http_reset_timedout": schema.BoolAttribute{
			MarkdownDescription: "Reset timed out connections and connections closed with the non-standard code 444. When the socket is closed, TCP RST is sent to the client, and all memory occupied by this socket is released. This helps avoid keeping an already closed socket with filled buffers in a FIN_WAIT1 state for a long time.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"http_default_type": schema.StringAttribute{
			MarkdownDescription: "This content type is sent if the file extension is unknown.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"http_server_names_hash_bucket_size": schema.Int64Attribute{
			MarkdownDescription: "Server Names Hash Bucket Size. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"http_server_names_hash_max_size": schema.Int64Attribute{
			MarkdownDescription: "Server Names Hash Max Size. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"http_variables_hash_max_size": schema.Int64Attribute{
			MarkdownDescription: "Variables Hash Max Size. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"http_variables_hash_bucket_size": schema.Int64Attribute{
			MarkdownDescription: "Variables Hash Bucket Size. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"http_ban_response": schema.StringAttribute{
			MarkdownDescription: "Select a response code for auto-blocking requests (bot user-agent or honeypot location). The default code is 403. 444 is a special response code that closes the connection without a response to the client. One of `403`, `444`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"http_log_perm_ban": schema.BoolAttribute{
			MarkdownDescription: "Log requests that led to auto-blocking (bots and honeypots).",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"http_bots_ua": schema.SetAttribute{
			MarkdownDescription: "List of bot user agents that will be used in bot protection. Uses a well known list by default.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"http_headers_more_enable": schema.BoolAttribute{
			MarkdownDescription: "Enhanced version of the standard headers module. Allows to add, set, or clear any output or input header.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"webgui_limit_networks": schema.BoolAttribute{
			MarkdownDescription: "Limit networks to directly connected networks. Enabling this option is recommended if your nginx instance is reachable via the internet to prevent remote access to the web interface for security reasons. Please note that you can lock yourself out if you are accessing OPNsense via a router without SNAT.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `nginx_settings`.",
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

func convertSettingsSchemaToStruct(d *settingsResourceModel) *nginx.Settings {
	return &nginx.Settings{
		General: nginx.SettingsGeneral{
			Enabled: settingsBool(d.Enabled),
			BanTtl:  settingsInt(d.BanTtl),
		},
		Http: nginx.SettingsHttp{
			HttpWorkerProcesses:           settingsInt(d.HttpWorkerProcesses),
			HttpWorkerConnections:         settingsInt(d.HttpWorkerConnections),
			HttpSendfile:                  settingsBool(d.HttpSendfile),
			HttpKeepaliveTimeout:          settingsInt(d.HttpKeepaliveTimeout),
			HttpResetTimedout:             settingsBool(d.HttpResetTimedout),
			HttpDefaultType:               d.HttpDefaultType.ValueString(),
			HttpServerNamesHashBucketSize: settingsInt(d.HttpServerNamesHashBucketSize),
			HttpServerNamesHashMaxSize:    settingsInt(d.HttpServerNamesHashMaxSize),
			HttpVariablesHashMaxSize:      settingsInt(d.HttpVariablesHashMaxSize),
			HttpVariablesHashBucketSize:   settingsInt(d.HttpVariablesHashBucketSize),
			HttpBanResponse:               api.SelectedMap(d.HttpBanResponse.ValueString()),
			HttpLogPermBan:                settingsBool(d.HttpLogPermBan),
			HttpBotsUa:                    api.SelectedMapList(settingsList(d.HttpBotsUa)),
			HttpHeadersMoreEnable:         settingsBool(d.HttpHeadersMoreEnable),
		},
		Webgui: nginx.SettingsWebgui{
			WebguiLimitNetworks: settingsBool(d.WebguiLimitNetworks),
		},
	}
}

func convertSettingsStructToSchema(s *nginx.Settings) *settingsResourceModel {
	return &settingsResourceModel{
		Enabled:                       types.BoolValue(tools.StringToBool(s.General.Enabled)),
		BanTtl:                        settingsReadInt(s.General.BanTtl),
		HttpWorkerProcesses:           settingsReadInt(s.Http.HttpWorkerProcesses),
		HttpWorkerConnections:         settingsReadInt(s.Http.HttpWorkerConnections),
		HttpSendfile:                  types.BoolValue(tools.StringToBool(s.Http.HttpSendfile)),
		HttpKeepaliveTimeout:          settingsReadInt(s.Http.HttpKeepaliveTimeout),
		HttpResetTimedout:             types.BoolValue(tools.StringToBool(s.Http.HttpResetTimedout)),
		HttpDefaultType:               types.StringValue(s.Http.HttpDefaultType),
		HttpServerNamesHashBucketSize: settingsReadInt(s.Http.HttpServerNamesHashBucketSize),
		HttpServerNamesHashMaxSize:    settingsReadInt(s.Http.HttpServerNamesHashMaxSize),
		HttpVariablesHashMaxSize:      settingsReadInt(s.Http.HttpVariablesHashMaxSize),
		HttpVariablesHashBucketSize:   settingsReadInt(s.Http.HttpVariablesHashBucketSize),
		HttpBanResponse:               types.StringValue(s.Http.HttpBanResponse.String()),
		HttpLogPermBan:                types.BoolValue(tools.StringToBool(s.Http.HttpLogPermBan)),
		HttpBotsUa:                    tools.StringSliceToSet(settingsNonEmpty(s.Http.HttpBotsUa)),
		HttpHeadersMoreEnable:         types.BoolValue(tools.StringToBool(s.Http.HttpHeadersMoreEnable)),
		WebguiLimitNetworks:           types.BoolValue(tools.StringToBool(s.Webgui.WebguiLimitNetworks)),
	}
}
