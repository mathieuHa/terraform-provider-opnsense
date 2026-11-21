package haproxy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHaproxyCpuResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/haproxy/settings/get", "plugin os-haproxy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_haproxy_cpu\" \"test\" {\n  name = \"tfacc\"\n  thread_id = \"all\"\n  cpu_id = [\"all\"]\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_cpu.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_haproxy_cpu.test", "thread_id", "all"),
					resource.TestCheckResourceAttr("opnsense_haproxy_cpu.test", "cpu_id.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_haproxy_cpu.test", "cpu_id.*", "all"),
					resource.TestCheckResourceAttrSet("opnsense_haproxy_cpu.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_haproxy_cpu.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_haproxy_cpu\" \"test\" {\n  name = \"tfacc2\"\n  thread_id = \"odd\"\n  cpu_id = [\"odd\"]\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_cpu.test", "name", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_haproxy_cpu.test", "thread_id", "odd"),
					resource.TestCheckResourceAttr("opnsense_haproxy_cpu.test", "cpu_id.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_haproxy_cpu.test", "cpu_id.*", "odd"),
				),
			},
			{
				Config: "\nresource \"opnsense_haproxy_cpu\" \"test\" {\n  name = \"tfacc2\"\n  thread_id = \"odd\"\n  cpu_id = [\"odd\"]\n}\n\ndata \"opnsense_haproxy_cpu\" \"test\" {\n  id = opnsense_haproxy_cpu.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_cpu.test", "name", "opnsense_haproxy_cpu.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_cpu.test", "thread_id", "opnsense_haproxy_cpu.test", "thread_id"),
				),
			},
		},
	})
}
