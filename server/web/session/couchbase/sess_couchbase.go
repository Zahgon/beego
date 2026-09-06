package couchbase

import (
	"context"
	"net/http"
	"sync"

	couchbase "github.com/couchbase/go-couchbase"

	"github.com/beego/beego/v2/server/web/session"
)

var couchbpder = &Provider{}

type SessionStore struct {
	b           *couchbase.Bucket
	sid         string
	lock        sync.RWMutex
	values      map[interface{}]interface{}
	maxlifetime int64
}

type Provider struct {
	maxlifetime int64
	SavePath    string `json:"save_path"`
	Pool        string `json:"pool"`
	Bucket      string `json:"bucket"`
	b           *couchbase.Bucket
}

func (cs *SessionStore) Set(ctx context.Context, key, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *SessionStore) Get(ctx context.Context, key interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (cs *SessionStore) Delete(ctx context.Context, key interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *SessionStore) Flush(context.Context) error { _ = "STUB: not implemented"; return nil }

func (cs *SessionStore) SessionID(context.Context) string { _ = "STUB: not implemented"; return "" }

func (cs *SessionStore) SessionRelease(_ context.Context, _ http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

func (cs *SessionStore) SessionReleaseIfPresent(c context.Context, w http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

func (cp *Provider) getBucket() *couchbase.Bucket { _ = "STUB: not implemented"; return nil }

func (cp *Provider) SessionInit(ctx context.Context, maxlifetime int64, cfg string) error {
	_ = "STUB: not implemented"
	return nil
}

func (cp *Provider) initOldStyle(savePath string) error { _ = "STUB: not implemented"; return nil }

func (cp *Provider) SessionRead(ctx context.Context, sid string) (session.Store, error) {
	_ = "STUB: not implemented"
	return *new(session.Store), nil
}

func (cp *Provider) SessionExist(ctx context.Context, sid string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cp *Provider) SessionRegenerate(ctx context.Context, oldsid, sid string) (session.Store, error) {
	_ = "STUB: not implemented"
	return *new(session.Store), nil
}

func (cp *Provider) SessionDestroy(ctx context.Context, sid string) error {
	_ = "STUB: not implemented"
	return nil
}

func (cp *Provider) SessionGC(context.Context) { _ = "STUB: not implemented"; return }

func (cp *Provider) SessionAll(context.Context) int { _ = "STUB: not implemented"; return 0 }

func init() {
	session.Register("couchbase", couchbpder)
}
