// Package kagi is a simple Go utility for managing application keys.
//
// It enables the application to generate and use a secure key on
// production, while allowing the developers to use hardcoded key
// that can be shared for the whole team during development.
//
// Here is a basic example for this package:
//
// 	package main
//
// 	import (
// 		"fmt"
//
// 		"github.com/h4ckedneko/kagi"
// 	)
//
// 	func main() {
// 		// Generate a new secure key.
// 		key := kagi.New(32)
// 		fmt.Println(key)
//
// 		// Decode a generated key.
// 		keyd := kagi.Decode(key)
// 		fmt.Println(string(keyd))
//
// 		// Decode a hardcoded key.
// 		keyd = kagi.Decode("abc123")
// 		fmt.Println(string(keyd))
// 	}
//
// See https://github.com/h4ckedneko/kagi/tree/master/examples for more advanced real-world examples.
//
package kagi
