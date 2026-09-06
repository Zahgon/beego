package cache

import (
	"context"
	"time"
)

type WriteDeleteCache struct {
	Cache
	storeFunc func(ctx context.Context, key string, val any) error
}

func NewWriteDeleteCache(cache Cache, fn func(ctx context.Context, key string, val any) error) (*WriteDeleteCache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *WriteDeleteCache) Set(ctx context.Context, key string, val any) error {
	_ = "STUB: not implemented"
	return nil
}

type WriteDoubleDeleteCache struct {
	Cache
	interval  time.Duration
	timeout   time.Duration
	storeFunc func(ctx context.Context, key string, val any) error
}

type WriteDoubleDeleteCacheOption func(c *WriteDoubleDeleteCache)

func NewWriteDoubleDeleteCache(cache Cache, interval, timeout time.Duration,
	fn func(ctx context.Context, key string, val any) error) (*WriteDoubleDeleteCache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *WriteDoubleDeleteCache) Set(
	ctx context.Context, key string, val any) error {
	_ = "STUB: not implemented"
	return nil
}
