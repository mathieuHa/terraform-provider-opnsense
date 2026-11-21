package radvd_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRadvdEntryResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_radvd_entry\" \"test\" {\n  interface = \"wan\"\n  mode = \"router\"\n  rdnss = [\"2001:db8::53\"]\n  dnssl = [\"example.com\"]\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_radvd_entry.test", "interface", "wan"),
					resource.TestCheckResourceAttr("opnsense_radvd_entry.test", "mode", "router"),
					resource.TestCheckResourceAttr("opnsense_radvd_entry.test", "rdnss.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_radvd_entry.test", "rdnss.*", "2001:db8::53"),
					resource.TestCheckResourceAttr("opnsense_radvd_entry.test", "dnssl.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_radvd_entry.test", "dnssl.*", "example.com"),
					resource.TestCheckResourceAttrSet("opnsense_radvd_entry.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_radvd_entry.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_radvd_entry\" \"test\" {\n  interface = \"wan\"\n  mode = \"stateless\"\n  enabled = false\n  min_interval = 30\n  max_interval = 120\n  default_preference = \"high\"\n  routes = [\"2001:db8:1::/48\"]\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_radvd_entry.test", "interface", "wan"),
					resource.TestCheckResourceAttr("opnsense_radvd_entry.test", "mode", "stateless"),
					resource.TestCheckResourceAttr("opnsense_radvd_entry.test", "enabled", "false"),
					resource.TestCheckResourceAttr("opnsense_radvd_entry.test", "min_interval", "30"),
					resource.TestCheckResourceAttr("opnsense_radvd_entry.test", "max_interval", "120"),
					resource.TestCheckResourceAttr("opnsense_radvd_entry.test", "default_preference", "high"),
					resource.TestCheckResourceAttr("opnsense_radvd_entry.test", "routes.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_radvd_entry.test", "routes.*", "2001:db8:1::/48"),
				),
			},
			{
				Config: "\nresource \"opnsense_radvd_entry\" \"test\" {\n  interface = \"wan\"\n  mode = \"stateless\"\n  enabled = false\n  min_interval = 30\n  max_interval = 120\n  default_preference = \"high\"\n  routes = [\"2001:db8:1::/48\"]\n}\n\ndata \"opnsense_radvd_entry\" \"test\" {\n  id = opnsense_radvd_entry.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_radvd_entry.test", "interface", "opnsense_radvd_entry.test", "interface"),
					resource.TestCheckResourceAttrPair("data.opnsense_radvd_entry.test", "mode", "opnsense_radvd_entry.test", "mode"),
					resource.TestCheckResourceAttrPair("data.opnsense_radvd_entry.test", "enabled", "opnsense_radvd_entry.test", "enabled"),
					resource.TestCheckResourceAttrPair("data.opnsense_radvd_entry.test", "min_interval", "opnsense_radvd_entry.test", "min_interval"),
					resource.TestCheckResourceAttrPair("data.opnsense_radvd_entry.test", "max_interval", "opnsense_radvd_entry.test", "max_interval"),
					resource.TestCheckResourceAttrPair("data.opnsense_radvd_entry.test", "default_preference", "opnsense_radvd_entry.test", "default_preference"),
				),
			},
		},
	})
}
