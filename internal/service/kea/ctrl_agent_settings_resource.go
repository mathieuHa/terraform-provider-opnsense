package kea

import (
	"context"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const ctrlAgentSettingsID = "kea_ctrl_agent_settings"

var _ resource.Resource = &ctrlAgentSettingsResource{}
var _ resource.ResourceWithConfigure = &ctrlAgentSettingsResource{}
var _ resource.ResourceWithImportState = &ctrlAgentSettingsResource{}

func newCtrlAgentSettingsResource() resource.Resource {
	return &ctrlAgentSettingsResource{}
}

type ctrlAgentSettingsResource struct {
	client opnsense.Client
}

func (r *ctrlAgentSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kea_ctrl_agent_settings"
}

func (r *ctrlAgentSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Kea control agent settings (**Services ▸ Kea DHCP ▸ Control Agent**). The control agent is the REST API Kea HA peers use to talk to each other. This is a singleton: import it with the ID `kea_ctrl_agent_settings` before managing it. Attributes left unset keep their current value.",
		Attributes:          ctrlAgentSettingsAttributes(),
	}
}

func (r *ctrlAgentSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	apiClient, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = opnsense.NewClient(apiClient)
}

func (r *ctrlAgentSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Cannot Create Singleton Resource",
		"This resource manages the existing Kea control agent settings, which cannot be created or destroyed.\n\n"+
			"Import it first, for example with an import block:\n"+
			"  import {\n    to = opnsense_kea_ctrl_agent_settings.<name>\n    id = \""+ctrlAgentSettingsID+"\"\n  }",
	)
}

func (r *ctrlAgentSettingsResource) read(ctx context.Context) (*ctrlAgentSettingsResourceModel, error) {
	result, err := r.client.Kea().CtrlAgentSettingsGet(ctx)
	if err != nil {
		return nil, err
	}
	model := convertCtrlAgentSettingsStructToSchema(&result.Monad)
	model.Id = types.StringValue(ctrlAgentSettingsID)
	return model, nil
}

func (r *ctrlAgentSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read Kea control agent settings, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *ctrlAgentSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ctrlAgentSettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.Kea().CtrlAgentSettingsSet(ctx, convertCtrlAgentSettingsSchemaToStruct(data))
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update Kea control agent settings, got error: %s", err))
		return
	}
	if result.Result != "saved" {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update Kea control agent settings, result: %s", result.Result))
		return
	}

	if err := r.client.Kea().Client().ReconfigureService(ctx, "/kea/service/reconfigure"); err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to apply Kea control agent settings, got error: %s", err))
		return
	}

	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read Kea control agent settings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "updated Kea control agent settings")

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *ctrlAgentSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning(
		"Singleton Resource Removed From State Only",
		"opnsense_kea_ctrl_agent_settings was removed from the state. The Kea control agent settings on the firewall were not changed.",
	)
}

func (r *ctrlAgentSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != ctrlAgentSettingsID {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("This is a singleton resource and must be imported with the ID %q, got %q.", ctrlAgentSettingsID, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
