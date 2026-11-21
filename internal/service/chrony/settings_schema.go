package chrony

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/chrony"
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
	Enabled         types.Bool   `tfsdk:"enabled"`
	Port            types.Int64  `tfsdk:"port"`
	NtsClient       types.Bool   `tfsdk:"nts_client"`
	NtsNoCertCheck  types.Bool   `tfsdk:"nts_no_cert_check"`
	Peers           types.Set    `tfsdk:"peers"`
	FallbackPeers   types.String `tfsdk:"fallback_peers"`
	AllowedNetworks types.Set    `tfsdk:"allowed_networks"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable chrony. Disable the built-in NTP server first, both cannot listen at the same time.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"port": schema.Int64Attribute{
			MarkdownDescription: "Command port of chronyc. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"nts_client": schema.BoolAttribute{
			MarkdownDescription: "Use Network Time Security (NTS) with the peers.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"nts_no_cert_check": schema.BoolAttribute{
			MarkdownDescription: "Do not check the time validity of the NTS server certificates.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"peers": schema.SetAttribute{
			MarkdownDescription: "NTP or NTS servers, e.g. `[\"time.cloudflare.com\"]`.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"fallback_peers": schema.StringAttribute{
			MarkdownDescription: "Server used when the peers are unreachable.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"allowed_networks": schema.SetAttribute{
			MarkdownDescription: "Networks allowed to use this server for time.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `chrony_settings`.",
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

func convertSettingsSchemaToStruct(d *settingsResourceModel) *chrony.Settings {
	return &chrony.Settings{
		Enabled:         settingsBool(d.Enabled),
		Port:            settingsInt(d.Port),
		NtsClient:       settingsBool(d.NtsClient),
		NtsNoCertCheck:  settingsBool(d.NtsNoCertCheck),
		Peers:           api.SelectedMapList(settingsList(d.Peers)),
		FallbackPeers:   d.FallbackPeers.ValueString(),
		AllowedNetworks: api.SelectedMapList(settingsList(d.AllowedNetworks)),
	}
}

func convertSettingsStructToSchema(s *chrony.Settings) *settingsResourceModel {
	return &settingsResourceModel{
		Enabled:         types.BoolValue(tools.StringToBool(s.Enabled)),
		Port:            settingsReadInt(s.Port),
		NtsClient:       types.BoolValue(tools.StringToBool(s.NtsClient)),
		NtsNoCertCheck:  types.BoolValue(tools.StringToBool(s.NtsNoCertCheck)),
		Peers:           tools.StringSliceToSet(settingsNonEmpty(s.Peers)),
		FallbackPeers:   types.StringValue(s.FallbackPeers),
		AllowedNetworks: tools.StringSliceToSet(settingsNonEmpty(s.AllowedNetworks)),
	}
}
