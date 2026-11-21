package hasync

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/hasync"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type settingsResourceModel struct {
	DisablePreempt  types.Bool   `tfsdk:"disable_preempt"`
	DisconnectPpps  types.Bool   `tfsdk:"disconnect_ppps"`
	PfsyncInterface types.String `tfsdk:"pfsync_interface"`
	PfsyncPeerIp    types.String `tfsdk:"pfsync_peer_ip"`
	PfsyncVersion   types.String `tfsdk:"pfsync_version"`
	PfsyncDefer     types.Bool   `tfsdk:"pfsync_defer"`
	SynchronizeToIp types.String `tfsdk:"synchronize_to_ip"`
	VerifyPeer      types.Bool   `tfsdk:"verify_peer"`
	Username        types.String `tfsdk:"username"`
	Password        types.String `tfsdk:"password"`
	SyncItems       types.Set    `tfsdk:"sync_items"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"disable_preempt": schema.BoolAttribute{
			MarkdownDescription: "Do not let a CARP member preempt the master role automatically.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"disconnect_ppps": schema.BoolAttribute{
			MarkdownDescription: "Disconnect PPP links when the firewall becomes CARP backup.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"pfsync_interface": schema.StringAttribute{
			MarkdownDescription: "Interface used to synchronise firewall states (pfsync), e.g. `opt1`. Empty disables state synchronisation.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"pfsync_peer_ip": schema.StringAttribute{
			MarkdownDescription: "Unicast pfsync peer address. Empty uses multicast.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"pfsync_version": schema.StringAttribute{
			MarkdownDescription: "pfsync protocol version, e.g. `1400`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"pfsync_defer": schema.BoolAttribute{
			MarkdownDescription: "Defer packets until the peer acknowledges the state (pfsync defer).",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"synchronize_to_ip": schema.StringAttribute{
			MarkdownDescription: "Address of the backup firewall the configuration is pushed to. Empty disables configuration synchronisation.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"verify_peer": schema.BoolAttribute{
			MarkdownDescription: "Verify the TLS certificate of the backup firewall.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"username": schema.StringAttribute{
			MarkdownDescription: "User on the backup firewall used for synchronisation.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"password": schema.StringAttribute{
			MarkdownDescription: "Password of that user.",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"sync_items": schema.SetAttribute{
			MarkdownDescription: "Configuration sections pushed to the backup, e.g. `[\"aliases\", \"rules\", \"nat\"]`.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `hasync_settings`.",
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
	if i.IsNull() || i.IsUnknown() {
		return ""
	}
	return tools.Int64ToString(i.ValueInt64())
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

func convertSettingsSchemaToStruct(d *settingsResourceModel) *hasync.Settings {
	return &hasync.Settings{
		DisablePreempt:  settingsBool(d.DisablePreempt),
		DisconnectPpps:  settingsBool(d.DisconnectPpps),
		PfsyncInterface: api.SelectedMap(d.PfsyncInterface.ValueString()),
		PfsyncPeerIp:    d.PfsyncPeerIp.ValueString(),
		PfsyncVersion:   api.SelectedMap(d.PfsyncVersion.ValueString()),
		PfsyncDefer:     settingsBool(d.PfsyncDefer),
		SynchronizeToIp: d.SynchronizeToIp.ValueString(),
		VerifyPeer:      settingsBool(d.VerifyPeer),
		Username:        d.Username.ValueString(),
		Password:        d.Password.ValueString(),
		SyncItems:       api.SelectedMapList(settingsList(d.SyncItems)),
	}
}

func convertSettingsStructToSchema(s *hasync.Settings) *settingsResourceModel {
	return &settingsResourceModel{
		DisablePreempt:  types.BoolValue(tools.StringToBool(s.DisablePreempt)),
		DisconnectPpps:  types.BoolValue(tools.StringToBool(s.DisconnectPpps)),
		PfsyncInterface: types.StringValue(s.PfsyncInterface.String()),
		PfsyncPeerIp:    types.StringValue(s.PfsyncPeerIp),
		PfsyncVersion:   types.StringValue(s.PfsyncVersion.String()),
		PfsyncDefer:     types.BoolValue(tools.StringToBool(s.PfsyncDefer)),
		SynchronizeToIp: types.StringValue(s.SynchronizeToIp),
		VerifyPeer:      types.BoolValue(tools.StringToBool(s.VerifyPeer)),
		Username:        types.StringValue(s.Username),
		Password:        types.StringValue(s.Password),
		SyncItems:       tools.StringSliceToSet(settingsNonEmpty(s.SyncItems)),
	}
}
