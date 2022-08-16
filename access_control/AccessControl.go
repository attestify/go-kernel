package access_control

import "github.com/attestify/go-kernel/identity/id"

type AccessControl struct {
	userId     id.Id
	resourceId id.Id
	entityType string
}

func New(userId int64, resourceId int64, entityType string) (AccessControl, error) {
	return AccessControl{
		userId:     id.New(userId),
		resourceId: id.New(resourceId),
		entityType: entityType,
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
