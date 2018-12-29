package dhcpv4

import (
	"fmt"
	"time"

	"github.com/u-root/u-root/pkg/uio"
)

// This option implements the IP Address Lease Time option
// https://tools.ietf.org/html/rfc2132

// optIPAddressLeaseTime represents the IP Address Lease Time option.
type optIPAddressLeaseTime struct {
	LeaseTime time.Duration
}

// parseOptIPAddressLeaseTime constructs an OptIPAddressLeaseTime struct from a
// sequence of bytes and returns it, or an error.
func parseOptIPAddressLeaseTime(data []byte) (*optIPAddressLeaseTime, error) {
	buf := uio.NewBigEndianBuffer(data)
	return &optIPAddressLeaseTime{time.Duration(buf.Read32()) * time.Second}, buf.FinError()
}

// ToBytes returns a serialized stream of bytes for this option.
func (o *optIPAddressLeaseTime) ToBytes() []byte {
	buf := uio.NewBigEndianBuffer(nil)
	buf.Write32(uint32(o.LeaseTime / time.Second))
	return buf.Data()
}

// String returns a human-readable string for this option.
func (o *optIPAddressLeaseTime) String() string {
	return fmt.Sprintf("%s", o.LeaseTime)
}

func OptIPAddressLeaseTime(d time.Duration) Option {
	return Option{Code: OptionIPAddressLeaseTime, Value: &optIPAddressLeaseTime{d}}
}

func GetIPAddressLeaseTime(o Options, def time.Duration) time.Duration {
	v := o.Get(OptionIPAddressLeaseTime)
	if v == nil {
		return def
	}
	d, err := parseOptIPAddressLeaseTime(v)
	if err != nil {
		return def
	}
	return d.LeaseTime
}
