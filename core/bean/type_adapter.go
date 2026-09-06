package bean

import (
	"context"
)

type TypeAdapter interface {
	DefaultValue(ctx context.Context, dftValue string) (interface{}, error)
}
