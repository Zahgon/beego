package ratelimit

import "time"

type bucket interface {
	take(amount uint) bool
	getCapacity() uint
	getRemaining() uint
	getRate() time.Duration
}

type bucketOption func(bucket)
