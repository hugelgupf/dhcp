package dhcpv4

import (
	"encoding/binary"
	"fmt"

	"github.com/u-root/u-root/pkg/uio"
)

// This option implements the Maximum DHCP Message size option
// https://tools.ietf.org/html/rfc2132

// OptMaximumDHCPMessageSize represents the Maximum DHCP Message size option.
type OptMaximumDHCPMessageSize struct {
	Size uint16
}

// ParseOptMaximumDHCPMessageSize constructs an OptMaximumDHCPMessageSize struct from a sequence of
// bytes and returns it, or an error.
func ParseOptMaximumDHCPMessageSize(data []byte) (*OptMaximumDHCPMessageSize, error) {
	buf := uio.NewBigEndianBuffer(data)
	return &OptMaximumDHCPMessageSize{Size: buf.Read16()}, buf.FinError()
}

// Code returns the option code.
func (o *OptMaximumDHCPMessageSize) Code() OptionCode {
	return OptionMaximumDHCPMessageSize
}

// ToBytes returns a serialized stream of bytes for this option.
func (o *OptMaximumDHCPMessageSize) ToBytes() []byte {
	serializedSize := make([]byte, 2)
	binary.BigEndian.PutUint16(serializedSize, o.Size)
	serializedOpt := []byte{byte(o.Code()), byte(o.Length())}
	return append(serializedOpt, serializedSize...)
}

// String returns a human-readable string for this option.
func (o *OptMaximumDHCPMessageSize) String() string {
	return fmt.Sprintf("Maximum DHCP Message Size -> %v", o.Size)
}

// Length returns the length of the data portion (excluding option code and byte
// for length, if any).
func (o *OptMaximumDHCPMessageSize) Length() int {
	return 2
}
