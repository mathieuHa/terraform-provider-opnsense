package kea

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/kea"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type dhcpv4OptionResourceModel struct {
	Code          types.String `tfsdk:"code"`
	Encoding      types.String `tfsdk:"encoding"`
	Data          types.String `tfsdk:"data"`
	Force         types.Bool   `tfsdk:"force"`
	MatchCode     types.String `tfsdk:"match_code"`
	MatchEncoding types.String `tfsdk:"match_encoding"`
	MatchData     types.String `tfsdk:"match_data"`
	Description   types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func dhcpv4OptionResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Custom Kea DHCPv4 options (**Services ▸ Kea DHCP ▸ Kea DHCPv4 ▸ Options**), attached to subnets and reservations through their `options` attribute.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"code": schema.StringAttribute{
				MarkdownDescription: "DHCPv4 option code, e.g. `42`; the GUI lists the known codes with their names.",
				Required:            true,
			},
			"encoding": schema.StringAttribute{
				MarkdownDescription: "Choose the encoding type. \"Hex\" supports all encapsulated and structured options generically via payload in hexadecimal byte pairs. One of `hex`, `ipv4-address`, `ipv6-address`, `uint8`, `uint16`, `uint32`, `int32`, `boolean`, `string`, `fqdn`. Defaults to `hex`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("hex"),
				Validators:          []validator.String{stringvalidator.OneOf("hex", "ipv4-address", "ipv6-address", "uint8", "uint16", "uint32", "int32", "boolean", "string", "fqdn")},
			},
			"data": schema.StringAttribute{
				MarkdownDescription: "Payload to send to a client.",
				Required:            true,
			},
			"force": schema.BoolAttribute{
				MarkdownDescription: "Always send the option, also when the client does not ask for it in the parameter request list. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"match_code": schema.StringAttribute{
				MarkdownDescription: "Only send the option to clients that sent this option code. Empty to send it to every client. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"match_encoding": schema.StringAttribute{
				MarkdownDescription: "Encoding used to evaluate the match condition. \"Hex\" supports all encapsulated and structured options generically via payload in hexadecimal byte pairs. One of `hex`, `ipv4-address`, `ipv6-address`, `uint8`, `uint16`, `uint32`, `int32`, `boolean`, `string`, `fqdn`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "hex", "ipv4-address", "ipv6-address", "uint8", "uint16", "uint32", "int32", "boolean", "string", "fqdn")},
			},
			"match_data": schema.StringAttribute{
				MarkdownDescription: "Data to match against the selected DHCP option. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "You must enter a description here. It is used to reference this option inside reservations and subnets.",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the Kea DHCPv4 option.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func dhcpv4OptionDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Custom Kea DHCPv4 options (**Services ▸ Kea DHCP ▸ Kea DHCPv4 ▸ Options**), attached to subnets and reservations through their `options` attribute.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the Kea DHCPv4 option.",
				Required:            true,
			},
			"code":           dschema.StringAttribute{MarkdownDescription: "DHCPv4 option code, e.g. `42`; the GUI lists the known codes with their names.", Computed: true},
			"encoding":       dschema.StringAttribute{MarkdownDescription: "Choose the encoding type. \"Hex\" supports all encapsulated and structured options generically via payload in hexadecimal byte pairs. One of `hex`, `ipv4-address`, `ipv6-address`, `uint8`, `uint16`, `uint32`, `int32`, `boolean`, `string`, `fqdn`.", Computed: true},
			"data":           dschema.StringAttribute{MarkdownDescription: "Payload to send to a client.", Computed: true},
			"force":          dschema.BoolAttribute{MarkdownDescription: "Always send the option, also when the client does not ask for it in the parameter request list.", Computed: true},
			"match_code":     dschema.StringAttribute{MarkdownDescription: "Only send the option to clients that sent this option code. Empty to send it to every client.", Computed: true},
			"match_encoding": dschema.StringAttribute{MarkdownDescription: "Encoding used to evaluate the match condition. \"Hex\" supports all encapsulated and structured options generically via payload in hexadecimal byte pairs. One of `hex`, `ipv4-address`, `ipv6-address`, `uint8`, `uint16`, `uint32`, `int32`, `boolean`, `string`, `fqdn`.", Computed: true},
			"match_data":     dschema.StringAttribute{MarkdownDescription: "Data to match against the selected DHCP option.", Computed: true},
			"description":    dschema.StringAttribute{MarkdownDescription: "You must enter a description here. It is used to reference this option inside reservations and subnets.", Computed: true},
		},
	}
}

func convertDhcpv4OptionSchemaToStruct(d *dhcpv4OptionResourceModel) (*kea.OptionV4, error) {
	return &kea.OptionV4{
		Code:          api.SelectedMap(d.Code.ValueString()),
		Encoding:      api.SelectedMap(d.Encoding.ValueString()),
		Data:          d.Data.ValueString(),
		Force:         tools.BoolToString(d.Force.ValueBool()),
		MatchCode:     api.SelectedMap(d.MatchCode.ValueString()),
		MatchEncoding: api.SelectedMap(d.MatchEncoding.ValueString()),
		MatchData:     d.MatchData.ValueString(),
		Description:   d.Description.ValueString(),
	}, nil
}

func convertDhcpv4OptionStructToSchema(d *kea.OptionV4) (*dhcpv4OptionResourceModel, error) {
	return &dhcpv4OptionResourceModel{
		Code:          types.StringValue(d.Code.String()),
		Encoding:      types.StringValue(d.Encoding.String()),
		Data:          types.StringValue(d.Data),
		Force:         types.BoolValue(tools.StringToBool(d.Force)),
		MatchCode:     types.StringValue(d.MatchCode.String()),
		MatchEncoding: types.StringValue(d.MatchEncoding.String()),
		MatchData:     types.StringValue(d.MatchData),
		Description:   types.StringValue(d.Description),
	}, nil
}
