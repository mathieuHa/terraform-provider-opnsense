package interfaces_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInterfacesLaggResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.OptionPreCheck(t, "/interfaces/lagg_settings/get_item", "lagg", "members", "vtnet2", "vtnet3")
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_interfaces_lagg\" \"test\" {\n  members = [\"vtnet2\"]\n  proto = \"failover\"\n  description = \"tf-acc lagg\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_interfaces_lagg.test", "members.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_interfaces_lagg.test", "members.*", "vtnet2"),
					resource.TestCheckResourceAttr("opnsense_interfaces_lagg.test", "proto", "failover"),
					resource.TestCheckResourceAttr("opnsense_interfaces_lagg.test", "description", "tf-acc lagg"),
					resource.TestCheckResourceAttrSet("opnsense_interfaces_lagg.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_interfaces_lagg.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_interfaces_lagg\" \"test\" {\n  members = [\"vtnet2\", \"vtnet3\"]\n  proto = \"loadbalance\"\n  hash = [\"l3\", \"l4\"]\n  mtu = 1400\n  description = \"tf-acc lagg 2\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_interfaces_lagg.test", "members.#", "2"),
					resource.TestCheckTypeSetElemAttr("opnsense_interfaces_lagg.test", "members.*", "vtnet2"),
					resource.TestCheckTypeSetElemAttr("opnsense_interfaces_lagg.test", "members.*", "vtnet3"),
					resource.TestCheckResourceAttr("opnsense_interfaces_lagg.test", "proto", "loadbalance"),
					resource.TestCheckResourceAttr("opnsense_interfaces_lagg.test", "hash.#", "2"),
					resource.TestCheckTypeSetElemAttr("opnsense_interfaces_lagg.test", "hash.*", "l3"),
					resource.TestCheckTypeSetElemAttr("opnsense_interfaces_lagg.test", "hash.*", "l4"),
					resource.TestCheckResourceAttr("opnsense_interfaces_lagg.test", "mtu", "1400"),
					resource.TestCheckResourceAttr("opnsense_interfaces_lagg.test", "description", "tf-acc lagg 2"),
				),
			},
			{
				Config: "\nresource \"opnsense_interfaces_lagg\" \"test\" {\n  members = [\"vtnet2\", \"vtnet3\"]\n  proto = \"loadbalance\"\n  hash = [\"l3\", \"l4\"]\n  mtu = 1400\n  description = \"tf-acc lagg 2\"\n}\n\ndata \"opnsense_interfaces_lagg\" \"test\" {\n  id = opnsense_interfaces_lagg.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_lagg.test", "proto", "opnsense_interfaces_lagg.test", "proto"),
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_lagg.test", "mtu", "opnsense_interfaces_lagg.test", "mtu"),
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_lagg.test", "description", "opnsense_interfaces_lagg.test", "description"),
				),
			},
		},
	})
}
