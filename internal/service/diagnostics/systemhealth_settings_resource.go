package diagnostics

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

const systemhealthSettingsID = "systemhealth_settings"

var _ resource.Resource = &systemhealthSettingsResource{}
var _ resource.ResourceWithConfigure = &systemhealthSettingsResource{}
var _ resource.ResourceWithImportState = &systemhealthSettingsResource{}

func newSystemhealthSettingsResource() resource.Resource {
	return &systemhealthSettingsResource{}
}

type systemhealthSettingsResource struct {
	client opnsense.Client
}

func (r *systemhealthSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemhealth_settings"
}

func (r *systemhealthSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Health graph collection (**Reporting ▸ Settings**). This is a singleton: import it with the ID `systemhealth_settings` before managing it. Attributes left unset keep their current value.",
		Attributes:          systemhealthSettingsAttributes(),
	}
}

func (r *systemhealthSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *systemhealthSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Cannot Create Singleton Resource",
		"This resource manages the existing system health settings, which cannot be created or destroyed.\n\n"+
			"Import it first, for example with an import block:\n"+
			"  import {\n    to = opnsense_systemhealth_settings.<name>\n    id = \""+systemhealthSettingsID+"\"\n  }",
	)
}

func (r *systemhealthSettingsResource) read(ctx context.Context) (*systemhealthSettingsResourceModel, error) {
	result, err := r.client.Diagnostics().SystemHealthConfigGet(ctx)
	if err != nil {
		return nil, err
	}
	model := convertSystemhealthSettingsStructToSchema(&result.Monad)
	model.Id = types.StringValue(systemhealthSettingsID)
	return model, nil
}

func (r *systemhealthSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read system health settings, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *systemhealthSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *systemhealthSettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.Diagnostics().SystemHealthConfigSet(ctx, convertSystemhealthSettingsSchemaToStruct(data))
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update system health settings, got error: %s", err))
		return
	}
	if result.Result != "saved" {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update system health settings, result: %s", result.Result))
		return
	}

	if err := r.client.Diagnostics().Client().ReconfigureService(ctx, "/diagnostics/systemhealth/reconfigure"); err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to apply system health settings, got error: %s", err))
		return
	}

	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read system health settings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "updated system health settings")

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *systemhealthSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning(
		"Singleton Resource Removed From State Only",
		"opnsense_systemhealth_settings was removed from the state. The system health settings on the firewall were not changed.",
	)
}

func (r *systemhealthSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != systemhealthSettingsID {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("This is a singleton resource and must be imported with the ID %q, got %q.", systemhealthSettingsID, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
