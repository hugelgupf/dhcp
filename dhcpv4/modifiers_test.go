package dhcpv4

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTransactionIDModifier(t *testing.T) {
	d, err := New()
	require.NoError(t, err)
	WithTransactionID(TransactionID{0xdd, 0xcc, 0xbb, 0xaa})(d)
	require.Equal(t, TransactionID{0xdd, 0xcc, 0xbb, 0xaa}, d.TransactionID)
}

func TestBroadcastModifier(t *testing.T) {
	d, err := New()
	require.NoError(t, err)

	// set and test broadcast
	WithBroadcast(true)(d)
	require.Equal(t, true, d.IsBroadcast())

	// set and test unicast
	WithBroadcast(false)(d)
	require.Equal(t, true, d.IsUnicast())
}

func TestHwAddrModifier(t *testing.T) {
	hwaddr := net.HardwareAddr{0xa, 0xb, 0xc, 0xd, 0xe, 0xf}
	d, err := New(WithHwAddr(hwaddr))
	require.NoError(t, err)
	require.Equal(t, hwaddr, d.ClientHWAddr)
}

func TestWithOptionModifier(t *testing.T) {
	d, err := New(WithOption(OptDomainName("slackware.it")))
	require.NoError(t, err)

	dnOpt := GetDomainName(d.Options)
	require.NotNil(t, dnOpt)
	require.Equal(t, "slackware.it", dnOpt)
}

func TestUserClassModifier(t *testing.T) {
	d, err := New(WithUserClass([]byte("linuxboot"), false))
	require.NoError(t, err)

	expected := []byte{
		'l', 'i', 'n', 'u', 'x', 'b', 'o', 'o', 't',
	}
	require.Equal(t, expected, d.GetOneOption(OptionUserClassInformation))
}

func TestUserClassModifierRFC(t *testing.T) {
	d, err := New(WithUserClass([]byte("linuxboot"), true))
	require.NoError(t, err)

	expected := []byte{
		9, 'l', 'i', 'n', 'u', 'x', 'b', 'o', 'o', 't',
	}
	require.Equal(t, expected, d.GetOneOption(OptionUserClassInformation))
}

func TestWithNetboot(t *testing.T) {
	d, err := New(WithNetboot)
	require.NoError(t, err)

	require.Equal(t, "TFTP Server Name, Bootfile Name", GetParameterRequestList(d.Options).String())
}

func TestWithNetbootExistingTFTP(t *testing.T) {
	d, _ := New()
	d.UpdateOption(OptParameterRequestList(OptionTFTPServerName))
	WithNetboot(d)
	require.Equal(t, "TFTP Server Name, Bootfile Name", GetParameterRequestList(d.Options).String())
}

func TestWithNetbootExistingBootfileName(t *testing.T) {
	d, _ := New()
	d.UpdateOption(OptParameterRequestList(OptionBootfileName))
	WithNetboot(d)
	require.Equal(t, "TFTP Server Name, Bootfile Name", GetParameterRequestList(d.Options).String())
}

func TestWithNetbootExistingBoth(t *testing.T) {
	d, _ := New()
	d.UpdateOption(OptParameterRequestList(OptionBootfileName, OptionTFTPServerName))
	WithNetboot(d)
	require.Equal(t, "TFTP Server Name, Bootfile Name", GetParameterRequestList(d.Options).String())
}

func TestWithRequestedOptions(t *testing.T) {
	// Check if OptionParameterRequestList is created when not present
	d, err := New(WithRequestedOptions(OptionFQDN))
	require.NoError(t, err)
	require.NotNil(t, d)

	opts := GetParameterRequestList(d.Options)
	require.NotNil(t, opts)
	require.ElementsMatch(t, opts, []OptionCode{OptionFQDN})
	// Check if already set options are preserved
	WithRequestedOptions(OptionHostName)(d)
	require.NotNil(t, d)
	opts = GetParameterRequestList(d.Options)
	require.NotNil(t, opts)
	require.ElementsMatch(t, opts, []OptionCode{OptionFQDN, OptionHostName})
}

func TestWithRelay(t *testing.T) {
	ip := net.IP{10, 0, 0, 1}
	d, err := New(WithRelay(ip))
	require.NoError(t, err)

	require.True(t, d.IsUnicast(), "expected unicast")
	require.Equal(t, ip, d.GatewayIPAddr)
	require.Equal(t, uint8(1), d.HopCount)
}

func TestWithNetmask(t *testing.T) {
	d, err := New(WithNetmask(net.IPv4Mask(255, 255, 255, 0)))
	require.NoError(t, err)

	osm := GetSubnetMask(d.Options)
	require.Equal(t, net.IPv4Mask(255, 255, 255, 0), osm)
}

func TestWithLeaseTime(t *testing.T) {
	d, err := New(WithLeaseTime(uint32(3600)))
	require.NoError(t, err)

	require.True(t, d.Options.Has(OptionIPAddressLeaseTime))
	olt := GetIPAddressLeaseTime(d.Options, 10*time.Second)
	require.Equal(t, 3600*time.Second, olt)
}

func TestWithDNS(t *testing.T) {
	d, err := New(WithDNS(net.ParseIP("10.0.0.1"), net.ParseIP("10.0.0.2")))
	require.NoError(t, err)

	dns := GetDNS(d.Options)
	require.Equal(t, net.ParseIP("10.0.0.1").To4(), dns[0])
	require.Equal(t, net.ParseIP("10.0.0.2").To4(), dns[1])
}

func TestWithDomainSearchList(t *testing.T) {
	d, err := New(WithDomainSearchList("slackware.it", "dhcp.slackware.it"))
	require.NoError(t, err)

	osl := GetDomainSearch(d.Options)
	require.NotNil(t, osl)
	require.Equal(t, 2, len(osl.Labels))
	require.Equal(t, "slackware.it", osl.Labels[0])
	require.Equal(t, "dhcp.slackware.it", osl.Labels[1])
}

func TestWithRouter(t *testing.T) {
	rtr := net.ParseIP("10.0.0.254").To4()
	d, err := New(WithRouter(rtr))
	require.NoError(t, err)

	ortr := GetRouter(d.Options)
	require.Equal(t, rtr, ortr[0])
}
