package bsdp

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/u-root/u-root/pkg/uio"
)

func bigEndianToBytes(m uio.Marshaler) []byte {
	buf := uio.NewBigEndianBuffer(nil)
	m.Marshal(buf)
	return buf.Data()
}

func bigEndianFromBytes(m uio.Unmarshaler, b []byte) error {
	buf := uio.NewBigEndianBuffer(b)
	return m.Unmarshal(buf)
}

func TestBootImageIDToBytes(t *testing.T) {
	b := BootImageID{
		IsInstall: true,
		ImageType: BootImageTypeMacOSX,
		Index:     0x1000,
	}
	actual := b.ToBytes()
	expected := []byte{0x81, 0, 0x10, 0}
	require.Equal(t, expected, actual)

	b.IsInstall = false
	actual = b.ToBytes()
	expected = []byte{0x01, 0, 0x10, 0}
	require.Equal(t, expected, actual)
}

func TestBootImageIDFromBytes(t *testing.T) {
	b := BootImageID{
		IsInstall: false,
		ImageType: BootImageTypeMacOSX,
		Index:     0x1000,
	}
	var newBootImage BootImageID
	require.NoError(t, bigEndianFromBytes(&newBootImage, b.ToBytes()))
	require.Equal(t, b, newBootImage)

	b = BootImageID{
		IsInstall: true,
		ImageType: BootImageTypeMacOSX,
		Index:     0x1011,
	}
	require.NoError(t, bigEndianFromBytes(&newBootImage, b.ToBytes()))
	require.Equal(t, b, newBootImage)
}

func TestBootImageIDFromBytesFail(t *testing.T) {
	serialized := []byte{0x81, 0, 0x10} // intentionally left short
	var deserialized BootImageID
	require.Error(t, bigEndianFromBytes(&deserialized, serialized))
}

func TestBootImageIDString(t *testing.T) {
	b := BootImageID{IsInstall: false, ImageType: BootImageTypeMacOSX, Index: 1001}
	require.Equal(t, "[1001] uninstallable macOS image", b.String())
}

/*
 * BootImage
 */
func TestBootImageToBytes(t *testing.T) {
	b := BootImage{
		ID: BootImageID{
			IsInstall: true,
			ImageType: BootImageTypeMacOSX,
			Index:     0x1000,
		},
		Name: "bsdp-1",
	}
	expected := []byte{
		0x81, 0, 0x10, 0, // boot image ID
		6,                         // len(Name)
		98, 115, 100, 112, 45, 49, // byte-encoding of Name
	}
	actual := b.ToBytes()
	require.Equal(t, expected, actual)

	b = BootImage{
		ID: BootImageID{
			IsInstall: false,
			ImageType: BootImageTypeMacOSX,
			Index:     0x1010,
		},
		Name: "bsdp-21",
	}
	expected = []byte{
		0x1, 0, 0x10, 0x10, // boot image ID
		7,                             // len(Name)
		98, 115, 100, 112, 45, 50, 49, // byte-encoding of Name
	}
	actual = b.ToBytes()
	require.Equal(t, expected, actual)
}

func TestBootImageFromBytes(t *testing.T) {
	input := []byte{
		0x1, 0, 0x10, 0x10, // boot image ID
		7,                             // len(Name)
		98, 115, 100, 112, 45, 50, 49, // byte-encoding of Name
	}
	var b BootImage
	require.NoError(t, bigEndianFromBytes(&b, input))
	expectedBootImage := BootImage{
		ID: BootImageID{
			IsInstall: false,
			ImageType: BootImageTypeMacOSX,
			Index:     0x1010,
		},
		Name: "bsdp-21",
	}
	require.Equal(t, expectedBootImage, b)
}

func TestBootImageFromBytesOnlyBootImageID(t *testing.T) {
	// Only a BootImageID, nothing else.
	input := []byte{0x1, 0, 0x10, 0x10}
	var b BootImage
	require.Error(t, bigEndianFromBytes(&b, input))
}

func TestBootImageFromBytesShortBootImage(t *testing.T) {
	input := []byte{
		0x1, 0, 0x10, 0x10, // boot image ID
		7,                         // len(Name)
		98, 115, 100, 112, 45, 50, // Name bytes (intentionally off-by-one)
	}
	var b BootImage
	require.Error(t, bigEndianFromBytes(&b, input))
}

func TestBootImageString(t *testing.T) {
	b := BootImage{
		ID: BootImageID{
			IsInstall: false,
			ImageType: BootImageTypeMacOSX,
			Index:     0x1010,
		},
		Name: "bsdp-21",
	}
	require.Equal(t, "bsdp-21 [4112] uninstallable macOS image", b.String())
}
