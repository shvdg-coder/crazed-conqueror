package shared

import (
	"log"
	characterformationinfra "shvdg/crazed-conquerer/internal/domains/character-formation/infrastructure"
	characterunitinfra "shvdg/crazed-conquerer/internal/domains/character-unit/infrastructure"
	characterinfra "shvdg/crazed-conquerer/internal/domains/character/infrastructure"
	formationinfra "shvdg/crazed-conquerer/internal/domains/formation/infrastructure"
	unitinfra "shvdg/crazed-conquerer/internal/domains/unit/infrastructure"
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
		suite.AddSchema(unitinfra.NewUnitSchema(suite.Database))
		suite.AddSchema(usercharacterinfra.NewUserCharacterSchema(suite.Database))
		suite.AddSchema(characterunitinfra.NewCharacterUnitSchema(suite.Database))
		suite.AddSchema(formationinfra.NewFormationSchema(suite.Database))
		suite.AddSchema(characterformationinfra.NewCharacterFormationSchema(suite.Database))

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
