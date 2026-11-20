package trafficshaper

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var defaultsSuffix = regexp.MustCompile(` Defaults to [^.]*\.$`)

var maskValues = []string{"none", "src-ip", "dst-ip", "src-ip6", "dst-ip6"}

func idAttribute(what string) schema.StringAttribute {
	return schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "UUID of the " + what + ".",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
}

func numberAttribute(what string) schema.Int64Attribute {
	return schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Number OPNsense assigned to the " + what + " (ipfw " + what + " number).",
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
	}
}

func enabledAttribute() schema.BoolAttribute {
	return optionalBool("Whether this entry is enabled. Defaults to `true`.", true)
}

func optionalBool(description string, def bool) schema.BoolAttribute {
	return schema.BoolAttribute{
		MarkdownDescription: description,
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(def),
	}
}

func optionalSentinelInt64(description string, min, max int64) schema.Int64Attribute {
	return schema.Int64Attribute{
		MarkdownDescription: description + " Defaults to `-1` (OPNsense default).",
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(-1),
		Validators: []validator.Int64{
			int64validator.Any(int64validator.OneOf(-1), int64validator.Between(min, max)),
		},
	}
}

func maskAttribute() schema.StringAttribute {
	return schema.StringAttribute{
		MarkdownDescription: "Create a dynamic queue per address: `none`, `src-ip`, `dst-ip`, `src-ip6` or `dst-ip6`. Masking by source or destination shares the bandwidth per host instead of across all traffic. Defaults to `none`.",
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("none"),
		Validators: []validator.String{
			stringvalidator.OneOf(maskValues...),
		},
	}
}

func codelAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"codel_enable":     optionalBool("When enabled, CoDel active queue management is used. Cannot be combined with `pie_enable`. Defaults to `false`.", false),
		"codel_target":     optionalSentinelInt64("CoDel target queue delay in milliseconds.", 1, 10000),
		"codel_interval":   optionalSentinelInt64("CoDel interval in milliseconds.", 1, 10000),
		"codel_ecn_enable": optionalBool("When enabled, CoDel marks packets with ECN instead of dropping them. Defaults to `false`.", false),
		"pie_enable":       optionalBool("When enabled, PIE active queue management is used. Cannot be combined with `codel_enable`. Defaults to `false`.", false),
	}
}

func dataSourceComputed(attrs map[string]schema.Attribute) map[string]dschema.Attribute {
	out := map[string]dschema.Attribute{}
	for name, a := range attrs {
		description := defaultsSuffix.ReplaceAllString(a.GetMarkdownDescription(), "")
		switch a.(type) {
		case schema.BoolAttribute:
			out[name] = dschema.BoolAttribute{MarkdownDescription: description, Computed: true}
		case schema.Int64Attribute:
			out[name] = dschema.Int64Attribute{MarkdownDescription: description, Computed: true}
		case schema.SetAttribute:
			out[name] = dschema.SetAttribute{MarkdownDescription: description, Computed: true, ElementType: a.(schema.SetAttribute).ElementType}
		default:
			out[name] = dschema.StringAttribute{MarkdownDescription: description, Computed: true}
		}
	}
	out["id"] = dschema.StringAttribute{MarkdownDescription: "UUID of the resource.", Required: true}
	return out
}
