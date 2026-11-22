package nut_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNutSettingsResource(t *testing.T) {
	a := "opnsense_nut_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/nut/settings/get", "plugin os-nut") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_nut_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "nut_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_nut_settings" "test" {
  general_enable         = false
  general_name           = "tfacc"
  general_listen         = ["127.0.0.1", "::1", "192.0.2.10"]
  account_admin_password = "tfacc-admin"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "general_name", "tfacc"),
					resource.TestCheckResourceAttr(a, "general_listen.#", "3"),
					resource.TestCheckResourceAttr(a, "account_admin_password", "tfacc-admin"),
				),
			},
			{
				Config: `
resource "opnsense_nut_settings" "test" {
  general_enable         = false
  general_name           = "UPSName"
  general_listen         = ["127.0.0.1", "::1"]
  account_admin_password = "tfacc-admin-2"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "general_name", "UPSName"),
				),
			},
		},
	})
}
