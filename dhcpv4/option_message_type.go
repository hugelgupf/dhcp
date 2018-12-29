package dhcpv4

import (
	"github.com/u-root/u-root/pkg/uio"
)

// This option implements the message type option
// https://tools.ietf.org/html/rfc2132

// optMessageType represents the DHCP message type option.
type optMessageType struct {
	MessageType MessageType
}

// parseOptMessageType constructs an OptMessageType struct from a sequence of
// bytes and returns it, or an error.
func parseOptMessageType(data []byte) (*optMessageType, error) {
	buf := uio.NewBigEndianBuffer(data)
	return &optMessageType{MessageType(buf.Read8())}, buf.FinError()
}

// ToBytes returns a serialized stream of bytes for this option.
func (o *optMessageType) ToBytes() []byte {
	return []byte{byte(o.MessageType)}
}

// String returns a human-readable string for this option.
func (o *optMessageType) String() string {
	return o.MessageType.String()
}

// OptMessageType returns a new DHCPv4 Message Type option.
func OptMessageType(m MessageType) Option {
	return Option{Code: OptionDHCPMessageType, Value: &optMessageType{m}}
}

// GetMessageType returns the DHCPv4 Message Type option in o.
func GetMessageType(o Options) MessageType {
	v := o.Get(OptionDHCPMessageType)
	if v == nil {
		return MessageTypeNone
	}
	mt, err := parseOptMessageType(v)
	if err != nil {
		return MessageTypeNone
	}
	return mt.MessageType
}
