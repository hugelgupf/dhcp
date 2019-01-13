package dhcpv4

import (
	"fmt"
	"net"

	"github.com/u-root/u-root/pkg/uio"
)

// OptBroadcastAddress implements the broadcast address option described in RFC
// 2132, Section 5.3.
type OptBroadcastAddress struct {
	BroadcastAddress net.IP
}

// ParseOptBroadcastAddress returns a new OptBroadcastAddress from a byte
// stream, or error if any.
func ParseOptBroadcastAddress(data []byte) (*OptBroadcastAddress, error) {
	buf := uio.NewBigEndianBuffer(data)
	return &OptBroadcastAddress{BroadcastAddress: net.IP(buf.CopyN(net.IPv4len))}, buf.FinError()
}

// Code returns the option code.
func (o *OptBroadcastAddress) Code() OptionCode {
	return OptionBroadcastAddress
}

// ToBytes returns a serialized stream of bytes for this option.
func (o *OptBroadcastAddress) ToBytes() []byte {
	return []byte(o.BroadcastAddress.To4())
}

// String returns a human-readable string.
func (o *OptBroadcastAddress) String() string {
	return fmt.Sprintf("Broadcast Address -> %v", o.BroadcastAddress.String())
}
