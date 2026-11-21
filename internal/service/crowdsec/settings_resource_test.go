package crowdsec_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCrowdSecSettingsResource(t *testing.T) {
	a := "opnsense_crowdsec_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/crowdsec/general/get", "plugin os-crowdsec") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_crowdsec_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "crowdsec_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_crowdsec_settings" "test" {
  agent_enabled           = false
  firewall_bouncer_enabled = false
  lapi_listen_port        = 8081
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "lapi_listen_port", "8081"),
				),
			},
			{
				Config: `
resource "opnsense_crowdsec_settings" "test" {
  agent_enabled           = false
  firewall_bouncer_enabled = false
  lapi_listen_port        = 8080
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "lapi_listen_port", "8080"),
				),
			},
		},
	})
}
