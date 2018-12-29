package dhcpv4

import (
	"net"

	"github.com/u-root/u-root/pkg/uio"
)

// optSubnetMask represents an option encapsulating the subnet mask.
//
// This option implements the subnet mask option
// https://tools.ietf.org/html/rfc2132
type optSubnetMask struct {
	SubnetMask net.IPMask
}

// ToBytes returns a serialized stream of bytes for this option.
func (o *optSubnetMask) ToBytes() []byte {
	return o.SubnetMask[:net.IPv4len]
}

// String returns a human-readable string.
func (o *optSubnetMask) String() string {
	return o.SubnetMask.String()
}

// parseOptSubnetMask returns a new OptSubnetMask from a byte
// stream, or error if any.
func parseOptSubnetMask(data []byte) (*optSubnetMask, error) {
	buf := uio.NewBigEndianBuffer(data)
	return &optSubnetMask{net.IPMask(buf.CopyN(net.IPv4len))}, buf.FinError()
}

// GetSubnetMask returns a subnet mask option contained in o, if there is one.
func GetSubnetMask(o Options) net.IPMask {
	v := o.Get(OptionSubnetMask)
	if v == nil {
		return nil
	}
	m, err := parseOptSubnetMask(v)
	if err != nil {
		return nil
	}
	return m.SubnetMask
}

// OptSubnetMask returns a DHCPv4 SubnetMask option per RFC 2132, Section XX.
func OptSubnetMask(mask net.IPMask) Option {
	return Option{
		Code:  OptionSubnetMask,
		Value: &optSubnetMask{mask},
	}
}

// WithNetmask adds or updates an OptionSubnetMask.
func WithNetmask(mask net.IPMask) Modifier {
	return WithOption(OptSubnetMask(mask))
}
