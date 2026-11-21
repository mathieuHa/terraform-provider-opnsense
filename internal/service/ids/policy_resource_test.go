package ids_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIdsPolicyResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_ids_policy\" \"test\" {\n  description = \"tfacc policy\"\n  new_action = \"drop\"\n  action = [\"alert\"]\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_ids_policy.test", "description", "tfacc policy"),
					resource.TestCheckResourceAttr("opnsense_ids_policy.test", "new_action", "drop"),
					resource.TestCheckResourceAttr("opnsense_ids_policy.test", "action.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_ids_policy.test", "action.*", "alert"),
					resource.TestCheckResourceAttrSet("opnsense_ids_policy.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_ids_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_ids_policy\" \"test\" {\n  description = \"tfacc policy 2\"\n  priority = 5\n  new_action = \"disable\"\n  action = [\"alert\", \"drop\"]\n  enabled = false\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_ids_policy.test", "description", "tfacc policy 2"),
					resource.TestCheckResourceAttr("opnsense_ids_policy.test", "priority", "5"),
					resource.TestCheckResourceAttr("opnsense_ids_policy.test", "new_action", "disable"),
					resource.TestCheckResourceAttr("opnsense_ids_policy.test", "action.#", "2"),
					resource.TestCheckTypeSetElemAttr("opnsense_ids_policy.test", "action.*", "alert"),
					resource.TestCheckTypeSetElemAttr("opnsense_ids_policy.test", "action.*", "drop"),
					resource.TestCheckResourceAttr("opnsense_ids_policy.test", "enabled", "false"),
				),
			},
			{
				Config: "\nresource \"opnsense_ids_policy\" \"test\" {\n  description = \"tfacc policy 2\"\n  priority = 5\n  new_action = \"disable\"\n  action = [\"alert\", \"drop\"]\n  enabled = false\n}\n\ndata \"opnsense_ids_policy\" \"test\" {\n  id = opnsense_ids_policy.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_ids_policy.test", "description", "opnsense_ids_policy.test", "description"),
					resource.TestCheckResourceAttrPair("data.opnsense_ids_policy.test", "priority", "opnsense_ids_policy.test", "priority"),
					resource.TestCheckResourceAttrPair("data.opnsense_ids_policy.test", "new_action", "opnsense_ids_policy.test", "new_action"),
					resource.TestCheckResourceAttrPair("data.opnsense_ids_policy.test", "enabled", "opnsense_ids_policy.test", "enabled"),
				),
			},
		},
	})
}
