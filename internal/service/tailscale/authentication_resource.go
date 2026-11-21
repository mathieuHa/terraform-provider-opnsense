package tailscale

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

const authenticationID = "tailscale_authentication"

var _ resource.Resource = &authenticationResource{}
var _ resource.ResourceWithConfigure = &authenticationResource{}
var _ resource.ResourceWithImportState = &authenticationResource{}

func newAuthenticationResource() resource.Resource {
	return &authenticationResource{}
}

type authenticationResource struct {
	client opnsense.Client
}

func (r *authenticationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tailscale_authentication"
}

func (r *authenticationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Tailscale login settings (**VPN ▸ Tailscale ▸ Authentication**, plugin `os-tailscale`): coordination server and pre-authentication key. This is a singleton: import it with the ID `tailscale_authentication` before managing it. Attributes left unset keep their current value.",
		Attributes:          authenticationAttributes(),
	}
}

func (r *authenticationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *authenticationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Cannot Create Singleton Resource",
		"This resource manages the existing Tailscale authentication, which cannot be created or destroyed.\n\n"+
			"Import it first, for example with an import block:\n"+
			"  import {\n    to = opnsense_tailscale_authentication.<name>\n    id = \""+authenticationID+"\"\n  }",
	)
}

func (r *authenticationResource) read(ctx context.Context) (*authenticationResourceModel, error) {
	result, err := r.client.Tailscale().AuthenticationGet(ctx)
	if err != nil {
		return nil, err
	}
	model := convertAuthenticationStructToSchema(&result.Monad)
	model.Id = types.StringValue(authenticationID)
	return model, nil
}

func (r *authenticationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read Tailscale authentication, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *authenticationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *authenticationResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.Tailscale().AuthenticationSet(ctx, convertAuthenticationSchemaToStruct(data))
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update Tailscale authentication, got error: %s", err))
		return
	}
	if result.Result != "saved" {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update Tailscale authentication, result: %s", result.Result))
		return
	}

	if err := r.client.Tailscale().Client().ReconfigureService(ctx, "/tailscale/service/reconfigure"); err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to apply Tailscale authentication, got error: %s", err))
		return
	}

	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read Tailscale authentication, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "updated Tailscale authentication")

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *authenticationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning(
		"Singleton Resource Removed From State Only",
		"opnsense_tailscale_authentication was removed from the state. The Tailscale authentication on the firewall were not changed.",
	)
}

func (r *authenticationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != authenticationID {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("This is a singleton resource and must be imported with the ID %q, got %q.", authenticationID, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
