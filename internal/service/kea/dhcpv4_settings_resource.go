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

const dhcpv4SettingsID = "kea_dhcpv4_settings"

var _ resource.Resource = &dhcpv4SettingsResource{}
var _ resource.ResourceWithConfigure = &dhcpv4SettingsResource{}
var _ resource.ResourceWithImportState = &dhcpv4SettingsResource{}

func newDhcpv4SettingsResource() resource.Resource {
	return &dhcpv4SettingsResource{}
}

type dhcpv4SettingsResource struct {
	client opnsense.Client
}

func (r *dhcpv4SettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kea_dhcpv4_settings"
}

func (r *dhcpv4SettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Kea DHCPv4 server settings (**Services ▸ Kea DHCP ▸ Kea DHCPv4 ▸ Settings**): service, interfaces, lease lifetime, lease expiry processing and HA. This is a singleton: import it with the ID `kea_dhcpv4_settings` before managing it. Attributes left unset keep their current value.",
		Attributes:          dhcpv4SettingsAttributes(),
	}
}

func (r *dhcpv4SettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *dhcpv4SettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Cannot Create Singleton Resource",
		"This resource manages the existing Kea DHCPv4 settings, which cannot be created or destroyed.\n\n"+
			"Import it first, for example with an import block:\n"+
			"  import {\n    to = opnsense_kea_dhcpv4_settings.<name>\n    id = \""+dhcpv4SettingsID+"\"\n  }",
	)
}

func (r *dhcpv4SettingsResource) read(ctx context.Context) (*dhcpv4SettingsResourceModel, error) {
	result, err := r.client.Kea().Dhcpv4ConfigGet(ctx)
	if err != nil {
		return nil, err
	}
	model := convertDhcpv4SettingsStructToSchema(&result.Monad)
	model.Id = types.StringValue(dhcpv4SettingsID)
	return model, nil
}

func (r *dhcpv4SettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read Kea DHCPv4 settings, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *dhcpv4SettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *dhcpv4SettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.Kea().Dhcpv4ConfigSet(ctx, convertDhcpv4SettingsSchemaToStruct(data))
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update Kea DHCPv4 settings, got error: %s", err))
		return
	}
	if result.Result != "saved" {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to update Kea DHCPv4 settings, result: %s", result.Result))
		return
	}

	if err := r.client.Kea().Client().ReconfigureService(ctx, "/kea/service/reconfigure"); err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to apply Kea DHCPv4 settings, got error: %s", err))
		return
	}

	model, err := r.read(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read Kea DHCPv4 settings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "updated Kea DHCPv4 settings")

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *dhcpv4SettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning(
		"Singleton Resource Removed From State Only",
		"opnsense_kea_dhcpv4_settings was removed from the state. The Kea DHCPv4 settings on the firewall were not changed.",
	)
}

func (r *dhcpv4SettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != dhcpv4SettingsID {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("This is a singleton resource and must be imported with the ID %q, got %q.", dhcpv4SettingsID, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
