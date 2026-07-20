package web

import (
	"sync"
)

const (
	DEV = "dev"

	PROD = "prod"
)

type M map[string]interface{}

type hookfunc func() error

var hooks = make([]hookfunc, 0)

func AddAPPStartHook(hf ...hookfunc) { _ = "STUB: not implemented"; return }

func Run(params ...string) { _ = "STUB: not implemented"; return }

func RunWithMiddleWares(addr string, mws ...MiddleWare) { _ = "STUB: not implemented"; return }

var initHttpOnce sync.Once

func initBeforeHTTPRun() { _ = "STUB: not implemented"; return }

func TestBeegoInit(ap string) { _ = "STUB: not implemented"; return }

func InitBeegoBeforeTest(appConfigPath string) { _ = "STUB: not implemented"; return }
