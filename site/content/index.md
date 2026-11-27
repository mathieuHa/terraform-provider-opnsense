# OPNsense provider for OpenTofu and Terraform

Manage your OPNsense firewalls as code: firewall rules and aliases, NAT, VPNs, DHCP, DNS, routing, high availability, intrusion detection and the most used plugins, through the OPNsense API.

<div class="grid cards" markdown>

-   **{{ resources }} resources, {{ data_sources }} data sources**

    ---

    {{ services }} services: {{ core }} built into OPNsense and {{ plugins }} plugins.

    [→ Supported resources](resources.md)

-   **Five minutes to the first rule**

    ---

    Create an API key, point the provider at your firewall and apply an alias, a rule and a DNS override.

    [→ Quickstart](quickstart.md)

-   **Tested recipes**

    ---

    WireGuard between two sites, HAProxy and Caddy reverse proxies, ACME certificates, Kea DHCP.

    [→ Guides](guides/singletons.md)

-   **Signed releases**

    ---

    Published on the OpenTofu registry as `mathieuHa/opnsense`, every release signed and tested against the latest OPNsense 26.7 patch release.

    [→ Releases](https://github.com/mathieuHa/terraform-provider-opnsense/releases)

</div>

## Install

```terraform
terraform {
  required_providers {
    opnsense = {
      source  = "mathieuHa/opnsense"
      version = "~> 0.37"
    }
  }
}

provider "opnsense" {
  uri        = "https://firewall.example.com"
  api_key    = var.opnsense_api_key
  api_secret = var.opnsense_api_secret
}
```

With Terraform, use the full registry address: `source = "registry.opentofu.org/mathieuha/opnsense"`.

## What you can manage

| Area | Examples |
|---|---|
| Firewall | Rules, aliases, categories, NAT (port forward, one-to-one, NPT), groups |
| Interfaces | Assignments, VLANs, LAGG, bridges, GRE, GIF, VXLAN, loopbacks, virtual IPs and CARP, static ARP/NDP |
| VPN | WireGuard, IPsec (connections, children, VTI, pools, keys), OpenVPN, Tailscale, ZeroTier |
| DHCP and DNS | Kea DHCPv4/v6 with DDNS, Dnsmasq, Unbound overrides, forwards and blocklists, router advertisements, DHCP relay |
| Routing | Gateways and groups, static routes, BGP with FRR |
| Services | IDS/Suricata, captive portal, traffic shaper, Monit, cron, syslog, HA sync, certificates |
| Plugins | HAProxy, Caddy, ACME client, Squid, CrowdSec, Telegraf, node_exporter, chrony, NUT — installed with `opnsense_firmware_plugin` |

## About this provider

This is a maintained fork of [browningluke/terraform-provider-opnsense](https://github.com/browningluke/terraform-provider-opnsense). It extends the coverage of the OPNsense API, ships each release with release notes and runs the acceptance suite against real OPNsense instances before tagging. Resources and attributes keep the upstream names, so existing configurations only need a new provider source.
