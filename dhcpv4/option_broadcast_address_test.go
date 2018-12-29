package dhcpv4

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptBroadcastAddressInterfaceMethods(t *testing.T) {
	ip := net.IP{192, 168, 0, 1}
	o := OptBroadcastAddress{BroadcastAddress: ip}

	require.Equal(t, OptionBroadcastAddress, o.Code(), "Code")

	expectedBytes := []byte{192, 168, 0, 1}
	require.Equal(t, expectedBytes, o.ToBytes(), "ToBytes")

	require.Equal(t, 4, o.Length(), "Length")

	require.Equal(t, "Broadcast Address -> 192.168.0.1", o.String(), "String")
}

func TestParseOptBroadcastAddress(t *testing.T) {
	var (
		o   *OptBroadcastAddress
		err error
	)
	o, err = ParseOptBroadcastAddress([]byte{})
	require.Error(t, err, "empty byte stream")

	o, err = ParseOptBroadcastAddress([]byte{192, 168, 0})
	require.Error(t, err, "wrong IP length")

	o, err = ParseOptBroadcastAddress([]byte{192, 168, 0, 1})
	require.NoError(t, err)
	require.Equal(t, net.IP{192, 168, 0, 1}, o.BroadcastAddress)
}
