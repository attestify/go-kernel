package grant_all_test

import (
	"errors"
	"github.com/attestify/go-kernel/access_control"
	"github.com/attestify/go-kernel/access_control/grant_all"
	"github.com/attestify/go-kernel/access_control/permission"
	"github.com/attestify/go-kernel/error/internal_error"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path **/

// Given there is a valid GrantAllGateway dependency
// When the GrantAll usecase is instantiated
// Then there should not be any errors
func Test_Instantiate_GrantAll_Successfully(t *testing.T) {
	setup(t)

	// Assemble
	var gateway = NewMockGrantAllGateway()

	// Act
	usecase := grant_all.New(&gateway)

	// Assert
	if usecase.HasError() {
		t.Errorf("Encountered an unexpected error.\n Error: %s\n", usecase.Error())
	}

	if usecase.Error() != nil {
		t.Errorf("An error was genrated when no eror was exptected.\n Error: %s", usecase.Error())
	}
}

// Given the GrantAll usecase is instantiated without error
// When the .Grant(..) is invoked with resourceId of 1, and permissions of "read"
// Then there should not be any errors
func Test_Invoke_Grant_Successfully(t *testing.T) {
	setup(t)

	// Assemble
	var gateway = NewMockGrantAllGateway()
	usecase := grant_all.New(&gateway)
	if usecase.HasError() {
		t.Errorf("An error was returned when no error was expected: \n %s", usecase.Error())
	}

	// Act
	var resourceId int64 = 1
	var permissions = []string{permission.Read}
	usecase.Grant(resourceId, permissions)

	// Assert
	if usecase.HasError() {
		t.Errorf("Encountered an unexpected error.\n Error: %s\n", usecase.Error())
	}
}

/** Sad Path **/

/** Side Effects of Nil GrantAllGateway **/

// Given there is a nil GrantAllGateway dependency
// When the GrantAll usecase is instantiated
// Then there should be an InternalError with the error message:
//   "The provided GrantAllGateway is nil. Please provide a valid instance of an GrantAllGateway."
func Test_Handle_InternalError_With_Nil_GrantAllGateway_When_Instantiated(t *testing.T) {
	setup(t)

	// Assemble
	var gateway grant_all.GrantAllGateway = nil

	// Act
	usecase := grant_all.New(gateway)

	// Assert
	if usecase.HasError() != true {
		t.Fatal("Expected an error, although none was provided.")
	}

	if !errors.As(usecase.Error(), &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}

	// This assertion ensure the InternalError is generated from the nil gateway check
	actualMessage := usecase.Error().Error()
	expectedMessage := "The provided GrantAllGateway is nil. Please provide a valid instance of an GrantAllGateway."
	if expectedMessage != actualMessage {
		t.Errorf("The returned error message was not expected: \n Expected: %s \n Actual %s", expectedMessage, actualMessage)
	}
}

// NOTE - this tests that error guards function properly
// Given the GrantAll usecase is instantiated with a nil GrantAllGateway dependency
// When the GrantAll.Grant(...) is invoked with valid parameters
// Then there should be an InternalError with the error message:
//   "The provided GrantAllGateway is nil. Please provide a valid instance of an GrantAllGateway."
func Test_Returns_InternalError_With_Nil_GrantAllGateway_When_Grant_Invoked(t *testing.T) {
	setup(t)

	// Assemble
	var gateway grant_all.GrantAllGateway = nil
	var resourceId int64 = 1
	var permissions = []string{permission.Read}

	// Act
	usecase := grant_all.New(gateway)
	usecase.Grant(resourceId, permissions)

	// Assert
	if usecase.HasError() != true {
		t.Fatal("Expected an error, although none was provided.")
	}

	if !errors.As(usecase.Error(), &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}

	// This assertion ensure the InternalError is generated from the nil gateway check
	actualMessage := usecase.Error().Error()
	expectedMessage := "The provided GrantAllGateway is nil. Please provide a valid instance of an GrantAllGateway."
	if expectedMessage != actualMessage {
		t.Errorf("The returned error message was not expected: \n Expected: %s \n Actual %s", expectedMessage, actualMessage)
	}
}

/** Side Effects of GrantAllGateway errors **/

// Given the GrantAll usecase is instantiated without error
// When .Grant(...) is invoked with valid parameters
// Then the GrantAllGateway returns an InternalError
//  and the GrantAll usecase should reflect InternalError returned by the GrantAllGateway
func Test_Handle_GrantAllGateway_Returns_InternalError_When_Grant_Invoked(t *testing.T) {
	setup(t)

	// Assemble
	gateway := NewMockGrantAllGateway()
	gateway.GrantAllGatewayInternalError()
	usecase := grant_all.New(&gateway)
	if usecase.HasError() == true {
		t.Fatalf("An error was thrown when non was expected.\n Error: %s\n", usecase.Error())
	}

	// Act
	var resourceId int64 = 1
	var permissions = []string{permission.Read}
	usecase.Grant(resourceId, permissions)

	// Assert
	if usecase.HasError() != true {
		t.Fatal("Expected an error, although none was provided.")
	}

	if !errors.As(usecase.Error(), &internal_error.InternalError{}) {
		t.Errorf("did not get the epected error of type InternalError")
	}

	// This assertion ensure the InternalError is generate when the GrantAll.Grant(...) method is invoked
	// If the message is other than what's here, the InternalError was generated by something else than .Grant(...)
	// By testing this, we can expect that and InternalError generated  by any other implementation of the
	// GrantAllGateway will be properly handled
	actualMessage := usecase.Error().Error()
	expectedMessage := "Error generated from GrantAll.Grant() invocation."
	if expectedMessage != actualMessage {
		t.Errorf("The returned error message was not expected: \n Expected: %s \n Actual %s", expectedMessage, actualMessage)
	}

}

/** Testing Tools **/

type MockGrantAllGateway struct {
	grantAllGatewayInternalError bool
	mockError                    error
}

func NewMockGrantAllGateway() MockGrantAllGateway {
	return MockGrantAllGateway{}
}

func (gateway *MockGrantAllGateway) Grant(control access_control.AccessControl) {
	if gateway.grantAllGatewayInternalError {
		gateway.mockError = internal_error.New("Error generated from GrantAll.Grant() invocation.")
	}
}

func (gateway *MockGrantAllGateway) GrantAllGatewayInternalError() {
	gateway.grantAllGatewayInternalError = true
}

func (gateway MockGrantAllGateway) HasError() bool { return gateway.mockError != nil }

func (gateway MockGrantAllGateway) Error() error { return gateway.mockError }