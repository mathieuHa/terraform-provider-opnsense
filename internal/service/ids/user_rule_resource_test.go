package ids_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIdsUserRuleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_ids_user_rule\" \"test\" {\n  description = \"tfacc user rule\"\n  source = \"192.0.2.10\"\n  action = \"alert\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_ids_user_rule.test", "description", "tfacc user rule"),
					resource.TestCheckResourceAttr("opnsense_ids_user_rule.test", "source", "192.0.2.10"),
					resource.TestCheckResourceAttr("opnsense_ids_user_rule.test", "action", "alert"),
					resource.TestCheckResourceAttrSet("opnsense_ids_user_rule.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_ids_user_rule.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_ids_user_rule\" \"test\" {\n  description = \"tfacc user rule 2\"\n  source = \"192.0.2.0/24\"\n  destination = \"198.51.100.1\"\n  action = \"pass\"\n  bypass = true\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_ids_user_rule.test", "description", "tfacc user rule 2"),
					resource.TestCheckResourceAttr("opnsense_ids_user_rule.test", "source", "192.0.2.0/24"),
					resource.TestCheckResourceAttr("opnsense_ids_user_rule.test", "destination", "198.51.100.1"),
					resource.TestCheckResourceAttr("opnsense_ids_user_rule.test", "action", "pass"),
					resource.TestCheckResourceAttr("opnsense_ids_user_rule.test", "bypass", "true"),
				),
			},
			{
				Config: "\nresource \"opnsense_ids_user_rule\" \"test\" {\n  description = \"tfacc user rule 2\"\n  source = \"192.0.2.0/24\"\n  destination = \"198.51.100.1\"\n  action = \"pass\"\n  bypass = true\n}\n\ndata \"opnsense_ids_user_rule\" \"test\" {\n  id = opnsense_ids_user_rule.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_ids_user_rule.test", "description", "opnsense_ids_user_rule.test", "description"),
					resource.TestCheckResourceAttrPair("data.opnsense_ids_user_rule.test", "source", "opnsense_ids_user_rule.test", "source"),
					resource.TestCheckResourceAttrPair("data.opnsense_ids_user_rule.test", "destination", "opnsense_ids_user_rule.test", "destination"),
					resource.TestCheckResourceAttrPair("data.opnsense_ids_user_rule.test", "action", "opnsense_ids_user_rule.test", "action"),
					resource.TestCheckResourceAttrPair("data.opnsense_ids_user_rule.test", "bypass", "opnsense_ids_user_rule.test", "bypass"),
				),
			},
		},
	})
}
