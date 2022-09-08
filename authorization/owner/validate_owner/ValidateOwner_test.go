package validate_owner_test

import (
	"errors"
	"github.com/attestify/go-kernel/authorization/owner/validate_owner"
	"github.com/attestify/go-kernel/authorization/owner_control"
	"github.com/attestify/go-kernel/error/internal_error"
	"github.com/attestify/go-kernel/identity/id"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path **/

// Given there is a valid ValidateOwnerGateway dependency
// When the ValidateOwner usecase is instantiated
// Then there should not be any errors
func Test_Instantiate_ValidateOwner_Successfully(t *testing.T) {
	setup(t)

	// Assemble
	gateway := NewMockValidateOwnerGateway()

	// Act
	usecase := validate_owner.New(&gateway)

	// Assert
	if usecase.HasError() != false {
		t.Errorf("Encountered an unexpected error.\n Error: %s\n", usecase.Error())
	}
}

// Given the ValidateOwnerGateway is expected to return an OwnerControl object set to TRUE
//  with successfully instantiated ValidateOwner usecase
// When the .Validate(...) is invoked
// Then an OwnerControl showing as owner must be returned
func Test_Instantiate_Invoke_Validate_Successfully_With_True_OwnerControl(t *testing.T) {
	setup(t)

	// Assemble
	gateway := NewMockValidateOwnerGateway()
	gateway.ReturnTrueOwnerControl()
	usecase := validate_owner.New(&gateway)
	if usecase.HasError() != false {
		t.Fatalf("Encountered an unexpected error.\n Error: %s\n", usecase.Error())
	}

	// Act
	var userId int64 = 0
	var resourceId int64 = 1
	control := usecase.Validate(userId, resourceId)

	// Assert
	if control.IsOwner() != true {
		t.Error("Expected .IsOwner() to be true.")
	}

}

// Given the ValidateOwnerGateway is expected to return an OwnerControl object set to FALSE
//  with successfully instantiated ValidateOwner usecase
// When the .Validate(...) is invoked
// Then an OwnerControl showing as NOT owner must be returned
func Test_Instantiate_Invoke_Validate_Successfully_With_False_OwnerControl(t *testing.T) {
	setup(t)

	// Assemble
	gateway := NewMockValidateOwnerGateway()
	gateway.ReturnFalseOwnerControl()
	usecase := validate_owner.New(&gateway)
	if usecase.HasError() != false {
		t.Fatalf("Encountered an unexpected error.\n Error: %s\n", usecase.Error())
	}

	// Act
	var userId int64 = 0
	var resourceId int64 = 1
	control := usecase.Validate(userId, resourceId)

	// Assert
	if control.IsOwner() != false {
		t.Error("Expected .IsOwner() to be false.")
	}

}

/** Sad Path **/

/** Side Effects of nil ValidateOwnerGateway **/

// Given there is a nil ValidateOwnerGateway dependency
// When the ValidateOwner usecase is instantiated
// Then there should be an InternalError with the error message:
//   "The provided ValidateOwnerGateway is nil. Please provide an instance of a ValidateOwnerGateway."
func Test_Handle_InternalError_With_Nil_ValidateOwnerGateway_When_Instantiated(t *testing.T) {
	setup(t)

	// Assemble
	var gateway validate_owner.ValidateOwnerGateway = nil

	// Act
	usecase := validate_owner.New(gateway)

	// Assert
	if usecase.HasError() != true {
		t.Fatalf("Expceted an error although no error was recorded.")
	}

	actualMessage := usecase.Error().Error()
	expectedMessage := "The provided ValidateOwnerGateway is nil. Please provide an instance of a ValidateOwnerGateway."
	if expectedMessage != actualMessage {
		t.Errorf("Did not get the expeceted error message.\n Expected: %s\n Actual: %s\n", expectedMessage, actualMessage)
	}
}

// NOTE - this tests that error guards function properly
// Given the ValidateOwner usecase is instantiated with a nil ValidateOwnerGateway dependency
// When the ValidateOwner.Validate(...) is invoked with valid parameters
// Then there should be an InternalError with the error message:
//   "The provided ValidateOwnerGateway is nil. Please provide an instance of a ValidateOwnerGateway."
func Test_Returns_InternalError_With_Nil_ValidateOwnerGateway_When_Validate_Invoked(t *testing.T) {
	setup(t)

	// Assemble
	var gateway validate_owner.ValidateOwnerGateway = nil
	usecase := validate_owner.New(gateway)

	// Act
	var userId int64 = 0
	var resourceId int64 = 1
	actualOwnerControl := usecase.Validate(userId, resourceId)

	// Assert
	if usecase.HasError() != true {
		t.Fatalf("Expceted an error although no error was recorded.")
	}

	actualMessage := usecase.Error().Error()
	expectedMessage := "The provided ValidateOwnerGateway is nil. Please provide an instance of a ValidateOwnerGateway."
	if expectedMessage != actualMessage {
		t.Errorf("Did not get the expeceted error message.\n Expected: %s\n Actual: %s\n", expectedMessage, actualMessage)
	}

	expectedOwnerControl := owner_control.MarkAsNotOwner(0, 1)
	if expectedOwnerControl.Equals(actualOwnerControl) == false {
		t.Errorf("The expected owner control was not regured.\n Expected: %v\n, Actual: %v\n", expectedOwnerControl,
			actualOwnerControl)
	}

}

/** Side Effects of ValidateOwnerGateway errors **/

// Given the ValidateOwner usecase is instantiated without error
// When .Validate(...) is invoked with valid parameters
// Then the ValidateOwnerGateway returns an InternalError
//  and the ValidateOwner usecase should reflect InternalError returned by the ValidateOwnerGateway
func Test_Handle_ValidateOwnerGateway_Returns_InternalError_When_Validate_Invoked(t *testing.T) {
	setup(t)

	// Assemble
	gateway := NewMockValidateOwnerGateway()
	gateway.ValidateOwnerGatewayInternalError()
	usecase := validate_owner.New(&gateway)
	if usecase.HasError() == true {
		t.Fatalf("An error was thrown when non was expected.\n Error: %s\n", usecase.Error())
	}

	// Act
	var userId int64 = 0
	var resourceId int64 = 1
	usecase.Validate(userId, resourceId)

	//Assert
	if usecase.HasError() != true {
		t.Fatal("An error was expected, but no error was returned")
	}

	if !errors.As(usecase.Error(), &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}

	// This assertion ensure the InternalError is generate when the ValidateOwner.Validate(...) method is invoked
	// If the message is other than what's here, the InternalError was generated by something else than .Validate(...)
	// By testing this, we can expect that and InternalError generated  by any other implementation of the
	// ValidateOwnerGateway will be properly handled
	actualMessage := usecase.Error().Error()
	expectedMessage := "Error generated from ValidateOwnerGateway.Validate() invocation."
	if expectedMessage != actualMessage {
		t.Errorf("The returned error message was not expected: \n Expected: %s \n Actual %s", expectedMessage, actualMessage)
	}
}

/** Testing Tools **/

type MockValidateOwnerGateway struct {
	isOwner bool
	validateOwnerGatewayInternalError bool
	mockError                      error
}

func NewMockValidateOwnerGateway() MockValidateOwnerGateway{
	return MockValidateOwnerGateway{}
}

func (mock *MockValidateOwnerGateway) Validate(userId id.Id, resourceId id.Id) owner_control.OwnerControl {
	if mock.validateOwnerGatewayInternalError {
		mock.mockError = internal_error.New("Error generated from ValidateOwnerGateway.Validate() invocation.")
		return owner_control.MarkAsNotOwner(userId.AsInteger(), resourceId.AsInteger())
	}

	var control owner_control.OwnerControl
	if mock.isOwner {
		control = owner_control.MarkAsOwner(userId.AsInteger(), resourceId.AsInteger())
	}

	return control
}

func (mock MockValidateOwnerGateway) Error() error {
	return mock.mockError
}

func (mock MockValidateOwnerGateway) HasError() bool {
	return mock.mockError != nil
}

func (mock *MockValidateOwnerGateway) ReturnTrueOwnerControl() {
	mock.isOwner = true
}

func (mock *MockValidateOwnerGateway) ReturnFalseOwnerControl() {
	mock.isOwner = false
}

func (mock *MockValidateOwnerGateway) ValidateOwnerGatewayInternalError() {
	mock.validateOwnerGatewayInternalError = true
}
