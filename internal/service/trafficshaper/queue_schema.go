package trafficshaper

import (
	"maps"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/trafficshaper"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type queueResourceModel struct {
	Enabled        types.Bool   `tfsdk:"enabled"`
	Number         types.Int64  `tfsdk:"number"`
	Pipe           types.String `tfsdk:"pipe"`
	Weight         types.Int64  `tfsdk:"weight"`
	Mask           types.String `tfsdk:"mask"`
	Buckets        types.Int64  `tfsdk:"buckets"`
	CodelEnable    types.Bool   `tfsdk:"codel_enable"`
	CodelTarget    types.Int64  `tfsdk:"codel_target"`
	CodelInterval  types.Int64  `tfsdk:"codel_interval"`
	CodelECNEnable types.Bool   `tfsdk:"codel_ecn_enable"`
	PIEEnable      types.Bool   `tfsdk:"pie_enable"`
	Description    types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func queueAttributes() map[string]schema.Attribute {
	attrs := map[string]schema.Attribute{
		"enabled": enabledAttribute(),
		"number":  numberAttribute("queue"),
		"pipe": schema.StringAttribute{
			MarkdownDescription: "UUID of the pipe this queue draws its bandwidth from.",
			Required:            true,
		},
		"weight": schema.Int64Attribute{
			MarkdownDescription: "Share of the pipe bandwidth this queue gets relative to the other queues on the same pipe, from `1` to `100`. Defaults to `100`.",
			Optional:            true,
			Computed:            true,
			Default:             int64default.StaticInt64(100),
			Validators:          []validator.Int64{int64validator.Between(1, 100)},
		},
		"mask":    maskAttribute(),
		"buckets": optionalSentinelInt64("Number of hash buckets for dynamic queues.", 1, 65535),
		"description": schema.StringAttribute{
			MarkdownDescription: "Description of the queue. Rules refer to the queue by UUID, the GUI shows this text.",
			Required:            true,
		},
		"id": idAttribute("queue"),
	}
	maps.Copy(attrs, codelAttributes())
	return attrs
}

func queueResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Traffic shaper queues share the bandwidth of a pipe between traffic classes by weight.",
		Version:             1,
		Attributes:          queueAttributes(),
	}
}

func queueDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Traffic shaper queues share the bandwidth of a pipe between traffic classes by weight.",
		Attributes:          dataSourceComputed(queueAttributes()),
	}
}

func convertQueueSchemaToStruct(d *queueResourceModel) (*trafficshaper.Queue, error) {
	number := ""
	if !d.Number.IsNull() && !d.Number.IsUnknown() {
		number = tools.Int64ToString(d.Number.ValueInt64())
	}
	return &trafficshaper.Queue{
		Number:         number,
		Enabled:        tools.BoolToString(d.Enabled.ValueBool()),
		Pipe:           api.SelectedMap(d.Pipe.ValueString()),
		Weight:         tools.Int64ToString(d.Weight.ValueInt64()),
		Mask:           api.SelectedMap(d.Mask.ValueString()),
		Buckets:        tools.Int64ToStringNegative(d.Buckets.ValueInt64()),
		CodelEnable:    tools.BoolToString(d.CodelEnable.ValueBool()),
		CodelTarget:    tools.Int64ToStringNegative(d.CodelTarget.ValueInt64()),
		CodelInterval:  tools.Int64ToStringNegative(d.CodelInterval.ValueInt64()),
		CodelECNEnable: tools.BoolToString(d.CodelECNEnable.ValueBool()),
		PIEEnable:      tools.BoolToString(d.PIEEnable.ValueBool()),
		Description:    d.Description.ValueString(),
	}, nil
}

func convertQueueStructToSchema(d *trafficshaper.Queue) (*queueResourceModel, error) {
	return &queueResourceModel{
		Enabled:        types.BoolValue(tools.StringToBool(d.Enabled)),
		Number:         tools.StringToInt64Null(d.Number),
		Pipe:           types.StringValue(d.Pipe.String()),
		Weight:         types.Int64Value(tools.StringToInt64(d.Weight)),
		Mask:           types.StringValue(d.Mask.String()),
		Buckets:        types.Int64Value(tools.StringToInt64(d.Buckets)),
		CodelEnable:    types.BoolValue(tools.StringToBool(d.CodelEnable)),
		CodelTarget:    types.Int64Value(tools.StringToInt64(d.CodelTarget)),
		CodelInterval:  types.Int64Value(tools.StringToInt64(d.CodelInterval)),
		CodelECNEnable: types.BoolValue(tools.StringToBool(d.CodelECNEnable)),
		PIEEnable:      types.BoolValue(tools.StringToBool(d.PIEEnable)),
		Description:    types.StringValue(d.Description),
	}, nil
}
