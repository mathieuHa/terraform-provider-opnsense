package ids

import (
	"context"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &rulesetResource{}
var _ resource.ResourceWithConfigure = &rulesetResource{}
var _ resource.ResourceWithImportState = &rulesetResource{}

type rulesetResourceModel struct {
	Filename    types.String `tfsdk:"filename"`
	Description types.String `tfsdk:"description"`
	UUID        types.String `tfsdk:"uuid"`
	Id          types.String `tfsdk:"id"`
}

func newRulesetResource() resource.Resource {
	return &rulesetResource{}
}

type rulesetResource struct {
	client opnsense.Client
}

func (r *rulesetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ids_ruleset"
}

func (r *rulesetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Enables one of the installable Suricata rulesets (**Services ▸ Intrusion Detection ▸ Administration ▸ Download**). Destroying the resource disables the ruleset. Rules are downloaded by the daily IDS update job or the **Download & Update Rules** button, then loaded when the IDS is reconfigured.",
		Attributes: map[string]schema.Attribute{
			"filename": schema.StringAttribute{
				MarkdownDescription: "Ruleset file name, e.g. `emerging-scan.rules` or `abuse.ch.urlhaus.rules`.",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Ruleset description.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"uuid": schema.StringAttribute{
				MarkdownDescription: "UUID of the ruleset, for `opnsense_ids_policy.rulesets`.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Same as `filename`.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *rulesetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *rulesetResource) toggle(ctx context.Context, filename, enabled string) error {
	res, err := r.client.Ids().RulesetToggle(ctx, filename, enabled)
	if err != nil {
		return err
	}
	if res.Status != enabled {
		return fmt.Errorf("unknown ruleset %q (status %q)", filename, res.Status)
	}
	return r.client.Ids().Client().ReconfigureService(ctx, "/ids/service/reconfigure")
}

func (r *rulesetResource) read(ctx context.Context, filename string) (*rulesetResourceModel, bool, error) {
	rs, err := r.client.Ids().RulesetGet(ctx, filename)
	if err != nil {
		return nil, false, err
	}
	if rs.Filename != filename || rs.Enabled != "1" {
		return nil, false, nil
	}
	opts, err := r.client.Ids().RulesetOptions(ctx)
	if err != nil {
		return nil, false, err
	}
	uuid, _ := opts.Policy.Rulesets.KeyOf(filename)
	return &rulesetResourceModel{
		Filename:    types.StringValue(filename),
		Description: types.StringValue(rs.Description),
		UUID:        types.StringValue(uuid),
		Id:          types.StringValue(filename),
	}, true, nil
}

func (r *rulesetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data rulesetResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	filename := data.Filename.ValueString()
	if err := r.toggle(ctx, filename, "1"); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to enable IDS ruleset, got error: %s", err))
		return
	}
	model, ok, err := r.read(ctx, filename)
	if err != nil || !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read IDS ruleset %q after enabling it: %v", filename, err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *rulesetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data rulesetResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	model, ok, err := r.read(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read IDS ruleset, got error: %s", err))
		return
	}
	if !ok {
		resp.State.RemoveResource(ctx)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *rulesetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data rulesetResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *rulesetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data rulesetResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.toggle(ctx, data.Id.ValueString(), "0"); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to disable IDS ruleset, got error: %s", err))
	}
}

func (r *rulesetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("filename"), req.ID)...)
}
