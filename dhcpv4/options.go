package dhcpv4

import (
	"errors"
	"fmt"
	"io"
	"math"
	"sort"
	"strings"

	"github.com/u-root/u-root/pkg/uio"
)

var (
	// ErrShortByteStream is an error that is thrown any time a short byte stream is
	// detected during option parsing.
	ErrShortByteStream = errors.New("short byte stream")

	// ErrZeroLengthByteStream is an error that is thrown any time a zero-length
	// byte stream is encountered.
	ErrZeroLengthByteStream = errors.New("zero-length byte stream")

	// ErrInvalidOptions is returned when invalid options data is
	// encountered during parsing. The data could report an incorrect
	// length or have trailing bytes which are not part of the option.
	ErrInvalidOptions = errors.New("invalid options data")
)

// magicCookie is the magic 4-byte value at the beginning of the list of options
// in a DHCPv4 packet.
var magicCookie = [4]byte{99, 130, 83, 99}

// OptionCode is a single byte representing the code for a given Option.
type OptionCode byte

// Option is an interface that all DHCP v4 options adhere to.
type OptionValue interface {
	ToBytes() []byte
	String() string
}

type Option struct {
	Code  OptionCode
	Value OptionValue
}

func (o Option) String() string {
	return fmt.Sprintf("%s: %s", o.Code, o.Value)
}

// ParseOption parses a sequence of bytes as a single DHCPv4 option, returning
// the specific option structure or error, if any.
/*func ParseOption(code OptionCode, data []byte) (Option, error) {
	if err != nil {
		return nil, err
	}
	return opt, nil
}*/

// OptionsFromBytesWithoutMagicCookie parses a sequence of bytes until the end
// and builds a list of options from it. The sequence should not contain the
// DHCP magic cookie. Returns an error if any invalid option or length is found.
func OptionsFromBytesWithoutMagicCookie(data []byte) (Options, error) {
	return OptionsFromBytesWithParser(data, true)
}

// OptionsFromBytesWithParser parses Options from byte sequences using the
// parsing function that is passed in as a paremeter
func OptionsFromBytesWithParser(data []byte, checkEndOption bool) (Options, error) {
	if len(data) == 0 {
		return nil, nil
	}
	buf := uio.NewBigEndianBuffer(data)

	options := make(Options)
	// Due to RFC 3396 allowing an option to be specified multiple times,
	// we have to collect all option data first, and then parse it.
	var end bool
	for buf.Len() >= 1 {
		// 1 byte: option code
		// 1 byte: option length n
		// n bytes: data
		code := OptionCode(buf.Read8())

		if code == OptionPad {
			continue
		} else if code == OptionEnd {
			end = true
			break
		}
		length := int(buf.Read8())

		// N bytes: option data
		data := buf.Consume(length)
		if data == nil {
			return nil, fmt.Errorf("error collecting options: %v", buf.Error())
		}
		data = data[:length:length]

		// RFC 3396: Just concatenate the data if the option code was
		// specified multiple times.
		options[code] = append(options[code], data...)
	}

	// If we never read the End option, the sender of this packet screwed
	// up.
	if !end && checkEndOption {
		return nil, io.ErrUnexpectedEOF
	}

	// Any bytes left must be padding.
	for buf.Len() >= 1 {
		if OptionCode(buf.Read8()) != OptionPad {
			return nil, ErrInvalidOptions
		}
	}
	return options, nil
}

// Options is a collection of options.
type Options map[OptionCode][]byte

// Get will attempt to get all options that match a DHCPv4 option
// from its OptionCode.  If the option was not found it will return an
// empty list.
//
// According to RFC 3396, options that are specified more than once are
// concatenated, and hence this should always just return one option. This
// currently returns a list to be API compatible.
func (o Options) Get(code OptionCode) []byte {
	return o[code]
}

// Has checks whether o has the given opcode.
func (o Options) Has(opcode OptionCode) bool {
	_, ok := o[opcode]
	return ok
}

