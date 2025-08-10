package shared

import (
	"log"
	"shvdg/crazed-conquerer/internal/shared/testing"
	"shvdg/crazed-conquerer/internal/user/infrastructure"
	"sync"
)

var (
	suite     *testing.Suite
	setupOnce sync.Once
	mutex     sync.RWMutex
)

// GetSharedSuite returns the shared test suite instance with infrastructure schemas, creating it if necessary
func GetSharedSuite() *testing.Suite {
	setupOnce.Do(func() {
		suite = testing.NewTestSuite()
		suite.AddSchema(infrastructure.NewUserSchema(suite.Database))

		err := suite.CreateAllTables(suite.Context)
		if err != nil {
			log.Fatalf("failed to create tables: %s", err)
		}
	})

	return suite
}

// CleanupSharedSuite terminates the shared test suite
func CleanupSharedSuite() {
	mutex.Lock()
	defer mutex.Unlock()

	if suite != nil {
		suite.Terminate()
		suite = nil
	}
}
