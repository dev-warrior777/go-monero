package main

///////////////////////////
// generate from keys, etc.
///////////////////////////

// spend key: e42a6f98f92faa98b695e7f05725bc736834280800b154fb145c778c74fbfe01
// view key: 4ada170972a137d2c2efd1c819d13e5465415677ee28350addc507c28459c900
// mnemonic_: joking sapling goat else acquire okay turnip nudged mechanic rafts bite necklace dauntless vitals withdrawn gimmick dodge dedicated moisture unhappy jigsaw colony syringe tarnished unhappy
// primary address: 47rDCGk2ndWBySuLHcfFVbEQARPAiVHcq3nbV2FuNggrMTMqDq1QdX4enkxAG1sPRvT1FyXWKtA2F8ZhweZgsCDhNyPmLb3

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dev-warrior777/go-monero/rpc"
)

var spendKey = "e42a6f98f92faa98b695e7f05725bc736834280800b154fb145c778c74fbfe01"
var viewKey = "4ada170972a137d2c2efd1c819d13e5465415677ee28350addc507c28459c900"
var primaryAddress = "47rDCGk2ndWBySuLHcfFVbEQARPAiVHcq3nbV2FuNggrMTMqDq1QdX4enkxAG1sPRvT1FyXWKtA2F8ZhweZgsCDhNyPmLb3"

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

	// https://web.getmonero.org/resources/developer-guides/wallet-rpc.html#generate_from_keys
	gfk_req := &rpc.GenerateFromKeysRequest{
		RestoreHeight: 0,
		Filename:      "own",
		Address:       primaryAddress,
		SpendKey:      spendKey,
		ViewKey:       viewKey,
		Password:      "",
	}
	gfk_resp, err := own.GenerateFromKeys(ctx, gfk_req)
	if err != nil {
		fmt.Printf("generate from keys: %v\n", err)
		own.CloseWallet(ctx)
		os.Exit(1)
	}
	fmt.Printf("primary address: %s\n", gfk_resp.Address)
	fmt.Printf("info: %s\n", gfk_resp.Info)

	rf_req := &rpc.RefreshRequest{
		StartHeight: 0, // (Optional) The block height from which to start. Real wallet put this in
	}
	rf_resp, err := own.Refresh(ctx, rf_req)
	if err != nil {
		fmt.Printf("generate from keys: %v\n", err)
		own.CloseWallet(ctx)
		os.Exit(1)
	}
	fmt.Printf("blocks fetched from monerod: %d\n", rf_resp.BlocksFetched)
	fmt.Printf("did we get more money? %v\n", rf_resp.ReceivedMoney)

	own.CloseWallet(ctx)
}
