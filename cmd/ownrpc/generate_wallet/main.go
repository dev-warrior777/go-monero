package main

///////////////////////////
// generate from keys, etc.
///////////////////////////

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dev-warrior777/go-monero/rpc"
)

// Generated from create wallet
var spendKey = "e42a6f98f92faa98b695e7f05725bc736834280800b154fb145c778c74fbfe01"
var viewKey = "4ada170972a137d2c2efd1c819d13e5465415677ee28350addc507c28459c900"
var primaryAddress = "47rDCGk2ndWBySuLHcfFVbEQARPAiVHcq3nbV2FuNggrMTMqDq1QdX4enkxAG1sPRvT1FyXWKtA2F8ZhweZgsCDhNyPmLb3"

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

	// https: //web.getmonero.org/resources/developer-guides/wallet-rpc.html#generate_from_keys

	// Generate from primary address, spend key, view key above
	gfk_req := &rpc.GenerateFromKeysRequest{
		// for a real wallet this should be the birthday if known
		RestoreHeight: 0,
		Filename:      "bob",
		Address:       primaryAddress,
		SpendKey:      spendKey,
		ViewKey:       viewKey,
		Password:      "",
	}

	gfk_resp, err := own.GenerateFromKeys(ctx, gfk_req)
	if err != nil {
		fmt.Printf("generate from keys: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("primary address: %s\n", gfk_resp.Address)
	fmt.Printf("info: %s\n", gfk_resp.Info)

	own.CloseWallet(ctx)
}
