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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type mailerResourceModel struct {
	Enabled     types.Bool   `tfsdk:"enabled"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Mailservers types.Set    `tfsdk:"mailservers"`
	Sender      types.String `tfsdk:"sender"`
	Recipient   types.String `tfsdk:"recipient"`
	Loglevel    types.String `tfsdk:"loglevel"`
	Timeout     types.String `tfsdk:"timeout"`
	Hostname    types.String `tfsdk:"hostname"`

	Id types.String `tfsdk:"id"`
}

func mailerResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "HAProxy e-mail alerts (**Services ▸ HAProxy ▸ Settings ▸ Advanced ▸ E-Mail Alerts**, plugin `os-haproxy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this mailer configuration. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Choose a name for this mailer configuration.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Choose a optional description for this mailer configuration. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"mailservers": schema.SetAttribute{
				MarkdownDescription: "Add mailservers to this mailer configuration, i.e. 192.168.1.1:25. Use TAB key to complete typing.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"sender": schema.StringAttribute{
				MarkdownDescription: "Declare the from email address to be used in both the envelope and header of email alerts.",
				Required:            true,
			},
			"recipient": schema.StringAttribute{
				MarkdownDescription: "Declare both the recipient address in the envelope and to address in the header of email alerts.",
				Required:            true,
			},
			"loglevel": schema.StringAttribute{
				MarkdownDescription: "Declare the maximum log level of messages for which email alerts will be sent. This acts as a filter on the sending of email alerts. One of `emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`, `debug`. Defaults to `alert`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("alert"),
				Validators:          []validator.String{stringvalidator.OneOf("emerg", "alert", "crit", "err", "warning", "notice", "info", "debug")},
			},
			"timeout": schema.StringAttribute{
				MarkdownDescription: "Defines the time (in seconds) available for a mail/connection to be made and send to the mail server. Defaults to `30`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("30"),
			},
			"hostname": schema.StringAttribute{
				MarkdownDescription: "Declare the to hostname address to be used when communicating with mailers. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the e-mail alert.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func mailerDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "HAProxy e-mail alerts (**Services ▸ HAProxy ▸ Settings ▸ Advanced ▸ E-Mail Alerts**, plugin `os-haproxy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the e-mail alert.",
				Required:            true,
			},
			"enabled":     dschema.BoolAttribute{MarkdownDescription: "Enable this mailer configuration.", Computed: true},
			"name":        dschema.StringAttribute{MarkdownDescription: "Choose a name for this mailer configuration.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "Choose a optional description for this mailer configuration.", Computed: true},
			"mailservers": dschema.SetAttribute{MarkdownDescription: "Add mailservers to this mailer configuration, i.e. 192.168.1.1:25. Use TAB key to complete typing.", Computed: true, ElementType: types.StringType},
			"sender":      dschema.StringAttribute{MarkdownDescription: "Declare the from email address to be used in both the envelope and header of email alerts.", Computed: true},
			"recipient":   dschema.StringAttribute{MarkdownDescription: "Declare both the recipient address in the envelope and to address in the header of email alerts.", Computed: true},
			"loglevel":    dschema.StringAttribute{MarkdownDescription: "Declare the maximum log level of messages for which email alerts will be sent. This acts as a filter on the sending of email alerts. One of `emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`, `debug`.", Computed: true},
			"timeout":     dschema.StringAttribute{MarkdownDescription: "Defines the time (in seconds) available for a mail/connection to be made and send to the mail server.", Computed: true},
			"hostname":    dschema.StringAttribute{MarkdownDescription: "Declare the to hostname address to be used when communicating with mailers.", Computed: true},
		},
	}
}

func convertMailerSchemaToStruct(d *mailerResourceModel) (*haproxy.Mailer, error) {
	return &haproxy.Mailer{
		Enabled:     tools.BoolToString(d.Enabled.ValueBool()),
		Name:        d.Name.ValueString(),
		Description: d.Description.ValueString(),
		Mailservers: api.SelectedMapList(setToSlice(d.Mailservers)),
		Sender:      d.Sender.ValueString(),
		Recipient:   d.Recipient.ValueString(),
		Loglevel:    api.SelectedMap(d.Loglevel.ValueString()),
		Timeout:     d.Timeout.ValueString(),
		Hostname:    d.Hostname.ValueString(),
	}, nil
}

func convertMailerStructToSchema(d *haproxy.Mailer) (*mailerResourceModel, error) {
	return &mailerResourceModel{
		Enabled:     types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:        types.StringValue(d.Name),
		Description: types.StringValue(d.Description),
		Mailservers: tools.StringSliceToSet(d.Mailservers),
		Sender:      types.StringValue(d.Sender),
		Recipient:   types.StringValue(d.Recipient),
		Loglevel:    types.StringValue(d.Loglevel.String()),
		Timeout:     types.StringValue(d.Timeout),
		Hostname:    types.StringValue(d.Hostname),
	}, nil
}
