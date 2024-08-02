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
	"log"
	"net/http"
	"os"

	"github.com/dev-warrior777/go-monero/rpc"
)

// "Own"
// =====
// var spendKey = "e42a6f98f92faa98b695e7f05725bc736834280800b154fb145c778c74fbfe01"
// var viewKey = "4ada170972a137d2c2efd1c819d13e5465415677ee28350addc507c28459c900"
// var primaryAddress = "47rDCGk2ndWBySuLHcfFVbEQARPAiVHcq3nbV2FuNggrMTMqDq1QdX4enkxAG1sPRvT1FyXWKtA2F8ZhweZgsCDhNyPmLb3"

/////////////////////////
// Run Regtest Harness //
/////////////////////////

func main() {
	ctx := context.Background()

	bill := rpc.New(rpc.Config{
		// bill rpc address
		Address: "http://127.0.0.1:28184/json_rpc",
		Client:  &http.Client{ /*default no auth HTTP client*/ },
	})

	// bill sends 0.000333 XMR -> own primary address
	tr_req := rpc.TransferRequest{
		Destinations: []rpc.Destination{
			{
				Amount: 333000000,
				// "own" address. No need to have own wallet loaded (opened) for the send
				Address: "47rDCGk2ndWBySuLHcfFVbEQARPAiVHcq3nbV2FuNggrMTMqDq1QdX4enkxAG1sPRvT1FyXWKtA2F8ZhweZgsCDhNyPmLb3",
			},
		},
		AccountIndex:   0,
		SubaddrIndices: []uint64{0},
		Priority:       0,
		RingSize:       11,
		UnlockTime:     0, // locked tx .. maybe for atomic swap can put n blocks here
		GetTxKey:       true,
	}
	tr_resp, err := bill.Transfer(ctx, &tr_req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx_hash:", tr_resp.TxHash)

	// ----------------------------------------------------------
	//
	// ----------------------------------------------------------

	own := rpc.New(rpc.Config{
		// own rpc port
		Address: "http://127.0.0.1:28884/json_rpc",
		Client:  &http.Client{ /*default no auth HTTP client*/ },
	})

	ow_req := &rpc.OpenWalletRequest{
		Filename: "own",
		Password: "",
	}

	own.OpenWallet(ctx, ow_req)

	rf_req := &rpc.RefreshRequest{
		StartHeight: 0, // (Optional) The block height from which to start. Real wallet put this in
	}
	rf_resp, err := own.Refresh(ctx, rf_req)
	if err != nil {
		fmt.Printf("refresh: %v\n", err)
		own.CloseWallet(ctx)
		os.Exit(1)
	}
	fmt.Printf("blocks fetched from monerod: %d\n", rf_resp.BlocksFetched)
	fmt.Printf("did we get more money? %v\n", rf_resp.ReceivedMoney)

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

	fmt.Println("Now mine_to_bill on the harness at least 10 blocks to make the money spendable")

	own.CloseWallet(ctx)
}
