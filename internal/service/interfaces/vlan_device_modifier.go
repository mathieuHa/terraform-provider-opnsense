package interfaces

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

type keepGeneratedDevice struct{}

func (keepGeneratedDevice) Description(context.Context) string {
	return "Keeps the device name OPNsense generated when the configuration does not set it."
}

func (m keepGeneratedDevice) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (keepGeneratedDevice) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() || req.StateValue.ValueString() == "" {
		return
	}
	if req.ConfigValue.IsNull() {
		resp.PlanValue = req.StateValue
	}
}
