package access_control_test

import (
	"github.com/attestify/go-kernel/access_control"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

func Test_Instantiate_AccessControl_Successfully(t *testing.T) {
	//Assemble
	var userId int64 = 1541815603606036480
	var resourceId int64 = 1541815603606036481
	var resource = "io:attestify::entity::some-entity"

	// Act
	ac, err := access_control.New(userId, resourceId, resource)

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
}
