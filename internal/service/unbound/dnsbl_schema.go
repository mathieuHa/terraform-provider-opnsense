package unbound

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/unbound"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type dnsblResourceModel struct {
	Enabled     types.Bool   `tfsdk:"enabled"`
	Type        types.Set    `tfsdk:"type"`
	Lists       types.Set    `tfsdk:"lists"`
	Allowlists  types.Set    `tfsdk:"allowlists"`
	Blocklists  types.Set    `tfsdk:"blocklists"`
	Wildcards   types.Set    `tfsdk:"wildcards"`
	SourceNets  types.Set    `tfsdk:"source_nets"`
	Address     types.String `tfsdk:"address"`
	NXDomain    types.Bool   `tfsdk:"nxdomain"`
	CacheTTL    types.Int64  `tfsdk:"cache_ttl"`
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func dnsblSet(description string) schema.SetAttribute {
	return schema.SetAttribute{
		MarkdownDescription: description + " Defaults to `[]`.",
		Optional:            true,
		Computed:            true,
		ElementType:         types.StringType,
		Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
	}
}

func dnsblResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Unbound DNS blocklist policies (**Services ▸ Unbound DNS ▸ Blocklist**). Each policy combines predefined and custom lists, and can be limited to some client networks.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether this policy is enabled. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"type":        dnsblSet("Predefined blocklists by OPNsense key, e.g. `oisd1` or `hgz002`; see the GUI for the full list."),
			"lists":       dnsblSet("URLs of additional blocklists to download."),
			"allowlists":  dnsblSet("Domains that are never blocked by this policy."),
			"blocklists":  dnsblSet("Domains that are always blocked by this policy."),
			"wildcards":   dnsblSet("Domains blocked together with all their subdomains."),
			"source_nets": dnsblSet("Client networks this policy applies to. Empty means all clients."),
			"address": schema.StringAttribute{
				MarkdownDescription: "Address returned for blocked names. Leave unset for `0.0.0.0` / `::`.",
				Optional:            true,
			},
			"nxdomain": schema.BoolAttribute{
				MarkdownDescription: "When enabled, blocked names are answered with NXDOMAIN instead of an address. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"cache_ttl": schema.Int64Attribute{
				MarkdownDescription: "How long downloaded lists are cached, in seconds. Defaults to `72000`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(72000),
				Validators:          []validator.Int64{int64validator.AtLeast(0)},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the policy.",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the blocklist policy.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func dnsblDataSourceSchema() dschema.Schema {
	set := func(d string) dschema.SetAttribute {
		return dschema.SetAttribute{MarkdownDescription: d, Computed: true, ElementType: types.StringType}
	}
	return dschema.Schema{
		MarkdownDescription: "Unbound DNS blocklist policies.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the blocklist policy.",
				Required:            true,
			},
			"enabled":     dschema.BoolAttribute{MarkdownDescription: "Whether this policy is enabled.", Computed: true},
			"type":        set("Predefined blocklists by OPNsense key."),
			"lists":       set("URLs of additional blocklists."),
			"allowlists":  set("Domains that are never blocked."),
			"blocklists":  set("Domains that are always blocked."),
			"wildcards":   set("Domains blocked together with all their subdomains."),
			"source_nets": set("Client networks this policy applies to."),
			"address":     dschema.StringAttribute{MarkdownDescription: "Address returned for blocked names.", Computed: true},
			"nxdomain":    dschema.BoolAttribute{MarkdownDescription: "Whether blocked names are answered with NXDOMAIN.", Computed: true},
			"cache_ttl":   dschema.Int64Attribute{MarkdownDescription: "How long downloaded lists are cached, in seconds.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Description of the policy.", Computed: true},
		},
	}
}

func setToList(s types.Set) api.SelectedMapList {
	out := tools.SetToStringSlice(s)
	if out == nil {
		out = []string{}
	}
	return api.SelectedMapList(out)
}

func convertDnsblSchemaToStruct(d *dnsblResourceModel) (*unbound.Dnsbl, error) {
	return &unbound.Dnsbl{
		Enabled:     tools.BoolToString(d.Enabled.ValueBool()),
		Type:        setToList(d.Type),
		Lists:       setToList(d.Lists),
		Allowlists:  setToList(d.Allowlists),
		Blocklists:  setToList(d.Blocklists),
		Wildcards:   setToList(d.Wildcards),
		SourceNets:  setToList(d.SourceNets),
		Address:     d.Address.ValueString(),
		NXDomain:    tools.BoolToString(d.NXDomain.ValueBool()),
		CacheTTL:    tools.Int64ToString(d.CacheTTL.ValueInt64()),
		Description: d.Description.ValueString(),
	}, nil
}

func convertDnsblStructToSchema(d *unbound.Dnsbl) (*dnsblResourceModel, error) {
	return &dnsblResourceModel{
		Enabled:     types.BoolValue(tools.StringToBool(d.Enabled)),
		Type:        tools.StringSliceToSet(d.Type),
		Lists:       tools.StringSliceToSet(d.Lists),
		Allowlists:  tools.StringSliceToSet(d.Allowlists),
		Blocklists:  tools.StringSliceToSet(d.Blocklists),
		Wildcards:   tools.StringSliceToSet(d.Wildcards),
		SourceNets:  tools.StringSliceToSet(d.SourceNets),
		Address:     tools.StringOrNull(d.Address),
		NXDomain:    types.BoolValue(tools.StringToBool(d.NXDomain)),
		CacheTTL:    types.Int64Value(tools.StringToInt64(d.CacheTTL)),
		Description: types.StringValue(d.Description),
	}, nil
}
