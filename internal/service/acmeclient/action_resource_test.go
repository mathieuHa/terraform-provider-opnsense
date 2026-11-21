package acmeclient_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAcmeclientActionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/acmeclient/settings/get", "plugin os-acme-client") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_acmeclient_action\" \"test\" {\n  name = \"tfacc\"\n  type = \"configd_restart_gui\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_acmeclient_action.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_action.test", "type", "configd_restart_gui"),
					resource.TestCheckResourceAttrSet("opnsense_acmeclient_action.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_acmeclient_action.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_acmeclient_action\" \"test\" {\n  name = \"tfacc2\"\n  type = \"configd_restart_gui\"\n  description = \"tfacc\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_acmeclient_action.test", "name", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_action.test", "type", "configd_restart_gui"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_action.test", "description", "tfacc"),
				),
			},
			{
				Config: "\nresource \"opnsense_acmeclient_action\" \"test\" {\n  name = \"tfacc2\"\n  type = \"configd_restart_gui\"\n  description = \"tfacc\"\n}\n\ndata \"opnsense_acmeclient_action\" \"test\" {\n  id = opnsense_acmeclient_action.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_acmeclient_action.test", "name", "opnsense_acmeclient_action.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_acmeclient_action.test", "type", "opnsense_acmeclient_action.test", "type"),
					resource.TestCheckResourceAttrPair("data.opnsense_acmeclient_action.test", "description", "opnsense_acmeclient_action.test", "description"),
				),
			},
		},
	})
}
