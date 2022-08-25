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

// Given a non-nil RevokeAccessGateway is provided
// When a RevokeAccess use case is instantiated
// Then no error should be returned
//  and .HasError() should be false
func Test_Instantiate_RevokeAccess_Successfully(t *testing.T) {
	setup(t)

	// Assemble
	var gateway revoke_access.RevokeAccessGateway = NewMockRevokeAccessGateway()

	// Act
	usecase := revoke_access.New(gateway)

	// Assert
	if usecase.HasError() == true {
		t.Error("Expected .HasError() to be false, although it returns true")
	}

	if usecase.Error() != nil {
		t.Errorf("An error was genrated when no eror was exptected.\n Error: %s", usecase.Error())
	}

}

// Given a non-nil RevokeAccessGateway is provided
//  and a user id of 0
//  and a resourceId of 1
//  and a resource of "test-resource"
//  and a permission_list of "write"
// When a RevokeAccess.Revoke(...) is invoked
// Then there should be no error
func Test_RevokeAccess_Successfully(t *testing.T) {
	setup(t)

	// Assemble
	var gateway revoke_access.RevokeAccessGateway = NewMockRevokeAccessGateway()
	var userId int64 = 0
	var resourceId int64 = 1
	resource := "test-resource"
	permissions := []string{"read"}

	// Act
	usecase := revoke_access.New(gateway)
	usecase.Revoke(userId, resourceId, resource, permissions)

	// Assert
	if usecase.HasError() {
		t.Errorf("An error was genrated when no eror was exptected.\n Error: %s", usecase.Error())
	}

}

/** Sad Path **/

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
	resource := "test-resource"
	permissions := []string{"read"}

	// Act
	usecase := revoke_access.New(gateway)
	usecase.Revoke(userId, resourceId, resource, permissions)

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

// Given we expect the RevokeAccessGateway to return InternalError
// When .Revoke(...) is invoked with the proper arguments
// Then the RevokeAccess use case must return an InternalError
func Test_Invoke_Assign_Returns_InternalError(t *testing.T) {
	setup(t)
	// Assemble
	gateway := NewMockRevokeAccessGateway()
	gateway.ReturnInternalError()
	usecase := revoke_access.New(gateway)

	if usecase.HasError() {
		t.Errorf("An error was returned when no error was expected: \n %s", usecase.Error())
	}

	var userId int64 = 0
	var resourceId int64 = 1
	resource := "test-entity"
	permissions := []string{"read"}

	// Act
	usecase.Revoke(userId, resourceId, resource, permissions)

	// Assert
	if !errors.As(usecase.Error(), &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}
}

/** Test tooling **/

type MockRevokeAccessGateway struct {
	gatewayError error
}

func NewMockRevokeAccessGateway() MockRevokeAccessGateway {
	return MockRevokeAccessGateway{}
}

func (gateway MockRevokeAccessGateway) Revoke(accessControl access_control.AccessControl) {}

func (gateway MockRevokeAccessGateway) Error() error {
	return gateway.gatewayError
}

func (gateway MockRevokeAccessGateway) HasError() bool {
	return gateway.gatewayError != nil
}

func (gateway *MockRevokeAccessGateway) ReturnInternalError() {
	gateway.gatewayError = internal_error.New("Some message")
}
