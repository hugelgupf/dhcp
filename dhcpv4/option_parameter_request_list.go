package dhcpv4

import (
	"fmt"
	"strings"

	"github.com/u-root/u-root/pkg/uio"
)

// This option implements the parameter request list option
// https://tools.ietf.org/html/rfc2132

// optParameterRequestList represents the parameter request list option.
type optParameterRequestList struct {
	RequestedOpts []OptionCode
}

// parseOptParameterRequestList returns a new OptParameterRequestList from a
// byte stream, or error if any.
func parseOptParameterRequestList(data []byte) (*optParameterRequestList, error) {
	buf := uio.NewBigEndianBuffer(data)
	requestedOpts := make([]OptionCode, 0, buf.Len())
	for buf.Len() > 0 {
		requestedOpts = append(requestedOpts, OptionCode(buf.Read8()))
	}
	return &optParameterRequestList{requestedOpts}, buf.Error()
}

// ToBytes returns a serialized stream of bytes for this option.
func (o *optParameterRequestList) ToBytes() []byte {
	buf := uio.NewBigEndianBuffer(nil)
	for _, req := range o.RequestedOpts {
		buf.Write8(uint8(req))
	}
	return buf.Data()
}

// String returns a human-readable string for this option.
func (o *optParameterRequestList) String() string {
	var optNames []string
	for _, ro := range o.RequestedOpts {
		name := ro.String()
		if name == "Unknown" {
			name += fmt.Sprintf("%s (%v)", name, ro)
		}
		optNames = append(optNames, name)
	}
	return strings.Join(optNames, ", ")
}

// GetParameterRequestList returns the DHCPv4 Parameter Request List in o.
func GetParameterRequestList(o Options) []OptionCode {
	v := o.Get(OptionParameterRequestList)
	if v == nil {
		return nil
	}
	codes, err := parseOptParameterRequestList(v)
	if err != nil {
		return nil
	}
	return codes.RequestedOpts
}

// OptParameterRequestList returns a new DHCPv4 Parameter Request List.
func OptParameterRequestList(codes ...OptionCode) Option {
	return Option{Code: OptionParameterRequestList, Value: &optParameterRequestList{codes}}
}
