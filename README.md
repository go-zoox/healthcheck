# pinyin - 汉字转拼音

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/healthcheck)](https://pkg.go.dev/github.com/go-zoox/healthcheck)
[![Build Status](https://github.com/go-zoox/healthcheck/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/healthcheck/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/healthcheck)](https://goreportcard.com/report/github.com/go-zoox/healthcheck)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/pinyin/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/pinyin?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/pinyin.svg)](https://github.com/go-zoox/healthcheck/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/pinyin.svg?label=Release)](https://github.com/go-zoox/healthcheck/releases)

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
	fmt.Println(pinyin.String("你好世界"))
	// Output: nihaoshijie

	fmt.Println(pinyin.Slice("你好世界"))
	// Output: [ni hao shi jie]

	fmt.Println(pinyin.Abbr("你好世界"))
	// Output: nhsj
}
```

## License
GoZoox is released under the [MIT License](./LICENSE).
