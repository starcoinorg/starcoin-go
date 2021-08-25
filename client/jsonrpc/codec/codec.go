package codec

import (
	"encoding/json"
	"fmt"
)

// Request is a jsonrpc request
type Request struct {
	ID      uint64          `json:"id"`
	Version string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

// Response is a jsonrpc response
type Response struct {
	ID     uint64          `json:"id"`
	Result json.RawMessage `json:"result"`
	Error  *ErrorObject    `json:"error,omitempty"`
}

// Response is a jsonrpc response
type WsResponse struct {
	ID     uint64       `json:"id"`
	Result int          `json:"result"`
	Error  *ErrorObject `json:"error,omitempty"`
}

// ErrorObject is a jsonrpc error
type ErrorObject struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Subscription is a jsonrpc subscription
type Subscription struct {
	ID     int             `json:"subscription"`
	Result json.RawMessage `json:"result"`
}

// Error implements error interface
func (e *ErrorObject) Error() string {
	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("jsonrpc.internal marshal error: %v", err)
	}
	return string(data)
}
