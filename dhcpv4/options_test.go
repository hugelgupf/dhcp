package dhcpv4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseOption(t *testing.T) {
	// Generic
	option := []byte{192, 168, 1, 254} // DNS option
	opt, err := ParseOption(OptionNameServer, option)
	require.NoError(t, err)
	generic := opt.(*OptionGeneric)
	require.Equal(t, OptionNameServer, generic.Code())
	require.Equal(t, []byte{192, 168, 1, 254}, generic.Data)
	require.Equal(t, 4, generic.Length())
	require.Equal(t, "Name Server -> [192 168 1 254]", generic.String())

	// Option subnet mask
	option = []byte{255, 255, 255, 0}
	opt, err = ParseOption(OptionSubnetMask, option)
	require.NoError(t, err)
	require.Equal(t, OptionSubnetMask, opt.Code(), "Code")
	require.Equal(t, 4, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Option router
	option = []byte{192, 168, 1, 1}
	opt, err = ParseOption(OptionRouter, option)
	require.NoError(t, err)
	require.Equal(t, OptionRouter, opt.Code(), "Code")
	require.Equal(t, 4, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Option domain name server
	option = []byte{192, 168, 1, 1}
	opt, err = ParseOption(OptionDomainNameServer, option)
	require.NoError(t, err)
	require.Equal(t, OptionDomainNameServer, opt.Code(), "Code")
	require.Equal(t, 4, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Option host name
	option = []byte{'t', 'e', 's', 't'}
	opt, err = ParseOption(OptionHostName, option)
	require.NoError(t, err)
	require.Equal(t, OptionHostName, opt.Code(), "Code")
	require.Equal(t, 4, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Option domain name
	option = []byte{'t', 'e', 's', 't'}
	opt, err = ParseOption(OptionDomainName, option)
	require.NoError(t, err)
	require.Equal(t, OptionDomainName, opt.Code(), "Code")
	require.Equal(t, 4, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Option root path
	option = []byte{'/', 'f', 'o', 'o'}
	opt, err = ParseOption(OptionRootPath, option)
	require.NoError(t, err)
	require.Equal(t, OptionRootPath, opt.Code(), "Code")
	require.Equal(t, 4, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Option broadcast address
	option = []byte{255, 255, 255, 255}
	opt, err = ParseOption(OptionBroadcastAddress, option)
	require.NoError(t, err)
	require.Equal(t, OptionBroadcastAddress, opt.Code(), "Code")
	require.Equal(t, 4, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Option NTP servers
	option = []byte{10, 10, 10, 10}
	opt, err = ParseOption(OptionNTPServers, option)
	require.NoError(t, err)
	require.Equal(t, OptionNTPServers, opt.Code(), "Code")
	require.Equal(t, 4, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Requested IP address
	option = []byte{1, 2, 3, 4}
	opt, err = ParseOption(OptionRequestedIPAddress, option)
	require.NoError(t, err)
	require.Equal(t, OptionRequestedIPAddress, opt.Code(), "Code")
	require.Equal(t, 4, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Requested IP address lease time
	option = []byte{0, 0, 0, 0}
	opt, err = ParseOption(OptionIPAddressLeaseTime, option)
	require.NoError(t, err)
	require.Equal(t, OptionIPAddressLeaseTime, opt.Code(), "Code")
	require.Equal(t, 4, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Message type
	option = []byte{1}
	opt, err = ParseOption(OptionDHCPMessageType, option)
	require.NoError(t, err)
	require.Equal(t, OptionDHCPMessageType, opt.Code(), "Code")
	require.Equal(t, 1, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Option server ID
	option = []byte{1, 2, 3, 4}
	opt, err = ParseOption(OptionServerIdentifier, option)
	require.NoError(t, err)
	require.Equal(t, OptionServerIdentifier, opt.Code(), "Code")
	require.Equal(t, 4, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Parameter request list
	option = []byte{55, 53, 61}
	opt, err = ParseOption(OptionParameterRequestList, option)
	require.NoError(t, err)
	require.Equal(t, OptionParameterRequestList, opt.Code(), "Code")
	require.Equal(t, 3, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Option max message size
	option = []byte{1, 2}
	opt, err = ParseOption(OptionMaximumDHCPMessageSize, option)
	require.NoError(t, err)
	require.Equal(t, OptionMaximumDHCPMessageSize, opt.Code(), "Code")
	require.Equal(t, 2, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Option class identifier
	option = []byte{'t', 'e', 's', 't'}
	opt, err = ParseOption(OptionClassIdentifier, option)
	require.NoError(t, err)
	require.Equal(t, OptionClassIdentifier, opt.Code(), "Code")
	require.Equal(t, 4, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Option TFTP server name
	option = []byte{'t', 'e', 's', 't'}
	opt, err = ParseOption(OptionTFTPServerName, option)
	require.NoError(t, err)
	require.Equal(t, OptionTFTPServerName, opt.Code(), "Code")
	require.Equal(t, 4, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Option Bootfile name
	option = []byte{'l', 'i', 'n', 'u', 'x', 'b', 'o', 'o', 't'}
	opt, err = ParseOption(OptionBootfileName, option)
	require.NoError(t, err)
	require.Equal(t, OptionBootfileName, opt.Code(), "Code")
	require.Equal(t, 9, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Option user class information
	option = []byte{4, 't', 'e', 's', 't'}
	opt, err = ParseOption(OptionUserClassInformation, option)
	require.NoError(t, err)
	require.Equal(t, OptionUserClassInformation, opt.Code(), "Code")
	require.Equal(t, 5, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Option relay agent information
	option = []byte{1, 4, 129, 168, 0, 1}
	opt, err = ParseOption(OptionRelayAgentInformation, option)
	require.NoError(t, err)
	require.Equal(t, OptionRelayAgentInformation, opt.Code(), "Code")
	require.Equal(t, 6, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")

	// Option client system architecture type option
	option = []byte{'t', 'e', 's', 't'}
	opt, err = ParseOption(OptionClientSystemArchitectureType, option)
	require.NoError(t, err)
	require.Equal(t, OptionClientSystemArchitectureType, opt.Code(), "Code")
	require.Equal(t, 4, opt.Length(), "Length")
	require.Equal(t, option, opt.ToBytes(), "ToBytes")
}
