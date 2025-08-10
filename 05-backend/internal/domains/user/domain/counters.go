package domain

import "sync/atomic"

var userCounter uint64 = 1

// NextUserNumber generates and returns the next unique user number.
func NextUserNumber() uint64 { return atomic.AddUint64(&userCounter, 1) }
