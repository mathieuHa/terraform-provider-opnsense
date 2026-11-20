package firewall

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/firewall"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type nptResourceModel struct {
	Enabled        types.Bool   `tfsdk:"enabled"`
	Log            types.Bool   `tfsdk:"log"`
	Sequence       types.Int64  `tfsdk:"sequence"`
	Interface      types.String `tfsdk:"interface"`
	SourceNet      types.String `tfsdk:"source_net"`
	DestinationNet types.String `tfsdk:"destination_net"`
	TrackInterface types.String `tfsdk:"track_interface"`
	Categories     types.Set    `tfsdk:"categories"`
	Description    types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func nptResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "NPTv6 rules translate an internal IPv6 prefix to an external one (RFC 6296), one-to-one and without keeping state.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether this rule is enabled. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"log": schema.BoolAttribute{
				MarkdownDescription: "When enabled, packets matching this rule are logged. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"sequence": schema.Int64Attribute{
				MarkdownDescription: "Order in which rules are evaluated, lowest first. Defaults to `1`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1),
				Validators:          []validator.Int64{int64validator.AtLeast(1)},
			},
			"interface": schema.StringAttribute{
				MarkdownDescription: "Interface the translation happens on, usually the WAN, e.g. `wan`.",
				Required:            true,
			},
			"source_net": schema.StringAttribute{
				MarkdownDescription: "Internal IPv6 prefix, e.g. `fd00:1::/64`.",
				Required:            true,
			},
			"destination_net": schema.StringAttribute{
				MarkdownDescription: "External IPv6 prefix of the same length, e.g. `2001:db8:1::/64`. Leave unset when `track_interface` provides it.",
				Optional:            true,
			},
			"track_interface": schema.StringAttribute{
				MarkdownDescription: "Interface whose delegated prefix is used as the external prefix, instead of a fixed `destination_net`.",
				Optional:            true,
			},
			"categories": schema.SetAttribute{
				MarkdownDescription: "UUIDs of the firewall categories this rule belongs to. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the NPT rule.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func nptDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "NPTv6 rules translate an internal IPv6 prefix to an external one (RFC 6296), one-to-one and without keeping state.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the NPT rule.",
				Required:            true,
			},
			"enabled":         dschema.BoolAttribute{MarkdownDescription: "Whether this rule is enabled.", Computed: true},
			"log":             dschema.BoolAttribute{MarkdownDescription: "Whether matching packets are logged.", Computed: true},
			"sequence":        dschema.Int64Attribute{MarkdownDescription: "Order in which rules are evaluated.", Computed: true},
			"interface":       dschema.StringAttribute{MarkdownDescription: "Interface the translation happens on.", Computed: true},
			"source_net":      dschema.StringAttribute{MarkdownDescription: "Internal IPv6 prefix.", Computed: true},
			"destination_net": dschema.StringAttribute{MarkdownDescription: "External IPv6 prefix.", Computed: true},
			"track_interface": dschema.StringAttribute{MarkdownDescription: "Interface whose delegated prefix is used as the external prefix.", Computed: true},
			"categories":      dschema.SetAttribute{MarkdownDescription: "UUIDs of the firewall categories this rule belongs to.", Computed: true, ElementType: types.StringType},
			"description":     dschema.StringAttribute{MarkdownDescription: "Optional description here for your reference (not parsed).", Computed: true},
		},
	}
}

func convertNPTSchemaToStruct(d *nptResourceModel) (*firewall.NPT, error) {
	categories := tools.SetToStringSlice(d.Categories)
	if categories == nil {
		categories = []string{}
	}
	return &firewall.NPT{
		Enabled:        tools.BoolToString(d.Enabled.ValueBool()),
		Log:            tools.BoolToString(d.Log.ValueBool()),
		Sequence:       tools.Int64ToString(d.Sequence.ValueInt64()),
		Interface:      api.SelectedMap(d.Interface.ValueString()),
		SourceNet:      d.SourceNet.ValueString(),
		DestinationNet: d.DestinationNet.ValueString(),
		TrackInterface: api.SelectedMap(d.TrackInterface.ValueString()),
		Categories:     api.SelectedMapList(categories),
		Description:    d.Description.ValueString(),
	}, nil
}

func convertNPTStructToSchema(d *firewall.NPT) (*nptResourceModel, error) {
	return &nptResourceModel{
		Enabled:        types.BoolValue(tools.StringToBool(d.Enabled)),
		Log:            types.BoolValue(tools.StringToBool(d.Log)),
		Sequence:       types.Int64Value(tools.StringToInt64(d.Sequence)),
		Interface:      types.StringValue(d.Interface.String()),
		SourceNet:      types.StringValue(d.SourceNet),
		DestinationNet: tools.StringOrNull(d.DestinationNet),
		TrackInterface: tools.StringOrNull(d.TrackInterface.String()),
		Categories:     tools.StringSliceToSet(d.Categories),
		Description:    tools.StringOrNull(d.Description),
	}, nil
}
