package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
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

type naxsiRuleResourceModel struct {
	Description      types.String `tfsdk:"description"`
	Ruletype         types.String `tfsdk:"ruletype"`
	Message          types.String `tfsdk:"message"`
	Identifier       types.Set    `tfsdk:"identifier"`
	DollarURL        types.String `tfsdk:"dollar_url"`
	MatchValue       types.String `tfsdk:"match_value"`
	MatchType        types.String `tfsdk:"match_type"`
	Negate           types.Bool   `tfsdk:"negate"`
	Score            types.Int64  `tfsdk:"score"`
	Regex            types.Bool   `tfsdk:"regex"`
	Args             types.Bool   `tfsdk:"args"`
	URL              types.Bool   `tfsdk:"url"`
	Headers          types.Bool   `tfsdk:"headers"`
	Body             types.Bool   `tfsdk:"body"`
	DollarArgsVar    types.String `tfsdk:"dollar_args_var"`
	DollarBodyVar    types.String `tfsdk:"dollar_body_var"`
	DollarHeadersVar types.String `tfsdk:"dollar_headers_var"`
	FileExtension    types.Bool   `tfsdk:"file_extension"`
	RawBody          types.Bool   `tfsdk:"raw_body"`
	Name             types.Bool   `tfsdk:"name"`

	Id types.String `tfsdk:"id"`
}

func naxsiRuleResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx NAXSI rules (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ Naxsi WAF Rule**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				MarkdownDescription: "Description.",
				Required:            true,
			},
			"ruletype": schema.StringAttribute{
				MarkdownDescription: "Main rules increase the policiy's score if matched. Basic rules are added directly to the location and match immeditaly. Basic rules are specifically useful for whitelist rules. Also basic rules need to be added to a policy to allow association with a location. One of `main`, `basic`.",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf("main", "basic")},
			},
			"message": schema.StringAttribute{
				MarkdownDescription: "A string describing the pattern. This is mostly used for analyzing and to have some human-understandable text. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"identifier": schema.SetAttribute{
				MarkdownDescription: "For blacklist rules, the unique numerical ID of the rule, that will be used in logs and in whitelists. IDs inferior to 1000 are reserved for Naxsi internal rules (protocol mismatch etc.)For whitelists, specify one or more comma separated IDs to whitelist. 0 whitelists all rules.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"dollar_url": schema.StringAttribute{
				MarkdownDescription: "Restrict match to a specific URL (supports regular expressions if enabled for this rule). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"match_value": schema.StringAttribute{
				MarkdownDescription: "Match Value. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"match_type": schema.StringAttribute{
				MarkdownDescription: "A blacklist rule increases the policy's score if the pattern matches. A whitelist rule instructs Naxsi to ignore specific rules under specific conditions. One of `id`, `wl`. Defaults to `id`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("id"),
				Validators:          []validator.String{stringvalidator.OneOf("id", "wl")},
			},
			"negate": schema.BoolAttribute{
				MarkdownDescription: "If this is box is checked, the score will be added if the rule does not match. Use with care. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"score": schema.Int64Attribute{
				MarkdownDescription: "If the rule matches the counter of each policy containing the rule will increase. If a policy's counter exceeds the configured \"value\" (to be configured at the policy) the policy's action gets triggered. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"regex": schema.BoolAttribute{
				MarkdownDescription: "If enabled, the match value, the URL, named parameters and headers are matched using regular expressions; otherwise only exact matches trigger the rule. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"args": schema.BoolAttribute{
				MarkdownDescription: "Search for matchs in a request's GET arguments. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"url": schema.BoolAttribute{
				MarkdownDescription: "Search for matchs in a request's URL (everything before ?). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"headers": schema.BoolAttribute{
				MarkdownDescription: "Search for matchs in a request's HTTP headers. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"body": schema.BoolAttribute{
				MarkdownDescription: "Search for matchs in a request's POST arguments and its raw (unparsed) body. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"dollar_args_var": schema.StringAttribute{
				MarkdownDescription: "Enter the name of a GET parameter to restrict matches to (supports regular expressions if enabled for this rule). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dollar_body_var": schema.StringAttribute{
				MarkdownDescription: "Enter the name of a POST parameter to restrict matches to (supports regular expressions if enabled for this rule). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dollar_headers_var": schema.StringAttribute{
				MarkdownDescription: "Enter the name of a HTTP header to restrict matches to (supports regular expressions if enabled for this rule). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"file_extension": schema.BoolAttribute{
				MarkdownDescription: "Search for matchs in a multipart POST request's filenames. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"raw_body": schema.BoolAttribute{
				MarkdownDescription: "Search for matchs in a request's raw unparsed body. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"name": schema.BoolAttribute{
				MarkdownDescription: "Check this box to match the variable name and not its content when matching any of the above checkboxes. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx NAXSI rule.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func naxsiRuleDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx NAXSI rules (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ Naxsi WAF Rule**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx NAXSI rule.",
				Required:            true,
			},
			"description":        dschema.StringAttribute{MarkdownDescription: "Description.", Computed: true},
			"ruletype":           dschema.StringAttribute{MarkdownDescription: "Main rules increase the policiy's score if matched. Basic rules are added directly to the location and match immeditaly. Basic rules are specifically useful for whitelist rules. Also basic rules need to be added to a policy to allow association with a location. One of `main`, `basic`.", Computed: true},
			"message":            dschema.StringAttribute{MarkdownDescription: "A string describing the pattern. This is mostly used for analyzing and to have some human-understandable text.", Computed: true},
			"identifier":         dschema.SetAttribute{MarkdownDescription: "For blacklist rules, the unique numerical ID of the rule, that will be used in logs and in whitelists. IDs inferior to 1000 are reserved for Naxsi internal rules (protocol mismatch etc.)For whitelists, specify one or more comma separated IDs to whitelist. 0 whitelists all rules.", Computed: true, ElementType: types.StringType},
			"dollar_url":         dschema.StringAttribute{MarkdownDescription: "Restrict match to a specific URL (supports regular expressions if enabled for this rule).", Computed: true},
			"match_value":        dschema.StringAttribute{MarkdownDescription: "Match Value.", Computed: true},
			"match_type":         dschema.StringAttribute{MarkdownDescription: "A blacklist rule increases the policy's score if the pattern matches. A whitelist rule instructs Naxsi to ignore specific rules under specific conditions. One of `id`, `wl`.", Computed: true},
			"negate":             dschema.BoolAttribute{MarkdownDescription: "If this is box is checked, the score will be added if the rule does not match. Use with care.", Computed: true},
			"score":              dschema.Int64Attribute{MarkdownDescription: "If the rule matches the counter of each policy containing the rule will increase. If a policy's counter exceeds the configured \"value\" (to be configured at the policy) the policy's action gets triggered.", Computed: true},
			"regex":              dschema.BoolAttribute{MarkdownDescription: "If enabled, the match value, the URL, named parameters and headers are matched using regular expressions; otherwise only exact matches trigger the rule.", Computed: true},
			"args":               dschema.BoolAttribute{MarkdownDescription: "Search for matchs in a request's GET arguments.", Computed: true},
			"url":                dschema.BoolAttribute{MarkdownDescription: "Search for matchs in a request's URL (everything before ?).", Computed: true},
			"headers":            dschema.BoolAttribute{MarkdownDescription: "Search for matchs in a request's HTTP headers.", Computed: true},
			"body":               dschema.BoolAttribute{MarkdownDescription: "Search for matchs in a request's POST arguments and its raw (unparsed) body.", Computed: true},
			"dollar_args_var":    dschema.StringAttribute{MarkdownDescription: "Enter the name of a GET parameter to restrict matches to (supports regular expressions if enabled for this rule).", Computed: true},
			"dollar_body_var":    dschema.StringAttribute{MarkdownDescription: "Enter the name of a POST parameter to restrict matches to (supports regular expressions if enabled for this rule).", Computed: true},
			"dollar_headers_var": dschema.StringAttribute{MarkdownDescription: "Enter the name of a HTTP header to restrict matches to (supports regular expressions if enabled for this rule).", Computed: true},
			"file_extension":     dschema.BoolAttribute{MarkdownDescription: "Search for matchs in a multipart POST request's filenames.", Computed: true},
			"raw_body":           dschema.BoolAttribute{MarkdownDescription: "Search for matchs in a request's raw unparsed body.", Computed: true},
			"name":               dschema.BoolAttribute{MarkdownDescription: "Check this box to match the variable name and not its content when matching any of the above checkboxes.", Computed: true},
		},
	}
}

