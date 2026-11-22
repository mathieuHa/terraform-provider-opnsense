package diagnostics

import (
	"github.com/browningluke/opnsense-go/pkg/diagnostics"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type systemhealthSettingsResourceModel struct {
	Enabled types.Bool `tfsdk:"enabled"`

	Id types.String `tfsdk:"id"`
}

func systemhealthSettingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Enables the RRD graphing backend. Note that if disabled, existing RRD data is not deleted. When enabled, starts a cron job that runs every minute, collecting helpful system metrics to be shown in the graph below.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `systemhealth_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertSystemhealthSettingsSchemaToStruct(d *systemhealthSettingsResourceModel) *diagnostics.SystemHealthConfig {
	return &diagnostics.SystemHealthConfig{
		Enabled: settingsBool(d.Enabled),
	}
}

func convertSystemhealthSettingsStructToSchema(s *diagnostics.SystemHealthConfig) *systemhealthSettingsResourceModel {
	return &systemhealthSettingsResourceModel{
		Enabled: types.BoolValue(tools.StringToBool(s.Enabled)),
	}
}
