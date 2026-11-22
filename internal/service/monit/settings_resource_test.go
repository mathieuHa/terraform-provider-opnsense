package monit_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMonitSettingsResource(t *testing.T) {

	c := opnsense.NewClient(acctest.Client())
	cur, err := c.Monit().GeneralConfigGet(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	a := "opnsense_monit_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_monit_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "monit_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_monit_settings" "test" {
  interval   = 90
  mailserver = ["192.0.2.25", "192.0.2.26"]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "interval", "90"),
					resource.TestCheckResourceAttr(a, "mailserver.#", "2"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "opnsense_monit_settings" "test" {
  interval   = %s
  mailserver = %s
}
`, settingsIntArg(cur.Monad.General.Interval), hclList(cur.Monad.General.Mailserver)),
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
