package telegraf

import (
	"github.com/browningluke/opnsense-go/pkg/telegraf"
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
	Enabled           types.Bool   `tfsdk:"enabled"`
	Wheelgroup        types.Bool   `tfsdk:"wheelgroup"`
	Interval          types.Int64  `tfsdk:"interval"`
	Roundinterval     types.Bool   `tfsdk:"roundinterval"`
	MetricBatchSize   types.Int64  `tfsdk:"metric_batch_size"`
	MetricBufferLimit types.Int64  `tfsdk:"metric_buffer_limit"`
	CollectionJitter  types.Int64  `tfsdk:"collection_jitter"`
	FlushInterval     types.Int64  `tfsdk:"flush_interval"`
	FlushJitter       types.Int64  `tfsdk:"flush_jitter"`
	Hostname          types.String `tfsdk:"hostname"`
	OmitHostname      types.Bool   `tfsdk:"omit_hostname"`
	Quiet             types.Bool   `tfsdk:"quiet"`
	Debug             types.Bool   `tfsdk:"debug"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "This will activate the telegraf agent.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"wheelgroup": schema.BoolAttribute{
			MarkdownDescription: "This will start the process with wheel group and root user permission. Please use this with care, currently only needed for Unbound and Suricata.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"interval": schema.Int64Attribute{
			MarkdownDescription: "Default data collection interval for all inputs in seconds. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"roundinterval": schema.BoolAttribute{
			MarkdownDescription: "Rounds collection interval to 'interval', e.g. if interval=\"10s\" then always collect on :00, :10, :20, etc.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"metric_batch_size": schema.Int64Attribute{
			MarkdownDescription: "Telegraf will send metrics to outputs in batches of at most metric_batch_size metrics. This controls the size of writes that Telegraf sends to output plugins. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"metric_buffer_limit": schema.Int64Attribute{
			MarkdownDescription: "For failed writes, Telegraf will cache metric_buffer_limit metrics for each output, and will flush this buffer on a successful write. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"collection_jitter": schema.Int64Attribute{
			MarkdownDescription: "Collection jitter is used to jitter the collection by a random amount. Each plugin will sleep for a random time within jitter before collecting. Input value are seconds. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"flush_interval": schema.Int64Attribute{
			MarkdownDescription: "Default flushing interval for all outputs. You should not set this below interval. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"flush_jitter": schema.Int64Attribute{
			MarkdownDescription: "Jitter the flush interval by a random amount. This is primarily to avoid large write spikes for users running a large number of telegraf instances. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"hostname": schema.StringAttribute{
			MarkdownDescription: "Override default hostname, if empty use system hostname.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"omit_hostname": schema.BoolAttribute{
			MarkdownDescription: "If set to true, do no set the \"host\" tag in the Telegraf agent.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"quiet": schema.BoolAttribute{
			MarkdownDescription: "Log only error level messages.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"debug": schema.BoolAttribute{
			MarkdownDescription: "Log at debug level.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `telegraf_settings`.",
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

func convertSettingsSchemaToStruct(d *settingsResourceModel) *telegraf.Settings {
	return &telegraf.Settings{
		Enabled:           settingsBool(d.Enabled),
		Wheelgroup:        settingsBool(d.Wheelgroup),
		Interval:          settingsInt(d.Interval),
		Roundinterval:     settingsBool(d.Roundinterval),
		MetricBatchSize:   settingsInt(d.MetricBatchSize),
		MetricBufferLimit: settingsInt(d.MetricBufferLimit),
		CollectionJitter:  settingsInt(d.CollectionJitter),
		FlushInterval:     settingsInt(d.FlushInterval),
		FlushJitter:       settingsInt(d.FlushJitter),
		Hostname:          d.Hostname.ValueString(),
		OmitHostname:      settingsBool(d.OmitHostname),
		Quiet:             settingsBool(d.Quiet),
		Debug:             settingsBool(d.Debug),
	}
}

func convertSettingsStructToSchema(s *telegraf.Settings) *settingsResourceModel {
	return &settingsResourceModel{
		Enabled:           types.BoolValue(tools.StringToBool(s.Enabled)),
		Wheelgroup:        types.BoolValue(tools.StringToBool(s.Wheelgroup)),
		Interval:          settingsReadInt(s.Interval),
		Roundinterval:     types.BoolValue(tools.StringToBool(s.Roundinterval)),
		MetricBatchSize:   settingsReadInt(s.MetricBatchSize),
		MetricBufferLimit: settingsReadInt(s.MetricBufferLimit),
		CollectionJitter:  settingsReadInt(s.CollectionJitter),
		FlushInterval:     settingsReadInt(s.FlushInterval),
		FlushJitter:       settingsReadInt(s.FlushJitter),
		Hostname:          types.StringValue(s.Hostname),
		OmitHostname:      types.BoolValue(tools.StringToBool(s.OmitHostname)),
		Quiet:             types.BoolValue(tools.StringToBool(s.Quiet)),
		Debug:             types.BoolValue(tools.StringToBool(s.Debug)),
	}
}
