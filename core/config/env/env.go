package env

import (
	"os"
	"strings"
	"sync"
)

var env sync.Map

func init() {
	for _, e := range os.Environ() {
		splits := strings.Split(e, "=")
		env.Store(splits[0], os.Getenv(splits[0]))
	}
	env.Store("GOBIN", GetGOBIN())
	env.Store("GOPATH", GetGOPATH())
}

func Get(key string, defVal string) string { _ = "STUB: not implemented"; return "" }

func MustGet(key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func Set(key string, value string) { _ = "STUB: not implemented"; return }

func MustSet(key string, value string) error { _ = "STUB: not implemented"; return nil }

func GetAll() map[string]string { _ = "STUB: not implemented"; return nil }

func envFile() (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetRuntimeEnv(key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetGOBIN() string { _ = "STUB: not implemented"; return "" }

func GetGOPATH() string { _ = "STUB: not implemented"; return "" }
