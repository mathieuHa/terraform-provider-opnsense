package captiveportal_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCaptiveportalZoneResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.OptionPreCheck(t, "/captiveportal/settings/get_zone", "zone", "interfaces", "opt1") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_captiveportal_zone\" \"test\" {\n  interfaces = [\"opt1\"]\n  description = \"tfacc zone\"\n  auth_servers = [\"Local Database\"]\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_captiveportal_zone.test", "interfaces.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_captiveportal_zone.test", "interfaces.*", "opt1"),
					resource.TestCheckResourceAttr("opnsense_captiveportal_zone.test", "description", "tfacc zone"),
					resource.TestCheckResourceAttr("opnsense_captiveportal_zone.test", "auth_servers.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_captiveportal_zone.test", "auth_servers.*", "Local Database"),
					resource.TestCheckResourceAttrSet("opnsense_captiveportal_zone.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_captiveportal_zone.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_captiveportal_zone\" \"test\" {\n  interfaces = [\"opt1\"]\n  description = \"tfacc zone 2\"\n  idle_timeout = 30\n  hard_timeout = 120\n  allowed_addresses = [\"10.99.2.50\"]\n  allowed_mac_addresses = [\"02:00:00:00:00:01\"]\n  roaming = false\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_captiveportal_zone.test", "interfaces.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_captiveportal_zone.test", "interfaces.*", "opt1"),
					resource.TestCheckResourceAttr("opnsense_captiveportal_zone.test", "description", "tfacc zone 2"),
					resource.TestCheckResourceAttr("opnsense_captiveportal_zone.test", "idle_timeout", "30"),
					resource.TestCheckResourceAttr("opnsense_captiveportal_zone.test", "hard_timeout", "120"),
					resource.TestCheckResourceAttr("opnsense_captiveportal_zone.test", "allowed_addresses.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_captiveportal_zone.test", "allowed_addresses.*", "10.99.2.50"),
					resource.TestCheckResourceAttr("opnsense_captiveportal_zone.test", "allowed_mac_addresses.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_captiveportal_zone.test", "allowed_mac_addresses.*", "02:00:00:00:00:01"),
					resource.TestCheckResourceAttr("opnsense_captiveportal_zone.test", "roaming", "false"),
				),
			},
			{
				Config: "\nresource \"opnsense_captiveportal_zone\" \"test\" {\n  interfaces = [\"opt1\"]\n  description = \"tfacc zone 2\"\n  idle_timeout = 30\n  hard_timeout = 120\n  allowed_addresses = [\"10.99.2.50\"]\n  allowed_mac_addresses = [\"02:00:00:00:00:01\"]\n  roaming = false\n}\n\ndata \"opnsense_captiveportal_zone\" \"test\" {\n  id = opnsense_captiveportal_zone.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_captiveportal_zone.test", "description", "opnsense_captiveportal_zone.test", "description"),
					resource.TestCheckResourceAttrPair("data.opnsense_captiveportal_zone.test", "idle_timeout", "opnsense_captiveportal_zone.test", "idle_timeout"),
					resource.TestCheckResourceAttrPair("data.opnsense_captiveportal_zone.test", "hard_timeout", "opnsense_captiveportal_zone.test", "hard_timeout"),
					resource.TestCheckResourceAttrPair("data.opnsense_captiveportal_zone.test", "roaming", "opnsense_captiveportal_zone.test", "roaming"),
				),
			},
		},
	})
}
