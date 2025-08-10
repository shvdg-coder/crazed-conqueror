package domain

import (
	sharedDomain "shvdg/crazed-conquerer/internal/shared/types"
)

// Definitions contains all ability definitions for a combatant
type Definitions struct {
	Dispels  []DispelDefinition
	Afflicts []AfflictDefinition
	Attacks  []AttackDefinition
	Moves    []MoveDefinition
	Heals    []HealDefinition
	Summons  []SummonDefinition
}

// AttackDefinition contains all data needed to execute an attack
type AttackDefinition struct {
	Name                  string
	AttackPowerPercentage int32
	Range                 []sharedDomain.Coordinates
}

// MoveDefinition contains all data needed to execute a move
type MoveDefinition struct {
	Name  string
	Range []sharedDomain.Coordinates
}

// HealDefinition contains all data needed to execute a heal
type HealDefinition struct {
	Name  string
	Range []sharedDomain.Coordinates
}

// AfflictDefinition contains all data needed to execute an affliction
type AfflictDefinition struct {
	Name  string
	Range []sharedDomain.Coordinates
}

// DispelDefinition contains all data needed to execute a dispel
type DispelDefinition struct {
	Name  string
	Range []sharedDomain.Coordinates
}

// SummonDefinition contains all data needed to execute a summon
type SummonDefinition struct {
	Name  string
	Range []sharedDomain.Coordinates
}
