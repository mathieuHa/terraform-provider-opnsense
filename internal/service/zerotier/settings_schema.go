package zerotier

import (
	"github.com/browningluke/opnsense-go/pkg/zerotier"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type settingsResourceModel struct {
	Enabled        types.Bool   `tfsdk:"enabled"`
	ApiAccessToken types.String `tfsdk:"api_access_token"`
	Localconf      types.String `tfsdk:"localconf"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "This will activate the Zerotier service.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"api_access_token": schema.StringAttribute{
			MarkdownDescription: "Access Token used for the Zerotier API service.",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"localconf": schema.StringAttribute{
			MarkdownDescription: "A JSON document containing configuration options that apply to the local node.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `zerotier_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func settingsBool(b types.Bool) string {
	if b.IsNull() || b.IsUnknown() {
		return ""
	}
	return tools.BoolToString(b.ValueBool())
}

func settingsInt(i types.Int64) string {
	if i.IsNull() || i.IsUnknown() || i.ValueInt64() == -1 {
		return ""
	}
	return tools.Int64ToString(i.ValueInt64())
}

func settingsReadInt(s string) types.Int64 {
	if s == "" {
		return types.Int64Value(-1)
	}
	return tools.StringToInt64Null(s)
}

func settingsList(s types.Set) []string {
	out := tools.SetToStringSlice(s)
	if out == nil {
		return []string{}
	}
	return out
}

func settingsNonEmpty(in []string) []string {
	out := []string{}
	for _, v := range in {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

func convertSettingsSchemaToStruct(d *settingsResourceModel) *zerotier.Settings {
	return &zerotier.Settings{
		Enabled:        settingsBool(d.Enabled),
		ApiAccessToken: d.ApiAccessToken.ValueString(),
		Localconf:      d.Localconf.ValueString(),
	}
}

func convertSettingsStructToSchema(s *zerotier.Settings) *settingsResourceModel {
	return &settingsResourceModel{
		Enabled:        types.BoolValue(tools.StringToBool(s.Enabled)),
		ApiAccessToken: types.StringValue(s.ApiAccessToken),
		Localconf:      types.StringValue(s.Localconf),
	}
}
