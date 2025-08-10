package domain

import (
	sharedDomain "shvdg/crazed-conquerer/internal/shared/types"
)

// Combatant represents a combat unit with all its properties
type Combatant struct {
	UnitId        string
	Team          sharedDomain.Team
	Definitions   Definitions
	State         State
	Modifications Modifications
}
