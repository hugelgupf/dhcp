package dhcpv4

import (
	"fmt"
	"net"

	"github.com/u-root/u-root/pkg/uio"
)

// OptSubnetMask implements the subnet mask option described by RFC 2132,
// Section 3.3.
type OptSubnetMask struct {
	SubnetMask net.IPMask
}

// ParseOptSubnetMask returns a new OptSubnetMask from a byte
// stream, or error if any.
func ParseOptSubnetMask(data []byte) (*OptSubnetMask, error) {
	buf := uio.NewBigEndianBuffer(data)
	return &OptSubnetMask{SubnetMask: net.IPMask(buf.CopyN(net.IPv4len))}, buf.FinError()
}

// Code returns the option code.
func (o *OptSubnetMask) Code() OptionCode {
	return OptionSubnetMask
}

// ToBytes returns a serialized stream of bytes for this option.
func (o *OptSubnetMask) ToBytes() []byte {
	return o.SubnetMask[:net.IPv4len]
}

// String returns a human-readable string.
func (o *OptSubnetMask) String() string {
	return fmt.Sprintf("Subnet Mask -> %v", o.SubnetMask.String())
}
