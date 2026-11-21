package ipsec

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/ipsec"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type keyPairResourceModel struct {
	Name           types.String `tfsdk:"name"`
	KeyType        types.String `tfsdk:"key_type"`
	PublicKey      types.String `tfsdk:"public_key"`
	PrivateKey     types.String `tfsdk:"private_key"`
	KeySize        types.String `tfsdk:"key_size"`
	KeyFingerprint types.String `tfsdk:"key_fingerprint"`

	Id types.String `tfsdk:"id"`
}

func keyPairResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Raw public/private key pairs for IPsec public key authentication (**VPN ▸ IPsec ▸ Key Pairs**).",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the key pair.",
				Required:            true,
			},
			"key_type": schema.StringAttribute{
				MarkdownDescription: "Key type. Defaults to `rsa`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("rsa"),
				Validators:          []validator.String{stringvalidator.OneOf("rsa", "ecdsa")},
			},
			"public_key": schema.StringAttribute{
				MarkdownDescription: "PEM public key.",
				Required:            true,
			},
			"private_key": schema.StringAttribute{
				MarkdownDescription: "PEM private key. Leave unset for a peer key used only for verification.",
				Optional:            true,
				Sensitive:           true,
			},
			"key_size": schema.StringAttribute{
				MarkdownDescription: "Key size, read from the key.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"key_fingerprint": schema.StringAttribute{
				MarkdownDescription: "Key fingerprint, read from the key.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the IPsec key pair.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func keyPairDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Raw public/private key pairs for IPsec public key authentication (**VPN ▸ IPsec ▸ Key Pairs**).",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the IPsec key pair.",
				Required:            true,
			},
			"name":            dschema.StringAttribute{MarkdownDescription: "Name of the key pair.", Computed: true},
			"key_type":        dschema.StringAttribute{MarkdownDescription: "Key type.", Computed: true},
			"public_key":      dschema.StringAttribute{MarkdownDescription: "PEM public key.", Computed: true},
			"private_key":     dschema.StringAttribute{MarkdownDescription: "PEM private key. Never returned by OPNsense, always null.", Computed: true, Sensitive: true},
			"key_size":        dschema.StringAttribute{MarkdownDescription: "Key size, read from the key.", Computed: true},
			"key_fingerprint": dschema.StringAttribute{MarkdownDescription: "Key fingerprint, read from the key.", Computed: true},
		},
	}
}

func convertKeyPairSchemaToStruct(d *keyPairResourceModel) (*ipsec.KeyPair, error) {
	return &ipsec.KeyPair{
		Name:           d.Name.ValueString(),
		KeyType:        api.SelectedMap(d.KeyType.ValueString()),
		PublicKey:      d.PublicKey.ValueString(),
		PrivateKey:     d.PrivateKey.ValueString(),
		KeySize:        d.KeySize.ValueString(),
		KeyFingerprint: d.KeyFingerprint.ValueString(),
	}, nil
}

func convertKeyPairStructToSchema(d *ipsec.KeyPair) (*keyPairResourceModel, error) {
	return &keyPairResourceModel{
		Name:           types.StringValue(d.Name),
		KeyType:        types.StringValue(d.KeyType.String()),
		PublicKey:      types.StringValue(d.PublicKey),
		PrivateKey:     types.StringNull(),
		KeySize:        tools.StringOrNull(d.KeySize),
		KeyFingerprint: tools.StringOrNull(d.KeyFingerprint),
	}, nil
}
