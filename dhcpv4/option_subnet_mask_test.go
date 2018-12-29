package dhcpv4

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptSubnetMaskInterfaceMethods(t *testing.T) {
	mask := net.IPMask{255, 255, 255, 0}
	o := optSubnetMask{mask}
	expectedBytes := []byte{255, 255, 255, 0}
	require.Equal(t, expectedBytes, o.ToBytes(), "ToBytes")
	require.Equal(t, "ffffff00", o.String(), "String")
}

func TestOptSubnetMask(t *testing.T) {
	mask := net.IPMask{255, 255, 255, 0}
	o := OptSubnetMask(mask)
	require.Equal(t, o.Code, OptionSubnetMask, "Code")
	require.Equal(t, "Subnet Mask: ffffff00", o.String(), "String")
}

func TestParseOptSubnetMask(t *testing.T) {
	o, err := parseOptSubnetMask([]byte{})
	require.Error(t, err, "empty byte stream")

	o, err = parseOptSubnetMask([]byte{255})
	require.Error(t, err, "short byte stream")

	o, err = parseOptSubnetMask([]byte{255, 255, 255, 0})
	require.NoError(t, err)
	require.Equal(t, net.IPMask{255, 255, 255, 0}, o.SubnetMask)
}
