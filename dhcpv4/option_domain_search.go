package dhcpv4

// This module defines the OptDomainSearch structure.
// https://tools.ietf.org/html/rfc3397

import (
	"fmt"

	"github.com/insomniacslk/dhcp/rfc1035label"
)

// FIXME rename OptDomainSearch to OptDomainSearchList, and DomainSearch to
// SearchList, for consistency with the equivalent v6 option

// optDomainSearch represents an option encapsulating a domain search list.
type optDomainSearch struct {
	DomainSearch *rfc1035label.Labels
}

func OptDomainSearch(labels *rfc1035label.Labels) Option {
	return Option{Code: OptionDNSDomainSearchList, Value: &optDomainSearch{labels}}
}

func GetDomainSearch(o Options) *rfc1035label.Labels {
	v := o.Get(OptionDNSDomainSearchList)
	if v == nil {
		return nil
	}
	ds, err := parseDomainSearch(v)
	if err != nil {
		return nil
	}
	return ds.DomainSearch
}

func parseDomainSearch(v []byte) (*optDomainSearch, error) {
	labels, err := rfc1035label.FromBytes(v)
	if err != nil {
		return nil, err
	}
	return &optDomainSearch{labels}, nil
}

// ToBytes returns a serialized stream of bytes for this option.
func (op *optDomainSearch) ToBytes() []byte {
	return op.DomainSearch.ToBytes()
}

// String returns a human-readable string.
func (op *optDomainSearch) String() string {
	return fmt.Sprintf("%v", op.DomainSearch.Labels)
}
