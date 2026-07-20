package ledis

import (
	"context"
	"net/http"
	"sync"

	"github.com/ledisdb/ledisdb/ledis"

	"github.com/beego/beego/v2/server/web/session"
)

var (
	ledispder = &Provider{}
	c         *ledis.DB
)

type SessionStore struct {
	sid         string
	lock        sync.RWMutex
	values      map[interface{}]interface{}
	maxlifetime int64
}

func (ls *SessionStore) Set(ctx context.Context, key, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (ls *SessionStore) Get(ctx context.Context, key interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (ls *SessionStore) Delete(ctx context.Context, key interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (ls *SessionStore) Flush(context.Context) error { _ = "STUB: not implemented"; return nil }

func (ls *SessionStore) SessionID(context.Context) string { _ = "STUB: not implemented"; return "" }

func (ls *SessionStore) SessionRelease(_ context.Context, _ http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

func (ls *SessionStore) SessionReleaseIfPresent(c context.Context, w http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

type Provider struct {
	maxlifetime int64
	SavePath    string `json:"save_path"`
	Db          int    `json:"db"`
}

func (lp *Provider) SessionInit(ctx context.Context, maxlifetime int64, cfgStr string) error {
	_ = "STUB: not implemented"
	return nil
}

func (lp *Provider) initOldStyle(cfgStr string) error { _ = "STUB: not implemented"; return nil }

func (lp *Provider) SessionRead(ctx context.Context, sid string) (session.Store, error) {
	_ = "STUB: not implemented"
	return *new(session.Store), nil
}

func (lp *Provider) SessionExist(ctx context.Context, sid string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (lp *Provider) SessionRegenerate(ctx context.Context, oldsid, sid string) (session.Store, error) {
	_ = "STUB: not implemented"
	return *new(session.Store), nil
}

func (lp *Provider) SessionDestroy(ctx context.Context, sid string) error {
	_ = "STUB: not implemented"
	return nil
}

func (lp *Provider) SessionGC(context.Context) { _ = "STUB: not implemented"; return }

func (lp *Provider) SessionAll(context.Context) int { _ = "STUB: not implemented"; return 0 }

func init() {
	session.Register("ledis", ledispder)
}
