package modify_access_test

import (
	"errors"
	"github.com/attestify/go-kernel/access_control"
	"github.com/attestify/go-kernel/access_control/modify_access"
	"github.com/attestify/go-kernel/error/internal_error"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path **/

// Given an instance of a ModifyAccessGateway is provided
// When the ModifyAccess use case is instantiated
// Then there should be no error
func Test_Instantiate_AssignRole_Successfully(t *testing.T) {
	setup(t)
	assignRoleGateway := NewAssignRoleGatewayMock()
	usecase := modify_access.New(assignRoleGateway)
	if usecase.HasError() {
		t.Errorf("An error was returned when no error was expected: \n %s", usecase.Error())
	}
}

// Given a valid instance of ModifyAccess exists
//  and the user id of "0" is provided,
//  and the entity if of "1" is provided,
//  and the entity of "test-entity" is provided
//  and a permission_list of "write" is provided
// When .Modify(...) is invoked
// Then there should be no error
func Test_Invoke_Assign_Successfully(t *testing.T) {
	setup(t)
	// Assemble
	gateway := NewAssignRoleGatewayMock()
	usecase := modify_access.New(gateway)
	if usecase.HasError() {
		t.Errorf("An error was returned when no error was expected: \n %s", usecase.Error())
	}
	var userId int64 = 0
	var resourceId int64 = 1
	permissions := []string{"read"}

	// Act
	usecase.Modify(userId, resourceId, permissions)

	// Assert
	if usecase.HasError() {
		t.Error("an error was returned when no error was expected")
	}
}

/** Sad Path **/

// Given an nil instance of a ModifyAccessGateway is provided
// When the ModifyAccess use case is instantiated
// Then an InternalError with the text "the provided ModifyAccessGateway is nil, please provide a valid instance of an ModifyAccessGateway" should be returned
func Test_Instantiate_AssignRole_With_Nil_AssignRoleGateway(t *testing.T) {
	setup(t)
	// Assemble
	var assignRoleGateway modify_access.ModifyAccessGateway = nil

	// Act
	usecase := modify_access.New(assignRoleGateway)

	// Assert
	if usecase.HasError() == false {
		t.Fatalf("An error was expected, although no error was returned.")
	}

	if !errors.As(usecase.Error(), &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}

	actualMessage := usecase.Error().Error()
	expectedMessage := "the provided ModifyAccessGateway is nil, please provide a valid instance of an ModifyAccessGateway"
	if expectedMessage != actualMessage {
		t.Errorf("The returned error message was not expected: \n Expected: %s \n Actual %s", expectedMessage, actualMessage)
	}
}

// Given an nil instance of a ModifyAccessGateway is provided
// When .Modify() is invoked
// Then an InternalError with the text "the provided ModifyAccessGateway is nil, please provide a valid instance of an ModifyAccessGateway" should be returned
func Test_Returns_InternalError_With_Nil_AssignRoleGateway_When_Modify_Invoked(t *testing.T) {
	setup(t)
	// Assemble
	var assignRoleGateway modify_access.ModifyAccessGateway = nil

	// Act
	usecase := modify_access.New(assignRoleGateway)
	var userId int64 = 0
	var resourceId int64 = 1
	permissions := []string{"read"}

	// Act
	usecase.Modify(userId, resourceId, permissions)

	// Assert
	if usecase.HasError() == false {
		t.Fatalf("An error was expected, although no error was returned.")
	}

	if !errors.As(usecase.Error(), &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}

	actualMessage := usecase.Error().Error()
	expectedMessage := "the provided ModifyAccessGateway is nil, please provide a valid instance of an ModifyAccessGateway"
	if expectedMessage != actualMessage {
		t.Errorf("The returned error message was not expected: \n Expected: %s \n Actual %s", expectedMessage, actualMessage)
	}
}

// Given we expect the ModifyAccessGateway to return InternalError
// When .Modify(...) is invoked with the proper arguments
// Then the ModifyAccess use case must return an InternalError
func Test_Invoke_Assign_Returns_InternalError(t *testing.T) {
	setup(t)
	// Assemble
	gateway := NewAssignRoleGatewayMock()
	gateway.ReturnInternalError()
	usecase := modify_access.New(gateway)
	if usecase.HasError() {
		t.Errorf("An error was returned when no error was expected: \n %s", usecase.Error())
	}
	var userId int64 = 0
	var resourceId int64 = 1
	permissions := []string{"read"}

	// Act
	usecase.Modify(userId, resourceId, permissions)

	// Assert
	if !errors.As(usecase.Error(), &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}
}

type AssignRoleGatewayMock struct {
	gatewayError error
}

func NewAssignRoleGatewayMock() AssignRoleGatewayMock {
	return AssignRoleGatewayMock{}
}

func (gateway AssignRoleGatewayMock) Modify(accessControl access_control.AccessControl) {}

func (gateway *AssignRoleGatewayMock) ReturnInternalError() {
	gateway.gatewayError = internal_error.New("some internal error")
}

func (gateway AssignRoleGatewayMock) HasError() bool {
	return gateway.gatewayError != nil
}

func (gateway AssignRoleGatewayMock) Error() error {
	return gateway.gatewayError
}
