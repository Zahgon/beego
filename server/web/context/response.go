package context

import (
	"net/http"
)

const (
	BadRequest StatusCode = http.StatusBadRequest

	NotFound StatusCode = http.StatusNotFound
)

type StatusCode int

func (s StatusCode) Error() string { _ = "STUB: not implemented"; return "" }

func (s StatusCode) Render(ctx *Context) { _ = "STUB: not implemented"; return }
