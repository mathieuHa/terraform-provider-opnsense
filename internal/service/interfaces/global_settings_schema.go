package interfaces

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/interfaces"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type globalSettingsResourceModel struct {
	DisableChecksumOffloading     types.Bool   `tfsdk:"disable_checksum_offloading"`
	DisableSegmentationOffloading types.Bool   `tfsdk:"disable_segmentation_offloading"`
	DisableLargeReceiveOffloading types.Bool   `tfsdk:"disable_large_receive_offloading"`
	VlanHwFilter                  types.String `tfsdk:"vlan_hw_filter"`
	DisableIpv6                   types.Bool   `tfsdk:"disable_ipv6"`
	Dhcp6NoRelease                types.Bool   `tfsdk:"dhcp6_no_release"`
	Dhcp6Debug                    types.Bool   `tfsdk:"dhcp6_debug"`
	Dhcp6Duid                     types.String `tfsdk:"dhcp6_duid"`
	Dhcp6RaTimeout                types.Int64  `tfsdk:"dhcp6_ra_timeout"`

	Id types.String `tfsdk:"id"`
}

func globalSettingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"disable_checksum_offloading": schema.BoolAttribute{
			MarkdownDescription: "Disable hardware checksum offloading.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"disable_segmentation_offloading": schema.BoolAttribute{
			MarkdownDescription: "Disable hardware TCP segmentation offloading (TSO).",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"disable_large_receive_offloading": schema.BoolAttribute{
			MarkdownDescription: "Disable hardware large receive offloading (LRO).",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"vlan_hw_filter": schema.StringAttribute{
			MarkdownDescription: "VLAN hardware filtering: `0` enable, `1` disable, `2` leave the driver default.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"disable_ipv6": schema.BoolAttribute{
			MarkdownDescription: "Block all IPv6 traffic.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"dhcp6_no_release": schema.BoolAttribute{
			MarkdownDescription: "Do not send a DHCPv6 release when the WAN goes down.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"dhcp6_debug": schema.BoolAttribute{
			MarkdownDescription: "Debug logging of the DHCPv6 client.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"dhcp6_duid": schema.StringAttribute{
			MarkdownDescription: "DHCPv6 DUID, empty to let the system generate one.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"dhcp6_ra_timeout": schema.Int64Attribute{
			MarkdownDescription: "Seconds to wait for a router advertisement before starting the DHCPv6 client. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `interfaces_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertGlobalSettingsSchemaToStruct(d *globalSettingsResourceModel) *interfaces.GlobalSettings {
	return &interfaces.GlobalSettings{
		DisableChecksumOffloading:     settingsBool(d.DisableChecksumOffloading),
		DisableSegmentationOffloading: settingsBool(d.DisableSegmentationOffloading),
		DisableLargeReceiveOffloading: settingsBool(d.DisableLargeReceiveOffloading),
		VlanHwFilter:                  api.SelectedMap(d.VlanHwFilter.ValueString()),
		DisableIpv6:                   settingsBool(d.DisableIpv6),
		Dhcp6NoRelease:                settingsBool(d.Dhcp6NoRelease),
		Dhcp6Debug:                    settingsBool(d.Dhcp6Debug),
		Dhcp6Duid:                     d.Dhcp6Duid.ValueString(),
		Dhcp6RaTimeout:                settingsInt(d.Dhcp6RaTimeout),
	}
}

func convertGlobalSettingsStructToSchema(s *interfaces.GlobalSettings) *globalSettingsResourceModel {
	return &globalSettingsResourceModel{
		DisableChecksumOffloading:     types.BoolValue(tools.StringToBool(s.DisableChecksumOffloading)),
		DisableSegmentationOffloading: types.BoolValue(tools.StringToBool(s.DisableSegmentationOffloading)),
		DisableLargeReceiveOffloading: types.BoolValue(tools.StringToBool(s.DisableLargeReceiveOffloading)),
		VlanHwFilter:                  types.StringValue(s.VlanHwFilter.String()),
		DisableIpv6:                   types.BoolValue(tools.StringToBool(s.DisableIpv6)),
		Dhcp6NoRelease:                types.BoolValue(tools.StringToBool(s.Dhcp6NoRelease)),
		Dhcp6Debug:                    types.BoolValue(tools.StringToBool(s.Dhcp6Debug)),
		Dhcp6Duid:                     types.StringValue(s.Dhcp6Duid),
		Dhcp6RaTimeout:                settingsReadInt(s.Dhcp6RaTimeout),
	}
}
