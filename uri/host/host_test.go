package host

import (
	"github.com/attestify/go-kernel/uri/registered_name"
	"testing"
)

/** Happy Path Tests **/

// Instantiate a Host object using the NewFromRegisteredName constructor without an error
// and expect the .Value() to be "attestify.io" and .HostType() to be "reg-name".
func Test_InstantiateHost(t *testing.T) {

	regname, _ := registered_name.NewFromString("attestify.io")
	host, err := NewFromRegisteredName(*regname)

	// Fatal use to end test if an error object was not returned because the expressions after this evaluate the error object
	if err != nil {
		t.Fatalf("An error was returned when no error was expected.\n Error: %s ", err.Error())
	}

	actualValue := host.Value()
	actualType  := host.HostType()

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