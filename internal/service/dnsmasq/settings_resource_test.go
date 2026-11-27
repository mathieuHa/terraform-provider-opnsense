package dnsmasq_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDnsmasqSettingsResource(t *testing.T) {
	acctest.AccPreCheck(t)

	c := opnsense.NewClient(acctest.Client())
	cur, err := c.Dnsmasq().DnsmasqConfigGet(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	a := "opnsense_dnsmasq_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_dnsmasq_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "dnsmasq_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_dnsmasq_settings" "test" {
  cache_size  = 1234
  log_queries = true
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "cache_size", "1234"),
					resource.TestCheckResourceAttr(a, "log_queries", "true"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "opnsense_dnsmasq_settings" "test" {
  cache_size  = %s
  log_queries = %t
}
`, settingsIntArg(cur.Monad.CacheSize), cur.Monad.LogQueries == "1"),
			},
		},
	})
}

func settingsIntArg(s string) string {
	if s == "" {
		return "-1"
	}
	return s
}

func hclList(v []string) string {
	out := "["
	for i, s := range v {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("%q", s)
	}
	return out + "]"
}
