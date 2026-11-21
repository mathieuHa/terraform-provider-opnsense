package acmeclient_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const acmeStack = `
resource "opnsense_acmeclient_account" "le" {
  name  = "tfacc-le"
  email = "tfacc@example.com"
  ca    = "letsencrypt_test"
}

resource "opnsense_acmeclient_validation" "cf" {
  name         = "tfacc-cf"
  method       = "dns01"
  dns_service  = "dns_cf"
  dns_cf_token = "cf-token-tfacc"
}

resource "opnsense_acmeclient_action" "gui" {
  name = "tfacc-gui"
  type = "configd_restart_gui"
}

resource "opnsense_acmeclient_certificate" "test" {
  name              = "%[1]s"
  alt_names         = ["www.%[1]s"]
  account           = opnsense_acmeclient_account.le.id
  validation_method = opnsense_acmeclient_validation.cf.id
  key_length        = "%[2]s"
  restart_actions   = [opnsense_acmeclient_action.gui.id]
  renew_interval    = %[3]d
}
`

func TestAccAcmeClientCertificateResource(t *testing.T) {
	c := "opnsense_acmeclient_certificate.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/acmeclient/settings/get", "plugin os-acme-client") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(acmeStack, "tfacc.example.com", "key_ec256", 60),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(c, "account", "opnsense_acmeclient_account.le", "id"),
					resource.TestCheckResourceAttrPair(c, "validation_method", "opnsense_acmeclient_validation.cf", "id"),
					resource.TestCheckResourceAttr(c, "alt_names.#", "1"),
					resource.TestCheckResourceAttr(c, "restart_actions.#", "1"),
					resource.TestCheckResourceAttr(c, "key_length", "key_ec256"),
					resource.TestCheckResourceAttr(c, "auto_renewal", "true"),
				),
			},
			{
				ResourceName:      c,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: fmt.Sprintf(acmeStack, "tfacc2.example.com", "key_4096", 30),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(c, "name", "tfacc2.example.com"),
					resource.TestCheckResourceAttr(c, "renew_interval", "30"),
				),
			},
		},
	})
}

func TestAccAcmeClientSettingsResource(t *testing.T) {
	a := "opnsense_acmeclient_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/acmeclient/settings/get", "plugin os-acme-client") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_acmeclient_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "acmeclient_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_acmeclient_settings" "test" {
  enabled        = false
  environment    = "stg"
  challenge_port = 43581
  log_level      = "debug"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "environment", "stg"),
					resource.TestCheckResourceAttr(a, "challenge_port", "43581"),
					resource.TestCheckResourceAttr(a, "log_level", "debug"),
				),
			},
			{
				Config: `
resource "opnsense_acmeclient_settings" "test" {
  enabled        = false
  environment    = ""
  challenge_port = 43580
  log_level      = "normal"
}
`,
				Check: resource.TestCheckResourceAttr(a, "environment", ""),
			},
		},
	})
}
