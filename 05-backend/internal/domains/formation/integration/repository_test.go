package integration

import (
	"context"
	"shvdg/crazed-conquerer/internal/domains/formation/domain"
	infra "shvdg/crazed-conquerer/internal/domains/formation/infrastructure"
	"shvdg/crazed-conquerer/internal/shared/contexts"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/testing"
	"shvdg/crazed-conquerer/internal/shared/testing/shared"

	"github.com/jackc/pgx/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Formation Repository", Ordered, func() {
	var err error
	var transaction pgx.Tx
	var ctx context.Context

	var suite *testing.Suite
	var formationRepo *infra.FormationRepositoryImpl

	BeforeAll(func() {
		suite = shared.GetSharedSuite()
		transaction, err = suite.StartTransaction()
		Expect(err).ToNot(HaveOccurred(), "failed to start transaction")

		ctx = contexts.SetTransaction(suite.Context, transaction)
		formationRepo = infra.NewFormationRepositoryImpl(suite.Database)
	})

	AfterAll(func() {
		err := transaction.Rollback(ctx)
		Expect(err).ToNot(HaveOccurred(), "failed to rollback transaction")
	})

	Context("When one formation is created", func() {
		var formation *domain.FormationEntity

		BeforeAll(func() {
			formation = domain.NewFormationEntity().
				WithDefaults().
				WithRowsFromJson(rowsJson).
				Build()
		})

		It("should successfully store the formation in the database", func() {
			err := formationRepo.Create(ctx, formation)
			Expect(err).ToNot(HaveOccurred(), "failed to create formation")

			fields := []string{infra.FieldId}
			values := []any{formation.GetId()}

			count, err := database.Count(ctx, suite.Database, infra.TableName, fields, values)
			Expect(err).ToNot(HaveOccurred(), "failed to count formations")
			Expect(count).To(Equal(1), "expected 1 formation to be created")
		})
	})
})
