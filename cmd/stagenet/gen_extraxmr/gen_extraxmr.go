package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dev-warrior777/go-monero/rpc"
)

// spend key: c6d38a6fcb21818efa45d7069269adc63be8339f950b29bea7c60938beb76e0c
// view key: c05281200f0521e0aa9e61bdd2691da9addd641f28ebea75fd5db683efc35204
// mnemonic_: moment guarded smog obnoxious twice jerseys certain names nozzle innocent coils vibrate oars tepid kiosk otter gasp arsenic bomb ridges vigilant nirvana knife madness names
// primary address: 5AxjLzo9oL7NcoM6gFe6HzVZMvEDoisnT8mFsjd2dMBJSK4tWcSutCKgq2kRG1oGSD2CLZ3JFibhtB5RMrDHtH36Vn2aPPb

var spendKey = "c6d38a6fcb21818efa45d7069269adc63be8339f950b29bea7c60938beb76e0c"
var viewKey = "c05281200f0521e0aa9e61bdd2691da9addd641f28ebea75fd5db683efc35204"
var primaryAddress = "5AxjLzo9oL7NcoM6gFe6HzVZMvEDoisnT8mFsjd2dMBJSK4tWcSutCKgq2kRG1oGSD2CLZ3JFibhtB5RMrDHtH36Vn2aPPb"

//////////////////
// Run Stagenet //
//////////////////

func main() {
	ctx := context.Background()

	extraRpc := rpc.New(rpc.Config{
		// Alice: monero-wallet-rpc server
		Address: "http://127.0.0.1:38484/json_rpc",
		Client:  &http.Client{ /*default no auth HTTP client*/ },
	})

	// Generate from primary address, spend key, view key above
	gfk_req := &rpc.GenerateFromKeysRequest{
		RestoreHeight: 1600000, // birthday if known
		Filename:      "extraxmr",
		Address:       primaryAddress,
		SpendKey:      spendKey,
		ViewKey:       viewKey,
		Password:      "",
	}

	gfk_resp, err := extraRpc.GenerateFromKeys(ctx, gfk_req)
	if err != nil {
		fmt.Printf("generate from keys: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("primary address: %s\n", gfk_resp.Address)
	fmt.Printf("info: %s\n", gfk_resp.Info)

	extraRpc.CloseWallet(ctx)
}
