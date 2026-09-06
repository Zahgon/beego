package cache

import (
	"context"
	"time"
)

type RandomExpireCacheOption func(*RandomExpireCache)

func WithRandomExpireOffsetFunc(fn func() time.Duration) RandomExpireCacheOption {
	_ = "STUB: not implemented"
	return *new(RandomExpireCacheOption)
}

type RandomExpireCache struct {
	Cache
	offset func() time.Duration
}

func (rec *RandomExpireCache) Put(ctx context.Context, key string, val interface{}, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRandomExpireCache(adapter Cache, opts ...RandomExpireCacheOption) Cache {
	_ = "STUB: not implemented"
	return *new(Cache)
}

func defaultExpiredFunc() func() time.Duration { _ = "STUB: not implemented"; return nil }
