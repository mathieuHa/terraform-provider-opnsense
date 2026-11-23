package bind

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/bind"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type dnsblSettingsResourceModel struct {
	Enabled             types.Bool `tfsdk:"enabled"`
	Type                types.Set  `tfsdk:"type"`
	Whitelists          types.Set  `tfsdk:"whitelists"`
	Forcesafegoogle     types.Bool `tfsdk:"forcesafegoogle"`
	Forcesafeduckduckgo types.Bool `tfsdk:"forcesafeduckduckgo"`
	Forcesafeyoutube    types.Bool `tfsdk:"forcesafeyoutube"`
	Forcestrictbing     types.Bool `tfsdk:"forcestrictbing"`

	Id types.String `tfsdk:"id"`
}

func dnsblSettingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "This will enable the use of DNS Blocklists for ADs, Malware, or both.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"type": schema.SetAttribute{
			MarkdownDescription: "Select which kind of DNSBL you want to use. One of `aa`, `ag`, `bla`, `blf`, `blp`, `ca`, `el`, `emd`, `ep`, `hpa`, `hpf`, `hpp`, `hup`, `mw`, `nc`, `pt`, `rw`, `sa`, `st`, `sb`, `ws`, `wsu`, `wse`, `yy`.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"whitelists": schema.SetAttribute{
			MarkdownDescription: "List of domains to whitelist. It will add a entry for the domains itself and all sub domains.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"forcesafegoogle": schema.BoolAttribute{
			MarkdownDescription: "This will force SafeSearch.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"forcesafeduckduckgo": schema.BoolAttribute{
			MarkdownDescription: "This will force SafeSearch.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"forcesafeyoutube": schema.BoolAttribute{
			MarkdownDescription: "This will force safe Youtube browsing.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"forcestrictbing": schema.BoolAttribute{
			MarkdownDescription: "This will force safe search when using Bing.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `bind_dnsbl_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertDnsblSettingsSchemaToStruct(d *dnsblSettingsResourceModel) *bind.DnsblSettings {
	return &bind.DnsblSettings{
		Enabled:             settingsBool(d.Enabled),
		Type:                api.SelectedMapList(settingsList(d.Type)),
		Whitelists:          api.SelectedMapList(settingsList(d.Whitelists)),
		Forcesafegoogle:     settingsBool(d.Forcesafegoogle),
		Forcesafeduckduckgo: settingsBool(d.Forcesafeduckduckgo),
		Forcesafeyoutube:    settingsBool(d.Forcesafeyoutube),
		Forcestrictbing:     settingsBool(d.Forcestrictbing),
	}
}

func convertDnsblSettingsStructToSchema(s *bind.DnsblSettings) *dnsblSettingsResourceModel {
	return &dnsblSettingsResourceModel{
		Enabled:             types.BoolValue(tools.StringToBool(s.Enabled)),
		Type:                tools.StringSliceToSet(settingsNonEmpty(s.Type)),
		Whitelists:          tools.StringSliceToSet(settingsNonEmpty(s.Whitelists)),
		Forcesafegoogle:     types.BoolValue(tools.StringToBool(s.Forcesafegoogle)),
		Forcesafeduckduckgo: types.BoolValue(tools.StringToBool(s.Forcesafeduckduckgo)),
		Forcesafeyoutube:    types.BoolValue(tools.StringToBool(s.Forcesafeyoutube)),
		Forcestrictbing:     types.BoolValue(tools.StringToBool(s.Forcestrictbing)),
	}
}
