package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type upstreamServerResourceModel struct {
	Description types.String `tfsdk:"description"`
	Server      types.String `tfsdk:"server"`
	Port        types.Int64  `tfsdk:"port"`
	Priority    types.Int64  `tfsdk:"priority"`
	MaxConns    types.Int64  `tfsdk:"max_conns"`
	MaxFails    types.Int64  `tfsdk:"max_fails"`
	FailTimeout types.Int64  `tfsdk:"fail_timeout"`
	NoUse       types.String `tfsdk:"no_use"`

	Id types.String `tfsdk:"id"`
}

func upstreamServerResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx upstream servers (**Services ▸ Nginx ▸ Configuration ▸ Upstream ▸ Upstream Server**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				MarkdownDescription: "Description.",
				Required:            true,
			},
			"server": schema.StringAttribute{
				MarkdownDescription: "Server.",
				Required:            true,
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: "Port.",
				Required:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "Server Priority.",
				Required:            true,
			},
			"max_conns": schema.Int64Attribute{
				MarkdownDescription: "Maximum Connections. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"max_fails": schema.Int64Attribute{
				MarkdownDescription: "Maximum Failures. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"fail_timeout": schema.Int64Attribute{
				MarkdownDescription: "Fail Timeout. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"no_use": schema.StringAttribute{
				MarkdownDescription: "Do Not Use. One of `down`, `backup`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "down", "backup")},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx upstream server.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func upstreamServerDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx upstream servers (**Services ▸ Nginx ▸ Configuration ▸ Upstream ▸ Upstream Server**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx upstream server.",
				Required:            true,
			},
			"description":  dschema.StringAttribute{MarkdownDescription: "Description.", Computed: true},
			"server":       dschema.StringAttribute{MarkdownDescription: "Server.", Computed: true},
			"port":         dschema.Int64Attribute{MarkdownDescription: "Port.", Computed: true},
			"priority":     dschema.Int64Attribute{MarkdownDescription: "Server Priority.", Computed: true},
			"max_conns":    dschema.Int64Attribute{MarkdownDescription: "Maximum Connections.", Computed: true},
			"max_fails":    dschema.Int64Attribute{MarkdownDescription: "Maximum Failures.", Computed: true},
			"fail_timeout": dschema.Int64Attribute{MarkdownDescription: "Fail Timeout.", Computed: true},
			"no_use":       dschema.StringAttribute{MarkdownDescription: "Do Not Use. One of `down`, `backup`.", Computed: true},
		},
	}
}

func convertUpstreamServerSchemaToStruct(d *upstreamServerResourceModel) (*nginx.UpstreamServer, error) {
	return &nginx.UpstreamServer{
		Description: d.Description.ValueString(),
		Server:      d.Server.ValueString(),
		Port:        tools.Int64ToString(d.Port.ValueInt64()),
		Priority:    tools.Int64ToString(d.Priority.ValueInt64()),
		MaxConns:    tools.Int64ToStringNegative(d.MaxConns.ValueInt64()),
		MaxFails:    tools.Int64ToStringNegative(d.MaxFails.ValueInt64()),
		FailTimeout: tools.Int64ToStringNegative(d.FailTimeout.ValueInt64()),
		NoUse:       api.SelectedMap(d.NoUse.ValueString()),
	}, nil
}

func convertUpstreamServerStructToSchema(d *nginx.UpstreamServer) (*upstreamServerResourceModel, error) {
	return &upstreamServerResourceModel{
		Description: types.StringValue(d.Description),
		Server:      types.StringValue(d.Server),
		Port:        types.Int64Value(tools.StringToInt64(d.Port)),
		Priority:    types.Int64Value(tools.StringToInt64(d.Priority)),
		MaxConns:    types.Int64Value(tools.StringToInt64(d.MaxConns)),
		MaxFails:    types.Int64Value(tools.StringToInt64(d.MaxFails)),
		FailTimeout: types.Int64Value(tools.StringToInt64(d.FailTimeout)),
		NoUse:       types.StringValue(d.NoUse.String()),
	}, nil
}
