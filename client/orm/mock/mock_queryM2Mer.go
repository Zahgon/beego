package mock

import (
	"context"

	"github.com/beego/beego/v2/client/orm"
)

type DoNothingQueryM2Mer struct{}

func (d *DoNothingQueryM2Mer) AddWithCtx(ctx context.Context, i ...interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQueryM2Mer) RemoveWithCtx(ctx context.Context, i ...interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQueryM2Mer) ExistWithCtx(ctx context.Context, i interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *DoNothingQueryM2Mer) ClearWithCtx(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQueryM2Mer) CountWithCtx(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQueryM2Mer) Add(i ...interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQueryM2Mer) Remove(i ...interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingQueryM2Mer) Exist(i interface{}) bool { _ = "STUB: not implemented"; return false }

func (d *DoNothingQueryM2Mer) Clear() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *DoNothingQueryM2Mer) Count() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

type QueryM2MerCondition struct {
	tableName string
	name      string
}

func NewQueryM2MerCondition(tableName string, name string) *QueryM2MerCondition {
	_ = "STUB: not implemented"
	return nil
}

func (q *QueryM2MerCondition) Match(ctx context.Context, inv *orm.Invocation) bool {
	_ = "STUB: not implemented"
	return false
}
