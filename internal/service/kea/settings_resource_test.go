package kea_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func singletonSteps(name, id, set, restore string, checks, restoreChecks resource.TestCheckFunc) []resource.TestStep {
	addr := "opnsense_" + name + ".test"
	return []resource.TestStep{
		{
			Config:             `resource "opnsense_` + name + `" "test" {}`,
			ResourceName:       addr,
			ImportState:        true,
			ImportStateId:      id,
			ImportStatePersist: true,
		},
		{
			Config: `resource "opnsense_` + name + `" "test" {` + set + `}`,
			Check:  checks,
		},
		{
			Config: `resource "opnsense_` + name + `" "test" {` + restore + `}`,
			Check:  restoreChecks,
		},
	}
}

func TestAccKeaCtrlAgentSettingsResource(t *testing.T) {
	a := "opnsense_kea_ctrl_agent_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: singletonSteps("kea_ctrl_agent_settings", "kea_ctrl_agent_settings",
			"\n  enabled = false\n  http_host = \"127.0.0.1\"\n  http_port = 8001\n",
			"\n  enabled = false\n  http_host = \"127.0.0.1\"\n  http_port = 8000\n",
			resource.TestCheckResourceAttr(a, "http_port", "8001"),
			resource.TestCheckResourceAttr(a, "http_port", "8000")),
	})
}

func TestAccKeaDdnsSettingsResource(t *testing.T) {
	a := "opnsense_kea_ddns_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: singletonSteps("kea_ddns_settings", "kea_ddns_settings",
			"\n  enabled = true\n  server_ip = \"127.0.0.1\"\n  server_port = 53002\n",
			"\n  enabled = false\n  manual_config = false\n  server_ip = \"127.0.0.1\"\n  server_port = 53001\n",
			resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(a, "enabled", "true"),
				resource.TestCheckResourceAttr(a, "server_port", "53002"),
				resource.TestCheckResourceAttr(a, "manual_config", "false")),
			resource.TestCheckResourceAttr(a, "enabled", "false")),
	})
}

func TestAccKeaDhcpv4SettingsResource(t *testing.T) {
	a := "opnsense_kea_dhcpv4_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: singletonSteps("kea_dhcpv4_settings", "kea_dhcpv4_settings",
			"\n  enabled = false\n  valid_lifetime = 7200\n  dhcp_socket_type = \"udp\"\n  compatibility = [\"lenient-option-parsing\"]\n  max_reclaim_leases = 50\n  ha_max_unacked_clients = 3\n  ha_this_server_name = \"tfacc\"\n",
			"\n  enabled = false\n  valid_lifetime = 4000\n  dhcp_socket_type = \"raw\"\n  compatibility = []\n  max_reclaim_leases = -1\n  ha_max_unacked_clients = 2\n  ha_this_server_name = \"\"\n",
			resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(a, "valid_lifetime", "7200"),
				resource.TestCheckResourceAttr(a, "dhcp_socket_type", "udp"),
				resource.TestCheckTypeSetElemAttr(a, "compatibility.*", "lenient-option-parsing"),
				resource.TestCheckResourceAttr(a, "max_reclaim_leases", "50"),
				resource.TestCheckResourceAttr(a, "ha_this_server_name", "tfacc"),
				resource.TestCheckResourceAttr(a, "firewall_rules", "true")),
			resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(a, "valid_lifetime", "4000"),
				resource.TestCheckResourceAttr(a, "compatibility.#", "0"),
				resource.TestCheckResourceAttr(a, "max_reclaim_leases", "-1"))),
	})
}

func TestAccKeaDhcpv6SettingsResource(t *testing.T) {
	a := "opnsense_kea_dhcpv6_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: singletonSteps("kea_dhcpv6_settings", "kea_dhcpv6_settings",
			"\n  enabled = false\n  valid_lifetime = 9000\n  mac_sources = [\"duid\", \"ipv6-link-local\"]\n  hold_reclaimed_time = 600\n",
			"\n  enabled = false\n  valid_lifetime = 4000\n  mac_sources = [\"ipv6-link-local\"]\n  hold_reclaimed_time = -1\n",
			resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(a, "valid_lifetime", "9000"),
				resource.TestCheckResourceAttr(a, "mac_sources.#", "2"),
				resource.TestCheckResourceAttr(a, "hold_reclaimed_time", "600")),
			resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(a, "mac_sources.#", "1"),
				resource.TestCheckResourceAttr(a, "hold_reclaimed_time", "-1"))),
	})
}
