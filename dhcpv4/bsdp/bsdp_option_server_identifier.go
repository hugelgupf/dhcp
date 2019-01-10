package bsdp

import (
	"fmt"
	"net"

	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/u-root/u-root/pkg/uio"
)

// OptServerIdentifier represents an option encapsulating the server identifier.
type OptServerIdentifier struct {
	ServerID net.IP
}

// ParseOptServerIdentifier returns a new OptServerIdentifier from a byte
// stream, or error if any.
func ParseOptServerIdentifier(data []byte) (*OptServerIdentifier, error) {
	buf := uio.NewBigEndianBuffer(data)
	return &OptServerIdentifier{ServerID: net.IP(buf.CopyN(net.IPv4len))}, buf.FinError()
}

// Code returns the option code.
func (o *OptServerIdentifier) Code() dhcpv4.OptionCode {
	return OptionServerIdentifier
}

// ToBytes returns a serialized stream of bytes for this option.
func (o *OptServerIdentifier) ToBytes() []byte {
	ret := []byte{byte(o.Code()), byte(o.Length())}
	return append(ret, o.ServerID.To4()...)
}

// String returns a human-readable string.
func (o *OptServerIdentifier) String() string {
	return fmt.Sprintf("BSDP Server Identifier -> %v", o.ServerID.String())
}

// Length returns the length of the data portion (excluding option code an byte
// length).
func (o *OptServerIdentifier) Length() int {
	return len(o.ServerID.To4())
}
