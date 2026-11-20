package core

import (
	"github.com/browningluke/opnsense-go/pkg/core"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type tunableResourceModel struct {
	Tunable     types.String `tfsdk:"tunable"`
	Value       types.String `tfsdk:"value"`
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func tunableResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "System tunables set kernel `sysctl` values and boot-time loader variables (**System ▸ Settings ▸ Tunables**). Loader variables only take effect after a reboot.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"tunable": schema.StringAttribute{
				MarkdownDescription: "Name of the tunable, e.g. `net.inet.tcp.delayed_ack`.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "Value to set. Leave unset to keep the system default while still listing the tunable.",
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed). When unset, OPNsense fills in the kernel description of known tunables.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the tunable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func tunableDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "System tunables set kernel `sysctl` values and boot-time loader variables.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the tunable.",
				Required:            true,
			},
			"tunable": dschema.StringAttribute{
				MarkdownDescription: "Name of the tunable.",
				Computed:            true,
			},
			"value": dschema.StringAttribute{
				MarkdownDescription: "Configured value.",
				Computed:            true,
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Computed:            true,
			},
		},
	}
}

func convertTunableSchemaToStruct(d *tunableResourceModel) (*core.Tunable, error) {
	return &core.Tunable{
		Tunable:     d.Tunable.ValueString(),
		Value:       d.Value.ValueString(),
		Description: d.Description.ValueString(),
	}, nil
}

func convertTunableStructToSchema(d *core.Tunable) (*tunableResourceModel, error) {
	return &tunableResourceModel{
		Tunable:     types.StringValue(d.Tunable),
		Value:       tools.StringOrNull(d.Value),
		Description: tools.StringOrNull(d.Description),
	}, nil
}
