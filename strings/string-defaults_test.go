package strings

import (
	"testing"
)

// Remove any leading spaces for a string
func Test_CleanStringEmptyLeading(t *testing.T) {
	inputString := "  hello there!"
	outputString := CleanString(inputString)

	if outputString != "hello there!" {
		t.Error("'  hello there!' did not default to 'hello there!'; leading spaces not truncated")
	}
}

// Remove all trailing spaces for a string
func Test_CleanStringEmptyTrailing(t *testing.T) {
	inputString := "hello there!  "
	outputString := CleanString(inputString)

	if outputString != "hello there!" {
		t.Error("'hello there!  ' did not default to 'hello there!'; trailing spaces not truncated.")
	}
}

// Ensure both leading and trailing spaces are removed.
func Test_CleanStringEmptyLeadingAndTrailing(t *testing.T) {
	inputString := "  hello there!  "
	outputString := CleanString(inputString)

	if outputString != "hello there!" {
		t.Error("'  hello there!  ' did not default to 'hello there!'; leading and trailing spaces not truncated.")
	}
}

// Ensure both leading and trailing spaces are removed,
// while preserving the line break in the middle of the string.
func Test_CleanStringEmptyLeadingAndTrailingWithSpaceInMiddle(t *testing.T) {
	inputString := "  hello \n there!  "
	outputString := CleanString(inputString)

	if outputString != "hello \n there!" {
		t.Error("'  hello \n there!  ' did not default to 'hello there!'; leading and trailing spaces not truncated while preserving the middle line break.")
	}
}

// Remove the line break from the beginning of the string.
func Test_CleanStringLeadingLineBreak(t *testing.T) {
	inputString := "\nhello there!"
	outputString := CleanString(inputString)

	if outputString != "hello there!" {
		t.Error("'\nhello there!' did not default to 'hello there!'; leading line break not truncated")
	}
}

// Remove the trailing line break from the end of the string
func Test_CleanStringTrailingLineBreak(t *testing.T) {
	inputString := "hello there!\n"
	outputString := CleanString(inputString)

	if outputString != "hello there!" {
		t.Error("'hello there!\n' did not default to 'hello there!'; trailing line break not truncated")
	}
}

// Ensure all upper case letters are lower case
func Test_CleanStringAndLowerCase(t *testing.T) {
	actualValue := CleanAndLower("Attestify.io")
	expectedValue := "attestify.io"
	if expectedValue != actualValue {
		t.Errorf("Did not return the expected value.\nActual: %s\nExpected: %s",
			actualValue,
			expectedValue)
	}
}
