package haproxy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHaproxyMailerResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/haproxy/settings/get", "plugin os-haproxy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_haproxy_mailer\" \"test\" {\n  name = \"tfacc\"\n  mailservers = [\"192.0.2.25:25\"]\n  sender = \"haproxy@example.com\"\n  recipient = \"ops@example.com\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_mailer.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_haproxy_mailer.test", "mailservers.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_haproxy_mailer.test", "mailservers.*", "192.0.2.25:25"),
					resource.TestCheckResourceAttr("opnsense_haproxy_mailer.test", "sender", "haproxy@example.com"),
					resource.TestCheckResourceAttr("opnsense_haproxy_mailer.test", "recipient", "ops@example.com"),
					resource.TestCheckResourceAttrSet("opnsense_haproxy_mailer.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_haproxy_mailer.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_haproxy_mailer\" \"test\" {\n  name = \"tfacc2\"\n  mailservers = [\"192.0.2.25:25\"]\n  sender = \"haproxy@example.com\"\n  recipient = \"ops2@example.com\"\n  description = \"tfacc\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_mailer.test", "name", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_haproxy_mailer.test", "mailservers.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_haproxy_mailer.test", "mailservers.*", "192.0.2.25:25"),
					resource.TestCheckResourceAttr("opnsense_haproxy_mailer.test", "sender", "haproxy@example.com"),
					resource.TestCheckResourceAttr("opnsense_haproxy_mailer.test", "recipient", "ops2@example.com"),
					resource.TestCheckResourceAttr("opnsense_haproxy_mailer.test", "description", "tfacc"),
				),
			},
			{
				Config: "\nresource \"opnsense_haproxy_mailer\" \"test\" {\n  name = \"tfacc2\"\n  mailservers = [\"192.0.2.25:25\"]\n  sender = \"haproxy@example.com\"\n  recipient = \"ops2@example.com\"\n  description = \"tfacc\"\n}\n\ndata \"opnsense_haproxy_mailer\" \"test\" {\n  id = opnsense_haproxy_mailer.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_mailer.test", "name", "opnsense_haproxy_mailer.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_mailer.test", "sender", "opnsense_haproxy_mailer.test", "sender"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_mailer.test", "recipient", "opnsense_haproxy_mailer.test", "recipient"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_mailer.test", "description", "opnsense_haproxy_mailer.test", "description"),
				),
			},
		},
	})
}
