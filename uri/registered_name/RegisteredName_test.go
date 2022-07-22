package registered_name

import (
	"testing"
)

/** Happy Path Tests **/

// Instantiate a RegisteredName object using a top level domain of "io", and a doming name
// "attestify", without receiving an error. Validate state of the object by expecting the
// .Value() to return "attestify.io"
func Test_InstantiateRegisteredName(t *testing.T) {

	registeredName, err := New("io", "attestify")

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

	registeredName, err := NewFromString("attestify.io")

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

	registeredName, err := NewFromString("subdomain.attestify.io")

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

	// Act
	rn1, err := New("io", "attestify")
	rn2, err := New("io", "attestify")

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

	// Act
	rn1, err := New("io", "attestify")
	rn2, err := New("com", "attestify")

	// Assert
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	areEqual := rn1.Equals(rn2)
	if areEqual {
		t.Errorf("The Registered Names should not equal.\nBase: %s\nComparator: %s", rn1.Value(), rn2.Value())
	}

}

func Test_HandleTopLevelDomainError(t *testing.T) {

	// Arrange & Act - provide a bad top level domain
	registeredName, err := New("bad!", "attestify")

	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err == nil {
		t.Fatalf("An error was expected, but no error was returned")
	}

	if registeredName == nil {
		t.Error("Expected an instantiated, empty, RegisteredName object, but got a 'nil' value.")
	}


}

func Test_HandleDomainNameError(t *testing.T) {

	// Arrange & Act - provide a bad domain name
	registeredName, err := New("io", "-attestify")

	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err == nil {
		t.Fatalf("An error was expected, but no error was returned")
	}

	if registeredName == nil {
		t.Error("Expected an instantiated, empty, RegisteredName object, but got a 'nil' value.")
	}
}
