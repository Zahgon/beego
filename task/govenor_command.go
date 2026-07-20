package task

import (
	"github.com/beego/beego/v2/core/admin"
)

type listTaskCommand struct{}

func (l *listTaskCommand) Execute(params ...interface{}) *admin.Result {
	_ = "STUB: not implemented"
	return nil
}

type runTaskCommand struct{}

func (r *runTaskCommand) Execute(params ...interface{}) *admin.Result {
	_ = "STUB: not implemented"
	return nil
}

func registerCommands() { _ = "STUB: not implemented"; return }
