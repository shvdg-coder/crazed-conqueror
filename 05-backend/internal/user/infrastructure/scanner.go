package infrastructure

import (
	"fmt"
	"shvdg/crazed-conquerer/internal/shared/converters"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/user/domain"

	"github.com/jackc/pgx/v5/pgtype"
)

// ScanUserEntity scans database row data into a UserEntity
func ScanUserEntity(scanner database.RowScanner) (*domain.UserEntity, error) {
	var user domain.UserEntity
	var lastLoginAt, createdAt, updatedAt pgtype.Timestamp

	err := scanner.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.DisplayName,
		&createdAt,
		&updatedAt,
		&lastLoginAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to scan user entity: %w", err)
	}

	if createdAt.Valid {
		user.CreatedAt = converters.TimeToTimestamp(createdAt.Time)
	}
	if updatedAt.Valid {
		user.UpdatedAt = converters.TimeToTimestamp(updatedAt.Time)
	}
	if lastLoginAt.Valid {
		user.LastLoginAt = converters.TimeToTimestamp(lastLoginAt.Time)
	}

	return &user, nil
}
