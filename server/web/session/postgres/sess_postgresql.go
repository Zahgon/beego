package postgres

import (
	"context"
	"database/sql"
	"net/http"
	"sync"

	_ "github.com/lib/pq"

	"github.com/beego/beego/v2/server/web/session"
)

var postgresqlpder = &Provider{}

type SessionStore struct {
	c      *sql.DB
	sid    string
	lock   sync.RWMutex
	values map[interface{}]interface{}
}

func (st *SessionStore) Set(ctx context.Context, key, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (st *SessionStore) Get(ctx context.Context, key interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (st *SessionStore) Delete(ctx context.Context, key interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (st *SessionStore) Flush(context.Context) error { _ = "STUB: not implemented"; return nil }

func (st *SessionStore) SessionID(context.Context) string { _ = "STUB: not implemented"; return "" }

func (st *SessionStore) SessionRelease(_ context.Context, _ http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

func (st *SessionStore) SessionReleaseIfPresent(ctx context.Context, w http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

type Provider struct {
	maxlifetime int64
	savePath    string
}

func (mp *Provider) connectInit() *sql.DB { _ = "STUB: not implemented"; return nil }

func (mp *Provider) SessionInit(ctx context.Context, maxlifetime int64, savePath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (mp *Provider) SessionRead(ctx context.Context, sid string) (session.Store, error) {
	_ = "STUB: not implemented"
	return *new(session.Store), nil
}

func (mp *Provider) SessionExist(ctx context.Context, sid string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (mp *Provider) SessionRegenerate(ctx context.Context, oldsid, sid string) (session.Store, error) {
	_ = "STUB: not implemented"
	return *new(session.Store), nil
}

func (mp *Provider) SessionDestroy(ctx context.Context, sid string) error {
	_ = "STUB: not implemented"
	return nil
}

func (mp *Provider) SessionGC(context.Context) { _ = "STUB: not implemented"; return }

func (mp *Provider) SessionAll(context.Context) int { _ = "STUB: not implemented"; return 0 }

func init() {
	session.Register("postgresql", postgresqlpder)
}
