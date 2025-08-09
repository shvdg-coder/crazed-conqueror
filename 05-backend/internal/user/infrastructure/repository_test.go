package infrastructure

import (
	"context"
	"shvdg/crazed-conquerer/internal/shared/contexts"
	"shvdg/crazed-conquerer/internal/shared/testing"
	"shvdg/crazed-conquerer/internal/user/domain"

	"github.com/jackc/pgx/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Repository", Ordered, func() {
	var err error
	var transaction pgx.Tx
	var ctx context.Context

	var suite *testing.Suite
	var userRepo *UserRepositoryImpl

	BeforeAll(func() {
		suite = testing.GetSharedSuite()
		transaction, err = suite.StartTransaction()
		Expect(err).ToNot(HaveOccurred(), "failed to start transaction")

		ctx = contexts.SetTransaction(suite.Context, transaction)
		userRepo = NewUserRepositoryImpl(suite.Database)
	})

	AfterAll(func() {
		err := transaction.Rollback(ctx)
		Expect(err).ToNot(HaveOccurred(), "failed to rollback transaction")
	})

	Context("When one user is created", func() {
		var user *domain.UserEntity

		BeforeAll(func() {
			user = domain.NewUserEntity().WithDefaults().
				WithEmail("hero@domain.com").
				WithPassword("securePa$$123").
				WithDisplayName("Hero").
				Build()
		})

		It("should successfully store the user in the database", func() {
			err := userRepo.Create(ctx, user)
			Expect(err).ToNot(HaveOccurred(), "failed to create user")
		})

		AfterAll(func() {
			_, err := transaction.Exec(ctx, "DELETE FROM users")
			Expect(err).ToNot(HaveOccurred(), "failed to cleanup users")
		})
	})
})
