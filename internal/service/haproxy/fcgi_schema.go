package haproxy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type fcgiResourceModel struct {
	Enabled       types.Bool   `tfsdk:"enabled"`
	Name          types.String `tfsdk:"name"`
	Description   types.String `tfsdk:"description"`
	Docroot       types.String `tfsdk:"docroot"`
	Index         types.String `tfsdk:"index"`
	PathInfo      types.String `tfsdk:"path_info"`
	LogStderr     types.Bool   `tfsdk:"log_stderr"`
	KeepConn      types.Bool   `tfsdk:"keep_conn"`
	GetValues     types.Bool   `tfsdk:"get_values"`
	MpxsConns     types.Bool   `tfsdk:"mpxs_conns"`
	MaxReqs       types.Int64  `tfsdk:"max_reqs"`
	LinkedActions types.Set    `tfsdk:"linked_actions"`

	Id types.String `tfsdk:"id"`
}

func fcgiResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "HAProxy FastCGI applications (**Services ▸ HAProxy ▸ Settings ▸ Advanced ▸ FastCGI Applications**, plugin `os-haproxy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this FastCGI application. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify this FastCGI application.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this FastCGI application. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"docroot": schema.StringAttribute{
				MarkdownDescription: "Define the document root on the remote host. It will be used to build the default value of FastCGI parameters SCRIPT_FILENAME and PATH_TRANSLATED. It is a mandatory setting.",
				Required:            true,
			},
			"index": schema.StringAttribute{
				MarkdownDescription: "Define the script name that will be appended after an URI that ends with a slash (\"/\") to set the default value of the FastCGI parameter SCRIPT_NAME. It is an optional setting. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"path_info": schema.StringAttribute{
				MarkdownDescription: "Define a regular expression to extract the script-name and the path-info from the URL-decoded path, see HAProxy's documentation for further details and examples. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"log_stderr": schema.BoolAttribute{
				MarkdownDescription: "Enable logging of STDERR messages reported by the FastCGI application. It is an optional setting. By default STDERR messages are ignored. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"keep_conn": schema.BoolAttribute{
				MarkdownDescription: "Instruct the FastCGI application to keep the connection open after sending a response. If disabled, the FastCGI application closes the connection after responding to a request. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"get_values": schema.BoolAttribute{
				MarkdownDescription: "Enable the retrieval of variables about connection management from the FastCGI application by sending the record FCGI_GET_VALUES on connection establishment. Some FastCGI applications do not support this feature and others close the connection immediately after sending their response. As a result, this option is disabled by default. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"mpxs_conns": schema.BoolAttribute{
				MarkdownDescription: "Enable the support for connection multiplexing. This option may be overwritten if the variable FCGI_MPXS_CONNS is retrieved during connection establishment. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"max_reqs": schema.Int64Attribute{
				MarkdownDescription: "Define the maximum number of concurrent requests this application will accept. This option may be overwritten if the variable FCGI_MAX_REQS is retrieved during connection establishment. Furthermore, if the application does not support connection multiplexing, this option will be ignored. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"linked_actions": schema.SetAttribute{
				MarkdownDescription: "Choose rules to be included in this FastCGI application. Only FastCGI rules are supported. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the FastCGI application.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func fcgiDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "HAProxy FastCGI applications (**Services ▸ HAProxy ▸ Settings ▸ Advanced ▸ FastCGI Applications**, plugin `os-haproxy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the FastCGI application.",
				Required:            true,
			},
			"enabled":        dschema.BoolAttribute{MarkdownDescription: "Enable this FastCGI application.", Computed: true},
			"name":           dschema.StringAttribute{MarkdownDescription: "Name to identify this FastCGI application.", Computed: true},
			"description":    dschema.StringAttribute{MarkdownDescription: "Description for this FastCGI application.", Computed: true},
			"docroot":        dschema.StringAttribute{MarkdownDescription: "Define the document root on the remote host. It will be used to build the default value of FastCGI parameters SCRIPT_FILENAME and PATH_TRANSLATED. It is a mandatory setting.", Computed: true},
			"index":          dschema.StringAttribute{MarkdownDescription: "Define the script name that will be appended after an URI that ends with a slash (\"/\") to set the default value of the FastCGI parameter SCRIPT_NAME. It is an optional setting.", Computed: true},
			"path_info":      dschema.StringAttribute{MarkdownDescription: "Define a regular expression to extract the script-name and the path-info from the URL-decoded path, see HAProxy's documentation for further details and examples.", Computed: true},
			"log_stderr":     dschema.BoolAttribute{MarkdownDescription: "Enable logging of STDERR messages reported by the FastCGI application. It is an optional setting. By default STDERR messages are ignored.", Computed: true},
			"keep_conn":      dschema.BoolAttribute{MarkdownDescription: "Instruct the FastCGI application to keep the connection open after sending a response. If disabled, the FastCGI application closes the connection after responding to a request.", Computed: true},
			"get_values":     dschema.BoolAttribute{MarkdownDescription: "Enable the retrieval of variables about connection management from the FastCGI application by sending the record FCGI_GET_VALUES on connection establishment. Some FastCGI applications do not support this feature and others close the connection immediately after sending their response. As a result, this option is disabled by default.", Computed: true},
			"mpxs_conns":     dschema.BoolAttribute{MarkdownDescription: "Enable the support for connection multiplexing. This option may be overwritten if the variable FCGI_MPXS_CONNS is retrieved during connection establishment.", Computed: true},
			"max_reqs":       dschema.Int64Attribute{MarkdownDescription: "Define the maximum number of concurrent requests this application will accept. This option may be overwritten if the variable FCGI_MAX_REQS is retrieved during connection establishment. Furthermore, if the application does not support connection multiplexing, this option will be ignored.", Computed: true},
			"linked_actions": dschema.SetAttribute{MarkdownDescription: "Choose rules to be included in this FastCGI application. Only FastCGI rules are supported.", Computed: true, ElementType: types.StringType},
		},
	}
}

func convertFcgiSchemaToStruct(d *fcgiResourceModel) (*haproxy.Fcgi, error) {
	return &haproxy.Fcgi{
		Enabled:       tools.BoolToString(d.Enabled.ValueBool()),
		Name:          d.Name.ValueString(),
		Description:   d.Description.ValueString(),
		Docroot:       d.Docroot.ValueString(),
		Index:         d.Index.ValueString(),
		PathInfo:      d.PathInfo.ValueString(),
		LogStderr:     tools.BoolToString(d.LogStderr.ValueBool()),
		KeepConn:      tools.BoolToString(d.KeepConn.ValueBool()),
		GetValues:     tools.BoolToString(d.GetValues.ValueBool()),
		MpxsConns:     tools.BoolToString(d.MpxsConns.ValueBool()),
		MaxReqs:       tools.Int64ToStringNegative(d.MaxReqs.ValueInt64()),
		LinkedActions: api.SelectedMapList(setToSlice(d.LinkedActions)),
	}, nil
}

func convertFcgiStructToSchema(d *haproxy.Fcgi) (*fcgiResourceModel, error) {
	return &fcgiResourceModel{
		Enabled:       types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:          types.StringValue(d.Name),
		Description:   types.StringValue(d.Description),
		Docroot:       types.StringValue(d.Docroot),
		Index:         types.StringValue(d.Index),
		PathInfo:      types.StringValue(d.PathInfo),
		LogStderr:     types.BoolValue(tools.StringToBool(d.LogStderr)),
		KeepConn:      types.BoolValue(tools.StringToBool(d.KeepConn)),
		GetValues:     types.BoolValue(tools.StringToBool(d.GetValues)),
		MpxsConns:     types.BoolValue(tools.StringToBool(d.MpxsConns)),
		MaxReqs:       types.Int64Value(tools.StringToInt64(d.MaxReqs)),
		LinkedActions: tools.StringSliceToSet(d.LinkedActions),
	}, nil
}
