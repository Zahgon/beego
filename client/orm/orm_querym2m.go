package orm

import (
	"context"
	"reflect"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

type queryM2M struct {
	md  interface{}
	mi  *models.ModelInfo
	fi  *models.FieldInfo
	qs  *querySet
	ind reflect.Value
}

func (o *queryM2M) Add(mds ...interface{}) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (o *queryM2M) AddWithCtx(ctx context.Context, mds ...interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *queryM2M) Remove(mds ...interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *queryM2M) RemoveWithCtx(ctx context.Context, mds ...interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *queryM2M) Exist(md interface{}) bool { _ = "STUB: not implemented"; return false }

func (o *queryM2M) ExistWithCtx(ctx context.Context, md interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (o *queryM2M) Clear() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (o *queryM2M) ClearWithCtx(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *queryM2M) Count() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (o *queryM2M) CountWithCtx(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var _ QueryM2Mer = new(queryM2M)

func newQueryM2M(md interface{}, o *ormBase, mi *models.ModelInfo, fi *models.FieldInfo, ind reflect.Value) QueryM2Mer {
	_ = "STUB: not implemented"
	return *new(QueryM2Mer)
}
