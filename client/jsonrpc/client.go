package jsonrpc

import (
	"context"
	"github.com/starcoinorg/starcoin-go/client/jsonrpc/transport"
)

// Client is the jsonrpc client
type Client struct {
	transport transport.Transport
}

// NewClient creates a new client
func NewClient(addr string) (*Client, error) {
	c := &Client{}

	t, err := transport.NewTransport(addr)
	if err != nil {
		return nil, err
	}
	c.transport = t
	return c, nil
}

// Close closes the tranport
func (c *Client) Close() error {
	return c.transport.Close()
}

// Call makes a jsonrpc call
func (c *Client) Call(context context.Context,method string, out interface{}, params interface{}) error {
	return c.transport.Call(context,method, out, params)
}
