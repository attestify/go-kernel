package top_level_domain

import (
	"testing"
)

/** Happy Path **/

func Test_InstantiateTopLevelDomain(t *testing.T) {

	tld, err := New("io")

	if err != nil {
		t.Error("An error was returned when no error was expected.")
	}
	if tld.Value() != "io" {
		t.Error("The value must return 'io'.")
	}

}

func Test_MustReturnErrorForEmptyString(t *testing.T) {


	_, err := New("")

	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err == nil {
		t.Fatal("An error was expected, but an error was not returned.")
	}

	if err.Error() != "The top level domain value must be atleast one (1) character." {
		t.Error("The following expected error message was not returned: 'The top level domain value must be atleast one (1) character.'.")
	}

}

func Test_TwoSameTopLevelDomainMustEqual(t *testing.T) {

	// Act
	tld1, err := New("io")
	tld2, err := New("io")

	// Assert
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	areEqual := tld1.Equals(tld2)
	if !areEqual {
		t.Errorf("The Top Level Domains do not equal.\nBase: %s\nComparator: %s", tld1.Value(), tld2.Value())
	}

}

/** Happy Path **/

func Test_MustReturnErrorForNumberInString(t *testing.T) {

	tld, err := New("1io")

	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err == nil {
		t.Fatal("An error was expected, but an error was not returned.")
	}

	if err.Error() != "The top level domain value can only be letters." {
		t.Error("The following expected error message was not returned: 'The top level domain value can only be letters.'.")
	}

}

func Test_MustReturnErrorForSymbolInString(t *testing.T) {

	tld, err := New("com-")

	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err == nil {
		t.Fatal("An error was expected, but an error was not returned.")
	}

	if err.Error() != "The top level domain value can only be letters." {
		t.Error("The following expected error message was not returned: 'The top level domain value can only be letters.'.")
	}

}

func Test_TwoDifferentDomainNameMustNotEqual(t *testing.T) {

	// Act
	tld1, err := New("io")
	tld2, err := New("com")

	// Assert
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	areEqual := tld1.Equals(tld2)
	if areEqual {
		t.Errorf("The Top Level Domains should not equal.\nBase: %s\nComparator: %s", tld1.Value(), tld2.Value())
	}

}
