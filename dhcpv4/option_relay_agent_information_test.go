package dhcpv4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseOptRelayAgentInformation(t *testing.T) {
	o := Options{
		OptionRelayAgentInformation: []byte{
			1, 5, 'l', 'i', 'n', 'u', 'x',
			2, 4, 'b', 'o', 'o', 't',
		},
	}

	opt := GetRelayAgentInfo(o)
	require.NotNil(t, opt)
	require.Equal(t, len(opt), 2)

	circuit := opt.Get(1)
	remote := opt.Get(2)
	require.Equal(t, circuit, []byte("linux"))
	require.Equal(t, remote, []byte("boot"))
}

func TestParseOptRelayAgentInformationToBytes(t *testing.T) {
	opt := OptRelayAgentInfo(
		Options{
			1: []byte("linux"),
			2: []byte("boot"),
		},
	)
	data := opt.Value.ToBytes()
	expected := []byte{
		1, 5, 'l', 'i', 'n', 'u', 'x',
		2, 4, 'b', 'o', 'o', 't',
	}
	require.Equal(t, expected, data)
}

func TestOptRelayAgentInformationToBytesString(t *testing.T) {
	o := OptRelayAgentInfo(nil)
	require.Equal(t, "Relay Agent Information: ", o.String())
}
