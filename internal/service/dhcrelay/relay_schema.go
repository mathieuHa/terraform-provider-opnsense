package dhcrelay

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dhcrelay"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type relayResourceModel struct {
	Enabled      types.Bool   `tfsdk:"enabled"`
	Interface    types.String `tfsdk:"interface"`
	Destination  types.String `tfsdk:"destination"`
	AgentInfo    types.Bool   `tfsdk:"agent_info"`
	CarpDependOn types.String `tfsdk:"carp_depend_on"`

	Id types.String `tfsdk:"id"`
}

func relayResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "DHCP relays (**Services ▸ DHCP Relay ▸ Relays**), one per interface.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether the relay is enabled. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"interface": schema.StringAttribute{
				MarkdownDescription: "Interface DHCP requests are relayed from.",
				Required:            true,
			},
			"destination": schema.StringAttribute{
				MarkdownDescription: "ID of the `opnsense_dhcrelay_destination` to relay to.",
				Required:            true,
			},
			"agent_info": schema.BoolAttribute{
				MarkdownDescription: "Append the relay agent information option (82). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"carp_depend_on": schema.StringAttribute{
				MarkdownDescription: "CARP VIP the relay follows; it only runs while that VIP is master.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the DHCP relay.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func relayDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "DHCP relays (**Services ▸ DHCP Relay ▸ Relays**), one per interface.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the DHCP relay.",
				Required:            true,
			},
			"enabled":        dschema.BoolAttribute{MarkdownDescription: "Whether the relay is enabled.", Computed: true},
			"interface":      dschema.StringAttribute{MarkdownDescription: "Interface DHCP requests are relayed from.", Computed: true},
			"destination":    dschema.StringAttribute{MarkdownDescription: "ID of the `opnsense_dhcrelay_destination` to relay to.", Computed: true},
			"agent_info":     dschema.BoolAttribute{MarkdownDescription: "Append the relay agent information option (82).", Computed: true},
			"carp_depend_on": dschema.StringAttribute{MarkdownDescription: "CARP VIP the relay follows; it only runs while that VIP is master.", Computed: true},
		},
	}
}

func convertRelaySchemaToStruct(d *relayResourceModel) (*dhcrelay.Relay, error) {
	return &dhcrelay.Relay{
		Enabled:      tools.BoolToString(d.Enabled.ValueBool()),
		Interface:    api.SelectedMap(d.Interface.ValueString()),
		Destination:  api.SelectedMap(d.Destination.ValueString()),
		AgentInfo:    tools.BoolToString(d.AgentInfo.ValueBool()),
		CarpDependOn: api.SelectedMap(d.CarpDependOn.ValueString()),
	}, nil
}

func convertRelayStructToSchema(d *dhcrelay.Relay) (*relayResourceModel, error) {
	return &relayResourceModel{
		Enabled:      types.BoolValue(tools.StringToBool(d.Enabled)),
		Interface:    types.StringValue(d.Interface.String()),
		Destination:  types.StringValue(d.Destination.String()),
		AgentInfo:    types.BoolValue(tools.StringToBool(d.AgentInfo)),
		CarpDependOn: tools.StringOrNull(d.CarpDependOn.String()),
	}, nil
}
