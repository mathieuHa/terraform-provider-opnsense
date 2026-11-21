package syslog

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/syslog"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type destinationResourceModel struct {
	Enabled     types.Bool   `tfsdk:"enabled"`
	Transport   types.String `tfsdk:"transport"`
	Hostname    types.String `tfsdk:"hostname"`
	Port        types.Int64  `tfsdk:"port"`
	Level       types.Set    `tfsdk:"level"`
	Facility    types.Set    `tfsdk:"facility"`
	Program     types.Set    `tfsdk:"program"`
	Certificate types.String `tfsdk:"certificate"`
	Rfc5424     types.Bool   `tfsdk:"rfc5424"`
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func destinationResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Remote syslog destinations (**System ▸ Settings ▸ Logging ▸ Remote**). Each destination forwards the selected programs, levels and facilities to a remote collector.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether this destination is enabled. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"transport": schema.StringAttribute{
				MarkdownDescription: "Transport protocol and address family. Defaults to `udp4`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("udp4"),
				Validators:          []validator.String{stringvalidator.OneOf("udp4", "tcp4", "tls4", "udp6", "tcp6", "tls6")},
			},
			"hostname": schema.StringAttribute{
				MarkdownDescription: "Hostname or IP address of the remote collector.",
				Required:            true,
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: "Port of the remote collector. Defaults to `514`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(514),
				Validators:          []validator.Int64{int64validator.Between(1, 65535)},
			},
			"level": schema.SetAttribute{
				MarkdownDescription: "Severities to forward. Empty forwards all. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
				Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("emerg", "alert", "crit", "err", "warn", "notice", "info", "debug"))},
			},
			"facility": schema.SetAttribute{
				MarkdownDescription: "Facilities to forward, e.g. `local0`. Empty forwards all. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"program": schema.SetAttribute{
				MarkdownDescription: "Programs to forward, e.g. `filterlog`. Empty forwards all. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"certificate": schema.StringAttribute{
				MarkdownDescription: "Refid of the client certificate used for TLS transports.",
				Optional:            true,
			},
			"rfc5424": schema.BoolAttribute{
				MarkdownDescription: "When enabled, messages are sent in RFC 5424 format instead of the BSD format. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the syslog destination.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func destinationDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Remote syslog destinations (**System ▸ Settings ▸ Logging ▸ Remote**). Each destination forwards the selected programs, levels and facilities to a remote collector.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the syslog destination.",
				Required:            true,
			},
			"enabled":     dschema.BoolAttribute{MarkdownDescription: "Whether this destination is enabled.", Computed: true},
			"transport":   dschema.StringAttribute{MarkdownDescription: "Transport protocol and address family.", Computed: true},
			"hostname":    dschema.StringAttribute{MarkdownDescription: "Hostname or IP address of the remote collector.", Computed: true},
			"port":        dschema.Int64Attribute{MarkdownDescription: "Port of the remote collector.", Computed: true},
			"level":       dschema.SetAttribute{MarkdownDescription: "Severities to forward. Empty forwards all.", Computed: true, ElementType: types.StringType},
			"facility":    dschema.SetAttribute{MarkdownDescription: "Facilities to forward, e.g. `local0`. Empty forwards all.", Computed: true, ElementType: types.StringType},
			"program":     dschema.SetAttribute{MarkdownDescription: "Programs to forward, e.g. `filterlog`. Empty forwards all.", Computed: true, ElementType: types.StringType},
			"certificate": dschema.StringAttribute{MarkdownDescription: "Refid of the client certificate used for TLS transports.", Computed: true},
			"rfc5424":     dschema.BoolAttribute{MarkdownDescription: "When enabled, messages are sent in RFC 5424 format instead of the BSD format.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Optional description here for your reference (not parsed).", Computed: true},
		},
	}
}

func convertDestinationSchemaToStruct(d *destinationResourceModel) (*syslog.Destination, error) {
	return &syslog.Destination{
		Enabled:     tools.BoolToString(d.Enabled.ValueBool()),
		Transport:   api.SelectedMap(d.Transport.ValueString()),
		Hostname:    d.Hostname.ValueString(),
		Port:        tools.Int64ToString(d.Port.ValueInt64()),
		Level:       api.SelectedMapList(setToSlice(d.Level)),
		Facility:    api.SelectedMapList(setToSlice(d.Facility)),
		Program:     api.SelectedMapList(setToSlice(d.Program)),
		Certificate: api.SelectedMap(d.Certificate.ValueString()),
		Rfc5424:     tools.BoolToString(d.Rfc5424.ValueBool()),
		Description: d.Description.ValueString(),
	}, nil
}

func convertDestinationStructToSchema(d *syslog.Destination) (*destinationResourceModel, error) {
	return &destinationResourceModel{
		Enabled:     types.BoolValue(tools.StringToBool(d.Enabled)),
		Transport:   types.StringValue(d.Transport.String()),
		Hostname:    types.StringValue(d.Hostname),
		Port:        types.Int64Value(tools.StringToInt64(d.Port)),
		Level:       tools.StringSliceToSet(d.Level),
		Facility:    tools.StringSliceToSet(d.Facility),
		Program:     tools.StringSliceToSet(d.Program),
		Certificate: tools.StringOrNull(d.Certificate.String()),
		Rfc5424:     types.BoolValue(tools.StringToBool(d.Rfc5424)),
		Description: tools.StringOrNull(d.Description),
	}, nil
}
