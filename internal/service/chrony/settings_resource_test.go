package chrony_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccChronySettingsResource(t *testing.T) {
	a := "opnsense_chrony_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/chrony/general/get", "plugin os-chrony") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_chrony_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "chrony_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_chrony_settings" "test" {
  enabled          = false
  nts_client       = true
  peers            = ["time.cloudflare.com", "nts.netnod.se"]
  allowed_networks = ["10.0.0.0/8"]
  fallback_peers   = "pool.ntp.org"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "nts_client", "true"),
					resource.TestCheckResourceAttr(a, "peers.#", "2"),
					resource.TestCheckResourceAttr(a, "allowed_networks.#", "1"),
					resource.TestCheckResourceAttr(a, "fallback_peers", "pool.ntp.org"),
				),
			},
			{
				Config: `
resource "opnsense_chrony_settings" "test" {
  enabled          = false
  nts_client       = false
  peers            = ["0.opnsense.pool.ntp.org"]
  allowed_networks = []
  fallback_peers   = ""
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "peers.#", "1"),
					resource.TestCheckResourceAttr(a, "allowed_networks.#", "0"),
				),
			},
		},
	})
}
