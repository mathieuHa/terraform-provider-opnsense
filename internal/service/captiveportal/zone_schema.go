package captiveportal

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/captiveportal"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type zoneResourceModel struct {
	Enabled                  types.Bool   `tfsdk:"enabled"`
	ZoneID                   types.Int64  `tfsdk:"zone_id"`
	Interfaces               types.Set    `tfsdk:"interfaces"`
	DisableRules             types.Bool   `tfsdk:"disable_rules"`
	AuthServers              types.Set    `tfsdk:"auth_servers"`
	Roaming                  types.Bool   `tfsdk:"roaming"`
	AlwaysSendAccountingReqs types.Bool   `tfsdk:"always_send_accounting_reqs"`
	AuthEnforceGroup         types.String `tfsdk:"auth_enforce_group"`
	IdleTimeout              types.Int64  `tfsdk:"idle_timeout"`
	HardTimeout              types.Int64  `tfsdk:"hard_timeout"`
	ConcurrentLogins         types.Bool   `tfsdk:"concurrent_logins"`
	Certificate              types.String `tfsdk:"certificate"`
	ServerName               types.String `tfsdk:"server_name"`
	AllowedAddresses         types.Set    `tfsdk:"allowed_addresses"`
	AllowedMACAddresses      types.Set    `tfsdk:"allowed_mac_addresses"`
	ExtendedPreAuthData      types.Bool   `tfsdk:"extended_pre_auth_data"`
	Template                 types.String `tfsdk:"template"`
	Description              types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func zoneResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Captive portal zones (**Services ▸ Captive Portal ▸ Administration**). Interfaces must have a static IPv4 address. Changes are applied by reloading the captive portal and ipfw.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether the zone is active. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"zone_id": schema.Int64Attribute{
				MarkdownDescription: "Zone number (0-19), assigned by OPNsense when not set.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
				Validators:          []validator.Int64{int64validator.Between(0, 19)},
			},
			"interfaces": schema.SetAttribute{
				MarkdownDescription: "Interfaces of the zone, e.g. `[\"opt1\"]`.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"disable_rules": schema.BoolAttribute{
				MarkdownDescription: "Do not add the automatic firewall rules for the portal; write them yourself. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"auth_servers": schema.SetAttribute{
				MarkdownDescription: "Authentication servers, e.g. `[\"Local Database\"]`. Empty lets everyone in after clicking through. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"roaming": schema.BoolAttribute{
				MarkdownDescription: "Keep a session when the client changes IP address (same MAC). Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"always_send_accounting_reqs": schema.BoolAttribute{
				MarkdownDescription: "Send RADIUS accounting requests for every authentication server, not only RADIUS ones. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"auth_enforce_group": schema.StringAttribute{
				MarkdownDescription: "GID of the local group users must belong to, empty for none. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"idle_timeout": schema.Int64Attribute{
				MarkdownDescription: "Minutes of inactivity before a client is disconnected, 0 for never. Defaults to `0`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators:          []validator.Int64{int64validator.Between(0, 10080)},
			},
			"hard_timeout": schema.Int64Attribute{
				MarkdownDescription: "Minutes after which a client is disconnected regardless of activity, 0 for never. Defaults to `0`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators:          []validator.Int64{int64validator.Between(0, 10080)},
			},
			"concurrent_logins": schema.BoolAttribute{
				MarkdownDescription: "Allow the same user to be logged in from several devices. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"certificate": schema.StringAttribute{
				MarkdownDescription: "Reference ID of the certificate used to serve the portal over HTTPS, empty for HTTP. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"server_name": schema.StringAttribute{
				MarkdownDescription: "Host name the portal redirects to, needed when using a certificate. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"allowed_addresses": schema.SetAttribute{
				MarkdownDescription: "Addresses or networks that bypass the portal. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"allowed_mac_addresses": schema.SetAttribute{
				MarkdownDescription: "MAC addresses that bypass the portal. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"extended_pre_auth_data": schema.BoolAttribute{
				MarkdownDescription: "Make the client IP and MAC available to the login page. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"template": schema.StringAttribute{
				MarkdownDescription: "UUID of an `opnsense_captiveportal_template`, empty for the default page. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the zone.",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the captive portal zone.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func zoneDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Captive portal zones (**Services ▸ Captive Portal ▸ Administration**). Interfaces must have a static IPv4 address. Changes are applied by reloading the captive portal and ipfw.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the captive portal zone.",
				Required:            true,
			},
			"enabled":                     dschema.BoolAttribute{MarkdownDescription: "Whether the zone is active.", Computed: true},
			"zone_id":                     dschema.Int64Attribute{MarkdownDescription: "Zone number (0-19), assigned by OPNsense when not set.", Computed: true},
			"interfaces":                  dschema.SetAttribute{MarkdownDescription: "Interfaces of the zone, e.g. `[\"opt1\"]`.", Computed: true, ElementType: types.StringType},
			"disable_rules":               dschema.BoolAttribute{MarkdownDescription: "Do not add the automatic firewall rules for the portal; write them yourself.", Computed: true},
			"auth_servers":                dschema.SetAttribute{MarkdownDescription: "Authentication servers, e.g. `[\"Local Database\"]`. Empty lets everyone in after clicking through.", Computed: true, ElementType: types.StringType},
			"roaming":                     dschema.BoolAttribute{MarkdownDescription: "Keep a session when the client changes IP address (same MAC).", Computed: true},
			"always_send_accounting_reqs": dschema.BoolAttribute{MarkdownDescription: "Send RADIUS accounting requests for every authentication server, not only RADIUS ones.", Computed: true},
			"auth_enforce_group":          dschema.StringAttribute{MarkdownDescription: "GID of the local group users must belong to, empty for none.", Computed: true},
			"idle_timeout":                dschema.Int64Attribute{MarkdownDescription: "Minutes of inactivity before a client is disconnected, 0 for never.", Computed: true},
			"hard_timeout":                dschema.Int64Attribute{MarkdownDescription: "Minutes after which a client is disconnected regardless of activity, 0 for never.", Computed: true},
			"concurrent_logins":           dschema.BoolAttribute{MarkdownDescription: "Allow the same user to be logged in from several devices.", Computed: true},
			"certificate":                 dschema.StringAttribute{MarkdownDescription: "Reference ID of the certificate used to serve the portal over HTTPS, empty for HTTP.", Computed: true},
			"server_name":                 dschema.StringAttribute{MarkdownDescription: "Host name the portal redirects to, needed when using a certificate.", Computed: true},
			"allowed_addresses":           dschema.SetAttribute{MarkdownDescription: "Addresses or networks that bypass the portal.", Computed: true, ElementType: types.StringType},
			"allowed_mac_addresses":       dschema.SetAttribute{MarkdownDescription: "MAC addresses that bypass the portal.", Computed: true, ElementType: types.StringType},
			"extended_pre_auth_data":      dschema.BoolAttribute{MarkdownDescription: "Make the client IP and MAC available to the login page.", Computed: true},
			"template":                    dschema.StringAttribute{MarkdownDescription: "UUID of an `opnsense_captiveportal_template`, empty for the default page.", Computed: true},
			"description":                 dschema.StringAttribute{MarkdownDescription: "Description of the zone.", Computed: true},
		},
	}
}

