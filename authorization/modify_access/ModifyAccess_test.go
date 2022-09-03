package modify_access_test

import (
	"errors"
	"github.com/attestify/go-kernel/authorization"
	"github.com/attestify/go-kernel/authorization/modify_access"
	"github.com/attestify/go-kernel/error/internal_error"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path **/

// Given there is a valid ModifyAccessGateway dependency
// When the ModifyAccess usecase is instantiated
// Then there should not be any errors
func Test_Instantiate_ModifyAccess_Successfully(t *testing.T) {
	setup(t)

	// Assemble
	assignRoleGateway := NewMockModifyAccessGateway()

	// Act
	usecase := modify_access.New(&assignRoleGateway)

	// Assert
	if usecase.HasError() {
		t.Errorf("An error was returned when no error was expected: \n %s", usecase.Error())
	}

	if usecase.Error() != nil {
		t.Errorf("An error was genrated when no eror was exptected.\n Error: %s", usecase.Error())
	}
}

// Given the ModifyAccess usecase is instantiated without error
// When .Modify(...) is invoked  with the user id of "0",
//  and an entity of "1" ,
//  and a permission_list of "read" is provided
// Then there should be no error
func Test_Invoke_Modify_Successfully(t *testing.T) {
	setup(t)
	// Assemble
	gateway := NewMockModifyAccessGateway()
	usecase := modify_access.New(&gateway)
	if usecase.HasError() {
		t.Errorf("An error was returned when no error was expected: \n %s", usecase.Error())
	}


	// Act
	var userId int64 = 0
	var resourceId int64 = 1
	permissions := []string{"read"}
	usecase.Modify(userId, resourceId, permissions)

	// Assert
	if usecase.HasError() {
		t.Error("an error was returned when no error was expected")
	}

	if usecase.Error() != nil {
		t.Errorf("An error was genrated when no eror was exptected.\n Error: %s", usecase.Error())
	}
}

/** Sad Path **/

/** Side Effects of nil ModifyAccessGateway **/

// Given there is a nil ModifyAccessGateway dependency
// When the ModifyAccess usecase is instantiated
// Then there should be an InternalError with the message:
//   "the provided ModifyAccessGateway is nil, please provide a valid instance of an ModifyAccessGateway" should be returned
func Test_Handle_InternalError_With_Nil_ModifyAccessGateway_When_Instantiated(t *testing.T) {
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

	// This assertion ensure the InternalError is generated from the nil gateway check
	actualMessage := usecase.Error().Error()
	expectedMessage := "the provided ModifyAccessGateway is nil, please provide a valid instance of an ModifyAccessGateway"
	if expectedMessage != actualMessage {
		t.Errorf("The returned error message was not expected: \n Expected: %s \n Actual %s", expectedMessage, actualMessage)
	}
}

// NOTE - this tests that error guards function properly
// Given the ModifyAccess usecase is instantiated with a nil ModifyAccessGateway dependency
// When .Modify() is invoked with valid parameters
// Then there should be an InternalError with the message:
//  "the provided ModifyAccessGateway is nil, please provide a valid instance of an ModifyAccessGateway" should be returned
func Test_Returns_InternalError_With_Nil_ModifyAccessGateway_When_Modify_Invoked(t *testing.T) {
	setup(t)

	// Assemble
	var assignRoleGateway modify_access.ModifyAccessGateway = nil
	var userId int64 = 0
	var resourceId int64 = 1
	permissions := []string{"read"}

	// Act
	usecase := modify_access.New(assignRoleGateway)
	usecase.Modify(userId, resourceId, permissions)

	// Assert
	if usecase.HasError() == false {
		t.Fatalf("An error was expected, although no error was returned.")
	}

	if !errors.As(usecase.Error(), &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}

	// This assertion ensure the InternalError is generated from the nil gateway check
	actualMessage := usecase.Error().Error()
	expectedMessage := "the provided ModifyAccessGateway is nil, please provide a valid instance of an ModifyAccessGateway"
	if expectedMessage != actualMessage {
		t.Errorf("The returned error message was not expected: \n Expected: %s \n Actual %s", expectedMessage, actualMessage)
	}
}

/** Side Effects of ModifyAccessGateway errors **/

// Given the ModifyAccess usecase is instantiated without error
// When .Modify(...) is invoked with the valid parameters
// Then the ModifyAccessGateway returns an InternalError
//  and the ModifyAccess usecase should reflect InternalError returned by the ModifyAccessGateway
func Test_Handle_ModifyAccessGateway_Returns_InternalError_When_Modify_Invoked(t *testing.T) {
	setup(t)
	// Assemble
	gateway := NewMockModifyAccessGateway()
	gateway.ModifyAccessGatewayInternalError()
	usecase := modify_access.New(&gateway)
	if usecase.HasError() {
		t.Errorf("An error was returned when no error was expected: \n %s", usecase.Error())
	}

	// Act
	var userId int64 = 0
	var resourceId int64 = 1
	permissions := []string{"read"}
	usecase.Modify(userId, resourceId, permissions)

	// Assert
	if !errors.As(usecase.Error(), &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}

	// This assertion ensure the InternalError is generate when the ModifyAccess.Modify(...) method is invoked
	// If the message is other than what's here, the InternalError was generated by something else than .Modify(...)
	// By testing this, we can expect that and InternalError generated  by any other implementation of the
	// ModifyAccessGateway will be properly handled
	actualMessage := usecase.Error().Error()
	expectedMessage := "Error generated from ModifyAccess.Modify() invocation."
	if expectedMessage != actualMessage {
		t.Errorf("The returned error message was not expected: \n Expected: %s \n Actual %s", expectedMessage, actualMessage)
	}
}

/** Testing Tools **/

type MockModifyAccessGateway struct {
	modifyAccessGatewayInternalError bool
	mockError error
}

func NewMockModifyAccessGateway() MockModifyAccessGateway {
	return MockModifyAccessGateway{}
}

func (gateway *MockModifyAccessGateway) Modify(accessControl authorization.AccessControl) {
	if gateway.modifyAccessGatewayInternalError {
		gateway.mockError = internal_error.New("Error generated from ModifyAccess.Modify() invocation.")
	}
}

func (gateway *MockModifyAccessGateway) ModifyAccessGatewayInternalError () {
	gateway.modifyAccessGatewayInternalError = true
}

func (gateway MockModifyAccessGateway) HasError() bool {
	return gateway.mockError != nil
}

func (gateway MockModifyAccessGateway) Error() error {
	return gateway.mockError
}