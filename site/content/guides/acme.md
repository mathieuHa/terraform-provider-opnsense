# ACME certificates

Issue a Let's Encrypt certificate for the firewall with a DNS-01 challenge at Cloudflare, and restart the web GUI whenever it is renewed.

```terraform
--8<-- "examples/guides/acme/main.tf"
```

- `opnsense_acmeclient_account` registers with the CA on first use.
- `opnsense_acmeclient_validation` holds the challenge: `method = "dns01"`, a `dns_service` from the acme.sh providers and that provider's attributes (`dns_cf_token` for Cloudflare). The credentials are sensitive and never shown in plans.
- `opnsense_acmeclient_certificate` requests the certificate with its alternative names and runs the automations listed in `restart_actions` after each renewal.

The ACME client issues and renews the certificate itself; the provider manages the configuration. Issued certificates appear in **System ▸ Trust ▸ Certificates**, ready for the web GUI, HAProxy or OpenVPN.
