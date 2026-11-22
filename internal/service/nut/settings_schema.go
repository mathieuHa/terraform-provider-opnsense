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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type settingsResourceModel struct {
	GeneralEnable        types.Bool   `tfsdk:"general_enable"`
	GeneralMode          types.String `tfsdk:"general_mode"`
	GeneralName          types.String `tfsdk:"general_name"`
	GeneralListen        types.Set    `tfsdk:"general_listen"`
	AccountAdminPassword types.String `tfsdk:"account_admin_password"`
	AccountMonPassword   types.String `tfsdk:"account_mon_password"`
	UsbhidEnable         types.Bool   `tfsdk:"usbhid_enable"`
	UsbhidArgs           types.Set    `tfsdk:"usbhid_args"`
	ApcsmartEnable       types.Bool   `tfsdk:"apcsmart_enable"`
	ApcsmartArgs         types.Set    `tfsdk:"apcsmart_args"`
	ApcupsdEnable        types.Bool   `tfsdk:"apcupsd_enable"`
	ApcupsdHostname      types.String `tfsdk:"apcupsd_hostname"`
	ApcupsdPort          types.Int64  `tfsdk:"apcupsd_port"`
	BcmxcpusbEnable      types.Bool   `tfsdk:"bcmxcpusb_enable"`
	BcmxcpusbArgs        types.Set    `tfsdk:"bcmxcpusb_args"`
	BlazerusbEnable      types.Bool   `tfsdk:"blazerusb_enable"`
	BlazerusbArgs        types.Set    `tfsdk:"blazerusb_args"`
	BlazerserEnable      types.Bool   `tfsdk:"blazerser_enable"`
	BlazerserArgs        types.Set    `tfsdk:"blazerser_args"`
	NetclientEnable      types.Bool   `tfsdk:"netclient_enable"`
	NetclientAddress     types.String `tfsdk:"netclient_address"`
	NetclientPort        types.Int64  `tfsdk:"netclient_port"`
	NetclientUser        types.String `tfsdk:"netclient_user"`
	NetclientPassword    types.String `tfsdk:"netclient_password"`
	QxEnable             types.Bool   `tfsdk:"qx_enable"`
	QxArgs               types.Set    `tfsdk:"qx_args"`
	RielloEnable         types.Bool   `tfsdk:"riello_enable"`
	RielloArgs           types.Set    `tfsdk:"riello_args"`
	SnmpEnable           types.Bool   `tfsdk:"snmp_enable"`
	SnmpArgs             types.Set    `tfsdk:"snmp_args"`

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
		"general_listen": schema.SetAttribute{
			MarkdownDescription: "Set the addresses this service listen on.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
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
		"usbhid_args": schema.SetAttribute{
			MarkdownDescription: "Set extra arguments for this UPS, e.g. \"port=auto\".",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"apcsmart_enable": schema.BoolAttribute{
			MarkdownDescription: "Enable the APCSMART driver.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"apcsmart_args": schema.SetAttribute{
			MarkdownDescription: "Set extra arguments for this UPS, e.g. \"port=auto\".",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
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
		"bcmxcpusb_args": schema.SetAttribute{
			MarkdownDescription: "Set extra arguments for this UPS, e.g. \"port=auto\".",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"blazerusb_enable": schema.BoolAttribute{
			MarkdownDescription: "Enable the BlazerUSB driver.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"blazerusb_args": schema.SetAttribute{
			MarkdownDescription: "Set extra arguments for this UPS, e.g. \"port=auto\".",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"blazerser_enable": schema.BoolAttribute{
			MarkdownDescription: "Enable the BlazerSerial driver. Please be aware that this driver needs to run nut-tools as root.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"blazerser_args": schema.SetAttribute{
			MarkdownDescription: "Set extra arguments for this UPS, e.g. \"port=auto\".",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
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
		"qx_args": schema.SetAttribute{
			MarkdownDescription: "Set extra arguments for this UPS, e.g. \"port=auto\".",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"riello_enable": schema.BoolAttribute{
			MarkdownDescription: "Enable the Riello driver.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"riello_args": schema.SetAttribute{
			MarkdownDescription: "Set extra arguments for this UPS, e.g. \"port=auto\".",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"snmp_enable": schema.BoolAttribute{
			MarkdownDescription: "Enable the SNMP driver.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"snmp_args": schema.SetAttribute{
			MarkdownDescription: "Set extra arguments for this UPS, e.g. \"community=public\".",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
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
			GeneralListen: api.SelectedMapList(settingsList(d.GeneralListen)),
		},
		Account: nut.SettingsAccount{
			AccountAdminPassword: d.AccountAdminPassword.ValueString(),
			AccountMonPassword:   d.AccountMonPassword.ValueString(),
		},
		Usbhid: nut.SettingsUsbhid{
			UsbhidEnable: settingsBool(d.UsbhidEnable),
			UsbhidArgs:   api.SelectedMapList(settingsList(d.UsbhidArgs)),
		},
		Apcsmart: nut.SettingsApcsmart{
			ApcsmartEnable: settingsBool(d.ApcsmartEnable),
			ApcsmartArgs:   api.SelectedMapList(settingsList(d.ApcsmartArgs)),
		},
		Apcupsd: nut.SettingsApcupsd{
			ApcupsdEnable:   settingsBool(d.ApcupsdEnable),
			ApcupsdHostname: d.ApcupsdHostname.ValueString(),
			ApcupsdPort:     settingsInt(d.ApcupsdPort),
		},
		Bcmxcpusb: nut.SettingsBcmxcpusb{
			BcmxcpusbEnable: settingsBool(d.BcmxcpusbEnable),
			BcmxcpusbArgs:   api.SelectedMapList(settingsList(d.BcmxcpusbArgs)),
		},
		Blazerusb: nut.SettingsBlazerusb{
			BlazerusbEnable: settingsBool(d.BlazerusbEnable),
			BlazerusbArgs:   api.SelectedMapList(settingsList(d.BlazerusbArgs)),
		},
		Blazerser: nut.SettingsBlazerser{
			BlazerserEnable: settingsBool(d.BlazerserEnable),
			BlazerserArgs:   api.SelectedMapList(settingsList(d.BlazerserArgs)),
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
			QxArgs:   api.SelectedMapList(settingsList(d.QxArgs)),
		},
		Riello: nut.SettingsRiello{
			RielloEnable: settingsBool(d.RielloEnable),
			RielloArgs:   api.SelectedMapList(settingsList(d.RielloArgs)),
		},
		Snmp: nut.SettingsSnmp{
			SnmpEnable: settingsBool(d.SnmpEnable),
			SnmpArgs:   api.SelectedMapList(settingsList(d.SnmpArgs)),
		},
	}
}

func convertSettingsStructToSchema(s *nut.Settings) *settingsResourceModel {
	return &settingsResourceModel{
		GeneralEnable:        types.BoolValue(tools.StringToBool(s.General.GeneralEnable)),
		GeneralMode:          types.StringValue(s.General.GeneralMode.String()),
		GeneralName:          types.StringValue(s.General.GeneralName),
		GeneralListen:        tools.StringSliceToSet(settingsNonEmpty(s.General.GeneralListen)),
		AccountAdminPassword: types.StringValue(s.Account.AccountAdminPassword),
		AccountMonPassword:   types.StringValue(s.Account.AccountMonPassword),
		UsbhidEnable:         types.BoolValue(tools.StringToBool(s.Usbhid.UsbhidEnable)),
		UsbhidArgs:           tools.StringSliceToSet(settingsNonEmpty(s.Usbhid.UsbhidArgs)),
		ApcsmartEnable:       types.BoolValue(tools.StringToBool(s.Apcsmart.ApcsmartEnable)),
		ApcsmartArgs:         tools.StringSliceToSet(settingsNonEmpty(s.Apcsmart.ApcsmartArgs)),
		ApcupsdEnable:        types.BoolValue(tools.StringToBool(s.Apcupsd.ApcupsdEnable)),
		ApcupsdHostname:      types.StringValue(s.Apcupsd.ApcupsdHostname),
		ApcupsdPort:          settingsReadInt(s.Apcupsd.ApcupsdPort),
		BcmxcpusbEnable:      types.BoolValue(tools.StringToBool(s.Bcmxcpusb.BcmxcpusbEnable)),
		BcmxcpusbArgs:        tools.StringSliceToSet(settingsNonEmpty(s.Bcmxcpusb.BcmxcpusbArgs)),
		BlazerusbEnable:      types.BoolValue(tools.StringToBool(s.Blazerusb.BlazerusbEnable)),
		BlazerusbArgs:        tools.StringSliceToSet(settingsNonEmpty(s.Blazerusb.BlazerusbArgs)),
		BlazerserEnable:      types.BoolValue(tools.StringToBool(s.Blazerser.BlazerserEnable)),
		BlazerserArgs:        tools.StringSliceToSet(settingsNonEmpty(s.Blazerser.BlazerserArgs)),
		NetclientEnable:      types.BoolValue(tools.StringToBool(s.Netclient.NetclientEnable)),
		NetclientAddress:     types.StringValue(s.Netclient.NetclientAddress),
		NetclientPort:        settingsReadInt(s.Netclient.NetclientPort),
		NetclientUser:        types.StringValue(s.Netclient.NetclientUser),
		NetclientPassword:    types.StringValue(s.Netclient.NetclientPassword),
		QxEnable:             types.BoolValue(tools.StringToBool(s.Qx.QxEnable)),
		QxArgs:               tools.StringSliceToSet(settingsNonEmpty(s.Qx.QxArgs)),
		RielloEnable:         types.BoolValue(tools.StringToBool(s.Riello.RielloEnable)),
		RielloArgs:           tools.StringSliceToSet(settingsNonEmpty(s.Riello.RielloArgs)),
		SnmpEnable:           types.BoolValue(tools.StringToBool(s.Snmp.SnmpEnable)),
		SnmpArgs:             tools.StringSliceToSet(settingsNonEmpty(s.Snmp.SnmpArgs)),
	}
}
