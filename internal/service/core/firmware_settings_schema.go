package core

import (
	"github.com/browningluke/opnsense-go/pkg/core"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type firmwareSettingsResourceModel struct {
	Mirror       types.String `tfsdk:"mirror"`
	Flavour      types.String `tfsdk:"flavour"`
	Type         types.String `tfsdk:"type"`
	Subscription types.String `tfsdk:"subscription"`

	Id types.String `tfsdk:"id"`
}

func firmwareSettingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"mirror": schema.StringAttribute{
			MarkdownDescription: "Package mirror URL, e.g. `https://pkg.opnsense.org`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"flavour": schema.StringAttribute{
			MarkdownDescription: "Package flavour, empty for the default.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"type": schema.StringAttribute{
			MarkdownDescription: "Release type, empty for the community release.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"subscription": schema.StringAttribute{
			MarkdownDescription: "Subscription key for business mirrors.",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `firmware_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertFirmwareSettingsSchemaToStruct(d *firmwareSettingsResourceModel) *core.FirmwareSettings {
	return &core.FirmwareSettings{
		Mirror:       d.Mirror.ValueString(),
		Flavour:      d.Flavour.ValueString(),
		Type:         d.Type.ValueString(),
		Subscription: d.Subscription.ValueString(),
	}
}

func convertFirmwareSettingsStructToSchema(s *core.FirmwareSettings) *firmwareSettingsResourceModel {
	return &firmwareSettingsResourceModel{
		Mirror:       types.StringValue(s.Mirror),
		Flavour:      types.StringValue(s.Flavour),
		Type:         types.StringValue(s.Type),
		Subscription: types.StringValue(s.Subscription),
	}
}
