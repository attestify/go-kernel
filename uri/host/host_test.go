package host

import (
	"errors"
	"github.com/attestify/go-kernel/error/validation_error"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path Tests **/

// Instantiate a Host object using the NewFromRegisteredName constructor without an error
// and expect the .Value() to be "attestify.io" and .HostType() to be "reg-name".
func Test_Instantiate_Host_Successfully(t *testing.T) {
	setup(t)
	host, err := NewFromRegisteredName("attestify.io")

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the error object
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	actualValue := host.Value()
	actualType := host.HostType()

	expectedValue := "attestify.io"
	expectedType := "reg-name"

	if expectedValue != actualValue {
		t.Errorf("Did not return the expected value.\nActual: %s\nExpected: %s",
			host.Value(),
			expectedValue)
	}

	if expectedType != actualType {
		t.Errorf("Did not return the expected type.\nActual: %s\nExpected: %s",
			host.HostType(),
			expectedType)
	}
}

/** Sad Path Tests **/

func Test_Instantiate_Host_With_Bad_Registered_Name(t *testing.T) {
	setup(t)
	badRegisteredName := "attestify.io1"
	_, err := NewFromRegisteredName(badRegisteredName)

	if err == nil {
		t.Fatalf("An error is expected, but no error was returned.")
	}

	if !errors.As(err, &validation_error.ValidationError{}) {
		t.Errorf("did not get the epected error of type ValidationError")
	}

}
