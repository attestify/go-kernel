package permission_list_test

import (
	"errors"
	"github.com/attestify/go-kernel/access_control/permission_list"
	"github.com/attestify/go-kernel/error/validation_error"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path **/

// Given a permission_list class is instantiated
// When "write" is added to the list
// Then .Contains("write") should return true
func Test_Instantiate_And_Add_Permission_Successful(t *testing.T) {
	setup(t)
	// Assemble
	permissionList := permission_list.New()

	// Act
	permissionList.AddPermission("write")

	// Assert
	if permissionList.HasError() {
		t.Error("The permission list has an error when an error was not expected")
	}

	expectedValue := "write"
	if permissionList.Contains(expectedValue) != true {
		t.Errorf("Expected [%s] but it was not found in the list.", expectedValue)
	}

}

// Given a permission_list class is instantiated
// When "write" is added to the list
// Then .Contains("write") should return true
func Test_Add_Slice_Of_Permissions_Successful(t *testing.T) {
	setup(t)
	// Assemble
	permissionList := permission_list.New()

	// Act
	manyPermissions := []string{"write", "read"}
	permissionList.AddManyPermissions(manyPermissions)

	// Assert
	if permissionList.HasError() {
		t.Error("The permission list has an error when an error was not expected")
	}

	actualValues := permissionList.GetAllPermissions()
	expectedValues := []string{"write", "read"}
	if stringSlicesEqual(actualValues, expectedValues) != true {
		t.Errorf("Expected %s but it was not found in the list.\n Actual Values: %s\n", expectedValues, actualValues)
	}

}

// Given a permission_list class is instantiated
// When "write" and "read" are added to the list
// Then the .GetAllPermissions should return a slice of ["write". "read"]
func Test_GetAllPermissions_Successful(t *testing.T) {
	setup(t)
	// Assemble
	permissionList := permission_list.New()

	// Act
	permissionList.AddPermission("write")
	permissionList.AddPermission("read")

	// Assert
	if permissionList.HasError() {
		t.Error("The permission list has an error when an error was not expected")
	}

	expectedList := []string{"write", "read"}
	actualList := permissionList.GetAllPermissions()
	if stringSlicesEqual(expectedList, actualList) != true {
		t.Errorf("The actual permissions did not matched the expected permissions.\n Expected: %s\n, Actual: %s\n", expectedList, actualList)
	}

}

// Given a permission_list class is instantiated
// When "Write" is added to the list
// Then all capital letters should be lower cased
//  and the .Contains("write") should return true
func Test_Lower_Case_All_Letters(t *testing.T) {
	setup(t)
	// Assemble
	permissionList := permission_list.New()

	// Act
	permissionList.AddPermission("Write")

	// Assert
	if permissionList.HasError() {
		t.Error("The permission list has an error when an error was not expected")
	}

	expectedValue := "write"
	if permissionList.Contains(expectedValue) != true {
		t.Errorf("Expected [%s] but it was not found in the list.", expectedValue)
	}
}

// Given a permission_list class is instantiated
// When " write " is added to the list
// Then all whitespace should be removed
//  and the .Contains("write") should return true
func Test_Remove_All_Whitespace(t *testing.T) {
	setup(t)
	// Assemble
	permissionList := permission_list.New()

	// Act
	permissionList.AddPermission(" write ")

	// Assert
	if permissionList.HasError() {
		t.Error("The permission list has an error when an error was not expected")
	}

	expectedValue := "write"
	if permissionList.Contains(expectedValue) != true {
		t.Errorf("Expected [%s] but it was not found in the list.", expectedValue)
	}
}

// Given a permission_list class is instantiated
// When "write%0allow$90984find" is added to the list
// Then all numbers should be removed
//  and all special characters should be replaced with a dash
//  and the .Contains("write-allow-find") should return true
func Test_Remove_All_Numbers_Update_Special_Characters_With_Dash(t *testing.T) {
	setup(t)
	// Assemble
	permissionList := permission_list.New()

	// Act
	permissionList.AddPermission("write%0allow$90984find")

	// Assert
	if permissionList.HasError() {
		t.Error("The permission list has an error when an error was not expected")
	}

	expectedValue := "write-allow-find"
	if permissionList.Contains(expectedValue) != true {
		t.Errorf("Expected [%s] but it was not found in the list.", expectedValue)
	}
}

// Given a permission_list class is instantiated
// When "*write%0allow$90984find-" is added to the list
// Then all leading and trailing special characters should be removed
//  and the .Contains("write-allow-find") should return true
func Test_Remove_All_Leading_And_Trailing_Special_Characters(t *testing.T) {
	setup(t)
	// Assemble
	permissionList := permission_list.New()

	// Act
	permissionList.AddPermission("*write%0allow$90984find-")

	// Assert
	if permissionList.HasError() {
		t.Error("The permission list has an error when an error was not expected")
	}

	expectedValue := "write-allow-find"
	if permissionList.Contains(expectedValue) != true {
		t.Errorf("Expected [%s] but it was not found in the list.", expectedValue)
	}
}

/** Sad Path **/

// Given a permission_list class is instantiated
// When and empty string for a permission is added to the list
// Then and ValidationError should be generated
//  with the text "The permissions must be at least one alphabetical character."
func Test_Generate_Error_For_Empty_String_AddPermission(t *testing.T) {
	setup(t)
	// Assemble
	permissionList := permission_list.New()

	// Act
	permissionList.AddPermission(" ")

	// Assert
	if permissionList.HasError() != true {
		t.Error("Expected and error, although no error was provided")
	}

	if !errors.As(permissionList.Error(), &validation_error.ValidationError{}) {
		t.Errorf("did not get the epected error of type ValidationError")
	}

	actualMessage := permissionList.Error().Error()
	expectedMessage := "The permissions must be at least one alphabetical character."
	if expectedMessage != actualMessage {
		t.Errorf("The expected error message was not returned.\n Expected: %s\n Actual: %s\n", expectedMessage, actualMessage)
	}
}

// Given a permission_list class is instantiated
// When "write" is added to the list twice
// Then .Contains("write") should return true
//   and .GetAllPermissions should return a slice with only a single (1) write value
func Test_Add_Two_Of_Same_Permission_Without_Duplication(t *testing.T) {
	setup(t)
	// Assemble
	permissionList := permission_list.New()

	// Act
	permissionList.AddPermission("write")
	permissionList.AddPermission("write")

	// Assert
	if permissionList.HasError() {
		t.Error("The permission list has an error when an error was not expected")
	}

	actualValue := permissionList.GetAllPermissions()
	expectedValue := []string{"write"}
	if stringSlicesEqual(expectedValue, actualValue) != true {
		t.Errorf("The expectced value does match the actual value.\n Expected: %s\n Actual: %s\n", expectedValue, actualValue)
	}

}

// todo - Test to ensure duplicates of adding many permissions with mulitupls permissions are not added

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
