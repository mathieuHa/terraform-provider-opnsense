package bind

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/bind"
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
	Enabled            types.Bool   `tfsdk:"enabled"`
	Disablev6          types.Bool   `tfsdk:"disablev6"`
	Enablerpz          types.Bool   `tfsdk:"enablerpz"`
	Listenv4           types.Set    `tfsdk:"listenv4"`
	Listenv6           types.Set    `tfsdk:"listenv6"`
	Querysource        types.String `tfsdk:"querysource"`
	Querysourcev6      types.String `tfsdk:"querysourcev6"`
	Transfersource     types.String `tfsdk:"transfersource"`
	Transfersourcev6   types.String `tfsdk:"transfersourcev6"`
	Port               types.Int64  `tfsdk:"port"`
	Forwarders         types.Set    `tfsdk:"forwarders"`
	Filteraaaav4       types.Bool   `tfsdk:"filteraaaav4"`
	Filteraaaav6       types.Bool   `tfsdk:"filteraaaav6"`
	Filteraaaaacl      types.Set    `tfsdk:"filteraaaaacl"`
	Logsize            types.Int64  `tfsdk:"logsize"`
	GeneralLogLevel    types.String `tfsdk:"general_log_level"`
	Maxcachesize       types.Int64  `tfsdk:"maxcachesize"`
	Recursion          types.Set    `tfsdk:"recursion"`
	Allowtransfer      types.Set    `tfsdk:"allowtransfer"`
	Allowquery         types.Set    `tfsdk:"allowquery"`
	Dnssecvalidation   types.String `tfsdk:"dnssecvalidation"`
	Hidehostname       types.Bool   `tfsdk:"hidehostname"`
	Hideversion        types.Bool   `tfsdk:"hideversion"`
	Disableprefetch    types.Bool   `tfsdk:"disableprefetch"`
	Enableratelimiting types.Bool   `tfsdk:"enableratelimiting"`
	Ratelimitcount     types.Int64  `tfsdk:"ratelimitcount"`
	Ratelimitexcept    types.Set    `tfsdk:"ratelimitexcept"`
	Rndcalgo           types.String `tfsdk:"rndcalgo"`
	Rndcsecret         types.String `tfsdk:"rndcsecret"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "This will activate the BIND daemon.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"disablev6": schema.BoolAttribute{
			MarkdownDescription: "This will run BIND in IPv4-only mode.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"enablerpz": schema.BoolAttribute{
			MarkdownDescription: "enablerpz.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"listenv4": schema.SetAttribute{
			MarkdownDescription: "Set the IPv4 addresses the service should listen to.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"listenv6": schema.SetAttribute{
			MarkdownDescription: "Set the IPv6 addresses the service should listen to.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"querysource": schema.StringAttribute{
			MarkdownDescription: "Specify the IPv4 address used as a source for outbound queries.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"querysourcev6": schema.StringAttribute{
			MarkdownDescription: "Specify the IPv6 address used as a source for outbound queries.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"transfersource": schema.StringAttribute{
			MarkdownDescription: "Specify the IPv4 address used as a source for zone transfers.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"transfersourcev6": schema.StringAttribute{
			MarkdownDescription: "Specify the IPv6 address used as a source for zone transfers.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"port": schema.Int64Attribute{
			MarkdownDescription: "Set the port the service should listen to. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"forwarders": schema.SetAttribute{
			MarkdownDescription: "Set one or more hosts to send your DNS queries if the request is unknown.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"filteraaaav4": schema.BoolAttribute{
			MarkdownDescription: "This will filter AAAA records on IPv4 Clients. Set \"DNSSEC Validation\" to \"No\" and AAAA records will be omitted even if they are signed.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"filteraaaav6": schema.BoolAttribute{
			MarkdownDescription: "This will filter AAAA records on IPv6 Clients. Set \"DNSSEC Validation\" to \"No\" and AAAA records will be omitted even if they are signed.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"filteraaaaacl": schema.SetAttribute{
			MarkdownDescription: "Specifies a list of client addresses for which AAAA filtering is to be applied.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"logsize": schema.Int64Attribute{
			MarkdownDescription: "Set the amount how big a logfile can growth. For Query and Blocked logs. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"general_log_level": schema.StringAttribute{
			MarkdownDescription: "Select General Log level. Log levels are listed in the order of increasing verbosity. Setting a certain log level will cause all messages of the specified and more severe log levels to be logged. One of `critical`, `error`, `warning`, `notice`, `info`, `debug`, `dynamic`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"maxcachesize": schema.Int64Attribute{
			MarkdownDescription: "How much memory in percent the cache can use from the system. Default is 80%. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"recursion": schema.SetAttribute{
			MarkdownDescription: "Define an ACL where you allow which clients can resolve via this service. Usually use your local LAN.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"allowtransfer": schema.SetAttribute{
			MarkdownDescription: "Define the ACLs where you allow which server can retrieve zones.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"allowquery": schema.SetAttribute{
			MarkdownDescription: "Define the ACLs where you allow which client are allowed to query this server.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"dnssecvalidation": schema.StringAttribute{
			MarkdownDescription: "Default is \"No\". Set to \"Auto\" to use the static trust anchor configuration by the system. One of `no`, `auto`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"hidehostname": schema.BoolAttribute{
			MarkdownDescription: "This will hide the system hostname for DNS queries.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"hideversion": schema.BoolAttribute{
			MarkdownDescription: "This will hide the local BIND version in DNS queries.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"disableprefetch": schema.BoolAttribute{
			MarkdownDescription: "This will disable prefetching of domains before they time out.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"enableratelimiting": schema.BoolAttribute{
			MarkdownDescription: "This will enable rate-limiting for DNS replies.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"ratelimitcount": schema.Int64Attribute{
			MarkdownDescription: "Set how many replies per second are allowed. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"ratelimitexcept": schema.SetAttribute{
			MarkdownDescription: "Except a list of IPs from rate-limiting like ::1.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"rndcalgo": schema.StringAttribute{
			MarkdownDescription: "Set the authentication algorithm for the RNDC key. This requires a restart of the Bind Service. One of `hmac-sha512`, `hmac-sha384`, `hmac-sha256`, `hmac-sha224`, `hmac-sha1`, `hmac-md5`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"rndcsecret": schema.StringAttribute{
			MarkdownDescription: "The base64-encoded RNDC key. This requires a restart of the Bind Service.",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `bind_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertSettingsSchemaToStruct(d *settingsResourceModel) *bind.GeneralSettings {
	return &bind.GeneralSettings{
		Enabled:            settingsBool(d.Enabled),
		Disablev6:          settingsBool(d.Disablev6),
		Enablerpz:          settingsBool(d.Enablerpz),
		Listenv4:           api.SelectedMapList(settingsList(d.Listenv4)),
		Listenv6:           api.SelectedMapList(settingsList(d.Listenv6)),
		Querysource:        d.Querysource.ValueString(),
		Querysourcev6:      d.Querysourcev6.ValueString(),
		Transfersource:     d.Transfersource.ValueString(),
		Transfersourcev6:   d.Transfersourcev6.ValueString(),
		Port:               settingsInt(d.Port),
		Forwarders:         api.SelectedMapList(settingsList(d.Forwarders)),
		Filteraaaav4:       settingsBool(d.Filteraaaav4),
		Filteraaaav6:       settingsBool(d.Filteraaaav6),
		Filteraaaaacl:      api.SelectedMapList(settingsList(d.Filteraaaaacl)),
		Logsize:            settingsInt(d.Logsize),
		GeneralLogLevel:    api.SelectedMap(d.GeneralLogLevel.ValueString()),
		Maxcachesize:       settingsInt(d.Maxcachesize),
		Recursion:          api.SelectedMapList(settingsList(d.Recursion)),
		Allowtransfer:      api.SelectedMapList(settingsList(d.Allowtransfer)),
		Allowquery:         api.SelectedMapList(settingsList(d.Allowquery)),
		Dnssecvalidation:   api.SelectedMap(d.Dnssecvalidation.ValueString()),
		Hidehostname:       settingsBool(d.Hidehostname),
		Hideversion:        settingsBool(d.Hideversion),
		Disableprefetch:    settingsBool(d.Disableprefetch),
		Enableratelimiting: settingsBool(d.Enableratelimiting),
		Ratelimitcount:     settingsInt(d.Ratelimitcount),
		Ratelimitexcept:    api.SelectedMapList(settingsList(d.Ratelimitexcept)),
		Rndcalgo:           api.SelectedMap(d.Rndcalgo.ValueString()),
		Rndcsecret:         d.Rndcsecret.ValueString(),
	}
}

