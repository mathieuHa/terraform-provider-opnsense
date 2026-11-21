package acmeclient_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAcmeclientValidationResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/acmeclient/settings/get", "plugin os-acme-client") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_acmeclient_validation\" \"test\" {\n  name = \"tfacc\"\n  method = \"dns01\"\n  dns_service = \"dns_cf\"\n  dns_cf_token = \"cf-token-tfacc\"\n  dns_cf_account_id = \"acct-tfacc\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_acmeclient_validation.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_validation.test", "method", "dns01"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_validation.test", "dns_service", "dns_cf"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_validation.test", "dns_cf_token", "cf-token-tfacc"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_validation.test", "dns_cf_account_id", "acct-tfacc"),
					resource.TestCheckResourceAttrSet("opnsense_acmeclient_validation.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_acmeclient_validation.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_acmeclient_validation\" \"test\" {\n  name = \"tfacc2\"\n  method = \"dns01\"\n  dns_service = \"dns_cf\"\n  dns_cf_token = \"cf-token-tfacc2\"\n  dns_sleep = 120\n  description = \"tfacc\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_acmeclient_validation.test", "name", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_validation.test", "method", "dns01"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_validation.test", "dns_service", "dns_cf"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_validation.test", "dns_cf_token", "cf-token-tfacc2"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_validation.test", "dns_sleep", "120"),
					resource.TestCheckResourceAttr("opnsense_acmeclient_validation.test", "description", "tfacc"),
				),
			},
			{
				Config: "\nresource \"opnsense_acmeclient_validation\" \"test\" {\n  name = \"tfacc2\"\n  method = \"dns01\"\n  dns_service = \"dns_cf\"\n  dns_cf_token = \"cf-token-tfacc2\"\n  dns_sleep = 120\n  description = \"tfacc\"\n}\n\ndata \"opnsense_acmeclient_validation\" \"test\" {\n  id = opnsense_acmeclient_validation.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_acmeclient_validation.test", "name", "opnsense_acmeclient_validation.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_acmeclient_validation.test", "method", "opnsense_acmeclient_validation.test", "method"),
					resource.TestCheckResourceAttrPair("data.opnsense_acmeclient_validation.test", "dns_service", "opnsense_acmeclient_validation.test", "dns_service"),
					resource.TestCheckResourceAttrPair("data.opnsense_acmeclient_validation.test", "dns_cf_token", "opnsense_acmeclient_validation.test", "dns_cf_token"),
					resource.TestCheckResourceAttrPair("data.opnsense_acmeclient_validation.test", "dns_sleep", "opnsense_acmeclient_validation.test", "dns_sleep"),
					resource.TestCheckResourceAttrPair("data.opnsense_acmeclient_validation.test", "description", "opnsense_acmeclient_validation.test", "description"),
				),
			},
		},
	})
}
