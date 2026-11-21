package radvd

import (
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func setToSlice(s types.Set) []string {
	out := tools.SetToStringSlice(s)
	if out == nil {
		return []string{}
	}
	return out
}
