package orm

import (
	"context"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

type insertSet struct {
	mi     *models.ModelInfo
	orm    *ormBase
	stmt   stmtQuerier
	closed bool
}

var _ Inserter = new(insertSet)

func (o *insertSet) Insert(md interface{}) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (o *insertSet) InsertWithCtx(ctx context.Context, md interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *insertSet) Close() error { _ = "STUB: not implemented"; return nil }

func newInsertSet(ctx context.Context, orm *ormBase, mi *models.ModelInfo) (Inserter, error) {
	_ = "STUB: not implemented"
	return *new(Inserter), nil
}
