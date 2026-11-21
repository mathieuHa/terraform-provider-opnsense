package haproxy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type resolverResourceModel struct {
	Enabled             types.Bool   `tfsdk:"enabled"`
	Name                types.String `tfsdk:"name"`
	Description         types.String `tfsdk:"description"`
	Nameservers         types.Set    `tfsdk:"nameservers"`
	ParseResolvConf     types.Bool   `tfsdk:"parse_resolv_conf"`
	ResolveRetries      types.Int64  `tfsdk:"resolve_retries"`
	TimeoutResolve      types.String `tfsdk:"timeout_resolve"`
	TimeoutRetry        types.String `tfsdk:"timeout_retry"`
	AcceptedPayloadSize types.Int64  `tfsdk:"accepted_payload_size"`
	HoldValid           types.String `tfsdk:"hold_valid"`
	HoldObsolete        types.String `tfsdk:"hold_obsolete"`
	HoldRefused         types.String `tfsdk:"hold_refused"`
	HoldNx              types.String `tfsdk:"hold_nx"`
	HoldTimeout         types.String `tfsdk:"hold_timeout"`
	HoldOther           types.String `tfsdk:"hold_other"`

	Id types.String `tfsdk:"id"`
}

func resolverResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "HAProxy DNS resolvers (**Services ▸ HAProxy ▸ Settings ▸ Advanced ▸ Resolvers**, plugin `os-haproxy`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this resolver configuration. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Choose a name for this resolver configuration.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Choose a optional description for this resolver configuration. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"nameservers": schema.SetAttribute{
				MarkdownDescription: "Add nameservers to this resolver configuration. They may be prefixed with either tcp@ or udp@ to use the TCP or UDP protocol respectively. For example 192.168.1.1:53, tcp@192.168.1.1:53 or udp@192.168.1.1:53. Use TAB key to complete typing. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"parse_resolv_conf": schema.BoolAttribute{
				MarkdownDescription: "Add all nameservers found in /etc/resolv.conf to this resolver configuration. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"resolve_retries": schema.Int64Attribute{
				MarkdownDescription: "This configures the number of queries to send to resolve a server name before giving up. Defaults to `3`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(3),
			},
			"timeout_resolve": schema.StringAttribute{
				MarkdownDescription: "This configures the default time to trigger name resolutions when no other time applied. Enter a number followed by one of the supported suffixes \"d\" (days), \"h\" (hour), \"m\" (minute), \"s\" (seconds), \"ms\" (miliseconds). Defaults to `1s`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("1s"),
			},
			"timeout_retry": schema.StringAttribute{
				MarkdownDescription: "This configures the default time between two DNS queries, when no valid response has been received. Enter a number followed by one of the supported suffixes \"d\" (days), \"h\" (hour), \"m\" (minute), \"s\" (seconds), \"ms\" (miliseconds). Defaults to `1s`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("1s"),
			},
			"accepted_payload_size": schema.Int64Attribute{
				MarkdownDescription: "Defines the maximum payload size accepted by HAProxy and announced to all the name servers configured in this resolvers section. The default is 512 bytes, the maximum allowed is 8192 for UDP and 65535 for TCP. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"hold_valid": schema.StringAttribute{
				MarkdownDescription: "When haproxy receives a valid NS response it will not query DNS until valid time expires. Default is \"10s\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"hold_obsolete": schema.StringAttribute{
				MarkdownDescription: "As a DNS server may not answer all the IPs in one DNS request, haproxy keeps a cache of previous answers. An answer will be considered obsolete after [hold obsolete] seconds without the IP returned. Default is \"30s\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"hold_refused": schema.StringAttribute{
				MarkdownDescription: "When the DNS server refuses the resolve request haproxy will not retry until [hold refused] elapses. Default \"30s\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"hold_nx": schema.StringAttribute{
				MarkdownDescription: "When haproxy receives a NXDOMAIN error message (domain does not exist) from the resolver it will not retry until [hold nx] elapses. Default \"30s\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"hold_timeout": schema.StringAttribute{
				MarkdownDescription: "When a DNS resolve request times out haproxy will not retry until [hold timeout] elapses. Default \"30s\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"hold_other": schema.StringAttribute{
				MarkdownDescription: "Sets the \"hold other\" timeout value for the resolver. Default \"30s\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the DNS resolver.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func resolverDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "HAProxy DNS resolvers (**Services ▸ HAProxy ▸ Settings ▸ Advanced ▸ Resolvers**, plugin `os-haproxy`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the DNS resolver.",
				Required:            true,
			},
			"enabled":               dschema.BoolAttribute{MarkdownDescription: "Enable this resolver configuration.", Computed: true},
			"name":                  dschema.StringAttribute{MarkdownDescription: "Choose a name for this resolver configuration.", Computed: true},
			"description":           dschema.StringAttribute{MarkdownDescription: "Choose a optional description for this resolver configuration.", Computed: true},
			"nameservers":           dschema.SetAttribute{MarkdownDescription: "Add nameservers to this resolver configuration. They may be prefixed with either tcp@ or udp@ to use the TCP or UDP protocol respectively. For example 192.168.1.1:53, tcp@192.168.1.1:53 or udp@192.168.1.1:53. Use TAB key to complete typing.", Computed: true, ElementType: types.StringType},
			"parse_resolv_conf":     dschema.BoolAttribute{MarkdownDescription: "Add all nameservers found in /etc/resolv.conf to this resolver configuration.", Computed: true},
			"resolve_retries":       dschema.Int64Attribute{MarkdownDescription: "This configures the number of queries to send to resolve a server name before giving up.", Computed: true},
			"timeout_resolve":       dschema.StringAttribute{MarkdownDescription: "This configures the default time to trigger name resolutions when no other time applied. Enter a number followed by one of the supported suffixes \"d\" (days), \"h\" (hour), \"m\" (minute), \"s\" (seconds), \"ms\" (miliseconds).", Computed: true},
			"timeout_retry":         dschema.StringAttribute{MarkdownDescription: "This configures the default time between two DNS queries, when no valid response has been received. Enter a number followed by one of the supported suffixes \"d\" (days), \"h\" (hour), \"m\" (minute), \"s\" (seconds), \"ms\" (miliseconds).", Computed: true},
			"accepted_payload_size": dschema.Int64Attribute{MarkdownDescription: "Defines the maximum payload size accepted by HAProxy and announced to all the name servers configured in this resolvers section. The default is 512 bytes, the maximum allowed is 8192 for UDP and 65535 for TCP.", Computed: true},
			"hold_valid":            dschema.StringAttribute{MarkdownDescription: "When haproxy receives a valid NS response it will not query DNS until valid time expires. Default is \"10s\".", Computed: true},
			"hold_obsolete":         dschema.StringAttribute{MarkdownDescription: "As a DNS server may not answer all the IPs in one DNS request, haproxy keeps a cache of previous answers. An answer will be considered obsolete after [hold obsolete] seconds without the IP returned. Default is \"30s\".", Computed: true},
			"hold_refused":          dschema.StringAttribute{MarkdownDescription: "When the DNS server refuses the resolve request haproxy will not retry until [hold refused] elapses. Default \"30s\".", Computed: true},
			"hold_nx":               dschema.StringAttribute{MarkdownDescription: "When haproxy receives a NXDOMAIN error message (domain does not exist) from the resolver it will not retry until [hold nx] elapses. Default \"30s\".", Computed: true},
			"hold_timeout":          dschema.StringAttribute{MarkdownDescription: "When a DNS resolve request times out haproxy will not retry until [hold timeout] elapses. Default \"30s\".", Computed: true},
			"hold_other":            dschema.StringAttribute{MarkdownDescription: "Sets the \"hold other\" timeout value for the resolver. Default \"30s\".", Computed: true},
		},
	}
}

