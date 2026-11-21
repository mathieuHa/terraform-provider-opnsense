package monit

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/monit"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type alertResourceModel struct {
	Enabled     types.Bool   `tfsdk:"enabled"`
	Recipient   types.String `tfsdk:"recipient"`
	NotOn       types.Bool   `tfsdk:"not_on"`
	Events      types.Set    `tfsdk:"events"`
	Format      types.String `tfsdk:"format"`
	Reminder    types.Int64  `tfsdk:"reminder"`
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func alertResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Monit alert recipients (**Services ▸ Monit ▸ Settings ▸ Alert Settings**).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether this alert is enabled. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"recipient": schema.StringAttribute{
				MarkdownDescription: "Email address that receives the alerts.",
				Required:            true,
			},
			"not_on": schema.BoolAttribute{
				MarkdownDescription: "When enabled, `events` lists the events that do **not** trigger an alert. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"events": schema.SetAttribute{
				MarkdownDescription: "Events that trigger (or, with `not_on`, do not trigger) an alert. Empty means all events. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"format": schema.StringAttribute{
				MarkdownDescription: "Custom mail format, see the Monit documentation.",
				Optional:            true,
			},
			"reminder": schema.Int64Attribute{
				MarkdownDescription: "Send a reminder every this many cycles while the error persists. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
				Validators:          []validator.Int64{int64validator.Any(int64validator.OneOf(-1), int64validator.AtLeast(0))},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the Monit alert.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func alertDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Monit alert recipients (**Services ▸ Monit ▸ Settings ▸ Alert Settings**).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the Monit alert.",
				Required:            true,
			},
			"enabled":     dschema.BoolAttribute{MarkdownDescription: "Whether this alert is enabled.", Computed: true},
			"recipient":   dschema.StringAttribute{MarkdownDescription: "Email address that receives the alerts.", Computed: true},
			"not_on":      dschema.BoolAttribute{MarkdownDescription: "When enabled, `events` lists the events that do **not** trigger an alert.", Computed: true},
			"events":      dschema.SetAttribute{MarkdownDescription: "Events that trigger (or, with `not_on`, do not trigger) an alert. Empty means all events.", Computed: true, ElementType: types.StringType},
			"format":      dschema.StringAttribute{MarkdownDescription: "Custom mail format, see the Monit documentation.", Computed: true},
			"reminder":    dschema.Int64Attribute{MarkdownDescription: "Send a reminder every this many cycles while the error persists.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Optional description here for your reference (not parsed).", Computed: true},
		},
	}
}

func convertAlertSchemaToStruct(d *alertResourceModel) (*monit.Alert, error) {
	return &monit.Alert{
		Enabled:     tools.BoolToString(d.Enabled.ValueBool()),
		Recipient:   d.Recipient.ValueString(),
		Noton:       tools.BoolToString(d.NotOn.ValueBool()),
		Events:      api.SelectedMapList(setToSlice(d.Events)),
		Format:      d.Format.ValueString(),
		Reminder:    tools.Int64ToStringNegative(d.Reminder.ValueInt64()),
		Description: d.Description.ValueString(),
	}, nil
}

func convertAlertStructToSchema(d *monit.Alert) (*alertResourceModel, error) {
	return &alertResourceModel{
		Enabled:     types.BoolValue(tools.StringToBool(d.Enabled)),
		Recipient:   types.StringValue(d.Recipient),
		NotOn:       types.BoolValue(tools.StringToBool(d.Noton)),
		Events:      tools.StringSliceToSet(d.Events),
		Format:      tools.StringOrNull(d.Format),
		Reminder:    types.Int64Value(tools.StringToInt64(d.Reminder)),
		Description: tools.StringOrNull(d.Description),
	}, nil
}
