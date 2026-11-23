package bind

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

const dnsblSettingsID = "bind_dnsbl_settings"

var _ resource.Resource = &dnsblSettingsResource{}
var _ resource.ResourceWithConfigure = &dnsblSettingsResource{}
var _ resource.ResourceWithImportState = &dnsblSettingsResource{}

func newDnsblSettingsResource() resource.Resource {
	return &dnsblSettingsResource{}
}

type dnsblSettingsResource struct {
	client opnsense.Client
}

func (r *dnsblSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bind_dnsbl_settings"
}

func (r *dnsblSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "BIND DNS blocklists and safe-search enforcement (**Services ▸ BIND ▸ Configuration ▸ DNSBL**, plugin `os-bind`). This is a singleton: import it with the ID `bind_dnsbl_settings` before managing it. Attributes left unset keep their current value.",
		Attributes:          dnsblSettingsAttributes(),
	}
}

func (r *dnsblSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *dnsblSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Cannot Create Singleton Resource",
		"This resource manages the existing BIND DNSBL settings, which cannot be created or destroyed.\n\n"+
			"Import it first, for example with an import block:\n"+
			"  import {\n    to = opnsense_bind_dnsbl_settings.<name>\n    id = \""+dnsblSettingsID+"\"\n  }",
	)
}

func (r *dnsblSettingsResource) read(ctx context.Context) (*dnsblSettingsResourceModel, error) {
	result, err := r.client.Bind().DnsblSettingsGet(ctx)
	if err != nil {
		return nil, err
	}
	model := convertDnsblSettingsStructToSchema(&result.Monad)
	model.Id = types.StringValue(dnsblSettingsID)
	return model, nil
}

func (r *dnsblSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read BIND DNSBL settings, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *dnsblSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *dnsblSettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.Bind().DnsblSettingsSet(ctx, convertDnsblSettingsSchemaToStruct(data))
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update BIND DNSBL settings, got error: %s", err))
		return
	}
	if result.Result != "saved" {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update BIND DNSBL settings, result: %s", result.Result))
		return
	}

	if err := r.client.Bind().Client().ReconfigureService(ctx, "/bind/service/reconfigure"); err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to apply BIND DNSBL settings, got error: %s", err))
		return
	}

	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read BIND DNSBL settings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "updated BIND DNSBL settings")

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *dnsblSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning(
		"Singleton Resource Removed From State Only",
		"opnsense_bind_dnsbl_settings was removed from the state. The BIND DNSBL settings on the firewall were not changed.",
	)
}

func (r *dnsblSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != dnsblSettingsID {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("This is a singleton resource and must be imported with the ID %q, got %q.", dnsblSettingsID, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
