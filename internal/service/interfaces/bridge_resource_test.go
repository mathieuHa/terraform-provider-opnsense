package interfaces_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInterfacesBridgeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.InterfacePreCheck(t, "opt1") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_interfaces_bridge\" \"test\" {\n  members = [\"opt1\"]\n  description = \"tf-test-bridge\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_interfaces_bridge.test", "members.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_interfaces_bridge.test", "members.*", "opt1"),
					resource.TestCheckResourceAttr("opnsense_interfaces_bridge.test", "description", "tf-test-bridge"),
					resource.TestCheckResourceAttrSet("opnsense_interfaces_bridge.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_interfaces_bridge.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_interfaces_bridge\" \"test\" {\n  members = [\"opt1\"]\n  enable_stp = true\n  stp_interfaces = [\"opt1\"]\n  max_addresses = 500\n  description = \"tf-test-bridge-2\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_interfaces_bridge.test", "members.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_interfaces_bridge.test", "members.*", "opt1"),
					resource.TestCheckResourceAttr("opnsense_interfaces_bridge.test", "enable_stp", "true"),
					resource.TestCheckResourceAttr("opnsense_interfaces_bridge.test", "stp_interfaces.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_interfaces_bridge.test", "stp_interfaces.*", "opt1"),
					resource.TestCheckResourceAttr("opnsense_interfaces_bridge.test", "max_addresses", "500"),
					resource.TestCheckResourceAttr("opnsense_interfaces_bridge.test", "description", "tf-test-bridge-2"),
				),
			},
			{
				Config: "\nresource \"opnsense_interfaces_bridge\" \"test\" {\n  members = [\"opt1\"]\n  enable_stp = true\n  stp_interfaces = [\"opt1\"]\n  max_addresses = 500\n  description = \"tf-test-bridge-2\"\n}\n\ndata \"opnsense_interfaces_bridge\" \"test\" {\n  id = opnsense_interfaces_bridge.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_bridge.test", "enable_stp", "opnsense_interfaces_bridge.test", "enable_stp"),
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_bridge.test", "max_addresses", "opnsense_interfaces_bridge.test", "max_addresses"),
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_bridge.test", "description", "opnsense_interfaces_bridge.test", "description"),
				),
			},
		},
	})
}
