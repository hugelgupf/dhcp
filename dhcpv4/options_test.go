package dhcpv4

import (
	"bytes"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/u-root/u-root/pkg/uio"
)

func TestOptionToBytes(t *testing.T) {
	o := Option{
		Code:  OptionDHCPMessageType,
		Value: &optionGeneric{[]byte{byte(MessageTypeDiscover)}},
	}
	serialized := o.Value.ToBytes()
	expected := []byte{1}
	require.Equal(t, expected, serialized)
}

func TestOptionString(t *testing.T) {
	o := Option{
		Code:  OptionDHCPMessageType,
		Value: &optMessageType{MessageTypeDiscover},
	}
	require.Equal(t, "DHCP Message Type: DISCOVER", o.String())
}

func TestOptionStringUnknown(t *testing.T) {
	o := Option{
		Code:  102, // Returend option code.
		Value: &optionGeneric{[]byte{byte(MessageTypeDiscover)}},
	}
	require.Equal(t, "Unknown: [1]", o.String())
}

func TestOptionsMarshal(t *testing.T) {
	for i, tt := range []struct {
		opts Options
		want []byte
	}{
		{
			opts: nil,
			want: []byte{255},
		},
		{
			opts: Options{
				5: []byte{1, 2, 3, 4},
			},
			want: []byte{
				5 /* key */, 4 /* length */, 1, 2, 3, 4,
				255, /* end key */
			},
		},
		{
			// Test sorted key order.
			opts: Options{
				5:   []byte{1, 2, 3},
				100: []byte{101, 102, 103},
			},
			want: []byte{
				5, 3, 1, 2, 3,
				100, 3, 101, 102, 103,
				255,
			},
		},
		{
			// Test RFC 3396.
			opts: Options{
				5: bytes.Repeat([]byte{10}, math.MaxUint8+1),
			},
			want: append(append(
				[]byte{5, math.MaxUint8}, bytes.Repeat([]byte{10}, math.MaxUint8)...),
				5, 1, 10,
				255,
			),
		},
	} {
		t.Run(fmt.Sprintf("Test %02d", i), func(t *testing.T) {
			b := uio.NewBigEndianBuffer(nil)
			tt.opts.Marshal(b, true)
			require.Equal(t, b.Data(), tt.want)
		})
	}
}

func TestOptionsUnmarshal(t *testing.T) {
	for i, tt := range []struct {
		input     []byte
		want      Options
		wantError bool
	}{
		{
			// Buffer missing data.
			input: []byte{
				3 /* key */, 3 /* length */, 1,
			},
			wantError: true,
		},
		{
			input: []byte{
				// This may look too long, but 0 is padding.
				// The issue here is the missing OptionEnd.
				3, 3, 0, 0, 0, 0, 0, 0, 0,
			},
			wantError: true,
		},
		{
			// Only OptionPad and OptionEnd can stand on their own
			// without a length field. So this is too short.
			input: []byte{
				3,
			},
			wantError: true,
		},
		{
			// Option present after the End is a nono.
			input:     []byte{byte(OptionEnd), 3},
			wantError: true,
		},
		{
			input: []byte{byte(OptionEnd)},
			want:  Options{},
		},
		{
			input: []byte{
				3, 2, 5, 6,
				byte(OptionEnd),
			},
			want: Options{
				3: []byte{5, 6},
			},
		},
		{
			// Test RFC 3396.
			input: append(
				append([]byte{3, math.MaxUint8}, bytes.Repeat([]byte{10}, math.MaxUint8)...),
				3, 5, 10, 10, 10, 10, 10,
				byte(OptionEnd),
			),
			want: Options{
				3: bytes.Repeat([]byte{10}, math.MaxUint8+5),
			},
		},
		{
			input: []byte{
				10, 2, 255, 254,
				11, 3, 5, 5, 5,
				byte(OptionEnd),
			},
			want: Options{
				10: []byte{255, 254},
				11: []byte{5, 5, 5},
			},
		},
		{
			input: append(
				append([]byte{10, 2, 255, 254}, bytes.Repeat([]byte{byte(OptionPad)}, 255)...),
				byte(OptionEnd),
			),
			want: Options{
				10: []byte{255, 254},
			},
		},
	} {
		t.Run(fmt.Sprintf("Test %02d", i), func(t *testing.T) {
			opt, err := OptionsFromBytesWithParser(tt.input, true)
			if tt.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, opt, tt.want)
			}
		})
	}
}
