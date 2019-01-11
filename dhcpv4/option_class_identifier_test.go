package dhcpv4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptClassIdentifierInterfaceMethods(t *testing.T) {
	o := OptClassIdentifier{Identifier: "foo"}
	require.Equal(t, OptionClassIdentifier, o.Code(), "Code")
	require.Equal(t, []byte{'f', 'o', 'o'}, o.ToBytes(), "ToBytes")
}

func TestParseOptClassIdentifier(t *testing.T) {
	data := []byte{'t', 'e', 's', 't'}
	o, err := ParseOptClassIdentifier(data)
	require.NoError(t, err)
	require.Equal(t, &OptClassIdentifier{Identifier: "test"}, o)
}

func TestOptClassIdentifierString(t *testing.T) {
	o := OptClassIdentifier{Identifier: "testy test"}
	require.Equal(t, "Class Identifier -> testy test", o.String())
}
