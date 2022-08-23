package permission_list

import (
	"github.com/attestify/go-kernel/strings"
)

// Permission
// Expected behaviour
// - Only characters allowed are alphabet characters
// - Will remove any numeric values
// - Replace any special characters with a dash
// - Must only start with an alpha character, and can only end with an alpha character
type Permission struct {
	value string
}

// todo - FUTURE - Implement Error pattern for a string with no values
func New(value string) Permission {
	value = strings.CleanAndLower(value)
	value = strings.RemoveAllNumbers(value)
	value = strings.ReplaceSpecialCharactersWithDash(value)
	value = strings.CleanLeadAndTrailSpecialCharacter(value)
	return Permission{
		value: value,
	}
}

func (permission Permission) Value() string {
	return permission.value
}
