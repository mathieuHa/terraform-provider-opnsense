package haproxy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHaproxyAclResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/haproxy/settings/get", "plugin os-haproxy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_haproxy_acl\" \"test\" {\n  name = \"tfacc\"\n  expression = \"hdr_beg\"\n  hdr_beg = \"tfacc.\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_acl.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_haproxy_acl.test", "expression", "hdr_beg"),
					resource.TestCheckResourceAttr("opnsense_haproxy_acl.test", "hdr_beg", "tfacc."),
					resource.TestCheckResourceAttrSet("opnsense_haproxy_acl.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_haproxy_acl.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_haproxy_acl\" \"test\" {\n  name = \"tfacc2\"\n  expression = \"hdr_beg\"\n  hdr_beg = \"tfacc2.\"\n  negate = true\n  description = \"tfacc\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_acl.test", "name", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_haproxy_acl.test", "expression", "hdr_beg"),
					resource.TestCheckResourceAttr("opnsense_haproxy_acl.test", "hdr_beg", "tfacc2."),
					resource.TestCheckResourceAttr("opnsense_haproxy_acl.test", "negate", "true"),
					resource.TestCheckResourceAttr("opnsense_haproxy_acl.test", "description", "tfacc"),
				),
			},
			{
				Config: "\nresource \"opnsense_haproxy_acl\" \"test\" {\n  name = \"tfacc2\"\n  expression = \"hdr_beg\"\n  hdr_beg = \"tfacc2.\"\n  negate = true\n  description = \"tfacc\"\n}\n\ndata \"opnsense_haproxy_acl\" \"test\" {\n  id = opnsense_haproxy_acl.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_acl.test", "name", "opnsense_haproxy_acl.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_acl.test", "expression", "opnsense_haproxy_acl.test", "expression"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_acl.test", "hdr_beg", "opnsense_haproxy_acl.test", "hdr_beg"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_acl.test", "negate", "opnsense_haproxy_acl.test", "negate"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_acl.test", "description", "opnsense_haproxy_acl.test", "description"),
				),
			},
		},
	})
}
