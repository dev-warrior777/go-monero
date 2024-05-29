package daemonrpc

import "context"

type TxEntry struct {
	AsHex           string   `json:"as_hex"`
	AsJson          string   `json:"as_json"`
	BlockHeight     uint64   `json:"block_height"`
	BlockTimestamp  uint64   `json:"block_timestamp"`
	DoublespendSeen bool     `json:"double_spend_seen"`
	InMempool       bool     `json:"in_pool"`
	OutputIndeces   []uint64 `json:"output_indices"`
	TxHash          string   `json:"tx_hash"`
}

type GetTransactionsRequest struct {
	TxsHashes    []TxHash `json:"txs_hashes"`
	DecodeAsJson bool     `json:"decode_as_json"`
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

func (c *Client) GetTransactions(ctx context.Context, req *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	resp := &GetTransactionsResponse{}
	err := c.Do(ctx, "get_transactions", req, resp)
	if req.DecodeAsJson {
		// TODO: parse 'as_json' raw json returned if
	}
	return resp, err
}
