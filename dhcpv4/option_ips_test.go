package dhcpv4

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseIPs(t *testing.T) {
	data := []byte{
		192, 168, 0, 10, // DNS #1
		192, 168, 0, 20, // DNS #2
	}
	o, err := parseIPs(data)
	require.NoError(t, err)
	servers := []net.IP{
		net.IP{192, 168, 0, 10},
		net.IP{192, 168, 0, 20},
	}
	require.Equal(t, servers, []net.IP(o))

	// Bad length
	data = []byte{1, 1, 1}
	_, err = parseIPs(data)
	require.Error(t, err, "should get error from bad length")

	// RFC2132 requires that at least one IP is specified for each IP field.
	_, err = parseIPs([]byte{})
	require.Error(t, err)
}

func TestOptDomainNameServerString(t *testing.T) {
	o := OptDNS(net.IPv4(192, 168, 0, 1), net.IPv4(192, 168, 0, 10))
	require.Equal(t, "Domain Name Server: 192.168.0.1, 192.168.0.10", o.String())
}

func TestOptNTPServersString(t *testing.T) {
	o := OptNTPServers(net.IPv4(192, 168, 0, 1), net.IPv4(192, 168, 0, 10))
	require.Equal(t, "NTP Servers: 192.168.0.1, 192.168.0.10", o.String())
}

func TestOptRouterString(t *testing.T) {
	o := OptRouter(net.IP{192, 168, 0, 1}, net.IP{192, 168, 0, 10})
	require.Equal(t, "Router: 192.168.0.1, 192.168.0.10", o.String())
}
