package hasync_test

import (
	"regexp"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHasyncSettingsResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_hasync_settings" "test" {}`,
				ResourceName:       "opnsense_hasync_settings.test",
				ImportState:        true,
				ImportStateId:      "hasync_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_hasync_settings" "test" {
  disable_preempt   = true
  synchronize_to_ip = "192.0.2.10"
  username          = "tfacc"
  password          = "tfacc-secret"
  sync_items        = ["aliases", "rules"]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_hasync_settings.test", "disable_preempt", "true"),
					resource.TestCheckResourceAttr("opnsense_hasync_settings.test", "synchronize_to_ip", "192.0.2.10"),
					resource.TestCheckResourceAttr("opnsense_hasync_settings.test", "sync_items.#", "2"),
				),
			},
			{
				Config: `
resource "opnsense_hasync_settings" "test" {
  disable_preempt   = false
  synchronize_to_ip = ""
  username          = ""
  password          = ""
  sync_items        = []
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_hasync_settings.test", "synchronize_to_ip", ""),
					resource.TestCheckResourceAttr("opnsense_hasync_settings.test", "sync_items.#", "0"),
				),
			},
		},
	})
}

func TestAccHasyncSettingsResource_CreateBlocked(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      `resource "opnsense_hasync_settings" "test" {}`,
				ExpectError: regexp.MustCompile("Cannot Create Singleton Resource"),
			},
		},
	})
}
