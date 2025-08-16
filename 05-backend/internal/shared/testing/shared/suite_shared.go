package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

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
)

type suiteData struct {
	DSN   string `json:"dsn"`
	Host  string `json:"host"`
	Port  string `json:"port"`
	Error string `json:"error,omitempty"`
}

var (
	sharedSuite *testing.Suite
	initOnce    sync.Once
	mutex       sync.RWMutex
	processID   = fmt.Sprintf("PID-%d", os.Getpid())
)

// InitializeGlobalSuite sets up the shared test suite and returns serialized connection data
// Returns error-prefixed bytes if setup fails
func InitializeGlobalSuite() []byte {
	log.Printf("[%s] Starting global test suite initialization...", processID)

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Printf("[%s] Received shutdown signal, cleaning up...", processID)
		CleanupGlobalSuite()
		os.Exit(1)
	}()

	var initErr error
	initOnce.Do(func() {
		defer func() {
			if r := recover(); r != nil {
				initErr = fmt.Errorf("panic during initialization: %v", r)
				log.Printf("[%s] Panic during suite initialization: %v", processID, r)
			}
		}()

		log.Printf("[%s] Creating test suite containers and database...", processID)
		sharedSuite = testing.NewTestSuite()

		if sharedSuite == nil {
			initErr = fmt.Errorf("failed to create test suite")
			return
		}

		// Add all schemas
		log.Printf("[%s] Adding domain schemas...", processID)
		sharedSuite.AddSchema(userinfra.NewUserSchema(sharedSuite.Database))
		sharedSuite.AddSchema(characterinfra.NewCharacterSchema(sharedSuite.Database))
		sharedSuite.AddSchema(unitinfra.NewUnitSchema(sharedSuite.Database))
		sharedSuite.AddSchema(usercharacterinfra.NewUserCharacterSchema(sharedSuite.Database))
		sharedSuite.AddSchema(characterunitinfra.NewCharacterUnitSchema(sharedSuite.Database))
		sharedSuite.AddSchema(formationinfra.NewFormationSchema(sharedSuite.Database))
		sharedSuite.AddSchema(characterformationinfra.NewCharacterFormationSchema(sharedSuite.Database))

		// Create tables
		log.Printf("[%s] Creating database tables...", processID)
		err := sharedSuite.CreateAllTables(sharedSuite.Context)
		if err != nil {
			initErr = fmt.Errorf("failed to create tables: %w", err)
			return
		}

		log.Printf("[%s] Global test suite initialized successfully", processID)
	})

	// Prepare response data
	data := suiteData{}
	if initErr != nil {
		data.Error = initErr.Error()
		log.Printf("[%s] Suite initialization failed: %v", processID, initErr)
	} else if sharedSuite != nil {
		data.DSN = sharedSuite.Dsn
		data.Host = sharedSuite.Postgres.Host
		data.Port = sharedSuite.Postgres.Port
		log.Printf("[%s] Returning suite data - Host: %s, Port: %s", processID, data.Host, data.Port)
	} else {
		data.Error = "suite is nil after initialization"
	}

	// Serialize and return data
	bytes, err := json.Marshal(data)
	if err != nil {
		errorData := suiteData{Error: fmt.Sprintf("failed to marshal suite data: %v", err)}
		if errorBytes, marshalErr := json.Marshal(errorData); marshalErr == nil {
			return errorBytes
		}
		return []byte(fmt.Sprintf(`{"error":"failed to marshal suite data: %v"}`, err))
	}

	return bytes
}

// GetSharedSuite returns the shared test suite instance (deprecated - use InitializeGlobalSuite/SetupLocalSuite)
func GetSharedSuite() *testing.Suite {
	mutex.RLock()
	defer mutex.RUnlock()
	return sharedSuite
}

// SetupLocalSuite establishes connection to the shared database using provided data
// Includes connection validation and retry logic
func SetupLocalSuite(data []byte) error {
	log.Printf("[%s] Setting up local suite connection...", processID)

	if len(data) == 0 {
		return fmt.Errorf("no suite data provided")
	}

	var suiteInfo suiteData
	if err := json.Unmarshal(data, &suiteInfo); err != nil {
		return fmt.Errorf("failed to unmarshal suite data: %w", err)
	}

	// Check for initialization errors from global setup
	if suiteInfo.Error != "" {
		return fmt.Errorf("global suite initialization failed: %s", suiteInfo.Error)
	}

	// Validate required connection data
	if suiteInfo.DSN == "" || suiteInfo.Host == "" || suiteInfo.Port == "" {
		return fmt.Errorf("invalid suite data: missing connection information")
	}

	// Create local database connection using shared DSN
	log.Printf("[%s] Creating local database connection with DSN...", processID)
	ctx := context.Background()

	db, err := database.NewService("postgres", suiteInfo.DSN, database.WithConnection(ctx))
	if err != nil {
		return fmt.Errorf("failed to create database connection: %w", err)
	}

	// Test the connection
	log.Printf("[%s] Testing database connection...", processID)
	if err := db.GetPool().Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Create minimal local suite
	log.Printf("[%s] Creating local suite instance...", processID)
	mutex.Lock()
	sharedSuite = &testing.Suite{
		Context:  ctx,
		Database: db,
		Dsn:      suiteInfo.DSN,
		Schemas:  schemas.NewService(db),
	}
	mutex.Unlock()

	log.Printf("[%s] Local suite connection established successfully", processID)
	return nil
}

// CleanupGlobalSuite terminates the global test suite with proper resource cleanup
func CleanupGlobalSuite() {
	log.Printf("[%s] Starting global test suite cleanup...", processID)

	mutex.Lock()
	defer mutex.Unlock()

	if sharedSuite != nil {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[%s] Panic during cleanup: %v", processID, r)
			}
		}()

		log.Printf("[%s] Terminating containers and cleaning up resources...", processID)

		// Use defer statements for proper cleanup order
		defer func() {
			sharedSuite = nil
			log.Printf("[%s] Global test suite cleanup completed", processID)
		}()

		defer func() {
			if sharedSuite != nil {
				sharedSuite.Terminate()
			}
		}()

		// Additional cleanup can be added here if needed
	} else {
		log.Printf("[%s] No global test suite to clean up", processID)
	}
}

// CleanupSharedSuite is deprecated - use CleanupGlobalSuite instead
func CleanupSharedSuite() {
	CleanupGlobalSuite()
}
