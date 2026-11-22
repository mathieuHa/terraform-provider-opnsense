package auth_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccAuthUserApiKeyResource(t *testing.T) {
	var key string
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy: func(s *terraform.State) error {
			res, err := opnsense.NewClient(acctest.Client()).Auth().UserGetAllApiKeys(context.Background())
			if err != nil {
				return err
			}
			for _, row := range res.Rows {
				if row.Key == key {
					return fmt.Errorf("API key %s still exists", key)
				}
			}
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: `
resource "opnsense_auth_user" "test" {
  name               = "tfacc-apikey"
  scrambled_password = true
}

resource "opnsense_auth_user_api_key" "test" {
  username = opnsense_auth_user.test.name
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_auth_user_api_key.test", "username", "tfacc-apikey"),
					resource.TestCheckResourceAttrSet("opnsense_auth_user_api_key.test", "key"),
					resource.TestCheckResourceAttrSet("opnsense_auth_user_api_key.test", "secret"),
					func(s *terraform.State) error {
						attrs := s.RootModule().Resources["opnsense_auth_user_api_key.test"].Primary.Attributes
						key = attrs["key"]
						req, _ := http.NewRequest(http.MethodGet, os.Getenv("OPNSENSE_URI")+"/api/core/firmware/status", nil)
						req.SetBasicAuth(attrs["key"], attrs["secret"])
						c := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: os.Getenv("OPNSENSE_ALLOW_INSECURE") == "true"}}}
						resp, err := c.Do(req)
						if err != nil {
							return err
						}
						resp.Body.Close()
						if resp.StatusCode == http.StatusUnauthorized {
							return fmt.Errorf("the new API key was rejected (HTTP 401)")
						}
						return nil
					},
				),
			},
			{
				ResourceName:            "opnsense_auth_user_api_key.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"secret"},
			},
		},
	})
}
