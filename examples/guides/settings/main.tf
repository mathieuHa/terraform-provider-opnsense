terraform {
  required_providers {
    opnsense = {
      source  = "mathieuHa/opnsense"
      version = "~> 0.37"
    }
  }
}

provider "opnsense" {}

import {
  to = opnsense_interfaces_settings.this
  id = "interfaces_settings"
}

resource "opnsense_interfaces_settings" "this" {
  disable_checksum_offloading      = true
  disable_segmentation_offloading  = true
  disable_large_receive_offloading = true
  dhcp6_ra_timeout                 = 10
}
