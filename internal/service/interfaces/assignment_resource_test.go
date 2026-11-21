package interfaces_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInterfacesAssignmentResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAssignmentConfig("TFACC", false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr("opnsense_interfaces_assignment.test", "id", regexp.MustCompile(`^opt\d+$`)),
					resource.TestCheckResourceAttrPair("opnsense_interfaces_assignment.test", "device", "opnsense_interfaces_loopback.test", "device"),
					resource.TestCheckResourceAttr("opnsense_interfaces_assignment.test", "description", "TFACC"),
					resource.TestCheckResourceAttr("opnsense_interfaces_assignment.test", "lock", "false"),
				),
			},
			{
				ResourceName:      "opnsense_interfaces_assignment.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAssignmentConfig("TFACC2", true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_interfaces_assignment.test", "description", "TFACC2"),
					resource.TestCheckResourceAttr("opnsense_interfaces_assignment.test", "lock", "true"),
				),
			},
		},
	})
}

func testAccAssignmentConfig(description string, lock bool) string {
	return fmt.Sprintf(`
resource "opnsense_interfaces_loopback" "test" {
  description = "tf-acc assignment"
}

resource "opnsense_interfaces_assignment" "test" {
  device      = opnsense_interfaces_loopback.test.device
  description = %q
  lock        = %t
}
`, description, lock)
}
