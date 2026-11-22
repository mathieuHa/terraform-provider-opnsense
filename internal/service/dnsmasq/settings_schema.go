package dnsmasq

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dnsmasq"
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
	Enable             types.Bool   `tfsdk:"enable"`
	Regdhcp            types.Bool   `tfsdk:"regdhcp"`
	Regdhcpstatic      types.Bool   `tfsdk:"regdhcpstatic"`
	Dhcpfirst          types.Bool   `tfsdk:"dhcpfirst"`
	StrictOrder        types.Bool   `tfsdk:"strict_order"`
	DomainNeeded       types.Bool   `tfsdk:"domain_needed"`
	NoPrivateReverse   types.Bool   `tfsdk:"no_private_reverse"`
	NoResolv           types.Bool   `tfsdk:"no_resolv"`
	LogQueries         types.Bool   `tfsdk:"log_queries"`
	NoHosts            types.Bool   `tfsdk:"no_hosts"`
	Strictbind         types.Bool   `tfsdk:"strictbind"`
	Dnssec             types.Bool   `tfsdk:"dnssec"`
	Regdhcpdomain      types.String `tfsdk:"regdhcpdomain"`
	Interface          types.Set    `tfsdk:"interface"`
	Port               types.Int64  `tfsdk:"port"`
	DnsForwardMax      types.Int64  `tfsdk:"dns_forward_max"`
	CacheSize          types.Int64  `tfsdk:"cache_size"`
	LocalTtl           types.Int64  `tfsdk:"local_ttl"`
	AddMac             types.String `tfsdk:"add_mac"`
	AddSubnet          types.Bool   `tfsdk:"add_subnet"`
	StripSubnet        types.Bool   `tfsdk:"strip_subnet"`
	DhcpNoInterface    types.Set    `tfsdk:"dhcp_no_interface"`
	DhcpFqdn           types.Bool   `tfsdk:"dhcp_fqdn"`
	DhcpDomain         types.String `tfsdk:"dhcp_domain"`
	DhcpLocal          types.Bool   `tfsdk:"dhcp_local"`
	DhcpLeaseMax       types.Int64  `tfsdk:"dhcp_lease_max"`
	DhcpAuthoritative  types.Bool   `tfsdk:"dhcp_authoritative"`
	DhcpDefaultFwRules types.Bool   `tfsdk:"dhcp_default_fw_rules"`
	DhcpReplyDelay     types.Int64  `tfsdk:"dhcp_reply_delay"`
	DhcpEnableRa       types.Bool   `tfsdk:"dhcp_enable_ra"`
	DhcpHostPing       types.Bool   `tfsdk:"dhcp_host_ping"`
	DhcpNosync         types.Bool   `tfsdk:"dhcp_nosync"`
	DhcpLogDhcp        types.Bool   `tfsdk:"dhcp_log_dhcp"`
	DhcpLogQuiet       types.Bool   `tfsdk:"dhcp_log_quiet"`
	NoIdent            types.Bool   `tfsdk:"no_ident"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enable": schema.BoolAttribute{
			MarkdownDescription: "Enable Dnsmasq.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"regdhcp": schema.BoolAttribute{
			MarkdownDescription: "If this option is set, then machines that specify their hostname when requesting a DHCP lease will be registered, so that their name can be resolved.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"regdhcpstatic": schema.BoolAttribute{
			MarkdownDescription: "If this option is set, then DHCP static mappings will be registered, so that their name can be resolved.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"dhcpfirst": schema.BoolAttribute{
			MarkdownDescription: "If this option is set, then DHCP mappings will be resolved before the manual list of names below. This only affects the name given for a reverse lookup (PTR).",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"strict_order": schema.BoolAttribute{
			MarkdownDescription: "If this option is set, we will query the DNS servers sequentially in the order specified (System: General Setup: DNS Servers), rather than all at once in parallel.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"domain_needed": schema.BoolAttribute{
			MarkdownDescription: "If this option is set, we will not forward A or AAAA queries for plain names, without dots or domain parts, to upstream name servers. If the name is not known from /etc/hosts or DHCP then a \"not found\" answer is returned.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"no_private_reverse": schema.BoolAttribute{
			MarkdownDescription: "If this option is set, we will not forward reverse DNS lookups (PTR) for private addresses (RFC 1918) to upstream name servers. Any entries in the Domain Overrides section forwarding private \"n.n.n.in-addr.arpa\" names to a specific server are still forwarded. If the IP to name is not known from /etc/hosts, DHCP or a specific domain override then a \"not found\" answer is immediately returned.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"no_resolv": schema.BoolAttribute{
			MarkdownDescription: "If this option is set, DNS forwarding to system nameservers (defined in System: General Setup: DNS Servers) will be disabled. Upstream servers defined in Services: Dnsmasq DNS & DHCP: Domains will still be used. This option is recommended when Unbound forwards local domain queries to Dnsmasq, so that all queries terminate without further lookups if they are unknown.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"log_queries": schema.BoolAttribute{
			MarkdownDescription: "Log the results of DNS queries.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"no_hosts": schema.BoolAttribute{
			MarkdownDescription: "Do not read hostnames in /etc/hosts.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"strictbind": schema.BoolAttribute{
			MarkdownDescription: "By default we bind the wildcard address, even when listening on some interfaces. Requests that should not be handled are discarded, this has the advantage of working even when interfaces come and go and change address. This option forces binding to only the interfaces we are listening on, which is less stable in non static environments.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"dnssec": schema.BoolAttribute{
			MarkdownDescription: "DNSSEC.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"regdhcpdomain": schema.StringAttribute{
			MarkdownDescription: "The domain name to use for DHCP hostname registration. If empty, the default system domain is used. Note that all DHCP leases will be assigned to the same domain. If this is undesired, static DHCP lease registration is able to provide coherent mappings.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"interface": schema.SetAttribute{
			MarkdownDescription: "Interface IPs used to respond to queries from clients. If no interfaces are selected, Dnsmasq will listen on all available IPv4 and IPv6 addresses by default. However, DHCP related firewall rules will only be added for explicitly selected interfaces, never for all interfaces.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"port": schema.Int64Attribute{
			MarkdownDescription: "The port used for responding to DNS queries. It should normally be left blank unless another service needs to bind to TCP/UDP port 53. Setting this to zero (0) completely disables DNS function. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"dns_forward_max": schema.Int64Attribute{
			MarkdownDescription: "Set the maximum number of concurrent DNS queries. On configurations with tight resources, this value may need to be reduced. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"cache_size": schema.Int64Attribute{
			MarkdownDescription: "Set the size of the cache. Setting the cache size to zero disables caching. Please note that huge cache size impacts performance. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"local_ttl": schema.Int64Attribute{
			MarkdownDescription: "This option allows a time-to-live (in seconds) to be given for local DNS entries, i.e. /etc/hosts or DHCP leases. This will reduce the load on the server at the expense of clients using stale data under some circumstances. A value of zero will disable client-side caching. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"add_mac": schema.StringAttribute{
			MarkdownDescription: "Add the MAC address of the requestor to DNS queries which are forwarded upstream. The MAC address will only be added if the upstream DNS Server is in the same subnet as the requestor. Since this is not standardized, it should be considered experimental. This is useful for selective DNS filtering on the upstream DNS server. One of `standard`, `base64`, `text`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"add_subnet": schema.BoolAttribute{
			MarkdownDescription: "Add the real client IPv4 and IPv6 addresses (add-subnet=32,128) to DNS queries which are forwarded upstream. Be careful setting this option as it can undermine privacy. This is useful for selective DNS filtering on the upstream DNS server.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"strip_subnet": schema.BoolAttribute{
			MarkdownDescription: "Strip the subnet received by a downstream DNS server. If \"Add subnet\" is used and the downstream DNS server already added a subnet, Dnsmasq will not replace it without setting \"Strip subnet\".",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"dhcp_no_interface": schema.SetAttribute{
			MarkdownDescription: "Do not provide DHCP, TFTP or router advertisement on the specified interfaces, but do provide DNS service. Please note that Dnsmasq continues to listen on the default DHCP ports as long as any DHCP ranges are configured; setting this option only ignores these packets on the selected interfaces.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"dhcp_fqdn": schema.BoolAttribute{
			MarkdownDescription: "If disabled, registers the unqualified names of DHCP clients into the DNS (e.g. 'smartphone'), in which case they should be unique. If enabled, the qualified name (e.g. 'smartphone.lan.internal') is registered. This option must be enabled if you are forwarding from Unbound to Dnsmasq for specific local domains.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"dhcp_domain": schema.StringAttribute{
			MarkdownDescription: "To ensure that all names have a domain part, there must be a default domain specified when DHCP FQDN is set. Leave empty to use the system domain.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"dhcp_local": schema.BoolAttribute{
			MarkdownDescription: "Sets all DHCP domains as local. This will configure this DNS server as authoritative; it will not forward queries to any upstream servers for these domains.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"dhcp_lease_max": schema.Int64Attribute{
			MarkdownDescription: "Limits Dnsmasq to the specified maximum number of DHCP leases. This limit is to prevent DoS attacks from hosts which create thousands of leases and use lots of memory in the Dnsmasq process. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"dhcp_authoritative": schema.BoolAttribute{
			MarkdownDescription: "Should be set when Dnsmasq is definitely the only DHCP server on a network. For DHCPv4, it changes the behaviour from strict RFC compliance so that DHCP requests on unknown leases from unknown hosts are not ignored.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"dhcp_default_fw_rules": schema.BoolAttribute{
			MarkdownDescription: "Automatically register firewall rules to allow DHCP traffic for all explicitly selected interfaces, can be disabled for more fine grained control if needed.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"dhcp_reply_delay": schema.Int64Attribute{
			MarkdownDescription: "Delays sending DHCPOFFER and PROXYDHCP replies for at least the specified number of seconds. This can be practical for split DHCP solutions, to make sure the secondary server answers slower than the primary. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"dhcp_enable_ra": schema.BoolAttribute{
			MarkdownDescription: "Setting this will enable Router Advertisements for all configured DHCPv6 ranges with the managed address bits set, and the use SLAAC bit reset. To change this default, select a combination of the possible options in the individual DHCPv6 ranges. Keep in mind that this is a global option; if there are configured DHCPv6 ranges, RAs will be sent unconditionally and cannot be deactivated selectively. Setting Router Advertisement modes in DHCPv6 ranges will have no effect without this global option enabled.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"dhcp_host_ping": schema.BoolAttribute{
			MarkdownDescription: "By default, the DHCP server will use a ping to ensure that an address is not in use before allocating it to a host.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"dhcp_nosync": schema.BoolAttribute{
			MarkdownDescription: "Ignore the DHCP general settings from being updated using HA sync.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"dhcp_log_dhcp": schema.BoolAttribute{
			MarkdownDescription: "Extra logging for DHCP, log all the options sent to DHCP clients and the tags used to determine them.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"dhcp_log_quiet": schema.BoolAttribute{
			MarkdownDescription: "Suppress logging of the routine operation of DHCP, RA and TFTP. Errors and problems will still be logged.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"no_ident": schema.BoolAttribute{
			MarkdownDescription: "Do not respond to class CHAOS and type TXT in domain bind queries. Without this option being set, the cache statistics are also available in the DNS as answers to queries of class CHAOS and type TXT in domain bind.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `dnsmasq_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertSettingsSchemaToStruct(d *settingsResourceModel) *dnsmasq.DnsmasqConfig {
	return &dnsmasq.DnsmasqConfig{
		Enable:           settingsBool(d.Enable),
		Regdhcp:          settingsBool(d.Regdhcp),
		Regdhcpstatic:    settingsBool(d.Regdhcpstatic),
		Dhcpfirst:        settingsBool(d.Dhcpfirst),
		StrictOrder:      settingsBool(d.StrictOrder),
		DomainNeeded:     settingsBool(d.DomainNeeded),
		NoPrivateReverse: settingsBool(d.NoPrivateReverse),
		NoResolv:         settingsBool(d.NoResolv),
		LogQueries:       settingsBool(d.LogQueries),
		NoHosts:          settingsBool(d.NoHosts),
		Strictbind:       settingsBool(d.Strictbind),
		Dnssec:           settingsBool(d.Dnssec),
		Regdhcpdomain:    d.Regdhcpdomain.ValueString(),
		Interface:        api.SelectedMapList(settingsList(d.Interface)),
		Port:             settingsInt(d.Port),
		DnsForwardMax:    settingsInt(d.DnsForwardMax),
		CacheSize:        settingsInt(d.CacheSize),
		LocalTtl:         settingsInt(d.LocalTtl),
		AddMac:           api.SelectedMap(d.AddMac.ValueString()),
		AddSubnet:        settingsBool(d.AddSubnet),
		StripSubnet:      settingsBool(d.StripSubnet),
		NoIdent:          settingsBool(d.NoIdent),
		Dhcp: dnsmasq.DnsmasqConfigDhcp{
			DhcpNoInterface:    api.SelectedMapList(settingsList(d.DhcpNoInterface)),
			DhcpFqdn:           settingsBool(d.DhcpFqdn),
			DhcpDomain:         d.DhcpDomain.ValueString(),
			DhcpLocal:          settingsBool(d.DhcpLocal),
			DhcpLeaseMax:       settingsInt(d.DhcpLeaseMax),
			DhcpAuthoritative:  settingsBool(d.DhcpAuthoritative),
			DhcpDefaultFwRules: settingsBool(d.DhcpDefaultFwRules),
			DhcpReplyDelay:     settingsInt(d.DhcpReplyDelay),
			DhcpEnableRa:       settingsBool(d.DhcpEnableRa),
			DhcpHostPing:       settingsBool(d.DhcpHostPing),
			DhcpNosync:         settingsBool(d.DhcpNosync),
			DhcpLogDhcp:        settingsBool(d.DhcpLogDhcp),
			DhcpLogQuiet:       settingsBool(d.DhcpLogQuiet),
		},
	}
}

