package revoke_access_test

import (
	"errors"
	"github.com/attestify/go-kernel/access_control"
	"github.com/attestify/go-kernel/access_control/revoke_access"
	"github.com/attestify/go-kernel/error/internal_error"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path **/

// Given there is a valid RevokeAccessGateway dependency
// When the RevokeAccess usecase is instantiated
// Then there should not be any errors
func Test_Instantiate_RevokeAccess_Successfully(t *testing.T) {
	setup(t)

	// Assemble
	gateway := NewMockRevokeAccessGateway()

	// Act
	usecase := revoke_access.New(&gateway)

	// Assert
	if usecase.HasError() == true {
		t.Error("Expected .HasError() to be false, although it returns true")
	}

	if usecase.Error() != nil {
		t.Errorf("An error was genrated when no eror was exptected.\n Error: %s", usecase.Error())
	}

}

// Given the RevokeAccess usecase is instantiated without error
// When the .Revoke(..) is invoked with proper parameter
// Then there should not be any errors
func Test_Invoke_Revoke_Successfully(t *testing.T) {
	setup(t)

	// Assemble
	gateway := NewMockRevokeAccessGateway()
	usecase := revoke_access.New(&gateway)
	if usecase.HasError() {
		t.Errorf("An error was returned when no error was expected: \n %s", usecase.Error())
	}

	// Act
	var userId int64 = 0
	var resourceId int64 = 1
	permissions := []string{"read"}
	usecase.Revoke(userId, resourceId, permissions)

	// Assert
	if usecase.HasError() {
		t.Errorf("An error was genrated when no eror was exptected.\n Error: %s", usecase.Error())
	}

	if usecase.Error() != nil {
		t.Errorf("An error was genrated when no eror was exptected.\n Error: %s", usecase.Error())
	}

}

/** Sad Path **/

/** Side Effects of Nil RevokeAccessGateway **/

// Given an nil instance of a RevokeAccessGateway is provided
// When the RevokeAccess use case is instantiated
// Then an InternalError should be returned with the text
// "the provided RevokeAccessGateway is nil, please provide a valid instance of an RevokeAccessGateway"
func Test_Instantiate_RevokeAccess_With_Nil_RevokeAccessGateway(t *testing.T) {
	setup(t)

	// Assemble
	var gateway revoke_access.RevokeAccessGateway = nil

	// Act
	usecase := revoke_access.New(gateway)

	// Assert
	if usecase.HasError() != true {
		t.Fatalf("Expected an error, although no error exists.")
	}

	if !errors.As(usecase.Error(), &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}

	actualMessage := usecase.Error().Error()
	expectedMessage := "the provided RevokeAccessGateway is nil, please provide a valid instance of an RevokeAccessGateway"

	if expectedMessage != actualMessage {
		t.Errorf("The actual error message was not the expected error message.\n Expected: %s\n Actual: %s\n", expectedMessage, actualMessage)
	}

}

// NOTE - this tests that error guards function properly
// Given an nil instance of a RevokeAccessGateway is provided
// When .Revoke() is invoked
// Then an InternalError should be returned with the text
// "the provided RevokeAccessGateway is nil, please provide a valid instance of an RevokeAccessGateway"
func Test_Returns_InternalError_With_Nil_RevokeAccessGateway_When_Revoke_Invoked(t *testing.T) {
	setup(t)

	// Assemble
	var gateway revoke_access.RevokeAccessGateway = nil
	var userId int64 = 0
	var resourceId int64 = 1
	permissions := []string{"read"}

	// Act
	usecase := revoke_access.New(gateway)
	usecase.Revoke(userId, resourceId, permissions)

	// Assert
	if usecase.HasError() != true {
		t.Fatalf("Expected an error, although no error exists.")
	}

	if !errors.As(usecase.Error(), &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}

	actualMessage := usecase.Error().Error()
	expectedMessage := "the provided RevokeAccessGateway is nil, please provide a valid instance of an RevokeAccessGateway"

	if expectedMessage != actualMessage {
		t.Errorf("The actual error message was not the expected error message.\n Expected: %s\n Actual: %s\n", expectedMessage, actualMessage)
	}

}

/** Side Effects of RevokeAccessGateway errors **/

// Given the RevokeAccess usecase is instantiated without error
// When .Revoke(...) is invoked with valid parameters
// Then the RevokeAccessGateway returns an InternalError
//  and the RevokeAccess usecase should reflect InternalError returned by the RevokeAccessGateway
func Test_Handle_RevokeAccessGateway_Returns_InternalError_When_Revoke_Invoked(t *testing.T) {
	setup(t)

	// Assemble
	gateway := NewMockRevokeAccessGateway()
	gateway.RevokeAccessGatewayInternalError()
	usecase := revoke_access.New(&gateway)
	if usecase.HasError() {
		t.Fatalf("An error was returned when no error was expected: \n %s", usecase.Error())
	}

	// Act
	var userId int64 = 0
	var resourceId int64 = 1
	permissions := []string{"read"}
	usecase.Revoke(userId, resourceId, permissions)

	// Assert
	if !errors.As(usecase.Error(), &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}

	// This assertion ensure the InternalError is generate when the RevokeAccess.Revoke(...) method is invoked
	// If the message is other than what's here, the InternalError was generated by something else than .Revoke(...)
	// By testing this, we can expect that and InternalError generated  by any other implementation of the
	// RevokeAccessGateway will be properly handled
	actualMessage := usecase.Error().Error()
	expectedMessage := "Error generated from RevokeAccess.Revoke() invocation."
	if expectedMessage != actualMessage {
		t.Errorf("The returned error message was not expected: \n Expected: %s \n Actual %s", expectedMessage, actualMessage)
	}
}

/** Test tooling **/

type MockRevokeAccessGateway struct {
	revokeAccessGatewayInternalError bool
	gatewayError                     error
}

func NewMockRevokeAccessGateway() MockRevokeAccessGateway {
	return MockRevokeAccessGateway{}
}

func (gateway *MockRevokeAccessGateway) Revoke(accessControl access_control.AccessControl) {
	if gateway.revokeAccessGatewayInternalError {
		gateway.gatewayError = internal_error.New("Error generated from RevokeAccess.Revoke() invocation.")
	}
}

func (gateway *MockRevokeAccessGateway) RevokeAccessGatewayInternalError() {
	gateway.revokeAccessGatewayInternalError = true
}

func (gateway MockRevokeAccessGateway) Error() error {
	return gateway.gatewayError
}

func (gateway MockRevokeAccessGateway) HasError() bool {
	return gateway.gatewayError != nil
}


