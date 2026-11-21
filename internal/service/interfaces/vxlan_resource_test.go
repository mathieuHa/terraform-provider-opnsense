package interfaces_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInterfacesVxlanResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_interfaces_vxlan\" \"test\" {\n  vni = 4242\n  local = \"192.0.2.50\"\n  remote = \"192.0.2.51\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_interfaces_vxlan.test", "vni", "4242"),
					resource.TestCheckResourceAttr("opnsense_interfaces_vxlan.test", "local", "192.0.2.50"),
					resource.TestCheckResourceAttr("opnsense_interfaces_vxlan.test", "remote", "192.0.2.51"),
					resource.TestCheckResourceAttrSet("opnsense_interfaces_vxlan.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_interfaces_vxlan.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_interfaces_vxlan\" \"test\" {\n  vni = 4243\n  local = \"192.0.2.50\"\n  remote = \"192.0.2.52\"\n  remote_port = 4789\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_interfaces_vxlan.test", "vni", "4243"),
					resource.TestCheckResourceAttr("opnsense_interfaces_vxlan.test", "local", "192.0.2.50"),
					resource.TestCheckResourceAttr("opnsense_interfaces_vxlan.test", "remote", "192.0.2.52"),
					resource.TestCheckResourceAttr("opnsense_interfaces_vxlan.test", "remote_port", "4789"),
				),
			},
			{
				Config: "\nresource \"opnsense_interfaces_vxlan\" \"test\" {\n  vni = 4243\n  local = \"192.0.2.50\"\n  remote = \"192.0.2.52\"\n  remote_port = 4789\n}\n\ndata \"opnsense_interfaces_vxlan\" \"test\" {\n  id = opnsense_interfaces_vxlan.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_vxlan.test", "vni", "opnsense_interfaces_vxlan.test", "vni"),
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_vxlan.test", "local", "opnsense_interfaces_vxlan.test", "local"),
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_vxlan.test", "remote", "opnsense_interfaces_vxlan.test", "remote"),
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_vxlan.test", "remote_port", "opnsense_interfaces_vxlan.test", "remote_port"),
				),
			},
		},
	})
}
