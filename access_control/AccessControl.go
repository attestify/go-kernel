package access_control

import (
	"github.com/attestify/go-kernel/access_control/permission_list"
	"github.com/attestify/go-kernel/identity/id"
)

type AccessControl struct {
	userId     id.Id
	resourceId id.Id
	entityType string
	permission_list.PermissionList
}

func New(userId int64, resourceId int64, entityType string, permissions []string) AccessControl {
	ac := AccessControl{
		userId:     id.New(userId),
		resourceId: id.New(resourceId),
		entityType: entityType,
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

func (ac AccessControl) EntityType() string {
	return ac.entityType
}

func (ac AccessControl) HasError() bool {
	return ac.PermissionList.HasError()
}

func (ac AccessControl) Error() error {
	return ac.PermissionList.Error()
}
