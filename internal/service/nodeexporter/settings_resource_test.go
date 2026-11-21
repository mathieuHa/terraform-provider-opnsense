package nodeexporter_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNodeExporterSettingsResource(t *testing.T) {
	a := "opnsense_nodeexporter_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/nodeexporter/general/get", "plugin os-node_exporter") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_nodeexporter_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "nodeexporter_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_nodeexporter_settings" "test" {
  enabled        = true
  listen_address = "0.0.0.0"
  listen_port    = 9101
  cpu         = true
  zfs         = true
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "enabled", "true"),
					resource.TestCheckResourceAttr(a, "listen_port", "9101"),
					resource.TestCheckResourceAttr(a, "cpu", "true"),
					resource.TestCheckResourceAttr(a, "listen_address", "0.0.0.0"),
				),
			},
			{
				Config: `
resource "opnsense_nodeexporter_settings" "test" {
  enabled        = false
  listen_address = "0.0.0.0"
  listen_port    = 9100
  cpu         = false
  zfs         = false
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "enabled", "false"),
					resource.TestCheckResourceAttr(a, "listen_port", "9100"),
				),
			},
		},
	})
}
