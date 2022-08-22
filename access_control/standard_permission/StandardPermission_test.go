package standard_permission_test

import (
	"github.com/attestify/go-kernel/access_control/standard_permission"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

// Validated that standard_permission.Create returns the value of "create"
func Test_Validate_PermissionList_Create(t *testing.T) {
	setup(t)
	actualValue := standard_permission.Create
	expectedValue := "create"

	if expectedValue != actualValue {
		t.Errorf("The actual permission list value is not the expected value.\n Expected: %s\n, Actual: %s\n", expectedValue, actualValue)
	}
}

// Validated that standard_permission.Read returns the value of "read"
func Test_Validate_PermissionList_Read(t *testing.T) {
	setup(t)
	actualValue := standard_permission.Read
	expectedValue := "read"

	if expectedValue != actualValue {
		t.Errorf("The actual permission list value is not the expected value.\n Expected: %s\n, Actual: %s\n", expectedValue, actualValue)
	}
}

// Validated that standard_permission.Update returns the value of "update"
func Test_Validate_PermissionList_Update(t *testing.T) {
	setup(t)
	actualValue := standard_permission.Update
	expectedValue := "update"

	if expectedValue != actualValue {
		t.Errorf("The actual permission list value is not the expected value.\n Expected: %s\n, Actual: %s\n", expectedValue, actualValue)
	}
}

// Validated that standard_permission.Delete returns the value of "delete"
func Test_Validate_PermissionList_Delete(t *testing.T) {
	setup(t)
	actualValue := standard_permission.Delete
	expectedValue := "delete"

	if expectedValue != actualValue {
		t.Errorf("The actual permission list value is not the expected value.\n Expected: %s\n, Actual: %s\n", expectedValue, actualValue)
	}
}

// Validated that standard_permission.CreateChild returns the value of "create-child"
func Test_Validate_PermissionList_CreateChild(t *testing.T) {
	setup(t)
	actualValue := standard_permission.CreateChild
	expectedValue := "create-child"

	if expectedValue != actualValue {
		t.Errorf("The actual permission list value is not the expected value.\n Expected: %s\n Actual: %s\n", expectedValue, actualValue)
	}
}

// Validated that standard_permission.ReadChild returns the value of "read-child"
func Test_Validate_PermissionList_ReadChild(t *testing.T) {
	setup(t)
	actualValue := standard_permission.ReadChild
	expectedValue := "read-child"

	if expectedValue != actualValue {
		t.Errorf("The actual permission list value is not the expected value.\n Expected: %s\n Actual: %s\n", expectedValue, actualValue)
	}
}

// Validated that standard_permission.UpdateChild returns the value of "update-child"
func Test_Validate_PermissionList_UpdateChild(t *testing.T) {
	setup(t)
	actualValue := standard_permission.UpdateChild
	expectedValue := "update-child"

	if expectedValue != actualValue {
		t.Errorf("The actual permission list value is not the expected value.\n Expected: %s\n Actual: %s\n", expectedValue, actualValue)
	}
}

// Validated that standard_permission.DeleteChild returns the value of "delete-child"
func Test_Validate_PermissionList_DeleteChild(t *testing.T) {
	setup(t)
	actualValue := standard_permission.DeleteChild
	expectedValue := "delete-child"

	if expectedValue != actualValue {
		t.Errorf("The actual permission list value is not the expected value.\n Expected: %s\n Actual: %s\n", expectedValue, actualValue)
	}
}
