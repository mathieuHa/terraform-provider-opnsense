package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type syslogTargetResourceModel struct {
	Description types.String `tfsdk:"description"`
	Host        types.String `tfsdk:"host"`
	Port        types.Int64  `tfsdk:"port"`
	Facility    types.String `tfsdk:"facility"`
	Severity    types.String `tfsdk:"severity"`
	Tag         types.String `tfsdk:"tag"`
	Nohostname  types.Bool   `tfsdk:"nohostname"`

	Id types.String `tfsdk:"id"`
}

func syslogTargetResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx syslog targets (**Services ▸ Nginx ▸ Configuration ▸ Other ▸ Syslog Target**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				MarkdownDescription: "Description.",
				Required:            true,
			},
			"host": schema.StringAttribute{
				MarkdownDescription: "Host.",
				Required:            true,
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: "Port. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"facility": schema.StringAttribute{
				MarkdownDescription: "Facility. One of `kern`, `user`, `mail`, `daemon`, `auth`, `intern`, `lpr`, `news`, `uucp`, `clock`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `cron`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`. Defaults to `local7`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("local7"),
				Validators:          []validator.String{stringvalidator.OneOf("kern", "user", "mail", "daemon", "auth", "intern", "lpr", "news", "uucp", "clock", "authpriv", "ftp", "ntp", "audit", "alert", "cron", "local0", "local1", "local2", "local3", "local4", "local5", "local6", "local7")},
			},
			"severity": schema.StringAttribute{
				MarkdownDescription: "Severity. One of `debug`, `info`, `notice`, `warn`, `error`, `crit`, `alert`, `emerg`. Defaults to `error`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("error"),
				Validators:          []validator.String{stringvalidator.OneOf("debug", "info", "notice", "warn", "error", "crit", "alert", "emerg")},
			},
			"tag": schema.StringAttribute{
				MarkdownDescription: "Tag. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"nohostname": schema.BoolAttribute{
				MarkdownDescription: "No Hostname. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx syslog target.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func syslogTargetDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx syslog targets (**Services ▸ Nginx ▸ Configuration ▸ Other ▸ Syslog Target**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx syslog target.",
				Required:            true,
			},
			"description": dschema.StringAttribute{MarkdownDescription: "Description.", Computed: true},
			"host":        dschema.StringAttribute{MarkdownDescription: "Host.", Computed: true},
			"port":        dschema.Int64Attribute{MarkdownDescription: "Port.", Computed: true},
			"facility":    dschema.StringAttribute{MarkdownDescription: "Facility. One of `kern`, `user`, `mail`, `daemon`, `auth`, `intern`, `lpr`, `news`, `uucp`, `clock`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `cron`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.", Computed: true},
			"severity":    dschema.StringAttribute{MarkdownDescription: "Severity. One of `debug`, `info`, `notice`, `warn`, `error`, `crit`, `alert`, `emerg`.", Computed: true},
			"tag":         dschema.StringAttribute{MarkdownDescription: "Tag.", Computed: true},
			"nohostname":  dschema.BoolAttribute{MarkdownDescription: "No Hostname.", Computed: true},
		},
	}
}

func convertSyslogTargetSchemaToStruct(d *syslogTargetResourceModel) (*nginx.SyslogTarget, error) {
	return &nginx.SyslogTarget{
		Description: d.Description.ValueString(),
		Host:        d.Host.ValueString(),
		Port:        tools.Int64ToStringNegative(d.Port.ValueInt64()),
		Facility:    api.SelectedMap(d.Facility.ValueString()),
		Severity:    api.SelectedMap(d.Severity.ValueString()),
		Tag:         d.Tag.ValueString(),
		Nohostname:  tools.BoolToString(d.Nohostname.ValueBool()),
	}, nil
}

func convertSyslogTargetStructToSchema(d *nginx.SyslogTarget) (*syslogTargetResourceModel, error) {
	return &syslogTargetResourceModel{
		Description: types.StringValue(d.Description),
		Host:        types.StringValue(d.Host),
		Port:        types.Int64Value(tools.StringToInt64(d.Port)),
		Facility:    types.StringValue(d.Facility.String()),
		Severity:    types.StringValue(d.Severity.String()),
		Tag:         types.StringValue(d.Tag),
		Nohostname:  types.BoolValue(tools.StringToBool(d.Nohostname)),
	}, nil
}
