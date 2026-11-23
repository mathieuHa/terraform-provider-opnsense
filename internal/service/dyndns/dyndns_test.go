package dyndns_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDyndnsAccountResource(t *testing.T) {
	a := "opnsense_dyndns_account.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/dyndns/settings/get", "plugin os-ddclient") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "opnsense_dyndns_account" "test" {
  enabled     = false
  service     = "cloudflare"
  username    = "token"
  password    = "tfacc-api-token"
  zone        = "example.com"
  hostnames   = ["fw.example.com"]
  checkip     = "cloudflare-ipv4"
  description = "tfacc"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "service", "cloudflare"),
					resource.TestCheckResourceAttr(a, "hostnames.#", "1"),
				),
			},
			{
				ResourceName:            a,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
			{
				Config: `
resource "opnsense_dyndns_account" "test" {
  enabled     = false
  service     = "cloudflare"
  username    = "token"
  password    = "tfacc-api-token"
  zone        = "example.com"
  hostnames   = ["fw.example.com", "vpn.example.com"]
  checkip     = "cloudflare-ipv4"
  ttl         = 120
  description = "tfacc 2"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "hostnames.#", "2"),
					resource.TestCheckResourceAttr(a, "ttl", "120"),
				),
			},
		},
	})
}

func TestAccDyndnsSettingsResource(t *testing.T) {
	a := "opnsense_dyndns_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/dyndns/settings/get", "plugin os-ddclient") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_dyndns_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "dyndns_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_dyndns_settings" "test" {
  enabled      = false
  daemon_delay = 600
  verbose      = true
}
`,
				Check: resource.TestCheckResourceAttr(a, "daemon_delay", "600"),
			},
			{
				Config: `
resource "opnsense_dyndns_settings" "test" {
  enabled      = false
  daemon_delay = 300
  verbose      = false
}
`,
				Check: resource.TestCheckResourceAttr(a, "daemon_delay", "300"),
			},
		},
	})
}
