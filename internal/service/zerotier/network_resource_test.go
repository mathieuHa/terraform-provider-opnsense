package zerotier_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func zerotierConfig(service bool, network string) string {
	return fmt.Sprintf(`
import {
  to = opnsense_zerotier_settings.test
  id = "zerotier_settings"
}

resource "opnsense_zerotier_settings" "test" {
  enabled = %t
}
%s`, service, network)
}

func TestAccZerotierNetworkResource(t *testing.T) {
	n := "opnsense_zerotier_network.test"
	network := func(desc string, enabled bool) string {
		return fmt.Sprintf(`
resource "opnsense_zerotier_network" "test" {
  network_id  = "8056c2e21c00beef"
  description = %q
  enabled     = %t
  depends_on  = [opnsense_zerotier_settings.test]
}
`, desc, enabled)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/zerotier/settings/get", "plugin os-zerotier") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: zerotierConfig(true, network("tfacc", false)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr(n, "id", acctest.UUIDRegexp()),
					resource.TestCheckResourceAttr(n, "enabled", "false"),
				),
			},
			{
				ResourceName:      n,
				ImportState:       true,
				ImportStateVerify: true,
				Config:            zerotierConfig(true, network("tfacc", false)),
			},
			{
				Config: zerotierConfig(true, network("tfacc2", true)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(n, "description", "tfacc2"),
					resource.TestCheckResourceAttr(n, "enabled", "true"),
				),
			},
			{
				Config: zerotierConfig(true, network("tfacc2", false)),
				Check:  resource.TestCheckResourceAttr(n, "enabled", "false"),
			},
			{
				Config: zerotierConfig(true, ""),
			},
			{
				Config: zerotierConfig(false, ""),
				Check:  resource.TestCheckResourceAttr("opnsense_zerotier_settings.test", "enabled", "false"),
			},
		},
	})
}
