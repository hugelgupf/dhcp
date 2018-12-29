package dhcpv4

// This option implements the Client System Architecture Type option
// https://tools.ietf.org/html/rfc4578

import (
	"fmt"

	"github.com/insomniacslk/dhcp/iana"
	"github.com/u-root/u-root/pkg/uio"
)

// optClientArchType represents an option encapsulating the Client System
// Architecture Type option Definition.
type optClientArchType struct {
	ArchTypes []iana.Arch
}

// OptClientArch returns a new Client System Architecture Type option.
func OptClientArch(archs ...iana.Arch) Option {
	return Option{Code: OptionClientSystemArchitectureType, Value: &optClientArchType{archs}}
}

// GetClientArch returns the Client System Architecture Type option.
func GetClientArch(o Options) []iana.Arch {
	v := o.Get(OptionClientSystemArchitectureType)
	if v == nil {
		return nil
	}
	archs, err := parseOptClientArchType(v)
	if err != nil {
		return nil
	}
	return archs.ArchTypes
}

// ToBytes returns a serialized stream of bytes for this option.
func (o *optClientArchType) ToBytes() []byte {
	buf := uio.NewBigEndianBuffer(nil)
	for _, at := range o.ArchTypes {
		buf.Write16(uint16(at))
	}
	return buf.Data()
}

// String returns a human-readable string.
func (o *optClientArchType) String() string {
	var archTypes string
	for idx, at := range o.ArchTypes {
		archTypes += at.String()
		if idx < len(o.ArchTypes)-1 {
			archTypes += ", "
		}
	}
	return archTypes
}

// parseOptClientArchType returns a new OptClientArchType from a byte stream,
// or error if any.
func parseOptClientArchType(data []byte) (*optClientArchType, error) {
	buf := uio.NewBigEndianBuffer(data)
	if buf.Len() == 0 {
		return nil, fmt.Errorf("must have at least one archtype if option is present")
	}

	archTypes := make([]iana.Arch, 0, buf.Len()/2)
	for buf.Has(2) {
		archTypes = append(archTypes, iana.Arch(buf.Read16()))
	}
	return &optClientArchType{archTypes}, buf.FinError()
}
