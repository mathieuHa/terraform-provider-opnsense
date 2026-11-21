package zerotier

import (
	"context"
	"fmt"
	"regexp"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/browningluke/opnsense-go/pkg/zerotier"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var networkIDRegexp = regexp.MustCompile(`^[0-9a-fA-F]{16}$`)

var _ resource.Resource = &networkResource{}
var _ resource.ResourceWithConfigure = &networkResource{}
var _ resource.ResourceWithImportState = &networkResource{}

type networkResourceModel struct {
	NetworkId   types.String `tfsdk:"network_id"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	Description types.String `tfsdk:"description"`
	Id          types.String `tfsdk:"id"`
}

func newNetworkResource() resource.Resource {
	return &networkResource{}
}

type networkResource struct {
	client opnsense.Client
}

func (r *networkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zerotier_network"
}

func (r *networkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ZeroTier networks this firewall joins (**VPN ▸ ZeroTier ▸ Networks**, plugin `os-zerotier`). The plugin only joins, leaves and deletes networks while the ZeroTier service is enabled (`opnsense_zerotier_settings.enabled`), so enable it before creating enabled networks or destroying any network.",
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: "16-character ZeroTier network ID.",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Validators:          []validator.String{stringvalidator.RegexMatches(networkIDRegexp, "must be 16 hexadecimal characters")},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Join the network. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "UUID of the network entry.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *networkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	apiClient, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = opnsense.NewClient(apiClient)
}

func (r *networkResource) find(ctx context.Context, match func(zerotier.NetworkRow) bool) (*zerotier.NetworkRow, error) {
	res, err := r.client.Zerotier().NetworkSearch(ctx)
	if err != nil {
		return nil, err
	}
	for _, row := range res.Rows {
		if match(row) {
			return &row, nil
		}
	}
	return nil, nil
}

func (r *networkResource) toggle(ctx context.Context, uuid string) error {
	res, err := r.client.Zerotier().NetworkToggle(ctx, uuid)
	if err != nil {
		return err
	}
	if res.Result == "service_not_enabled" {
		return fmt.Errorf("the ZeroTier service is disabled; enable it with opnsense_zerotier_settings first")
	}
	return nil
}

func (r *networkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data networkResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	nid := data.NetworkId.ValueString()
	if existing, err := r.find(ctx, func(n zerotier.NetworkRow) bool { return n.NetworkId == nid }); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to list ZeroTier networks, got error: %s", err))
		return
	} else if existing != nil {
		resp.Diagnostics.AddError("Network Already Exists", fmt.Sprintf("ZeroTier network %s is already configured (%s); import it instead.", nid, existing.UUID))
		return
	}
	if _, err := r.client.Zerotier().NetworkAdd(ctx, &zerotier.Network{Enabled: "0", NetworkId: nid, Description: data.Description.ValueString()}); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ZeroTier network, got error: %s", err))
		return
	}
	row, err := r.find(ctx, func(n zerotier.NetworkRow) bool { return n.NetworkId == nid })
	if err != nil || row == nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to find ZeroTier network %s after creating it: %v", nid, err))
		return
	}
	data.Id = types.StringValue(row.UUID)
	if data.Enabled.ValueBool() {
		if err := r.toggle(ctx, row.UUID); err != nil {
			resp.Diagnostics.Append(resp.State.Set(ctx, &networkResourceModel{NetworkId: data.NetworkId, Enabled: types.BoolValue(false), Description: data.Description, Id: data.Id})...)
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to join ZeroTier network %s: %s", nid, err))
			return
		}
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *networkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data networkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	id := data.Id.ValueString()
	row, err := r.find(ctx, func(n zerotier.NetworkRow) bool { return n.UUID == id })
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ZeroTier network, got error: %s", err))
		return
	}
	if row == nil {
		resp.State.RemoveResource(ctx)
		return
	}
	data.NetworkId = types.StringValue(row.NetworkId)
	data.Enabled = types.BoolValue(tools.StringToBool(row.Enabled))
	data.Description = types.StringValue(row.Description)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *networkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state networkResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	uuid := state.Id.ValueString()
	if _, err := r.client.Zerotier().NetworkSet(ctx, uuid, &zerotier.Network{
		Enabled:     tools.BoolToString(state.Enabled.ValueBool()),
		NetworkId:   data.NetworkId.ValueString(),
		Description: data.Description.ValueString(),
	}); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ZeroTier network, got error: %s", err))
		return
	}
	if data.Enabled.ValueBool() != state.Enabled.ValueBool() {
		if err := r.toggle(ctx, uuid); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to join or leave ZeroTier network %s: %s", data.NetworkId.ValueString(), err))
			return
		}
	}
	data.Id = state.Id
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *networkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data networkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Zerotier().NetworkDelete(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ZeroTier network, got error: %s", err))
		return
	}
	switch res.Result {
	case "deleted", "not_found":
	case "service_not_enabled":
		resp.Diagnostics.AddError("ZeroTier Service Disabled",
			"The os-zerotier plugin only deletes networks while the ZeroTier service is enabled. Enable it with opnsense_zerotier_settings, then destroy again.")
	default:
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ZeroTier network, result: %s", res.Result))
	}
}

func (r *networkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
