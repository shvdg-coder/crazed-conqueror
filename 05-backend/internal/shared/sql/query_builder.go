package sql

import (
	"strconv"
	"strings"
)

// QueryBuilder provides a fluent API for building SQL queries
type QueryBuilder struct {
	query        strings.Builder
	args         []any
	batchArgs    [][]any
	paramIndex   int
	fields       []QueryField
	hasWhere     bool
	hasSet       bool
	insertFields []string
}

// NewQuery creates a new QueryBuilder instance
func NewQuery() *QueryBuilder {
	return &QueryBuilder{
		args:       make([]any, 0),
		paramIndex: 1,
		fields:     make([]QueryField, 0),
	}
}

// String returns the built query string
func (qb *QueryBuilder) String() string {
	return qb.query.String()
}

// Arguments returns the query args
func (qb *QueryBuilder) Arguments() []any {
	return qb.args
}

// hasWhere returns true if WHERE clause has been added
func (qb *QueryBuilder) hasWhereClause() bool {
	return qb.hasWhere
}

// hasSet returns true if SET clause has been added
func (qb *QueryBuilder) hasSetClause() bool {
	return qb.hasSet
}

// Build returns the final query string and args
func (qb *QueryBuilder) Build() (string, []any) {
	return qb.String(), qb.Arguments()
}

// BuildBatch returns the final query string and batch args
func (qb *QueryBuilder) BuildBatch() (string, [][]any) {
	return qb.String(), qb.batchArgs
}

// SELECT Methods

// Select adds a SELECT clause with field names
func (qb *QueryBuilder) Select(fields ...string) *QueryBuilder {
	qb.query.WriteString("SELECT ")
	qb.query.WriteString(strings.Join(fields, ", "))
	return qb
}

// SelectFields adds a SELECT clause with QueryField objects
func (qb *QueryBuilder) SelectFields(fields ...QueryField) *QueryBuilder {
	qb.query.WriteString("SELECT ")
	fieldStrings := make([]string, len(fields))
	for i, field := range fields {
		fieldStrings[i] = field.String()
	}
	qb.query.WriteString(strings.Join(fieldStrings, ", "))
	return qb
}

// From adds a FROM clause
func (qb *QueryBuilder) From(table string) *QueryBuilder {
	qb.query.WriteString(" FROM ")
	qb.query.WriteString(table)
	return qb
}

// Where adds a WHERE condition
func (qb *QueryBuilder) Where(field string, value any) *QueryBuilder {
	if !qb.hasWhereClause() {
		qb.query.WriteString(" WHERE ")
		qb.hasWhere = true
	} else {
		qb.query.WriteString(" AND ")
	}

	qb.query.WriteString(field)
	qb.query.WriteString(" = $")
	qb.query.WriteString(strconv.Itoa(qb.paramIndex))
	qb.args = append(qb.args, value)
	qb.paramIndex++

	return qb
}

// WhereIn adds a WHERE IN condition
func (qb *QueryBuilder) WhereIn(field string, values ...any) *QueryBuilder {
	if len(values) == 0 {
		return qb
	}

	if !qb.hasWhereClause() {
		qb.query.WriteString(" WHERE ")
		qb.hasWhere = true
	} else {
		qb.query.WriteString(" AND ")
	}

	qb.query.WriteString(field)
	qb.query.WriteString(" IN (")

	placeholders := make([]string, len(values))
	for i, value := range values {
		placeholders[i] = "$" + strconv.Itoa(qb.paramIndex)
		qb.args = append(qb.args, value)
		qb.paramIndex++
	}

	qb.query.WriteString(strings.Join(placeholders, ", "))
	qb.query.WriteString(")")

	return qb
}

// WhereTupleIn adds a WHERE tuple IN condition for composite keys
func (qb *QueryBuilder) WhereTupleIn(tuples [][]any, fields ...string) *QueryBuilder {
	if len(tuples) == 0 || len(fields) == 0 {
		return qb
	}

	if !qb.hasWhereClause() {
		qb.query.WriteString(" WHERE ")
		qb.hasWhere = true
	} else {
		qb.query.WriteString(" AND ")
	}

	qb.query.WriteString("(")
	qb.query.WriteString(strings.Join(fields, ", "))
	qb.query.WriteString(") IN (")

	tuplePlaceholders := make([]string, len(tuples))
	for i, tuple := range tuples {
		valuePlaceholders := make([]string, len(tuple))
		for j, value := range tuple {
			valuePlaceholders[j] = "$" + strconv.Itoa(qb.paramIndex)
			qb.args = append(qb.args, value)
			qb.paramIndex++
		}
		tuplePlaceholders[i] = "(" + strings.Join(valuePlaceholders, ", ") + ")"
	}

	qb.query.WriteString(strings.Join(tuplePlaceholders, ", "))
	qb.query.WriteString(")")

	return qb
}

