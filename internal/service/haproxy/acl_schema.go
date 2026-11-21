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

type aclResourceModel struct {
	Name                      types.String `tfsdk:"name"`
	Description               types.String `tfsdk:"description"`
	Expression                types.String `tfsdk:"expression"`
	Negate                    types.Bool   `tfsdk:"negate"`
	CaseSensitive             types.Bool   `tfsdk:"case_sensitive"`
	HdrBeg                    types.String `tfsdk:"hdr_beg"`
	HdrEnd                    types.String `tfsdk:"hdr_end"`
	Hdr                       types.String `tfsdk:"hdr"`
	HdrReg                    types.String `tfsdk:"hdr_reg"`
	HdrSub                    types.String `tfsdk:"hdr_sub"`
	PathBeg                   types.String `tfsdk:"path_beg"`
	PathEnd                   types.String `tfsdk:"path_end"`
	Path                      types.String `tfsdk:"path"`
	PathReg                   types.String `tfsdk:"path_reg"`
	PathDir                   types.String `tfsdk:"path_dir"`
	PathSub                   types.String `tfsdk:"path_sub"`
	CustHdrBegName            types.String `tfsdk:"cust_hdr_beg_name"`
	CustHdrBeg                types.String `tfsdk:"cust_hdr_beg"`
	CustHdrEndName            types.String `tfsdk:"cust_hdr_end_name"`
	CustHdrEnd                types.String `tfsdk:"cust_hdr_end"`
	CustHdrName               types.String `tfsdk:"cust_hdr_name"`
	CustHdr                   types.String `tfsdk:"cust_hdr"`
	CustHdrRegName            types.String `tfsdk:"cust_hdr_reg_name"`
	CustHdrReg                types.String `tfsdk:"cust_hdr_reg"`
	CustHdrSubName            types.String `tfsdk:"cust_hdr_sub_name"`
	CustHdrSub                types.String `tfsdk:"cust_hdr_sub"`
	URLParam                  types.String `tfsdk:"url_param"`
	URLParamValue             types.String `tfsdk:"url_param_value"`
	Var                       types.String `tfsdk:"var"`
	VarValue                  types.String `tfsdk:"var_value"`
	VarComparison             types.String `tfsdk:"var_comparison"`
	SslCVerifyCode            types.Int64  `tfsdk:"ssl_c_verify_code"`
	SslCCaCommonname          types.String `tfsdk:"ssl_c_ca_commonname"`
	SslHelloType              types.String `tfsdk:"ssl_hello_type"`
	Src                       types.String `tfsdk:"src"`
	SrcBytesInRateComparison  types.String `tfsdk:"src_bytes_in_rate_comparison"`
	SrcBytesInRate            types.Int64  `tfsdk:"src_bytes_in_rate"`
	SrcBytesOutRateComparison types.String `tfsdk:"src_bytes_out_rate_comparison"`
	SrcBytesOutRate           types.Int64  `tfsdk:"src_bytes_out_rate"`
	SrcConnCntComparison      types.String `tfsdk:"src_conn_cnt_comparison"`
	SrcConnCnt                types.Int64  `tfsdk:"src_conn_cnt"`
	SrcConnCurComparison      types.String `tfsdk:"src_conn_cur_comparison"`
	SrcConnCur                types.Int64  `tfsdk:"src_conn_cur"`
	SrcConnRateComparison     types.String `tfsdk:"src_conn_rate_comparison"`
	SrcConnRate               types.Int64  `tfsdk:"src_conn_rate"`
	SrcHttpErrCntComparison   types.String `tfsdk:"src_http_err_cnt_comparison"`
	SrcHttpErrCnt             types.Int64  `tfsdk:"src_http_err_cnt"`
	SrcHttpErrRateComparison  types.String `tfsdk:"src_http_err_rate_comparison"`
	SrcHttpErrRate            types.Int64  `tfsdk:"src_http_err_rate"`
	SrcHttpReqCntComparison   types.String `tfsdk:"src_http_req_cnt_comparison"`
	SrcHttpReqCnt             types.Int64  `tfsdk:"src_http_req_cnt"`
	SrcHttpReqRateComparison  types.String `tfsdk:"src_http_req_rate_comparison"`
	SrcHttpReqRate            types.Int64  `tfsdk:"src_http_req_rate"`
	SrcKbytesInComparison     types.String `tfsdk:"src_kbytes_in_comparison"`
	SrcKbytesIn               types.Int64  `tfsdk:"src_kbytes_in"`
	SrcKbytesOutComparison    types.String `tfsdk:"src_kbytes_out_comparison"`
	SrcKbytesOut              types.Int64  `tfsdk:"src_kbytes_out"`
	SrcPortComparison         types.String `tfsdk:"src_port_comparison"`
	SrcPort                   types.Int64  `tfsdk:"src_port"`
	SrcSessCntComparison      types.String `tfsdk:"src_sess_cnt_comparison"`
	SrcSessCnt                types.Int64  `tfsdk:"src_sess_cnt"`
	SrcSessRateComparison     types.String `tfsdk:"src_sess_rate_comparison"`
	SrcSessRate               types.Int64  `tfsdk:"src_sess_rate"`
	Nbsrv                     types.Int64  `tfsdk:"nbsrv"`
	NbsrvBackend              types.String `tfsdk:"nbsrv_backend"`
	SslFcSni                  types.String `tfsdk:"ssl_fc_sni"`
	SslSni                    types.String `tfsdk:"ssl_sni"`
	SslSniSub                 types.String `tfsdk:"ssl_sni_sub"`
	SslSniBeg                 types.String `tfsdk:"ssl_sni_beg"`
	SslSniEnd                 types.String `tfsdk:"ssl_sni_end"`
	SslSniReg                 types.String `tfsdk:"ssl_sni_reg"`
	CustomAcl                 types.String `tfsdk:"custom_acl"`
	Value                     types.String `tfsdk:"value"`
	Urlparam                  types.String `tfsdk:"urlparam"`
	QueryBackend              types.String `tfsdk:"query_backend"`
	AllowedUsers              types.Set    `tfsdk:"allowed_users"`
	AllowedGroups             types.Set    `tfsdk:"allowed_groups"`
	HttpMethod                types.Set    `tfsdk:"http_method"`
	ScBytesInRateComparison   types.String `tfsdk:"sc_bytes_in_rate_comparison"`
	ScBytesInRate             types.Int64  `tfsdk:"sc_bytes_in_rate"`
	ScBytesOutRateComparison  types.String `tfsdk:"sc_bytes_out_rate_comparison"`
	ScBytesOutRate            types.Int64  `tfsdk:"sc_bytes_out_rate"`
	ScClrGpcComparison        types.String `tfsdk:"sc_clr_gpc_comparison"`
	ScClrGpc                  types.Int64  `tfsdk:"sc_clr_gpc"`
	ScConnCntComparison       types.String `tfsdk:"sc_conn_cnt_comparison"`
	ScConnCnt                 types.Int64  `tfsdk:"sc_conn_cnt"`
	ScConnCurComparison       types.String `tfsdk:"sc_conn_cur_comparison"`
	ScConnCur                 types.Int64  `tfsdk:"sc_conn_cur"`
	ScConnRateComparison      types.String `tfsdk:"sc_conn_rate_comparison"`
	ScConnRate                types.Int64  `tfsdk:"sc_conn_rate"`
	ScGetGpcComparison        types.String `tfsdk:"sc_get_gpc_comparison"`
	ScGetGpc                  types.Int64  `tfsdk:"sc_get_gpc"`
	ScGlitchCntComparison     types.String `tfsdk:"sc_glitch_cnt_comparison"`
	ScGlitchCnt               types.Int64  `tfsdk:"sc_glitch_cnt"`
	ScGlitchRateComparison    types.String `tfsdk:"sc_glitch_rate_comparison"`
	ScGlitchRate              types.Int64  `tfsdk:"sc_glitch_rate"`
	ScGpcRateComparison       types.String `tfsdk:"sc_gpc_rate_comparison"`
	ScGpcRate                 types.Int64  `tfsdk:"sc_gpc_rate"`
	ScHttpErrCntComparison    types.String `tfsdk:"sc_http_err_cnt_comparison"`
	ScHttpErrCnt              types.Int64  `tfsdk:"sc_http_err_cnt"`
	ScHttpErrRateComparison   types.String `tfsdk:"sc_http_err_rate_comparison"`
	ScHttpErrRate             types.Int64  `tfsdk:"sc_http_err_rate"`
	ScHttpFailCntComparison   types.String `tfsdk:"sc_http_fail_cnt_comparison"`
	ScHttpFailCnt             types.Int64  `tfsdk:"sc_http_fail_cnt"`
	ScHttpFailRateComparison  types.String `tfsdk:"sc_http_fail_rate_comparison"`
	ScHttpFailRate            types.Int64  `tfsdk:"sc_http_fail_rate"`
	ScHttpReqCntComparison    types.String `tfsdk:"sc_http_req_cnt_comparison"`
	ScHttpReqCnt              types.Int64  `tfsdk:"sc_http_req_cnt"`
	ScHttpReqRateComparison   types.String `tfsdk:"sc_http_req_rate_comparison"`
	ScHttpReqRate             types.Int64  `tfsdk:"sc_http_req_rate"`
	ScIncGpcComparison        types.String `tfsdk:"sc_inc_gpc_comparison"`
	ScIncGpc                  types.Int64  `tfsdk:"sc_inc_gpc"`
	ScSessCntComparison       types.String `tfsdk:"sc_sess_cnt_comparison"`
	ScSessCnt                 types.Int64  `tfsdk:"sc_sess_cnt"`
	ScSessRateComparison      types.String `tfsdk:"sc_sess_rate_comparison"`
	ScSessRate                types.Int64  `tfsdk:"sc_sess_rate"`
	SrcGetGpcComparison       types.String `tfsdk:"src_get_gpc_comparison"`
	SrcGetGpc                 types.Int64  `tfsdk:"src_get_gpc"`
	SrcGetGptComparison       types.String `tfsdk:"src_get_gpt_comparison"`
	SrcGetGpt                 types.Int64  `tfsdk:"src_get_gpt"`
	SrcGlitchCntComparison    types.String `tfsdk:"src_glitch_cnt_comparison"`
	SrcGlitchCnt              types.Int64  `tfsdk:"src_glitch_cnt"`
	SrcGlitchRateComparison   types.String `tfsdk:"src_glitch_rate_comparison"`
	SrcGlitchRate             types.Int64  `tfsdk:"src_glitch_rate"`
	SrcGpcRateComparison      types.String `tfsdk:"src_gpc_rate_comparison"`
	SrcGpcRate                types.Int64  `tfsdk:"src_gpc_rate"`
	SrcHttpFailCntComparison  types.String `tfsdk:"src_http_fail_cnt_comparison"`
	SrcHttpFailCnt            types.Int64  `tfsdk:"src_http_fail_cnt"`
	SrcHttpFailRateComparison types.String `tfsdk:"src_http_fail_rate_comparison"`
	SrcHttpFailRate           types.Int64  `tfsdk:"src_http_fail_rate"`
	SrcIncGpcComparison       types.String `tfsdk:"src_inc_gpc_comparison"`
	SrcIncGpc                 types.Int64  `tfsdk:"src_inc_gpc"`
	ScClrGpc0Comparison       types.String `tfsdk:"sc_clr_gpc0_comparison"`
	ScClrGpc0                 types.Int64  `tfsdk:"sc_clr_gpc0"`
	ScClrGpc1Comparison       types.String `tfsdk:"sc_clr_gpc1_comparison"`
	ScClrGpc1                 types.Int64  `tfsdk:"sc_clr_gpc1"`
	Sc0ClrGpc0Comparison      types.String `tfsdk:"sc0_clr_gpc0_comparison"`
	Sc0ClrGpc0                types.Int64  `tfsdk:"sc0_clr_gpc0"`
	Sc0ClrGpc1Comparison      types.String `tfsdk:"sc0_clr_gpc1_comparison"`
	Sc0ClrGpc1                types.Int64  `tfsdk:"sc0_clr_gpc1"`
	Sc1ClrGpcComparison       types.String `tfsdk:"sc1_clr_gpc_comparison"`
	Sc1ClrGpc                 types.Int64  `tfsdk:"sc1_clr_gpc"`
	Sc1ClrGpc0Comparison      types.String `tfsdk:"sc1_clr_gpc0_comparison"`
	Sc1ClrGpc0                types.Int64  `tfsdk:"sc1_clr_gpc0"`
	Sc1ClrGpc1Comparison      types.String `tfsdk:"sc1_clr_gpc1_comparison"`
	Sc1ClrGpc1                types.Int64  `tfsdk:"sc1_clr_gpc1"`
	Sc2ClrGpcComparison       types.String `tfsdk:"sc2_clr_gpc_comparison"`
	Sc2ClrGpc                 types.Int64  `tfsdk:"sc2_clr_gpc"`
	Sc2ClrGpc0Comparison      types.String `tfsdk:"sc2_clr_gpc0_comparison"`
	Sc2ClrGpc0                types.Int64  `tfsdk:"sc2_clr_gpc0"`
	Sc2ClrGpc1Comparison      types.String `tfsdk:"sc2_clr_gpc1_comparison"`
	Sc2ClrGpc1                types.Int64  `tfsdk:"sc2_clr_gpc1"`
	ScGetGpc0Comparison       types.String `tfsdk:"sc_get_gpc0_comparison"`
	ScGetGpc0                 types.Int64  `tfsdk:"sc_get_gpc0"`
	ScGetGpc1Comparison       types.String `tfsdk:"sc_get_gpc1_comparison"`
	ScGetGpc1                 types.Int64  `tfsdk:"sc_get_gpc1"`
	Sc0GetGpc0Comparison      types.String `tfsdk:"sc0_get_gpc0_comparison"`
	Sc0GetGpc0                types.Int64  `tfsdk:"sc0_get_gpc0"`
	Sc0GetGpc1Comparison      types.String `tfsdk:"sc0_get_gpc1_comparison"`
	Sc0GetGpc1                types.Int64  `tfsdk:"sc0_get_gpc1"`
	Sc1GetGpc0Comparison      types.String `tfsdk:"sc1_get_gpc0_comparison"`
	Sc1GetGpc0                types.Int64  `tfsdk:"sc1_get_gpc0"`
	Sc1GetGpc1Comparison      types.String `tfsdk:"sc1_get_gpc1_comparison"`
	Sc1GetGpc1                types.Int64  `tfsdk:"sc1_get_gpc1"`
	Sc2GetGpc0Comparison      types.String `tfsdk:"sc2_get_gpc0_comparison"`
	Sc2GetGpc0                types.Int64  `tfsdk:"sc2_get_gpc0"`
	Sc2GetGpc1Comparison      types.String `tfsdk:"sc2_get_gpc1_comparison"`
	Sc2GetGpc1                types.Int64  `tfsdk:"sc2_get_gpc1"`
	ScGetGptComparison        types.String `tfsdk:"sc_get_gpt_comparison"`
	ScGetGpt                  types.Int64  `tfsdk:"sc_get_gpt"`
	ScGetGpt0Comparison       types.String `tfsdk:"sc_get_gpt0_comparison"`
	ScGetGpt0                 types.Int64  `tfsdk:"sc_get_gpt0"`
	Sc0GetGpt0Comparison      types.String `tfsdk:"sc0_get_gpt0_comparison"`
	Sc0GetGpt0                types.Int64  `tfsdk:"sc0_get_gpt0"`
	Sc1GetGpt0Comparison      types.String `tfsdk:"sc1_get_gpt0_comparison"`
	Sc1GetGpt0                types.Int64  `tfsdk:"sc1_get_gpt0"`
	Sc2GetGpt0Comparison      types.String `tfsdk:"sc2_get_gpt0_comparison"`
	Sc2GetGpt0                types.Int64  `tfsdk:"sc2_get_gpt0"`
	ScGpc0RateComparison      types.String `tfsdk:"sc_gpc0_rate_comparison"`
	ScGpc0Rate                types.Int64  `tfsdk:"sc_gpc0_rate"`
	ScGpc1RateComparison      types.String `tfsdk:"sc_gpc1_rate_comparison"`
	ScGpc1Rate                types.Int64  `tfsdk:"sc_gpc1_rate"`
	Sc0Gpc0RateComparison     types.String `tfsdk:"sc0_gpc0_rate_comparison"`
	Sc0Gpc0Rate               types.Int64  `tfsdk:"sc0_gpc0_rate"`
	Sc0Gpc1RateComparison     types.String `tfsdk:"sc0_gpc1_rate_comparison"`
	Sc0Gpc1Rate               types.Int64  `tfsdk:"sc0_gpc1_rate"`
	Sc1Gpc0RateComparison     types.String `tfsdk:"sc1_gpc0_rate_comparison"`
	Sc1Gpc0Rate               types.Int64  `tfsdk:"sc1_gpc0_rate"`
	Sc1Gpc1RateComparison     types.String `tfsdk:"sc1_gpc1_rate_comparison"`
	Sc1Gpc1Rate               types.Int64  `tfsdk:"sc1_gpc1_rate"`
	Sc2Gpc0RateComparison     types.String `tfsdk:"sc2_gpc0_rate_comparison"`
	Sc2Gpc0Rate               types.Int64  `tfsdk:"sc2_gpc0_rate"`
	Sc2Gpc1RateComparison     types.String `tfsdk:"sc2_gpc1_rate_comparison"`
	Sc2Gpc1Rate               types.Int64  `tfsdk:"sc2_gpc1_rate"`
	ScIncGpc0Comparison       types.String `tfsdk:"sc_inc_gpc0_comparison"`
	ScIncGpc0                 types.Int64  `tfsdk:"sc_inc_gpc0"`
	ScIncGpc1Comparison       types.String `tfsdk:"sc_inc_gpc1_comparison"`
	ScIncGpc1                 types.Int64  `tfsdk:"sc_inc_gpc1"`
	Sc0IncGpc0Comparison      types.String `tfsdk:"sc0_inc_gpc0_comparison"`
	Sc0IncGpc0                types.Int64  `tfsdk:"sc0_inc_gpc0"`
	Sc0IncGpc1Comparison      types.String `tfsdk:"sc0_inc_gpc1_comparison"`
	Sc0IncGpc1                types.Int64  `tfsdk:"sc0_inc_gpc1"`
	Sc1IncGpc0Comparison      types.String `tfsdk:"sc1_inc_gpc0_comparison"`
	Sc1IncGpc0                types.Int64  `tfsdk:"sc1_inc_gpc0"`
	Sc1IncGpc1Comparison      types.String `tfsdk:"sc1_inc_gpc1_comparison"`
	Sc1IncGpc1                types.Int64  `tfsdk:"sc1_inc_gpc1"`
	Sc2IncGpc0Comparison      types.String `tfsdk:"sc2_inc_gpc0_comparison"`
	Sc2IncGpc0                types.Int64  `tfsdk:"sc2_inc_gpc0"`
	Sc2IncGpc1Comparison      types.String `tfsdk:"sc2_inc_gpc1_comparison"`
	Sc2IncGpc1                types.Int64  `tfsdk:"sc2_inc_gpc1"`
	SrcClrGpc0Comparison      types.String `tfsdk:"src_clr_gpc0_comparison"`
	SrcClrGpc0                types.Int64  `tfsdk:"src_clr_gpc0"`
	SrcClrGpc1Comparison      types.String `tfsdk:"src_clr_gpc1_comparison"`
	SrcClrGpc1                types.Int64  `tfsdk:"src_clr_gpc1"`
	SrcGetGpc0Comparison      types.String `tfsdk:"src_get_gpc0_comparison"`
	SrcGetGpc0                types.Int64  `tfsdk:"src_get_gpc0"`
	SrcGetGpc1Comparison      types.String `tfsdk:"src_get_gpc1_comparison"`
	SrcGetGpc1                types.Int64  `tfsdk:"src_get_gpc1"`
	SrcGpc0RateComparison     types.String `tfsdk:"src_gpc0_rate_comparison"`
	SrcGpc0Rate               types.Int64  `tfsdk:"src_gpc0_rate"`
	SrcGpc1RateComparison     types.String `tfsdk:"src_gpc1_rate_comparison"`
	SrcGpc1Rate               types.Int64  `tfsdk:"src_gpc1_rate"`
	SrcIncGpc0Comparison      types.String `tfsdk:"src_inc_gpc0_comparison"`
	SrcIncGpc0                types.Int64  `tfsdk:"src_inc_gpc0"`
	SrcIncGpc1Comparison      types.String `tfsdk:"src_inc_gpc1_comparison"`
	SrcIncGpc1                types.Int64  `tfsdk:"src_inc_gpc1"`
	GpcNumber                 types.Int64  `tfsdk:"gpc_number"`
	GptNumber                 types.Int64  `tfsdk:"gpt_number"`
	ScNumber                  types.Int64  `tfsdk:"sc_number"`
	TableName                 types.String `tfsdk:"table_name"`
	Mapfile                   types.String `tfsdk:"mapfile"`
	Converter                 types.String `tfsdk:"converter"`

	Id types.String `tfsdk:"id"`
}

func aclResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "HAProxy conditions (**Services ▸ HAProxy ▸ Settings ▸ Rules & Checks ▸ Conditions**, plugin `os-haproxy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify this condition.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this condition. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"expression": schema.StringAttribute{
				MarkdownDescription: "Choose the condition type. Note that a wide range of sample fetch methods is either available in layer 4 or layer 7. The syntax check will show errors when using an incompatible sample fetch method. One of `cust_hdr_beg`, `cust_hdr_end`, `cust_hdr`, `cust_hdr_reg`, `cust_hdr_sub`, `hdr_beg`, `hdr_end`, `hdr`, `hdr_reg`, `hdr_sub`, `http_auth`, `http_method`, `nbsrv`, `path_beg`, `path_dir`, `path_end`, `path`, `path_reg`, `path_sub`, `quic_enabled`, `traffic_is_http`, `traffic_is_ssl`, `sc_bytes_in_rate`, `sc_bytes_out_rate`, `sc_clr_gpc`, `sc_clr_gpc0`, `sc_clr_gpc1`, `sc0_clr_gpc0`, `sc0_clr_gpc1`, `sc1_clr_gpc`, `sc1_clr_gpc0`, `sc1_clr_gpc1`, `sc2_clr_gpc`, `sc2_clr_gpc0`, `sc2_clr_gpc1`, `sc_conn_cnt`, `sc_conn_cur`, `sc_conn_rate`, `sc_get_gpc`, `sc_get_gpc0`, `sc_get_gpc1`, `sc0_get_gpc0`, `sc0_get_gpc1`, `sc1_get_gpc0`, `sc1_get_gpc1`, `sc2_get_gpc0`, `sc2_get_gpc1`, `sc_get_gpt`, `sc_get_gpt0`, `sc0_get_gpt0`, `sc1_get_gpt0`, `sc2_get_gpt0`, `sc_glitch_cnt`, `sc_glitch_rate`, `sc_gpc_rate`, `sc_gpc0_rate`, `sc_gpc1_rate`, `sc0_gpc0_rate`, `sc0_gpc1_rate`, `sc1_gpc0_rate`, `sc1_gpc1_rate`, `sc2_gpc0_rate`, `sc2_gpc1_rate`, `sc_http_err_cnt`, `sc_http_err_rate`, `sc_http_fail_cnt`, `sc_http_fail_rate`, `sc_http_req_cnt`, `sc_http_req_rate`, `sc_inc_gpc`, `sc_inc_gpc0`, `sc_inc_gpc1`, `sc0_inc_gpc0`, `sc0_inc_gpc1`, `sc1_inc_gpc0`, `sc1_inc_gpc1`, `sc2_inc_gpc0`, `sc2_inc_gpc1`, `sc_sess_cnt`, `sc_sess_rate`, `src`, `src_bytes_in_rate`, `src_bytes_out_rate`, `src_clr_gpc`, `src_clr_gpc0`, `src_clr_gpc1`, `src_conn_cnt`, `src_conn_cur`, `src_conn_rate`, `src_get_gpc`, `src_get_gpc0`, `src_get_gpc1`, `src_get_gpt`, `src_glitch_cnt`, `src_glitch_rate`, `src_gpc_rate`, `src_gpc0_rate`, `src_gpc1_rate`, `src_http_err_cnt`, `src_http_err_rate`, `src_http_fail_cnt`, `src_http_fail_rate`, `src_http_req_cnt`, `src_http_req_rate`, `src_inc_gpc`, `src_inc_gpc0`, `src_inc_gpc1`, `src_is_local`, `src_kbytes_in`, `src_kbytes_out`, `src_port`, `src_sess_cnt`, `src_sess_rate`, `ssl_c_ca_commonname`, `ssl_c_verify_code`, `ssl_c_verify`, `ssl_fc_sni`, `ssl_fc`, `ssl_hello_type`, `ssl_sni_beg`, `ssl_sni_end`, `ssl_sni_reg`, `ssl_sni`, `ssl_sni_sub`, `stopping`, `url_param`, `var`, `wait_end`, `custom_acl`.",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf("cust_hdr_beg", "cust_hdr_end", "cust_hdr", "cust_hdr_reg", "cust_hdr_sub", "hdr_beg", "hdr_end", "hdr", "hdr_reg", "hdr_sub", "http_auth", "http_method", "nbsrv", "path_beg", "path_dir", "path_end", "path", "path_reg", "path_sub", "quic_enabled", "traffic_is_http", "traffic_is_ssl", "sc_bytes_in_rate", "sc_bytes_out_rate", "sc_clr_gpc", "sc_clr_gpc0", "sc_clr_gpc1", "sc0_clr_gpc0", "sc0_clr_gpc1", "sc1_clr_gpc", "sc1_clr_gpc0", "sc1_clr_gpc1", "sc2_clr_gpc", "sc2_clr_gpc0", "sc2_clr_gpc1", "sc_conn_cnt", "sc_conn_cur", "sc_conn_rate", "sc_get_gpc", "sc_get_gpc0", "sc_get_gpc1", "sc0_get_gpc0", "sc0_get_gpc1", "sc1_get_gpc0", "sc1_get_gpc1", "sc2_get_gpc0", "sc2_get_gpc1", "sc_get_gpt", "sc_get_gpt0", "sc0_get_gpt0", "sc1_get_gpt0", "sc2_get_gpt0", "sc_glitch_cnt", "sc_glitch_rate", "sc_gpc_rate", "sc_gpc0_rate", "sc_gpc1_rate", "sc0_gpc0_rate", "sc0_gpc1_rate", "sc1_gpc0_rate", "sc1_gpc1_rate", "sc2_gpc0_rate", "sc2_gpc1_rate", "sc_http_err_cnt", "sc_http_err_rate", "sc_http_fail_cnt", "sc_http_fail_rate", "sc_http_req_cnt", "sc_http_req_rate", "sc_inc_gpc", "sc_inc_gpc0", "sc_inc_gpc1", "sc0_inc_gpc0", "sc0_inc_gpc1", "sc1_inc_gpc0", "sc1_inc_gpc1", "sc2_inc_gpc0", "sc2_inc_gpc1", "sc_sess_cnt", "sc_sess_rate", "src", "src_bytes_in_rate", "src_bytes_out_rate", "src_clr_gpc", "src_clr_gpc0", "src_clr_gpc1", "src_conn_cnt", "src_conn_cur", "src_conn_rate", "src_get_gpc", "src_get_gpc0", "src_get_gpc1", "src_get_gpt", "src_glitch_cnt", "src_glitch_rate", "src_gpc_rate", "src_gpc0_rate", "src_gpc1_rate", "src_http_err_cnt", "src_http_err_rate", "src_http_fail_cnt", "src_http_fail_rate", "src_http_req_cnt", "src_http_req_rate", "src_inc_gpc", "src_inc_gpc0", "src_inc_gpc1", "src_is_local", "src_kbytes_in", "src_kbytes_out", "src_port", "src_sess_cnt", "src_sess_rate", "ssl_c_ca_commonname", "ssl_c_verify_code", "ssl_c_verify", "ssl_fc_sni", "ssl_fc", "ssl_hello_type", "ssl_sni_beg", "ssl_sni_end", "ssl_sni_reg", "ssl_sni", "ssl_sni_sub", "stopping", "url_param", "var", "wait_end", "custom_acl")},
			},
			"negate": schema.BoolAttribute{
				MarkdownDescription: "Use this to invert the meaning of the expression. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"case_sensitive": schema.BoolAttribute{
				MarkdownDescription: "Enable to make the condition case-sensitive (if applicable). By default all conditions are not case-sensitive. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"hdr_beg": schema.StringAttribute{
				MarkdownDescription: "HTTP host header starts with string (prefix match). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"hdr_end": schema.StringAttribute{
				MarkdownDescription: "HTTP host header ends with string (suffix match). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"hdr": schema.StringAttribute{
				MarkdownDescription: "HTTP host header matches exact string. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"hdr_reg": schema.StringAttribute{
				MarkdownDescription: "HTTP host header matches regular expression. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"hdr_sub": schema.StringAttribute{
				MarkdownDescription: "HTTP host header contains string (substring match). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"path_beg": schema.StringAttribute{
				MarkdownDescription: "HTTP request URL path starts with string (prefix match). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"path_end": schema.StringAttribute{
				MarkdownDescription: "HTTP request URL path ends with string (suffix match). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"path": schema.StringAttribute{
				MarkdownDescription: "HTTP request URL path matches exact string. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"path_reg": schema.StringAttribute{
				MarkdownDescription: "HTTP request URL path matches regular expression. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"path_dir": schema.StringAttribute{
				MarkdownDescription: "HTTP request URL path contains directory (subdir match). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"path_sub": schema.StringAttribute{
				MarkdownDescription: "HTTP request URL path contains string (substring match). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"cust_hdr_beg_name": schema.StringAttribute{
				MarkdownDescription: "The name of the HTTP Header. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"cust_hdr_beg": schema.StringAttribute{
				MarkdownDescription: "HTTP Header starts with string (prefix match). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"cust_hdr_end_name": schema.StringAttribute{
				MarkdownDescription: "The name of the HTTP Header. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"cust_hdr_end": schema.StringAttribute{
				MarkdownDescription: "HTTP Header ends with string (suffix match). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"cust_hdr_name": schema.StringAttribute{
				MarkdownDescription: "The name of the HTTP Header. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"cust_hdr": schema.StringAttribute{
				MarkdownDescription: "HTTP Header matches exact string. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"cust_hdr_reg_name": schema.StringAttribute{
				MarkdownDescription: "The name of the HTTP Header. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"cust_hdr_reg": schema.StringAttribute{
				MarkdownDescription: "HTTP Header matches regular expression. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"cust_hdr_sub_name": schema.StringAttribute{
				MarkdownDescription: "The name of the HTTP Header. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"cust_hdr_sub": schema.StringAttribute{
				MarkdownDescription: "HTTP Header contains string (substring match). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"url_param": schema.StringAttribute{
				MarkdownDescription: "Specify the URL parameter to be checked for the value specified below. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"url_param_value": schema.StringAttribute{
				MarkdownDescription: "Specify the value for the URL parameter. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"var": schema.StringAttribute{
				MarkdownDescription: "The name of a variable, followed by an optional default value, e.g. (req.rate_limit) or (req.rate_limit,10). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"var_value": schema.StringAttribute{
				MarkdownDescription: "Value. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"var_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"ssl_c_verify_code": schema.Int64Attribute{
				MarkdownDescription: "Specify the SSL/TLS error ID that should be checked for the incoming connection. Please refer to your SSL library's documentation for an exhaustive list of error codes. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"ssl_c_ca_commonname": schema.StringAttribute{
				MarkdownDescription: "Verify the CA Common-Name of the certificate presented by the client against the specified string. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_hello_type": schema.StringAttribute{
				MarkdownDescription: "An integer value containing the type of the SSL hello message found in the request buffer if the buffer contains data that parse as a complete SSL (v3 or superior) client hello message. One of `x0`, `x1`, `x2`. Defaults to `x1`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("x1"),
				Validators:          []validator.String{stringvalidator.OneOf("", "x0", "x1", "x2")},
			},
			"src": schema.StringAttribute{
				MarkdownDescription: "Verify the source IPv4 address of the client of the session matches the specified IPv4 or IPv6 address. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"src_bytes_in_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_bytes_in_rate": schema.Int64Attribute{
				MarkdownDescription: "The average bytes rate from the incoming connection's source address. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_bytes_out_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_bytes_out_rate": schema.Int64Attribute{
				MarkdownDescription: "The average bytes rate to the incoming connection's source address. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_conn_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_conn_cnt": schema.Int64Attribute{
				MarkdownDescription: "The cumulative number of connections initiated from the current incoming connection's source address. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_conn_cur_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_conn_cur": schema.Int64Attribute{
				MarkdownDescription: "The current amount of concurrent connections initiated from the current incoming connection's source address. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_conn_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_conn_rate": schema.Int64Attribute{
				MarkdownDescription: "The average connection rate from the incoming connection's source address. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_http_err_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_http_err_cnt": schema.Int64Attribute{
				MarkdownDescription: "The cumulative number of HTTP errors from the incoming connection's source address. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_http_err_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_http_err_rate": schema.Int64Attribute{
				MarkdownDescription: "The average rate of HTTP errors from the incoming connection's source address. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_http_req_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_http_req_cnt": schema.Int64Attribute{
				MarkdownDescription: "The cumulative number of HTTP requests from the incoming connection's source address. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_http_req_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_http_req_rate": schema.Int64Attribute{
				MarkdownDescription: "The average rate of HTTP requests from the incoming connection's source address. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_kbytes_in_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_kbytes_in": schema.Int64Attribute{
				MarkdownDescription: "The total amount of data received from the incoming connection's source address (in kilobytes). Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_kbytes_out_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_kbytes_out": schema.Int64Attribute{
				MarkdownDescription: "The total amount of data sent to the incoming connection's source address (in kilobytes). Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_port_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_port": schema.Int64Attribute{
				MarkdownDescription: "An integer value corresponding to the TCP source port of the connection on the client side, which is the port the client connected from. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_sess_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_sess_cnt": schema.Int64Attribute{
				MarkdownDescription: "The cumulative number of connections initiated from the incoming connection's source address. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_sess_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "src_sess_rate_comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_sess_rate": schema.Int64Attribute{
				MarkdownDescription: "src_sess_rate. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"nbsrv": schema.Int64Attribute{
				MarkdownDescription: "Verify the minimum number of usable servers in the named backend matches the specified value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"nbsrv_backend": schema.StringAttribute{
				MarkdownDescription: "Use the specified backend to count usable servers. Leave empty to use the current backend. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_fc_sni": schema.StringAttribute{
				MarkdownDescription: "The value of the Server Name TLS extension sent by a client matches the exact string. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_sni": schema.StringAttribute{
				MarkdownDescription: "The value of the Server Name TLS extension sent by a client matches the exact string. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_sni_sub": schema.StringAttribute{
				MarkdownDescription: "The value of the Server Name TLS extension sent by a client contains the specified string (substring match). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_sni_beg": schema.StringAttribute{
				MarkdownDescription: "The value of the Server Name TLS extension sent by a client starts with the specified string (prefix match). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_sni_end": schema.StringAttribute{
				MarkdownDescription: "The value of the Server Name TLS extension sent by a client ends with the specified string (suffix match). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_sni_reg": schema.StringAttribute{
				MarkdownDescription: "The value of the Server Name TLS extension sent by a client matches with the specified regular expression. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"custom_acl": schema.StringAttribute{
				MarkdownDescription: "Specify a HAProxy condition/ACL that is currently not supported by the GUI. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "value. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"urlparam": schema.StringAttribute{
				MarkdownDescription: "urlparam. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"query_backend": schema.StringAttribute{
				MarkdownDescription: "queryBackend. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"allowed_users": schema.SetAttribute{
				MarkdownDescription: "Allowed Users. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"allowed_groups": schema.SetAttribute{
				MarkdownDescription: "Allowed Groups. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"http_method": schema.SetAttribute{
				MarkdownDescription: "HTTP Method. One of `CONNECT`, `DELETE`, `GET`, `HEAD`, `OPTIONS`, `PATCH`, `POST`, `PUT`, `TRACE`. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
				Validators:          []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf("CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT", "TRACE"))},
			},
			"sc_bytes_in_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_bytes_in_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_bytes_out_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_bytes_out_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_clr_gpc_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_clr_gpc": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_conn_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_conn_cnt": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_conn_cur_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_conn_cur": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_conn_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_conn_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_get_gpc_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_get_gpc": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_glitch_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_glitch_cnt": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_glitch_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_glitch_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_gpc_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_gpc_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_http_err_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_http_err_cnt": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_http_err_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_http_err_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_http_fail_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_http_fail_cnt": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_http_fail_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_http_fail_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_http_req_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_http_req_cnt": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_http_req_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_http_req_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_inc_gpc_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_inc_gpc": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_sess_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_sess_cnt": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_sess_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_sess_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_get_gpc_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_get_gpc": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_get_gpt_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_get_gpt": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_glitch_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_glitch_cnt": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_glitch_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_glitch_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_gpc_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_gpc_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_http_fail_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_http_fail_cnt": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_http_fail_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_http_fail_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_inc_gpc_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_inc_gpc": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_clr_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_clr_gpc0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_clr_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_clr_gpc1": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc0_clr_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_clr_gpc0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc0_clr_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_clr_gpc1": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc1_clr_gpc_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_clr_gpc": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc1_clr_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_clr_gpc0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc1_clr_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_clr_gpc1": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc2_clr_gpc_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_clr_gpc": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc2_clr_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_clr_gpc0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc2_clr_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_clr_gpc1": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_get_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_get_gpc0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_get_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_get_gpc1": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc0_get_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_get_gpc0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc0_get_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_get_gpc1": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc1_get_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_get_gpc0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc1_get_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_get_gpc1": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc2_get_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_get_gpc0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc2_get_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_get_gpc1": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_get_gpt_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_get_gpt": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_get_gpt0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_get_gpt0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc0_get_gpt0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_get_gpt0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc1_get_gpt0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_get_gpt0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc2_get_gpt0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_get_gpt0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_gpc0_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_gpc0_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_gpc1_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_gpc1_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc0_gpc0_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_gpc0_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc0_gpc1_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_gpc1_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc1_gpc0_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_gpc0_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc1_gpc1_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_gpc1_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc2_gpc0_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_gpc0_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc2_gpc1_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_gpc1_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_inc_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_inc_gpc0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_inc_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_inc_gpc1": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc0_inc_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_inc_gpc0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc0_inc_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_inc_gpc1": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc1_inc_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_inc_gpc0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc1_inc_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_inc_gpc1": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc2_inc_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_inc_gpc0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc2_inc_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_inc_gpc1": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_clr_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_clr_gpc0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_clr_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_clr_gpc1": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_get_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_get_gpc0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_get_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_get_gpc1": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_gpc0_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_gpc0_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_gpc1_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_gpc1_rate": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_inc_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_inc_gpc0": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"src_inc_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`. Defaults to `gt`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("gt"),
				Validators:          []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_inc_gpc1": schema.Int64Attribute{
				MarkdownDescription: "Value. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"gpc_number": schema.Int64Attribute{
				MarkdownDescription: "Refers to the number of the General Purpose Counter, e.g. 0 or 1 when using gpc0 or gpc1 respectively. This value will be ignored in conditions that do not support GPC. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"gpt_number": schema.Int64Attribute{
				MarkdownDescription: "Refers to the number of the General Purpose Tag, e.g. 0 or 1 when using gpt0 or gpt1 respectively. This value will be ignored in conditions that do not support GPT. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"sc_number": schema.Int64Attribute{
				MarkdownDescription: "Refers to the number of the Stick Counter, e.g. 0 or 1 when using sc0 or sc1 respectively. This value will be ignored in conditions that do not support SC. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"table_name": schema.StringAttribute{
				MarkdownDescription: "Specifies the name of a stick-table that will be used for key lookups. This is usually the name of the backend/frontend where the stick-stable is configured. When nothing is specified, the stick-table of the current backend/frontend will be used. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"mapfile": schema.StringAttribute{
				MarkdownDescription: "Load patterns from a map file. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"converter": schema.StringAttribute{
				MarkdownDescription: "Transforms or processes the output of a sample fetch (e.g. variable, mapping, arithmetic, string manipulation) and passes the result to the next converter or final evaluation. Make more complex conditions possible, e.g. acl rate_abuse var(req.rate_limit),sub(req.request_rate) lt 0. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the condition.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func aclDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "HAProxy conditions (**Services ▸ HAProxy ▸ Settings ▸ Rules & Checks ▸ Conditions**, plugin `os-haproxy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the condition.",
				Required:            true,
			},
			"name":                          dschema.StringAttribute{MarkdownDescription: "Name to identify this condition.", Computed: true},
			"description":                   dschema.StringAttribute{MarkdownDescription: "Description for this condition.", Computed: true},
			"expression":                    dschema.StringAttribute{MarkdownDescription: "Choose the condition type. Note that a wide range of sample fetch methods is either available in layer 4 or layer 7. The syntax check will show errors when using an incompatible sample fetch method. One of `cust_hdr_beg`, `cust_hdr_end`, `cust_hdr`, `cust_hdr_reg`, `cust_hdr_sub`, `hdr_beg`, `hdr_end`, `hdr`, `hdr_reg`, `hdr_sub`, `http_auth`, `http_method`, `nbsrv`, `path_beg`, `path_dir`, `path_end`, `path`, `path_reg`, `path_sub`, `quic_enabled`, `traffic_is_http`, `traffic_is_ssl`, `sc_bytes_in_rate`, `sc_bytes_out_rate`, `sc_clr_gpc`, `sc_clr_gpc0`, `sc_clr_gpc1`, `sc0_clr_gpc0`, `sc0_clr_gpc1`, `sc1_clr_gpc`, `sc1_clr_gpc0`, `sc1_clr_gpc1`, `sc2_clr_gpc`, `sc2_clr_gpc0`, `sc2_clr_gpc1`, `sc_conn_cnt`, `sc_conn_cur`, `sc_conn_rate`, `sc_get_gpc`, `sc_get_gpc0`, `sc_get_gpc1`, `sc0_get_gpc0`, `sc0_get_gpc1`, `sc1_get_gpc0`, `sc1_get_gpc1`, `sc2_get_gpc0`, `sc2_get_gpc1`, `sc_get_gpt`, `sc_get_gpt0`, `sc0_get_gpt0`, `sc1_get_gpt0`, `sc2_get_gpt0`, `sc_glitch_cnt`, `sc_glitch_rate`, `sc_gpc_rate`, `sc_gpc0_rate`, `sc_gpc1_rate`, `sc0_gpc0_rate`, `sc0_gpc1_rate`, `sc1_gpc0_rate`, `sc1_gpc1_rate`, `sc2_gpc0_rate`, `sc2_gpc1_rate`, `sc_http_err_cnt`, `sc_http_err_rate`, `sc_http_fail_cnt`, `sc_http_fail_rate`, `sc_http_req_cnt`, `sc_http_req_rate`, `sc_inc_gpc`, `sc_inc_gpc0`, `sc_inc_gpc1`, `sc0_inc_gpc0`, `sc0_inc_gpc1`, `sc1_inc_gpc0`, `sc1_inc_gpc1`, `sc2_inc_gpc0`, `sc2_inc_gpc1`, `sc_sess_cnt`, `sc_sess_rate`, `src`, `src_bytes_in_rate`, `src_bytes_out_rate`, `src_clr_gpc`, `src_clr_gpc0`, `src_clr_gpc1`, `src_conn_cnt`, `src_conn_cur`, `src_conn_rate`, `src_get_gpc`, `src_get_gpc0`, `src_get_gpc1`, `src_get_gpt`, `src_glitch_cnt`, `src_glitch_rate`, `src_gpc_rate`, `src_gpc0_rate`, `src_gpc1_rate`, `src_http_err_cnt`, `src_http_err_rate`, `src_http_fail_cnt`, `src_http_fail_rate`, `src_http_req_cnt`, `src_http_req_rate`, `src_inc_gpc`, `src_inc_gpc0`, `src_inc_gpc1`, `src_is_local`, `src_kbytes_in`, `src_kbytes_out`, `src_port`, `src_sess_cnt`, `src_sess_rate`, `ssl_c_ca_commonname`, `ssl_c_verify_code`, `ssl_c_verify`, `ssl_fc_sni`, `ssl_fc`, `ssl_hello_type`, `ssl_sni_beg`, `ssl_sni_end`, `ssl_sni_reg`, `ssl_sni`, `ssl_sni_sub`, `stopping`, `url_param`, `var`, `wait_end`, `custom_acl`.", Computed: true},
			"negate":                        dschema.BoolAttribute{MarkdownDescription: "Use this to invert the meaning of the expression.", Computed: true},
			"case_sensitive":                dschema.BoolAttribute{MarkdownDescription: "Enable to make the condition case-sensitive (if applicable). By default all conditions are not case-sensitive.", Computed: true},
			"hdr_beg":                       dschema.StringAttribute{MarkdownDescription: "HTTP host header starts with string (prefix match).", Computed: true},
			"hdr_end":                       dschema.StringAttribute{MarkdownDescription: "HTTP host header ends with string (suffix match).", Computed: true},
			"hdr":                           dschema.StringAttribute{MarkdownDescription: "HTTP host header matches exact string.", Computed: true},
			"hdr_reg":                       dschema.StringAttribute{MarkdownDescription: "HTTP host header matches regular expression.", Computed: true},
			"hdr_sub":                       dschema.StringAttribute{MarkdownDescription: "HTTP host header contains string (substring match).", Computed: true},
			"path_beg":                      dschema.StringAttribute{MarkdownDescription: "HTTP request URL path starts with string (prefix match).", Computed: true},
			"path_end":                      dschema.StringAttribute{MarkdownDescription: "HTTP request URL path ends with string (suffix match).", Computed: true},
			"path":                          dschema.StringAttribute{MarkdownDescription: "HTTP request URL path matches exact string.", Computed: true},
			"path_reg":                      dschema.StringAttribute{MarkdownDescription: "HTTP request URL path matches regular expression.", Computed: true},
			"path_dir":                      dschema.StringAttribute{MarkdownDescription: "HTTP request URL path contains directory (subdir match).", Computed: true},
			"path_sub":                      dschema.StringAttribute{MarkdownDescription: "HTTP request URL path contains string (substring match).", Computed: true},
			"cust_hdr_beg_name":             dschema.StringAttribute{MarkdownDescription: "The name of the HTTP Header.", Computed: true},
			"cust_hdr_beg":                  dschema.StringAttribute{MarkdownDescription: "HTTP Header starts with string (prefix match).", Computed: true},
			"cust_hdr_end_name":             dschema.StringAttribute{MarkdownDescription: "The name of the HTTP Header.", Computed: true},
			"cust_hdr_end":                  dschema.StringAttribute{MarkdownDescription: "HTTP Header ends with string (suffix match).", Computed: true},
			"cust_hdr_name":                 dschema.StringAttribute{MarkdownDescription: "The name of the HTTP Header.", Computed: true},
			"cust_hdr":                      dschema.StringAttribute{MarkdownDescription: "HTTP Header matches exact string.", Computed: true},
			"cust_hdr_reg_name":             dschema.StringAttribute{MarkdownDescription: "The name of the HTTP Header.", Computed: true},
			"cust_hdr_reg":                  dschema.StringAttribute{MarkdownDescription: "HTTP Header matches regular expression.", Computed: true},
			"cust_hdr_sub_name":             dschema.StringAttribute{MarkdownDescription: "The name of the HTTP Header.", Computed: true},
			"cust_hdr_sub":                  dschema.StringAttribute{MarkdownDescription: "HTTP Header contains string (substring match).", Computed: true},
			"url_param":                     dschema.StringAttribute{MarkdownDescription: "Specify the URL parameter to be checked for the value specified below.", Computed: true},
			"url_param_value":               dschema.StringAttribute{MarkdownDescription: "Specify the value for the URL parameter.", Computed: true},
			"var":                           dschema.StringAttribute{MarkdownDescription: "The name of a variable, followed by an optional default value, e.g. (req.rate_limit) or (req.rate_limit,10).", Computed: true},
			"var_value":                     dschema.StringAttribute{MarkdownDescription: "Value.", Computed: true},
			"var_comparison":                dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"ssl_c_verify_code":             dschema.Int64Attribute{MarkdownDescription: "Specify the SSL/TLS error ID that should be checked for the incoming connection. Please refer to your SSL library's documentation for an exhaustive list of error codes.", Computed: true},
			"ssl_c_ca_commonname":           dschema.StringAttribute{MarkdownDescription: "Verify the CA Common-Name of the certificate presented by the client against the specified string.", Computed: true},
			"ssl_hello_type":                dschema.StringAttribute{MarkdownDescription: "An integer value containing the type of the SSL hello message found in the request buffer if the buffer contains data that parse as a complete SSL (v3 or superior) client hello message. One of `x0`, `x1`, `x2`.", Computed: true},
			"src":                           dschema.StringAttribute{MarkdownDescription: "Verify the source IPv4 address of the client of the session matches the specified IPv4 or IPv6 address.", Computed: true},
			"src_bytes_in_rate_comparison":  dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_bytes_in_rate":             dschema.Int64Attribute{MarkdownDescription: "The average bytes rate from the incoming connection's source address.", Computed: true},
			"src_bytes_out_rate_comparison": dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_bytes_out_rate":            dschema.Int64Attribute{MarkdownDescription: "The average bytes rate to the incoming connection's source address.", Computed: true},
			"src_conn_cnt_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_conn_cnt":                  dschema.Int64Attribute{MarkdownDescription: "The cumulative number of connections initiated from the current incoming connection's source address.", Computed: true},
			"src_conn_cur_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_conn_cur":                  dschema.Int64Attribute{MarkdownDescription: "The current amount of concurrent connections initiated from the current incoming connection's source address.", Computed: true},
			"src_conn_rate_comparison":      dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_conn_rate":                 dschema.Int64Attribute{MarkdownDescription: "The average connection rate from the incoming connection's source address.", Computed: true},
			"src_http_err_cnt_comparison":   dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_http_err_cnt":              dschema.Int64Attribute{MarkdownDescription: "The cumulative number of HTTP errors from the incoming connection's source address.", Computed: true},
			"src_http_err_rate_comparison":  dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_http_err_rate":             dschema.Int64Attribute{MarkdownDescription: "The average rate of HTTP errors from the incoming connection's source address.", Computed: true},
			"src_http_req_cnt_comparison":   dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_http_req_cnt":              dschema.Int64Attribute{MarkdownDescription: "The cumulative number of HTTP requests from the incoming connection's source address.", Computed: true},
			"src_http_req_rate_comparison":  dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_http_req_rate":             dschema.Int64Attribute{MarkdownDescription: "The average rate of HTTP requests from the incoming connection's source address.", Computed: true},
			"src_kbytes_in_comparison":      dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_kbytes_in":                 dschema.Int64Attribute{MarkdownDescription: "The total amount of data received from the incoming connection's source address (in kilobytes).", Computed: true},
			"src_kbytes_out_comparison":     dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_kbytes_out":                dschema.Int64Attribute{MarkdownDescription: "The total amount of data sent to the incoming connection's source address (in kilobytes).", Computed: true},
			"src_port_comparison":           dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_port":                      dschema.Int64Attribute{MarkdownDescription: "An integer value corresponding to the TCP source port of the connection on the client side, which is the port the client connected from.", Computed: true},
			"src_sess_cnt_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_sess_cnt":                  dschema.Int64Attribute{MarkdownDescription: "The cumulative number of connections initiated from the incoming connection's source address.", Computed: true},
			"src_sess_rate_comparison":      dschema.StringAttribute{MarkdownDescription: "src_sess_rate_comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_sess_rate":                 dschema.Int64Attribute{MarkdownDescription: "src_sess_rate.", Computed: true},
			"nbsrv":                         dschema.Int64Attribute{MarkdownDescription: "Verify the minimum number of usable servers in the named backend matches the specified value.", Computed: true},
			"nbsrv_backend":                 dschema.StringAttribute{MarkdownDescription: "Use the specified backend to count usable servers. Leave empty to use the current backend.", Computed: true},
			"ssl_fc_sni":                    dschema.StringAttribute{MarkdownDescription: "The value of the Server Name TLS extension sent by a client matches the exact string.", Computed: true},
			"ssl_sni":                       dschema.StringAttribute{MarkdownDescription: "The value of the Server Name TLS extension sent by a client matches the exact string.", Computed: true},
			"ssl_sni_sub":                   dschema.StringAttribute{MarkdownDescription: "The value of the Server Name TLS extension sent by a client contains the specified string (substring match).", Computed: true},
			"ssl_sni_beg":                   dschema.StringAttribute{MarkdownDescription: "The value of the Server Name TLS extension sent by a client starts with the specified string (prefix match).", Computed: true},
			"ssl_sni_end":                   dschema.StringAttribute{MarkdownDescription: "The value of the Server Name TLS extension sent by a client ends with the specified string (suffix match).", Computed: true},
			"ssl_sni_reg":                   dschema.StringAttribute{MarkdownDescription: "The value of the Server Name TLS extension sent by a client matches with the specified regular expression.", Computed: true},
			"custom_acl":                    dschema.StringAttribute{MarkdownDescription: "Specify a HAProxy condition/ACL that is currently not supported by the GUI.", Computed: true},
			"value":                         dschema.StringAttribute{MarkdownDescription: "value.", Computed: true},
			"urlparam":                      dschema.StringAttribute{MarkdownDescription: "urlparam.", Computed: true},
			"query_backend":                 dschema.StringAttribute{MarkdownDescription: "queryBackend.", Computed: true},
			"allowed_users":                 dschema.SetAttribute{MarkdownDescription: "Allowed Users.", Computed: true, ElementType: types.StringType},
			"allowed_groups":                dschema.SetAttribute{MarkdownDescription: "Allowed Groups.", Computed: true, ElementType: types.StringType},
			"http_method":                   dschema.SetAttribute{MarkdownDescription: "HTTP Method. One of `CONNECT`, `DELETE`, `GET`, `HEAD`, `OPTIONS`, `PATCH`, `POST`, `PUT`, `TRACE`.", Computed: true, ElementType: types.StringType},
			"sc_bytes_in_rate_comparison":   dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_bytes_in_rate":              dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_bytes_out_rate_comparison":  dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_bytes_out_rate":             dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_clr_gpc_comparison":         dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_clr_gpc":                    dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_conn_cnt_comparison":        dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_conn_cnt":                   dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_conn_cur_comparison":        dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_conn_cur":                   dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_conn_rate_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_conn_rate":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_get_gpc_comparison":         dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_get_gpc":                    dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_glitch_cnt_comparison":      dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_glitch_cnt":                 dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_glitch_rate_comparison":     dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_glitch_rate":                dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_gpc_rate_comparison":        dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_gpc_rate":                   dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_http_err_cnt_comparison":    dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_http_err_cnt":               dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_http_err_rate_comparison":   dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_http_err_rate":              dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_http_fail_cnt_comparison":   dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_http_fail_cnt":              dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_http_fail_rate_comparison":  dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_http_fail_rate":             dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_http_req_cnt_comparison":    dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_http_req_cnt":               dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_http_req_rate_comparison":   dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_http_req_rate":              dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_inc_gpc_comparison":         dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_inc_gpc":                    dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_sess_cnt_comparison":        dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_sess_cnt":                   dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_sess_rate_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_sess_rate":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"src_get_gpc_comparison":        dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_get_gpc":                   dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"src_get_gpt_comparison":        dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_get_gpt":                   dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"src_glitch_cnt_comparison":     dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_glitch_cnt":                dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"src_glitch_rate_comparison":    dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_glitch_rate":               dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"src_gpc_rate_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_gpc_rate":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"src_http_fail_cnt_comparison":  dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_http_fail_cnt":             dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"src_http_fail_rate_comparison": dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_http_fail_rate":            dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"src_inc_gpc_comparison":        dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_inc_gpc":                   dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_clr_gpc0_comparison":        dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_clr_gpc0":                   dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_clr_gpc1_comparison":        dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_clr_gpc1":                   dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc0_clr_gpc0_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc0_clr_gpc0":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc0_clr_gpc1_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc0_clr_gpc1":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc1_clr_gpc_comparison":        dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc1_clr_gpc":                   dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc1_clr_gpc0_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc1_clr_gpc0":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc1_clr_gpc1_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc1_clr_gpc1":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc2_clr_gpc_comparison":        dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc2_clr_gpc":                   dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc2_clr_gpc0_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc2_clr_gpc0":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc2_clr_gpc1_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc2_clr_gpc1":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_get_gpc0_comparison":        dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_get_gpc0":                   dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_get_gpc1_comparison":        dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_get_gpc1":                   dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc0_get_gpc0_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc0_get_gpc0":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc0_get_gpc1_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc0_get_gpc1":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc1_get_gpc0_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc1_get_gpc0":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc1_get_gpc1_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc1_get_gpc1":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc2_get_gpc0_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc2_get_gpc0":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc2_get_gpc1_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc2_get_gpc1":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_get_gpt_comparison":         dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_get_gpt":                    dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_get_gpt0_comparison":        dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_get_gpt0":                   dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc0_get_gpt0_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc0_get_gpt0":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc1_get_gpt0_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc1_get_gpt0":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc2_get_gpt0_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc2_get_gpt0":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_gpc0_rate_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_gpc0_rate":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_gpc1_rate_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_gpc1_rate":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc0_gpc0_rate_comparison":      dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc0_gpc0_rate":                 dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc0_gpc1_rate_comparison":      dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc0_gpc1_rate":                 dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc1_gpc0_rate_comparison":      dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc1_gpc0_rate":                 dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc1_gpc1_rate_comparison":      dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc1_gpc1_rate":                 dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc2_gpc0_rate_comparison":      dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc2_gpc0_rate":                 dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc2_gpc1_rate_comparison":      dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc2_gpc1_rate":                 dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_inc_gpc0_comparison":        dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_inc_gpc0":                   dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc_inc_gpc1_comparison":        dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc_inc_gpc1":                   dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc0_inc_gpc0_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc0_inc_gpc0":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc0_inc_gpc1_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc0_inc_gpc1":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc1_inc_gpc0_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc1_inc_gpc0":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc1_inc_gpc1_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc1_inc_gpc1":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc2_inc_gpc0_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc2_inc_gpc0":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"sc2_inc_gpc1_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"sc2_inc_gpc1":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"src_clr_gpc0_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_clr_gpc0":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"src_clr_gpc1_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_clr_gpc1":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"src_get_gpc0_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_get_gpc0":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"src_get_gpc1_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_get_gpc1":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"src_gpc0_rate_comparison":      dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_gpc0_rate":                 dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"src_gpc1_rate_comparison":      dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_gpc1_rate":                 dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"src_inc_gpc0_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_inc_gpc0":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"src_inc_gpc1_comparison":       dschema.StringAttribute{MarkdownDescription: "Comparison. One of `gt`, `ge`, `eq`, `lt`, `le`.", Computed: true},
			"src_inc_gpc1":                  dschema.Int64Attribute{MarkdownDescription: "Value.", Computed: true},
			"gpc_number":                    dschema.Int64Attribute{MarkdownDescription: "Refers to the number of the General Purpose Counter, e.g. 0 or 1 when using gpc0 or gpc1 respectively. This value will be ignored in conditions that do not support GPC.", Computed: true},
			"gpt_number":                    dschema.Int64Attribute{MarkdownDescription: "Refers to the number of the General Purpose Tag, e.g. 0 or 1 when using gpt0 or gpt1 respectively. This value will be ignored in conditions that do not support GPT.", Computed: true},
			"sc_number":                     dschema.Int64Attribute{MarkdownDescription: "Refers to the number of the Stick Counter, e.g. 0 or 1 when using sc0 or sc1 respectively. This value will be ignored in conditions that do not support SC.", Computed: true},
			"table_name":                    dschema.StringAttribute{MarkdownDescription: "Specifies the name of a stick-table that will be used for key lookups. This is usually the name of the backend/frontend where the stick-stable is configured. When nothing is specified, the stick-table of the current backend/frontend will be used.", Computed: true},
			"mapfile":                       dschema.StringAttribute{MarkdownDescription: "Load patterns from a map file.", Computed: true},
			"converter":                     dschema.StringAttribute{MarkdownDescription: "Transforms or processes the output of a sample fetch (e.g. variable, mapping, arithmetic, string manipulation) and passes the result to the next converter or final evaluation. Make more complex conditions possible, e.g. acl rate_abuse var(req.rate_limit),sub(req.request_rate) lt 0.", Computed: true},
		},
	}
}

func convertAclSchemaToStruct(d *aclResourceModel) (*haproxy.Acl, error) {
	return &haproxy.Acl{
		Name:                      d.Name.ValueString(),
		Description:               d.Description.ValueString(),
		Expression:                api.SelectedMap(d.Expression.ValueString()),
		Negate:                    tools.BoolToString(d.Negate.ValueBool()),
		CaseSensitive:             tools.BoolToString(d.CaseSensitive.ValueBool()),
		HdrBeg:                    d.HdrBeg.ValueString(),
		HdrEnd:                    d.HdrEnd.ValueString(),
		Hdr:                       d.Hdr.ValueString(),
		HdrReg:                    d.HdrReg.ValueString(),
		HdrSub:                    d.HdrSub.ValueString(),
		PathBeg:                   d.PathBeg.ValueString(),
		PathEnd:                   d.PathEnd.ValueString(),
		Path:                      d.Path.ValueString(),
		PathReg:                   d.PathReg.ValueString(),
		PathDir:                   d.PathDir.ValueString(),
		PathSub:                   d.PathSub.ValueString(),
		CustHdrBegName:            d.CustHdrBegName.ValueString(),
		CustHdrBeg:                d.CustHdrBeg.ValueString(),
		CustHdrEndName:            d.CustHdrEndName.ValueString(),
		CustHdrEnd:                d.CustHdrEnd.ValueString(),
		CustHdrName:               d.CustHdrName.ValueString(),
		CustHdr:                   d.CustHdr.ValueString(),
		CustHdrRegName:            d.CustHdrRegName.ValueString(),
		CustHdrReg:                d.CustHdrReg.ValueString(),
		CustHdrSubName:            d.CustHdrSubName.ValueString(),
		CustHdrSub:                d.CustHdrSub.ValueString(),
		UrlParam:                  d.URLParam.ValueString(),
		UrlParamValue:             d.URLParamValue.ValueString(),
		Var:                       d.Var.ValueString(),
		VarValue:                  d.VarValue.ValueString(),
		VarComparison:             api.SelectedMap(d.VarComparison.ValueString()),
		SslCVerifyCode:            tools.Int64ToStringNegative(d.SslCVerifyCode.ValueInt64()),
		SslCCaCommonname:          d.SslCCaCommonname.ValueString(),
		SslHelloType:              api.SelectedMap(d.SslHelloType.ValueString()),
		Src:                       d.Src.ValueString(),
		SrcBytesInRateComparison:  api.SelectedMap(d.SrcBytesInRateComparison.ValueString()),
		SrcBytesInRate:            tools.Int64ToStringNegative(d.SrcBytesInRate.ValueInt64()),
		SrcBytesOutRateComparison: api.SelectedMap(d.SrcBytesOutRateComparison.ValueString()),
		SrcBytesOutRate:           tools.Int64ToStringNegative(d.SrcBytesOutRate.ValueInt64()),
		SrcConnCntComparison:      api.SelectedMap(d.SrcConnCntComparison.ValueString()),
		SrcConnCnt:                tools.Int64ToStringNegative(d.SrcConnCnt.ValueInt64()),
		SrcConnCurComparison:      api.SelectedMap(d.SrcConnCurComparison.ValueString()),
		SrcConnCur:                tools.Int64ToStringNegative(d.SrcConnCur.ValueInt64()),
		SrcConnRateComparison:     api.SelectedMap(d.SrcConnRateComparison.ValueString()),
		SrcConnRate:               tools.Int64ToStringNegative(d.SrcConnRate.ValueInt64()),
		SrcHttpErrCntComparison:   api.SelectedMap(d.SrcHttpErrCntComparison.ValueString()),
		SrcHttpErrCnt:             tools.Int64ToStringNegative(d.SrcHttpErrCnt.ValueInt64()),
		SrcHttpErrRateComparison:  api.SelectedMap(d.SrcHttpErrRateComparison.ValueString()),
		SrcHttpErrRate:            tools.Int64ToStringNegative(d.SrcHttpErrRate.ValueInt64()),
		SrcHttpReqCntComparison:   api.SelectedMap(d.SrcHttpReqCntComparison.ValueString()),
		SrcHttpReqCnt:             tools.Int64ToStringNegative(d.SrcHttpReqCnt.ValueInt64()),
		SrcHttpReqRateComparison:  api.SelectedMap(d.SrcHttpReqRateComparison.ValueString()),
		SrcHttpReqRate:            tools.Int64ToStringNegative(d.SrcHttpReqRate.ValueInt64()),
		SrcKbytesInComparison:     api.SelectedMap(d.SrcKbytesInComparison.ValueString()),
		SrcKbytesIn:               tools.Int64ToStringNegative(d.SrcKbytesIn.ValueInt64()),
		SrcKbytesOutComparison:    api.SelectedMap(d.SrcKbytesOutComparison.ValueString()),
		SrcKbytesOut:              tools.Int64ToStringNegative(d.SrcKbytesOut.ValueInt64()),
		SrcPortComparison:         api.SelectedMap(d.SrcPortComparison.ValueString()),
		SrcPort:                   tools.Int64ToStringNegative(d.SrcPort.ValueInt64()),
		SrcSessCntComparison:      api.SelectedMap(d.SrcSessCntComparison.ValueString()),
		SrcSessCnt:                tools.Int64ToStringNegative(d.SrcSessCnt.ValueInt64()),
		SrcSessRateComparison:     api.SelectedMap(d.SrcSessRateComparison.ValueString()),
		SrcSessRate:               tools.Int64ToStringNegative(d.SrcSessRate.ValueInt64()),
		Nbsrv:                     tools.Int64ToStringNegative(d.Nbsrv.ValueInt64()),
		NbsrvBackend:              api.SelectedMap(d.NbsrvBackend.ValueString()),
		SslFcSni:                  d.SslFcSni.ValueString(),
		SslSni:                    d.SslSni.ValueString(),
		SslSniSub:                 d.SslSniSub.ValueString(),
		SslSniBeg:                 d.SslSniBeg.ValueString(),
		SslSniEnd:                 d.SslSniEnd.ValueString(),
		SslSniReg:                 d.SslSniReg.ValueString(),
		CustomAcl:                 d.CustomAcl.ValueString(),
		Value:                     d.Value.ValueString(),
		Urlparam:                  d.Urlparam.ValueString(),
		QueryBackend:              api.SelectedMap(d.QueryBackend.ValueString()),
		AllowedUsers:              api.SelectedMapList(setToSlice(d.AllowedUsers)),
		AllowedGroups:             api.SelectedMapList(setToSlice(d.AllowedGroups)),
		HttpMethod:                api.SelectedMapList(setToSlice(d.HttpMethod)),
		ScBytesInRateComparison:   api.SelectedMap(d.ScBytesInRateComparison.ValueString()),
		ScBytesInRate:             tools.Int64ToStringNegative(d.ScBytesInRate.ValueInt64()),
		ScBytesOutRateComparison:  api.SelectedMap(d.ScBytesOutRateComparison.ValueString()),
		ScBytesOutRate:            tools.Int64ToStringNegative(d.ScBytesOutRate.ValueInt64()),
		ScClrGpcComparison:        api.SelectedMap(d.ScClrGpcComparison.ValueString()),
		ScClrGpc:                  tools.Int64ToStringNegative(d.ScClrGpc.ValueInt64()),
		ScConnCntComparison:       api.SelectedMap(d.ScConnCntComparison.ValueString()),
		ScConnCnt:                 tools.Int64ToStringNegative(d.ScConnCnt.ValueInt64()),
		ScConnCurComparison:       api.SelectedMap(d.ScConnCurComparison.ValueString()),
		ScConnCur:                 tools.Int64ToStringNegative(d.ScConnCur.ValueInt64()),
		ScConnRateComparison:      api.SelectedMap(d.ScConnRateComparison.ValueString()),
		ScConnRate:                tools.Int64ToStringNegative(d.ScConnRate.ValueInt64()),
		ScGetGpcComparison:        api.SelectedMap(d.ScGetGpcComparison.ValueString()),
		ScGetGpc:                  tools.Int64ToStringNegative(d.ScGetGpc.ValueInt64()),
		ScGlitchCntComparison:     api.SelectedMap(d.ScGlitchCntComparison.ValueString()),
		ScGlitchCnt:               tools.Int64ToStringNegative(d.ScGlitchCnt.ValueInt64()),
		ScGlitchRateComparison:    api.SelectedMap(d.ScGlitchRateComparison.ValueString()),
		ScGlitchRate:              tools.Int64ToStringNegative(d.ScGlitchRate.ValueInt64()),
		ScGpcRateComparison:       api.SelectedMap(d.ScGpcRateComparison.ValueString()),
		ScGpcRate:                 tools.Int64ToStringNegative(d.ScGpcRate.ValueInt64()),
		ScHttpErrCntComparison:    api.SelectedMap(d.ScHttpErrCntComparison.ValueString()),
		ScHttpErrCnt:              tools.Int64ToStringNegative(d.ScHttpErrCnt.ValueInt64()),
		ScHttpErrRateComparison:   api.SelectedMap(d.ScHttpErrRateComparison.ValueString()),
		ScHttpErrRate:             tools.Int64ToStringNegative(d.ScHttpErrRate.ValueInt64()),
		ScHttpFailCntComparison:   api.SelectedMap(d.ScHttpFailCntComparison.ValueString()),
		ScHttpFailCnt:             tools.Int64ToStringNegative(d.ScHttpFailCnt.ValueInt64()),
		ScHttpFailRateComparison:  api.SelectedMap(d.ScHttpFailRateComparison.ValueString()),
		ScHttpFailRate:            tools.Int64ToStringNegative(d.ScHttpFailRate.ValueInt64()),
		ScHttpReqCntComparison:    api.SelectedMap(d.ScHttpReqCntComparison.ValueString()),
		ScHttpReqCnt:              tools.Int64ToStringNegative(d.ScHttpReqCnt.ValueInt64()),
		ScHttpReqRateComparison:   api.SelectedMap(d.ScHttpReqRateComparison.ValueString()),
		ScHttpReqRate:             tools.Int64ToStringNegative(d.ScHttpReqRate.ValueInt64()),
		ScIncGpcComparison:        api.SelectedMap(d.ScIncGpcComparison.ValueString()),
		ScIncGpc:                  tools.Int64ToStringNegative(d.ScIncGpc.ValueInt64()),
		ScSessCntComparison:       api.SelectedMap(d.ScSessCntComparison.ValueString()),
		ScSessCnt:                 tools.Int64ToStringNegative(d.ScSessCnt.ValueInt64()),
		ScSessRateComparison:      api.SelectedMap(d.ScSessRateComparison.ValueString()),
		ScSessRate:                tools.Int64ToStringNegative(d.ScSessRate.ValueInt64()),
		SrcGetGpcComparison:       api.SelectedMap(d.SrcGetGpcComparison.ValueString()),
		SrcGetGpc:                 tools.Int64ToStringNegative(d.SrcGetGpc.ValueInt64()),
		SrcGetGptComparison:       api.SelectedMap(d.SrcGetGptComparison.ValueString()),
		SrcGetGpt:                 tools.Int64ToStringNegative(d.SrcGetGpt.ValueInt64()),
		SrcGlitchCntComparison:    api.SelectedMap(d.SrcGlitchCntComparison.ValueString()),
		SrcGlitchCnt:              tools.Int64ToStringNegative(d.SrcGlitchCnt.ValueInt64()),
		SrcGlitchRateComparison:   api.SelectedMap(d.SrcGlitchRateComparison.ValueString()),
		SrcGlitchRate:             tools.Int64ToStringNegative(d.SrcGlitchRate.ValueInt64()),
		SrcGpcRateComparison:      api.SelectedMap(d.SrcGpcRateComparison.ValueString()),
		SrcGpcRate:                tools.Int64ToStringNegative(d.SrcGpcRate.ValueInt64()),
		SrcHttpFailCntComparison:  api.SelectedMap(d.SrcHttpFailCntComparison.ValueString()),
		SrcHttpFailCnt:            tools.Int64ToStringNegative(d.SrcHttpFailCnt.ValueInt64()),
		SrcHttpFailRateComparison: api.SelectedMap(d.SrcHttpFailRateComparison.ValueString()),
		SrcHttpFailRate:           tools.Int64ToStringNegative(d.SrcHttpFailRate.ValueInt64()),
		SrcIncGpcComparison:       api.SelectedMap(d.SrcIncGpcComparison.ValueString()),
		SrcIncGpc:                 tools.Int64ToStringNegative(d.SrcIncGpc.ValueInt64()),
		ScClrGpc0Comparison:       api.SelectedMap(d.ScClrGpc0Comparison.ValueString()),
		ScClrGpc0:                 tools.Int64ToStringNegative(d.ScClrGpc0.ValueInt64()),
		ScClrGpc1Comparison:       api.SelectedMap(d.ScClrGpc1Comparison.ValueString()),
		ScClrGpc1:                 tools.Int64ToStringNegative(d.ScClrGpc1.ValueInt64()),
		Sc0ClrGpc0Comparison:      api.SelectedMap(d.Sc0ClrGpc0Comparison.ValueString()),
		Sc0ClrGpc0:                tools.Int64ToStringNegative(d.Sc0ClrGpc0.ValueInt64()),
		Sc0ClrGpc1Comparison:      api.SelectedMap(d.Sc0ClrGpc1Comparison.ValueString()),
		Sc0ClrGpc1:                tools.Int64ToStringNegative(d.Sc0ClrGpc1.ValueInt64()),
		Sc1ClrGpcComparison:       api.SelectedMap(d.Sc1ClrGpcComparison.ValueString()),
		Sc1ClrGpc:                 tools.Int64ToStringNegative(d.Sc1ClrGpc.ValueInt64()),
		Sc1ClrGpc0Comparison:      api.SelectedMap(d.Sc1ClrGpc0Comparison.ValueString()),
		Sc1ClrGpc0:                tools.Int64ToStringNegative(d.Sc1ClrGpc0.ValueInt64()),
		Sc1ClrGpc1Comparison:      api.SelectedMap(d.Sc1ClrGpc1Comparison.ValueString()),
		Sc1ClrGpc1:                tools.Int64ToStringNegative(d.Sc1ClrGpc1.ValueInt64()),
		Sc2ClrGpcComparison:       api.SelectedMap(d.Sc2ClrGpcComparison.ValueString()),
		Sc2ClrGpc:                 tools.Int64ToStringNegative(d.Sc2ClrGpc.ValueInt64()),
		Sc2ClrGpc0Comparison:      api.SelectedMap(d.Sc2ClrGpc0Comparison.ValueString()),
		Sc2ClrGpc0:                tools.Int64ToStringNegative(d.Sc2ClrGpc0.ValueInt64()),
		Sc2ClrGpc1Comparison:      api.SelectedMap(d.Sc2ClrGpc1Comparison.ValueString()),
		Sc2ClrGpc1:                tools.Int64ToStringNegative(d.Sc2ClrGpc1.ValueInt64()),
		ScGetGpc0Comparison:       api.SelectedMap(d.ScGetGpc0Comparison.ValueString()),
		ScGetGpc0:                 tools.Int64ToStringNegative(d.ScGetGpc0.ValueInt64()),
		ScGetGpc1Comparison:       api.SelectedMap(d.ScGetGpc1Comparison.ValueString()),
		ScGetGpc1:                 tools.Int64ToStringNegative(d.ScGetGpc1.ValueInt64()),
		Sc0GetGpc0Comparison:      api.SelectedMap(d.Sc0GetGpc0Comparison.ValueString()),
		Sc0GetGpc0:                tools.Int64ToStringNegative(d.Sc0GetGpc0.ValueInt64()),
		Sc0GetGpc1Comparison:      api.SelectedMap(d.Sc0GetGpc1Comparison.ValueString()),
		Sc0GetGpc1:                tools.Int64ToStringNegative(d.Sc0GetGpc1.ValueInt64()),
		Sc1GetGpc0Comparison:      api.SelectedMap(d.Sc1GetGpc0Comparison.ValueString()),
		Sc1GetGpc0:                tools.Int64ToStringNegative(d.Sc1GetGpc0.ValueInt64()),
		Sc1GetGpc1Comparison:      api.SelectedMap(d.Sc1GetGpc1Comparison.ValueString()),
		Sc1GetGpc1:                tools.Int64ToStringNegative(d.Sc1GetGpc1.ValueInt64()),
		Sc2GetGpc0Comparison:      api.SelectedMap(d.Sc2GetGpc0Comparison.ValueString()),
		Sc2GetGpc0:                tools.Int64ToStringNegative(d.Sc2GetGpc0.ValueInt64()),
		Sc2GetGpc1Comparison:      api.SelectedMap(d.Sc2GetGpc1Comparison.ValueString()),
		Sc2GetGpc1:                tools.Int64ToStringNegative(d.Sc2GetGpc1.ValueInt64()),
		ScGetGptComparison:        api.SelectedMap(d.ScGetGptComparison.ValueString()),
		ScGetGpt:                  tools.Int64ToStringNegative(d.ScGetGpt.ValueInt64()),
		ScGetGpt0Comparison:       api.SelectedMap(d.ScGetGpt0Comparison.ValueString()),
		ScGetGpt0:                 tools.Int64ToStringNegative(d.ScGetGpt0.ValueInt64()),
		Sc0GetGpt0Comparison:      api.SelectedMap(d.Sc0GetGpt0Comparison.ValueString()),
		Sc0GetGpt0:                tools.Int64ToStringNegative(d.Sc0GetGpt0.ValueInt64()),
		Sc1GetGpt0Comparison:      api.SelectedMap(d.Sc1GetGpt0Comparison.ValueString()),
		Sc1GetGpt0:                tools.Int64ToStringNegative(d.Sc1GetGpt0.ValueInt64()),
		Sc2GetGpt0Comparison:      api.SelectedMap(d.Sc2GetGpt0Comparison.ValueString()),
		Sc2GetGpt0:                tools.Int64ToStringNegative(d.Sc2GetGpt0.ValueInt64()),
		ScGpc0RateComparison:      api.SelectedMap(d.ScGpc0RateComparison.ValueString()),
		ScGpc0Rate:                tools.Int64ToStringNegative(d.ScGpc0Rate.ValueInt64()),
		ScGpc1RateComparison:      api.SelectedMap(d.ScGpc1RateComparison.ValueString()),
		ScGpc1Rate:                tools.Int64ToStringNegative(d.ScGpc1Rate.ValueInt64()),
		Sc0Gpc0RateComparison:     api.SelectedMap(d.Sc0Gpc0RateComparison.ValueString()),
		Sc0Gpc0Rate:               tools.Int64ToStringNegative(d.Sc0Gpc0Rate.ValueInt64()),
		Sc0Gpc1RateComparison:     api.SelectedMap(d.Sc0Gpc1RateComparison.ValueString()),
		Sc0Gpc1Rate:               tools.Int64ToStringNegative(d.Sc0Gpc1Rate.ValueInt64()),
		Sc1Gpc0RateComparison:     api.SelectedMap(d.Sc1Gpc0RateComparison.ValueString()),
		Sc1Gpc0Rate:               tools.Int64ToStringNegative(d.Sc1Gpc0Rate.ValueInt64()),
		Sc1Gpc1RateComparison:     api.SelectedMap(d.Sc1Gpc1RateComparison.ValueString()),
		Sc1Gpc1Rate:               tools.Int64ToStringNegative(d.Sc1Gpc1Rate.ValueInt64()),
		Sc2Gpc0RateComparison:     api.SelectedMap(d.Sc2Gpc0RateComparison.ValueString()),
		Sc2Gpc0Rate:               tools.Int64ToStringNegative(d.Sc2Gpc0Rate.ValueInt64()),
		Sc2Gpc1RateComparison:     api.SelectedMap(d.Sc2Gpc1RateComparison.ValueString()),
		Sc2Gpc1Rate:               tools.Int64ToStringNegative(d.Sc2Gpc1Rate.ValueInt64()),
		ScIncGpc0Comparison:       api.SelectedMap(d.ScIncGpc0Comparison.ValueString()),
		ScIncGpc0:                 tools.Int64ToStringNegative(d.ScIncGpc0.ValueInt64()),
		ScIncGpc1Comparison:       api.SelectedMap(d.ScIncGpc1Comparison.ValueString()),
		ScIncGpc1:                 tools.Int64ToStringNegative(d.ScIncGpc1.ValueInt64()),
		Sc0IncGpc0Comparison:      api.SelectedMap(d.Sc0IncGpc0Comparison.ValueString()),
		Sc0IncGpc0:                tools.Int64ToStringNegative(d.Sc0IncGpc0.ValueInt64()),
		Sc0IncGpc1Comparison:      api.SelectedMap(d.Sc0IncGpc1Comparison.ValueString()),
		Sc0IncGpc1:                tools.Int64ToStringNegative(d.Sc0IncGpc1.ValueInt64()),
		Sc1IncGpc0Comparison:      api.SelectedMap(d.Sc1IncGpc0Comparison.ValueString()),
		Sc1IncGpc0:                tools.Int64ToStringNegative(d.Sc1IncGpc0.ValueInt64()),
		Sc1IncGpc1Comparison:      api.SelectedMap(d.Sc1IncGpc1Comparison.ValueString()),
		Sc1IncGpc1:                tools.Int64ToStringNegative(d.Sc1IncGpc1.ValueInt64()),
		Sc2IncGpc0Comparison:      api.SelectedMap(d.Sc2IncGpc0Comparison.ValueString()),
		Sc2IncGpc0:                tools.Int64ToStringNegative(d.Sc2IncGpc0.ValueInt64()),
		Sc2IncGpc1Comparison:      api.SelectedMap(d.Sc2IncGpc1Comparison.ValueString()),
		Sc2IncGpc1:                tools.Int64ToStringNegative(d.Sc2IncGpc1.ValueInt64()),
		SrcClrGpc0Comparison:      api.SelectedMap(d.SrcClrGpc0Comparison.ValueString()),
		SrcClrGpc0:                tools.Int64ToStringNegative(d.SrcClrGpc0.ValueInt64()),
		SrcClrGpc1Comparison:      api.SelectedMap(d.SrcClrGpc1Comparison.ValueString()),
		SrcClrGpc1:                tools.Int64ToStringNegative(d.SrcClrGpc1.ValueInt64()),
		SrcGetGpc0Comparison:      api.SelectedMap(d.SrcGetGpc0Comparison.ValueString()),
		SrcGetGpc0:                tools.Int64ToStringNegative(d.SrcGetGpc0.ValueInt64()),
		SrcGetGpc1Comparison:      api.SelectedMap(d.SrcGetGpc1Comparison.ValueString()),
		SrcGetGpc1:                tools.Int64ToStringNegative(d.SrcGetGpc1.ValueInt64()),
		SrcGpc0RateComparison:     api.SelectedMap(d.SrcGpc0RateComparison.ValueString()),
		SrcGpc0Rate:               tools.Int64ToStringNegative(d.SrcGpc0Rate.ValueInt64()),
		SrcGpc1RateComparison:     api.SelectedMap(d.SrcGpc1RateComparison.ValueString()),
		SrcGpc1Rate:               tools.Int64ToStringNegative(d.SrcGpc1Rate.ValueInt64()),
		SrcIncGpc0Comparison:      api.SelectedMap(d.SrcIncGpc0Comparison.ValueString()),
		SrcIncGpc0:                tools.Int64ToStringNegative(d.SrcIncGpc0.ValueInt64()),
		SrcIncGpc1Comparison:      api.SelectedMap(d.SrcIncGpc1Comparison.ValueString()),
		SrcIncGpc1:                tools.Int64ToStringNegative(d.SrcIncGpc1.ValueInt64()),
		GpcNumber:                 tools.Int64ToStringNegative(d.GpcNumber.ValueInt64()),
		GptNumber:                 tools.Int64ToStringNegative(d.GptNumber.ValueInt64()),
		ScNumber:                  tools.Int64ToStringNegative(d.ScNumber.ValueInt64()),
		TableName:                 d.TableName.ValueString(),
		Mapfile:                   api.SelectedMap(d.Mapfile.ValueString()),
		Converter:                 d.Converter.ValueString(),
	}, nil
}

func convertAclStructToSchema(d *haproxy.Acl) (*aclResourceModel, error) {
	return &aclResourceModel{
		Name:                      types.StringValue(d.Name),
		Description:               types.StringValue(d.Description),
		Expression:                types.StringValue(d.Expression.String()),
		Negate:                    types.BoolValue(tools.StringToBool(d.Negate)),
		CaseSensitive:             types.BoolValue(tools.StringToBool(d.CaseSensitive)),
		HdrBeg:                    types.StringValue(d.HdrBeg),
		HdrEnd:                    types.StringValue(d.HdrEnd),
		Hdr:                       types.StringValue(d.Hdr),
		HdrReg:                    types.StringValue(d.HdrReg),
		HdrSub:                    types.StringValue(d.HdrSub),
		PathBeg:                   types.StringValue(d.PathBeg),
		PathEnd:                   types.StringValue(d.PathEnd),
		Path:                      types.StringValue(d.Path),
		PathReg:                   types.StringValue(d.PathReg),
		PathDir:                   types.StringValue(d.PathDir),
		PathSub:                   types.StringValue(d.PathSub),
		CustHdrBegName:            types.StringValue(d.CustHdrBegName),
		CustHdrBeg:                types.StringValue(d.CustHdrBeg),
		CustHdrEndName:            types.StringValue(d.CustHdrEndName),
		CustHdrEnd:                types.StringValue(d.CustHdrEnd),
		CustHdrName:               types.StringValue(d.CustHdrName),
		CustHdr:                   types.StringValue(d.CustHdr),
		CustHdrRegName:            types.StringValue(d.CustHdrRegName),
		CustHdrReg:                types.StringValue(d.CustHdrReg),
		CustHdrSubName:            types.StringValue(d.CustHdrSubName),
		CustHdrSub:                types.StringValue(d.CustHdrSub),
		URLParam:                  types.StringValue(d.UrlParam),
		URLParamValue:             types.StringValue(d.UrlParamValue),
		Var:                       types.StringValue(d.Var),
		VarValue:                  types.StringValue(d.VarValue),
		VarComparison:             types.StringValue(d.VarComparison.String()),
		SslCVerifyCode:            types.Int64Value(tools.StringToInt64(d.SslCVerifyCode)),
		SslCCaCommonname:          types.StringValue(d.SslCCaCommonname),
		SslHelloType:              types.StringValue(d.SslHelloType.String()),
		Src:                       types.StringValue(d.Src),
		SrcBytesInRateComparison:  types.StringValue(d.SrcBytesInRateComparison.String()),
		SrcBytesInRate:            types.Int64Value(tools.StringToInt64(d.SrcBytesInRate)),
		SrcBytesOutRateComparison: types.StringValue(d.SrcBytesOutRateComparison.String()),
		SrcBytesOutRate:           types.Int64Value(tools.StringToInt64(d.SrcBytesOutRate)),
		SrcConnCntComparison:      types.StringValue(d.SrcConnCntComparison.String()),
		SrcConnCnt:                types.Int64Value(tools.StringToInt64(d.SrcConnCnt)),
		SrcConnCurComparison:      types.StringValue(d.SrcConnCurComparison.String()),
		SrcConnCur:                types.Int64Value(tools.StringToInt64(d.SrcConnCur)),
		SrcConnRateComparison:     types.StringValue(d.SrcConnRateComparison.String()),
		SrcConnRate:               types.Int64Value(tools.StringToInt64(d.SrcConnRate)),
		SrcHttpErrCntComparison:   types.StringValue(d.SrcHttpErrCntComparison.String()),
		SrcHttpErrCnt:             types.Int64Value(tools.StringToInt64(d.SrcHttpErrCnt)),
		SrcHttpErrRateComparison:  types.StringValue(d.SrcHttpErrRateComparison.String()),
		SrcHttpErrRate:            types.Int64Value(tools.StringToInt64(d.SrcHttpErrRate)),
		SrcHttpReqCntComparison:   types.StringValue(d.SrcHttpReqCntComparison.String()),
		SrcHttpReqCnt:             types.Int64Value(tools.StringToInt64(d.SrcHttpReqCnt)),
		SrcHttpReqRateComparison:  types.StringValue(d.SrcHttpReqRateComparison.String()),
		SrcHttpReqRate:            types.Int64Value(tools.StringToInt64(d.SrcHttpReqRate)),
		SrcKbytesInComparison:     types.StringValue(d.SrcKbytesInComparison.String()),
		SrcKbytesIn:               types.Int64Value(tools.StringToInt64(d.SrcKbytesIn)),
		SrcKbytesOutComparison:    types.StringValue(d.SrcKbytesOutComparison.String()),
		SrcKbytesOut:              types.Int64Value(tools.StringToInt64(d.SrcKbytesOut)),
		SrcPortComparison:         types.StringValue(d.SrcPortComparison.String()),
		SrcPort:                   types.Int64Value(tools.StringToInt64(d.SrcPort)),
		SrcSessCntComparison:      types.StringValue(d.SrcSessCntComparison.String()),
		SrcSessCnt:                types.Int64Value(tools.StringToInt64(d.SrcSessCnt)),
		SrcSessRateComparison:     types.StringValue(d.SrcSessRateComparison.String()),
		SrcSessRate:               types.Int64Value(tools.StringToInt64(d.SrcSessRate)),
		Nbsrv:                     types.Int64Value(tools.StringToInt64(d.Nbsrv)),
		NbsrvBackend:              types.StringValue(d.NbsrvBackend.String()),
		SslFcSni:                  types.StringValue(d.SslFcSni),
		SslSni:                    types.StringValue(d.SslSni),
		SslSniSub:                 types.StringValue(d.SslSniSub),
		SslSniBeg:                 types.StringValue(d.SslSniBeg),
		SslSniEnd:                 types.StringValue(d.SslSniEnd),
		SslSniReg:                 types.StringValue(d.SslSniReg),
		CustomAcl:                 types.StringValue(d.CustomAcl),
		Value:                     types.StringValue(d.Value),
		Urlparam:                  types.StringValue(d.Urlparam),
		QueryBackend:              types.StringValue(d.QueryBackend.String()),
		AllowedUsers:              tools.StringSliceToSet(d.AllowedUsers),
		AllowedGroups:             tools.StringSliceToSet(d.AllowedGroups),
		HttpMethod:                tools.StringSliceToSet(d.HttpMethod),
		ScBytesInRateComparison:   types.StringValue(d.ScBytesInRateComparison.String()),
		ScBytesInRate:             types.Int64Value(tools.StringToInt64(d.ScBytesInRate)),
		ScBytesOutRateComparison:  types.StringValue(d.ScBytesOutRateComparison.String()),
		ScBytesOutRate:            types.Int64Value(tools.StringToInt64(d.ScBytesOutRate)),
		ScClrGpcComparison:        types.StringValue(d.ScClrGpcComparison.String()),
		ScClrGpc:                  types.Int64Value(tools.StringToInt64(d.ScClrGpc)),
		ScConnCntComparison:       types.StringValue(d.ScConnCntComparison.String()),
		ScConnCnt:                 types.Int64Value(tools.StringToInt64(d.ScConnCnt)),
		ScConnCurComparison:       types.StringValue(d.ScConnCurComparison.String()),
		ScConnCur:                 types.Int64Value(tools.StringToInt64(d.ScConnCur)),
		ScConnRateComparison:      types.StringValue(d.ScConnRateComparison.String()),
		ScConnRate:                types.Int64Value(tools.StringToInt64(d.ScConnRate)),
		ScGetGpcComparison:        types.StringValue(d.ScGetGpcComparison.String()),
		ScGetGpc:                  types.Int64Value(tools.StringToInt64(d.ScGetGpc)),
		ScGlitchCntComparison:     types.StringValue(d.ScGlitchCntComparison.String()),
		ScGlitchCnt:               types.Int64Value(tools.StringToInt64(d.ScGlitchCnt)),
		ScGlitchRateComparison:    types.StringValue(d.ScGlitchRateComparison.String()),
		ScGlitchRate:              types.Int64Value(tools.StringToInt64(d.ScGlitchRate)),
		ScGpcRateComparison:       types.StringValue(d.ScGpcRateComparison.String()),
		ScGpcRate:                 types.Int64Value(tools.StringToInt64(d.ScGpcRate)),
		ScHttpErrCntComparison:    types.StringValue(d.ScHttpErrCntComparison.String()),
		ScHttpErrCnt:              types.Int64Value(tools.StringToInt64(d.ScHttpErrCnt)),
		ScHttpErrRateComparison:   types.StringValue(d.ScHttpErrRateComparison.String()),
		ScHttpErrRate:             types.Int64Value(tools.StringToInt64(d.ScHttpErrRate)),
		ScHttpFailCntComparison:   types.StringValue(d.ScHttpFailCntComparison.String()),
		ScHttpFailCnt:             types.Int64Value(tools.StringToInt64(d.ScHttpFailCnt)),
		ScHttpFailRateComparison:  types.StringValue(d.ScHttpFailRateComparison.String()),
		ScHttpFailRate:            types.Int64Value(tools.StringToInt64(d.ScHttpFailRate)),
		ScHttpReqCntComparison:    types.StringValue(d.ScHttpReqCntComparison.String()),
		ScHttpReqCnt:              types.Int64Value(tools.StringToInt64(d.ScHttpReqCnt)),
		ScHttpReqRateComparison:   types.StringValue(d.ScHttpReqRateComparison.String()),
		ScHttpReqRate:             types.Int64Value(tools.StringToInt64(d.ScHttpReqRate)),
		ScIncGpcComparison:        types.StringValue(d.ScIncGpcComparison.String()),
		ScIncGpc:                  types.Int64Value(tools.StringToInt64(d.ScIncGpc)),
		ScSessCntComparison:       types.StringValue(d.ScSessCntComparison.String()),
		ScSessCnt:                 types.Int64Value(tools.StringToInt64(d.ScSessCnt)),
		ScSessRateComparison:      types.StringValue(d.ScSessRateComparison.String()),
		ScSessRate:                types.Int64Value(tools.StringToInt64(d.ScSessRate)),
		SrcGetGpcComparison:       types.StringValue(d.SrcGetGpcComparison.String()),
		SrcGetGpc:                 types.Int64Value(tools.StringToInt64(d.SrcGetGpc)),
		SrcGetGptComparison:       types.StringValue(d.SrcGetGptComparison.String()),
		SrcGetGpt:                 types.Int64Value(tools.StringToInt64(d.SrcGetGpt)),
		SrcGlitchCntComparison:    types.StringValue(d.SrcGlitchCntComparison.String()),
		SrcGlitchCnt:              types.Int64Value(tools.StringToInt64(d.SrcGlitchCnt)),
		SrcGlitchRateComparison:   types.StringValue(d.SrcGlitchRateComparison.String()),
		SrcGlitchRate:             types.Int64Value(tools.StringToInt64(d.SrcGlitchRate)),
		SrcGpcRateComparison:      types.StringValue(d.SrcGpcRateComparison.String()),
		SrcGpcRate:                types.Int64Value(tools.StringToInt64(d.SrcGpcRate)),
		SrcHttpFailCntComparison:  types.StringValue(d.SrcHttpFailCntComparison.String()),
		SrcHttpFailCnt:            types.Int64Value(tools.StringToInt64(d.SrcHttpFailCnt)),
		SrcHttpFailRateComparison: types.StringValue(d.SrcHttpFailRateComparison.String()),
		SrcHttpFailRate:           types.Int64Value(tools.StringToInt64(d.SrcHttpFailRate)),
		SrcIncGpcComparison:       types.StringValue(d.SrcIncGpcComparison.String()),
		SrcIncGpc:                 types.Int64Value(tools.StringToInt64(d.SrcIncGpc)),
		ScClrGpc0Comparison:       types.StringValue(d.ScClrGpc0Comparison.String()),
		ScClrGpc0:                 types.Int64Value(tools.StringToInt64(d.ScClrGpc0)),
		ScClrGpc1Comparison:       types.StringValue(d.ScClrGpc1Comparison.String()),
		ScClrGpc1:                 types.Int64Value(tools.StringToInt64(d.ScClrGpc1)),
		Sc0ClrGpc0Comparison:      types.StringValue(d.Sc0ClrGpc0Comparison.String()),
		Sc0ClrGpc0:                types.Int64Value(tools.StringToInt64(d.Sc0ClrGpc0)),
		Sc0ClrGpc1Comparison:      types.StringValue(d.Sc0ClrGpc1Comparison.String()),
		Sc0ClrGpc1:                types.Int64Value(tools.StringToInt64(d.Sc0ClrGpc1)),
		Sc1ClrGpcComparison:       types.StringValue(d.Sc1ClrGpcComparison.String()),
		Sc1ClrGpc:                 types.Int64Value(tools.StringToInt64(d.Sc1ClrGpc)),
		Sc1ClrGpc0Comparison:      types.StringValue(d.Sc1ClrGpc0Comparison.String()),
		Sc1ClrGpc0:                types.Int64Value(tools.StringToInt64(d.Sc1ClrGpc0)),
		Sc1ClrGpc1Comparison:      types.StringValue(d.Sc1ClrGpc1Comparison.String()),
		Sc1ClrGpc1:                types.Int64Value(tools.StringToInt64(d.Sc1ClrGpc1)),
		Sc2ClrGpcComparison:       types.StringValue(d.Sc2ClrGpcComparison.String()),
		Sc2ClrGpc:                 types.Int64Value(tools.StringToInt64(d.Sc2ClrGpc)),
		Sc2ClrGpc0Comparison:      types.StringValue(d.Sc2ClrGpc0Comparison.String()),
		Sc2ClrGpc0:                types.Int64Value(tools.StringToInt64(d.Sc2ClrGpc0)),
		Sc2ClrGpc1Comparison:      types.StringValue(d.Sc2ClrGpc1Comparison.String()),
		Sc2ClrGpc1:                types.Int64Value(tools.StringToInt64(d.Sc2ClrGpc1)),
		ScGetGpc0Comparison:       types.StringValue(d.ScGetGpc0Comparison.String()),
		ScGetGpc0:                 types.Int64Value(tools.StringToInt64(d.ScGetGpc0)),
		ScGetGpc1Comparison:       types.StringValue(d.ScGetGpc1Comparison.String()),
		ScGetGpc1:                 types.Int64Value(tools.StringToInt64(d.ScGetGpc1)),
		Sc0GetGpc0Comparison:      types.StringValue(d.Sc0GetGpc0Comparison.String()),
		Sc0GetGpc0:                types.Int64Value(tools.StringToInt64(d.Sc0GetGpc0)),
		Sc0GetGpc1Comparison:      types.StringValue(d.Sc0GetGpc1Comparison.String()),
		Sc0GetGpc1:                types.Int64Value(tools.StringToInt64(d.Sc0GetGpc1)),
		Sc1GetGpc0Comparison:      types.StringValue(d.Sc1GetGpc0Comparison.String()),
		Sc1GetGpc0:                types.Int64Value(tools.StringToInt64(d.Sc1GetGpc0)),
		Sc1GetGpc1Comparison:      types.StringValue(d.Sc1GetGpc1Comparison.String()),
		Sc1GetGpc1:                types.Int64Value(tools.StringToInt64(d.Sc1GetGpc1)),
		Sc2GetGpc0Comparison:      types.StringValue(d.Sc2GetGpc0Comparison.String()),
		Sc2GetGpc0:                types.Int64Value(tools.StringToInt64(d.Sc2GetGpc0)),
		Sc2GetGpc1Comparison:      types.StringValue(d.Sc2GetGpc1Comparison.String()),
		Sc2GetGpc1:                types.Int64Value(tools.StringToInt64(d.Sc2GetGpc1)),
		ScGetGptComparison:        types.StringValue(d.ScGetGptComparison.String()),
		ScGetGpt:                  types.Int64Value(tools.StringToInt64(d.ScGetGpt)),
		ScGetGpt0Comparison:       types.StringValue(d.ScGetGpt0Comparison.String()),
		ScGetGpt0:                 types.Int64Value(tools.StringToInt64(d.ScGetGpt0)),
		Sc0GetGpt0Comparison:      types.StringValue(d.Sc0GetGpt0Comparison.String()),
		Sc0GetGpt0:                types.Int64Value(tools.StringToInt64(d.Sc0GetGpt0)),
		Sc1GetGpt0Comparison:      types.StringValue(d.Sc1GetGpt0Comparison.String()),
		Sc1GetGpt0:                types.Int64Value(tools.StringToInt64(d.Sc1GetGpt0)),
		Sc2GetGpt0Comparison:      types.StringValue(d.Sc2GetGpt0Comparison.String()),
		Sc2GetGpt0:                types.Int64Value(tools.StringToInt64(d.Sc2GetGpt0)),
		ScGpc0RateComparison:      types.StringValue(d.ScGpc0RateComparison.String()),
		ScGpc0Rate:                types.Int64Value(tools.StringToInt64(d.ScGpc0Rate)),
		ScGpc1RateComparison:      types.StringValue(d.ScGpc1RateComparison.String()),
		ScGpc1Rate:                types.Int64Value(tools.StringToInt64(d.ScGpc1Rate)),
		Sc0Gpc0RateComparison:     types.StringValue(d.Sc0Gpc0RateComparison.String()),
		Sc0Gpc0Rate:               types.Int64Value(tools.StringToInt64(d.Sc0Gpc0Rate)),
		Sc0Gpc1RateComparison:     types.StringValue(d.Sc0Gpc1RateComparison.String()),
		Sc0Gpc1Rate:               types.Int64Value(tools.StringToInt64(d.Sc0Gpc1Rate)),
		Sc1Gpc0RateComparison:     types.StringValue(d.Sc1Gpc0RateComparison.String()),
		Sc1Gpc0Rate:               types.Int64Value(tools.StringToInt64(d.Sc1Gpc0Rate)),
		Sc1Gpc1RateComparison:     types.StringValue(d.Sc1Gpc1RateComparison.String()),
		Sc1Gpc1Rate:               types.Int64Value(tools.StringToInt64(d.Sc1Gpc1Rate)),
		Sc2Gpc0RateComparison:     types.StringValue(d.Sc2Gpc0RateComparison.String()),
		Sc2Gpc0Rate:               types.Int64Value(tools.StringToInt64(d.Sc2Gpc0Rate)),
		Sc2Gpc1RateComparison:     types.StringValue(d.Sc2Gpc1RateComparison.String()),
		Sc2Gpc1Rate:               types.Int64Value(tools.StringToInt64(d.Sc2Gpc1Rate)),
		ScIncGpc0Comparison:       types.StringValue(d.ScIncGpc0Comparison.String()),
		ScIncGpc0:                 types.Int64Value(tools.StringToInt64(d.ScIncGpc0)),
		ScIncGpc1Comparison:       types.StringValue(d.ScIncGpc1Comparison.String()),
		ScIncGpc1:                 types.Int64Value(tools.StringToInt64(d.ScIncGpc1)),
		Sc0IncGpc0Comparison:      types.StringValue(d.Sc0IncGpc0Comparison.String()),
		Sc0IncGpc0:                types.Int64Value(tools.StringToInt64(d.Sc0IncGpc0)),
		Sc0IncGpc1Comparison:      types.StringValue(d.Sc0IncGpc1Comparison.String()),
		Sc0IncGpc1:                types.Int64Value(tools.StringToInt64(d.Sc0IncGpc1)),
		Sc1IncGpc0Comparison:      types.StringValue(d.Sc1IncGpc0Comparison.String()),
		Sc1IncGpc0:                types.Int64Value(tools.StringToInt64(d.Sc1IncGpc0)),
		Sc1IncGpc1Comparison:      types.StringValue(d.Sc1IncGpc1Comparison.String()),
		Sc1IncGpc1:                types.Int64Value(tools.StringToInt64(d.Sc1IncGpc1)),
		Sc2IncGpc0Comparison:      types.StringValue(d.Sc2IncGpc0Comparison.String()),
		Sc2IncGpc0:                types.Int64Value(tools.StringToInt64(d.Sc2IncGpc0)),
		Sc2IncGpc1Comparison:      types.StringValue(d.Sc2IncGpc1Comparison.String()),
		Sc2IncGpc1:                types.Int64Value(tools.StringToInt64(d.Sc2IncGpc1)),
		SrcClrGpc0Comparison:      types.StringValue(d.SrcClrGpc0Comparison.String()),
		SrcClrGpc0:                types.Int64Value(tools.StringToInt64(d.SrcClrGpc0)),
		SrcClrGpc1Comparison:      types.StringValue(d.SrcClrGpc1Comparison.String()),
		SrcClrGpc1:                types.Int64Value(tools.StringToInt64(d.SrcClrGpc1)),
		SrcGetGpc0Comparison:      types.StringValue(d.SrcGetGpc0Comparison.String()),
		SrcGetGpc0:                types.Int64Value(tools.StringToInt64(d.SrcGetGpc0)),
		SrcGetGpc1Comparison:      types.StringValue(d.SrcGetGpc1Comparison.String()),
		SrcGetGpc1:                types.Int64Value(tools.StringToInt64(d.SrcGetGpc1)),
		SrcGpc0RateComparison:     types.StringValue(d.SrcGpc0RateComparison.String()),
		SrcGpc0Rate:               types.Int64Value(tools.StringToInt64(d.SrcGpc0Rate)),
		SrcGpc1RateComparison:     types.StringValue(d.SrcGpc1RateComparison.String()),
		SrcGpc1Rate:               types.Int64Value(tools.StringToInt64(d.SrcGpc1Rate)),
		SrcIncGpc0Comparison:      types.StringValue(d.SrcIncGpc0Comparison.String()),
		SrcIncGpc0:                types.Int64Value(tools.StringToInt64(d.SrcIncGpc0)),
		SrcIncGpc1Comparison:      types.StringValue(d.SrcIncGpc1Comparison.String()),
		SrcIncGpc1:                types.Int64Value(tools.StringToInt64(d.SrcIncGpc1)),
		GpcNumber:                 types.Int64Value(tools.StringToInt64(d.GpcNumber)),
		GptNumber:                 types.Int64Value(tools.StringToInt64(d.GptNumber)),
		ScNumber:                  types.Int64Value(tools.StringToInt64(d.ScNumber)),
		TableName:                 types.StringValue(d.TableName),
		Mapfile:                   types.StringValue(d.Mapfile.String()),
		Converter:                 types.StringValue(d.Converter),
	}, nil
}
