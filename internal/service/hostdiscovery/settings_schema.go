package hostdiscovery

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/hostdiscovery"
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
	Enabled         types.Bool  `tfsdk:"enabled"`
	Promiscuous     types.Bool  `tfsdk:"promiscuous"`
	Verbose         types.Bool  `tfsdk:"verbose"`
	SkipNetworks    types.Set   `tfsdk:"skip_networks"`
	Expire4Interval types.Int64 `tfsdk:"expire4_interval"`
	Expire6Interval types.Int64 `tfsdk:"expire6_interval"`
	Interfaces      types.Set   `tfsdk:"interfaces"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable host discovery (hostwatch), which records the hosts seen on the interfaces.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"promiscuous": schema.BoolAttribute{
			MarkdownDescription: "Capture in promiscuous mode.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"verbose": schema.BoolAttribute{
			MarkdownDescription: "Verbose logging.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"skip_networks": schema.SetAttribute{
			MarkdownDescription: "Networks that are not recorded.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"expire4_interval": schema.Int64Attribute{
			MarkdownDescription: "Seconds before an IPv4 host is forgotten (3600 or more). `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"expire6_interval": schema.Int64Attribute{
			MarkdownDescription: "Seconds before an IPv6 host is forgotten (3600 or more). `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"interfaces": schema.SetAttribute{
			MarkdownDescription: "Interfaces to watch, empty for all.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `hostdiscovery_settings`.",
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

func convertSettingsSchemaToStruct(d *settingsResourceModel) *hostdiscovery.Settings {
	return &hostdiscovery.Settings{
		General: hostdiscovery.SettingsGeneral{
			Enabled:         settingsBool(d.Enabled),
			Promiscuous:     settingsBool(d.Promiscuous),
			Verbose:         settingsBool(d.Verbose),
			SkipNetworks:    api.SelectedMapList(settingsList(d.SkipNetworks)),
			Expire4Interval: settingsInt(d.Expire4Interval),
			Expire6Interval: settingsInt(d.Expire6Interval),
			Interfaces:      api.SelectedMapList(settingsList(d.Interfaces)),
		},
	}
}

func convertSettingsStructToSchema(s *hostdiscovery.Settings) *settingsResourceModel {
	return &settingsResourceModel{
		Enabled:         types.BoolValue(tools.StringToBool(s.General.Enabled)),
		Promiscuous:     types.BoolValue(tools.StringToBool(s.General.Promiscuous)),
		Verbose:         types.BoolValue(tools.StringToBool(s.General.Verbose)),
		SkipNetworks:    tools.StringSliceToSet(settingsNonEmpty(s.General.SkipNetworks)),
		Expire4Interval: settingsReadInt(s.General.Expire4Interval),
		Expire6Interval: settingsReadInt(s.General.Expire6Interval),
		Interfaces:      tools.StringSliceToSet(settingsNonEmpty(s.General.Interfaces)),
	}
}
