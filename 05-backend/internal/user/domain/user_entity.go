package domain

import (
	"shvdg/crazed-conquerer/internal/shared/converters"
	"time"
)

// GetLastLoginAtAsTime retrieves the timestamp as time.Time.
func (u *UserEntity) GetLastLoginAtAsTime() time.Time {
	return converters.TimestampToTime(u.GetLastLoginAt())
}

// GetCreatedAtAsTime retrieves the timestamp as time.Time.
func (u *UserEntity) GetCreatedAtAsTime() time.Time {
	return converters.TimestampToTime(u.GetCreatedAt())
}

// GetUpdatedAtAsTime retrieves the timestamp as time.Time.
func (u *UserEntity) GetUpdatedAtAsTime() time.Time {
	return converters.TimestampToTime(u.GetUpdatedAt())
}
