package mock

import (
	"context"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/client/orm/clauses/order_clause"
)

type DoNothingQuerySetter struct{}

func (d *DoNothingQuerySetter) OrderClauses(orders ...*order_clause.Order) orm.QuerySeter {
	_ = "STUB: not implemented"
	return *new(orm.QuerySeter)
}

func (d *DoNothingQuerySetter) CountWithCtx(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQuerySetter) ExistWithCtx(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *DoNothingQuerySetter) UpdateWithCtx(ctx context.Context, values orm.Params) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQuerySetter) DeleteWithCtx(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQuerySetter) PrepareInsertWithCtx(ctx context.Context) (orm.Inserter, error) {
	_ = "STUB: not implemented"
	return *new(orm.Inserter), nil
}

func (d *DoNothingQuerySetter) AllWithCtx(ctx context.Context, container interface{}, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQuerySetter) OneWithCtx(ctx context.Context, container interface{}, cols ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DoNothingQuerySetter) ValuesWithCtx(ctx context.Context, results *[]orm.Params, exprs ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQuerySetter) ValuesListWithCtx(ctx context.Context, results *[]orm.ParamsList, exprs ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQuerySetter) ValuesFlatWithCtx(ctx context.Context, result *orm.ParamsList, expr string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQuerySetter) Aggregate(s string) orm.QuerySeter {
	_ = "STUB: not implemented"
	return *new(orm.QuerySeter)
}

func (d *DoNothingQuerySetter) Filter(s string, i ...interface{}) orm.QuerySeter {
	_ = "STUB: not implemented"
	return *new(orm.QuerySeter)
}

func (d *DoNothingQuerySetter) FilterRaw(s string, s2 string) orm.QuerySeter {
	_ = "STUB: not implemented"
	return *new(orm.QuerySeter)
}

func (d *DoNothingQuerySetter) Exclude(s string, i ...interface{}) orm.QuerySeter {
	_ = "STUB: not implemented"
	return *new(orm.QuerySeter)
}

func (d *DoNothingQuerySetter) SetCond(condition *orm.Condition) orm.QuerySeter {
	_ = "STUB: not implemented"
	return *new(orm.QuerySeter)
}

func (d *DoNothingQuerySetter) GetCond() *orm.Condition { _ = "STUB: not implemented"; return nil }

func (d *DoNothingQuerySetter) Limit(limit interface{}, args ...interface{}) orm.QuerySeter {
	_ = "STUB: not implemented"
	return *new(orm.QuerySeter)
}

func (d *DoNothingQuerySetter) Offset(offset interface{}) orm.QuerySeter {
	_ = "STUB: not implemented"
	return *new(orm.QuerySeter)
}

func (d *DoNothingQuerySetter) GroupBy(exprs ...string) orm.QuerySeter {
	_ = "STUB: not implemented"
	return *new(orm.QuerySeter)
}

func (d *DoNothingQuerySetter) OrderBy(exprs ...string) orm.QuerySeter {
	_ = "STUB: not implemented"
	return *new(orm.QuerySeter)
}

func (d *DoNothingQuerySetter) ForceIndex(indexes ...string) orm.QuerySeter {
	_ = "STUB: not implemented"
	return *new(orm.QuerySeter)
}

func (d *DoNothingQuerySetter) UseIndex(indexes ...string) orm.QuerySeter {
	_ = "STUB: not implemented"
	return *new(orm.QuerySeter)
}

func (d *DoNothingQuerySetter) IgnoreIndex(indexes ...string) orm.QuerySeter {
	_ = "STUB: not implemented"
	return *new(orm.QuerySeter)
}

func (d *DoNothingQuerySetter) RelatedSel(params ...interface{}) orm.QuerySeter {
	_ = "STUB: not implemented"
	return *new(orm.QuerySeter)
}

func (d *DoNothingQuerySetter) Distinct() orm.QuerySeter {
	_ = "STUB: not implemented"
	return *new(orm.QuerySeter)
}

func (d *DoNothingQuerySetter) ForUpdate() orm.QuerySeter {
	_ = "STUB: not implemented"
	return *new(orm.QuerySeter)
}

func (d *DoNothingQuerySetter) Count() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *DoNothingQuerySetter) Exist() bool { _ = "STUB: not implemented"; return false }

func (d *DoNothingQuerySetter) Update(values orm.Params) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQuerySetter) Delete() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *DoNothingQuerySetter) PrepareInsert() (orm.Inserter, error) {
	_ = "STUB: not implemented"
	return *new(orm.Inserter), nil
}

func (d *DoNothingQuerySetter) All(container interface{}, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQuerySetter) One(container interface{}, cols ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DoNothingQuerySetter) Values(results *[]orm.Params, exprs ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQuerySetter) ValuesList(results *[]orm.ParamsList, exprs ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQuerySetter) ValuesFlat(result *orm.ParamsList, expr string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQuerySetter) RowsToMap(result *orm.Params, keyCol, valueCol string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQuerySetter) RowsToStruct(ptrStruct interface{}, keyCol, valueCol string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
