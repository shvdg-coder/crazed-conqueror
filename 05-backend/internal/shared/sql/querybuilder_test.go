package sql

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("QueryBuilder", func() {
	var qb *QueryBuilder

	BeforeEach(func() {
		qb = NewQuery()
	})

	Describe("SELECT queries", func() {
		Context("simple SELECT", func() {
			It("should build basic SELECT query", func() {
				query, args := qb.Select("id", "name", "email").
					From("users").
					Build()

				Expect(query).To(Equal("SELECT id, name, email FROM users"))
				Expect(args).To(BeEmpty())
			})

			It("should build SELECT with WHERE condition", func() {
				query, args := qb.Select("id", "name").
					From("users").
					Where("email", "test@example.com").
					Build()

				Expect(query).To(Equal("SELECT id, name FROM users WHERE email = $1"))
				Expect(args).To(HaveLen(1))
				Expect(args[0]).To(Equal("test@example.com"))
			})

			It("should build SELECT with multiple WHERE conditions", func() {
				query, args := qb.Select("id", "name").
					From("users").
					Where("email", "test@example.com").
					Where("active", true).
					Build()

				Expect(query).To(Equal("SELECT id, name FROM users WHERE email = $1 AND active = $2"))
				Expect(args).To(HaveLen(2))
				Expect(args[0]).To(Equal("test@example.com"))
				Expect(args[1]).To(Equal(true))
			})

			It("should build SELECT with WHERE IN condition", func() {
				query, args := qb.Select("id", "name").
					From("users").
					WhereIn("id", 1, 2, 3).
					Build()

				Expect(query).To(Equal("SELECT id, name FROM users WHERE id IN ($1, $2, $3)"))
				Expect(args).To(HaveLen(3))
				Expect(args[0]).To(Equal(1))
				Expect(args[1]).To(Equal(2))
				Expect(args[2]).To(Equal(3))
			})

			It("should build SELECT with WHERE IN and additional WHERE", func() {
				query, args := qb.Select("id", "name").
					From("users").
					WhereIn("status", "active", "pending").
					Where("verified", true).
					Build()

				Expect(query).To(Equal("SELECT id, name FROM users WHERE status IN ($1, $2) AND verified = $3"))
				Expect(args).To(HaveLen(3))
				Expect(args[0]).To(Equal("active"))
				Expect(args[1]).To(Equal("pending"))
				Expect(args[2]).To(Equal(true))
			})

			It("should handle empty WHERE IN gracefully", func() {
				query, args := qb.Select("id", "name").
					From("users").
					WhereIn("id"). // No values
					Build()

				Expect(query).To(Equal("SELECT id, name FROM users"))
				Expect(args).To(BeEmpty())
			})
		})

		Context("SELECT with QueryFields", func() {
			It("should build SELECT with QueryField objects", func() {
				fields := []QueryField{
					NewQueryField("id"),
					NewQueryField("created_at", "::timestamptz"),
					NewQueryField("data", "::jsonb"),
				}

				query, args := qb.SelectFields(fields...).
					From("users").
					Build()

				Expect(query).To(Equal("SELECT id, created_at::timestamptz, data::jsonb FROM users"))
				Expect(args).To(BeEmpty())
			})
		})

		Context("SELECT with RETURNING", func() {
			It("should build SELECT with RETURNING clause", func() {
				query, args := qb.Select("*").
					From("users").
					Where("id", 123).
					Returning("id", "updated_at").
					Build()

				Expect(query).To(Equal("SELECT * FROM users WHERE id = $1 RETURNING id, updated_at"))
				Expect(args).To(HaveLen(1))
				Expect(args[0]).To(Equal(123))
			})
		})
	})

	Describe("INSERT queries", func() {
		Context("simple INSERT", func() {
			It("should build basic INSERT query", func() {
				query, args := qb.InsertInto("users").
					ValueFields("id", "email", "name").
					Values(1, "test@example.com", "Test User").
					Build()

				Expect(query).To(Equal("INSERT INTO users (id, email, name) VALUES ($1, $2, $3)"))
				Expect(args).To(HaveLen(3))
				Expect(args[0]).To(Equal(1))
				Expect(args[1]).To(Equal("test@example.com"))
				Expect(args[2]).To(Equal("Test User"))
			})

			It("should build INSERT with RETURNING", func() {
				query, args := qb.InsertInto("users").
					ValueFields("email", "name").
					Values("test@example.com", "Test User").
					Returning("id", "created_at").
					Build()

				Expect(query).To(Equal("INSERT INTO users (email, name) VALUES ($1, $2) RETURNING id, created_at"))
				Expect(args).To(HaveLen(2))
				Expect(args[0]).To(Equal("test@example.com"))
				Expect(args[1]).To(Equal("Test User"))
			})
		})

	})

	Describe("UPDATE queries", func() {
		Context("simple UPDATE", func() {
			It("should build basic UPDATE query", func() {
				query, args := qb.Update("users").
					Set("email", "new@example.com").
					Set("name", "New Name").
					Where("id", 1).
					Build()

				Expect(query).To(Equal("UPDATE users SET email = $1, name = $2 WHERE id = $3"))
				Expect(args).To(HaveLen(3))
				Expect(args[0]).To(Equal("new@example.com"))
				Expect(args[1]).To(Equal("New Name"))
				Expect(args[2]).To(Equal(1))
			})

			It("should build UPDATE with RETURNING", func() {
				query, args := qb.Update("users").
					Set("email", "new@example.com").
					Where("id", 1).
					Returning("id", "updated_at").
					Build()

				Expect(query).To(Equal("UPDATE users SET email = $1 WHERE id = $2 RETURNING id, updated_at"))
				Expect(args).To(HaveLen(2))
				Expect(args[0]).To(Equal("new@example.com"))
				Expect(args[1]).To(Equal(1))
			})

			It("should build UPDATE with multiple WHERE conditions", func() {
				query, args := qb.Update("users").
					Set("active", false).
					Where("email", "test@example.com").
					Where("verified", true).
					Build()

				Expect(query).To(Equal("UPDATE users SET active = $1 WHERE email = $2 AND verified = $3"))
				Expect(args).To(HaveLen(3))
				Expect(args[0]).To(Equal(false))
				Expect(args[1]).To(Equal("test@example.com"))
				Expect(args[2]).To(Equal(true))
			})
		})
	})

	Describe("parameter indexing", func() {
		It("should maintain proper parameter indexing across multiple operations", func() {
			query, args := qb.Select("id", "name").
				From("users").
				Where("status", "active").        // $1
				Where("verified", true).          // $2
				WhereIn("role", "admin", "user"). // $3, $4
				Build()

			Expect(query).To(Equal("SELECT id, name FROM users WHERE status = $1 AND verified = $2 AND role IN ($3, $4)"))
			Expect(args).To(HaveLen(4))
			Expect(args[0]).To(Equal("active"))
			Expect(args[1]).To(Equal(true))
			Expect(args[2]).To(Equal("admin"))
			Expect(args[3]).To(Equal("user"))
		})

		It("should maintain parameter indexing in UPDATE queries", func() {
			query, args := qb.Update("users").
				Set("name", "Updated Name").     // $1
				Set("email", "new@example.com"). // $2
				Where("id", 123).                // $3
				Where("active", true).           // $4
				Build()

			Expect(query).To(Equal("UPDATE users SET name = $1, email = $2 WHERE id = $3 AND active = $4"))
			Expect(args).To(HaveLen(4))
			Expect(args[0]).To(Equal("Updated Name"))
			Expect(args[1]).To(Equal("new@example.com"))
			Expect(args[2]).To(Equal(123))
			Expect(args[3]).To(Equal(true))
		})
	})
})
