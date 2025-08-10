package domain

// CreateDefaultBattlefield creates a new default battlefield.
func CreateDefaultBattlefield() *Battlefield {
	return NewBattlefieldBuilder().AddEmptyRows(1, 10).Build()
}
