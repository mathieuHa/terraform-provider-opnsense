package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type limitRequestConnectionResourceModel struct {
	LimitZone       types.Set    `tfsdk:"limit_zone"`
	ConnectionCount types.Int64  `tfsdk:"connection_count"`
	Burst           types.Int64  `tfsdk:"burst"`
	Nodelay         types.Bool   `tfsdk:"nodelay"`
	Description     types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func limitRequestConnectionResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx connection limits (**Services ▸ Nginx ▸ Configuration ▸ Access ▸ Connection Limits**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"limit_zone": schema.SetAttribute{
				MarkdownDescription: "Limit Zone.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"connection_count": schema.Int64Attribute{
				MarkdownDescription: "Connection Count (Streams Only). Defaults to `5`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(5),
			},
			"burst": schema.Int64Attribute{
				MarkdownDescription: "Burst (HTTP Only). Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"nodelay": schema.BoolAttribute{
				MarkdownDescription: "No Delay (HTTP Only). Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description.",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx connection limit.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func limitRequestConnectionDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx connection limits (**Services ▸ Nginx ▸ Configuration ▸ Access ▸ Connection Limits**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx connection limit.",
				Required:            true,
			},
			"limit_zone":       dschema.SetAttribute{MarkdownDescription: "Limit Zone.", Computed: true, ElementType: types.StringType},
			"connection_count": dschema.Int64Attribute{MarkdownDescription: "Connection Count (Streams Only).", Computed: true},
			"burst":            dschema.Int64Attribute{MarkdownDescription: "Burst (HTTP Only).", Computed: true},
			"nodelay":          dschema.BoolAttribute{MarkdownDescription: "No Delay (HTTP Only).", Computed: true},
			"description":      dschema.StringAttribute{MarkdownDescription: "Description.", Computed: true},
		},
	}
}

func convertLimitRequestConnectionSchemaToStruct(d *limitRequestConnectionResourceModel) (*nginx.LimitRequestConnection, error) {
	return &nginx.LimitRequestConnection{
		LimitZone:       api.SelectedMapList(setToSlice(d.LimitZone)),
		ConnectionCount: tools.Int64ToString(d.ConnectionCount.ValueInt64()),
		Burst:           tools.Int64ToStringNegative(d.Burst.ValueInt64()),
		Nodelay:         tools.BoolToString(d.Nodelay.ValueBool()),
		Description:     d.Description.ValueString(),
	}, nil
}

func convertLimitRequestConnectionStructToSchema(d *nginx.LimitRequestConnection) (*limitRequestConnectionResourceModel, error) {
	return &limitRequestConnectionResourceModel{
		LimitZone:       tools.StringSliceToSet(d.LimitZone),
		ConnectionCount: types.Int64Value(tools.StringToInt64(d.ConnectionCount)),
		Burst:           types.Int64Value(tools.StringToInt64(d.Burst)),
		Nodelay:         types.BoolValue(tools.StringToBool(d.Nodelay)),
		Description:     types.StringValue(d.Description),
	}, nil
}
