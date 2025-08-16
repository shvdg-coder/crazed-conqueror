package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	characterformationinfra "shvdg/crazed-conquerer/internal/domains/character-formation/infrastructure"
	characterunitinfra "shvdg/crazed-conquerer/internal/domains/character-unit/infrastructure"
	characterinfra "shvdg/crazed-conquerer/internal/domains/character/infrastructure"
	formationinfra "shvdg/crazed-conquerer/internal/domains/formation/infrastructure"
	unitinfra "shvdg/crazed-conquerer/internal/domains/unit/infrastructure"
	usercharacterinfra "shvdg/crazed-conquerer/internal/domains/user-character/infrastructure"
	userinfra "shvdg/crazed-conquerer/internal/domains/user/infrastructure"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/schemas"
	"shvdg/crazed-conquerer/internal/shared/testing"
	"sync"
)

// suiteData represents the serialized connection data for the shared test suite
type suiteData struct {
	DSN   string `json:"dsn"`
	Host  string `json:"host"`
	Port  string `json:"port"`
	Error string `json:"error,omitempty"`
}

// Global variables that are shared between the tests
var (
	sharedSuite *testing.Suite
	initOnce    sync.Once
	mutex       sync.RWMutex
	processId   = fmt.Sprintf("PID-%d", os.Getpid())
)

// InitializeTestSuite sets up the shared test suite and returns serialized connection data
func InitializeTestSuite() []byte {
	log.Printf("[%s] Create test suite", processId)

	var initErr error
	initOnce.Do(func() {
		defer recoverFromPanic(&initErr)

		sharedSuite = testing.NewTestSuite()
		if sharedSuite == nil {
			initErr = fmt.Errorf("failed to create test suite")
			return
		}

		addDomainSchemas()

		if err := sharedSuite.CreateAllTables(sharedSuite.Context); err != nil {
			initErr = fmt.Errorf("failed to create tables: %w", err)
			return
		}

		log.Printf("[%s] Suite initialized successfully", processId)
	})

	return marshalSuiteData(initErr)
}

// addDomainSchemas adds schemas for all domains to the shared test suite
func addDomainSchemas() {
	sharedSuite.AddSchema(userinfra.NewUserSchema(sharedSuite.Database))
	sharedSuite.AddSchema(characterinfra.NewCharacterSchema(sharedSuite.Database))
	sharedSuite.AddSchema(unitinfra.NewUnitSchema(sharedSuite.Database))
	sharedSuite.AddSchema(usercharacterinfra.NewUserCharacterSchema(sharedSuite.Database))
	sharedSuite.AddSchema(characterunitinfra.NewCharacterUnitSchema(sharedSuite.Database))
	sharedSuite.AddSchema(formationinfra.NewFormationSchema(sharedSuite.Database))
	sharedSuite.AddSchema(characterformationinfra.NewCharacterFormationSchema(sharedSuite.Database))
}

// GetSharedSuite returns the shared test suite instance
func GetSharedSuite() *testing.Suite {
	mutex.RLock()
	defer mutex.RUnlock()
	return sharedSuite
}

// SetupLocalSuite establishes connection to the shared database
func SetupLocalSuite(data []byte) error {
	if len(data) == 0 {
		return fmt.Errorf("no suite data provided")
	}

	var suiteInfo suiteData
	if err := json.Unmarshal(data, &suiteInfo); err != nil {
		return fmt.Errorf("failed to unmarshal suite data: %w", err)
	}

	if suiteInfo.Error != "" {
		return fmt.Errorf("test suite initialization failed: %s", suiteInfo.Error)
	}

	if suiteInfo.DSN == "" || suiteInfo.Host == "" || suiteInfo.Port == "" {
		return fmt.Errorf("invalid suite data: missing connection information")
	}

	return createLocalConnection(suiteInfo)
}

// CleanupSharedSuite terminates the shared test suite
func CleanupSharedSuite() {
	log.Printf("[%s] Cleaning up shared test suite", processId)

	mutex.Lock()
	defer mutex.Unlock()

	if sharedSuite != nil {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[%s] Panic during cleanup: %v", processId, r)
			}
			sharedSuite = nil
		}()

		sharedSuite.Terminate()
	}
}

// recoverFromPanic recovers from panics during initialization
func recoverFromPanic(initErr *error) {
	if r := recover(); r != nil {
		*initErr = fmt.Errorf("panic during initialization: %v", r)
		log.Printf("[%s] Panic during initialization: %v", processId, r)
	}
}

// marshalSuiteData serializes the shared test suite connection data
func marshalSuiteData(initErr error) []byte {
	data := suiteData{}

	if initErr != nil {
		data.Error = initErr.Error()
	} else if sharedSuite != nil {
		data.DSN = sharedSuite.DSN
		data.Host = sharedSuite.Postgres.Host
		data.Port = sharedSuite.Postgres.Port
	} else {
		data.Error = "suite is nil after initialization"
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return []byte(fmt.Sprintf(`{"error":"marshal error: %v"}`, err))
	}
	return bytes
}

// createLocalConnection establishes a connection to the shared database
func createLocalConnection(suiteInfo suiteData) error {
	ctx := context.Background()
	db, err := database.NewService("postgres", suiteInfo.DSN, database.WithConnection(ctx))
	if err != nil {
		return fmt.Errorf("failed to create database connection: %w", err)
	}

	if err := db.GetPool().Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	mutex.Lock()
	sharedSuite = &testing.Suite{
		Context:  ctx,
		Database: db,
		DSN:      suiteInfo.DSN,
		Schemas:  schemas.NewService(db),
	}
	mutex.Unlock()

	return nil
}
