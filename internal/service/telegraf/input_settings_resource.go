package telegraf

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

const inputSettingsID = "telegraf_input"

var _ resource.Resource = &inputSettingsResource{}
var _ resource.ResourceWithConfigure = &inputSettingsResource{}
var _ resource.ResourceWithImportState = &inputSettingsResource{}

func newInputSettingsResource() resource.Resource {
	return &inputSettingsResource{}
}

type inputSettingsResource struct {
	client opnsense.Client
}

func (r *inputSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_telegraf_input"
}

func (r *inputSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Telegraf input plugins (**Services ▸ Telegraf ▸ Input**, plugin `os-telegraf`): which metrics the agent collects. This is a singleton: import it with the ID `telegraf_input` before managing it. Attributes left unset keep their current value.",
		Attributes:          inputSettingsAttributes(),
	}
}

func (r *inputSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *inputSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Cannot Create Singleton Resource",
		"This resource manages the existing Telegraf inputs, which cannot be created or destroyed.\n\n"+
			"Import it first, for example with an import block:\n"+
			"  import {\n    to = opnsense_telegraf_input.<name>\n    id = \""+inputSettingsID+"\"\n  }",
	)
}

func (r *inputSettingsResource) read(ctx context.Context) (*inputSettingsResourceModel, error) {
	result, err := r.client.Telegraf().InputSettingsGet(ctx)
	if err != nil {
		return nil, err
	}
	model := convertInputSettingsStructToSchema(&result.Monad)
	model.Id = types.StringValue(inputSettingsID)
	return model, nil
}

func (r *inputSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read Telegraf inputs, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *inputSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *inputSettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.Telegraf().InputSettingsSet(ctx, convertInputSettingsSchemaToStruct(data))
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update Telegraf inputs, got error: %s", err))
		return
	}
	if result.Result != "saved" {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update Telegraf inputs, result: %s", result.Result))
		return
	}

	if err := r.client.Telegraf().Client().ReconfigureService(ctx, "/telegraf/service/reconfigure"); err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to apply Telegraf inputs, got error: %s", err))
		return
	}

	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read Telegraf inputs, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "updated Telegraf inputs")

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *inputSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning(
		"Singleton Resource Removed From State Only",
		"opnsense_telegraf_input was removed from the state. The Telegraf inputs on the firewall were not changed.",
	)
}

func (r *inputSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != inputSettingsID {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("This is a singleton resource and must be imported with the ID %q, got %q.", inputSettingsID, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
