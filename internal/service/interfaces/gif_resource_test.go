package interfaces_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInterfacesGifResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_interfaces_gif\" \"test\" {\n  local_address = \"wan\"\n  remote_address = \"192.0.2.1\"\n  tunnel_local_address = \"10.94.0.1\"\n  tunnel_remote_address = \"10.94.0.2\"\n  description = \"tf-acc gif\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_interfaces_gif.test", "local_address", "wan"),
					resource.TestCheckResourceAttr("opnsense_interfaces_gif.test", "remote_address", "192.0.2.1"),
					resource.TestCheckResourceAttr("opnsense_interfaces_gif.test", "tunnel_local_address", "10.94.0.1"),
					resource.TestCheckResourceAttr("opnsense_interfaces_gif.test", "tunnel_remote_address", "10.94.0.2"),
					resource.TestCheckResourceAttr("opnsense_interfaces_gif.test", "description", "tf-acc gif"),
					resource.TestCheckResourceAttrSet("opnsense_interfaces_gif.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_interfaces_gif.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_interfaces_gif\" \"test\" {\n  local_address = \"wan\"\n  remote_address = \"192.0.2.2\"\n  tunnel_local_address = \"10.94.1.1\"\n  tunnel_remote_address = \"10.94.1.2\"\n  tunnel_remote_net = 30\n  description = \"tf-acc gif 2\"\n  ecn_friendly = true\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_interfaces_gif.test", "local_address", "wan"),
					resource.TestCheckResourceAttr("opnsense_interfaces_gif.test", "remote_address", "192.0.2.2"),
					resource.TestCheckResourceAttr("opnsense_interfaces_gif.test", "tunnel_local_address", "10.94.1.1"),
					resource.TestCheckResourceAttr("opnsense_interfaces_gif.test", "tunnel_remote_address", "10.94.1.2"),
					resource.TestCheckResourceAttr("opnsense_interfaces_gif.test", "tunnel_remote_net", "30"),
					resource.TestCheckResourceAttr("opnsense_interfaces_gif.test", "description", "tf-acc gif 2"),
					resource.TestCheckResourceAttr("opnsense_interfaces_gif.test", "ecn_friendly", "true"),
				),
			},
			{
				Config: "\nresource \"opnsense_interfaces_gif\" \"test\" {\n  local_address = \"wan\"\n  remote_address = \"192.0.2.2\"\n  tunnel_local_address = \"10.94.1.1\"\n  tunnel_remote_address = \"10.94.1.2\"\n  tunnel_remote_net = 30\n  description = \"tf-acc gif 2\"\n  ecn_friendly = true\n}\n\ndata \"opnsense_interfaces_gif\" \"test\" {\n  id = opnsense_interfaces_gif.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_gif.test", "local_address", "opnsense_interfaces_gif.test", "local_address"),
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_gif.test", "remote_address", "opnsense_interfaces_gif.test", "remote_address"),
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_gif.test", "tunnel_local_address", "opnsense_interfaces_gif.test", "tunnel_local_address"),
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_gif.test", "tunnel_remote_address", "opnsense_interfaces_gif.test", "tunnel_remote_address"),
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_gif.test", "tunnel_remote_net", "opnsense_interfaces_gif.test", "tunnel_remote_net"),
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_gif.test", "description", "opnsense_interfaces_gif.test", "description"),
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_gif.test", "ecn_friendly", "opnsense_interfaces_gif.test", "ecn_friendly"),
				),
			},
		},
	})
}
