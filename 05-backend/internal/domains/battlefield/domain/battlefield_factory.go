package domain

// CreateDefaultBattlefield creates a new default battlefield.
func CreateDefaultBattlefield() *Battlefield {
	return NewBattlefieldBuilder().WithDimensions(1, 10).Build()
}
