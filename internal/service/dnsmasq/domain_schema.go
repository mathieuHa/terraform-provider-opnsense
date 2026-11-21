package dnsmasq

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dnsmasq"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type domainResourceModel struct {
	Sequence      types.Int64  `tfsdk:"sequence"`
	Domain        types.String `tfsdk:"domain"`
	IP            types.String `tfsdk:"ip"`
	Port          types.Int64  `tfsdk:"port"`
	SourceIP      types.String `tfsdk:"source_ip"`
	FirewallAlias types.String `tfsdk:"firewall_alias"`
	Description   types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func domainResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Dnsmasq domain overrides send queries for a domain to a specific DNS server.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"sequence": schema.Int64Attribute{
				MarkdownDescription: "Order in which overrides are evaluated. Defaults to `1`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1),
				Validators:          []validator.Int64{int64validator.AtLeast(1)},
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "Domain to override, e.g. `corp.example.com`.",
				Required:            true,
			},
			"ip": schema.StringAttribute{
				MarkdownDescription: "IP address of the DNS server queries are forwarded to.",
				Required:            true,
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: "Port of that DNS server. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.Between(1, 65535))},
			},
			"source_ip": schema.StringAttribute{
				MarkdownDescription: "Source address used to reach the server.",
				Optional:            true,
			},
			"firewall_alias": schema.StringAttribute{
				MarkdownDescription: "Firewall alias filled with the addresses resolved for this domain (ipset).",
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the dnsmasq domain override.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func domainDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Dnsmasq domain overrides send queries for a domain to a specific DNS server.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the dnsmasq domain override.",
				Required:            true,
			},
			"sequence":       dschema.Int64Attribute{MarkdownDescription: "Order in which overrides are evaluated.", Computed: true},
			"domain":         dschema.StringAttribute{MarkdownDescription: "Domain to override, e.g. `corp.example.com`.", Computed: true},
			"ip":             dschema.StringAttribute{MarkdownDescription: "IP address of the DNS server queries are forwarded to.", Computed: true},
			"port":           dschema.Int64Attribute{MarkdownDescription: "Port of that DNS server.", Computed: true},
			"source_ip":      dschema.StringAttribute{MarkdownDescription: "Source address used to reach the server.", Computed: true},
			"firewall_alias": dschema.StringAttribute{MarkdownDescription: "Firewall alias filled with the addresses resolved for this domain (ipset).", Computed: true},
			"description":    dschema.StringAttribute{MarkdownDescription: "Optional description here for your reference (not parsed).", Computed: true},
		},
	}
}

func convertDomainSchemaToStruct(d *domainResourceModel) (*dnsmasq.Domain, error) {
	return &dnsmasq.Domain{
		Sequence:      tools.Int64ToString(d.Sequence.ValueInt64()),
		Domain:        d.Domain.ValueString(),
		Ip:            d.IP.ValueString(),
		Port:          tools.Int64ToStringNegative(d.Port.ValueInt64()),
		SourceIp:      d.SourceIP.ValueString(),
		FirewallAlias: api.SelectedMap(d.FirewallAlias.ValueString()),
		Description:   d.Description.ValueString(),
	}, nil
}

func convertDomainStructToSchema(d *dnsmasq.Domain) (*domainResourceModel, error) {
	return &domainResourceModel{
		Sequence:      types.Int64Value(tools.StringToInt64(d.Sequence)),
		Domain:        types.StringValue(d.Domain),
		IP:            types.StringValue(d.Ip),
		Port:          types.Int64Value(tools.StringToInt64(d.Port)),
		SourceIP:      tools.StringOrNull(d.SourceIp),
		FirewallAlias: tools.StringOrNull(d.FirewallAlias.String()),
		Description:   tools.StringOrNull(d.Description),
	}, nil
}
