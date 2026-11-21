package ids_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIdsSettingsResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_ids_settings" "test" {}`,
				ResourceName:       "opnsense_ids_settings.test",
				ImportState:        true,
				ImportStateId:      "ids_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_ids_settings" "test" {
  enabled          = false
  mode             = "pcap"
  interfaces       = ["wan"]
  home_networks    = ["10.0.0.0/8", "192.0.2.0/24"]
  detect_profile   = "medium"
  alert_save_logs  = 7
  eve_tls_enabled  = true
  eve_tls_custom   = ["sni", "ja4"]
  syslog_alerts    = true
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_ids_settings.test", "detect_profile", "medium"),
					resource.TestCheckResourceAttr("opnsense_ids_settings.test", "alert_save_logs", "7"),
					resource.TestCheckResourceAttr("opnsense_ids_settings.test", "home_networks.#", "2"),
					resource.TestCheckTypeSetElemAttr("opnsense_ids_settings.test", "eve_tls_custom.*", "ja4"),
					resource.TestCheckResourceAttr("opnsense_ids_settings.test", "syslog_alerts", "true"),
					resource.TestCheckResourceAttrSet("opnsense_ids_settings.test", "update_cron"),
				),
			},
			{
				Config: `
resource "opnsense_ids_settings" "test" {
  enabled          = false
  mode             = "pcap"
  interfaces       = ["wan"]
  home_networks    = ["192.168.0.0/16", "10.0.0.0/8", "172.16.0.0/12"]
  detect_profile   = ""
  alert_save_logs  = 4
  eve_tls_enabled  = false
  eve_tls_custom   = []
  syslog_alerts    = false
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_ids_settings.test", "detect_profile", ""),
					resource.TestCheckResourceAttr("opnsense_ids_settings.test", "home_networks.#", "3"),
					resource.TestCheckResourceAttr("opnsense_ids_settings.test", "eve_tls_custom.#", "0"),
				),
			},
		},
	})
}
