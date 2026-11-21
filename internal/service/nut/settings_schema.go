package nut

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nut"
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
	GeneralEnable        types.Bool   `tfsdk:"general_enable"`
	GeneralMode          types.String `tfsdk:"general_mode"`
	GeneralName          types.String `tfsdk:"general_name"`
	GeneralListen        types.String `tfsdk:"general_listen"`
	AccountAdminPassword types.String `tfsdk:"account_admin_password"`
	AccountMonPassword   types.String `tfsdk:"account_mon_password"`
	UsbhidEnable         types.Bool   `tfsdk:"usbhid_enable"`
	UsbhidArgs           types.String `tfsdk:"usbhid_args"`
	ApcsmartEnable       types.Bool   `tfsdk:"apcsmart_enable"`
	ApcsmartArgs         types.String `tfsdk:"apcsmart_args"`
	ApcupsdEnable        types.Bool   `tfsdk:"apcupsd_enable"`
	ApcupsdHostname      types.String `tfsdk:"apcupsd_hostname"`
	ApcupsdPort          types.Int64  `tfsdk:"apcupsd_port"`
	BcmxcpusbEnable      types.Bool   `tfsdk:"bcmxcpusb_enable"`
	BcmxcpusbArgs        types.String `tfsdk:"bcmxcpusb_args"`
	BlazerusbEnable      types.Bool   `tfsdk:"blazerusb_enable"`
	BlazerusbArgs        types.String `tfsdk:"blazerusb_args"`
	BlazerserEnable      types.Bool   `tfsdk:"blazerser_enable"`
	BlazerserArgs        types.String `tfsdk:"blazerser_args"`
	NetclientEnable      types.Bool   `tfsdk:"netclient_enable"`
	NetclientAddress     types.String `tfsdk:"netclient_address"`
	NetclientPort        types.Int64  `tfsdk:"netclient_port"`
	NetclientUser        types.String `tfsdk:"netclient_user"`
	NetclientPassword    types.String `tfsdk:"netclient_password"`
	QxEnable             types.Bool   `tfsdk:"qx_enable"`
	QxArgs               types.String `tfsdk:"qx_args"`
	RielloEnable         types.Bool   `tfsdk:"riello_enable"`
	RielloArgs           types.String `tfsdk:"riello_args"`
	SnmpEnable           types.Bool   `tfsdk:"snmp_enable"`
	SnmpArgs             types.String `tfsdk:"snmp_args"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"general_enable": schema.BoolAttribute{
			MarkdownDescription: "Enable or disable the nut service. If enabled, the system will shutdown when the UPS emits a low battery warning.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"general_mode": schema.StringAttribute{
			MarkdownDescription: "Set the service mode. Currently only standalone and netclient are available. One of `standalone`, `netclient`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"general_name": schema.StringAttribute{
			MarkdownDescription: "Set a name for your UPS.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"general_listen": schema.StringAttribute{
			MarkdownDescription: "Set the addresses this service listen on.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"account_admin_password": schema.StringAttribute{
			MarkdownDescription: "Set the password for admin user \"admin\".",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"account_mon_password": schema.StringAttribute{
			MarkdownDescription: "Set the password for monitoring user \"monuser\".",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"usbhid_enable": schema.BoolAttribute{
			MarkdownDescription: "Enable the USBHID driver.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"usbhid_args": schema.StringAttribute{
			MarkdownDescription: "Set extra arguments for this UPS, e.g. \"port=auto\".",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"apcsmart_enable": schema.BoolAttribute{
			MarkdownDescription: "Enable the APCSMART driver.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"apcsmart_args": schema.StringAttribute{
			MarkdownDescription: "Set extra arguments for this UPS, e.g. \"port=auto\".",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"apcupsd_enable": schema.BoolAttribute{
			MarkdownDescription: "Enable the APCUPSD controlled devices driver.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"apcupsd_hostname": schema.StringAttribute{
			MarkdownDescription: "Set the hostname or ip of the remote apcupsd server.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"apcupsd_port": schema.Int64Attribute{
			MarkdownDescription: "Set the port of the remote apcupsd server (optional). `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"bcmxcpusb_enable": schema.BoolAttribute{
			MarkdownDescription: "Enable the PowerWare BCMXCPUSB driver.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"bcmxcpusb_args": schema.StringAttribute{
			MarkdownDescription: "Set extra arguments for this UPS, e.g. \"port=auto\".",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"blazerusb_enable": schema.BoolAttribute{
			MarkdownDescription: "Enable the BlazerUSB driver.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"blazerusb_args": schema.StringAttribute{
			MarkdownDescription: "Set extra arguments for this UPS, e.g. \"port=auto\".",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"blazerser_enable": schema.BoolAttribute{
			MarkdownDescription: "Enable the BlazerSerial driver. Please be aware that this driver needs to run nut-tools as root.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"blazerser_args": schema.StringAttribute{
			MarkdownDescription: "Set extra arguments for this UPS, e.g. \"port=auto\".",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"netclient_enable": schema.BoolAttribute{
			MarkdownDescription: "Enable the Netclient driver.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"netclient_address": schema.StringAttribute{
			MarkdownDescription: "Set the IP address of the remote NUT server.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"netclient_port": schema.Int64Attribute{
			MarkdownDescription: "Set the TCP port of the remote NUT server. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"netclient_user": schema.StringAttribute{
			MarkdownDescription: "Set the username of the remote NUT server.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"netclient_password": schema.StringAttribute{
			MarkdownDescription: "Set the password of the remote NUT server.",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"qx_enable": schema.BoolAttribute{
			MarkdownDescription: "Enable the QX driver.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"qx_args": schema.StringAttribute{
			MarkdownDescription: "Set extra arguments for this UPS, e.g. \"port=auto\".",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"riello_enable": schema.BoolAttribute{
			MarkdownDescription: "Enable the Riello driver.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"riello_args": schema.StringAttribute{
			MarkdownDescription: "Set extra arguments for this UPS, e.g. \"port=auto\".",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"snmp_enable": schema.BoolAttribute{
			MarkdownDescription: "Enable the SNMP driver.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"snmp_args": schema.StringAttribute{
			MarkdownDescription: "Set extra arguments for this UPS, e.g. \"community=public\".",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `nut_settings`.",
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

func convertSettingsSchemaToStruct(d *settingsResourceModel) *nut.Settings {
	return &nut.Settings{
		General: nut.SettingsGeneral{
			GeneralEnable: settingsBool(d.GeneralEnable),
			GeneralMode:   api.SelectedMap(d.GeneralMode.ValueString()),
			GeneralName:   d.GeneralName.ValueString(),
			GeneralListen: api.SelectedMap(d.GeneralListen.ValueString()),
		},
		Account: nut.SettingsAccount{
			AccountAdminPassword: d.AccountAdminPassword.ValueString(),
			AccountMonPassword:   d.AccountMonPassword.ValueString(),
		},
		Usbhid: nut.SettingsUsbhid{
			UsbhidEnable: settingsBool(d.UsbhidEnable),
			UsbhidArgs:   api.SelectedMap(d.UsbhidArgs.ValueString()),
		},
		Apcsmart: nut.SettingsApcsmart{
			ApcsmartEnable: settingsBool(d.ApcsmartEnable),
			ApcsmartArgs:   api.SelectedMap(d.ApcsmartArgs.ValueString()),
		},
		Apcupsd: nut.SettingsApcupsd{
			ApcupsdEnable:   settingsBool(d.ApcupsdEnable),
			ApcupsdHostname: d.ApcupsdHostname.ValueString(),
			ApcupsdPort:     settingsInt(d.ApcupsdPort),
		},
		Bcmxcpusb: nut.SettingsBcmxcpusb{
			BcmxcpusbEnable: settingsBool(d.BcmxcpusbEnable),
			BcmxcpusbArgs:   api.SelectedMap(d.BcmxcpusbArgs.ValueString()),
		},
		Blazerusb: nut.SettingsBlazerusb{
			BlazerusbEnable: settingsBool(d.BlazerusbEnable),
			BlazerusbArgs:   api.SelectedMap(d.BlazerusbArgs.ValueString()),
		},
		Blazerser: nut.SettingsBlazerser{
			BlazerserEnable: settingsBool(d.BlazerserEnable),
			BlazerserArgs:   api.SelectedMap(d.BlazerserArgs.ValueString()),
		},
		Netclient: nut.SettingsNetclient{
			NetclientEnable:   settingsBool(d.NetclientEnable),
			NetclientAddress:  d.NetclientAddress.ValueString(),
			NetclientPort:     settingsInt(d.NetclientPort),
			NetclientUser:     d.NetclientUser.ValueString(),
			NetclientPassword: d.NetclientPassword.ValueString(),
		},
		Qx: nut.SettingsQx{
			QxEnable: settingsBool(d.QxEnable),
			QxArgs:   api.SelectedMap(d.QxArgs.ValueString()),
		},
		Riello: nut.SettingsRiello{
			RielloEnable: settingsBool(d.RielloEnable),
			RielloArgs:   api.SelectedMap(d.RielloArgs.ValueString()),
		},
		Snmp: nut.SettingsSnmp{
			SnmpEnable: settingsBool(d.SnmpEnable),
			SnmpArgs:   api.SelectedMap(d.SnmpArgs.ValueString()),
		},
	}
}

func convertSettingsStructToSchema(s *nut.Settings) *settingsResourceModel {
	return &settingsResourceModel{
		GeneralEnable:        types.BoolValue(tools.StringToBool(s.General.GeneralEnable)),
		GeneralMode:          types.StringValue(s.General.GeneralMode.String()),
		GeneralName:          types.StringValue(s.General.GeneralName),
		GeneralListen:        types.StringValue(s.General.GeneralListen.String()),
		AccountAdminPassword: types.StringValue(s.Account.AccountAdminPassword),
		AccountMonPassword:   types.StringValue(s.Account.AccountMonPassword),
		UsbhidEnable:         types.BoolValue(tools.StringToBool(s.Usbhid.UsbhidEnable)),
		UsbhidArgs:           types.StringValue(s.Usbhid.UsbhidArgs.String()),
		ApcsmartEnable:       types.BoolValue(tools.StringToBool(s.Apcsmart.ApcsmartEnable)),
		ApcsmartArgs:         types.StringValue(s.Apcsmart.ApcsmartArgs.String()),
		ApcupsdEnable:        types.BoolValue(tools.StringToBool(s.Apcupsd.ApcupsdEnable)),
		ApcupsdHostname:      types.StringValue(s.Apcupsd.ApcupsdHostname),
		ApcupsdPort:          settingsReadInt(s.Apcupsd.ApcupsdPort),
		BcmxcpusbEnable:      types.BoolValue(tools.StringToBool(s.Bcmxcpusb.BcmxcpusbEnable)),
		BcmxcpusbArgs:        types.StringValue(s.Bcmxcpusb.BcmxcpusbArgs.String()),
		BlazerusbEnable:      types.BoolValue(tools.StringToBool(s.Blazerusb.BlazerusbEnable)),
		BlazerusbArgs:        types.StringValue(s.Blazerusb.BlazerusbArgs.String()),
		BlazerserEnable:      types.BoolValue(tools.StringToBool(s.Blazerser.BlazerserEnable)),
		BlazerserArgs:        types.StringValue(s.Blazerser.BlazerserArgs.String()),
		NetclientEnable:      types.BoolValue(tools.StringToBool(s.Netclient.NetclientEnable)),
		NetclientAddress:     types.StringValue(s.Netclient.NetclientAddress),
		NetclientPort:        settingsReadInt(s.Netclient.NetclientPort),
		NetclientUser:        types.StringValue(s.Netclient.NetclientUser),
		NetclientPassword:    types.StringValue(s.Netclient.NetclientPassword),
		QxEnable:             types.BoolValue(tools.StringToBool(s.Qx.QxEnable)),
		QxArgs:               types.StringValue(s.Qx.QxArgs.String()),
		RielloEnable:         types.BoolValue(tools.StringToBool(s.Riello.RielloEnable)),
		RielloArgs:           types.StringValue(s.Riello.RielloArgs.String()),
		SnmpEnable:           types.BoolValue(tools.StringToBool(s.Snmp.SnmpEnable)),
		SnmpArgs:             types.StringValue(s.Snmp.SnmpArgs.String()),
	}
}
