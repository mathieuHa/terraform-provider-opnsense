package core_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCoreTunableResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTunableConfig("0"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_core_tunable.test", "tunable", "net.inet.tcp.delayed_ack"),
					resource.TestCheckResourceAttr("opnsense_core_tunable.test", "value", "0"),
					resource.TestCheckResourceAttrSet("opnsense_core_tunable.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_core_tunable.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccTunableConfig("1") + `
data "opnsense_core_tunable" "test" {
  id = opnsense_core_tunable.test.id
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_core_tunable.test", "value", "1"),
					resource.TestCheckResourceAttr("data.opnsense_core_tunable.test", "value", "1"),
				),
			},
		},
	})
}

func testAccTunableConfig(value string) string {
	return fmt.Sprintf(`
resource "opnsense_core_tunable" "test" {
  tunable     = "net.inet.tcp.delayed_ack"
  value       = %q
  description = "tf-test-tunable"
}
`, value)
}
