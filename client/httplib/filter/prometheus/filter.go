package prometheus

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/beego/beego/v2/client/httplib"
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

func (builder *FilterChainBuilder) FilterChain(next httplib.Filter) httplib.Filter {
	_ = "STUB: not implemented"
	return *new(httplib.Filter)
}

func (builder *FilterChainBuilder) report(startTime time.Time, endTime time.Time,
	ctx context.Context, req *httplib.BeegoHTTPRequest, resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return
}
