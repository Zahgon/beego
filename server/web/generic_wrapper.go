package web

import (
	"github.com/beego/beego/v2/server/web/context"
)

type bizFunc[T any] func(ctx *context.Context, param T) (any, error)

type extractFunc[T any] func(ctx *context.Context) (params T, err error)

func WrapperFromJson[T any](
	biz bizFunc[T]) func(ctx *context.Context) {
	_ = "STUB: not implemented"
	return nil
}

func WrapperFromForm[T any](
	biz bizFunc[T]) func(ctx *context.Context) {
	_ = "STUB: not implemented"
	return nil
}

func Wrapper[T any](
	biz bizFunc[T]) func(ctx *context.Context) {
	_ = "STUB: not implemented"
	return nil
}

func internalWrapper[T any](
	biz bizFunc[T],
	ef extractFunc[T]) func(ctx *context.Context) {
	_ = "STUB: not implemented"
	return nil
}
