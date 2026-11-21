package proxy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/proxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type remoteBlacklistResourceModel struct {
	Enabled     types.Bool   `tfsdk:"enabled"`
	Filename    types.String `tfsdk:"filename"`
	URL         types.String `tfsdk:"url"`
	Username    types.String `tfsdk:"username"`
	Password    types.String `tfsdk:"password"`
	Filter      types.Set    `tfsdk:"filter"`
	SslNoVerify types.Bool   `tfsdk:"ssl_no_verify"`
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func remoteBlacklistResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Remote access control lists of the web proxy (**Services ▸ Squid Web Proxy ▸ Access Control List ▸ Remote ACLs**). The list is downloaded from `url` and denied by the proxy.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether this list is enabled. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"filename": schema.StringAttribute{
				MarkdownDescription: "Unique local file name of the list (letters and digits, optional extension).",
				Required:            true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "URL the list is downloaded from.",
				Required:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "User name for the download, if the server requires one.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for the download.",
				Optional:            true,
				Sensitive:           true,
			},
			"filter": schema.SetAttribute{
				MarkdownDescription: "Categories of the list to use. Empty uses all categories. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"ssl_no_verify": schema.BoolAttribute{
				MarkdownDescription: "Do not verify the TLS certificate of the download server. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the list. Letters, digits, spaces, `.`, `,` and `_` only.",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the remote blacklist.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func remoteBlacklistDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Remote access control lists of the web proxy (**Services ▸ Squid Web Proxy ▸ Access Control List ▸ Remote ACLs**). The list is downloaded from `url` and denied by the proxy.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the remote blacklist.",
				Required:            true,
			},
			"enabled":       dschema.BoolAttribute{MarkdownDescription: "Whether this list is enabled.", Computed: true},
			"filename":      dschema.StringAttribute{MarkdownDescription: "Unique local file name of the list (letters and digits, optional extension).", Computed: true},
			"url":           dschema.StringAttribute{MarkdownDescription: "URL the list is downloaded from.", Computed: true},
			"username":      dschema.StringAttribute{MarkdownDescription: "User name for the download, if the server requires one.", Computed: true},
			"password":      dschema.StringAttribute{MarkdownDescription: "Password for the download. Never returned by OPNsense, always null.", Computed: true, Sensitive: true},
			"filter":        dschema.SetAttribute{MarkdownDescription: "Categories of the list to use. Empty uses all categories.", Computed: true, ElementType: types.StringType},
			"ssl_no_verify": dschema.BoolAttribute{MarkdownDescription: "Do not verify the TLS certificate of the download server.", Computed: true},
			"description":   dschema.StringAttribute{MarkdownDescription: "Description of the list. Letters, digits, spaces, `.`, `,` and `_` only.", Computed: true},
		},
	}
}

func convertRemoteBlacklistSchemaToStruct(d *remoteBlacklistResourceModel) (*proxy.RemoteBlacklist, error) {
	return &proxy.RemoteBlacklist{
		Enabled:     tools.BoolToString(d.Enabled.ValueBool()),
		Filename:    d.Filename.ValueString(),
		URL:         d.URL.ValueString(),
		Username:    d.Username.ValueString(),
		Password:    d.Password.ValueString(),
		Filter:      api.SelectedMapList(setToSlice(d.Filter)),
		SSLNoVerify: tools.BoolToString(d.SslNoVerify.ValueBool()),
		Description: d.Description.ValueString(),
	}, nil
}

func convertRemoteBlacklistStructToSchema(d *proxy.RemoteBlacklist) (*remoteBlacklistResourceModel, error) {
	return &remoteBlacklistResourceModel{
		Enabled:     types.BoolValue(tools.StringToBool(d.Enabled)),
		Filename:    types.StringValue(d.Filename),
		URL:         types.StringValue(d.URL),
		Username:    tools.StringOrNull(d.Username),
		Password:    types.StringNull(),
		Filter:      tools.StringSliceToSet(d.Filter),
		SslNoVerify: types.BoolValue(tools.StringToBool(d.SSLNoVerify)),
		Description: types.StringValue(d.Description),
	}, nil
}