func convertSettingsStructToSchema(s *dnsmasq.DnsmasqConfig) *settingsResourceModel {
	return &settingsResourceModel{
		Enable:             types.BoolValue(tools.StringToBool(s.Enable)),
		Regdhcp:            types.BoolValue(tools.StringToBool(s.Regdhcp)),
		Regdhcpstatic:      types.BoolValue(tools.StringToBool(s.Regdhcpstatic)),
		Dhcpfirst:          types.BoolValue(tools.StringToBool(s.Dhcpfirst)),
		StrictOrder:        types.BoolValue(tools.StringToBool(s.StrictOrder)),
		DomainNeeded:       types.BoolValue(tools.StringToBool(s.DomainNeeded)),
		NoPrivateReverse:   types.BoolValue(tools.StringToBool(s.NoPrivateReverse)),
		NoResolv:           types.BoolValue(tools.StringToBool(s.NoResolv)),
		LogQueries:         types.BoolValue(tools.StringToBool(s.LogQueries)),
		NoHosts:            types.BoolValue(tools.StringToBool(s.NoHosts)),
		Strictbind:         types.BoolValue(tools.StringToBool(s.Strictbind)),
		Dnssec:             types.BoolValue(tools.StringToBool(s.Dnssec)),
		Regdhcpdomain:      types.StringValue(s.Regdhcpdomain),
		Interface:          tools.StringSliceToSet(settingsNonEmpty(s.Interface)),
		Port:               settingsReadInt(s.Port),
		DnsForwardMax:      settingsReadInt(s.DnsForwardMax),
		CacheSize:          settingsReadInt(s.CacheSize),
		LocalTtl:           settingsReadInt(s.LocalTtl),
		AddMac:             types.StringValue(s.AddMac.String()),
		AddSubnet:          types.BoolValue(tools.StringToBool(s.AddSubnet)),
		StripSubnet:        types.BoolValue(tools.StringToBool(s.StripSubnet)),
		DhcpNoInterface:    tools.StringSliceToSet(settingsNonEmpty(s.Dhcp.DhcpNoInterface)),
		DhcpFqdn:           types.BoolValue(tools.StringToBool(s.Dhcp.DhcpFqdn)),
		DhcpDomain:         types.StringValue(s.Dhcp.DhcpDomain),
		DhcpLocal:          types.BoolValue(tools.StringToBool(s.Dhcp.DhcpLocal)),
		DhcpLeaseMax:       settingsReadInt(s.Dhcp.DhcpLeaseMax),
		DhcpAuthoritative:  types.BoolValue(tools.StringToBool(s.Dhcp.DhcpAuthoritative)),
		DhcpDefaultFwRules: types.BoolValue(tools.StringToBool(s.Dhcp.DhcpDefaultFwRules)),
		DhcpReplyDelay:     settingsReadInt(s.Dhcp.DhcpReplyDelay),
		DhcpEnableRa:       types.BoolValue(tools.StringToBool(s.Dhcp.DhcpEnableRa)),
		DhcpHostPing:       types.BoolValue(tools.StringToBool(s.Dhcp.DhcpHostPing)),
		DhcpNosync:         types.BoolValue(tools.StringToBool(s.Dhcp.DhcpNosync)),
		DhcpLogDhcp:        types.BoolValue(tools.StringToBool(s.Dhcp.DhcpLogDhcp)),
		DhcpLogQuiet:       types.BoolValue(tools.StringToBool(s.Dhcp.DhcpLogQuiet)),
		NoIdent:            types.BoolValue(tools.StringToBool(s.NoIdent)),
	}
}
