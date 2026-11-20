package trafficshaper

import (
	"maps"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/trafficshaper"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type pipeResourceModel struct {
	Enabled         types.Bool   `tfsdk:"enabled"`
	Number          types.Int64  `tfsdk:"number"`
	Bandwidth       types.Int64  `tfsdk:"bandwidth"`
	BandwidthMetric types.String `tfsdk:"bandwidth_metric"`
	Queue           types.Int64  `tfsdk:"queue"`
	Mask            types.String `tfsdk:"mask"`
	Buckets         types.Int64  `tfsdk:"buckets"`
	Scheduler       types.String `tfsdk:"scheduler"`
	CodelEnable     types.Bool   `tfsdk:"codel_enable"`
	CodelTarget     types.Int64  `tfsdk:"codel_target"`
	CodelInterval   types.Int64  `tfsdk:"codel_interval"`
	CodelECNEnable  types.Bool   `tfsdk:"codel_ecn_enable"`
	PIEEnable       types.Bool   `tfsdk:"pie_enable"`
	FQCodelQuantum  types.Int64  `tfsdk:"fqcodel_quantum"`
	FQCodelLimit    types.Int64  `tfsdk:"fqcodel_limit"`
	FQCodelFlows    types.Int64  `tfsdk:"fqcodel_flows"`
	Delay           types.Int64  `tfsdk:"delay"`
	Description     types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func pipeAttributes() map[string]schema.Attribute {
	attrs := map[string]schema.Attribute{
		"enabled": enabledAttribute(),
		"number":  numberAttribute("pipe"),
		"bandwidth": schema.Int64Attribute{
			MarkdownDescription: "Total bandwidth of the pipe, in `bandwidth_metric` units.",
			Required:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(1)},
		},
		"bandwidth_metric": schema.StringAttribute{
			MarkdownDescription: "Unit of `bandwidth`: `bit`, `Kbit`, `Mbit` or `Gbit` (per second). Defaults to `Kbit`.",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("Kbit"),
			Validators: []validator.String{
				stringvalidator.OneOf("bit", "Kbit", "Mbit", "Gbit"),
			},
		},
		"queue":   optionalSentinelInt64("Queue size in slots.", 2, 100),
		"mask":    maskAttribute(),
		"buckets": optionalSentinelInt64("Number of hash buckets for dynamic queues.", 1, 65535),
		"scheduler": schema.StringAttribute{
			MarkdownDescription: "Scheduler type: `fifo`, `rr` (deficit round robin), `qfq`, `fq_codel` (FlowQueue-CoDel) or `fq_pie` (FlowQueue-PIE). Leave unset for weighted fair queueing.",
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("fifo", "rr", "qfq", "fq_codel", "fq_pie"),
			},
		},
		"fqcodel_quantum": optionalSentinelInt64("FQ-CoDel quantum in bytes.", 1, 65535),
		"fqcodel_limit":   optionalSentinelInt64("FQ-CoDel queue limit in packets.", 1, 65535),
		"fqcodel_flows":   optionalSentinelInt64("FQ-CoDel number of flows.", 1, 65535),
		"delay":           optionalSentinelInt64("Added delay in milliseconds.", 1, 3000),
		"description": schema.StringAttribute{
			MarkdownDescription: "Description of the pipe. Queues and rules refer to the pipe by UUID, the GUI shows this text.",
			Required:            true,
		},
		"id": idAttribute("pipe"),
	}
	maps.Copy(attrs, codelAttributes())
	return attrs
}

func pipeResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Traffic shaper pipes limit the total bandwidth available to the traffic sent into them, either directly by a rule or through queues.",
		Version:             1,
		Attributes:          pipeAttributes(),
	}
}

func pipeDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Traffic shaper pipes limit the total bandwidth available to the traffic sent into them, either directly by a rule or through queues.",
		Attributes:          dataSourceComputed(pipeAttributes()),
	}
}

func convertPipeSchemaToStruct(d *pipeResourceModel) (*trafficshaper.Pipe, error) {
	number := ""
	if !d.Number.IsNull() && !d.Number.IsUnknown() {
		number = tools.Int64ToString(d.Number.ValueInt64())
	}
	return &trafficshaper.Pipe{
		Number:          number,
		Enabled:         tools.BoolToString(d.Enabled.ValueBool()),
		Bandwidth:       tools.Int64ToString(d.Bandwidth.ValueInt64()),
		BandwidthMetric: api.SelectedMap(d.BandwidthMetric.ValueString()),
		Queue:           tools.Int64ToStringNegative(d.Queue.ValueInt64()),
		Mask:            api.SelectedMap(d.Mask.ValueString()),
		Buckets:         tools.Int64ToStringNegative(d.Buckets.ValueInt64()),
		Scheduler:       api.SelectedMap(d.Scheduler.ValueString()),
		CodelEnable:     tools.BoolToString(d.CodelEnable.ValueBool()),
		CodelTarget:     tools.Int64ToStringNegative(d.CodelTarget.ValueInt64()),
		CodelInterval:   tools.Int64ToStringNegative(d.CodelInterval.ValueInt64()),
		CodelECNEnable:  tools.BoolToString(d.CodelECNEnable.ValueBool()),
		PIEEnable:       tools.BoolToString(d.PIEEnable.ValueBool()),
		FQCodelQuantum:  tools.Int64ToStringNegative(d.FQCodelQuantum.ValueInt64()),
		FQCodelLimit:    tools.Int64ToStringNegative(d.FQCodelLimit.ValueInt64()),
		FQCodelFlows:    tools.Int64ToStringNegative(d.FQCodelFlows.ValueInt64()),
		Delay:           tools.Int64ToStringNegative(d.Delay.ValueInt64()),
		Description:     d.Description.ValueString(),
	}, nil
}

func convertPipeStructToSchema(d *trafficshaper.Pipe) (*pipeResourceModel, error) {
	return &pipeResourceModel{
		Enabled:         types.BoolValue(tools.StringToBool(d.Enabled)),
		Number:          tools.StringToInt64Null(d.Number),
		Bandwidth:       types.Int64Value(tools.StringToInt64(d.Bandwidth)),
		BandwidthMetric: types.StringValue(d.BandwidthMetric.String()),
		Queue:           types.Int64Value(tools.StringToInt64(d.Queue)),
		Mask:            types.StringValue(d.Mask.String()),
		Buckets:         types.Int64Value(tools.StringToInt64(d.Buckets)),
		Scheduler:       tools.StringOrNull(d.Scheduler.String()),
		CodelEnable:     types.BoolValue(tools.StringToBool(d.CodelEnable)),
		CodelTarget:     types.Int64Value(tools.StringToInt64(d.CodelTarget)),
		CodelInterval:   types.Int64Value(tools.StringToInt64(d.CodelInterval)),
		CodelECNEnable:  types.BoolValue(tools.StringToBool(d.CodelECNEnable)),
		PIEEnable:       types.BoolValue(tools.StringToBool(d.PIEEnable)),
		FQCodelQuantum:  types.Int64Value(tools.StringToInt64(d.FQCodelQuantum)),
		FQCodelLimit:    types.Int64Value(tools.StringToInt64(d.FQCodelLimit)),
		FQCodelFlows:    types.Int64Value(tools.StringToInt64(d.FQCodelFlows)),
		Delay:           types.Int64Value(tools.StringToInt64(d.Delay)),
		Description:     types.StringValue(d.Description),
	}, nil
}
