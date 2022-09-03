package permission_list_test

import (
	"errors"
	"github.com/attestify/go-kernel/authorization/permission_list"
	"github.com/attestify/go-kernel/error/validation_error"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path **/

// Given a permission_list class is instantiated
// When "write" is added to the list
// Then .ContainsPermission("write") should return true
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
	if permissionList.ContainsPermission(expectedValue) != true {
		t.Errorf("Expected [%s] but it was not found in the list.", expectedValue)
	}

}

// Given a permission_list class is instantiated
// When "write" is added to the list
// Then .ContainsPermission("write") should return true
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
//  and the .ContainsPermission("write") should return true
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
	if permissionList.ContainsPermission(expectedValue) != true {
		t.Errorf("Expected [%s] but it was not found in the list.", expectedValue)
	}
}

// Given a permission_list class is instantiated
// When " write " is added to the list
// Then all whitespace should be removed
//  and the .ContainsPermission("write") should return true
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
	if permissionList.ContainsPermission(expectedValue) != true {
		t.Errorf("Expected [%s] but it was not found in the list.", expectedValue)
	}
}

// Given a permission_list class is instantiated
// When "write%0allow$90984find" is added to the list
// Then all numbers should be removed
//  and all special characters should be replaced with a dash
//  and the .ContainsPermission("write-allow-find") should return true
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
	if permissionList.ContainsPermission(expectedValue) != true {
		t.Errorf("Expected [%s] but it was not found in the list.", expectedValue)
	}
}

// Given a permission_list class is instantiated
// When "*write%0allow$90984find-" is added to the list
// Then all leading and trailing special characters should be removed
//  and the .ContainsPermission("write-allow-find") should return true
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
	if permissionList.ContainsPermission(expectedValue) != true {
		t.Errorf("Expected [%s] but it was not found in the list.", expectedValue)
	}
}

// Given a permission_list class contains permissions "write" and "read"
//  and a second permission_list class contains "write" and "read"
// When the first permission class is compared to the second permission class
// Then they will equal each other
func Test_Two_PermissionLists_Equal(t *testing.T) {
	listOne := permission_list.New()
	listOne.AddPermission("write")
	listOne.AddPermission("read")

	listTwo := permission_list.New()
	listTwo.AddPermission("read")
	listTwo.AddPermission("write")

	if listOne.Equals(listTwo) != true {
		t.Errorf("Expected both permissions list to equal.\n ListOne: %s\n, ListTwo: %s\n",
			listOne.GetAllPermissions(), listTwo.GetAllPermissions())
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
// When "write" is added, and an empty string is added, and "read" is added
// Then a .HasError() should be true
//  and .Error() returns a ValidationError
//  and there should only be "write" when you get all permissions
func Test_Generate_Error_Stop_Adding_Once_Error_Occurred_AddPermission(t *testing.T) {
	setup(t)
	// Assemble
	permissionList := permission_list.New()

	// Act
	permissionList.AddPermission("write")
	permissionList.AddPermission("")
	permissionList.AddPermission("read")

	// Assert
	if permissionList.HasError() != true {
		t.Error("Expected and error, although no error was provided")
	}

	if !errors.As(permissionList.Error(), &validation_error.ValidationError{}) {
		t.Errorf("did not get the epected error of type ValidationError")
	}

	actual := permissionList.GetAllPermissions()
	expected := []string{"write"}
	if stringSlicesEqual(expected, actual) != true {
		t.Errorf("The expected values were returned.\n Expected: %s\n Actual: %s\n", expected, actual)
	}
}

// Given a permission_list class is instantiated
// When "write" is added, and an empty string is added, and "read" is added using Add Many feature
// Then a .HasError() should be true
//  and .Error() returns a ValidationError
//  and there should only be "write" when you get all permissions
func Test_Generate_Error_Stop_Adding_Once_Error_Occurred_AddManyPermission(t *testing.T) {
	setup(t)
	// Assemble
	permissionList := permission_list.New()

	// Act
	permissionList.AddManyPermissions([]string{"write", "", "read", "copy", "paste"})

	// Assert
	if permissionList.HasError() != true {
		t.Error("Expected and error, although no error was provided")
	}

	if !errors.As(permissionList.Error(), &validation_error.ValidationError{}) {
		t.Errorf("did not get the epected error of type ValidationError")
	}

	actual := permissionList.GetAllPermissions()
	expected := []string{"write"}
	if stringSlicesEqual(expected, actual) != true {
		t.Errorf("The expected values were returned.\n Expected: %s\n Actual: %s\n", expected, actual)
	}
}

// Given a permission_list class is instantiated
// When "write" is added to the list twice
// Then .GetAllPermissions() should return no duplicate copies of the permissions
//   and the only permissions that should be returned is "write"
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

// Given a permission_list class is instantiated
// When "write", "read", "delete", and "update" are added one-by-one
//  and "create", "delete", "copy", "write" are added as many permissions
// Then .GetAllPermissions() should return no duplicate copies of the permissions
//   and the only permissions that should be returned are "write", "read", "delete", "update", "create", "copy".
func Test_Add_Many_Of_Same_Permission_Without_Duplication(t *testing.T) {
	setup(t)
	// Assemble
	permissionList := permission_list.New()

	// Act
	permissionList.AddPermission("write")
	permissionList.AddPermission("read")
	permissionList.AddPermission("delete")
	permissionList.AddPermission("update")
	permissionList.AddManyPermissions([]string{"create", "delete", "copy", "write"})

	// Assert
	if permissionList.HasError() {
		t.Error("The permission list has an error when an error was not expected")
	}

	actualValue := permissionList.GetAllPermissions()
	expectedValue := []string{"write", "read", "delete", "update", "create", "copy"}
	if stringSlicesEqual(expectedValue, actualValue) != true {
		t.Errorf("The expectced value does match the actual value.\n Expected: %s\n Actual: %s\n", expectedValue, actualValue)
	}

}

// Given a permission_list class contains permissions "write", "read", and "delete"
//  and a second permission_list class contains "write" and "read"
// When the first permission class is compared to the second permission class
// Then they will not equal each other
func Test_Two_PermissionLists_Different_Size_Do_Not_Equal(t *testing.T) {
	listOne := permission_list.New()
	listOne.AddPermission("write")
	listOne.AddPermission("read")
	listOne.AddPermission("delete")

	listTwo := permission_list.New()
	listTwo.AddPermission("write")
	listTwo.AddPermission("read")

	if listOne.Equals(listTwo) {
		t.Errorf("Expected both permissions list to NOT equal.\n ListOne: %s\n, ListTwo: %s\n",
			listOne.GetAllPermissions(), listTwo.GetAllPermissions())
	}
}

// Given a permission_list class contains permissions "write" and "read"
//  and a second permission_list class contains "read" and "delete"
// When the first permission class is compared to the second permission class
// Then they will not equal each other
func Test_Two_PermissionLists_Same_Size_Do_Not_Equal(t *testing.T) {
	listOne := permission_list.New()
	listOne.AddPermission("write")
	listOne.AddPermission("read")

	listTwo := permission_list.New()
	listTwo.AddPermission("read")
	listTwo.AddPermission("delete")

	if listOne.Equals(listTwo) {
		t.Errorf("Expected both permissions list to NOT equal.\n ListOne: %s\n, ListTwo: %s\n",
			listOne.GetAllPermissions(), listTwo.GetAllPermissions())
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
