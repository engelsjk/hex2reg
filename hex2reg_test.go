package hex2reg

import (
	"testing"
)

type Code struct {
	Hex int
	N   string
}

func testCodes(t *testing.T, codes []Code) {
	h2r := NewHex2Reg()
	for _, code := range codes {
		want := code.N
		got := h2r.Lookup(code.Hex)
		if want != got {
			t.Fatalf(`wanted %s, got %s for %x`, want, got, code.Hex)
		}
	}
}
func TestHex2RegUSA(t *testing.T) {
	codes := []Code{
		{Hex: 0xA029D9, N: "N11"}, // USA
		{Hex: 0xAA0DB8, N: "N747NA"},
		{Hex: 0xADF782, N: "N9998"},
	}
	testCodes(t, codes)
}

func TestHex2RegJapan(t *testing.T) {
	codes := []Code{
		{Hex: 0x86CF38, N: "JA807A"}, // Japan
		{Hex: 0x84B5AC, N: "JA20JJ"},
	}
	testCodes(t, codes)
}

func TestHex2RegSouthKorea(t *testing.T) {
	codes := []Code{
		{Hex: 0x71C225, N: "HL8225"}, // South Korea
		{Hex: 0x71C337, N: "HL8337"},
	}
	testCodes(t, codes)
}

func TestHex2RegPortugal(t *testing.T) {
	codes := []Code{
		{Hex: 0x491145, N: "CS-DJE"}, // Portugal
	}
	testCodes(t, codes)
}

func TestHex2RegSingapore(t *testing.T) {
	codes := []Code{
		{Hex: 0x76CDB4, N: "9V-SMT"}, // Singapore
	}
	testCodes(t, codes)
}
