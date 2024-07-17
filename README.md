Windows DNS module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with Windows DNS

## Caddy module name

```
dns.providers.windns
```

## Config examples

Usage with the Caddyfile:

```
# one site
tls {
    dns windns {
        host     192.168.1.100
        user     Administrator
        password SecurePassword123!
        zone     example.com
    }
}
```
