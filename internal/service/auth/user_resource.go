package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/errs"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = &userResource{}
var _ resource.ResourceWithConfigure = &userResource{}
var _ resource.ResourceWithImportState = &userResource{}

func newUserResource() resource.Resource {
	return &userResource{}
}

type userResource struct {
	client opnsense.Client
}

func (r *userResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_user"
}

func (r *userResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = userResourceSchema()
}

func (r *userResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	apiClient, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *opnsense.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = opnsense.NewClient(apiClient)
}

func (r *userResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *userResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	obj, err := convertUserSchemaToStruct(data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to parse user, got error: %s", err))
		return
	}

	id, err := r.client.Auth().AddUser(ctx, obj)
	if err != nil {
		if id != "" {
			data.Id = types.StringValue(id)

			if readStruct, readErr := r.client.Auth().GetUser(ctx, id); readErr == nil {
				if readModel, convErr := convertUserStructToSchema(readStruct); convErr == nil {
					readModel.Id = data.Id
					data = readModel
				}
			}

			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to create user, got error: %s", err))
		return
	}

	data.Id = types.StringValue(id)

	readStruct, err := r.client.Auth().GetUser(ctx, id)
	if err != nil {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read back user, got error: %s", err))
		return
	}
	readModel, err := convertUserStructToSchema(readStruct)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read back user, got error: %s", err))
		return
	}
	readModel.Id = data.Id
	readModel.Password = data.Password
	readModel.ScrambledPassword = data.ScrambledPassword
	data = readModel

	tflog.Trace(ctx, "created a resource")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *userResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *userResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	obj, err := r.client.Auth().GetUser(ctx, data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, fmt.Sprintf("user not present in remote, removing from state"))
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read user, got error: %s", err))
		return
	}

	userModel, err := convertUserStructToSchema(obj)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read user, got error: %s", err))
		return
	}

	userModel.Id = data.Id
	userModel.Password = data.Password
	userModel.ScrambledPassword = data.ScrambledPassword
	if userModel.ScrambledPassword.IsNull() {
		userModel.ScrambledPassword = types.BoolValue(false)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &userModel)...)
}

func (r *userResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *userResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	obj, err := convertUserSchemaToStruct(data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to parse user, got error: %s", err))
		return
	}

	err = r.client.Auth().UpdateUser(ctx, data.Id.ValueString(), obj)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update user, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *userResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *userResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Auth().DeleteUser(ctx, data.Id.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to delete user, got error: %s", err))
		return
	}
}

func (r *userResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
