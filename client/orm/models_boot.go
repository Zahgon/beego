package orm

import (
	imodels "github.com/beego/beego/v2/client/orm/internal/models"
)

var defaultModelCache = imodels.NewModelCacheHandler()

func RegisterModel(models ...interface{}) { _ = "STUB: not implemented"; return }

func RegisterModelWithPrefix(prefix string, models ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func RegisterModelWithSuffix(suffix string, models ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func BootStrap() { _ = "STUB: not implemented"; return }

func BootStrapWithAlias(alias string) { _ = "STUB: not implemented"; return }

func ResetModelCache() { _ = "STUB: not implemented"; return }
