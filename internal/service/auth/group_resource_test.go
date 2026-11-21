package auth_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAuthGroupResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_auth_group\" \"test\" {\n  name = \"tftestgroup\"\n  description = \"tf-test-group\"\n  privileges = [\"page-diagnostics-ping\"]\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_auth_group.test", "name", "tftestgroup"),
					resource.TestCheckResourceAttr("opnsense_auth_group.test", "description", "tf-test-group"),
					resource.TestCheckResourceAttr("opnsense_auth_group.test", "privileges.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_auth_group.test", "privileges.*", "page-diagnostics-ping"),
					resource.TestCheckResourceAttrSet("opnsense_auth_group.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_auth_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_auth_group\" \"test\" {\n  name = \"tftestgroup\"\n  description = \"tf-test-group-2\"\n  privileges = [\"page-diagnostics-ping\", \"page-diagnostics-traceroute\"]\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_auth_group.test", "name", "tftestgroup"),
					resource.TestCheckResourceAttr("opnsense_auth_group.test", "description", "tf-test-group-2"),
					resource.TestCheckResourceAttr("opnsense_auth_group.test", "privileges.#", "2"),
					resource.TestCheckTypeSetElemAttr("opnsense_auth_group.test", "privileges.*", "page-diagnostics-ping"),
					resource.TestCheckTypeSetElemAttr("opnsense_auth_group.test", "privileges.*", "page-diagnostics-traceroute"),
				),
			},
			{
				Config: "\nresource \"opnsense_auth_group\" \"test\" {\n  name = \"tftestgroup\"\n  description = \"tf-test-group-2\"\n  privileges = [\"page-diagnostics-ping\", \"page-diagnostics-traceroute\"]\n}\n\ndata \"opnsense_auth_group\" \"test\" {\n  id = opnsense_auth_group.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_auth_group.test", "name", "opnsense_auth_group.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_auth_group.test", "description", "opnsense_auth_group.test", "description"),
				),
			},
		},
	})
}
