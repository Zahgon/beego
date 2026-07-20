package ssdb

import (
	"context"
	"net/http"
	"sync"

	"github.com/ssdb/gossdb/ssdb"

	"github.com/beego/beego/v2/server/web/session"
)

var ssdbProvider = &Provider{}

type Provider struct {
	client      *ssdb.Client
	Host        string `json:"host"`
	Port        int    `json:"port"`
	maxLifetime int64
}

func (p *Provider) connectInit() error { _ = "STUB: not implemented"; return nil }

func (p *Provider) SessionInit(ctx context.Context, maxLifetime int64, cfg string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) initOldStyle(savePath string) error { _ = "STUB: not implemented"; return nil }

func (p *Provider) SessionRead(ctx context.Context, sid string) (session.Store, error) {
	_ = "STUB: not implemented"
	return *new(session.Store), nil
}

func (p *Provider) SessionExist(ctx context.Context, sid string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *Provider) SessionRegenerate(ctx context.Context, oldsid, sid string) (session.Store, error) {
	_ = "STUB: not implemented"
	return *new(session.Store), nil
}

func (p *Provider) SessionDestroy(ctx context.Context, sid string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) SessionGC(context.Context) { _ = "STUB: not implemented"; return }

func (p *Provider) SessionAll(context.Context) int { _ = "STUB: not implemented"; return 0 }

type SessionStore struct {
	sid         string
	lock        sync.RWMutex
	values      map[interface{}]interface{}
	maxLifetime int64
	client      *ssdb.Client
}

func (s *SessionStore) Set(ctx context.Context, key, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SessionStore) Get(ctx context.Context, key interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (s *SessionStore) Delete(ctx context.Context, key interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SessionStore) Flush(context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *SessionStore) SessionID(context.Context) string { _ = "STUB: not implemented"; return "" }

func (s *SessionStore) SessionRelease(_ context.Context, _ http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

func (s *SessionStore) SessionReleaseIfPresent(c context.Context, w http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

func init() {
	session.Register("ssdb", ssdbProvider)
}
