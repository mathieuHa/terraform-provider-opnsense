package bind_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const bindStack = `
resource "opnsense_bind_acl" "lan" {
  name     = "tfacc_lan"
  networks = ["10.0.0.0/8"]
}

resource "opnsense_bind_primary_domain" "lab" {
  domainname = "tfacc.lab.example"
  allowquery = [opnsense_bind_acl.lan.id]
  ttl        = %[1]d
}

resource "opnsense_bind_record" "www" {
  domain = opnsense_bind_primary_domain.lab.id
  name   = "www"
  type   = "A"
  value  = "%[2]s"
}

resource "opnsense_bind_secondary_domain" "partner" {
  domainname      = "partner.tfacc.example"
  primaryip       = ["192.0.2.53"]
  transferkeyalgo = "hmac-sha256"
  transferkeyname = "tfacc-xfer"
  transferkey     = "dGZhY2MtdHNpZy1rZXk="
}

resource "opnsense_bind_forward_domain" "corp" {
  domainname    = "corp.tfacc.example"
  forwardserver = ["192.0.2.10", "192.0.2.11"]
}
`

func TestAccBindStack(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/bind/general/get", "plugin os-bind") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(bindStack, 3600, "192.0.2.80"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("opnsense_bind_record.www", "domain", "opnsense_bind_primary_domain.lab", "id"),
					resource.TestCheckResourceAttr("opnsense_bind_primary_domain.lab", "allowquery.#", "1"),
					resource.TestCheckResourceAttr("opnsense_bind_secondary_domain.partner", "transferkeyalgo", "hmac-sha256"),
					resource.TestCheckResourceAttr("opnsense_bind_forward_domain.corp", "forwardserver.#", "2"),
				),
			},
			{
				ResourceName:      "opnsense_bind_primary_domain.lab",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "opnsense_bind_secondary_domain.partner",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: fmt.Sprintf(bindStack, 7200, "192.0.2.81"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_bind_primary_domain.lab", "ttl", "7200"),
					resource.TestCheckResourceAttr("opnsense_bind_record.www", "value", "192.0.2.81"),
				),
			},
		},
	})
}

func TestAccBindSettingsResource(t *testing.T) {
	a := "opnsense_bind_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/bind/general/get", "plugin os-bind") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_bind_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "bind_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_bind_settings" "test" {
  enabled    = false
  forwarders = ["192.0.2.53", "192.0.2.54"]
  hideversion = true
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "forwarders.#", "2"),
					resource.TestCheckResourceAttr(a, "hideversion", "true"),
				),
			},
			{
				Config: `
resource "opnsense_bind_settings" "test" {
  enabled     = false
  forwarders  = []
  hideversion = false
}
`,
				Check: resource.TestCheckResourceAttr(a, "forwarders.#", "0"),
			},
		},
	})
}
