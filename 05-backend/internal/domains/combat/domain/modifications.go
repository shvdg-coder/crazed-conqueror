package domain

import (
	sharedDomain "shvdg/crazed-conquerer/internal/shared/types"
)

// Modifications contain all active modifiers for a combatant
type Modifications struct {
	HealthModifiers                 []HealthModifier
	AttackPowerModifiers            []AttackPowerModifier
	LocationModifiers               []LocationModifier
	TicksBetweenAttackingModifiers  []TicksBetweenAttackingModifier
	TicksBetweenMovingModifiers     []TicksBetweenMovingModifier
	TicksBetweenDispellingModifiers []TicksBetweenDispellingModifier
	TicksBetweenAfflictingModifiers []TicksBetweenAfflictingModifier
	TicksBetweenHealingModifiers    []TicksBetweenHealingModifier
	TicksBetweenSummoningModifiers  []TicksBetweenSummoningModifier
}

// HealthModifier represents a health modification effect
type HealthModifier struct {
	Name          string
	HealthDelta   int32
	DurationTicks int32
}

// AttackPowerModifier represents an attack power modification effect
type AttackPowerModifier struct {
	Name          string
	AttackDelta   int32
	DurationTicks int32
}

// LocationModifier represents a location modification effect
type LocationModifier struct {
	Name             string
	CoordinatesDelta sharedDomain.Coordinates
	DurationTicks    int32
}

// TicksBetweenAttackingModifier represents an attack speed modification effect
type TicksBetweenAttackingModifier struct {
	Name          string
	TicksDelta    int32
	DurationTicks int32
}

// TicksBetweenMovingModifier represents a movement speed modification effect
type TicksBetweenMovingModifier struct {
	Name          string
	TicksDelta    int32
	DurationTicks int32
}

// TicksBetweenDispellingModifier represents a dispelling speed modification effect
type TicksBetweenDispellingModifier struct {
	Name          string
	TicksDelta    int32
	DurationTicks int32
}

// TicksBetweenAfflictingModifier represents an afflicting speed modification effect
type TicksBetweenAfflictingModifier struct {
	Name          string
	TicksDelta    int32
	DurationTicks int32
}

// TicksBetweenHealingModifier represents a healing speed modification effect
type TicksBetweenHealingModifier struct {
	Name          string
	TicksDelta    int32
	DurationTicks int32
}

// TicksBetweenSummoningModifier represents a summoning speed modification effect
type TicksBetweenSummoningModifier struct {
	Name          string
	TicksDelta    int32
	DurationTicks int32
}
