package kagi_test

import (
	"bytes"
	"encoding/base64"
	"testing"

	"github.com/h4ckedneko/kagi"
)

func testKeys(t *testing.T, keys []string, size int) {
	for i, k := range keys {
		s := len(kagi.Decode(k))
		if size != s {
			t.Errorf("length should be %d not %d", size, s)
		}
		for i2, k2 := range keys {
			if i != i2 && k == k2 {
				t.Errorf("not unique: %d:%q and %d:%q", i, k, i2, k2)
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
	if !bytes.Equal(key, keyd) {
		t.Errorf("expected %q but got %q", key, keyd)
	}
}

func TestDecodeHardcoded(t *testing.T) {
	key := []byte{'a', 'b', 'c', '1', '2', '3'}
	keyd := kagi.Decode(string(key))
	if !bytes.Equal(key, keyd) {
		t.Errorf("expected %q but got %q", key, keyd)
	}
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

func BenchmarkDecodeStringGenerated(b *testing.B) {
	key := "base64:YWJjMTIz"
	for i := 0; i < b.N; i++ {
		kagi.DecodeString(key)
	}
}

func BenchmarkDecodeStringHardcoded(b *testing.B) {
	key := "abc123"
	for i := 0; i < b.N; i++ {
		kagi.DecodeString(key)
	}
}
