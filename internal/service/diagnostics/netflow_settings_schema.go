package diagnostics

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/diagnostics"
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

type netflowSettingsResourceModel struct {
	CaptureInterfaces types.Set    `tfsdk:"capture_interfaces"`
	CaptureEgressOnly types.Set    `tfsdk:"capture_egress_only"`
	CaptureVersion    types.String `tfsdk:"capture_version"`
	CaptureTargets    types.Set    `tfsdk:"capture_targets"`
	CollectEnable     types.Bool   `tfsdk:"collect_enable"`
	ActiveTimeout     types.Int64  `tfsdk:"active_timeout"`
	InactiveTimeout   types.Int64  `tfsdk:"inactive_timeout"`

	Id types.String `tfsdk:"id"`
}

func netflowSettingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"capture_interfaces": schema.SetAttribute{
			MarkdownDescription: "Select all the interfaces to enable NetFlow on.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"capture_egress_only": schema.SetAttribute{
			MarkdownDescription: "Select interfaces used for WAN traffic to avoid counting NAT traffic twice.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"capture_version": schema.StringAttribute{
			MarkdownDescription: "Select NetFlow version to use. One of `v5`, `v9`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"capture_targets": schema.SetAttribute{
			MarkdownDescription: "Select destinations to send NetFlow data to (ip address:port, eg 192.168.0.1:2550).",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"collect_enable": schema.BoolAttribute{
			MarkdownDescription: "Collect NetFlow data on this firewall for use with Insight. Note that the local cache only holds the latest 100 MB of data.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"active_timeout": schema.Int64Attribute{
			MarkdownDescription: "Split long running flows into smaller parts. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"inactive_timeout": schema.Int64Attribute{
			MarkdownDescription: "Expire idle flows. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `netflow_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertNetflowSettingsSchemaToStruct(d *netflowSettingsResourceModel) *diagnostics.NetflowConfig {
	return &diagnostics.NetflowConfig{
		ActiveTimeout:   settingsInt(d.ActiveTimeout),
		InactiveTimeout: settingsInt(d.InactiveTimeout),
		Capture: diagnostics.NetflowConfigCapture{
			CaptureInterfaces: api.SelectedMapList(settingsList(d.CaptureInterfaces)),
			CaptureEgressOnly: api.SelectedMapList(settingsList(d.CaptureEgressOnly)),
			CaptureVersion:    api.SelectedMap(d.CaptureVersion.ValueString()),
			CaptureTargets:    api.SelectedMapList(settingsList(d.CaptureTargets)),
		},
		Collect: diagnostics.NetflowConfigCollect{
			CollectEnable: settingsBool(d.CollectEnable),
		},
	}
}

func convertNetflowSettingsStructToSchema(s *diagnostics.NetflowConfig) *netflowSettingsResourceModel {
	return &netflowSettingsResourceModel{
		CaptureInterfaces: tools.StringSliceToSet(settingsNonEmpty(s.Capture.CaptureInterfaces)),
		CaptureEgressOnly: tools.StringSliceToSet(settingsNonEmpty(s.Capture.CaptureEgressOnly)),
		CaptureVersion:    types.StringValue(s.Capture.CaptureVersion.String()),
		CaptureTargets:    tools.StringSliceToSet(settingsNonEmpty(s.Capture.CaptureTargets)),
		CollectEnable:     types.BoolValue(tools.StringToBool(s.Collect.CollectEnable)),
		ActiveTimeout:     settingsReadInt(s.ActiveTimeout),
		InactiveTimeout:   settingsReadInt(s.InactiveTimeout),
	}
}
