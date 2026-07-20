package web

import (
	"github.com/beego/beego/v2/server/web/context"
)

type FilterChain func(next FilterFunc) FilterFunc

type FilterFunc = HandleFunc

type FilterRouter struct {
	filterFunc     FilterFunc
	next           *FilterRouter
	tree           *Tree
	pattern        string
	returnOnOutput bool
	resetParams    bool
}

func newFilterRouter(pattern string, filter FilterFunc, opts ...FilterOpt) *FilterRouter {
	_ = "STUB: not implemented"
	return nil
}

func (f *FilterRouter) filter(ctx *context.Context, urlPath string, preFilterParams map[string]string) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (f *FilterRouter) ValidRouter(url string, ctx *context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

type filterOpts struct {
	returnOnOutput      bool
	resetParams         bool
	routerCaseSensitive bool
}

type FilterOpt func(opts *filterOpts)

func WithReturnOnOutput(ret bool) FilterOpt { _ = "STUB: not implemented"; return *new(FilterOpt) }

func WithResetParams(reset bool) FilterOpt { _ = "STUB: not implemented"; return *new(FilterOpt) }

func WithCaseSensitive(sensitive bool) FilterOpt { _ = "STUB: not implemented"; return *new(FilterOpt) }
