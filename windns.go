package windns

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	windns "github.com/tony-butchart/libdns-windns"
)

// Provider wraps the provider implementation as a Caddy module.
type Provider struct{ *windns.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.windns",
		New: func() caddy.Module { return &Provider{new(windns.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.Host = caddy.NewReplacer().ReplaceAll(p.Provider.Host, "")
	p.Provider.User = caddy.NewReplacer().ReplaceAll(p.Provider.User, "")
	p.Provider.Password = caddy.NewReplacer().ReplaceAll(p.Provider.Password, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	windns {
//	    host     <windows_dns_server_ip>
//	    user     <username>
//	    password <password>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Provider.Host = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "host":
				if p.Provider.Host != "" {
					return d.Err("Host already set")
				}
				if d.NextArg() {
					p.Provider.Host = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "user":
				if p.Provider.User != "" {
					return d.Err("User already set")
				}
				if d.NextArg() {
					p.Provider.User = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "password":
				if p.Provider.Password != "" {
					return d.Err("Password already set")
				}
				if d.NextArg() {
					p.Provider.Password = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.Host == "" || p.Provider.User == "" || p.Provider.Password == "" {
		return d.Err("missing required fields")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
