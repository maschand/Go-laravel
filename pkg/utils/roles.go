package utils

import (
	"fmt"

	"gitlab.com/d6825/golang_template/pkg/const"
)

// VerifyRole func for verifying a given role.
func VerifyRole(role string) (string, error) {
	// Switch given role.
	switch role {
	case _const.AdminRoleName:
		// Nothing to do, verified successfully.
	case _const.ModeratorRoleName:
		// Nothing to do, verified successfully.
	case _const.UserRoleName:
		// Nothing to do, verified successfully.
	default:
		// Return error message.
		return "", fmt.Errorf("role '%v' does not exist", role)
	}

	return role, nil
}
