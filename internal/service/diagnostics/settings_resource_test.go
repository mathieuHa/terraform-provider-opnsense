package diagnostics_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetflowSettingsResource(t *testing.T) {

	c := opnsense.NewClient(acctest.Client())
	cur, err := c.Diagnostics().NetflowConfigGet(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	a := "opnsense_netflow_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_netflow_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "netflow_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_netflow_settings" "test" {
  active_timeout   = 900
  inactive_timeout = 30
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "active_timeout", "900"),
					resource.TestCheckResourceAttr(a, "inactive_timeout", "30"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "opnsense_netflow_settings" "test" {
  active_timeout   = %s
  inactive_timeout = %s
}
`, settingsIntArg(cur.Monad.ActiveTimeout), settingsIntArg(cur.Monad.InactiveTimeout)),
			},
		},
	})
}

func TestAccSystemhealthSettingsResource(t *testing.T) {

	c := opnsense.NewClient(acctest.Client())
	cur, err := c.Diagnostics().SystemHealthConfigGet(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	a := "opnsense_systemhealth_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_systemhealth_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "systemhealth_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_systemhealth_settings" "test" {
  enabled = false
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "enabled", "false"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "opnsense_systemhealth_settings" "test" {
  enabled = %t
}
`, cur.Monad.Enabled == "1"),
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
