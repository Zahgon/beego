package es

import (
	"github.com/beego/beego/v2/core/logs"
)

type IndexNaming interface {
	IndexName(lm *logs.LogMsg) string
}

var indexNaming IndexNaming = &defaultIndexNaming{}

func SetIndexNaming(i IndexNaming) { _ = "STUB: not implemented"; return }

type defaultIndexNaming struct{}

func (d *defaultIndexNaming) IndexName(lm *logs.LogMsg) string {
	_ = "STUB: not implemented"
	return ""
}
