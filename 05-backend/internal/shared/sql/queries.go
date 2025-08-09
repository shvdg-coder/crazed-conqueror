package sql

import "fmt"

// CreateDeleteQuery creates a SQL delete statement
func CreateDeleteQuery(tableName, where string) string {
	return fmt.Sprintf("DELETE FROM %s WHERE %s", tableName, where)
}
