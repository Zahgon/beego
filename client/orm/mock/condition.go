package mock

import (
	"context"

	"github.com/beego/beego/v2/client/orm"
)

type Mock struct {
	cond Condition
	resp []interface{}
	cb   func(inv *orm.Invocation)
}

func NewMock(cond Condition, resp []interface{}, cb func(inv *orm.Invocation)) *Mock {
	_ = "STUB: not implemented"
	return nil
}

type Condition interface {
	Match(ctx context.Context, inv *orm.Invocation) bool
}

type SimpleCondition struct {
	tableName string
	method    string
}

func NewSimpleCondition(tableName string, methodName string) Condition {
	_ = "STUB: not implemented"
	return *new(Condition)
}

func (s *SimpleCondition) Match(ctx context.Context, inv *orm.Invocation) bool {
	_ = "STUB: not implemented"
	return false
}
