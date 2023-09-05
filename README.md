# healthcheck - A Go package for healthcheck

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/healthcheck)](https://pkg.go.dev/github.com/go-zoox/healthcheck)
[![Build Status](https://github.com/go-zoox/healthcheck/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/healthcheck/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/healthcheck)](https://goreportcard.com/report/github.com/go-zoox/healthcheck)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/healthcheck/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/healthcheck?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/healthcheck.svg)](https://github.com/go-zoox/healthcheck/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/healthcheck.svg?label=Release)](https://github.com/go-zoox/healthcheck/releases)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/healthcheck
```

## Getting Started

```go
package main

import (
	"context"

	"github.com/go-zoox/healthcheck"
)

func main() {
	// HTTP
	ok, err := healthcheck.HTTP("https://github.com")
	if err != nil {
		panic(err)
	}
	fmt.Println("HTTP => https://github.com", ok)

	// TCP
	ok, err = healthcheck.TCP("127.0.0.1", 443)
	if err != nil {
		panic(err)
	}
	fmt.Println("TCP => 127.0.0.1:443", ok)

	// Ping
	ok, err = healthcheck.Ping("github.com")
	if err != nil {
		panic(err)
	}
	fmt.Println("Ping => github.com", ok)

	// Keyword in HTTP response
	ok, err = healthcheck.Keyword("https://github.com", "GitHub")
	if err != nil {
		panic(err)
	}
	fmt.Println("Keyword => https://github.com", ok)
}
```

## License
GoZoox is released under the [MIT License](./LICENSE).
