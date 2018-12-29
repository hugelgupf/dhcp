package dhcpv4

import (
	"testing"

	"github.com/insomniacslk/dhcp/iana"
	"github.com/stretchr/testify/require"
)

func TestParseOptClientArchType(t *testing.T) {
	data := []byte{
		0, 6, // EFI_IA32
	}
	opt, err := parseOptClientArchType(data)
	require.NoError(t, err)
	require.Equal(t, opt.ArchTypes[0], iana.EFI_IA32)
}

func TestParseOptClientArchTypeMultiple(t *testing.T) {
	data := []byte{
		0, 6, // EFI_IA32
		0, 2, // EFI_ITANIUM
	}
	opt, err := parseOptClientArchType(data)
	require.NoError(t, err)
	require.Equal(t, opt.ArchTypes[0], iana.EFI_IA32)
	require.Equal(t, opt.ArchTypes[1], iana.EFI_ITANIUM)
}

func TestParseOptClientArchTypeInvalid(t *testing.T) {
	data := []byte{42}
	_, err := parseOptClientArchType(data)
	require.Error(t, err)
}

func TestOptClientArchTypeParseAndToBytesMultiple(t *testing.T) {
	data := []byte{
		0, 6, // EFI_IA32
		0, 8, // EFI_XSCALE
	}
	opt := OptClientArch(iana.EFI_IA32, iana.EFI_XSCALE)
	require.Equal(t, opt.Value.ToBytes(), data)
	require.Equal(t, opt.Code, OptionClientSystemArchitectureType)
}
