package id

import (
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

// Test_Instantiate_Id instantiates Id struct successfully without error
func Test_Instantiate_Id(t *testing.T) {
	setup(t)
	New(1541815603606036480)
}

// Test_Instantiate_Id_Get_Value returns the string value of "1541815603606036480"
// when Id is instantiated with the int64 argument of 1541815603606036480.
func Test_Instantiate_Id_Get_Value(t *testing.T) {
	setup(t)
	id := New(1541815603606036480)

	expectedValue := "1541815603606036480"
	actualValue := id.Value()
	if expectedValue != actualValue {
		t.Errorf("Did not return the expected value.\nActual: %s\nExpected: %s",
			actualValue,
			expectedValue)
	}

}

// Test_Instantiate_Id_Get_Value returns the integer value of 1541815603606036480
// when Id is instantiated with the int64 argument of 1541815603606036480.
func Test_Instantiate_Id_Get_As_Integer(t *testing.T) {
	setup(t)
	id := New(1541815603606036480)

	var expectedInt int64
	expectedInt = 1541815603606036480
	actualInt := id.AsInteger()
	if expectedInt != actualInt {
		t.Errorf("Did not return the expected value.\nActual: %d\nExpected: %d",
			actualInt,
			expectedInt)
	}

}
