package owner_control_test

import (
	"github.com/attestify/go-kernel/authorization/owner_control"
	"testing"
)

func setup(t *testing.T) {
	t.Parallel()
}

/** Happy Path **/

func Test_Mark_User_Owner_Of_Resource_Successfully(t *testing.T) {
	setup(t)

	// Assemble
	var userId int64 = 1541815603606036480
	var resourceId int64 = 1541815603606036481

	// Act
	ownerControl := owner_control.MarkAsOwner(userId, resourceId)

	// Assert
	if ownerControl.IsOwner() != true {
		t.Error("Expected .IsOwner() to be true.  It is false.")
	}

	if ownerControl.IsNotOwner() != false {
		t.Error("Expected .IsNotOwner() to be false.  It is true.")
	}

	owner := ownerControl.Owner()
	actualOwner := owner.AsInteger()
	var expectedOwner int64 = 1541815603606036480
	if expectedOwner != actualOwner {
		t.Errorf("Owner IDs mismatch.\n Expected: %d\n, Actual: %d\n", expectedOwner, actualOwner)
	}

	resource := ownerControl.Resource()
	actualResource := resource.AsInteger()
	var expectedResource int64 = 1541815603606036481
	if actualResource != expectedResource {
		t.Errorf("Resource IDs mismatch.\n Expected: %d\n, Actual: %d\n", expectedResource, actualResource)
	}
}

func Test_Mark_User_Not_Owner_Of_Resource_Successfully(t *testing.T) {
	setup(t)

	// Assemble
	var userId int64 = 1541815603606036480
	var resourceId int64 = 1541815603606036481

	// Act
	ownerControl := owner_control.MarkAsNotOwner(userId, resourceId)

	// Assert
	if ownerControl.IsOwner() != false {
		t.Error("Expected .IsOwner() to be false.  It is true.")
	}

	if ownerControl.IsNotOwner() != true {
		t.Error("Expected .IsNotOwner() to be true.  It is false.")
	}

	owner := ownerControl.Owner()
	actualOwner := owner.AsInteger()
	var expectedOwner int64 = 1541815603606036480
	if expectedOwner != actualOwner {
		t.Errorf("Owner IDs mismatch.\n Expected: %d\n, Actual: %d\n", expectedOwner, actualOwner)
	}

	resource := ownerControl.Resource()
	actualResource := resource.AsInteger()
	var expectedResource int64 = 1541815603606036481
	if actualResource != expectedResource {
		t.Errorf("Resource IDs mismatch.\n Expected: %d\n, Actual: %d\n", expectedResource, actualResource)
	}
}

func Test_Two_Same_OwnerControl_Objects_Equal(t *testing.T) {
	setup(t)

	// Assemble
	var userId int64 = 1541815603606036480
	var resourceId int64 = 1541815603606036481
	control1 := owner_control.MarkAsOwner(userId, resourceId)

	var userId2 int64 = 1541815603606036480
	var resourceId2 int64 = 1541815603606036481
	control2 := owner_control.MarkAsOwner(userId2, resourceId2)

	// Act
	controlsEqual := control1.Equals(control2)

	// Assert
	if controlsEqual != true {
		t.Errorf("Expected the controls to equal.\n Control1: %v\n, Control2: %v\n", control1, control2)
	}


}

/** Sad Path **/

