package zerotier_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccZerotierSettingsResource(t *testing.T) {
	a := "opnsense_zerotier_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/zerotier/settings/get", "plugin os-zerotier") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_zerotier_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "zerotier_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_zerotier_settings" "test" {
  enabled          = false
  api_access_token = "tfacc-token"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "api_access_token", "tfacc-token"),
				),
			},
			{
				Config: `
resource "opnsense_zerotier_settings" "test" {
  enabled          = false
  api_access_token = ""
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "api_access_token", ""),
				),
			},
		},
	})
}
