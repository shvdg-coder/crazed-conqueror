package integration

import (
	"context"
	"shvdg/crazed-conquerer/internal/shared/contexts"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/testing"
	"shvdg/crazed-conquerer/internal/shared/testing/shared"
	"shvdg/crazed-conquerer/internal/user/domain"
	infra "shvdg/crazed-conquerer/internal/user/infrastructure"

	"github.com/jackc/pgx/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Repository", Ordered, func() {
	var err error
	var transaction pgx.Tx
	var ctx context.Context

	var suite *testing.Suite
	var userRepo *infra.UserRepositoryImpl

	BeforeAll(func() {
		suite = shared.GetSharedSuite()
		transaction, err = suite.StartTransaction()
		Expect(err).ToNot(HaveOccurred(), "failed to start transaction")

		ctx = contexts.SetTransaction(suite.Context, transaction)
		userRepo = infra.NewUserRepositoryImpl(suite.Database)
	})

	AfterAll(func() {
		err := transaction.Rollback(ctx)
		Expect(err).ToNot(HaveOccurred(), "failed to rollback transaction")
	})

	Context("When one user is created", func() {
		var user *domain.UserEntity

		BeforeAll(func() {
			user = domain.NewUserEntity().WithDefaults().Build()
		})

		It("should successfully store the user in the database", func() {
			err := userRepo.Create(ctx, user)
			Expect(err).ToNot(HaveOccurred(), "failed to create user")

			fields := []string{infra.FieldId, infra.FieldEmail, infra.FieldDisplayName}
			values := []any{user.GetId(), user.GetEmail(), user.GetDisplayName()}

			count, err := database.Count(ctx, suite.Database, infra.TableName, fields, values)
			Expect(err).ToNot(HaveOccurred(), "failed to count users")
			Expect(count).To(Equal(1), "expected 1 user to be created")
		})

	})

	Context("When retrieving a user by email", func() {
		var user *domain.UserEntity

		BeforeAll(func() {
			user = domain.NewUserEntity().WithDefaults().
				WithEmail("findme@domain.com").
				Build()
			err := userRepo.Create(ctx, user)
			Expect(err).ToNot(HaveOccurred(), "failed to create user")
		})

		It("should return the correct user", func() {
			foundUser, err := userRepo.GetByEmail(ctx, "findme@domain.com")
			Expect(err).ToNot(HaveOccurred(), "failed to get user by email")
			Expect(foundUser).ToNot(BeNil(), "expected to find a user")
			Expect(foundUser.GetId()).To(Equal(user.GetId()))
		})

		It("should return nil when no user is found", func() {
			foundUser, err := userRepo.GetByEmail(ctx, "nosuchuser@domain.com")
			Expect(err).To(HaveOccurred(), "expected an error for no user found")
			Expect(foundUser).To(BeNil(), "expected not to find a user")
		})
	})

	Context("When one user is updated", func() {
		var user *domain.UserEntity

		BeforeAll(func() {
			user = domain.NewUserEntity().WithDefaults().Build()
			err := userRepo.Create(ctx, user)
			Expect(err).ToNot(HaveOccurred(), "failed to create user")
		})

		It("should successfully update the user in the database", func() {
			user.DisplayName = "A New Name"
			err := userRepo.Update(ctx, user)
			Expect(err).ToNot(HaveOccurred(), "failed to update user")

			updatedUser, err := userRepo.GetByEmail(ctx, user.GetEmail())
			Expect(err).ToNot(HaveOccurred(), "failed to retrieve updated user")
			Expect(updatedUser.GetDisplayName()).To(Equal("A New Name"))
		})
	})

	Context("When a user is upserted", func() {
		var user *domain.UserEntity

		BeforeAll(func() {
			user = domain.NewUserEntity().WithDefaults().
				WithEmail("upsert@example.com").
				WithDisplayName("Initial Name").
				Build()
		})

		It("should insert the user if they do not exist", func() {
			err := userRepo.Upsert(ctx, user)
			Expect(err).ToNot(HaveOccurred(), "failed to upsert new user")

			retrieved, err := userRepo.GetByEmail(ctx, "upsert@example.com")
			Expect(err).ToNot(HaveOccurred())
			Expect(retrieved).ToNot(BeNil())
			Expect(retrieved.GetDisplayName()).To(Equal("Initial Name"))
		})

		It("should update the user if they already exist", func() {
			user.DisplayName = "Updated Name"
			err := userRepo.Upsert(ctx, user)
			Expect(err).ToNot(HaveOccurred(), "failed to upsert existing user")

			retrieved, err := userRepo.GetByEmail(ctx, "upsert@example.com")
			Expect(err).ToNot(HaveOccurred())
			Expect(retrieved).ToNot(BeNil())
			Expect(retrieved.GetDisplayName()).To(Equal("Updated Name"))
		})
	})

	Context("When one user is deleted", func() {
		var user *domain.UserEntity

		BeforeAll(func() {
			user = domain.NewUserEntity().WithDefaults().Build()
			err := userRepo.Create(ctx, user)
			Expect(err).ToNot(HaveOccurred(), "failed to create user")
		})

		It("should successfully remove the user from the database", func() {
			err := userRepo.Delete(ctx, user)
			Expect(err).ToNot(HaveOccurred(), "failed to delete user")

			count, err := database.Count(ctx, suite.Database, infra.TableName, []string{infra.FieldId}, []any{user.GetId()})
			Expect(err).ToNot(HaveOccurred(), "failed to count users")
			Expect(count).To(BeZero(), "expected user to be deleted")
		})
	})
})
