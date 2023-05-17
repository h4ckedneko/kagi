# kagi (かぎ)

[![Godoc Reference](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/h4ckedneko/kagi)
[![Latest Version](https://img.shields.io/github/v/release/h4ckedneko/kagi?label=latest)](https://github.com/h4ckedneko/kagi/releases)
[![License Name](https://img.shields.io/github/license/h4ckedneko/kagi?color=blue)](https://github.com/h4ckedneko/kagi/blob/master/LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/h4ckedneko/kagi/testing.yml)](https://github.com/h4ckedneko/kagi/actions/workflows/testing.yml)
[![Coverage Status](https://img.shields.io/codecov/c/github/h4ckedneko/kagi)](https://app.codecov.io/gh/h4ckedneko/kagi)
[![Go Report Card Status](https://goreportcard.com/badge/github.com/h4ckedneko/kagi)](https://goreportcard.com/report/github.com/h4ckedneko/kagi)

Package kagi is a simple Go utility for managing application keys. It enables the application to generate and use a secure key on production, while allowing the developers to use hardcoded key that can be shared for the whole team during development.

**Features and benefits:**

-   Minimal API, it has only two functions.
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
goos: linux
goarch: amd64
pkg: github.com/h4ckedneko/kagi
cpu: AMD Ryzen 5 5600G with Radeon Graphics
BenchmarkNew8-12                         3424884                374.7 ns/op           64 B/op          4 allocs/op
BenchmarkNew16-12                        3072880                406.5 ns/op           96 B/op          4 allocs/op
BenchmarkNew32-12                        2774564                459.2 ns/op          192 B/op          4 allocs/op
BenchmarkNew64-12                        1826727                648.1 ns/op          352 B/op          4 allocs/op
BenchmarkDecodeGenerated-12             33482005                33.15 ns/op            8 B/op          1 allocs/op
BenchmarkDecodeHardcoded-12             63619456                20.19 ns/op            8 B/op          1 allocs/op
BenchmarkDecodeStringGenerated-12       31331952                39.12 ns/op            8 B/op          1 allocs/op
BenchmarkDecodeStringHardcoded-12       50978620                26.69 ns/op            8 B/op          1 allocs/op
PASS
ok      github.com/h4ckedneko/kagi      11.968s
```

## License

MIT © Lyntor Paul Figueroa. See [LICENSE](https://github.com/h4ckedneko/kagi/blob/master/LICENSE) for full license text.
