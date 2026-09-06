package mock

import (
	"database/sql"

	"github.com/beego/beego/v2/client/orm"
)

type DoNothingRawSetter struct{}

func (d *DoNothingRawSetter) Exec() (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (d *DoNothingRawSetter) QueryRow(containers ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DoNothingRawSetter) QueryRows(containers ...interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingRawSetter) SetArgs(i ...interface{}) orm.RawSeter {
	_ = "STUB: not implemented"
	return *new(orm.RawSeter)
}

func (d *DoNothingRawSetter) Values(container *[]orm.Params, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingRawSetter) ValuesList(container *[]orm.ParamsList, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingRawSetter) ValuesFlat(container *orm.ParamsList, cols ...string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingRawSetter) RowsToMap(result *orm.Params, keyCol, valueCol string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingRawSetter) RowsToStruct(ptrStruct interface{}, keyCol, valueCol string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *DoNothingRawSetter) Prepare() (orm.RawPreparer, error) {
	_ = "STUB: not implemented"
	return *new(orm.RawPreparer), nil
}
