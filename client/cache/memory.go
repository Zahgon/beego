package cache

import (
	"context"
	"sync"
	"time"
)

var DefaultEvery = 60

type MemoryItem struct {
	val         interface{}
	createdTime time.Time
	lifespan    time.Duration
}

func (mi *MemoryItem) isExpire() bool { _ = "STUB: not implemented"; return false }

type MemoryCache struct {
	sync.RWMutex
	dur   time.Duration
	items map[string]*MemoryItem
	Every int
}

func NewMemoryCache() Cache { _ = "STUB: not implemented"; return *new(Cache) }

func (bc *MemoryCache) Get(ctx context.Context, key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bc *MemoryCache) GetMulti(ctx context.Context, keys []string) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bc *MemoryCache) Put(ctx context.Context, key string, val interface{}, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (bc *MemoryCache) Delete(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (bc *MemoryCache) Incr(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (bc *MemoryCache) Decr(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (bc *MemoryCache) IsExist(ctx context.Context, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (bc *MemoryCache) ClearAll(context.Context) error { _ = "STUB: not implemented"; return nil }

func (bc *MemoryCache) StartAndGC(config string) error { _ = "STUB: not implemented"; return nil }

func (bc *MemoryCache) vacuum() { _ = "STUB: not implemented"; return }

func (bc *MemoryCache) expiredKeys() (keys []string) { _ = "STUB: not implemented"; return nil }

func (bc *MemoryCache) clearItems(keys []string) { _ = "STUB: not implemented"; return }

func init() {
	Register("memory", NewMemoryCache)
}
