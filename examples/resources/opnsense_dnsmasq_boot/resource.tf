resource "opnsense_dnsmasq_boot" "pxe" {
  tags           = [opnsense_dnsmasq_tag.pxe.id]
  filename       = "ipxe.efi"
  server_address = "192.168.1.10"
}
