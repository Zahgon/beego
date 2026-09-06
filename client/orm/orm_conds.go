package orm

import (
	"github.com/beego/beego/v2/client/orm/clauses"
)

const (
	ExprSep = clauses.ExprSep
)

type condValue struct {
	exprs  []string
	args   []interface{}
	cond   *Condition
	isOr   bool
	isNot  bool
	isCond bool
	isRaw  bool
	sql    string
}

type Condition struct {
	params []condValue
}

func NewCondition() *Condition { _ = "STUB: not implemented"; return nil }

func (c Condition) Raw(expr string, sql string) *Condition { _ = "STUB: not implemented"; return nil }

func (c Condition) And(expr string, args ...interface{}) *Condition {
	_ = "STUB: not implemented"
	return nil
}

func (c Condition) AndNot(expr string, args ...interface{}) *Condition {
	_ = "STUB: not implemented"
	return nil
}

func (c *Condition) AndCond(cond *Condition) *Condition { _ = "STUB: not implemented"; return nil }

func (c *Condition) AndNotCond(cond *Condition) *Condition { _ = "STUB: not implemented"; return nil }

func (c Condition) Or(expr string, args ...interface{}) *Condition {
	_ = "STUB: not implemented"
	return nil
}

func (c Condition) OrNot(expr string, args ...interface{}) *Condition {
	_ = "STUB: not implemented"
	return nil
}

func (c *Condition) OrCond(cond *Condition) *Condition { _ = "STUB: not implemented"; return nil }

func (c *Condition) OrNotCond(cond *Condition) *Condition { _ = "STUB: not implemented"; return nil }

func (c *Condition) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (c Condition) clone() *Condition { _ = "STUB: not implemented"; return nil }
