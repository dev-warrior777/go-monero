package walletrpc

import "context"

type VerifyRequest struct {
	// What should have been signed.
	Data string `json:"data"`

	// Public address of the wallet used to sign the data.
	Address string `json:"address"`

	// Signature generated by sign method.
	Signature string `json:"signature"`
}

type VerifyResponse struct {
	Good bool `json:"good"`
}

// Verify a signature on a string.
func (c *Client) Verify(ctx context.Context, req *VerifyRequest) (*VerifyResponse, error) {
	resp := &VerifyResponse{}
	err := c.Do(ctx, "verify", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}