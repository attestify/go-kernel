package domain_name_test

import (
	"errors"
	"github.com/attestify/go-kernel/error/validation_error"
	"github.com/attestify/go-kernel/uri/domain_name"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path Tests **/

func Test_InstantiateDomainName(t *testing.T) {
	setup(t)
	tld, err := domain_name.New("attestify")

	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	if tld.Value() != "attestify" {
		t.Error("The value must return 'attestify'.")
	}

}

func Test_InstantiateForStringOfExactly255Characters(t *testing.T) {
	setup(t)
	// Arrange - Generate a string of eactly 255 characters
	testString := ""
	for i := 1; i <= 255; i++ {
		testString = testString + "a"
	}

	// Act
	_, err := domain_name.New(testString)

	// Assert
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

}

func Test_InstantiateForStringWithDash(t *testing.T) {
	setup(t)
	// Act
	tld, err := domain_name.New("attestify-site")

	// Assert
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}
	if tld.Value() != "attestify-site" {
		t.Error("The value must return 'attestify-site'.")
	}

}

func Test_TwoSameDomainNameMustEqual(t *testing.T) {
	setup(t)
	// Act
	dn1, err := domain_name.New("attestify")
	dn2, err := domain_name.New("attestify")

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
	setup(t)
	// Arrange & Act
	_, err := domain_name.New("")

	// Assert
	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the
	// error object
	if err == nil {
		t.Fatal("An error was expected, but an error was not returned.")
	}

	expectedMessage := "The domain name value must be at least one (1) character, " +
		"and no greater than two-hundred fifty-five (255) characters."
	if err.Error() != expectedMessage {
		t.Errorf("The exptected error was not returned. \n Actual: %s \n Expected: %s", err.Error(), expectedMessage)
	}

	if !errors.As(err, &validation_error.ValidationError{}) {
		t.Errorf("did not get the epected error of type ValidationError")
	}


}

func Test_MustReturnErrorForStringMoreThan255Characters(t *testing.T) {
	setup(t)
	// Arrange - Generate a string of 256 characters
	testString := ""
	for i := 1; i <= 256; i++ {
		testString = testString + "a"
	}

	// Act
	_, err := domain_name.New(testString)

	// Assert
	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err == nil {
		t.Fatal("An error was expected, but an error was not returned.")
	}

	expectedMessage := "The domain name value must be at least one (1) character, " +
		"and no greater than two-hundred fifty-five (255) characters."
	if err.Error() != expectedMessage {
		t.Errorf("The exptected error was not returned. \n Actual: %s \n Expected: %s", err.Error(), expectedMessage)
	}

	if !errors.As(err, &validation_error.ValidationError{}) {
		t.Errorf("did not get the epected error of type ValidationError")
	}

}

func Test_MustReturnErrorStringStartingWithDash(t *testing.T) {
	setup(t)
	// Arrange & Act
	_, err := domain_name.New("-attestify")

	// Assert
	// Fatal use to end test if an error obejct was not returned because the expessions after this evaluate the error object
	if err == nil {
		t.Fatal("An error was expected, but an error was not returned.")
	}

	expectedMessage := "The domain name can only be ASCII characters and hyphens.  The domain name cannot start with a hyphen."
	if err.Error() != expectedMessage {
		t.Errorf("The exptected error was not returned. \n Actual: %s \n Expected: %s", err.Error(), expectedMessage)
	}

}

func Test_TwoDifferentTopLevelDomainMustNotEqual(t *testing.T) {
	setup(t)
	// Act
	dn1, err := domain_name.New("attestify")
	dn2, err := domain_name.New("billbensing")

	// Assert
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	areEqual := dn1.Equals(dn2)
	if areEqual {
		t.Errorf("The Top Level Domains should not equal.\nBase: %s\nComparator: %s", dn1.Value(), dn2.Value())
	}

}
