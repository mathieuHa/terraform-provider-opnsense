package haproxy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/haproxy"
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

type healthcheckResourceModel struct {
	Name                  types.String `tfsdk:"name"`
	Description           types.String `tfsdk:"description"`
	Type                  types.String `tfsdk:"type"`
	Interval              types.String `tfsdk:"interval"`
	Ssl                   types.String `tfsdk:"ssl"`
	SslSni                types.String `tfsdk:"ssl_sni"`
	ForceSsl              types.Bool   `tfsdk:"force_ssl"`
	Checkport             types.Int64  `tfsdk:"checkport"`
	HttpMethod            types.String `tfsdk:"http_method"`
	HttpUri               types.String `tfsdk:"http_uri"`
	HttpVersion           types.String `tfsdk:"http_version"`
	HttpHost              types.String `tfsdk:"http_host"`
	HttpExpressionEnabled types.Bool   `tfsdk:"http_expression_enabled"`
	HttpExpression        types.String `tfsdk:"http_expression"`
	HttpNegate            types.Bool   `tfsdk:"http_negate"`
	HttpValue             types.String `tfsdk:"http_value"`
	TcpEnabled            types.Bool   `tfsdk:"tcp_enabled"`
	TcpSendValue          types.String `tfsdk:"tcp_send_value"`
	TcpMatchType          types.String `tfsdk:"tcp_match_type"`
	TcpNegate             types.Bool   `tfsdk:"tcp_negate"`
	TcpMatchValue         types.String `tfsdk:"tcp_match_value"`
	AgentPort             types.Int64  `tfsdk:"agent_port"`
	MysqlUser             types.String `tfsdk:"mysql_user"`
	MysqlPost41           types.Bool   `tfsdk:"mysql_post41"`
	PgsqlUser             types.String `tfsdk:"pgsql_user"`
	SmtpDomain            types.String `tfsdk:"smtp_domain"`
	EsmtpDomain           types.String `tfsdk:"esmtp_domain"`
	DbUser                types.String `tfsdk:"db_user"`

	Id types.String `tfsdk:"id"`
}

func healthcheckResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "HAProxy health monitors (**Services ▸ HAProxy ▸ Settings ▸ Rules & Checks ▸ Health Monitors**, plugin `os-haproxy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify this Health Monitor.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this Health Monitor. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Select type of health check. One of `tcp`, `http`, `agent`, `ldap`, `mysql`, `pgsql`, `redis`, `smtp`, `esmtp`, `ssl`. Defaults to `http`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http"),
				Validators:          []validator.String{stringvalidator.OneOf("tcp", "http", "agent", "ldap", "mysql", "pgsql", "redis", "smtp", "esmtp", "ssl")},
			},
			"interval": schema.StringAttribute{
				MarkdownDescription: "Select interval (in milliseconds) between two consecutive health checks. This value can be overriden in backend pool and real server configuration. Defaults to `2s`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("2s"),
			},
			"ssl": schema.StringAttribute{
				MarkdownDescription: "This option forces encryption of all health checks over SSL, regardless of whether the server uses SSL or not for the normal traffic. One of `nopref`, `ssl`, `sslsni`, `nossl`. Defaults to `nopref`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("nopref"),
				Validators:          []validator.String{stringvalidator.OneOf("", "nopref", "ssl", "sslsni", "nossl")},
			},
			"ssl_sni": schema.StringAttribute{
				MarkdownDescription: "The host name sent in the SNI TLS extension to the server. Overrides the server's SNI setting for health checks. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"force_ssl": schema.BoolAttribute{
				MarkdownDescription: "force_ssl. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"checkport": schema.Int64Attribute{
				MarkdownDescription: "Provide an optional TCP communication port to use during health checks, i.e. 80 or 443. The server port will used for health checks when no port is specified. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"http_method": schema.StringAttribute{
				MarkdownDescription: "Select HTTP method for health check. One of `options`, `head`, `get`, `put`, `post`, `delete`, `trace`. Defaults to `options`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("options"),
				Validators:          []validator.String{stringvalidator.OneOf("", "options", "head", "get", "put", "post", "delete", "trace")},
			},
			"http_uri": schema.StringAttribute{
				MarkdownDescription: "Specify HTTP request URI for health check. Defaults to `/`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/"),
			},
			"http_version": schema.StringAttribute{
				MarkdownDescription: "Select HTTP version for a HTTP health check. One of `http10`, `http11`, `http2`. Defaults to `http10`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http10"),
				Validators:          []validator.String{stringvalidator.OneOf("", "http10", "http11", "http2")},
			},
			"http_host": schema.StringAttribute{
				MarkdownDescription: "Specify HTTP host to use for health check. Requires HTTP/1.1. Defaults to `localhost`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("localhost"),
			},
			"http_expression_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enabled. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_expression": schema.StringAttribute{
				MarkdownDescription: "Select health check expression. One of `status`, `rstatus`, `string`, `rstring`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "status", "rstatus", "string", "rstring")},
			},
			"http_negate": schema.BoolAttribute{
				MarkdownDescription: "Use this to invert the meaning of the expression. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_value": schema.StringAttribute{
				MarkdownDescription: "Specify a value to match with the expression. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tcp_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enabled. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tcp_send_value": schema.StringAttribute{
				MarkdownDescription: "Specify a value to match with the expression. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tcp_match_type": schema.StringAttribute{
				MarkdownDescription: "Select how to look for a specific pattern in the response. One of `string`, `rstring`, `binary`. Defaults to `string`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("string"),
				Validators:          []validator.String{stringvalidator.OneOf("", "string", "rstring", "binary")},
			},
			"tcp_negate": schema.BoolAttribute{
				MarkdownDescription: "Use this to invert the meaning of the expression. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tcp_match_value": schema.StringAttribute{
				MarkdownDescription: "Specify the pattern to look for in the response buffer. If the match is set to binary, then the pattern must be passed as a serie of hexadecimal digits in an even number. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"agent_port": schema.Int64Attribute{
				MarkdownDescription: "Specify the TCP port used for agent checks. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"mysql_user": schema.StringAttribute{
				MarkdownDescription: "Specify the username to be used for database health checks. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"mysql_post41": schema.BoolAttribute{
				MarkdownDescription: "Send checks compatible with MySQL server 4.1 and later. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pgsql_user": schema.StringAttribute{
				MarkdownDescription: "Specify the username to be used for database health checks. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"smtp_domain": schema.StringAttribute{
				MarkdownDescription: "Specify the domain name to present to the server for SMTP health checks. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"esmtp_domain": schema.StringAttribute{
				MarkdownDescription: "Specify the domain name to present to the server for ESMTP health checks. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"db_user": schema.StringAttribute{
				MarkdownDescription: "dbUser. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the health monitor.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func healthcheckDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "HAProxy health monitors (**Services ▸ HAProxy ▸ Settings ▸ Rules & Checks ▸ Health Monitors**, plugin `os-haproxy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the health monitor.",
				Required:            true,
			},
			"name":                    dschema.StringAttribute{MarkdownDescription: "Name to identify this Health Monitor.", Computed: true},
			"description":             dschema.StringAttribute{MarkdownDescription: "Description for this Health Monitor.", Computed: true},
			"type":                    dschema.StringAttribute{MarkdownDescription: "Select type of health check. One of `tcp`, `http`, `agent`, `ldap`, `mysql`, `pgsql`, `redis`, `smtp`, `esmtp`, `ssl`.", Computed: true},
			"interval":                dschema.StringAttribute{MarkdownDescription: "Select interval (in milliseconds) between two consecutive health checks. This value can be overriden in backend pool and real server configuration.", Computed: true},
			"ssl":                     dschema.StringAttribute{MarkdownDescription: "This option forces encryption of all health checks over SSL, regardless of whether the server uses SSL or not for the normal traffic. One of `nopref`, `ssl`, `sslsni`, `nossl`.", Computed: true},
			"ssl_sni":                 dschema.StringAttribute{MarkdownDescription: "The host name sent in the SNI TLS extension to the server. Overrides the server's SNI setting for health checks.", Computed: true},
			"force_ssl":               dschema.BoolAttribute{MarkdownDescription: "force_ssl.", Computed: true},
			"checkport":               dschema.Int64Attribute{MarkdownDescription: "Provide an optional TCP communication port to use during health checks, i.e. 80 or 443. The server port will used for health checks when no port is specified.", Computed: true},
			"http_method":             dschema.StringAttribute{MarkdownDescription: "Select HTTP method for health check. One of `options`, `head`, `get`, `put`, `post`, `delete`, `trace`.", Computed: true},
			"http_uri":                dschema.StringAttribute{MarkdownDescription: "Specify HTTP request URI for health check.", Computed: true},
			"http_version":            dschema.StringAttribute{MarkdownDescription: "Select HTTP version for a HTTP health check. One of `http10`, `http11`, `http2`.", Computed: true},
			"http_host":               dschema.StringAttribute{MarkdownDescription: "Specify HTTP host to use for health check. Requires HTTP/1.1.", Computed: true},
			"http_expression_enabled": dschema.BoolAttribute{MarkdownDescription: "Enabled.", Computed: true},
			"http_expression":         dschema.StringAttribute{MarkdownDescription: "Select health check expression. One of `status`, `rstatus`, `string`, `rstring`.", Computed: true},
			"http_negate":             dschema.BoolAttribute{MarkdownDescription: "Use this to invert the meaning of the expression.", Computed: true},
			"http_value":              dschema.StringAttribute{MarkdownDescription: "Specify a value to match with the expression.", Computed: true},
			"tcp_enabled":             dschema.BoolAttribute{MarkdownDescription: "Enabled.", Computed: true},
			"tcp_send_value":          dschema.StringAttribute{MarkdownDescription: "Specify a value to match with the expression.", Computed: true},
			"tcp_match_type":          dschema.StringAttribute{MarkdownDescription: "Select how to look for a specific pattern in the response. One of `string`, `rstring`, `binary`.", Computed: true},
			"tcp_negate":              dschema.BoolAttribute{MarkdownDescription: "Use this to invert the meaning of the expression.", Computed: true},
			"tcp_match_value":         dschema.StringAttribute{MarkdownDescription: "Specify the pattern to look for in the response buffer. If the match is set to binary, then the pattern must be passed as a serie of hexadecimal digits in an even number.", Computed: true},
			"agent_port":              dschema.Int64Attribute{MarkdownDescription: "Specify the TCP port used for agent checks.", Computed: true},
			"mysql_user":              dschema.StringAttribute{MarkdownDescription: "Specify the username to be used for database health checks.", Computed: true},
			"mysql_post41":            dschema.BoolAttribute{MarkdownDescription: "Send checks compatible with MySQL server 4.1 and later.", Computed: true},
			"pgsql_user":              dschema.StringAttribute{MarkdownDescription: "Specify the username to be used for database health checks.", Computed: true},
			"smtp_domain":             dschema.StringAttribute{MarkdownDescription: "Specify the domain name to present to the server for SMTP health checks.", Computed: true},
			"esmtp_domain":            dschema.StringAttribute{MarkdownDescription: "Specify the domain name to present to the server for ESMTP health checks.", Computed: true},
			"db_user":                 dschema.StringAttribute{MarkdownDescription: "dbUser.", Computed: true},
		},
	}
}

func convertHealthcheckSchemaToStruct(d *healthcheckResourceModel) (*haproxy.Healthcheck, error) {
	return &haproxy.Healthcheck{
		Name:                  d.Name.ValueString(),
		Description:           d.Description.ValueString(),
		Type:                  api.SelectedMap(d.Type.ValueString()),
		Interval:              d.Interval.ValueString(),
		Ssl:                   api.SelectedMap(d.Ssl.ValueString()),
		SslSni:                d.SslSni.ValueString(),
		ForceSsl:              tools.BoolToString(d.ForceSsl.ValueBool()),
		Checkport:             tools.Int64ToStringNegative(d.Checkport.ValueInt64()),
		HttpMethod:            api.SelectedMap(d.HttpMethod.ValueString()),
		HttpUri:               d.HttpUri.ValueString(),
		HttpVersion:           api.SelectedMap(d.HttpVersion.ValueString()),
		HttpHost:              d.HttpHost.ValueString(),
		HttpExpressionEnabled: tools.BoolToString(d.HttpExpressionEnabled.ValueBool()),
		HttpExpression:        api.SelectedMap(d.HttpExpression.ValueString()),
		HttpNegate:            tools.BoolToString(d.HttpNegate.ValueBool()),
		HttpValue:             d.HttpValue.ValueString(),
		TcpEnabled:            tools.BoolToString(d.TcpEnabled.ValueBool()),
		TcpSendValue:          d.TcpSendValue.ValueString(),
		TcpMatchType:          api.SelectedMap(d.TcpMatchType.ValueString()),
		TcpNegate:             tools.BoolToString(d.TcpNegate.ValueBool()),
		TcpMatchValue:         d.TcpMatchValue.ValueString(),
		AgentPort:             tools.Int64ToStringNegative(d.AgentPort.ValueInt64()),
		MysqlUser:             d.MysqlUser.ValueString(),
		MysqlPost41:           tools.BoolToString(d.MysqlPost41.ValueBool()),
		PgsqlUser:             d.PgsqlUser.ValueString(),
		SmtpDomain:            d.SmtpDomain.ValueString(),
		EsmtpDomain:           d.EsmtpDomain.ValueString(),
		DbUser:                d.DbUser.ValueString(),
	}, nil
}

func convertHealthcheckStructToSchema(d *haproxy.Healthcheck) (*healthcheckResourceModel, error) {
	return &healthcheckResourceModel{
		Name:                  types.StringValue(d.Name),
		Description:           types.StringValue(d.Description),
		Type:                  types.StringValue(d.Type.String()),
		Interval:              types.StringValue(d.Interval),
		Ssl:                   types.StringValue(d.Ssl.String()),
		SslSni:                types.StringValue(d.SslSni),
		ForceSsl:              types.BoolValue(tools.StringToBool(d.ForceSsl)),
		Checkport:             types.Int64Value(tools.StringToInt64(d.Checkport)),
		HttpMethod:            types.StringValue(d.HttpMethod.String()),
		HttpUri:               types.StringValue(d.HttpUri),
		HttpVersion:           types.StringValue(d.HttpVersion.String()),
		HttpHost:              types.StringValue(d.HttpHost),
		HttpExpressionEnabled: types.BoolValue(tools.StringToBool(d.HttpExpressionEnabled)),
		HttpExpression:        types.StringValue(d.HttpExpression.String()),
		HttpNegate:            types.BoolValue(tools.StringToBool(d.HttpNegate)),
		HttpValue:             types.StringValue(d.HttpValue),
		TcpEnabled:            types.BoolValue(tools.StringToBool(d.TcpEnabled)),
		TcpSendValue:          types.StringValue(d.TcpSendValue),
		TcpMatchType:          types.StringValue(d.TcpMatchType.String()),
		TcpNegate:             types.BoolValue(tools.StringToBool(d.TcpNegate)),
		TcpMatchValue:         types.StringValue(d.TcpMatchValue),
		AgentPort:             types.Int64Value(tools.StringToInt64(d.AgentPort)),
		MysqlUser:             types.StringValue(d.MysqlUser),
		MysqlPost41:           types.BoolValue(tools.StringToBool(d.MysqlPost41)),
		PgsqlUser:             types.StringValue(d.PgsqlUser),
		SmtpDomain:            types.StringValue(d.SmtpDomain),
		EsmtpDomain:           types.StringValue(d.EsmtpDomain),
		DbUser:                types.StringValue(d.DbUser),
	}, nil
}
