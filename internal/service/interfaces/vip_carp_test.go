package interfaces_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInterfacesVipResource_Carp(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVipCarpConfig(0, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_interfaces_vip.test", "mode", "carp"),
					resource.TestCheckResourceAttr("opnsense_interfaces_vip.test", "vhid", "77"),
					resource.TestCheckResourceAttr("opnsense_interfaces_vip.test", "advbase", "1"),
					resource.TestCheckResourceAttr("opnsense_interfaces_vip.test", "advskew", "0"),
					resource.TestCheckResourceAttr("opnsense_interfaces_vip.test", "no_sync", "false"),
					resource.TestCheckResourceAttrSet("opnsense_interfaces_vip.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_interfaces_vip.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccVipCarpConfig(100, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_interfaces_vip.test", "advskew", "100"),
					resource.TestCheckResourceAttr("opnsense_interfaces_vip.test", "no_sync", "true"),
				),
			},
		},
	})
}

func testAccVipCarpConfig(skew int, noSync bool) string {
	return fmt.Sprintf(`
resource "opnsense_interfaces_vip" "test" {
  mode        = "carp"
  interface   = "wan"
  network     = "203.0.113.77/24"
  vhid        = 77
  advskew     = %d
  password    = "tf-acc-carp-secret"
  no_sync     = %t
  description = "tf-acc carp vip"
}
`, skew, noSync)
}
