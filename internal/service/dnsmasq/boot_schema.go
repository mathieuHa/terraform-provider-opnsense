package dnsmasq

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dnsmasq"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type bootResourceModel struct {
	Interface     types.String `tfsdk:"interface"`
	Tags          types.Set    `tfsdk:"tags"`
	Filename      types.String `tfsdk:"filename"`
	ServerName    types.String `tfsdk:"server_name"`
	ServerAddress types.String `tfsdk:"server_address"`
	Description   types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func bootResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Dnsmasq network boot entries (**Services ▸ Dnsmasq DNS & DHCP ▸ DHCP boot**) point PXE clients at a boot file and TFTP server.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"interface": schema.StringAttribute{
				MarkdownDescription: "Only apply on this interface.",
				Optional:            true,
			},
			"tags": schema.SetAttribute{
				MarkdownDescription: "UUIDs of tags a client must carry, e.g. one matching the client architecture. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"filename": schema.StringAttribute{
				MarkdownDescription: "Boot file name, e.g. `pxelinux.0`.",
				Required:            true,
			},
			"server_name": schema.StringAttribute{
				MarkdownDescription: "Name of the boot server.",
				Optional:            true,
			},
			"server_address": schema.StringAttribute{
				MarkdownDescription: "Address of the TFTP server.",
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the dnsmasq boot entry.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func bootDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Dnsmasq network boot entries (**Services ▸ Dnsmasq DNS & DHCP ▸ DHCP boot**) point PXE clients at a boot file and TFTP server.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the dnsmasq boot entry.",
				Required:            true,
			},
			"interface":      dschema.StringAttribute{MarkdownDescription: "Only apply on this interface.", Computed: true},
			"tags":           dschema.SetAttribute{MarkdownDescription: "UUIDs of tags a client must carry, e.g. one matching the client architecture.", Computed: true, ElementType: types.StringType},
			"filename":       dschema.StringAttribute{MarkdownDescription: "Boot file name, e.g. `pxelinux.0`.", Computed: true},
			"server_name":    dschema.StringAttribute{MarkdownDescription: "Name of the boot server.", Computed: true},
			"server_address": dschema.StringAttribute{MarkdownDescription: "Address of the TFTP server.", Computed: true},
			"description":    dschema.StringAttribute{MarkdownDescription: "Optional description here for your reference (not parsed).", Computed: true},
		},
	}
}

func convertBootSchemaToStruct(d *bootResourceModel) (*dnsmasq.Boot, error) {
	return &dnsmasq.Boot{
		Interface:     api.SelectedMap(d.Interface.ValueString()),
		Tag:           api.SelectedMapList(setToSlice(d.Tags)),
		Filename:      d.Filename.ValueString(),
		Servername:    d.ServerName.ValueString(),
		ServerAddress: d.ServerAddress.ValueString(),
		Description:   d.Description.ValueString(),
	}, nil
}

func convertBootStructToSchema(d *dnsmasq.Boot) (*bootResourceModel, error) {
	return &bootResourceModel{
		Interface:     tools.StringOrNull(d.Interface.String()),
		Tags:          tools.StringSliceToSet(d.Tag),
		Filename:      types.StringValue(d.Filename),
		ServerName:    tools.StringOrNull(d.Servername),
		ServerAddress: tools.StringOrNull(d.ServerAddress),
		Description:   tools.StringOrNull(d.Description),
	}, nil
}
