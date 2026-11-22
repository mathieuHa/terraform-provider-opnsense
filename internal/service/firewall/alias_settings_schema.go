package firewall

import (
	"github.com/browningluke/opnsense-go/pkg/firewall"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type aliasSettingsResourceModel struct {
	GeoipUrl types.String `tfsdk:"geoip_url"`

	Id types.String `tfsdk:"id"`
}

func aliasSettingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"geoip_url": schema.StringAttribute{
			MarkdownDescription: "Location to fetch geoip address ranges from. Refer to the OPNsense documentation for further details.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `firewall_alias_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertAliasSettingsSchemaToStruct(d *aliasSettingsResourceModel) *firewall.AliasConfig {
	return &firewall.AliasConfig{
		Geoip: firewall.AliasConfigGeoip{
			GeoipUrl: d.GeoipUrl.ValueString(),
		},
	}
}

func convertAliasSettingsStructToSchema(s *firewall.AliasConfig) *aliasSettingsResourceModel {
	return &aliasSettingsResourceModel{
		GeoipUrl: types.StringValue(s.Geoip.GeoipUrl),
	}
}
