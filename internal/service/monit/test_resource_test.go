package monit_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMonitTestResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_monit_test\" \"test\" {\n  name = \"tf_test_cpu\"\n  condition = \"cpu usage is greater than 90%\"\n  action = \"alert\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_monit_test.test", "name", "tf_test_cpu"),
					resource.TestCheckResourceAttr("opnsense_monit_test.test", "condition", "cpu usage is greater than 90%"),
					resource.TestCheckResourceAttr("opnsense_monit_test.test", "action", "alert"),
					resource.TestCheckResourceAttrSet("opnsense_monit_test.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_monit_test.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_monit_test\" \"test\" {\n  name = \"tf_test_cpu\"\n  condition = \"cpu usage is greater than 95%\"\n  action = \"alert\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_monit_test.test", "name", "tf_test_cpu"),
					resource.TestCheckResourceAttr("opnsense_monit_test.test", "condition", "cpu usage is greater than 95%"),
					resource.TestCheckResourceAttr("opnsense_monit_test.test", "action", "alert"),
				),
			},
			{
				Config: "\nresource \"opnsense_monit_test\" \"test\" {\n  name = \"tf_test_cpu\"\n  condition = \"cpu usage is greater than 95%\"\n  action = \"alert\"\n}\n\ndata \"opnsense_monit_test\" \"test\" {\n  id = opnsense_monit_test.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_monit_test.test", "name", "opnsense_monit_test.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_monit_test.test", "condition", "opnsense_monit_test.test", "condition"),
					resource.TestCheckResourceAttrPair("data.opnsense_monit_test.test", "action", "opnsense_monit_test.test", "action"),
				),
			},
		},
	})
}
