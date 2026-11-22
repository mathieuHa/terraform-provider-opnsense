package firewall

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

const aliasSettingsID = "firewall_alias_settings"

var _ resource.Resource = &aliasSettingsResource{}
var _ resource.ResourceWithConfigure = &aliasSettingsResource{}
var _ resource.ResourceWithImportState = &aliasSettingsResource{}

func newAliasSettingsResource() resource.Resource {
	return &aliasSettingsResource{}
}

type aliasSettingsResource struct {
	client opnsense.Client
}

func (r *aliasSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall_alias_settings"
}

func (r *aliasSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Firewall alias settings (**Firewall ▸ Aliases ▸ GeoIP settings**): the source of the GeoIP database used by `geoip` aliases. This is a singleton: import it with the ID `firewall_alias_settings` before managing it. Attributes left unset keep their current value.",
		Attributes:          aliasSettingsAttributes(),
	}
}

func (r *aliasSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *aliasSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Cannot Create Singleton Resource",
		"This resource manages the existing alias settings, which cannot be created or destroyed.\n\n"+
			"Import it first, for example with an import block:\n"+
			"  import {\n    to = opnsense_firewall_alias_settings.<name>\n    id = \""+aliasSettingsID+"\"\n  }",
	)
}

func (r *aliasSettingsResource) read(ctx context.Context) (*aliasSettingsResourceModel, error) {
	result, err := r.client.Firewall().AliasConfigGet(ctx)
	if err != nil {
		return nil, err
	}
	model := convertAliasSettingsStructToSchema(&result.Monad)
	model.Id = types.StringValue(aliasSettingsID)
	return model, nil
}

func (r *aliasSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read alias settings, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *aliasSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *aliasSettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.Firewall().AliasConfigSet(ctx, convertAliasSettingsSchemaToStruct(data))
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update alias settings, got error: %s", err))
		return
	}
	if result.Result != "saved" {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update alias settings, result: %s", result.Result))
		return
	}

	if err := r.client.Firewall().Client().ReconfigureService(ctx, "/firewall/alias/reconfigure"); err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to apply alias settings, got error: %s", err))
		return
	}

	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read alias settings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "updated alias settings")

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *aliasSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning(
		"Singleton Resource Removed From State Only",
		"opnsense_firewall_alias_settings was removed from the state. The alias settings on the firewall were not changed.",
	)
}

func (r *aliasSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != aliasSettingsID {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("This is a singleton resource and must be imported with the ID %q, got %q.", aliasSettingsID, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
