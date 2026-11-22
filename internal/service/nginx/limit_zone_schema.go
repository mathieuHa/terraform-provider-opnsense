package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type limitZoneResourceModel struct {
	Description types.String `tfsdk:"description"`
	Key         types.String `tfsdk:"key"`
	RateUnit    types.String `tfsdk:"rate_unit"`
	Size        types.Int64  `tfsdk:"size"`
	Rate        types.Int64  `tfsdk:"rate"`

	Id types.String `tfsdk:"id"`
}

func limitZoneResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx limit zones (**Services ▸ Nginx ▸ Configuration ▸ Access ▸ Limit Zone**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				MarkdownDescription: "Description.",
				Required:            true,
			},
			"key": schema.StringAttribute{
				MarkdownDescription: "Key. One of `binary_remote_addr`. Defaults to `binary_remote_addr`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("binary_remote_addr"),
				Validators:          []validator.String{stringvalidator.OneOf("binary_remote_addr")},
			},
			"rate_unit": schema.StringAttribute{
				MarkdownDescription: "Rate Unit. One of `r/s`, `r/m`. Defaults to `r/s`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("r/s"),
				Validators:          []validator.String{stringvalidator.OneOf("r/s", "r/m")},
			},
			"size": schema.Int64Attribute{
				MarkdownDescription: "Size (MB). Defaults to `10`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(10),
			},
			"rate": schema.Int64Attribute{
				MarkdownDescription: "Rate. Defaults to `20`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(20),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx limit zone.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func limitZoneDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx limit zones (**Services ▸ Nginx ▸ Configuration ▸ Access ▸ Limit Zone**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx limit zone.",
				Required:            true,
			},
			"description": dschema.StringAttribute{MarkdownDescription: "Description.", Computed: true},
			"key":         dschema.StringAttribute{MarkdownDescription: "Key. One of `binary_remote_addr`.", Computed: true},
			"rate_unit":   dschema.StringAttribute{MarkdownDescription: "Rate Unit. One of `r/s`, `r/m`.", Computed: true},
			"size":        dschema.Int64Attribute{MarkdownDescription: "Size (MB).", Computed: true},
			"rate":        dschema.Int64Attribute{MarkdownDescription: "Rate.", Computed: true},
		},
	}
}

func convertLimitZoneSchemaToStruct(d *limitZoneResourceModel) (*nginx.LimitZone, error) {
	return &nginx.LimitZone{
		Description: d.Description.ValueString(),
		Key:         api.SelectedMap(d.Key.ValueString()),
		RateUnit:    api.SelectedMap(d.RateUnit.ValueString()),
		Size:        tools.Int64ToString(d.Size.ValueInt64()),
		Rate:        tools.Int64ToString(d.Rate.ValueInt64()),
	}, nil
}

func convertLimitZoneStructToSchema(d *nginx.LimitZone) (*limitZoneResourceModel, error) {
	return &limitZoneResourceModel{
		Description: types.StringValue(d.Description),
		Key:         types.StringValue(d.Key.String()),
		RateUnit:    types.StringValue(d.RateUnit.String()),
		Size:        types.Int64Value(tools.StringToInt64(d.Size)),
		Rate:        types.Int64Value(tools.StringToInt64(d.Rate)),
	}, nil
}
