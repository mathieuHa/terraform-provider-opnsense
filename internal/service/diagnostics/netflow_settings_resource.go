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

const netflowSettingsID = "netflow_settings"

var _ resource.Resource = &netflowSettingsResource{}
var _ resource.ResourceWithConfigure = &netflowSettingsResource{}
var _ resource.ResourceWithImportState = &netflowSettingsResource{}

func newNetflowSettingsResource() resource.Resource {
	return &netflowSettingsResource{}
}

type netflowSettingsResource struct {
	client opnsense.Client
}

func (r *netflowSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_netflow_settings"
}

func (r *netflowSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "NetFlow export and local collection (**Reporting ▸ NetFlow**). When `collect_enable` is set, OPNsense adds the local collector `127.0.0.1:2056` to `capture_targets`; include it in the configuration to avoid a diff. This is a singleton: import it with the ID `netflow_settings` before managing it. Attributes left unset keep their current value.",
		Attributes:          netflowSettingsAttributes(),
	}
}

func (r *netflowSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *netflowSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Cannot Create Singleton Resource",
		"This resource manages the existing NetFlow settings, which cannot be created or destroyed.\n\n"+
			"Import it first, for example with an import block:\n"+
			"  import {\n    to = opnsense_netflow_settings.<name>\n    id = \""+netflowSettingsID+"\"\n  }",
	)
}

func (r *netflowSettingsResource) read(ctx context.Context) (*netflowSettingsResourceModel, error) {
	result, err := r.client.Diagnostics().NetflowConfigGet(ctx)
	if err != nil {
		return nil, err
	}
	model := convertNetflowSettingsStructToSchema(&result.Monad)
	model.Id = types.StringValue(netflowSettingsID)
	return model, nil
}

func (r *netflowSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read NetFlow settings, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *netflowSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *netflowSettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.Diagnostics().NetflowConfigSet(ctx, convertNetflowSettingsSchemaToStruct(data))
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update NetFlow settings, got error: %s", err))
		return
	}
	if result.Result != "saved" {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update NetFlow settings, result: %s", result.Result))
		return
	}

	if err := r.client.Diagnostics().Client().ReconfigureService(ctx, "/diagnostics/netflow/reconfigure"); err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to apply NetFlow settings, got error: %s", err))
		return
	}

	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read NetFlow settings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "updated NetFlow settings")

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *netflowSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning(
		"Singleton Resource Removed From State Only",
		"opnsense_netflow_settings was removed from the state. The NetFlow settings on the firewall were not changed.",
	)
}

func (r *netflowSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != netflowSettingsID {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("This is a singleton resource and must be imported with the ID %q, got %q.", netflowSettingsID, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
