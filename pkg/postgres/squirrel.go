package postgres

import sq "github.com/Masterminds/squirrel"

type SqlBuilder struct {
	sq.StatementBuilderType
}

func newSqlBuilder() *SqlBuilder {
	return &SqlBuilder{sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}
}

func (s *SqlBuilder) Equal(key string, value interface{}) sq.Eq {
	return sq.Eq{key: value}
}
