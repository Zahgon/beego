package httplib

import (
	"io"
	"net/http"
)

type Client struct {
	Name       string
	Endpoint   string
	CommonOpts []BeegoHTTPRequestOption

	Setting BeegoHTTPSettings
}

type HTTPResponseCarrier interface {
	SetHTTPResponse(resp *http.Response)
}

type HTTPBodyCarrier interface {
	SetReader(r io.ReadCloser)
}

type HTTPBytesCarrier interface {
	SetBytes(bytes []byte)
}

type HTTPStatusCarrier interface {
	SetStatusCode(status int)
}

type HTTPHeadersCarrier interface {
	SetHeader(header map[string][]string)
}

func NewClient(name string, endpoint string, opts ...ClientOption) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) customReq(req *BeegoHTTPRequest, opts []BeegoHTTPRequestOption) {
	_ = "STUB: not implemented"
	return
}

func (c *Client) handleResponse(value interface{}, req *BeegoHTTPRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) handleCarrier(value interface{}, req *BeegoHTTPRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Get(value interface{}, path string, opts ...BeegoHTTPRequestOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Post(value interface{}, path string, body interface{}, opts ...BeegoHTTPRequestOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Put(value interface{}, path string, body interface{}, opts ...BeegoHTTPRequestOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Delete(value interface{}, path string, opts ...BeegoHTTPRequestOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Head(value interface{}, path string, opts ...BeegoHTTPRequestOption) error {
	_ = "STUB: not implemented"
	return nil
}
