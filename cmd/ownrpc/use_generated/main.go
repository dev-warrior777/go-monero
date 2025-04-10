package main

///////////////////////////////////////////////////////////////////////////////
// use wallet created with generate from keys, etc.
//
// spend key: e42a6f98f92faa98b695e7f05725bc736834280800b154fb145c778c74fbfe01
// view key: 4ada170972a137d2c2efd1c819d13e5465415677ee28350addc507c28459c900
// mnemonic_: joking sapling goat else acquire okay turnip nudged mechanic rafts bite necklace dauntless vitals withdrawn gimmick dodge dedicated moisture unhappy jigsaw colony syringe tarnished unhappy
// primary address: 47rDCGk2ndWBySuLHcfFVbEQARPAiVHcq3nbV2FuNggrMTMqDq1QdX4enkxAG1sPRvT1FyXWKtA2F8ZhweZgsCDhNyPmLb3
///////////////////////////////////////////////////////////////////////////////

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/bisoncraft/go-monero/rpc"
)

/////////////////////////
// Run Regtest Harness //
/////////////////////////

func main() {
	ctx := context.Background()

	own := rpc.New(rpc.Config{
		// own: monero-wallet-rpc server
		Address: "http://127.0.0.1:28484/json_rpc",
		Client:  &http.Client{ /*default no auth HTTP client*/ },
	})

	ow_req := &rpc.OpenWalletRequest{
		Filename: "bob",
		Password: "",
	}

	fmt.Println("Opening wallet...")

	err := own.OpenWallet(ctx, ow_req)
	if err != nil {
		fmt.Printf("open: %v\n", err)
		own.CloseWallet(ctx)
		os.Exit(1)
	}
	fmt.Println("OPENED!")

	rf_req := &rpc.RefreshRequest{
		// (Optional) The block height from which to start.
		// Real wallet put wallet birthday if you know it
		StartHeight: 0,
	}

	rf_resp, err := own.Refresh(ctx, rf_req)
	if err != nil {
		fmt.Printf("refresh: %v\n", err)
		own.CloseWallet(ctx)
		os.Exit(2)
	}
	fmt.Println("REFRESHED!")
	fmt.Printf("blocks fetched from monerod: %d\n", rf_resp.BlocksFetched)
	fmt.Printf("did we get money? %v\n", rf_resp.ReceivedMoney)

	gb_req := &rpc.GetBalanceRequest{
		AccountIndex: 0,
	}

	gb_resp, err := own.GetBalance(ctx, gb_req)
	if err != nil {
		fmt.Printf("refresh: %v\n", err)
		own.CloseWallet(ctx)
		os.Exit(2)
	}
	fmt.Printf("balance: %d\n", gb_resp.Balance)
	fmt.Printf("unlocked balance: %d\n", gb_resp.UnlockedBalance)

	//////////////////////
	// Use the wallet here
	//////////////////////
	fmt.Println("DO STUFF!")

	own.CloseWallet(ctx)

	fmt.Println("CLOSED!")
}
