package httplib

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)

type (
	ClientOption           func(client *Client)
	BeegoHTTPRequestOption func(request *BeegoHTTPRequest)
)

func WithEnableCookie(enable bool) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithUserAgent(userAgent string) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithTLSClientConfig(config *tls.Config) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithTransport(transport http.RoundTripper) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithProxy(proxy func(*http.Request) (*url.URL, error)) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithCheckRedirect(redirect func(req *http.Request, via []*http.Request) error) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithHTTPSetting(setting BeegoHTTPSettings) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithEnableGzip(enable bool) ClientOption { _ = "STUB: not implemented"; return *new(ClientOption) }

func WithTimeout(connectTimeout, readWriteTimeout time.Duration) BeegoHTTPRequestOption {
	_ = "STUB: not implemented"
	return *new(BeegoHTTPRequestOption)
}

func WithHeader(key, value string) BeegoHTTPRequestOption {
	_ = "STUB: not implemented"
	return *new(BeegoHTTPRequestOption)
}

func WithCookie(cookie *http.Cookie) BeegoHTTPRequestOption {
	_ = "STUB: not implemented"
	return *new(BeegoHTTPRequestOption)
}

func WithTokenFactory(tokenFactory func() string) BeegoHTTPRequestOption {
	_ = "STUB: not implemented"
	return *new(BeegoHTTPRequestOption)
}

func WithBasicAuth(basicAuth func() (string, string)) BeegoHTTPRequestOption {
	_ = "STUB: not implemented"
	return *new(BeegoHTTPRequestOption)
}

func WithFilters(fcs ...FilterChain) BeegoHTTPRequestOption {
	_ = "STUB: not implemented"
	return *new(BeegoHTTPRequestOption)
}

func WithContentType(contentType string) BeegoHTTPRequestOption {
	_ = "STUB: not implemented"
	return *new(BeegoHTTPRequestOption)
}

func WithParam(key, value string) BeegoHTTPRequestOption {
	_ = "STUB: not implemented"
	return *new(BeegoHTTPRequestOption)
}

func WithRetry(times int, delay time.Duration) BeegoHTTPRequestOption {
	_ = "STUB: not implemented"
	return *new(BeegoHTTPRequestOption)
}
