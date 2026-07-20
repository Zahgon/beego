package bean

import (
	"context"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/bean"
)

type DefaultValueFilterChainBuilder struct {
	factory                bean.AutoWireBeanFactory
	compatibleWithOldStyle bool

	includeInsertOrUpdate bool
}

func NewDefaultValueFilterChainBuilder(typeAdapters map[string]bean.TypeAdapter,
	includeInsertOrUpdate bool,
	compatibleWithOldStyle bool) *DefaultValueFilterChainBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DefaultValueFilterChainBuilder) FilterChain(next orm.Filter) orm.Filter {
	_ = "STUB: not implemented"
	return *new(orm.Filter)
}

func (d *DefaultValueFilterChainBuilder) handleInsert(ctx context.Context, inv *orm.Invocation) {
	_ = "STUB: not implemented"
	return
}

func (d *DefaultValueFilterChainBuilder) handleInsertOrUpdate(ctx context.Context, inv *orm.Invocation) {
	_ = "STUB: not implemented"
	return
}

func (d *DefaultValueFilterChainBuilder) handleInsertMulti(ctx context.Context, inv *orm.Invocation) {
	_ = "STUB: not implemented"
	return
}

func (d *DefaultValueFilterChainBuilder) setDefaultValue(ctx context.Context, ins interface{}) {
	_ = "STUB: not implemented"
	return
}
