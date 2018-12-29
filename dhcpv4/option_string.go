package dhcpv4

// optString represents an option encapsulating the domain name.
//
// This option implements the domain name option
// https://tools.ietf.org/html/rfc2132
type optString struct {
	S string
}

// ToBytes returns a serialized stream of bytes for this option.
func (o optString) ToBytes() []byte {
	return []byte(o.S)
}

// String returns a human-readable string.
func (o optString) String() string {
	return o.S
}

// OptDomainName returns a DHCPv4 Domain Name option.
func OptDomainName(name string) Option {
	return Option{Code: OptionDomainName, Value: optString{name}}
}

// GetDomainName finds and parses the DHCPv4 Domain Name option from o.
func GetDomainName(o Options) string {
	v := o.Get(OptionDomainName)
	if v == nil {
		return ""
	}
	return string(v)
}

// OptHostName returns a DHCPv4 Host Name option.
func OptHostName(name string) Option {
	return Option{Code: OptionHostName, Value: optString{name}}
}

// GetHostName finds and parses the DHCPv4 Host Name option from o.
func GetHostName(o Options) string {
	v := o.Get(OptionHostName)
	if v == nil {
		return ""
	}
	return string(v)
}

// OptRootPath returns a DHCPv4 Root Path option.
func OptRootPath(name string) Option {
	return Option{Code: OptionRootPath, Value: optString{name}}
}

// GetRootPath finds and parses the DHCPv4 Root Path option from o.
func GetRootPath(o Options) string {
	v := o.Get(OptionRootPath)
	if v == nil {
		return ""
	}
	return string(v)
}

// OptBootFileName returns a DHCPv4 Boot File Name option.
func OptBootFileName(name string) Option {
	return Option{Code: OptionBootfileName, Value: optString{name}}
}

// GetBootFileName finds and parses the DHCPv4 Boot File Name option from o.
func GetBootFileName(o Options) string {
	v := o.Get(OptionBootfileName)
	if v == nil {
		return ""
	}
	return string(v)
}

// OptTFTPServerName returns a DHCPv4 TFTP Server Name option.
func OptTFTPServerName(name string) Option {
	return Option{Code: OptionTFTPServerName, Value: optString{name}}
}

// GetTFTPServerName finds and parses the DHCPv4 TFTP Server Name option from o.
func GetTFTPServerName(o Options) string {
	v := o.Get(OptionTFTPServerName)
	if v == nil {
		return ""
	}
	return string(v)
}

// OptClassIdentifier returns a DHCPv4 Class Identifier option.
func OptClassIdentifier(name string) Option {
	return Option{Code: OptionClassIdentifier, Value: optString{name}}
}

// GetClassIdentifier finds and parses the DHCPv4 Class Identifier option from o.
func GetClassIdentifier(o Options) string {
	v := o.Get(OptionClassIdentifier)
	if v == nil {
		return ""
	}
	return string(v)
}
