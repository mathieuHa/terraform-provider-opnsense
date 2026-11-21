package proxy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/proxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type pacRuleResourceModel struct {
	Enabled     types.Bool   `tfsdk:"enabled"`
	Description types.String `tfsdk:"description"`
	Matches     types.Set    `tfsdk:"matches"`
	JoinType    types.String `tfsdk:"join_type"`
	MatchType   types.String `tfsdk:"match_type"`
	Proxies     types.Set    `tfsdk:"proxies"`

	Id types.String `tfsdk:"id"`
}

func pacRuleResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Rules of the proxy auto-config file (**Services ▸ Squid Web Proxy ▸ Proxy Auto Config ▸ Rules**). A rule sends clients to its proxies when its matches apply.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether this rule is enabled. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"matches": schema.SetAttribute{
				MarkdownDescription: "IDs of `opnsense_proxy_pac_match` resources.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"join_type": schema.StringAttribute{
				MarkdownDescription: "How the matches are combined. Defaults to `and`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("and"),
				Validators:          []validator.String{stringvalidator.OneOf("and", "or")},
			},
			"match_type": schema.StringAttribute{
				MarkdownDescription: "Use the proxies `if` the matches apply, or `unless` they do. Defaults to `if`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("if"),
				Validators:          []validator.String{stringvalidator.OneOf("if", "unless")},
			},
			"proxies": schema.SetAttribute{
				MarkdownDescription: "IDs of `opnsense_proxy_pac_proxy` resources, in order of preference.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the PAC rule.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func pacRuleDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Rules of the proxy auto-config file (**Services ▸ Squid Web Proxy ▸ Proxy Auto Config ▸ Rules**). A rule sends clients to its proxies when its matches apply.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the PAC rule.",
				Required:            true,
			},
			"enabled":     dschema.BoolAttribute{MarkdownDescription: "Whether this rule is enabled.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Optional description here for your reference (not parsed).", Computed: true},
			"matches":     dschema.SetAttribute{MarkdownDescription: "IDs of `opnsense_proxy_pac_match` resources.", Computed: true, ElementType: types.StringType},
			"join_type":   dschema.StringAttribute{MarkdownDescription: "How the matches are combined.", Computed: true},
			"match_type":  dschema.StringAttribute{MarkdownDescription: "Use the proxies `if` the matches apply, or `unless` they do.", Computed: true},
			"proxies":     dschema.SetAttribute{MarkdownDescription: "IDs of `opnsense_proxy_pac_proxy` resources, in order of preference.", Computed: true, ElementType: types.StringType},
		},
	}
}

func convertPacRuleSchemaToStruct(d *pacRuleResourceModel) (*proxy.PacRule, error) {
	return &proxy.PacRule{
		Enabled:     tools.BoolToString(d.Enabled.ValueBool()),
		Description: d.Description.ValueString(),
		Matches:     api.SelectedMapList(setToSlice(d.Matches)),
		JoinType:    api.SelectedMap(d.JoinType.ValueString()),
		MatchType:   api.SelectedMap(d.MatchType.ValueString()),
		Proxies:     api.SelectedMapList(setToSlice(d.Proxies)),
	}, nil
}

func convertPacRuleStructToSchema(d *proxy.PacRule) (*pacRuleResourceModel, error) {
	return &pacRuleResourceModel{
		Enabled:     types.BoolValue(tools.StringToBool(d.Enabled)),
		Description: tools.StringOrNull(d.Description),
		Matches:     tools.StringSliceToSet(d.Matches),
		JoinType:    types.StringValue(d.JoinType.String()),
		MatchType:   types.StringValue(d.MatchType.String()),
		Proxies:     tools.StringSliceToSet(d.Proxies),
	}, nil
}
