resource "opnsense_firewall_nat_one_to_one" "example_one" {
  external_net = "220.110.81.9/32"
  source = {
    net = "10.10.2.9/32"
  }
  description = "Example one"
}


resource "opnsense_firewall_nat_one_to_one" "example_two" {
  enabled      = false
  log          = true
  external_net = "220.110.81.9/32"
  type         = "binat"

  source = {
    net    = "10.10.2.9/32"
    invert = false
  }

  destination = {
    net    = "any"
    invert = false
  }

  nat_reflection = "enable"
  categories     = [opnsense_firewall_category.nat.id]
  description    = "Example two"
}

resource "opnsense_firewall_nat_one_to_one" "example_three" {
  enabled      = true
  log          = true
  external_net = "220.110.81.9/32"
  type         = "nat"

  source = {
    net = "__lan_network" # aliases are only allowed in type nat rules
  }

  nat_reflection = "default"
  description    = "Example three"
}

resource "opnsense_firewall_category" "nat" {
  name = "nat"
}
