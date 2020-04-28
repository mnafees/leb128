package leb128

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnsignedEncode(t *testing.T) {
	assert.Equal(t, []byte{0x81, 0x01}, UnsignedEncode(129))
	assert.Equal(t, []byte{0xE5, 0x8E, 0x26}, UnsignedEncode(624485))
	assert.Equal(t, []byte{0x80, 0x80, 0x80, 0xFD, 0x07}, UnsignedEncode(2141192192))
}

func TestUnsignedDecode(t *testing.T) {
	assert.Equal(t, uint64(129), UnsignedDecode([]byte{0x81, 0x01}))
	assert.Equal(t, uint64(624485), UnsignedDecode([]byte{0xE5, 0x8E, 0x26}))
	assert.Equal(t, uint64(2141192192), UnsignedDecode([]byte{0x80, 0x80, 0x80, 0xFD, 0x07}))
}

func TestSignedEncode(t *testing.T) {
	assert.Equal(t, []byte{0xC0, 0xBB, 0x78}, SignedEncode(-123456))
	assert.Equal(t, []byte{0x9B, 0xF1, 0x59}, SignedEncode(-624485))
	assert.Equal(t, []byte{0x80, 0x80, 0x80, 0xFD, 0x07}, SignedEncode(2141192192))
}

func TestSignedDecode(t *testing.T) {
	assert.Equal(t, int64(-123456), SignedDecode([]byte{0xC0, 0xBB, 0x78}))
	assert.Equal(t, int64(-624485), SignedDecode([]byte{0x9B, 0xF1, 0x59}))
	assert.Equal(t, int64(2141192192), SignedDecode([]byte{0x80, 0x80, 0x80, 0xFD, 0x07}))
}
