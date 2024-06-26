package rpc

import "context"

type LabelAccountRequest struct {
	// Apply label to account at this index.
	AccountIndex uint64 `json:"account_index"`

	// Label for the account.
	Label string `json:"label"`
}

// Label an account.
func (c *Client) LabelAccount(ctx context.Context, req *LabelAccountRequest) error {
	return c.Do(ctx, "label_account", &req, nil)
}
