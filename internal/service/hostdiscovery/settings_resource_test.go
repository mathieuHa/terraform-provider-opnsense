package hostdiscovery_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHostdiscoverySettingsResource(t *testing.T) {
	a := "opnsense_hostdiscovery_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/hostdiscovery/settings/get", "OPNsense 26.1 or later") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_hostdiscovery_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "hostdiscovery_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_hostdiscovery_settings" "test" {
  enabled          = true
  skip_networks    = ["192.0.2.0/24"]
  expire4_interval = 7200
  interfaces       = ["wan"]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "skip_networks.#", "1"),
					resource.TestCheckResourceAttr(a, "expire4_interval", "7200"),
					resource.TestCheckTypeSetElemAttr(a, "interfaces.*", "wan"),
				),
			},
			{
				Config: `
resource "opnsense_hostdiscovery_settings" "test" {
  enabled          = true
  skip_networks    = []
  expire4_interval = -1
  interfaces       = []
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "skip_networks.#", "0"),
					resource.TestCheckResourceAttr(a, "expire4_interval", "-1"),
					resource.TestCheckResourceAttr(a, "interfaces.#", "0"),
				),
			},
		},
	})
}
