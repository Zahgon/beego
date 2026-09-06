package bean

import (
	"context"
)

type AutoWireBeanFactory interface {
	AutoWire(ctx context.Context, appCtx ApplicationContext, bean interface{}) error
}
