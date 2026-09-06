package cache

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) (interface{}, error)

	GetMulti(ctx context.Context, keys []string) ([]interface{}, error)

	Put(ctx context.Context, key string, val interface{}, timeout time.Duration) error

	Delete(ctx context.Context, key string) error

	Incr(ctx context.Context, key string) error

	Decr(ctx context.Context, key string) error

	IsExist(ctx context.Context, key string) (bool, error)

	ClearAll(ctx context.Context) error

	StartAndGC(config string) error
}

type Instance func() Cache

var adapters = make(map[string]Instance)

func Register(name string, adapter Instance) { _ = "STUB: not implemented"; return }

func NewCache(adapterName, config string) (adapter Cache, err error) {
	_ = "STUB: not implemented"
	return *new(Cache), nil
}
