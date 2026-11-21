package proxy_test

import (
	"regexp"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccProxySettingsResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.EndpointPreCheck(t, "/proxy/settings/get", "requires the os-squid plugin")
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_proxy_settings" "test" {}`,
				ResourceName:       "opnsense_proxy_settings.test",
				ImportState:        true,
				ImportStateId:      "proxy_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_proxy_settings" "test" {
  enabled          = true
  interfaces       = ["lo0"]
  port             = 3128
  visible_hostname = "tf-acc-proxy"
  cache_size       = 200
  acl_whitelist    = ["\\.example\\.com"]
  acl_ssl_ports    = ["443:https", "8443:alt-https"]
  icap_preview_size = 2048
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_proxy_settings.test", "id", "proxy_settings"),
					resource.TestCheckResourceAttr("opnsense_proxy_settings.test", "enabled", "true"),
					resource.TestCheckTypeSetElemAttr("opnsense_proxy_settings.test", "interfaces.*", "lo0"),
					resource.TestCheckResourceAttr("opnsense_proxy_settings.test", "visible_hostname", "tf-acc-proxy"),
					resource.TestCheckResourceAttr("opnsense_proxy_settings.test", "cache_size", "200"),
					resource.TestCheckTypeSetElemAttr("opnsense_proxy_settings.test", "acl_whitelist.*", "\\.example\\.com"),
					resource.TestCheckResourceAttr("opnsense_proxy_settings.test", "acl_ssl_ports.#", "2"),
					resource.TestCheckResourceAttr("opnsense_proxy_settings.test", "icap_preview_size", "2048"),
				),
			},
			{
				Config: `
resource "opnsense_proxy_settings" "test" {
  enabled          = false
  interfaces       = []
  visible_hostname = ""
  cache_size       = 100
  acl_whitelist    = []
  acl_ssl_ports    = ["443:https"]
  icap_preview_size = 1024
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_proxy_settings.test", "enabled", "false"),
					resource.TestCheckResourceAttr("opnsense_proxy_settings.test", "acl_whitelist.#", "0"),
					resource.TestCheckResourceAttr("opnsense_proxy_settings.test", "acl_ssl_ports.#", "1"),
				),
			},
		},
	})
}

func TestAccProxySettingsResource_CreateBlocked(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/proxy/settings/get", "requires the os-squid plugin") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      `resource "opnsense_proxy_settings" "test" {}`,
				ExpectError: regexp.MustCompile("Cannot Create Singleton Resource"),
			},
		},
	})
}
