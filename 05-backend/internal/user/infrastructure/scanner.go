package infrastructure

import (
	"fmt"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/user/domain"
)

// ScanUserEntity scans database row data into a UserEntity
func ScanUserEntity(scanner database.RowScanner) (*domain.UserEntity, error) {
	var user domain.UserEntity

	err := scanner.Scan(
		&user.Id,
		&user.Email,
		&user.DisplayName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLoginAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to scan user entity: %w", err)
	}

	return &user, nil
}
