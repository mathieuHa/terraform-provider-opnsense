package kea

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func optionsList(s types.Set) api.SelectedMapList {
	out := tools.SetToStringSlice(s)
	if out == nil {
		return api.SelectedMapList{}
	}
	return api.SelectedMapList(out)
}
