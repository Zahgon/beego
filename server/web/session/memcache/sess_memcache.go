package memcache

import (
	"context"
	"net/http"
	"sync"

	"github.com/bradfitz/gomemcache/memcache"

	"github.com/beego/beego/v2/server/web/session"
)

var (
	mempder = &MemProvider{}
	client  *memcache.Client
)

type SessionStore struct {
	sid         string
	lock        sync.RWMutex
	values      map[interface{}]interface{}
	maxlifetime int64
}

func (rs *SessionStore) Set(ctx context.Context, key, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *SessionStore) Get(ctx context.Context, key interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (rs *SessionStore) Delete(ctx context.Context, key interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (rs *SessionStore) Flush(context.Context) error { _ = "STUB: not implemented"; return nil }

func (rs *SessionStore) SessionID(context.Context) string { _ = "STUB: not implemented"; return "" }

func (rs *SessionStore) SessionRelease(ctx context.Context, w http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

func (rs *SessionStore) SessionReleaseIfPresent(ctx context.Context, w http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

func (rs *SessionStore) releaseSession(_ context.Context, _ http.ResponseWriter, requirePresent bool) {
	_ = "STUB: not implemented"
	return
}

type MemProvider struct {
	maxlifetime int64
	conninfo    []string
	poolsize    int
	password    string
}

func (rp *MemProvider) SessionInit(ctx context.Context, maxlifetime int64, savePath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (rp *MemProvider) SessionRead(ctx context.Context, sid string) (session.Store, error) {
	_ = "STUB: not implemented"
	return *new(session.Store), nil
}

func (rp *MemProvider) SessionExist(ctx context.Context, sid string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (rp *MemProvider) SessionRegenerate(ctx context.Context, oldsid, sid string) (session.Store, error) {
	_ = "STUB: not implemented"
	return *new(session.Store), nil
}

func (rp *MemProvider) SessionDestroy(ctx context.Context, sid string) error {
	_ = "STUB: not implemented"
	return nil
}

func (rp *MemProvider) connectInit() error { _ = "STUB: not implemented"; return nil }

func (rp *MemProvider) SessionGC(context.Context) { _ = "STUB: not implemented"; return }

func (rp *MemProvider) SessionAll(context.Context) int { _ = "STUB: not implemented"; return 0 }

func init() {
	session.Register("memcache", mempder)
}
