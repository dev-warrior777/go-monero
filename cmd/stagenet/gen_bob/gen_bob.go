package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dev-warrior777/go-monero/rpc"
)

// spend key: f87748fb2abd3e112d07fd2bec4d6d7bb63295fa110394f19cc4b78f12dc3902
// view key: 9ca0a18760517b877b57b534c70810348f85c29e226d52b0205b684d50aa390c
// mnemonic_: ramped error elapse elbow possible ruthless duke gumball megabyte reruns visited irony deodorant frying fever neither motherly lexicon sober skirting gnome duckling goblet gown duke
// primary address: 53Rh4JwXa5n22hdnwKvZev2wDfVhDxYpcd7EimFAGe7PNu5unZ6UQvRCg7SuoE8FdJTHcFudxZVRx9vWbgYWoYdrQRkUdbV

var spendKey = "f87748fb2abd3e112d07fd2bec4d6d7bb63295fa110394f19cc4b78f12dc3902"
var viewKey = "9ca0a18760517b877b57b534c70810348f85c29e226d52b0205b684d50aa390c"
var primaryAddress = "53Rh4JwXa5n22hdnwKvZev2wDfVhDxYpcd7EimFAGe7PNu5unZ6UQvRCg7SuoE8FdJTHcFudxZVRx9vWbgYWoYdrQRkUdbV"

//////////////////
// Run Stagenet //
//////////////////

func main() {
	ctx := context.Background()

	bobRpc := rpc.New(rpc.Config{
		// Alice: monero-wallet-rpc server
		Address: "http://127.0.0.1:38184/json_rpc",
		Client:  &http.Client{ /*default no auth HTTP client*/ },
	})

	// Generate from primary address, spend key, view key above
	gfk_req := &rpc.GenerateFromKeysRequest{
		RestoreHeight: 1600000, // birthday if known
		Filename:      "bob",
		Address:       primaryAddress,
		SpendKey:      spendKey,
		ViewKey:       viewKey,
		Password:      "",
	}

	gfk_resp, err := bobRpc.GenerateFromKeys(ctx, gfk_req)
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

	rf_resp, err := bobRpc.Refresh(ctx, rf_req)
	if err != nil {
		fmt.Printf("refresh: %v\n", err)
		bobRpc.CloseWallet(ctx)
		os.Exit(2)
	}
	fmt.Println("REFRESHED!")
	fmt.Printf("blocks fetched from monerod: %d\n", rf_resp.BlocksFetched)
	fmt.Printf("did we get money? %v\n", rf_resp.ReceivedMoney)

	gb_req := &rpc.GetBalanceRequest{
		AccountIndex: 0,
	}

	gb_resp, err := bobRpc.GetBalance(ctx, gb_req)
	if err != nil {
		fmt.Printf("getbalance: %v\n", err)
		bobRpc.CloseWallet(ctx)
		os.Exit(2)
	}
	fmt.Printf("balance: %d\n", gb_resp.Balance)
	fmt.Printf("unlocked balance: %d\n", gb_resp.UnlockedBalance)

	//////////////////////
	// Use the wallet here
	//////////////////////
	fmt.Println("DO STUFF!")

	bobRpc.CloseWallet(ctx)

	fmt.Println("CLOSED!")
}