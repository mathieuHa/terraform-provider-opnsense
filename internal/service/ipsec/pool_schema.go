package ipsec

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/ipsec"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type poolResourceModel struct {
	Enabled   types.Bool   `tfsdk:"enabled"`
	Name      types.String `tfsdk:"name"`
	Addresses types.String `tfsdk:"addresses"`
	DNS       types.Set    `tfsdk:"dns"`

	Id types.String `tfsdk:"id"`
}

func poolResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Address pools for IPsec road-warrior clients (**VPN ▸ IPsec ▸ Pools**). Reference them from `opnsense_ipsec_connection.ip_pools`.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether the pool is enabled. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the pool.",
				Required:            true,
			},
			"addresses": schema.StringAttribute{
				MarkdownDescription: "Network the client addresses are taken from, in CIDR notation.",
				Required:            true,
			},
			"dns": schema.SetAttribute{
				MarkdownDescription: "DNS servers pushed to the clients. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the IPsec pool.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func poolDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Address pools for IPsec road-warrior clients (**VPN ▸ IPsec ▸ Pools**). Reference them from `opnsense_ipsec_connection.ip_pools`.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the IPsec pool.",
				Required:            true,
			},
			"enabled":   dschema.BoolAttribute{MarkdownDescription: "Whether the pool is enabled.", Computed: true},
			"name":      dschema.StringAttribute{MarkdownDescription: "Name of the pool.", Computed: true},
			"addresses": dschema.StringAttribute{MarkdownDescription: "Network the client addresses are taken from, in CIDR notation.", Computed: true},
			"dns":       dschema.SetAttribute{MarkdownDescription: "DNS servers pushed to the clients.", Computed: true, ElementType: types.StringType},
		},
	}
}

func convertPoolSchemaToStruct(d *poolResourceModel) (*ipsec.Pool, error) {
	return &ipsec.Pool{
		Enabled:   tools.BoolToString(d.Enabled.ValueBool()),
		Name:      d.Name.ValueString(),
		Addresses: d.Addresses.ValueString(),
		DNS:       api.SelectedMapList(setToSlice(d.DNS)),
	}, nil
}

func convertPoolStructToSchema(d *ipsec.Pool) (*poolResourceModel, error) {
	return &poolResourceModel{
		Enabled:   types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:      types.StringValue(d.Name),
		Addresses: types.StringValue(d.Addresses),
		DNS:       tools.StringSliceToSet(d.DNS),
	}, nil
}
