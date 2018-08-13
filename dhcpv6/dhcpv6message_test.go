package dhcpv6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsNetboot(t *testing.T) {
	msg1 := DHCPv6Message{}
	require.False(t, msg1.IsNetboot())

	msg2 := DHCPv6Message{}
	optro := OptRequestedOption{}
	optro.AddRequestedOption(OptionBootfileURL)
	msg2.AddOption(&optro)
	require.True(t, msg2.IsNetboot())

	msg3 := DHCPv6Message{}
	optbf := OptBootFileURL{}
	msg3.AddOption(&optbf)
	require.True(t, msg3.IsNetboot())
}

func TestIsOptionRequested(t *testing.T) {
	msg1 := DHCPv6Message{}
	require.False(t, msg1.IsOptionRequested(OptionDNSRecursiveNameServer))

	msg2 := DHCPv6Message{}
	optro := OptRequestedOption{}
	optro.AddRequestedOption(OptionDNSRecursiveNameServer)
	msg2.AddOption(&optro)
	require.True(t, msg2.IsOptionRequested(OptionDNSRecursiveNameServer))
}
