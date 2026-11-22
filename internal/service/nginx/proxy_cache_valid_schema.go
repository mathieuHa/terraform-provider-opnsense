package nginx

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type proxyCacheValidResourceModel struct {
	Description types.String `tfsdk:"description"`
	Code        types.Set    `tfsdk:"code"`
	Valid       types.Int64  `tfsdk:"valid"`

	Id types.String `tfsdk:"id"`
}

func proxyCacheValidResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx cache validity rules (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ Cache Validity**, plugin `os-nginx`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				MarkdownDescription: "Brief description for reference.",
				Required:            true,
			},
			"code": schema.SetAttribute{
				MarkdownDescription: "Enter a Respone codes or use \"any\" to cache any responses. Defaults to `[\"any\"]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(types.SetValueMust(types.StringType, []attr.Value{types.StringValue("any")})),
			},
			"valid": schema.Int64Attribute{
				MarkdownDescription: "Specify caching time in minutes.",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the nginx cache validity rule.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func proxyCacheValidDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx cache validity rules (**Services ▸ Nginx ▸ Configuration ▸ HTTP(S) ▸ Cache Validity**, plugin `os-nginx`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the nginx cache validity rule.",
				Required:            true,
			},
			"description": dschema.StringAttribute{MarkdownDescription: "Brief description for reference.", Computed: true},
			"code":        dschema.SetAttribute{MarkdownDescription: "Enter a Respone codes or use \"any\" to cache any responses.", Computed: true, ElementType: types.StringType},
			"valid":       dschema.Int64Attribute{MarkdownDescription: "Specify caching time in minutes.", Computed: true},
		},
	}
}

func convertProxyCacheValidSchemaToStruct(d *proxyCacheValidResourceModel) (*nginx.ProxyCacheValid, error) {
	return &nginx.ProxyCacheValid{
		Description: d.Description.ValueString(),
		Code:        api.SelectedMapList(setToSlice(d.Code)),
		Valid:       tools.Int64ToString(d.Valid.ValueInt64()),
	}, nil
}

func convertProxyCacheValidStructToSchema(d *nginx.ProxyCacheValid) (*proxyCacheValidResourceModel, error) {
	return &proxyCacheValidResourceModel{
		Description: types.StringValue(d.Description),
		Code:        tools.StringSliceToSet(d.Code),
		Valid:       types.Int64Value(tools.StringToInt64(d.Valid)),
	}, nil
}
