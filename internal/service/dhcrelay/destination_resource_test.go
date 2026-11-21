package dhcrelay_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDhcrelayDestinationResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_dhcrelay_destination\" \"test\" {\n  name = \"tfacc\"\n  server = [\"192.0.2.67\"]\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dhcrelay_destination.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_dhcrelay_destination.test", "server.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_dhcrelay_destination.test", "server.*", "192.0.2.67"),
					resource.TestCheckResourceAttrSet("opnsense_dhcrelay_destination.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_dhcrelay_destination.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_dhcrelay_destination\" \"test\" {\n  name = \"tfacc2\"\n  server = [\"192.0.2.67\", \"192.0.2.68\"]\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dhcrelay_destination.test", "name", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_dhcrelay_destination.test", "server.#", "2"),
					resource.TestCheckTypeSetElemAttr("opnsense_dhcrelay_destination.test", "server.*", "192.0.2.67"),
					resource.TestCheckTypeSetElemAttr("opnsense_dhcrelay_destination.test", "server.*", "192.0.2.68"),
				),
			},
			{
				Config: "\nresource \"opnsense_dhcrelay_destination\" \"test\" {\n  name = \"tfacc2\"\n  server = [\"192.0.2.67\", \"192.0.2.68\"]\n}\n\ndata \"opnsense_dhcrelay_destination\" \"test\" {\n  id = opnsense_dhcrelay_destination.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_dhcrelay_destination.test", "name", "opnsense_dhcrelay_destination.test", "name"),
				),
			},
		},
	})
}
