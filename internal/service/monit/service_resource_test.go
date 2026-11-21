package monit_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMonitServiceResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "resource \"opnsense_monit_test\" \"t\" {\n  name      = \"tf_test_svc_mem\"\n  condition = \"memory usage is greater than 90%\"\n  action    = \"alert\"\n}\n\nresource \"opnsense_monit_service\" \"test\" {\n  name = \"tf_test_svc\"\n  type = \"system\"\n  tests = [opnsense_monit_test.t.id]\n  description = \"tf-test-svc\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_monit_service.test", "name", "tf_test_svc"),
					resource.TestCheckResourceAttr("opnsense_monit_service.test", "type", "system"),
					resource.TestCheckResourceAttr("opnsense_monit_service.test", "description", "tf-test-svc"),
					resource.TestCheckResourceAttrSet("opnsense_monit_service.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_monit_service.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "resource \"opnsense_monit_test\" \"t\" {\n  name      = \"tf_test_svc_mem\"\n  condition = \"memory usage is greater than 90%\"\n  action    = \"alert\"\n}\n\nresource \"opnsense_monit_service\" \"test\" {\n  name = \"tf_test_svc\"\n  type = \"system\"\n  start_timeout = 60\n  tests = []\n  description = \"tf-test-svc-2\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_monit_service.test", "name", "tf_test_svc"),
					resource.TestCheckResourceAttr("opnsense_monit_service.test", "type", "system"),
					resource.TestCheckResourceAttr("opnsense_monit_service.test", "start_timeout", "60"),
					resource.TestCheckResourceAttr("opnsense_monit_service.test", "description", "tf-test-svc-2"),
				),
			},
			{
				Config: "resource \"opnsense_monit_test\" \"t\" {\n  name      = \"tf_test_svc_mem\"\n  condition = \"memory usage is greater than 90%\"\n  action    = \"alert\"\n}\n\nresource \"opnsense_monit_service\" \"test\" {\n  name = \"tf_test_svc\"\n  type = \"system\"\n  start_timeout = 60\n  tests = []\n  description = \"tf-test-svc-2\"\n}\n\ndata \"opnsense_monit_service\" \"test\" {\n  id = opnsense_monit_service.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_monit_service.test", "name", "opnsense_monit_service.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_monit_service.test", "type", "opnsense_monit_service.test", "type"),
					resource.TestCheckResourceAttrPair("data.opnsense_monit_service.test", "start_timeout", "opnsense_monit_service.test", "start_timeout"),
					resource.TestCheckResourceAttrPair("data.opnsense_monit_service.test", "description", "opnsense_monit_service.test", "description"),
				),
			},
		},
	})
}
