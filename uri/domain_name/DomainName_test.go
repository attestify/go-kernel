package domain_name

import (
	"testing"
)

/** Happy Path Tests **/

func Test_InstantiateDomainName(t *testing.T) {

	tld, err := New("attestify")

	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	if tld.Value() != "attestify" {
		t.Error("The value must return 'attestify'.")
	}

}

func Test_InstantiateForStringOfExactly255Characters(t *testing.T) {

	// Arrange - Generate a string of eactly 255 characters
	testString := ""
	for i := 1; i <= 255; i++ {
		testString = testString + "a"
	}

	// Act
	_, err := New(testString)

	// Assert
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

}

func Test_InstantiateForStringWithDash(t *testing.T) {

	// Act
	tld, err := New("attestify-site")

	// Assert
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}
	if tld.Value() != "attestify-site" {
		t.Error("The value must return 'attestify-site'.")
	}

}

func Test_TwoSameDomainNameMustEqual(t *testing.T) {

	// Act
	dn1, err := New("attestify")
	dn2, err := New("attestify")

	// Assert
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	areEqual := dn1.Equals(dn2)
	if !areEqual {
		t.Errorf("The Top Level Domains do not equal.\nBase: %s\nComparator: %s", dn1.Value(), dn2.Value())
	}

}

/** Sad Path Tests **/

func Test_MustReturnErrorForEmptyString(t *testing.T) {

	// Arrange & Act
	tld, err := New("")

	// Assert
	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err == nil {
		t.Fatal("An error was expected, but an error was not returned.")
	}

	expectedMessage := "The domain name value must be atleast one (1) character, and no greather than two-hundred fifty-five (255) characters."
	if err.Error() != expectedMessage {
		t.Errorf("The exptected error was not returned. \n Actual: %s \n Expected: %s", err.Error(), expectedMessage)
	}
	if tld == nil {
		t.Error("Expected an instantiated, empty, DomainName object, but got a 'nil' value.")
	}

}

func Test_MustReturnErrorForStringMoreThan255Characters(t *testing.T) {

	// Arrange - Generate a string of 256 characters
	testString := ""
	for i := 1; i <= 256; i++ {
		testString = testString + "a"
	}

	// Act
	tld, err := New(testString)

	// Assert
	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err == nil {
		t.Fatal("An error was expected, but an error was not returned.")
	}

	expectedMessage := "The domain name value must be atleast one (1) character, and no greather than two-hundred fifty-five (255) characters."
	if err.Error() != expectedMessage {
		t.Errorf("The exptected error was not returned. \n Actual: %s \n Expected: %s", err.Error(), expectedMessage)
	}
	if tld == nil {
		t.Error("Expected an instantiated, empty, TopLeveLDomain object, but got a 'nil' value.")
	}

}

func Test_MustReturnErrorStringStartingWithDash(t *testing.T) {

	// Arrange & Act
	tld, err := New("-attestify")

	// Assert
	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err == nil {
		t.Fatal("An error was expected, but an error was not returned.")
	}

	expectedMessage := "The domain name can only be ASCII characters and hyphens.  The domain name cannot start with a hyphen."
	if err.Error() != expectedMessage {
		t.Errorf("The exptected error was not returned. \n Actual: %s \n Expected: %s", err.Error(), expectedMessage)
	}
	if tld == nil {
		t.Error("Expected an instantiated, empty, TopLeveLDomain object, but got a 'nil' value.")
	}

}

func Test_TwoDifferentTopLevelDomainMustNotEqual(t *testing.T) {

	// Act
	dn1, err := New("attestify")
	dn2, err := New("billbensing")

	// Assert
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	areEqual := dn1.Equals(dn2)
	if areEqual {
		t.Errorf("The Top Level Domains should not equal.\nBase: %s\nComparator: %s", dn1.Value(), dn2.Value())
	}

}
