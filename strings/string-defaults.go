package strings

import "strings"

// CleanString removes all leading & trailing spaces, and line breaks in a string,
// while retaining any other spaces or line breaks which are not leading elements of the string.
func CleanString(value string) string {
	return strings.TrimSpace(value)
}

// CleanAndLower cleans the string with CleanString, then lowers all capital-cased letters.
func CleanAndLower(value string) string {
	clean := CleanString(value)
	cleanAndLow := strings.ToLower(clean)
	return cleanAndLow
}
