package old_rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

type OutPubKey string

type EcdhInfo struct {
	// Diffie-Helman elliptic curves' structures
	Amount string `json:"amount"`
	Mask   string `json:"mask"`
}

type RctSignatures struct {
	// List of signatures used in the ring signatures truct - to hide the true origin
	// of the transaction inputs.
	Type   uint64      `json:"type"`
	TxnFee uint64      `json:"txnFee"`
	Ecdh   []EcdhInfo  `json:"ecdhInfo"`
	OutPk  []OutPubKey `json:"outPk"`
}

type TaggedKey struct {
	StealthKey string `json:"key"`
	ViewTag    string `json:"view_tag"`
}

type Target struct {
	TaggedKey `json:"tagged_key"`
}
type Output struct {
	AmountCoinbase uint64 `json:"amount"`
	Target         Target `json:"target"`
}

type PubKeyPrevOut struct {
	Amount     uint64   `json:"amount"`
	KeyOffsets []uint64 `json:"key_offsets"`
	KeyImage   string   `json:"k_image"`
}

type Input struct {
	Key PubKeyPrevOut `json:"key"`
}

type JsonDetails struct {
	Version       uint64        `json:"version"`
	TxUnlockTime  uint64        `json:"unlock_time"` // unlock time for txs we lock for n blocks
	Vin           []Input       `json:"vin"`
	Vout          []Output      `json:"vout"`
	Extra         []uint8       `json:"extra"` // bytes: Payment ID, miner info, random bytes, etc.
	RctSignatures RctSignatures `json:"rct_signatures"`
}

type TxEntry struct {
	TxHash          string   `json:"tx_hash"`
	BlockHeight     uint64   `json:"block_height"`
	BlockTimestamp  uint64   `json:"block_timestamp"`
	DoublespendSeen bool     `json:"double_spend_seen"`
	InMempool       bool     `json:"in_pool"`
	OutputIndeces   []uint64 `json:"output_indices"`
	AsHex           string   `json:"as_hex"`
	AsJson          string   `json:"as_json"`
	// Additional tx details if DecodeAsJson is requested:
	// - sent as a raw json string in the 'as_json' field which we need to further
	//   decode into JsonDetails
	// - Discards all the 'rctsig_prunable' info as probably not useful for us.
	JsonDetails JsonDetails
}

type GetTransactionsResponse struct {
	Status    string    `json:"status"`
	Credits   uint64    `json:"credits"`
	MissedTxs []string  `json:"missed_tx"`
	TopHash   string    `json:"top_hash"`
	Untrusted bool      `json:"untrusted"`
	Txs       []TxEntry `json:"txs"`
	TxsAsHex  []string  `json:"txs_as_hex"`
}

type GetTransactionsRequest struct {
	TxsHashes    []TxHash `json:"txs_hashes"`
	DecodeAsJson bool     `json:"decode_as_json,omitempty"`
}

func (c *Client) GetTransactions(ctx context.Context, req *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	resp := &GetTransactionsResponse{}
	err := c.Do(ctx, "get_transactions", req, resp)
	if err != nil {
		return nil, err
	}
	if req.DecodeAsJson {
		err = decodeAsJson(resp)
		if err != nil {
			return nil, err
		}
	}
	return resp, err
}

// Parse 'as_json' (which has been returned as raw json) into TxEntry.Jsondetails
func decodeAsJson(resp *GetTransactionsResponse) error {
	if resp.Txs == nil {
		return errors.New("json decode error: failed to get Txs array")
	}
	for i, tx := range resp.Txs {
		asJson := tx.AsJson
		if asJson == "" {
			return fmt.Errorf("json decode error: no AsJson info for tx #%d", i)
		}
		r := bytes.NewBufferString(asJson)
		dec := json.NewDecoder(r)
		err := dec.Decode(&tx.JsonDetails)
		if err != nil {
			return err
		}
		resp.Txs[i].JsonDetails = tx.JsonDetails
	}
	return nil
}
