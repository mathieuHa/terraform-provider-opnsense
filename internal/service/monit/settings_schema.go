package monit

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/monit"
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
	Enabled                   types.Bool   `tfsdk:"enabled"`
	Interval                  types.Int64  `tfsdk:"interval"`
	Startdelay                types.Int64  `tfsdk:"startdelay"`
	Mailserver                types.Set    `tfsdk:"mailserver"`
	Port                      types.Int64  `tfsdk:"port"`
	Username                  types.String `tfsdk:"username"`
	Password                  types.String `tfsdk:"password"`
	Ssl                       types.Bool   `tfsdk:"ssl"`
	Sslversion                types.String `tfsdk:"sslversion"`
	Sslverify                 types.Bool   `tfsdk:"sslverify"`
	Logfile                   types.String `tfsdk:"logfile"`
	Statefile                 types.String `tfsdk:"statefile"`
	EventqueuePath            types.String `tfsdk:"eventqueue_path"`
	EventqueueSlots           types.Int64  `tfsdk:"eventqueue_slots"`
	HttpdEnabled              types.Bool   `tfsdk:"httpd_enabled"`
	HttpdUsername             types.String `tfsdk:"httpd_username"`
	HttpdPassword             types.String `tfsdk:"httpd_password"`
	HttpdPort                 types.Int64  `tfsdk:"httpd_port"`
	HttpdAllow                types.Set    `tfsdk:"httpd_allow"`
	MmonitUrl                 types.String `tfsdk:"mmonit_url"`
	MmonitTimeout             types.Int64  `tfsdk:"mmonit_timeout"`
	MmonitRegisterCredentials types.Bool   `tfsdk:"mmonit_register_credentials"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable or disable monit.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"interval": schema.Int64Attribute{
			MarkdownDescription: "Polling interval in seconds. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"startdelay": schema.Int64Attribute{
			MarkdownDescription: "On system boot wait before Monit starts checking services. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"mailserver": schema.SetAttribute{
			MarkdownDescription: "Comma-separated list of SMTP servers for alert delivery.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"port": schema.Int64Attribute{
			MarkdownDescription: "The port of the mail server. Typically 465 for SSL or 25 for TLS and non secure connections. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"username": schema.StringAttribute{
			MarkdownDescription: "The username for mail server authentication.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"password": schema.StringAttribute{
			MarkdownDescription: "The password for mail server authentication.",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"ssl": schema.BoolAttribute{
			MarkdownDescription: "Enable encryption for mail server communication.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"sslversion": schema.StringAttribute{
			MarkdownDescription: "Set the SSL version to use. In AUTO mode, only TLS is used. One of `auto`, `tlsv1`, `tlsv11`, `tlsv12`, `tlsv13`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"sslverify": schema.BoolAttribute{
			MarkdownDescription: "Enable or disable SSL server certificate verification.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"logfile": schema.StringAttribute{
			MarkdownDescription: "The log file of the Monit process. This can be the keyword syslog or the path to a file.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"statefile": schema.StringAttribute{
			MarkdownDescription: "The state file of the Monit process.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"eventqueue_path": schema.StringAttribute{
			MarkdownDescription: "The path to the eventqueue directory.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"eventqueue_slots": schema.Int64Attribute{
			MarkdownDescription: "The number of eventqueue slots. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"httpd_enabled": schema.BoolAttribute{
			MarkdownDescription: "Start the Monit httpd service.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"httpd_username": schema.StringAttribute{
			MarkdownDescription: "Monit HTTPD Username.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"httpd_password": schema.StringAttribute{
			MarkdownDescription: "Monit HTTPD Password.",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"httpd_port": schema.Int64Attribute{
			MarkdownDescription: "Monit HTTPD Port. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"httpd_allow": schema.SetAttribute{
			MarkdownDescription: "The username:password or host/network etc. for accessing the Monit httpd service.See Monit Authentication.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"mmonit_url": schema.StringAttribute{
			MarkdownDescription: "The M/Monit URL.https://user:pass@192.168.1.10:8443/collectorIf you want to control Monit services from your M/Monit instance you have to configure the Monit Port too and add corresponding firewall rules as well.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"mmonit_timeout": schema.Int64Attribute{
			MarkdownDescription: "M/Monit Timeout. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"mmonit_register_credentials": schema.BoolAttribute{
			MarkdownDescription: "Automatically register in M/Monit by sending Monit credentials (see Monit Access List above).",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `monit_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertSettingsSchemaToStruct(d *settingsResourceModel) *monit.GeneralConfig {
	return &monit.GeneralConfig{
		General: monit.GeneralConfigGeneral{
			Enabled:                   settingsBool(d.Enabled),
			Interval:                  settingsInt(d.Interval),
			Startdelay:                settingsInt(d.Startdelay),
			Mailserver:                api.SelectedMapList(settingsList(d.Mailserver)),
			Port:                      settingsInt(d.Port),
			Username:                  d.Username.ValueString(),
			Password:                  d.Password.ValueString(),
			Ssl:                       settingsBool(d.Ssl),
			Sslversion:                api.SelectedMap(d.Sslversion.ValueString()),
			Sslverify:                 settingsBool(d.Sslverify),
			Logfile:                   d.Logfile.ValueString(),
			Statefile:                 d.Statefile.ValueString(),
			EventqueuePath:            d.EventqueuePath.ValueString(),
			EventqueueSlots:           settingsInt(d.EventqueueSlots),
			HttpdEnabled:              settingsBool(d.HttpdEnabled),
			HttpdUsername:             d.HttpdUsername.ValueString(),
			HttpdPassword:             d.HttpdPassword.ValueString(),
			HttpdPort:                 settingsInt(d.HttpdPort),
			HttpdAllow:                api.SelectedMapList(settingsList(d.HttpdAllow)),
			MmonitUrl:                 d.MmonitUrl.ValueString(),
			MmonitTimeout:             settingsInt(d.MmonitTimeout),
			MmonitRegisterCredentials: settingsBool(d.MmonitRegisterCredentials),
		},
	}
}

