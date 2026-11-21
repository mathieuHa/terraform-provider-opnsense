package monit

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/monit"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type serviceResourceModel struct {
	Enabled      types.Bool   `tfsdk:"enabled"`
	Name         types.String `tfsdk:"name"`
	Type         types.String `tfsdk:"type"`
	Pidfile      types.String `tfsdk:"pidfile"`
	Match        types.String `tfsdk:"match"`
	Path         types.String `tfsdk:"path"`
	Timeout      types.Int64  `tfsdk:"timeout"`
	StartTimeout types.Int64  `tfsdk:"start_timeout"`
	Address      types.String `tfsdk:"address"`
	Interface    types.String `tfsdk:"interface"`
	Start        types.String `tfsdk:"start"`
	Stop         types.String `tfsdk:"stop"`
	Tests        types.Set    `tfsdk:"tests"`
	Depends      types.Set    `tfsdk:"depends"`
	Polltime     types.String `tfsdk:"polltime"`
	Description  types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func serviceResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Services monitored by Monit (**Services ▸ Monit ▸ Settings ▸ Service Settings**), each checked against a set of tests.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether this service is monitored. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Unique name of the service.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "What is monitored.",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf("process", "file", "fifo", "filesystem", "directory", "host", "system", "custom", "network")},
			},
			"pidfile": schema.StringAttribute{
				MarkdownDescription: "PID file of the process, for `process` services.",
				Optional:            true,
			},
			"match": schema.StringAttribute{
				MarkdownDescription: "Process name pattern, for `process` services without a PID file.",
				Optional:            true,
			},
			"path": schema.StringAttribute{
				MarkdownDescription: "Path to monitor, or the script for `custom` services.",
				Optional:            true,
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: "Script timeout in seconds, for `custom` services. Defaults to `300`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(300),
				Validators:          []validator.Int64{int64validator.AtLeast(1)},
			},
			"start_timeout": schema.Int64Attribute{
				MarkdownDescription: "Seconds to wait for the service to start. Defaults to `30`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(30),
				Validators:          []validator.Int64{int64validator.AtLeast(1)},
			},
			"address": schema.StringAttribute{
				MarkdownDescription: "Address to check, for `host` services.",
				Optional:            true,
			},
			"interface": schema.StringAttribute{
				MarkdownDescription: "Interface to check, for `network` services.",
				Optional:            true,
			},
			"start": schema.StringAttribute{
				MarkdownDescription: "Command used to start the service.",
				Optional:            true,
			},
			"stop": schema.StringAttribute{
				MarkdownDescription: "Command used to stop the service.",
				Optional:            true,
			},
			"tests": schema.SetAttribute{
				MarkdownDescription: "UUIDs of the tests applied to this service. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"depends": schema.SetAttribute{
				MarkdownDescription: "UUIDs of the services this one depends on. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"polltime": schema.StringAttribute{
				MarkdownDescription: "Polling interval, e.g. `every 2 cycles`.",
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the Monit service.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func serviceDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Services monitored by Monit (**Services ▸ Monit ▸ Settings ▸ Service Settings**), each checked against a set of tests.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the Monit service.",
				Required:            true,
			},
			"enabled":       dschema.BoolAttribute{MarkdownDescription: "Whether this service is monitored.", Computed: true},
			"name":          dschema.StringAttribute{MarkdownDescription: "Unique name of the service.", Computed: true},
			"type":          dschema.StringAttribute{MarkdownDescription: "What is monitored.", Computed: true},
			"pidfile":       dschema.StringAttribute{MarkdownDescription: "PID file of the process, for `process` services.", Computed: true},
			"match":         dschema.StringAttribute{MarkdownDescription: "Process name pattern, for `process` services without a PID file.", Computed: true},
			"path":          dschema.StringAttribute{MarkdownDescription: "Path to monitor, or the script for `custom` services.", Computed: true},
			"timeout":       dschema.Int64Attribute{MarkdownDescription: "Script timeout in seconds, for `custom` services.", Computed: true},
			"start_timeout": dschema.Int64Attribute{MarkdownDescription: "Seconds to wait for the service to start.", Computed: true},
			"address":       dschema.StringAttribute{MarkdownDescription: "Address to check, for `host` services.", Computed: true},
			"interface":     dschema.StringAttribute{MarkdownDescription: "Interface to check, for `network` services.", Computed: true},
			"start":         dschema.StringAttribute{MarkdownDescription: "Command used to start the service.", Computed: true},
			"stop":          dschema.StringAttribute{MarkdownDescription: "Command used to stop the service.", Computed: true},
			"tests":         dschema.SetAttribute{MarkdownDescription: "UUIDs of the tests applied to this service.", Computed: true, ElementType: types.StringType},
			"depends":       dschema.SetAttribute{MarkdownDescription: "UUIDs of the services this one depends on.", Computed: true, ElementType: types.StringType},
			"polltime":      dschema.StringAttribute{MarkdownDescription: "Polling interval, e.g. `every 2 cycles`.", Computed: true},
			"description":   dschema.StringAttribute{MarkdownDescription: "Optional description here for your reference (not parsed).", Computed: true},
		},
	}
}

func convertServiceSchemaToStruct(d *serviceResourceModel) (*monit.Service, error) {
	return &monit.Service{
		Enabled:      tools.BoolToString(d.Enabled.ValueBool()),
		Name:         d.Name.ValueString(),
		Type:         api.SelectedMap(d.Type.ValueString()),
		Pidfile:      d.Pidfile.ValueString(),
		Match:        d.Match.ValueString(),
		Path:         d.Path.ValueString(),
		Timeout:      tools.Int64ToString(d.Timeout.ValueInt64()),
		Starttimeout: tools.Int64ToString(d.StartTimeout.ValueInt64()),
		Address:      d.Address.ValueString(),
		Interface:    api.SelectedMap(d.Interface.ValueString()),
		Start:        d.Start.ValueString(),
		Stop:         d.Stop.ValueString(),
		Tests:        api.SelectedMapList(setToSlice(d.Tests)),
		Depends:      api.SelectedMapList(setToSlice(d.Depends)),
		Polltime:     d.Polltime.ValueString(),
		Description:  d.Description.ValueString(),
	}, nil
}

func convertServiceStructToSchema(d *monit.Service) (*serviceResourceModel, error) {
	return &serviceResourceModel{
		Enabled:      types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:         types.StringValue(d.Name),
		Type:         types.StringValue(d.Type.String()),
		Pidfile:      tools.StringOrNull(d.Pidfile),
		Match:        tools.StringOrNull(d.Match),
		Path:         tools.StringOrNull(d.Path),
		Timeout:      types.Int64Value(tools.StringToInt64(d.Timeout)),
		StartTimeout: types.Int64Value(tools.StringToInt64(d.Starttimeout)),
		Address:      tools.StringOrNull(d.Address),
		Interface:    tools.StringOrNull(d.Interface.String()),
		Start:        tools.StringOrNull(d.Start),
		Stop:         tools.StringOrNull(d.Stop),
		Tests:        tools.StringSliceToSet(d.Tests),
		Depends:      tools.StringSliceToSet(d.Depends),
		Polltime:     tools.StringOrNull(d.Polltime),
		Description:  tools.StringOrNull(d.Description),
	}, nil
}
