# Quickstart

This walks you from an empty directory to a firewall alias, a filter rule and a DNS override applied on your OPNsense, in about five minutes.

## Requirements

- [OpenTofu](https://opentofu.org/docs/intro/install/) 1.6 or later, or Terraform 1.5 or later.
- OPNsense 26.x. Each release is tested against the latest OPNsense 26.7 patch release (26.7.10 for v0.40.1); some resources use API endpoints that only exist in recent releases.
- Network access from where you run `tofu` to the OPNsense web GUI (HTTPS).

## 1. Create an API key

1. In OPNsense, open **System ▸ Access ▸ Users** and add a user, for example `automation`. Give it the privileges it needs: **GUI - All pages** is the simplest, or narrower pages if you only manage some services.
2. Edit the user and, under **API keys**, click **+**. Your browser downloads `apikey.txt` with a `key` and a `secret`.

Keep the key and secret in your secret store: anyone holding them can change the firewall.

## 2. Configure the provider

The provider reads its settings from the environment, so nothing secret ends up in your configuration:

```bash
export OPNSENSE_URI="https://192.168.1.1"
export OPNSENSE_API_KEY="<key from apikey.txt>"
export OPNSENSE_API_SECRET="<secret from apikey.txt>"
export OPNSENSE_ALLOW_INSECURE=true   # only while the GUI uses its self-signed certificate
```

All settings can also be set in the `provider` block; see the [provider configuration](reference/provider.md).

## 3. Write the configuration

Create `main.tf`:

```terraform
--8<-- "examples/guides/quickstart/main.tf"
```

It creates:

- an alias `admins` holding the admin network,
- a rule on `lan` that lets that alias reach the web GUI (`(self)` is *This Firewall*) over HTTPS, with logging,
- an Unbound override resolving `nas.home.example` to `10.0.10.20`.

## 4. Apply

```bash
tofu init
tofu plan
tofu apply
```

`tofu init` downloads the provider from the OpenTofu registry and checks its signature. `tofu apply` creates the objects and applies them on the firewall: the rule is active and the DNS override answers as soon as the apply finishes. Run `tofu plan` again and it reports no changes.

!!! warning "Do not lock yourself out"

    Rules apply immediately. Before managing rules on the interface you reach the firewall through, keep the built-in anti-lockout rule enabled, test on another interface first, and keep console access to the firewall.

## 5. Clean up

```bash
tofu destroy
```

## Next steps

- [Settings resources and imports](guides/singletons.md): manage global settings and bring existing objects under management.
- [Plugins](guides/plugins.md): install and configure HAProxy, Caddy, the ACME client and others.
- [Supported resources](resources.md): everything the provider can manage.

??? note "Using Terraform instead of OpenTofu"

    The provider is published on the OpenTofu registry. Terraform installs it from there when the source uses the full address:

    ```terraform
    terraform {
      required_providers {
        opnsense = {
          source  = "registry.opentofu.org/mathieuha/opnsense"
          version = "~> 0.37"
        }
      }
    }
    ```

    Everything else in this documentation is the same with `terraform` in place of `tofu`.
