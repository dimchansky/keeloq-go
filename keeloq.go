package keeloq

// 1-bit lookup table for the NLF
const (
	lut    = uint32(0x3A5C742E)
	cycles = 528
)

// Encrypt encrypts a 32-bit block of plaintext using the KeeLoq algorithm.
// `block`: 32-bit plaintext block
// `key`: 64-bit key
// returns 32-bit ciphertext block
func Encrypt(block uint32, key uint64) uint32 {
	for i := 0; i < cycles; i++ {
		// Calculate LUT key
		lutKey := (block>>1)&1 | (block>>8)&2 | (block>>18)&4 | (block>>23)&8 | (block>>27)&16

		// Calculate next bit to feed
		msb := (block >> 16 & 1) ^ (block & 1) ^ (lut >> lutKey & 1) ^ uint32(key&1)

		// Feed it
		block = msb<<31 | block>>1

		// Rotate key right
		key = (key&1)<<63 | key>>1
	}

	return block
}

// Decrypt decrypts a 32-bit block of ciphertext using the KeeLoq algorithm.
// block: 32-bit ciphertext block
// key: 64-bit key
// returns 32-bit plaintext block
func Decrypt(block uint32, key uint64) uint32 {
	for i := 0; i < cycles; i++ {
		// Calculate LUT key
		lutKey := (block>>0)&1 | (block>>7)&2 | (block>>17)&4 | (block>>22)&8 | (block>>26)&16

		// Calculate next bit to feed
		lsb := (block >> 31) ^ (block >> 15 & 1) ^ (lut >> lutKey & 1) ^ uint32(key>>15&1)

		// Feed it
		block = (block&0x7FFFFFFF)<<1 | lsb

		// Rotate key left
		key = (key&0x7FFFFFFFFFFFFFFF)<<1 | key>>63
	}

	return block
}
