package revoke_access

import (
	"github.com/attestify/go-kernel/access_control"
	"github.com/attestify/go-kernel/error/internal_error"
)

type RevokeAccess struct {
	gateway           RevokeAccessGateway
	accessControl 	access_control.AccessControl
	revokeAccessError error
}

func New(revokeAccessGateway RevokeAccessGateway) RevokeAccess {
	if revokeAccessGateway == nil {
		err := internal_error.New("the provided RevokeAccessGateway is nil, please provide a valid instance of an RevokeAccessGateway")
		return RevokeAccess{
			revokeAccessError: err,
		}
	}
	return RevokeAccess{
		gateway: revokeAccessGateway,
	}
}

func (usecase *RevokeAccess) Revoke(userId int64, resourceId int64, resource string, permissions []string)  {
	usecase.validateAccessControl(userId, resourceId, resource, permissions)
	usecase.revokeAccessControl()
}

func (usecase RevokeAccess) Error() error {
	return usecase.revokeAccessError
}

func (usecase RevokeAccess) HasError() bool {
	return usecase.revokeAccessError != nil
}

func (usecase *RevokeAccess) validateAccessControl(userId int64, resourceId int64, resource string, permissions []string) {
	if usecase.HasError() { return }
	_accessControl, _err := access_control.New(userId, resourceId, resource, permissions)
	usecase.accessControl = _accessControl
	usecase.revokeAccessError = _err
}

func (usecase *RevokeAccess) revokeAccessControl() {
	if usecase.HasError() { return }
	usecase.gateway.Revoke(usecase.accessControl)
	if usecase.gateway.HasError() {
		usecase.revokeAccessError = usecase.gateway.Error()
	}
}