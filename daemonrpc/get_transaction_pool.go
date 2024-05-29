package daemonrpc

import "context"

type SpentKeyImage struct {
	IdHash   string   `json:"id_hash"`
	TxHashes []TxHash `json:"tx_hashes"`
}

type Transaction struct {
	BlobSize           uint64 `json:"blob_size"`
	DoNotRelay         bool   `json:"do_not_relay"`
	DoubleSpendSeen    bool   `json:"double_spend_seen"`
	Fee                uint64 `json:"fee"`
	IdHash             string `json:"id_hash"`
	KeptByBlock        bool   `json:"kept_by_block"`
	LastFailedHeight   uint64 `json:"last_failed_height"`
	LastFailedIdHash   string `json:"last_failed_id_hash"`
	LastRelayedTime    uint64 `json:"last_relayed_time"`
	MaxUsedBlockHeight uint64 `json:"max_used_block_height"`
	MaxUsedBlockIdHash string `json:"max_used_block_id_hash"`
	ReceiveTime        uint64 `json:"receive_time"`
	Relayed            bool   `json:"relayed"`
	TxBlob             string `json:"tx_blob"`
	TxJson             string `json:"tx_json"`
}

type GetTransactionPoolResponse struct {
	Status         string          `json:"status"`
	Credits        uint64          `json:"credits"`
	SpentKeyImages []SpentKeyImage `json:"spent_key_images"`
	Transactions   []Transaction   `json:"transactions"`
	Untrusted      bool            `json:"untrusted"`
}

// Connect to the Monero daemon.
func (c *Client) GetTransactionPool(ctx context.Context) (*GetTransactionPoolResponse, error) {
	resp := &GetTransactionPoolResponse{}
	err := c.Do(ctx, "get_transaction_pool", nil, resp)
	return resp, err
}
