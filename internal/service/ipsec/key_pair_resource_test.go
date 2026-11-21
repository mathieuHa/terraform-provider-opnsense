package ipsec_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testKeyPairPEM(t *testing.T) (string, string) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	pub, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		t.Fatal(err)
	}
	return string(pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pub})),
		string(pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}))
}

func TestAccIpsecKeyPairResource(t *testing.T) {
	pub, priv := testKeyPairPEM(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccKeyPairConfig("tfacc", pub, priv),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_ipsec_key_pair.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_ipsec_key_pair.test", "key_type", "rsa"),
					resource.TestCheckResourceAttr("opnsense_ipsec_key_pair.test", "key_size", "2048"),
					resource.TestCheckResourceAttrSet("opnsense_ipsec_key_pair.test", "key_fingerprint"),
				),
			},
			{
				ResourceName:            "opnsense_ipsec_key_pair.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"private_key"},
			},
			{
				Config: testAccKeyPairConfig("tfacc2", pub, priv),
				Check:  resource.TestCheckResourceAttr("opnsense_ipsec_key_pair.test", "name", "tfacc2"),
			},
		},
	})
}

func testAccKeyPairConfig(name, pub, priv string) string {
	return fmt.Sprintf(`
resource "opnsense_ipsec_key_pair" "test" {
  name        = %q
  public_key  = %q
  private_key = %q
}
`, name, pub, priv)
}
