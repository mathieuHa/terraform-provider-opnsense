package kea

import (
	"github.com/browningluke/opnsense-go/pkg/kea"
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

type ddnsSettingsResourceModel struct {
	Enabled      types.Bool   `tfsdk:"enabled"`
	ManualConfig types.Bool   `tfsdk:"manual_config"`
	ServerIp     types.String `tfsdk:"server_ip"`
	ServerPort   types.Int64  `tfsdk:"server_port"`

	Id types.String `tfsdk:"id"`
}

func ddnsSettingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable the Kea DHCP-DDNS service, which sends DNS updates for the subnets that have `ddns_forward_zone` set.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"manual_config": schema.BoolAttribute{
			MarkdownDescription: "Use a hand-written kea-dhcp-ddns.conf instead of the generated one.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"server_ip": schema.StringAttribute{
			MarkdownDescription: "IPv4 address the DDNS service listens on for requests from the DHCP servers.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"server_port": schema.Int64Attribute{
			MarkdownDescription: "Port the DDNS service listens on. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `kea_ddns_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertDdnsSettingsSchemaToStruct(d *ddnsSettingsResourceModel) *kea.DdnsSettings {
	return &kea.DdnsSettings{
		General: kea.DdnsSettingsGeneral{
			Enabled:      settingsBool(d.Enabled),
			ManualConfig: settingsBool(d.ManualConfig),
			ServerIp:     d.ServerIp.ValueString(),
			ServerPort:   settingsInt(d.ServerPort),
		},
	}
}

func convertDdnsSettingsStructToSchema(s *kea.DdnsSettings) *ddnsSettingsResourceModel {
	return &ddnsSettingsResourceModel{
		Enabled:      types.BoolValue(tools.StringToBool(s.General.Enabled)),
		ManualConfig: types.BoolValue(tools.StringToBool(s.General.ManualConfig)),
		ServerIp:     types.StringValue(s.General.ServerIp),
		ServerPort:   settingsReadInt(s.General.ServerPort),
	}
}
