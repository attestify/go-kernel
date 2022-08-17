package grant_access_test

import (
	"errors"
	"github.com/attestify/go-kernel/access_control"
	"github.com/attestify/go-kernel/access_control/grant_access"
	"github.com/attestify/go-kernel/error/internal_error"
	"github.com/attestify/go-kernel/identity/id"
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
	_, err := grant_access.New(assignRoleGateway)
	if err != nil {
		t.Errorf("An error was returned when no error was expected: \n %s", err.Error())
	}
}

// Given a valid instance of GrantAccess exists
//   and the user id of "[x]" is provided,
//   and the entity if of "[y]" is provided,
//   and the entity of "test-entity" is provided
// When .Grant(...) is invoked
// Then there should be no error
func Test_Invoke_Assign_Successfully(t *testing.T) {
	setup(t)
	// Assemble
	gateway := NewAssignRoleGatewayMock()
	usecase, err := grant_access.New(gateway)
	if err != nil {
		t.Errorf("An error was returned when no error was expected: \n %s", err.Error())
	}
	var userId int64 = 0
	var entityId int64 = 1
	entity := "test-entity"

	// Act
	err = usecase.Grant(userId, entityId, entity)

	// Assert
	if err != nil {
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
	var assignRoleGateway access_control.GrantAccessGateway = nil

	// Act
	_, err := grant_access.New(assignRoleGateway)

	// Assert
	if err == nil {
		t.Fatalf("An error was expected, although no error was returned.")
	}

	if !errors.As(err, &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}

	actualMessage := err.Error()
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
	usecase, err := grant_access.New(gateway)
	if err != nil {
		t.Errorf("An error was returned when no error was expected: \n %s", err.Error())
	}
	var userId int64 = 0
	var entityId int64 = 1
	entity := "test-entity"

	// Act
	err = usecase.Grant(userId, entityId, entity)

	// Assert
	if !errors.As(err, &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}
}

type AssignRoleGatewayMock struct {
	internalError bool
}

func NewAssignRoleGatewayMock() AssignRoleGatewayMock {
	return AssignRoleGatewayMock{
		internalError: false,
	}
}

func (gateway AssignRoleGatewayMock) Grant(userId id.Id, entityId id.Id, entity string) error {
	if gateway.internalError {
		return internal_error.New("some internal error")
	}
	return nil
}

func (gateway *AssignRoleGatewayMock) ReturnInternalError() {
	gateway.internalError = true
}
