package session

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

type Store interface {
	Set(ctx context.Context, key, value interface{}) error
	Get(ctx context.Context, key interface{}) interface{}
	Delete(ctx context.Context, key interface{}) error
	SessionID(ctx context.Context) string
	SessionReleaseIfPresent(ctx context.Context, w http.ResponseWriter)
	SessionRelease(ctx context.Context, w http.ResponseWriter)
	Flush(ctx context.Context) error
}

type Provider interface {
	SessionInit(ctx context.Context, gclifetime int64, config string) error
	SessionRead(ctx context.Context, sid string) (Store, error)
	SessionExist(ctx context.Context, sid string) (bool, error)
	SessionRegenerate(ctx context.Context, oldsid, sid string) (Store, error)
	SessionDestroy(ctx context.Context, sid string) error
	SessionAll(ctx context.Context) int
	SessionGC(ctx context.Context)
}

var provides = make(map[string]Provider)

var SLogger = NewSessionLog(os.Stderr)

func Register(name string, provide Provider) { _ = "STUB: not implemented"; return }

func GetProvider(name string) (Provider, error) {
	_ = "STUB: not implemented"
	return *new(Provider), nil
}

type Manager struct {
	provider Provider
	config   *ManagerConfig
}

func NewManager(provideName string, cf *ManagerConfig) (*Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (manager *Manager) GetProvider() Provider { _ = "STUB: not implemented"; return *new(Provider) }

func (manager *Manager) getSid(r *http.Request) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Store, err error) {
	_ = "STUB: not implemented"
	return *new(Store), nil
}

func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (manager *Manager) GetSessionStore(sid string) (sessions Store, err error) {
	_ = "STUB: not implemented"
	return *new(Store), nil
}

func (manager *Manager) GC() { _ = "STUB: not implemented"; return }

func (manager *Manager) SessionRegenerateID(w http.ResponseWriter, r *http.Request) (Store, error) {
	_ = "STUB: not implemented"
	return *new(Store), nil
}

func (manager *Manager) GetActiveSession() int { _ = "STUB: not implemented"; return 0 }

func (manager *Manager) SetSecure(secure bool) { _ = "STUB: not implemented"; return }

func (manager *Manager) sessionID() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (manager *Manager) isSecure(req *http.Request) bool { _ = "STUB: not implemented"; return false }

type Log struct {
	*log.Logger
}

func NewSessionLog(out io.Writer) *Log { _ = "STUB: not implemented"; return nil }
