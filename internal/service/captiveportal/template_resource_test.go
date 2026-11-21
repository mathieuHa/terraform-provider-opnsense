package captiveportal_test

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"regexp"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func portalZip(t *testing.T, body string) string {
	var buf bytes.Buffer
	z := zip.NewWriter(&buf)
	w, err := z.Create("index.html")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Fprintf(w, "<html><body>%s</body></html>", body)
	if err := z.Close(); err != nil {
		t.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func TestAccCaptivePortalTemplateResource(t *testing.T) {
	v1, v2 := portalZip(t, "tfacc portal"), portalZip(t, "tfacc portal v2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.OptionPreCheck(t, "/captiveportal/settings/get_zone", "zone", "interfaces", "opt1") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "opnsense_captiveportal_template" "test" {
  name    = "tfacc"
  content = %q
}

resource "opnsense_captiveportal_zone" "test" {
  interfaces  = ["opt1"]
  description = "tfacc template zone"
  template    = opnsense_captiveportal_template.test.id
}
`, v1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr("opnsense_captiveportal_template.test", "id", acctest.UUIDRegexp()),
					resource.TestCheckResourceAttrSet("opnsense_captiveportal_template.test", "file_id"),
					resource.TestCheckResourceAttrPair("opnsense_captiveportal_zone.test", "template", "opnsense_captiveportal_template.test", "id"),
					resource.TestCheckResourceAttrSet("opnsense_captiveportal_zone.test", "zone_id"),
				),
			},
			{
				ResourceName:            "opnsense_captiveportal_template.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"content"},
			},
			{
				Config: fmt.Sprintf(`
resource "opnsense_captiveportal_template" "test" {
  name    = "tfacc2"
  content = %q
}
`, v2),
				Check: resource.TestCheckResourceAttr("opnsense_captiveportal_template.test", "name", "tfacc2"),
			},
			{
				Config: `
resource "opnsense_captiveportal_template" "bad" {
  name    = "tfacc-bad"
  content = "bm90IGEgemlwIGZpbGUgYXQgYWxs"
}
`,
				ExpectError: regexp.MustCompile(`must be a base64-encoded ZIP`),
			},
		},
	})
}
