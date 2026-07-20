package prometheus

import (
	"context"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/beego/beego/v2/client/orm"
)

type FilterChainBuilder struct {
	AppName    string
	ServerName string
	RunMode    string
}

var (
	summaryVec     prometheus.ObserverVec
	initSummaryVec sync.Once
)

func (builder *FilterChainBuilder) FilterChain(next orm.Filter) orm.Filter {
	_ = "STUB: not implemented"
	return *new(orm.Filter)
}

func (builder *FilterChainBuilder) report(ctx context.Context, inv *orm.Invocation, dur time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (builder *FilterChainBuilder) reportTxn(ctx context.Context, inv *orm.Invocation) {
	_ = "STUB: not implemented"
	return
}
