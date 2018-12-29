package dhcpv4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptMessageType(t *testing.T) {
	o := OptMessageType(MessageTypeDiscover)
	require.Equal(t, OptionDHCPMessageType, o.Code, "Code")
	require.Equal(t, []byte{1}, o.Value.ToBytes(), "ToBytes")
	require.Equal(t, "DHCP Message Type: DISCOVER", o.String())
}

func TestParseOptMessageType(t *testing.T) {
	data := []byte{1} // DISCOVER
	mt, err := parseOptMessageType(data)
	require.NoError(t, err)
	require.Equal(t, MessageTypeDiscover, mt.MessageType)

	// Bad length
	data = []byte{1, 2}
	_, err = parseOptMessageType(data)
	require.Error(t, err, "should get error from bad length")
}
