package trafficshaper

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/trafficshaper"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var dscpValues = []string{
	"be", "ef",
	"af11", "af12", "af13", "af21", "af22", "af23", "af31", "af32", "af33", "af41", "af42", "af43",
	"cs1", "cs2", "cs3", "cs4", "cs5", "cs6", "cs7",
}

type ruleResourceModel struct {
	Enabled         types.Bool   `tfsdk:"enabled"`
	Sequence        types.Int64  `tfsdk:"sequence"`
	Interface       types.String `tfsdk:"interface"`
	Interface2      types.String `tfsdk:"interface2"`
	Protocol        types.String `tfsdk:"protocol"`
	IPLength        types.Int64  `tfsdk:"ip_length"`
	Source          types.Set    `tfsdk:"source"`
	SourceNot       types.Bool   `tfsdk:"source_not"`
	SourcePort      types.String `tfsdk:"source_port"`
	Destination     types.Set    `tfsdk:"destination"`
	DestinationNot  types.Bool   `tfsdk:"destination_not"`
	DestinationPort types.String `tfsdk:"destination_port"`
	DSCP            types.Set    `tfsdk:"dscp"`
	Direction       types.String `tfsdk:"direction"`
	Target          types.String `tfsdk:"target"`
	Description     types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func anySetValue() types.Set {
	return types.SetValueMust(types.StringType, []attr.Value{types.StringValue("any")})
}

func networkSetAttribute(description string) schema.SetAttribute {
	return schema.SetAttribute{
		MarkdownDescription: description + " Defaults to `[\"any\"]`.",
		Optional:            true,
		Computed:            true,
		ElementType:         types.StringType,
		Default:             setdefault.StaticValue(anySetValue()),
		Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
	}
}

func portAttribute(description string) schema.StringAttribute {
	return schema.StringAttribute{
		MarkdownDescription: description + " A port, a range (`8000-8080`), a well-known name or `any`. Defaults to `any`.",
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("any"),
	}
}

func ruleAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": enabledAttribute(),
		"sequence": schema.Int64Attribute{
			MarkdownDescription: "Order in which rules are evaluated, lowest first; the first match wins. Defaults to `1`.",
			Optional:            true,
			Computed:            true,
			Default:             int64default.StaticInt64(1),
			Validators:          []validator.Int64{int64validator.Between(1, 1000000)},
		},
		"interface": schema.StringAttribute{
			MarkdownDescription: "Interface the rule matches on, e.g. `wan`.",
			Required:            true,
		},
		"interface2": schema.StringAttribute{
			MarkdownDescription: "Second interface, to only match traffic flowing between `interface` and this one.",
			Optional:            true,
		},
		"protocol": schema.StringAttribute{
			MarkdownDescription: "Protocol to match: `ip`, `ip4`, `ip6`, `udp`, `tcp`, `tcp_ack`, `tcp_ack_not`, `icmp`, `ipv6-icmp`, `igmp`, `esp`, `ah` or `gre`. Defaults to `ip`.",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("ip"),
			Validators: []validator.String{
				stringvalidator.OneOf("ip", "ip4", "ip6", "udp", "tcp", "tcp_ack", "tcp_ack_not", "icmp", "ipv6-icmp", "igmp", "esp", "ah", "gre"),
			},
		},
		"ip_length":        optionalSentinelInt64("Only match packets up to this size in bytes, e.g. small ACKs.", 2, 65535),
		"source":           networkSetAttribute("Source networks or addresses to match, or `any`."),
		"source_not":       optionalBool("When enabled, the source match is inverted. Defaults to `false`.", false),
		"source_port":      portAttribute("Source port to match."),
		"destination":      networkSetAttribute("Destination networks or addresses to match, or `any`."),
		"destination_not":  optionalBool("When enabled, the destination match is inverted. Defaults to `false`.", false),
		"destination_port": portAttribute("Destination port to match."),
		"dscp": schema.SetAttribute{
			MarkdownDescription: "DSCP values to match, e.g. `ef` or `af41`. Defaults to `[]` (any).",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf(dscpValues...))},
		},
		"direction": schema.StringAttribute{
			MarkdownDescription: "Traffic direction relative to `interface`: `in` or `out`. Leave unset to match both.",
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("in", "out"),
			},
		},
		"target": schema.StringAttribute{
			MarkdownDescription: "UUID of the pipe or queue matching traffic is sent to.",
			Required:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "Optional description here for your reference (not parsed).",
			Optional:            true,
		},
		"id": idAttribute("rule"),
	}
}

func ruleResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Traffic shaper rules select traffic and send it to a pipe or a queue.",
		Version:             1,
		Attributes:          ruleAttributes(),
	}
}

func ruleDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Traffic shaper rules select traffic and send it to a pipe or a queue.",
		Attributes:          dataSourceComputed(ruleAttributes()),
	}
}

func convertRuleSchemaToStruct(d *ruleResourceModel) (*trafficshaper.Rule, error) {
	return &trafficshaper.Rule{
		Enabled:         tools.BoolToString(d.Enabled.ValueBool()),
		Sequence:        tools.Int64ToString(d.Sequence.ValueInt64()),
		Interface:       api.SelectedMap(d.Interface.ValueString()),
		Interface2:      api.SelectedMap(d.Interface2.ValueString()),
		Protocol:        api.SelectedMap(d.Protocol.ValueString()),
		IPLength:        tools.Int64ToStringNegative(d.IPLength.ValueInt64()),
		Source:          api.SelectedMapList(tools.SetToStringSlice(d.Source)),
		SourceNot:       tools.BoolToString(d.SourceNot.ValueBool()),
		SourcePort:      d.SourcePort.ValueString(),
		Destination:     api.SelectedMapList(tools.SetToStringSlice(d.Destination)),
		DestinationNot:  tools.BoolToString(d.DestinationNot.ValueBool()),
		DestinationPort: d.DestinationPort.ValueString(),
		DSCP:            api.SelectedMapList(setOrEmpty(d.DSCP)),
		Direction:       api.SelectedMap(d.Direction.ValueString()),
		Target:          api.SelectedMap(d.Target.ValueString()),
		Description:     d.Description.ValueString(),
	}, nil
}

func setOrEmpty(s types.Set) []string {
	out := tools.SetToStringSlice(s)
	if out == nil {
		return []string{}
	}
	return out
}

func convertRuleStructToSchema(d *trafficshaper.Rule) (*ruleResourceModel, error) {
	return &ruleResourceModel{
		Enabled:         types.BoolValue(tools.StringToBool(d.Enabled)),
		Sequence:        types.Int64Value(tools.StringToInt64(d.Sequence)),
		Interface:       types.StringValue(d.Interface.String()),
		Interface2:      tools.StringOrNull(d.Interface2.String()),
		Protocol:        types.StringValue(d.Protocol.String()),
		IPLength:        types.Int64Value(tools.StringToInt64(d.IPLength)),
		Source:          tools.StringSliceToSet(d.Source),
		SourceNot:       types.BoolValue(tools.StringToBool(d.SourceNot)),
		SourcePort:      types.StringValue(d.SourcePort),
		Destination:     tools.StringSliceToSet(d.Destination),
		DestinationNot:  types.BoolValue(tools.StringToBool(d.DestinationNot)),
		DestinationPort: types.StringValue(d.DestinationPort),
		DSCP:            tools.StringSliceToSet(d.DSCP),
		Direction:       tools.StringOrNull(d.Direction.String()),
		Target:          types.StringValue(d.Target.String()),
		Description:     tools.StringOrNull(d.Description),
	}, nil
}
