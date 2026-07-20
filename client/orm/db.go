package orm

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/buffers"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

var ErrMissPK = errors.New("missed pk value")

var operators = map[string]bool{
	"exact":       true,
	"iexact":      true,
	"strictexact": true,
	"contains":    true,
	"icontains":   true,

	"gt":          true,
	"gte":         true,
	"lt":          true,
	"lte":         true,
	"eq":          true,
	"nq":          true,
	"ne":          true,
	"startswith":  true,
	"endswith":    true,
	"istartswith": true,
	"iendswith":   true,
	"in":          true,
	"between":     true,

	"isnull": true,
}

type dbBase struct {
	ins dbBaser
}

var _ dbBaser = new(dbBase)

func (d *dbBase) collectValues(mi *models.ModelInfo, ind reflect.Value, cols []string, skipAuto bool, insert bool, names *[]string, tz *time.Location) (values []interface{}, autoFields []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (d *dbBase) collectFieldValue(mi *models.ModelInfo, fi *models.FieldInfo, ind reflect.Value, insert bool, tz *time.Location) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *dbBase) PrepareInsert(ctx context.Context, q dbQuerier, mi *models.ModelInfo) (stmtQuerier, string, error) {
	_ = "STUB: not implemented"
	return *new(stmtQuerier), "", nil
}

func (d *dbBase) InsertStmt(ctx context.Context, stmt stmtQuerier, mi *models.ModelInfo, ind reflect.Value, tz *time.Location) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *dbBase) Read(ctx context.Context, q dbQuerier, mi *models.ModelInfo, ind reflect.Value, tz *time.Location, cols []string, isForUpdate bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *dbBase) Insert(ctx context.Context, q dbQuerier, mi *models.ModelInfo, ind reflect.Value, tz *time.Location) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *dbBase) InsertMulti(ctx context.Context, q dbQuerier, mi *models.ModelInfo, sind reflect.Value, bulk int, tz *time.Location) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *dbBase) InsertValue(ctx context.Context, q dbQuerier, mi *models.ModelInfo, isMulti bool, names []string, values []interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *dbBase) InsertValueSQL(names []string, values []interface{}, isMulti bool, mi *models.ModelInfo) string {
	_ = "STUB: not implemented"
	return ""
}

func (d *dbBase) InsertOrUpdate(ctx context.Context, q dbQuerier, mi *models.ModelInfo, ind reflect.Value, a *alias, args ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *dbBase) InsertOrUpdateSQL(names []string, values *[]interface{}, mi *models.ModelInfo, a *alias, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *dbBase) Update(ctx context.Context, q dbQuerier, mi *models.ModelInfo, ind reflect.Value, tz *time.Location, cols []string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *dbBase) UpdateSQL(setNames []string, pkName string, mi *models.ModelInfo) string {
	_ = "STUB: not implemented"
	return ""
}

func (d *dbBase) Delete(ctx context.Context, q dbQuerier, mi *models.ModelInfo, ind reflect.Value, tz *time.Location, cols []string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *dbBase) DeleteSQL(whereCols []string, mi *models.ModelInfo) string {
	_ = "STUB: not implemented"
	return ""
}

func (d *dbBase) UpdateBatch(ctx context.Context, q dbQuerier, qs *querySet, mi *models.ModelInfo, cond *Condition, params Params, tz *time.Location) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *dbBase) UpdateBatchSQL(mi *models.ModelInfo, cols []string, values []interface{}, specifyIndexes, join, where string) string {
	_ = "STUB: not implemented"
	return ""
}

func (d *dbBase) buildSetSQL(buf buffers.Buffer, cols []string, values []interface{}) {
	_ = "STUB: not implemented"
	return
}

func (d *dbBase) deleteRels(ctx context.Context, q dbQuerier, mi *models.ModelInfo, args []interface{}, tz *time.Location) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *dbBase) DeleteBatch(ctx context.Context, q dbQuerier, qs *querySet, mi *models.ModelInfo, cond *Condition, tz *time.Location) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *dbBase) ReadBatch(ctx context.Context, q dbQuerier, qs querySet, mi *models.ModelInfo, cond *Condition, container interface{}, tz *time.Location, cols []string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *dbBase) readBatchSQL(tables *dbTables, tCols []string, cond *Condition, qs querySet, mi *models.ModelInfo, tz *time.Location) (string, []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *dbBase) preProcCols(cols []string) []string { _ = "STUB: not implemented"; return nil }

func (d *dbBase) readSQL(buf buffers.Buffer, tables *dbTables, tCols []string, cond *Condition, qs querySet, mi *models.ModelInfo, tz *time.Location) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (d *dbBase) Count(ctx context.Context, q dbQuerier, qs querySet, mi *models.ModelInfo, cond *Condition, tz *time.Location) (cnt int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *dbBase) countSQL(qs querySet, mi *models.ModelInfo, cond *Condition, tz *time.Location) (string, []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *dbBase) GenerateOperatorSQL(mi *models.ModelInfo, fi *models.FieldInfo, operator string, args []interface{}, tz *time.Location) (string, []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *dbBase) GenerateOperatorLeftCol(*models.FieldInfo, string, *string) {
	_ = "STUB: not implemented"
	return
}

func (d *dbBase) setColsValues(mi *models.ModelInfo, ind *reflect.Value, cols []string, values []interface{}, tz *time.Location) {
	_ = "STUB: not implemented"
	return
}

func (d *dbBase) convertValueFromDB(fi *models.FieldInfo, val interface{}, tz *time.Location) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *dbBase) setFieldValue(fi *models.FieldInfo, value interface{}, field reflect.Value) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *dbBase) ReadValues(ctx context.Context, q dbQuerier, qs querySet, mi *models.ModelInfo, cond *Condition, exprs []string, container interface{}, tz *time.Location) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *dbBase) readValuesSQL(tables *dbTables, cols []string, qs querySet, mi *models.ModelInfo, cond *Condition, tz *time.Location) (string, []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *dbBase) SupportUpdateJoin() bool { _ = "STUB: not implemented"; return false }

func (d *dbBase) MaxLimit() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *dbBase) TableQuote() string { _ = "STUB: not implemented"; return "" }

func (d *dbBase) ReplaceMarks(query *string) { _ = "STUB: not implemented"; return }

func (d *dbBase) HasReturningID(*models.ModelInfo, *string) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *dbBase) setval(ctx context.Context, db dbQuerier, mi *models.ModelInfo, autoFields []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *dbBase) TimeFromDB(t *time.Time, tz *time.Location) { _ = "STUB: not implemented"; return }

func (d *dbBase) TimeToDB(t *time.Time, tz *time.Location) { _ = "STUB: not implemented"; return }

func (d *dbBase) DbTypes() map[string]string { _ = "STUB: not implemented"; return nil }

func (d *dbBase) GetTables(db dbQuerier) (map[string]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *dbBase) GetColumns(ctx context.Context, db dbQuerier, table string) (map[string][3]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *dbBase) OperatorSQL(operator string) string { _ = "STUB: not implemented"; return "" }

func (d *dbBase) ShowTablesQuery() string { _ = "STUB: not implemented"; return "" }

func (d *dbBase) ShowColumnsQuery(table string) string { _ = "STUB: not implemented"; return "" }

func (d *dbBase) IndexExists(context.Context, dbQuerier, string, string) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *dbBase) GenerateSpecifyIndex(tableName string, useIndex int, indexes []string) string {
	_ = "STUB: not implemented"
	return ""
}
