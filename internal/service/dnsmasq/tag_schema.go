package dnsmasq

import (
	"github.com/browningluke/opnsense-go/pkg/dnsmasq"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type tagResourceModel struct {
	Tag types.String `tfsdk:"tag"`

	Id types.String `tfsdk:"id"`
}

func tagResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Dnsmasq tags (**Services ▸ Dnsmasq DNS & DHCP ▸ DHCP tags**) label DHCP clients, ranges and options so options can be sent selectively.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"tag": schema.StringAttribute{
				MarkdownDescription: "Name of the tag.",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the dnsmasq tag.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func tagDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Dnsmasq tags (**Services ▸ Dnsmasq DNS & DHCP ▸ DHCP tags**) label DHCP clients, ranges and options so options can be sent selectively.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the dnsmasq tag.",
				Required:            true,
			},
			"tag": dschema.StringAttribute{MarkdownDescription: "Name of the tag.", Computed: true},
		},
	}
}

func convertTagSchemaToStruct(d *tagResourceModel) (*dnsmasq.Tag, error) {
	return &dnsmasq.Tag{
		Tag: d.Tag.ValueString(),
	}, nil
}

func convertTagStructToSchema(d *dnsmasq.Tag) (*tagResourceModel, error) {
	return &tagResourceModel{
		Tag: types.StringValue(d.Tag),
	}, nil
}
