package orm

import (
	"context"
	"reflect"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

var sqliteOperators = map[string]string{
	"exact":       "= ?",
	"iexact":      "LIKE ? ESCAPE '\\'",
	"contains":    "LIKE ? ESCAPE '\\'",
	"icontains":   "LIKE ? ESCAPE '\\'",
	"gt":          "> ?",
	"gte":         ">= ?",
	"lt":          "< ?",
	"lte":         "<= ?",
	"eq":          "= ?",
	"ne":          "!= ?",
	"startswith":  "LIKE ? ESCAPE '\\'",
	"endswith":    "LIKE ? ESCAPE '\\'",
	"istartswith": "LIKE ? ESCAPE '\\'",
	"iendswith":   "LIKE ? ESCAPE '\\'",
}

var sqliteTypes = map[string]string{
	"auto":                "integer NOT NULL PRIMARY KEY AUTOINCREMENT",
	"pk":                  "NOT NULL PRIMARY KEY",
	"bool":                "bool",
	"string":              "varchar(%d)",
	"string-char":         "character(%d)",
	"string-text":         "text",
	"time.Time-date":      "date",
	"time.Time":           "datetime",
	"time.Time-precision": "datetime(%d)",
	"int8":                "tinyint",
	"int16":               "smallint",
	"int32":               "integer",
	"int64":               "bigint",
	"uint8":               "tinyint unsigned",
	"uint16":              "smallint unsigned",
	"uint32":              "integer unsigned",
	"uint64":              "bigint unsigned",
	"float64":             "real",
	"float64-decimal":     "decimal",
}

type dbBaseSqlite struct {
	dbBase
}

var _ dbBaser = new(dbBaseSqlite)

func (d *dbBaseSqlite) Read(ctx context.Context, q dbQuerier, mi *models.ModelInfo, ind reflect.Value, tz *time.Location, cols []string, isForUpdate bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *dbBaseSqlite) OperatorSQL(operator string) string { _ = "STUB: not implemented"; return "" }

func (d *dbBaseSqlite) GenerateOperatorLeftCol(fi *models.FieldInfo, operator string, leftCol *string) {
	_ = "STUB: not implemented"
	return
}

func (d *dbBaseSqlite) SupportUpdateJoin() bool { _ = "STUB: not implemented"; return false }

func (d *dbBaseSqlite) MaxLimit() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *dbBaseSqlite) DbTypes() map[string]string { _ = "STUB: not implemented"; return nil }

func (d *dbBaseSqlite) ShowTablesQuery() string { _ = "STUB: not implemented"; return "" }

func (d *dbBaseSqlite) GetColumns(ctx context.Context, db dbQuerier, table string) (map[string][3]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *dbBaseSqlite) ShowColumnsQuery(table string) string { _ = "STUB: not implemented"; return "" }

func (d *dbBaseSqlite) IndexExists(ctx context.Context, db dbQuerier, table string, name string) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *dbBaseSqlite) GenerateSpecifyIndex(tableName string, useIndex int, indexes []string) string {
	_ = "STUB: not implemented"
	return ""
}

func newdbBaseSqlite() dbBaser { _ = "STUB: not implemented"; return *new(dbBaser) }
