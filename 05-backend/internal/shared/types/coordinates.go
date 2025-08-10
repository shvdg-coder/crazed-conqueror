package types

// NewCoordinates initializes a new Coordinates object.
func NewCoordinates(x, y int32) *Coordinates {
	return &Coordinates{X: x, Y: y}
}
