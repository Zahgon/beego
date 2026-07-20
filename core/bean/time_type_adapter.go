package bean

import (
	"context"
)

type TimeTypeAdapter struct {
	Layout string
}

func (t *TimeTypeAdapter) DefaultValue(ctx context.Context, dftValue string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
