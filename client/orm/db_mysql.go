package orm

import (
	"context"
	"reflect"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

var mysqlOperators = map[string]string{
	"exact":       "= ?",
	"iexact":      "LIKE ?",
	"strictexact": "= BINARY ?",
	"contains":    "LIKE BINARY ?",
	"icontains":   "LIKE ?",

	"gt":          "> ?",
	"gte":         ">= ?",
	"lt":          "< ?",
	"lte":         "<= ?",
	"eq":          "= ?",
	"ne":          "!= ?",
	"startswith":  "LIKE BINARY ?",
	"endswith":    "LIKE BINARY ?",
	"istartswith": "LIKE ?",
	"iendswith":   "LIKE ?",
}

var mysqlTypes = map[string]string{
	"auto":                "AUTO_INCREMENT NOT NULL PRIMARY KEY",
	"pk":                  "NOT NULL PRIMARY KEY",
	"bool":                "bool",
	"string":              "varchar(%d)",
	"string-char":         "char(%d)",
	"string-text":         "longtext",
	"time.Time-date":      "date",
	"time.Time":           "datetime",
	"int8":                "tinyint",
	"int16":               "smallint",
	"int32":               "integer",
	"int64":               "bigint",
	"uint8":               "tinyint unsigned",
	"uint16":              "smallint unsigned",
	"uint32":              "integer unsigned",
	"uint64":              "bigint unsigned",
	"float64":             "double precision",
	"float64-decimal":     "numeric(%d, %d)",
	"time.Time-precision": "datetime(%d)",
}

type dbBaseMysql struct {
	dbBase
}

var _ dbBaser = new(dbBaseMysql)

func (d *dbBaseMysql) OperatorSQL(operator string) string { _ = "STUB: not implemented"; return "" }

func (d *dbBaseMysql) DbTypes() map[string]string { _ = "STUB: not implemented"; return nil }

func (d *dbBaseMysql) ShowTablesQuery() string { _ = "STUB: not implemented"; return "" }

func (d *dbBaseMysql) ShowColumnsQuery(table string) string { _ = "STUB: not implemented"; return "" }

func (d *dbBaseMysql) IndexExists(ctx context.Context, db dbQuerier, table string, name string) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *dbBaseMysql) InsertOrUpdate(ctx context.Context, q dbQuerier, mi *models.ModelInfo, ind reflect.Value, a *alias, args ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func newdbBaseMysql() dbBaser { _ = "STUB: not implemented"; return *new(dbBaser) }
