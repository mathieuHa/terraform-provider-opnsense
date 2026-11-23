package dyndns

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dyndns"
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
	Enabled     types.Bool   `tfsdk:"enabled"`
	Verbose     types.Bool   `tfsdk:"verbose"`
	Allowipv6   types.Bool   `tfsdk:"allowipv6"`
	DaemonDelay types.Int64  `tfsdk:"daemon_delay"`
	Backend     types.String `tfsdk:"backend"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable ddclient.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"verbose": schema.BoolAttribute{
			MarkdownDescription: "Enable verbose logging.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"allowipv6": schema.BoolAttribute{
			MarkdownDescription: "Allow IPv6 for updates.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"daemon_delay": schema.Int64Attribute{
			MarkdownDescription: "Interval in seconds to check for address changes. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"backend": schema.StringAttribute{
			MarkdownDescription: "Select the backend to use. One of `ddclient`, `opnsense`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `dyndns_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertSettingsSchemaToStruct(d *settingsResourceModel) *dyndns.GeneralSettings {
	return &dyndns.GeneralSettings{
		General: dyndns.GeneralSettingsGeneral{
			Enabled:     settingsBool(d.Enabled),
			Verbose:     settingsBool(d.Verbose),
			Allowipv6:   settingsBool(d.Allowipv6),
			DaemonDelay: settingsInt(d.DaemonDelay),
			Backend:     api.SelectedMap(d.Backend.ValueString()),
		},
	}
}

func convertSettingsStructToSchema(s *dyndns.GeneralSettings) *settingsResourceModel {
	return &settingsResourceModel{
		Enabled:     types.BoolValue(tools.StringToBool(s.General.Enabled)),
		Verbose:     types.BoolValue(tools.StringToBool(s.General.Verbose)),
		Allowipv6:   types.BoolValue(tools.StringToBool(s.General.Allowipv6)),
		DaemonDelay: settingsReadInt(s.General.DaemonDelay),
		Backend:     types.StringValue(s.General.Backend.String()),
	}
}
