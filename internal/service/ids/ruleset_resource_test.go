package ids_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testRuleset = "emerging-scan.rules"

func TestAccIdsRulesetResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccRulesetDisabled,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "opnsense_ids_ruleset" "test" {
  filename = %q
}

resource "opnsense_ids_policy" "test" {
  description = "tfacc ruleset policy"
  rulesets    = [opnsense_ids_ruleset.test.uuid]
  new_action  = "drop"
}
`, testRuleset),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_ids_ruleset.test", "id", testRuleset),
					resource.TestCheckResourceAttr("opnsense_ids_ruleset.test", "description", "ET open/emerging-scan"),
					resource.TestMatchResourceAttr("opnsense_ids_ruleset.test", "uuid", acctest.UUIDRegexp()),
					resource.TestCheckResourceAttrPair("opnsense_ids_policy.test", "rulesets.0", "opnsense_ids_ruleset.test", "uuid"),
				),
			},
			{
				ResourceName:      "opnsense_ids_ruleset.test",
				ImportState:       true,
				ImportStateId:     testRuleset,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIdsRulesetResource_Unknown(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      `resource "opnsense_ids_ruleset" "test" { filename = "tfacc-missing.rules" }`,
				ExpectError: regexp.MustCompile(`unknown ruleset`),
			},
		},
	})
}

func testAccRulesetDisabled(s *terraform.State) error {
	c := opnsense.NewClient(acctest.Client())
	rs, err := c.Ids().RulesetGet(context.Background(), testRuleset)
	if err != nil {
		return err
	}
	if rs.Enabled != "0" {
		return fmt.Errorf("ruleset %s still enabled", testRuleset)
	}
	return nil
}
