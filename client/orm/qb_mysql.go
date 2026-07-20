package orm

const CommaSpace = ", "

type MySQLQueryBuilder struct {
	tokens []string
}

func (qb *MySQLQueryBuilder) Select(fields ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) ForUpdate() QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) From(tables ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) InnerJoin(table string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) LeftJoin(table string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) RightJoin(table string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) On(cond string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) Where(cond string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) And(cond string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) Or(cond string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) In(vals ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) OrderBy(fields ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) Asc() QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) Desc() QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) Limit(limit int) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) Offset(offset int) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) GroupBy(fields ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) Having(cond string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) Update(tables ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) Set(kv ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) Delete(tables ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) InsertInto(table string, fields ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) Values(vals ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *MySQLQueryBuilder) Subquery(sub string, alias string) string {
	_ = "STUB: not implemented"
	return ""
}

func (qb *MySQLQueryBuilder) String() string { _ = "STUB: not implemented"; return "" }
