package ipsec

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/ipsec"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

type manualSpdResourceModel struct {
	Enabled         types.Bool   `tfsdk:"enabled"`
	Protocol        types.String `tfsdk:"protocol"`
	RequestID       types.Int64  `tfsdk:"request_id"`
	ConnectionChild types.String `tfsdk:"connection_child"`
	Source          types.String `tfsdk:"source"`
	Destination     types.String `tfsdk:"destination"`
	Description     types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func manualSpdResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manual security policy database entries (**VPN ▸ IPsec ▸ Manual SPD**), for example to route extra networks through a route-based tunnel.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether the entry is enabled. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: "IPsec protocol. Defaults to `esp`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("esp"),
				Validators:          []validator.String{stringvalidator.OneOf("esp", "ah")},
			},
			"request_id": schema.Int64Attribute{
				MarkdownDescription: "Request ID of the child SA the policy is bound to.",
				Required:            true,
				Validators:          []validator.Int64{int64validator.AtLeast(1)},
			},
			"connection_child": schema.StringAttribute{
				MarkdownDescription: "ID of the `opnsense_ipsec_child` the entry belongs to.",
				Optional:            true,
			},
			"source": schema.StringAttribute{
				MarkdownDescription: "Source network, in CIDR notation.",
				Required:            true,
			},
			"destination": schema.StringAttribute{
				MarkdownDescription: "Destination network, in CIDR notation.",
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the IPsec manual SPD entry.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func manualSpdDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Manual security policy database entries (**VPN ▸ IPsec ▸ Manual SPD**), for example to route extra networks through a route-based tunnel.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the IPsec manual SPD entry.",
				Required:            true,
			},
			"enabled":          dschema.BoolAttribute{MarkdownDescription: "Whether the entry is enabled.", Computed: true},
			"protocol":         dschema.StringAttribute{MarkdownDescription: "IPsec protocol.", Computed: true},
			"request_id":       dschema.Int64Attribute{MarkdownDescription: "Request ID of the child SA the policy is bound to.", Computed: true},
			"connection_child": dschema.StringAttribute{MarkdownDescription: "ID of the `opnsense_ipsec_child` the entry belongs to.", Computed: true},
			"source":           dschema.StringAttribute{MarkdownDescription: "Source network, in CIDR notation.", Computed: true},
			"destination":      dschema.StringAttribute{MarkdownDescription: "Destination network, in CIDR notation.", Computed: true},
			"description":      dschema.StringAttribute{MarkdownDescription: "Optional description here for your reference (not parsed).", Computed: true},
		},
	}
}

func convertManualSPDSchemaToStruct(d *manualSpdResourceModel) (*ipsec.ManualSPD, error) {
	return &ipsec.ManualSPD{
		Enabled:         tools.BoolToString(d.Enabled.ValueBool()),
		Protocol:        api.SelectedMap(d.Protocol.ValueString()),
		ReqID:           tools.Int64ToString(d.RequestID.ValueInt64()),
		ConnectionChild: api.SelectedMap(d.ConnectionChild.ValueString()),
		Source:          d.Source.ValueString(),
		Destination:     d.Destination.ValueString(),
		Description:     d.Description.ValueString(),
	}, nil
}

func convertManualSPDStructToSchema(d *ipsec.ManualSPD) (*manualSpdResourceModel, error) {
	return &manualSpdResourceModel{
		Enabled:         types.BoolValue(tools.StringToBool(d.Enabled)),
		Protocol:        types.StringValue(d.Protocol.String()),
		RequestID:       types.Int64Value(tools.StringToInt64(d.ReqID)),
		ConnectionChild: tools.StringOrNull(d.ConnectionChild.String()),
		Source:          types.StringValue(d.Source),
		Destination:     tools.StringOrNull(d.Destination),
		Description:     tools.StringOrNull(d.Description),
	}, nil
}
