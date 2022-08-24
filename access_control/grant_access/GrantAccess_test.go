package grant_access_test

import (
	"errors"
	"github.com/attestify/go-kernel/access_control"
	"github.com/attestify/go-kernel/access_control/grant_access"
	"github.com/attestify/go-kernel/error/internal_error"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path **/

// Given an instance of a GrantAccessGateway is provided
// When the GrantAccess use case is instantiated
// Then there should be no error
func Test_Instantiate_AssignRole_Successfully(t *testing.T) {
	setup(t)
	assignRoleGateway := NewAssignRoleGatewayMock()
	usecase := grant_access.New(assignRoleGateway)
	if usecase.HasError() {
		t.Errorf("An error was returned when no error was expected: \n %s", usecase.Error())
	}
}

// Given a valid instance of GrantAccess exists
//  and the user id of "0" is provided,
//  and the entity if of "1" is provided,
//  and the entity of "test-entity" is provided
//  and a permission_list of "write" is provided
// When .Grant(...) is invoked
// Then there should be no error
func Test_Invoke_Assign_Successfully(t *testing.T) {
	setup(t)
	// Assemble
	gateway := NewAssignRoleGatewayMock()
	usecase := grant_access.New(gateway)
	if usecase.HasError() {
		t.Errorf("An error was returned when no error was expected: \n %s", usecase.Error())
	}
	var userId int64 = 0
	var resourceId int64 = 1
	resource := "test-entity"
	permissions := []string{"read"}

	// Act
	usecase.Grant(userId, resourceId, resource, permissions)

	// Assert
	if usecase.HasError() {
		t.Error("an error was returned when no error was expected")
	}
}

/** Sad Path **/

// Given an nil instance of a GrantAccessGateway is provided
// When the GrantAccess use case is instantiated
// Then an InternalError with the text "the provided GrantAccessGateway is nil, please provide a valid instance of an GrantAccessGateway" should be returned
func Test_Instantiate_AssignRole_With_Nil_AssignRoleGateway(t *testing.T) {
	setup(t)
	// Assemble
	var assignRoleGateway grant_access.GrantAccessGateway = nil

	// Act
	usecase := grant_access.New(assignRoleGateway)

	// Assert
	if usecase.HasError() == false {
		t.Fatalf("An error was expected, although no error was returned.")
	}

	if !errors.As(usecase.Error(), &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}

	actualMessage := usecase.Error().Error()
	expectedMessage := "the provided GrantAccessGateway is nil, please provide a valid instance of an GrantAccessGateway"
	if expectedMessage != actualMessage {
		t.Errorf("The returned error message was not expected: \n Expected: %s \n Actual %s", expectedMessage, actualMessage)
	}
}

// Given an nil instance of a GrantAccessGateway is provided
// When .Grant() is invoked
// Then an InternalError with the text "the provided GrantAccessGateway is nil, please provide a valid instance of an GrantAccessGateway" should be returned
func Test_Returns_InternalError_With_Nil_AssignRoleGateway_When_Grant_Invoked(t *testing.T) {
	setup(t)
	// Assemble
	var assignRoleGateway grant_access.GrantAccessGateway = nil

	// Act
	usecase := grant_access.New(assignRoleGateway)
	var userId int64 = 0
	var resourceId int64 = 1
	resource := "test-entity"
	permissions := []string{"read"}

	// Act
	usecase.Grant(userId, resourceId, resource, permissions)

	// Assert
	if usecase.HasError() == false {
		t.Fatalf("An error was expected, although no error was returned.")
	}

	if !errors.As(usecase.Error(), &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}

	actualMessage := usecase.Error().Error()
	expectedMessage := "the provided GrantAccessGateway is nil, please provide a valid instance of an GrantAccessGateway"
	if expectedMessage != actualMessage {
		t.Errorf("The returned error message was not expected: \n Expected: %s \n Actual %s", expectedMessage, actualMessage)
	}
}

// Given we expect the GrantAccessGateway to return InternalError
// When .Grant(...) is invoked with the proper arguments
// Then the GrantAccess use case must return an InternalError
func Test_Invoke_Assign_Returns_InternalError(t *testing.T) {
	setup(t)
	// Assemble
	gateway := NewAssignRoleGatewayMock()
	gateway.ReturnInternalError()
	usecase := grant_access.New(gateway)
	if usecase.HasError() {
		t.Errorf("An error was returned when no error was expected: \n %s", usecase.Error())
	}
	var userId int64 = 0
	var resourceId int64 = 1
	resource := "test-entity"
	permissions := []string{"read"}

	// Act
	usecase.Grant(userId, resourceId, resource, permissions)

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

func (gateway AssignRoleGatewayMock) Grant(accessControl access_control.AccessControl) {}

func (gateway *AssignRoleGatewayMock) ReturnInternalError() {
	gateway.gatewayError =  internal_error.New("some internal error")
}

func (gateway AssignRoleGatewayMock) HasError() bool {
	return gateway.gatewayError != nil
}

func (gateway AssignRoleGatewayMock) Error() error {
	return gateway.gatewayError
}
