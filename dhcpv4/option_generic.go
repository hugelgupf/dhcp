package dhcpv4

import (
	"fmt"
)

// optionGeneric is an option that only contains the option code and associated
// data. Every option that does not have a specific implementation will fall
// back to this option.
type optionGeneric struct {
	Data []byte
}

// ToBytes returns a serialized generic option as a slice of bytes.
func (o optionGeneric) ToBytes() []byte {
	return o.Data
}

// String returns a human-readable representation of a generic option.
func (o optionGeneric) String() string {
	return fmt.Sprintf("%v", o.Data)
}

// OptGeneric returns a generic option.
func OptGeneric(code OptionCode, value []byte) Option {
	return Option{Code: code, Value: &optionGeneric{value}}
}
