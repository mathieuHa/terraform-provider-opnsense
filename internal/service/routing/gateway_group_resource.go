package routing

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

var _ resource.Resource = &gatewayGroupResource{}
var _ resource.ResourceWithConfigure = &gatewayGroupResource{}
var _ resource.ResourceWithImportState = &gatewayGroupResource{}

func newGatewayGroupResource() resource.Resource {
	return &gatewayGroupResource{}
}

type gatewayGroupResource struct {
	client opnsense.Client
}

func (r *gatewayGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_gateway_group"
}

func (r *gatewayGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = gatewayGroupResourceSchema()
}

func (r *gatewayGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *gatewayGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *gatewayGroupResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	group, err := convertGatewayGroupSchemaToStruct(data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to parse gateway group, got error: %s", err))
		return
	}

	var id string
	err = retryWhileGatewayListStale(ctx, func() error {
		var addErr error
		id, addErr = r.client.Routing().AddGatewayGroup(ctx, group)
		return addErr
	})
	if err != nil {
		if id != "" {
			data.Id = types.StringValue(id)

			if readStruct, readErr := r.client.Routing().GetGatewayGroup(ctx, id); readErr == nil {
				if readModel, convErr := convertGatewayGroupStructToSchema(readStruct); convErr == nil {
					readModel.Id = data.Id
					data = readModel
				}
			}

			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to create gateway group, got error: %s", err))
		return
	}

	data.Id = types.StringValue(id)

	tflog.Trace(ctx, "created a resource")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *gatewayGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *gatewayGroupResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	group, err := r.client.Routing().GetGatewayGroup(ctx, data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, fmt.Sprintf("gateway group not present in remote, removing from state"))
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read gateway group, got error: %s", err))
		return
	}

	gatewayGroupModel, err := convertGatewayGroupStructToSchema(group)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read gateway group, got error: %s", err))
		return
	}

	gatewayGroupModel.Id = data.Id

	resp.Diagnostics.Append(resp.State.Set(ctx, &gatewayGroupModel)...)
}

func (r *gatewayGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *gatewayGroupResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	group, err := convertGatewayGroupSchemaToStruct(data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to parse gateway group, got error: %s", err))
		return
	}

	err = retryWhileGatewayListStale(ctx, func() error {
		return r.client.Routing().UpdateGatewayGroup(ctx, data.Id.ValueString(), group)
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update gateway group, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *gatewayGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *gatewayGroupResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Routing().DeleteGatewayGroup(ctx, data.Id.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to delete gateway group, got error: %s", err))
		return
	}
}

func (r *gatewayGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
