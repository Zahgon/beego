package orm

import (
	"context"
	"database/sql"
	"time"

	"github.com/beego/beego/v2/core/utils"
)

const (
	TxNameKey = "TxName"
)

var (
	_ Ormer   = new(filterOrmDecorator)
	_ TxOrmer = new(filterOrmDecorator)
)

type filterOrmDecorator struct {
	ormer
	TxBeginner
	TxCommitter

	root Filter

	insideTx    bool
	txStartTime time.Time
	txName      string
}

func NewFilterOrmDecorator(delegate Ormer, filterChains ...FilterChain) Ormer {
	_ = "STUB: not implemented"
	return *new(Ormer)
}

func NewFilterTxOrmDecorator(delegate TxOrmer, root Filter, txName string) TxOrmer {
	_ = "STUB: not implemented"
	return *new(TxOrmer)
}

func (f *filterOrmDecorator) Read(md interface{}, cols ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *filterOrmDecorator) ReadWithCtx(ctx context.Context, md interface{}, cols ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *filterOrmDecorator) ReadForUpdate(md interface{}, cols ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *filterOrmDecorator) ReadForUpdateWithCtx(ctx context.Context, md interface{}, cols ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *filterOrmDecorator) ReadOrCreate(md interface{}, col1 string, cols ...string) (bool, int64, error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

func (f *filterOrmDecorator) ReadOrCreateWithCtx(ctx context.Context, md interface{}, col1 string, cols ...string) (bool, int64, error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

func (f *filterOrmDecorator) LoadRelated(md interface{}, name string, args ...utils.KV) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *filterOrmDecorator) LoadRelatedWithCtx(ctx context.Context, md interface{}, name string, args ...utils.KV) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *filterOrmDecorator) QueryM2M(md interface{}, name string) QueryM2Mer {
	_ = "STUB: not implemented"
	return *new(QueryM2Mer)
}

func (f *filterOrmDecorator) QueryM2MWithCtx(_ context.Context, md interface{}, name string) QueryM2Mer {
	_ = "STUB: not implemented"
	return *new(QueryM2Mer)
}

func (f *filterOrmDecorator) QueryTable(ptrStructOrTableName interface{}) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (f *filterOrmDecorator) QueryTableWithCtx(_ context.Context, ptrStructOrTableName interface{}) QuerySeter {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (f *filterOrmDecorator) DBStats() *sql.DBStats { _ = "STUB: not implemented"; return nil }

func (f *filterOrmDecorator) Insert(md interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *filterOrmDecorator) InsertWithCtx(ctx context.Context, md interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *filterOrmDecorator) InsertOrUpdate(md interface{}, colConflitAndArgs ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *filterOrmDecorator) InsertOrUpdateWithCtx(ctx context.Context, md interface{}, colConflitAndArgs ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *filterOrmDecorator) InsertMulti(bulk int, mds interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *filterOrmDecorator) InsertMultiWithCtx(ctx context.Context, bulk int, mds interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *filterOrmDecorator) Update(md interface{}, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *filterOrmDecorator) UpdateWithCtx(ctx context.Context, md interface{}, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *filterOrmDecorator) Delete(md interface{}, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *filterOrmDecorator) DeleteWithCtx(ctx context.Context, md interface{}, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *filterOrmDecorator) Raw(query string, args ...interface{}) RawSeter {
	_ = "STUB: not implemented"
	return *new(RawSeter)
}

func (f *filterOrmDecorator) RawWithCtx(ctx context.Context, query string, args ...interface{}) RawSeter {
	_ = "STUB: not implemented"
	return *new(RawSeter)
}

func (f *filterOrmDecorator) Driver() Driver { _ = "STUB: not implemented"; return *new(Driver) }

func (f *filterOrmDecorator) Begin() (TxOrmer, error) {
	_ = "STUB: not implemented"
	return *new(TxOrmer), nil
}

func (f *filterOrmDecorator) BeginWithCtx(ctx context.Context) (TxOrmer, error) {
	_ = "STUB: not implemented"
	return *new(TxOrmer), nil
}

func (f *filterOrmDecorator) BeginWithOpts(opts *sql.TxOptions) (TxOrmer, error) {
	_ = "STUB: not implemented"
	return *new(TxOrmer), nil
}

func (f *filterOrmDecorator) BeginWithCtxAndOpts(ctx context.Context, opts *sql.TxOptions) (TxOrmer, error) {
	_ = "STUB: not implemented"
	return *new(TxOrmer), nil
}

func (f *filterOrmDecorator) DoTx(task func(ctx context.Context, txOrm TxOrmer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *filterOrmDecorator) DoTxWithCtx(ctx context.Context, task func(ctx context.Context, txOrm TxOrmer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *filterOrmDecorator) DoTxWithOpts(opts *sql.TxOptions, task func(ctx context.Context, txOrm TxOrmer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *filterOrmDecorator) DoTxWithCtxAndOpts(ctx context.Context, opts *sql.TxOptions, task func(ctx context.Context, txOrm TxOrmer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *filterOrmDecorator) Commit() error { _ = "STUB: not implemented"; return nil }

func (f *filterOrmDecorator) Rollback() error { _ = "STUB: not implemented"; return nil }

func (f *filterOrmDecorator) RollbackUnlessCommit() error { _ = "STUB: not implemented"; return nil }

func (*filterOrmDecorator) convertError(v interface{}) error { _ = "STUB: not implemented"; return nil }

func getTxNameFromCtx(ctx context.Context) string { _ = "STUB: not implemented"; return "" }
