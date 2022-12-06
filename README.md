# Startup

[![Go Report Card](https://goreportcard.com/badge/github.com/xuender/startup)](https://goreportcard.com/report/github.com/xuender/startup)
[![GoDoc](https://godoc.org/github.com/xuender/startup?status.svg)](https://pkg.go.dev/github.com/xuender/startup)
[![codecov](https://codecov.io/gh/xuender/startup/branch/main/graph/badge.svg?token=WHAOPUYKCM)](https://codecov.io/gh/xuender/startup)
[![GitHub license](https://img.shields.io/github/license/xuender/startup)](https://github.com/xuender/startup/blob/main/LICENSE)
[![tag](https://img.shields.io/github/tag/xuender/startup.svg)](https://github.com/xuender/startup/releases)

✨ **`xuender/startup` is a startup Go library based on Go 1.18+ Generics.**

![codecov](https://codecov.io/gh/xuender/startup/branch/main/graphs/tree.svg?token=WHAOPUYKCM)

## 🚀 Install

```sh
go get github.com/xuender/startup@latest
```

## 💡 Usage

You can import `startup` using:

```go
import "github.com/xuender/startup"
```

### Install

Then install startup:

```go
startup.Install()
// or
startup.Install(args...)
```

### Uninstall

Then uninstall startup:

```go
startup.Uninstall()
```

### Status

Then startup status:

```go
startup.Status()

// true or false
```

## 📝 License

Copyright © 2022~time.Now [Xuender](https://github.com/xuender).

This project is [MIT](./LICENSE) licensed.
