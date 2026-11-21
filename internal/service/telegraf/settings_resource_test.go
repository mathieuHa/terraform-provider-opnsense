package telegraf_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTelegrafGeneralResource(t *testing.T) {
	a := "opnsense_telegraf_general.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/telegraf/general/get", "plugin os-telegraf") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_telegraf_general" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "telegraf_general",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_telegraf_general" "test" {
  enabled  = false
  interval = 20
  debug    = true
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "interval", "20"),
					resource.TestCheckResourceAttr(a, "debug", "true"),
				),
			},
			{
				Config: `
resource "opnsense_telegraf_general" "test" {
  enabled  = false
  interval = 10
  debug    = false
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "interval", "10"),
				),
			},
		},
	})
}

func TestAccTelegrafInputResource(t *testing.T) {
	a := "opnsense_telegraf_input.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/telegraf/input/get", "plugin os-telegraf") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_telegraf_input" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "telegraf_input",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_telegraf_input" "test" {
  cpu  = true
  disk = true
  pf   = true
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "pf", "true"),
				),
			},
			{
				Config: `
resource "opnsense_telegraf_input" "test" {
  cpu  = true
  disk = true
  pf   = false
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "pf", "false"),
				),
			},
		},
	})
}

func TestAccTelegrafOutputResource(t *testing.T) {
	a := "opnsense_telegraf_output.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/telegraf/output/get", "plugin os-telegraf") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_telegraf_output" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "telegraf_output",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_telegraf_output" "test" {
  influx_v2_enable = true
  influx_v2_url    = "http://192.0.2.86:8086"
  influx_v2_token  = "tfacc-token"
  influx_v2_organization = "tfacc"
  influx_v2_bucket = "tfacc"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "influx_v2_enable", "true"),
					resource.TestCheckResourceAttr(a, "influx_v2_url", "http://192.0.2.86:8086"),
				),
			},
			{
				Config: `
resource "opnsense_telegraf_output" "test" {
  influx_v2_enable = false
  influx_v2_url    = ""
  influx_v2_token  = ""
  influx_v2_organization = ""
  influx_v2_bucket = ""
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "influx_v2_enable", "false"),
				),
			},
		},
	})
}
