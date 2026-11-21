package caddy_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const caddyStack = `
resource "opnsense_caddy_reverse_proxy" "wild" {
  from_domain = "*.tfacc.example.com"
  description = "tfacc wildcard"
}

resource "opnsense_caddy_subdomain" "app" {
  reverse     = opnsense_caddy_reverse_proxy.wild.id
  from_domain = "app.tfacc.example.com"
  description = "tfacc app"
}

resource "opnsense_caddy_access_list" "lan" {
  name       = "tfacc-lan"
  client_ips = ["10.0.0.0/8"]
}

resource "opnsense_caddy_basic_auth" "user" {
  user     = "tfacc"
  password = "tfacc-secret"
}

resource "opnsense_caddy_header" "xff" {
  header = "X-Tfacc"
  value  = "%[1]s"
}

resource "opnsense_caddy_handle" "app" {
  reverse     = opnsense_caddy_reverse_proxy.wild.id
  subdomain   = opnsense_caddy_subdomain.app.id
  handle_path = "/api/*"
  accesslist  = opnsense_caddy_access_list.lan.id
  basicauth   = [opnsense_caddy_basic_auth.user.id]
  header      = [opnsense_caddy_header.xff.id]
  to_domain   = ["10.99.2.10", "10.99.2.11"]
  to_port     = %[2]d
  lb_policy   = "round_robin"
  description = "tfacc handler"
}

resource "opnsense_caddy_layer4" "ssh" {
  type        = "global"
  matchers    = "ssh"
  from_port   = 2222
  to_domain   = ["10.99.2.10"]
  to_port     = 22
  description = "tfacc ssh"
}
`

func TestAccCaddyStack(t *testing.T) {
	h := "opnsense_caddy_handle.app"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/caddy/general/get", "plugin os-caddy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(caddyStack, "one", 8080),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(h, "reverse", "opnsense_caddy_reverse_proxy.wild", "id"),
					resource.TestCheckResourceAttrPair(h, "subdomain", "opnsense_caddy_subdomain.app", "id"),
					resource.TestCheckResourceAttr(h, "to_domain.#", "2"),
					resource.TestCheckResourceAttr(h, "to_port", "8080"),
					resource.TestCheckResourceAttr(h, "handle_directive", "reverse_proxy"),
					resource.TestCheckResourceAttr("opnsense_caddy_layer4.ssh", "protocol", "tcp"),
					resource.TestCheckResourceAttr("opnsense_caddy_layer4.ssh", "matchers", "ssh"),
				),
			},
			{
				ResourceName:      h,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:            "opnsense_caddy_basic_auth.user",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
			{
				Config: fmt.Sprintf(caddyStack, "two", 9090),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(h, "to_port", "9090"),
					resource.TestCheckResourceAttr("opnsense_caddy_header.xff", "value", "two"),
				),
			},
		},
	})
}

func TestAccCaddySettingsResource(t *testing.T) {
	a := "opnsense_caddy_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/caddy/general/get", "plugin os-caddy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_caddy_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "caddy_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_caddy_settings" "test" {
  enabled       = false
  tls_email     = "admin@example.com"
  http_versions = ["h1", "h2", "h3"]
  grace_period  = 20
  log_level     = "WARN"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "tls_email", "admin@example.com"),
					resource.TestCheckResourceAttr(a, "http_versions.#", "3"),
					resource.TestCheckResourceAttr(a, "grace_period", "20"),
					resource.TestCheckResourceAttr(a, "log_level", "WARN"),
				),
			},
			{
				Config: `
resource "opnsense_caddy_settings" "test" {
  enabled       = false
  tls_email     = ""
  http_versions = ["h1", "h2"]
  grace_period  = 10
  log_level     = ""
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "tls_email", ""),
					resource.TestCheckResourceAttr(a, "http_versions.#", "2"),
				),
			},
		},
	})
}
