package dhcpv4

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/u-root/u-root/pkg/uio"
)

// This option implements the Vendor-Identifying Vendor Class Option
// https://tools.ietf.org/html/rfc3925

// VIVCIdentifier represents one Vendor-Identifying vendor class option.
type VIVCIdentifier struct {
	EntID uint32
	Data  []byte
}

// OptVIVC represents the DHCP message type option.
type OptVIVC struct {
	Identifiers []VIVCIdentifier
}

// ParseOptVIVC contructs an OptVIVC tsruct from a sequence of bytes and returns
// it, or an error.
func ParseOptVIVC(data []byte) (*OptVIVC, error) {
	buf := uio.NewBigEndianBuffer(data)

	var ids []VIVCIdentifier
	for buf.Has(5) {
		entID := buf.Read32()
		idLen := int(buf.Read8())
		ids = append(ids, VIVCIdentifier{EntID: entID, Data: buf.CopyN(idLen)})
	}

	return &OptVIVC{Identifiers: ids}, buf.FinError()
}

// Code returns the option code.
func (o *OptVIVC) Code() OptionCode {
	return OptionVendorIdentifyingVendorClass
}

// ToBytes returns a serialized stream of bytes for this option.
func (o *OptVIVC) ToBytes() []byte {
	buf := make([]byte, o.Length()+2)
	copy(buf[0:], []byte{byte(o.Code()), byte(o.Length())})

	b := buf[2:]
	for _, id := range o.Identifiers {
		binary.BigEndian.PutUint32(b[0:4], id.EntID)
		b[4] = byte(len(id.Data))
		copy(b[5:], id.Data)
		b = b[len(id.Data)+5:]
	}
	return buf
}

// String returns a human-readable string for this option.
func (o *OptVIVC) String() string {
	buf := bytes.Buffer{}
	fmt.Fprintf(&buf, "Vendor-Identifying Vendor Class ->")

	for _, id := range o.Identifiers {
		fmt.Fprintf(&buf, " %d:'%s',", id.EntID, id.Data)
	}

	return buf.String()[:buf.Len()-1]
}

// Length returns the length of the data portion (excluding option code and byte
// for length, if any).
func (o *OptVIVC) Length() int {
	n := 0
	for _, id := range o.Identifiers {
		// each class has a header of endID (4 bytes) and length (1 byte)
		n += 5 + len(id.Data)
	}
	return n
}
