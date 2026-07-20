package opentracing

import (
	"github.com/opentracing/opentracing-go"

	"github.com/beego/beego/v2/server/web"
	beegoCtx "github.com/beego/beego/v2/server/web/context"
)

type FilterChainBuilder struct {
	CustomSpanFunc func(span opentracing.Span, ctx *beegoCtx.Context)
}

func (builder *FilterChainBuilder) FilterChain(next web.FilterFunc) web.FilterFunc {
	_ = "STUB: not implemented"
	return *new(web.FilterFunc)
}

func (builder *FilterChainBuilder) operationName(ctx *beegoCtx.Context) string {
	_ = "STUB: not implemented"
	return ""
}
