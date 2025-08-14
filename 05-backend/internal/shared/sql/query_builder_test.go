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

	Describe("COUNT queries", func() {
		Context("basic COUNT", func() {
			It("should build simple COUNT query", func() {
				query, args := qb.Count().
					From("users").
					Build()

				Expect(query).To(Equal("SELECT COUNT(*) FROM users"))
				Expect(args).To(BeEmpty())
			})

			It("should build COUNT with multiple WHERE conditions", func() {
				query, args := qb.Count().
					From("users").
					Where("active", true).
					Where("role", "admin").
					Build()

				Expect(query).To(Equal("SELECT COUNT(*) FROM users WHERE active = $1 AND role = $2"))
				Expect(args).To(Equal([]any{true, "admin"}))
			})
		})
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
				Expect(args).To(Equal([]any{"test@example.com"}))
			})

			It("should build SELECT with multiple WHERE conditions", func() {
				query, args := qb.Select("id", "name").
					From("users").
					Where("email", "test@example.com").
					Where("active", true).
					Build()

				Expect(query).To(Equal("SELECT id, name FROM users WHERE email = $1 AND active = $2"))
				Expect(args).To(Equal([]any{"test@example.com", true}))
			})

			It("should build SELECT with WHERE IN condition", func() {
				query, args := qb.Select("id", "name").
					From("users").
					WhereIn("id", 1, 2, 3).
					Build()

				Expect(query).To(Equal("SELECT id, name FROM users WHERE id IN ($1, $2, $3)"))
				Expect(args).To(Equal([]any{1, 2, 3}))
			})

			It("should build SELECT with WHERE IN and additional WHERE", func() {
				query, args := qb.Select("id", "name").
					From("users").
					WhereIn("status", "active", "pending").
					Where("verified", true).
					Build()

				Expect(query).To(Equal("SELECT id, name FROM users WHERE status IN ($1, $2) AND verified = $3"))
				Expect(args).To(Equal([]any{"active", "pending", true}))
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
	})

	Describe("INSERT queries", func() {
		Context("simple INSERT", func() {
			It("should build basic INSERT query", func() {
				query, args := qb.InsertInto("users").
					InsertFields("id", "email", "name").
					Values(1, "test@example.com", "Test User").
					Build()

				Expect(query).To(Equal("INSERT INTO users (id, email, name) VALUES ($1, $2, $3)"))
				Expect(args).To(Equal([]any{1, "test@example.com", "Test User"}))
			})
		})

		Context("batch INSERT", func() {
			It("should build batch INSERT with single argument set", func() {
				argumentSets := [][]any{
					{"user1", "char1"},
				}

				query, batchArgs := qb.InsertInto("user_characters").
					InsertFields("user_id", "character_id").
					BatchValues(argumentSets).
					BuildBatch()

				Expect(query).To(Equal("INSERT INTO user_characters (user_id, character_id) VALUES ($1, $2)"))
				Expect(batchArgs).To(Equal([][]any{{"user1", "char1"}}))
			})

			It("should build batch INSERT with multiple argument sets", func() {
				argumentSets := [][]any{
					{"user1", "char1"},
					{"user2", "char2"},
					{"user3", "char3"},
				}

				query, batchArgs := qb.InsertInto("user_characters").
					InsertFields("user_id", "character_id").
					BatchValues(argumentSets).
					BuildBatch()

				Expect(query).To(Equal("INSERT INTO user_characters (user_id, character_id) VALUES ($1, $2)"))
				Expect(batchArgs).To(Equal([][]any{
					{"user1", "char1"},
					{"user2", "char2"},
					{"user3", "char3"},
				}))
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
				Expect(args).To(Equal([]any{"new@example.com", "New Name", 1}))
			})

			It("should build UPDATE with multiple WHERE conditions", func() {
				query, args := qb.Update("users").
					Set("active", false).
					Where("email", "test@example.com").
					Where("verified", true).
					Build()

				Expect(query).To(Equal("UPDATE users SET active = $1 WHERE email = $2 AND verified = $3"))
				Expect(args).To(Equal([]any{false, "test@example.com", true}))
			})
		})

		Context("batch UPDATE", func() {
			It("should build batch UPDATE with single argument set", func() {
				argumentSets := [][]any{
					{"NewName1", "new1@email.com"},
				}

				query, batchArgs := qb.Update("users").
					BatchSets(argumentSets, "name", "email").
					Where("active", true).
					BuildBatch()

				Expect(query).To(Equal("UPDATE users SET name = $1, email = $2 WHERE active = $3"))
				Expect(batchArgs).To(Equal([][]any{{"NewName1", "new1@email.com"}}))
			})

			It("should build batch UPDATE with multiple argument sets", func() {
				argumentSets := [][]any{
					{"NewName1", "new1@email.com"},
					{"NewName2", "new2@email.com"},
					{"NewName3", "new3@email.com"},
				}

				query, batchArgs := qb.Update("users").
					BatchSets(argumentSets, "name", "email").
					Where("active", true).
					BuildBatch()

				Expect(query).To(Equal("UPDATE users SET name = $1, email = $2 WHERE active = $3"))
				Expect(batchArgs).To(Equal([][]any{
					{"NewName1", "new1@email.com"},
					{"NewName2", "new2@email.com"},
					{"NewName3", "new3@email.com"},
				}))
			})

			It("should build batch UPDATE with multiple WHERE conditions", func() {
				argumentSets := [][]any{
					{"NewName1"},
				}

				query, batchArgs := qb.Update("users").
					BatchSets(argumentSets, "name").
					Where("active", true).
					Where("verified", true).
					BuildBatch()

				Expect(query).To(Equal("UPDATE users SET name = $1 WHERE active = $2 AND verified = $3"))
				Expect(batchArgs).To(Equal([][]any{{"NewName1"}}))
			})
		})
	})

	Describe("UPSERT queries", func() {
		Context("simple UPSERT", func() {
			It("should build basic UPSERT with DO NOTHING", func() {
				query, args := qb.InsertInto("users").
					InsertFields("id", "email", "name").
					Values(1, "test@example.com", "Test User").
					OnConflict("id").
					DoNothing().
					Build()

				Expect(query).To(Equal("INSERT INTO users (id, email, name) VALUES ($1, $2, $3) ON CONFLICT (id) DO NOTHING"))
				Expect(args).To(Equal([]any{1, "test@example.com", "Test User"}))
			})

			It("should build basic UPSERT with DO UPDATE", func() {
				query, args := qb.InsertInto("users").
					InsertFields("id", "email", "name").
					Values(1, "test@example.com", "Test User").
					OnConflict("id").
					DoUpdate("email", "name").
					Build()

				Expect(query).To(Equal("INSERT INTO users (id, email, name) VALUES ($1, $2, $3) ON CONFLICT (id) DO UPDATE SET email = EXCLUDED.email, name = EXCLUDED.name"))
				Expect(args).To(Equal([]any{1, "test@example.com", "Test User"}))
			})
		})

		Context("batch UPSERT", func() {
			It("should build batch UPSERT with DO NOTHING", func() {
				argumentSets := [][]any{
					{"user1", "char1"},
					{"user2", "char2"},
				}

				query, batchArgs := qb.InsertInto("user_characters").
					InsertFields("user_id", "character_id").
					BatchUpsert(argumentSets, []string{"user_id", "character_id"}).
					BuildBatch()

				Expect(query).To(Equal("INSERT INTO user_characters (user_id, character_id) VALUES ($1, $2) ON CONFLICT (user_id, character_id) DO NOTHING"))
				Expect(batchArgs).To(Equal([][]any{
					{"user1", "char1"},
					{"user2", "char2"},
				}))
			})

			It("should build batch UPSERT with DO UPDATE", func() {
				argumentSets := [][]any{
					{"user1", "char1", 10},
					{"user2", "char2", 15},
				}

				query, batchArgs := qb.InsertInto("user_characters").
					InsertFields("user_id", "character_id", "level").
					BatchUpsert(argumentSets, []string{"user_id", "character_id"}, "level").
					BuildBatch()

				Expect(query).To(Equal("INSERT INTO user_characters (user_id, character_id, level) VALUES ($1, $2, $3) ON CONFLICT (user_id, character_id) DO UPDATE SET level = EXCLUDED.level"))
				Expect(batchArgs).To(Equal([][]any{
					{"user1", "char1", 10},
					{"user2", "char2", 15},
				}))
			})
		})
	})

	Describe("DELETE queries", func() {
		Context("simple DELETE", func() {
			It("should build basic DELETE query", func() {
				query, args := qb.DeleteFrom("users").
					Where("id", 123).
					Build()

				Expect(query).To(Equal("DELETE FROM users WHERE id = $1"))
				Expect(args).To(Equal([]any{123}))
			})

			It("should build DELETE with multiple WHERE conditions", func() {
				query, args := qb.DeleteFrom("users").
					Where("email", "test@example.com").
					Where("active", false).
					Build()

				Expect(query).To(Equal("DELETE FROM users WHERE email = $1 AND active = $2"))
				Expect(args).To(Equal([]any{"test@example.com", false}))
			})

			It("should build DELETE with WHERE IN", func() {
				query, args := qb.DeleteFrom("users").
					WhereIn("id", 1, 2, 3).
					Build()

				Expect(query).To(Equal("DELETE FROM users WHERE id IN ($1, $2, $3)"))
				Expect(args).To(Equal([]any{1, 2, 3}))
			})
		})
	})

	Describe("WhereTupleIn queries", func() {
		Context("tuple IN conditions", func() {
			It("should handle single tuple", func() {
				tuples := [][]any{
					{"user1", "char1"},
				}
				query, args := qb.Select("*").
					From("user_characters").
					WhereTupleIn(tuples, "user_id", "character_id").
					Build()

				Expect(query).To(Equal("SELECT * FROM user_characters WHERE (user_id, character_id) IN (($1, $2))"))
				Expect(args).To(Equal([]any{"user1", "char1"}))
			})

			It("should work with three-field tuples", func() {
				tuples := [][]any{
					{"user1", "char1", "slot1"},
					{"user2", "char2", "slot2"},
				}
				query, args := qb.Select("*").
					From("user_character_slots").
					WhereTupleIn(tuples, "user_id", "character_id", "slot_id").
					Build()

				Expect(query).To(Equal("SELECT * FROM user_character_slots WHERE (user_id, character_id, slot_id) IN (($1, $2, $3), ($4, $5, $6))"))
				Expect(args).To(Equal([]any{"user1", "char1", "slot1", "user2", "char2", "slot2"}))
			})
		})
	})
})
