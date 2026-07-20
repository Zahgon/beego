package httplib

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type BeegoHTTPSettings struct {
	UserAgent        string
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
	TLSClientConfig  *tls.Config
	Proxy            func(*http.Request) (*url.URL, error)
	Transport        http.RoundTripper
	CheckRedirect    func(req *http.Request, via []*http.Request) error
	EnableCookie     bool
	Gzip             bool
	Retries          int
	RetryDelay       time.Duration
	FilterChains     []FilterChain
	EscapeHTML       bool
}

func createDefaultCookie() { _ = "STUB: not implemented"; return }

func SetDefaultSetting(setting BeegoHTTPSettings) { _ = "STUB: not implemented"; return }

func GetDefaultSetting() BeegoHTTPSettings {
	_ = "STUB: not implemented"
	return *new(BeegoHTTPSettings)
}

var defaultSetting = BeegoHTTPSettings{
	UserAgent:        "beegoServer",
	ConnectTimeout:   60 * time.Second,
	ReadWriteTimeout: 60 * time.Second,
	Gzip:             true,
	FilterChains:     make([]FilterChain, 0, 4),
	EscapeHTML:       true,
}

var (
	defaultCookieJar http.CookieJar
	settingMutex     sync.Mutex
)

func AddDefaultFilter(fc FilterChain) { _ = "STUB: not implemented"; return }
