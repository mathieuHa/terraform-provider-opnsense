package ipsec_test

import (
	"regexp"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIpsecSettingsResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_ipsec_settings" "test" {}`,
				ResourceName:       "opnsense_ipsec_settings.test",
				ImportState:        true,
				ImportStateId:      "ipsec_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_ipsec_settings" "test" {
  enabled              = true
  make_before_break    = true
  threads              = 8
  passthrough_networks = ["10.199.0.0/24"]
  retransmit_timeout   = "4.5"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_ipsec_settings.test", "id", "ipsec_settings"),
					resource.TestCheckResourceAttr("opnsense_ipsec_settings.test", "enabled", "true"),
					resource.TestCheckResourceAttr("opnsense_ipsec_settings.test", "make_before_break", "true"),
					resource.TestCheckResourceAttr("opnsense_ipsec_settings.test", "threads", "8"),
					resource.TestCheckTypeSetElemAttr("opnsense_ipsec_settings.test", "passthrough_networks.*", "10.199.0.0/24"),
					resource.TestCheckResourceAttr("opnsense_ipsec_settings.test", "retransmit_timeout", "4.5"),
				),
			},
			{
				Config: `
resource "opnsense_ipsec_settings" "test" {
  enabled              = false
  make_before_break    = false
  threads              = 16
  passthrough_networks = []
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_ipsec_settings.test", "enabled", "false"),
					resource.TestCheckResourceAttr("opnsense_ipsec_settings.test", "threads", "16"),
					resource.TestCheckResourceAttr("opnsense_ipsec_settings.test", "passthrough_networks.#", "0"),
					resource.TestCheckResourceAttr("opnsense_ipsec_settings.test", "retransmit_timeout", "4.5"),
				),
			},
		},
	})
}

func TestAccIpsecSettingsResource_CreateBlocked(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      `resource "opnsense_ipsec_settings" "test" {}`,
				ExpectError: regexp.MustCompile("Cannot Create Singleton Resource"),
			},
		},
	})
}
