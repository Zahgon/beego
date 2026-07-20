package orm

import (
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"

	"github.com/beego/beego/v2/client/orm/clauses/order_clause"
)

type dbTable struct {
	id    int
	index string
	name  string
	names []string
	sel   bool
	inner bool
	mi    *models.ModelInfo
	fi    *models.FieldInfo
	jtl   *dbTable
}

type dbTables struct {
	tablesM map[string]*dbTable
	tables  []*dbTable
	mi      *models.ModelInfo
	base    dbBaser
	skipEnd bool
}

func (t *dbTables) set(names []string, mi *models.ModelInfo, fi *models.FieldInfo, inner bool) *dbTable {
	_ = "STUB: not implemented"
	return nil
}

func (t *dbTables) add(names []string, mi *models.ModelInfo, fi *models.FieldInfo, inner bool) (*dbTable, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (t *dbTables) get(name string) (*dbTable, bool) { _ = "STUB: not implemented"; return nil, false }

func (t *dbTables) loopDepth(depth int, prefix string, fi *models.FieldInfo, related []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (t *dbTables) parseRelated(rels []string, depth int) { _ = "STUB: not implemented"; return }

func (t *dbTables) getJoinSQL() (join string) { _ = "STUB: not implemented"; return "" }

func (t *dbTables) parseExprs(mi *models.ModelInfo, exprs []string) (index, name string, info *models.FieldInfo, success bool) {
	_ = "STUB: not implemented"
	return "", "", nil, false
}

func (t *dbTables) getCondSQL(cond *Condition, sub bool, tz *time.Location) (where string, params []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

func (t *dbTables) getGroupSQL(groups []string) (groupSQL string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *dbTables) getOrderSQL(orders []*order_clause.Order) (orderSQL string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *dbTables) getLimitSQL(mi *models.ModelInfo, offset int64, limit int64) (limits string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *dbTables) getIndexSql(tableName string, useIndex int, indexes []string) (clause string) {
	_ = "STUB: not implemented"
	return ""
}

func newDbTables(mi *models.ModelInfo, base dbBaser) *dbTables {
	_ = "STUB: not implemented"
	return nil
}
