# A Supermarket Checkout tool



[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/github.com/web-slinger/go-checkout)

This repository contains a package that demonstrates common functionality of a supermarket checkout using standard libraries.

## Clone the project

```bash
git clone https://github.com/web-slinger/go-checkout
cd go-checkout
```

## Use this project

```bash
go get github.com/web-slinger/go-checkout
```

```go
import (
    gocheckout "github.com/web-slinger/go-checkout"
)
```

```go
// get pricing model
priceScheme := gocheckout.GetPricingModel()

// initialise checkout
checkout := gocheckout.NewCheckout(priceScheme)

// scan your items
err := checkout.Scan("A")
if err != nil {
    panic(err)
}

// get total price 
totalPrice, err := checkout.GetTotalPrice()
if err != nil {
    panic(err)
}

fmt.Printf("total price '%d'", totalPrice)
```

## Testing this project

The tested methods are on the `Checkout` type, `Scan()` and `GetTotalPrice()`

```bash
go test ./...
```

Coverage currently stands at 95.5%

```bash
go test -cover ./...
```

## Mocking checkout

For some tests of your application you may want to mock certain functions, the `Checkout` type implements the `ICheckout` interface below

```go
type ICheckout interface {
    Scan(SKU string) (err error)
    GetTotalPrice() (totalPrice int, err error)
}
```

## Building this project

This project is to be used as a package to be part of an application.

```
go build -o bin/checkout
```