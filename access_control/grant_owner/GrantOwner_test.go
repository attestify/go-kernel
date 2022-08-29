package grant_owner_test

import (
	"errors"
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
	usecase := grant_owner.New(gateway)

	//Assert
	if usecase.HasError() {
		t.Errorf("An unexpected error occurred.\n Error: %s\n", usecase.Error())
	}

}

//func Test_Invoke_Grant_Successfully(t *testing.T) {
//	setup(t)
//	// Assemble
//	gateway := NewGrantOwnerGatewayMock()
//	usecase := grant_owner.New(gateway)
//	var userId int64 = 0
//	var resourceId int64 = 1
//
//	//Act
//	usecase.Grant(userId, resourceId)
//
//	//Assert
//	if usecase.HasError() {
//		t.Errorf("An unexpected error occurred.\n Error: %s\n", usecase.Error())
//	}
//
//}

/** Sad Path **/

// todo - Left off here - need to complete implementation to fulfill the test
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
	if !usecase.HasError() {
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



/** Testing Tools **/

type GrantOwnerGatewayMock struct {}

func NewGrantOwnerGatewayMock()  GrantOwnerGatewayMock {
	return GrantOwnerGatewayMock{}
}
