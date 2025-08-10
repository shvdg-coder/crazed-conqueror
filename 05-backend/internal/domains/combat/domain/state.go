package domain

import (
	sharedDomain "shvdg/crazed-conquerer/internal/shared/types"
)

// State represents the current state of a combatant
type State struct {
	Health      int32
	AttackPower int32

	Coordinates sharedDomain.Coordinates
	Facing      sharedDomain.Direction
	Stance      Stance

	TicksBetweenDispelling int32
	TicksBetweenAfflicting int32
	TicksBetweenAttacking  int32
	TicksBetweenMoving     int32
	TicksBetweenHealing    int32
	TicksBetweenSummoning  int32
}
