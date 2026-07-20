package orm

import (
	"github.com/beego/beego/v2/client/orm/internal/models"
)

type dbIndex struct {
	Table string
	Name  string
	SQL   string
}

func getColumnTyp(al *alias, fi *models.FieldInfo) (col string) {
	_ = "STUB: not implemented"
	return ""
}

func getColumnAddQuery(al *alias, fi *models.FieldInfo) string {
	_ = "STUB: not implemented"
	return ""
}

func getColumnDefault(fi *models.FieldInfo) string { _ = "STUB: not implemented"; return "" }
