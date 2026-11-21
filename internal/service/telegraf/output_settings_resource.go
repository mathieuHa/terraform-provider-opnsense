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

const outputSettingsID = "telegraf_output"

var _ resource.Resource = &outputSettingsResource{}
var _ resource.ResourceWithConfigure = &outputSettingsResource{}
var _ resource.ResourceWithImportState = &outputSettingsResource{}

func newOutputSettingsResource() resource.Resource {
	return &outputSettingsResource{}
}

type outputSettingsResource struct {
	client opnsense.Client
}

func (r *outputSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_telegraf_output"
}

func (r *outputSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Telegraf output plugins (**Services ▸ Telegraf ▸ Output**, plugin `os-telegraf`): InfluxDB v1/v2, Elasticsearch, Graylog, Prometheus, Datadog, MQTT, OpenTelemetry and more. This is a singleton: import it with the ID `telegraf_output` before managing it. Attributes left unset keep their current value.",
		Attributes:          outputSettingsAttributes(),
	}
}

func (r *outputSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *outputSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Cannot Create Singleton Resource",
		"This resource manages the existing Telegraf outputs, which cannot be created or destroyed.\n\n"+
			"Import it first, for example with an import block:\n"+
			"  import {\n    to = opnsense_telegraf_output.<name>\n    id = \""+outputSettingsID+"\"\n  }",
	)
}

func (r *outputSettingsResource) read(ctx context.Context) (*outputSettingsResourceModel, error) {
	result, err := r.client.Telegraf().OutputSettingsGet(ctx)
	if err != nil {
		return nil, err
	}
	model := convertOutputSettingsStructToSchema(&result.Monad)
	model.Id = types.StringValue(outputSettingsID)
	return model, nil
}

func (r *outputSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read Telegraf outputs, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *outputSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *outputSettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.Telegraf().OutputSettingsSet(ctx, convertOutputSettingsSchemaToStruct(data))
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update Telegraf outputs, got error: %s", err))
		return
	}
	if result.Result != "saved" {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update Telegraf outputs, result: %s", result.Result))
		return
	}

	if err := r.client.Telegraf().Client().ReconfigureService(ctx, "/telegraf/service/reconfigure"); err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to apply Telegraf outputs, got error: %s", err))
		return
	}

	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read Telegraf outputs, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "updated Telegraf outputs")

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *outputSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning(
		"Singleton Resource Removed From State Only",
		"opnsense_telegraf_output was removed from the state. The Telegraf outputs on the firewall were not changed.",
	)
}

func (r *outputSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != outputSettingsID {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("This is a singleton resource and must be imported with the ID %q, got %q.", outputSettingsID, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
