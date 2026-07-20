package opentelemetry

import (
	"context"

	otelTrace "go.opentelemetry.io/otel/trace"

	"github.com/beego/beego/v2/client/orm"
)

type (
	CustomSpanFunc    func(ctx context.Context, span otelTrace.Span, inv *orm.Invocation)
	FilterChainOption func(fcv *FilterChainBuilder)
)

type FilterChainBuilder struct {
	customSpanFunc CustomSpanFunc
}

func NewFilterChainBuilder(options ...FilterChainOption) *FilterChainBuilder {
	_ = "STUB: not implemented"
	return nil
}

func WithCustomSpanFunc(customSpanFunc CustomSpanFunc) FilterChainOption {
	_ = "STUB: not implemented"
	return *new(FilterChainOption)
}

func (builder *FilterChainBuilder) FilterChain(next orm.Filter) orm.Filter {
	_ = "STUB: not implemented"
	return *new(orm.Filter)
}

func (builder *FilterChainBuilder) buildSpan(ctx context.Context, span otelTrace.Span, inv *orm.Invocation) {
	_ = "STUB: not implemented"
	return
}

func invOperationName(ctx context.Context, inv *orm.Invocation) string {
	_ = "STUB: not implemented"
	return ""
}
