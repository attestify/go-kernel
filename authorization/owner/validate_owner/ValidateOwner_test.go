package validate_owner_test

import (
	"github.com/attestify/go-kernel/authorization/owner/validate_owner"
	"github.com/attestify/go-kernel/authorization/owner_control"
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
	usecase := validate_owner.New(gateway)

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
	usecase := validate_owner.New(gateway)
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
	usecase := validate_owner.New(gateway)
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

// todo - handle InternalError as response to ValidateOwnerGateway.Validate()

// todo - Error gates - execute .Validate() with nil gateway

/** Testing Tools **/

type MockValidateOwnerGateway struct {
	isOwner bool
	result owner_control.OwnerControl
}

func NewMockValidateOwnerGateway() MockValidateOwnerGateway{
	return MockValidateOwnerGateway{}
}

func (mock MockValidateOwnerGateway) Validate(userId id.Id, resourceId id.Id) owner_control.OwnerControl {
	var control owner_control.OwnerControl
	if mock.isOwner {
		control = owner_control.MarkAsOwner(userId.AsInteger(), resourceId.AsInteger())
	}
	return control
}

func (mock MockValidateOwnerGateway) Error() error {
	return nil
}

func (mock MockValidateOwnerGateway) HasError() bool {
	return false
}

func (mock *MockValidateOwnerGateway) ReturnTrueOwnerControl() {
	mock.isOwner = true
}

func (mock *MockValidateOwnerGateway) ReturnFalseOwnerControl() {
	mock.isOwner = false
}
