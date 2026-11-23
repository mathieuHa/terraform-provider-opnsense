package dyndns

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

var _ resource.Resource = &accountResource{}
var _ resource.ResourceWithConfigure = &accountResource{}
var _ resource.ResourceWithImportState = &accountResource{}

func newAccountResource() resource.Resource {
	return &accountResource{}
}

type accountResource struct {
	client opnsense.Client
}

func (r *accountResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dyndns_account"
}

func (r *accountResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = accountResourceSchema()
}

func (r *accountResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *accountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *accountResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	obj, err := convertAccountSchemaToStruct(data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to parse dynamic DNS account, got error: %s", err))
		return
	}

	id, err := r.client.Dyndns().AddAccount(ctx, obj)
	if err != nil {
		if id != "" {
			data.Id = types.StringValue(id)

			if readStruct, readErr := r.client.Dyndns().GetAccount(ctx, id); readErr == nil {
				if readModel, convErr := convertAccountStructToSchema(readStruct); convErr == nil {
					readModel.Id = data.Id
					data = readModel
				}
			}

			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to create dynamic DNS account, got error: %s", err))
		return
	}

	data.Id = types.StringValue(id)

	readStruct, err := r.client.Dyndns().GetAccount(ctx, id)
	if err != nil {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read back dynamic DNS account, got error: %s", err))
		return
	}
	readModel, err := convertAccountStructToSchema(readStruct)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read back dynamic DNS account, got error: %s", err))
		return
	}
	readModel.Id = data.Id
	readModel.Password = data.Password
	data = readModel

	tflog.Trace(ctx, "created a resource")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *accountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *accountResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	obj, err := r.client.Dyndns().GetAccount(ctx, data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, fmt.Sprintf("dynamic DNS account not present in remote, removing from state"))
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read dynamic DNS account, got error: %s", err))
		return
	}

	accountModel, err := convertAccountStructToSchema(obj)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read dynamic DNS account, got error: %s", err))
		return
	}

	accountModel.Id = data.Id
	accountModel.Password = data.Password

	resp.Diagnostics.Append(resp.State.Set(ctx, &accountModel)...)
}

func (r *accountResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *accountResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	obj, err := convertAccountSchemaToStruct(data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to parse dynamic DNS account, got error: %s", err))
		return
	}

	err = r.client.Dyndns().UpdateAccount(ctx, data.Id.ValueString(), obj)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update dynamic DNS account, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *accountResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *accountResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Dyndns().DeleteAccount(ctx, data.Id.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to delete dynamic DNS account, got error: %s", err))
		return
	}
}

func (r *accountResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
