package proxy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/proxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

type pacMatchResourceModel struct {
	Name            types.String `tfsdk:"name"`
	Description     types.String `tfsdk:"description"`
	Negate          types.Bool   `tfsdk:"negate"`
	MatchType       types.String `tfsdk:"match_type"`
	Hostname        types.String `tfsdk:"hostname"`
	URL             types.String `tfsdk:"url"`
	Network         types.String `tfsdk:"network"`
	DomainLevelFrom types.Int64  `tfsdk:"domain_level_from"`
	DomainLevelTo   types.Int64  `tfsdk:"domain_level_to"`
	TimeFrom        types.Int64  `tfsdk:"time_from"`
	TimeTo          types.Int64  `tfsdk:"time_to"`
	DateFrom        types.String `tfsdk:"date_from"`
	DateTo          types.String `tfsdk:"date_to"`
	WeekdayFrom     types.String `tfsdk:"weekday_from"`
	WeekdayTo       types.String `tfsdk:"weekday_to"`

	Id types.String `tfsdk:"id"`
}

func pacMatchResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Conditions used by proxy auto-config rules (**Services ▸ Squid Web Proxy ▸ Proxy Auto Config ▸ Matches**).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the match.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"negate": schema.BoolAttribute{
				MarkdownDescription: "Invert the match. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"match_type": schema.StringAttribute{
				MarkdownDescription: "What is matched.",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf("url_matches", "hostname_matches", "dns_domain_is", "destination_in_net", "my_ip_in_net", "plain_hostname", "is_resolvable", "dns_domain_levels", "weekday_range", "date_range", "time_range")},
			},
			"hostname": schema.StringAttribute{
				MarkdownDescription: "Host name pattern, for the hostname and domain matches.",
				Optional:            true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "URL pattern, for `url_matches`.",
				Optional:            true,
			},
			"network": schema.StringAttribute{
				MarkdownDescription: "Network in CIDR notation, for the network matches.",
				Optional:            true,
			},
			"domain_level_from": schema.Int64Attribute{
				MarkdownDescription: "Minimum number of domain levels, for `dns_domain_levels`. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.AtLeast(0))},
			},
			"domain_level_to": schema.Int64Attribute{
				MarkdownDescription: "Maximum number of domain levels, for `dns_domain_levels`. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.AtLeast(0))},
			},
			"time_from": schema.Int64Attribute{
				MarkdownDescription: "Start hour, for `time_range`. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.Between(0, 23))},
			},
			"time_to": schema.Int64Attribute{
				MarkdownDescription: "End hour, for `time_range`. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.Between(0, 23))},
			},
			"date_from": schema.StringAttribute{
				MarkdownDescription: "Start month, for `date_range`. Defaults to `JAN`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("JAN"),
				Validators:          []validator.String{stringvalidator.OneOf("JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC")},
			},
			"date_to": schema.StringAttribute{
				MarkdownDescription: "End month, for `date_range`. Defaults to `DEC`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("DEC"),
				Validators:          []validator.String{stringvalidator.OneOf("JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC")},
			},
			"weekday_from": schema.StringAttribute{
				MarkdownDescription: "First weekday, for `weekday_range`. Defaults to `MON`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("MON"),
				Validators:          []validator.String{stringvalidator.OneOf("MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN")},
			},
			"weekday_to": schema.StringAttribute{
				MarkdownDescription: "Last weekday, for `weekday_range`. Defaults to `SUN`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("SUN"),
				Validators:          []validator.String{stringvalidator.OneOf("MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN")},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the PAC match.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func pacMatchDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Conditions used by proxy auto-config rules (**Services ▸ Squid Web Proxy ▸ Proxy Auto Config ▸ Matches**).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the PAC match.",
				Required:            true,
			},
			"name":              dschema.StringAttribute{MarkdownDescription: "Name of the match.", Computed: true},
			"description":       dschema.StringAttribute{MarkdownDescription: "Optional description here for your reference (not parsed).", Computed: true},
			"negate":            dschema.BoolAttribute{MarkdownDescription: "Invert the match.", Computed: true},
			"match_type":        dschema.StringAttribute{MarkdownDescription: "What is matched.", Computed: true},
			"hostname":          dschema.StringAttribute{MarkdownDescription: "Host name pattern, for the hostname and domain matches.", Computed: true},
			"url":               dschema.StringAttribute{MarkdownDescription: "URL pattern, for `url_matches`.", Computed: true},
			"network":           dschema.StringAttribute{MarkdownDescription: "Network in CIDR notation, for the network matches.", Computed: true},
			"domain_level_from": dschema.Int64Attribute{MarkdownDescription: "Minimum number of domain levels, for `dns_domain_levels`.", Computed: true},
			"domain_level_to":   dschema.Int64Attribute{MarkdownDescription: "Maximum number of domain levels, for `dns_domain_levels`.", Computed: true},
			"time_from":         dschema.Int64Attribute{MarkdownDescription: "Start hour, for `time_range`.", Computed: true},
			"time_to":           dschema.Int64Attribute{MarkdownDescription: "End hour, for `time_range`.", Computed: true},
			"date_from":         dschema.StringAttribute{MarkdownDescription: "Start month, for `date_range`.", Computed: true},
			"date_to":           dschema.StringAttribute{MarkdownDescription: "End month, for `date_range`.", Computed: true},
			"weekday_from":      dschema.StringAttribute{MarkdownDescription: "First weekday, for `weekday_range`.", Computed: true},
			"weekday_to":        dschema.StringAttribute{MarkdownDescription: "Last weekday, for `weekday_range`.", Computed: true},
		},
	}
}

