package nodeexporter

import (
	"github.com/browningluke/opnsense-go/pkg/nodeexporter"
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
	Enabled       types.Bool   `tfsdk:"enabled"`
	ListenAddress types.String `tfsdk:"listen_address"`
	ListenPort    types.Int64  `tfsdk:"listen_port"`
	Cpu           types.Bool   `tfsdk:"cpu"`
	Exec          types.Bool   `tfsdk:"exec"`
	Filesystem    types.Bool   `tfsdk:"filesystem"`
	Loadavg       types.Bool   `tfsdk:"loadavg"`
	Meminfo       types.Bool   `tfsdk:"meminfo"`
	Netdev        types.Bool   `tfsdk:"netdev"`
	Time          types.Bool   `tfsdk:"time"`
	Devstat       types.Bool   `tfsdk:"devstat"`
	Interrupts    types.Bool   `tfsdk:"interrupts"`
	Ntp           types.Bool   `tfsdk:"ntp"`
	Zfs           types.Bool   `tfsdk:"zfs"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable the Prometheus node exporter.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"listen_address": schema.StringAttribute{
			MarkdownDescription: "Address to listen on, `0.0.0.0` for all.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"listen_port": schema.Int64Attribute{
			MarkdownDescription: "Port to listen on. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"cpu": schema.BoolAttribute{
			MarkdownDescription: "Enable the cpu collector.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"exec": schema.BoolAttribute{
			MarkdownDescription: "Enable the exec collector.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"filesystem": schema.BoolAttribute{
			MarkdownDescription: "Enable the filesystem collector.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"loadavg": schema.BoolAttribute{
			MarkdownDescription: "Enable the loadavg collector.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"meminfo": schema.BoolAttribute{
			MarkdownDescription: "Enable the meminfo collector.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"netdev": schema.BoolAttribute{
			MarkdownDescription: "Enable the netdev collector.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"time": schema.BoolAttribute{
			MarkdownDescription: "Enable the time collector.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"devstat": schema.BoolAttribute{
			MarkdownDescription: "Enable the devstat collector.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"interrupts": schema.BoolAttribute{
			MarkdownDescription: "Enable the interrupts collector.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"ntp": schema.BoolAttribute{
			MarkdownDescription: "Enable the ntp collector.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"zfs": schema.BoolAttribute{
			MarkdownDescription: "Enable the zfs collector.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `nodeexporter_settings`.",
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

func convertSettingsSchemaToStruct(d *settingsResourceModel) *nodeexporter.Settings {
	return &nodeexporter.Settings{
		Enabled:       settingsBool(d.Enabled),
		ListenAddress: d.ListenAddress.ValueString(),
		ListenPort:    settingsInt(d.ListenPort),
		Cpu:           settingsBool(d.Cpu),
		Exec:          settingsBool(d.Exec),
		Filesystem:    settingsBool(d.Filesystem),
		Loadavg:       settingsBool(d.Loadavg),
		Meminfo:       settingsBool(d.Meminfo),
		Netdev:        settingsBool(d.Netdev),
		Time:          settingsBool(d.Time),
		Devstat:       settingsBool(d.Devstat),
		Interrupts:    settingsBool(d.Interrupts),
		Ntp:           settingsBool(d.Ntp),
		Zfs:           settingsBool(d.Zfs),
	}
}

func convertSettingsStructToSchema(s *nodeexporter.Settings) *settingsResourceModel {
	return &settingsResourceModel{
		Enabled:       types.BoolValue(tools.StringToBool(s.Enabled)),
		ListenAddress: types.StringValue(s.ListenAddress),
		ListenPort:    settingsReadInt(s.ListenPort),
		Cpu:           types.BoolValue(tools.StringToBool(s.Cpu)),
		Exec:          types.BoolValue(tools.StringToBool(s.Exec)),
		Filesystem:    types.BoolValue(tools.StringToBool(s.Filesystem)),
		Loadavg:       types.BoolValue(tools.StringToBool(s.Loadavg)),
		Meminfo:       types.BoolValue(tools.StringToBool(s.Meminfo)),
		Netdev:        types.BoolValue(tools.StringToBool(s.Netdev)),
		Time:          types.BoolValue(tools.StringToBool(s.Time)),
		Devstat:       types.BoolValue(tools.StringToBool(s.Devstat)),
		Interrupts:    types.BoolValue(tools.StringToBool(s.Interrupts)),
		Ntp:           types.BoolValue(tools.StringToBool(s.Ntp)),
		Zfs:           types.BoolValue(tools.StringToBool(s.Zfs)),
	}
}
