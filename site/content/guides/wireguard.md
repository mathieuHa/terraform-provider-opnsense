# WireGuard between two sites

Connect two OPNsense firewalls with a WireGuard tunnel and let their LANs talk to each other, managing both firewalls from one configuration.

## Layout

| | Paris | Lyon |
|---|---|---|
| WAN address | `wan_ip` of `paris` | `wan_ip` of `lyon` |
| Tunnel address | `10.99.0.1` | `10.99.0.2` |
| LAN | `10.1.0.0/24` | `10.2.0.0/24` |

## Configuration

One provider per firewall, and a small module applied to each:

=== "main.tf"

    ```terraform
    --8<-- "examples/guides/wireguard-site-to-site/main.tf"
    ```

=== "site/main.tf"

    ```terraform
    --8<-- "examples/guides/wireguard-site-to-site/site/main.tf"
    ```

=== "terraform.tfvars"

    ```terraform
    sites = {
      paris = {
        uri        = "https://fw-paris.example.com"
        api_key    = "..."
        api_secret = "..."
        wan_ip     = "203.0.113.10"
        tunnel_ip  = "10.99.0.1"
        lan        = "10.1.0.0/24"
        wg_private = "..."
        wg_public  = "..."
      }
      lyon = {
        uri        = "https://fw-lyon.example.com"
        api_key    = "..."
        api_secret = "..."
        wan_ip     = "198.51.100.20"
        tunnel_ip  = "10.99.0.2"
        lan        = "10.2.0.0/24"
        wg_private = "..."
        wg_public  = "..."
      }
    }
    ```

Generate one key pair per site with `wg genkey | tee private.key | wg pubkey > public.key`, and keep the variables file out of version control (or feed it from your secret store).

## What it does

On each firewall:

- a WAN rule accepts WireGuard (UDP 51820) from the other site only,
- `opnsense_wireguard_client` describes the other site: its public key, endpoint, and the addresses allowed through the tunnel (its tunnel address and LAN),
- `opnsense_wireguard_server` is the local tunnel instance with its keys and tunnel address,
- `opnsense_wireguard_settings` turns WireGuard on (a [singleton](singletons.md), imported from the root module),
- an alias holds the other site's networks and a floating rule lets them in.

After `tofu apply`, the tunnel comes up and each firewall reaches the other's tunnel address and LAN.
