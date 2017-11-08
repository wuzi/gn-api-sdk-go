# gn-api-sdk-go

> A go library for integration of your backend with the payment services
provided by [Gerencianet](http://gerencianet.com.br).

[![Build Status](https://travis-ci.org/gerencianet/gn-api-sdk-go.svg)](https://travis-ci.org/gerencianet/gn-api-sdk-go)
[![Coverage Status](https://coveralls.io/repos/github/gerencianet/gn-api-sdk-go/badge.svg?branch=master)](https://coveralls.io/github/gerencianet/gn-api-sdk-go?branch=master)
[![Code Climate](https://codeclimate.com/github/gerencianet/gn-api-sdk-go/badges/gpa.svg)](https://codeclimate.com/github/gerencianet/gn-api-sdk-go)

## Installation

Install with:

```bash
$ go get github.com/gerencianet/gn-api-sdk-go/gerencianet
```
## Tested with
```
go 1.8
```
## Basic usage

```go

import (
    "github.com/gerencianet/gn-api-sdk-go/gerencianet"
)

credentials := map[string]interface{} {
    "client_id": "client_id",
    "client_secret": "client_secret",
    "sandbox": true,
    "timeout": 10,
}

gn := gerencianet.NewGerencianet(credentials)

body = {
    "items": [{
        "name": "Product 1",
        "value": 1000,
        "amount": 2,
    }],
    "shippings": [{
        "name": "Default Shipping Cost",
        "value": 100,
    }]
}

res, err := gn.CreateCharge(body)

```

## Examples

You can run the examples inside `_examples` with
`$ go run example.go`:

```bash
$ go run charge/create_charge.go
```

Just remember to set the correct credentials inside `examples/configs.go` before running.

## Tests

To run the tests, just run:

```bash
$ go test
```

## Additional documentation

The full documentation with all available endpoints is in https://dev.gerencianet.com.br/.

## Changelog

[CHANGELOG](CHANGELOG.md)

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/gerencianet/gn-api-sdk-go. This project is intended to be a safe, welcoming space for collaboration.

## License

The library is available as open source under the terms of the [MIT License](LICENSE).
