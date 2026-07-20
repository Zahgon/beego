package orm

import (
	"reflect"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func getDbAlias(name string) *alias { _ = "STUB: not implemented"; return nil }

func getExistPk(mi *models.ModelInfo, ind reflect.Value) (column string, value interface{}, exist bool) {
	_ = "STUB: not implemented"
	return "", nil, false
}

func getFlatParams(fi *models.FieldInfo, args []interface{}, tz *time.Location) (params []interface{}) {
	_ = "STUB: not implemented"
	return nil
}
