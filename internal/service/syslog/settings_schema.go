package syslog

import (
	"github.com/browningluke/opnsense-go/pkg/syslog"
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
	Enabled     types.Bool  `tfsdk:"enabled"`
	Loglocal    types.Bool  `tfsdk:"loglocal"`
	Maxpreserve types.Int64 `tfsdk:"maxpreserve"`
	Maxfilesize types.Int64 `tfsdk:"maxfilesize"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "enabled.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"loglocal": schema.BoolAttribute{
			MarkdownDescription: "Enable writing log files to the local disk.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"maxpreserve": schema.Int64Attribute{
			MarkdownDescription: "Number of logs to preserve. When no maximum file size is offered or the logs are smaller than the the size requested, this equals the number of days. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"maxfilesize": schema.Int64Attribute{
			MarkdownDescription: "Maximum file size per log file. When set and a log file exceeds the amount specified, it will be rotated. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `syslog_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertSettingsSchemaToStruct(d *settingsResourceModel) *syslog.SyslogConfig {
	return &syslog.SyslogConfig{
		General: syslog.SyslogConfigGeneral{
			Enabled:     settingsBool(d.Enabled),
			Loglocal:    settingsBool(d.Loglocal),
			Maxpreserve: settingsInt(d.Maxpreserve),
			Maxfilesize: settingsInt(d.Maxfilesize),
		},
	}
}

func convertSettingsStructToSchema(s *syslog.SyslogConfig) *settingsResourceModel {
	return &settingsResourceModel{
		Enabled:     types.BoolValue(tools.StringToBool(s.General.Enabled)),
		Loglocal:    types.BoolValue(tools.StringToBool(s.General.Loglocal)),
		Maxpreserve: settingsReadInt(s.General.Maxpreserve),
		Maxfilesize: settingsReadInt(s.General.Maxfilesize),
	}
}
