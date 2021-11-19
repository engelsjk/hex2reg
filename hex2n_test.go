package hex2n

import (
	"testing"
)

type Code struct {
	Hex int
	N   string
}

func TestHex2N(t *testing.T) {

	codes := []Code{
		{Hex: 0xA029D9, N: "N11"},
		{Hex: 0xAD21BC, N: "N945WN"},
		{Hex: 0xAC65B8, N: "N8982P"},
		{Hex: 0xA32C7C, N: "N303WS"},
		{Hex: 0xA1A649, N: "N205TK"},
		{Hex: 0xADF782, N: "N9998"},
	}

	for _, code := range codes {
		want := code.N
		got := Hex2N(code.Hex)
		if want != got {
			t.Fatalf(`wanted %s, got %s for %x`, want, got, code.Hex)
		}
	}
}
