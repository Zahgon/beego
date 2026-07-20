package orm

import (
	"context"
	"database/sql"
	"reflect"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"

	"github.com/beego/beego/v2/client/orm/clauses/order_clause"
	"github.com/beego/beego/v2/core/utils"
)

type TableNameI interface {
	TableName() string
}

type TableEngineI interface {
	TableEngine() string
}

type TableIndexI interface {
	TableIndex() [][]string
}

type TableUniqueI interface {
	TableUnique() [][]string
}

type IsApplicableTableForDB interface {
	IsApplicableTableForDB(db string) bool
}

type Driver interface {
	Name() string
	Type() DriverType
}

type Fielder = models.Fielder

type TxBeginner interface {
	Begin() (TxOrmer, error)
	BeginWithCtx(ctx context.Context) (TxOrmer, error)
	BeginWithOpts(opts *sql.TxOptions) (TxOrmer, error)
	BeginWithCtxAndOpts(ctx context.Context, opts *sql.TxOptions) (TxOrmer, error)

	DoTx(task func(ctx context.Context, txOrm TxOrmer) error) error
	DoTxWithCtx(ctx context.Context, task func(ctx context.Context, txOrm TxOrmer) error) error
	DoTxWithOpts(opts *sql.TxOptions, task func(ctx context.Context, txOrm TxOrmer) error) error
	DoTxWithCtxAndOpts(ctx context.Context, opts *sql.TxOptions, task func(ctx context.Context, txOrm TxOrmer) error) error
}

type TxCommitter interface {
	txEnder
}

