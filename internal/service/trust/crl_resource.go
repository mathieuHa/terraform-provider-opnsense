package trust

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/browningluke/opnsense-go/pkg/trust"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var crlReasons = []string{"unspecified", "keyCompromise", "cACompromise", "affiliationChanged", "superseded", "cessationOfOperation", "certificateHold"}

var crlRevokedTypes = map[string]attr.Type{
	"certificate": types.StringType,
	"reason":      types.StringType,
}

var _ resource.Resource = &crlResource{}
var _ resource.ResourceWithConfigure = &crlResource{}
var _ resource.ResourceWithImportState = &crlResource{}

type crlResourceModel struct {
	Ca          types.String `tfsdk:"ca"`
	Description types.String `tfsdk:"description"`
	Method      types.String `tfsdk:"method"`
	Lifetime    types.Int64  `tfsdk:"lifetime"`
	ImportPem   types.String `tfsdk:"import_pem"`
	Revoked     types.Set    `tfsdk:"revoked"`
	Serial      types.String `tfsdk:"serial"`
	Pem         types.String `tfsdk:"pem"`
	Id          types.String `tfsdk:"id"`
}

type crlRevokedModel struct {
	Certificate types.String `tfsdk:"certificate"`
	Reason      types.String `tfsdk:"reason"`
}

func newCrlResource() resource.Resource {
	return &crlResource{}
}

type crlResource struct {
	client opnsense.Client
}

func (r *crlResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_trust_crl"
}

func (r *crlResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Certificate revocation list of a certificate authority (**System ▸ Trust ▸ Revocation**). With `method = \"internal\"`, OPNsense signs the list with the CA's private key and lists the `revoked` certificates; with `method = \"existing\"`, the list given in `import_pem` is stored as is. There is one list per CA.",
		Attributes: map[string]schema.Attribute{
			"ca": schema.StringAttribute{
				MarkdownDescription: "Reference ID of the CA, e.g. `opnsense_trust_ca.internal.ref_id`.",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the list.",
				Required:            true,
			},
			"method": schema.StringAttribute{
				MarkdownDescription: "`internal` (generated and signed by OPNsense, needs the CA's private key) or `existing` (imported from `import_pem`). Defaults to `internal`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("internal"),
				Validators:          []validator.String{stringvalidator.OneOf("internal", "existing")},
			},
			"lifetime": schema.Int64Attribute{
				MarkdownDescription: "Validity of a generated list, in days. Defaults to `9999`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(9999),
			},
			"import_pem": schema.StringAttribute{
				MarkdownDescription: "PEM-encoded CRL to store, with `method = \"existing\"`.",
				Optional:            true,
			},
			"revoked": schema.SetNestedAttribute{
				MarkdownDescription: "Certificates revoked by a generated list. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.ObjectType{AttrTypes: crlRevokedTypes})),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"certificate": schema.StringAttribute{
							MarkdownDescription: "Reference ID of the certificate, e.g. `opnsense_trust_cert.laptop.ref_id`.",
							Required:            true,
						},
						"reason": schema.StringAttribute{
							MarkdownDescription: "Revocation reason: `unspecified`, `keyCompromise`, `cACompromise`, `affiliationChanged`, `superseded`, `cessationOfOperation` or `certificateHold`.",
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf(crlReasons...)},
						},
					},
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: "Serial number of the list, increased by OPNsense on every change.",
				Computed:            true,
			},
			"pem": schema.StringAttribute{
				MarkdownDescription: "Current list, PEM-encoded.",
				Computed:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Same as `ca`.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *crlResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *crlResource) save(ctx context.Context, data *crlResourceModel) error {
	payload := &trust.CrlData{
		Descr:     data.Description.ValueString(),
		Crlmethod: api.SelectedMap(data.Method.ValueString()),
		Lifetime:  strconv.FormatInt(data.Lifetime.ValueInt64(), 10),
		Text:      data.ImportPem.ValueString(),
	}
	var revoked []crlRevokedModel
	data.Revoked.ElementsAs(ctx, &revoked, false)
	byReason := make([]api.SelectedMapList, len(crlReasons))
	for i := range byReason {
		byReason[i] = api.SelectedMapList{}
	}
	for _, rv := range revoked {
		for i, name := range crlReasons {
			if name == rv.Reason.ValueString() {
				byReason[i] = append(byReason[i], rv.Certificate.ValueString())
			}
		}
	}
	payload.RevokedReason0, payload.RevokedReason1, payload.RevokedReason2 = byReason[0], byReason[1], byReason[2]
	payload.RevokedReason3, payload.RevokedReason4, payload.RevokedReason5 = byReason[3], byReason[4], byReason[5]
	payload.RevokedReason6 = byReason[6]
	res, err := r.client.Trust().CrlSet(ctx, data.Ca.ValueString(), payload)
	if err != nil {
		return err
	}
	if res.Status != "saved" {
		msgs := []string{}
		for k, v := range res.Validations {
			msgs = append(msgs, k+": "+v)
		}
		sort.Strings(msgs)
		return fmt.Errorf("status %q: %s", res.Status, strings.Join(msgs, "; "))
	}
	return nil
}

func (r *crlResource) read(ctx context.Context, data *crlResourceModel) (bool, error) {
	res, err := r.client.Trust().CrlGet(ctx, data.Ca.ValueString())
	if err != nil {
		return false, err
	}
	c := res.Crl
	if c.Descr == "" && c.Text == "" {
		return false, nil
	}
	revoked := []crlRevokedModel{}
	for i, refs := range [][]string{c.RevokedReason0, c.RevokedReason1, c.RevokedReason2, c.RevokedReason3, c.RevokedReason4, c.RevokedReason5, c.RevokedReason6} {
		for _, ref := range refs {
			revoked = append(revoked, crlRevokedModel{Certificate: types.StringValue(ref), Reason: types.StringValue(crlReasons[i])})
		}
	}
	set, diags := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: crlRevokedTypes}, revoked)
	if diags.HasError() {
		return false, fmt.Errorf("unable to build the revoked set")
	}
	lifetime, _ := strconv.ParseInt(c.Lifetime, 10, 64)
	data.Description = types.StringValue(c.Descr)
	data.Method = types.StringValue(c.Crlmethod.String())
	data.Lifetime = types.Int64Value(lifetime)
	data.Revoked = set
	data.Serial = types.StringValue(c.Serial)
	data.Pem = types.StringValue(c.Text)
	data.Id = data.Ca
	return true, nil
}

func (r *crlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data crlResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.save(ctx, &data); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create the CRL, got error: %s", err))
		return
	}
	if ok, err := r.read(ctx, &data); err != nil || !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read the CRL after creating it: %v", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *crlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data crlResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	ok, err := r.read(ctx, &data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read the CRL, got error: %s", err))
		return
	}
	if !ok {
		resp.State.RemoveResource(ctx)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *crlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data crlResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.save(ctx, &data); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update the CRL, got error: %s", err))
		return
	}
	if ok, err := r.read(ctx, &data); err != nil || !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read the CRL after updating it: %v", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *crlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data crlResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Trust().CrlDelete(ctx, data.Ca.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete the CRL, got error: %s", err))
		return
	}
	if res.Status != "deleted" && res.Status != "not found" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete the CRL, status %q", res.Status))
	}
}

func (r *crlResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ca"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("revoked"), types.SetNull(types.ObjectType{AttrTypes: crlRevokedTypes}))...)
}
