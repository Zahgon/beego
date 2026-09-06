package admin

import (
	"errors"
)

type Command interface {
	Execute(params ...interface{}) *Result
}

var CommandNotFound = errors.New("Command not found")

type Result struct {
	Status  int
	Error   error
	Content interface{}
}

func (r *Result) IsSuccess() bool { _ = "STUB: not implemented"; return false }

type moduleCommands map[string]Command

func (m moduleCommands) Get(name string) Command { _ = "STUB: not implemented"; return *new(Command) }

type commandRegistry map[string]moduleCommands

func (c commandRegistry) Get(moduleName string) moduleCommands {
	_ = "STUB: not implemented"
	return *new(moduleCommands)
}

var cmdRegistry = make(commandRegistry)

func RegisterCommand(module string, commandName string, command Command) {
	_ = "STUB: not implemented"
	return
}

func GetCommand(module string, cmdName string) Command {
	_ = "STUB: not implemented"
	return *new(Command)
}

type doNothingCommand struct{}

func (d *doNothingCommand) Execute(params ...interface{}) *Result {
	_ = "STUB: not implemented"
	return nil
}
