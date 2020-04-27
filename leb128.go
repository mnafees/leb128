package leb128

import "math/bits"

// Based on the explanation here: https://en.wikipedia.org/wiki/LEB128.

// UnsignedEncode encodes an uint64 to LEB128 encoded byte array
func UnsignedEncode(value uint64) []byte {
	var enc []byte
	for value > 0 {
		bits := byte(value & 0x7F)
		if value>>7 > 0 {
			bits |= 0x80
		}
		enc = append(enc, bits)
		value >>= 7
	}
	return enc
}

// UnsignedDecode decodes a LEB128 encoded byte array back to uint64
func UnsignedDecode(enc []byte) uint64 {
	if len(enc) > 8 {
		panic("Error decoding byte array. Cannot fit in uint64.")
	}

	var dec uint64
	for i := len(enc) - 1; i >= 0; i-- {
		bits := enc[i] & 0x7F
		dec = (dec << 7) | uint64(bits)
	}
	return dec
}

// SignedEncode encodes an int64 to LEB128 encoded byte array
func SignedEncode(value int64) []byte {
	if value >= 0 {
		return UnsignedEncode(uint64(value))
	}

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
	panic("Not implemented yet")
}
