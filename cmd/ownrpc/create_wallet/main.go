package main

/////////////////////////////////////////////////////////
// Create an initial random wallet
//
// - creates inside the dextest/xmr/wallets/own directory
/////////////////////////////////////////////////////////

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dev-warrior777/go-monero/rpc"
)

/////////////////////////
// Run Regtest Harness //
/////////////////////////

func main() {
	ctx := context.Background()

	own := rpc.New(rpc.Config{
		// own rpc port
		Address: "http://127.0.0.1:28884/json_rpc",
		Client:  &http.Client{ /*default no auth HTTP client*/ },
	})

	cw_req := &rpc.CreateWalletRequest{
		Filename: "own",
		Password: "",
		Language: "English",
	}

	// once only!
	// .. Or
	// delete dextest/xmr/wallets/own/own
	// delete dextest/xmr/wallets/own/own.keys
	err := own.CreateWallet(ctx, cw_req)
	if err != nil {
		fmt.Printf("create wallet: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("wallet created")

	qk_spend_req := &rpc.QueryKeyRequest{
		KeyType: "spend_key",
	}
	qk_resp, err := own.QueryKey(ctx, qk_spend_req)
	if err != nil {
		fmt.Printf("query spend key: %v\n", err)
		own.CloseWallet(ctx)
		os.Exit(2)
	}
	fmt.Printf("spend key: %s\n", qk_resp.Key)

	qk_view_req := &rpc.QueryKeyRequest{
		KeyType: "view_key",
	}
	qk_view_resp, err := own.QueryKey(ctx, qk_view_req)
	if err != nil {
		fmt.Printf("query view key: %v\n", err)
		own.CloseWallet(ctx)
		os.Exit(3)
	}
	fmt.Printf("view key: %s\n", qk_view_resp.Key)

	qk_mnemonic_req := &rpc.QueryKeyRequest{
		KeyType: "mnemonic",
	}
	qk_mnemonic_resp, err := own.QueryKey(ctx, qk_mnemonic_req)
	if err != nil {
		fmt.Printf("query mnemonic: %v\n", err)
		own.CloseWallet(ctx)
		os.Exit(4)
	}
	fmt.Printf("mnemonic_: %s\n", qk_mnemonic_resp.Key)

	ga_req := &rpc.GetAddressRequest{
		AccountIndex: 0,
	}

	ga_resp, err := own.GetAddress(ctx, ga_req)
	if err != nil {
		fmt.Printf("get address: %v\n", err)
		own.CloseWallet(ctx)
		os.Exit(5)
	}
	fmt.Printf("primary address: %s\n", ga_resp.Address)

	own.CloseWallet(ctx)
}
