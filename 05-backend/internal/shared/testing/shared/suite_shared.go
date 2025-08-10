package shared

import (
	"log"
	characterinfra "shvdg/crazed-conquerer/internal/domains/character/infrastructure"
	usercharacterinfra "shvdg/crazed-conquerer/internal/domains/user-character/infrastructure"
	userinfra "shvdg/crazed-conquerer/internal/domains/user/infrastructure"
	"shvdg/crazed-conquerer/internal/shared/testing"
	"sync"
)

var (
	suite     *testing.Suite
	setupOnce sync.Once
	mutex     sync.RWMutex
)

// GetSharedSuite returns the shared test suite instance with userinfra schemas, creating it if necessary
func GetSharedSuite() *testing.Suite {
	setupOnce.Do(func() {
		suite = testing.NewTestSuite()
		suite.AddSchema(userinfra.NewUserSchema(suite.Database))
		suite.AddSchema(characterinfra.NewCharacterSchema(suite.Database))
		suite.AddSchema(usercharacterinfra.NewUserCharacterSchema(suite.Database))

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
