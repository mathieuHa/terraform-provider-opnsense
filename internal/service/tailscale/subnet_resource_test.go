package tailscale_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTailscaleSubnetResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/tailscale/settings/get", "plugin os-tailscale") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_tailscale_subnet\" \"test\" {\n  subnet = \"10.99.10.0/24\"\n  description = \"tfacc subnet\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_tailscale_subnet.test", "subnet", "10.99.10.0/24"),
					resource.TestCheckResourceAttr("opnsense_tailscale_subnet.test", "description", "tfacc subnet"),
					resource.TestCheckResourceAttrSet("opnsense_tailscale_subnet.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_tailscale_subnet.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_tailscale_subnet\" \"test\" {\n  subnet = \"10.99.11.0/24\"\n  description = \"tfacc subnet 2\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_tailscale_subnet.test", "subnet", "10.99.11.0/24"),
					resource.TestCheckResourceAttr("opnsense_tailscale_subnet.test", "description", "tfacc subnet 2"),
				),
			},
			{
				Config: "\nresource \"opnsense_tailscale_subnet\" \"test\" {\n  subnet = \"10.99.11.0/24\"\n  description = \"tfacc subnet 2\"\n}\n\ndata \"opnsense_tailscale_subnet\" \"test\" {\n  id = opnsense_tailscale_subnet.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_tailscale_subnet.test", "subnet", "opnsense_tailscale_subnet.test", "subnet"),
					resource.TestCheckResourceAttrPair("data.opnsense_tailscale_subnet.test", "description", "opnsense_tailscale_subnet.test", "description"),
				),
			},
		},
	})
}
