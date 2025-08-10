package integration

import (
	"context"
	"shvdg/crazed-conquerer/internal/domains/character-unit/domain"
	infra "shvdg/crazed-conquerer/internal/domains/character-unit/infrastructure"
	characterDomain "shvdg/crazed-conquerer/internal/domains/character/domain"
	characterInfra "shvdg/crazed-conquerer/internal/domains/character/infrastructure"
	unitDomain "shvdg/crazed-conquerer/internal/domains/unit/domain"
	unitInfra "shvdg/crazed-conquerer/internal/domains/unit/infrastructure"
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
	var characterRepo *characterInfra.CharacterRepositoryImpl
	var unitRepo *unitInfra.UnitRepositoryImpl
	var characterUnitRepo *infra.CharacterUnitRepositoryImpl

	// Test entities
	var dummyCharacter *characterDomain.CharacterEntity
	var dummyUnit *unitDomain.UnitEntity
	var dummyCharacterUnit *domain.CharacterUnitEntity

	BeforeAll(func() {
		suite = shared.GetSharedSuite()
		transaction, err = suite.StartTransaction()
		Expect(err).ToNot(HaveOccurred(), "failed to start transaction")

		ctx = contexts.SetTransaction(suite.Context, transaction)
		characterUnitRepo = infra.NewCharacterUnitRepositoryImpl(suite.Database)
		characterRepo = characterInfra.NewCharacterRepositoryImpl(suite.Database)
		unitRepo = unitInfra.NewUnitRepositoryImpl(suite.Database)

		By("Creating test character")
		dummyCharacter = characterDomain.NewCharacterEntity().WithDefaults().Build()
		err = characterRepo.Create(ctx, dummyCharacter)
		Expect(err).ToNot(HaveOccurred(), "failed to create test character")

		By("Creating test unit")
		dummyUnit = unitDomain.NewUnitEntity().WithDefaults().Build()
		err = unitRepo.Create(ctx, dummyUnit)
		Expect(err).ToNot(HaveOccurred(), "failed to create test unit")

		By("Creating character-unit association")
		dummyCharacterUnit = domain.NewCharacterUnitEntity().
			WithCharacterId(dummyCharacter.GetId()).
			WithUnitId(dummyUnit.GetId()).
			Build()
	})

	AfterAll(func() {
		err := transaction.Rollback(ctx)
		Expect(err).ToNot(HaveOccurred(), "failed to rollback transaction")
	})

	Context("When one character unit is created", func() {
		It("should successfully store the character unit in the database", func() {
			err := characterUnitRepo.Create(ctx, dummyCharacterUnit)
			Expect(err).ToNot(HaveOccurred(), "failed to create character unit")

			fields := []string{infra.FieldCharacterId, infra.FieldUnitId}
			values := []any{dummyCharacterUnit.GetCharacterId(), dummyCharacterUnit.GetUnitId()}

			count, err := database.Count(ctx, suite.Database, infra.TableName, fields, values)
			Expect(err).ToNot(HaveOccurred(), "failed to count character units")
			Expect(count).To(Equal(1), "expected 1 character unit to be created")
		})
	})

	Context("When retrieving character units by character ID", func() {
		It("should return character unit associations for the given character ID", func() {
			characterUnits, err := characterUnitRepo.GetByCharacterID(ctx, dummyCharacter.GetId())
			Expect(err).ToNot(HaveOccurred(), "failed to get character units by character ID")
			Expect(characterUnits).ToNot(BeNil(), "expected to find character units")
			Expect(characterUnits).To(HaveLen(1), "expected exactly one character unit association")

			characterUnit := characterUnits[0]
			Expect(characterUnit.GetCharacterId()).To(Equal(dummyCharacter.GetId()), "character ID should match")
			Expect(characterUnit.GetUnitId()).To(Equal(dummyUnit.GetId()), "unit ID should match")
		})

		It("should return empty slice when no associations are found", func() {
			nonExistentCharacter := characterDomain.NewCharacterEntity().WithDefaults().Build()

			characterUnits, err := characterUnitRepo.GetByCharacterID(ctx, nonExistentCharacter.GetId())
			Expect(err).ToNot(HaveOccurred(), "should not error when no associations found")
			Expect(characterUnits).To(BeEmpty(), "expected empty slice for non-existent character")
		})
	})

	Context("When retrieving character units by unit ID", func() {
		It("should return character unit associations for the given unit ID", func() {
			characterUnits, err := characterUnitRepo.GetByUnitID(ctx, dummyUnit.GetId())
			Expect(err).ToNot(HaveOccurred(), "failed to get character units by unit ID")
			Expect(characterUnits).ToNot(BeNil(), "expected to find character units")
			Expect(characterUnits).To(HaveLen(1), "expected exactly one character unit association")

			characterUnit := characterUnits[0]
			Expect(characterUnit.GetCharacterId()).To(Equal(dummyCharacter.GetId()), "character ID should match")
			Expect(characterUnit.GetUnitId()).To(Equal(dummyUnit.GetId()), "unit ID should match")
		})

		It("should return empty slice when no associations are found", func() {
			nonExistentUnit := unitDomain.NewUnitEntity().WithDefaults().Build()

			characterUnits, err := characterUnitRepo.GetByUnitID(ctx, nonExistentUnit.GetId())
			Expect(err).ToNot(HaveOccurred(), "should not error when no associations found")
			Expect(characterUnits).To(BeEmpty(), "expected empty slice for non-existent unit")
		})
	})

	Context("When one character unit is deleted", func() {
		It("should successfully remove the character unit from the database", func() {
			err := characterUnitRepo.Delete(ctx, dummyCharacterUnit)
			Expect(err).ToNot(HaveOccurred(), "failed to delete character unit")

			fields := []string{infra.FieldCharacterId, infra.FieldUnitId}
			values := []any{dummyCharacterUnit.GetCharacterId(), dummyCharacterUnit.GetUnitId()}
			count, err := database.Count(ctx, suite.Database, infra.TableName, fields, values)
			Expect(err).ToNot(HaveOccurred(), "failed to count character units")
			Expect(count).To(BeZero(), "expected character unit to be deleted")
		})
	})
})
