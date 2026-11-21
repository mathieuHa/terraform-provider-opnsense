package ids_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIdsPolicyRuleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_ids_policy_rule\" \"test\" {\n  sid = 2100498\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_ids_policy_rule.test", "sid", "2100498"),
					resource.TestCheckResourceAttrSet("opnsense_ids_policy_rule.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_ids_policy_rule.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_ids_policy_rule\" \"test\" {\n  sid = 2100498\n  action = \"drop\"\n  enabled = false\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_ids_policy_rule.test", "sid", "2100498"),
					resource.TestCheckResourceAttr("opnsense_ids_policy_rule.test", "action", "drop"),
					resource.TestCheckResourceAttr("opnsense_ids_policy_rule.test", "enabled", "false"),
				),
			},
			{
				Config: "\nresource \"opnsense_ids_policy_rule\" \"test\" {\n  sid = 2100498\n  action = \"drop\"\n  enabled = false\n}\n\ndata \"opnsense_ids_policy_rule\" \"test\" {\n  id = opnsense_ids_policy_rule.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_ids_policy_rule.test", "sid", "opnsense_ids_policy_rule.test", "sid"),
					resource.TestCheckResourceAttrPair("data.opnsense_ids_policy_rule.test", "action", "opnsense_ids_policy_rule.test", "action"),
					resource.TestCheckResourceAttrPair("data.opnsense_ids_policy_rule.test", "enabled", "opnsense_ids_policy_rule.test", "enabled"),
				),
			},
		},
	})
}
