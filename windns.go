package windns

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	winddns "github.com/tony-butchart/libdns-windns"
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
	repl := caddy.NewReplacer()
	p.Provider.Host = repl.ReplaceAll(p.Provider.Host, "")
	p.Provider.User = repl.ReplaceAll(p.Provider.User, "")
	p.Provider.Password = repl.ReplaceAll(p.Provider.Password, "")
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
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "host":
				if p.Provider.Host != "" {
					return d.Err("Host already set")
				}
				p.Provider.Host = d.Val()
			case "user":
				if p.Provider.User != "" {
					return d.Err("User already set")
				}
				p.Provider.User = d.Val()
			case "password":
				if p.Provider.Password != "" {
					return d.Err("Password already set")
				}
				p.Provider.Password = d.Val()
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
