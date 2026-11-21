package kea

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

const ddnsSettingsID = "kea_ddns_settings"

var _ resource.Resource = &ddnsSettingsResource{}
var _ resource.ResourceWithConfigure = &ddnsSettingsResource{}
var _ resource.ResourceWithImportState = &ddnsSettingsResource{}

func newDdnsSettingsResource() resource.Resource {
	return &ddnsSettingsResource{}
}

type ddnsSettingsResource struct {
	client opnsense.Client
}

func (r *ddnsSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kea_ddns_settings"
}

func (r *ddnsSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Kea DHCP-DDNS settings (**Services ▸ Kea DHCP ▸ DDNS**). Zones, DNS servers and TSIG keys are set per subnet on `opnsense_kea_dhcpv4_subnet`. This is a singleton: import it with the ID `kea_ddns_settings` before managing it. Attributes left unset keep their current value.",
		Attributes:          ddnsSettingsAttributes(),
	}
}

func (r *ddnsSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ddnsSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Cannot Create Singleton Resource",
		"This resource manages the existing Kea DDNS settings, which cannot be created or destroyed.\n\n"+
			"Import it first, for example with an import block:\n"+
			"  import {\n    to = opnsense_kea_ddns_settings.<name>\n    id = \""+ddnsSettingsID+"\"\n  }",
	)
}

func (r *ddnsSettingsResource) read(ctx context.Context) (*ddnsSettingsResourceModel, error) {
	result, err := r.client.Kea().DdnsSettingsGet(ctx)
	if err != nil {
		return nil, err
	}
	model := convertDdnsSettingsStructToSchema(&result.Monad)
	model.Id = types.StringValue(ddnsSettingsID)
	return model, nil
}

func (r *ddnsSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read Kea DDNS settings, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *ddnsSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ddnsSettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.Kea().DdnsSettingsSet(ctx, convertDdnsSettingsSchemaToStruct(data))
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update Kea DDNS settings, got error: %s", err))
		return
	}
	if result.Result != "saved" {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update Kea DDNS settings, result: %s", result.Result))
		return
	}

	if err := r.client.Kea().Client().ReconfigureService(ctx, "/kea/service/reconfigure"); err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to apply Kea DDNS settings, got error: %s", err))
		return
	}

	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read Kea DDNS settings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "updated Kea DDNS settings")

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *ddnsSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning(
		"Singleton Resource Removed From State Only",
		"opnsense_kea_ddns_settings was removed from the state. The Kea DDNS settings on the firewall were not changed.",
	)
}

func (r *ddnsSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != ddnsSettingsID {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("This is a singleton resource and must be imported with the ID %q, got %q.", ddnsSettingsID, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
