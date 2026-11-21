package dhcrelay_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDhcrelayRelayResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "resource \"opnsense_dhcrelay_destination\" \"d\" {\n  name   = \"tfacc-relay\"\n  server = [\"192.0.2.67\"]\n}\n\nresource \"opnsense_dhcrelay_relay\" \"test\" {\n  interface = \"wan\"\n  destination = opnsense_dhcrelay_destination.d.id\n  enabled = false\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dhcrelay_relay.test", "interface", "wan"),
					resource.TestCheckResourceAttr("opnsense_dhcrelay_relay.test", "enabled", "false"),
					resource.TestCheckResourceAttrSet("opnsense_dhcrelay_relay.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_dhcrelay_relay.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "resource \"opnsense_dhcrelay_destination\" \"d\" {\n  name   = \"tfacc-relay\"\n  server = [\"192.0.2.67\"]\n}\n\nresource \"opnsense_dhcrelay_relay\" \"test\" {\n  interface = \"wan\"\n  destination = opnsense_dhcrelay_destination.d.id\n  enabled = false\n  agent_info = true\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dhcrelay_relay.test", "interface", "wan"),
					resource.TestCheckResourceAttr("opnsense_dhcrelay_relay.test", "enabled", "false"),
					resource.TestCheckResourceAttr("opnsense_dhcrelay_relay.test", "agent_info", "true"),
				),
			},
			{
				Config: "resource \"opnsense_dhcrelay_destination\" \"d\" {\n  name   = \"tfacc-relay\"\n  server = [\"192.0.2.67\"]\n}\n\nresource \"opnsense_dhcrelay_relay\" \"test\" {\n  interface = \"wan\"\n  destination = opnsense_dhcrelay_destination.d.id\n  enabled = false\n  agent_info = true\n}\n\ndata \"opnsense_dhcrelay_relay\" \"test\" {\n  id = opnsense_dhcrelay_relay.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_dhcrelay_relay.test", "interface", "opnsense_dhcrelay_relay.test", "interface"),
					resource.TestCheckResourceAttrPair("data.opnsense_dhcrelay_relay.test", "enabled", "opnsense_dhcrelay_relay.test", "enabled"),
					resource.TestCheckResourceAttrPair("data.opnsense_dhcrelay_relay.test", "agent_info", "opnsense_dhcrelay_relay.test", "agent_info"),
				),
			},
		},
	})
}
