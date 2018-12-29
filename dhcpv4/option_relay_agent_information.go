package dhcpv4

import (
	"github.com/u-root/u-root/pkg/uio"
)

// optRelayAgentInformation is a "container" option for specific agent-supplied
// sub-options.
type optRelayAgentInformation struct {
	Options Options
}

func OptRelayAgentInfo(o Options) Option {
	return Option{Code: OptionRelayAgentInformation, Value: &optRelayAgentInformation{o}}
}

// This option implements the relay agent information option
// https://tools.ietf.org/html/rfc3046
func GetRelayAgentInfo(o Options) Options {
	v := o.Get(OptionRelayAgentInformation)
	if v == nil {
		return nil
	}
	r, err := parseRelayAgentInfo(v)
	if err != nil {
		return nil
	}
	return r.Options
}

func parseRelayAgentInfo(v []byte) (*optRelayAgentInformation, error) {
	options, err := OptionsFromBytesWithParser(v, false)
	if err != nil {
		return nil, err
	}
	return &optRelayAgentInformation{options}, nil
}

// ToBytes returns a serialized stream of bytes for this option.
func (o *optRelayAgentInformation) ToBytes() []byte {
	buf := uio.NewBigEndianBuffer(nil)
	o.Options.Marshal(buf, false)
	return buf.Data()
}

// String returns a human-readable string for this option.
func (o *optRelayAgentInformation) String() string {
	return o.Options.String()
}
