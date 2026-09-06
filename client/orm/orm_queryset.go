package orm

import (
	"context"

	"github.com/beego/beego/v2/client/orm/internal/models"

	"github.com/beego/beego/v2/client/orm/clauses/order_clause"
)

type colValue struct {
	value int64
	opt   operator
}

type operator int

const (
	ColAdd operator = iota
	ColMinus
	ColMultiply
	ColExcept
	ColBitAnd
	ColBitRShift
	ColBitLShift
	ColBitXOR
	ColBitOr
)

func ColValue(opt operator, value interface{}) interface{} { _ = "STUB: not implemented"; return nil }

type querySet struct {
	mi        *models.ModelInfo
	cond      *Condition
	related   []string
	relDepth  int
	limit     int64
	offset    int64
	groups    []string
	orders    []*order_clause.Order
	distinct  bool
	forUpdate bool
	useIndex  int
	indexes   []string
	orm       *ormBase
	aggregate string
}

var _ QuerySeter = new(querySet)

func (o querySet) Filter(expr string, args ...interface{}) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (o querySet) FilterRaw(expr string, sql string) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (o querySet) Exclude(expr string, args ...interface{}) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (o *querySet) setOffset(num interface{}) { _ = "STUB: not implemented"; return }

func (o querySet) Limit(limit interface{}, args ...interface{}) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (o querySet) Offset(offset interface{}) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (o querySet) GroupBy(exprs ...string) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (o querySet) OrderBy(expressions ...string) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (o querySet) OrderClauses(orders ...*order_clause.Order) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (o querySet) Distinct() QuerySeter { _ = "STUB: not implemented"; return *new(QuerySeter) }

func (o querySet) ForUpdate() QuerySeter { _ = "STUB: not implemented"; return *new(QuerySeter) }

func (o querySet) ForceIndex(indexes ...string) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (o querySet) UseIndex(indexes ...string) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (o querySet) IgnoreIndex(indexes ...string) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (o querySet) RelatedSel(params ...interface{}) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (o querySet) SetCond(cond *Condition) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (o querySet) GetCond() *Condition { _ = "STUB: not implemented"; return nil }

func (o querySet) Count() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (o querySet) CountWithCtx(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o querySet) Exist() bool { _ = "STUB: not implemented"; return false }

func (o querySet) ExistWithCtx(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func (o querySet) Update(values Params) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (o querySet) UpdateWithCtx(ctx context.Context, values Params) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o querySet) Delete() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (o querySet) DeleteWithCtx(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o querySet) PrepareInsert() (Inserter, error) {
	_ = "STUB: not implemented"
	return *new(Inserter), nil
}

func (o querySet) PrepareInsertWithCtx(ctx context.Context) (Inserter, error) {
	_ = "STUB: not implemented"
	return *new(Inserter), nil
}

func (o querySet) All(container interface{}, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o querySet) AllWithCtx(ctx context.Context, container interface{}, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o querySet) One(container interface{}, cols ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (o querySet) OneWithCtx(ctx context.Context, container interface{}, cols ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (o querySet) Values(results *[]Params, exprs ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o querySet) ValuesWithCtx(ctx context.Context, results *[]Params, exprs ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o querySet) ValuesList(results *[]ParamsList, exprs ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o querySet) ValuesListWithCtx(ctx context.Context, results *[]ParamsList, exprs ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o querySet) ValuesFlat(result *ParamsList, expr string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o querySet) ValuesFlatWithCtx(ctx context.Context, result *ParamsList, expr string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o querySet) RowsToMap(result *Params, keyCol, valueCol string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o querySet) RowsToStruct(ptrStruct interface{}, keyCol, valueCol string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func newQuerySet(orm *ormBase, mi *models.ModelInfo) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (o querySet) Aggregate(s string) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}
