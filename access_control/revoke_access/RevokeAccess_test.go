package revoke_access_test

// todo - complete RevokeAccess Testing

import (
	"github.com/attestify/go-kernel/access_control"
	"github.com/attestify/go-kernel/access_control/revoke_access"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path **/

// Given a non-nil RevokeAccessGateway is provided
// When a RevokeAccess use case is instantiated
// Then no error should be returned
func Test_Instantiate_RevokeAccess_Successfully(t *testing.T) {
	setup(t)

	// Assemble
	var gateway revoke_access.RevokeAccessGateway = NewMockRevokeAccessGateway()

	// Act
	usecase := revoke_access.New(gateway)

	// Assert
	if usecase.Error() != nil {
		t.Errorf("An error was genrated when no eror was exptected.\n Error: %s", usecase.Error())
	}

}

// Given a non-nil RevokeAccessGateway is provided
//  and a user id of 0
//  and a resourceId of 1
//  and a resource of "test-resource"
//  and a permission of "write"
// When a RevokeAccess.Revoke(...) is invoked
// Then there should be no error
func Test_RevokeAccess_Successfully(t *testing.T) {
	setup(t)

	// Assemble
	var gateway revoke_access.RevokeAccessGateway = NewMockRevokeAccessGateway()
	var userId int64 = 0
	var resourceId int64 = 1
	resource := "test-resource"
	permissions := []string{"read"}

	// Act
	usecase := revoke_access.New(gateway)
	usecase.Revoke(userId, resourceId, resource, permissions)

	// Assert
	if usecase.Error() != nil {
		t.Errorf("An error was genrated when no eror was exptected.\n Error: %s", usecase.Error())
	}

}

/** Sad Path **/




/** Test tooling **/

type MockRevokeAccessGateway struct {}

func NewMockRevokeAccessGateway() MockRevokeAccessGateway {
	return MockRevokeAccessGateway{}
}

func (gateway MockRevokeAccessGateway) Revoke(accessControl access_control.AccessControl) error {
	return nil
}

