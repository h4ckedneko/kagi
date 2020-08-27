# kagi (かぎ)

[![GoDoc Reference](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/h4ckedneko/kagi?tab=doc)
[![Latest Version](https://img.shields.io/github/v/release/h4ckedneko/kagi?label=latest)](https://github.com/h4ckedneko/kagi/releases)
[![License Name](https://img.shields.io/github/license/h4ckedneko/kagi?color=blue)](https://github.com/h4ckedneko/kagi/blob/master/LICENSE)
[![Build Status](https://img.shields.io/github/workflow/status/h4ckedneko/kagi/Testing)](https://github.com/h4ckedneko/kagi/actions?query=workflow:Testing)
[![Coverage Status](https://gocover.io/_badge/github.com/h4ckedneko/kagi)](https://gocover.io/github.com/h4ckedneko/kagi)
[![Report Card Status](https://goreportcard.com/badge/github.com/h4ckedneko/kagi)](https://goreportcard.com/report/github.com/h4ckedneko/kagi)

Package kagi is a simple utility to manage application keys in Go. It enables the application to generate and use a secure key for production, while allowing the developers to use hardcoded key that can be shared for the whole team during development.

**Features and benefits:**

-   Minimal API, it only has two functions.
-   Generates cryptographically secure keys.
-   Supports either generated or hardcoded keys.
-   No imported external dependencies.
-   Properly tested with benchmarks.

## Installation

Make sure you have a working [Go](https://golang.org/doc/install) workspace, then:

```
go get github.com/h4ckedneko/kagi
```

For updating to latest stable release, do:

```
go get -u github.com/h4ckedneko/kagi
```

## Usage

Here is a basic example for this package:

```go
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
```

See [examples](https://github.com/h4ckedneko/kagi/tree/master/examples) for more advanced real-world examples.

## Performance

You can run benchmarks by yourself using `make bench` command.

```
BenchmarkNew8-2                  1003131              1168 ns/op              72 B/op          4 allocs/op
BenchmarkNew16-2                 1033360              1193 ns/op             112 B/op          4 allocs/op
BenchmarkNew32-2                  972768              1323 ns/op             192 B/op          4 allocs/op
BenchmarkNew64-2                  627580              1802 ns/op             352 B/op          4 allocs/op
BenchmarkDecodeGenerated-2       8482894               124 ns/op               8 B/op          1 allocs/op
BenchmarkDecodeHardcoded-2      17643516              63.0 ns/op               8 B/op          1 allocs/op
```

## License

MIT © Lyntor Paul Figueroa. See [LICENSE](https://github.com/h4ckedneko/kagi/blob/master/LICENSE) for full license text.
