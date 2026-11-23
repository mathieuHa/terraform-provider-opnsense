package bind

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/bind"
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

type primaryDomainResourceModel struct {
	Enabled           types.Bool   `tfsdk:"enabled"`
	Domainname        types.String `tfsdk:"domainname"`
	Allowtransfer     types.Set    `tfsdk:"allowtransfer"`
	Allowrndctransfer types.Bool   `tfsdk:"allowrndctransfer"`
	Allowquery        types.Set    `tfsdk:"allowquery"`
	Allowrndcupdate   types.Bool   `tfsdk:"allowrndcupdate"`
	TTL               types.Int64  `tfsdk:"ttl"`
	Refresh           types.Int64  `tfsdk:"refresh"`
	Retry             types.Int64  `tfsdk:"retry"`
	Expire            types.Int64  `tfsdk:"expire"`
	Negative          types.Int64  `tfsdk:"negative"`
	Mailadmin         types.String `tfsdk:"mailadmin"`
	Dnsserver         types.String `tfsdk:"dnsserver"`

	Id types.String `tfsdk:"id"`
}

func primaryDomainResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Zones BIND is authoritative for (**Services ▸ BIND ▸ Configuration ▸ Primary Zones**, plugin `os-bind`). Records are `opnsense_bind_record` resources; BIND only loads a zone that has at least one NS record (`name = \"@\"`, `type = \"NS\"`), which OPNsense does not add by itself.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "This will enable or disable this zone. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"domainname": schema.StringAttribute{
				MarkdownDescription: "Set the name for this zone. Both forward and reverse zones may be specified, i.e. example.com or 0.168.192.in-addr.arpa.",
				Required:            true,
			},
			"allowtransfer": schema.SetAttribute{
				MarkdownDescription: "Define the ACLs where you allow which server can retrieve this zone. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"allowrndctransfer": schema.BoolAttribute{
				MarkdownDescription: "Allow transfers via the RDNC key named \"rndc-key\". The key is shown in the general tab. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allowquery": schema.SetAttribute{
				MarkdownDescription: "Define the ACLs where you allow which client are allowed to query this zone. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"allowrndcupdate": schema.BoolAttribute{
				MarkdownDescription: "Allow updates via the RDNC key named \"rndc-key\". The key is shown in the general tab. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ttl": schema.Int64Attribute{
				MarkdownDescription: "Set the general Time To Live for this zone. Defaults to `86400`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(86400),
			},
			"refresh": schema.Int64Attribute{
				MarkdownDescription: "Set the time in seconds after which name servers should refresh the zone information. Defaults to `21600`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(21600),
			},
			"retry": schema.Int64Attribute{
				MarkdownDescription: "Set the time in seconds after which name servers should retry requests if the master does not respond. Defaults to `3600`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(3600),
			},
			"expire": schema.Int64Attribute{
				MarkdownDescription: "Set the time in seconds after which name servers should stop answering requests if the master does not respond. Defaults to `3542400`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(3542400),
			},
			"negative": schema.Int64Attribute{
				MarkdownDescription: "Set the time in seconds after which an entry for a non-existent record should expire from cache. Defaults to `3600`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(3600),
			},
			"mailadmin": schema.StringAttribute{
				MarkdownDescription: "Set the mail address of zone admin. A @-sign will automatically be replaced with a dot in the zone data. Defaults to `mail.opnsense.localdomain`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("mail.opnsense.localdomain"),
			},
			"dnsserver": schema.StringAttribute{
				MarkdownDescription: "Set the DNS server hosting this file. This should usually be the FQDN of your firewall where the BIND plugin is installed. Defaults to `opnsense.localdomain`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("opnsense.localdomain"),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the BIND primary zone.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func primaryDomainDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Zones BIND is authoritative for (**Services ▸ BIND ▸ Configuration ▸ Primary Zones**, plugin `os-bind`). Records are `opnsense_bind_record` resources; BIND only loads a zone that has at least one NS record (`name = \"@\"`, `type = \"NS\"`), which OPNsense does not add by itself.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the BIND primary zone.",
				Required:            true,
			},
			"enabled":           dschema.BoolAttribute{MarkdownDescription: "This will enable or disable this zone.", Computed: true},
			"domainname":        dschema.StringAttribute{MarkdownDescription: "Set the name for this zone. Both forward and reverse zones may be specified, i.e. example.com or 0.168.192.in-addr.arpa.", Computed: true},
			"allowtransfer":     dschema.SetAttribute{MarkdownDescription: "Define the ACLs where you allow which server can retrieve this zone.", Computed: true, ElementType: types.StringType},
			"allowrndctransfer": dschema.BoolAttribute{MarkdownDescription: "Allow transfers via the RDNC key named \"rndc-key\". The key is shown in the general tab.", Computed: true},
			"allowquery":        dschema.SetAttribute{MarkdownDescription: "Define the ACLs where you allow which client are allowed to query this zone.", Computed: true, ElementType: types.StringType},
			"allowrndcupdate":   dschema.BoolAttribute{MarkdownDescription: "Allow updates via the RDNC key named \"rndc-key\". The key is shown in the general tab.", Computed: true},
			"ttl":               dschema.Int64Attribute{MarkdownDescription: "Set the general Time To Live for this zone.", Computed: true},
			"refresh":           dschema.Int64Attribute{MarkdownDescription: "Set the time in seconds after which name servers should refresh the zone information.", Computed: true},
			"retry":             dschema.Int64Attribute{MarkdownDescription: "Set the time in seconds after which name servers should retry requests if the master does not respond.", Computed: true},
			"expire":            dschema.Int64Attribute{MarkdownDescription: "Set the time in seconds after which name servers should stop answering requests if the master does not respond.", Computed: true},
			"negative":          dschema.Int64Attribute{MarkdownDescription: "Set the time in seconds after which an entry for a non-existent record should expire from cache.", Computed: true},
			"mailadmin":         dschema.StringAttribute{MarkdownDescription: "Set the mail address of zone admin. A @-sign will automatically be replaced with a dot in the zone data.", Computed: true},
			"dnsserver":         dschema.StringAttribute{MarkdownDescription: "Set the DNS server hosting this file. This should usually be the FQDN of your firewall where the BIND plugin is installed.", Computed: true},
		},
	}
}

