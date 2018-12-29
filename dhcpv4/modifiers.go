package dhcpv4

import (
	"net"
	"time"

	"github.com/insomniacslk/dhcp/rfc1035label"
)

// WithTransactionID sets the Transaction ID for the DHCPv4 packet
func WithTransactionID(xid TransactionID) Modifier {
	return func(d *DHCPv4) *DHCPv4 {
		d.TransactionID = xid
		return d
	}
}

// WithBroadcast sets the packet to be broadcast or unicast
func WithBroadcast(broadcast bool) Modifier {
	return func(d *DHCPv4) *DHCPv4 {
		if broadcast {
			d.SetBroadcast()
		} else {
			d.SetUnicast()
		}
		return d
	}
}

// WithHwAddr sets the hardware address for a packet
func WithHwAddr(hwaddr net.HardwareAddr) Modifier {
	return func(d *DHCPv4) *DHCPv4 {
		d.ClientHWAddr = hwaddr
		return d
	}
}

// WithOption appends a DHCPv4 option provided by the user
func WithOption(opt Option) Modifier {
	return func(d *DHCPv4) *DHCPv4 {
		d.AddOption(opt)
		return d
	}
}

// WithUserClass adds a user class option to the packet.
// The rfc parameter allows you to specify if the userclass should be
// rfc compliant or not. More details in issue #113
func WithUserClass(uc []byte, rfc bool) Modifier {
	// TODO let the user specify multiple user classes
	return func(d *DHCPv4) *DHCPv4 {
		if rfc {
			d.AddOption(OptRFC3004UserClass([][]byte{uc}))
		} else {
			d.AddOption(OptUserClass(uc))
		}
		return d
	}
}

// WithNetboot adds bootfile URL and bootfile param options to a DHCPv4 packet.
func WithNetboot(d *DHCPv4) *DHCPv4 {
	d.AddOption(OptParameterRequestList(OptionTFTPServerName, OptionBootfileName))
	return d
}

// WithRequestedOptions adds requested options to the packet.
func WithRequestedOptions(optionCodes ...OptionCode) Modifier {
	return WithOption(OptParameterRequestList(optionCodes...))
}

// WithRelay adds parameters required for DHCPv4 to be relayed by the relay
// server with given ip
func WithRelay(ip net.IP) Modifier {
	return func(d *DHCPv4) *DHCPv4 {
		d.SetUnicast()
		d.GatewayIPAddr = ip
		d.HopCount = 1
		return d
	}
}

// WithLeaseTime adds or updates an OptIPAddressLeaseTime
func WithLeaseTime(leaseDuration time.Duration) Modifier {
	return func(d *DHCPv4) *DHCPv4 {
		d.UpdateOption(OptIPAddressLeaseTime(leaseDuration))
		return d
	}
}

// WithDomainSearchList adds or updates an OptionDomainSearch
func WithDomainSearchList(searchList ...string) Modifier {
	return func(d *DHCPv4) *DHCPv4 {
		labels := rfc1035label.Labels{
			Labels: searchList,
		}
		odsl := OptDomainSearch(&labels)
		d.UpdateOption(odsl)
		return d
	}
}
