package dnsmasq

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dnsmasq"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type optionResourceModel struct {
	Type        types.String `tfsdk:"type"`
	Option      types.String `tfsdk:"option"`
	Option6     types.String `tfsdk:"option6"`
	Interface   types.String `tfsdk:"interface"`
	Tags        types.Set    `tfsdk:"tags"`
	SetTag      types.String `tfsdk:"set_tag"`
	Value       types.String `tfsdk:"value"`
	Force       types.Bool   `tfsdk:"force"`
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func optionResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Dnsmasq DHCP options (**Services ▸ Dnsmasq DNS & DHCP ▸ DHCP options**). A `set` option is sent to clients; a `match` option tags clients that send it.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"type": schema.StringAttribute{
				MarkdownDescription: "`set` sends the option, `match` tags clients sending it. Defaults to `set`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("set"),
				Validators:          []validator.String{stringvalidator.OneOf("set", "match")},
			},
			"option": schema.StringAttribute{
				MarkdownDescription: "DHCPv4 option number, e.g. `6` for DNS servers. Set either this or `option6`.",
				Optional:            true,
			},
			"option6": schema.StringAttribute{
				MarkdownDescription: "DHCPv6 option number, e.g. `23` for DNS servers.",
				Optional:            true,
			},
			"interface": schema.StringAttribute{
				MarkdownDescription: "Only send the option on this interface.",
				Optional:            true,
			},
			"tags": schema.SetAttribute{
				MarkdownDescription: "UUIDs of tags a client must carry to receive the option (type `set`). Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"set_tag": schema.StringAttribute{
				MarkdownDescription: "UUID of the tag set on matching clients (type `match`).",
				Optional:            true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "Option value. IPv6 addresses must be in brackets, e.g. `[2001:db8::53]`.",
				Optional:            true,
			},
			"force": schema.BoolAttribute{
				MarkdownDescription: "When enabled, the option is sent even if the client did not ask for it. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the dnsmasq DHCP option.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func optionDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Dnsmasq DHCP options (**Services ▸ Dnsmasq DNS & DHCP ▸ DHCP options**). A `set` option is sent to clients; a `match` option tags clients that send it.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the dnsmasq DHCP option.",
				Required:            true,
			},
			"type":        dschema.StringAttribute{MarkdownDescription: "`set` sends the option, `match` tags clients sending it.", Computed: true},
			"option":      dschema.StringAttribute{MarkdownDescription: "DHCPv4 option number, e.g. `6` for DNS servers. Set either this or `option6`.", Computed: true},
			"option6":     dschema.StringAttribute{MarkdownDescription: "DHCPv6 option number, e.g. `23` for DNS servers.", Computed: true},
			"interface":   dschema.StringAttribute{MarkdownDescription: "Only send the option on this interface.", Computed: true},
			"tags":        dschema.SetAttribute{MarkdownDescription: "UUIDs of tags a client must carry to receive the option (type `set`).", Computed: true, ElementType: types.StringType},
			"set_tag":     dschema.StringAttribute{MarkdownDescription: "UUID of the tag set on matching clients (type `match`).", Computed: true},
			"value":       dschema.StringAttribute{MarkdownDescription: "Option value. IPv6 addresses must be in brackets, e.g. `[2001:db8::53]`.", Computed: true},
			"force":       dschema.BoolAttribute{MarkdownDescription: "When enabled, the option is sent even if the client did not ask for it.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Optional description here for your reference (not parsed).", Computed: true},
		},
	}
}

func convertOptionSchemaToStruct(d *optionResourceModel) (*dnsmasq.Option, error) {
	return &dnsmasq.Option{
		Type:         api.SelectedMap(d.Type.ValueString()),
		OptionV4:     api.SelectedMap(d.Option.ValueString()),
		OptionV6:     api.SelectedMap(d.Option6.ValueString()),
		Interface:    api.SelectedMap(d.Interface.ValueString()),
		TypeSetTags:  api.SelectedMapList(setToSlice(d.Tags)),
		TypeMatchTag: api.SelectedMap(d.SetTag.ValueString()),
		Value:        d.Value.ValueString(),
		Force:        tools.BoolToString(d.Force.ValueBool()),
		Description:  d.Description.ValueString(),
	}, nil
}

func convertOptionStructToSchema(d *dnsmasq.Option) (*optionResourceModel, error) {
	return &optionResourceModel{
		Type:        types.StringValue(d.Type.String()),
		Option:      tools.StringOrNull(d.OptionV4.String()),
		Option6:     tools.StringOrNull(d.OptionV6.String()),
		Interface:   tools.StringOrNull(d.Interface.String()),
		Tags:        tools.StringSliceToSet(d.TypeSetTags),
		SetTag:      tools.StringOrNull(d.TypeMatchTag.String()),
		Value:       tools.StringOrNull(d.Value),
		Force:       types.BoolValue(tools.StringToBool(d.Force)),
		Description: tools.StringOrNull(d.Description),
	}, nil
}
