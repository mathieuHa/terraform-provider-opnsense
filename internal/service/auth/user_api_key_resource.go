package auth

import (
	"context"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/auth"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &userApiKeyResource{}
var _ resource.ResourceWithConfigure = &userApiKeyResource{}
var _ resource.ResourceWithImportState = &userApiKeyResource{}

type userApiKeyResourceModel struct {
	Username types.String `tfsdk:"username"`
	Key      types.String `tfsdk:"key"`
	Secret   types.String `tfsdk:"secret"`
	Id       types.String `tfsdk:"id"`
}

func newUserApiKeyResource() resource.Resource {
	return &userApiKeyResource{}
}

type userApiKeyResource struct {
	client opnsense.Client
}

func (r *userApiKeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_user_api_key"
}

func (r *userApiKeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "API key of a user (**System ▸ Access ▸ Users ▸ API keys**). OPNsense returns the secret only when the key is created: it is stored in the state, so keep the state protected, and it is unknown after an import. Destroying the resource revokes the key.",
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				MarkdownDescription: "Name of the user the key belongs to, e.g. `opnsense_auth_user.automation.name`.",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"key": schema.StringAttribute{
				MarkdownDescription: "API key.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"secret": schema.StringAttribute{
				MarkdownDescription: "API secret, only known when the key is created by this resource.",
				Computed:            true,
				Sensitive:           true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Same as `key`.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *userApiKeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	apiClient, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = opnsense.NewClient(apiClient)
}

func (r *userApiKeyResource) find(ctx context.Context, key string) (*auth.ApiKeyResult, error) {
	res, err := r.client.Auth().UserGetAllApiKeys(ctx)
	if err != nil {
		return nil, err
	}
	for _, row := range res.Rows {
		if row.Key == key {
			return &row, nil
		}
	}
	return nil, nil
}

func (r *userApiKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data userApiKeyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Auth().UserAddApiKey(ctx, data.Username.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create API key, got error: %s", err))
		return
	}
	if res.Result != "ok" || res.Key == "" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create API key for user %q (result %q); does the user exist?", data.Username.ValueString(), res.Result))
		return
	}
	data.Key = types.StringValue(res.Key)
	data.Secret = types.StringValue(res.Secret)
	data.Id = types.StringValue(res.Key)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *userApiKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data userApiKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	row, err := r.find(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read API keys, got error: %s", err))
		return
	}
	if row == nil {
		resp.State.RemoveResource(ctx)
		return
	}
	data.Username = types.StringValue(row.Username)
	data.Key = types.StringValue(row.Key)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *userApiKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data userApiKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *userApiKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data userApiKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	row, err := r.find(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read API keys, got error: %s", err))
		return
	}
	if row == nil {
		return
	}
	if _, err := r.client.Auth().UserDeleteApiKey(ctx, row.Id); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete API key, got error: %s", err))
	}
}

func (r *userApiKeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("secret"), types.StringNull())...)
}
