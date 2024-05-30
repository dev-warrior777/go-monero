# go-monero

[![GoDoc](https://godoc.org/gitlab.com/moneropay/go-monero/walletrpc?status.svg)](https://godoc.org/gitlab.com/moneropay/go-monero/walletrpc)

This package provides Golang bindings for monero-wallet-rpc calls.
It also contains functions for XMR/atomic unit conversion.

Unlike other wallet RPC bindings, this package is actively maintained and it is complete
with all the RPC methods. While using this package for our other project [MoneroPay](https://gitlab.com/moneropay/moneropay),
we have realized that all the other forks don't handle optional parameters correctly and send them anyway,
therefore causing bugs.

## This Fork

Thanks to the developers of MoneroPay!

This fork is based on the above and is for integration into dex. It updates by introducing daemon rpc functions with both the older RPC's (see the README.md in `old_rpc` directory) and the newer JSON 2.0 RPC calls (see the README.md in `rpc` directory)

Tested on `stagenet` and `regtest`

## MoneroPay Example

Mainnet:

```sh
# monero-wallet-rpc --detach \
#  --rpc-bind-port 18083 \
#  --wallet-file /home/moneropay/wallet \
#  --password s3cure \
#  --daemon-login kernal:s3cure \
#  --rpc-login kernal:s3cure
```

```golang
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"

    "github.com/gabstv/httpdigest"
)

func main() {
    // username: kernal, password: s3cure
    client := walletrpc.New(walletrpc.Config{
        Address: "http://127.0.0.1:18083/json_rpc",
        Client: &http.Client{
            Transport: httpdigest.New("kernal", "s3cure"), // Remove if no auth; use default round-tripper
        },
    })
    resp, err := client.GetBalance(context.Background(), &walletrpc.GetBalanceRequest{})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Total balance:", walletrpc.XMRToDecimal(resp.Balance))
    fmt.Println("Unlocked balance:", walletrpc.XMRToDecimal(resp.UnlockedBalance))
}
```
