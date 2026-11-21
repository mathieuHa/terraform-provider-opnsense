package hasync

import (
	"context"
	"fmt"
	"time"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &syncResource{}
var _ resource.ResourceWithConfigure = &syncResource{}

type syncResourceModel struct {
	Triggers types.Map    `tfsdk:"triggers"`
	SyncedAt types.String `tfsdk:"synced_at"`
	Id       types.String `tfsdk:"id"`
}

func newSyncResource() resource.Resource {
	return &syncResource{}
}

type syncResource struct {
	client opnsense.Client
}

func (r *syncResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hasync_sync"
}

func (r *syncResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Pushes the configuration to the backup firewall and reconfigures its services, like **System ▸ High Availability ▸ Status ▸ Synchronize and reconfigure all**. The push happens when the resource is created and whenever `triggers` change; destroying the resource does nothing on the firewalls.",
		Attributes: map[string]schema.Attribute{
			"triggers": schema.MapAttribute{
				MarkdownDescription: "Arbitrary values; any change pushes the configuration again, e.g. the IDs of the rules managed on the master.",
				Optional:            true,
				ElementType:         types.StringType,
				PlanModifiers:       []planmodifier.Map{mapplanmodifier.RequiresReplace()},
			},
			"synced_at": schema.StringAttribute{
				MarkdownDescription: "Time of the last push (RFC 3339).",
				Computed:            true,
			},
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *syncResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *syncResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data syncResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, err := r.client.Hasync().SettingsSyncAll(ctx); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to synchronise the backup firewall, got error: %s", err))
		return
	}

	now := time.Now().UTC().Format(time.RFC3339)
	data.SyncedAt = types.StringValue(now)
	data.Id = types.StringValue(now)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *syncResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *syncResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data syncResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *syncResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