func convertPrimaryDomainSchemaToStruct(d *primaryDomainResourceModel) (*bind.PrimaryDomain, error) {
	return &bind.PrimaryDomain{
		Enabled:           tools.BoolToString(d.Enabled.ValueBool()),
		Domainname:        d.Domainname.ValueString(),
		Allowtransfer:     api.SelectedMapList(setToSlice(d.Allowtransfer)),
		Allowrndctransfer: tools.BoolToString(d.Allowrndctransfer.ValueBool()),
		Allowquery:        api.SelectedMapList(setToSlice(d.Allowquery)),
		Allowrndcupdate:   tools.BoolToString(d.Allowrndcupdate.ValueBool()),
		Ttl:               tools.Int64ToString(d.TTL.ValueInt64()),
		Refresh:           tools.Int64ToString(d.Refresh.ValueInt64()),
		Retry:             tools.Int64ToString(d.Retry.ValueInt64()),
		Expire:            tools.Int64ToString(d.Expire.ValueInt64()),
		Negative:          tools.Int64ToString(d.Negative.ValueInt64()),
		Mailadmin:         d.Mailadmin.ValueString(),
		Dnsserver:         d.Dnsserver.ValueString(),
	}, nil
}

func convertPrimaryDomainStructToSchema(d *bind.PrimaryDomain) (*primaryDomainResourceModel, error) {
	return &primaryDomainResourceModel{
		Enabled:           types.BoolValue(tools.StringToBool(d.Enabled)),
		Domainname:        types.StringValue(d.Domainname),
		Allowtransfer:     tools.StringSliceToSet(d.Allowtransfer),
		Allowrndctransfer: types.BoolValue(tools.StringToBool(d.Allowrndctransfer)),
		Allowquery:        tools.StringSliceToSet(d.Allowquery),
		Allowrndcupdate:   types.BoolValue(tools.StringToBool(d.Allowrndcupdate)),
		TTL:               types.Int64Value(tools.StringToInt64(d.Ttl)),
		Refresh:           types.Int64Value(tools.StringToInt64(d.Refresh)),
		Retry:             types.Int64Value(tools.StringToInt64(d.Retry)),
		Expire:            types.Int64Value(tools.StringToInt64(d.Expire)),
		Negative:          types.Int64Value(tools.StringToInt64(d.Negative)),
		Mailadmin:         types.StringValue(d.Mailadmin),
		Dnsserver:         types.StringValue(d.Dnsserver),
	}, nil
}
