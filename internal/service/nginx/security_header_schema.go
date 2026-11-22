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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type securityHeaderResourceModel struct {
	Description                              types.String `tfsdk:"description"`
	Referrer                                 types.String `tfsdk:"referrer"`
	Xssprotection                            types.String `tfsdk:"xssprotection"`
	ContentTypeOptions                       types.Bool   `tfsdk:"content_type_options"`
	StrictTransportSecurityTime              types.Int64  `tfsdk:"strict_transport_security_time"`
	StrictTransportSecurityIncludeSubdomains types.Bool   `tfsdk:"strict_transport_security_include_subdomains"`
	StrictTransportSecurityPreload           types.Bool   `tfsdk:"strict_transport_security_preload"`
	EnableCsp                                types.Bool   `tfsdk:"enable_csp"`
	CspLogViolations                         types.Bool   `tfsdk:"csp_log_violations"`
	CspReportOnly                            types.Bool   `tfsdk:"csp_report_only"`
	CspDefaultSrcEnabled                     types.Bool   `tfsdk:"csp_default_src_enabled"`
	CspDefaultSrcDataUrls                    types.Bool   `tfsdk:"csp_default_src_data_urls"`
	CspDefaultSrcHttpUrls                    types.Set    `tfsdk:"csp_default_src_http_urls"`
	CspDefaultSrcInline                      types.Bool   `tfsdk:"csp_default_src_inline"`
	CspDefaultSrcEval                        types.Bool   `tfsdk:"csp_default_src_eval"`
	CspDefaultSrcSelf                        types.Bool   `tfsdk:"csp_default_src_self"`
	CspDefaultSrcBlob                        types.Bool   `tfsdk:"csp_default_src_blob"`
	CspDefaultSrcMediastream                 types.Bool   `tfsdk:"csp_default_src_mediastream"`
	CspDefaultSrcFilesystem                  types.Bool   `tfsdk:"csp_default_src_filesystem"`
	CspDefaultSrcNone                        types.Bool   `tfsdk:"csp_default_src_none"`
	CspScriptSrcEnabled                      types.Bool   `tfsdk:"csp_script_src_enabled"`
	CspScriptSrcDataUrls                     types.Bool   `tfsdk:"csp_script_src_data_urls"`
	CspScriptSrcHttpUrls                     types.Set    `tfsdk:"csp_script_src_http_urls"`
	CspScriptSrcInline                       types.Bool   `tfsdk:"csp_script_src_inline"`
	CspScriptSrcEval                         types.Bool   `tfsdk:"csp_script_src_eval"`
	CspScriptSrcSelf                         types.Bool   `tfsdk:"csp_script_src_self"`
	CspScriptSrcBlob                         types.Bool   `tfsdk:"csp_script_src_blob"`
	CspScriptSrcMediastream                  types.Bool   `tfsdk:"csp_script_src_mediastream"`
	CspScriptSrcFilesystem                   types.Bool   `tfsdk:"csp_script_src_filesystem"`
	CspScriptSrcNone                         types.Bool   `tfsdk:"csp_script_src_none"`
	CspImgSrcEnabled                         types.Bool   `tfsdk:"csp_img_src_enabled"`
	CspImgSrcDataUrls                        types.Bool   `tfsdk:"csp_img_src_data_urls"`
	CspImgSrcHttpUrls                        types.Set    `tfsdk:"csp_img_src_http_urls"`
	CspImgSrcInline                          types.Bool   `tfsdk:"csp_img_src_inline"`
	CspImgSrcEval                            types.Bool   `tfsdk:"csp_img_src_eval"`
	CspImgSrcSelf                            types.Bool   `tfsdk:"csp_img_src_self"`
	CspImgSrcBlob                            types.Bool   `tfsdk:"csp_img_src_blob"`
	CspImgSrcMediastream                     types.Bool   `tfsdk:"csp_img_src_mediastream"`
	CspImgSrcFilesystem                      types.Bool   `tfsdk:"csp_img_src_filesystem"`
	CspImgSrcNone                            types.Bool   `tfsdk:"csp_img_src_none"`
	CspStyleSrcEnabled                       types.Bool   `tfsdk:"csp_style_src_enabled"`
	CspStyleSrcDataUrls                      types.Bool   `tfsdk:"csp_style_src_data_urls"`
	CspStyleSrcHttpUrls                      types.Set    `tfsdk:"csp_style_src_http_urls"`
	CspStyleSrcInline                        types.Bool   `tfsdk:"csp_style_src_inline"`
	CspStyleSrcEval                          types.Bool   `tfsdk:"csp_style_src_eval"`
	CspStyleSrcSelf                          types.Bool   `tfsdk:"csp_style_src_self"`
	CspStyleSrcBlob                          types.Bool   `tfsdk:"csp_style_src_blob"`
	CspStyleSrcMediastream                   types.Bool   `tfsdk:"csp_style_src_mediastream"`
	CspStyleSrcFilesystem                    types.Bool   `tfsdk:"csp_style_src_filesystem"`
	CspStyleSrcNone                          types.Bool   `tfsdk:"csp_style_src_none"`
	CspMediaSrcEnabled                       types.Bool   `tfsdk:"csp_media_src_enabled"`
	CspMediaSrcDataUrls                      types.Bool   `tfsdk:"csp_media_src_data_urls"`
	CspMediaSrcHttpUrls                      types.Set    `tfsdk:"csp_media_src_http_urls"`
	CspMediaSrcInline                        types.Bool   `tfsdk:"csp_media_src_inline"`
	CspMediaSrcEval                          types.Bool   `tfsdk:"csp_media_src_eval"`
	CspMediaSrcSelf                          types.Bool   `tfsdk:"csp_media_src_self"`
	CspMediaSrcBlob                          types.Bool   `tfsdk:"csp_media_src_blob"`
	CspMediaSrcMediastream                   types.Bool   `tfsdk:"csp_media_src_mediastream"`
	CspMediaSrcFilesystem                    types.Bool   `tfsdk:"csp_media_src_filesystem"`
	CspMediaSrcNone                          types.Bool   `tfsdk:"csp_media_src_none"`
	CspFontSrcEnabled                        types.Bool   `tfsdk:"csp_font_src_enabled"`
	CspFontSrcDataUrls                       types.Bool   `tfsdk:"csp_font_src_data_urls"`
	CspFontSrcHttpUrls                       types.Set    `tfsdk:"csp_font_src_http_urls"`
	CspFontSrcInline                         types.Bool   `tfsdk:"csp_font_src_inline"`
	CspFontSrcEval                           types.Bool   `tfsdk:"csp_font_src_eval"`
	CspFontSrcSelf                           types.Bool   `tfsdk:"csp_font_src_self"`
	CspFontSrcBlob                           types.Bool   `tfsdk:"csp_font_src_blob"`
	CspFontSrcMediastream                    types.Bool   `tfsdk:"csp_font_src_mediastream"`
	CspFontSrcFilesystem                     types.Bool   `tfsdk:"csp_font_src_filesystem"`
	CspFontSrcNone                           types.Bool   `tfsdk:"csp_font_src_none"`
	CspFrameSrcEnabled                       types.Bool   `tfsdk:"csp_frame_src_enabled"`
	CspFrameSrcDataUrls                      types.Bool   `tfsdk:"csp_frame_src_data_urls"`
	CspFrameSrcHttpUrls                      types.Set    `tfsdk:"csp_frame_src_http_urls"`
	CspFrameSrcInline                        types.Bool   `tfsdk:"csp_frame_src_inline"`
	CspFrameSrcEval                          types.Bool   `tfsdk:"csp_frame_src_eval"`
	CspFrameSrcSelf                          types.Bool   `tfsdk:"csp_frame_src_self"`
	CspFrameSrcBlob                          types.Bool   `tfsdk:"csp_frame_src_blob"`
	CspFrameSrcMediastream                   types.Bool   `tfsdk:"csp_frame_src_mediastream"`
	CspFrameSrcFilesystem                    types.Bool   `tfsdk:"csp_frame_src_filesystem"`
	CspFrameSrcNone                          types.Bool   `tfsdk:"csp_frame_src_none"`
	CspFrameAncestorsEnabled                 types.Bool   `tfsdk:"csp_frame_ancestors_enabled"`
	CspFrameAncestorsDataUrls                types.Bool   `tfsdk:"csp_frame_ancestors_data_urls"`
	CspFrameAncestorsHttpUrls                types.Set    `tfsdk:"csp_frame_ancestors_http_urls"`
	CspFrameAncestorsSelf                    types.Bool   `tfsdk:"csp_frame_ancestors_self"`
	CspFrameAncestorsBlob                    types.Bool   `tfsdk:"csp_frame_ancestors_blob"`
	CspFrameAncestorsMediastream             types.Bool   `tfsdk:"csp_frame_ancestors_mediastream"`
	CspFrameAncestorsFilesystem              types.Bool   `tfsdk:"csp_frame_ancestors_filesystem"`
	CspFrameAncestorsNone                    types.Bool   `tfsdk:"csp_frame_ancestors_none"`
	CspConnectSrcEnabled                     types.Bool   `tfsdk:"csp_connect_src_enabled"`
	CspConnectSrcHttpUrls                    types.Set    `tfsdk:"csp_connect_src_http_urls"`
	CspConnectSrcNone                        types.Bool   `tfsdk:"csp_connect_src_none"`
	CspWorkerSrcEnabled                      types.Bool   `tfsdk:"csp_worker_src_enabled"`
	CspWorkerSrcDataUrls                     types.Bool   `tfsdk:"csp_worker_src_data_urls"`
	CspWorkerSrcHttpUrls                     types.Set    `tfsdk:"csp_worker_src_http_urls"`
	CspWorkerSrcInline                       types.Bool   `tfsdk:"csp_worker_src_inline"`
	CspWorkerSrcEval                         types.Bool   `tfsdk:"csp_worker_src_eval"`
	CspWorkerSrcSelf                         types.Bool   `tfsdk:"csp_worker_src_self"`
	CspWorkerSrcBlob                         types.Bool   `tfsdk:"csp_worker_src_blob"`
	CspWorkerSrcFilesystem                   types.Bool   `tfsdk:"csp_worker_src_filesystem"`
	CspWorkerSrcNone                         types.Bool   `tfsdk:"csp_worker_src_none"`
	CspFormActionEnabled                     types.Bool   `tfsdk:"csp_form_action_enabled"`
	CspFormActionDataUrls                    types.Bool   `tfsdk:"csp_form_action_data_urls"`
	CspFormActionHttpUrls                    types.Set    `tfsdk:"csp_form_action_http_urls"`
	CspFormActionInline                      types.Bool   `tfsdk:"csp_form_action_inline"`
	CspFormActionEval                        types.Bool   `tfsdk:"csp_form_action_eval"`
	CspFormActionSelf                        types.Bool   `tfsdk:"csp_form_action_self"`
	CspFormActionBlob                        types.Bool   `tfsdk:"csp_form_action_blob"`
	CspFormActionMediastream                 types.Bool   `tfsdk:"csp_form_action_mediastream"`
	CspFormActionFilesystem                  types.Bool   `tfsdk:"csp_form_action_filesystem"`
	CspFormActionNone                        types.Bool   `tfsdk:"csp_form_action_none"`

	Id types.String `tfsdk:"id"`
}

func securityHeaderResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx security header sets (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ Security Headers**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				MarkdownDescription: "This is only for your reference.",
				Required:            true,
			},
			"referrer": schema.StringAttribute{
				MarkdownDescription: "Same Origin: The header will be sent if you stay on the same server using the same protocol (no data leak) No Referrer When Downgrade: Prevents sending a referrer when switching from HTTPS to HTTP Origin, Strict-Origin: Always send the header but no path or query information. Strict Origin additionally suppressed the header on downgrades. (Strict) Origin When Cross Origin: Full Referrer on the same origin, and like (Strict) Origin when cross domain. Unsafe URL: Sends the full URL to all pages. One of `no-referrer`, `no-referrer-when-downgrade`, `same-origin`, `origin`, `strict-origin`, `strict-origin-when-cross-origin`, `origin-when-cross-origin`, `unsafe-url`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "no-referrer", "no-referrer-when-downgrade", "same-origin", "origin", "strict-origin", "strict-origin-when-cross-origin", "origin-when-cross-origin", "unsafe-url")},
			},
			"xssprotection": schema.StringAttribute{
				MarkdownDescription: "Block: The browser should block the response Off: Allow Anything On: The Browser decides how to handle it. One of `1; mode=block`, `0`, `1`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "1; mode=block", "0", "1")},
			},
			"content_type_options": schema.BoolAttribute{
				MarkdownDescription: "Don't Sniff Content Type. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"strict_transport_security_time": schema.Int64Attribute{
				MarkdownDescription: "A time in seconds in which the transport security (TLS) should be enforced. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"strict_transport_security_include_subdomains": schema.BoolAttribute{
				MarkdownDescription: "If checked, also subdomains are affected. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"strict_transport_security_preload": schema.BoolAttribute{
				MarkdownDescription: "If checked, the preload directive is added. Google maintains an HSTS preload service. By following the guidelines and successfully submitting your domain, browsers will never connect to your domain using an insecure connection. While the service is hosted by Google, all browsers have stated an intent to use (or actually started using) the preload list. However, it is not part of the HSTS specification and should not be treated as official. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"enable_csp": schema.BoolAttribute{
				MarkdownDescription: "If checked, the Content Security Policy (CSP) header is enabled. A detailed configuration is still required via the other tabs of this sheet. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_log_violations": schema.BoolAttribute{
				MarkdownDescription: "If checked, the plugin collects CSP violation reports and stores one JSON document per line under /var/log/nginx/csp_violations.log. You can use that file to check for XSS attempts or broken web pages where the CSP denied access to a resource. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_report_only": schema.BoolAttribute{
				MarkdownDescription: "If checked, the CSP is not enforced (learning mode). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_default_src_enabled": schema.BoolAttribute{
				MarkdownDescription: "If checked, this part of the CSP is enabled. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_default_src_data_urls": schema.BoolAttribute{
				MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_default_src_http_urls": schema.SetAttribute{
				MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"csp_default_src_inline": schema.BoolAttribute{
				MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_default_src_eval": schema.BoolAttribute{
				MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_default_src_self": schema.BoolAttribute{
				MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_default_src_blob": schema.BoolAttribute{
				MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_default_src_mediastream": schema.BoolAttribute{
				MarkdownDescription: "Enable Media Streams. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_default_src_filesystem": schema.BoolAttribute{
				MarkdownDescription: "Enable File System URLs. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_default_src_none": schema.BoolAttribute{
				MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_script_src_enabled": schema.BoolAttribute{
				MarkdownDescription: "If checked, this part of the CSP is enabled. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_script_src_data_urls": schema.BoolAttribute{
				MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_script_src_http_urls": schema.SetAttribute{
				MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"csp_script_src_inline": schema.BoolAttribute{
				MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_script_src_eval": schema.BoolAttribute{
				MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_script_src_self": schema.BoolAttribute{
				MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_script_src_blob": schema.BoolAttribute{
				MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_script_src_mediastream": schema.BoolAttribute{
				MarkdownDescription: "Enable Media Streams. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_script_src_filesystem": schema.BoolAttribute{
				MarkdownDescription: "Enable File System URLs. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_script_src_none": schema.BoolAttribute{
				MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_img_src_enabled": schema.BoolAttribute{
				MarkdownDescription: "If checked, this part of the CSP is enabled. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_img_src_data_urls": schema.BoolAttribute{
				MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_img_src_http_urls": schema.SetAttribute{
				MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"csp_img_src_inline": schema.BoolAttribute{
				MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_img_src_eval": schema.BoolAttribute{
				MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_img_src_self": schema.BoolAttribute{
				MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_img_src_blob": schema.BoolAttribute{
				MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_img_src_mediastream": schema.BoolAttribute{
				MarkdownDescription: "Enable Media Streams. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_img_src_filesystem": schema.BoolAttribute{
				MarkdownDescription: "Enable File System URLs. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_img_src_none": schema.BoolAttribute{
				MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_style_src_enabled": schema.BoolAttribute{
				MarkdownDescription: "If checked, this part of the CSP is enabled. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_style_src_data_urls": schema.BoolAttribute{
				MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_style_src_http_urls": schema.SetAttribute{
				MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"csp_style_src_inline": schema.BoolAttribute{
				MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_style_src_eval": schema.BoolAttribute{
				MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_style_src_self": schema.BoolAttribute{
				MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_style_src_blob": schema.BoolAttribute{
				MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_style_src_mediastream": schema.BoolAttribute{
				MarkdownDescription: "Enable Media Streams. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_style_src_filesystem": schema.BoolAttribute{
				MarkdownDescription: "Enable File System URLs. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_style_src_none": schema.BoolAttribute{
				MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_media_src_enabled": schema.BoolAttribute{
				MarkdownDescription: "If checked, this part of the CSP is enabled. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_media_src_data_urls": schema.BoolAttribute{
				MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_media_src_http_urls": schema.SetAttribute{
				MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"csp_media_src_inline": schema.BoolAttribute{
				MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_media_src_eval": schema.BoolAttribute{
				MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_media_src_self": schema.BoolAttribute{
				MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_media_src_blob": schema.BoolAttribute{
				MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_media_src_mediastream": schema.BoolAttribute{
				MarkdownDescription: "Enable Media Streams. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_media_src_filesystem": schema.BoolAttribute{
				MarkdownDescription: "Enable File System URLs. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_media_src_none": schema.BoolAttribute{
				MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_font_src_enabled": schema.BoolAttribute{
				MarkdownDescription: "If checked, this part of the CSP is enabled. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_font_src_data_urls": schema.BoolAttribute{
				MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_font_src_http_urls": schema.SetAttribute{
				MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"csp_font_src_inline": schema.BoolAttribute{
				MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_font_src_eval": schema.BoolAttribute{
				MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_font_src_self": schema.BoolAttribute{
				MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_font_src_blob": schema.BoolAttribute{
				MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_font_src_mediastream": schema.BoolAttribute{
				MarkdownDescription: "Enable Media Streams. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_font_src_filesystem": schema.BoolAttribute{
				MarkdownDescription: "Enable File System URLs. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_font_src_none": schema.BoolAttribute{
				MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_frame_src_enabled": schema.BoolAttribute{
				MarkdownDescription: "If checked, this part of the CSP is enabled. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_frame_src_data_urls": schema.BoolAttribute{
				MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_frame_src_http_urls": schema.SetAttribute{
				MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"csp_frame_src_inline": schema.BoolAttribute{
				MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_frame_src_eval": schema.BoolAttribute{
				MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_frame_src_self": schema.BoolAttribute{
				MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_frame_src_blob": schema.BoolAttribute{
				MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_frame_src_mediastream": schema.BoolAttribute{
				MarkdownDescription: "Enable Media Streams. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_frame_src_filesystem": schema.BoolAttribute{
				MarkdownDescription: "Enable File System URLs. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_frame_src_none": schema.BoolAttribute{
				MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_frame_ancestors_enabled": schema.BoolAttribute{
				MarkdownDescription: "If checked, this part of the CSP is enabled. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_frame_ancestors_data_urls": schema.BoolAttribute{
				MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_frame_ancestors_http_urls": schema.SetAttribute{
				MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"csp_frame_ancestors_self": schema.BoolAttribute{
				MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_frame_ancestors_blob": schema.BoolAttribute{
				MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_frame_ancestors_mediastream": schema.BoolAttribute{
				MarkdownDescription: "Enable Media Streams. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_frame_ancestors_filesystem": schema.BoolAttribute{
				MarkdownDescription: "Enable File System URLs. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_frame_ancestors_none": schema.BoolAttribute{
				MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_connect_src_enabled": schema.BoolAttribute{
				MarkdownDescription: "If checked, this part of the CSP is enabled. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_connect_src_http_urls": schema.SetAttribute{
				MarkdownDescription: "Allow connecting to websockets. You can use wildcards here like wss://*.exmaple.com. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"csp_connect_src_none": schema.BoolAttribute{
				MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_worker_src_enabled": schema.BoolAttribute{
				MarkdownDescription: "If checked, this part of the CSP is enabled. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_worker_src_data_urls": schema.BoolAttribute{
				MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_worker_src_http_urls": schema.SetAttribute{
				MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"csp_worker_src_inline": schema.BoolAttribute{
				MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_worker_src_eval": schema.BoolAttribute{
				MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_worker_src_self": schema.BoolAttribute{
				MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_worker_src_blob": schema.BoolAttribute{
				MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_worker_src_filesystem": schema.BoolAttribute{
				MarkdownDescription: "Enable File System URLs. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_worker_src_none": schema.BoolAttribute{
				MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_form_action_enabled": schema.BoolAttribute{
				MarkdownDescription: "If checked, this part of the CSP is enabled. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_form_action_data_urls": schema.BoolAttribute{
				MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_form_action_http_urls": schema.SetAttribute{
				MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"csp_form_action_inline": schema.BoolAttribute{
				MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_form_action_eval": schema.BoolAttribute{
				MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_form_action_self": schema.BoolAttribute{
				MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_form_action_blob": schema.BoolAttribute{
				MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_form_action_mediastream": schema.BoolAttribute{
				MarkdownDescription: "Enable Media Streams. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_form_action_filesystem": schema.BoolAttribute{
				MarkdownDescription: "Enable File System URLs. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"csp_form_action_none": schema.BoolAttribute{
				MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx security header set.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func securityHeaderDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx security header sets (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ Security Headers**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx security header set.",
				Required:            true,
			},
			"description":                    dschema.StringAttribute{MarkdownDescription: "This is only for your reference.", Computed: true},
			"referrer":                       dschema.StringAttribute{MarkdownDescription: "Same Origin: The header will be sent if you stay on the same server using the same protocol (no data leak) No Referrer When Downgrade: Prevents sending a referrer when switching from HTTPS to HTTP Origin, Strict-Origin: Always send the header but no path or query information. Strict Origin additionally suppressed the header on downgrades. (Strict) Origin When Cross Origin: Full Referrer on the same origin, and like (Strict) Origin when cross domain. Unsafe URL: Sends the full URL to all pages. One of `no-referrer`, `no-referrer-when-downgrade`, `same-origin`, `origin`, `strict-origin`, `strict-origin-when-cross-origin`, `origin-when-cross-origin`, `unsafe-url`.", Computed: true},
			"xssprotection":                  dschema.StringAttribute{MarkdownDescription: "Block: The browser should block the response Off: Allow Anything On: The Browser decides how to handle it. One of `1; mode=block`, `0`, `1`.", Computed: true},
			"content_type_options":           dschema.BoolAttribute{MarkdownDescription: "Don't Sniff Content Type.", Computed: true},
			"strict_transport_security_time": dschema.Int64Attribute{MarkdownDescription: "A time in seconds in which the transport security (TLS) should be enforced.", Computed: true},
			"strict_transport_security_include_subdomains": dschema.BoolAttribute{MarkdownDescription: "If checked, also subdomains are affected.", Computed: true},
			"strict_transport_security_preload":            dschema.BoolAttribute{MarkdownDescription: "If checked, the preload directive is added. Google maintains an HSTS preload service. By following the guidelines and successfully submitting your domain, browsers will never connect to your domain using an insecure connection. While the service is hosted by Google, all browsers have stated an intent to use (or actually started using) the preload list. However, it is not part of the HSTS specification and should not be treated as official.", Computed: true},
			"enable_csp":                                   dschema.BoolAttribute{MarkdownDescription: "If checked, the Content Security Policy (CSP) header is enabled. A detailed configuration is still required via the other tabs of this sheet.", Computed: true},
			"csp_log_violations":                           dschema.BoolAttribute{MarkdownDescription: "If checked, the plugin collects CSP violation reports and stores one JSON document per line under /var/log/nginx/csp_violations.log. You can use that file to check for XSS attempts or broken web pages where the CSP denied access to a resource.", Computed: true},
			"csp_report_only":                              dschema.BoolAttribute{MarkdownDescription: "If checked, the CSP is not enforced (learning mode).", Computed: true},
			"csp_default_src_enabled":                      dschema.BoolAttribute{MarkdownDescription: "If checked, this part of the CSP is enabled.", Computed: true},
			"csp_default_src_data_urls":                    dschema.BoolAttribute{MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute).", Computed: true},
			"csp_default_src_http_urls":                    dschema.SetAttribute{MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com.", Computed: true, ElementType: types.StringType},
			"csp_default_src_inline":                       dschema.BoolAttribute{MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags.", Computed: true},
			"csp_default_src_eval":                         dschema.BoolAttribute{MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS.", Computed: true},
			"csp_default_src_self":                         dschema.BoolAttribute{MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same).", Computed: true},
			"csp_default_src_blob":                         dschema.BoolAttribute{MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript.", Computed: true},
			"csp_default_src_mediastream":                  dschema.BoolAttribute{MarkdownDescription: "Enable Media Streams.", Computed: true},
			"csp_default_src_filesystem":                   dschema.BoolAttribute{MarkdownDescription: "Enable File System URLs.", Computed: true},
			"csp_default_src_none":                         dschema.BoolAttribute{MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden.", Computed: true},
			"csp_script_src_enabled":                       dschema.BoolAttribute{MarkdownDescription: "If checked, this part of the CSP is enabled.", Computed: true},
			"csp_script_src_data_urls":                     dschema.BoolAttribute{MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute).", Computed: true},
			"csp_script_src_http_urls":                     dschema.SetAttribute{MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com.", Computed: true, ElementType: types.StringType},
			"csp_script_src_inline":                        dschema.BoolAttribute{MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags.", Computed: true},
			"csp_script_src_eval":                          dschema.BoolAttribute{MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS.", Computed: true},
			"csp_script_src_self":                          dschema.BoolAttribute{MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same).", Computed: true},
			"csp_script_src_blob":                          dschema.BoolAttribute{MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript.", Computed: true},
			"csp_script_src_mediastream":                   dschema.BoolAttribute{MarkdownDescription: "Enable Media Streams.", Computed: true},
			"csp_script_src_filesystem":                    dschema.BoolAttribute{MarkdownDescription: "Enable File System URLs.", Computed: true},
			"csp_script_src_none":                          dschema.BoolAttribute{MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden.", Computed: true},
			"csp_img_src_enabled":                          dschema.BoolAttribute{MarkdownDescription: "If checked, this part of the CSP is enabled.", Computed: true},
			"csp_img_src_data_urls":                        dschema.BoolAttribute{MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute).", Computed: true},
			"csp_img_src_http_urls":                        dschema.SetAttribute{MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com.", Computed: true, ElementType: types.StringType},
			"csp_img_src_inline":                           dschema.BoolAttribute{MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags.", Computed: true},
			"csp_img_src_eval":                             dschema.BoolAttribute{MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS.", Computed: true},
			"csp_img_src_self":                             dschema.BoolAttribute{MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same).", Computed: true},
			"csp_img_src_blob":                             dschema.BoolAttribute{MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript.", Computed: true},
			"csp_img_src_mediastream":                      dschema.BoolAttribute{MarkdownDescription: "Enable Media Streams.", Computed: true},
			"csp_img_src_filesystem":                       dschema.BoolAttribute{MarkdownDescription: "Enable File System URLs.", Computed: true},
			"csp_img_src_none":                             dschema.BoolAttribute{MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden.", Computed: true},
			"csp_style_src_enabled":                        dschema.BoolAttribute{MarkdownDescription: "If checked, this part of the CSP is enabled.", Computed: true},
			"csp_style_src_data_urls":                      dschema.BoolAttribute{MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute).", Computed: true},
			"csp_style_src_http_urls":                      dschema.SetAttribute{MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com.", Computed: true, ElementType: types.StringType},
			"csp_style_src_inline":                         dschema.BoolAttribute{MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags.", Computed: true},
			"csp_style_src_eval":                           dschema.BoolAttribute{MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS.", Computed: true},
			"csp_style_src_self":                           dschema.BoolAttribute{MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same).", Computed: true},
			"csp_style_src_blob":                           dschema.BoolAttribute{MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript.", Computed: true},
			"csp_style_src_mediastream":                    dschema.BoolAttribute{MarkdownDescription: "Enable Media Streams.", Computed: true},
			"csp_style_src_filesystem":                     dschema.BoolAttribute{MarkdownDescription: "Enable File System URLs.", Computed: true},
			"csp_style_src_none":                           dschema.BoolAttribute{MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden.", Computed: true},
			"csp_media_src_enabled":                        dschema.BoolAttribute{MarkdownDescription: "If checked, this part of the CSP is enabled.", Computed: true},
			"csp_media_src_data_urls":                      dschema.BoolAttribute{MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute).", Computed: true},
			"csp_media_src_http_urls":                      dschema.SetAttribute{MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com.", Computed: true, ElementType: types.StringType},
			"csp_media_src_inline":                         dschema.BoolAttribute{MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags.", Computed: true},
			"csp_media_src_eval":                           dschema.BoolAttribute{MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS.", Computed: true},
			"csp_media_src_self":                           dschema.BoolAttribute{MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same).", Computed: true},
			"csp_media_src_blob":                           dschema.BoolAttribute{MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript.", Computed: true},
			"csp_media_src_mediastream":                    dschema.BoolAttribute{MarkdownDescription: "Enable Media Streams.", Computed: true},
			"csp_media_src_filesystem":                     dschema.BoolAttribute{MarkdownDescription: "Enable File System URLs.", Computed: true},
			"csp_media_src_none":                           dschema.BoolAttribute{MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden.", Computed: true},
			"csp_font_src_enabled":                         dschema.BoolAttribute{MarkdownDescription: "If checked, this part of the CSP is enabled.", Computed: true},
			"csp_font_src_data_urls":                       dschema.BoolAttribute{MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute).", Computed: true},
			"csp_font_src_http_urls":                       dschema.SetAttribute{MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com.", Computed: true, ElementType: types.StringType},
			"csp_font_src_inline":                          dschema.BoolAttribute{MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags.", Computed: true},
			"csp_font_src_eval":                            dschema.BoolAttribute{MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS.", Computed: true},
			"csp_font_src_self":                            dschema.BoolAttribute{MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same).", Computed: true},
			"csp_font_src_blob":                            dschema.BoolAttribute{MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript.", Computed: true},
			"csp_font_src_mediastream":                     dschema.BoolAttribute{MarkdownDescription: "Enable Media Streams.", Computed: true},
			"csp_font_src_filesystem":                      dschema.BoolAttribute{MarkdownDescription: "Enable File System URLs.", Computed: true},
			"csp_font_src_none":                            dschema.BoolAttribute{MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden.", Computed: true},
			"csp_frame_src_enabled":                        dschema.BoolAttribute{MarkdownDescription: "If checked, this part of the CSP is enabled.", Computed: true},
			"csp_frame_src_data_urls":                      dschema.BoolAttribute{MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute).", Computed: true},
			"csp_frame_src_http_urls":                      dschema.SetAttribute{MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com.", Computed: true, ElementType: types.StringType},
			"csp_frame_src_inline":                         dschema.BoolAttribute{MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags.", Computed: true},
			"csp_frame_src_eval":                           dschema.BoolAttribute{MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS.", Computed: true},
			"csp_frame_src_self":                           dschema.BoolAttribute{MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same).", Computed: true},
			"csp_frame_src_blob":                           dschema.BoolAttribute{MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript.", Computed: true},
			"csp_frame_src_mediastream":                    dschema.BoolAttribute{MarkdownDescription: "Enable Media Streams.", Computed: true},
			"csp_frame_src_filesystem":                     dschema.BoolAttribute{MarkdownDescription: "Enable File System URLs.", Computed: true},
			"csp_frame_src_none":                           dschema.BoolAttribute{MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden.", Computed: true},
			"csp_frame_ancestors_enabled":                  dschema.BoolAttribute{MarkdownDescription: "If checked, this part of the CSP is enabled.", Computed: true},
			"csp_frame_ancestors_data_urls":                dschema.BoolAttribute{MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute).", Computed: true},
			"csp_frame_ancestors_http_urls":                dschema.SetAttribute{MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com.", Computed: true, ElementType: types.StringType},
			"csp_frame_ancestors_self":                     dschema.BoolAttribute{MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same).", Computed: true},
			"csp_frame_ancestors_blob":                     dschema.BoolAttribute{MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript.", Computed: true},
			"csp_frame_ancestors_mediastream":              dschema.BoolAttribute{MarkdownDescription: "Enable Media Streams.", Computed: true},
			"csp_frame_ancestors_filesystem":               dschema.BoolAttribute{MarkdownDescription: "Enable File System URLs.", Computed: true},
			"csp_frame_ancestors_none":                     dschema.BoolAttribute{MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden.", Computed: true},
			"csp_connect_src_enabled":                      dschema.BoolAttribute{MarkdownDescription: "If checked, this part of the CSP is enabled.", Computed: true},
			"csp_connect_src_http_urls":                    dschema.SetAttribute{MarkdownDescription: "Allow connecting to websockets. You can use wildcards here like wss://*.exmaple.com.", Computed: true, ElementType: types.StringType},
			"csp_connect_src_none":                         dschema.BoolAttribute{MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden.", Computed: true},
			"csp_worker_src_enabled":                       dschema.BoolAttribute{MarkdownDescription: "If checked, this part of the CSP is enabled.", Computed: true},
			"csp_worker_src_data_urls":                     dschema.BoolAttribute{MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute).", Computed: true},
			"csp_worker_src_http_urls":                     dschema.SetAttribute{MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com.", Computed: true, ElementType: types.StringType},
			"csp_worker_src_inline":                        dschema.BoolAttribute{MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags.", Computed: true},
			"csp_worker_src_eval":                          dschema.BoolAttribute{MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS.", Computed: true},
			"csp_worker_src_self":                          dschema.BoolAttribute{MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same).", Computed: true},
			"csp_worker_src_blob":                          dschema.BoolAttribute{MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript.", Computed: true},
			"csp_worker_src_filesystem":                    dschema.BoolAttribute{MarkdownDescription: "Enable File System URLs.", Computed: true},
			"csp_worker_src_none":                          dschema.BoolAttribute{MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden.", Computed: true},
			"csp_form_action_enabled":                      dschema.BoolAttribute{MarkdownDescription: "If checked, this part of the CSP is enabled.", Computed: true},
			"csp_form_action_data_urls":                    dschema.BoolAttribute{MarkdownDescription: "Data URLs are used to embed files into HTML (for example images written directly into the src attribute).", Computed: true},
			"csp_form_action_http_urls":                    dschema.SetAttribute{MarkdownDescription: "Allow loading files over HTTP(S) allows downloading of content over other domains or CDNs. You can use wildcards here like https://*.exmaple.com.", Computed: true, ElementType: types.StringType},
			"csp_form_action_inline":                       dschema.BoolAttribute{MarkdownDescription: "Checking this directive allows to use scripts or styles directly embedded in in the HTML content. Examples are the script and the style tags.", Computed: true},
			"csp_form_action_eval":                         dschema.BoolAttribute{MarkdownDescription: "Checking this box allows functions like eval or createFunction in JS, or style attributes for CSS.", Computed: true},
			"csp_form_action_self":                         dschema.BoolAttribute{MarkdownDescription: "Allows everything from the same site (path can differ, but host, protocol and port need to be the same).", Computed: true},
			"csp_form_action_blob":                         dschema.BoolAttribute{MarkdownDescription: "Allows to use blobs as a data source. This usually is content, which is somehow generated in JavaScript.", Computed: true},
			"csp_form_action_mediastream":                  dschema.BoolAttribute{MarkdownDescription: "Enable Media Streams.", Computed: true},
			"csp_form_action_filesystem":                   dschema.BoolAttribute{MarkdownDescription: "Enable File System URLs.", Computed: true},
			"csp_form_action_none":                         dschema.BoolAttribute{MarkdownDescription: "If this checkbox is checked, all other settings for this directive are ignored and everything will be forbidden.", Computed: true},
		},
	}
}

func convertSecurityHeaderSchemaToStruct(d *securityHeaderResourceModel) (*nginx.SecurityHeader, error) {
	return &nginx.SecurityHeader{
		Description:                              d.Description.ValueString(),
		Referrer:                                 api.SelectedMap(d.Referrer.ValueString()),
		Xssprotection:                            api.SelectedMap(d.Xssprotection.ValueString()),
		ContentTypeOptions:                       tools.BoolToString(d.ContentTypeOptions.ValueBool()),
		StrictTransportSecurityTime:              tools.Int64ToStringNegative(d.StrictTransportSecurityTime.ValueInt64()),
		StrictTransportSecurityIncludeSubdomains: tools.BoolToString(d.StrictTransportSecurityIncludeSubdomains.ValueBool()),
		StrictTransportSecurityPreload:           tools.BoolToString(d.StrictTransportSecurityPreload.ValueBool()),
		EnableCsp:                                tools.BoolToString(d.EnableCsp.ValueBool()),
		CspLogViolations:                         tools.BoolToString(d.CspLogViolations.ValueBool()),
		CspReportOnly:                            tools.BoolToString(d.CspReportOnly.ValueBool()),
		CspDefaultSrcEnabled:                     tools.BoolToString(d.CspDefaultSrcEnabled.ValueBool()),
		CspDefaultSrcDataUrls:                    tools.BoolToString(d.CspDefaultSrcDataUrls.ValueBool()),
		CspDefaultSrcHttpUrls:                    api.SelectedMapList(setToSlice(d.CspDefaultSrcHttpUrls)),
		CspDefaultSrcInline:                      tools.BoolToString(d.CspDefaultSrcInline.ValueBool()),
		CspDefaultSrcEval:                        tools.BoolToString(d.CspDefaultSrcEval.ValueBool()),
		CspDefaultSrcSelf:                        tools.BoolToString(d.CspDefaultSrcSelf.ValueBool()),
		CspDefaultSrcBlob:                        tools.BoolToString(d.CspDefaultSrcBlob.ValueBool()),
		CspDefaultSrcMediastream:                 tools.BoolToString(d.CspDefaultSrcMediastream.ValueBool()),
		CspDefaultSrcFilesystem:                  tools.BoolToString(d.CspDefaultSrcFilesystem.ValueBool()),
		CspDefaultSrcNone:                        tools.BoolToString(d.CspDefaultSrcNone.ValueBool()),
		CspScriptSrcEnabled:                      tools.BoolToString(d.CspScriptSrcEnabled.ValueBool()),
		CspScriptSrcDataUrls:                     tools.BoolToString(d.CspScriptSrcDataUrls.ValueBool()),
		CspScriptSrcHttpUrls:                     api.SelectedMapList(setToSlice(d.CspScriptSrcHttpUrls)),
		CspScriptSrcInline:                       tools.BoolToString(d.CspScriptSrcInline.ValueBool()),
		CspScriptSrcEval:                         tools.BoolToString(d.CspScriptSrcEval.ValueBool()),
		CspScriptSrcSelf:                         tools.BoolToString(d.CspScriptSrcSelf.ValueBool()),
		CspScriptSrcBlob:                         tools.BoolToString(d.CspScriptSrcBlob.ValueBool()),
		CspScriptSrcMediastream:                  tools.BoolToString(d.CspScriptSrcMediastream.ValueBool()),
		CspScriptSrcFilesystem:                   tools.BoolToString(d.CspScriptSrcFilesystem.ValueBool()),
		CspScriptSrcNone:                         tools.BoolToString(d.CspScriptSrcNone.ValueBool()),
		CspImgSrcEnabled:                         tools.BoolToString(d.CspImgSrcEnabled.ValueBool()),
		CspImgSrcDataUrls:                        tools.BoolToString(d.CspImgSrcDataUrls.ValueBool()),
		CspImgSrcHttpUrls:                        api.SelectedMapList(setToSlice(d.CspImgSrcHttpUrls)),
		CspImgSrcInline:                          tools.BoolToString(d.CspImgSrcInline.ValueBool()),
		CspImgSrcEval:                            tools.BoolToString(d.CspImgSrcEval.ValueBool()),
		CspImgSrcSelf:                            tools.BoolToString(d.CspImgSrcSelf.ValueBool()),
		CspImgSrcBlob:                            tools.BoolToString(d.CspImgSrcBlob.ValueBool()),
		CspImgSrcMediastream:                     tools.BoolToString(d.CspImgSrcMediastream.ValueBool()),
		CspImgSrcFilesystem:                      tools.BoolToString(d.CspImgSrcFilesystem.ValueBool()),
		CspImgSrcNone:                            tools.BoolToString(d.CspImgSrcNone.ValueBool()),
		CspStyleSrcEnabled:                       tools.BoolToString(d.CspStyleSrcEnabled.ValueBool()),
		CspStyleSrcDataUrls:                      tools.BoolToString(d.CspStyleSrcDataUrls.ValueBool()),
		CspStyleSrcHttpUrls:                      api.SelectedMapList(setToSlice(d.CspStyleSrcHttpUrls)),
		CspStyleSrcInline:                        tools.BoolToString(d.CspStyleSrcInline.ValueBool()),
		CspStyleSrcEval:                          tools.BoolToString(d.CspStyleSrcEval.ValueBool()),
		CspStyleSrcSelf:                          tools.BoolToString(d.CspStyleSrcSelf.ValueBool()),
		CspStyleSrcBlob:                          tools.BoolToString(d.CspStyleSrcBlob.ValueBool()),
		CspStyleSrcMediastream:                   tools.BoolToString(d.CspStyleSrcMediastream.ValueBool()),
		CspStyleSrcFilesystem:                    tools.BoolToString(d.CspStyleSrcFilesystem.ValueBool()),
		CspStyleSrcNone:                          tools.BoolToString(d.CspStyleSrcNone.ValueBool()),
		CspMediaSrcEnabled:                       tools.BoolToString(d.CspMediaSrcEnabled.ValueBool()),
		CspMediaSrcDataUrls:                      tools.BoolToString(d.CspMediaSrcDataUrls.ValueBool()),
		CspMediaSrcHttpUrls:                      api.SelectedMapList(setToSlice(d.CspMediaSrcHttpUrls)),
		CspMediaSrcInline:                        tools.BoolToString(d.CspMediaSrcInline.ValueBool()),
		CspMediaSrcEval:                          tools.BoolToString(d.CspMediaSrcEval.ValueBool()),
		CspMediaSrcSelf:                          tools.BoolToString(d.CspMediaSrcSelf.ValueBool()),
		CspMediaSrcBlob:                          tools.BoolToString(d.CspMediaSrcBlob.ValueBool()),
		CspMediaSrcMediastream:                   tools.BoolToString(d.CspMediaSrcMediastream.ValueBool()),
		CspMediaSrcFilesystem:                    tools.BoolToString(d.CspMediaSrcFilesystem.ValueBool()),
		CspMediaSrcNone:                          tools.BoolToString(d.CspMediaSrcNone.ValueBool()),
		CspFontSrcEnabled:                        tools.BoolToString(d.CspFontSrcEnabled.ValueBool()),
		CspFontSrcDataUrls:                       tools.BoolToString(d.CspFontSrcDataUrls.ValueBool()),
		CspFontSrcHttpUrls:                       api.SelectedMapList(setToSlice(d.CspFontSrcHttpUrls)),
		CspFontSrcInline:                         tools.BoolToString(d.CspFontSrcInline.ValueBool()),
		CspFontSrcEval:                           tools.BoolToString(d.CspFontSrcEval.ValueBool()),
		CspFontSrcSelf:                           tools.BoolToString(d.CspFontSrcSelf.ValueBool()),
		CspFontSrcBlob:                           tools.BoolToString(d.CspFontSrcBlob.ValueBool()),
		CspFontSrcMediastream:                    tools.BoolToString(d.CspFontSrcMediastream.ValueBool()),
		CspFontSrcFilesystem:                     tools.BoolToString(d.CspFontSrcFilesystem.ValueBool()),
		CspFontSrcNone:                           tools.BoolToString(d.CspFontSrcNone.ValueBool()),
		CspFrameSrcEnabled:                       tools.BoolToString(d.CspFrameSrcEnabled.ValueBool()),
		CspFrameSrcDataUrls:                      tools.BoolToString(d.CspFrameSrcDataUrls.ValueBool()),
		CspFrameSrcHttpUrls:                      api.SelectedMapList(setToSlice(d.CspFrameSrcHttpUrls)),
		CspFrameSrcInline:                        tools.BoolToString(d.CspFrameSrcInline.ValueBool()),
		CspFrameSrcEval:                          tools.BoolToString(d.CspFrameSrcEval.ValueBool()),
		CspFrameSrcSelf:                          tools.BoolToString(d.CspFrameSrcSelf.ValueBool()),
		CspFrameSrcBlob:                          tools.BoolToString(d.CspFrameSrcBlob.ValueBool()),
		CspFrameSrcMediastream:                   tools.BoolToString(d.CspFrameSrcMediastream.ValueBool()),
		CspFrameSrcFilesystem:                    tools.BoolToString(d.CspFrameSrcFilesystem.ValueBool()),
		CspFrameSrcNone:                          tools.BoolToString(d.CspFrameSrcNone.ValueBool()),
		CspFrameAncestorsEnabled:                 tools.BoolToString(d.CspFrameAncestorsEnabled.ValueBool()),
		CspFrameAncestorsDataUrls:                tools.BoolToString(d.CspFrameAncestorsDataUrls.ValueBool()),
		CspFrameAncestorsHttpUrls:                api.SelectedMapList(setToSlice(d.CspFrameAncestorsHttpUrls)),
		CspFrameAncestorsSelf:                    tools.BoolToString(d.CspFrameAncestorsSelf.ValueBool()),
		CspFrameAncestorsBlob:                    tools.BoolToString(d.CspFrameAncestorsBlob.ValueBool()),
		CspFrameAncestorsMediastream:             tools.BoolToString(d.CspFrameAncestorsMediastream.ValueBool()),
		CspFrameAncestorsFilesystem:              tools.BoolToString(d.CspFrameAncestorsFilesystem.ValueBool()),
		CspFrameAncestorsNone:                    tools.BoolToString(d.CspFrameAncestorsNone.ValueBool()),
		CspConnectSrcEnabled:                     tools.BoolToString(d.CspConnectSrcEnabled.ValueBool()),
		CspConnectSrcHttpUrls:                    api.SelectedMapList(setToSlice(d.CspConnectSrcHttpUrls)),
		CspConnectSrcNone:                        tools.BoolToString(d.CspConnectSrcNone.ValueBool()),
		CspWorkerSrcEnabled:                      tools.BoolToString(d.CspWorkerSrcEnabled.ValueBool()),
		CspWorkerSrcDataUrls:                     tools.BoolToString(d.CspWorkerSrcDataUrls.ValueBool()),
		CspWorkerSrcHttpUrls:                     api.SelectedMapList(setToSlice(d.CspWorkerSrcHttpUrls)),
		CspWorkerSrcInline:                       tools.BoolToString(d.CspWorkerSrcInline.ValueBool()),
		CspWorkerSrcEval:                         tools.BoolToString(d.CspWorkerSrcEval.ValueBool()),
		CspWorkerSrcSelf:                         tools.BoolToString(d.CspWorkerSrcSelf.ValueBool()),
		CspWorkerSrcBlob:                         tools.BoolToString(d.CspWorkerSrcBlob.ValueBool()),
		CspWorkerSrcFilesystem:                   tools.BoolToString(d.CspWorkerSrcFilesystem.ValueBool()),
		CspWorkerSrcNone:                         tools.BoolToString(d.CspWorkerSrcNone.ValueBool()),
		CspFormActionEnabled:                     tools.BoolToString(d.CspFormActionEnabled.ValueBool()),
		CspFormActionDataUrls:                    tools.BoolToString(d.CspFormActionDataUrls.ValueBool()),
		CspFormActionHttpUrls:                    api.SelectedMapList(setToSlice(d.CspFormActionHttpUrls)),
		CspFormActionInline:                      tools.BoolToString(d.CspFormActionInline.ValueBool()),
		CspFormActionEval:                        tools.BoolToString(d.CspFormActionEval.ValueBool()),
		CspFormActionSelf:                        tools.BoolToString(d.CspFormActionSelf.ValueBool()),
		CspFormActionBlob:                        tools.BoolToString(d.CspFormActionBlob.ValueBool()),
		CspFormActionMediastream:                 tools.BoolToString(d.CspFormActionMediastream.ValueBool()),
		CspFormActionFilesystem:                  tools.BoolToString(d.CspFormActionFilesystem.ValueBool()),
		CspFormActionNone:                        tools.BoolToString(d.CspFormActionNone.ValueBool()),
	}, nil
}

func convertSecurityHeaderStructToSchema(d *nginx.SecurityHeader) (*securityHeaderResourceModel, error) {
	return &securityHeaderResourceModel{
		Description:                              types.StringValue(d.Description),
		Referrer:                                 types.StringValue(d.Referrer.String()),
		Xssprotection:                            types.StringValue(d.Xssprotection.String()),
		ContentTypeOptions:                       types.BoolValue(tools.StringToBool(d.ContentTypeOptions)),
		StrictTransportSecurityTime:              types.Int64Value(tools.StringToInt64(d.StrictTransportSecurityTime)),
		StrictTransportSecurityIncludeSubdomains: types.BoolValue(tools.StringToBool(d.StrictTransportSecurityIncludeSubdomains)),
		StrictTransportSecurityPreload:           types.BoolValue(tools.StringToBool(d.StrictTransportSecurityPreload)),
		EnableCsp:                                types.BoolValue(tools.StringToBool(d.EnableCsp)),
		CspLogViolations:                         types.BoolValue(tools.StringToBool(d.CspLogViolations)),
		CspReportOnly:                            types.BoolValue(tools.StringToBool(d.CspReportOnly)),
		CspDefaultSrcEnabled:                     types.BoolValue(tools.StringToBool(d.CspDefaultSrcEnabled)),
		CspDefaultSrcDataUrls:                    types.BoolValue(tools.StringToBool(d.CspDefaultSrcDataUrls)),
		CspDefaultSrcHttpUrls:                    tools.StringSliceToSet(d.CspDefaultSrcHttpUrls),
		CspDefaultSrcInline:                      types.BoolValue(tools.StringToBool(d.CspDefaultSrcInline)),
		CspDefaultSrcEval:                        types.BoolValue(tools.StringToBool(d.CspDefaultSrcEval)),
		CspDefaultSrcSelf:                        types.BoolValue(tools.StringToBool(d.CspDefaultSrcSelf)),
		CspDefaultSrcBlob:                        types.BoolValue(tools.StringToBool(d.CspDefaultSrcBlob)),
		CspDefaultSrcMediastream:                 types.BoolValue(tools.StringToBool(d.CspDefaultSrcMediastream)),
		CspDefaultSrcFilesystem:                  types.BoolValue(tools.StringToBool(d.CspDefaultSrcFilesystem)),
		CspDefaultSrcNone:                        types.BoolValue(tools.StringToBool(d.CspDefaultSrcNone)),
		CspScriptSrcEnabled:                      types.BoolValue(tools.StringToBool(d.CspScriptSrcEnabled)),
		CspScriptSrcDataUrls:                     types.BoolValue(tools.StringToBool(d.CspScriptSrcDataUrls)),
		CspScriptSrcHttpUrls:                     tools.StringSliceToSet(d.CspScriptSrcHttpUrls),
		CspScriptSrcInline:                       types.BoolValue(tools.StringToBool(d.CspScriptSrcInline)),
		CspScriptSrcEval:                         types.BoolValue(tools.StringToBool(d.CspScriptSrcEval)),
		CspScriptSrcSelf:                         types.BoolValue(tools.StringToBool(d.CspScriptSrcSelf)),
		CspScriptSrcBlob:                         types.BoolValue(tools.StringToBool(d.CspScriptSrcBlob)),
		CspScriptSrcMediastream:                  types.BoolValue(tools.StringToBool(d.CspScriptSrcMediastream)),
		CspScriptSrcFilesystem:                   types.BoolValue(tools.StringToBool(d.CspScriptSrcFilesystem)),
		CspScriptSrcNone:                         types.BoolValue(tools.StringToBool(d.CspScriptSrcNone)),
		CspImgSrcEnabled:                         types.BoolValue(tools.StringToBool(d.CspImgSrcEnabled)),
		CspImgSrcDataUrls:                        types.BoolValue(tools.StringToBool(d.CspImgSrcDataUrls)),
		CspImgSrcHttpUrls:                        tools.StringSliceToSet(d.CspImgSrcHttpUrls),
		CspImgSrcInline:                          types.BoolValue(tools.StringToBool(d.CspImgSrcInline)),
		CspImgSrcEval:                            types.BoolValue(tools.StringToBool(d.CspImgSrcEval)),
		CspImgSrcSelf:                            types.BoolValue(tools.StringToBool(d.CspImgSrcSelf)),
		CspImgSrcBlob:                            types.BoolValue(tools.StringToBool(d.CspImgSrcBlob)),
		CspImgSrcMediastream:                     types.BoolValue(tools.StringToBool(d.CspImgSrcMediastream)),
		CspImgSrcFilesystem:                      types.BoolValue(tools.StringToBool(d.CspImgSrcFilesystem)),
		CspImgSrcNone:                            types.BoolValue(tools.StringToBool(d.CspImgSrcNone)),
		CspStyleSrcEnabled:                       types.BoolValue(tools.StringToBool(d.CspStyleSrcEnabled)),
		CspStyleSrcDataUrls:                      types.BoolValue(tools.StringToBool(d.CspStyleSrcDataUrls)),
		CspStyleSrcHttpUrls:                      tools.StringSliceToSet(d.CspStyleSrcHttpUrls),
		CspStyleSrcInline:                        types.BoolValue(tools.StringToBool(d.CspStyleSrcInline)),
		CspStyleSrcEval:                          types.BoolValue(tools.StringToBool(d.CspStyleSrcEval)),
		CspStyleSrcSelf:                          types.BoolValue(tools.StringToBool(d.CspStyleSrcSelf)),
		CspStyleSrcBlob:                          types.BoolValue(tools.StringToBool(d.CspStyleSrcBlob)),
		CspStyleSrcMediastream:                   types.BoolValue(tools.StringToBool(d.CspStyleSrcMediastream)),
		CspStyleSrcFilesystem:                    types.BoolValue(tools.StringToBool(d.CspStyleSrcFilesystem)),
		CspStyleSrcNone:                          types.BoolValue(tools.StringToBool(d.CspStyleSrcNone)),
		CspMediaSrcEnabled:                       types.BoolValue(tools.StringToBool(d.CspMediaSrcEnabled)),
		CspMediaSrcDataUrls:                      types.BoolValue(tools.StringToBool(d.CspMediaSrcDataUrls)),
		CspMediaSrcHttpUrls:                      tools.StringSliceToSet(d.CspMediaSrcHttpUrls),
		CspMediaSrcInline:                        types.BoolValue(tools.StringToBool(d.CspMediaSrcInline)),
		CspMediaSrcEval:                          types.BoolValue(tools.StringToBool(d.CspMediaSrcEval)),
		CspMediaSrcSelf:                          types.BoolValue(tools.StringToBool(d.CspMediaSrcSelf)),
		CspMediaSrcBlob:                          types.BoolValue(tools.StringToBool(d.CspMediaSrcBlob)),
		CspMediaSrcMediastream:                   types.BoolValue(tools.StringToBool(d.CspMediaSrcMediastream)),
		CspMediaSrcFilesystem:                    types.BoolValue(tools.StringToBool(d.CspMediaSrcFilesystem)),
		CspMediaSrcNone:                          types.BoolValue(tools.StringToBool(d.CspMediaSrcNone)),
		CspFontSrcEnabled:                        types.BoolValue(tools.StringToBool(d.CspFontSrcEnabled)),
		CspFontSrcDataUrls:                       types.BoolValue(tools.StringToBool(d.CspFontSrcDataUrls)),
		CspFontSrcHttpUrls:                       tools.StringSliceToSet(d.CspFontSrcHttpUrls),
		CspFontSrcInline:                         types.BoolValue(tools.StringToBool(d.CspFontSrcInline)),
		CspFontSrcEval:                           types.BoolValue(tools.StringToBool(d.CspFontSrcEval)),
		CspFontSrcSelf:                           types.BoolValue(tools.StringToBool(d.CspFontSrcSelf)),
		CspFontSrcBlob:                           types.BoolValue(tools.StringToBool(d.CspFontSrcBlob)),
		CspFontSrcMediastream:                    types.BoolValue(tools.StringToBool(d.CspFontSrcMediastream)),
		CspFontSrcFilesystem:                     types.BoolValue(tools.StringToBool(d.CspFontSrcFilesystem)),
		CspFontSrcNone:                           types.BoolValue(tools.StringToBool(d.CspFontSrcNone)),
		CspFrameSrcEnabled:                       types.BoolValue(tools.StringToBool(d.CspFrameSrcEnabled)),
		CspFrameSrcDataUrls:                      types.BoolValue(tools.StringToBool(d.CspFrameSrcDataUrls)),
		CspFrameSrcHttpUrls:                      tools.StringSliceToSet(d.CspFrameSrcHttpUrls),
		CspFrameSrcInline:                        types.BoolValue(tools.StringToBool(d.CspFrameSrcInline)),
		CspFrameSrcEval:                          types.BoolValue(tools.StringToBool(d.CspFrameSrcEval)),
		CspFrameSrcSelf:                          types.BoolValue(tools.StringToBool(d.CspFrameSrcSelf)),
		CspFrameSrcBlob:                          types.BoolValue(tools.StringToBool(d.CspFrameSrcBlob)),
		CspFrameSrcMediastream:                   types.BoolValue(tools.StringToBool(d.CspFrameSrcMediastream)),
		CspFrameSrcFilesystem:                    types.BoolValue(tools.StringToBool(d.CspFrameSrcFilesystem)),
		CspFrameSrcNone:                          types.BoolValue(tools.StringToBool(d.CspFrameSrcNone)),
		CspFrameAncestorsEnabled:                 types.BoolValue(tools.StringToBool(d.CspFrameAncestorsEnabled)),
		CspFrameAncestorsDataUrls:                types.BoolValue(tools.StringToBool(d.CspFrameAncestorsDataUrls)),
		CspFrameAncestorsHttpUrls:                tools.StringSliceToSet(d.CspFrameAncestorsHttpUrls),
		CspFrameAncestorsSelf:                    types.BoolValue(tools.StringToBool(d.CspFrameAncestorsSelf)),
		CspFrameAncestorsBlob:                    types.BoolValue(tools.StringToBool(d.CspFrameAncestorsBlob)),
		CspFrameAncestorsMediastream:             types.BoolValue(tools.StringToBool(d.CspFrameAncestorsMediastream)),
		CspFrameAncestorsFilesystem:              types.BoolValue(tools.StringToBool(d.CspFrameAncestorsFilesystem)),
		CspFrameAncestorsNone:                    types.BoolValue(tools.StringToBool(d.CspFrameAncestorsNone)),
		CspConnectSrcEnabled:                     types.BoolValue(tools.StringToBool(d.CspConnectSrcEnabled)),
		CspConnectSrcHttpUrls:                    tools.StringSliceToSet(d.CspConnectSrcHttpUrls),
		CspConnectSrcNone:                        types.BoolValue(tools.StringToBool(d.CspConnectSrcNone)),
		CspWorkerSrcEnabled:                      types.BoolValue(tools.StringToBool(d.CspWorkerSrcEnabled)),
		CspWorkerSrcDataUrls:                     types.BoolValue(tools.StringToBool(d.CspWorkerSrcDataUrls)),
		CspWorkerSrcHttpUrls:                     tools.StringSliceToSet(d.CspWorkerSrcHttpUrls),
		CspWorkerSrcInline:                       types.BoolValue(tools.StringToBool(d.CspWorkerSrcInline)),
		CspWorkerSrcEval:                         types.BoolValue(tools.StringToBool(d.CspWorkerSrcEval)),
		CspWorkerSrcSelf:                         types.BoolValue(tools.StringToBool(d.CspWorkerSrcSelf)),
		CspWorkerSrcBlob:                         types.BoolValue(tools.StringToBool(d.CspWorkerSrcBlob)),
		CspWorkerSrcFilesystem:                   types.BoolValue(tools.StringToBool(d.CspWorkerSrcFilesystem)),
		CspWorkerSrcNone:                         types.BoolValue(tools.StringToBool(d.CspWorkerSrcNone)),
		CspFormActionEnabled:                     types.BoolValue(tools.StringToBool(d.CspFormActionEnabled)),
		CspFormActionDataUrls:                    types.BoolValue(tools.StringToBool(d.CspFormActionDataUrls)),
		CspFormActionHttpUrls:                    tools.StringSliceToSet(d.CspFormActionHttpUrls),
		CspFormActionInline:                      types.BoolValue(tools.StringToBool(d.CspFormActionInline)),
		CspFormActionEval:                        types.BoolValue(tools.StringToBool(d.CspFormActionEval)),
		CspFormActionSelf:                        types.BoolValue(tools.StringToBool(d.CspFormActionSelf)),
		CspFormActionBlob:                        types.BoolValue(tools.StringToBool(d.CspFormActionBlob)),
		CspFormActionMediastream:                 types.BoolValue(tools.StringToBool(d.CspFormActionMediastream)),
		CspFormActionFilesystem:                  types.BoolValue(tools.StringToBool(d.CspFormActionFilesystem)),
		CspFormActionNone:                        types.BoolValue(tools.StringToBool(d.CspFormActionNone)),
	}, nil
}
