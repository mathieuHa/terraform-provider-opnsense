package bind

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/bind"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type secondaryDomainResourceModel struct {
	Enabled              types.Bool   `tfsdk:"enabled"`
	Primaryip            types.Set    `tfsdk:"primaryip"`
	Transferkeyalgo      types.String `tfsdk:"transferkeyalgo"`
	Transferkeyname      types.String `tfsdk:"transferkeyname"`
	Transferkey          types.String `tfsdk:"transferkey"`
	Allownotifysecondary types.Set    `tfsdk:"allownotifysecondary"`
	Domainname           types.String `tfsdk:"domainname"`
	Allowtransfer        types.Set    `tfsdk:"allowtransfer"`
	Allowquery           types.Set    `tfsdk:"allowquery"`

	Id types.String `tfsdk:"id"`
}

func secondaryDomainResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Zones BIND copies from another server (**Services ▸ BIND ▸ Configuration ▸ Secondary Zones**, plugin `os-bind`).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "This will enable or disable this zone. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"primaryip": schema.SetAttribute{
				MarkdownDescription: "Set the IP address of primary server. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"transferkeyalgo": schema.StringAttribute{
				MarkdownDescription: "Set the authentication algorithm for the TSIG key used to transfer domain data from the primary server. One of `hmac-sha512`, `hmac-sha384`, `hmac-sha256`, `hmac-sha224`, `hmac-sha1`, `hmac-md5`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "hmac-sha512", "hmac-sha384", "hmac-sha256", "hmac-sha224", "hmac-sha1", "hmac-md5")},
			},
			"transferkeyname": schema.StringAttribute{
				MarkdownDescription: "The name of the TSIG key, which must match the value on the primary server. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"transferkey": schema.StringAttribute{
				MarkdownDescription: "The TSIG key used to transfer domain data from the master server in Base64 encoding. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"allownotifysecondary": schema.SetAttribute{
				MarkdownDescription: "A list of allowed IP addresses to receive notifies from (in addition to the primary server.). Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
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
			"allowquery": schema.SetAttribute{
				MarkdownDescription: "Define the ACLs where you allow which client are allowed to query this zone. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the BIND secondary zone.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func secondaryDomainDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Zones BIND copies from another server (**Services ▸ BIND ▸ Configuration ▸ Secondary Zones**, plugin `os-bind`).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the BIND secondary zone.",
				Required:            true,
			},
			"enabled":              dschema.BoolAttribute{MarkdownDescription: "This will enable or disable this zone.", Computed: true},
			"primaryip":            dschema.SetAttribute{MarkdownDescription: "Set the IP address of primary server.", Computed: true, ElementType: types.StringType},
			"transferkeyalgo":      dschema.StringAttribute{MarkdownDescription: "Set the authentication algorithm for the TSIG key used to transfer domain data from the primary server. One of `hmac-sha512`, `hmac-sha384`, `hmac-sha256`, `hmac-sha224`, `hmac-sha1`, `hmac-md5`.", Computed: true},
			"transferkeyname":      dschema.StringAttribute{MarkdownDescription: "The name of the TSIG key, which must match the value on the primary server.", Computed: true},
			"transferkey":          dschema.StringAttribute{MarkdownDescription: "The TSIG key used to transfer domain data from the master server in Base64 encoding.", Computed: true, Sensitive: true},
			"allownotifysecondary": dschema.SetAttribute{MarkdownDescription: "A list of allowed IP addresses to receive notifies from (in addition to the primary server.).", Computed: true, ElementType: types.StringType},
			"domainname":           dschema.StringAttribute{MarkdownDescription: "Set the name for this zone. Both forward and reverse zones may be specified, i.e. example.com or 0.168.192.in-addr.arpa.", Computed: true},
			"allowtransfer":        dschema.SetAttribute{MarkdownDescription: "Define the ACLs where you allow which server can retrieve this zone.", Computed: true, ElementType: types.StringType},
			"allowquery":           dschema.SetAttribute{MarkdownDescription: "Define the ACLs where you allow which client are allowed to query this zone.", Computed: true, ElementType: types.StringType},
		},
	}
}

func convertSecondaryDomainSchemaToStruct(d *secondaryDomainResourceModel) (*bind.SecondaryDomain, error) {
	return &bind.SecondaryDomain{
		Enabled:              tools.BoolToString(d.Enabled.ValueBool()),
		Primaryip:            api.SelectedMapList(setToSlice(d.Primaryip)),
		Transferkeyalgo:      api.SelectedMap(d.Transferkeyalgo.ValueString()),
		Transferkeyname:      d.Transferkeyname.ValueString(),
		Transferkey:          d.Transferkey.ValueString(),
		Allownotifysecondary: api.SelectedMapList(setToSlice(d.Allownotifysecondary)),
		Domainname:           d.Domainname.ValueString(),
		Allowtransfer:        api.SelectedMapList(setToSlice(d.Allowtransfer)),
		Allowquery:           api.SelectedMapList(setToSlice(d.Allowquery)),
	}, nil
}

func convertSecondaryDomainStructToSchema(d *bind.SecondaryDomain) (*secondaryDomainResourceModel, error) {
	return &secondaryDomainResourceModel{
		Enabled:              types.BoolValue(tools.StringToBool(d.Enabled)),
		Primaryip:            tools.StringSliceToSet(d.Primaryip),
		Transferkeyalgo:      types.StringValue(d.Transferkeyalgo.String()),
		Transferkeyname:      types.StringValue(d.Transferkeyname),
		Transferkey:          types.StringValue(d.Transferkey),
		Allownotifysecondary: tools.StringSliceToSet(d.Allownotifysecondary),
		Domainname:           types.StringValue(d.Domainname),
		Allowtransfer:        tools.StringSliceToSet(d.Allowtransfer),
		Allowquery:           tools.StringSliceToSet(d.Allowquery),
	}, nil
}
