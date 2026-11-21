# HAProxy reverse proxy

Install HAProxy, balance an application over two servers with a health check, and route requests by host name.

```terraform
--8<-- "examples/guides/haproxy/main.tf"
```

How the pieces fit:

- `opnsense_haproxy_server`: the real servers, created with `for_each`.
- `opnsense_haproxy_backend`: the pool, linking the servers and the health check, balanced round-robin.
- `opnsense_haproxy_acl` (a *condition* in the GUI) matches the `Host` header; `opnsense_haproxy_action` (a *rule*) sends matching requests to the pool.
- `opnsense_haproxy_frontend` listens on port 80, applies the rule and falls back to a backend without servers, so other host names get a `503`.
- `opnsense_haproxy_settings` turns the service on.

Listening on ports 80 or 443 on an interface where the web GUI also listens requires moving the GUI to other ports first (**System ▸ Settings ▸ Administration**).

For HTTPS, set `ssl_enabled = true` and `ssl_certificates` on the frontend, with certificates from `opnsense_trust_cert` or the [ACME client](acme.md).
