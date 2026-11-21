package tailscale

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/tailscale"
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

type settingsResourceModel struct {
	Enabled            types.Bool   `tfsdk:"enabled"`
	LoginTimeout       types.Int64  `tfsdk:"login_timeout"`
	ListenPort         types.Int64  `tfsdk:"listen_port"`
	AcceptDns          types.Bool   `tfsdk:"accept_dns"`
	AdvertiseExitNode  types.Bool   `tfsdk:"advertise_exit_node"`
	ExitNode           types.String `tfsdk:"exit_node"`
	AcceptSubnetRoutes types.Bool   `tfsdk:"accept_subnet_routes"`
	EnableSsh          types.Bool   `tfsdk:"enable_ssh"`
	DisableSnat        types.Bool   `tfsdk:"disable_snat"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable Tailscale. The node must be authenticated, see `opnsense_tailscale_authentication`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"login_timeout": schema.Int64Attribute{
			MarkdownDescription: "Seconds to wait for the login when the service starts. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"listen_port": schema.Int64Attribute{
			MarkdownDescription: "UDP port of the WireGuard tunnel. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"accept_dns": schema.BoolAttribute{
			MarkdownDescription: "Use the DNS settings of the tailnet.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"advertise_exit_node": schema.BoolAttribute{
			MarkdownDescription: "Offer this firewall as an exit node.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"exit_node": schema.StringAttribute{
			MarkdownDescription: "Exit node to route internet traffic through, empty for none.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"accept_subnet_routes": schema.BoolAttribute{
			MarkdownDescription: "Accept the subnet routes advertised by other nodes.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"enable_ssh": schema.BoolAttribute{
			MarkdownDescription: "Enable Tailscale SSH on this node.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"disable_snat": schema.BoolAttribute{
			MarkdownDescription: "Do not masquerade traffic to the advertised subnets.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `tailscale_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func settingsBool(b types.Bool) string {
	if b.IsNull() || b.IsUnknown() {
		return ""
	}
	return tools.BoolToString(b.ValueBool())
}

func settingsInt(i types.Int64) string {
	if i.IsNull() || i.IsUnknown() || i.ValueInt64() == -1 {
		return ""
	}
	return tools.Int64ToString(i.ValueInt64())
}

func settingsReadInt(s string) types.Int64 {
	if s == "" {
		return types.Int64Value(-1)
	}
	return tools.StringToInt64Null(s)
}

func settingsList(s types.Set) []string {
	out := tools.SetToStringSlice(s)
	if out == nil {
		return []string{}
	}
	return out
}

func settingsNonEmpty(in []string) []string {
	out := []string{}
	for _, v := range in {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

func convertSettingsSchemaToStruct(d *settingsResourceModel) *tailscale.Settings {
	return &tailscale.Settings{
		Enabled:            settingsBool(d.Enabled),
		LoginTimeout:       settingsInt(d.LoginTimeout),
		ListenPort:         settingsInt(d.ListenPort),
		AcceptDns:          settingsBool(d.AcceptDns),
		AdvertiseExitNode:  settingsBool(d.AdvertiseExitNode),
		ExitNode:           api.SelectedMap(d.ExitNode.ValueString()),
		AcceptSubnetRoutes: settingsBool(d.AcceptSubnetRoutes),
		EnableSsh:          settingsBool(d.EnableSsh),
		DisableSnat:        settingsBool(d.DisableSnat),
	}
}

func convertSettingsStructToSchema(s *tailscale.Settings) *settingsResourceModel {
	return &settingsResourceModel{
		Enabled:            types.BoolValue(tools.StringToBool(s.Enabled)),
		LoginTimeout:       settingsReadInt(s.LoginTimeout),
		ListenPort:         settingsReadInt(s.ListenPort),
		AcceptDns:          types.BoolValue(tools.StringToBool(s.AcceptDns)),
		AdvertiseExitNode:  types.BoolValue(tools.StringToBool(s.AdvertiseExitNode)),
		ExitNode:           types.StringValue(s.ExitNode.String()),
		AcceptSubnetRoutes: types.BoolValue(tools.StringToBool(s.AcceptSubnetRoutes)),
		EnableSsh:          types.BoolValue(tools.StringToBool(s.EnableSsh)),
		DisableSnat:        types.BoolValue(tools.StringToBool(s.DisableSnat)),
	}
}
