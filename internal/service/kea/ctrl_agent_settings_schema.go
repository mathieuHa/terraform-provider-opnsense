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

type ctrlAgentSettingsResourceModel struct {
	Enabled  types.Bool   `tfsdk:"enabled"`
	HttpHost types.String `tfsdk:"http_host"`
	HttpPort types.Int64  `tfsdk:"http_port"`

	Id types.String `tfsdk:"id"`
}

func ctrlAgentSettingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable the Kea control agent (REST API used by HA peers).",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"http_host": schema.StringAttribute{
			MarkdownDescription: "IPv4 address the control agent listens on.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"http_port": schema.Int64Attribute{
			MarkdownDescription: "Port the control agent listens on. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `kea_ctrl_agent_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertCtrlAgentSettingsSchemaToStruct(d *ctrlAgentSettingsResourceModel) *kea.CtrlAgentSettings {
	return &kea.CtrlAgentSettings{
		General: kea.CtrlAgentSettingsGeneral{
			Enabled:  settingsBool(d.Enabled),
			HttpHost: d.HttpHost.ValueString(),
			HttpPort: settingsInt(d.HttpPort),
		},
	}
}

func convertCtrlAgentSettingsStructToSchema(s *kea.CtrlAgentSettings) *ctrlAgentSettingsResourceModel {
	return &ctrlAgentSettingsResourceModel{
		Enabled:  types.BoolValue(tools.StringToBool(s.General.Enabled)),
		HttpHost: types.StringValue(s.General.HttpHost),
		HttpPort: settingsReadInt(s.General.HttpPort),
	}
}
