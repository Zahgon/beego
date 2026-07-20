package orm

import (
	"context"
)

type dbBaseTidb struct {
	dbBase
}

var _ dbBaser = new(dbBaseTidb)

func (d *dbBaseTidb) OperatorSQL(operator string) string { _ = "STUB: not implemented"; return "" }

func (d *dbBaseTidb) DbTypes() map[string]string { _ = "STUB: not implemented"; return nil }

func (d *dbBaseTidb) ShowTablesQuery() string { _ = "STUB: not implemented"; return "" }

func (d *dbBaseTidb) ShowColumnsQuery(table string) string { _ = "STUB: not implemented"; return "" }

func (d *dbBaseTidb) IndexExists(ctx context.Context, db dbQuerier, table string, name string) bool {
	_ = "STUB: not implemented"
	return false
}

func newdbBaseTidb() dbBaser { _ = "STUB: not implemented"; return *new(dbBaser) }
