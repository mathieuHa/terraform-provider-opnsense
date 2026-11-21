package telegraf

import (
	"github.com/browningluke/opnsense-go/pkg/telegraf"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type keyResourceModel struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	Name    types.String `tfsdk:"name"`
	Value   types.String `tfsdk:"value"`

	Id types.String `tfsdk:"id"`
}

func keyResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Global tags added to every Telegraf metric (**Services ▸ Telegraf ▸ Tags**, plugin `os-telegraf`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "This will enable or disable the key mapping. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Set the name of this mapping. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "Set the value of the defined key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the Telegraf global tag.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func keyDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Global tags added to every Telegraf metric (**Services ▸ Telegraf ▸ Tags**, plugin `os-telegraf`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the Telegraf global tag.",
				Required:            true,
			},
			"enabled": dschema.BoolAttribute{MarkdownDescription: "This will enable or disable the key mapping.", Computed: true},
			"name":    dschema.StringAttribute{MarkdownDescription: "Set the name of this mapping.", Computed: true},
			"value":   dschema.StringAttribute{MarkdownDescription: "Set the value of the defined key.", Computed: true},
		},
	}
}

func convertKeySchemaToStruct(d *keyResourceModel) (*telegraf.Key, error) {
	return &telegraf.Key{
		Enabled: tools.BoolToString(d.Enabled.ValueBool()),
		Name:    d.Name.ValueString(),
		Value:   d.Value.ValueString(),
	}, nil
}

func convertKeyStructToSchema(d *telegraf.Key) (*keyResourceModel, error) {
	return &keyResourceModel{
		Enabled: types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:    types.StringValue(d.Name),
		Value:   types.StringValue(d.Value),
	}, nil
}
