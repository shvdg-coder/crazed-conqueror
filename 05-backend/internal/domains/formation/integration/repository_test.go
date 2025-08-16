package integration

import (
	"context"
	"shvdg/crazed-conquerer/internal/domains/formation/domain"
	infra "shvdg/crazed-conquerer/internal/domains/formation/infrastructure"
	"shvdg/crazed-conquerer/internal/shared/contexts"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/sql"
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
			formation = domain.NewFormationEntity().WithDefaults().
				WithRowsFromJson(simpleRowsJson).
				Build()
		})

		It("should successfully store the formation in the database", func() {
			err := formationRepo.Create(ctx, formation)
			Expect(err).ToNot(HaveOccurred(), "failed to create formation")

			query, args := sql.NewQuery().Count().From(infra.TableName).Where(infra.FieldId, formation.GetId()).Build()
			count, err := database.QueryOne(ctx, suite.Database, query, args, database.ScanInt)

			Expect(err).ToNot(HaveOccurred(), "failed to count formations")
			Expect(count).To(Equal(1), "expected 1 formation to be created")
		})
	})

	Context("When retrieving formation by ID", func() {
		BeforeAll(func() {
			formation := domain.NewFormationEntity().WithDefaults().
				WithId("find-me-123").
				WithRowsFromJson(simpleRowsJson).
				Build()
			err := formationRepo.Create(ctx, formation)
			Expect(err).ToNot(HaveOccurred(), "failed to create formation")
		})

		It("should return the correct formation", func() {
			foundFormation, err := formationRepo.GetById(ctx, "find-me-123")
			Expect(err).ToNot(HaveOccurred(), "failed to get formation by ID")
			Expect(foundFormation).ToNot(BeNil(), "expected to find a formation")
			Expect(foundFormation.GetId()).To(Equal("find-me-123"))
			Expect(foundFormation.GetRows()).To(HaveLen(1), "expected 1 row")
			Expect(foundFormation.GetRows()[0].GetColumns()).To(HaveLen(2), "expected 2 columns")

			firstColumn := foundFormation.GetRows()[0].GetColumns()[0]
			Expect(firstColumn.GetPositionX()).To(Equal(int32(0)))
			Expect(firstColumn.GetPositionY()).To(Equal(int32(0)))
			Expect(firstColumn.GetUnitId()).To(Equal("unit_1"))
		})
	})
})
