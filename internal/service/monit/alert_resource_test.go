package monit_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMonitAlertResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_monit_alert\" \"test\" {\n  recipient = \"ops@example.com\"\n  events = [\"nonexist\", \"pid\"]\n  description = \"tf-test-alert\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_monit_alert.test", "recipient", "ops@example.com"),
					resource.TestCheckResourceAttr("opnsense_monit_alert.test", "events.#", "2"),
					resource.TestCheckTypeSetElemAttr("opnsense_monit_alert.test", "events.*", "nonexist"),
					resource.TestCheckTypeSetElemAttr("opnsense_monit_alert.test", "events.*", "pid"),
					resource.TestCheckResourceAttr("opnsense_monit_alert.test", "description", "tf-test-alert"),
					resource.TestCheckResourceAttrSet("opnsense_monit_alert.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_monit_alert.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_monit_alert\" \"test\" {\n  recipient = \"ops2@example.com\"\n  not_on = true\n  events = [\"instance\"]\n  reminder = 10\n  description = \"tf-test-alert-2\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_monit_alert.test", "recipient", "ops2@example.com"),
					resource.TestCheckResourceAttr("opnsense_monit_alert.test", "not_on", "true"),
					resource.TestCheckResourceAttr("opnsense_monit_alert.test", "events.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_monit_alert.test", "events.*", "instance"),
					resource.TestCheckResourceAttr("opnsense_monit_alert.test", "reminder", "10"),
					resource.TestCheckResourceAttr("opnsense_monit_alert.test", "description", "tf-test-alert-2"),
				),
			},
			{
				Config: "\nresource \"opnsense_monit_alert\" \"test\" {\n  recipient = \"ops2@example.com\"\n  not_on = true\n  events = [\"instance\"]\n  reminder = 10\n  description = \"tf-test-alert-2\"\n}\n\ndata \"opnsense_monit_alert\" \"test\" {\n  id = opnsense_monit_alert.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_monit_alert.test", "recipient", "opnsense_monit_alert.test", "recipient"),
					resource.TestCheckResourceAttrPair("data.opnsense_monit_alert.test", "not_on", "opnsense_monit_alert.test", "not_on"),
					resource.TestCheckResourceAttrPair("data.opnsense_monit_alert.test", "reminder", "opnsense_monit_alert.test", "reminder"),
					resource.TestCheckResourceAttrPair("data.opnsense_monit_alert.test", "description", "opnsense_monit_alert.test", "description"),
				),
			},
		},
	})
}
