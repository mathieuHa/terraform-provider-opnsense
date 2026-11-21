package interfaces_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInterfacesLoopbackResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_interfaces_loopback\" \"test\" {\n  description = \"tf-acc loopback\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_interfaces_loopback.test", "description", "tf-acc loopback"),
					resource.TestCheckResourceAttrSet("opnsense_interfaces_loopback.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_interfaces_loopback.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_interfaces_loopback\" \"test\" {\n  description = \"tf-acc loopback 2\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_interfaces_loopback.test", "description", "tf-acc loopback 2"),
				),
			},
			{
				Config: "\nresource \"opnsense_interfaces_loopback\" \"test\" {\n  description = \"tf-acc loopback 2\"\n}\n\ndata \"opnsense_interfaces_loopback\" \"test\" {\n  id = opnsense_interfaces_loopback.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_interfaces_loopback.test", "description", "opnsense_interfaces_loopback.test", "description"),
				),
			},
		},
	})
}