// Add appends an option to the existing ones.
func (o Options) Add(option Option) {
	c := option.Code
	o[c] = append(o[c], option.Value.ToBytes()...)
}

// Update updates the existing options with the passed option, adding it
// at the end if not present already
func (o Options) Update(option Option) {
	o[option.Code] = option.Value.ToBytes()
}

// Marshal writes options binary representations to b.
func (o Options) Marshal(b *uio.Lexer, writeEnd bool) {
	for _, c := range o.sortedKeys() {
		code := OptionCode(c)
		// Even if the End option is in there, don't marshal it until
		// the end.
		if code == OptionEnd {
			continue
		}

		data := o[code]

		// RFC 3396: If more than 256 bytes of data are given, the
		// option is simply listed multiple times.
		for len(data) > 0 {
			// 1 byte: option code
			b.Write8(uint8(code))

			// Some DHCPv4 options have fixed length and do not put
			// length on the wire.
			if code == OptionPad {
				continue
			}

			n := len(data)
			if n > math.MaxUint8 {
				n = math.MaxUint8
			}

			// 1 byte: option length
			b.Write8(uint8(n))

			// N bytes: option data
			b.WriteBytes(data[:n])
			data = data[n:]
		}
	}

	if writeEnd {
		b.Write8(uint8(OptionEnd))
	}
}

func getOption(code OptionCode, data []byte) *Option {
	option := &Option{
		Code:  code,
		Value: &optionGeneric{data},
	}

	var (
		opt OptionValue
		err error
	)
	switch code {
	case OptionSubnetMask:
		opt, err = parseOptSubnetMask(data)

	case OptionRouter, OptionDomainNameServer, OptionNTPServers, OptionServerIdentifier:
		opt, err = parseIPs(data)

	case OptionHostName, OptionDomainName, OptionRootPath,
		OptionClassIdentifier, OptionTFTPServerName, OptionBootfileName:
		opt = &optString{string(data)}

	case OptionBroadcastAddress, OptionRequestedIPAddress:
		opt, err = parseIP(data)

	case OptionIPAddressLeaseTime:
		opt, err = parseOptIPAddressLeaseTime(data)

	case OptionDHCPMessageType:
		opt, err = parseOptMessageType(data)

	case OptionParameterRequestList:
		opt, err = parseOptParameterRequestList(data)

	case OptionUserClassInformation:
		opt, err = parseOptUserClass(data)

	case OptionRelayAgentInformation:
		opt, err = parseRelayAgentInfo(data)

	case OptionClientSystemArchitectureType:
		opt, err = parseOptClientArchType(data)

	case OptionMaximumDHCPMessageSize:
		opt, err = ParseOptMaximumDHCPMessageSize(data)
	case OptionDNSDomainSearchList:
		opt, err = parseDomainSearch(data)
	case OptionVendorIdentifyingVendorClass:
		opt, err = ParseOptVIVC(data)
	}
	if err == nil {
		option.Value = opt
	}
	return option
}

func (o Options) String() string {
	return o.ToString(getOption)
}

// OptionParser is a function signature for option parsing
type OptionParser func(code OptionCode, data []byte) *Option

func (o Options) ToString(parse OptionParser) string {
	var ret string
	for code, v := range o {
		option := parse(code, v)
		optString := option.String()
		// If this option has sub structures, offset them accordingly.
		if strings.Contains(optString, "\n") {
			optString = strings.Replace(optString, "\n  ", "\n      ", -1)
		}
		ret += fmt.Sprintf("    %v\n", optString)
	}
	return ret
}

// sortedKeys returns an ordered slice of option keys from the Options map, for
// use in serializing options to binary.
func (o Options) sortedKeys() []int {
	// Send all values for a given key
	var codes []int
	for k := range o {
		codes = append(codes, int(k))
	}

	sort.Sort(sort.IntSlice(codes))
	return codes
}
