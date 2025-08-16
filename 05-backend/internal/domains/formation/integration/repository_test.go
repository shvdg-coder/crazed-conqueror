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
				WithRowsFromJson(smallRowsJson).
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
				WithRowsFromJson(smallRowsJson).
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

			secondColumn := foundFormation.GetRows()[0].GetColumns()[1]
			Expect(secondColumn.GetPositionX()).To(Equal(int32(1)))
			Expect(secondColumn.GetPositionY()).To(Equal(int32(0)))
			Expect(secondColumn.GetUnitId()).To(Equal("unit_2"))
		})
	})

	Context("When one formation is updated", func() {
		var originalFormation, updatedFormation *domain.FormationEntity

		BeforeAll(func() {
			originalFormation = domain.NewFormationEntity().WithDefaults().
				WithId("update-test-123").
				WithRowsFromJson(smallRowsJson).
				Build()

			err := formationRepo.Create(ctx, originalFormation)
			Expect(err).ToNot(HaveOccurred(), "failed to create original formation")

			updatedFormation = domain.NewFormationEntity().WithDefaults().
				WithId("update-test-123").
				WithRowsFromJson(mediumRowsJson).
				Build()
		})

		It("should successfully update the formation rows", func() {
			err := formationRepo.Update(ctx, updatedFormation)
			Expect(err).ToNot(HaveOccurred(), "failed to update formation")

			foundFormation, err := formationRepo.GetById(ctx, "update-test-123")
			Expect(err).ToNot(HaveOccurred(), "failed to get updated formation")
			Expect(foundFormation).ToNot(BeNil(), "expected to find the updated formation")

			Expect(foundFormation.GetRows()).To(HaveLen(2), "expected 2 row")
			Expect(foundFormation.GetRows()[0].GetColumns()).To(HaveLen(2), "expected 2 columns")

			// Verify columns of the first row
			firstRow := foundFormation.GetRows()[0].GetColumns()
			Expect(firstRow[0].GetPositionX()).To(Equal(int32(0)))
			Expect(firstRow[0].GetPositionY()).To(Equal(int32(0)))
			Expect(firstRow[0].GetUnitId()).To(Equal("unit_1"))

			Expect(firstRow[1].GetPositionX()).To(Equal(int32(1)))
			Expect(firstRow[1].GetPositionY()).To(Equal(int32(0)))
			Expect(firstRow[1].GetUnitId()).To(Equal("unit_2"))

			// Verify columns of the second row
			secondRow := foundFormation.GetRows()[1].GetColumns()
			Expect(secondRow[0].GetPositionX()).To(Equal(int32(0)))
			Expect(secondRow[0].GetPositionY()).To(Equal(int32(1)))
			Expect(secondRow[0].GetUnitId()).To(Equal("unit_4"))

			Expect(secondRow[1].GetPositionX()).To(Equal(int32(1)))
			Expect(secondRow[1].GetPositionY()).To(Equal(int32(1)))
			Expect(secondRow[1].GetUnitId()).To(Equal("unit_5"))
		})
	})

	Context("When one formation is upserted", func() {
		var upsertFormation *domain.FormationEntity

		BeforeAll(func() {
			upsertFormation = domain.NewFormationEntity().WithDefaults().
				WithId("upsert-test-123").
				WithRowsFromJson(smallRowsJson).
				Build()
		})

		It("should successfully upsert the formation", func() {
			err := formationRepo.Upsert(ctx, upsertFormation)
			Expect(err).ToNot(HaveOccurred(), "failed to upsert formation")

			foundFormation, err := formationRepo.GetById(ctx, "upsert-test-123")
			Expect(err).ToNot(HaveOccurred(), "failed to get formation by ID")
			Expect(foundFormation).ToNot(BeNil(), "expected to find a formation")
			Expect(foundFormation.GetId()).To(Equal("upsert-test-123"))
			Expect(foundFormation.GetRows()).To(HaveLen(1), "expected 1 row")
			Expect(foundFormation.GetRows()[0].GetColumns()).To(HaveLen(2), "expected 2 columns")

			firstColumn := foundFormation.GetRows()[0].GetColumns()[0]
			Expect(firstColumn.GetPositionX()).To(Equal(int32(0)))
			Expect(firstColumn.GetPositionY()).To(Equal(int32(0)))
			Expect(firstColumn.GetUnitId()).To(Equal("unit_1"))

			secondColumn := foundFormation.GetRows()[0].GetColumns()[1]
			Expect(secondColumn.GetPositionX()).To(Equal(int32(1)))
			Expect(secondColumn.GetPositionY()).To(Equal(int32(0)))
			Expect(secondColumn.GetUnitId()).To(Equal("unit_2"))
		})
	})
})
