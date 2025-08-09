package environment

import (
	"os"

	_ "github.com/joho/godotenv/autoload" // Load environment variables from a .env file
)

// EnvStr retrieves an environment value as a string.
func EnvStr(key string) string {
	return os.Getenv(key)
}
