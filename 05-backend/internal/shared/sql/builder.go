package sql

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// CreateDollarHolders returns an array of dollar placeholders ($1, $2, etc.) for query parameters
func CreateDollarHolders(startIndex int, fields []string) []string {
	holders := make([]string, len(fields))
	for i := range holders {
		holders[i] = "$" + strconv.Itoa(startIndex+i)
	}
	return holders
}

// CreateDollarClause returns an array of field equality clauses using dollar placeholders
func CreateDollarClause(startIndex int, fields []string) []string {
	clauses := make([]string, len(fields))
	for i, field := range fields {
		clauses[i] = field + " = $" + strconv.Itoa(startIndex+i)
	}
	return clauses
}

// CreateNamedClause returns an array of field equality clauses using table-qualified field names
func CreateNamedClause(fields []string, targetTable, sourceTable string) []string {
	clauses := make([]string, len(fields))
	for i, field := range fields {
		if targetTable != "" {
			clauses[i] = targetTable + "." + field + " = " + sourceTable + "." + field
		} else {
			clauses[i] = field + " = " + sourceTable + "." + field
		}
	}
	return clauses
}

// CreateTupleInClause builds a composite IN clause
func CreateTupleInClause(columnNames []string, tupleCount, startIndex int) string {
	var tuples []string
	argIndex := startIndex
	for i := 0; i < tupleCount; i++ {
		holders := make([]string, len(columnNames))
		for j := range columnNames {
			holders[j] = "$" + strconv.Itoa(argIndex)
			argIndex++
		}
		tuples = append(tuples, "("+strings.Join(holders, ", ")+")")
	}
	return "(" + strings.Join(columnNames, ", ") + ") IN (" + strings.Join(tuples, ", ") + ")"
}

// GenerateTempTableName returns a unique temporary table name based on the base table name
func GenerateTempTableName(baseTableName string) string {
	uuidStr := strings.ReplaceAll(uuid.New().String(), "-", "")
	return baseTableName + "_temp_" + uuidStr
}

// BuildTempTableQuery returns a query string to create an empty temporary table
func BuildTempTableQuery(tempTableName, targetTableName string, columns []string) string {
	return "CREATE TEMPORARY TABLE " + tempTableName + " AS SELECT " + strings.Join(columns, ", ") + " FROM " + targetTableName + " WHERE 1=0"
}

// BuildDropTableQuery returns a query string to drop a table if it exists
func BuildDropTableQuery(table string) string {
	return "DROP TABLE IF EXISTS " + table + " CASCADE"
}

// BuildCountQuery returns a COUNT query string with the specified WHERE clauses
func BuildCountQuery(table string, whereClauses []string) string {
	return "SELECT COUNT(*) FROM " + table + " WHERE " + strings.Join(whereClauses, " AND ")
}

// BuildInsertQuery returns an INSERT query string for the specified table and fields
func BuildInsertQuery(table string, fields []string) string {
	return "INSERT INTO " + table + " (" + strings.Join(fields, ", ") + ") VALUES (" + strings.Join(CreateDollarHolders(1, fields), ", ") + ")"
}

// BuildInsertReturningQuery returns an INSERT query string with a RETURNING clause
func BuildInsertReturningQuery(table string, insertFields, returnFields []string) string {
	insertQuery := BuildInsertQuery(table, insertFields)
	return insertQuery + " RETURNING " + strings.Join(returnFields, ", ")
}

// BuildSelectQuery returns a SELECT query string with optional WHERE clauses
func BuildSelectQuery(table string, fields []string, whereClauses ...string) string {
	query := "SELECT " + strings.Join(fields, ", ") + " FROM " + table
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}
	return query
}

// BuildUpdateQuery returns an UPDATE query string with SET and WHERE clauses
func BuildUpdateQuery(table string, setClauses, whereClauses []string) string {
	return "UPDATE " + table + " SET " + strings.Join(setClauses, ", ") + " WHERE " + strings.Join(whereClauses, " AND ")
}

// BuildUpdateReturningQuery returns an UPDATE query string with a RETURNING clause
func BuildUpdateReturningQuery(table string, setClauses, whereClauses []string, returnFields []string) string {
	updateQuery := BuildUpdateQuery(table, setClauses, whereClauses)
	return updateQuery + " RETURNING " + strings.Join(returnFields, ", ")
}

// BuildDeleteQuery returns a DELETE query string with WHERE clauses
func BuildDeleteQuery(table string, whereClauses []string) string {
	return "DELETE FROM " + table + " WHERE " + strings.Join(whereClauses, " AND ")
}

// BuildDeleteReturningQuery returns a DELETE query string with RETURNING clause
func BuildDeleteReturningQuery(table string, whereClauses, returnFields []string) string {
	deleteQuery := BuildDeleteQuery(table, whereClauses)
	return deleteQuery + " RETURNING " + strings.Join(returnFields, ", ")
}

// BuildSelectFromQuery returns a SELECT query string with an EXISTS subquery
func BuildSelectFromQuery(targetTable, sourceTable string, fields, whereClauses []string) string {
	return "SELECT " + strings.Join(fields, ", ") + " FROM " + targetTable + " WHERE EXISTS (SELECT 1 FROM " + sourceTable + " WHERE " + strings.Join(whereClauses, " AND ") + ")"
}

// BuildUpdateFromQuery returns an UPDATE query string with a FROM clause
func BuildUpdateFromQuery(targetTable, sourceTable string, setClauses, whereClauses []string) string {
	return "UPDATE " + targetTable + " SET " + strings.Join(setClauses, ", ") + " FROM " + sourceTable + " WHERE " + strings.Join(whereClauses, " AND ")
}

// BuildDeleteFromQuery returns a DELETE query string with an EXISTS subquery
func BuildDeleteFromQuery(targetTable, sourceTable string, whereClauses []string) string {
	return "DELETE FROM " + targetTable + " WHERE EXISTS (SELECT 1 FROM " + sourceTable + " WHERE " + strings.Join(whereClauses, " AND ") + ")"
}

// BuildUpsertQuery returns an INSERT query string with ON CONFLICT DO UPDATE clause
func BuildUpsertQuery(table string, insertFields []string, keyFields, updateFields []string) string {
	insertQuery := BuildInsertQuery(table, insertFields)
	if len(updateFields) == 0 {
		return insertQuery + " ON CONFLICT (" + strings.Join(keyFields, ", ") + ") DO NOTHING"
	}

	setClauses := make([]string, len(updateFields))
	for i, field := range updateFields {
		setClauses[i] = field + " = EXCLUDED." + field
	}

	return insertQuery + " ON CONFLICT (" + strings.Join(keyFields, ", ") + ") DO UPDATE SET " + strings.Join(setClauses, ", ")
}

// BuildUpsertReturningQuery returns an INSERT query string with ON CONFLICT DO UPDATE and RETURNING clauses
func BuildUpsertReturningQuery(table string, insertFields, keyFields, updateFields, returnFields []string) string {
	upsertQuery := BuildUpsertQuery(table, insertFields, keyFields, updateFields)
	return upsertQuery + " RETURNING " + strings.Join(returnFields, ", ")
}