func convertZoneSchemaToStruct(d *zoneResourceModel) (*captiveportal.Zone, error) {
	return &captiveportal.Zone{
		Enabled: tools.BoolToString(d.Enabled.ValueBool()),
		ZoneId: func() string {
			if d.ZoneID.IsNull() || d.ZoneID.IsUnknown() {
				return ""
			}
			return tools.Int64ToString(d.ZoneID.ValueInt64())
		}(),
		Interfaces:               api.SelectedMapList(setToSlice(d.Interfaces)),
		DisableRules:             tools.BoolToString(d.DisableRules.ValueBool()),
		AuthServers:              api.SelectedMapList(setToSlice(d.AuthServers)),
		Roaming:                  tools.BoolToString(d.Roaming.ValueBool()),
		AlwaysSendAccountingReqs: tools.BoolToString(d.AlwaysSendAccountingReqs.ValueBool()),
		AuthEnforceGroup:         api.SelectedMap(d.AuthEnforceGroup.ValueString()),
		IdleTimeout:              tools.Int64ToString(d.IdleTimeout.ValueInt64()),
		HardTimeout:              tools.Int64ToString(d.HardTimeout.ValueInt64()),
		ConcurrentLogins:         tools.BoolToString(d.ConcurrentLogins.ValueBool()),
		Certificate:              api.SelectedMap(d.Certificate.ValueString()),
		ServerName:               d.ServerName.ValueString(),
		AllowedAddresses:         api.SelectedMapList(setToSlice(d.AllowedAddresses)),
		AllowedMACAddresses:      api.SelectedMapList(setToSlice(d.AllowedMACAddresses)),
		ExtendedPreAuthData:      tools.BoolToString(d.ExtendedPreAuthData.ValueBool()),
		Template:                 api.SelectedMap(d.Template.ValueString()),
		Description:              d.Description.ValueString(),
	}, nil
}

func convertZoneStructToSchema(d *captiveportal.Zone) (*zoneResourceModel, error) {
	return &zoneResourceModel{
		Enabled:                  types.BoolValue(tools.StringToBool(d.Enabled)),
		ZoneID:                   tools.StringToInt64Null(d.ZoneId),
		Interfaces:               tools.StringSliceToSet(d.Interfaces),
		DisableRules:             types.BoolValue(tools.StringToBool(d.DisableRules)),
		AuthServers:              tools.StringSliceToSet(d.AuthServers),
		Roaming:                  types.BoolValue(tools.StringToBool(d.Roaming)),
		AlwaysSendAccountingReqs: types.BoolValue(tools.StringToBool(d.AlwaysSendAccountingReqs)),
		AuthEnforceGroup:         types.StringValue(d.AuthEnforceGroup.String()),
		IdleTimeout:              types.Int64Value(tools.StringToInt64(d.IdleTimeout)),
		HardTimeout:              types.Int64Value(tools.StringToInt64(d.HardTimeout)),
		ConcurrentLogins:         types.BoolValue(tools.StringToBool(d.ConcurrentLogins)),
		Certificate:              types.StringValue(d.Certificate.String()),
		ServerName:               types.StringValue(d.ServerName),
		AllowedAddresses:         tools.StringSliceToSet(d.AllowedAddresses),
		AllowedMACAddresses:      tools.StringSliceToSet(d.AllowedMACAddresses),
		ExtendedPreAuthData:      types.BoolValue(tools.StringToBool(d.ExtendedPreAuthData)),
		Template:                 types.StringValue(d.Template.String()),
		Description:              types.StringValue(d.Description),
	}, nil
}

func setToSlice(s types.Set) []string {
	out := tools.SetToStringSlice(s)
	if out == nil {
		return []string{}
	}
	return out
}
