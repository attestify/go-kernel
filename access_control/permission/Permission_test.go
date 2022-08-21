package permission_test

import (
	"github.com/attestify/go-kernel/access_control/permission"
	"testing"
)

func setup(t *testing.T){
	t.Parallel()
}

/** Happy Path **/

// Given a value of "write" is provided
// When a permission class is instantiated
// Then the .Value() method should return "write"
func Test_Instantiate_Permission_Successful(t *testing.T)  {
	setup(t)
	perm := permission.New("write")

	actual := perm.Value()
	expected := "write"
	if expected != actual {
		t.Errorf("Actual permission is not what was expected.\n Expected: %s\n Actual: %s\n", expected, actual)
	}

}

// Given a value of "Write" is provided
// When a permission class is instantiated
// Then the .Value() method should return "write"
//  and all letters should be lower cased
func Test_Lower_Case_All_Letters(t *testing.T) {
	setup(t)
	perm := permission.New("Write")

	actual := perm.Value()
	expected := "write"
	if expected != actual {
		t.Errorf("Actual permission is not what was expected.\n Expected: %s\n Actual: %s\n", expected, actual)
	}
}

// todo - Test_Remove_All_Whitespace - provide description
func Test_Remove_All_Whitespace(t *testing.T) {
	setup(t)
	perm := permission.New(" Write ")

	actual := perm.Value()
	expected := "write"
	if expected != actual {
		t.Errorf("Actual permission is not what was expected.\n Expected: %s\n Actual: %s\n", expected, actual)
	}
}

// todo - Test_Remove_All_Numbers_Update_Special_Characters_With_Dash - provide description
func Test_Remove_All_Numbers_Update_Special_Characters_With_Dash(t *testing.T) {
	setup(t)
	perm := permission.New("write%0allow$90984find")

	actual := perm.Value()
	expected := "write-allow-find"
	if expected != actual {
		t.Errorf("Actual permission is not what was expected.\n Expected: %s\n Actual: %s\n", expected, actual)
	}
}

// todo - Test_Remove_All_Leading_And_Trailing_Special_Characters - provide description
func Test_Remove_All_Leading_And_Trailing_Special_Characters(t *testing.T) {
	setup(t)
	perm := permission.New("*write%0allow$90984find-")

	actual := perm.Value()
	expected := "write-allow-find"
	if expected != actual {
		t.Errorf("Actual permission is not what was expected.\n Expected: %s\n Actual: %s\n", expected, actual)
	}
}

/** Sad Path **/
