package tailscale_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTailscaleSettingsResource(t *testing.T) {
	a := "opnsense_tailscale_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/tailscale/settings/get", "plugin os-tailscale") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_tailscale_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "tailscale_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_tailscale_settings" "test" {
  enabled              = false
  listen_port          = 41642
  accept_dns           = false
  advertise_exit_node  = true
  accept_subnet_routes = true
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "listen_port", "41642"),
					resource.TestCheckResourceAttr(a, "accept_dns", "false"),
					resource.TestCheckResourceAttr(a, "advertise_exit_node", "true"),
					resource.TestCheckResourceAttr(a, "exit_node", ""),
				),
			},
			{
				Config: `
resource "opnsense_tailscale_settings" "test" {
  enabled              = false
  listen_port          = 41641
  accept_dns           = true
  advertise_exit_node  = false
  accept_subnet_routes = false
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "listen_port", "41641"),
					resource.TestCheckResourceAttr(a, "accept_dns", "true"),
				),
			},
		},
	})
}

func TestAccTailscaleAuthenticationResource(t *testing.T) {
	a := "opnsense_tailscale_authentication.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/tailscale/authentication/get", "plugin os-tailscale") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_tailscale_authentication" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "tailscale_authentication",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_tailscale_authentication" "test" {
  login_server = "https://headscale.example.com"
  pre_auth_key = "tskey-auth-tfacc"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "login_server", "https://headscale.example.com"),
					resource.TestCheckResourceAttr(a, "pre_auth_key", "tskey-auth-tfacc"),
				),
			},
			{
				Config: `
resource "opnsense_tailscale_authentication" "test" {
  login_server = "https://controlplane.tailscale.com"
  pre_auth_key = ""
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "login_server", "https://controlplane.tailscale.com"),
					resource.TestCheckResourceAttr(a, "pre_auth_key", ""),
				),
			},
		},
	})
}
