package interfaces

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/interfaces"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type vxlanResourceModel struct {
	DeviceID   types.Int64  `tfsdk:"device_id"`
	Device     types.String `tfsdk:"device"`
	Vni        types.Int64  `tfsdk:"vni"`
	Local      types.String `tfsdk:"local"`
	LocalPort  types.Int64  `tfsdk:"local_port"`
	Remote     types.String `tfsdk:"remote"`
	RemotePort types.Int64  `tfsdk:"remote_port"`
	Group      types.String `tfsdk:"group"`
	Interface  types.String `tfsdk:"interface"`

	Id types.String `tfsdk:"id"`
}

func vxlanResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "VXLAN devices (**Interfaces ▸ Devices ▸ VXLAN**). The device is `vxlan<device_id>`; set either `remote` (unicast) or `group` (multicast).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "Device name, `vxlan<device_id>`.",
				Computed:            true,
			},
			"device_id": schema.Int64Attribute{
				MarkdownDescription: "Number of the device (`vxlan<device_id>`). Assigned by OPNsense when unset.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
				Validators:          []validator.Int64{int64validator.AtLeast(0)},
			},
			"vni": schema.Int64Attribute{
				MarkdownDescription: "VXLAN network identifier.",
				Required:            true,
				Validators:          []validator.Int64{int64validator.Between(0, 16777215)},
			},
			"local": schema.StringAttribute{
				MarkdownDescription: "Local address the VXLAN traffic is sent from.",
				Required:            true,
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: "Local UDP port. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.Between(1, 65535))},
			},
			"remote": schema.StringAttribute{
				MarkdownDescription: "Unicast remote address.",
				Optional:            true,
			},
			"remote_port": schema.Int64Attribute{
				MarkdownDescription: "Remote UDP port. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.Between(1, 65535))},
			},
			"group": schema.StringAttribute{
				MarkdownDescription: "Multicast group address.",
				Optional:            true,
			},
			"interface": schema.StringAttribute{
				MarkdownDescription: "Interface used for multicast, e.g. `wan`.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the VXLAN.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func vxlanDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "VXLAN devices (**Interfaces ▸ Devices ▸ VXLAN**). The device is `vxlan<device_id>`; set either `remote` (unicast) or `group` (multicast).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the VXLAN.",
				Required:            true,
			},
			"device":      dschema.StringAttribute{MarkdownDescription: "Device name, `vxlan<device_id>`.", Computed: true},
			"device_id":   dschema.Int64Attribute{MarkdownDescription: "Number of the device (`vxlan<device_id>`). Assigned by OPNsense when unset.", Computed: true},
			"vni":         dschema.Int64Attribute{MarkdownDescription: "VXLAN network identifier.", Computed: true},
			"local":       dschema.StringAttribute{MarkdownDescription: "Local address the VXLAN traffic is sent from.", Computed: true},
			"local_port":  dschema.Int64Attribute{MarkdownDescription: "Local UDP port.", Computed: true},
			"remote":      dschema.StringAttribute{MarkdownDescription: "Unicast remote address.", Computed: true},
			"remote_port": dschema.Int64Attribute{MarkdownDescription: "Remote UDP port.", Computed: true},
			"group":       dschema.StringAttribute{MarkdownDescription: "Multicast group address.", Computed: true},
			"interface":   dschema.StringAttribute{MarkdownDescription: "Interface used for multicast, e.g. `wan`.", Computed: true},
		},
	}
}

func convertVxlanSchemaToStruct(d *vxlanResourceModel) (*interfaces.Vxlan, error) {
	return &interfaces.Vxlan{
		DeviceID: func() string {
			if d.DeviceID.IsNull() || d.DeviceID.IsUnknown() {
				return ""
			}
			return tools.Int64ToString(d.DeviceID.ValueInt64())
		}(),
		VNI:        tools.Int64ToString(d.Vni.ValueInt64()),
		Local:      d.Local.ValueString(),
		LocalPort:  tools.Int64ToStringNegative(d.LocalPort.ValueInt64()),
		Remote:     d.Remote.ValueString(),
		RemotePort: tools.Int64ToStringNegative(d.RemotePort.ValueInt64()),
		Group:      d.Group.ValueString(),
		Interface:  api.SelectedMap(d.Interface.ValueString()),
	}, nil
}

func convertVxlanStructToSchema(d *interfaces.Vxlan) (*vxlanResourceModel, error) {
	return &vxlanResourceModel{
		DeviceID: tools.StringToInt64Null(d.DeviceID),
		Device: func() types.String {
			if d.DeviceID == "" {
				return types.StringNull()
			}
			return types.StringValue("vxlan" + d.DeviceID)
		}(),
		Vni:        types.Int64Value(tools.StringToInt64(d.VNI)),
		Local:      types.StringValue(d.Local),
		LocalPort:  types.Int64Value(tools.StringToInt64(d.LocalPort)),
		Remote:     tools.StringOrNull(d.Remote),
		RemotePort: types.Int64Value(tools.StringToInt64(d.RemotePort)),
		Group:      tools.StringOrNull(d.Group),
		Interface:  tools.StringOrNull(d.Interface.String()),
	}, nil
}
