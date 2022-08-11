package top_level_domain_test

import (
	"errors"
	"github.com/attestify/go-kernel/error/validation_error"
	"github.com/attestify/go-kernel/uri/top_level_domain"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path **/

func Test_InstantiateTopLevelDomain(t *testing.T) {
	setup(t)
	tld, err := top_level_domain.New("io")

	if err != nil {
		t.Error("An error was returned when no error was expected.")
	}
	if tld.Value() != "io" {
		t.Error("The value must return 'io'.")
	}

}

func Test_TwoSameTopLevelDomainMustEqual(t *testing.T) {
	setup(t)
	// Act
	tld1, err := top_level_domain.New("io")
	tld2, err := top_level_domain.New("io")

	// Assert
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	areEqual := tld1.Equals(tld2)
	if !areEqual {
		t.Errorf("The Top Level Domains do not equal.\nBase: %s\nComparator: %s", tld1.Value(), tld2.Value())
	}

}

/** Sad Path **/

func Test_Return_Error_For_Zero_Characters(t *testing.T) {
	setup(t)
	_, err := top_level_domain.New("")

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the error object
	if err == nil {
		t.Fatal("An error was expected, but an error was not returned.")
	}

	if !errors.As(err, &validation_error.ValidationError{}) {
		t.Errorf("did not get the epected error of type ValidationError")
	}

	expectedMessage := "The top level domain value must be at least one (1) character."
	if err.Error() != expectedMessage {
		t.Errorf("The the expected error message was not returned: \n Expected: %s \n Actual: %s", expectedMessage,
			err.Error())
	}
}

func Test_MustReturnErrorForNumberInString(t *testing.T) {
	setup(t)
	_, err := top_level_domain.New("1io")

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the error object
	if err == nil {
		t.Fatal("An error was expected, but an error was not returned.")
	}

	if !errors.As(err, &validation_error.ValidationError{}) {
		t.Errorf("did not get the epected error of type ValidationError")
	}

	if err.Error() != "The top level domain value can only be letters." {
		t.Error("The following expected error message was not returned: 'The top level domain value can only be letters.'.")
	}

}

func Test_MustReturnErrorForSymbolInString(t *testing.T) {
	setup(t)
	_, err := top_level_domain.New("com-")

	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err == nil {
		t.Fatal("An error was expected, but an error was not returned.")
	}

	if !errors.As(err, &validation_error.ValidationError{}) {
		t.Errorf("did not get the epected error of type ValidationError")
	}

	if err.Error() != "The top level domain value can only be letters." {
		t.Error("The following expected error message was not returned: 'The top level domain value can only be letters.'.")
	}

}

func Test_TwoDifferentDomainNameMustNotEqual(t *testing.T) {
	setup(t)
	// Act
	tld1, err := top_level_domain.New("io")
	tld2, err := top_level_domain.New("com")

	// Assert
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	areEqual := tld1.Equals(tld2)
	if areEqual {
		t.Errorf("The Top Level Domains should not equal.\nBase: %s\nComparator: %s", tld1.Value(), tld2.Value())
	}

}
