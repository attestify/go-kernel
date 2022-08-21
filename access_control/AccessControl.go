package access_control

import (
	"github.com/attestify/go-kernel/access_control/permission"
	"github.com/attestify/go-kernel/identity/id"
)

type AccessControl struct {
	userId     id.Id
	resourceId id.Id
	entityType string
	permissions []permission.Permission
}

func New(userId int64, resourceId int64, entityType string, permissions []string) (AccessControl, error) {
	var validPermissions []permission.Permission
	for _, perm := range permissions {
		_perm := permission.New(perm)
		validPermissions = append(validPermissions, _perm)
	}
	return AccessControl{
		userId:     id.New(userId),
		resourceId: id.New(resourceId),
		entityType: entityType,
		permissions: validPermissions,
	}, nil
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

func (ac AccessControl) Permissions() []string {
	var result []string
	for _, perm := range ac.permissions {
		result = append(result, perm.Value())
	}
	return result
}
