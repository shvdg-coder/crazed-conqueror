package domain

import "sync/atomic"

var unitCounter uint64 = 1

// NextUnitNumber generates and returns the next unique unit number.
func NextUnitNumber() uint64 { return atomic.AddUint64(&unitCounter, 1) }
