package interfaces_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInterfacesNeighborResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_interfaces_neighbor\" \"test\" {\n  mac_address = \"02:00:00:00:00:01\"\n  ip_address = \"192.168.150.201\"\n  description = \"tfacc neighbor\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_interfaces_neighbor.test", "mac_address", "02:00:00:00:00:01"),
					resource.TestCheckResourceAttr("opnsense_interfaces_neighbor.test", "ip_address", "192.168.150.201"),
					resource.TestCheckResourceAttr("opnsense_interfaces_neighbor.test", "description", "tfacc neighbor"),
					resource.TestCheckResourceAttrSet("opnsense_interfaces_neighbor.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_interfaces_neighbor.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_interfaces_neighbor\" \"test\" {\n  mac_address = \"02:00:00:00:00:02\"\n  ip_address = \"192.168.150.202\"\n  description = \"tfacc neighbor 2\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_interfaces_neighbor.test", "mac_address", "02:00:00:00:00:02"),
					resource.TestCheckResourceAttr("opnsense_interfaces_neighbor.test", "ip_address", "192.168.150.202"),
					resource.TestCheckResourceAttr("opnsense_interfaces_neighbor.test", "description", "tfacc neighbor 2"),
				),
			},
			{
				Config: "\nresource \"opnsense_interfaces_neighbor\" \"test\" {\n  mac_address = \"02:00:00:00:00:02\"\n  ip_address = \"192.168.150.202\"\n  description = \"tfacc neighbor 2\"\n}\n\ndata \"opnsense_interfaces_neighbor\" \"test\" {\n  id = opnsense_interfaces_neighbor.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_neighbor.test", "mac_address", "opnsense_interfaces_neighbor.test", "mac_address"),
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_neighbor.test", "ip_address", "opnsense_interfaces_neighbor.test", "ip_address"),
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_neighbor.test", "description", "opnsense_interfaces_neighbor.test", "description"),
				),
			},
		},
	})
}
