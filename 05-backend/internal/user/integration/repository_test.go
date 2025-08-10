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
			user = domain.NewUserEntity().WithDefaults().
				WithEmail("hero@domain.com").
				WithPassword("securePa$$123").
				WithDisplayName("Hero").
				Build()
		})

		It("should successfully store the user in the database", func() {
			err := userRepo.Create(ctx, user)
			Expect(err).ToNot(HaveOccurred(), "failed to create user")

			fields := []string{infra.FieldId, infra.FieldEmail, infra.FieldDisplayName}
			values := []any{user.GetId(), user.GetEmail(), user.GetDisplayName()}

			// TODO: find a way to make the use of executor easier
			executor, cleanup, err := suite.Database.GetExecutor(ctx)
			Expect(err).ToNot(HaveOccurred(), "failed to get executor")
			defer cleanup()

			count, err := database.Count(ctx, executor, infra.TableName, fields, values)
			Expect(err).ToNot(HaveOccurred(), "failed to count users")
			Expect(count).To(Equal(1), "expected 1 user to be created")
		})

		AfterAll(func() {
			_, err := transaction.Exec(ctx, infra.DropTableQuery)
			Expect(err).ToNot(HaveOccurred(), "failed to cleanup users")
		})
	})
})
