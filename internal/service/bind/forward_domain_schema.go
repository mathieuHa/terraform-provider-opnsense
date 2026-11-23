package bind

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/bind"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type forwardDomainResourceModel struct {
	Enabled       types.Bool   `tfsdk:"enabled"`
	Forwardserver types.Set    `tfsdk:"forwardserver"`
	Domainname    types.String `tfsdk:"domainname"`

	Id types.String `tfsdk:"id"`
}

func forwardDomainResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Zones whose queries BIND forwards to other servers (**Services ▸ BIND ▸ Configuration ▸ Forward Zones**, plugin `os-bind`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "This will enable or disable this zone. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"forwardserver": schema.SetAttribute{
				MarkdownDescription: "Set the IP address of server to forward requests to. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"domainname": schema.StringAttribute{
				MarkdownDescription: "Set the name for this zone. Both forward and reverse zones may be specified, i.e. example.com or 0.168.192.in-addr.arpa.",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the BIND forward zone.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func forwardDomainDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Zones whose queries BIND forwards to other servers (**Services ▸ BIND ▸ Configuration ▸ Forward Zones**, plugin `os-bind`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the BIND forward zone.",
				Required:            true,
			},
			"enabled":       dschema.BoolAttribute{MarkdownDescription: "This will enable or disable this zone.", Computed: true},
			"forwardserver": dschema.SetAttribute{MarkdownDescription: "Set the IP address of server to forward requests to.", Computed: true, ElementType: types.StringType},
			"domainname":    dschema.StringAttribute{MarkdownDescription: "Set the name for this zone. Both forward and reverse zones may be specified, i.e. example.com or 0.168.192.in-addr.arpa.", Computed: true},
		},
	}
}

func convertForwardDomainSchemaToStruct(d *forwardDomainResourceModel) (*bind.ForwardDomain, error) {
	return &bind.ForwardDomain{
		Enabled:       tools.BoolToString(d.Enabled.ValueBool()),
		Forwardserver: api.SelectedMapList(setToSlice(d.Forwardserver)),
		Domainname:    d.Domainname.ValueString(),
	}, nil
}

func convertForwardDomainStructToSchema(d *bind.ForwardDomain) (*forwardDomainResourceModel, error) {
	return &forwardDomainResourceModel{
		Enabled:       types.BoolValue(tools.StringToBool(d.Enabled)),
		Forwardserver: tools.StringSliceToSet(d.Forwardserver),
		Domainname:    types.StringValue(d.Domainname),
	}, nil
}
