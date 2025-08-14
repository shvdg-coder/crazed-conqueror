package integration

import (
	"context"
	"shvdg/crazed-conquerer/internal/domains/character/domain"
	infra "shvdg/crazed-conquerer/internal/domains/character/infrastructure"
	"shvdg/crazed-conquerer/internal/shared/contexts"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/testing"
	"shvdg/crazed-conquerer/internal/shared/testing/shared"

	"github.com/jackc/pgx/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CharacterEntity Repository", Ordered, func() {
	var err error
	var transaction pgx.Tx
	var ctx context.Context

	var suite *testing.Suite
	var characterRepo *infra.CharacterRepositoryImpl

	BeforeAll(func() {
		suite = shared.GetSharedSuite()
		transaction, err = suite.StartTransaction()
		Expect(err).ToNot(HaveOccurred(), "failed to start transaction")

		ctx = contexts.SetTransaction(suite.Context, transaction)
		characterRepo = infra.NewCharacterRepositoryImpl(suite.Database)
	})

	AfterAll(func() {
		err := transaction.Rollback(ctx)
		Expect(err).ToNot(HaveOccurred(), "failed to rollback transaction")
	})

	Context("When one character is created", func() {
		var character *domain.CharacterEntity

		BeforeAll(func() {
			character = domain.NewCharacterEntity().WithDefaults().Build()
		})

		It("should successfully store the character in the database", func() {
			err := characterRepo.Create(ctx, character)
			Expect(err).ToNot(HaveOccurred(), "failed to create character")

			fields := []string{infra.FieldId, infra.FieldName}
			values := []any{character.GetId(), character.GetName()}

			count, err := database.Count(ctx, suite.Database, infra.TableName, fields, values)
			Expect(err).ToNot(HaveOccurred(), "failed to count characters")
			Expect(count).To(Equal(1), "expected 1 character to be created")
		})
	})

	Context("When one character is updated", func() {
		var character *domain.CharacterEntity

		BeforeAll(func() {
			character = domain.NewCharacterEntity().WithDefaults().Build()
			err := characterRepo.Create(ctx, character)
			Expect(err).ToNot(HaveOccurred(), "failed to create character")
		})

		It("should successfully update the character in the database", func() {
			character.Name = "UpdatedName"
			err := characterRepo.Update(ctx, character)
			Expect(err).ToNot(HaveOccurred(), "failed to update character")

			updatedCharacter, err := characterRepo.GetById(ctx, character.GetId())
			Expect(err).ToNot(HaveOccurred(), "failed to retrieve updated character")
			Expect(updatedCharacter.GetName()).To(Equal("UpdatedName"))
		})
	})

	Context("When a character is upserted", func() {
		var character *domain.CharacterEntity

		BeforeAll(func() {
			character = domain.NewCharacterEntity().WithDefaults().
				WithId("upsert-test-id").
				WithName("InitialName").
				Build()
		})

		It("should insert the character if they do not exist", func() {
			err := characterRepo.Upsert(ctx, character)
			Expect(err).ToNot(HaveOccurred(), "failed to upsert new character")

			retrieved, err := characterRepo.GetById(ctx, "upsert-test-id")
			Expect(err).ToNot(HaveOccurred())
			Expect(retrieved).ToNot(BeNil())
			Expect(retrieved.GetName()).To(Equal("InitialName"))
		})

		It("should update the character if they already exist", func() {
			character.Name = "UpdatedName"
			err := characterRepo.Upsert(ctx, character)
			Expect(err).ToNot(HaveOccurred(), "failed to upsert existing character")

			retrieved, err := characterRepo.GetById(ctx, "upsert-test-id")
			Expect(err).ToNot(HaveOccurred())
			Expect(retrieved).ToNot(BeNil())
			Expect(retrieved.GetName()).To(Equal("UpdatedName"))
		})
	})

	Context("When one character is deleted", func() {
		var character *domain.CharacterEntity

		BeforeAll(func() {
			character = domain.NewCharacterEntity().WithDefaults().Build()
			err := characterRepo.Create(ctx, character)
			Expect(err).ToNot(HaveOccurred(), "failed to create character")
		})

		It("should successfully remove the character from the database", func() {
			err := characterRepo.Delete(ctx, character)
			Expect(err).ToNot(HaveOccurred(), "failed to delete character")

			count, err := database.Count(ctx, suite.Database, infra.TableName, []string{infra.FieldId}, []any{character.GetId()})
			Expect(err).ToNot(HaveOccurred(), "failed to count characters")
			Expect(count).To(BeZero(), "expected character to be deleted")
		})
	})
})
