package sql

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SQL Utils", Ordered, func() {

	Describe("Query building", func() {
		var fields []string
		var table string

		BeforeAll(func() {
			fields = []string{"id", "name", "email"}
			table = "users"
		})

		It("should build a valid INSERT query", func() {
			query := BuildInsertQuery(table, fields)
			Expect(query).To(Equal("INSERT INTO users (id, name, email) VALUES ($1, $2, $3)"))
		})

		It("should build a valid INSERT RETURNING query", func() {
			query := BuildInsertReturningQuery(table, fields, fields)
			Expect(query).To(Equal("INSERT INTO users (id, name, email) VALUES ($1, $2, $3) RETURNING id, name, email"))
		})

		It("should build a valid SELECT query", func() {
			whereClauses := CreateDollarClause(1, []string{"id"})
			query := BuildSelectQuery(table, fields, whereClauses...)
			Expect(query).To(Equal("SELECT id, name, email FROM users WHERE id = $1"))
		})

		It("should build a valid UPDATE query", func() {
			setClauses := CreateDollarClause(1, []string{"name", "email"})
			whereClauses := CreateDollarClause(3, []string{"id"})
			query := BuildUpdateQuery(table, setClauses, whereClauses)
			Expect(query).To(Equal("UPDATE users SET name = $1, email = $2 WHERE id = $3"))
		})

		It("should build a valid UPDATE RETURNING query", func() {
			setClauses := CreateDollarClause(1, []string{"name", "email"})
			whereClauses := CreateDollarClause(3, []string{"id"})
			returnFields := []string{"id", "name", "email"}
			query := BuildUpdateReturningQuery(table, setClauses, whereClauses, returnFields)
			Expect(query).To(Equal("UPDATE users SET name = $1, email = $2 WHERE id = $3 RETURNING id, name, email"))
		})

		It("should build a valid DELETE query", func() {
			whereClauses := CreateDollarClause(1, []string{"id"})
			query := BuildDeleteQuery(table, whereClauses)
			Expect(query).To(Equal("DELETE FROM users WHERE id = $1"))
		})

		It("should build a valid DELETE RETURNING query", func() {
			whereClauses := CreateDollarClause(1, []string{"id"})
			returnFields := []string{"id", "name", "email"}
			query := BuildDeleteReturningQuery(table, whereClauses, returnFields)
			Expect(query).To(Equal("DELETE FROM users WHERE id = $1 RETURNING id, name, email"))
		})

		It("should build a valid SELECT FROM query", func() {
			sourceTable := "temp"
			whereClauses := CreateNamedClause([]string{"id"}, table, sourceTable)
			query := BuildSelectFromQuery(table, sourceTable, fields, whereClauses)
			Expect(query).To(Equal("SELECT id, name, email FROM users WHERE EXISTS (SELECT 1 FROM temp WHERE users.id = temp.id)"))
		})

		It("should build a valid UPDATE FROM query", func() {
			sourceTable := "temp"
			setClauses := CreateNamedClause([]string{"name", "email"}, "", sourceTable)
			whereClauses := CreateNamedClause([]string{"id"}, table, sourceTable)
			query := BuildUpdateFromQuery(table, sourceTable, setClauses, whereClauses)
			Expect(query).To(Equal("UPDATE users SET name = temp.name, email = temp.email FROM temp WHERE users.id = temp.id"))
		})

		It("should build a valid DELETE FROM query", func() {
			sourceTable := "temp"
			whereClauses := CreateNamedClause([]string{"id"}, table, sourceTable)
			query := BuildDeleteFromQuery(table, sourceTable, whereClauses)
			Expect(query).To(Equal("DELETE FROM users WHERE EXISTS (SELECT 1 FROM temp WHERE users.id = temp.id)"))
		})

		It("should build a valid UPSERT query with update fields", func() {
			insertFields := []string{"id", "name", "email"}
			keyFields := []string{"id"}
			updateFields := []string{"name", "email"}
			query := BuildUpsertQuery(table, insertFields, keyFields, updateFields)
			Expect(query).To(Equal("INSERT INTO users (id, name, email) VALUES ($1, $2, $3) ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name, email = EXCLUDED.email"))
		})

		It("should build a valid UPSERT query without update fields", func() {
			insertFields := []string{"id", "name", "email"}
			keyFields := []string{"id"}
			var updateFields []string
			query := BuildUpsertQuery(table, insertFields, keyFields, updateFields)
			Expect(query).To(Equal("INSERT INTO users (id, name, email) VALUES ($1, $2, $3) ON CONFLICT (id) DO NOTHING"))
		})

		It("should build a valid UPSERT RETURNING query", func() {
			insertFields := []string{"id", "name", "email"}
			keyFields := []string{"id"}
			updateFields := []string{"name", "email"}
			returnFields := []string{"id", "name", "email", "created_at", "updated_at"}
			query := BuildUpsertReturningQuery(table, insertFields, keyFields, updateFields, returnFields)
			Expect(query).To(Equal("INSERT INTO users (id, name, email) VALUES ($1, $2, $3) ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name, email = EXCLUDED.email RETURNING id, name, email, created_at, updated_at"))
		})

		It("should build a valid COUNT query", func() {
			whereClauses := CreateDollarClause(1, []string{"id"})
			query := BuildCountQuery(table, whereClauses)
			Expect(query).To(Equal("SELECT COUNT(*) FROM users WHERE id = $1"))
		})

		It("should build a valid composite IN clause with tuples", func() {
			cols := []string{"col1", "col2"}
			count := 2
			startIdx := 1
			clause := CreateTupleInClause(cols, count, startIdx)
			Expect(clause).To(Equal("(col1, col2) IN (($1, $2), ($3, $4))"))
		})
	})

	Describe("Helper functions", func() {
		It("should generate a unique temp table name", func() {
			tableName := GenerateTempTableName("users")
			Expect(tableName).To(HavePrefix("users_temp_"))
			Expect(tableName).To(HaveLen(43)) // "users_temp_" + 32-character UUID
		})

		It("should build a temp table query", func() {
			query := BuildTempTableQuery("users_temp", "users", []string{"id", "name"})
			Expect(query).To(Equal("CREATE TEMPORARY TABLE users_temp AS SELECT id, name FROM users WHERE 1=0"))
		})

		It("should build a drop table query", func() {
			query := BuildDropTableQuery("users_temp")
			Expect(query).To(Equal("DROP TABLE IF EXISTS users_temp CASCADE"))
		})
	})
})