func convertResolverSchemaToStruct(d *resolverResourceModel) (*haproxy.Resolver, error) {
	return &haproxy.Resolver{
		Enabled:             tools.BoolToString(d.Enabled.ValueBool()),
		Name:                d.Name.ValueString(),
		Description:         d.Description.ValueString(),
		Nameservers:         api.SelectedMapList(setToSlice(d.Nameservers)),
		ParseResolvConf:     tools.BoolToString(d.ParseResolvConf.ValueBool()),
		ResolveRetries:      tools.Int64ToString(d.ResolveRetries.ValueInt64()),
		TimeoutResolve:      d.TimeoutResolve.ValueString(),
		TimeoutRetry:        d.TimeoutRetry.ValueString(),
		AcceptedPayloadSize: tools.Int64ToStringNegative(d.AcceptedPayloadSize.ValueInt64()),
		HoldValid:           d.HoldValid.ValueString(),
		HoldObsolete:        d.HoldObsolete.ValueString(),
		HoldRefused:         d.HoldRefused.ValueString(),
		HoldNx:              d.HoldNx.ValueString(),
		HoldTimeout:         d.HoldTimeout.ValueString(),
		HoldOther:           d.HoldOther.ValueString(),
	}, nil
}

func convertResolverStructToSchema(d *haproxy.Resolver) (*resolverResourceModel, error) {
	return &resolverResourceModel{
		Enabled:             types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:                types.StringValue(d.Name),
		Description:         types.StringValue(d.Description),
		Nameservers:         tools.StringSliceToSet(d.Nameservers),
		ParseResolvConf:     types.BoolValue(tools.StringToBool(d.ParseResolvConf)),
		ResolveRetries:      types.Int64Value(tools.StringToInt64(d.ResolveRetries)),
		TimeoutResolve:      types.StringValue(d.TimeoutResolve),
		TimeoutRetry:        types.StringValue(d.TimeoutRetry),
		AcceptedPayloadSize: types.Int64Value(tools.StringToInt64(d.AcceptedPayloadSize)),
		HoldValid:           types.StringValue(d.HoldValid),
		HoldObsolete:        types.StringValue(d.HoldObsolete),
		HoldRefused:         types.StringValue(d.HoldRefused),
		HoldNx:              types.StringValue(d.HoldNx),
		HoldTimeout:         types.StringValue(d.HoldTimeout),
		HoldOther:           types.StringValue(d.HoldOther),
	}, nil
}
