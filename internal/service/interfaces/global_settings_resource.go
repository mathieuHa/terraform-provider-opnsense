package interfaces

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

const globalSettingsID = "interfaces_settings"

var _ resource.Resource = &globalSettingsResource{}
var _ resource.ResourceWithConfigure = &globalSettingsResource{}
var _ resource.ResourceWithImportState = &globalSettingsResource{}

func newGlobalSettingsResource() resource.Resource {
	return &globalSettingsResource{}
}

type globalSettingsResource struct {
	client opnsense.Client
}

func (r *globalSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interfaces_settings"
}

func (r *globalSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Global interface settings (**Interfaces ▸ Settings**): hardware offloading, VLAN hardware filtering, IPv6 and the DHCPv6 client. This is a singleton: import it with the ID `interfaces_settings` before managing it. Attributes left unset keep their current value.",
		Attributes:          globalSettingsAttributes(),
	}
}

func (r *globalSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *globalSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Cannot Create Singleton Resource",
		"This resource manages the existing global interface settings, which cannot be created or destroyed.\n\n"+
			"Import it first, for example with an import block:\n"+
			"  import {\n    to = opnsense_interfaces_settings.<name>\n    id = \""+globalSettingsID+"\"\n  }",
	)
}

func (r *globalSettingsResource) read(ctx context.Context) (*globalSettingsResourceModel, error) {
	result, err := r.client.Interfaces().GlobalSettingsGet(ctx)
	if err != nil {
		return nil, err
	}
	model := convertGlobalSettingsStructToSchema(&result.Monad)
	model.Id = types.StringValue(globalSettingsID)
	return model, nil
}

func (r *globalSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read global interface settings, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *globalSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *globalSettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.Interfaces().GlobalSettingsSet(ctx, convertGlobalSettingsSchemaToStruct(data))
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update global interface settings, got error: %s", err))
		return
	}
	if result.Result != "saved" {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update global interface settings, result: %s", result.Result))
		return
	}

	if err := r.client.Interfaces().Client().ReconfigureService(ctx, "/interfaces/settings/reconfigure"); err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to apply global interface settings, got error: %s", err))
		return
	}

	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read global interface settings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "updated global interface settings")

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *globalSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning(
		"Singleton Resource Removed From State Only",
		"opnsense_interfaces_settings was removed from the state. The global interface settings on the firewall were not changed.",
	)
}

func (r *globalSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != globalSettingsID {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("This is a singleton resource and must be imported with the ID %q, got %q.", globalSettingsID, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
