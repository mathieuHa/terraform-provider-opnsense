# Kea DHCP

Serve DHCP on the LAN with Kea, with a fixed address for a printer.

```terraform
--8<-- "examples/guides/kea/main.tf"
```

- `opnsense_kea_dhcpv4_settings` enables the DHCPv4 server on `lan` (a [singleton](singletons.md)).
- `opnsense_kea_dhcpv4_subnet` defines the range and the options sent to clients. With `auto_collect = false`, the routers, DNS servers and domain come from the configuration instead of the interface settings.
- `opnsense_kea_dhcpv4_reservation` pins an address to a MAC address.

To register clients in DNS, enable `opnsense_kea_ddns_settings` and set the `ddns_*` attributes of the subnet. For two firewalls, add `opnsense_kea_dhcpv4_peer` and the HA attributes of the settings.
