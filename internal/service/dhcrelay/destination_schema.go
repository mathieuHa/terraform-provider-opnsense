package dhcrelay

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dhcrelay"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type destinationResourceModel struct {
	Name   types.String `tfsdk:"name"`
	Server types.Set    `tfsdk:"server"`

	Id types.String `tfsdk:"id"`
}

func destinationResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "DHCP servers that relayed requests are forwarded to (**Services ▸ DHCP Relay ▸ Destinations**).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the destination.",
				Required:            true,
			},
			"server": schema.SetAttribute{
				MarkdownDescription: "DHCP server addresses.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the DHCP relay destination.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func destinationDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "DHCP servers that relayed requests are forwarded to (**Services ▸ DHCP Relay ▸ Destinations**).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the DHCP relay destination.",
				Required:            true,
			},
			"name":   dschema.StringAttribute{MarkdownDescription: "Name of the destination.", Computed: true},
			"server": dschema.SetAttribute{MarkdownDescription: "DHCP server addresses.", Computed: true, ElementType: types.StringType},
		},
	}
}

func convertDestinationSchemaToStruct(d *destinationResourceModel) (*dhcrelay.Destination, error) {
	return &dhcrelay.Destination{
		Name:   d.Name.ValueString(),
		Server: api.SelectedMapList(setToSlice(d.Server)),
	}, nil
}

func convertDestinationStructToSchema(d *dhcrelay.Destination) (*destinationResourceModel, error) {
	return &destinationResourceModel{
		Name:   types.StringValue(d.Name),
		Server: tools.StringSliceToSet(d.Server),
	}, nil
}
