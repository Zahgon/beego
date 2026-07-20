package orm

var quote string = `"`

type PostgresQueryBuilder struct {
	tokens []string
}

func processingStr(str []string) string { _ = "STUB: not implemented"; return "" }

func (qb *PostgresQueryBuilder) Select(fields ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) ForUpdate() QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) From(tables ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) InnerJoin(table string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) LeftJoin(table string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) RightJoin(table string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) On(cond string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) Where(cond string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) And(cond string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) Or(cond string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) In(vals ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) OrderBy(fields ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) Asc() QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) Desc() QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) Limit(limit int) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) Offset(offset int) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) GroupBy(fields ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) Having(cond string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) Update(tables ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) Set(kv ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) Delete(tables ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) InsertInto(table string, fields ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) Values(vals ...string) QueryBuilder {
	_ = "STUB: not implemented"
	return *new(QueryBuilder)
}

func (qb *PostgresQueryBuilder) Subquery(sub string, alias string) string {
	_ = "STUB: not implemented"
	return ""
}

func (qb *PostgresQueryBuilder) String() string { _ = "STUB: not implemented"; return "" }
