package berror

import (
	"sync"
)

type Code interface {
	Code() uint32
	Module() string
	Desc() string
	Name() string
}

var defaultCodeRegistry = &codeRegistry{
	codes: make(map[uint32]*codeDefinition, 127),
}

func DefineCode(code uint32, module string, name string, desc string) Code {
	_ = "STUB: not implemented"
	return *new(Code)
}

type codeRegistry struct {
	lock  sync.RWMutex
	codes map[uint32]*codeDefinition
}

func (cr *codeRegistry) Get(code uint32) (Code, bool) {
	_ = "STUB: not implemented"
	return *new(Code), false
}

type codeDefinition struct {
	code   uint32
	module string
	desc   string
	name   string
}

func (c *codeDefinition) Name() string { _ = "STUB: not implemented"; return "" }

func (c *codeDefinition) Code() uint32 { _ = "STUB: not implemented"; return 0 }

func (c *codeDefinition) Module() string { _ = "STUB: not implemented"; return "" }

func (c *codeDefinition) Desc() string { _ = "STUB: not implemented"; return "" }
