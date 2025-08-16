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
	sharedSuite *testing.Suite
	initOnce    sync.Once
	mutex       sync.RWMutex
)

// GetSharedSuite returns the shared test suite instance with all schemas, creating it if necessary
func GetSharedSuite() *testing.Suite {
	initOnce.Do(func() {
		log.Println("Initializing global test suite...")
		sharedSuite = testing.NewTestSuite()
		sharedSuite.AddSchema(userinfra.NewUserSchema(sharedSuite.Database))
		sharedSuite.AddSchema(characterinfra.NewCharacterSchema(sharedSuite.Database))
		sharedSuite.AddSchema(unitinfra.NewUnitSchema(sharedSuite.Database))
		sharedSuite.AddSchema(usercharacterinfra.NewUserCharacterSchema(sharedSuite.Database))
		sharedSuite.AddSchema(characterunitinfra.NewCharacterUnitSchema(sharedSuite.Database))
		sharedSuite.AddSchema(formationinfra.NewFormationSchema(sharedSuite.Database))
		sharedSuite.AddSchema(characterformationinfra.NewCharacterFormationSchema(sharedSuite.Database))

		err := sharedSuite.CreateAllTables(sharedSuite.Context)
		if err != nil {
			log.Fatalf("failed to create tables: %s", err)
		}
	})

	return sharedSuite
}

// CleanupSharedSuite terminates the global test suite
func CleanupSharedSuite() {
	mutex.Lock()
	defer mutex.Unlock()

	if sharedSuite != nil {
		log.Println("Cleaning up global test suite...")
		sharedSuite.Terminate()
		sharedSuite = nil
	}
}