func convertPacMatchSchemaToStruct(d *pacMatchResourceModel) (*proxy.PacMatch, error) {
	return &proxy.PacMatch{
		Name:            d.Name.ValueString(),
		Description:     d.Description.ValueString(),
		Negate:          tools.BoolToString(d.Negate.ValueBool()),
		MatchType:       api.SelectedMap(d.MatchType.ValueString()),
		Hostname:        d.Hostname.ValueString(),
		URL:             d.URL.ValueString(),
		Network:         d.Network.ValueString(),
		DomainLevelFrom: tools.Int64ToStringNegative(d.DomainLevelFrom.ValueInt64()),
		DomainLevelTo:   tools.Int64ToStringNegative(d.DomainLevelTo.ValueInt64()),
		TimeFrom:        tools.Int64ToStringNegative(d.TimeFrom.ValueInt64()),
		TimeTo:          tools.Int64ToStringNegative(d.TimeTo.ValueInt64()),
		DateFrom:        api.SelectedMap(d.DateFrom.ValueString()),
		DateTo:          api.SelectedMap(d.DateTo.ValueString()),
		WeekdayFrom:     api.SelectedMap(d.WeekdayFrom.ValueString()),
		WeekdayTo:       api.SelectedMap(d.WeekdayTo.ValueString()),
	}, nil
}

func convertPacMatchStructToSchema(d *proxy.PacMatch) (*pacMatchResourceModel, error) {
	return &pacMatchResourceModel{
		Name:            types.StringValue(d.Name),
		Description:     tools.StringOrNull(d.Description),
		Negate:          types.BoolValue(tools.StringToBool(d.Negate)),
		MatchType:       types.StringValue(d.MatchType.String()),
		Hostname:        tools.StringOrNull(d.Hostname),
		URL:             tools.StringOrNull(d.URL),
		Network:         tools.StringOrNull(d.Network),
		DomainLevelFrom: types.Int64Value(tools.StringToInt64(d.DomainLevelFrom)),
		DomainLevelTo:   types.Int64Value(tools.StringToInt64(d.DomainLevelTo)),
		TimeFrom:        types.Int64Value(tools.StringToInt64(d.TimeFrom)),
		TimeTo:          types.Int64Value(tools.StringToInt64(d.TimeTo)),
		DateFrom:        types.StringValue(d.DateFrom.String()),
		DateTo:          types.StringValue(d.DateTo.String()),
		WeekdayFrom:     types.StringValue(d.WeekdayFrom.String()),
		WeekdayTo:       types.StringValue(d.WeekdayTo.String()),
	}, nil
}
