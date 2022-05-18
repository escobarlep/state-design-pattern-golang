# Golang studies
## _State Desing Pattern in Golang_

This repository represent an application for a commom order life-cicle

## Features

- Initialize order in Pending Payment state
- Cicle through payment, prepare, ship and cancel state

## Tech

- [Golang](https://go.dev/) - Go was created at Google in 2007, and since then, engineering teams across Google have adopted Go to build products and services at massive scale. !

## Execution

This application requires [Golang](https://go.dev/dl/) to run.

To run this app, simply run:
```sh
go run main.go
```

## Expected output

```sh

Actual state is *main.PaymentPendingState

Cannot ship the order when payment is pending!

Payment verify! Shipping soon.

Actual state is *main.OrderBeingPreparedState

Payment already verified.

Actual state is *main.OrderBeingPreparedState

Shipping your order now.

Actual state is *main.ShippedOrderState

Actual state is *main.PaymentPendingState

Cannot ship the order when payment is pending!

Cancelling your unpaid order

Actual state is *main.CancelledOrderState
```

## License

MIT

**Keep Studying!**
