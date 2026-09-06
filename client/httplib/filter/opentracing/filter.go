package opentracing

import (
	"context"
	"net/http"

	"github.com/opentracing/opentracing-go"

	"github.com/beego/beego/v2/client/httplib"
)

type FilterChainBuilder struct {
	TagURL bool

	CustomSpanFunc func(span opentracing.Span, ctx context.Context,
		req *httplib.BeegoHTTPRequest, resp *http.Response, err error)
}

func (builder *FilterChainBuilder) FilterChain(next httplib.Filter) httplib.Filter {
	_ = "STUB: not implemented"
	return *new(httplib.Filter)
}
