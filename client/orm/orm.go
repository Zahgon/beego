package orm

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"reflect"

	iutils "github.com/beego/beego/v2/client/orm/internal/utils"

	"github.com/beego/beego/v2/client/orm/internal/models"

	"github.com/beego/beego/v2/core/utils"
)

const (
	DebugQueries = iota
)

var (
	Debug            = false
	DebugLog         = NewLog(os.Stdout)
	DefaultRowsLimit = -1
	DefaultRelsDepth = 2
	DefaultTimeLoc   = iutils.DefaultTimeLoc
	ErrTxDone        = errors.New("<TxOrmer.Commit/Rollback> transaction already done")
	ErrMultiRows     = errors.New("<QuerySeter> return multi rows")
	ErrNoRows        = errors.New("<QuerySeter> no row found")
	ErrStmtClosed    = errors.New("<QuerySeter> stmt already closed")
	ErrArgs          = errors.New("<Ormer> args error may be empty")
	ErrNotImplement  = errors.New("have not implement")

	ErrLastInsertIdUnavailable = errors.New("<Ormer> last insert id is unavailable")
)

type Params map[string]interface{}

type ParamsList []interface{}

type ormBase struct {
	alias *alias
	db    dbQuerier
}

var (
	_ DQL          = new(ormBase)
	_ DML          = new(ormBase)
	_ DriverGetter = new(ormBase)
)

func (*ormBase) getMi(md interface{}) (mi *models.ModelInfo) { _ = "STUB: not implemented"; return nil }

func (*ormBase) getPtrMiInd(md interface{}) (mi *models.ModelInfo, ind reflect.Value) {
	_ = "STUB: not implemented"
	return nil, *new(reflect.Value)
}

func getTypeMi(mdTyp reflect.Type) *models.ModelInfo { _ = "STUB: not implemented"; return nil }

func (*ormBase) getFieldInfo(mi *models.ModelInfo, name string) *models.FieldInfo {
	_ = "STUB: not implemented"
	return nil
}

func (o *ormBase) Read(md interface{}, cols ...string) error { _ = "STUB: not implemented"; return nil }

