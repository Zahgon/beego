package ratelimit

import (
	"sync"
	"time"
)

type tokenBucket struct {
	sync.RWMutex
	remaining   uint
	capacity    uint
	lastCheckAt time.Time
	rate        time.Duration
}

func newTokenBucket(opts ...bucketOption) bucket { _ = "STUB: not implemented"; return *new(bucket) }

func withCapacity(capacity uint) bucketOption { _ = "STUB: not implemented"; return *new(bucketOption) }

func withRate(rate time.Duration) bucketOption {
	_ = "STUB: not implemented"
	return *new(bucketOption)
}

func (b *tokenBucket) getRemaining() uint { _ = "STUB: not implemented"; return 0 }

func (b *tokenBucket) getRate() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (b *tokenBucket) getCapacity() uint { _ = "STUB: not implemented"; return 0 }

func (b *tokenBucket) take(amount uint) bool { _ = "STUB: not implemented"; return false }
