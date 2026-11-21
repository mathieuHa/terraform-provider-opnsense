package proxy

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

var _ resource.Resource = &remoteBlacklistResource{}
var _ resource.ResourceWithConfigure = &remoteBlacklistResource{}
var _ resource.ResourceWithImportState = &remoteBlacklistResource{}

func newRemoteBlacklistResource() resource.Resource {
	return &remoteBlacklistResource{}
}

type remoteBlacklistResource struct {
	client opnsense.Client
}

func (r *remoteBlacklistResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_proxy_remote_blacklist"
}

func (r *remoteBlacklistResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = remoteBlacklistResourceSchema()
}

func (r *remoteBlacklistResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *remoteBlacklistResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *remoteBlacklistResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	obj, err := convertRemoteBlacklistSchemaToStruct(data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to parse remote blacklist, got error: %s", err))
		return
	}

	id, err := r.client.Proxy().AddRemoteBlacklist(ctx, obj)
	if err != nil {
		if id != "" {
			data.Id = types.StringValue(id)

			if readStruct, readErr := r.client.Proxy().GetRemoteBlacklist(ctx, id); readErr == nil {
				if readModel, convErr := convertRemoteBlacklistStructToSchema(readStruct); convErr == nil {
					readModel.Id = data.Id
					data = readModel
				}
			}

			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to create remote blacklist, got error: %s", err))
		return
	}

	data.Id = types.StringValue(id)

	readStruct, err := r.client.Proxy().GetRemoteBlacklist(ctx, id)
	if err != nil {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read back remote blacklist, got error: %s", err))
		return
	}
	readModel, err := convertRemoteBlacklistStructToSchema(readStruct)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read back remote blacklist, got error: %s", err))
		return
	}
	readModel.Id = data.Id
	readModel.Password = data.Password
	data = readModel

	tflog.Trace(ctx, "created a resource")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *remoteBlacklistResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *remoteBlacklistResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	obj, err := r.client.Proxy().GetRemoteBlacklist(ctx, data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, fmt.Sprintf("remote blacklist not present in remote, removing from state"))
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read remote blacklist, got error: %s", err))
		return
	}

	remoteBlacklistModel, err := convertRemoteBlacklistStructToSchema(obj)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read remote blacklist, got error: %s", err))
		return
	}

	remoteBlacklistModel.Id = data.Id
	remoteBlacklistModel.Password = data.Password

	resp.Diagnostics.Append(resp.State.Set(ctx, &remoteBlacklistModel)...)
}

func (r *remoteBlacklistResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *remoteBlacklistResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	obj, err := convertRemoteBlacklistSchemaToStruct(data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to parse remote blacklist, got error: %s", err))
		return
	}

	err = r.client.Proxy().UpdateRemoteBlacklist(ctx, data.Id.ValueString(), obj)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update remote blacklist, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *remoteBlacklistResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *remoteBlacklistResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Proxy().DeleteRemoteBlacklist(ctx, data.Id.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to delete remote blacklist, got error: %s", err))
		return
	}
}

func (r *remoteBlacklistResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
