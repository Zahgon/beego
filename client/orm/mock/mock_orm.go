package mock

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/beego/beego/v2/client/orm"
)

func init() {
	RegisterMockDB("default")
}

func RegisterMockDB(name string) { _ = "STUB: not implemented"; return }

func MockTable(tableName string, resp ...interface{}) *Mock { _ = "STUB: not implemented"; return nil }

func MockMethod(method string, resp ...interface{}) *Mock { _ = "STUB: not implemented"; return nil }

func MockRead(tableName string, cb func(data interface{}), err error) *Mock {
	_ = "STUB: not implemented"
	return nil
}

func MockReadForUpdateWithCtx(tableName string, cb func(data interface{}), err error) *Mock {
	_ = "STUB: not implemented"
	return nil
}

func MockReadOrCreateWithCtx(tableName string,
	cb func(data interface{}),
	insert bool, id int64, err error) *Mock {
	_ = "STUB: not implemented"
	return nil
}

func MockInsertWithCtx(tableName string, id int64, err error) *Mock {
	_ = "STUB: not implemented"
	return nil
}

func MockInsertMultiWithCtx(tableName string, cnt int64, err error) *Mock {
	_ = "STUB: not implemented"
	return nil
}

func MockInsertOrUpdateWithCtx(tableName string, id int64, err error) *Mock {
	_ = "STUB: not implemented"
	return nil
}

func MockUpdateWithCtx(tableName string, affectedRow int64, err error) *Mock {
	_ = "STUB: not implemented"
	return nil
}

func MockDeleteWithCtx(tableName string, affectedRow int64, err error) *Mock {
	_ = "STUB: not implemented"
	return nil
}

func MockQueryM2MWithCtx(tableName string, name string, res orm.QueryM2Mer) *Mock {
	_ = "STUB: not implemented"
	return nil
}

func MockLoadRelatedWithCtx(tableName string, name string, rows int64, err error) *Mock {
	_ = "STUB: not implemented"
	return nil
}

func MockQueryTableWithCtx(tableName string, qs orm.QuerySeter) *Mock {
	_ = "STUB: not implemented"
	return nil
}

func MockRawWithCtx(rs orm.RawSeter) *Mock { _ = "STUB: not implemented"; return nil }

func MockDBStats(stats *sql.DBStats) *Mock { _ = "STUB: not implemented"; return nil }

func MockCommit(err error) *Mock { _ = "STUB: not implemented"; return nil }

func MockRollback(err error) *Mock { _ = "STUB: not implemented"; return nil }

func MockRollbackUnlessCommit(err error) *Mock { _ = "STUB: not implemented"; return nil }
