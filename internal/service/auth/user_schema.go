package auth

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/auth"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type userResourceModel struct {
	Name              types.String `tfsdk:"name"`
	Uid               types.String `tfsdk:"uid"`
	Enabled           types.Bool   `tfsdk:"enabled"`
	Password          types.String `tfsdk:"password"`
	ScrambledPassword types.Bool   `tfsdk:"scrambled_password"`
	FullName          types.String `tfsdk:"full_name"`
	Email             types.String `tfsdk:"email"`
	Comment           types.String `tfsdk:"comment"`
	Expires           types.String `tfsdk:"expires"`
	AuthorizedKeys    types.String `tfsdk:"authorized_keys"`
	Shell             types.String `tfsdk:"shell"`
	LandingPage       types.String `tfsdk:"landing_page"`
	Language          types.String `tfsdk:"language"`
	Privileges        types.Set    `tfsdk:"privileges"`
	GroupMemberships  types.Set    `tfsdk:"group_memberships"`

	Id types.String `tfsdk:"id"`
}

func userResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Local users (**System ▸ Access ▸ Users**). API keys and OTP enrolment made outside Terraform are kept on update.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Login name.",
				Required:            true,
			},
			"uid": schema.StringAttribute{
				MarkdownDescription: "User ID assigned by OPNsense.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether the account can log in. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password. Never read back from OPNsense; leave unset to keep the current one.",
				Optional:            true,
				Sensitive:           true,
			},
			"scrambled_password": schema.BoolAttribute{
				MarkdownDescription: "When enabled, OPNsense sets a random password, for accounts that only use API keys or SSH keys. The password is scrambled when the user is created with it, or when it changes from `false` to `true`; an imported user is recorded as `true`, so importing never scrambles an existing password. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"full_name": schema.StringAttribute{
				MarkdownDescription: "Full name of the user.",
				Optional:            true,
			},
			"email": schema.StringAttribute{
				MarkdownDescription: "Email address.",
				Optional:            true,
			},
			"comment": schema.StringAttribute{
				MarkdownDescription: "Free-form comment.",
				Optional:            true,
			},
			"expires": schema.StringAttribute{
				MarkdownDescription: "Expiration date in `MM/DD/YYYY` format; leave unset for no expiry.",
				Optional:            true,
			},
			"authorized_keys": schema.StringAttribute{
				MarkdownDescription: "SSH public keys allowed for this user, one per line.",
				Optional:            true,
			},
			"shell": schema.StringAttribute{
				MarkdownDescription: "Login shell for SSH and console access; unset means no shell.",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("/bin/csh", "/bin/sh", "/bin/tcsh")},
			},
			"landing_page": schema.StringAttribute{
				MarkdownDescription: "Page shown after login.",
				Optional:            true,
			},
			"language": schema.StringAttribute{
				MarkdownDescription: "GUI language, e.g. `en_US`; unset means the system default.",
				Optional:            true,
			},
			"privileges": schema.SetAttribute{
				MarkdownDescription: "Privileges granted to the user directly, on top of those of its groups. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"group_memberships": schema.SetAttribute{
				MarkdownDescription: "Group IDs (`opnsense_auth_group.gid`) the user belongs to. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the user.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func userDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Local users (**System ▸ Access ▸ Users**). API keys and OTP enrolment made outside Terraform are kept on update.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the user.",
				Required:            true,
			},
			"name":               dschema.StringAttribute{MarkdownDescription: "Login name.", Computed: true},
			"uid":                dschema.StringAttribute{MarkdownDescription: "User ID assigned by OPNsense.", Computed: true},
			"enabled":            dschema.BoolAttribute{MarkdownDescription: "Whether the account can log in.", Computed: true},
			"password":           dschema.StringAttribute{MarkdownDescription: "Password. Never returned by OPNsense, always null.", Computed: true, Sensitive: true},
			"scrambled_password": dschema.BoolAttribute{MarkdownDescription: "When enabled, OPNsense sets a random password, for accounts that only use API keys or SSH keys.", Computed: true},
			"full_name":          dschema.StringAttribute{MarkdownDescription: "Full name of the user.", Computed: true},
			"email":              dschema.StringAttribute{MarkdownDescription: "Email address.", Computed: true},
			"comment":            dschema.StringAttribute{MarkdownDescription: "Free-form comment.", Computed: true},
			"expires":            dschema.StringAttribute{MarkdownDescription: "Expiration date in `MM/DD/YYYY` format; leave unset for no expiry.", Computed: true},
			"authorized_keys":    dschema.StringAttribute{MarkdownDescription: "SSH public keys allowed for this user, one per line.", Computed: true},
			"shell":              dschema.StringAttribute{MarkdownDescription: "Login shell for SSH and console access; unset means no shell.", Computed: true},
			"landing_page":       dschema.StringAttribute{MarkdownDescription: "Page shown after login.", Computed: true},
			"language":           dschema.StringAttribute{MarkdownDescription: "GUI language, e.g. `en_US`; unset means the system default.", Computed: true},
			"privileges":         dschema.SetAttribute{MarkdownDescription: "Privileges granted to the user directly, on top of those of its groups.", Computed: true, ElementType: types.StringType},
			"group_memberships":  dschema.SetAttribute{MarkdownDescription: "Group IDs (`opnsense_auth_group.gid`) the user belongs to.", Computed: true, ElementType: types.StringType},
		},
	}
}

func convertUserSchemaToStruct(d *userResourceModel) (*auth.User, error) {
	return &auth.User{
		Name:              d.Name.ValueString(),
		UserId:            d.Uid.ValueString(),
		Disabled:          tools.BoolToString(!d.Enabled.ValueBool()),
		Password:          d.Password.ValueString(),
		ScrambledPassword: tools.BoolToString(d.ScrambledPassword.ValueBool()),
		Fullname:          d.FullName.ValueString(),
		Email:             d.Email.ValueString(),
		Comment:           d.Comment.ValueString(),
		Expires:           d.Expires.ValueString(),
		AuthorizedKeys:    d.AuthorizedKeys.ValueString(),
		Shell:             api.SelectedMap(d.Shell.ValueString()),
		LandingPage:       d.LandingPage.ValueString(),
		Language:          api.SelectedMap(d.Language.ValueString()),
		Privilege:         api.SelectedMapList(setToSlice(d.Privileges)),
		GroupMemberships:  api.SelectedMapList(setToSlice(d.GroupMemberships)),
	}, nil
}

func convertUserStructToSchema(d *auth.User) (*userResourceModel, error) {
	return &userResourceModel{
		Name:              types.StringValue(d.Name),
		Uid:               tools.StringOrNull(d.UserId),
		Enabled:           types.BoolValue(!tools.StringToBool(d.Disabled)),
		Password:          types.StringNull(),
		ScrambledPassword: types.BoolValue(tools.StringToBool(d.ScrambledPassword)),
		FullName:          tools.StringOrNull(d.Fullname),
		Email:             tools.StringOrNull(d.Email),
		Comment:           tools.StringOrNull(d.Comment),
		Expires:           tools.StringOrNull(d.Expires),
		AuthorizedKeys:    tools.StringOrNull(d.AuthorizedKeys),
		Shell:             tools.StringOrNull(d.Shell.String()),
		LandingPage:       tools.StringOrNull(d.LandingPage),
		Language:          tools.StringOrNull(d.Language.String()),
		Privileges:        tools.StringSliceToSet(d.Privilege),
		GroupMemberships:  tools.StringSliceToSet(d.GroupMemberships),
	}, nil
}
