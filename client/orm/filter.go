package orm

import (
	"context"
)

type FilterChain func(next Filter) Filter

type Filter func(ctx context.Context, inv *Invocation) []interface{}

var globalFilterChains = make([]FilterChain, 0, 4)

func AddGlobalFilterChain(filterChain ...FilterChain) { _ = "STUB: not implemented"; return }
