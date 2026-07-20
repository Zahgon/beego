package session

import (
	"context"
	"crypto/cipher"
	"net/http"
	"sync"
)

var cookiepder = &CookieProvider{}

type CookieSessionStore struct {
	sid    string
	values map[interface{}]interface{}
	lock   sync.RWMutex
}

func (st *CookieSessionStore) Set(ctx context.Context, key, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (st *CookieSessionStore) Get(ctx context.Context, key interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (st *CookieSessionStore) Delete(ctx context.Context, key interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (st *CookieSessionStore) Flush(context.Context) error { _ = "STUB: not implemented"; return nil }

func (st *CookieSessionStore) SessionID(context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

func (st *CookieSessionStore) SessionRelease(_ context.Context, w http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

func (st *CookieSessionStore) SessionReleaseIfPresent(ctx context.Context, w http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

type cookieConfig struct {
	SecurityKey  string `json:"securityKey"`
	BlockKey     string `json:"blockKey"`
	SecurityName string `json:"securityName"`
	CookieName   string `json:"cookieName"`
	Secure       bool   `json:"secure"`
	Maxage       int    `json:"maxage"`
}

type CookieProvider struct {
	maxlifetime int64
	config      *cookieConfig
	block       cipher.Block
}

func (pder *CookieProvider) SessionInit(ctx context.Context, maxlifetime int64, config string) error {
	_ = "STUB: not implemented"
	return nil
}

func (pder *CookieProvider) SessionRead(ctx context.Context, sid string) (Store, error) {
	_ = "STUB: not implemented"
	return *new(Store), nil
}

func (pder *CookieProvider) SessionExist(ctx context.Context, sid string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (pder *CookieProvider) SessionRegenerate(ctx context.Context, oldsid, sid string) (Store, error) {
	_ = "STUB: not implemented"
	return *new(Store), nil
}

func (pder *CookieProvider) SessionDestroy(ctx context.Context, sid string) error {
	_ = "STUB: not implemented"
	return nil
}

func (pder *CookieProvider) SessionGC(context.Context) { _ = "STUB: not implemented"; return }

func (pder *CookieProvider) SessionAll(context.Context) int { _ = "STUB: not implemented"; return 0 }

func (pder *CookieProvider) SessionUpdate(ctx context.Context, sid string) error {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	Register("cookie", cookiepder)
}
