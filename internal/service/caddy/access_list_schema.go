package caddy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/caddy"
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

type accessListResourceModel struct {
	Name                types.String `tfsdk:"name"`
	ClientIps           types.Set    `tfsdk:"client_ips"`
	Invert              types.Bool   `tfsdk:"invert"`
	HttpResponseCode    types.Int64  `tfsdk:"http_response_code"`
	HttpResponseMessage types.String `tfsdk:"http_response_message"`
	RequestMatcher      types.String `tfsdk:"request_matcher"`
	Description         types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func accessListResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Caddy access lists (**Services ▸ Caddy Web Server ▸ Reverse Proxy ▸ Access**, plugin `os-caddy`), used by domains, subdomains and handlers.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Enter a name for this access list.",
				Required:            true,
			},
			"client_ips": schema.SetAttribute{
				MarkdownDescription: "Enter the client IP addresses or networks for this access list. Accepts multiple IPv4/IPv6 addresses or networks.",
				ElementType:         types.StringType,
				Required:            true,
				Validators:          []validator.Set{setvalidator.SizeAtLeast(1)},
			},
			"invert": schema.BoolAttribute{
				MarkdownDescription: "If checked, the access list logic will be inverted (i.e., the listed IPs will be blocked instead of allowed). Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_response_code": schema.Int64Attribute{
				MarkdownDescription: "Set a custom HTTP response code that should be returned to the requesting client when the access list does not match. If empty, the connection is aborted with no response. Defaults to `-1` (OPNsense default).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"http_response_message": schema.StringAttribute{
				MarkdownDescription: "Set a custom HTTP response message in addition to the HTTP response code. If empty, the default HTTP response message for the chosen HTTP response code will be used. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"request_matcher": schema.StringAttribute{
				MarkdownDescription: "Set the request matcher that will be applied to this access list. One of `client_ip`, `remote_ip`. Defaults to `client_ip`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("client_ip"),
				Validators:          []validator.String{stringvalidator.OneOf("client_ip", "remote_ip")},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Enter a description for this access list. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the Caddy access list.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func accessListDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Caddy access lists (**Services ▸ Caddy Web Server ▸ Reverse Proxy ▸ Access**, plugin `os-caddy`), used by domains, subdomains and handlers.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the Caddy access list.",
				Required:            true,
			},
			"name":                  dschema.StringAttribute{MarkdownDescription: "Enter a name for this access list.", Computed: true},
			"client_ips":            dschema.SetAttribute{MarkdownDescription: "Enter the client IP addresses or networks for this access list. Accepts multiple IPv4/IPv6 addresses or networks.", Computed: true, ElementType: types.StringType},
			"invert":                dschema.BoolAttribute{MarkdownDescription: "If checked, the access list logic will be inverted (i.e., the listed IPs will be blocked instead of allowed).", Computed: true},
			"http_response_code":    dschema.Int64Attribute{MarkdownDescription: "Set a custom HTTP response code that should be returned to the requesting client when the access list does not match. If empty, the connection is aborted with no response.", Computed: true},
			"http_response_message": dschema.StringAttribute{MarkdownDescription: "Set a custom HTTP response message in addition to the HTTP response code. If empty, the default HTTP response message for the chosen HTTP response code will be used.", Computed: true},
			"request_matcher":       dschema.StringAttribute{MarkdownDescription: "Set the request matcher that will be applied to this access list. One of `client_ip`, `remote_ip`.", Computed: true},
			"description":           dschema.StringAttribute{MarkdownDescription: "Enter a description for this access list.", Computed: true},
		},
	}
}

func convertAccessListSchemaToStruct(d *accessListResourceModel) (*caddy.AccessList, error) {
	return &caddy.AccessList{
		AccesslistName:      d.Name.ValueString(),
		ClientIps:           api.SelectedMapList(setToSlice(d.ClientIps)),
		AccesslistInvert:    tools.BoolToString(d.Invert.ValueBool()),
		HttpResponseCode:    tools.Int64ToStringNegative(d.HttpResponseCode.ValueInt64()),
		HttpResponseMessage: d.HttpResponseMessage.ValueString(),
		RequestMatcher:      api.SelectedMap(d.RequestMatcher.ValueString()),
		Description:         d.Description.ValueString(),
	}, nil
}

func convertAccessListStructToSchema(d *caddy.AccessList) (*accessListResourceModel, error) {
	return &accessListResourceModel{
		Name:                types.StringValue(d.AccesslistName),
		ClientIps:           tools.StringSliceToSet(d.ClientIps),
		Invert:              types.BoolValue(tools.StringToBool(d.AccesslistInvert)),
		HttpResponseCode:    types.Int64Value(tools.StringToInt64(d.HttpResponseCode)),
		HttpResponseMessage: types.StringValue(d.HttpResponseMessage),
		RequestMatcher:      types.StringValue(d.RequestMatcher.String()),
		Description:         types.StringValue(d.Description),
	}, nil
}
