package firewall

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/firewall"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type outboundNatSettingsResourceModel struct {
	Mode types.String `tfsdk:"mode"`

	Id types.String `tfsdk:"id"`
}

func outboundNatSettingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"mode": schema.StringAttribute{
			MarkdownDescription: "Outbound NAT mode: `automatic` (rules generated for every interface), `hybrid` (generated rules plus manual ones), `advanced` (manual rules only) or `disabled`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `firewall_outbound_nat_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertOutboundNatSettingsSchemaToStruct(d *outboundNatSettingsResourceModel) *firewall.OutboundNatConfig {
	return &firewall.OutboundNatConfig{
		General: firewall.OutboundNatConfigGeneral{
			Mode: api.SelectedMap(d.Mode.ValueString()),
		},
	}
}

func convertOutboundNatSettingsStructToSchema(s *firewall.OutboundNatConfig) *outboundNatSettingsResourceModel {
	return &outboundNatSettingsResourceModel{
		Mode: types.StringValue(s.General.Mode.String()),
	}
}
