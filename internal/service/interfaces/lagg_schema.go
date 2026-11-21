package interfaces

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/interfaces"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type laggResourceModel struct {
	Device          types.String `tfsdk:"device"`
	Members         types.Set    `tfsdk:"members"`
	PrimaryMember   types.String `tfsdk:"primary_member"`
	Proto           types.String `tfsdk:"proto"`
	LacpFastTimeout types.Bool   `tfsdk:"lacp_fast_timeout"`
	UseFlowid       types.String `tfsdk:"use_flowid"`
	Hash            types.Set    `tfsdk:"hash"`
	LacpStrict      types.String `tfsdk:"lacp_strict"`
	MTU             types.Int64  `tfsdk:"mtu"`
	Description     types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func laggResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Link aggregation devices (**Interfaces ▸ Devices ▸ LAGG**). Members must be unassigned physical interfaces. Never move the interface you manage OPNsense through into a LAGG without console access.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "Device name, e.g. `lagg0`. Assigned by OPNsense when unset.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"members": schema.SetAttribute{
				MarkdownDescription: "Member interfaces, e.g. `[\"vtnet2\", \"vtnet3\"]`.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"primary_member": schema.StringAttribute{
				MarkdownDescription: "Member preferred by the `failover` protocol.",
				Optional:            true,
			},
			"proto": schema.StringAttribute{
				MarkdownDescription: "Aggregation protocol. Defaults to `lacp`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("lacp"),
				Validators:          []validator.String{stringvalidator.OneOf("none", "lacp", "failover", "fec", "loadbalance", "roundrobin")},
			},
			"lacp_fast_timeout": schema.BoolAttribute{
				MarkdownDescription: "Use the fast LACP timeout. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"use_flowid": schema.StringAttribute{
				MarkdownDescription: "Use the RSS hash of the network card for balancing: `1`, `0`, or empty for the system default. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "1", "0")},
			},
			"hash": schema.SetAttribute{
				MarkdownDescription: "Packet layers hashed for balancing. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
				Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("l2", "l3", "l4"))},
			},
			"lacp_strict": schema.StringAttribute{
				MarkdownDescription: "LACP strict mode: `1`, `0`, or empty for the system default. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "1", "0")},
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: "MTU of the device. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.Between(576, 65535))},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the LAGG.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func laggDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Link aggregation devices (**Interfaces ▸ Devices ▸ LAGG**). Members must be unassigned physical interfaces. Never move the interface you manage OPNsense through into a LAGG without console access.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the LAGG.",
				Required:            true,
			},
			"device":            dschema.StringAttribute{MarkdownDescription: "Device name, e.g. `lagg0`. Assigned by OPNsense when unset.", Computed: true},
			"members":           dschema.SetAttribute{MarkdownDescription: "Member interfaces, e.g. `[\"vtnet2\", \"vtnet3\"]`.", Computed: true, ElementType: types.StringType},
			"primary_member":    dschema.StringAttribute{MarkdownDescription: "Member preferred by the `failover` protocol.", Computed: true},
			"proto":             dschema.StringAttribute{MarkdownDescription: "Aggregation protocol.", Computed: true},
			"lacp_fast_timeout": dschema.BoolAttribute{MarkdownDescription: "Use the fast LACP timeout.", Computed: true},
			"use_flowid":        dschema.StringAttribute{MarkdownDescription: "Use the RSS hash of the network card for balancing: `1`, `0`, or empty for the system default.", Computed: true},
			"hash":              dschema.SetAttribute{MarkdownDescription: "Packet layers hashed for balancing.", Computed: true, ElementType: types.StringType},
			"lacp_strict":       dschema.StringAttribute{MarkdownDescription: "LACP strict mode: `1`, `0`, or empty for the system default.", Computed: true},
			"mtu":               dschema.Int64Attribute{MarkdownDescription: "MTU of the device.", Computed: true},
			"description":       dschema.StringAttribute{MarkdownDescription: "Optional description here for your reference (not parsed).", Computed: true},
		},
	}
}

func convertLaggSchemaToStruct(d *laggResourceModel) (*interfaces.Lagg, error) {
	return &interfaces.Lagg{
		Device:          d.Device.ValueString(),
		Members:         api.SelectedMapList(setToSlice(d.Members)),
		PrimaryMember:   api.SelectedMap(d.PrimaryMember.ValueString()),
		Proto:           api.SelectedMap(d.Proto.ValueString()),
		LacpFastTimeout: tools.BoolToString(d.LacpFastTimeout.ValueBool()),
		UseFlowID:       api.SelectedMap(d.UseFlowid.ValueString()),
		Hash:            api.SelectedMapList(setToSlice(d.Hash)),
		LacpStrict:      api.SelectedMap(d.LacpStrict.ValueString()),
		MTU:             tools.Int64ToStringNegative(d.MTU.ValueInt64()),
		Description:     d.Description.ValueString(),
	}, nil
}

func convertLaggStructToSchema(d *interfaces.Lagg) (*laggResourceModel, error) {
	return &laggResourceModel{
		Device:          tools.StringOrNull(d.Device),
		Members:         tools.StringSliceToSet(d.Members),
		PrimaryMember:   tools.StringOrNull(d.PrimaryMember.String()),
		Proto:           types.StringValue(d.Proto.String()),
		LacpFastTimeout: types.BoolValue(tools.StringToBool(d.LacpFastTimeout)),
		UseFlowid:       types.StringValue(d.UseFlowID.String()),
		Hash:            tools.StringSliceToSet(d.Hash),
		LacpStrict:      types.StringValue(d.LacpStrict.String()),
		MTU:             types.Int64Value(tools.StringToInt64(d.MTU)),
		Description:     tools.StringOrNull(d.Description),
	}, nil
}