func convertNaxsiRuleSchemaToStruct(d *naxsiRuleResourceModel) (*nginx.NaxsiRule, error) {
	return &nginx.NaxsiRule{
		Description:      d.Description.ValueString(),
		Ruletype:         api.SelectedMap(d.Ruletype.ValueString()),
		Message:          d.Message.ValueString(),
		Identifier:       api.SelectedMapList(setToSlice(d.Identifier)),
		DollarUrl:        d.DollarURL.ValueString(),
		MatchValue:       d.MatchValue.ValueString(),
		MatchType:        api.SelectedMap(d.MatchType.ValueString()),
		Negate:           tools.BoolToString(d.Negate.ValueBool()),
		Score:            tools.Int64ToStringNegative(d.Score.ValueInt64()),
		Regex:            tools.BoolToString(d.Regex.ValueBool()),
		Args:             tools.BoolToString(d.Args.ValueBool()),
		Url:              tools.BoolToString(d.URL.ValueBool()),
		Headers:          tools.BoolToString(d.Headers.ValueBool()),
		Body:             tools.BoolToString(d.Body.ValueBool()),
		DollarArgsVar:    d.DollarArgsVar.ValueString(),
		DollarBodyVar:    d.DollarBodyVar.ValueString(),
		DollarHeadersVar: d.DollarHeadersVar.ValueString(),
		FileExtension:    tools.BoolToString(d.FileExtension.ValueBool()),
		RawBody:          tools.BoolToString(d.RawBody.ValueBool()),
		Name:             tools.BoolToString(d.Name.ValueBool()),
	}, nil
}

func convertNaxsiRuleStructToSchema(d *nginx.NaxsiRule) (*naxsiRuleResourceModel, error) {
	return &naxsiRuleResourceModel{
		Description:      types.StringValue(d.Description),
		Ruletype:         types.StringValue(d.Ruletype.String()),
		Message:          types.StringValue(d.Message),
		Identifier:       tools.StringSliceToSet(d.Identifier),
		DollarURL:        types.StringValue(d.DollarUrl),
		MatchValue:       types.StringValue(d.MatchValue),
		MatchType:        types.StringValue(d.MatchType.String()),
		Negate:           types.BoolValue(tools.StringToBool(d.Negate)),
		Score:            types.Int64Value(tools.StringToInt64(d.Score)),
		Regex:            types.BoolValue(tools.StringToBool(d.Regex)),
		Args:             types.BoolValue(tools.StringToBool(d.Args)),
		URL:              types.BoolValue(tools.StringToBool(d.Url)),
		Headers:          types.BoolValue(tools.StringToBool(d.Headers)),
		Body:             types.BoolValue(tools.StringToBool(d.Body)),
		DollarArgsVar:    types.StringValue(d.DollarArgsVar),
		DollarBodyVar:    types.StringValue(d.DollarBodyVar),
		DollarHeadersVar: types.StringValue(d.DollarHeadersVar),
		FileExtension:    types.BoolValue(tools.StringToBool(d.FileExtension)),
		RawBody:          types.BoolValue(tools.StringToBool(d.RawBody)),
		Name:             types.BoolValue(tools.StringToBool(d.Name)),
	}, nil
}