type txer interface {
	Begin() (*sql.Tx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type txEnder interface {
	Commit() error
	Rollback() error

	RollbackUnlessCommit() error
}

type DML interface {
	Insert(md interface{}) (int64, error)
	InsertWithCtx(ctx context.Context, md interface{}) (int64, error)

	InsertOrUpdate(md interface{}, colConflitAndArgs ...string) (int64, error)
	InsertOrUpdateWithCtx(ctx context.Context, md interface{}, colConflitAndArgs ...string) (int64, error)

	InsertMulti(bulk int, mds interface{}) (int64, error)
	InsertMultiWithCtx(ctx context.Context, bulk int, mds interface{}) (int64, error)

	Update(md interface{}, cols ...string) (int64, error)
	UpdateWithCtx(ctx context.Context, md interface{}, cols ...string) (int64, error)

	Delete(md interface{}, cols ...string) (int64, error)
	DeleteWithCtx(ctx context.Context, md interface{}, cols ...string) (int64, error)

	Raw(query string, args ...interface{}) RawSeter
	RawWithCtx(ctx context.Context, query string, args ...interface{}) RawSeter
}

type DQL interface {
	Read(md interface{}, cols ...string) error
	ReadWithCtx(ctx context.Context, md interface{}, cols ...string) error

	ReadForUpdate(md interface{}, cols ...string) error
	ReadForUpdateWithCtx(ctx context.Context, md interface{}, cols ...string) error

	ReadOrCreate(md interface{}, col1 string, cols ...string) (bool, int64, error)
	ReadOrCreateWithCtx(ctx context.Context, md interface{}, col1 string, cols ...string) (bool, int64, error)

	LoadRelated(md interface{}, name string, args ...utils.KV) (int64, error)
	LoadRelatedWithCtx(ctx context.Context, md interface{}, name string, args ...utils.KV) (int64, error)

	QueryM2M(md interface{}, name string) QueryM2Mer

	QueryM2MWithCtx(ctx context.Context, md interface{}, name string) QueryM2Mer

	QueryTable(ptrStructOrTableName interface{}) QuerySeter

	QueryTableWithCtx(ctx context.Context, ptrStructOrTableName interface{}) QuerySeter

	DBStats() *sql.DBStats
}

type DriverGetter interface {
	Driver() Driver
}

type ormer interface {
	DQL
	DML
	DriverGetter
}

type QueryExecutor interface {
	ormer
}

type Ormer interface {
	QueryExecutor
	TxBeginner
}

type TxOrmer interface {
	QueryExecutor
	TxCommitter
}

type Inserter interface {
	Insert(interface{}) (int64, error)
	InsertWithCtx(context.Context, interface{}) (int64, error)
	Close() error
}

type QuerySeter interface {
	Filter(string, ...interface{}) QuerySeter

	FilterRaw(string, string) QuerySeter

	Exclude(string, ...interface{}) QuerySeter

	SetCond(*Condition) QuerySeter

	GetCond() *Condition

	Limit(limit interface{}, args ...interface{}) QuerySeter

	Offset(offset interface{}) QuerySeter

	GroupBy(exprs ...string) QuerySeter

	OrderBy(exprs ...string) QuerySeter

	OrderClauses(orders ...*order_clause.Order) QuerySeter

	ForceIndex(indexes ...string) QuerySeter

	UseIndex(indexes ...string) QuerySeter

	IgnoreIndex(indexes ...string) QuerySeter

	RelatedSel(params ...interface{}) QuerySeter

	Distinct() QuerySeter

	ForUpdate() QuerySeter

	Count() (int64, error)
	CountWithCtx(context.Context) (int64, error)

	Exist() bool
	ExistWithCtx(context.Context) bool

	Update(values Params) (int64, error)
	UpdateWithCtx(ctx context.Context, values Params) (int64, error)

	Delete() (int64, error)
	DeleteWithCtx(context.Context) (int64, error)

	PrepareInsert() (Inserter, error)
	PrepareInsertWithCtx(context.Context) (Inserter, error)

	All(container interface{}, cols ...string) (int64, error)
	AllWithCtx(ctx context.Context, container interface{}, cols ...string) (int64, error)

	One(container interface{}, cols ...string) error
	OneWithCtx(ctx context.Context, container interface{}, cols ...string) error

	Values(results *[]Params, exprs ...string) (int64, error)
	ValuesWithCtx(ctx context.Context, results *[]Params, exprs ...string) (int64, error)

	ValuesList(results *[]ParamsList, exprs ...string) (int64, error)
	ValuesListWithCtx(ctx context.Context, results *[]ParamsList, exprs ...string) (int64, error)

	ValuesFlat(result *ParamsList, expr string) (int64, error)
	ValuesFlatWithCtx(ctx context.Context, result *ParamsList, expr string) (int64, error)

	RowsToMap(result *Params, keyCol, valueCol string) (int64, error)

	RowsToStruct(ptrStruct interface{}, keyCol, valueCol string) (int64, error)

	Aggregate(s string) QuerySeter
}

type QueryM2Mer interface {
	Add(...interface{}) (int64, error)
	AddWithCtx(context.Context, ...interface{}) (int64, error)

	Remove(...interface{}) (int64, error)
	RemoveWithCtx(context.Context, ...interface{}) (int64, error)

	Exist(interface{}) bool
	ExistWithCtx(context.Context, interface{}) bool

	Clear() (int64, error)
	ClearWithCtx(context.Context) (int64, error)

	Count() (int64, error)
	CountWithCtx(context.Context) (int64, error)
}

type RawPreparer interface {
	Exec(...interface{}) (sql.Result, error)
	Close() error
}

type RawSeter interface {
	Exec() (sql.Result, error)

	QueryRow(containers ...interface{}) error

	QueryRows(containers ...interface{}) (int64, error)
	SetArgs(...interface{}) RawSeter

	Values(container *[]Params, cols ...string) (int64, error)

	ValuesList(container *[]ParamsList, cols ...string) (int64, error)

	ValuesFlat(container *ParamsList, cols ...string) (int64, error)

	RowsToMap(result *Params, keyCol, valueCol string) (int64, error)

	RowsToStruct(ptrStruct interface{}, keyCol, valueCol string) (int64, error)

	Prepare() (RawPreparer, error)
}

type stmtQuerier interface {
	Close() error
	Exec(args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error)
	Query(args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, args ...interface{}) (*sql.Rows, error)
	QueryRow(args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row
}

type dbQuerier interface {
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type dbBaser interface {
	Read(context.Context, dbQuerier, *models.ModelInfo, reflect.Value, *time.Location, []string, bool) error
	ReadBatch(context.Context, dbQuerier, querySet, *models.ModelInfo, *Condition, interface{}, *time.Location, []string) (int64, error)
	Count(context.Context, dbQuerier, querySet, *models.ModelInfo, *Condition, *time.Location) (int64, error)
	ReadValues(context.Context, dbQuerier, querySet, *models.ModelInfo, *Condition, []string, interface{}, *time.Location) (int64, error)

	Insert(context.Context, dbQuerier, *models.ModelInfo, reflect.Value, *time.Location) (int64, error)
	InsertOrUpdate(context.Context, dbQuerier, *models.ModelInfo, reflect.Value, *alias, ...string) (int64, error)
	InsertMulti(context.Context, dbQuerier, *models.ModelInfo, reflect.Value, int, *time.Location) (int64, error)
	InsertValue(context.Context, dbQuerier, *models.ModelInfo, bool, []string, []interface{}) (int64, error)
	InsertStmt(context.Context, stmtQuerier, *models.ModelInfo, reflect.Value, *time.Location) (int64, error)

	Update(context.Context, dbQuerier, *models.ModelInfo, reflect.Value, *time.Location, []string) (int64, error)
	UpdateBatch(context.Context, dbQuerier, *querySet, *models.ModelInfo, *Condition, Params, *time.Location) (int64, error)

	Delete(context.Context, dbQuerier, *models.ModelInfo, reflect.Value, *time.Location, []string) (int64, error)
	DeleteBatch(context.Context, dbQuerier, *querySet, *models.ModelInfo, *Condition, *time.Location) (int64, error)

	SupportUpdateJoin() bool
	OperatorSQL(string) string
	GenerateOperatorSQL(*models.ModelInfo, *models.FieldInfo, string, []interface{}, *time.Location) (string, []interface{})
	GenerateOperatorLeftCol(*models.FieldInfo, string, *string)
	PrepareInsert(context.Context, dbQuerier, *models.ModelInfo) (stmtQuerier, string, error)
	MaxLimit() uint64
	TableQuote() string
	ReplaceMarks(*string)
	HasReturningID(*models.ModelInfo, *string) bool
	TimeFromDB(*time.Time, *time.Location)
	TimeToDB(*time.Time, *time.Location)
	DbTypes() map[string]string
	GetTables(dbQuerier) (map[string]bool, error)
	GetColumns(context.Context, dbQuerier, string) (map[string][3]string, error)
	ShowTablesQuery() string
	ShowColumnsQuery(string) string
	IndexExists(context.Context, dbQuerier, string, string) bool
	collectFieldValue(*models.ModelInfo, *models.FieldInfo, reflect.Value, bool, *time.Location) (interface{}, error)
	setval(context.Context, dbQuerier, *models.ModelInfo, []string) error

	GenerateSpecifyIndex(tableName string, useIndex int, indexes []string) string
}
