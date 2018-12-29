package dhcpv4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptParameterRequestListInterfaceMethods(t *testing.T) {
	opts := []OptionCode{OptionBootfileName, OptionNameServer}
	o := OptParameterRequestList(opts...)

	require.Equal(t, OptionParameterRequestList, o.Code, "Code")

	expectedBytes := []byte{67, 5}
	require.Equal(t, expectedBytes, o.Value.ToBytes(), "ToBytes")

	expectedString := "Parameter Request List: Bootfile Name, Name Server"
	require.Equal(t, expectedString, o.String(), "String")
}

func TestParseOptParameterRequestList(t *testing.T) {
	o, err := parseOptParameterRequestList([]byte{67, 5})
	require.NoError(t, err)
	expectedOpts := []OptionCode{OptionBootfileName, OptionNameServer}
	require.Equal(t, expectedOpts, o.RequestedOpts)
}
