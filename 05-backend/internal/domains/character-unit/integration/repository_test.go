package integration

import (
	"context"
	"shvdg/crazed-conquerer/internal/domains/character-unit/domain"
	infra "shvdg/crazed-conquerer/internal/domains/character-unit/infrastructure"
	"shvdg/crazed-conquerer/internal/shared/contexts"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/testing"
	"shvdg/crazed-conquerer/internal/shared/testing/shared"

	"github.com/jackc/pgx/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CharacterUnit Repository", Ordered, func() {
	var err error
	var transaction pgx.Tx
	var ctx context.Context

	var suite *testing.Suite
	var characterUnitRepo *infra.CharacterUnitRepositoryImpl

	BeforeAll(func() {
		suite = shared.GetSharedSuite()
		transaction, err = suite.StartTransaction()
		Expect(err).ToNot(HaveOccurred(), "failed to start transaction")

		ctx = contexts.SetTransaction(suite.Context, transaction)
		characterUnitRepo = infra.NewCharacterUnitRepositoryImpl(suite.Database)
	})

	AfterAll(func() {
		err := transaction.Rollback(ctx)
		Expect(err).ToNot(HaveOccurred(), "failed to rollback transaction")
	})

	Context("When one character unit is created", func() {
		var characterUnit *domain.CharacterUnitEntity

		BeforeAll(func() {
			characterUnit = domain.NewCharacterUnitEntity().WithDefaults().Build()
		})

		It("should successfully store the character unit in the database", func() {
			err := characterUnitRepo.Create(ctx, characterUnit)
			Expect(err).ToNot(HaveOccurred(), "failed to create character unit")

			fields := []string{infra.FieldCharacterId, infra.FieldUnitId}
			values := []any{characterUnit.GetCharacterId(), characterUnit.GetUnitId()}

			count, err := database.Count(ctx, suite.Database, infra.TableName, fields, values)
			Expect(err).ToNot(HaveOccurred(), "failed to count character units")
			Expect(count).To(Equal(1), "expected 1 character unit to be created")
		})
	})

	Context("When one character unit is deleted", func() {
		var characterUnit *domain.CharacterUnitEntity

		BeforeAll(func() {
			characterUnit = domain.NewCharacterUnitEntity().WithDefaults().Build()
			err := characterUnitRepo.Create(ctx, characterUnit)
			Expect(err).ToNot(HaveOccurred(), "failed to create character unit")
		})

		It("should successfully remove the character unit from the database", func() {
			err := characterUnitRepo.Delete(ctx, characterUnit)
			Expect(err).ToNot(HaveOccurred(), "failed to delete character unit")

			fields := []string{infra.FieldCharacterId, infra.FieldUnitId}
			values := []any{characterUnit.GetCharacterId(), characterUnit.GetUnitId()}
			count, err := database.Count(ctx, suite.Database, infra.TableName, fields, values)
			Expect(err).ToNot(HaveOccurred(), "failed to count character units")
			Expect(count).To(BeZero(), "expected character unit to be deleted")
		})
	})
})
