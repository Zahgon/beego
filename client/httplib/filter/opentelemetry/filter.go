package opentelemetry

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/trace"

	"github.com/beego/beego/v2/client/httplib"
)

type CustomSpanFunc func(span trace.Span, ctx context.Context, req *httplib.BeegoHTTPRequest, resp *http.Response, err error)

type OtelFilterChainBuilder struct {
	tagURL bool

	customSpanFunc CustomSpanFunc
}

func NewOpenTelemetryFilter(tagURL bool, spanFunc CustomSpanFunc) *OtelFilterChainBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (builder *OtelFilterChainBuilder) FilterChain(next httplib.Filter) httplib.Filter {
	_ = "STUB: not implemented"
	return *new(httplib.Filter)
}
