package cache

import (
	"context"
	"time"
)

type WriteThroughCache struct {
	Cache
	storeFunc func(ctx context.Context, key string, val any) error
}

func NewWriteThroughCache(cache Cache, fn func(ctx context.Context, key string, val any) error) (*WriteThroughCache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *WriteThroughCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
