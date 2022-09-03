package owner_control

import "github.com/attestify/go-kernel/identity/id"

type OwnerControl struct {
	userId id.Id
	resourceId id.Id
	isOwner bool
}

func MarkAsOwner(userId int64, resourceId int64) OwnerControl {
	return OwnerControl{
		userId: id.New(userId),
		resourceId: id.New(resourceId),
		isOwner: true,
	}
}

func MarkAsNotOwner(userId int64, resourceId int64) OwnerControl {
	return OwnerControl{
		userId: id.New(userId),
		resourceId: id.New(resourceId),
		isOwner: false,
	}
}

func (control OwnerControl) IsOwner() bool {
	return control.isOwner
}

func (control OwnerControl) IsNotOwner() bool {
	return !control.IsOwner()
}

func (control OwnerControl) Owner() id.Id {
	return control.userId
}

func (control OwnerControl) Resource() id.Id {
	return control.resourceId
}

func (control OwnerControl) Equals(compare OwnerControl) bool {
	isOwnerEquals := control.IsOwner() == compare.IsOwner()
	sameOwner := control.Owner() == compare.Owner()
	sameResource := control.Resource() == compare. Resource()

	return isOwnerEquals && sameOwner && sameResource
}
