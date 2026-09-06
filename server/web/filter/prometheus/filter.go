package prometheus

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

const unknownRouterPattern = "UnknownRouterPattern"

type FilterChainBuilder struct{}

var (
	summaryVec     prometheus.ObserverVec
	initSummaryVec sync.Once
)

func (builder *FilterChainBuilder) FilterChain(next web.FilterFunc) web.FilterFunc {
	_ = "STUB: not implemented"
	return *new(web.FilterFunc)
}

func (builder *FilterChainBuilder) buildVec() *prometheus.SummaryVec {
	_ = "STUB: not implemented"
	return nil
}

func registerBuildInfo() { _ = "STUB: not implemented"; return }

func report(dur time.Duration, ctx *context.Context, vec prometheus.ObserverVec) {
	_ = "STUB: not implemented"
	return
}
