package kagi_test

import (
	"encoding/base64"
	"testing"

	"github.com/h4ckedneko/kagi"
	"github.com/stretchr/testify/assert"
)

func testKeys(t *testing.T, keys []string, size int) {
	for i1, key1 := range keys {
		assert.Equal(t, size, len(kagi.Decode(key1)))
		for i2, key2 := range keys {
			if i1 != i2 {
				assert.NotEqual(t, key1, key2)
			}
		}
	}
}

func TestNew8(t *testing.T) {
	keys := make([]string, 1000)
	for i := range keys {
		keys[i] = kagi.New8()
	}
	testKeys(t, keys, 8)
}

func TestNew16(t *testing.T) {
	keys := make([]string, 1000)
	for i := range keys {
		keys[i] = kagi.New16()
	}
	testKeys(t, keys, 16)
}

func TestNew32(t *testing.T) {
	keys := make([]string, 1000)
	for i := range keys {
		keys[i] = kagi.New32()
	}
	testKeys(t, keys, 32)
}

func TestNew64(t *testing.T) {
	keys := make([]string, 1000)
	for i := range keys {
		keys[i] = kagi.New64()
	}
	testKeys(t, keys, 64)
}

func TestDecodeGenerated(t *testing.T) {
	key := []byte{'a', 'b', 'c', '1', '2', '3'}
	keyd := kagi.Decode("base64:" + base64.RawStdEncoding.EncodeToString(key))
	assert.Equal(t, key, keyd)
}

func TestDecodeHardcoded(t *testing.T) {
	key := []byte{'a', 'b', 'c', '1', '2', '3'}
	keyd := kagi.Decode(string(key))
	assert.Equal(t, key, keyd)
}

func BenchmarkNew8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kagi.New8()
	}
}

func BenchmarkNew16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kagi.New16()
	}
}

func BenchmarkNew32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kagi.New32()
	}
}

func BenchmarkNew64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kagi.New64()
	}
}

func BenchmarkDecodeGenerated(b *testing.B) {
	key := "base64:YWJjMTIz"
	for i := 0; i < b.N; i++ {
		kagi.Decode(key)
	}
}

func BenchmarkDecodeHardcoded(b *testing.B) {
	key := "abc123"
	for i := 0; i < b.N; i++ {
		kagi.Decode(key)
	}
}
