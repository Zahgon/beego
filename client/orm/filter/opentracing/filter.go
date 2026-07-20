package opentracing

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/beego/beego/v2/client/orm"
)

type FilterChainBuilder struct {
	CustomSpanFunc func(span opentracing.Span, ctx context.Context, inv *orm.Invocation)
}

func (builder *FilterChainBuilder) FilterChain(next orm.Filter) orm.Filter {
	_ = "STUB: not implemented"
	return *new(orm.Filter)
}

func (builder *FilterChainBuilder) buildSpan(span opentracing.Span, ctx context.Context, inv *orm.Invocation) {
	_ = "STUB: not implemented"
	return
}

func (builder *FilterChainBuilder) operationName(ctx context.Context, inv *orm.Invocation) string {
	_ = "STUB: not implemented"
	return ""
}
