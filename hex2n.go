package hex2n

import (
	"math"
	"strconv"
)

// Reference: https://github.com/wiedehopf/tar1090/blob/master/html/registrations.js#L202

// Hex2N converts an ICAO 24-bit Mode S code (in base 16 hexadecimal integer format) to an N-Number.
// Only valid for Mode S codes assigned to US aircraft registered with the FAA.
func Hex2N(hex int) string {

	offset := hex - 0xA00001
	if offset < 0 || offset >= 915399 {
		return ""
	}

	digit1 := int(math.Floor(float64(offset)/101711.0)) + 1
	reg := "N" + strconv.Itoa(digit1)
	offset = offset % 101711

	if offset <= 600 {
		// Na, NaA .. NaZ, NaAA .. NaZZ
		return reg + nletters(offset, true)
	}

	// Na0* .. Na9*
	offset -= 601

	digit2 := int(math.Floor(float64(offset) / 10111.0))
	reg += strconv.Itoa(digit2)
	offset = offset % 10111

	if offset <= 600 {
		// Nab, NabA..NabZ, NabAA..NabZZ
		return reg + nletters(offset, true)
	}

	// Nab0* .. Nab9*
	offset -= 601

	digit3 := int(math.Floor(float64(offset) / 951.0))
	reg += strconv.Itoa(digit3)
	offset = offset % 951

	if offset <= 600 {
		// Nabc, NabcA .. NabcZ, NabcAA .. NabcZZ
		return reg + nletters(offset, true)
	}

	// Nabc0* .. Nabc9*
	offset -= 601

	digit4 := int(math.Floor(float64(offset) / 35.0))
	reg += strconv.Itoa(digit4)
	offset = offset % 35

	if offset <= 24 {
		// Nabcd, NabcdA .. NabcdZ
		return reg + nletters(offset, false)
	}

	// Nabcd0 .. Nabcd9
	offset -= 25
	return reg + strconv.Itoa(offset)
}

func nletters(rem int, s bool) string {

	var limitedAlphabet = []string{
		"A", "B", "C", "D", "E", "F",
		"G", "H", "J", "K", "L", "M",
		"N", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y", "Z"}

	if rem == 0 {
		return ""
	}

	rem -= 1

	if !s {
		return limitedAlphabet[rem]
	}

	index := int(math.Floor(float64(rem) / 25.0))
	return limitedAlphabet[index] + nletters(rem%25, false)
}
