package strings

import (
	"regexp"
	"strings"
)

// CleanString removes all leading & trailing spaces, and line breaks in a string,
// while retaining any other spaces or line breaks which are not leading elements of the string.
func CleanString(value string) string {
	if LengthIsZero(value) {
		return ""
	}
	return strings.TrimSpace(value)
}

// CleanAndLower cleans the string with CleanString, then lowers all capital-cased letters.
func CleanAndLower(value string) string {
	if LengthIsZero(value) {
		return ""
	}
	clean := CleanString(value)
	cleanAndLow := strings.ToLower(clean)
	return cleanAndLow
}

// RemoveAllNumbers removes all numbers from a string
func RemoveAllNumbers(value string) string {
	if LengthIsZero(value) {
		return ""
	}
	reg, _ := regexp.Compile("[0-9]")
	return reg.ReplaceAllString(value, "")
}

// ReplaceSpecialCharactersWithDash Replaces all special characters with a dash
func ReplaceSpecialCharactersWithDash(value string) string {
	if LengthIsZero(value) {
		return ""
	}
	reg, _ := regexp.Compile(`[^\w]`)
	return reg.ReplaceAllString(value, "-")
}

// CleanLeadAndTrailSpecialCharacter replaces the first, and last, characters of a string if they are special characters
func CleanLeadAndTrailSpecialCharacter(value string) string {
	if LengthIsZero(value) {
		return ""
	}
	// Declare regular expresion to identify a special characters
	reg, _ := regexp.Compile(`[^\w]`)

	// check first character, if a special character remove.
	firstCharacter := value[0:1]
	if reg.MatchString(firstCharacter) {
		value = strings.TrimPrefix(value, firstCharacter)
	}

	// check last character, if a special character remove.
	lastCharacter := value[len(value)-1:]
	if reg.MatchString(lastCharacter) {
		value = strings.TrimSuffix(value, lastCharacter)
	}

	return value
}

func LengthIsZero(value string) bool {
	length := len([]rune(value))
	if length == 0 {
		return true
	}
	return false
}
