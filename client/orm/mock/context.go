package mock

import (
	"context"
)

type mockCtxKeyType string

const mockCtxKey = mockCtxKeyType("beego-orm-mock")

func CtxWithMock(ctx context.Context, mock ...*Mock) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func mockFromCtx(ctx context.Context) []*Mock { _ = "STUB: not implemented"; return nil }
