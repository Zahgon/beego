package memcache

import (
	"context"
	"time"

	"github.com/bradfitz/gomemcache/memcache"

	"github.com/beego/beego/v2/client/cache"
)

type Cache struct {
	conn     *memcache.Client
	conninfo []string
}

func NewMemCache() cache.Cache { _ = "STUB: not implemented"; return *new(cache.Cache) }

func (rc *Cache) Get(ctx context.Context, key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rc *Cache) GetMulti(ctx context.Context, keys []string) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rc *Cache) Put(ctx context.Context, key string, val interface{}, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (rc *Cache) Delete(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (rc *Cache) Incr(ctx context.Context, key string) error { _ = "STUB: not implemented"; return nil }

func (rc *Cache) Decr(ctx context.Context, key string) error { _ = "STUB: not implemented"; return nil }

func (rc *Cache) IsExist(ctx context.Context, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (rc *Cache) ClearAll(context.Context) error { _ = "STUB: not implemented"; return nil }

func (rc *Cache) StartAndGC(config string) error { _ = "STUB: not implemented"; return nil }

func init() {
	cache.Register("memcache", NewMemCache)
}
