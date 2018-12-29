package dhcpv4

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestOptIPAddressLeaseTimeInterfaceMethods(t *testing.T) {
	o := OptIPAddressLeaseTime(43200 * time.Second)
	require.Equal(t, OptionIPAddressLeaseTime, o.Code, "Code")
	require.Equal(t, []byte{0, 0, 168, 192}, o.Value.ToBytes(), "ToBytes")
}

func TestParseOptIPAddressLeaseTime(t *testing.T) {
	data := []byte{0, 0, 168, 192}
	o, err := parseOptIPAddressLeaseTime(data)
	require.NoError(t, err)
	require.Equal(t, 43200*time.Second, o.LeaseTime)

	// Short byte stream
	data = []byte{168, 192}
	_, err = parseOptIPAddressLeaseTime(data)
	require.Error(t, err, "should get error from short byte stream")

	// Bad length
	data = []byte{1, 1, 1, 1, 1}
	_, err = parseOptIPAddressLeaseTime(data)
	require.Error(t, err, "should get error from bad length")
}

func TestOptIPAddressLeaseTimeString(t *testing.T) {
	o := OptIPAddressLeaseTime(43200 * time.Second)
	require.Equal(t, "IP Addresses Lease Time: 12h0m0s", o.String())
}
