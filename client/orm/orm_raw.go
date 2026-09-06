package orm

import (
	"database/sql"
	"reflect"
)

type rawPrepare struct {
	rs     *rawSet
	stmt   stmtQuerier
	closed bool
}

func (o *rawPrepare) Exec(args ...interface{}) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (o *rawPrepare) Close() error { _ = "STUB: not implemented"; return nil }

func newRawPreparer(rs *rawSet) (RawPreparer, error) {
	_ = "STUB: not implemented"
	return *new(RawPreparer), nil
}

type rawSet struct {
	query string
	args  []interface{}
	orm   *ormBase
}

var _ RawSeter = new(rawSet)

func (o rawSet) SetArgs(args ...interface{}) RawSeter {
	_ = "STUB: not implemented"
	return *new(RawSeter)
}

func (o *rawSet) Exec() (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (o *rawSet) setFieldValue(ind reflect.Value, value interface{}) {
	_ = "STUB: not implemented"
	return
}

func (o *rawSet) loopSetRefs(refs []interface{}, sInds []reflect.Value, nIndsPtr *[]reflect.Value, eTyps []reflect.Type, init bool) {
	_ = "STUB: not implemented"
	return
}

func (o *rawSet) QueryRow(containers ...interface{}) error { _ = "STUB: not implemented"; return nil }

func (o *rawSet) QueryRows(containers ...interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *rawSet) readValues(container interface{}, needCols []string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *rawSet) queryRowsTo(container interface{}, keyCol, valueCol string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *rawSet) Values(container *[]Params, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *rawSet) ValuesList(container *[]ParamsList, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *rawSet) ValuesFlat(container *ParamsList, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *rawSet) RowsToMap(result *Params, keyCol, valueCol string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *rawSet) RowsToStruct(ptrStruct interface{}, keyCol, valueCol string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *rawSet) Prepare() (RawPreparer, error) {
	_ = "STUB: not implemented"
	return *new(RawPreparer), nil
}

func newRawSet(orm *ormBase, query string, args []interface{}) RawSeter {
	_ = "STUB: not implemented"
	return *new(RawSeter)
}
