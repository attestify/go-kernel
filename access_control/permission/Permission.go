package permission

import (
	"github.com/attestify/go-kernel/strings"
)

// todo - Create Control value object
// Permission
// Expected behaviour
// - Only characters allowed are alphabet characters
// - Will remove any numeric values
// - Replace any special characters with a dash
type Permission struct {
	value string
}

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