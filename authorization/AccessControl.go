package authorization

import (
	"github.com/attestify/go-kernel/authorization/permission_list"
	"github.com/attestify/go-kernel/identity/id"
)

type AccessControl struct {
	userId     id.Id
	resourceId id.Id
	permission_list.PermissionList
}

func New(userId int64, resourceId int64, permissions []string) AccessControl {
	ac := AccessControl{
		userId:     id.New(userId),
		resourceId: id.New(resourceId),
	}
	ac.AddManyPermissions(permissions)
	return ac
}

func (ac AccessControl) UserId() int64 {
	return ac.userId.AsInteger()
}

func (ac AccessControl) ResourceId() int64 {
	return ac.resourceId.AsInteger()
}

func (ac AccessControl) HasError() bool {
	return ac.PermissionList.HasError()
}

func (ac AccessControl) Error() error {
	return ac.PermissionList.Error()
}

func (ac AccessControl) Equals(compare AccessControl) bool {
	sameUserId := ac.userId == compare.userId
	sameResourceId := ac.resourceId == compare.resourceId
	samePermissionList := ac.PermissionList.Equals(compare.PermissionList)
	return sameUserId && sameResourceId && samePermissionList
}
