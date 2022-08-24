package strings

import "testing"

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path **/

// Remove any leading spaces for a string
func Test_CleanStringEmptyLeading(t *testing.T) {
	setup(t)
	inputString := "  hello there!"
	outputString := CleanString(inputString)

	if outputString != "hello there!" {
		t.Error("'  hello there!' did not default to 'hello there!'; leading spaces not truncated")
	}
}

// Remove all trailing spaces for a string
func Test_CleanStringEmptyTrailing(t *testing.T) {
	setup(t)
	inputString := "hello there!  "
	outputString := CleanString(inputString)

	if outputString != "hello there!" {
		t.Error("'hello there!  ' did not default to 'hello there!'; trailing spaces not truncated.")
	}
}

// Ensure both leading and trailing spaces are removed.
func Test_CleanStringEmptyLeadingAndTrailing(t *testing.T) {
	setup(t)
	inputString := "  hello there!  "
	outputString := CleanString(inputString)

	if outputString != "hello there!" {
		t.Error("'  hello there!  ' did not default to 'hello there!'; leading and trailing spaces not truncated.")
	}
}

// Ensure both leading and trailing spaces are removed,
// while preserving the line break in the middle of the string.
func Test_CleanStringEmptyLeadingAndTrailingWithSpaceInMiddle(t *testing.T) {
	setup(t)
	inputString := "  hello \n there!  "
	outputString := CleanString(inputString)

	if outputString != "hello \n there!" {
		t.Error("'  hello \n there!  ' did not default to 'hello there!'; leading and trailing spaces not truncated while preserving the middle line break.")
	}
}

// Remove the line break from the beginning of the string.
func Test_CleanStringLeadingLineBreak(t *testing.T) {
	setup(t)
	inputString := "\nhello there!"
	outputString := CleanString(inputString)

	if outputString != "hello there!" {
		t.Error("'\nhello there!' did not default to 'hello there!'; leading line break not truncated")
	}
}

// Remove the trailing line break from the end of the string
func Test_CleanStringTrailingLineBreak(t *testing.T) {
	setup(t)
	inputString := "hello there!\n"
	outputString := CleanString(inputString)

	if outputString != "hello there!" {
		t.Error("'hello there!\n' did not default to 'hello there!'; trailing line break not truncated")
	}
}

// Ensure all upper case letters are lower case
func Test_CleanStringAndLowerCase(t *testing.T) {
	setup(t)
	actualValue := CleanAndLower("Attestify.io")
	expectedValue := "attestify.io"
	if expectedValue != actualValue {
		t.Errorf("Did not return the expected value.\nActual: %s\nExpected: %s",
			actualValue,
			expectedValue)
	}
}

// Remove all numbers from a string
func Test_RemoveAllNumbers(t *testing.T) {
	setup(t)
	input := "work9"
	actual := RemoveAllNumbers(input)

	expected := "work"
	if expected != actual {
		t.Errorf("Actual value was differnet from expected value.\n Expected: %s\n Actual: %s\n", expected, actual)
	}
}

// Replace all special characters with a dash
func Test_ReplaceSpecialCharactersWithDash(t *testing.T) {
	setup(t)
	input := "$There%Should*Now-Be@all(dashes)"
	actual := ReplaceSpecialCharactersWithDash(input)

	expected := "-There-Should-Now-Be-all-dashes-"
	if expected != actual {
		t.Errorf("Actual value was differnet from expected value.\n Expected: %s\n Actual: %s\n", expected, actual)
	}
}

// Replace the first and last character of a string if it's a special character
func Test_CleanLeadAndTrailSpecialCharacter(t *testing.T) {
	setup(t)
	input := "-should-remove-first-and-last-special-character-"
	actual := CleanLeadAndTrailSpecialCharacter(input)

	expected := "should-remove-first-and-last-special-character"
	if expected != actual {
		t.Errorf("Actual value was differnet from expected value.\n Expected: %s\n Actual: %s\n", expected, actual)
	}
}

/** Sad Path **/
