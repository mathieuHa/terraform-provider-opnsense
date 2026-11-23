package dyndns

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dyndns"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type accountResourceModel struct {
	Enabled        types.Bool   `tfsdk:"enabled"`
	Service        types.String `tfsdk:"service"`
	Protocol       types.String `tfsdk:"protocol"`
	Server         types.String `tfsdk:"server"`
	Username       types.String `tfsdk:"username"`
	Password       types.String `tfsdk:"password"`
	ResourceID     types.String `tfsdk:"resource_id"`
	Hostnames      types.Set    `tfsdk:"hostnames"`
	Wildcard       types.Bool   `tfsdk:"wildcard"`
	Zone           types.String `tfsdk:"zone"`
	Checkip        types.String `tfsdk:"checkip"`
	Dynipv6host    types.String `tfsdk:"dynipv6host"`
	CheckipTimeout types.Int64  `tfsdk:"checkip_timeout"`
	ForceSsl       types.Bool   `tfsdk:"force_ssl"`
	TTL            types.Int64  `tfsdk:"ttl"`
	Interface      types.String `tfsdk:"interface"`
	Description    types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func accountResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Dynamic DNS accounts (**Services ▸ Dynamic DNS**, plugin `os-ddclient`): the provider, credentials and host names updated with the firewall's address.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this virtual server. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"service": schema.StringAttribute{
				MarkdownDescription: "Select the service to use. One of `hosting1984`, `changeip`, `cloudflare`, `cloudns`, `digitalocean`, `dinahosting`, `dnsmadeeasy`, `dns-o-matic`, `dnsexit2`, `dyndns2`, `dnspark`, `dnspodcn`, `dslreports1`, `dondominio`, `duckdns`, `dynu`, `easydns`, `freedns`, `freemyip`, `gandi`, `godaddy`, `googledomains`, `he-net`, `he-net-tunnel`, `hetzner`, `inwx`, `keysystems`, `loopia`, `mythicdyn`, `namecheap`, `nfsn`, `njalla`, `noip`, `nsupdatev4`, `nsupdatev6`, `ovh`, `porkbun`, `regfishde`, `servercow`, `sitelutions`, `spdyn`, `strato`, `woima`, `yandex`, `zoneedit1`, `custom`.",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf("hosting1984", "changeip", "cloudflare", "cloudns", "digitalocean", "dinahosting", "dnsmadeeasy", "dns-o-matic", "dnsexit2", "dyndns2", "dnspark", "dnspodcn", "dslreports1", "dondominio", "duckdns", "dynu", "easydns", "freedns", "freemyip", "gandi", "godaddy", "googledomains", "he-net", "he-net-tunnel", "hetzner", "inwx", "keysystems", "loopia", "mythicdyn", "namecheap", "nfsn", "njalla", "noip", "nsupdatev4", "nsupdatev6", "ovh", "porkbun", "regfishde", "servercow", "sitelutions", "spdyn", "strato", "woima", "yandex", "zoneedit1", "custom")},
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: "Select the protocol to use. One of `dyndns1`, `dyndns2`, `get`, `post`, `put`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "dyndns1", "dyndns2", "get", "post", "put")},
			},
			"server": schema.StringAttribute{
				MarkdownDescription: "DynDNS Server hostname or uri to use (depending on the protocol). When a URI is provided, the tag __MYIP__ will be replaced with the current detected address for this service and __HOSTNAME__ will contain the (comma separated) list of hostnames provided. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "Username or login to use. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password associated with this account.",
				Optional:            true,
				Sensitive:           true,
			},
			"resource_id": schema.StringAttribute{
				MarkdownDescription: "resourceId. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"hostnames": schema.SetAttribute{
				MarkdownDescription: "Hostname to update.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"wildcard": schema.BoolAttribute{
				MarkdownDescription: "add a DNS wildcard CNAME record that points to the configured host. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"zone": schema.StringAttribute{
				MarkdownDescription: "Zone containing the host entry. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"checkip": schema.StringAttribute{
				MarkdownDescription: "How the firewall's current address is found: `if` / `if6` (address of `interface`), or a web service such as `cloudflare-ipv4` or `web_ipify-ipv4`. The values depend on the backend chosen in `opnsense_dyndns_settings`.",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf("web_akamai", "akamai-ipv4", "akamai-ipv6", "web_cloudflare", "cloudflare-ipv4", "cloudflare-ipv6", "dns_cloudflare-dns", "dynu-ipv4", "dynu-ipv6", "web_freedns", "web_he", "web_icanhazip", "web_ipify-ipv4", "web_ipify-ipv6", "web_loopia", "web_myonlineportal", "web_noip-ipv4", "web_noip-ipv6", "web_nsupdate.info-ipv4", "web_nsupdate.info-ipv6", "web_zoneedit", "if")},
			},
			"dynipv6host": schema.StringAttribute{
				MarkdownDescription: "Swap the interface identifier of the ipv6 address with the given partial ipv6 address (the least significant 64 bits of the address). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"checkip_timeout": schema.Int64Attribute{
				MarkdownDescription: "How long to wait before the checkip process times out. Defaults to `10`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(10),
			},
			"force_ssl": schema.BoolAttribute{
				MarkdownDescription: "Force update using HTTPS, please note setting this option will enforce https updates on all accounts as ddclient only supports SSL=yes on a global level (the check ip service may still use HTTP on other services). Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ttl": schema.Int64Attribute{
				MarkdownDescription: "Time to Live for the DNS entry. Defaults to `300`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(300),
			},
			"interface": schema.StringAttribute{
				MarkdownDescription: "Interface to monitor. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the dynamic DNS account.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func accountDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Dynamic DNS accounts (**Services ▸ Dynamic DNS**, plugin `os-ddclient`): the provider, credentials and host names updated with the firewall's address.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the dynamic DNS account.",
				Required:            true,
			},
			"enabled":         dschema.BoolAttribute{MarkdownDescription: "Enable this virtual server.", Computed: true},
			"service":         dschema.StringAttribute{MarkdownDescription: "Select the service to use. One of `hosting1984`, `changeip`, `cloudflare`, `cloudns`, `digitalocean`, `dinahosting`, `dnsmadeeasy`, `dns-o-matic`, `dnsexit2`, `dyndns2`, `dnspark`, `dnspodcn`, `dslreports1`, `dondominio`, `duckdns`, `dynu`, `easydns`, `freedns`, `freemyip`, `gandi`, `godaddy`, `googledomains`, `he-net`, `he-net-tunnel`, `hetzner`, `inwx`, `keysystems`, `loopia`, `mythicdyn`, `namecheap`, `nfsn`, `njalla`, `noip`, `nsupdatev4`, `nsupdatev6`, `ovh`, `porkbun`, `regfishde`, `servercow`, `sitelutions`, `spdyn`, `strato`, `woima`, `yandex`, `zoneedit1`, `custom`.", Computed: true},
			"protocol":        dschema.StringAttribute{MarkdownDescription: "Select the protocol to use. One of `dyndns1`, `dyndns2`, `get`, `post`, `put`.", Computed: true},
			"server":          dschema.StringAttribute{MarkdownDescription: "DynDNS Server hostname or uri to use (depending on the protocol). When a URI is provided, the tag __MYIP__ will be replaced with the current detected address for this service and __HOSTNAME__ will contain the (comma separated) list of hostnames provided.", Computed: true},
			"username":        dschema.StringAttribute{MarkdownDescription: "Username or login to use.", Computed: true},
			"password":        dschema.StringAttribute{MarkdownDescription: "Password associated with this account. Never returned by OPNsense, always null.", Computed: true, Sensitive: true},
			"resource_id":     dschema.StringAttribute{MarkdownDescription: "resourceId.", Computed: true},
			"hostnames":       dschema.SetAttribute{MarkdownDescription: "Hostname to update.", Computed: true, ElementType: types.StringType},
			"wildcard":        dschema.BoolAttribute{MarkdownDescription: "add a DNS wildcard CNAME record that points to the configured host.", Computed: true},
			"zone":            dschema.StringAttribute{MarkdownDescription: "Zone containing the host entry.", Computed: true},
			"checkip":         dschema.StringAttribute{MarkdownDescription: "How the firewall's current address is found: `if` / `if6` (address of `interface`), or a web service such as `cloudflare-ipv4` or `web_ipify-ipv4`. The values depend on the backend chosen in `opnsense_dyndns_settings`.", Computed: true},
			"dynipv6host":     dschema.StringAttribute{MarkdownDescription: "Swap the interface identifier of the ipv6 address with the given partial ipv6 address (the least significant 64 bits of the address).", Computed: true},
			"checkip_timeout": dschema.Int64Attribute{MarkdownDescription: "How long to wait before the checkip process times out.", Computed: true},
			"force_ssl":       dschema.BoolAttribute{MarkdownDescription: "Force update using HTTPS, please note setting this option will enforce https updates on all accounts as ddclient only supports SSL=yes on a global level (the check ip service may still use HTTP on other services).", Computed: true},
			"ttl":             dschema.Int64Attribute{MarkdownDescription: "Time to Live for the DNS entry.", Computed: true},
			"interface":       dschema.StringAttribute{MarkdownDescription: "Interface to monitor.", Computed: true},
			"description":     dschema.StringAttribute{MarkdownDescription: "Description.", Computed: true},
		},
	}
}

