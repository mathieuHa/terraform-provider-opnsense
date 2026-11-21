package captiveportal

import (
	"context"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/captiveportal"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &templateResource{}
var _ resource.ResourceWithConfigure = &templateResource{}
var _ resource.ResourceWithImportState = &templateResource{}

type templateResourceModel struct {
	Name    types.String `tfsdk:"name"`
	Content types.String `tfsdk:"content"`
	FileId  types.String `tfsdk:"file_id"`
	Id      types.String `tfsdk:"id"`
}

func newTemplateResource() resource.Resource {
	return &templateResource{}
}

type templateResource struct {
	client opnsense.Client
}

func (r *templateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_captiveportal_template"
}

func (r *templateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Custom captive portal page (**Services ▸ Captive Portal ▸ Templates**), used by `opnsense_captiveportal_zone.template`. OPNsense keeps only the files that differ from the default template, so `content` is not read back: changes made outside Terraform are not detected.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Template name, unique on the firewall.",
				Required:            true,
			},
			"content": schema.StringAttribute{
				MarkdownDescription: "Base64-encoded ZIP of the template files (at least `index.html`), e.g. `filebase64(\"portal.zip\")`. Start from the default template downloaded from the GUI.",
				Required:            true,
			},
			"file_id": schema.StringAttribute{
				MarkdownDescription: "Internal file ID of the template.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "UUID of the template.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *templateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *templateResource) find(ctx context.Context, match func(captiveportal.TemplateRow) bool) (*captiveportal.TemplateRow, error) {
	res, err := r.client.CaptivePortal().TemplateSearch(ctx)
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

func (r *templateResource) save(ctx context.Context, data *templateResourceModel, uuid string) error {
	var res *captiveportal.TemplateSaveResult
	var err error
	if uuid == "" {
		res, err = r.client.CaptivePortal().TemplateCreate(ctx, data.Name.ValueString(), data.Content.ValueString())
	} else {
		res, err = r.client.CaptivePortal().TemplateUpdate(ctx, uuid, data.Name.ValueString(), data.Content.ValueString())
	}
	if err != nil {
		return err
	}
	if res.Error != "" {
		return fmt.Errorf("%s", res.Error)
	}
	return nil
}

func (r *templateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data templateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	name := data.Name.ValueString()
	existing, err := r.find(ctx, func(t captiveportal.TemplateRow) bool { return t.Name == name })
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to list captive portal templates, got error: %s", err))
		return
	}
	if existing != nil {
		resp.Diagnostics.AddError("Template Already Exists",
			fmt.Sprintf("A captive portal template named %q already exists (%s); import it instead.", name, existing.UUID))
		return
	}
	if err := r.save(ctx, &data, ""); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create captive portal template, got error: %s", err))
		return
	}
	row, err := r.find(ctx, func(t captiveportal.TemplateRow) bool { return t.Name == name })
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read captive portal template %q after creating it, got error: %s", name, err))
		return
	}
	if row == nil {
		resp.Diagnostics.AddError("Template Not Saved",
			fmt.Sprintf("OPNsense did not save the captive portal template %q. The content must be a base64-encoded ZIP of the template files.", name))
		return
	}
	data.Id = types.StringValue(row.UUID)
	data.FileId = types.StringValue(row.FileId)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *templateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data templateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	id := data.Id.ValueString()
	row, err := r.find(ctx, func(t captiveportal.TemplateRow) bool { return t.UUID == id })
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read captive portal template, got error: %s", err))
		return
	}
	if row == nil {
		resp.State.RemoveResource(ctx)
		return
	}
	data.Name = types.StringValue(row.Name)
	data.FileId = types.StringValue(row.FileId)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *templateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state templateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.save(ctx, &data, state.Id.ValueString()); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update captive portal template, got error: %s", err))
		return
	}
	data.Id = state.Id
	data.FileId = state.FileId
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *templateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data templateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.client.CaptivePortal().TemplateDelete(ctx, data.Id.ValueString()); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete captive portal template, got error: %s", err))
	}
}

func (r *templateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
