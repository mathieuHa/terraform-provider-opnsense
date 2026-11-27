package auth_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/plancheck"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAuthUserResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "resource \"opnsense_auth_group\" \"g\" {\n  name = \"tftestusergrp\"\n}\n\nresource \"opnsense_auth_user\" \"test\" {\n  name = \"tftestuser\"\n  password = \"Tf-Test-Passw0rd!\"\n  full_name = \"TF Test\"\n  email = \"tf@example.com\"\n  group_memberships = [opnsense_auth_group.g.gid]\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_auth_user.test", "name", "tftestuser"),
					resource.TestCheckResourceAttr("opnsense_auth_user.test", "password", "Tf-Test-Passw0rd!"),
					resource.TestCheckResourceAttr("opnsense_auth_user.test", "full_name", "TF Test"),
					resource.TestCheckResourceAttr("opnsense_auth_user.test", "email", "tf@example.com"),
					resource.TestCheckResourceAttrSet("opnsense_auth_user.test", "id"),
				),
			},
			{
				ResourceName:            "opnsense_auth_user.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password", "scrambled_password"},
			},
			{
				Config: "resource \"opnsense_auth_group\" \"g\" {\n  name = \"tftestusergrp\"\n}\n\nresource \"opnsense_auth_user\" \"test\" {\n  name = \"tftestuser\"\n  password = \"Tf-Test-Passw0rd-2!\"\n  full_name = \"TF Test 2\"\n  shell = \"/bin/sh\"\n  privileges = [\"page-diagnostics-ping\"]\n  group_memberships = []\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_auth_user.test", "name", "tftestuser"),
					resource.TestCheckResourceAttr("opnsense_auth_user.test", "password", "Tf-Test-Passw0rd-2!"),
					resource.TestCheckResourceAttr("opnsense_auth_user.test", "full_name", "TF Test 2"),
					resource.TestCheckResourceAttr("opnsense_auth_user.test", "shell", "/bin/sh"),
					resource.TestCheckResourceAttr("opnsense_auth_user.test", "privileges.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_auth_user.test", "privileges.*", "page-diagnostics-ping"),
				),
			},
			{
				Config: "resource \"opnsense_auth_group\" \"g\" {\n  name = \"tftestusergrp\"\n}\n\nresource \"opnsense_auth_user\" \"test\" {\n  name = \"tftestuser\"\n  password = \"Tf-Test-Passw0rd-2!\"\n  full_name = \"TF Test 2\"\n  shell = \"/bin/sh\"\n  privileges = [\"page-diagnostics-ping\"]\n  group_memberships = []\n}\n\ndata \"opnsense_auth_user\" \"test\" {\n  id = opnsense_auth_user.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_auth_user.test", "name", "opnsense_auth_user.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_auth_user.test", "full_name", "opnsense_auth_user.test", "full_name"),
					resource.TestCheckResourceAttrPair("data.opnsense_auth_user.test", "shell", "opnsense_auth_user.test", "shell"),
				),
			},
		},
	})
}

func TestAccAuthUserImportKeepsPassword(t *testing.T) {
	user := func(scrambled string) string {
		return "resource \"opnsense_auth_user\" \"test\" {\n  name = \"tftestimport\"\n  password = \"Tf-Test-Passw0rd!\"\n" + scrambled + "}\n"
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: user(""),
			},
			{
				ResourceName:    "opnsense_auth_user.test",
				Config:          "resource \"opnsense_auth_user\" \"test\" {\n  name = \"tftestimport\"\n  scrambled_password = true\n}\n",
				ImportState:     true,
				ImportStateKind: resource.ImportBlockWithID,
				ImportPlanChecks: resource.ImportPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction("opnsense_auth_user.test", plancheck.ResourceActionNoop)},
				},
			},
		},
	})
}