func (o *ormBase) ReadWithCtx(ctx context.Context, md interface{}, cols ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *ormBase) ReadForUpdate(md interface{}, cols ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *ormBase) ReadForUpdateWithCtx(ctx context.Context, md interface{}, cols ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *ormBase) ReadOrCreate(md interface{}, col1 string, cols ...string) (bool, int64, error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

func (o *ormBase) ReadOrCreateWithCtx(ctx context.Context, md interface{}, col1 string, cols ...string) (bool, int64, error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

func (o *ormBase) Insert(md interface{}) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (o *ormBase) InsertWithCtx(ctx context.Context, md interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (*ormBase) setPk(mi *models.ModelInfo, ind reflect.Value, id int64) {
	_ = "STUB: not implemented"
	return
}

func (o *ormBase) InsertMulti(bulk int, mds interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *ormBase) InsertMultiWithCtx(ctx context.Context, bulk int, mds interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *ormBase) InsertOrUpdate(md interface{}, colConflictAndArgs ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *ormBase) InsertOrUpdateWithCtx(ctx context.Context, md interface{}, colConflitAndArgs ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *ormBase) Update(md interface{}, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *ormBase) UpdateWithCtx(ctx context.Context, md interface{}, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *ormBase) Delete(md interface{}, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *ormBase) DeleteWithCtx(ctx context.Context, md interface{}, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *ormBase) QueryM2M(md interface{}, name string) QueryM2Mer {
	_ = "STUB: not implemented"
	return *new(QueryM2Mer)
}

func (o *ormBase) QueryM2MWithCtx(_ context.Context, md interface{}, name string) QueryM2Mer {
	_ = "STUB: not implemented"
	return *new(QueryM2Mer)
}

func (o *ormBase) LoadRelated(md interface{}, name string, args ...utils.KV) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *ormBase) LoadRelatedWithCtx(_ context.Context, md interface{}, name string, args ...utils.KV) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *ormBase) queryRelated(md interface{}, name string) (*models.ModelInfo, *models.FieldInfo, reflect.Value, *querySet) {
	_ = "STUB: not implemented"
	return nil, nil, *new(reflect.Value), nil
}

func (o *ormBase) getReverseQs(md interface{}, mi *models.ModelInfo, fi *models.FieldInfo) *querySet {
	_ = "STUB: not implemented"
	return nil
}

func (o *ormBase) getRelQs(md interface{}, mi *models.ModelInfo, fi *models.FieldInfo) *querySet {
	_ = "STUB: not implemented"
	return nil
}

func (o *ormBase) QueryTable(ptrStructOrTableName interface{}) (qs QuerySeter) {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (o *ormBase) QueryTableWithCtx(_ context.Context, ptrStructOrTableName interface{}) (qs QuerySeter) {
	_ = "STUB: not implemented"
	return *new(QuerySeter)
}

func (o *ormBase) Raw(query string, args ...interface{}) RawSeter {
	_ = "STUB: not implemented"
	return *new(RawSeter)
}

func (o *ormBase) RawWithCtx(_ context.Context, query string, args ...interface{}) RawSeter {
	_ = "STUB: not implemented"
	return *new(RawSeter)
}

func (o *ormBase) Driver() Driver { _ = "STUB: not implemented"; return *new(Driver) }

func (o *ormBase) DBStats() *sql.DBStats { _ = "STUB: not implemented"; return nil }

type orm struct {
	ormBase
}

var _ Ormer = new(orm)

func (o *orm) Begin() (TxOrmer, error) { _ = "STUB: not implemented"; return *new(TxOrmer), nil }

func (o *orm) BeginWithCtx(ctx context.Context) (TxOrmer, error) {
	_ = "STUB: not implemented"
	return *new(TxOrmer), nil
}

func (o *orm) BeginWithOpts(opts *sql.TxOptions) (TxOrmer, error) {
	_ = "STUB: not implemented"
	return *new(TxOrmer), nil
}

func (o *orm) BeginWithCtxAndOpts(ctx context.Context, opts *sql.TxOptions) (TxOrmer, error) {
	_ = "STUB: not implemented"
	return *new(TxOrmer), nil
}

func (o *orm) DoTx(task func(ctx context.Context, txOrm TxOrmer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *orm) DoTxWithCtx(ctx context.Context, task func(ctx context.Context, txOrm TxOrmer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *orm) DoTxWithOpts(opts *sql.TxOptions, task func(ctx context.Context, txOrm TxOrmer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *orm) DoTxWithCtxAndOpts(ctx context.Context, opts *sql.TxOptions, task func(ctx context.Context, txOrm TxOrmer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func doTxTemplate(ctx context.Context, o TxBeginner, opts *sql.TxOptions,
	task func(ctx context.Context, txOrm TxOrmer) error) error {
	_ = "STUB: not implemented"
	return nil
}

type txOrm struct {
	ormBase
}

var _ TxOrmer = new(txOrm)

func (t *txOrm) Commit() error { _ = "STUB: not implemented"; return nil }

func (t *txOrm) Rollback() error { _ = "STUB: not implemented"; return nil }

func (t *txOrm) RollbackUnlessCommit() error { _ = "STUB: not implemented"; return nil }

func NewOrm() Ormer { _ = "STUB: not implemented"; return *new(Ormer) }

func NewOrmUsingDB(aliasName string) Ormer { _ = "STUB: not implemented"; return *new(Ormer) }

func NewOrmWithDB(driverName, aliasName string, db *sql.DB, params ...DBOption) (Ormer, error) {
	_ = "STUB: not implemented"
	return *new(Ormer), nil
}

func newDBWithAlias(al *alias) Ormer { _ = "STUB: not implemented"; return *new(Ormer) }
