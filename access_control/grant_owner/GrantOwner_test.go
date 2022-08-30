package grant_owner_test

import (
	"errors"
	"github.com/attestify/go-kernel/access_control"
	"github.com/attestify/go-kernel/access_control/grant_owner"
	"github.com/attestify/go-kernel/error/internal_error"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path **/

// Given there is a valid GrantOwnerGateway dependency
// When the GrantOwner usecase is instantiated
// Then there should not be any errors
func Test_Instantiate_GrantOwner_Successfully(t *testing.T) {
	setup(t)
	// Assemble
	gateway := NewGrantOwnerGatewayMock()

	//Act
	usecase := grant_owner.New(&gateway)

	//Assert
	if usecase.HasError() {
		t.Errorf("An unexpected error occurred.\n Error: %s\n", usecase.Error())
	}

}

// Given the GrantOwner usecase is instantiated without error
// When the .Grant(..) is invoked with a userId of 0 and a resourceId of 1
// Then there should not be any errors
func Test_Invoke_Grant_Successfully(t *testing.T) {
	setup(t)
	// Assemble
	gateway := NewGrantOwnerGatewayMock()
	var userId int64 = 0
	var resourceId int64 = 1

	//Act
	usecase := grant_owner.New(&gateway)
	usecase.Grant(userId, resourceId)

	//Assert
	if usecase.HasError() {
		t.Errorf("An unexpected error occurred.\n Error: %s\n", usecase.Error())
	}

}

/** Sad Path **/

// Given there is a valid GrantOwnerGateway dependency
// When the GrantOwner usecase is instantiated
// Then there should be an error
//  with the error message "the provided GrantOwnerGateway is nil, please provide a valid instance of an GrantOwnerGateway"
func Test_Instantiate_GrantOwner_With_Bad_GrantOwnerGateway(t *testing.T) {
	setup(t)
	// Assemble
	var gateway grant_owner.GrantOwnerGateway = nil

	//Act
	usecase := grant_owner.New(gateway)

	//Assert
	if usecase.HasError() == false {
		t.Fatal("An error was expected, but no error was returned")
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

// Given the GrantOwner usecase is instantiated without error
// When the .Grant(..) is invoked with a userId of 0 and a resourceId of 1
//	And the GrantOwnerGateway occurs an InternalError
// Then the GrantOwner usecase should reflect InternalError
func Test_Handle_InternalError_Invoking_GrantOwnerGateway(t *testing.T) {
	setup(t)
	// Assemble
	gateway := NewGrantOwnerGatewayMock()
	gateway.GenerateInternalError()
	usecase := grant_owner.New(&gateway)
	if usecase.HasError() == true {
		t.Fatalf("An error was thrown when non was expected.\n Error: %s\n", usecase.Error())
	}

	//Act
	var userId int64 = 0
	var resourceId int64 = 1
	usecase.Grant(userId, resourceId)

	//Assert
	if usecase.HasError() != true {
		t.Fatal("An error was expected, but no error was returned")
	}

	if !errors.As(usecase.Error(), &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}

}

// todo - test gate for setAccessControl
// todo - test gate for grantOwner

/** Testing Tools **/

type GrantOwnerGatewayMock struct {
	generateInternalError bool
	mockError error
}

func NewGrantOwnerGatewayMock()  GrantOwnerGatewayMock {
	return GrantOwnerGatewayMock{}
}

func (mock *GrantOwnerGatewayMock) Grant(accessControl access_control.AccessControl)  {}

func (mock GrantOwnerGatewayMock) HasError() bool {
	return mock.mockError != nil
}

func (mock GrantOwnerGatewayMock) Error() error {
	return mock.mockError
}

func (mock *GrantOwnerGatewayMock) GenerateInternalError() {
	mock.mockError = internal_error.New("Internal error when .Grant(...) was invoked")
}