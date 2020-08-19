package main

import (
	"fmt"

	"github.com/h4ckedneko/kagi"
)

func main() {
	// Generate a new secure key.
	key := kagi.New(32)
	fmt.Println(key)

	// Decode a generated key.
	keyd := kagi.Decode(key)
	fmt.Println(string(keyd))

	// Decode a hardcoded key.
	keyd = kagi.Decode("abc123")
	fmt.Println(string(keyd))
}