func convertSettingsStructToSchema(s *bind.GeneralSettings) *settingsResourceModel {
	return &settingsResourceModel{
		Enabled:            types.BoolValue(tools.StringToBool(s.Enabled)),
		Disablev6:          types.BoolValue(tools.StringToBool(s.Disablev6)),
		Enablerpz:          types.BoolValue(tools.StringToBool(s.Enablerpz)),
		Listenv4:           tools.StringSliceToSet(settingsNonEmpty(s.Listenv4)),
		Listenv6:           tools.StringSliceToSet(settingsNonEmpty(s.Listenv6)),
		Querysource:        types.StringValue(s.Querysource),
		Querysourcev6:      types.StringValue(s.Querysourcev6),
		Transfersource:     types.StringValue(s.Transfersource),
		Transfersourcev6:   types.StringValue(s.Transfersourcev6),
		Port:               settingsReadInt(s.Port),
		Forwarders:         tools.StringSliceToSet(settingsNonEmpty(s.Forwarders)),
		Filteraaaav4:       types.BoolValue(tools.StringToBool(s.Filteraaaav4)),
		Filteraaaav6:       types.BoolValue(tools.StringToBool(s.Filteraaaav6)),
		Filteraaaaacl:      tools.StringSliceToSet(settingsNonEmpty(s.Filteraaaaacl)),
		Logsize:            settingsReadInt(s.Logsize),
		GeneralLogLevel:    types.StringValue(s.GeneralLogLevel.String()),
		Maxcachesize:       settingsReadInt(s.Maxcachesize),
		Recursion:          tools.StringSliceToSet(settingsNonEmpty(s.Recursion)),
		Allowtransfer:      tools.StringSliceToSet(settingsNonEmpty(s.Allowtransfer)),
		Allowquery:         tools.StringSliceToSet(settingsNonEmpty(s.Allowquery)),
		Dnssecvalidation:   types.StringValue(s.Dnssecvalidation.String()),
		Hidehostname:       types.BoolValue(tools.StringToBool(s.Hidehostname)),
		Hideversion:        types.BoolValue(tools.StringToBool(s.Hideversion)),
		Disableprefetch:    types.BoolValue(tools.StringToBool(s.Disableprefetch)),
		Enableratelimiting: types.BoolValue(tools.StringToBool(s.Enableratelimiting)),
		Ratelimitcount:     settingsReadInt(s.Ratelimitcount),
		Ratelimitexcept:    tools.StringSliceToSet(settingsNonEmpty(s.Ratelimitexcept)),
		Rndcalgo:           types.StringValue(s.Rndcalgo.String()),
		Rndcsecret:         types.StringValue(s.Rndcsecret),
	}
}
