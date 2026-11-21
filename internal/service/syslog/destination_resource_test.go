package syslog_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSyslogDestinationResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_syslog_destination\" \"test\" {\n  hostname = \"192.0.2.10\"\n  transport = \"tcp4\"\n  port = 6514\n  program = [\"filterlog\"]\n  rfc5424 = true\n  description = \"tf-test-syslog\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_syslog_destination.test", "hostname", "192.0.2.10"),
					resource.TestCheckResourceAttr("opnsense_syslog_destination.test", "transport", "tcp4"),
					resource.TestCheckResourceAttr("opnsense_syslog_destination.test", "port", "6514"),
					resource.TestCheckResourceAttr("opnsense_syslog_destination.test", "program.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_syslog_destination.test", "program.*", "filterlog"),
					resource.TestCheckResourceAttr("opnsense_syslog_destination.test", "rfc5424", "true"),
					resource.TestCheckResourceAttr("opnsense_syslog_destination.test", "description", "tf-test-syslog"),
					resource.TestCheckResourceAttrSet("opnsense_syslog_destination.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_syslog_destination.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_syslog_destination\" \"test\" {\n  hostname = \"192.0.2.11\"\n  transport = \"udp4\"\n  port = 514\n  program = [\"filterlog\", \"firewall\"]\n  level = [\"err\", \"warn\"]\n  facility = [\"local0\"]\n  rfc5424 = false\n  description = \"tf-test-syslog-2\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_syslog_destination.test", "hostname", "192.0.2.11"),
					resource.TestCheckResourceAttr("opnsense_syslog_destination.test", "transport", "udp4"),
					resource.TestCheckResourceAttr("opnsense_syslog_destination.test", "port", "514"),
					resource.TestCheckResourceAttr("opnsense_syslog_destination.test", "program.#", "2"),
					resource.TestCheckTypeSetElemAttr("opnsense_syslog_destination.test", "program.*", "filterlog"),
					resource.TestCheckTypeSetElemAttr("opnsense_syslog_destination.test", "program.*", "firewall"),
					resource.TestCheckResourceAttr("opnsense_syslog_destination.test", "level.#", "2"),
					resource.TestCheckTypeSetElemAttr("opnsense_syslog_destination.test", "level.*", "err"),
					resource.TestCheckTypeSetElemAttr("opnsense_syslog_destination.test", "level.*", "warn"),
					resource.TestCheckResourceAttr("opnsense_syslog_destination.test", "facility.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_syslog_destination.test", "facility.*", "local0"),
					resource.TestCheckResourceAttr("opnsense_syslog_destination.test", "rfc5424", "false"),
					resource.TestCheckResourceAttr("opnsense_syslog_destination.test", "description", "tf-test-syslog-2"),
				),
			},
			{
				Config: "\nresource \"opnsense_syslog_destination\" \"test\" {\n  hostname = \"192.0.2.11\"\n  transport = \"udp4\"\n  port = 514\n  program = [\"filterlog\", \"firewall\"]\n  level = [\"err\", \"warn\"]\n  facility = [\"local0\"]\n  rfc5424 = false\n  description = \"tf-test-syslog-2\"\n}\n\ndata \"opnsense_syslog_destination\" \"test\" {\n  id = opnsense_syslog_destination.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_syslog_destination.test", "hostname", "opnsense_syslog_destination.test", "hostname"),
					resource.TestCheckResourceAttrPair("data.opnsense_syslog_destination.test", "transport", "opnsense_syslog_destination.test", "transport"),
					resource.TestCheckResourceAttrPair("data.opnsense_syslog_destination.test", "port", "opnsense_syslog_destination.test", "port"),
					resource.TestCheckResourceAttrPair("data.opnsense_syslog_destination.test", "rfc5424", "opnsense_syslog_destination.test", "rfc5424"),
					resource.TestCheckResourceAttrPair("data.opnsense_syslog_destination.test", "description", "opnsense_syslog_destination.test", "description"),
				),
			},
		},
	})
}
