package core

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

const firmwareSettingsID = "firmware_settings"

var _ resource.Resource = &firmwareSettingsResource{}
var _ resource.ResourceWithConfigure = &firmwareSettingsResource{}
var _ resource.ResourceWithImportState = &firmwareSettingsResource{}

func newFirmwareSettingsResource() resource.Resource {
	return &firmwareSettingsResource{}
}

type firmwareSettingsResource struct {
	client opnsense.Client
}

func (r *firmwareSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firmware_settings"
}

func (r *firmwareSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Firmware repository settings (**System ▸ Firmware ▸ Settings**): mirror, flavour, release type and subscription. Plugins are managed with `opnsense_firmware_plugin`. This is a singleton: import it with the ID `firmware_settings` before managing it. Attributes left unset keep their current value.",
		Attributes:          firmwareSettingsAttributes(),
	}
}

func (r *firmwareSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *firmwareSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Cannot Create Singleton Resource",
		"This resource manages the existing firmware settings, which cannot be created or destroyed.\n\n"+
			"Import it first, for example with an import block:\n"+
			"  import {\n    to = opnsense_firmware_settings.<name>\n    id = \""+firmwareSettingsID+"\"\n  }",
	)
}

func (r *firmwareSettingsResource) read(ctx context.Context) (*firmwareSettingsResourceModel, error) {
	result, err := r.client.Core().FirmwareSettingsGet(ctx)
	if err != nil {
		return nil, err
	}
	model := convertFirmwareSettingsStructToSchema(&result.Monad)
	model.Id = types.StringValue(firmwareSettingsID)
	return model, nil
}

func (r *firmwareSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read firmware settings, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *firmwareSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *firmwareSettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.Core().FirmwareSettingsSet(ctx, convertFirmwareSettingsSchemaToStruct(data))
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update firmware settings, got error: %s", err))
		return
	}
	if result.Status != "ok" {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update firmware settings, status: %s", result.Status))
		return
	}

	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read firmware settings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "updated firmware settings")

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *firmwareSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning(
		"Singleton Resource Removed From State Only",
		"opnsense_firmware_settings was removed from the state. The firmware settings on the firewall were not changed.",
	)
}

func (r *firmwareSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != firmwareSettingsID {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("This is a singleton resource and must be imported with the ID %q, got %q.", firmwareSettingsID, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
