package ipsec

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/ipsec"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type settingsResourceModel struct {
	Enabled             types.Bool   `tfsdk:"enabled"`
	PassthroughNetworks types.Set    `tfsdk:"passthrough_networks"`
	IgnoreAcquireTs     types.Bool   `tfsdk:"ignore_acquire_ts"`
	InstallRoutes       types.Bool   `tfsdk:"install_routes"`
	CiscoUnity          types.Bool   `tfsdk:"cisco_unity"`
	MakeBeforeBreak     types.Bool   `tfsdk:"make_before_break"`
	PreferredOldsa      types.Bool   `tfsdk:"preferred_oldsa"`
	DisableVpnRules     types.Bool   `tfsdk:"disable_vpn_rules"`
	MaxIkev1Exchanges   types.Int64  `tfsdk:"max_ikev1_exchanges"`
	Threads             types.Int64  `tfsdk:"threads"`
	IkesaTableSize      types.Int64  `tfsdk:"ikesa_table_size"`
	IkesaTableSegments  types.Int64  `tfsdk:"ikesa_table_segments"`
	InitLimitHalfOpen   types.Int64  `tfsdk:"init_limit_half_open"`
	RetransmitTries     types.Int64  `tfsdk:"retransmit_tries"`
	RetransmitJitter    types.Int64  `tfsdk:"retransmit_jitter"`
	RetransmitLimit     types.Int64  `tfsdk:"retransmit_limit"`
	ReqidBase           types.Int64  `tfsdk:"reqid_base"`
	RetransmitTimeout   types.String `tfsdk:"retransmit_timeout"`
	RetransmitBase      types.String `tfsdk:"retransmit_base"`

	Id types.String `tfsdk:"id"`
}

func settingsResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Global IPsec settings (**VPN ▸ IPsec ▸ Advanced Settings**) and the IPsec on/off switch. This is a singleton: import it with the ID `ipsec_settings` before managing it. Attributes left unset keep their current value.",

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether IPsec is enabled (**VPN ▸ IPsec ▸ Connections ▸ Enable IPsec**). Leave unset to keep the current value.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"passthrough_networks": schema.SetAttribute{
				MarkdownDescription: "Networks excluded from IPsec processing, e.g. local networks covered by a wider tunnel.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
			},
			"ignore_acquire_ts": schema.BoolAttribute{
				MarkdownDescription: "Ignore the traffic selectors from the kernel's acquire events.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"install_routes": schema.BoolAttribute{
				MarkdownDescription: "Let charon install routes into a separate routing table for established IPsec tunnels.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"cisco_unity": schema.BoolAttribute{
				MarkdownDescription: "Send the Cisco Unity vendor ID payload (IKEv1 only).",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"make_before_break": schema.BoolAttribute{
				MarkdownDescription: "Initiate IKEv2 reauthentication with a make-before-break instead of a break-before-make scheme.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"preferred_oldsa": schema.BoolAttribute{
				MarkdownDescription: "Prefer older SAs over newer ones.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"disable_vpn_rules": schema.BoolAttribute{
				MarkdownDescription: "Do not add the automatic firewall rules for IPsec.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"max_ikev1_exchanges": schema.Int64Attribute{
				MarkdownDescription: "Maximum number of IKEv1 phase 2 exchanges per IKE SA.",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.Int64{int64validator.AtLeast(0)},
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"threads": schema.Int64Attribute{
				MarkdownDescription: "Number of worker threads in charon.",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.Int64{int64validator.AtLeast(0)},
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"ikesa_table_size": schema.Int64Attribute{
				MarkdownDescription: "Size of the IKE SA hash table.",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.Int64{int64validator.AtLeast(0)},
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"ikesa_table_segments": schema.Int64Attribute{
				MarkdownDescription: "Number of exclusively locked segments in the IKE SA hash table.",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.Int64{int64validator.AtLeast(0)},
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"init_limit_half_open": schema.Int64Attribute{
				MarkdownDescription: "Limit of half-open IKE SAs before new initiation requests are dropped.",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.Int64{int64validator.AtLeast(0)},
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"retransmit_tries": schema.Int64Attribute{
				MarkdownDescription: "Number of times to retransmit a packet before giving up.",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.Int64{int64validator.AtLeast(0)},
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"retransmit_jitter": schema.Int64Attribute{
				MarkdownDescription: "Maximum jitter in percent applied randomly to calculated retransmission timeouts.",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.Int64{int64validator.AtLeast(0)},
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"retransmit_limit": schema.Int64Attribute{
				MarkdownDescription: "Upper limit in seconds for the calculated retransmission timeout.",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.Int64{int64validator.AtLeast(0)},
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"reqid_base": schema.Int64Attribute{
				MarkdownDescription: "First reqid to assign automatically to children.",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.Int64{int64validator.AtLeast(0)},
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"retransmit_timeout": schema.StringAttribute{
				MarkdownDescription: "Timeout in seconds before sending the first retransmit.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"retransmit_base": schema.StringAttribute{
				MarkdownDescription: "Base of the exponential backoff for retransmits.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Always `ipsec_settings`.",
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func boolOrEmpty(b types.Bool) string {
	if b.IsNull() || b.IsUnknown() {
		return ""
	}
	return tools.BoolToString(b.ValueBool())
}

func intOrEmpty(i types.Int64) string {
	if i.IsNull() || i.IsUnknown() {
		return ""
	}
	return tools.Int64ToString(i.ValueInt64())
}

func setOrEmpty(s types.Set) []string {
	out := tools.SetToStringSlice(s)
	if out == nil {
		return []string{}
	}
	return out
}

func nonEmpty(in []string) []string {
	out := []string{}
	for _, v := range in {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

func convertSettingsSchemaToStruct(d *settingsResourceModel) *ipsec.Settings {
	return &ipsec.Settings{
		General: ipsec.SettingsGeneral{
			PreferredOldSA:      boolOrEmpty(d.PreferredOldsa),
			DisableVPNRules:     boolOrEmpty(d.DisableVpnRules),
			PassthroughNetworks: api.SelectedMapList(setOrEmpty(d.PassthroughNetworks)),
		},
		Charon: ipsec.SettingsCharon{
			IgnoreAcquireTS:    boolOrEmpty(d.IgnoreAcquireTs),
			InstallRoutes:      boolOrEmpty(d.InstallRoutes),
			CiscoUnity:         boolOrEmpty(d.CiscoUnity),
			MakeBeforeBreak:    boolOrEmpty(d.MakeBeforeBreak),
			MaxIKEv1Exchanges:  intOrEmpty(d.MaxIkev1Exchanges),
			Threads:            intOrEmpty(d.Threads),
			IKESATableSize:     intOrEmpty(d.IkesaTableSize),
			IKESATableSegments: intOrEmpty(d.IkesaTableSegments),
			InitLimitHalfOpen:  intOrEmpty(d.InitLimitHalfOpen),
			RetransmitTries:    intOrEmpty(d.RetransmitTries),
			RetransmitJitter:   intOrEmpty(d.RetransmitJitter),
			RetransmitLimit:    intOrEmpty(d.RetransmitLimit),
			ReqIDBase:          intOrEmpty(d.ReqidBase),
			RetransmitTimeout:  d.RetransmitTimeout.ValueString(),
			RetransmitBase:     d.RetransmitBase.ValueString(),
		},
	}
}

func convertSettingsStructToSchema(s *ipsec.Settings, enabled bool) *settingsResourceModel {
	return &settingsResourceModel{
		Enabled:             types.BoolValue(enabled),
		PassthroughNetworks: tools.StringSliceToSet(nonEmpty(s.General.PassthroughNetworks)),
		PreferredOldsa:      types.BoolValue(tools.StringToBool(s.General.PreferredOldSA)),
		DisableVpnRules:     types.BoolValue(tools.StringToBool(s.General.DisableVPNRules)),
		IgnoreAcquireTs:     types.BoolValue(tools.StringToBool(s.Charon.IgnoreAcquireTS)),
		InstallRoutes:       types.BoolValue(tools.StringToBool(s.Charon.InstallRoutes)),
		CiscoUnity:          types.BoolValue(tools.StringToBool(s.Charon.CiscoUnity)),
		MakeBeforeBreak:     types.BoolValue(tools.StringToBool(s.Charon.MakeBeforeBreak)),
		MaxIkev1Exchanges:   tools.StringToInt64Null(s.Charon.MaxIKEv1Exchanges),
		Threads:             tools.StringToInt64Null(s.Charon.Threads),
		IkesaTableSize:      tools.StringToInt64Null(s.Charon.IKESATableSize),
		IkesaTableSegments:  tools.StringToInt64Null(s.Charon.IKESATableSegments),
		InitLimitHalfOpen:   tools.StringToInt64Null(s.Charon.InitLimitHalfOpen),
		RetransmitTries:     tools.StringToInt64Null(s.Charon.RetransmitTries),
		RetransmitJitter:    tools.StringToInt64Null(s.Charon.RetransmitJitter),
		RetransmitLimit:     tools.StringToInt64Null(s.Charon.RetransmitLimit),
		ReqidBase:           tools.StringToInt64Null(s.Charon.ReqIDBase),
		RetransmitTimeout:   tools.StringOrNull(s.Charon.RetransmitTimeout),
		RetransmitBase:      tools.StringOrNull(s.Charon.RetransmitBase),
	}
}
