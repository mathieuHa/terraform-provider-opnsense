package nginx_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const nginxStack = `
resource "opnsense_nginx_upstream_server" "a" {
  description = "tfacc-a"
  server      = "127.0.0.1"
  port        = %[2]d
  priority    = 1
}

resource "opnsense_nginx_upstream" "app" {
  description   = "tfacc-app"
  serverentries = [opnsense_nginx_upstream_server.a.id]
}

resource "opnsense_nginx_credential" "user" {
  username = "tfacc"
  password = "tfacc-pass"
}

resource "opnsense_nginx_userlist" "users" {
  name  = "tfacc"
  users = [opnsense_nginx_credential.user.id]
}

resource "opnsense_nginx_ip_acl" "lab" {
  description    = "tfacc-lab"
  default_action = "deny"
  entries = [
    { network = "127.0.0.0/8", action = "allow" },
    { network = "%[3]s", action = "allow" },
  ]
}

resource "opnsense_nginx_errorpage" "down" {
  name        = "tfacc-down"
  statuscodes = ["502", "503"]
  pagecontent = "<h1>tfacc down</h1>"
  response    = "503"
}

resource "opnsense_nginx_security_header" "strict" {
  description          = "tfacc-headers"
  referrer             = "no-referrer"
  content_type_options = true
}

resource "opnsense_nginx_location" "root" {
  description       = "tfacc-root"
  urlpattern        = "/"
  upstream          = opnsense_nginx_upstream.app.id
  authbasic         = "tfacc"
  authbasicuserfile = opnsense_nginx_userlist.users.id
  ip_acl            = opnsense_nginx_ip_acl.lab.id
  errorpages        = [opnsense_nginx_errorpage.down.id]
}

resource "opnsense_nginx_http_server" "web" {
  servername          = ["tfacc.lab.example"]
  listen_http_address = ["18081"]
  locations           = [opnsense_nginx_location.root.id]
  security_header     = opnsense_nginx_security_header.strict.id
}

resource "opnsense_nginx_sni_hostname_map" "tls" {
  description = "tfacc-sni"
  entries = [
    { hostname = "app.tfacc.example", upstream = opnsense_nginx_upstream.app.id },
  ]
}

resource "opnsense_nginx_stream_server" "tls" {
  listen_address   = ["18443"]
  route_field      = "sni_upstream_map"
  sni_upstream_map = opnsense_nginx_sni_hostname_map.tls.id
}
`

func TestAccNginxStack(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/nginx/settings/get", "plugin os-nginx") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(nginxStack, "", 9100, "10.0.0.0/8"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("opnsense_nginx_location.root", "upstream", "opnsense_nginx_upstream.app", "id"),
					resource.TestCheckResourceAttr("opnsense_nginx_http_server.web", "servername.#", "1"),
					resource.TestCheckResourceAttr("opnsense_nginx_http_server.web", "tls_protocols.#", "2"),
					resource.TestCheckResourceAttr("opnsense_nginx_ip_acl.lab", "entries.#", "2"),
					resource.TestCheckResourceAttr("opnsense_nginx_ip_acl.lab", "entries.1.network", "10.0.0.0/8"),
					resource.TestCheckResourceAttr("opnsense_nginx_errorpage.down", "pagecontent", "<h1>tfacc down</h1>"),
					resource.TestCheckResourceAttr("opnsense_nginx_stream_server.tls", "route_field", "sni_upstream_map"),
				),
			},
			{
				ResourceName:      "opnsense_nginx_ip_acl.lab",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "opnsense_nginx_errorpage.down",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "opnsense_nginx_http_server.web",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: fmt.Sprintf(nginxStack, "", 9101, "192.168.0.0/16"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_nginx_upstream_server.a", "port", "9101"),
					resource.TestCheckResourceAttr("opnsense_nginx_ip_acl.lab", "entries.1.network", "192.168.0.0/16"),
				),
			},
		},
	})
}

func TestAccNginxSettingsResource(t *testing.T) {
	a := "opnsense_nginx_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/nginx/settings/get", "plugin os-nginx") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_nginx_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "nginx_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_nginx_settings" "test" {
  enabled                 = false
  ban_ttl                 = 60
  http_worker_connections = 2048
  http_sendfile           = true
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "ban_ttl", "60"),
					resource.TestCheckResourceAttr(a, "http_worker_connections", "2048"),
					resource.TestCheckResourceAttr(a, "http_sendfile", "true"),
				),
			},
			{
				Config: `
resource "opnsense_nginx_settings" "test" {
  enabled                 = false
  ban_ttl                 = 4320
  http_worker_connections = 1024
  http_sendfile           = false
}
`,
				Check: resource.TestCheckResourceAttr(a, "ban_ttl", "4320"),
			},
		},
	})
}
