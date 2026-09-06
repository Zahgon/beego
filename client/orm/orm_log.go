package orm

import (
	"context"
	"database/sql"
	"io"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/logs"
)

type Log = logs.Log

func NewLog(out io.Writer) *logs.Log { _ = "STUB: not implemented"; return nil }

var LogFunc func(query map[string]interface{})

func debugLogQueies(alias *alias, operation, query string, t time.Time, err error, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

type stmtQueryLog struct {
	alias *alias
	query string
	stmt  stmtQuerier
}

var _ stmtQuerier = new(stmtQueryLog)

func (d *stmtQueryLog) Close() error { _ = "STUB: not implemented"; return nil }

func (d *stmtQueryLog) Exec(args ...interface{}) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (d *stmtQueryLog) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (d *stmtQueryLog) Query(args ...interface{}) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *stmtQueryLog) QueryContext(ctx context.Context, args ...interface{}) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *stmtQueryLog) QueryRow(args ...interface{}) *sql.Row {
	_ = "STUB: not implemented"
	return nil
}

func (d *stmtQueryLog) QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row {
	_ = "STUB: not implemented"
	return nil
}

func newStmtQueryLog(alias *alias, stmt stmtQuerier, query string) stmtQuerier {
	_ = "STUB: not implemented"
	return *new(stmtQuerier)
}

type dbQueryLog struct {
	alias *alias
	db    dbQuerier
	tx    txer
	txe   txEnder
}

var (
	_ dbQuerier = new(dbQueryLog)
	_ txer      = new(dbQueryLog)
	_ txEnder   = new(dbQueryLog)
)

func (d *dbQueryLog) Prepare(query string) (*sql.Stmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *dbQueryLog) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *dbQueryLog) Exec(query string, args ...interface{}) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (d *dbQueryLog) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (d *dbQueryLog) Query(query string, args ...interface{}) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *dbQueryLog) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *dbQueryLog) QueryRow(query string, args ...interface{}) *sql.Row {
	_ = "STUB: not implemented"
	return nil
}

func (d *dbQueryLog) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	_ = "STUB: not implemented"
	return nil
}

func (d *dbQueryLog) Begin() (*sql.Tx, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *dbQueryLog) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *dbQueryLog) Commit() error { _ = "STUB: not implemented"; return nil }

func (d *dbQueryLog) Rollback() error { _ = "STUB: not implemented"; return nil }

func (d *dbQueryLog) RollbackUnlessCommit() error { _ = "STUB: not implemented"; return nil }

func (d *dbQueryLog) SetDB(db dbQuerier) { _ = "STUB: not implemented"; return }

func newDbQueryLog(alias *alias, db dbQuerier) dbQuerier {
	_ = "STUB: not implemented"
	return *new(dbQuerier)
}
