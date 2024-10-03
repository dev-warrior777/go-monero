package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dev-warrior777/go-monero/rpc"
)

// turnip arises splendid obtains southern nylon nuns lakes
// aggravate wounded depth bugs vials yawning justice oval
// hire sleepless being rounded dolphin echo puck rapid nuns
// 55ps81tfB2JTdbHXYVJuZVeYCYagjLjBPgHtH6DRHXZ3eMLLtE7FECTMmGzJmFVqPz75KsVcVdGMfej8grDiEx5KDmU7NBA
// secret: 1c8eb5be5d2488caa8b9058d23e709f0e8a8677834f07e6702889979c4625304
// public: 696d1b57e07f939f3733fb16fc10f4e06cadd6c74e43faeae7faae9a95073adf
// secret: 894b7da3b33a9a35bd434461a70a56c519bcef11268a6efc5fb9eb3c83797c01
// public: 4df48d1418985c7c22d75c01243e752451486f0bb7c80ae18d322155e82d8e71

var spendKey = "1c8eb5be5d2488caa8b9058d23e709f0e8a8677834f07e6702889979c4625304"
var viewKey = "894b7da3b33a9a35bd434461a70a56c519bcef11268a6efc5fb9eb3c83797c01"
var primaryAddress = "55ps81tfB2JTdbHXYVJuZVeYCYagjLjBPgHtH6DRHXZ3eMLLtE7FECTMmGzJmFVqPz75KsVcVdGMfej8grDiEx5KDmU7NBA"

//////////////////
// Run Stagenet //
//////////////////

func main() {
	ctx := context.Background()

	aliceRpc := rpc.New(rpc.Config{
		// Alice: monero-wallet-rpc server
		Address: "http://127.0.0.1:38284/json_rpc",
		Client:  &http.Client{ /*default no auth HTTP client*/ },
	})

	// Generate from primary address, spend key, view key above
	gfk_req := &rpc.GenerateFromKeysRequest{
		RestoreHeight: 1600000, // birthday if known
		Filename:      "alice",
		Address:       primaryAddress,
		SpendKey:      spendKey,
		ViewKey:       viewKey,
		Password:      "",
	}

	gfk_resp, err := aliceRpc.GenerateFromKeys(ctx, gfk_req)
	if err != nil {
		fmt.Printf("generate from keys: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("primary address: %s\n", gfk_resp.Address)
	fmt.Printf("info: %s\n", gfk_resp.Info)

	rf_req := &rpc.RefreshRequest{
		// (Optional) The block height from which to start.
		// Real wallet put wallet birthday if you know it
		StartHeight: 1600000 - 1,
	}

	rf_resp, err := aliceRpc.Refresh(ctx, rf_req)
	if err != nil {
		fmt.Printf("refresh: %v\n", err)
		aliceRpc.CloseWallet(ctx)
		os.Exit(2)
	}
	fmt.Println("REFRESHED!")
	fmt.Printf("blocks fetched from monerod: %d\n", rf_resp.BlocksFetched)
	fmt.Printf("did we get money? %v\n", rf_resp.ReceivedMoney)

	gb_req := &rpc.GetBalanceRequest{
		AccountIndex: 0,
	}

	gb_resp, err := aliceRpc.GetBalance(ctx, gb_req)
	if err != nil {
		fmt.Printf("getbalance: %v\n", err)
		aliceRpc.CloseWallet(ctx)
		os.Exit(2)
	}
	fmt.Printf("balance: %d\n", gb_resp.Balance)
	fmt.Printf("unlocked balance: %d\n", gb_resp.UnlockedBalance)

	//////////////////////
	// Use the wallet here
	//////////////////////
	fmt.Println("DO STUFF!")

	aliceRpc.CloseWallet(ctx)
}
