package firewall_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccFirewallGroupResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallGroupConfig(`["wan"]`, "first"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "name", "tfgroup"),
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "members.#", "1"),
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "no_group", "false"),
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "sequence", "0"),
					resource.TestCheckResourceAttrSet("opnsense_firewall_group.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_firewall_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccFirewallGroupConfig(`["wan", "lo0"]`, "second"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "members.#", "2"),
					resource.TestCheckTypeSetElemAttr("opnsense_firewall_group.test", "members.*", "lo0"),
					resource.TestCheckResourceAttr("opnsense_firewall_group.test", "description", "second"),
				),
			},
			{
				Config: testAccFirewallGroupConfig(`["wan", "lo0"]`, "second") + `
data "opnsense_firewall_group" "test" {
  id = opnsense_firewall_group.test.id
}
`,
				Check: resource.TestCheckResourceAttr("data.opnsense_firewall_group.test", "members.#", "2"),
			},
		},
	})
}

func TestAccFirewallNPTResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallNPTConfig("2001:db8:1::/64", false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_firewall_npt.test", "interface", "wan"),
					resource.TestCheckResourceAttr("opnsense_firewall_npt.test", "source_net", "fd00:1::/64"),
					resource.TestCheckResourceAttr("opnsense_firewall_npt.test", "destination_net", "2001:db8:1::/64"),
					resource.TestCheckResourceAttr("opnsense_firewall_npt.test", "enabled", "true"),
					resource.TestCheckResourceAttr("opnsense_firewall_npt.test", "log", "false"),
					resource.TestCheckNoResourceAttr("opnsense_firewall_npt.test", "track_interface"),
					resource.TestCheckResourceAttr("opnsense_firewall_npt.test", "categories.#", "0"),
				),
			},
			{
				ResourceName:      "opnsense_firewall_npt.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccFirewallNPTConfig("2001:db8:2::/64", true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_firewall_npt.test", "destination_net", "2001:db8:2::/64"),
					resource.TestCheckResourceAttr("opnsense_firewall_npt.test", "log", "true"),
				),
			},
			{
				Config: testAccFirewallNPTConfig("2001:db8:2::/64", true) + `
data "opnsense_firewall_npt" "test" {
  id = opnsense_firewall_npt.test.id
}
`,
				Check: resource.TestCheckResourceAttr("data.opnsense_firewall_npt.test", "source_net", "fd00:1::/64"),
			},
		},
	})
}

func testAccFirewallGroupConfig(members, description string) string {
	return fmt.Sprintf(`
resource "opnsense_firewall_group" "test" {
  name        = "tfgroup"
  members     = %s
  description = %q
}
`, members, description)
}

func testAccFirewallNPTConfig(destination string, log bool) string {
	return fmt.Sprintf(`
resource "opnsense_firewall_npt" "test" {
  interface       = "wan"
  sequence        = 100
  source_net      = "fd00:1::/64"
  destination_net = %q
  log             = %t
  description     = "tf-test-npt"
}
`, destination, log)
}
