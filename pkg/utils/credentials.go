package utils

import (
	"fmt"

	"github.com/create-go-app/fiber-go-template/pkg/const"
)

// GetCredentialsByRole func for getting credentials from a role name.
func GetCredentialsByRole(role string) ([]string, error) {
	// Define credentials variable.
	var credentials []string

	// Switch given role.
	switch role {
	case _const.AdminRoleName:
		// Admin credentials (all access).
		credentials = []string{
			_const.BookCreateCredential,
			_const.BookUpdateCredential,
			_const.BookDeleteCredential,
		}
	case _const.ModeratorRoleName:
		// Moderator credentials (only book creation and update).
		credentials = []string{
			_const.BookCreateCredential,
			_const.BookUpdateCredential,
		}
	case _const.UserRoleName:
		// Simple user credentials (only book creation).
		credentials = []string{
			_const.BookCreateCredential,
		}
	default:
		// Return error message.
		return nil, fmt.Errorf("role '%v' does not exist", role)
	}

	return credentials, nil
}
