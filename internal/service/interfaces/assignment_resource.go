package interfaces

import (
	"context"
	"errors"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/errs"
	"github.com/browningluke/opnsense-go/pkg/interfaces"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &assignmentResource{}
var _ resource.ResourceWithConfigure = &assignmentResource{}
var _ resource.ResourceWithImportState = &assignmentResource{}

type assignmentResourceModel struct {
	Device      types.String `tfsdk:"device"`
	Description types.String `tfsdk:"description"`
	Lock        types.Bool   `tfsdk:"lock"`
	Id          types.String `tfsdk:"id"`
}

func newAssignmentResource() resource.Resource {
	return &assignmentResource{}
}

type assignmentResource struct {
	client opnsense.Client
}

func (r *assignmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interfaces_assignment"
}

func (r *assignmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Interface assignments (**Interfaces ▸ Assignments**) turn a device into an interface such as `opt2`, usable in firewall rules and services. OPNsense 26.7 has no API for the IP configuration of an interface, so addresses still have to be set in the GUI. Never reassign the interface you manage OPNsense through without console access.",
		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "Device to assign, e.g. `vtnet2`, `vlan0.10`, `lo1` or `gre0`.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Interface description, shown as its name in the GUI. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"lock": schema.BoolAttribute{
				MarkdownDescription: "Prevent removal of the interface in the GUI. The provider unlocks the interface before deleting it. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Interface identifier assigned by OPNsense, e.g. `opt2`.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *assignmentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func assignmentToStruct(d *assignmentResourceModel) *interfaces.Assignment {
	return &interfaces.Assignment{
		Device:      api.SelectedMap(d.Device.ValueString()),
		Description: d.Description.ValueString(),
		Lock:        tools.BoolToString(d.Lock.ValueBool()),
	}
}

func (r *assignmentResource) identifierFor(ctx context.Context, device string) (string, error) {
	rows, err := r.client.Interfaces().AssignmentsSearch(ctx)
	if err != nil {
		return "", err
	}
	for _, row := range rows.Rows {
		if row.Device == device {
			return row.Identifier, nil
		}
	}
	return "", fmt.Errorf("no interface is assigned to device %s", device)
}

func (r *assignmentResource) read(ctx context.Context, id string) (*assignmentResourceModel, error) {
	a, err := r.client.Interfaces().GetAssignment(ctx, id)
	if err != nil {
		return nil, err
	}
	return &assignmentResourceModel{
		Device:      types.StringValue(a.Device.String()),
		Description: types.StringValue(a.Description),
		Lock:        types.BoolValue(tools.StringToBool(a.Lock)),
		Id:          types.StringValue(id),
	}, nil
}

func (r *assignmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data assignmentResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, addErr := r.client.Interfaces().AddAssignment(ctx, assignmentToStruct(&data))
	id, err := r.identifierFor(ctx, data.Device.ValueString())
	if err != nil {
		if addErr != nil {
			err = addErr
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to assign %s, got error: %s", data.Device.ValueString(), err))
		return
	}

	model, err := r.read(ctx, id)
	if err != nil {
		data.Id = types.StringValue(id)
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read interface %s, got error: %s", id, err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
	if addErr != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Interface %s was assigned, but applying it failed: %s", id, addErr))
	}
}

func (r *assignmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data assignmentResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	model, err := r.read(ctx, data.Id.ValueString())
	if err != nil {
		var notFound *errs.NotFoundError
		if errors.As(err, &notFound) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read interface %s, got error: %s", data.Id.ValueString(), err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *assignmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data assignmentResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.Interfaces().UpdateAssignment(ctx, data.Id.ValueString(), assignmentToStruct(&data)); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update interface %s, got error: %s", data.Id.ValueString(), err))
		return
	}

	model, err := r.read(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read interface %s, got error: %s", data.Id.ValueString(), err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *assignmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data assignmentResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := data.Id.ValueString()
	current, err := r.client.Interfaces().GetAssignment(ctx, id)
	if err != nil {
		var notFound *errs.NotFoundError
		if errors.As(err, &notFound) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read interface %s, got error: %s", id, err))
		return
	}
	if tools.StringToBool(current.Lock) {
		current.Lock = "0"
		if err := r.client.Interfaces().UpdateAssignment(ctx, id, &interfaces.Assignment{Device: current.Device, Description: current.Description, Lock: "0"}); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unlock interface %s, got error: %s", id, err))
			return
		}
	}

	if err := r.client.Interfaces().DeleteAssignment(ctx, id); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete interface %s, got error: %s", id, err))
	}
}

func (r *assignmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
