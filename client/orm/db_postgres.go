package orm

import (
	"context"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

var postgresOperators = map[string]string{
	"exact":       "= ?",
	"iexact":      "= UPPER(?)",
	"contains":    "LIKE ?",
	"icontains":   "LIKE UPPER(?)",
	"gt":          "> ?",
	"gte":         ">= ?",
	"lt":          "< ?",
	"lte":         "<= ?",
	"eq":          "= ?",
	"ne":          "!= ?",
	"startswith":  "LIKE ?",
	"endswith":    "LIKE ?",
	"istartswith": "LIKE UPPER(?)",
	"iendswith":   "LIKE UPPER(?)",
}

var postgresTypes = map[string]string{
	"auto":                "bigserial NOT NULL PRIMARY KEY",
	"pk":                  "NOT NULL PRIMARY KEY",
	"bool":                "bool",
	"string":              "varchar(%d)",
	"string-char":         "char(%d)",
	"string-text":         "text",
	"time.Time-date":      "date",
	"time.Time":           "timestamp with time zone",
	"int8":                `smallint CHECK("%COL%" >= -127 AND "%COL%" <= 128)`,
	"int16":               "smallint",
	"int32":               "integer",
	"int64":               "bigint",
	"uint8":               `smallint CHECK("%COL%" >= 0 AND "%COL%" <= 255)`,
	"uint16":              `integer CHECK("%COL%" >= 0)`,
	"uint32":              `bigint CHECK("%COL%" >= 0)`,
	"uint64":              `bigint CHECK("%COL%" >= 0)`,
	"float64":             "double precision",
	"float64-decimal":     "numeric(%d, %d)",
	"json":                "json",
	"jsonb":               "jsonb",
	"time.Time-precision": "timestamp(%d) with time zone",
}

type dbBasePostgres struct {
	dbBase
}

var _ dbBaser = new(dbBasePostgres)

func (d *dbBasePostgres) OperatorSQL(operator string) string { _ = "STUB: not implemented"; return "" }

func (d *dbBasePostgres) GenerateOperatorLeftCol(fi *models.FieldInfo, operator string, leftCol *string) {
	_ = "STUB: not implemented"
	return
}

func (d *dbBasePostgres) SupportUpdateJoin() bool { _ = "STUB: not implemented"; return false }

func (d *dbBasePostgres) MaxLimit() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *dbBasePostgres) TableQuote() string { _ = "STUB: not implemented"; return "" }

func (d *dbBasePostgres) ReplaceMarks(query *string) { _ = "STUB: not implemented"; return }

func (d *dbBasePostgres) HasReturningID(mi *models.ModelInfo, query *string) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *dbBasePostgres) setval(ctx context.Context, db dbQuerier, mi *models.ModelInfo, autoFields []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *dbBasePostgres) ShowTablesQuery() string { _ = "STUB: not implemented"; return "" }

func (d *dbBasePostgres) ShowColumnsQuery(table string) string {
	_ = "STUB: not implemented"
	return ""
}

func (d *dbBasePostgres) DbTypes() map[string]string { _ = "STUB: not implemented"; return nil }

func (d *dbBasePostgres) IndexExists(ctx context.Context, db dbQuerier, table string, name string) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *dbBasePostgres) GenerateSpecifyIndex(tableName string, useIndex int, indexes []string) string {
	_ = "STUB: not implemented"
	return ""
}

func newdbBasePostgres() dbBaser { _ = "STUB: not implemented"; return *new(dbBaser) }
