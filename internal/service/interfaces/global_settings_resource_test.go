package interfaces_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInterfacesSettingsResource(t *testing.T) {
	a := "opnsense_interfaces_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_interfaces_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "interfaces_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_interfaces_settings" "test" {
  vlan_hw_filter   = "2"
  dhcp6_no_release = true
  dhcp6_ra_timeout = 15
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "vlan_hw_filter", "2"),
					resource.TestCheckResourceAttr(a, "dhcp6_no_release", "true"),
					resource.TestCheckResourceAttr(a, "dhcp6_ra_timeout", "15"),
					resource.TestCheckResourceAttr(a, "disable_checksum_offloading", "true"),
				),
			},
			{
				Config: `
resource "opnsense_interfaces_settings" "test" {
  vlan_hw_filter   = "1"
  dhcp6_no_release = false
  dhcp6_ra_timeout = 10
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "vlan_hw_filter", "1"),
					resource.TestCheckResourceAttr(a, "dhcp6_ra_timeout", "10"),
				),
			},
		},
	})
}
