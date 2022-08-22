package revoke_access

import (
	"github.com/attestify/go-kernel/access_control"
	"github.com/attestify/go-kernel/error/internal_error"
)

type RevokeAccess struct {
	gateway       RevokeAccessGateway
	accessControl access_control.AccessControl
	usecaseError  error
}

func New(revokeAccessGateway RevokeAccessGateway) RevokeAccess {
	var err error
	if revokeAccessGateway == nil {
		err = internal_error.New("the provided RevokeAccessGateway is nil, please provide a valid instance of an RevokeAccessGateway")
	}
	return RevokeAccess{
		gateway:      revokeAccessGateway,
		usecaseError: err,
	}
}

func (usecase *RevokeAccess) Revoke(userId int64, resourceId int64, resource string, permissions []string) {
	usecase.setAccessControl(userId, resourceId, resource, permissions)
	usecase.revokeAccessControl()
}

func (usecase *RevokeAccess) setAccessControl(userId int64, resourceId int64, resource string, permissions []string) {
	if usecase.HasError() {
		return
	}
	usecase.accessControl = access_control.New(userId, resourceId, resource, permissions)
}

func (usecase *RevokeAccess) revokeAccessControl() {
	if usecase.HasError() {
		return
	}
	usecase.gateway.Revoke(usecase.accessControl)
	if usecase.gateway.HasError() {
		usecase.usecaseError = usecase.gateway.Error()
	}
}

// Error returns the current error.  This can be nil.
func (usecase RevokeAccess) Error() error {
	return usecase.usecaseError
}

// HasError informs you if there is currently an error state
func (usecase RevokeAccess) HasError() bool {
	return usecase.usecaseError != nil
}
