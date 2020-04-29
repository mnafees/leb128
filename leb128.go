package leb128

import (
	"math/bits"
)

// Based on the explanation here: https://en.wikipedia.org/wiki/LEB128.

// UnsignedEncode encodes an uint64 to LEB128 encoded byte array
func UnsignedEncode(value uint64) []byte {
	if value == 0 { // Special case
		return []byte{0x00}
	}

	var enc []byte
	for value > 0 {
		bits := byte(value & 0x7F) // Extract the last 7 bits
		if value>>7 > 0 {          // Add high 1 bits on all but last
			bits |= 0x80
		}
		enc = append(enc, bits)
		value >>= 7 // Shift value right by 7 bits
	}
	return enc
}

// UnsignedDecode decodes a LEB128 encoded byte array back to uint64
func UnsignedDecode(enc []byte) uint64 {
	if len(enc) > 8 { // We expect a 64 bit unsigned number
		panic("Error decoding byte array. Cannot fit in uint64.")
	}

	if len(enc) == 1 && enc[0] == 0x00 { // Special case
		return 0
	}

	var dec uint64
	for i := len(enc) - 1; i >= 0; i-- {
		bits := enc[i] & 0x7F // Extract the last 7 bits
		dec = (dec << 7) | uint64(bits)
	}
	return dec
}

// SignedEncode encodes an int64 to LEB128 encoded byte array
func SignedEncode(value int64) []byte {
	if value >= 0 {
		return UnsignedEncode(uint64(value))
	}

	// Manually try to convert to 2's complement
	uvalue := uint64(-value)
	bitLength := bits.Len64(uvalue)
	i := 7
	for {
		if i > bitLength {
			bitLength = i
			break
		}
		i += 7
	}
	uvalue = ^uvalue
	uvalue++

	var enc []byte
	for i = 7; i <= bitLength; i += 7 {
		bits := byte(uvalue & 0x7F)
		// Add high 1 bits on all but last
		if i == bitLength {
			bits |= 0x00
		} else {
			bits |= 0x80
		}
		enc = append(enc, bits)
		uvalue >>= 7
	}
	return enc
}

// SignedDecode decodes a LEB128 encoded byte array back to int64
func SignedDecode(enc []byte) int64 {
	if len(enc) > 8 {
		panic("Error decoding byte array. Cannot fit in int64.")
	}

	// Use UnsignedDecode for positive numbers
	if enc[len(enc)-1]&0x40 != 0x40 {
		return int64(UnsignedDecode(enc))
	}

	var dec uint64
	for i := len(enc) - 1; i >= 0; i-- {
		bits := enc[i] & 0x7F
		dec = (dec << 7) | uint64(bits)
	}
	// Convert from 2's complement
	bitLength := bits.Len64(dec)
	dec = ^dec
	dec = (dec << (64 - bitLength)) >> (64 - bitLength)
	dec++
	return int64(-dec)
}
