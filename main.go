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

	"github.com/dev-warrior777/go-monero/old_rpc"
	"github.com/dev-warrior777/go-monero/rpc"
)

// Regtest
func main() {
	ctx := context.Background()

	daemon := old_rpc.New(old_rpc.Config{
		Address: "http://127.0.0.1:18081",
	})

	gtxs_req := old_rpc.GetTransactionsRequest{
		TxsHashes:    []old_rpc.TxHash{"73ed80efe6763276486232aa984dac69a6f74b17dd21ea917434237642bdb484"},
		DecodeAsJson: true,
	}
	gtxs_resp, err := daemon.GetTransactions(ctx, &gtxs_req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("status:", gtxs_resp.Status)

	srt_req := old_rpc.SendRawTransactionRequest{
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

	wallet := rpc.New(rpc.Config{
		Address: "http://127.0.0.1:28284/json_rpc", // charlie
		Client:  &http.Client{ /*default no auth HTTP client*/ },
	})

	// charlie sends 0.000333 XMR -> fred
	tr_req := rpc.TransferRequest{
		Destinations: []rpc.Destination{
			{
				Amount:  333000000,
				Address: "494aSG3QY1C4PJf7YyDjFc9n2uAuARWSoGQ3hrgPWXtEjgGrYDn2iUw8WJP5Dzm4GuMkY332N9WfbaKfu5tWM3wk8ZeSEC5",
			},
		},
		AccountIndex:   0,
		SubaddrIndices: []uint64{0},
		Priority:       0,
		RingSize:       11,
		UnlockTime:     20,
		GetTxKey:       true,
	}
	tr_resp, err := wallet.Transfer(ctx, &tr_req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx_hash:", tr_resp.TxHash)

	client := rpc.New(rpc.Config{
		Address: "http://127.0.0.1:18081/json_rpc",
		Client:  &http.Client{ /*default no auth HTTP client*/ },
	})

	dgi_resp, err := client.DaemonGetInfo(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("status:", dgi_resp.Status)

	drg_req := rpc.DemonRegtestGenerateRequest{
		AmountOfBlocks: 1,
		WalletAddress:  "494aSG3QY1C4PJf7YyDjFc9n2uAuARWSoGQ3hrgPWXtEjgGrYDn2iUw8WJP5Dzm4GuMkY332N9WfbaKfu5tWM3wk8ZeSEC5",
		StartingNonce:  0,
	}
	drg_resp, err := client.DaemonRegtestGenerate(ctx, &drg_req)
	if err != nil {
		// if not regtest daemon will say:
		// "json2.Error {Code: -13, Message: "Regtest required when generating blocks"}"
		log.Fatal(err)
	}
	fmt.Println("status:", drg_resp.Status)
}
