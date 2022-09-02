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
	gateway := NewMockGrantOwnerGateway()

	//Act
	usecase := grant_owner.New(&gateway)

	//Assert
	if usecase.HasError() {
		t.Errorf("An unexpected error occurred.\n Error: %s\n", usecase.Error())
	}

	if usecase.Error() != nil {
		t.Errorf("An error was genrated when no eror was exptected.\n Error: %s", usecase.Error())
	}

}

// Given the GrantOwner usecase is instantiated without error
// When the .Grant(..) is invoked with a userId of 0 and a resourceId of 1
// Then there should not be any errors
func Test_Invoke_Grant_Successfully(t *testing.T) {
	setup(t)

	// Assemble
	gateway := NewMockGrantOwnerGateway()
	usecase := grant_owner.New(&gateway)
	if usecase.HasError() {
		t.Errorf("An error was returned when no error was expected: \n %s", usecase.Error())
	}

	//Act
	var userId int64 = 0
	var resourceId int64 = 1
	usecase.Grant(userId, resourceId)

	//Assert
	if usecase.HasError() {
		t.Errorf("An unexpected error occurred.\n Error: %s\n", usecase.Error())
	}

	if usecase.Error() != nil {
		t.Errorf("An error was genrated when no eror was exptected.\n Error: %s", usecase.Error())
	}

}

/** Sad Path **/

/** Side Effects of nil GrantOwnerGateway **/

// Given there is a nil GrantOwnerGateway dependency
// When the GrantOwner usecase is instantiated
// Then there should be an InternalError with the error message:
//   "the provided GrantOwnerGateway is nil, please provide a valid instance of an GrantOwnerGateway"
func Test_Handle_InternalError_With_Nil_GrantOwnerGateway_When_Instantiated(t *testing.T) {
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

	// This assertion ensure the InternalError is generated from the nil gateway check
	actualMessage := usecase.Error().Error()
	expectedMessage := "the provided ModifyAccessGateway is nil, please provide a valid instance of an ModifyAccessGateway"
	if expectedMessage != actualMessage {
		t.Errorf("The returned error message was not expected: \n Expected: %s \n Actual %s", expectedMessage, actualMessage)
	}

}

// NOTE - this tests that error guards function properly
// Given the GrantOwner usecase is instantiated with a nil GrantOwnerGateway dependency
// When the GrantOwner.Grant(...) is invoked with valid parameters
// Then there should be an InternalError with the error message:
//   "the provided GrantOwnerGateway is nil, please provide a valid instance of an GrantOwnerGateway"
func Test_Returns_InternalError_With_Nil_GrantOwnerGateway_When_Grant_Invoked(t *testing.T) {
	setup(t)
	// Assemble
	var gateway grant_owner.GrantOwnerGateway = nil
	var userId int64 = 0
	var resourceId int64 = 1

	//Act
	usecase := grant_owner.New(gateway)
	usecase.Grant(userId, resourceId)

	//Assert
	if usecase.HasError() == false {
		t.Fatal("An error was expected, but no error was returned")
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

/** Side Effects of GrantAllGateway errors **/

// Given the GrantOwner usecase is instantiated without error
// When .Grant(...) is invoked with valid parameters
// Then  the GrantOwnerGateway returns an InternalError
//  and the GrantOwner usecase should reflect InternalError returned by the GrantOwnerGateway
func Test_Handle_GrantOwnerGateway_Returns_InternalError_When_Grant_Invoked(t *testing.T) {
	setup(t)

	// Assemble
	gateway := NewMockGrantOwnerGateway()
	gateway.GrantOwnerGatewayInternalError()
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

	// This assertion ensure the InternalError is generate when the GrantOwner.Grant(...) method is invoked
	// If the message is other than what's here, the InternalError was generated by something else than .Grant(...)
	// By testing this, we can expect that and InternalError generated  by any other implementation of the
	// GrantOwnerGateway will be properly handled
	actualMessage := usecase.Error().Error()
	expectedMessage := "Error generated from GrantOwner.Grant() invocation."
	if expectedMessage != actualMessage {
		t.Errorf("The returned error message was not expected: \n Expected: %s \n Actual %s", expectedMessage, actualMessage)
	}

}

/** Testing Tools **/

type MockGrantOwnerGateway struct {
	grantOwnerGatewayInternalError bool
	mockError                      error
}

func NewMockGrantOwnerGateway() MockGrantOwnerGateway {
	return MockGrantOwnerGateway{}
}

func (gateway *MockGrantOwnerGateway) Grant(accessControl access_control.AccessControl) {
	if gateway.grantOwnerGatewayInternalError {
		gateway.mockError =  internal_error.New("Error generated from GrantOwner.Grant() invocation.")
	}
}

func (gateway *MockGrantOwnerGateway) GrantOwnerGatewayInternalError() {
	gateway.grantOwnerGatewayInternalError = true
}

func (gateway MockGrantOwnerGateway) HasError() bool {
	return gateway.mockError != nil
}

func (gateway MockGrantOwnerGateway) Error() error {
	return gateway.mockError
}