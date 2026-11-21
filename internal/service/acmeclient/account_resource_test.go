package acmeclient_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAcmeclientAccountResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/acmeclient/settings/get", "plugin os-acme-client") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_acmeclient_account\" \"test\" {\n  name = \"tfacc\"\n  email = \"tfacc@example.com\"\n  ca = \"letsencrypt_test\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_acmeclient_account.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_account.test", "email", "tfacc@example.com"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_account.test", "ca", "letsencrypt_test"),
					resource.TestCheckResourceAttrSet("opnsense_acmeclient_account.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_acmeclient_account.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_acmeclient_account\" \"test\" {\n  name = \"tfacc2\"\n  email = \"tfacc2@example.com\"\n  ca = \"zerossl\"\n  eab_kid = \"kid-tfacc\"\n  eab_hmac = \"aG1hYy10ZmFjYw\"\n  description = \"tfacc\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_acmeclient_account.test", "name", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_account.test", "email", "tfacc2@example.com"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_account.test", "ca", "zerossl"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_account.test", "eab_kid", "kid-tfacc"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_account.test", "eab_hmac", "aG1hYy10ZmFjYw"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_account.test", "description", "tfacc"),
				),
			},
			{
				Config: "\nresource \"opnsense_acmeclient_account\" \"test\" {\n  name = \"tfacc2\"\n  email = \"tfacc2@example.com\"\n  ca = \"zerossl\"\n  eab_kid = \"kid-tfacc\"\n  eab_hmac = \"aG1hYy10ZmFjYw\"\n  description = \"tfacc\"\n}\n\ndata \"opnsense_acmeclient_account\" \"test\" {\n  id = opnsense_acmeclient_account.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_acmeclient_account.test", "name", "opnsense_acmeclient_account.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_acmeclient_account.test", "email", "opnsense_acmeclient_account.test", "email"),
					resource.TestCheckResourceAttrPair("data.opnsense_acmeclient_account.test", "ca", "opnsense_acmeclient_account.test", "ca"),
					resource.TestCheckResourceAttrPair("data.opnsense_acmeclient_account.test", "eab_kid", "opnsense_acmeclient_account.test", "eab_kid"),
					resource.TestCheckResourceAttrPair("data.opnsense_acmeclient_account.test", "eab_hmac", "opnsense_acmeclient_account.test", "eab_hmac"),
					resource.TestCheckResourceAttrPair("data.opnsense_acmeclient_account.test", "description", "opnsense_acmeclient_account.test", "description"),
				),
			},
		},
	})
}
