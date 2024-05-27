// Example
// monero-wallet-rpc --detach \
// 	--rpc-bind-port 18083 \
// 	--wallet-file /home/moneropay/wallet \
// 	--password s3cure \
// 	--daemon-login kernal:s3cure \
// 	--rpc-login kernal:s3cure

// stagenet: un-detached
//
// monero-wallet-rpc \
// --stagenet \
// --trusted-daemon \
// --rpc-bind-port 38383 \
// --disable-rpc-login \
// --wallet-file stagenet/wallet_03.bin \
// --password "" \
// --log-file stagenet/wallet_03.log \
// --log-level 1

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dev-warrior777/go-monero/walletrpc"
)

func main() {
	// username: kernal, password: s3cure
	client := walletrpc.New(walletrpc.Config{
		Address: "http://127.0.0.1:38383/json_rpc",
		Client:  &http.Client{
			// Transport: httpdigest.New("kernal", "s3cure"), // Remove if no auth.
		},
	})
	resp, err := client.GetBalance(context.Background(), &walletrpc.GetBalanceRequest{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total balance:", walletrpc.XMRToDecimal(resp.Balance))
	fmt.Println("Unlocked balance:", walletrpc.XMRToDecimal(resp.UnlockedBalance))
}
