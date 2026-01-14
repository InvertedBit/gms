package adminhandlers

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

// isDuplicateKeyError checks if the error is a unique constraint violation
// Works across different database drivers (PostgreSQL, SQLite, etc.)
func isDuplicateKeyError(err error) bool {
	if err == nil {
		return false
	}
	
	// Check GORM's built-in error type
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	}
	
	// Check error message for common patterns
	errMsg := strings.ToLower(err.Error())
	return strings.Contains(errMsg, "duplicate") ||
		strings.Contains(errMsg, "unique constraint") ||
		strings.Contains(errMsg, "violates unique") ||
		strings.Contains(errMsg, "23505") // PostgreSQL unique violation error code
}
