# Mailcow API client package for Go
[![PkgGoDev](https://pkg.go.dev/badge/github.com/jpbede/mailcowgo)](https://pkg.go.dev/github.com/jpbede/mailcowgo)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/96fda22fb4b94bde9975e247d0e6dc27)](https://www.codacy.com/gh/jpbede/mailcowgo/dashboard)
[![codecov](https://codecov.io/gh/jpbede/mailcowgo/branch/main/graph/badge.svg?token=1262YRM8ME)](https://codecov.io/gh/jpbede/mailcowgo)
![test](https://github.com/jpbede/mailcowgo/workflows/test/badge.svg)

This repository contains a Go package for accessing the [Mailcow API](https://mailcow.docs.apiary.io/).

## Installation
Install using go get:
```shell
go get github.com/jpbede/mailcowgo
```

## Usage
Import the lib as usual
```go
import "github.com/jpbede/mailcowgo"
```

Create a new client without options:
```go
mailcowClient, err := mailcow.New("https://my.awesome.email", "api key")
```
The client now offers more specific sub-clients, for example for managing domains or mailboxes.
For more information have a look at the [documentation](https://pkg.go.dev/github.com/jpbede/mailcowgo).
