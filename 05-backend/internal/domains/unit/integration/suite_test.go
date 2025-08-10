package integration

import (
	"shvdg/crazed-conquerer/internal/shared/testing/shared"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInfrastructure(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unit Infrastructure Tests")
}

// Executes the first block before and the second block after all the tests are run.
var _ = SynchronizedBeforeSuite(func() []byte {
	shared.GetSharedSuite()
	return nil
}, func(data []byte) {
	// N.A
})

// Executes the first block before and the second block after the teardown.
var _ = SynchronizedAfterSuite(func() {
	// N.A
}, func() {
	shared.CleanupSharedSuite()
})
