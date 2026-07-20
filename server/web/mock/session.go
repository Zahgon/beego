package mock

import (
	"context"
	"net/http"

	"github.com/beego/beego/v2/server/web/session"
)

func NewSessionProvider(name string) *SessionProvider { _ = "STUB: not implemented"; return nil }

type SessionProvider struct {
	Store *SessionStore
}

func newSessionProvider() *SessionProvider { _ = "STUB: not implemented"; return nil }

func (s *SessionProvider) SessionInit(ctx context.Context, gclifetime int64, config string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SessionProvider) SessionRead(ctx context.Context, sid string) (session.Store, error) {
	_ = "STUB: not implemented"
	return *new(session.Store), nil
}

func (s *SessionProvider) SessionExist(ctx context.Context, sid string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *SessionProvider) SessionRegenerate(ctx context.Context, oldsid, sid string) (session.Store, error) {
	_ = "STUB: not implemented"
	return *new(session.Store), nil
}

func (s *SessionProvider) SessionDestroy(ctx context.Context, sid string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SessionProvider) SessionAll(ctx context.Context) int { _ = "STUB: not implemented"; return 0 }

func (s *SessionProvider) SessionGC(ctx context.Context) { _ = "STUB: not implemented"; return }

type SessionStore struct {
	sid    string
	values map[interface{}]interface{}
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

func (s *SessionStore) SessionID(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func (s *SessionStore) SessionRelease(_ context.Context, _ http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

func (*SessionStore) SessionReleaseIfPresent(_ context.Context, _ http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

func (s *SessionStore) Flush(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func newSessionStore() *SessionStore { _ = "STUB: not implemented"; return nil }
