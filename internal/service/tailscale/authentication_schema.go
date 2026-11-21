package tailscale

import (
	"github.com/browningluke/opnsense-go/pkg/tailscale"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type authenticationResourceModel struct {
	LoginServer types.String `tfsdk:"login_server"`
	PreAuthKey  types.String `tfsdk:"pre_auth_key"`

	Id types.String `tfsdk:"id"`
}

func authenticationAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"login_server": schema.StringAttribute{
			MarkdownDescription: "Coordination server URL, e.g. a Headscale instance.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"pre_auth_key": schema.StringAttribute{
			MarkdownDescription: "Pre-authentication key used to join the tailnet.",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `tailscale_authentication`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertAuthenticationSchemaToStruct(d *authenticationResourceModel) *tailscale.Authentication {
	return &tailscale.Authentication{
		LoginServer: d.LoginServer.ValueString(),
		PreAuthKey:  d.PreAuthKey.ValueString(),
	}
}

func convertAuthenticationStructToSchema(s *tailscale.Authentication) *authenticationResourceModel {
	return &authenticationResourceModel{
		LoginServer: types.StringValue(s.LoginServer),
		PreAuthKey:  types.StringValue(s.PreAuthKey),
	}
}
