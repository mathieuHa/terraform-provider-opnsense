package haproxy_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const haproxyStack = `
resource "opnsense_haproxy_server" "web1" {
  name    = "tfacc-web1"
  address = "10.99.2.10"
  port    = 8080
}

resource "opnsense_haproxy_server" "web2" {
  name    = "tfacc-web2"
  address = "10.99.2.11"
  port    = 8080
}

resource "opnsense_haproxy_healthcheck" "http" {
  name = "tfacc-http"
  type = "http"
}

resource "opnsense_haproxy_backend" "web" {
  name           = "tfacc-web"
  mode           = "http"
  algorithm      = "%[1]s"
  linked_servers = [opnsense_haproxy_server.web1.id, opnsense_haproxy_server.web2.id]
  health_check   = opnsense_haproxy_healthcheck.http.id
}

resource "opnsense_haproxy_acl" "host" {
  name       = "tfacc-host"
  expression = "hdr"
  hdr        = "app.tfacc.example.com"
}

resource "opnsense_haproxy_action" "route" {
  name        = "tfacc-route"
  type        = "use_backend"
  linked_acls = [opnsense_haproxy_acl.host.id]
  use_backend = opnsense_haproxy_backend.web.id
}

resource "opnsense_haproxy_frontend" "http" {
  name            = "tfacc-http"
  bind            = ["127.0.0.1:%[2]d"]
  mode            = "http"
  default_backend = opnsense_haproxy_backend.web.id
  linked_actions  = [opnsense_haproxy_action.route.id]
}
`

func TestAccHAProxyStack(t *testing.T) {
	f := "opnsense_haproxy_frontend.http"
	b := "opnsense_haproxy_backend.web"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/haproxy/settings/get", "plugin os-haproxy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(haproxyStack, "roundrobin", 18080),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(f, "default_backend", b, "id"),
					resource.TestCheckResourceAttr(f, "linked_actions.#", "1"),
					resource.TestCheckResourceAttr(b, "linked_servers.#", "2"),
					resource.TestCheckResourceAttrPair(b, "health_check", "opnsense_haproxy_healthcheck.http", "id"),
					resource.TestCheckResourceAttr(b, "algorithm", "roundrobin"),
					resource.TestCheckResourceAttr(f, "ssl_min_version", "TLSv1.2"),
					resource.TestCheckResourceAttr(f, "advertised_protocols.#", "2"),
					resource.TestCheckTypeSetElemAttr(f, "advertised_protocols.*", "h2"),
					resource.TestCheckTypeSetElemAttr(f, "ssl_bind_options.*", "prefer-client-ciphers"),
				),
			},
			{
				ResourceName:      f,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      b,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: fmt.Sprintf(haproxyStack, "leastconn", 18081),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(b, "algorithm", "leastconn"),
					resource.TestCheckTypeSetElemAttr(f, "bind.*", "127.0.0.1:18081"),
				),
			},
		},
	})
}

func TestAccHAProxySettingsResource(t *testing.T) {
	a := "opnsense_haproxy_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/haproxy/settings/get", "plugin os-haproxy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_haproxy_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "haproxy_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_haproxy_settings" "test" {
  enabled                  = false
  defaults_max_connections = 2000
  stats_enabled            = true
  stats_users              = ["tfacc:tfacc-secret"]
  logging_facility         = "local1"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "defaults_max_connections", "2000"),
					resource.TestCheckResourceAttr(a, "stats_enabled", "true"),
					resource.TestCheckResourceAttr(a, "stats_users.#", "1"),
					resource.TestCheckResourceAttr(a, "logging_facility", "local1"),
				),
			},
			{
				Config: `
resource "opnsense_haproxy_settings" "test" {
  enabled                  = false
  defaults_max_connections = -1
  stats_enabled            = false
  stats_users              = []
  logging_facility         = "local0"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "defaults_max_connections", "-1"),
					resource.TestCheckResourceAttr(a, "stats_users.#", "0"),
				),
			},
		},
	})
}
