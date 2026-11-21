package core

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var pluginNameRegexp = regexp.MustCompile(`^[a-z0-9][a-z0-9._+-]*$`)

const (
	firmwarePollInterval = 3 * time.Second
	firmwareJobTimeout   = 20 * time.Minute
)

var _ resource.Resource = &firmwarePluginResource{}
var _ resource.ResourceWithConfigure = &firmwarePluginResource{}
var _ resource.ResourceWithImportState = &firmwarePluginResource{}

type firmwarePluginResourceModel struct {
	Name    types.String `tfsdk:"name"`
	Version types.String `tfsdk:"version"`
	Id      types.String `tfsdk:"id"`
}

func newFirmwarePluginResource() resource.Resource {
	return &firmwarePluginResource{}
}

type firmwarePluginResource struct {
	client opnsense.Client
}

func (r *firmwarePluginResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firmware_plugin"
}

func (r *firmwarePluginResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Installs an OPNsense plugin package (**System ▸ Firmware ▸ Plugins**), e.g. `os-squid` before managing `opnsense_proxy_settings`. The firewall must reach its package mirror, and OPNsense refuses to install plugins while the firmware itself is out of date. Destroying the resource removes the plugin.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Package name, e.g. `os-squid`.",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Validators:          []validator.String{stringvalidator.RegexMatches(pluginNameRegexp, "must be a package name such as os-squid")},
			},
			"version": schema.StringAttribute{
				MarkdownDescription: "Installed version.",
				Computed:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Same as `name`.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *firmwarePluginResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *firmwarePluginResource) installed(ctx context.Context, name string) (string, bool, error) {
	info, err := r.client.Core().FirmwareInfo(ctx)
	if err != nil {
		return "", false, err
	}
	for _, p := range info.Packages {
		if p.Name == name && p.Installed == "1" {
			return p.Version, true, nil
		}
	}
	return "", false, nil
}

func (r *firmwarePluginResource) wait(ctx context.Context) (string, error) {
	deadline := time.Now().Add(firmwareJobTimeout)
	log := ""
	for time.Now().Before(deadline) {
		select {
		case <-ctx.Done():
			return log, ctx.Err()
		case <-time.After(firmwarePollInterval):
		}
		st, err := r.client.Core().FirmwareUpgradeStatus(ctx)
		if err != nil {
			continue
		}
		if st.Log != "" {
			log = st.Log
		}
		if st.Status == "done" || st.Status == "reboot" {
			return log, nil
		}
	}
	return log, fmt.Errorf("firmware job did not finish within %s", firmwareJobTimeout)
}

func logTail(log string) string {
	lines := strings.Split(strings.TrimSpace(log), "\n")
	if len(lines) > 15 {
		lines = lines[len(lines)-15:]
	}
	return strings.Join(lines, "\n")
}

func (r *firmwarePluginResource) waitIdle(ctx context.Context) error {
	deadline := time.Now().Add(firmwareJobTimeout)
	for time.Now().Before(deadline) {
		st, err := r.client.Core().FirmwareRunning(ctx)
		if err == nil && st.Status == "ready" {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(firmwarePollInterval):
		}
	}
	return fmt.Errorf("another firmware job is still running after %s", firmwareJobTimeout)
}

func (r *firmwarePluginResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data firmwarePluginResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	name := data.Name.ValueString()
	if version, ok, err := r.installed(ctx, name); err == nil && ok {
		data.Version = types.StringValue(version)
		data.Id = data.Name
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}
	if err := r.waitIdle(ctx); err != nil {
		resp.Diagnostics.AddError("Firmware Busy", err.Error())
		return
	}
	res, err := r.client.Core().FirmwareInstall(ctx, name)
	if err != nil || res.Status != "ok" {
		status := ""
		if res != nil {
			status = res.Status
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to start the installation of %s: %v (status %q)", name, err, status))
		return
	}
	log, err := r.wait(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Installation of %s failed: %s\n\n%s", name, err, logTail(log)))
		return
	}
	version, ok, err := r.installed(ctx, name)
	if err != nil || !ok {
		resp.Diagnostics.AddError("Plugin Not Installed",
			fmt.Sprintf("%s is not installed after the firmware job finished (%v). Last lines of the log:\n\n%s", name, err, logTail(log)))
		return
	}
	data.Version = types.StringValue(version)
	data.Id = data.Name
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *firmwarePluginResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data firmwarePluginResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	name := data.Id.ValueString()
	version, ok, err := r.installed(ctx, name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read the installed packages, got error: %s", err))
		return
	}
	if !ok {
		resp.State.RemoveResource(ctx)
		return
	}
	data.Name = types.StringValue(name)
	data.Version = types.StringValue(version)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *firmwarePluginResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data firmwarePluginResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *firmwarePluginResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data firmwarePluginResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	name := data.Id.ValueString()
	if err := r.waitIdle(ctx); err != nil {
		resp.Diagnostics.AddError("Firmware Busy", err.Error())
		return
	}
	res, err := r.client.Core().FirmwareRemove(ctx, name)
	if err != nil || res.Status != "ok" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to start the removal of %s: %v", name, err))
		return
	}
	log, err := r.wait(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Removal of %s failed: %s\n\n%s", name, err, logTail(log)))
		return
	}
	if _, ok, err := r.installed(ctx, name); err == nil && ok {
		resp.Diagnostics.AddError("Plugin Still Installed", fmt.Sprintf("%s is still installed after the removal job. Last lines of the log:\n\n%s", name, logTail(log)))
	}
}

func (r *firmwarePluginResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), req.ID)...)
}
