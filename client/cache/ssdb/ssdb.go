package ssdb

import (
	"context"
	"time"

	"github.com/ssdb/gossdb/ssdb"

	"github.com/beego/beego/v2/client/cache"
)

type Cache struct {
	conn     *ssdb.Client
	conninfo []string
}

func NewSsdbCache() cache.Cache { _ = "STUB: not implemented"; return *new(cache.Cache) }

func (rc *Cache) Get(ctx context.Context, key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rc *Cache) GetMulti(ctx context.Context, keys []string) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rc *Cache) DelMulti(keys []string) error { _ = "STUB: not implemented"; return nil }

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

func (rc *Cache) Scan(keyStart string, keyEnd string, limit int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rc *Cache) StartAndGC(config string) error { _ = "STUB: not implemented"; return nil }

func (rc *Cache) connectInit() error { _ = "STUB: not implemented"; return nil }

func init() {
	cache.Register("ssdb", NewSsdbCache)
}
