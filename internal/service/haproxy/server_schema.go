package haproxy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
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

type serverResourceModel struct {
	Enabled              types.Bool   `tfsdk:"enabled"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	Address              types.String `tfsdk:"address"`
	Port                 types.Int64  `tfsdk:"port"`
	Checkport            types.Int64  `tfsdk:"checkport"`
	Mode                 types.String `tfsdk:"mode"`
	MultiplexerProtocol  types.String `tfsdk:"multiplexer_protocol"`
	Type                 types.String `tfsdk:"type"`
	ServiceName          types.String `tfsdk:"service_name"`
	Number               types.String `tfsdk:"number"`
	LinkedResolver       types.String `tfsdk:"linked_resolver"`
	ResolverOpts         types.Set    `tfsdk:"resolver_opts"`
	ResolvePrefer        types.String `tfsdk:"resolve_prefer"`
	Ssl                  types.Bool   `tfsdk:"ssl"`
	SslSni               types.String `tfsdk:"ssl_sni"`
	SslSniExpr           types.String `tfsdk:"ssl_sni_expr"`
	SslVerify            types.Bool   `tfsdk:"ssl_verify"`
	SslCa                types.Set    `tfsdk:"ssl_ca"`
	SslCrl               types.String `tfsdk:"ssl_crl"`
	SslClientCertificate types.String `tfsdk:"ssl_client_certificate"`
	MaxConnections       types.Int64  `tfsdk:"max_connections"`
	Weight               types.Int64  `tfsdk:"weight"`
	CheckInterval        types.String `tfsdk:"check_interval"`
	CheckDownInterval    types.String `tfsdk:"check_down_interval"`
	Source               types.String `tfsdk:"source"`
	Advanced             types.String `tfsdk:"advanced"`
	UnixSocket           types.String `tfsdk:"unix_socket"`

	Id types.String `tfsdk:"id"`
}

func serverResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "HAProxy servers (**Services ▸ HAProxy ▸ Settings ▸ Real Servers**, plugin `os-haproxy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this server. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify a static server. When creating a server template, then this prefix is used for the server names to be built.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this server. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"address": schema.StringAttribute{
				MarkdownDescription: "Provide either the FQDN or the IP address of this server. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: "Provide the TCP or UDP communication port for this server, i.e. 80 or 443. If set, all connections will be sent to this port. If unset, the same port the client connected to will be used. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"checkport": schema.Int64Attribute{
				MarkdownDescription: "Provide the TCP communication port to use during check, i.e. 80 or 443. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Sets the operation mode to use for this server. One of `active`, `backup`, `disabled`. Defaults to `active`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("active"),
				Validators:          []validator.String{stringvalidator.OneOf("", "active", "backup", "disabled")},
			},
			"multiplexer_protocol": schema.StringAttribute{
				MarkdownDescription: "Forces the multiplexer's protocol to use for the outgoing connections to this server. It must be compatible with the mode of the backend (TCP or HTTP). It must also be usable on the backend side. Idea behind this option is to bypass the selection of the best multiplexer's protocol for all connections established to this server. Use only when strictly necessary. One of `unspecified`, `fcgi`, `h2`, `h1`. Defaults to `unspecified`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("unspecified"),
				Validators:          []validator.String{stringvalidator.OneOf("", "unspecified", "fcgi", "h2", "h1")},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Either configure a static server or a template to initialize multiple servers with shared parameters. One of `static`, `template`, `unix`. Defaults to `static`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("static"),
				Validators:          []validator.String{stringvalidator.OneOf("static", "template", "unix")},
			},
			"service_name": schema.StringAttribute{
				MarkdownDescription: "Provide either the FQDN for all the servers this template initializes or a service name to discover the available services via DNS SRV records. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"number": schema.StringAttribute{
				MarkdownDescription: "The number of servers this template initializes, i.e. 5 or 1-5. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"linked_resolver": schema.StringAttribute{
				MarkdownDescription: "Specify the resolver that the server template should look at to discover available services via DNS. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"resolver_opts": schema.SetAttribute{
				MarkdownDescription: "Add resolver options. Use TAB key to complete typing. One of `allow-dup-ip`, `ignore-weight`, `prevent-dup-ip`. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
				Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("allow-dup-ip", "ignore-weight", "prevent-dup-ip"))},
			},
			"resolve_prefer": schema.StringAttribute{
				MarkdownDescription: "When DNS resolution is enabled for a server and multiple IP addresses from different families are returned, HAProxy will prefer using an IP address from the selected family. One of `ipv4`, `ipv6`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "ipv4", "ipv6")},
			},
			"ssl": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable SSL communication with this server. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_sni": schema.StringAttribute{
				MarkdownDescription: "The host name sent in the SNI TLS extension to the server. When present it will be preferred over the SNI expression. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_sni_expr": schema.StringAttribute{
				MarkdownDescription: "A HAProxy SNI expression to specify the data that will be sent in the SNI TLS extension to the server, e.g. req.hdr(host). When a SNI name is present it will be used instead and this option will be ignored. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_verify": schema.BoolAttribute{
				MarkdownDescription: "If disabled, server certificate is not verified. Otherwise the certificate provided by the server is verified using CAs and optional CRLs. Please note that this setting can be overriden by the global configuration.NOTE: It is critically important to verify server certificates when using SSL to connect to servers, otherwise the communication is prone to trivial man-in-the-middle attacks rendering SSL totally useless. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ssl_ca": schema.SetAttribute{
				MarkdownDescription: "The selected CAs will be used to verify the server certificate. When nothing is selected, the system-wide CA certificate bundle will be used. To import additional CAs, go to Authority Manager. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"ssl_crl": schema.StringAttribute{
				MarkdownDescription: "This certificate revocation list will be used to verify server's certificate. To import additional CRLs, go to Certificate Manager. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_client_certificate": schema.StringAttribute{
				MarkdownDescription: "This certificate will be sent if the server send a client certificate request.To import additional certificates, go to Certificate Manager. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"max_connections": schema.Int64Attribute{
				MarkdownDescription: "Specifies the maximal number of concurrent connections that will be sent to this server. The default value is \"0\" which means unlimited. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"weight": schema.Int64Attribute{
				MarkdownDescription: "Adjust the server's weight relative to other servers. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"check_interval": schema.StringAttribute{
				MarkdownDescription: "Sets the interval (in milliseconds) for running health checks on this server. This setting takes precedence over default values in health monitors. It can still be overwritten from the backend pool. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"check_down_interval": schema.StringAttribute{
				MarkdownDescription: "Sets the interval (in milliseconds) for running health checks on the server when the server state is DOWN. If it is not set HAProxy uses the check interval. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"source": schema.StringAttribute{
				MarkdownDescription: "Sets the source address which will be used when connecting to the server. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"advanced": schema.StringAttribute{
				MarkdownDescription: "A list of parameters that will be appended to the server line in every backend where this server will be used.Example: send-proxy. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"unix_socket": schema.StringAttribute{
				MarkdownDescription: "Select the frontend that provides the UNIX socket. This UNIX socket will be used as the server's address, making it possible to send connections to this frontend. Only frontends that provide the unix@ pattern as listen address can be selected. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the server.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func serverDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "HAProxy servers (**Services ▸ HAProxy ▸ Settings ▸ Real Servers**, plugin `os-haproxy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the server.",
				Required:            true,
			},
			"enabled":                dschema.BoolAttribute{MarkdownDescription: "Enable this server.", Computed: true},
			"name":                   dschema.StringAttribute{MarkdownDescription: "Name to identify a static server. When creating a server template, then this prefix is used for the server names to be built.", Computed: true},
			"description":            dschema.StringAttribute{MarkdownDescription: "Description for this server.", Computed: true},
			"address":                dschema.StringAttribute{MarkdownDescription: "Provide either the FQDN or the IP address of this server.", Computed: true},
			"port":                   dschema.Int64Attribute{MarkdownDescription: "Provide the TCP or UDP communication port for this server, i.e. 80 or 443. If set, all connections will be sent to this port. If unset, the same port the client connected to will be used.", Computed: true},
			"checkport":              dschema.Int64Attribute{MarkdownDescription: "Provide the TCP communication port to use during check, i.e. 80 or 443.", Computed: true},
			"mode":                   dschema.StringAttribute{MarkdownDescription: "Sets the operation mode to use for this server. One of `active`, `backup`, `disabled`.", Computed: true},
			"multiplexer_protocol":   dschema.StringAttribute{MarkdownDescription: "Forces the multiplexer's protocol to use for the outgoing connections to this server. It must be compatible with the mode of the backend (TCP or HTTP). It must also be usable on the backend side. Idea behind this option is to bypass the selection of the best multiplexer's protocol for all connections established to this server. Use only when strictly necessary. One of `unspecified`, `fcgi`, `h2`, `h1`.", Computed: true},
			"type":                   dschema.StringAttribute{MarkdownDescription: "Either configure a static server or a template to initialize multiple servers with shared parameters. One of `static`, `template`, `unix`.", Computed: true},
			"service_name":           dschema.StringAttribute{MarkdownDescription: "Provide either the FQDN for all the servers this template initializes or a service name to discover the available services via DNS SRV records.", Computed: true},
			"number":                 dschema.StringAttribute{MarkdownDescription: "The number of servers this template initializes, i.e. 5 or 1-5.", Computed: true},
			"linked_resolver":        dschema.StringAttribute{MarkdownDescription: "Specify the resolver that the server template should look at to discover available services via DNS.", Computed: true},
			"resolver_opts":          dschema.SetAttribute{MarkdownDescription: "Add resolver options. Use TAB key to complete typing. One of `allow-dup-ip`, `ignore-weight`, `prevent-dup-ip`.", Computed: true, ElementType: types.StringType},
			"resolve_prefer":         dschema.StringAttribute{MarkdownDescription: "When DNS resolution is enabled for a server and multiple IP addresses from different families are returned, HAProxy will prefer using an IP address from the selected family. One of `ipv4`, `ipv6`.", Computed: true},
			"ssl":                    dschema.BoolAttribute{MarkdownDescription: "Enable or disable SSL communication with this server.", Computed: true},
			"ssl_sni":                dschema.StringAttribute{MarkdownDescription: "The host name sent in the SNI TLS extension to the server. When present it will be preferred over the SNI expression.", Computed: true},
			"ssl_sni_expr":           dschema.StringAttribute{MarkdownDescription: "A HAProxy SNI expression to specify the data that will be sent in the SNI TLS extension to the server, e.g. req.hdr(host). When a SNI name is present it will be used instead and this option will be ignored.", Computed: true},
			"ssl_verify":             dschema.BoolAttribute{MarkdownDescription: "If disabled, server certificate is not verified. Otherwise the certificate provided by the server is verified using CAs and optional CRLs. Please note that this setting can be overriden by the global configuration.NOTE: It is critically important to verify server certificates when using SSL to connect to servers, otherwise the communication is prone to trivial man-in-the-middle attacks rendering SSL totally useless.", Computed: true},
			"ssl_ca":                 dschema.SetAttribute{MarkdownDescription: "The selected CAs will be used to verify the server certificate. When nothing is selected, the system-wide CA certificate bundle will be used. To import additional CAs, go to Authority Manager.", Computed: true, ElementType: types.StringType},
			"ssl_crl":                dschema.StringAttribute{MarkdownDescription: "This certificate revocation list will be used to verify server's certificate. To import additional CRLs, go to Certificate Manager.", Computed: true},
			"ssl_client_certificate": dschema.StringAttribute{MarkdownDescription: "This certificate will be sent if the server send a client certificate request.To import additional certificates, go to Certificate Manager.", Computed: true},
			"max_connections":        dschema.Int64Attribute{MarkdownDescription: "Specifies the maximal number of concurrent connections that will be sent to this server. The default value is \"0\" which means unlimited.", Computed: true},
			"weight":                 dschema.Int64Attribute{MarkdownDescription: "Adjust the server's weight relative to other servers.", Computed: true},
			"check_interval":         dschema.StringAttribute{MarkdownDescription: "Sets the interval (in milliseconds) for running health checks on this server. This setting takes precedence over default values in health monitors. It can still be overwritten from the backend pool.", Computed: true},
			"check_down_interval":    dschema.StringAttribute{MarkdownDescription: "Sets the interval (in milliseconds) for running health checks on the server when the server state is DOWN. If it is not set HAProxy uses the check interval.", Computed: true},
			"source":                 dschema.StringAttribute{MarkdownDescription: "Sets the source address which will be used when connecting to the server.", Computed: true},
			"advanced":               dschema.StringAttribute{MarkdownDescription: "A list of parameters that will be appended to the server line in every backend where this server will be used.Example: send-proxy.", Computed: true},
			"unix_socket":            dschema.StringAttribute{MarkdownDescription: "Select the frontend that provides the UNIX socket. This UNIX socket will be used as the server's address, making it possible to send connections to this frontend. Only frontends that provide the unix@ pattern as listen address can be selected.", Computed: true},
		},
	}
}

func convertServerSchemaToStruct(d *serverResourceModel) (*haproxy.Server, error) {
	return &haproxy.Server{
		Enabled:              tools.BoolToString(d.Enabled.ValueBool()),
		Name:                 d.Name.ValueString(),
		Description:          d.Description.ValueString(),
		Address:              d.Address.ValueString(),
		Port:                 tools.Int64ToStringNegative(d.Port.ValueInt64()),
		Checkport:            tools.Int64ToStringNegative(d.Checkport.ValueInt64()),
		Mode:                 api.SelectedMap(d.Mode.ValueString()),
		MultiplexerProtocol:  api.SelectedMap(d.MultiplexerProtocol.ValueString()),
		Type:                 api.SelectedMap(d.Type.ValueString()),
		ServiceName:          d.ServiceName.ValueString(),
		Number:               d.Number.ValueString(),
		LinkedResolver:       api.SelectedMap(d.LinkedResolver.ValueString()),
		ResolverOpts:         api.SelectedMapList(setToSlice(d.ResolverOpts)),
		ResolvePrefer:        api.SelectedMap(d.ResolvePrefer.ValueString()),
		Ssl:                  tools.BoolToString(d.Ssl.ValueBool()),
		SslSni:               d.SslSni.ValueString(),
		SslSniExpr:           d.SslSniExpr.ValueString(),
		SslVerify:            tools.BoolToString(d.SslVerify.ValueBool()),
		SslCa:                api.SelectedMapList(setToSlice(d.SslCa)),
		SslCrl:               api.SelectedMap(d.SslCrl.ValueString()),
		SslClientCertificate: api.SelectedMap(d.SslClientCertificate.ValueString()),
		MaxConnections:       tools.Int64ToStringNegative(d.MaxConnections.ValueInt64()),
		Weight:               tools.Int64ToStringNegative(d.Weight.ValueInt64()),
		CheckInterval:        d.CheckInterval.ValueString(),
		CheckDownInterval:    d.CheckDownInterval.ValueString(),
		Source:               d.Source.ValueString(),
		Advanced:             d.Advanced.ValueString(),
		UnixSocket:           api.SelectedMap(d.UnixSocket.ValueString()),
	}, nil
}

func convertServerStructToSchema(d *haproxy.Server) (*serverResourceModel, error) {
	return &serverResourceModel{
		Enabled:              types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:                 types.StringValue(d.Name),
		Description:          types.StringValue(d.Description),
		Address:              types.StringValue(d.Address),
		Port:                 types.Int64Value(tools.StringToInt64(d.Port)),
		Checkport:            types.Int64Value(tools.StringToInt64(d.Checkport)),
		Mode:                 types.StringValue(d.Mode.String()),
		MultiplexerProtocol:  types.StringValue(d.MultiplexerProtocol.String()),
		Type:                 types.StringValue(d.Type.String()),
		ServiceName:          types.StringValue(d.ServiceName),
		Number:               types.StringValue(d.Number),
		LinkedResolver:       types.StringValue(d.LinkedResolver.String()),
		ResolverOpts:         tools.StringSliceToSet(d.ResolverOpts),
		ResolvePrefer:        types.StringValue(d.ResolvePrefer.String()),
		Ssl:                  types.BoolValue(tools.StringToBool(d.Ssl)),
		SslSni:               types.StringValue(d.SslSni),
		SslSniExpr:           types.StringValue(d.SslSniExpr),
		SslVerify:            types.BoolValue(tools.StringToBool(d.SslVerify)),
		SslCa:                tools.StringSliceToSet(d.SslCa),
		SslCrl:               types.StringValue(d.SslCrl.String()),
		SslClientCertificate: types.StringValue(d.SslClientCertificate.String()),
		MaxConnections:       types.Int64Value(tools.StringToInt64(d.MaxConnections)),
		Weight:               types.Int64Value(tools.StringToInt64(d.Weight)),
		CheckInterval:        types.StringValue(d.CheckInterval),
		CheckDownInterval:    types.StringValue(d.CheckDownInterval),
		Source:               types.StringValue(d.Source),
		Advanced:             types.StringValue(d.Advanced),
		UnixSocket:           types.StringValue(d.UnixSocket.String()),
	}, nil
}