func convertSettingsStructToSchema(s *monit.GeneralConfig) *settingsResourceModel {
	return &settingsResourceModel{
		Enabled:                   types.BoolValue(tools.StringToBool(s.General.Enabled)),
		Interval:                  settingsReadInt(s.General.Interval),
		Startdelay:                settingsReadInt(s.General.Startdelay),
		Mailserver:                tools.StringSliceToSet(settingsNonEmpty(s.General.Mailserver)),
		Port:                      settingsReadInt(s.General.Port),
		Username:                  types.StringValue(s.General.Username),
		Password:                  types.StringValue(s.General.Password),
		Ssl:                       types.BoolValue(tools.StringToBool(s.General.Ssl)),
		Sslversion:                types.StringValue(s.General.Sslversion.String()),
		Sslverify:                 types.BoolValue(tools.StringToBool(s.General.Sslverify)),
		Logfile:                   types.StringValue(s.General.Logfile),
		Statefile:                 types.StringValue(s.General.Statefile),
		EventqueuePath:            types.StringValue(s.General.EventqueuePath),
		EventqueueSlots:           settingsReadInt(s.General.EventqueueSlots),
		HttpdEnabled:              types.BoolValue(tools.StringToBool(s.General.HttpdEnabled)),
		HttpdUsername:             types.StringValue(s.General.HttpdUsername),
		HttpdPassword:             types.StringValue(s.General.HttpdPassword),
		HttpdPort:                 settingsReadInt(s.General.HttpdPort),
		HttpdAllow:                tools.StringSliceToSet(settingsNonEmpty(s.General.HttpdAllow)),
		MmonitUrl:                 types.StringValue(s.General.MmonitUrl),
		MmonitTimeout:             settingsReadInt(s.General.MmonitTimeout),
		MmonitRegisterCredentials: types.BoolValue(tools.StringToBool(s.General.MmonitRegisterCredentials)),
	}
}
