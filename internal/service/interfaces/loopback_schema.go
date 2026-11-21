package interfaces

import (
	"github.com/browningluke/opnsense-go/pkg/interfaces"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type loopbackResourceModel struct {
	DeviceID    types.Int64  `tfsdk:"device_id"`
	Device      types.String `tfsdk:"device"`
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func loopbackResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Loopback devices (**Interfaces ▸ Devices ▸ Loopback**). The device is `lo<device_id>`.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "Device name, `lo<device_id>`.",
				Computed:            true,
			},
			"device_id": schema.Int64Attribute{
				MarkdownDescription: "Number of the device (`lo<device_id>`). Assigned by OPNsense when unset.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
				Validators:          []validator.Int64{int64validator.AtLeast(1)},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the loopback.",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the loopback.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func loopbackDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Loopback devices (**Interfaces ▸ Devices ▸ Loopback**). The device is `lo<device_id>`.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the loopback.",
				Required:            true,
			},
			"device":      dschema.StringAttribute{MarkdownDescription: "Device name, `lo<device_id>`.", Computed: true},
			"device_id":   dschema.Int64Attribute{MarkdownDescription: "Number of the device (`lo<device_id>`). Assigned by OPNsense when unset.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Description of the loopback.", Computed: true},
		},
	}
}

func convertLoopbackSchemaToStruct(d *loopbackResourceModel) (*interfaces.Loopback, error) {
	return &interfaces.Loopback{
		DeviceID: func() string {
			if d.DeviceID.IsNull() || d.DeviceID.IsUnknown() {
				return ""
			}
			return tools.Int64ToString(d.DeviceID.ValueInt64())
		}(),
		Description: d.Description.ValueString(),
	}, nil
}

func convertLoopbackStructToSchema(d *interfaces.Loopback) (*loopbackResourceModel, error) {
	return &loopbackResourceModel{
		DeviceID: tools.StringToInt64Null(d.DeviceID),
		Device: func() types.String {
			if d.DeviceID == "" {
				return types.StringNull()
			}
			return types.StringValue("lo" + d.DeviceID)
		}(),
		Description: types.StringValue(d.Description),
	}, nil
}
