package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/starcoinorg/starcoin-go/client/jsonrpc/codec"
	"io/ioutil"
	"net/http"
)

// HTTP is an http transport
type HTTP struct {
	addr   string
	client *http.Client
}

func newHTTP(addr string) *HTTP {
	return &HTTP{
		addr:   addr,
		client: http.DefaultClient,
	}
}

// Close implements the transport interface
func (h *HTTP) Close() error {
	return nil
}

// Call implements the transport interface
func (h *HTTP) Call(context context.Context,method string, out interface{}, params interface{}) error {
	// Encode json-rpc request
	request := codec.Request{
		Method:  method,
		Version: "2.0",
	}

	if params != nil {
		data, err := json.Marshal(params)
		if err != nil {
			return err
		}
		request.Params = data
	}
	raw, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req,err := http.NewRequest("POST",h.addr,bytes.NewReader(raw))
	req.WithContext(context)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := h.client.Do(req)
	if err != nil  {
		return errors.WithStack(err)
	}

	// Decode json-rpc response
	var response codec.Response
	defer res.Body.Close()
	body,err := ioutil.ReadAll(res.Body)
	if err != nil  {
		return errors.WithStack(err)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return errors.Wrap(err, "parse response failed")
	}
	if response.Error != nil {
		return response.Error
	}

	if err := json.Unmarshal(response.Result, out); err != nil {
		return errors.Wrap(err, "parse result failed")
	}
	return nil
}
