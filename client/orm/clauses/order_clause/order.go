package order_clause

type Sort int8

const (
	None       Sort = 0
	Ascending  Sort = 1
	Descending Sort = 2
)

type Option func(order *Order)

type Order struct {
	column string
	sort   Sort
	isRaw  bool
}

func Clause(options ...Option) *Order { _ = "STUB: not implemented"; return nil }

func (o *Order) GetColumn() string { _ = "STUB: not implemented"; return "" }

func (o *Order) GetSort() Sort { _ = "STUB: not implemented"; return *new(Sort) }

func (o *Order) SortString() string { _ = "STUB: not implemented"; return "" }

func (o *Order) IsRaw() bool { _ = "STUB: not implemented"; return false }

func ParseOrder(expressions ...string) []*Order { _ = "STUB: not implemented"; return nil }

func Column(column string) Option { _ = "STUB: not implemented"; return *new(Option) }

func sort(sort Sort) Option { _ = "STUB: not implemented"; return *new(Option) }

func SortAscending() Option { _ = "STUB: not implemented"; return *new(Option) }

func SortDescending() Option { _ = "STUB: not implemented"; return *new(Option) }

func SortNone() Option { _ = "STUB: not implemented"; return *new(Option) }

func Raw() Option { _ = "STUB: not implemented"; return *new(Option) }
