package main

///////////////////////////////////////////////////////////////////////////////
// use wallet created previously with generate from keys, etc.
///////////////////////////////////////////////////////////////////////////////

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dev-warrior777/go-monero/rpc"
)

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

	ow_req := &rpc.OpenWalletRequest{
		Filename: "alice",
		Password: "",
	}

	fmt.Println("Opening wallet...")

	err := aliceRpc.OpenWallet(ctx, ow_req)
	if err != nil {
		fmt.Printf("open: %v\n", err)
		aliceRpc.CloseWallet(ctx)
		os.Exit(1)
	}
	fmt.Println("OPENED!")

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
		fmt.Printf("refresh: %v\n", err)
		aliceRpc.CloseWallet(ctx)
		os.Exit(2)
	}
	fmt.Printf("balance: %d\n", gb_resp.Balance)
	fmt.Printf("unlocked balance: %d\n", gb_resp.UnlockedBalance)

	bobDest := rpc.Destination{
		Amount:  500000000000,
		Address: "53Rh4JwXa5n22hdnwKvZev2wDfVhDxYpcd7EimFAGe7PNu5unZ6UQvRCg7SuoE8FdJTHcFudxZVRx9vWbgYWoYdrQRkUdbV",
	}
	tr_req := &rpc.TransferRequest{
		Destinations: []rpc.Destination{bobDest},
		Priority:     0,
		RingSize:     16,
		UnlockTime:   0,
	}

	tr_resp, err := aliceRpc.Transfer(ctx, tr_req)
	if err != nil {
		fmt.Printf("transfer: %v\n", err)
		aliceRpc.CloseWallet(ctx)
		os.Exit(2)
	}
	fmt.Printf("amount sent: %d\n", tr_resp.Amount)
	fmt.Printf("fee: %d\n", tr_resp.Fee)

	aliceRpc.CloseWallet(ctx)

	fmt.Println("CLOSED!")
}
