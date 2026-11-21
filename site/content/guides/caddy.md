# Caddy reverse proxy

Caddy obtains and renews certificates automatically. This guide publishes two internal services, one of them behind basic authentication.

```terraform
--8<-- "examples/guides/caddy/main.tf"
```

- `opnsense_caddy_settings` turns Caddy on; `tls_email` is the address used for the certificates of Let's Encrypt or ZeroSSL.
- Each `opnsense_caddy_reverse_proxy` is a domain; each `opnsense_caddy_handle` sends its requests, optionally only some paths, to upstream servers.
- `opnsense_caddy_basic_auth` users protect a handler through its `basicauth` list.

Caddy needs ports 80 and 443 on the firewall for automatic certificates, so move the web GUI to other ports first (**System ▸ Settings ▸ Administration**). The domains must resolve to the firewall and the ports must be reachable from the internet, or use a DNS provider (`tls_dns_provider` in the settings and `dns_challenge` on the domain).