func convertAccountSchemaToStruct(d *accountResourceModel) (*dyndns.Account, error) {
	return &dyndns.Account{
		Enabled:        tools.BoolToString(d.Enabled.ValueBool()),
		Service:        api.SelectedMap(d.Service.ValueString()),
		Protocol:       api.SelectedMap(d.Protocol.ValueString()),
		Server:         d.Server.ValueString(),
		Username:       d.Username.ValueString(),
		Password:       d.Password.ValueString(),
		ResourceId:     d.ResourceID.ValueString(),
		Hostnames:      api.SelectedMapList(setToSlice(d.Hostnames)),
		Wildcard:       tools.BoolToString(d.Wildcard.ValueBool()),
		Zone:           d.Zone.ValueString(),
		Checkip:        api.SelectedMap(d.Checkip.ValueString()),
		Dynipv6host:    d.Dynipv6host.ValueString(),
		CheckipTimeout: tools.Int64ToString(d.CheckipTimeout.ValueInt64()),
		ForceSsl:       tools.BoolToString(d.ForceSsl.ValueBool()),
		Ttl:            tools.Int64ToString(d.TTL.ValueInt64()),
		Interface:      api.SelectedMap(d.Interface.ValueString()),
		Description:    d.Description.ValueString(),
	}, nil
}

func convertAccountStructToSchema(d *dyndns.Account) (*accountResourceModel, error) {
	return &accountResourceModel{
		Enabled:        types.BoolValue(tools.StringToBool(d.Enabled)),
		Service:        types.StringValue(d.Service.String()),
		Protocol:       types.StringValue(d.Protocol.String()),
		Server:         types.StringValue(d.Server),
		Username:       types.StringValue(d.Username),
		Password:       types.StringNull(),
		ResourceID:     types.StringValue(d.ResourceId),
		Hostnames:      tools.StringSliceToSet(d.Hostnames),
		Wildcard:       types.BoolValue(tools.StringToBool(d.Wildcard)),
		Zone:           types.StringValue(d.Zone),
		Checkip:        types.StringValue(d.Checkip.String()),
		Dynipv6host:    types.StringValue(d.Dynipv6host),
		CheckipTimeout: types.Int64Value(tools.StringToInt64(d.CheckipTimeout)),
		ForceSsl:       types.BoolValue(tools.StringToBool(d.ForceSsl)),
		TTL:            types.Int64Value(tools.StringToInt64(d.Ttl)),
		Interface:      types.StringValue(d.Interface.String()),
		Description:    types.StringValue(d.Description),
	}, nil
}
