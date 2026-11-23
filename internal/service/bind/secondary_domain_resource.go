package bind

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

var _ resource.Resource = &secondaryDomainResource{}
var _ resource.ResourceWithConfigure = &secondaryDomainResource{}
var _ resource.ResourceWithImportState = &secondaryDomainResource{}

func newSecondaryDomainResource() resource.Resource {
	return &secondaryDomainResource{}
}

type secondaryDomainResource struct {
	client opnsense.Client
}

func (r *secondaryDomainResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bind_secondary_domain"
}

func (r *secondaryDomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = secondaryDomainResourceSchema()
}

func (r *secondaryDomainResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *secondaryDomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *secondaryDomainResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	obj, err := convertSecondaryDomainSchemaToStruct(data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to parse BIND secondary zone, got error: %s", err))
		return
	}

	id, err := r.client.Bind().AddSecondaryDomain(ctx, obj)
	if err != nil {
		if id != "" {
			data.Id = types.StringValue(id)

			if readStruct, readErr := r.client.Bind().GetSecondaryDomain(ctx, id); readErr == nil {
				if readModel, convErr := convertSecondaryDomainStructToSchema(readStruct); convErr == nil {
					readModel.Id = data.Id
					data = readModel
				}
			}

			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to create BIND secondary zone, got error: %s", err))
		return
	}

	data.Id = types.StringValue(id)

	readStruct, err := r.client.Bind().GetSecondaryDomain(ctx, id)
	if err != nil {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read back BIND secondary zone, got error: %s", err))
		return
	}
	readModel, err := convertSecondaryDomainStructToSchema(readStruct)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read back BIND secondary zone, got error: %s", err))
		return
	}
	readModel.Id = data.Id
	data = readModel

	tflog.Trace(ctx, "created a resource")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *secondaryDomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *secondaryDomainResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	obj, err := r.client.Bind().GetSecondaryDomain(ctx, data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, fmt.Sprintf("BIND secondary zone not present in remote, removing from state"))
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read BIND secondary zone, got error: %s", err))
		return
	}

	secondaryDomainModel, err := convertSecondaryDomainStructToSchema(obj)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read BIND secondary zone, got error: %s", err))
		return
	}

	secondaryDomainModel.Id = data.Id

	resp.Diagnostics.Append(resp.State.Set(ctx, &secondaryDomainModel)...)
}

func (r *secondaryDomainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *secondaryDomainResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	obj, err := convertSecondaryDomainSchemaToStruct(data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to parse BIND secondary zone, got error: %s", err))
		return
	}

	err = r.client.Bind().UpdateSecondaryDomain(ctx, data.Id.ValueString(), obj)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update BIND secondary zone, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *secondaryDomainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *secondaryDomainResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Bind().DeleteSecondaryDomain(ctx, data.Id.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to delete BIND secondary zone, got error: %s", err))
		return
	}
}

func (r *secondaryDomainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
