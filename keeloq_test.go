package keeloq_test

import (
	"fmt"
	"testing"

	"github.com/dimchansky/keeloq-go"
)

type testCase struct {
	Key    uint64
	Plain  uint32
	Cipher uint32
}

var testCases = []testCase{
	{
		Key:    0xBEEFDEADBEEFDEAD,
		Plain:  0x2000C022,
		Cipher: 0x054C90C2,
	},
	{
		Key:    0x5CEC6701B79FD949,
		Plain:  0xF741E2DB,
		Cipher: 0xE44F4CDF,
	},
	{
		Key:    0x5CEC6701B79FD949,
		Plain:  0x0CA69B92,
		Cipher: 0xA6AC0EA2,
	},
}

func TestEncrypt(t *testing.T) {
	for _, tt := range testCases {
		t.Run(fmt.Sprintf("Key 0x%0000000000000000X Plain 0x%00000000X", tt.Key, tt.Plain), func(t *testing.T) {
			if got := keeloq.Encrypt(tt.Plain, tt.Key); got != tt.Cipher {
				t.Errorf("Encrypt() = %v, want %v", got, tt.Cipher)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	for _, tt := range testCases {
		t.Run(fmt.Sprintf("Key 0x%0000000000000000X Cipher 0x%00000000X", tt.Key, tt.Cipher), func(t *testing.T) {
			if got := keeloq.Decrypt(tt.Cipher, tt.Key); got != tt.Plain {
				t.Errorf("Decrypt() = %v, want %v", got, tt.Plain)
			}
		})
	}
}

func FuzzEncryptDecrypt(f *testing.F) {
	f.Fuzz(func(t *testing.T, key uint64, plaintext uint32) {
		cipher := keeloq.Encrypt(plaintext, key)
		plaintext2 := keeloq.Decrypt(cipher, key)
		if plaintext != plaintext2 {
			t.Errorf("Failed to encrypt/decrypt plaintext 0x%00000000X using key 0x%0000000000000000X", plaintext, key)
		}
	})
}

//goland:noinspection GoUnusedGlobalVariable
var encryptResult uint32

func BenchmarkEncrypt(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encryptResult = keeloq.Encrypt(0x2000C022, 0xBEEFDEADBEEFDEAD)
	}
}

//goland:noinspection GoUnusedGlobalVariable
var decryptResult uint32

func BenchmarkDecrypt(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decryptResult = keeloq.Decrypt(0x054C90C2, 0xBEEFDEADBEEFDEAD)
	}
}
