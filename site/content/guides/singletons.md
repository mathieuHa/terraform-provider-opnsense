# Settings resources and imports

Most resources create, update and delete list items (rules, aliases, servers…). OPNsense also has global settings pages that always exist and cannot be created or deleted: the provider exposes them as **singleton** resources.

## Managing a settings page

A singleton has to be imported before you manage it. Use an `import` block with its fixed ID:

```terraform
--8<-- "examples/guides/settings/main.tf"
```

On the first `tofu apply`, OpenTofu imports the current settings, then changes only what differs from your configuration.

Rules that apply to every singleton:

- **Attributes you leave out keep their current value.** You can manage one setting without listing the others.
- **Numbers that are not set on the firewall read back as `-1`** in the newer settings resources; set `-1` to clear a value.
- **`tofu destroy` only removes the resource from the state.** The settings on the firewall stay as they are.

!!! tip "Start from the current values"

    Put only the `import` block in your configuration and let OpenTofu write the resource for you:

    ```bash
    tofu plan -generate-config-out=generated.tf
    ```

    `generated.tf` then contains every attribute with its current value, and the next plan is a pure import.

## Singleton import IDs

{{ singletons_table }}

## Importing existing objects

List items (rules, aliases, servers…) are imported by their UUID. The search endpoint of each API returns it in the `uuid` field of every row, for example `GET /api/firewall/alias/search_item`:

```terraform
import {
  to = opnsense_firewall_alias.legacy
  id = "5c1d3b0e-8f7a-4f0c-9a51-1f2e3d4c5b6a"
}

resource "opnsense_firewall_alias" "legacy" {
  name    = "legacy_servers"
  type    = "host"
  content = ["10.0.20.10", "10.0.20.11"]
}
```

`tofu plan -generate-config-out=generated.tf` works here too and writes the configuration of every imported object.
