package registered_name_test

import (
	"errors"
	"github.com/attestify/go-kernel/error/validation_error"
	"github.com/attestify/go-kernel/uri/registered_name"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path Tests **/

// Instantiate a RegisteredName object using a top level domain of "io", and a doming name
// "attestify", without receiving an error. Validate state of the object by expecting the
// .Value() to return "attestify.io"
func Test_InstantiateRegisteredName(t *testing.T) {
	setup(t)
	registeredName, err := registered_name.New("io", "attestify")

	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	actualValue := registeredName.Value()
	expectedValue := "attestify.io"
	if expectedValue != actualValue {
		t.Errorf("Did not return the expected value.\nActual: %s\nExpected: %s",
			registeredName.Value(),
			expectedValue)
	}
}

func Test_InstantiateRegisteredNameFromString(t *testing.T) {
	setup(t)
	registeredName, err := registered_name.NewFromString("attestify.io")

	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	actualValue := registeredName.Value()
	expectedValue := "attestify.io"
	if expectedValue != actualValue {
		t.Errorf("Did not return the expected value.\nActual: %s\nExpected: %s",
			registeredName.Value(),
			expectedValue)
	}
}

func Test_InstantiateRegisteredNameFromStringWithSubDomain(t *testing.T) {
	setup(t)
	registeredName, err := registered_name.NewFromString("subdomain.attestify.io")

	// Fatal use to end test if an error objet was not returned because the expressions after this evaluate the error object
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	actualValue := registeredName.Value()
	expectedValue := "subdomain.attestify.io"
	if expectedValue != actualValue {
		t.Errorf("Did not return the expected value.\nActual: %s\nExpected: %s",
			registeredName.Value(),
			expectedValue)
	}
}

func Test_TwoSameRegisteredNameMustEqual(t *testing.T) {
	setup(t)
	// Act
	rn1, err := registered_name.New("io", "attestify")
	rn2, err := registered_name.New("io", "attestify")

	// Assert
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	areEqual := rn1.Equals(rn2)
	if !areEqual {
		t.Errorf("The Registered Names do not equal.\nBase: %s\nComparator: %s", rn1.Value(), rn2.Value())
	}
}

/** Sad Path Tests **/

func Test_TwoDifferentTopLevelDomainMustNotEqual(t *testing.T) {
	setup(t)
	// Act
	rn1, err := registered_name.New("io", "attestify")
	rn2, err := registered_name.New("com", "attestify")

	// Assert
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	areEqual := rn1.Equals(rn2)
	if areEqual {
		t.Errorf("The Registered Names should not equal.\nBase: %s\nComparator: %s", rn1.Value(), rn2.Value())
	}

}

func Test_Handle_Top_Level_Domain_Error(t *testing.T) {
	setup(t)
	// Arrange & Act - provide a bad top level domain
	_, err := registered_name.New("bad!", "attestify")

	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err == nil {
		t.Fatalf("An error was expected, but no error was returned")
	}

	if !errors.As(err, &validation_error.ValidationError{}) {
		t.Errorf("did not get the epected error of type ValidationError")
	}
}

func Test_Handle_Domain_Name_Error(t *testing.T) {
	setup(t)
	// Arrange & Act - provide a bad domain name
	_, err := registered_name.New("io", "-attestify")

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the
	//error object
	if err == nil {
		t.Fatalf("An error was expected, but no error was returned")
	}

	if !errors.As(err, &validation_error.ValidationError{}) {
		t.Errorf("did not get the epected error of type ValidationError")
	}

}
