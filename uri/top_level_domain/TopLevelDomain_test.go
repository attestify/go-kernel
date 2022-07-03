package top_level_domain

import (
	"testing"
)

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

	tld, err := New("")

	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err == nil {
		t.Fatal("An error was expected, but an error was not returned.")
	}

	if err.Error() != "The top level domain value must be atleast one (1) character." {
		t.Error("The following expected error message was not returned: 'The top level domain value must be atleast one (1) character.'.")
	}
	if tld == nil {
		t.Error("Expected an instantiated, empty, TopLeveLDomain object, but got a 'nil' value.")
	}

}

func Test_MustReturnErrorForNumberInString(t *testing.T) {

	tld, err := New("1io")

	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err == nil {
		t.Fatal("An error was expected, but an error was not returned.")
	}

	if err.Error() != "The top level domain value can only be letters." {
		t.Error("The following expected error message was not returned: 'The top level domain value can only be letters.'.")
	}
	if tld == nil {
		t.Error("Expected an instantiated, empty, TopLeveLDomain object, but got a 'nil' value.")
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
	if tld == nil {
		t.Error("Expected an instantiated, empty, TopLeveLDomain object, but got a 'nil' value.")
	}

}
