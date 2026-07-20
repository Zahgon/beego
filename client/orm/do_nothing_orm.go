package orm

import (
	"context"
	"database/sql"

	"github.com/beego/beego/v2/core/utils"
)

var _ Ormer = new(DoNothingOrm)

type DoNothingOrm struct{}

func (d *DoNothingOrm) Read(md interface{}, cols ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DoNothingOrm) ReadWithCtx(ctx context.Context, md interface{}, cols ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DoNothingOrm) ReadForUpdate(md interface{}, cols ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DoNothingOrm) ReadForUpdateWithCtx(ctx context.Context, md interface{}, cols ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DoNothingOrm) ReadOrCreate(md interface{}, col1 string, cols ...string) (bool, int64, error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

func (d *DoNothingOrm) ReadOrCreateWithCtx(ctx context.Context, md interface{}, col1 string, cols ...string) (bool, int64, error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

func (d *DoNothingOrm) LoadRelated(md interface{}, name string, args ...utils.KV) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingOrm) LoadRelatedWithCtx(ctx context.Context, md interface{}, name string, args ...utils.KV) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingOrm) QueryM2M(md interface{}, name string) QueryM2Mer {
	_ = "STUB: not implemented"
	return *new(QueryM2Mer)
}

func (d *DoNothingOrm) QueryM2MWithCtx(ctx context.Context, md interface{}, name string) QueryM2Mer {
	_ = "STUB: not implemented"
	return *new(QueryM2Mer)
}

func (d *DoNothingOrm) QueryTable(ptrStructOrTableName interface{}) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (d *DoNothingOrm) QueryTableWithCtx(ctx context.Context, ptrStructOrTableName interface{}) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (d *DoNothingOrm) DBStats() *sql.DBStats { _ = "STUB: not implemented"; return nil }

func (d *DoNothingOrm) Insert(md interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingOrm) InsertWithCtx(ctx context.Context, md interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingOrm) InsertOrUpdate(md interface{}, colConflitAndArgs ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingOrm) InsertOrUpdateWithCtx(ctx context.Context, md interface{}, colConflitAndArgs ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingOrm) InsertMulti(bulk int, mds interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingOrm) InsertMultiWithCtx(ctx context.Context, bulk int, mds interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingOrm) Update(md interface{}, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingOrm) UpdateWithCtx(ctx context.Context, md interface{}, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingOrm) Delete(md interface{}, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingOrm) DeleteWithCtx(ctx context.Context, md interface{}, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingOrm) Raw(query string, args ...interface{}) RawSeter {
	_ = "STUB: not implemented"
	return *new(RawSeter)
}

func (d *DoNothingOrm) RawWithCtx(ctx context.Context, query string, args ...interface{}) RawSeter {
	_ = "STUB: not implemented"
	return *new(RawSeter)
}

func (d *DoNothingOrm) Driver() Driver { _ = "STUB: not implemented"; return *new(Driver) }

func (d *DoNothingOrm) Begin() (TxOrmer, error) {
	_ = "STUB: not implemented"
	return *new(TxOrmer), nil
}

func (d *DoNothingOrm) BeginWithCtx(ctx context.Context) (TxOrmer, error) {
	_ = "STUB: not implemented"
	return *new(TxOrmer), nil
}

func (d *DoNothingOrm) BeginWithOpts(opts *sql.TxOptions) (TxOrmer, error) {
	_ = "STUB: not implemented"
	return *new(TxOrmer), nil
}

func (d *DoNothingOrm) BeginWithCtxAndOpts(ctx context.Context, opts *sql.TxOptions) (TxOrmer, error) {
	_ = "STUB: not implemented"
	return *new(TxOrmer), nil
}

func (d *DoNothingOrm) DoTx(task func(ctx context.Context, txOrm TxOrmer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DoNothingOrm) DoTxWithCtx(ctx context.Context, task func(ctx context.Context, txOrm TxOrmer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DoNothingOrm) DoTxWithOpts(opts *sql.TxOptions, task func(ctx context.Context, txOrm TxOrmer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DoNothingOrm) DoTxWithCtxAndOpts(ctx context.Context, opts *sql.TxOptions, task func(ctx context.Context, txOrm TxOrmer) error) error {
	_ = "STUB: not implemented"
	return nil
}

type DoNothingTxOrm struct {
	DoNothingOrm
}

func (d *DoNothingTxOrm) Commit() error { _ = "STUB: not implemented"; return nil }

func (d *DoNothingTxOrm) Rollback() error { _ = "STUB: not implemented"; return nil }
