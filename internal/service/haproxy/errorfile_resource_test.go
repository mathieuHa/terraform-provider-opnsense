package haproxy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHaproxyErrorfileResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/haproxy/settings/get", "plugin os-haproxy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_haproxy_errorfile\" \"test\" {\n  name = \"tfacc\"\n  code = \"x503\"\n  content = \"HTTP/1.0 503 Service Unavailable\\r\\nContent-Type: text/plain\\r\\n\\r\\ntfacc\\r\\n\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_errorfile.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_haproxy_errorfile.test", "code", "x503"),
					resource.TestCheckResourceAttr("opnsense_haproxy_errorfile.test", "content", "HTTP/1.0 503 Service Unavailable\r\nContent-Type: text/plain\r\n\r\ntfacc\r\n"),
					resource.TestCheckResourceAttrSet("opnsense_haproxy_errorfile.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_haproxy_errorfile.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_haproxy_errorfile\" \"test\" {\n  name = \"tfacc2\"\n  code = \"x503\"\n  content = \"HTTP/1.0 503 Service Unavailable\\r\\nContent-Type: text/plain\\r\\n\\r\\ntfacc2\\r\\n\"\n  description = \"tfacc\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_errorfile.test", "name", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_haproxy_errorfile.test", "code", "x503"),
					resource.TestCheckResourceAttr("opnsense_haproxy_errorfile.test", "content", "HTTP/1.0 503 Service Unavailable\r\nContent-Type: text/plain\r\n\r\ntfacc2\r\n"),
					resource.TestCheckResourceAttr("opnsense_haproxy_errorfile.test", "description", "tfacc"),
				),
			},
			{
				Config: "\nresource \"opnsense_haproxy_errorfile\" \"test\" {\n  name = \"tfacc2\"\n  code = \"x503\"\n  content = \"HTTP/1.0 503 Service Unavailable\\r\\nContent-Type: text/plain\\r\\n\\r\\ntfacc2\\r\\n\"\n  description = \"tfacc\"\n}\n\ndata \"opnsense_haproxy_errorfile\" \"test\" {\n  id = opnsense_haproxy_errorfile.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_errorfile.test", "name", "opnsense_haproxy_errorfile.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_errorfile.test", "code", "opnsense_haproxy_errorfile.test", "code"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_errorfile.test", "content", "opnsense_haproxy_errorfile.test", "content"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_errorfile.test", "description", "opnsense_haproxy_errorfile.test", "description"),
				),
			},
		},
	})
}
