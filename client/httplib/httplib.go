package httplib

import (
	"context"
	"crypto/tls"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"time"
)

const contentTypeKey = "Content-Type"

var doRequestFilter = func(ctx context.Context, req *BeegoHTTPRequest) (*http.Response, error) {
	return req.doRequest(ctx)
}

func NewBeegoRequest(rawurl, method string) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func NewBeegoRequestWithCtx(ctx context.Context, rawurl, method string) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func Get(url string) *BeegoHTTPRequest { _ = "STUB: not implemented"; return nil }

func Post(url string) *BeegoHTTPRequest { _ = "STUB: not implemented"; return nil }

func Put(url string) *BeegoHTTPRequest { _ = "STUB: not implemented"; return nil }

func Delete(url string) *BeegoHTTPRequest { _ = "STUB: not implemented"; return nil }

func Head(url string) *BeegoHTTPRequest { _ = "STUB: not implemented"; return nil }

type BeegoHTTPRequest struct {
	url     string
	req     *http.Request
	params  map[string][]string
	files   map[string]string
	setting BeegoHTTPSettings
	resp    *http.Response

	body []byte

	copyBody func() io.ReadCloser
}

func (b *BeegoHTTPRequest) GetRequest() *http.Request { _ = "STUB: not implemented"; return nil }

func (b *BeegoHTTPRequest) Setting(setting BeegoHTTPSettings) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) SetBasicAuth(username, password string) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) SetEnableCookie(enable bool) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) SetUserAgent(useragent string) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) Retries(times int) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) RetryDelay(delay time.Duration) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) SetTimeout(connectTimeout, readWriteTimeout time.Duration) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) SetTLSClientConfig(config *tls.Config) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) Header(key, value string) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) SetHost(host string) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) SetProtocolVersion(vers string) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) SetCookie(cookie *http.Cookie) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) SetTransport(transport http.RoundTripper) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) SetProxy(proxy func(*http.Request) (*url.URL, error)) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) SetCheckRedirect(redirect func(req *http.Request, via []*http.Request) error) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) SetFilters(fcs ...FilterChain) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) AddFilters(fcs ...FilterChain) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) SetEscapeHTML(isEscape bool) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) Param(key, value string) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) PostFile(formname, filename string) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) Body(data interface{}) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) reqBody(data []byte) *BeegoHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (b *BeegoHTTPRequest) XMLBody(obj interface{}) (*BeegoHTTPRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BeegoHTTPRequest) YAMLBody(obj interface{}) (*BeegoHTTPRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BeegoHTTPRequest) JSONBody(obj interface{}) (*BeegoHTTPRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BeegoHTTPRequest) JSONMarshal(obj interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BeegoHTTPRequest) buildURL(paramBody string) { _ = "STUB: not implemented"; return }

func (b *BeegoHTTPRequest) handleFiles() { _ = "STUB: not implemented"; return }

func (*BeegoHTTPRequest) handleFileToBody(bodyWriter *multipart.Writer, formname string, filename string) {
	_ = "STUB: not implemented"
	return
}

func (b *BeegoHTTPRequest) getResponse() (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BeegoHTTPRequest) DoRequest() (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BeegoHTTPRequest) DoRequestWithCtx(ctx context.Context) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BeegoHTTPRequest) doRequest(_ context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BeegoHTTPRequest) sendRequest(client *http.Client) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BeegoHTTPRequest) buildCookieJar() http.CookieJar {
	_ = "STUB: not implemented"
	return *new(http.CookieJar)
}

func (b *BeegoHTTPRequest) buildTrans() http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

func (b *BeegoHTTPRequest) buildParamBody() string { _ = "STUB: not implemented"; return "" }

func (b *BeegoHTTPRequest) String() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (b *BeegoHTTPRequest) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *BeegoHTTPRequest) ToFile(filename string) error { _ = "STUB: not implemented"; return nil }

func pathExistAndMkdir(filename string) (err error) { _ = "STUB: not implemented"; return nil }

func (b *BeegoHTTPRequest) ToJSON(v interface{}) error { _ = "STUB: not implemented"; return nil }

func (b *BeegoHTTPRequest) ToXML(v interface{}) error { _ = "STUB: not implemented"; return nil }

func (b *BeegoHTTPRequest) ToYAML(v interface{}) error { _ = "STUB: not implemented"; return nil }

func (b *BeegoHTTPRequest) ToValue(value interface{}) error { _ = "STUB: not implemented"; return nil }

func (b *BeegoHTTPRequest) Response() (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	_ = "STUB: not implemented"
	return nil
}

func TimeoutDialerCtx(cTimeout time.Duration,
	rwTimeout time.Duration) func(ctx context.Context, net, addr string) (c net.Conn, err error) {
	_ = "STUB: not implemented"
	return nil
}
