package orm

import (
	"context"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

var oracleOperators = map[string]string{
	"exact":       "= ?",
	"gt":          "> ?",
	"gte":         ">= ?",
	"lt":          "< ?",
	"lte":         "<= ?",
	"//iendswith": "LIKE ?",
}

var oracleTypes = map[string]string{
	"pk":                  "NOT NULL PRIMARY KEY",
	"bool":                "bool",
	"string":              "VARCHAR2(%d)",
	"string-char":         "CHAR(%d)",
	"string-text":         "VARCHAR2(%d)",
	"time.Time-date":      "DATE",
	"time.Time":           "TIMESTAMP",
	"int8":                "INTEGER",
	"int16":               "INTEGER",
	"int32":               "INTEGER",
	"int64":               "INTEGER",
	"uint8":               "INTEGER",
	"uint16":              "INTEGER",
	"uint32":              "INTEGER",
	"uint64":              "INTEGER",
	"float64":             "NUMBER",
	"float64-decimal":     "NUMBER(%d, %d)",
	"time.Time-precision": "TIMESTAMP(%d)",
}

type dbBaseOracle struct {
	dbBase
}

var _ dbBaser = new(dbBaseOracle)

func newdbBaseOracle() dbBaser { _ = "STUB: not implemented"; return *new(dbBaser) }

func (d *dbBaseOracle) OperatorSQL(operator string) string { _ = "STUB: not implemented"; return "" }

func (d *dbBaseOracle) DbTypes() map[string]string { _ = "STUB: not implemented"; return nil }

func (d *dbBaseOracle) ShowTablesQuery() string { _ = "STUB: not implemented"; return "" }

func (d *dbBaseOracle) ShowColumnsQuery(table string) string { _ = "STUB: not implemented"; return "" }

func (d *dbBaseOracle) IndexExists(ctx context.Context, db dbQuerier, table string, name string) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *dbBaseOracle) GenerateSpecifyIndex(tableName string, useIndex int, indexes []string) string {
	_ = "STUB: not implemented"
	return ""
}

func (d *dbBaseOracle) InsertValue(ctx context.Context, q dbQuerier, mi *models.ModelInfo, isMulti bool, names []string, values []interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
