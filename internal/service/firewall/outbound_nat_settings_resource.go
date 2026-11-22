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

const outboundNatSettingsID = "firewall_outbound_nat_settings"

var _ resource.Resource = &outboundNatSettingsResource{}
var _ resource.ResourceWithConfigure = &outboundNatSettingsResource{}
var _ resource.ResourceWithImportState = &outboundNatSettingsResource{}

func newOutboundNatSettingsResource() resource.Resource {
	return &outboundNatSettingsResource{}
}

type outboundNatSettingsResource struct {
	client opnsense.Client
}

func (r *outboundNatSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall_outbound_nat_settings"
}

func (r *outboundNatSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Outbound NAT mode (**Firewall ▸ NAT ▸ Outbound**). Manual rules are `opnsense_firewall_nat` resources. This is a singleton: import it with the ID `firewall_outbound_nat_settings` before managing it. Attributes left unset keep their current value.",
		Attributes:          outboundNatSettingsAttributes(),
	}
}

func (r *outboundNatSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *outboundNatSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Cannot Create Singleton Resource",
		"This resource manages the existing outbound NAT settings, which cannot be created or destroyed.\n\n"+
			"Import it first, for example with an import block:\n"+
			"  import {\n    to = opnsense_firewall_outbound_nat_settings.<name>\n    id = \""+outboundNatSettingsID+"\"\n  }",
	)
}

func (r *outboundNatSettingsResource) read(ctx context.Context) (*outboundNatSettingsResourceModel, error) {
	result, err := r.client.Firewall().OutboundNatConfigGet(ctx)
	if err != nil {
		return nil, err
	}
	model := convertOutboundNatSettingsStructToSchema(&result.Monad)
	model.Id = types.StringValue(outboundNatSettingsID)
	return model, nil
}

func (r *outboundNatSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read outbound NAT settings, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *outboundNatSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *outboundNatSettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.Firewall().OutboundNatConfigSet(ctx, convertOutboundNatSettingsSchemaToStruct(data))
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update outbound NAT settings, got error: %s", err))
		return
	}
	if result.Result != "saved" {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update outbound NAT settings, result: %s", result.Result))
		return
	}

	if err := r.client.Firewall().Client().ReconfigureService(ctx, "/firewall/source_nat/apply"); err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to apply outbound NAT settings, got error: %s", err))
		return
	}

	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read outbound NAT settings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "updated outbound NAT settings")

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *outboundNatSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning(
		"Singleton Resource Removed From State Only",
		"opnsense_firewall_outbound_nat_settings was removed from the state. The outbound NAT settings on the firewall were not changed.",
	)
}

func (r *outboundNatSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != outboundNatSettingsID {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("This is a singleton resource and must be imported with the ID %q, got %q.", outboundNatSettingsID, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
