package dhcpv4

import (
	"fmt"
	"net"
	"strings"

	"github.com/u-root/u-root/pkg/uio"
)

// This option implements the router option
// https://tools.ietf.org/html/rfc2132

// IPs are IPv4 addresses from a DHCP packet as used and specified by options
// in RFC 2132, Sections 3.5 through 3.13, 8.2, 8.3, 8.5, 8.6, 8.9, and 8.10.
//
// IPs implements the OptionValue type.
type IPs []net.IP

func parseIPs(data []byte) (IPs, error) {
	buf := uio.NewBigEndianBuffer(data)

	if buf.Len() == 0 {
		return nil, fmt.Errorf("IP DHCP options must always list at least one IP")
	}

	ips := make(IPs, 0, buf.Len()/net.IPv4len)
	for buf.Has(net.IPv4len) {
		ips = append(ips, net.IP(buf.CopyN(net.IPv4len)))
	}
	return ips, buf.FinError()
}

// ToBytes marshals IPv4 addresses to a DHCP packet as specified by RFC 2132,
// Section 3.5 et al.
func (i IPs) ToBytes() []byte {
	buf := uio.NewBigEndianBuffer(nil)

	for _, ip := range i {
		buf.WriteBytes(ip.To4())
	}
	return buf.Data()
}

// String returns a human-readable representation of a list of IPs.
func (i IPs) String() string {
	s := make([]string, 0, len(i))
	for _, ip := range i {
		s = append(s, ip.String())
	}
	return strings.Join(s, ", ")
}

// GetRouter finds and parses the DHCPv4 Router option.
func GetRouter(o Options) []net.IP {
	v := o.Get(OptionRouter)
	if v == nil {
		return nil
	}
	ips, err := parseIPs(v)
	if err != nil {
		return nil
	}
	return []net.IP(ips)
}

// OptRouter returns a DHCPv4 Router option.
func OptRouter(routers ...net.IP) Option {
	return Option{
		Code:  OptionRouter,
		Value: IPs(routers),
	}
}

// WithRouter updates a packet with the DHCPv4 Router option.
func WithRouter(routers ...net.IP) Modifier {
	return func(d *DHCPv4) *DHCPv4 {
		d.UpdateOption(OptRouter(routers...))
		return d
	}
}

// GetNTPServers finds and parses the DHCPv4 NTP Servers option.
func GetNTPServers(o Options) []net.IP {
	v := o.Get(OptionNTPServers)
	if v == nil {
		return nil
	}
	ips, err := parseIPs(v)
	if err != nil {
		return nil
	}
	return []net.IP(ips)
}

// OptNTPServers returns a DHCPv4 NTP Server option.
func OptNTPServers(ntpServers ...net.IP) Option {
	return Option{
		Code:  OptionNTPServers,
		Value: IPs(ntpServers),
	}
}

// GetDNS finds and parses the DHCPv4 Domain Name Server option.
func GetDNS(o Options) []net.IP {
	v := o.Get(OptionDomainNameServer)
	if v == nil {
		return nil
	}
	ips, err := parseIPs(v)
	if err != nil {
		return nil
	}
	return []net.IP(ips)
}

// OptDNS returns a DHCPv4 Domain Name Server option.
func OptDNS(servers ...net.IP) Option {
	return Option{
		Code:  OptionDomainNameServer,
		Value: IPs(servers),
	}
}

// WtihDNS modifies a packet with the DHCPv4 Domain Name Server option.
func WithDNS(servers ...net.IP) Modifier {
	return WithOption(OptDNS(servers...))
}
