package kea

import (
	"context"
	"strings"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/kea"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// dhcpv4SubnetResourceModel describes the resource data model.
type dhcpv4SubnetResourceModel struct {
	Subnet types.String `tfsdk:"subnet"`
	Pools  types.Set    `tfsdk:"pools"`

	MatchClientId types.Bool `tfsdk:"match_client_id"`

	AutoCollect types.Bool `tfsdk:"auto_collect"`

	Routers      types.Set `tfsdk:"routers"`
	StaticRoutes types.Set `tfsdk:"static_routes"`

	DomainNameServers types.Set    `tfsdk:"dns_servers"`
	DomainName        types.String `tfsdk:"domain_name"`
	DomainSearch      types.Set    `tfsdk:"domain_search"`

	NTPServers  types.Set `tfsdk:"ntp_servers"`
	TimeServers types.Set `tfsdk:"time_servers"`

	NextServer   types.String `tfsdk:"next_server"`
	TFTPServer   types.String `tfsdk:"tftp_server"`
	TFTPBootfile types.String `tfsdk:"tftp_bootfile"`

	ValidLifetime              types.Int64  `tfsdk:"valid_lifetime"`
	Allocator                  types.String `tfsdk:"allocator"`
	DdnsForwardZone            types.String `tfsdk:"ddns_forward_zone"`
	DdnsReverseZone            types.String `tfsdk:"ddns_reverse_zone"`
	DdnsQualifyingSuffix       types.String `tfsdk:"ddns_qualifying_suffix"`
	DdnsDnsServer              types.String `tfsdk:"ddns_dns_server"`
	DdnsDnsPort                types.Int64  `tfsdk:"ddns_dns_port"`
	DdnsDomainKeyName          types.String `tfsdk:"ddns_key_name"`
	DdnsDomainKeySecret        types.String `tfsdk:"ddns_key_secret"`
	DdnsDomainKeyAlgorithm     types.String `tfsdk:"ddns_key_algorithm"`
	DdnsOverrideNoUpdate       types.Bool   `tfsdk:"ddns_override_no_update"`
	DdnsOverrideClientUpdate   types.Bool   `tfsdk:"ddns_override_client_update"`
	DdnsUpdateOnRenew          types.Bool   `tfsdk:"ddns_update_on_renew"`
	DdnsConflictResolutionMode types.String `tfsdk:"ddns_conflict_resolution_mode"`

	Options types.Set `tfsdk:"options"`

	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func dhcpv4SubnetResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Configure DHCPv4 subnets for Kea.",

		Attributes: map[string]schema.Attribute{
			"subnet": schema.StringAttribute{
				MarkdownDescription: "Subnet to use (e.g. `\"192.0.2.64/26\"`), should be large enough to hold the specified pools and reservations.",
				Required:            true,
			},
			"pools": schema.SetAttribute{
				MarkdownDescription: "Set of pools in range or subnet format (e.g. `\"192.168.0.100 - 192.168.0.200\"` , `\"192.0.2.64/26\"`). Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"match_client_id": schema.BoolAttribute{
				MarkdownDescription: "By default, KEA uses client-identifiers instead of MAC addresses to locate clients, disabling this option changes back to matching on MAC address which is used by most dhcp implementations. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"auto_collect": schema.BoolAttribute{
				MarkdownDescription: "Automatically update option data from the GUI for relevant attributes. When set, values for `routers`, `dns_servers` and `ntp_servers` will be ignored. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"routers": schema.SetAttribute{
				MarkdownDescription: "Default gateways to offer to the clients. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"static_routes": schema.SetNestedAttribute{
				MarkdownDescription: "Static routes that the client should install in its routing cache. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"destination_ip": schema.StringAttribute{
							MarkdownDescription: "Destination IP address for static route.",
							Required:            true,
						},
						"router_ip": schema.StringAttribute{
							MarkdownDescription: "Gateway IP for static route.",
							Required:            true,
						},
					},
				},
				Default: setdefault.StaticValue(
					tools.EmptySetValue(
						types.ObjectType{
							AttrTypes: keaStaticRouteTypes,
						},
					),
				),
			},
			"dns_servers": schema.SetAttribute{
				MarkdownDescription: "DNS servers to offer to the clients. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"domain_name": schema.StringAttribute{
				MarkdownDescription: "Domain name to offer to the client, set to this firewall's domain name when left empty. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"domain_search": schema.SetAttribute{
				MarkdownDescription: "Set of Domain Names to be used by the client to locate not-fully-qualified domain names. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"ntp_servers": schema.SetAttribute{
				MarkdownDescription: "Set of IP addresses indicating NTP (RFC 5905) servers available to the client. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"time_servers": schema.SetAttribute{
				MarkdownDescription: "Set of RFC 868 time servers available to the client. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"next_server": schema.StringAttribute{
				MarkdownDescription: "Next server IP address. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tftp_server": schema.StringAttribute{
				MarkdownDescription: "TFTP server address or fqdn. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tftp_bootfile": schema.StringAttribute{
				MarkdownDescription: "Boot filename to request. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"options": schema.SetAttribute{
				MarkdownDescription: "UUIDs of `opnsense_kea_dhcpv4_option` resources sent to the clients. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"valid_lifetime": schema.Int64Attribute{
				MarkdownDescription: "Lease lifetime in seconds for this subnet, `-1` for the server default. Defaults to `-1`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"allocator": schema.StringAttribute{
				MarkdownDescription: "Address allocator: empty for the default, `iterative`, `random` or `flq` (free lease queue). Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "iterative", "random", "flq")},
			},
			"ddns_forward_zone": schema.StringAttribute{
				MarkdownDescription: "Forward DNS zone updated by Kea DDNS, e.g. `lan.example.`. Requires `ddns_dns_server` and an enabled `opnsense_kea_ddns_settings`. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ddns_reverse_zone": schema.StringAttribute{
				MarkdownDescription: "Reverse DNS zone, e.g. `1.168.192.in-addr.arpa.`. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ddns_qualifying_suffix": schema.StringAttribute{
				MarkdownDescription: "Suffix appended to unqualified client host names. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ddns_dns_server": schema.StringAttribute{
				MarkdownDescription: "Address of the DNS server that receives the updates. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ddns_dns_port": schema.Int64Attribute{
				MarkdownDescription: "Port of the DNS server, `-1` for 53. Defaults to `-1`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"ddns_key_name": schema.StringAttribute{
				MarkdownDescription: "TSIG key name. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ddns_key_secret": schema.StringAttribute{
				MarkdownDescription: "TSIG key secret, base64-encoded. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"ddns_key_algorithm": schema.StringAttribute{
				MarkdownDescription: "TSIG algorithm: `hmac-md5`, `hmac-sha1`, `hmac-sha224`, `hmac-sha256`, `hmac-sha384` or `hmac-sha512`. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "hmac-md5", "hmac-sha1", "hmac-sha224", "hmac-sha256", "hmac-sha384", "hmac-sha512")},
			},
			"ddns_override_no_update": schema.BoolAttribute{
				MarkdownDescription: "Update DNS even when the client asks for no update. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ddns_override_client_update": schema.BoolAttribute{
				MarkdownDescription: "Update DNS even when the client wants to do it itself. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ddns_update_on_renew": schema.BoolAttribute{
				MarkdownDescription: "Send DNS updates on every renewal, not only for new leases. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ddns_conflict_resolution_mode": schema.StringAttribute{
				MarkdownDescription: "Conflict resolution: empty for `check-with-dhcid`, `no-check-with-dhcid`, `check-exists-with-dhcid` or `no-check-without-dhcid`. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "no-check-with-dhcid", "check-exists-with-dhcid", "no-check-without-dhcid")},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the subnet.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func dhcpv4SubnetDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Configure DHCPv4 subnets for Kea.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"subnet": dschema.StringAttribute{
				MarkdownDescription: "Subnet in use (e.g. `\"192.0.2.64/26\"`).",
				Computed:            true,
			},
			"pools": dschema.SetAttribute{
				MarkdownDescription: "Set of pools in range or subnet format.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"match_client_id": dschema.BoolAttribute{
				MarkdownDescription: "Whether to match on client-identifier.",
				Computed:            true,
			},
			"auto_collect": dschema.BoolAttribute{
				MarkdownDescription: "Automatically update option data from the GUI for relevant attributes.",
				Computed:            true,
			},
			"routers": dschema.SetAttribute{
				MarkdownDescription: "Default gateways to offer to the clients.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"static_routes": dschema.SetNestedAttribute{
				MarkdownDescription: "Static routes that the client should install in its routing cache.",
				Computed:            true,
				NestedObject: dschema.NestedAttributeObject{
					Attributes: map[string]dschema.Attribute{
						"destination_ip": dschema.StringAttribute{
							MarkdownDescription: "Destination IP address for static route.",
							Computed:            true,
						},
						"router_ip": dschema.StringAttribute{
							MarkdownDescription: "Gateway IP for static route.",
							Computed:            true,
						},
					},
				},
			},
			"dns_servers": dschema.SetAttribute{
				MarkdownDescription: "DNS servers to offer to the clients.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"domain_name": dschema.StringAttribute{
				MarkdownDescription: "Domain name to offer to the client.",
				Computed:            true,
			},
			"domain_search": dschema.SetAttribute{
				MarkdownDescription: "Set of Domain Names to be used by the client to locate not-fully-qualified domain names.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"ntp_servers": dschema.SetAttribute{
				MarkdownDescription: "Set of IP addresses indicating NTP servers available to the client.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"time_servers": dschema.SetAttribute{
				MarkdownDescription: "Set of RFC 868 time servers available to the client.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"next_server": dschema.StringAttribute{
				MarkdownDescription: "Next server IP address.",
				Computed:            true,
			},
			"tftp_server": dschema.StringAttribute{
				MarkdownDescription: "TFTP server address or fqdn.",
				Computed:            true,
			},
			"tftp_bootfile": dschema.StringAttribute{
				MarkdownDescription: "Boot filename to request.",
				Computed:            true,
			},
			"options": dschema.SetAttribute{MarkdownDescription: "UUIDs of the custom options sent to the clients.", Computed: true, ElementType: types.StringType},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Computed:            true,
			},
			"valid_lifetime":                dschema.Int64Attribute{MarkdownDescription: "Lease lifetime in seconds for this subnet, `-1` for the server default.", Computed: true},
			"allocator":                     dschema.StringAttribute{MarkdownDescription: "Address allocator: empty for the default, `iterative`, `random` or `flq` (free lease queue).", Computed: true},
			"ddns_forward_zone":             dschema.StringAttribute{MarkdownDescription: "Forward DNS zone updated by Kea DDNS, e.g. `lan.example.`. Requires `ddns_dns_server` and an enabled `opnsense_kea_ddns_settings`.", Computed: true},
			"ddns_reverse_zone":             dschema.StringAttribute{MarkdownDescription: "Reverse DNS zone, e.g. `1.168.192.in-addr.arpa.`.", Computed: true},
			"ddns_qualifying_suffix":        dschema.StringAttribute{MarkdownDescription: "Suffix appended to unqualified client host names.", Computed: true},
			"ddns_dns_server":               dschema.StringAttribute{MarkdownDescription: "Address of the DNS server that receives the updates.", Computed: true},
			"ddns_dns_port":                 dschema.Int64Attribute{MarkdownDescription: "Port of the DNS server, `-1` for 53.", Computed: true},
			"ddns_key_name":                 dschema.StringAttribute{MarkdownDescription: "TSIG key name.", Computed: true},
			"ddns_key_secret":               dschema.StringAttribute{MarkdownDescription: "TSIG key secret, base64-encoded.", Computed: true, Sensitive: true},
			"ddns_key_algorithm":            dschema.StringAttribute{MarkdownDescription: "TSIG algorithm: `hmac-md5`, `hmac-sha1`, `hmac-sha224`, `hmac-sha256`, `hmac-sha384` or `hmac-sha512`.", Computed: true},
			"ddns_override_no_update":       dschema.BoolAttribute{MarkdownDescription: "Update DNS even when the client asks for no update.", Computed: true},
			"ddns_override_client_update":   dschema.BoolAttribute{MarkdownDescription: "Update DNS even when the client wants to do it itself.", Computed: true},
			"ddns_update_on_renew":          dschema.BoolAttribute{MarkdownDescription: "Send DNS updates on every renewal, not only for new leases.", Computed: true},
			"ddns_conflict_resolution_mode": dschema.StringAttribute{MarkdownDescription: "Conflict resolution: empty for `check-with-dhcid`, `no-check-with-dhcid`, `check-exists-with-dhcid` or `no-check-without-dhcid`.", Computed: true},
		},
	}
}

func convertDhcpv4SubnetSchemaToStruct(d *dhcpv4SubnetResourceModel) (*kea.SubnetV4, error) {
	var routesList []staticRouteModel
	d.StaticRoutes.ElementsAs(context.Background(), &routesList, false)

	var routePieceList []string
	for _, route := range routesList {
		routePieceList = append(routePieceList, route.DestinationIp.ValueString()+","+
			route.RouterIp.ValueString())
	}

	return &kea.SubnetV4{
		Subnet:                d.Subnet.ValueString(),
		NextServer:            d.NextServer.ValueString(),
		Pools:                 tools.SetToString(d.Pools, "\n"),
		MatchClientId:         tools.BoolToString(d.MatchClientId.ValueBool()),
		OptionDataAutoCollect: tools.BoolToString(d.AutoCollect.ValueBool()),
		OptionData: kea.OptionDataV4{
			DomainNameServers: tools.SetToStringSlice(d.DomainNameServers),
			DomainSearch:      tools.SetToStringSlice(d.DomainSearch),
			Routers:           tools.SetToStringSlice(d.Routers),
			StaticRoutes:      strings.Join(routePieceList, ";"),
			DomainName:        d.DomainName.ValueString(),
			NtpServers:        tools.SetToStringSlice(d.NTPServers),
			TimeServers:       tools.SetToStringSlice(d.TimeServers),
			TftpServerName:    d.TFTPServer.ValueString(),
			BootFileName:      d.TFTPBootfile.ValueString(),
		},
		ValidLifetime:              tools.Int64ToStringNegative(d.ValidLifetime.ValueInt64()),
		Allocator:                  api.SelectedMap(d.Allocator.ValueString()),
		DdnsForwardZone:            d.DdnsForwardZone.ValueString(),
		DdnsReverseZone:            d.DdnsReverseZone.ValueString(),
		DdnsQualifyingSuffix:       d.DdnsQualifyingSuffix.ValueString(),
		DdnsDnsServer:              d.DdnsDnsServer.ValueString(),
		DdnsDnsPort:                tools.Int64ToStringNegative(d.DdnsDnsPort.ValueInt64()),
		DdnsDomainKeyName:          d.DdnsDomainKeyName.ValueString(),
		DdnsDomainKeySecret:        d.DdnsDomainKeySecret.ValueString(),
		DdnsDomainKeyAlgorithm:     api.SelectedMap(d.DdnsDomainKeyAlgorithm.ValueString()),
		DdnsOverrideNoUpdate:       tools.BoolToString(d.DdnsOverrideNoUpdate.ValueBool()),
		DdnsOverrideClientUpdate:   tools.BoolToString(d.DdnsOverrideClientUpdate.ValueBool()),
		DdnsUpdateOnRenew:          tools.BoolToString(d.DdnsUpdateOnRenew.ValueBool()),
		DdnsConflictResolutionMode: api.SelectedMap(d.DdnsConflictResolutionMode.ValueString()),
		Options:                    optionsList(d.Options),
		Description:                d.Description.ValueString(),
	}, nil
}

func convertDhcpv4SubnetStructToSchema(d *kea.SubnetV4) (*dhcpv4SubnetResourceModel, error) {
	model := &dhcpv4SubnetResourceModel{
		Subnet:            types.StringValue(d.Subnet),
		Pools:             tools.StringSliceToSet(strings.Split(d.Pools, "\n")),
		MatchClientId:     types.BoolValue(tools.StringToBool(d.MatchClientId)),
		AutoCollect:       types.BoolValue(tools.StringToBool(d.OptionDataAutoCollect)),
		Routers:           tools.StringSliceToSet(d.OptionData.Routers),
		DomainNameServers: tools.StringSliceToSet(d.OptionData.DomainNameServers),
		DomainName:        types.StringValue(d.OptionData.DomainName),
		DomainSearch:      tools.StringSliceToSet(d.OptionData.DomainSearch),
		NTPServers:        tools.StringSliceToSet(d.OptionData.NtpServers),
		TimeServers:       tools.StringSliceToSet(d.OptionData.TimeServers),
		NextServer:        types.StringValue(d.NextServer),
		TFTPServer:        types.StringValue(d.OptionData.TftpServerName),
		TFTPBootfile:      types.StringValue(d.OptionData.BootFileName),
		Options:           tools.StringSliceToSet(d.Options),
		Description:       types.StringValue(d.Description),

		ValidLifetime:              types.Int64Value(tools.StringToInt64(d.ValidLifetime)),
		Allocator:                  types.StringValue(d.Allocator.String()),
		DdnsForwardZone:            types.StringValue(d.DdnsForwardZone),
		DdnsReverseZone:            types.StringValue(d.DdnsReverseZone),
		DdnsQualifyingSuffix:       types.StringValue(d.DdnsQualifyingSuffix),
		DdnsDnsServer:              types.StringValue(d.DdnsDnsServer),
		DdnsDnsPort:                types.Int64Value(tools.StringToInt64(d.DdnsDnsPort)),
		DdnsDomainKeyName:          types.StringValue(d.DdnsDomainKeyName),
		DdnsDomainKeySecret:        types.StringValue(d.DdnsDomainKeySecret),
		DdnsDomainKeyAlgorithm:     types.StringValue(d.DdnsDomainKeyAlgorithm.String()),
		DdnsOverrideNoUpdate:       types.BoolValue(tools.StringToBool(d.DdnsOverrideNoUpdate)),
		DdnsOverrideClientUpdate:   types.BoolValue(tools.StringToBool(d.DdnsOverrideClientUpdate)),
		DdnsUpdateOnRenew:          types.BoolValue(tools.StringToBool(d.DdnsUpdateOnRenew)),
		DdnsConflictResolutionMode: types.StringValue(d.DdnsConflictResolutionMode.String()),
	}

	v, _ := types.SetValue(
		types.ObjectType{
			AttrTypes: keaStaticRouteTypes,
		},
		[]attr.Value{},
	)

	routeSlice := strings.Split(d.OptionData.StaticRoutes, ";")
	routes := []staticRouteModel{}

	if len(routeSlice) > 0 {
		for _, i := range routeSlice {
			routePiece := strings.Split(i, ",")
			if len(routePiece) != 2 {
				continue
			}
			routes = append(routes, staticRouteModel{
				DestinationIp: types.StringValue(routePiece[0]),
				RouterIp:      types.StringValue(routePiece[1]),
			})
		}

		v, _ = types.SetValueFrom(
			context.Background(),
			types.ObjectType{
				AttrTypes: keaStaticRouteTypes,
			},
			routes,
		)
	}
	model.StaticRoutes = v

	return model, nil
}
