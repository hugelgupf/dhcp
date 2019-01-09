package dhcpv4

import (
	"errors"
)

// ErrShortByteStream is an error that is thrown any time a short byte stream is
// detected during option parsing.
var ErrShortByteStream = errors.New("short byte stream")

// ErrZeroLengthByteStream is an error that is thrown any time a zero-length
// byte stream is encountered.
var ErrZeroLengthByteStream = errors.New("zero-length byte stream")

// OptionCode is a single byte representing the code for a given Option.
type OptionCode byte

// Option is an interface that all DHCP v4 options adhere to.
type Option interface {
	Code() OptionCode
	ToBytes() []byte
	Length() int
	String() string
}

// ParseOption parses a sequence of bytes as a single DHCPv4 option, returning
// the specific option structure or error, if any.
func ParseOption(data []byte) (Option, error) {
	if len(data) == 0 {
		return nil, errors.New("invalid zero-length DHCPv4 option")
	}
	var (
		opt Option
		err error
	)
	switch OptionCode(data[0]) {
	case OptionSubnetMask:
		opt, err = ParseOptSubnetMask(data)
	case OptionRouter:
		opt, err = ParseOptRouter(data)
	case OptionDomainNameServer:
		opt, err = ParseOptDomainNameServer(data)
	case OptionHostName:
		opt, err = ParseOptHostName(data)
	case OptionDomainName:
		opt, err = ParseOptDomainName(data)
	case OptionRootPath:
		opt, err = ParseOptRootPath(data)
	case OptionBroadcastAddress:
		opt, err = ParseOptBroadcastAddress(data)
	case OptionNTPServers:
		opt, err = ParseOptNTPServers(data)
	case OptionRequestedIPAddress:
		opt, err = ParseOptRequestedIPAddress(data)
	case OptionIPAddressLeaseTime:
		opt, err = ParseOptIPAddressLeaseTime(data)
	case OptionDHCPMessageType:
		opt, err = ParseOptMessageType(data)
	case OptionServerIdentifier:
		opt, err = ParseOptServerIdentifier(data)
	case OptionParameterRequestList:
		opt, err = ParseOptParameterRequestList(data)
	case OptionMaximumDHCPMessageSize:
		opt, err = ParseOptMaximumDHCPMessageSize(data)
	case OptionClassIdentifier:
		opt, err = ParseOptClassIdentifier(data)
	case OptionTFTPServerName:
		opt, err = ParseOptTFTPServerName(data)
	case OptionBootfileName:
		opt, err = ParseOptBootfileName(data)
	case OptionUserClassInformation:
		opt, err = ParseOptUserClass(data)
	case OptionRelayAgentInformation:
		opt, err = ParseOptRelayAgentInformation(data)
	case OptionClientSystemArchitectureType:
		opt, err = ParseOptClientArchType(data)
	case OptionDNSDomainSearchList:
		opt, err = ParseOptDomainSearch(data)
	case OptionVendorIdentifyingVendorClass:
		opt, err = ParseOptVIVC(data)
	default:
		opt, err = ParseOptionGeneric(data)
	}
	if err != nil {
		return nil, err
	}
	return opt, nil
}

// OptionsFromBytes parses a sequence of bytes until the end and builds a list
// of options from it.
//
// The sequence should not contain the DHCP magic cookie.
//
// Returns an error if any invalid option or length is found.
func OptionsFromBytes(data []byte) ([]Option, error) {
	return OptionsFromBytesWithParser(data, ParseOption)
}

// OptionParser is a function signature for option parsing
type OptionParser func(data []byte) (Option, error)

// OptionsFromBytesWithParser parses Options from byte sequences using the
// parsing function that is passed in as a paremeter
func OptionsFromBytesWithParser(data []byte, parser OptionParser) ([]Option, error) {
	options := make([]Option, 0, 10)
	idx := 0
	for {
		if idx == len(data) {
			break
		}
		// This should never happen.
		if idx > len(data) {
			return nil, errors.New("read past the end of options")
		}
		opt, err := parser(data[idx:])
		idx++
		if err != nil {
			return nil, err
		}
		options = append(options, opt)
		if opt.Code() == OptionEnd {
			break
		}

		// Options with zero length have no length byte, so here we handle the
		// ones with nonzero length
		if opt.Code() != OptionPad {
			idx++
		}
		idx += opt.Length()
	}
	return options, nil
}
