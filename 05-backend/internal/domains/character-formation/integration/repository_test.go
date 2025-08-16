package integration

import (
	"context"
	"shvdg/crazed-conquerer/internal/domains/character-formation/domain"
	infra "shvdg/crazed-conquerer/internal/domains/character-formation/infrastructure"
	characterDomain "shvdg/crazed-conquerer/internal/domains/character/domain"
	characterInfra "shvdg/crazed-conquerer/internal/domains/character/infrastructure"
	formationDomain "shvdg/crazed-conquerer/internal/domains/formation/domain"
	formationInfra "shvdg/crazed-conquerer/internal/domains/formation/infrastructure"
	"shvdg/crazed-conquerer/internal/shared/contexts"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/sql"
	"shvdg/crazed-conquerer/internal/shared/testing"
	"shvdg/crazed-conquerer/internal/shared/testing/shared"

	"github.com/jackc/pgx/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CharacterFormation Repository", Ordered, func() {
	var err error
	var transaction pgx.Tx
	var ctx context.Context

	var suite *testing.Suite
	var characterRepo *characterInfra.CharacterRepositoryImpl
	var formationRepo *formationInfra.FormationRepositoryImpl
	var characterFormationRepo *infra.CharacterFormationRepositoryImpl

	// Test entities
	var dummyCharacter *characterDomain.CharacterEntity
	var dummyFormation *formationDomain.FormationEntity
	var dummyCharacterFormation *domain.CharacterFormationEntity

	BeforeAll(func() {
		suite = shared.GetSharedSuite()
		transaction, err = suite.StartTransaction()
		Expect(err).ToNot(HaveOccurred(), "failed to start transaction")

		ctx = contexts.SetTransaction(suite.Context, transaction)
		characterFormationRepo = infra.NewCharacterFormationRepositoryImpl(suite.Database)
		characterRepo = characterInfra.NewCharacterRepositoryImpl(suite.Database)
		formationRepo = formationInfra.NewFormationRepositoryImpl(suite.Database)

		By("Creating test character")
		dummyCharacter = characterDomain.NewCharacterEntity().WithDefaults().Build()
		err = characterRepo.Create(ctx, dummyCharacter)
		Expect(err).ToNot(HaveOccurred(), "failed to create test character")

		By("Creating test formation")
		dummyFormation = formationDomain.NewFormationEntity().WithDefaults().Build()
		err = formationRepo.Create(ctx, dummyFormation)
		Expect(err).ToNot(HaveOccurred(), "failed to create test formation")

		By("Creating character-formation association")
		dummyCharacterFormation = domain.NewCharacterFormationEntity().
			WithCharacterId(dummyCharacter.GetId()).
			WithId(dummyFormation.GetId()).
			Build()
	})

	AfterAll(func() {
		err := transaction.Rollback(ctx)
		Expect(err).ToNot(HaveOccurred(), "failed to rollback transaction")
	})

	Context("When one character formation is created", func() {
		It("should successfully store the character formation in the database", func() {
			err := characterFormationRepo.Create(ctx, dummyCharacterFormation)
			Expect(err).ToNot(HaveOccurred(), "failed to create character formation")

			query, args := sql.NewQuery().Count().From(infra.TableName).
				Where(infra.FieldCharacterId, dummyCharacterFormation.GetCharacterId()).
				Where(infra.FieldFormationId, dummyCharacterFormation.GetFormationId()).Build()
			count, err := database.QueryOne(ctx, suite.Database, query, args, database.ScanInt)

			Expect(err).ToNot(HaveOccurred(), "failed to count character formations")
			Expect(count).To(Equal(1), "expected 1 character formation to be created")
		})
	})

	Context("When retrieving character formations by character ID", func() {
		It("should return character formation associations for the given character ID", func() {
			characterFormations, err := characterFormationRepo.GetByCharacterId(ctx, dummyCharacter.GetId())
			Expect(err).ToNot(HaveOccurred(), "failed to get character formations by character ID")
			Expect(characterFormations).ToNot(BeNil(), "expected to find character formations")
			Expect(characterFormations).To(HaveLen(1), "expected exactly one character formation association")

			characterFormation := characterFormations[0]
			Expect(characterFormation.GetCharacterId()).To(Equal(dummyCharacter.GetId()), "character ID should match")
			Expect(characterFormation.GetFormationId()).To(Equal(dummyFormation.GetId()), "formation ID should match")
		})
	})

	Context("When retrieving character formations by formation ID", func() {
		It("should return character formation associations for the given formation ID", func() {
			characterFormations, err := characterFormationRepo.GetByFormationId(ctx, dummyFormation.GetId())
			Expect(err).ToNot(HaveOccurred(), "failed to get character formations by formation ID")
			Expect(characterFormations).ToNot(BeNil(), "expected to find character formations")
			Expect(characterFormations).To(HaveLen(1), "expected exactly one character formation association")

			characterFormation := characterFormations[0]
			Expect(characterFormation.GetCharacterId()).To(Equal(dummyCharacter.GetId()), "character ID should match")
			Expect(characterFormation.GetFormationId()).To(Equal(dummyFormation.GetId()), "formation ID should match")
		})
	})

	Context("When one character formation is deleted", func() {
		It("should successfully remove the character formation from the database", func() {
			err := characterFormationRepo.Delete(ctx, dummyCharacterFormation)
			Expect(err).ToNot(HaveOccurred(), "failed to delete character formation")

			query, args := sql.NewQuery().Count().From(infra.TableName).
				Where(infra.FieldCharacterId, dummyCharacterFormation.GetCharacterId()).
				Where(infra.FieldFormationId, dummyCharacterFormation.GetFormationId()).Build()
			count, err := database.QueryOne(ctx, suite.Database, query, args, database.ScanInt)

			Expect(err).ToNot(HaveOccurred(), "failed to count character formations")
			Expect(count).To(BeZero(), "expected character formation to be deleted")
		})
	})
})
