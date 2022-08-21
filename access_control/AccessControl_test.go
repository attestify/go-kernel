package access_control_test

import (
	"github.com/attestify/go-kernel/access_control"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

// Given a user Id of "1541815603606036480" is provided
//  and a resource Id of "1541815603606036481" is provided
//  and a resource of "io:attestify::entity::some-entity" is provided
//  and the permission of "write" is provided
// When the AccessControl class is instantiated
// Then .UserId() should return 1541815603606036480
//  and .ResourceId() should return 1541815603606036481
//  and .Resource should return "io:attestify::entity::some-entity"
//  and .Permissions should return a list of permissions with only one permission of "write"
func Test_Instantiate_AccessControl_Successfully(t *testing.T) {
	//Assemble
	setup(t)
	var userId int64 = 1541815603606036480
	var resourceId int64 = 1541815603606036481
	resource := "io:attestify::entity::some-entity"
	permissions := []string{"write"}

	// Act
	ac, err := access_control.New(userId, resourceId, resource, permissions)

	// Assert
	if err != nil {
		t.Error("An error was returned when no error was expected")
	}

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

	actualEntity := ac.EntityType()
	expectedEntity := "io:attestify::entity::some-entity"
	if actualEntity != expectedEntity {
		t.Errorf("The actual entity did not match the expected entity.\n Expected: %s\n Actual: %s\n", expectedEntity, actualEntity)
	}

	actualPermissions := ac.Permissions()
	expectedPermissions := []string{"write"}

	if !stringSlicesEqual(expectedPermissions, actualPermissions) {
		t.Errorf("The actual permissions did not matched the expected permissions.\n Expected: %s\n, Actual: %s\n", expectedPermissions, actualPermissions)
	}
}

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
