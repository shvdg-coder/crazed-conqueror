package integration

import (
	"shvdg/crazed-conquerer/internal/shared/testing"
	"shvdg/crazed-conquerer/internal/user/infrastructure"
)

// init registers the user schema with the shared suite.
func init() {
	suite := testing.GetSharedSuite()
	suite.AddSchema(infrastructure.NewUserSchema(suite.Database))
}
