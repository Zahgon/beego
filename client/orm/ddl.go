package orm

import (
	imodels "github.com/beego/beego/v2/client/orm/internal/models"
)

func getDbDropSQL(mc *imodels.ModelCache, al *alias) (queries []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDbCreateSQL(mc *imodels.ModelCache, al *alias) (queries []string, tableIndexes map[string][]dbIndex, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
