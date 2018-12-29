package dhcpv4

import (
	"net"

	"github.com/u-root/u-root/pkg/uio"
)

// This option implements the server identifier option
// https://tools.ietf.org/html/rfc2132
type optIP struct {
	IP net.IP
}

func parseIP(data []byte) (*optIP, error) {
	buf := uio.NewBigEndianBuffer(data)
	return &optIP{net.IP(buf.CopyN(net.IPv4len))}, buf.FinError()
}

// ToBytes returns a serialized stream of bytes for this option.
func (o *optIP) ToBytes() []byte {
	return []byte(o.IP.To4())
}

// String returns a human-readable string.
func (o *optIP) String() string {
	return o.IP.String()
}

// GetBroadcastAddress returns the DHCPv4 Broadcast Address value in o.
func GetBroadcastAddress(o Options) net.IP {
	v := o.Get(OptionBroadcastAddress)
	if v == nil {
		return nil
	}
	ip, err := parseIP(v)
	if err != nil {
		return nil
	}
	return ip.IP
}

// OptBroadcastAddress returns a new DHCPv4 Broadcast Address option.
func OptBroadcastAddress(ip net.IP) Option {
	return Option{Code: OptionBroadcastAddress, Value: &optIP{ip}}
}

// GetRequestedIPAddress returns the DHCPv4 Requested IP Address value in o.
func GetRequestedIPAddress(o Options) net.IP {
	v := o.Get(OptionRequestedIPAddress)
	if v == nil {
		return nil
	}
	ip, err := parseIP(v)
	if err != nil {
		return nil
	}
	return ip.IP
}

// OptRequestedIPAddress returns a new DHCPv4 Requested IP Address option.
func OptRequestedIPAddress(ip net.IP) Option {
	return Option{Code: OptionRequestedIPAddress, Value: &optIP{ip}}
}

// GetServerIdentifier returns the DHCPv4 Server Identifier value in o.
func GetServerIdentifier(o Options) net.IP {
	v := o.Get(OptionServerIdentifier)
	if v == nil {
		return nil
	}
	ip, err := parseIP(v)
	if err != nil {
		return nil
	}
	return ip.IP
}

// OptServerIdentifier returns a new DHCPv4 Server Identifier option.
func OptServerIdentifier(ip net.IP) Option {
	return Option{Code: OptionServerIdentifier, Value: &optIP{ip}}
}
