package redis_cluster

import (
	"context"
	"net/http"
	"sync"
	"time"

	rediss "github.com/redis/go-redis/v9"

	"github.com/beego/beego/v2/server/web/session"
)

var redispder = &Provider{}

var MaxPoolSize = 1000

type SessionStore struct {
	p           *rediss.ClusterClient
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

func (rs *SessionStore) releaseSession(ctx context.Context, _ http.ResponseWriter, requirePresent bool) {
	_ = "STUB: not implemented"
	return
}

type Provider struct {
	maxlifetime int64
	SavePath    string `json:"save_path"`
	Poolsize    int    `json:"poolsize"`
	Password    string `json:"password"`
	DbNum       int    `json:"db_num"`

	idleTimeout    time.Duration
	IdleTimeoutStr string `json:"idle_timeout"`

	idleCheckFrequency    time.Duration
	IdleCheckFrequencyStr string `json:"idle_check_frequency"`
	MaxRetries            int    `json:"max_retries"`
	poollist              *rediss.ClusterClient
}

func (rp *Provider) SessionInit(ctx context.Context, maxlifetime int64, cfgStr string) error {
	_ = "STUB: not implemented"
	return nil
}

func (rp *Provider) initOldStyle(savePath string) { _ = "STUB: not implemented"; return }

func (rp *Provider) SessionRead(ctx context.Context, sid string) (session.Store, error) {
	_ = "STUB: not implemented"
	return *new(session.Store), nil
}

func (rp *Provider) SessionExist(ctx context.Context, sid string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (rp *Provider) SessionRegenerate(ctx context.Context, oldsid, sid string) (session.Store, error) {
	_ = "STUB: not implemented"
	return *new(session.Store), nil
}

func (rp *Provider) SessionDestroy(ctx context.Context, sid string) error {
	_ = "STUB: not implemented"
	return nil
}

func (rp *Provider) SessionGC(context.Context) { _ = "STUB: not implemented"; return }

func (rp *Provider) SessionAll(context.Context) int { _ = "STUB: not implemented"; return 0 }

func init() {
	session.Register("redis_cluster", redispder)
}
