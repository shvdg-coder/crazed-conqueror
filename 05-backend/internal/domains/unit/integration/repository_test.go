package integration

import (
	"context"
	"shvdg/crazed-conquerer/internal/domains/unit/domain"
	infa "shvdg/crazed-conquerer/internal/domains/unit/infrastructure"
	"shvdg/crazed-conquerer/internal/shared/contexts"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/testing"
	"shvdg/crazed-conquerer/internal/shared/testing/shared"

	"github.com/jackc/pgx/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Unit Repository", Ordered, func() {
	var err error
	var transaction pgx.Tx
	var ctx context.Context

	var suite *testing.Suite
	var unitRepo *infa.UnitRepositoryImpl

	BeforeAll(func() {
		suite = shared.GetSharedSuite()
		transaction, err = suite.StartTransaction()
		Expect(err).ToNot(HaveOccurred(), "failed to start transaction")

		ctx = contexts.SetTransaction(suite.Context, transaction)
		unitRepo = infa.NewUnitRepositoryImpl(suite.Database)
	})

	AfterAll(func() {
		err := transaction.Rollback(ctx)
		Expect(err).ToNot(HaveOccurred(), "failed to rollback transaction")
	})

	Context("When one unit is created", func() {
		var unit *domain.UnitEntity

		BeforeAll(func() {
			unit = domain.NewUnitEntity().WithDefaults().Build()
		})

		It("should successfully store the unit in the database", func() {
			err := unitRepo.Create(ctx, unit)
			Expect(err).ToNot(HaveOccurred(), "failed to create unit")

			fields := []string{infa.FieldId, infa.FieldVocation, infa.FieldName}
			values := []any{unit.GetId(), unit.GetVocation(), unit.GetName()}

			count, err := database.Count(ctx, suite.Database, infa.TableName, fields, values)
			Expect(err).ToNot(HaveOccurred(), "failed to count units")
			Expect(count).To(Equal(1), "expected 1 unit to be created")
		})
	})

	Context("When retrieving a unit by ID", func() {
		var unit *domain.UnitEntity

		BeforeAll(func() {
			unit = domain.NewUnitEntity().WithDefaults().
				WithId("find-me-123").
				Build()
			err := unitRepo.Create(ctx, unit)
			Expect(err).ToNot(HaveOccurred(), "failed to create unit")
		})

		It("should return the correct unit", func() {
			foundUnit, err := unitRepo.GetById(ctx, "find-me-123")
			Expect(err).ToNot(HaveOccurred(), "failed to get unit by ID")
			Expect(foundUnit).ToNot(BeNil(), "expected to find a unit")
			Expect(foundUnit.GetId()).To(Equal(unit.GetId()))
		})
	})

	Context("When one unit is updated", func() {
		var unit *domain.UnitEntity

		BeforeAll(func() {
			unit = domain.NewUnitEntity().WithDefaults().Build()
			err := unitRepo.Create(ctx, unit)
			Expect(err).ToNot(HaveOccurred(), "failed to create unit")
		})

		It("should successfully update the unit in the database", func() {
			unit.Name = "A New Name"
			unit.Level = "50"
			err := unitRepo.Update(ctx, unit)
			Expect(err).ToNot(HaveOccurred(), "failed to update unit")

			updatedUnit, err := unitRepo.GetById(ctx, unit.GetId())
			Expect(err).ToNot(HaveOccurred(), "failed to retrieve updated unit")
			Expect(updatedUnit.GetName()).To(Equal("A New Name"))
			Expect(updatedUnit.GetLevel()).To(Equal("50"))
		})
	})

	Context("When a unit is upserted", func() {
		var unit *domain.UnitEntity

		BeforeAll(func() {
			unit = domain.NewUnitEntity().WithDefaults().
				WithId("upsert-unit-789").
				WithName("Initial Name").
				WithLevel("10").
				Build()
		})

		It("should insert the unit if they do not exist", func() {
			err := unitRepo.Upsert(ctx, unit)
			Expect(err).ToNot(HaveOccurred(), "failed to upsert new unit")

			retrieved, err := unitRepo.GetById(ctx, "upsert-unit-789")
			Expect(err).ToNot(HaveOccurred())
			Expect(retrieved).ToNot(BeNil())
			Expect(retrieved.GetName()).To(Equal("Initial Name"))
			Expect(retrieved.GetLevel()).To(Equal("10"))
		})

		It("should update the unit if they already exist", func() {
			unit.Name = "Updated Name"
			unit.Level = "25"
			err := unitRepo.Upsert(ctx, unit)
			Expect(err).ToNot(HaveOccurred(), "failed to upsert existing unit")

			retrieved, err := unitRepo.GetById(ctx, "upsert-unit-789")
			Expect(err).ToNot(HaveOccurred())
			Expect(retrieved).ToNot(BeNil())
			Expect(retrieved.GetName()).To(Equal("Updated Name"))
			Expect(retrieved.GetLevel()).To(Equal("25"))
		})
	})

	Context("When one unit is deleted", func() {
		var unit *domain.UnitEntity

		BeforeAll(func() {
			unit = domain.NewUnitEntity().WithDefaults().Build()
			err := unitRepo.Create(ctx, unit)
			Expect(err).ToNot(HaveOccurred(), "failed to create unit")
		})

		It("should successfully remove the unit from the database", func() {
			err := unitRepo.Delete(ctx, unit)
			Expect(err).ToNot(HaveOccurred(), "failed to delete unit")

			count, err := database.Count(ctx, suite.Database, infa.TableName, []string{infa.FieldId}, []any{unit.GetId()})
			Expect(err).ToNot(HaveOccurred(), "failed to count units")
			Expect(count).To(BeZero(), "expected unit to be deleted")
		})
	})
})
