terraform {
  required_providers {
    opnsense = {
      source  = "mathieuHa/opnsense"
      version = "~> 0.37"
    }
  }
}

variable "sites" {
  type = map(object({
    uri        = string
    api_key    = string
    api_secret = string
    wan_ip     = string
    tunnel_ip  = string
    lan        = string
    wg_private = string
    wg_public  = string
  }))
  sensitive = true
}

provider "opnsense" {
  alias      = "paris"
  uri        = var.sites.paris.uri
  api_key    = var.sites.paris.api_key
  api_secret = var.sites.paris.api_secret
}

provider "opnsense" {
  alias      = "lyon"
  uri        = var.sites.lyon.uri
  api_key    = var.sites.lyon.api_key
  api_secret = var.sites.lyon.api_secret
}

module "paris" {
  source    = "./site"
  providers = { opnsense = opnsense.paris }
  name      = "paris"
  local     = var.sites.paris
  peer      = var.sites.lyon
}

module "lyon" {
  source    = "./site"
  providers = { opnsense = opnsense.lyon }
  name      = "lyon"
  local     = var.sites.lyon
  peer      = var.sites.paris
}

import {
  to = module.paris.opnsense_wireguard_settings.this
  id = "wireguard_settings"
}

import {
  to = module.lyon.opnsense_wireguard_settings.this
  id = "wireguard_settings"
}
