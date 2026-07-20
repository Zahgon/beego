package web

import (
	"net/http"
	"time"
)

var beeAdminApp *adminApp

var FilterMonitorFunc func(string, string, time.Duration, string, int) bool

func init() {
	FilterMonitorFunc = func(string, string, time.Duration, string, int) bool { return true }
}

func list(root string, p interface{}, m M) { _ = "STUB: not implemented"; return }

func writeJSON(rw http.ResponseWriter, jsonData []byte) { _ = "STUB: not implemented"; return }

type adminApp struct {
	*HttpServer
}

func (admin *adminApp) Run() { _ = "STUB: not implemented"; return }

func registerAdmin() error { _ = "STUB: not implemented"; return nil }
