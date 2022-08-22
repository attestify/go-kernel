package revoke_access

import (
	"github.com/attestify/go-kernel/access_control"
)

type RevokeAccess struct {
	gateway           RevokeAccessGateway
	revokeAccessError error
}

func New(revokeAccessGateway RevokeAccessGateway) RevokeAccess {
	return RevokeAccess{
		gateway: revokeAccessGateway,
	}
}

func (usecase RevokeAccess) Revoke(userId int64, resourceId int64, resource string, permissions []string)  {
	accessControl, _ := access_control.New(userId, resourceId, resource, permissions)
	err := usecase.gateway.Revoke(accessControl)
	if err != nil {
		usecase.revokeAccessError = err
		return
	}
}

func (usecase RevokeAccess) Error() error {
	return usecase.revokeAccessError
}