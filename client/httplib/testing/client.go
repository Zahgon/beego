package testing

import (
	"github.com/beego/beego/v2/client/httplib"
)

var (
	port    = ""
	baseURL = "http://localhost:"
)

type TestHTTPRequest struct {
	httplib.BeegoHTTPRequest
}

func SetTestingPort(p string) { _ = "STUB: not implemented"; return }

func getPort() string { _ = "STUB: not implemented"; return "" }

func Get(path string) *TestHTTPRequest { _ = "STUB: not implemented"; return nil }

func Post(path string) *TestHTTPRequest { _ = "STUB: not implemented"; return nil }

func Put(path string) *TestHTTPRequest { _ = "STUB: not implemented"; return nil }

func Delete(path string) *TestHTTPRequest { _ = "STUB: not implemented"; return nil }

func Head(path string) *TestHTTPRequest { _ = "STUB: not implemented"; return nil }
