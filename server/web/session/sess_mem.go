package session

import (
	"container/list"
	"context"
	"net/http"
	"sync"
	"time"
)

var mempder = &MemProvider{list: list.New(), sessions: make(map[string]*list.Element)}

type MemSessionStore struct {
	sid          string
	timeAccessed time.Time
	value        map[interface{}]interface{}
	lock         sync.RWMutex
}

func (st *MemSessionStore) Set(ctx context.Context, key, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (st *MemSessionStore) Get(ctx context.Context, key interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (st *MemSessionStore) Delete(ctx context.Context, key interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (st *MemSessionStore) Flush(context.Context) error { _ = "STUB: not implemented"; return nil }

func (st *MemSessionStore) SessionID(context.Context) string { _ = "STUB: not implemented"; return "" }

func (st *MemSessionStore) SessionRelease(_ context.Context, _ http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

func (*MemSessionStore) SessionReleaseIfPresent(_ context.Context, _ http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

type MemProvider struct {
	lock        sync.RWMutex
	sessions    map[string]*list.Element
	list        *list.List
	maxlifetime int64
	savePath    string
}

func (pder *MemProvider) SessionInit(ctx context.Context, maxlifetime int64, savePath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (pder *MemProvider) SessionRead(ctx context.Context, sid string) (Store, error) {
	_ = "STUB: not implemented"
	return *new(Store), nil
}

func (pder *MemProvider) SessionExist(ctx context.Context, sid string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (pder *MemProvider) SessionRegenerate(ctx context.Context, oldsid, sid string) (Store, error) {
	_ = "STUB: not implemented"
	return *new(Store), nil
}

func (pder *MemProvider) SessionDestroy(ctx context.Context, sid string) error {
	_ = "STUB: not implemented"
	return nil
}

func (pder *MemProvider) SessionGC(context.Context) { _ = "STUB: not implemented"; return }

func (pder *MemProvider) SessionAll(context.Context) int { _ = "STUB: not implemented"; return 0 }

func (pder *MemProvider) SessionUpdate(ctx context.Context, sid string) error {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	Register("memory", mempder)
}
