package rpc

import "context"

type CreateAddressRequest struct {
	// Create a new address for this account.
	AccountIndex uint64 `json:"account_index"`

	// (Optional) Label for the new address.
	Label string `json:"label,omitempty"`

	// (Optional) Number of addresses to create (Defaults to 1).
	Count uint64 `json:"count"`
}

type CreateAddressResponse struct {
	// Newly created address. Base58 representation of the public keys.
	Address string `json:"address"`

	// Index of the new address under the input account.
	AddressIndex uint64 `json:"address_index"`

	// List of (count) address indeces
	AddressIndeces []uint64 `json:"address_indices"`

	// List of (count) addresses
	Addresses []string `json:"addresses"`
}

// Create a new address for an account. Optionally, label the new address.
func (c *Client) CreateAddress(ctx context.Context, req *CreateAddressRequest) (*CreateAddressResponse, error) {
	resp := &CreateAddressResponse{}
	err := c.Do(ctx, "create_address", &req, resp)
	return resp, err
}