// INSERT Methods

// InsertInto adds an INSERT INTO clause
func (qb *QueryBuilder) InsertInto(table string) *QueryBuilder {
	qb.query.WriteString("INSERT INTO ")
	qb.query.WriteString(table)
	return qb
}

// InsertFields sets the field names for INSERT operations
func (qb *QueryBuilder) InsertFields(fields ...string) *QueryBuilder {
	qb.insertFields = fields
	return qb
}

// Values adds a VALUES clause for INSERT using previously set fields
func (qb *QueryBuilder) Values(values ...any) *QueryBuilder {
	qb.query.WriteString(" (")
	qb.query.WriteString(strings.Join(qb.insertFields, ", "))
	qb.query.WriteString(") VALUES (")

	placeholders := make([]string, len(values))
	for i, value := range values {
		placeholders[i] = "$" + strconv.Itoa(qb.paramIndex)
		qb.args = append(qb.args, value)
		qb.paramIndex++
	}

	qb.query.WriteString(strings.Join(placeholders, ", "))
	qb.query.WriteString(")")

	return qb
}

// BatchValues adds a template VALUES clause for batch INSERT operations
func (qb *QueryBuilder) BatchValues(argumentSets [][]any) *QueryBuilder {
	if len(argumentSets) == 0 {
		return qb
	}

	qb.batchArgs = argumentSets

	qb.query.WriteString(" (")
	qb.query.WriteString(strings.Join(qb.insertFields, ", "))
	qb.query.WriteString(") VALUES (")

	fieldCount := len(qb.insertFields)
	placeholders := make([]string, fieldCount)
	for i := 0; i < fieldCount; i++ {
		placeholders[i] = "$" + strconv.Itoa(i+1)
	}

	qb.query.WriteString(strings.Join(placeholders, ", "))
	qb.query.WriteString(")")

	return qb
}

// UPDATE Methods

// Update adds an UPDATE clause
func (qb *QueryBuilder) Update(table string) *QueryBuilder {
	qb.query.WriteString("UPDATE ")
	qb.query.WriteString(table)
	return qb
}

// Set adds a SET condition for UPDATE
func (qb *QueryBuilder) Set(field string, value any) *QueryBuilder {
	if !qb.hasSetClause() {
		qb.query.WriteString(" SET ")
		qb.hasSet = true
	} else {
		qb.query.WriteString(", ")
	}

	qb.query.WriteString(field)
	qb.query.WriteString(" = $")
	qb.query.WriteString(strconv.Itoa(qb.paramIndex))
	qb.args = append(qb.args, value)
	qb.paramIndex++

	return qb
}

// BatchSets adds SET clauses using numbered placeholders for batch UPDATE operations
func (qb *QueryBuilder) BatchSets(argumentSets [][]any, setFields ...string) *QueryBuilder {
	if len(argumentSets) == 0 || len(setFields) == 0 {
		return qb
	}

	qb.batchArgs = argumentSets

	qb.query.WriteString(" SET ")

	setClauses := make([]string, len(setFields))
	for i, field := range setFields {
		setClauses[i] = field + " = $" + strconv.Itoa(qb.paramIndex)
		qb.paramIndex++
	}

	qb.query.WriteString(strings.Join(setClauses, ", "))
	qb.hasSet = true

	return qb
}

// DELETE Methods

// DeleteFrom adds a DELETE FROM clause
func (qb *QueryBuilder) DeleteFrom(table string) *QueryBuilder {
	qb.query.WriteString("DELETE FROM ")
	qb.query.WriteString(table)
	return qb
}

// Common Methods

// Returning adds a RETURNING clause
func (qb *QueryBuilder) Returning(fields ...string) *QueryBuilder {
	qb.query.WriteString(" RETURNING ")
	qb.query.WriteString(strings.Join(fields, ", "))
	return qb
}
