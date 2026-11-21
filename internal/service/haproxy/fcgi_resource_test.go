package haproxy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHaproxyFcgiResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/haproxy/settings/get", "plugin os-haproxy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_haproxy_fcgi\" \"test\" {\n  name = \"tfacc\"\n  docroot = \"/var/www\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_fcgi.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_haproxy_fcgi.test", "docroot", "/var/www"),
					resource.TestCheckResourceAttrSet("opnsense_haproxy_fcgi.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_haproxy_fcgi.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_haproxy_fcgi\" \"test\" {\n  name = \"tfacc2\"\n  docroot = \"/var/www/html\"\n  description = \"tfacc\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_fcgi.test", "name", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_haproxy_fcgi.test", "docroot", "/var/www/html"),
					resource.TestCheckResourceAttr("opnsense_haproxy_fcgi.test", "description", "tfacc"),
				),
			},
			{
				Config: "\nresource \"opnsense_haproxy_fcgi\" \"test\" {\n  name = \"tfacc2\"\n  docroot = \"/var/www/html\"\n  description = \"tfacc\"\n}\n\ndata \"opnsense_haproxy_fcgi\" \"test\" {\n  id = opnsense_haproxy_fcgi.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_fcgi.test", "name", "opnsense_haproxy_fcgi.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_fcgi.test", "docroot", "opnsense_haproxy_fcgi.test", "docroot"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_fcgi.test", "description", "opnsense_haproxy_fcgi.test", "description"),
				),
			},
		},
	})
}
