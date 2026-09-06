package orm

import (
	"context"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

type Invocation struct {
	Method string

	Md interface{}

	Args []interface{}

	mi *models.ModelInfo

	f func(ctx context.Context) []interface{}

	InsideTx    bool
	TxStartTime time.Time
	TxName      string
}

func (inv *Invocation) GetTableName() string { _ = "STUB: not implemented"; return "" }

func (inv *Invocation) execute(ctx context.Context) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (inv *Invocation) GetPkFieldName() string { _ = "STUB: not implemented"; return "" }
