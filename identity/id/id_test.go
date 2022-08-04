package id

import (
	"testing"
)

// Instantiate Id struct successfully without error
func Test_Instantiate_Id(t *testing.T) {
	id := New(1541815603606036480)
	if id == nil {
		t.Error("Should not have received a nil object.")
	}
}

func Test_Instantiate_Id_Get_Value(t *testing.T) {
	id := New(1541815603606036480)

	expectedValue := "1541815603606036480"
	actualValue := id.Value()
	if expectedValue != actualValue {
		t.Errorf("Did not return the expected value.\nActual: %s\nExpected: %s",
			actualValue,
			expectedValue)
	}

}

func Test_Instantiate_Id_Get_As_Value(t *testing.T) {
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
