package integration

import (
	"context"
	infra "shvdg/crazed-conquerer/internal/domains/character-formation/infrastructure"
	"shvdg/crazed-conquerer/internal/shared/contexts"
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
	var _ *infra.CharacterFormationRepositoryImpl

	BeforeAll(func() {
		suite = shared.GetSharedSuite()
		transaction, err = suite.StartTransaction()
		Expect(err).ToNot(HaveOccurred(), "failed to start transaction")

		ctx = contexts.SetTransaction(suite.Context, transaction)
		_ = infra.NewCharacterFormationRepositoryImpl(suite.Database)
	})

	AfterAll(func() {
		err := transaction.Rollback(ctx)
		Expect(err).ToNot(HaveOccurred(), "failed to rollback transaction")
	})
})
