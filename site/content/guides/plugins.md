# Plugins

OPNsense plugins add services such as HAProxy or Caddy. The provider can install them and manage their configuration.

## Installing a plugin

`opnsense_firmware_plugin` installs a package and waits for the firmware job to finish. Make the plugin's resources depend on it:

```terraform
resource "opnsense_firmware_plugin" "haproxy" {
  name = "os-haproxy"

  lifecycle {
    prevent_destroy = true
  }
}

resource "opnsense_haproxy_server" "app1" {
  name       = "app1"
  address    = "10.0.10.21"
  port       = 8080
  depends_on = [opnsense_firmware_plugin.haproxy]
}
```

- Destroying `opnsense_firmware_plugin` **removes the plugin**, and its configuration stops being applied. `prevent_destroy` protects against doing it by accident.
- If the plugin is already installed, creating the resource only records it in the state.
- OPNsense refuses to install plugins while the firmware itself has pending updates. The resource then fails with OPNsense's message: update the firewall first.
- The firewall needs access to its package mirror; `opnsense_firmware_settings` manages the mirror.

## Supported plugins

{{ plugins_table }}

When the plugin is not installed, its resources fail with `status code 404`.

## Guides

- [HAProxy reverse proxy](haproxy.md)
- [Caddy reverse proxy](caddy.md)
- [ACME certificates](acme.md)
