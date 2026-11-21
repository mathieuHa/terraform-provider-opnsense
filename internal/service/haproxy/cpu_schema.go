package haproxy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type cpuResourceModel struct {
	Enabled  types.Bool   `tfsdk:"enabled"`
	Name     types.String `tfsdk:"name"`
	ThreadID types.String `tfsdk:"thread_id"`
	CpuID    types.Set    `tfsdk:"cpu_id"`

	Id types.String `tfsdk:"id"`
}

func cpuResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "HAProxy CPU affinity rules (**Services ▸ HAProxy ▸ Settings ▸ Global Parameters ▸ CPU Affinity Rules**, plugin `os-haproxy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this CPU affinity rule. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Choose a name for this CPU affinity rule.",
				Required:            true,
			},
			"thread_id": schema.StringAttribute{
				MarkdownDescription: "Thread ID that should bind to a specific CPU set. Any thread IDs above nbthread are ignored. One of `all`, `odd`, `even`, `x1`, `x2`, `x3`, `x4`, `x5`, `x6`, `x7`, `x8`, `x9`, `x10`, `x11`, `x12`, `x13`, `x14`, `x15`, `x16`, `x17`, `x18`, `x19`, `x20`, `x21`, `x22`, `x23`, `x24`, `x25`, `x26`, `x27`, `x28`, `x29`, `x30`, `x31`, `x32`, `x33`, `x34`, `x35`, `x36`, `x37`, `x38`, `x39`, `x40`, `x41`, `x42`, `x43`, `x44`, `x45`, `x46`, `x47`, `x48`, `x49`, `x50`, `x51`, `x52`, `x53`, `x54`, `x55`, `x56`, `x57`, `x58`, `x59`, `x60`, `x61`, `x62`, `x63`.",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf("all", "odd", "even", "x1", "x2", "x3", "x4", "x5", "x6", "x7", "x8", "x9", "x10", "x11", "x12", "x13", "x14", "x15", "x16", "x17", "x18", "x19", "x20", "x21", "x22", "x23", "x24", "x25", "x26", "x27", "x28", "x29", "x30", "x31", "x32", "x33", "x34", "x35", "x36", "x37", "x38", "x39", "x40", "x41", "x42", "x43", "x44", "x45", "x46", "x47", "x48", "x49", "x50", "x51", "x52", "x53", "x54", "x55", "x56", "x57", "x58", "x59", "x60", "x61", "x62", "x63")},
			},
			"cpu_id": schema.SetAttribute{
				MarkdownDescription: "Bind the process/thread ID to this CPU. One of `all`, `odd`, `even`, `x0`, `x1`, `x2`, `x3`, `x4`, `x5`, `x6`, `x7`, `x8`, `x9`, `x10`, `x11`, `x12`, `x13`, `x14`, `x15`, `x16`, `x17`, `x18`, `x19`, `x20`, `x21`, `x22`, `x23`, `x24`, `x25`, `x26`, `x27`, `x28`, `x29`, `x30`, `x31`, `x32`, `x33`, `x34`, `x35`, `x36`, `x37`, `x38`, `x39`, `x40`, `x41`, `x42`, `x43`, `x44`, `x45`, `x46`, `x47`, `x48`, `x49`, `x50`, `x51`, `x52`, `x53`, `x54`, `x55`, `x56`, `x57`, `x58`, `x59`, `x60`, `x61`, `x62`, `x63`.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1), setvalidator.ValueStringsAre(stringvalidator.OneOf("all", "odd", "even", "x0", "x1", "x2", "x3", "x4", "x5", "x6", "x7", "x8", "x9", "x10", "x11", "x12", "x13", "x14", "x15", "x16", "x17", "x18", "x19", "x20", "x21", "x22", "x23", "x24", "x25", "x26", "x27", "x28", "x29", "x30", "x31", "x32", "x33", "x34", "x35", "x36", "x37", "x38", "x39", "x40", "x41", "x42", "x43", "x44", "x45", "x46", "x47", "x48", "x49", "x50", "x51", "x52", "x53", "x54", "x55", "x56", "x57", "x58", "x59", "x60", "x61", "x62", "x63"))},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the CPU affinity rule.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func cpuDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "HAProxy CPU affinity rules (**Services ▸ HAProxy ▸ Settings ▸ Global Parameters ▸ CPU Affinity Rules**, plugin `os-haproxy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the CPU affinity rule.",
				Required:            true,
			},
			"enabled":   dschema.BoolAttribute{MarkdownDescription: "Enable this CPU affinity rule.", Computed: true},
			"name":      dschema.StringAttribute{MarkdownDescription: "Choose a name for this CPU affinity rule.", Computed: true},
			"thread_id": dschema.StringAttribute{MarkdownDescription: "Thread ID that should bind to a specific CPU set. Any thread IDs above nbthread are ignored. One of `all`, `odd`, `even`, `x1`, `x2`, `x3`, `x4`, `x5`, `x6`, `x7`, `x8`, `x9`, `x10`, `x11`, `x12`, `x13`, `x14`, `x15`, `x16`, `x17`, `x18`, `x19`, `x20`, `x21`, `x22`, `x23`, `x24`, `x25`, `x26`, `x27`, `x28`, `x29`, `x30`, `x31`, `x32`, `x33`, `x34`, `x35`, `x36`, `x37`, `x38`, `x39`, `x40`, `x41`, `x42`, `x43`, `x44`, `x45`, `x46`, `x47`, `x48`, `x49`, `x50`, `x51`, `x52`, `x53`, `x54`, `x55`, `x56`, `x57`, `x58`, `x59`, `x60`, `x61`, `x62`, `x63`.", Computed: true},
			"cpu_id":    dschema.SetAttribute{MarkdownDescription: "Bind the process/thread ID to this CPU. One of `all`, `odd`, `even`, `x0`, `x1`, `x2`, `x3`, `x4`, `x5`, `x6`, `x7`, `x8`, `x9`, `x10`, `x11`, `x12`, `x13`, `x14`, `x15`, `x16`, `x17`, `x18`, `x19`, `x20`, `x21`, `x22`, `x23`, `x24`, `x25`, `x26`, `x27`, `x28`, `x29`, `x30`, `x31`, `x32`, `x33`, `x34`, `x35`, `x36`, `x37`, `x38`, `x39`, `x40`, `x41`, `x42`, `x43`, `x44`, `x45`, `x46`, `x47`, `x48`, `x49`, `x50`, `x51`, `x52`, `x53`, `x54`, `x55`, `x56`, `x57`, `x58`, `x59`, `x60`, `x61`, `x62`, `x63`.", Computed: true, ElementType: types.StringType},
		},
	}
}

func convertCpuSchemaToStruct(d *cpuResourceModel) (*haproxy.Cpu, error) {
	return &haproxy.Cpu{
		Enabled:  tools.BoolToString(d.Enabled.ValueBool()),
		Name:     d.Name.ValueString(),
		ThreadId: api.SelectedMap(d.ThreadID.ValueString()),
		CpuId:    api.SelectedMapList(setToSlice(d.CpuID)),
	}, nil
}

func convertCpuStructToSchema(d *haproxy.Cpu) (*cpuResourceModel, error) {
	return &cpuResourceModel{
		Enabled:  types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:     types.StringValue(d.Name),
		ThreadID: types.StringValue(d.ThreadId.String()),
		CpuID:    tools.StringSliceToSet(d.CpuId),
	}, nil
}
