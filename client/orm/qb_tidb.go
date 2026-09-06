package orm

type TiDBQueryBuilder struct {
	MySQLQueryBuilder
	tokens []string
}
