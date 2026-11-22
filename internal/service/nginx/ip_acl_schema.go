package nginx

import (
	"context"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/nginx"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ipAclResourceModel struct {
	Description   types.String `tfsdk:"description"`
	DefaultAction types.String `tfsdk:"default_action"`
	Entries       types.List   `tfsdk:"entries"`

	Id types.String `tfsdk:"id"`
}

type ipAclEntryModel struct {
	Network types.String `tfsdk:"network"`
	Action  types.String `tfsdk:"action"`
}

var ipAclEntryTypes = map[string]attr.Type{
	"network": types.StringType,
	"action":  types.StringType,
}

func ipAclResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "nginx IP access lists (**Services ▸ Nginx ▸ Configuration ▸ Access ▸ IP ACLs**, plugin `os-nginx`), used by HTTP servers, locations and stream servers. Entries are evaluated in order.",
		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				MarkdownDescription: "Name of the ACL, shown in the drop-downs.",
				Required:            true,
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: "Action for clients matching no entry: `deny`, `allow`, or empty for none. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "deny", "allow")},
			},
			"entries": schema.ListNestedAttribute{
				MarkdownDescription: "ACL entries, in order.",
				Required:            true,
				Validators:          []validator.List{listvalidator.SizeAtLeast(1)},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network": schema.StringAttribute{
							MarkdownDescription: "Address or network, e.g. `10.0.0.0/8`.",
							Required:            true,
						},
						"action": schema.StringAttribute{
							MarkdownDescription: "`allow` or `deny`.",
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf("allow", "deny")},
						},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the ACL.",
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func ipAclDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "nginx IP access lists (**Services ▸ Nginx ▸ Configuration ▸ Access ▸ IP ACLs**, plugin `os-nginx`).",
		Attributes: map[string]dschema.Attribute{
			"id":             dschema.StringAttribute{MarkdownDescription: "UUID of the ACL.", Required: true},
			"description":    dschema.StringAttribute{MarkdownDescription: "Name of the ACL.", Computed: true},
			"default_action": dschema.StringAttribute{MarkdownDescription: "Action for clients matching no entry.", Computed: true},
			"entries": dschema.ListNestedAttribute{
				MarkdownDescription: "ACL entries, in order.",
				Computed:            true,
				NestedObject: dschema.NestedAttributeObject{
					Attributes: map[string]dschema.Attribute{
						"network": dschema.StringAttribute{MarkdownDescription: "Address or network.", Computed: true},
						"action":  dschema.StringAttribute{MarkdownDescription: "`allow` or `deny`.", Computed: true},
					},
				},
			},
		},
	}
}

func convertIpAclSchemaToStruct(d *ipAclResourceModel) (*nginx.IpAcl, error) {
	var entries []ipAclEntryModel
	d.Entries.ElementsAs(context.Background(), &entries, false)
	data := []nginx.IpAclEntry{}
	for _, e := range entries {
		data = append(data, nginx.IpAclEntry{Network: e.Network.ValueString(), Action: e.Action.ValueString()})
	}
	return &nginx.IpAcl{
		Description:   d.Description.ValueString(),
		DefaultAction: api.SelectedMap(d.DefaultAction.ValueString()),
		Data:          data,
	}, nil
}

func convertIpAclStructToSchema(d *nginx.IpAcl) (*ipAclResourceModel, error) {
	entries := []ipAclEntryModel{}
	for _, e := range d.Data {
		entries = append(entries, ipAclEntryModel{Network: types.StringValue(e.Network), Action: types.StringValue(e.Action)})
	}
	list, _ := types.ListValueFrom(context.Background(), types.ObjectType{AttrTypes: ipAclEntryTypes}, entries)
	return &ipAclResourceModel{
		Description:   types.StringValue(d.Description),
		DefaultAction: types.StringValue(d.DefaultAction.String()),
		Entries:       list,
	}, nil
}
