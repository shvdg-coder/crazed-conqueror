package integration

import (
	"context"
	characterDomain "shvdg/crazed-conquerer/internal/domains/character/domain"
	characterInfra "shvdg/crazed-conquerer/internal/domains/character/infrastructure"
	"shvdg/crazed-conquerer/internal/domains/user-character/domain"
	"shvdg/crazed-conquerer/internal/domains/user-character/infrastructure"
	userDomain "shvdg/crazed-conquerer/internal/domains/user/domain"
	userInfra "shvdg/crazed-conquerer/internal/domains/user/infrastructure"
	"shvdg/crazed-conquerer/internal/shared/contexts"
	"shvdg/crazed-conquerer/internal/shared/testing"
	"shvdg/crazed-conquerer/internal/shared/testing/shared"

	"github.com/jackc/pgx/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Character Repository", Ordered, func() {
	var err error
	var transaction pgx.Tx
	var ctx context.Context

	var suite *testing.Suite
	var userRepo *userInfra.UserRepositoryImpl
	var characterRepo *characterInfra.CharacterRepositoryImpl
	var userCharacterRepo *infrastructure.UserCharacterRepositoryImpl

	// Test entities
	var dummyUser *userDomain.UserEntity
	var dummyCharacter *characterDomain.CharacterEntity
	var dummyUserCharacter *domain.UserCharacterEntity

	BeforeAll(func() {
		suite = shared.GetSharedSuite()
		transaction, err = suite.StartTransaction()
		Expect(err).ToNot(HaveOccurred(), "failed to start transaction")

		ctx = contexts.SetTransaction(suite.Context, transaction)

		userCharacterRepo = infrastructure.NewUserCharacterRepositoryImpl(suite.Database)
		userRepo = userInfra.NewUserRepositoryImpl(suite.Database)
		characterRepo = characterInfra.NewCharacterRepositoryImpl(suite.Database)

		By("Creating test user")
		dummyUser = userDomain.NewUserEntity().WithDefaults().Build()
		err = userRepo.Create(ctx, dummyUser)
		Expect(err).ToNot(HaveOccurred(), "failed to create test user")

		By("Creating test character")
		dummyCharacter = characterDomain.NewCharacterEntity().WithDefaults().Build()
		err = characterRepo.Create(ctx, dummyCharacter)
		Expect(err).ToNot(HaveOccurred(), "failed to create test character")

		By("Creating user-character association")
		dummyUserCharacter = domain.NewUserCharacterEntity().
			WithUserID(dummyUser.GetId()).
			WithCharacterID(dummyCharacter.GetId()).
			Build()
	})

	AfterAll(func() {
		err := transaction.Rollback(ctx)
		Expect(err).ToNot(HaveOccurred(), "failed to rollback transaction")
	})

	Context("When creating a user-character association", func() {
		It("should successfully store the association in the database", func() {
			err := userCharacterRepo.Create(ctx, dummyUserCharacter)
			Expect(err).ToNot(HaveOccurred(), "failed to create user-character association")
		})
	})

	Context("When retrieving user characters by user ID", func() {
		It("should return user character associations for the given user ID", func() {
			userCharacters, err := userCharacterRepo.GetByUserID(ctx, dummyUser.GetId())
			Expect(err).ToNot(HaveOccurred(), "failed to get user characters by user ID")
			Expect(userCharacters).ToNot(BeNil(), "expected to find user characters")
			Expect(userCharacters).To(HaveLen(1), "expected exactly one user character association")

			userCharacter := userCharacters[0]
			Expect(userCharacter.GetUserId()).To(Equal(dummyUser.GetId()), "user ID should match")
			Expect(userCharacter.GetCharacterId()).To(Equal(dummyCharacter.GetId()), "character ID should match")
		})
	})

	Context("When retrieving user characters by character ID", func() {
		It("should return user character associations for the given character ID", func() {
			userCharacters, err := userCharacterRepo.GetByCharacterID(ctx, dummyCharacter.GetId())
			Expect(err).ToNot(HaveOccurred(), "failed to get user characters by character ID")
			Expect(userCharacters).ToNot(BeNil(), "expected to find user characters")
			Expect(userCharacters).To(HaveLen(1), "expected exactly one user character association")

			userCharacter := userCharacters[0]
			Expect(userCharacter.GetUserId()).To(Equal(dummyUser.GetId()), "user ID should match")
			Expect(userCharacter.GetCharacterId()).To(Equal(dummyCharacter.GetId()), "character ID should match")
		})
	})

	Context("When deleting a user-character association", func() {
		It("should successfully remove the association from the database", func() {
			err := userCharacterRepo.Delete(ctx, dummyUserCharacter)
			Expect(err).ToNot(HaveOccurred(), "failed to delete user-character association")

			userCharacters, err := userCharacterRepo.GetByUserID(ctx, dummyUser.GetId())
			Expect(err).ToNot(HaveOccurred(), "failed to get user characters")
			Expect(userCharacters).To(BeEmpty(), "expected user-character association to be deleted")
		})
	})
})
