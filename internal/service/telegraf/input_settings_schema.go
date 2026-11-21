package telegraf

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/telegraf"
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

type inputSettingsResourceModel struct {
	Cpu                      types.Bool   `tfsdk:"cpu"`
	CpuPercpu                types.Bool   `tfsdk:"cpu_percpu"`
	CpuTotalcpu              types.Bool   `tfsdk:"cpu_totalcpu"`
	CollectCpuTime           types.Bool   `tfsdk:"collect_cpu_time"`
	Disk                     types.Bool   `tfsdk:"disk"`
	DiskMountPoints          types.String `tfsdk:"disk_mount_points"`
	DiskIgnoreFs             types.String `tfsdk:"disk_ignore_fs"`
	Diskio                   types.Bool   `tfsdk:"diskio"`
	InternetSpeed            types.Bool   `tfsdk:"internet_speed"`
	InternetSpeedSaving      types.Bool   `tfsdk:"internet_speed_saving"`
	InternetSpeedInterval    types.Int64  `tfsdk:"internet_speed_interval"`
	Mem                      types.Bool   `tfsdk:"mem"`
	Processes                types.Bool   `tfsdk:"processes"`
	Swap                     types.Bool   `tfsdk:"swap"`
	System                   types.Bool   `tfsdk:"system"`
	Network                  types.Bool   `tfsdk:"network"`
	Pf                       types.Bool   `tfsdk:"pf"`
	Ping                     types.Bool   `tfsdk:"ping"`
	PingCount                types.Int64  `tfsdk:"ping_count"`
	PingHosts                types.String `tfsdk:"ping_hosts"`
	Ping6                    types.Bool   `tfsdk:"ping6"`
	Ping6Count               types.Int64  `tfsdk:"ping6_count"`
	Ping6Hosts               types.String `tfsdk:"ping6_hosts"`
	Haproxy                  types.Bool   `tfsdk:"haproxy"`
	Zfs                      types.Bool   `tfsdk:"zfs"`
	Ntpq                     types.Bool   `tfsdk:"ntpq"`
	NtpqDnsLookup            types.Bool   `tfsdk:"ntpq_dns_lookup"`
	IntrusionDetectionAlerts types.Bool   `tfsdk:"intrusion_detection_alerts"`
	Unbound                  types.Bool   `tfsdk:"unbound"`
	Apcupsd                  types.Bool   `tfsdk:"apcupsd"`
	ApcupsdServer            types.String `tfsdk:"apcupsd_server"`

	Id types.String `tfsdk:"id"`
}

func inputSettingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"cpu": schema.BoolAttribute{
			MarkdownDescription: "Read metrics about CPU usage.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"cpu_percpu": schema.BoolAttribute{
			MarkdownDescription: "Whether to report Per-CPU stats or not.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"cpu_totalcpu": schema.BoolAttribute{
			MarkdownDescription: "Whether to report total system CPU stats or not.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"collect_cpu_time": schema.BoolAttribute{
			MarkdownDescription: "If true, collect raw CPU time metrics.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"disk": schema.BoolAttribute{
			MarkdownDescription: "Read metrics about disk usage by mount point.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"disk_mount_points": schema.StringAttribute{
			MarkdownDescription: "By default, Telegraf gather stats for all mountpoints. Setting mountpoints will restrict the stats to the specified mountpoints.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"disk_ignore_fs": schema.StringAttribute{
			MarkdownDescription: "Ignore some mountpoints by filesystem type. For example (dev)tmpfs (usually present on /run, /var/run, /dev/shm or /dev).",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"diskio": schema.BoolAttribute{
			MarkdownDescription: "Read metrics about disk IO by device.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"internet_speed": schema.BoolAttribute{
			MarkdownDescription: "Enable the collection of data about the internet speed on the system.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"internet_speed_saving": schema.BoolAttribute{
			MarkdownDescription: "Not recommended if speed tests work. Reduces memory and computation needs, but accuracy is compromised at connections above 30Mbps.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"internet_speed_interval": schema.Int64Attribute{
			MarkdownDescription: "Default internet speed test interval in seconds. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"mem": schema.BoolAttribute{
			MarkdownDescription: "Read metrics about memory usage.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"processes": schema.BoolAttribute{
			MarkdownDescription: "Get the number of processes and group them by status.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"swap": schema.BoolAttribute{
			MarkdownDescription: "Read metrics about swap memory usage.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"system": schema.BoolAttribute{
			MarkdownDescription: "Read metrics about system load and uptime.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"network": schema.BoolAttribute{
			MarkdownDescription: "Read network interface metrics.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"pf": schema.BoolAttribute{
			MarkdownDescription: "Read PF values via pfctl.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"ping": schema.BoolAttribute{
			MarkdownDescription: "Ping hosts using IPv4 and measure the metrics.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"ping_count": schema.Int64Attribute{
			MarkdownDescription: "Number of ping packets to send per interval. Corresponds to the \"-c\" option of the ping command. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"ping_hosts": schema.StringAttribute{
			MarkdownDescription: "Set the hosts to ping using IPv4 in a CSV list.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"ping6": schema.BoolAttribute{
			MarkdownDescription: "Ping hosts using IPv6 and measure the metrics.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"ping6_count": schema.Int64Attribute{
			MarkdownDescription: "Number of ping packets to send per interval. Corresponds to the \"-c\" option of the ping command. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"ping6_hosts": schema.StringAttribute{
			MarkdownDescription: "Set the hosts to ping using IPv6 in a CSV list.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"haproxy": schema.BoolAttribute{
			MarkdownDescription: "Enable the collection of HAProxy statistics.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"zfs": schema.BoolAttribute{
			MarkdownDescription: "Enable the collection of ZFS statistics.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"ntpq": schema.BoolAttribute{
			MarkdownDescription: "Enable the collection of NTP query metrics.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"ntpq_dns_lookup": schema.BoolAttribute{
			MarkdownDescription: "Can increase metric gather times.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"intrusion_detection_alerts": schema.BoolAttribute{
			MarkdownDescription: "Requires intrusion detection alerts are stored locally.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"unbound": schema.BoolAttribute{
			MarkdownDescription: "Enable the collection of Unbound metrics.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"apcupsd": schema.BoolAttribute{
			MarkdownDescription: "Enable the collection of apcupsd metrics.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"apcupsd_server": schema.StringAttribute{
			MarkdownDescription: "IP address or hostname of the apcupsd net information server. Default address is 127.0.0.1.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `telegraf_input`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertInputSettingsSchemaToStruct(d *inputSettingsResourceModel) *telegraf.InputSettings {
	return &telegraf.InputSettings{
		Cpu:                      settingsBool(d.Cpu),
		CpuPercpu:                settingsBool(d.CpuPercpu),
		CpuTotalcpu:              settingsBool(d.CpuTotalcpu),
		CollectCpuTime:           settingsBool(d.CollectCpuTime),
		Disk:                     settingsBool(d.Disk),
		DiskMountPoints:          d.DiskMountPoints.ValueString(),
		DiskIgnoreFs:             api.SelectedMap(d.DiskIgnoreFs.ValueString()),
		Diskio:                   settingsBool(d.Diskio),
		InternetSpeed:            settingsBool(d.InternetSpeed),
		InternetSpeedSaving:      settingsBool(d.InternetSpeedSaving),
		InternetSpeedInterval:    settingsInt(d.InternetSpeedInterval),
		Mem:                      settingsBool(d.Mem),
		Processes:                settingsBool(d.Processes),
		Swap:                     settingsBool(d.Swap),
		System:                   settingsBool(d.System),
		Network:                  settingsBool(d.Network),
		Pf:                       settingsBool(d.Pf),
		Ping:                     settingsBool(d.Ping),
		PingCount:                settingsInt(d.PingCount),
		PingHosts:                api.SelectedMap(d.PingHosts.ValueString()),
		Ping6:                    settingsBool(d.Ping6),
		Ping6Count:               settingsInt(d.Ping6Count),
		Ping6Hosts:               api.SelectedMap(d.Ping6Hosts.ValueString()),
		Haproxy:                  settingsBool(d.Haproxy),
		Zfs:                      settingsBool(d.Zfs),
		Ntpq:                     settingsBool(d.Ntpq),
		NtpqDnsLookup:            settingsBool(d.NtpqDnsLookup),
		IntrusionDetectionAlerts: settingsBool(d.IntrusionDetectionAlerts),
		Unbound:                  settingsBool(d.Unbound),
		Apcupsd:                  settingsBool(d.Apcupsd),
		ApcupsdServer:            d.ApcupsdServer.ValueString(),
	}
}

func convertInputSettingsStructToSchema(s *telegraf.InputSettings) *inputSettingsResourceModel {
	return &inputSettingsResourceModel{
		Cpu:                      types.BoolValue(tools.StringToBool(s.Cpu)),
		CpuPercpu:                types.BoolValue(tools.StringToBool(s.CpuPercpu)),
		CpuTotalcpu:              types.BoolValue(tools.StringToBool(s.CpuTotalcpu)),
		CollectCpuTime:           types.BoolValue(tools.StringToBool(s.CollectCpuTime)),
		Disk:                     types.BoolValue(tools.StringToBool(s.Disk)),
		DiskMountPoints:          types.StringValue(s.DiskMountPoints),
		DiskIgnoreFs:             types.StringValue(s.DiskIgnoreFs.String()),
		Diskio:                   types.BoolValue(tools.StringToBool(s.Diskio)),
		InternetSpeed:            types.BoolValue(tools.StringToBool(s.InternetSpeed)),
		InternetSpeedSaving:      types.BoolValue(tools.StringToBool(s.InternetSpeedSaving)),
		InternetSpeedInterval:    settingsReadInt(s.InternetSpeedInterval),
		Mem:                      types.BoolValue(tools.StringToBool(s.Mem)),
		Processes:                types.BoolValue(tools.StringToBool(s.Processes)),
		Swap:                     types.BoolValue(tools.StringToBool(s.Swap)),
		System:                   types.BoolValue(tools.StringToBool(s.System)),
		Network:                  types.BoolValue(tools.StringToBool(s.Network)),
		Pf:                       types.BoolValue(tools.StringToBool(s.Pf)),
		Ping:                     types.BoolValue(tools.StringToBool(s.Ping)),
		PingCount:                settingsReadInt(s.PingCount),
		PingHosts:                types.StringValue(s.PingHosts.String()),
		Ping6:                    types.BoolValue(tools.StringToBool(s.Ping6)),
		Ping6Count:               settingsReadInt(s.Ping6Count),
		Ping6Hosts:               types.StringValue(s.Ping6Hosts.String()),
		Haproxy:                  types.BoolValue(tools.StringToBool(s.Haproxy)),
		Zfs:                      types.BoolValue(tools.StringToBool(s.Zfs)),
		Ntpq:                     types.BoolValue(tools.StringToBool(s.Ntpq)),
		NtpqDnsLookup:            types.BoolValue(tools.StringToBool(s.NtpqDnsLookup)),
		IntrusionDetectionAlerts: types.BoolValue(tools.StringToBool(s.IntrusionDetectionAlerts)),
		Unbound:                  types.BoolValue(tools.StringToBool(s.Unbound)),
		Apcupsd:                  types.BoolValue(tools.StringToBool(s.Apcupsd)),
		ApcupsdServer:            types.StringValue(s.ApcupsdServer),
	}
}
