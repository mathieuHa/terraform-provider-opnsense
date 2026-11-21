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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type actionResourceModel struct {
	Enabled                        types.Bool   `tfsdk:"enabled"`
	Name                           types.String `tfsdk:"name"`
	Description                    types.String `tfsdk:"description"`
	TestType                       types.String `tfsdk:"test_type"`
	LinkedAcls                     types.Set    `tfsdk:"linked_acls"`
	Operator                       types.String `tfsdk:"operator"`
	Type                           types.String `tfsdk:"type"`
	UseBackend                     types.String `tfsdk:"use_backend"`
	UseServer                      types.String `tfsdk:"use_server"`
	FcgiPassHeader                 types.String `tfsdk:"fcgi_pass_header"`
	FcgiSetParam                   types.String `tfsdk:"fcgi_set_param"`
	MonitorFailUri                 types.String `tfsdk:"monitor_fail_uri"`
	Custom                         types.String `tfsdk:"custom"`
	HttpAfterResponseAction        types.String `tfsdk:"http_after_response_action"`
	HttpAfterResponseOption        types.String `tfsdk:"http_after_response_option"`
	HttpRequestAction              types.String `tfsdk:"http_request_action"`
	HttpRequestOption              types.String `tfsdk:"http_request_option"`
	HttpResponseAction             types.String `tfsdk:"http_response_action"`
	HttpResponseOption             types.String `tfsdk:"http_response_option"`
	TcpRequestAction               types.String `tfsdk:"tcp_request_action"`
	TcpRequestOption               types.String `tfsdk:"tcp_request_option"`
	TcpResponseAction              types.String `tfsdk:"tcp_response_action"`
	TcpResponseOption              types.String `tfsdk:"tcp_response_option"`
	HttpRequestAuth                types.String `tfsdk:"http_request_auth"`
	HttpRequestDenyStatus          types.Int64  `tfsdk:"http_request_deny_status"`
	HttpRequestRedirect            types.String `tfsdk:"http_request_redirect"`
	HttpRequestLua                 types.String `tfsdk:"http_request_lua"`
	HttpRequestUseService          types.String `tfsdk:"http_request_use_service"`
	HttpRequestAddHeaderName       types.String `tfsdk:"http_request_add_header_name"`
	HttpRequestAddHeaderContent    types.String `tfsdk:"http_request_add_header_content"`
	HttpRequestSetHeaderName       types.String `tfsdk:"http_request_set_header_name"`
	HttpRequestSetHeaderContent    types.String `tfsdk:"http_request_set_header_content"`
	HttpRequestDelHeaderName       types.String `tfsdk:"http_request_del_header_name"`
	HttpRequestReplaceHeaderName   types.String `tfsdk:"http_request_replace_header_name"`
	HttpRequestReplaceHeaderRegex  types.String `tfsdk:"http_request_replace_header_regex"`
	HttpRequestReplaceValueName    types.String `tfsdk:"http_request_replace_value_name"`
	HttpRequestReplaceValueRegex   types.String `tfsdk:"http_request_replace_value_regex"`
	HttpRequestSetPath             types.String `tfsdk:"http_request_set_path"`
	HttpRequestSetVarScope         types.String `tfsdk:"http_request_set_var_scope"`
	HttpRequestSetVarName          types.String `tfsdk:"http_request_set_var_name"`
	HttpRequestSetVarExpr          types.String `tfsdk:"http_request_set_var_expr"`
	HttpResponseLua                types.String `tfsdk:"http_response_lua"`
	HttpResponseAddHeaderName      types.String `tfsdk:"http_response_add_header_name"`
	HttpResponseAddHeaderContent   types.String `tfsdk:"http_response_add_header_content"`
	HttpResponseSetHeaderName      types.String `tfsdk:"http_response_set_header_name"`
	HttpResponseSetHeaderContent   types.String `tfsdk:"http_response_set_header_content"`
	HttpResponseDelHeaderName      types.String `tfsdk:"http_response_del_header_name"`
	HttpResponseReplaceHeaderName  types.String `tfsdk:"http_response_replace_header_name"`
	HttpResponseReplaceHeaderRegex types.String `tfsdk:"http_response_replace_header_regex"`
	HttpResponseReplaceValueName   types.String `tfsdk:"http_response_replace_value_name"`
	HttpResponseReplaceValueRegex  types.String `tfsdk:"http_response_replace_value_regex"`
	HttpResponseSetStatusCode      types.Int64  `tfsdk:"http_response_set_status_code"`
	HttpResponseSetStatusReason    types.String `tfsdk:"http_response_set_status_reason"`
	HttpResponseSetVarScope        types.String `tfsdk:"http_response_set_var_scope"`
	HttpResponseSetVarName         types.String `tfsdk:"http_response_set_var_name"`
	HttpResponseSetVarExpr         types.String `tfsdk:"http_response_set_var_expr"`
	TcpRequestContentLua           types.String `tfsdk:"tcp_request_content_lua"`
	TcpRequestContentUseService    types.String `tfsdk:"tcp_request_content_use_service"`
	TcpRequestInspectDelay         types.String `tfsdk:"tcp_request_inspect_delay"`
	TcpResponseContentLua          types.String `tfsdk:"tcp_response_content_lua"`
	TcpResponseInspectDelay        types.String `tfsdk:"tcp_response_inspect_delay"`
	MapDataUseBackendFile          types.String `tfsdk:"map_data_use_backend_file"`
	MapDataUseBackendDefault       types.String `tfsdk:"map_data_use_backend_default"`
	MapDataUseBackendInput         types.String `tfsdk:"map_data_use_backend_input"`
	MapUseBackendFile              types.String `tfsdk:"map_use_backend_file"`
	MapUseBackendDefault           types.String `tfsdk:"map_use_backend_default"`
	CompressionAlgoRes             types.String `tfsdk:"compression_algo_res"`
	CompressionAlgoReq             types.String `tfsdk:"compression_algo_req"`
	CompressionMimeRes             types.Set    `tfsdk:"compression_mime_res"`
	CompressionMimeReq             types.Set    `tfsdk:"compression_mime_req"`
	CompressionOffloading          types.Bool   `tfsdk:"compression_offloading"`
	CompressionMinsizeRes          types.Int64  `tfsdk:"compression_minsize_res"`
	CompressionMinsizeReq          types.Int64  `tfsdk:"compression_minsize_req"`
	CompressionDirection           types.String `tfsdk:"compression_direction"`
	GpcNumber                      types.Int64  `tfsdk:"gpc_number"`
	GptNumber                      types.Int64  `tfsdk:"gpt_number"`
	ScNumber                       types.Int64  `tfsdk:"sc_number"`
	Mapfile                        types.String `tfsdk:"mapfile"`
	MapDefault                     types.String `tfsdk:"map_default"`
	SampleFetch                    types.String `tfsdk:"sample_fetch"`

	Id types.String `tfsdk:"id"`
}

func actionResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "HAProxy rules (**Services ▸ HAProxy ▸ Settings ▸ Rules & Checks ▸ Rules**, plugin `os-haproxy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this rule. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify this rule.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this rule. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"test_type": schema.StringAttribute{
				MarkdownDescription: "Apply this rule when conditions are met (IF) or unmet (UNLESS). One of `if`, `unless`. Defaults to `if`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("if"),
				Validators:          []validator.String{stringvalidator.OneOf("if", "unless")},
			},
			"linked_acls": schema.SetAttribute{
				MarkdownDescription: "Select one or more conditions to be used for this rule. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"operator": schema.StringAttribute{
				MarkdownDescription: "Combine multiple conditions using AND (all must match) or OR (any must match). One of `and`, `or`. Defaults to `and`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("and"),
				Validators:          []validator.String{stringvalidator.OneOf("", "and", "or")},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Choose a HAProxy action that should be used if the condition evaluates to true. One of `compression`, `fcgi_pass_header`, `fcgi_set_param`, `http-after-response`, `http-request`, `http-response`, `map_data_use_backend`, `map_use_backend`, `monitor_fail`, `tcp-request`, `tcp-response`, `use_backend`, `use_server`, `custom`.",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf("compression", "fcgi_pass_header", "fcgi_set_param", "http-after-response", "http-request", "http-response", "map_data_use_backend", "map_use_backend", "monitor_fail", "tcp-request", "tcp-response", "use_backend", "use_server", "custom")},
			},
			"use_backend": schema.StringAttribute{
				MarkdownDescription: "HAProxy will use this backend pool if the condition evaluates to true. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"use_server": schema.StringAttribute{
				MarkdownDescription: "HAProxy will use this server instead of other servers that are specified in the Backend Pool. The server must exist in the context where this rule is applied. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"fcgi_pass_header": schema.StringAttribute{
				MarkdownDescription: "Specify the name of a request header which will be passed to the FastCGI application. Most request headers are already available to the FastCGI application, prefixed with \"HTTP_\". Thus, this directive is only required to pass headers that are purposefully omitted. Currently, the headers \"Authorization\", \"Proxy-Authorization\" and hop-by-hop headers are omitted. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"fcgi_set_param": schema.StringAttribute{
				MarkdownDescription: "Set a FastCGI parameter that should be passed to the application. Its value must follow HAProxy's Custom Log format rules. With this directive, it is possible to overwrite the value of default FastCGI parameters. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"monitor_fail_uri": schema.StringAttribute{
				MarkdownDescription: "Specifies a URI which will be intercepted to return HAProxy's health status instead of forwarding the request. When a HTTP request is received, HAProxy will return either \"HTTP/1.0 200 OK\" or \"HTTP/1.0 503 Service unavailable\", depending on the fined failure conditions. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"custom": schema.StringAttribute{
				MarkdownDescription: "Specify a HAProxy rule/ACL that is currently not supported by the GUI. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_after_response_action": schema.StringAttribute{
				MarkdownDescription: "A detailed explanation of each action and examples can be found in HAProxy's documentation. One of `add-header`, `allow`, `capture`, `del-header`, `del-map`, `do-log`, `replace-header`, `replace-value`, `sc-add-gpc`, `sc-inc-gpc`, `sc-inc-gpc0`, `sc-inc-gpc1`, `sc-set-gpt`, `sc-set-gpt0`, `set-header`, `set-log-level`, `set-map`, `set-status`, `set-var`, `set-var-fmt`, `strict-mode`, `unset-var`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "add-header", "allow", "capture", "del-header", "del-map", "do-log", "replace-header", "replace-value", "sc-add-gpc", "sc-inc-gpc", "sc-inc-gpc0", "sc-inc-gpc1", "sc-set-gpt", "sc-set-gpt0", "set-header", "set-log-level", "set-map", "set-status", "set-var", "set-var-fmt", "strict-mode", "unset-var")},
			},
			"http_after_response_option": schema.StringAttribute{
				MarkdownDescription: "Some actions require additional parameters and specific syntax. Examples and explanations can be found in the HAProxy's documentation. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_request_action": schema.StringAttribute{
				MarkdownDescription: "A detailed explanation of each action and examples can be found in HAProxy's documentation. One of `add-acl`, `add-header`, `allow`, `auth`, `cache-use`, `capture`, `del-acl`, `del-header`, `del-map`, `deny`, `disable-l7-retry`, `do-log`, `do-resolve`, `early-hint`, `lua`, `normalize-uri`, `redirect`, `reject`, `replace-header`, `replace-path`, `replace-pathq`, `replace-uri`, `replace-value`, `return`, `sc-add-gpc`, `sc-inc-gpc`, `sc-inc-gpc0`, `sc-inc-gpc1`, `sc-set-gpt`, `sc-set-gpt0`, `send-spoe-group`, `set-dst`, `set-dst-port`, `set-fc-mark`, `set-fc-tos`, `set-header`, `set-log-level`, `set-map`, `set-method`, `set-nice`, `set-path`, `set-pathq`, `set-priority-class`, `set-priority-offset`, `set-query`, `set-src`, `set-src-port`, `set-timeout`, `set-uri`, `set-var`, `set-var-fmt`, `silent-drop`, `strict-mode`, `tarpit`, `track-sc0`, `track-sc1`, `track-sc2`, `unset-var`, `use-service`, `wait-for-body`, `wait-for-handshake`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "add-acl", "add-header", "allow", "auth", "cache-use", "capture", "del-acl", "del-header", "del-map", "deny", "disable-l7-retry", "do-log", "do-resolve", "early-hint", "lua", "normalize-uri", "redirect", "reject", "replace-header", "replace-path", "replace-pathq", "replace-uri", "replace-value", "return", "sc-add-gpc", "sc-inc-gpc", "sc-inc-gpc0", "sc-inc-gpc1", "sc-set-gpt", "sc-set-gpt0", "send-spoe-group", "set-dst", "set-dst-port", "set-fc-mark", "set-fc-tos", "set-header", "set-log-level", "set-map", "set-method", "set-nice", "set-path", "set-pathq", "set-priority-class", "set-priority-offset", "set-query", "set-src", "set-src-port", "set-timeout", "set-uri", "set-var", "set-var-fmt", "silent-drop", "strict-mode", "tarpit", "track-sc0", "track-sc1", "track-sc2", "unset-var", "use-service", "wait-for-body", "wait-for-handshake")},
			},
			"http_request_option": schema.StringAttribute{
				MarkdownDescription: "Some actions require additional parameters and specific syntax. Examples and explanations can be found in the HAProxy's documentation. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_response_action": schema.StringAttribute{
				MarkdownDescription: "A detailed explanation of each action and examples can be found in HAProxy's documentation. One of `add-acl`, `add-header`, `allow`, `cache-store`, `capture`, `del-acl`, `del-header`, `del-map`, `deny`, `do-log`, `lua`, `redirect`, `replace-header`, `replace-value`, `return`, `sc-add-gpc`, `sc-inc-gpc`, `sc-inc-gpc0`, `sc-inc-gpc1`, `sc-set-gpt`, `sc-set-gpt0`, `send-spoe-group`, `set-fc-mark`, `set-fc-tos`, `set-header`, `set-log-level`, `set-map`, `set-nice`, `set-status`, `set-timeout`, `set-var`, `set-var-fmt`, `silent-drop`, `strict-mode`, `track-sc0`, `track-sc1`, `track-sc2`, `unset-var`, `wait-for-body`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "add-acl", "add-header", "allow", "cache-store", "capture", "del-acl", "del-header", "del-map", "deny", "do-log", "lua", "redirect", "replace-header", "replace-value", "return", "sc-add-gpc", "sc-inc-gpc", "sc-inc-gpc0", "sc-inc-gpc1", "sc-set-gpt", "sc-set-gpt0", "send-spoe-group", "set-fc-mark", "set-fc-tos", "set-header", "set-log-level", "set-map", "set-nice", "set-status", "set-timeout", "set-var", "set-var-fmt", "silent-drop", "strict-mode", "track-sc0", "track-sc1", "track-sc2", "unset-var", "wait-for-body")},
			},
			"http_response_option": schema.StringAttribute{
				MarkdownDescription: "Some actions require additional parameters and specific syntax. Examples and explanations can be found in the HAProxy's documentation. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tcp_request_action": schema.StringAttribute{
				MarkdownDescription: "A detailed explanation of each action and examples can be found in HAProxy's documentation. One of `connection_accept`, `connection_expect-netscaler-cip`, `connection_expect-proxy`, `connection_fc-silent-drop`, `connection_reject`, `connection_sc-add-gpc`, `connection_sc-inc-gpc`, `connection_sc-inc-gpc0`, `connection_sc-inc-gpc1`, `connection_sc-set-gpt`, `connection_sc-set-gpt0`, `connection_send-spoe-group`, `connection_set-dst`, `connection_set-dst-port`, `connection_set-fc-mark`, `connection_set-fc-tos`, `connection_set-log-level`, `connection_set-src`, `connection_set-src-port`, `connection_set-var`, `connection_set-var-fmt`, `connection_silent-drop`, `connection_track-sc0`, `connection_track-sc1`, `connection_track-sc2`, `connection_unset-var`, `content_accept`, `content_capture`, `content_do-resolve`, `content_lua`, `content_reject`, `content_sc-add-gpc`, `content_sc-inc-gpc`, `content_sc-inc-gpc0`, `content_sc-inc-gpc1`, `content_sc-set-gpt`, `content_sc-set-gpt0`, `content_send-spoe-group`, `content_set-dst`, `content_set-dst-port`, `content_set-fc-mark`, `content_set-fc-tos`, `content_set-log-level`, `content_set-nice`, `content_set-priority-class`, `content_set-priority-offset`, `content_set-src`, `content_set-src-port`, `content_set-var`, `content_set-var-fmt`, `content_silent-drop`, `content_switch-mode`, `content_track-sc0`, `content_track-sc1`, `content_track-sc2`, `content_unset-var`, `content_use-service`, `inspect-delay`, `session_accept`, `session_attach-srv`, `session_reject`, `session_sc-add-gpc`, `session_sc-inc-gpc`, `session_sc-inc-gpc0`, `session_sc-inc-gpc1`, `session_sc-set-gpt`, `session_sc-set-gpt0`, `session_send-spoe-group`, `session_set-dst`, `session_set-dst-port`, `session_set-fc-mark`, `session_set-fc-tos`, `session_set-log-level`, `session_set-src`, `session_set-src-port`, `session_set-var`, `session_set-var-fmt`, `session_silent-drop`, `session_track-sc0`, `session_track-sc1`, `session_track-sc2`, `session_unset-var`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "connection_accept", "connection_expect-netscaler-cip", "connection_expect-proxy", "connection_fc-silent-drop", "connection_reject", "connection_sc-add-gpc", "connection_sc-inc-gpc", "connection_sc-inc-gpc0", "connection_sc-inc-gpc1", "connection_sc-set-gpt", "connection_sc-set-gpt0", "connection_send-spoe-group", "connection_set-dst", "connection_set-dst-port", "connection_set-fc-mark", "connection_set-fc-tos", "connection_set-log-level", "connection_set-src", "connection_set-src-port", "connection_set-var", "connection_set-var-fmt", "connection_silent-drop", "connection_track-sc0", "connection_track-sc1", "connection_track-sc2", "connection_unset-var", "content_accept", "content_capture", "content_do-resolve", "content_lua", "content_reject", "content_sc-add-gpc", "content_sc-inc-gpc", "content_sc-inc-gpc0", "content_sc-inc-gpc1", "content_sc-set-gpt", "content_sc-set-gpt0", "content_send-spoe-group", "content_set-dst", "content_set-dst-port", "content_set-fc-mark", "content_set-fc-tos", "content_set-log-level", "content_set-nice", "content_set-priority-class", "content_set-priority-offset", "content_set-src", "content_set-src-port", "content_set-var", "content_set-var-fmt", "content_silent-drop", "content_switch-mode", "content_track-sc0", "content_track-sc1", "content_track-sc2", "content_unset-var", "content_use-service", "inspect-delay", "session_accept", "session_attach-srv", "session_reject", "session_sc-add-gpc", "session_sc-inc-gpc", "session_sc-inc-gpc0", "session_sc-inc-gpc1", "session_sc-set-gpt", "session_sc-set-gpt0", "session_send-spoe-group", "session_set-dst", "session_set-dst-port", "session_set-fc-mark", "session_set-fc-tos", "session_set-log-level", "session_set-src", "session_set-src-port", "session_set-var", "session_set-var-fmt", "session_silent-drop", "session_track-sc0", "session_track-sc1", "session_track-sc2", "session_unset-var")},
			},
			"tcp_request_option": schema.StringAttribute{
				MarkdownDescription: "Some actions require additional parameters and specific syntax. Examples and explanations can be found in the HAProxy's documentation. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tcp_response_action": schema.StringAttribute{
				MarkdownDescription: "A detailed explanation of each action and examples can be found in HAProxy's documentation. One of `content_accept`, `content_close`, `content_lua`, `content_reject`, `content_sc-add-gpc`, `content_sc-inc-gpc`, `content_sc-inc-gpc0`, `content_sc-inc-gpc1`, `content_sc-set-gpt`, `content_sc-set-gpt0`, `content_send-spoe-group`, `content_set-fc-mark`, `content_set-fc-tos`, `content_set-log-level`, `content_set-nice`, `content_set-var`, `content_set-var-fmt`, `content_silent-drop`, `content_unset-var`, `inspect-delay`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "content_accept", "content_close", "content_lua", "content_reject", "content_sc-add-gpc", "content_sc-inc-gpc", "content_sc-inc-gpc0", "content_sc-inc-gpc1", "content_sc-set-gpt", "content_sc-set-gpt0", "content_send-spoe-group", "content_set-fc-mark", "content_set-fc-tos", "content_set-log-level", "content_set-nice", "content_set-var", "content_set-var-fmt", "content_silent-drop", "content_unset-var", "inspect-delay")},
			},
			"tcp_response_option": schema.StringAttribute{
				MarkdownDescription: "Some actions require additional parameters and specific syntax. Examples and explanations can be found in the HAProxy's documentation. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_request_auth": schema.StringAttribute{
				MarkdownDescription: "http_request_auth. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_request_deny_status": schema.Int64Attribute{
				MarkdownDescription: "http_request_deny_status. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"http_request_redirect": schema.StringAttribute{
				MarkdownDescription: "http_request_redirect. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_request_lua": schema.StringAttribute{
				MarkdownDescription: "http_request_lua. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_request_use_service": schema.StringAttribute{
				MarkdownDescription: "http_request_use_service. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_request_add_header_name": schema.StringAttribute{
				MarkdownDescription: "http_request_add_header_name. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_request_add_header_content": schema.StringAttribute{
				MarkdownDescription: "http_request_add_header_content. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_request_set_header_name": schema.StringAttribute{
				MarkdownDescription: "http_request_set_header_name. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_request_set_header_content": schema.StringAttribute{
				MarkdownDescription: "http_request_set_header_content. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_request_del_header_name": schema.StringAttribute{
				MarkdownDescription: "http_request_del_header_name. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_request_replace_header_name": schema.StringAttribute{
				MarkdownDescription: "http_request_replace_header_name. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_request_replace_header_regex": schema.StringAttribute{
				MarkdownDescription: "http_request_replace_header_regex. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_request_replace_value_name": schema.StringAttribute{
				MarkdownDescription: "http_request_replace_value_name. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_request_replace_value_regex": schema.StringAttribute{
				MarkdownDescription: "http_request_replace_value_regex. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_request_set_path": schema.StringAttribute{
				MarkdownDescription: "http_request_set_path. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_request_set_var_scope": schema.StringAttribute{
				MarkdownDescription: "http_request_set_var_scope. One of `proc`, `sess`, `txn`, `req`, `res`. Defaults to `txn`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("txn"),
				Validators:          []validator.String{stringvalidator.OneOf("", "proc", "sess", "txn", "req", "res")},
			},
			"http_request_set_var_name": schema.StringAttribute{
				MarkdownDescription: "http_request_set_var_name. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_request_set_var_expr": schema.StringAttribute{
				MarkdownDescription: "http_request_set_var_expr. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_response_lua": schema.StringAttribute{
				MarkdownDescription: "http_response_lua. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_response_add_header_name": schema.StringAttribute{
				MarkdownDescription: "http_response_add_header_name. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_response_add_header_content": schema.StringAttribute{
				MarkdownDescription: "http_response_add_header_content. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_response_set_header_name": schema.StringAttribute{
				MarkdownDescription: "http_response_set_header_name. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_response_set_header_content": schema.StringAttribute{
				MarkdownDescription: "http_response_set_header_content. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_response_del_header_name": schema.StringAttribute{
				MarkdownDescription: "http_response_del_header_name. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_response_replace_header_name": schema.StringAttribute{
				MarkdownDescription: "http_response_replace_header_name. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_response_replace_header_regex": schema.StringAttribute{
				MarkdownDescription: "http_response_replace_header_regex. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_response_replace_value_name": schema.StringAttribute{
				MarkdownDescription: "http_response_replace_value_name. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_response_replace_value_regex": schema.StringAttribute{
				MarkdownDescription: "http_response_replace_value_regex. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_response_set_status_code": schema.Int64Attribute{
				MarkdownDescription: "http_response_set_status_code. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"http_response_set_status_reason": schema.StringAttribute{
				MarkdownDescription: "http_response_set_status_reason. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_response_set_var_scope": schema.StringAttribute{
				MarkdownDescription: "http_response_set_var_scope. One of `proc`, `sess`, `txn`, `req`, `res`. Defaults to `txn`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("txn"),
				Validators:          []validator.String{stringvalidator.OneOf("", "proc", "sess", "txn", "req", "res")},
			},
			"http_response_set_var_name": schema.StringAttribute{
				MarkdownDescription: "http_response_set_var_name. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_response_set_var_expr": schema.StringAttribute{
				MarkdownDescription: "http_response_set_var_expr. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tcp_request_content_lua": schema.StringAttribute{
				MarkdownDescription: "tcp_request_content_lua. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tcp_request_content_use_service": schema.StringAttribute{
				MarkdownDescription: "tcp_request_content_use_service. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tcp_request_inspect_delay": schema.StringAttribute{
				MarkdownDescription: "tcp_request_inspect_delay. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tcp_response_content_lua": schema.StringAttribute{
				MarkdownDescription: "tcp_response_content_lua. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tcp_response_inspect_delay": schema.StringAttribute{
				MarkdownDescription: "tcp_response_inspect_delay. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"map_data_use_backend_file": schema.StringAttribute{
				MarkdownDescription: "HAProxy will extract the Host header from the HTTP request and search the map file for a match. If a match is found, the backend pool from the map file will be used. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"map_data_use_backend_default": schema.StringAttribute{
				MarkdownDescription: "HAProxy will use this backend pool if no match is found in the map file. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"map_data_use_backend_input": schema.StringAttribute{
				MarkdownDescription: "The sample expression (e.g., path, req.ssl_sni, hdr(Host), str(active)) whose fetched value is used as the lookup key for the map file. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"map_use_backend_file": schema.StringAttribute{
				MarkdownDescription: "HAProxy will extract the Host header from the HTTP request and search the map file for a match. If a match is found, the backend pool from the map file will be used. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"map_use_backend_default": schema.StringAttribute{
				MarkdownDescription: "HAProxy will use this backend pool if no match is found in the map file. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"compression_algo_res": schema.StringAttribute{
				MarkdownDescription: "The compression algorithm used for HTTP responses. Compression will only be used depending on the Accept-Encoding request header. One of `gzip`, `deflate`, `raw-deflate`. Defaults to `gzip`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gzip"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gzip", "deflate", "raw-deflate")},
			},
			"compression_algo_req": schema.StringAttribute{
				MarkdownDescription: "The compression algorithm used for HTTP requests. Compression will only be used depending on the Accept-Encoding request header. One of `gzip`, `deflate`, `raw-deflate`. Defaults to `gzip`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gzip"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gzip", "deflate", "raw-deflate")},
			},
			"compression_mime_res": schema.SetAttribute{
				MarkdownDescription: "HTTP responses will use compression for the specified mime types, i.e. text/css, text/html, application/javascript. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"compression_mime_req": schema.SetAttribute{
				MarkdownDescription: "HTTP requests will use compression for the specified mime types, i.e. text/css, text/html, application/javascript. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"compression_offloading": schema.BoolAttribute{
				MarkdownDescription: "When enabled, the load balancer will remove the Accept-Encoding header before forwarding a request to backend servers. This prevents the servers from performing compression, offloading that work to the load balancer. This may impact HAProxy's performance. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"compression_minsize_res": schema.Int64Attribute{
				MarkdownDescription: "Specifies a minimum file size for HTTP response compression in bytes. By avoiding unnecessary compression work on very small files, this may reduce CPU overhead. Entering 0 disables this option. Defaults to `1500`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1500),
			},
			"compression_minsize_req": schema.Int64Attribute{
				MarkdownDescription: "Specifies a minimum file size for HTTP request compression in bytes. By avoiding unnecessary compression work on very small files, this may reduce CPU overhead. Entering 0 disables this option. Defaults to `1500`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1500),
			},
			"compression_direction": schema.StringAttribute{
				MarkdownDescription: "Whether HAProxy uses compression for HTTP responses, HTTP requests, or both. Be aware that most servers will not support request body decompression. For safekeeping only enable compression for HTTP responses. One of `response`, `request`, `both`. Defaults to `response`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("response"),
				Validators:          []validator.String{stringvalidator.OneOf("", "response", "request", "both")},
			},
			"gpc_number": schema.Int64Attribute{
				MarkdownDescription: "Refers to the number of the General Purpose Counter, e.g. 0 or 1 when using gpc0 or gpc1 respectively. This value will be ignored in rules that do not support GPC. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"gpt_number": schema.Int64Attribute{
				MarkdownDescription: "Refers to the number of the General Purpose Tag, e.g. 0 or 1 when using gpt0 or gpt1 respectively. This value will be ignored in rules that do not support GPT. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_number": schema.Int64Attribute{
				MarkdownDescription: "Refers to the number of the Stick Counter, e.g. 0 or 1 when using sc0 or sc1 respectively. This value will be ignored in rules that do not support SC. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"mapfile": schema.StringAttribute{
				MarkdownDescription: "Load patterns from a map file. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"map_default": schema.StringAttribute{
				MarkdownDescription: "When a map file is specified, this is the default value. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"sample_fetch": schema.StringAttribute{
				MarkdownDescription: "Extracts a value from the request or connection context (e.g. path, header, source IP) to be used as input for converters or matching rules. May be used in combination with set-var and map files to create more complex rules, e.g.http-request set-var(req.rate_limit) path,map_beg(/path/to/mapfile,20)http-request set-var(req.request_rate) base32+src,table_http_req_rate(). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the rule.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func actionDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "HAProxy rules (**Services ▸ HAProxy ▸ Settings ▸ Rules & Checks ▸ Rules**, plugin `os-haproxy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the rule.",
				Required:            true,
			},
			"enabled":                            dschema.BoolAttribute{MarkdownDescription: "Enable this rule.", Computed: true},
			"name":                               dschema.StringAttribute{MarkdownDescription: "Name to identify this rule.", Computed: true},
			"description":                        dschema.StringAttribute{MarkdownDescription: "Description for this rule.", Computed: true},
			"test_type":                          dschema.StringAttribute{MarkdownDescription: "Apply this rule when conditions are met (IF) or unmet (UNLESS). One of `if`, `unless`.", Computed: true},
			"linked_acls":                        dschema.SetAttribute{MarkdownDescription: "Select one or more conditions to be used for this rule.", Computed: true, ElementType: types.StringType},
			"operator":                           dschema.StringAttribute{MarkdownDescription: "Combine multiple conditions using AND (all must match) or OR (any must match). One of `and`, `or`.", Computed: true},
			"type":                               dschema.StringAttribute{MarkdownDescription: "Choose a HAProxy action that should be used if the condition evaluates to true. One of `compression`, `fcgi_pass_header`, `fcgi_set_param`, `http-after-response`, `http-request`, `http-response`, `map_data_use_backend`, `map_use_backend`, `monitor_fail`, `tcp-request`, `tcp-response`, `use_backend`, `use_server`, `custom`.", Computed: true},
			"use_backend":                        dschema.StringAttribute{MarkdownDescription: "HAProxy will use this backend pool if the condition evaluates to true.", Computed: true},
			"use_server":                         dschema.StringAttribute{MarkdownDescription: "HAProxy will use this server instead of other servers that are specified in the Backend Pool. The server must exist in the context where this rule is applied.", Computed: true},
			"fcgi_pass_header":                   dschema.StringAttribute{MarkdownDescription: "Specify the name of a request header which will be passed to the FastCGI application. Most request headers are already available to the FastCGI application, prefixed with \"HTTP_\". Thus, this directive is only required to pass headers that are purposefully omitted. Currently, the headers \"Authorization\", \"Proxy-Authorization\" and hop-by-hop headers are omitted.", Computed: true},
			"fcgi_set_param":                     dschema.StringAttribute{MarkdownDescription: "Set a FastCGI parameter that should be passed to the application. Its value must follow HAProxy's Custom Log format rules. With this directive, it is possible to overwrite the value of default FastCGI parameters.", Computed: true},
			"monitor_fail_uri":                   dschema.StringAttribute{MarkdownDescription: "Specifies a URI which will be intercepted to return HAProxy's health status instead of forwarding the request. When a HTTP request is received, HAProxy will return either \"HTTP/1.0 200 OK\" or \"HTTP/1.0 503 Service unavailable\", depending on the fined failure conditions.", Computed: true},
			"custom":                             dschema.StringAttribute{MarkdownDescription: "Specify a HAProxy rule/ACL that is currently not supported by the GUI.", Computed: true},
			"http_after_response_action":         dschema.StringAttribute{MarkdownDescription: "A detailed explanation of each action and examples can be found in HAProxy's documentation. One of `add-header`, `allow`, `capture`, `del-header`, `del-map`, `do-log`, `replace-header`, `replace-value`, `sc-add-gpc`, `sc-inc-gpc`, `sc-inc-gpc0`, `sc-inc-gpc1`, `sc-set-gpt`, `sc-set-gpt0`, `set-header`, `set-log-level`, `set-map`, `set-status`, `set-var`, `set-var-fmt`, `strict-mode`, `unset-var`.", Computed: true},
			"http_after_response_option":         dschema.StringAttribute{MarkdownDescription: "Some actions require additional parameters and specific syntax. Examples and explanations can be found in the HAProxy's documentation.", Computed: true},
			"http_request_action":                dschema.StringAttribute{MarkdownDescription: "A detailed explanation of each action and examples can be found in HAProxy's documentation. One of `add-acl`, `add-header`, `allow`, `auth`, `cache-use`, `capture`, `del-acl`, `del-header`, `del-map`, `deny`, `disable-l7-retry`, `do-log`, `do-resolve`, `early-hint`, `lua`, `normalize-uri`, `redirect`, `reject`, `replace-header`, `replace-path`, `replace-pathq`, `replace-uri`, `replace-value`, `return`, `sc-add-gpc`, `sc-inc-gpc`, `sc-inc-gpc0`, `sc-inc-gpc1`, `sc-set-gpt`, `sc-set-gpt0`, `send-spoe-group`, `set-dst`, `set-dst-port`, `set-fc-mark`, `set-fc-tos`, `set-header`, `set-log-level`, `set-map`, `set-method`, `set-nice`, `set-path`, `set-pathq`, `set-priority-class`, `set-priority-offset`, `set-query`, `set-src`, `set-src-port`, `set-timeout`, `set-uri`, `set-var`, `set-var-fmt`, `silent-drop`, `strict-mode`, `tarpit`, `track-sc0`, `track-sc1`, `track-sc2`, `unset-var`, `use-service`, `wait-for-body`, `wait-for-handshake`.", Computed: true},
			"http_request_option":                dschema.StringAttribute{MarkdownDescription: "Some actions require additional parameters and specific syntax. Examples and explanations can be found in the HAProxy's documentation.", Computed: true},
			"http_response_action":               dschema.StringAttribute{MarkdownDescription: "A detailed explanation of each action and examples can be found in HAProxy's documentation. One of `add-acl`, `add-header`, `allow`, `cache-store`, `capture`, `del-acl`, `del-header`, `del-map`, `deny`, `do-log`, `lua`, `redirect`, `replace-header`, `replace-value`, `return`, `sc-add-gpc`, `sc-inc-gpc`, `sc-inc-gpc0`, `sc-inc-gpc1`, `sc-set-gpt`, `sc-set-gpt0`, `send-spoe-group`, `set-fc-mark`, `set-fc-tos`, `set-header`, `set-log-level`, `set-map`, `set-nice`, `set-status`, `set-timeout`, `set-var`, `set-var-fmt`, `silent-drop`, `strict-mode`, `track-sc0`, `track-sc1`, `track-sc2`, `unset-var`, `wait-for-body`.", Computed: true},
			"http_response_option":               dschema.StringAttribute{MarkdownDescription: "Some actions require additional parameters and specific syntax. Examples and explanations can be found in the HAProxy's documentation.", Computed: true},
			"tcp_request_action":                 dschema.StringAttribute{MarkdownDescription: "A detailed explanation of each action and examples can be found in HAProxy's documentation. One of `connection_accept`, `connection_expect-netscaler-cip`, `connection_expect-proxy`, `connection_fc-silent-drop`, `connection_reject`, `connection_sc-add-gpc`, `connection_sc-inc-gpc`, `connection_sc-inc-gpc0`, `connection_sc-inc-gpc1`, `connection_sc-set-gpt`, `connection_sc-set-gpt0`, `connection_send-spoe-group`, `connection_set-dst`, `connection_set-dst-port`, `connection_set-fc-mark`, `connection_set-fc-tos`, `connection_set-log-level`, `connection_set-src`, `connection_set-src-port`, `connection_set-var`, `connection_set-var-fmt`, `connection_silent-drop`, `connection_track-sc0`, `connection_track-sc1`, `connection_track-sc2`, `connection_unset-var`, `content_accept`, `content_capture`, `content_do-resolve`, `content_lua`, `content_reject`, `content_sc-add-gpc`, `content_sc-inc-gpc`, `content_sc-inc-gpc0`, `content_sc-inc-gpc1`, `content_sc-set-gpt`, `content_sc-set-gpt0`, `content_send-spoe-group`, `content_set-dst`, `content_set-dst-port`, `content_set-fc-mark`, `content_set-fc-tos`, `content_set-log-level`, `content_set-nice`, `content_set-priority-class`, `content_set-priority-offset`, `content_set-src`, `content_set-src-port`, `content_set-var`, `content_set-var-fmt`, `content_silent-drop`, `content_switch-mode`, `content_track-sc0`, `content_track-sc1`, `content_track-sc2`, `content_unset-var`, `content_use-service`, `inspect-delay`, `session_accept`, `session_attach-srv`, `session_reject`, `session_sc-add-gpc`, `session_sc-inc-gpc`, `session_sc-inc-gpc0`, `session_sc-inc-gpc1`, `session_sc-set-gpt`, `session_sc-set-gpt0`, `session_send-spoe-group`, `session_set-dst`, `session_set-dst-port`, `session_set-fc-mark`, `session_set-fc-tos`, `session_set-log-level`, `session_set-src`, `session_set-src-port`, `session_set-var`, `session_set-var-fmt`, `session_silent-drop`, `session_track-sc0`, `session_track-sc1`, `session_track-sc2`, `session_unset-var`.", Computed: true},
			"tcp_request_option":                 dschema.StringAttribute{MarkdownDescription: "Some actions require additional parameters and specific syntax. Examples and explanations can be found in the HAProxy's documentation.", Computed: true},
			"tcp_response_action":                dschema.StringAttribute{MarkdownDescription: "A detailed explanation of each action and examples can be found in HAProxy's documentation. One of `content_accept`, `content_close`, `content_lua`, `content_reject`, `content_sc-add-gpc`, `content_sc-inc-gpc`, `content_sc-inc-gpc0`, `content_sc-inc-gpc1`, `content_sc-set-gpt`, `content_sc-set-gpt0`, `content_send-spoe-group`, `content_set-fc-mark`, `content_set-fc-tos`, `content_set-log-level`, `content_set-nice`, `content_set-var`, `content_set-var-fmt`, `content_silent-drop`, `content_unset-var`, `inspect-delay`.", Computed: true},
			"tcp_response_option":                dschema.StringAttribute{MarkdownDescription: "Some actions require additional parameters and specific syntax. Examples and explanations can be found in the HAProxy's documentation.", Computed: true},
			"http_request_auth":                  dschema.StringAttribute{MarkdownDescription: "http_request_auth.", Computed: true},
			"http_request_deny_status":           dschema.Int64Attribute{MarkdownDescription: "http_request_deny_status.", Computed: true},
			"http_request_redirect":              dschema.StringAttribute{MarkdownDescription: "http_request_redirect.", Computed: true},
			"http_request_lua":                   dschema.StringAttribute{MarkdownDescription: "http_request_lua.", Computed: true},
			"http_request_use_service":           dschema.StringAttribute{MarkdownDescription: "http_request_use_service.", Computed: true},
			"http_request_add_header_name":       dschema.StringAttribute{MarkdownDescription: "http_request_add_header_name.", Computed: true},
			"http_request_add_header_content":    dschema.StringAttribute{MarkdownDescription: "http_request_add_header_content.", Computed: true},
			"http_request_set_header_name":       dschema.StringAttribute{MarkdownDescription: "http_request_set_header_name.", Computed: true},
			"http_request_set_header_content":    dschema.StringAttribute{MarkdownDescription: "http_request_set_header_content.", Computed: true},
			"http_request_del_header_name":       dschema.StringAttribute{MarkdownDescription: "http_request_del_header_name.", Computed: true},
			"http_request_replace_header_name":   dschema.StringAttribute{MarkdownDescription: "http_request_replace_header_name.", Computed: true},
			"http_request_replace_header_regex":  dschema.StringAttribute{MarkdownDescription: "http_request_replace_header_regex.", Computed: true},
			"http_request_replace_value_name":    dschema.StringAttribute{MarkdownDescription: "http_request_replace_value_name.", Computed: true},
			"http_request_replace_value_regex":   dschema.StringAttribute{MarkdownDescription: "http_request_replace_value_regex.", Computed: true},
			"http_request_set_path":              dschema.StringAttribute{MarkdownDescription: "http_request_set_path.", Computed: true},
			"http_request_set_var_scope":         dschema.StringAttribute{MarkdownDescription: "http_request_set_var_scope. One of `proc`, `sess`, `txn`, `req`, `res`.", Computed: true},
			"http_request_set_var_name":          dschema.StringAttribute{MarkdownDescription: "http_request_set_var_name.", Computed: true},
			"http_request_set_var_expr":          dschema.StringAttribute{MarkdownDescription: "http_request_set_var_expr.", Computed: true},
			"http_response_lua":                  dschema.StringAttribute{MarkdownDescription: "http_response_lua.", Computed: true},
			"http_response_add_header_name":      dschema.StringAttribute{MarkdownDescription: "http_response_add_header_name.", Computed: true},
			"http_response_add_header_content":   dschema.StringAttribute{MarkdownDescription: "http_response_add_header_content.", Computed: true},
			"http_response_set_header_name":      dschema.StringAttribute{MarkdownDescription: "http_response_set_header_name.", Computed: true},
			"http_response_set_header_content":   dschema.StringAttribute{MarkdownDescription: "http_response_set_header_content.", Computed: true},
			"http_response_del_header_name":      dschema.StringAttribute{MarkdownDescription: "http_response_del_header_name.", Computed: true},
			"http_response_replace_header_name":  dschema.StringAttribute{MarkdownDescription: "http_response_replace_header_name.", Computed: true},
			"http_response_replace_header_regex": dschema.StringAttribute{MarkdownDescription: "http_response_replace_header_regex.", Computed: true},
			"http_response_replace_value_name":   dschema.StringAttribute{MarkdownDescription: "http_response_replace_value_name.", Computed: true},
			"http_response_replace_value_regex":  dschema.StringAttribute{MarkdownDescription: "http_response_replace_value_regex.", Computed: true},
			"http_response_set_status_code":      dschema.Int64Attribute{MarkdownDescription: "http_response_set_status_code.", Computed: true},
			"http_response_set_status_reason":    dschema.StringAttribute{MarkdownDescription: "http_response_set_status_reason.", Computed: true},
			"http_response_set_var_scope":        dschema.StringAttribute{MarkdownDescription: "http_response_set_var_scope. One of `proc`, `sess`, `txn`, `req`, `res`.", Computed: true},
			"http_response_set_var_name":         dschema.StringAttribute{MarkdownDescription: "http_response_set_var_name.", Computed: true},
			"http_response_set_var_expr":         dschema.StringAttribute{MarkdownDescription: "http_response_set_var_expr.", Computed: true},
			"tcp_request_content_lua":            dschema.StringAttribute{MarkdownDescription: "tcp_request_content_lua.", Computed: true},
			"tcp_request_content_use_service":    dschema.StringAttribute{MarkdownDescription: "tcp_request_content_use_service.", Computed: true},
			"tcp_request_inspect_delay":          dschema.StringAttribute{MarkdownDescription: "tcp_request_inspect_delay.", Computed: true},
			"tcp_response_content_lua":           dschema.StringAttribute{MarkdownDescription: "tcp_response_content_lua.", Computed: true},
			"tcp_response_inspect_delay":         dschema.StringAttribute{MarkdownDescription: "tcp_response_inspect_delay.", Computed: true},
			"map_data_use_backend_file":          dschema.StringAttribute{MarkdownDescription: "HAProxy will extract the Host header from the HTTP request and search the map file for a match. If a match is found, the backend pool from the map file will be used.", Computed: true},
			"map_data_use_backend_default":       dschema.StringAttribute{MarkdownDescription: "HAProxy will use this backend pool if no match is found in the map file.", Computed: true},
			"map_data_use_backend_input":         dschema.StringAttribute{MarkdownDescription: "The sample expression (e.g., path, req.ssl_sni, hdr(Host), str(active)) whose fetched value is used as the lookup key for the map file.", Computed: true},
			"map_use_backend_file":               dschema.StringAttribute{MarkdownDescription: "HAProxy will extract the Host header from the HTTP request and search the map file for a match. If a match is found, the backend pool from the map file will be used.", Computed: true},
			"map_use_backend_default":            dschema.StringAttribute{MarkdownDescription: "HAProxy will use this backend pool if no match is found in the map file.", Computed: true},
			"compression_algo_res":               dschema.StringAttribute{MarkdownDescription: "The compression algorithm used for HTTP responses. Compression will only be used depending on the Accept-Encoding request header. One of `gzip`, `deflate`, `raw-deflate`.", Computed: true},
			"compression_algo_req":               dschema.StringAttribute{MarkdownDescription: "The compression algorithm used for HTTP requests. Compression will only be used depending on the Accept-Encoding request header. One of `gzip`, `deflate`, `raw-deflate`.", Computed: true},
			"compression_mime_res":               dschema.SetAttribute{MarkdownDescription: "HTTP responses will use compression for the specified mime types, i.e. text/css, text/html, application/javascript.", Computed: true, ElementType: types.StringType},
			"compression_mime_req":               dschema.SetAttribute{MarkdownDescription: "HTTP requests will use compression for the specified mime types, i.e. text/css, text/html, application/javascript.", Computed: true, ElementType: types.StringType},
			"compression_offloading":             dschema.BoolAttribute{MarkdownDescription: "When enabled, the load balancer will remove the Accept-Encoding header before forwarding a request to backend servers. This prevents the servers from performing compression, offloading that work to the load balancer. This may impact HAProxy's performance.", Computed: true},
			"compression_minsize_res":            dschema.Int64Attribute{MarkdownDescription: "Specifies a minimum file size for HTTP response compression in bytes. By avoiding unnecessary compression work on very small files, this may reduce CPU overhead. Entering 0 disables this option.", Computed: true},
			"compression_minsize_req":            dschema.Int64Attribute{MarkdownDescription: "Specifies a minimum file size for HTTP request compression in bytes. By avoiding unnecessary compression work on very small files, this may reduce CPU overhead. Entering 0 disables this option.", Computed: true},
			"compression_direction":              dschema.StringAttribute{MarkdownDescription: "Whether HAProxy uses compression for HTTP responses, HTTP requests, or both. Be aware that most servers will not support request body decompression. For safekeeping only enable compression for HTTP responses. One of `response`, `request`, `both`.", Computed: true},
			"gpc_number":                         dschema.Int64Attribute{MarkdownDescription: "Refers to the number of the General Purpose Counter, e.g. 0 or 1 when using gpc0 or gpc1 respectively. This value will be ignored in rules that do not support GPC.", Computed: true},
			"gpt_number":                         dschema.Int64Attribute{MarkdownDescription: "Refers to the number of the General Purpose Tag, e.g. 0 or 1 when using gpt0 or gpt1 respectively. This value will be ignored in rules that do not support GPT.", Computed: true},
			"sc_number":                          dschema.Int64Attribute{MarkdownDescription: "Refers to the number of the Stick Counter, e.g. 0 or 1 when using sc0 or sc1 respectively. This value will be ignored in rules that do not support SC.", Computed: true},
			"mapfile":                            dschema.StringAttribute{MarkdownDescription: "Load patterns from a map file.", Computed: true},
			"map_default":                        dschema.StringAttribute{MarkdownDescription: "When a map file is specified, this is the default value.", Computed: true},
			"sample_fetch":                       dschema.StringAttribute{MarkdownDescription: "Extracts a value from the request or connection context (e.g. path, header, source IP) to be used as input for converters or matching rules. May be used in combination with set-var and map files to create more complex rules, e.g.http-request set-var(req.rate_limit) path,map_beg(/path/to/mapfile,20)http-request set-var(req.request_rate) base32+src,table_http_req_rate().", Computed: true},
		},
	}
}

func convertActionSchemaToStruct(d *actionResourceModel) (*haproxy.Action, error) {
	return &haproxy.Action{
		Enabled:                        tools.BoolToString(d.Enabled.ValueBool()),
		Name:                           d.Name.ValueString(),
		Description:                    d.Description.ValueString(),
		TestType:                       api.SelectedMap(d.TestType.ValueString()),
		LinkedAcls:                     api.SelectedMapList(setToSlice(d.LinkedAcls)),
		Operator:                       api.SelectedMap(d.Operator.ValueString()),
		Type:                           api.SelectedMap(d.Type.ValueString()),
		UseBackend:                     api.SelectedMap(d.UseBackend.ValueString()),
		UseServer:                      api.SelectedMap(d.UseServer.ValueString()),
		FcgiPassHeader:                 d.FcgiPassHeader.ValueString(),
		FcgiSetParam:                   d.FcgiSetParam.ValueString(),
		MonitorFailUri:                 d.MonitorFailUri.ValueString(),
		Custom:                         d.Custom.ValueString(),
		HttpAfterResponseAction:        api.SelectedMap(d.HttpAfterResponseAction.ValueString()),
		HttpAfterResponseOption:        d.HttpAfterResponseOption.ValueString(),
		HttpRequestAction:              api.SelectedMap(d.HttpRequestAction.ValueString()),
		HttpRequestOption:              d.HttpRequestOption.ValueString(),
		HttpResponseAction:             api.SelectedMap(d.HttpResponseAction.ValueString()),
		HttpResponseOption:             d.HttpResponseOption.ValueString(),
		TcpRequestAction:               api.SelectedMap(d.TcpRequestAction.ValueString()),
		TcpRequestOption:               d.TcpRequestOption.ValueString(),
		TcpResponseAction:              api.SelectedMap(d.TcpResponseAction.ValueString()),
		TcpResponseOption:              d.TcpResponseOption.ValueString(),
		HttpRequestAuth:                d.HttpRequestAuth.ValueString(),
		HttpRequestDenyStatus:          tools.Int64ToStringNegative(d.HttpRequestDenyStatus.ValueInt64()),
		HttpRequestRedirect:            d.HttpRequestRedirect.ValueString(),
		HttpRequestLua:                 d.HttpRequestLua.ValueString(),
		HttpRequestUseService:          d.HttpRequestUseService.ValueString(),
		HttpRequestAddHeaderName:       d.HttpRequestAddHeaderName.ValueString(),
		HttpRequestAddHeaderContent:    d.HttpRequestAddHeaderContent.ValueString(),
		HttpRequestSetHeaderName:       d.HttpRequestSetHeaderName.ValueString(),
		HttpRequestSetHeaderContent:    d.HttpRequestSetHeaderContent.ValueString(),
		HttpRequestDelHeaderName:       d.HttpRequestDelHeaderName.ValueString(),
		HttpRequestReplaceHeaderName:   d.HttpRequestReplaceHeaderName.ValueString(),
		HttpRequestReplaceHeaderRegex:  d.HttpRequestReplaceHeaderRegex.ValueString(),
		HttpRequestReplaceValueName:    d.HttpRequestReplaceValueName.ValueString(),
		HttpRequestReplaceValueRegex:   d.HttpRequestReplaceValueRegex.ValueString(),
		HttpRequestSetPath:             d.HttpRequestSetPath.ValueString(),
		HttpRequestSetVarScope:         api.SelectedMap(d.HttpRequestSetVarScope.ValueString()),
		HttpRequestSetVarName:          d.HttpRequestSetVarName.ValueString(),
		HttpRequestSetVarExpr:          d.HttpRequestSetVarExpr.ValueString(),
		HttpResponseLua:                d.HttpResponseLua.ValueString(),
		HttpResponseAddHeaderName:      d.HttpResponseAddHeaderName.ValueString(),
		HttpResponseAddHeaderContent:   d.HttpResponseAddHeaderContent.ValueString(),
		HttpResponseSetHeaderName:      d.HttpResponseSetHeaderName.ValueString(),
		HttpResponseSetHeaderContent:   d.HttpResponseSetHeaderContent.ValueString(),
		HttpResponseDelHeaderName:      d.HttpResponseDelHeaderName.ValueString(),
		HttpResponseReplaceHeaderName:  d.HttpResponseReplaceHeaderName.ValueString(),
		HttpResponseReplaceHeaderRegex: d.HttpResponseReplaceHeaderRegex.ValueString(),
		HttpResponseReplaceValueName:   d.HttpResponseReplaceValueName.ValueString(),
		HttpResponseReplaceValueRegex:  d.HttpResponseReplaceValueRegex.ValueString(),
		HttpResponseSetStatusCode:      tools.Int64ToStringNegative(d.HttpResponseSetStatusCode.ValueInt64()),
		HttpResponseSetStatusReason:    d.HttpResponseSetStatusReason.ValueString(),
		HttpResponseSetVarScope:        api.SelectedMap(d.HttpResponseSetVarScope.ValueString()),
		HttpResponseSetVarName:         d.HttpResponseSetVarName.ValueString(),
		HttpResponseSetVarExpr:         d.HttpResponseSetVarExpr.ValueString(),
		TcpRequestContentLua:           d.TcpRequestContentLua.ValueString(),
		TcpRequestContentUseService:    d.TcpRequestContentUseService.ValueString(),
		TcpRequestInspectDelay:         d.TcpRequestInspectDelay.ValueString(),
		TcpResponseContentLua:          d.TcpResponseContentLua.ValueString(),
		TcpResponseInspectDelay:        d.TcpResponseInspectDelay.ValueString(),
		MapDataUseBackendFile:          api.SelectedMap(d.MapDataUseBackendFile.ValueString()),
		MapDataUseBackendDefault:       api.SelectedMap(d.MapDataUseBackendDefault.ValueString()),
		MapDataUseBackendInput:         d.MapDataUseBackendInput.ValueString(),
		MapUseBackendFile:              api.SelectedMap(d.MapUseBackendFile.ValueString()),
		MapUseBackendDefault:           api.SelectedMap(d.MapUseBackendDefault.ValueString()),
		CompressionAlgoRes:             api.SelectedMap(d.CompressionAlgoRes.ValueString()),
		CompressionAlgoReq:             api.SelectedMap(d.CompressionAlgoReq.ValueString()),
		CompressionMimeRes:             api.SelectedMapList(setToSlice(d.CompressionMimeRes)),
		CompressionMimeReq:             api.SelectedMapList(setToSlice(d.CompressionMimeReq)),
		CompressionOffloading:          tools.BoolToString(d.CompressionOffloading.ValueBool()),
		CompressionMinsizeRes:          tools.Int64ToString(d.CompressionMinsizeRes.ValueInt64()),
		CompressionMinsizeReq:          tools.Int64ToString(d.CompressionMinsizeReq.ValueInt64()),
		CompressionDirection:           api.SelectedMap(d.CompressionDirection.ValueString()),
		GpcNumber:                      tools.Int64ToStringNegative(d.GpcNumber.ValueInt64()),
		GptNumber:                      tools.Int64ToStringNegative(d.GptNumber.ValueInt64()),
		ScNumber:                       tools.Int64ToStringNegative(d.ScNumber.ValueInt64()),
		Mapfile:                        api.SelectedMap(d.Mapfile.ValueString()),
		MapDefault:                     d.MapDefault.ValueString(),
		SampleFetch:                    d.SampleFetch.ValueString(),
	}, nil
}

func convertActionStructToSchema(d *haproxy.Action) (*actionResourceModel, error) {
	return &actionResourceModel{
		Enabled:                        types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:                           types.StringValue(d.Name),
		Description:                    types.StringValue(d.Description),
		TestType:                       types.StringValue(d.TestType.String()),
		LinkedAcls:                     tools.StringSliceToSet(d.LinkedAcls),
		Operator:                       types.StringValue(d.Operator.String()),
		Type:                           types.StringValue(d.Type.String()),
		UseBackend:                     types.StringValue(d.UseBackend.String()),
		UseServer:                      types.StringValue(d.UseServer.String()),
		FcgiPassHeader:                 types.StringValue(d.FcgiPassHeader),
		FcgiSetParam:                   types.StringValue(d.FcgiSetParam),
		MonitorFailUri:                 types.StringValue(d.MonitorFailUri),
		Custom:                         types.StringValue(d.Custom),
		HttpAfterResponseAction:        types.StringValue(d.HttpAfterResponseAction.String()),
		HttpAfterResponseOption:        types.StringValue(d.HttpAfterResponseOption),
		HttpRequestAction:              types.StringValue(d.HttpRequestAction.String()),
		HttpRequestOption:              types.StringValue(d.HttpRequestOption),
		HttpResponseAction:             types.StringValue(d.HttpResponseAction.String()),
		HttpResponseOption:             types.StringValue(d.HttpResponseOption),
		TcpRequestAction:               types.StringValue(d.TcpRequestAction.String()),
		TcpRequestOption:               types.StringValue(d.TcpRequestOption),
		TcpResponseAction:              types.StringValue(d.TcpResponseAction.String()),
		TcpResponseOption:              types.StringValue(d.TcpResponseOption),
		HttpRequestAuth:                types.StringValue(d.HttpRequestAuth),
		HttpRequestDenyStatus:          types.Int64Value(tools.StringToInt64(d.HttpRequestDenyStatus)),
		HttpRequestRedirect:            types.StringValue(d.HttpRequestRedirect),
		HttpRequestLua:                 types.StringValue(d.HttpRequestLua),
		HttpRequestUseService:          types.StringValue(d.HttpRequestUseService),
		HttpRequestAddHeaderName:       types.StringValue(d.HttpRequestAddHeaderName),
		HttpRequestAddHeaderContent:    types.StringValue(d.HttpRequestAddHeaderContent),
		HttpRequestSetHeaderName:       types.StringValue(d.HttpRequestSetHeaderName),
		HttpRequestSetHeaderContent:    types.StringValue(d.HttpRequestSetHeaderContent),
		HttpRequestDelHeaderName:       types.StringValue(d.HttpRequestDelHeaderName),
		HttpRequestReplaceHeaderName:   types.StringValue(d.HttpRequestReplaceHeaderName),
		HttpRequestReplaceHeaderRegex:  types.StringValue(d.HttpRequestReplaceHeaderRegex),
		HttpRequestReplaceValueName:    types.StringValue(d.HttpRequestReplaceValueName),
		HttpRequestReplaceValueRegex:   types.StringValue(d.HttpRequestReplaceValueRegex),
		HttpRequestSetPath:             types.StringValue(d.HttpRequestSetPath),
		HttpRequestSetVarScope:         types.StringValue(d.HttpRequestSetVarScope.String()),
		HttpRequestSetVarName:          types.StringValue(d.HttpRequestSetVarName),
		HttpRequestSetVarExpr:          types.StringValue(d.HttpRequestSetVarExpr),
		HttpResponseLua:                types.StringValue(d.HttpResponseLua),
		HttpResponseAddHeaderName:      types.StringValue(d.HttpResponseAddHeaderName),
		HttpResponseAddHeaderContent:   types.StringValue(d.HttpResponseAddHeaderContent),
		HttpResponseSetHeaderName:      types.StringValue(d.HttpResponseSetHeaderName),
		HttpResponseSetHeaderContent:   types.StringValue(d.HttpResponseSetHeaderContent),
		HttpResponseDelHeaderName:      types.StringValue(d.HttpResponseDelHeaderName),
		HttpResponseReplaceHeaderName:  types.StringValue(d.HttpResponseReplaceHeaderName),
		HttpResponseReplaceHeaderRegex: types.StringValue(d.HttpResponseReplaceHeaderRegex),
		HttpResponseReplaceValueName:   types.StringValue(d.HttpResponseReplaceValueName),
		HttpResponseReplaceValueRegex:  types.StringValue(d.HttpResponseReplaceValueRegex),
		HttpResponseSetStatusCode:      types.Int64Value(tools.StringToInt64(d.HttpResponseSetStatusCode)),
		HttpResponseSetStatusReason:    types.StringValue(d.HttpResponseSetStatusReason),
		HttpResponseSetVarScope:        types.StringValue(d.HttpResponseSetVarScope.String()),
		HttpResponseSetVarName:         types.StringValue(d.HttpResponseSetVarName),
		HttpResponseSetVarExpr:         types.StringValue(d.HttpResponseSetVarExpr),
		TcpRequestContentLua:           types.StringValue(d.TcpRequestContentLua),
		TcpRequestContentUseService:    types.StringValue(d.TcpRequestContentUseService),
		TcpRequestInspectDelay:         types.StringValue(d.TcpRequestInspectDelay),
		TcpResponseContentLua:          types.StringValue(d.TcpResponseContentLua),
		TcpResponseInspectDelay:        types.StringValue(d.TcpResponseInspectDelay),
		MapDataUseBackendFile:          types.StringValue(d.MapDataUseBackendFile.String()),
		MapDataUseBackendDefault:       types.StringValue(d.MapDataUseBackendDefault.String()),
		MapDataUseBackendInput:         types.StringValue(d.MapDataUseBackendInput),
		MapUseBackendFile:              types.StringValue(d.MapUseBackendFile.String()),
		MapUseBackendDefault:           types.StringValue(d.MapUseBackendDefault.String()),
		CompressionAlgoRes:             types.StringValue(d.CompressionAlgoRes.String()),
		CompressionAlgoReq:             types.StringValue(d.CompressionAlgoReq.String()),
		CompressionMimeRes:             tools.StringSliceToSet(d.CompressionMimeRes),
		CompressionMimeReq:             tools.StringSliceToSet(d.CompressionMimeReq),
		CompressionOffloading:          types.BoolValue(tools.StringToBool(d.CompressionOffloading)),
		CompressionMinsizeRes:          types.Int64Value(tools.StringToInt64(d.CompressionMinsizeRes)),
		CompressionMinsizeReq:          types.Int64Value(tools.StringToInt64(d.CompressionMinsizeReq)),
		CompressionDirection:           types.StringValue(d.CompressionDirection.String()),
		GpcNumber:                      types.Int64Value(tools.StringToInt64(d.GpcNumber)),
		GptNumber:                      types.Int64Value(tools.StringToInt64(d.GptNumber)),
		ScNumber:                       types.Int64Value(tools.StringToInt64(d.ScNumber)),
		Mapfile:                        types.StringValue(d.Mapfile.String()),
		MapDefault:                     types.StringValue(d.MapDefault),
		SampleFetch:                    types.StringValue(d.SampleFetch),
	}, nil
}
