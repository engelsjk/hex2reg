package hex2reg

import (
	"fmt"
	"math"
	"strconv"
)

type StrideMap struct {
	Start    int
	End      int
	S1       int
	S2       int
	Prefix   string
	Alphabet []string
	First    string
	Last     string
	Offset   int
}

type NumericMap struct {
	Start    int
	End      int
	First    int
	Count    int
	Template string
}

type Hex2Reg struct {
	la []string     // LimitedAlphabet
	fa []string     // FullAlphabet
	sm []StrideMap  // StrideMaps
	nm []NumericMap // NumericMaps
}

func New() Hex2Reg {

	h2r := Hex2Reg{}

	h2r.la = []string{
		"A", "B", "C", "D", "E", "F",
		"G", "H", "J", "K", "L", "M",
		"N", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y", "Z"}

	h2r.fa = []string{
		"A", "B", "C", "D", "E", "F",
		"G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y", "Z"}

	h2r.sm = []StrideMap{
		{Start: 0x390000, S1: 1024, S2: 32, Prefix: "F-G"},
		{Start: 0x398000, S1: 1024, S2: 32, Prefix: "F-H"},
		{Start: 0x3C4421, S1: 1024, S2: 32, Prefix: "D-A", First: "AAA", Last: "OZZ"},
		{Start: 0x3C0001, S1: 26 * 26, S2: 26, Prefix: "D-A", First: "PAA", Last: "ZZZ"},
		{Start: 0x3C8421, S1: 1024, S2: 32, Prefix: "D-B", First: "AAA", Last: "OZZ"},
		{Start: 0x3C2001, S1: 26 * 26, S2: 26, Prefix: "D-B", First: "PAA", Last: "ZZZ"},
		{Start: 0x3CC000, S1: 26 * 26, S2: 26, Prefix: "D-C"},
		{Start: 0x3D04A8, S1: 26 * 26, S2: 26, Prefix: "D-E"},
		{Start: 0x3D4950, S1: 26 * 26, S2: 26, Prefix: "D-F"},
		{Start: 0x3D8DF8, S1: 26 * 26, S2: 26, Prefix: "D-G"},
		{Start: 0x3DD2A0, S1: 26 * 26, S2: 26, Prefix: "D-H"},
		{Start: 0x3E1748, S1: 26 * 26, S2: 26, Prefix: "D-I"},
		{Start: 0x448421, S1: 1024, S2: 32, Prefix: "OO-"},
		{Start: 0x458421, S1: 1024, S2: 32, Prefix: "OY-"},
		{Start: 0x460000, S1: 26 * 26, S2: 26, Prefix: "OH-"},
		{Start: 0x468421, S1: 1024, S2: 32, Prefix: "SX-"},
		{Start: 0x490421, S1: 1024, S2: 32, Prefix: "CS-"},
		{Start: 0x4A0421, S1: 1024, S2: 32, Prefix: "YR-"},
		{Start: 0x4B8421, S1: 1024, S2: 32, Prefix: "TC-"},
		{Start: 0x740421, S1: 1024, S2: 32, Prefix: "JY-"},
		{Start: 0x760421, S1: 1024, S2: 32, Prefix: "AP-"},
		{Start: 0x768421, S1: 1024, S2: 32, Prefix: "9V-"},
		{Start: 0x778421, S1: 1024, S2: 32, Prefix: "YK-"},
		{Start: 0xC00001, S1: 26 * 26, S2: 26, Prefix: "D-I"},
		{Start: 0xC044A9, S1: 26 * 26, S2: 26, Prefix: "D-I"},
		{Start: 0xE01041, S1: 4096, S2: 64, Prefix: "D-I"},
	}

	h2r.nm = []NumericMap{
		{Start: 0x140000, First: 0, Count: 100000, Template: "RA-00000"},
		{Start: 0x0B03E8, First: 1000, Count: 1000, Template: "CU-T0000"},
	}

	h2r.initMaps()

	return h2r
}

func (h Hex2Reg) initMaps() {
	for i := 0; i < len(h.sm); i++ {

		if len(h.sm[i].Alphabet) == 0 {
			h.sm[i].Alphabet = h.fa
		}

		if h.sm[i].First != "" {
			first := []rune(h.sm[i].First)
			var c1, c2, c3 int
			for i, a := range h.sm[i].Alphabet {
				if a == string(first[0]) {
					c1 = i
				}
				if a == string(first[1]) {
					c2 = i
				}
				if a == string(first[2]) {
					c3 = i
				}
			}
			h.sm[i].Offset = c1*h.sm[i].S1 + c2*h.sm[i].S2 + c3
		}

		if h.sm[i].Last != "" {
			last := []rune(h.sm[i].Last)
			var c1, c2, c3 int
			for i, a := range h.sm[i].Alphabet {
				if a == string(last[0]) {
					c1 = i
				}
				if a == string(last[1]) {
					c2 = i
				}
				if a == string(last[2]) {
					c3 = i
				}
				h.sm[i].End = h.sm[i].Start - h.sm[i].Offset + c1*h.sm[i].S1 + c2*h.sm[i].S2 + c3
			}
		} else {
			c := len(h.sm[i].Alphabet) - 1
			h.sm[i].End = h.sm[i].Start - h.sm[i].Offset + c*h.sm[i].S1 + c*h.sm[i].S2 + c
		}
	}

	for i := 0; i < len(h.nm); i++ {
		h.nm[i].End = h.nm[i].Start + h.nm[i].Count - 1
	}

}

/////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////

func (h Hex2Reg) Lookup(hex int) string {

	if hex == 0 {
		return ""
	}

	if reg := h.usa(hex); reg != "" {
		return reg
	}

	if reg := h.kor(hex); reg != "" {
		return reg
	}

	if reg := h.jpn(hex); reg != "" {
		return reg
	}

	if reg := h.numeric(hex); reg != "" {
		return reg
	}

	if reg := h.stride(hex); reg != "" {
		return reg
	}

	return ""
}

/////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////

// Reference: https://github.com/wiedehopf/tar1090/blob/master/html/registrations.js#L202
func (h Hex2Reg) usa(hex int) string {

	offset := hex - 0xA00001
	if offset < 0 || offset >= 915399 {
		return ""
	}

	digit1 := int(math.Floor(float64(offset)/101711.0)) + 1
	reg := "N" + strconv.Itoa(digit1)
	offset = offset % 101711

	if offset <= 600 {
		// Na, NaA .. NaZ, NaAA .. NaZZ
		return reg + nletters(h.la, offset, true)
	}

	// Na0* .. Na9*
	offset -= 601

	digit2 := int(math.Floor(float64(offset) / 10111.0))
	reg += strconv.Itoa(digit2)
	offset = offset % 10111

	if offset <= 600 {
		// Nab, NabA..NabZ, NabAA..NabZZ
		return reg + nletters(h.la, offset, true)
	}

	// Nab0* .. Nab9*
	offset -= 601

	digit3 := int(math.Floor(float64(offset) / 951.0))
	reg += strconv.Itoa(digit3)
	offset = offset % 951

	if offset <= 600 {
		// Nabc, NabcA .. NabcZ, NabcAA .. NabcZZ
		return reg + nletters(h.la, offset, true)
	}

	// Nabc0* .. Nabc9*
	offset -= 601

	digit4 := int(math.Floor(float64(offset) / 35.0))
	reg += strconv.Itoa(digit4)
	offset = offset % 35

	if offset <= 24 {
		// Nabcd, NabcdA .. NabcdZ
		return reg + nletters(h.la, offset, false)
	}

	// Nabcd0 .. Nabcd9
	offset -= 25
	return reg + strconv.Itoa(offset)
}

// Reference: https://github.com/wiedehopf/tar1090/blob/3baff7b112f5d7ee9100ba7e4ea0cb9b6f4c5b42/html/registrations.js#L275
func (h Hex2Reg) jpn(hex int) string {

	offset := hex - 0x840000
	if offset < 0 || offset >= 229840 {
		return ""
	}

	reg := "JA"

	digit1 := int(math.Floor(float64(offset) / 22984.0))
	if digit1 < 0 || digit1 > 9 {
		return ""
	}

	reg += strconv.Itoa(digit1)
	offset = offset % 22984

	digit2 := int(math.Floor(float64(offset) / 916.0))
	if digit2 < 0 || digit2 > 9 {
		return ""
	}

	reg += strconv.Itoa(digit2)
	offset = offset % 916

	if offset < 340 {
		digit3 := int(math.Floor(float64(offset) / 34.0))
		reg += strconv.Itoa(digit3)
		offset = offset % 34

		if offset < 10 {
			reg += strconv.Itoa(offset)
			return reg
		}

		offset -= 10
		reg += h.la[offset]
		return reg
	}

	offset -= 340

	letter3 := int(math.Floor(float64(offset) / 24.0))
	reg += h.la[letter3] + h.la[letter3%24]

	return reg
}

// Reference: https://github.com/wiedehopf/tar1090/blob/3baff7b112f5d7ee9100ba7e4ea0cb9b6f4c5b42/html/registrations.js#L258
func (h Hex2Reg) kor(hex int) string {
	if hex >= 0x71BA00 && hex <= 0x71BF99 {
		return "HL" + fmt.Sprintf("%X", hex-0x71BA00+0x7200)
	}
	if hex >= 0x71C000 && hex <= 0x71C099 {
		return "HL" + fmt.Sprintf("%x", hex-0x71C000+0x8000)
	}
	if hex >= 0x71C200 && hex <= 0x71C299 {
		return "HL" + fmt.Sprintf("%x", hex-0x71C200+0x8200)
	}
	return ""
}

// Reference: https://github.com/wiedehopf/tar1090/blob/3baff7b112f5d7ee9100ba7e4ea0cb9b6f4c5b42/html/registrations.js#L169
func (h Hex2Reg) numeric(hex int) string {
	for i := 0; i < len(h.nm); i++ {
		if hex < h.nm[i].Start || hex > h.nm[i].End {
			continue
		}
		reg := strconv.Itoa((hex - h.nm[i].Start + h.nm[i].First))
		idx := len([]rune(h.nm[i].Template)) - len([]rune(reg))
		reg = h.nm[i].Template[0:idx] + reg
		return reg
	}
	return ""
}

// Reference: https://github.com/wiedehopf/tar1090/blob/3baff7b112f5d7ee9100ba7e4ea0cb9b6f4c5b42/html/registrations.js#L141
func (h Hex2Reg) stride(hex int) string {
	for i := 0; i < len(h.sm); i++ {
		if hex < h.sm[i].Start || hex > h.sm[i].End {
			continue
		}

		offset := hex - h.sm[i].Start + h.sm[i].Offset

		i1 := int(math.Floor(float64(offset) / float64(h.sm[i].S1)))
		offset = offset % h.sm[i].S1
		i2 := int(math.Floor(float64(offset) / float64(h.sm[i].S2)))
		offset = offset % h.sm[i].S2
		i3 := offset

		l := len(h.sm[i].Alphabet)
		if i1 < 0 || i1 >= l || i2 < 0 || i2 >= l || i3 < 0 || i3 >= l {
			continue
		}

		reg := h.sm[i].Prefix + h.sm[i].Alphabet[i1] + h.sm[i].Alphabet[i2] + h.sm[i].Alphabet[i3]
		return reg
	}
	return ""
}

func nletters(a []string, rem int, s bool) string {

	if rem == 0 {
		return ""
	}

	rem -= 1

	if !s {
		return a[rem]
	}

	index := int(math.Floor(float64(rem) / 25.0))
	return a[index] + nletters(a, rem%25, false)
}
