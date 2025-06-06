package rpc

type Address struct {
	// AddressIndex Index of the subaddress in the account.
	AddressIndex uint64 `json:"address_index"`

	// Address at this index. Base58 representation of the public keys.
	Address string `json:"address"`

	// Balance for the subaddress (locked or unlocked).
	Balance uint64 `json:"balance"`

	// UnlockedBalance Unlocked balance for the subaddress.
	UnlockedBalance uint64 `json:"unlocked_balance"`

	// Label of the subaddress.
	Label string `json:"label"`

	// Tag for filtering accounts.
	Tag string `json:"tag"`

	// NumUnspentOutputs Number of unspent outputs available for the subaddress.
	NumUnspentOutputs uint64 `json:"num_unspent_outputs"`

	// Used states if the (sub)address has already received funds.
	Used bool `json:"used"`
}

type SubaddressIndex struct {
	// Major  Account index.
	Major uint64 `json:"major"`

	// Minor  Address index.
	Minor uint64 `json:"minor"`
}

type Destination struct {
	// Amount to send to each destination, in atomic units.
	Amount uint64 `json:"amount"`

	// Destination public address.
	Address string `json:"address"`
}

type Payment struct {
	// Payment ID matching the input parameter.
	PaymentId string `json:"payment_id"`

	// Transaction hash used as the transaction ID.
	TxHash string `json:"tx_hash"`

	// Amount for this payment.
	Amount uint64 `json:"amount"`

	// Height of the block that first confirmed this payment.
	BlockHeight uint64 `json:"block_height"`

	// Time (in block height) until this payment is safe to spend.
	UnlockTime uint64 `json:"unlock_time"`

	// SubaddrIndex index:
	SubaddrIndex SubaddressIndex `json:"subaddr_index"`

	// Account index for the subaddress.
	Major uint64 `json:"major"`

	// Index of the subaddress in the account.
	Minor uint64 `json:"minor"`

	// Address receiving the payment; Base58 representation of the public keys.
	Address string `json:"address"`
}

type SignedKeyImage struct {
	KeyImage  string `json:"key_image"`
	Signature string `json:"signature"`
}

type Entry struct {
	// Public address of the entry.
	Address string `json:"address"`

	// Description of this address entry.
	Description string `json:"description"`

	Index uint64 `json:"index"`
}

type KeyImages struct {
	KeyImages []string `json:"key_images"`
}

type BlockHeader struct {
	BlockSize                 uint64 `json:"block_size"`
	BlockWeight               uint64 `json:"block_weight"`
	CumulativeDifficulty      uint64 `json:"cumulative_difficulty"`
	CumulativeDifficultyTop64 uint64 `json:"cumulative_difficulty_top64"`
	Depth                     uint64 `json:"depth"`
	Difficulty                uint64 `json:"difficulty"`
	DifficultyTop64           uint64 `json:"difficulty_top64"`
	Hash                      string `json:"hash"`
	Height                    uint64 `json:"height"`
	LongTermWeight            uint64 `json:"long_term_weight"`
	MajorVersion              uint64 `json:"major_version"`
	MinerTxHash               string `json:"miner_tx_hash"`
	MinorVersion              uint64 `json:"minor_version"`
	Nonce                     uint64 `json:"nonce"`
	NumTxes                   uint64 `json:"num_txes"`
	OrphanStatus              bool   `json:"orphan_status"`
	PowHash                   string `json:"pow_hash"`
	PrevHash                  string `json:"prev_hash"`
	Reward                    uint64 `json:"reward"`
	Timestamp                 uint64 `json:"timestamp"`
	WideCumulativeDifficulty  string `json:"wide_cumulative_difficulty"`
	WideDifficulty            string `json:"wide_difficulty"`
}
