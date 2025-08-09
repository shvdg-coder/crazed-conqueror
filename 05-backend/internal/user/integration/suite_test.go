package integration

import (
	"shvdg/crazed-conquerer/internal/shared/testing/integration"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInfrastructure(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Infrastructure Tests")
}

// Executes the first block before and the second block after all the tests are run.
var _ = SynchronizedBeforeSuite(func() []byte {
	integration.GetSharedSuite()
	return nil
}, func(data []byte) {
	// N.A
})

// Executes the first block before and the second block after the teardown.
var _ = SynchronizedAfterSuite(func() {
	// N.A
}, func() {
	integration.CleanupSharedSuite()
})
