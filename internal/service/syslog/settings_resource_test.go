package syslog_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSyslogSettingsResource(t *testing.T) {
	acctest.AccPreCheck(t)

	c := opnsense.NewClient(acctest.Client())
	cur, err := c.Syslog().SyslogConfigGet(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	a := "opnsense_syslog_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_syslog_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "syslog_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_syslog_settings" "test" {
  maxpreserve = 14
  maxfilesize = 4096
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "maxpreserve", "14"),
					resource.TestCheckResourceAttr(a, "maxfilesize", "4096"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "opnsense_syslog_settings" "test" {
  maxpreserve = %s
  maxfilesize = %s
}
`, settingsIntArg(cur.Monad.General.Maxpreserve), settingsIntArg(cur.Monad.General.Maxfilesize)),
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
