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

func (s *SqlBuilder) Gt(key string, value interface{}) sq.Gt {
	return sq.Gt{key: value}
}

func (s *SqlBuilder) GtOrEqual(key string, value interface{}) sq.GtOrEq {
	return sq.GtOrEq{key: value}
}

func (s *SqlBuilder) Lt(key string, value interface{}) sq.Lt {
	return sq.Lt{key: value}
}

func (s *SqlBuilder) LtOrEq(key string, value interface{}) sq.LtOrEq {
	return sq.LtOrEq{key: value}
}
