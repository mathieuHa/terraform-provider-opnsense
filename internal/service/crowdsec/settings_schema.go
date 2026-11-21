package crowdsec

import (
	"github.com/browningluke/opnsense-go/pkg/crowdsec"
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
	AgentEnabled            types.Bool   `tfsdk:"agent_enabled"`
	LapiEnabled             types.Bool   `tfsdk:"lapi_enabled"`
	FirewallBouncerEnabled  types.Bool   `tfsdk:"firewall_bouncer_enabled"`
	LapiManualConfiguration types.Bool   `tfsdk:"lapi_manual_configuration"`
	LapiListenAddress       types.String `tfsdk:"lapi_listen_address"`
	LapiListenPort          types.Int64  `tfsdk:"lapi_listen_port"`
	RulesEnabled            types.Bool   `tfsdk:"rules_enabled"`
	RulesLog                types.Bool   `tfsdk:"rules_log"`
	RulesTag                types.String `tfsdk:"rules_tag"`
	EnrollKey               types.String `tfsdk:"enroll_key"`
	CrowdsecFirewallVerbose types.Bool   `tfsdk:"crowdsec_firewall_verbose"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"agent_enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable/disable the CrowdSec agent. Keep this enabled to detect attacks and receive alerts from the CrowSec central service.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"lapi_enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable/disable the CrowdSec Local API. Keep this enabled unless you connect to a LAPI on another machine.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"firewall_bouncer_enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable/disable the remediation component. Keep this enabled to block packets from the attacking IP addresses.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"lapi_manual_configuration": schema.BoolAttribute{
			MarkdownDescription: "Avoid overwriting LAPI settings for config.yaml, local_api_credentials.yaml, crowdsec-firewall-bouncer.yaml. The next two configuration options (lapi_listen_address, lapi_listen_port) will be ignored. Allows unsupported configurations like linking together multiple opnsense instances or connecting to an existing crowdsec multi-server setup.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"lapi_listen_address": schema.StringAttribute{
			MarkdownDescription: "Where to listen for LAPI connections: IP address. The default value is 127.0.0.1. You can change it to a LAN address to connect from other agents/machines and bouncers. This is written in /usr/local/etc/crowdsec/config.yaml, local_api_credentials.yaml and bouncers/crowdsec-firewall-bouncer.yaml. To enable TLS, add the certificate information to config.yaml and change http to https in the other two files. Comments in YAML will not be preserved.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"lapi_listen_port": schema.Int64Attribute{
			MarkdownDescription: "Where to listen for LAPI connections: port. The default value is 8080, but you can change it to avoid conflicts with existing services. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"rules_enabled": schema.BoolAttribute{
			MarkdownDescription: "Generate block rules from the Crowdsec blocklists. They are applied to all interfaces, ipv4/v6, ingress and egress. If you disable this, you'll have to write your own rules to block anything.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"rules_log": schema.BoolAttribute{
			MarkdownDescription: "Enable log collection for CrowdSec's block rules.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"rules_tag": schema.StringAttribute{
			MarkdownDescription: "Add a tag to packets that are dropped by CrowdSec rules for diagnostic purposes.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"enroll_key": schema.StringAttribute{
			MarkdownDescription: "Click \"Enroll command\" on the the website and copy the key here.",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"crowdsec_firewall_verbose": schema.BoolAttribute{
			MarkdownDescription: "Verbose /var/log/crowdsec/crowdsec-firewall-bouncer.log. Enable this for debugging.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `crowdsec_settings`.",
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

func convertSettingsSchemaToStruct(d *settingsResourceModel) *crowdsec.Settings {
	return &crowdsec.Settings{
		AgentEnabled:            settingsBool(d.AgentEnabled),
		LapiEnabled:             settingsBool(d.LapiEnabled),
		FirewallBouncerEnabled:  settingsBool(d.FirewallBouncerEnabled),
		LapiManualConfiguration: settingsBool(d.LapiManualConfiguration),
		LapiListenAddress:       d.LapiListenAddress.ValueString(),
		LapiListenPort:          settingsInt(d.LapiListenPort),
		RulesEnabled:            settingsBool(d.RulesEnabled),
		RulesLog:                settingsBool(d.RulesLog),
		RulesTag:                d.RulesTag.ValueString(),
		EnrollKey:               d.EnrollKey.ValueString(),
		CrowdsecFirewallVerbose: settingsBool(d.CrowdsecFirewallVerbose),
	}
}

func convertSettingsStructToSchema(s *crowdsec.Settings) *settingsResourceModel {
	return &settingsResourceModel{
		AgentEnabled:            types.BoolValue(tools.StringToBool(s.AgentEnabled)),
		LapiEnabled:             types.BoolValue(tools.StringToBool(s.LapiEnabled)),
		FirewallBouncerEnabled:  types.BoolValue(tools.StringToBool(s.FirewallBouncerEnabled)),
		LapiManualConfiguration: types.BoolValue(tools.StringToBool(s.LapiManualConfiguration)),
		LapiListenAddress:       types.StringValue(s.LapiListenAddress),
		LapiListenPort:          settingsReadInt(s.LapiListenPort),
		RulesEnabled:            types.BoolValue(tools.StringToBool(s.RulesEnabled)),
		RulesLog:                types.BoolValue(tools.StringToBool(s.RulesLog)),
		RulesTag:                types.StringValue(s.RulesTag),
		EnrollKey:               types.StringValue(s.EnrollKey),
		CrowdsecFirewallVerbose: types.BoolValue(tools.StringToBool(s.CrowdsecFirewallVerbose)),
	}
}
