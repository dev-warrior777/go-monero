//
// Example:
// =======
//
// daemon stagenet: detached
//
// monero-wallet-rpc --detach \
//  --stagenet \
// 	--rpc-bind-port 38081 \
// 	--wallet-file /home/moneropay/wallet \
// 	--password s3cure \
// 	--daemon-login kernal:s3cure \
// 	--rpc-login kernal:s3cure
//
// wallet server stagenet: un-detached
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

	"github.com/dev-warrior777/go-monero/daemonrpc"
	"github.com/dev-warrior777/go-monero/walletrpc"
)

func main() {
	ctx := context.Background()

	daemon := daemonrpc.New(daemonrpc.Config{
		Address: "http://127.0.0.1:38081",
	})

	gtxs_req := daemonrpc.GetTransactionsRequest{
		TxsHashes:    []daemonrpc.TxHash{"c8cc2347935ca416ea8b3d0c83061f367be0536948563c694cb272958e02ab71"},
		DecodeAsJson: false,
	}
	gtxs_resp, err := daemon.GetTransactions(ctx, &gtxs_req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("status:", gtxs_resp.Status)

	srt_req := daemonrpc.SendRawTransactionRequest{
		TxAsHex:    gtxs_resp.TxsAsHex[0],
		DoNotRelay: false,
	}
	srt_resp, err := daemon.SendRawTransaction(ctx, &srt_req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("status:", srt_resp.Status)

	gtp_resp, err := daemon.GetTransactionPool(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("status:", gtp_resp.Status)

	wallet_client := walletrpc.New(walletrpc.Config{
		Address: "http://127.0.0.1:38383/json_rpc",
		Client:  &http.Client{ /*default no auth HTTP client*/ },
	})
	// wallet_resp, err := wallet_client.GetBalance(ctx, &walletrpc.GetBalanceRequest{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Total balance:", walletrpc.XMRToDecimal(wallet_resp.Balance))
	// fmt.Println("Unlocked balance:", walletrpc.XMRToDecimal(wallet_resp.UnlockedBalance))
}
