package dhcpv4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptDomainNameInterfaceMethods(t *testing.T) {
	o := OptDomainName("foo")
	require.Equal(t, OptionDomainName, o.Code, "Code")
	require.Equal(t, []byte{'f', 'o', 'o'}, o.Value.ToBytes(), "ToBytes")
	require.Equal(t, "Domain Name: foo", o.String())
}

func TestParseOptDomainName(t *testing.T) {
	o := Options{
		OptionDomainName: []byte{'t', 'e', 's', 't'},
	}
	domain := GetDomainName(o)
	require.Equal(t, "test", domain)
}

func TestOptHostNameInterfaceMethods(t *testing.T) {
	o := OptHostName("foo")
	require.Equal(t, OptionHostName, o.Code, "Code")
	require.Equal(t, []byte{'f', 'o', 'o'}, o.Value.ToBytes(), "ToBytes")
	require.Equal(t, "Host Name: foo", o.String())
}

func TestParseOptHostName(t *testing.T) {
	o := Options{
		OptionHostName: []byte{'t', 'e', 's', 't'},
	}
	host := GetHostName(o)
	require.Equal(t, "test", host)
}

func TestOptRootPathInterfaceMethods(t *testing.T) {
	o := OptRootPath("foo")
	require.Equal(t, OptionRootPath, o.Code, "Code")
	require.Equal(t, []byte{'f', 'o', 'o'}, o.Value.ToBytes(), "ToBytes")
	require.Equal(t, "Root Path: foo", o.String())
}

func TestParseOptRootPath(t *testing.T) {
	o := Options{
		OptionRootPath: []byte{'t', 'e', 's', 't'},
	}
	host := GetRootPath(o)
	require.Equal(t, "test", host)
}

func TestOptBootFileNameInterfaceMethods(t *testing.T) {
	o := OptBootFileName("foo")
	require.Equal(t, OptionBootfileName, o.Code, "Code")
	require.Equal(t, []byte{'f', 'o', 'o'}, o.Value.ToBytes(), "ToBytes")
	require.Equal(t, "Bootfile Name: foo", o.String())
}

func TestParseOptBootFileName(t *testing.T) {
	o := Options{
		OptionBootfileName: []byte{'t', 'e', 's', 't'},
	}
	host := GetBootFileName(o)
	require.Equal(t, "test", host)
}

func TestOptTFTPServerNameInterfaceMethods(t *testing.T) {
	o := OptTFTPServerName("foo")
	require.Equal(t, OptionTFTPServerName, o.Code, "Code")
	require.Equal(t, []byte{'f', 'o', 'o'}, o.Value.ToBytes(), "ToBytes")
	require.Equal(t, "TFTP Server Name: foo", o.String())
}

func TestParseOptTFTPServerName(t *testing.T) {
	o := Options{
		OptionTFTPServerName: []byte{'t', 'e', 's', 't'},
	}
	host := GetTFTPServerName(o)
	require.Equal(t, "test", host)
}

func TestOptClassIdentifierInterfaceMethods(t *testing.T) {
	o := OptClassIdentifier("foo")
	require.Equal(t, OptionClassIdentifier, o.Code, "Code")
	require.Equal(t, []byte{'f', 'o', 'o'}, o.Value.ToBytes(), "ToBytes")
	require.Equal(t, "Class Identifier: foo", o.String())
}

func TestParseOptClassIdentifier(t *testing.T) {
	o := Options{
		OptionClassIdentifier: []byte{'t', 'e', 's', 't'},
	}
	host := GetClassIdentifier(o)
	require.Equal(t, "test", host)
}
