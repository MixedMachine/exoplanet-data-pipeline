package api

import (
	"fmt"
)

type QueryBuilder struct {
	query string
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		query: "?query=",
	}
}

func (q *QueryBuilder) AddSelect(selectClause string) *QueryBuilder {
	q.query += "select+" + selectClause
	return q
}

func (q *QueryBuilder) AddFrom(fromClause string) *QueryBuilder {
	q.query += "+from+" + fromClause
	return q
}

func (q *QueryBuilder) AddWhere() *QueryBuilder {
	q.query += "+where"
	return q
}

func (q *QueryBuilder) AddWhereParameter(field, operator, value string) *QueryBuilder {
	q.query += fmt.Sprintf("+%s+%s+'%s'", field, operator, value)
	return q
}

func (q *QueryBuilder) AddAndWhereParameter(field, operator, value string) *QueryBuilder {
	q.query += fmt.Sprintf("+and+%s+%s+'%s'", field, operator, value)
	return q
}

func (q *QueryBuilder) AddOrWhereParameter(field, operator, value string) *QueryBuilder {
	q.query += fmt.Sprintf("+or+%s+%s+'%s'", field, operator, value)
	return q
}

func (q *QueryBuilder) AddFormat(formatClause string) *QueryBuilder {
	q.query += "&format=" + formatClause
	return q
}

func (q *QueryBuilder) Build() string {
	return q.query
}

func (q *QueryBuilder) GetQuery() string {
	return q.query
}

func (q *QueryBuilder) ResetQuery() {
	q.query = "?query="
}
