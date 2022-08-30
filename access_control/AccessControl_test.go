package access_control_test

import (
	"errors"
	"github.com/attestify/go-kernel/access_control"
	"github.com/attestify/go-kernel/error/validation_error"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path **/

// Given a user Id of "1541815603606036480" is provided
//  and a resource Id of "1541815603606036481" is provided
//  and a resource of "io:attestify::entity::some-entity" is provided
//  and the permission_list of "write" is provided
// When the AccessControl class is instantiated
// Then .UserId() should return 1541815603606036480
//  and .ResourceId() should return 1541815603606036481
//  and .Resource should return "io:attestify::entity::some-entity"
//  and .Permissions should return a list of permissions with only one permission_list of "write"
func Test_Instantiate_AccessControl_Successfully(t *testing.T) {
	//Assemble
	setup(t)
	var userId int64 = 1541815603606036480
	var resourceId int64 = 1541815603606036481
	permissions := []string{"write"}
	// Act
	ac := access_control.New(userId, resourceId, permissions)

	// Assert
	var actual int64 = ac.UserId()
	var expected int64 = 1541815603606036480
	if actual != expected {
		t.Errorf("The actual user id did not match the expected user id.\n Expected: %d\n Actual: %d\n", expected, actual)
	}

	actual = ac.ResourceId()
	expected = 1541815603606036481
	if actual != expected {
		t.Errorf("The actual resource id did not match the expected resource id.\n Expected: %d\n Actual: %d\n", expected, actual)
	}

	actualPermissions := ac.GetAllPermissions()
	expectedPermissions := []string{"write"}

	if !stringSlicesEqual(expectedPermissions, actualPermissions) {
		t.Errorf("The actual permissions did not matched the expected permissions.\n Expected: %s\n, Actual: %s\n", expectedPermissions, actualPermissions)
	}
}

// Given a user Id of "1541815603606036480" is provided
//  and a resource Id of "1541815603606036481" is provided
//  and the permission_list of "write", "read", "delete" is provided for the first AccessControl
//  AND given a user Id of "1541815603606036480" is provided
//  and a resource Id of "1541815603606036481" is provided
//  and the permission_list of "delete", "write", "read" is provided for the second AccessControl
// When both AccessControls are compared
// Then they must equal each other
func Test_AccessControlOne_Equals_AccessControlTwo_Successfully(t *testing.T) {
	//Assemble
	setup(t)
	var userId int64 = 1541815603606036480
	var resourceId int64 = 1541815603606036481
	permissions := []string{"write", "read", "delete"}

	var userId2 int64 = 1541815603606036480
	var resourceId2 int64 = 1541815603606036481
	permissions2 := []string{"delete", "write", "read"}

	// Act
	acOne := access_control.New(userId, resourceId, permissions)
	acTwo := access_control.New(userId2, resourceId2, permissions2)


	// Assert
	if acOne.Equals(acTwo) != true {
		t.Errorf("Expected both AccessControls equal.\n acOne: %+v\n, acTwo: %+v\n", acOne, acTwo)
	}
}

/** Sad Path **/

// Given a user Id of "1541815603606036480" is provided
//  and a resource Id of "1541815603606036481" is provided
//  and a resource of "io:attestify::entity::some-entity" is provided
//  and the permission_list of and empty string is provided
// When the AccessControl class is instantiated
// Then .HasError() should return true
//  and .Error() should return a ValidationError
func Test_Error_Propagation_For_PermissionList(t *testing.T) {
	//Assemble
	setup(t)
	var userId int64 = 1541815603606036480
	var resourceId int64 = 1541815603606036481
	permissions := []string{""}

	// Act
	ac := access_control.New(userId, resourceId, permissions)

	// Assert
	if ac.HasError() != true {
		t.Errorf("Expected an error, although no error exists")
	}

	if !errors.As(ac.Error(), &validation_error.ValidationError{}) {
		t.Errorf("Did not get the epected error of type ValidationError")
	}
}

// Given a user Id of "1541815603606036482" is provided
//  and a resource Id of "1541815603606036481" is provided
//  and the permission_list of "write", "read", "delete" is provided for the first AccessControl
//  AND given a user Id of "1541815603606036480" is provided
//  and a resource Id of "1541815603606036481" is provided
//  and the permission_list of "delete", "write", "read" is provided for the second AccessControl
// When both AccessControls are compared
// Then they must equal each other
func Test_AccessControlOne_DoesNotEqual_AccessControlTwo_Different_UserId(t *testing.T) {
	//Assemble
	setup(t)
	var userId int64 = 1541815603606036482
	var resourceId int64 = 1541815603606036481
	permissions := []string{"write", "read", "delete"}

	var userId2 int64 = 1541815603606036480
	var resourceId2 int64 = 1541815603606036481
	permissions2 := []string{"write", "read", "delete"}

	// Act
	acOne := access_control.New(userId, resourceId, permissions)
	acTwo := access_control.New(userId2, resourceId2, permissions2)

	// Assert
	if acOne.Equals(acTwo) {
		t.Errorf("Expected both AccessControls to NOT equal because of different userIds.\n acOne: %d\n, acTwo: %d\n",
			acOne.UserId(), acTwo.UserId())
	}
}

// todo - updated test to reflect different resourceId
// Given a user Id of "1541815603606036480" is provided
//  and a resource Id of "1541815603606036481" is provided
//  and the permission_list of "write", "read", "delete" is provided for the first AccessControl
//  AND given a user Id of "1541815603606036480" is provided
//  and a resource Id of "1541815603606036482" is provided
//  and the permission_list of "delete", "write", "read" is provided for the second AccessControl
// When both AccessControls are compared
// Then they must equal each other
func Test_AccessControlOne_DoesNotEqual_AccessControlTwo_Different_ResourceId(t *testing.T) {
	//Assemble
	setup(t)
	var userId int64 = 1541815603606036480
	var resourceId int64 = 1541815603606036481
	permissions := []string{"write", "read", "delete"}

	var userId2 int64 = 1541815603606036480
	var resourceId2 int64 = 1541815603606036482
	permissions2 := []string{"write", "read", "delete"}

	// Act
	acOne := access_control.New(userId, resourceId, permissions)
	acTwo := access_control.New(userId2, resourceId2, permissions2)

	// Assert
	if acOne.Equals(acTwo) {
		t.Errorf("Expected both AccessControls to NOT equal because of different resourceIds.\n acOne: %d\n, " +
			"acTwo: %d\n",
			acOne.ResourceId(), acTwo.ResourceId())
	}
}

// Given a user Id of "1541815603606036480" is provided
//  and a resource Id of "1541815603606036481" is provided
//  and the permission_list of "write" and "read", is provided for the first AccessControl
//  AND given a user Id of "1541815603606036480" is provided
//  and a resource Id of "1541815603606036481" is provided
//  and the permission_list of "delete" and "write" is provided for the second AccessControl
// When both AccessControls are compared
// Then they must equal each other
func Test_AccessControlOne_DoesNotEqual_AccessControlTwo_Different_Permissions(t *testing.T) {
	//Assemble
	setup(t)
	var userId int64 = 1541815603606036480
	var resourceId int64 = 1541815603606036481
	permissions := []string{"write", "read"}

	var userId2 int64 = 1541815603606036480
	var resourceId2 int64 = 1541815603606036481
	permissions2 := []string{"delete", "write"}

	// Act
	acOne := access_control.New(userId, resourceId, permissions)
	acTwo := access_control.New(userId2, resourceId2, permissions2)

	// Assert
	if acOne.Equals(acTwo) {
		t.Errorf("Expected both AccessControls to NOT equal because of different PermissonLists.\n acOne: %s\n, " +
			"acTwo: %s\n",
			acOne.GetAllPermissions(), acTwo.GetAllPermissions())
	}
}


/** Testing Tools **/
func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
