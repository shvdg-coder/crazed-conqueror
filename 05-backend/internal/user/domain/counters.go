package domain

import "sync/atomic"

var userCounter uint64 = 1

// NextUser generates and returns the next unique user number.
func NextUser() uint64 { return atomic.AddUint64(&userCounter, 1) }
